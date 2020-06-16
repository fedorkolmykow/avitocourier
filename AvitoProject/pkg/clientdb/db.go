package clientdb

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // to call init() in which driver will be registered for usage in database/sql

	v1 "AvitoProject/pkg/modeles/v1"
)

const(
	requestOrders = `SELECT 
	Orders.order_id,  
	Seller.name,
	Buyer.name,
	Courier.name,
	Notice.title,
	Notice.price,
	Orders.delivery_price
FROM Orders
LEFT JOIN Client AS Buyer ON Buyer.client_id = Orders.buyer_id
LEFT JOIN Courier ON Courier.courier_id = Orders.courier_id
LEFT JOIN Notice ON Notice.notice_id = Orders.notice_id
LEFT JOIN Client AS Seller ON Notice.seller_id = Seller.client_id
WHERE Seller.client_id = $1;`
	requestOrder = `SELECT 
	Orders.order_id, 
	Seller.phone, 
	Seller.name,
	Buyer.phone,
	Buyer.name,
	Courier.phone,
	Courier.name,
	City.city,
	Start_add.address,
	End_add.address,
	Notice.title,
	Notice.price,
	Orders.delivery_price
FROM Orders
LEFT JOIN Client AS Buyer ON Buyer.client_id = Orders.buyer_id
LEFT JOIN Courier ON Courier.courier_id = Orders.courier_id
LEFT JOIN Notice ON Notice.notice_id = Orders.notice_id
LEFT JOIN Client AS Seller ON Notice.seller_id = Seller.client_id
LEFT JOIN Address AS Start_add ON Start_add.address_id = Notice.start_address_id
LEFT JOIN Address AS End_add ON End_add.address_id = Orders.end_address_id
LEFT JOIN City ON End_add.city_id = City.city_id WHERE Orders.order_id = $1;`
	insertOrder = `INSERT INTO Orders (courier_id, buyer_id, end_address_id, notice_id, delivery_price)
	VALUES	
	($1, $2, $3, $4, $5);`
	requestMostFreeCourier = `SELECT Courier.courier_id, 0 as co FROM Courier
WHERE courier_id not in (SELECT courier_id FROM Orders)
UNION
SELECT courier_id, count(*) as co
FROM Orders
group by courier_id
order by co asc
limit 1;`
	requestSellerFromNotice = `SELECT seller_id From Notice
WHERE notice_id = $1;`
)

// Db ...
type Db interface {
	SelectAllOrders(sellerID int) (orders []v1.Order,err error)
	SelectOrder(orderID int) (order *v1.Order,err error)
	SelectSellerFromNotice(noticeID int) (sellerID int, err error)
	InsertOrder(courierID, buyerID, endAddID, noticeID, deliveryPrice int) (err error)
	SelectMostFreeCourier() (courierID int, err error)
}

type db struct {
	dbCon *sql.DB
}


func (d *db) SelectAllOrders(sellerID int) (orders []v1.Order,err error){
	orders = []v1.Order{}

	rows, err := d.dbCon.Query(requestOrders, sellerID)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		o := v1.Order{}
		err = rows.Scan(&o.OrderID, &o.SellerName, &o.BuyerName,
			&o.CourierName, &o.Title, &o.Price, &o.DeliveryPrice)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (d *db) SelectOrder(orderID int) (order *v1.Order, err error){
    o := v1.Order{}

	row := d.dbCon.QueryRow(requestOrder, orderID)
	err = row.Scan(&o.OrderID, &o.SellerPhone, &o.SellerName,
	&o.BuyerPhone, &o.BuyerName, &o.CourierPhone, &o.CourierName,
	&o.City, &o.StartAddr, &o.EndAddr, &o.Title, &o.Price, &o.DeliveryPrice)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &o, nil
}

func (d *db) InsertOrder(courierID, buyerID, endAddID, noticeID, deliveryPrice int) (err error){
	_, err = d.dbCon.Exec(insertOrder, courierID, buyerID, endAddID, noticeID, deliveryPrice)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *db) SelectMostFreeCourier() (courierID int, err error){
	row := d.dbCon.QueryRow(requestMostFreeCourier)
	err = row.Scan(&courierID)
	return
}

func (d *db) SelectSellerFromNotice(noticeID int) (sellerID int, err error){
	row := d.dbCon.QueryRow(requestSellerFromNotice, noticeID)
	err = row.Scan(&sellerID)
	return
}

// NewDb returns a new Db instance.
func NewDb() Db{
	connStr := "user=postgres password=avitopass dbname=avito sslmode=disable port=5555"
	dbCon, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = dbCon.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &db{dbCon:dbCon}
}