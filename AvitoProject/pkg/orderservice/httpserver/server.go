package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	v1 "AvitoProject/pkg/modeles/v1"
)

type service interface {
	GetOrder(orderID int) (order *v1.Order, err error)
	GetAllOrders(sellerID int) (orders []v1.Order, err error)
	SetOrder(oc *v1.OrderCreation) (orderID int, err error)
	CalculatePrice(addrID, noticeID int) (price int, err error)
}

type serializator interface {
	DecodeSetOrder(r *http.Request) (oc *v1.OrderCreation, err error)
	EncodeSetOrder(w http.ResponseWriter, orderID int)
	EncodeGetDeliveryPrice(w http.ResponseWriter, price int)
	EncodeGetOrder(w http.ResponseWriter, order *v1.Order)
	EncodeGetOrders(w http.ResponseWriter, orders []v1.Order)
}

type server struct {
	svc service
	ser serializator
}

func (s *server) HandleCalculate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		endAddrID, err := strconv.Atoi(r.FormValue("end_addr_id"))
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		noticeID, err := strconv.Atoi(r.FormValue("notice_id"))
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		price, err := s.svc.CalculatePrice(endAddrID, noticeID)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		s.ser.EncodeGetDeliveryPrice(w, price)

	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func (s *server) HandleOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		orderID, err := strconv.Atoi(vars["order_id"])
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		order, err := s.svc.GetOrder(orderID)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		s.ser.EncodeGetOrder(w, order)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func (s *server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
}

func (s *server) HandleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		oc, err := s.ser.DecodeSetOrder(r)
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		orderID, err := s.svc.SetOrder(oc)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		s.ser.EncodeSetOrder(w, orderID)

	case "GET":
		seller, err := strconv.Atoi(r.FormValue("seller"))
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		orders, err := s.svc.GetAllOrders(seller)
		if err != nil {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		s.ser.EncodeGetOrders(w, orders)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

// NewServer returns a new mux.Router instance.
func NewServer(svc service, ser serializator) (httpServer *mux.Router) {
	s := server{svc: svc, ser: ser}
	router := mux.NewRouter()

	router.HandleFunc("/", s.Handle)

	router.HandleFunc("/orders/price", s.HandleCalculate)

	router.HandleFunc("/orders/{order_id:[0-9]+}", s.HandleOrder)

	router.HandleFunc("/orders", s.HandleOrders)

	return router
}
