package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	v1 "AvitoProject/pkg/modeles/v1"
)

type service interface{
	GetOrder(orderID int) (order *v1.Order, err error)
	GetAllOrders(sellerID int) (orders []v1.Order, err error)
	SetOrder(buyerID, endAddrID, noticeID int) (err error)
	CalculatePrice(addrID, noticeID int) (price int, err error)
}

type server struct{
	svc service
}

func (s *server) HandleCalculate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "GET")
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

func (s *server) HandleOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		orderID, err := strconv.Atoi(vars["order_id"])
		if err != nil{
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		order, err := s.svc.GetOrder(orderID)
		if err != nil{
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, order)
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
		buyerID, err := strconv.Atoi(r.FormValue("buyer_id"))
		if err != nil{
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		endAddrID, err := strconv.Atoi(r.FormValue("end_addr_id"))
		if err != nil{
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		noticeID, err := strconv.Atoi(r.FormValue("notice_id"))
		if err != nil{
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		s.svc.SetOrder(buyerID, endAddrID, noticeID)
	case "GET":
		seller, err := strconv.Atoi(r.FormValue("seller"))
		if err != nil{
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}
		orders, err := s.svc.GetAllOrders(seller)
		if err != nil{
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, orders)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
}

// NewServer returns a new mux.Router instance.
func NewServer(svc service) (httpServer *mux.Router){
	s := server{svc: svc}
	router := mux.NewRouter()
	//mux := http.NewServeMux()
	router.HandleFunc("/", s.Handle)

	router.HandleFunc("/price", s.HandleCalculate)

	router.HandleFunc("/orders/{order_id:[0-9]+}", s.HandleOrder)

	router.HandleFunc("/orders", s.HandleOrders)

	return router
}