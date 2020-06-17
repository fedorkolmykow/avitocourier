package serialization

import (
	v1 "AvitoProject/pkg/modeles/v1"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeSetOrder(t *testing.T) {
	c := NewSerializator()
	exp_oc := &v1.OrderCreation{
		BuyerID:   1,
		EndAddrID: 2,
		NoticeID:  3,
	}
	body, err := json.Marshal(exp_oc)
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
	if *oc != *exp_oc {
		t.Errorf("got %v want %v",
			oc, exp_oc)
	}
}

func TestEncodeSetOrder(t *testing.T) {
	c := NewSerializator()
	w := httptest.NewRecorder()
	orderId := 1
	exp := `{"order_id":1}`
	c.EncodeSetOrder(w, orderId)
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

//func TestEncodeGetDeliveryPrice(t *testing.T) {
//	c := NewSerializator()
//
//}
//
//func TestEncodeGetOrder(t *testing.T) {
//	c := NewSerializator()
//
//}
//
//func TestEncodeGetOrders(t *testing.T) {
//	c := NewSerializator()
//
//}
