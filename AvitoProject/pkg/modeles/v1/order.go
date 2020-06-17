package v1

// Order ...
type Order struct {
	OrderID       int    `json:"order_id"`
	SellerPhone   string `json:"seller_phone,omitempty"`
	SellerName    string `json:"seller_name"`
	BuyerPhone    string `json:"buyer_phone,omitempty"`
	BuyerName     string `json:"buyer_name"`
	CourierPhone  string `json:"courier_phone,omitempty"`
	CourierName   string `json:"courier_name,omitempty"`
	City          string `json:"city,omitempty"`
	StartAddr     string `json:"start_addr,omitempty"`
	EndAddr       string `json:"end_addr,omitempty"`
	Title         string `json:"title"`
	Price         string `json:"price"`
	DeliveryPrice string `json:"delivery_price"`
}

// OrderCreationRequest ...
type OrderCreationRequest struct {
	BuyerID   int `json:"buyer_id"`
	EndAddrID int `json:"end_addr_id"`
	NoticeID  int `json:"notice_id"`
}

// OrderCreationResponse ...
type OrderCreationResponse struct {
	OrderID int `json:"order_id"`
}

// OrderDeliveryPriceResponse ...
type OrderDeliveryPriceResponse struct {
	DeliveryPrice int `json:"delivery_price"`
}

// ShortOrdersResponse ...
type ShortOrdersResponse struct {
	Orders []Order `json:"orders"`
}