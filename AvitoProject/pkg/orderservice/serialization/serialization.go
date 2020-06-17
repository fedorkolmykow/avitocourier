package serialization

import (
	v1 "AvitoProject/pkg/modeles/v1"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Serializator ...
type Serializator interface{
	DecodeSetOrder(r *http.Request) (oc *v1.OrderCreation, err error)
	EncodeSetOrder(w http.ResponseWriter, orderID int)
    EncodeGetDeliveryPrice(w http.ResponseWriter, price int)
	EncodeGetOrder(w http.ResponseWriter, order *v1.Order)
	EncodeGetOrders(w http.ResponseWriter, orders []v1.Order)
}

type serializator struct {}

type setOrderResponse struct{
	OrderID int `json:"order_id"`
}

type getDeliveryPriceResponse struct {
	DeliveryPrice int `json:"delivery_price"`
}

type getOrders struct{
	Orders []v1.Order `json:"orders"`
}

func (c* serializator) DecodeSetOrder(r *http.Request) (oc *v1.OrderCreation, err error){
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err = json.Unmarshal(body, oc)
	return
}

func (c *serializator) EncodeSetOrder(w http.ResponseWriter, orderID int){
	r := setOrderResponse{OrderID: orderID}
	Encode(w, r)
	return
}

func (c *serializator) EncodeGetDeliveryPrice(w http.ResponseWriter, price int){
	r := getDeliveryPriceResponse{DeliveryPrice: price}
	Encode(w, r)
	return
}

func (c *serializator) EncodeGetOrder(w http.ResponseWriter, order *v1.Order){
	Encode(w, order)
	return
}

func (c *serializator) EncodeGetOrders(w http.ResponseWriter, orders []v1.Order){
	r := getOrders{Orders:orders}
    Encode(w, r)
	return
}

func Encode(w http.ResponseWriter, r interface{}){
	res, err := json.Marshal(&r)
	if err != nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

// NewSerializator returns a new Serializator instance.
func NewSerializator() Serializator{
	return &serializator{}
}
