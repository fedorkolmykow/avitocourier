package v1

// Order ...
type Order struct {
	OrderID int
	SellerPhone string
	SellerName string
	BuyerPhone string
	BuyerName string
	CourierPhone string
	CourierName string
	City string
	StartAddr string
	EndAddr string
	Title string
	Price string
	DeliveryPrice string
}
