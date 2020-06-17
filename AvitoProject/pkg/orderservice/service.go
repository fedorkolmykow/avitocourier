package orderservice

import (
	"errors"
	"fmt"

	v1 "AvitoProject/pkg/modeles/v1"
)

type calculator interface {
	Calculate(ad1, ad2 int) (price int)
}

type db interface {
	SelectAllOrders(sellerID int) (orders []v1.Order, err error)
	SelectOrder(orderID int) (order *v1.Order, err error)
	SelectSellerFromNotice(noticeID int) (sellerID int, err error)
	InsertOrder(courierID, buyerID, endAddID, noticeID, deliveryPrice int) (orderID int, err error)
	SelectMostFreeCourier() (courierID int, err error)
}

// Service ...
type Service interface {
	GetOrder(orderID int) (order *v1.Order, err error)
	GetAllOrders(sellerID int) (orders []v1.Order, err error)
	SetOrder(oc *v1.OrderCreation) (orderID int, err error)
	CalculatePrice(addrID, noticeID int) (price int, err error)
}

type service struct {
	calc  calculator
	dbCon db
}

func (s *service) GetOrder(orderID int) (order *v1.Order, err error) {
	return s.dbCon.SelectOrder(orderID)
}

func (s *service) SetOrder(oc *v1.OrderCreation) (orderID int, err error) {
	sellerID, err := s.dbCon.SelectSellerFromNotice(oc.NoticeID)
	if err != nil {
		return
	}
	if sellerID == oc.BuyerID {
		err = errors.New("seller and buyer cannot be same")
		fmt.Println(err)
		return
	}
	deliveryPrice := s.calc.Calculate(oc.EndAddrID, oc.NoticeID)
	courierID, err := s.dbCon.SelectMostFreeCourier()
	orderID, err = s.dbCon.InsertOrder(courierID, oc.BuyerID, oc.EndAddrID, oc.NoticeID, deliveryPrice)
	return
}

func (s *service) GetAllOrders(sellerID int) (orders []v1.Order, err error) {
	return s.dbCon.SelectAllOrders(sellerID)
}

func (s *service) CalculatePrice(addrID, noticeID int) (price int, err error) {
	return s.calc.Calculate(addrID, noticeID), nil
}

// NewService returns a new Service instance.
func NewService(
	calc calculator,
	dbCon db,
) Service {
	return &service{
		calc:  calc,
		dbCon: dbCon,
	}
}
