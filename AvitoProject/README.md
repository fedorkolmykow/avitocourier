# Go API client for swagger

Заданиe от Авито на стажировку. Подробности можно найти тут https://github.com/avito-tech/safedeal-backend-trainee/

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OrdersApi* | **OrdersGet** | **Get** /orders | Получить список заказов на доставку
*OrdersApi* | **OrdersOrderIdGet** | **Get** /orders/{order_id} | Получить подробное описание заказа
*OrdersApi* | **OrdersPost** | **Post** /orders | Добавить новую доставку
*OrdersApi* | **OrdersPriceGet** | **Get** /orders/price | Получить цену доставки заказа


## Documentation For Models

 - [Order](docs/Order.md)
 - [OrderCreation](docs/OrderCreation.md)
 - [OrderCreationResp](docs/OrderCreationResp.md)
 - [OrderDeliveryPrice](docs/OrderDeliveryPrice.md)
 - [ShortOrder](docs/ShortOrder.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

fedorkolmykow@gmail.com

