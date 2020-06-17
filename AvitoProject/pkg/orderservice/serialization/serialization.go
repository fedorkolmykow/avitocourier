package serialization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	v1 "AvitoProject/pkg/modeles/v1"
)

// Serializator ...
type Serializator interface {
	DecodeSetOrder(r *http.Request) (oc *v1.OrderCreationRequest, err error)
	EncodeSetOrder(w http.ResponseWriter, orderID int)
	EncodeGetDeliveryPrice(w http.ResponseWriter, price int)
	EncodeGetOrder(w http.ResponseWriter, order *v1.Order)
	EncodeGetOrders(w http.ResponseWriter, orders []v1.Order)
}

type serializator struct{}



func (c *serializator) DecodeSetOrder(r *http.Request) (oc *v1.OrderCreationRequest, err error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err = json.Unmarshal(body, &oc)
	return
}

func (c *serializator) EncodeSetOrder(w http.ResponseWriter, orderID int) {
	r := v1.OrderCreationResponse{OrderID: orderID}
	encode(w, r)
	return
}

func (c *serializator) EncodeGetDeliveryPrice(w http.ResponseWriter, price int) {
	r := v1.OrderDeliveryPriceResponse{DeliveryPrice: price}
	encode(w, r)
	return
}

func (c *serializator) EncodeGetOrder(w http.ResponseWriter, order *v1.Order) {
	encode(w, order)
	return
}

func (c *serializator) EncodeGetOrders(w http.ResponseWriter, orders []v1.Order) {
	r := v1.ShortOrdersResponse{Orders: orders}
	encode(w, r)
	return
}

func encode(w http.ResponseWriter, r interface{}) {
	res, err := json.Marshal(&r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

// NewSerializator returns a new Serializator instance.
func NewSerializator() Serializator {
	return &serializator{}
}
