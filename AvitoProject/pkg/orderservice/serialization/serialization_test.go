package serialization

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "AvitoProject/pkg/modeles/v1"
)

const(

)



func TestDecodeSetOrder(t *testing.T) {
	c := NewSerializator()
	expOc := &v1.OrderCreationRequest{
		BuyerID:   1,
		EndAddrID: 2,
		NoticeID:  3,
	}
	body, err := json.Marshal(expOc)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	oc, err := c.DecodeSetOrder(req)
	if err != nil {
		t.Fatal(err)
	}
	if *oc != *expOc {
		t.Errorf("got %v want %v",
			oc, expOc)
	}
}

func TestEncodeSetOrder(t *testing.T) {
	c := NewSerializator()
	w := httptest.NewRecorder()
	orderID := 1
	exp := `{"order_id":1}`
	c.EncodeSetOrder(w, orderID)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := string(body)
	if bodyStr != exp {
		t.Errorf("got %v want %v",
			bodyStr, exp)
	}
}

func TestEncodeGetDeliveryPrice(t *testing.T) {
	c := NewSerializator()
	w := httptest.NewRecorder()
	price:= 100
	exp := `{"delivery_price":100}`
	c.EncodeGetDeliveryPrice(w, price)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := string(body)
	if bodyStr != exp {
		t.Errorf("got %v want %v",
			bodyStr, exp)
	}
}

func TestEncodeGetOrder(t *testing.T) {
	c := NewSerializator()
	w := httptest.NewRecorder()
	o := &v1.Order{
		OrderID:       0,
		SellerPhone:   "911",
		SellerName:    "Police",
		BuyerPhone:    "912",
		BuyerName:     "NotPolice",
		CourierPhone:  "913",
		CourierName:   "CertainlyNotPolice",
		City:          "Moscow",
		StartAddr:     "Red street",
		EndAddr:       "Green street",
		Title:         "Gun",
		Price:         "125",
		DeliveryPrice: "1000",
	}
	exp := `{"order_id":0,"seller_phone":"911","seller_name":"Police","buyer_phone":"912","buyer_name":"NotPolice","courier_phone":"913","courier_name":"CertainlyNotPolice","city":"Moscow","start_addr":"Red street","end_addr":"Green street","title":"Gun","price":"125","delivery_price":"1000"}`
	c.EncodeGetOrder(w, o)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := string(body)
	if bodyStr != exp {
		t.Errorf("got %v want %v",
			bodyStr, exp)
	}
}

func TestEncodeGetOrders(t *testing.T) {
	c := NewSerializator()
	w := httptest.NewRecorder()
	o := &v1.ShortOrdersResponse{Orders: []v1.Order{
		v1.Order{
			OrderID:       0,
			SellerName:    "Me",
			BuyerName:     "You",
			Title:         "Apple",
			Price:         "1349",
			DeliveryPrice: "12",
		},
		v1.Order{
			OrderID:       1,
			SellerName:    "Me",
			BuyerName:     "You",
			Title:         "Pencil",
			Price:         "2",
			DeliveryPrice: "12",
		},
	}}
	exp := `{"orders":[{"order_id":0,"seller_name":"Me","buyer_name":"You","title":"Apple","price":"1349","delivery_price":"12"},{"order_id":1,"seller_name":"Me","buyer_name":"You","title":"Pencil","price":"2","delivery_price":"12"}]}`
	c.EncodeGetOrders(w, o.Orders)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := string(body)
	if bodyStr != exp {
		t.Errorf("got %v want %v",
			bodyStr, exp)
	}
}

