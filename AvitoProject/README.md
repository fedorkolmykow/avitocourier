# Go API client for swagger

Задание от Авито на стажировку. Подробности можно найти тут https://github.com/avito-tech/safedeal-backend-trainee/

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
 - [OrderCreationRequest](docs/OrderCreationRequest.md)
 - [OrderCreationResponse](docs/OrderCreationResponse.md)
 - [OrderDeliveryPriceResponse](docs/OrderDeliveryPriceResponse.md)
 - [ShortOrdersResponse](docs/ShortOrdersResponse.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

fedorkolmykow@gmail.com

