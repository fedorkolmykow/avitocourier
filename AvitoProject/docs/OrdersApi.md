# \OrdersApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersGet**](OrdersApi.md#OrdersGet) | **Get** /orders | Получить список заказов на доставку
[**OrdersOrderIdGet**](OrdersApi.md#OrdersOrderIdGet) | **Get** /orders/{order_id} | Получить подробное описание заказа
[**OrdersPost**](OrdersApi.md#OrdersPost) | **Post** /orders | Добавить новую доставку
[**OrdersPriceGet**](OrdersApi.md#OrdersPriceGet) | **Get** /orders/price | Получить цену доставки заказа


# **OrdersGet**
> []ShortOrdersResponse OrdersGet(ctx, optional)
Получить список заказов на доставку

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrdersApiOrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seller** | **optional.Int32**| ID продавца | 

### Return type

[**[]ShortOrdersResponse**](ShortOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersOrderIdGet**
> Order OrdersOrderIdGet(ctx, orderId)
Получить подробное описание заказа

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**|  | 

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersPost**
> OrderCreationResponse OrdersPost(ctx, optional)
Добавить новую доставку

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrdersApiOrdersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OrderCreationRequest**](OrderCreationRequest.md)|  | 

### Return type

[**OrderCreationResponse**](OrderCreationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersPriceGet**
> OrderDeliveryPriceResponse OrdersPriceGet(ctx, optional)
Получить цену доставки заказа

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrdersApiOrdersPriceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersPriceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noticeId** | **optional.Int32**| ID объявления | 
 **endAddrId** | **optional.Int32**| ID адреса для доставки | 

### Return type

[**OrderDeliveryPriceResponse**](OrderDeliveryPriceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

