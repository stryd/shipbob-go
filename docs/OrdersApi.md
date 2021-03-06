# \OrdersApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrdersApi.md#CancelOrder) | **Post** /order/{orderId}/cancel | Cancel single Order by Order ID
[**CancelOrderShipment**](OrdersApi.md#CancelOrderShipment) | **Post** /order/{orderId}/shipment/{shipmentId}/cancel | Cancel one Shipment by Order Id and Shipment Id
[**CancelShipment**](OrdersApi.md#CancelShipment) | **Post** /shipment/{shipmentId}/cancel | Cancel one Shipment by Shipment Id
[**CancelShipmentBulk**](OrdersApi.md#CancelShipmentBulk) | **Post** /shipment/cancelbulk | Cancel multiple Shipments by Shipment Id
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /order | Create Order
[**CreateOrderEstimate**](OrdersApi.md#CreateOrderEstimate) | **Post** /order/estimate | Estimate Fulfillment Cost For Order
[**GetOrder**](OrdersApi.md#GetOrder) | **Get** /order/{orderId} | Get Order
[**GetOrderShipment**](OrdersApi.md#GetOrderShipment) | **Get** /order/{orderId}/shipment/{shipmentId} | Get one Shipment by Order Id and Shipment Id
[**GetOrderShipmentLogs**](OrdersApi.md#GetOrderShipmentLogs) | **Get** /order/{orderId}/shipment/{shipmentId}/logs | Get logs for one Shipment by Order Id and Shipment Id
[**GetOrderShipmentTimeline**](OrdersApi.md#GetOrderShipmentTimeline) | **Get** /order/{orderId}/shipment/{shipmentId}/timeline | Get one Shipment&#39;s status timeline by Order Id and Shipment Id
[**GetOrderShipments**](OrdersApi.md#GetOrderShipments) | **Get** /order/{orderId}/shipment | Get all Shipments for Order
[**GetOrders**](OrdersApi.md#GetOrders) | **Get** /order | Get Orders
[**GetShipment**](OrdersApi.md#GetShipment) | **Get** /shipment/{shipmentId} | Get one Shipment by Shipment Id
[**GetShipmentLogs**](OrdersApi.md#GetShipmentLogs) | **Get** /shipment/{shipmentId}/logs | Get logs for one Shipment by Shipment Id
[**GetShipmentTimeline**](OrdersApi.md#GetShipmentTimeline) | **Get** /shipment/{shipmentId}/timeline | Get one Shipment&#39;s status timeline by Shipment Id
[**GetShippingMethodCollection**](OrdersApi.md#GetShippingMethodCollection) | **Get** /shippingmethod | Get shipping methods



## CancelOrder

> CanceledOrder CancelOrder(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

Cancel single Order by Order ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | The order ID to cancel
    shipbobChannelId := int32(56) // int32 | Channel ID for Operation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CancelOrder(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrder`: CanceledOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order ID to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel ID for Operation | 

### Return type

[**CanceledOrder**](CanceledOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrderShipment

> Shipment CancelOrderShipment(ctx, shipmentId, orderId).ShipbobChannelId(shipbobChannelId).Execute()

Cancel one Shipment by Order Id and Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipmentId := int32(56) // int32 | The shipment id to get
    orderId := "orderId_example" // string | 
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CancelOrderShipment(context.Background(), shipmentId, orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelOrderShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrderShipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CancelOrderShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Shipment**](Shipment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelShipment

> Shipment CancelShipment(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Cancel one Shipment by Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CancelShipment(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelShipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CancelShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Shipment**](Shipment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelShipmentBulk

> InlineResponse200 CancelShipmentBulk(ctx).ShipbobChannelId(shipbobChannelId).CancelShipment(cancelShipment).Execute()

Cancel multiple Shipments by Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipbobChannelId := int32(56) // int32 | Channel ID for Operation
    cancelShipment := *openapiclient.NewCancelShipment() // CancelShipment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CancelShipmentBulk(context.Background()).ShipbobChannelId(shipbobChannelId).CancelShipment(cancelShipment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelShipmentBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelShipmentBulk`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CancelShipmentBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel ID for Operation | 
 **cancelShipment** | [**CancelShipment**](CancelShipment.md) |  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> Order CreateOrder(ctx).ShipbobChannelId(shipbobChannelId).CreateOrder(createOrder).Execute()

Create Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    createOrder := *openapiclient.NewCreateOrder([]openapiclient.CreateOrderProducts{*openapiclient.NewCreateOrderProducts(int32(123))}, *openapiclient.NewRecipientInfo(*openapiclient.NewOrderAddress("100 Nowhere Blvd", "Gotham City", "US"), "John Doe"), "ReferenceId_example", "ShippingMethod_example") // CreateOrder |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CreateOrder(context.Background()).ShipbobChannelId(shipbobChannelId).CreateOrder(createOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **createOrder** | [**CreateOrder**](CreateOrder.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrderEstimate

> OrderEstimate CreateOrderEstimate(ctx).ShipbobChannelId(shipbobChannelId).EstimateFulfillmentRequest(estimateFulfillmentRequest).Execute()

Estimate Fulfillment Cost For Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    estimateFulfillmentRequest := *openapiclient.NewEstimateFulfillmentRequest(*openapiclient.NewEstimationAddress("US"), []openapiclient.EstimateFulfillmentRequestProducts{*openapiclient.NewEstimateFulfillmentRequestProducts(int32(123))}) // EstimateFulfillmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.CreateOrderEstimate(context.Background()).ShipbobChannelId(shipbobChannelId).EstimateFulfillmentRequest(estimateFulfillmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderEstimate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderEstimate`: OrderEstimate
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderEstimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **estimateFulfillmentRequest** | [**EstimateFulfillmentRequest**](EstimateFulfillmentRequest.md) |  | 

### Return type

[**OrderEstimate**](OrderEstimate.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

Get Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | 
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrder(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Order**](Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderShipment

> Shipment GetOrderShipment(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get one Shipment by Order Id and Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | The order id to get the shipment for
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderShipment(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderShipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Shipment**](Shipment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderShipmentLogs

> []ShipmentLog GetOrderShipmentLogs(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get logs for one Shipment by Order Id and Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | The order id to get the shipment for
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderShipmentLogs(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderShipmentLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderShipmentLogs`: []ShipmentLog
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderShipmentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderShipmentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipmentLog**](ShipmentLog.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderShipmentTimeline

> []ShipmentLog GetOrderShipmentTimeline(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get one Shipment's status timeline by Order Id and Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | The order id to get the shipment for
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderShipmentTimeline(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderShipmentTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderShipmentTimeline`: []ShipmentLog
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderShipmentTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderShipmentTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipmentLog**](ShipmentLog.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderShipments

> []Shipment GetOrderShipments(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

Get all Shipments for Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := int32(56) // int32 | The order id to get shipments for
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrderShipments(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderShipments`: []Shipment
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderShipments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get shipments for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]Shipment**](Shipment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> []Order GetOrders(ctx).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).StartDate(startDate).EndDate(endDate).SortOrder(sortOrder).HasTracking(hasTracking).LastUpdateStartDate(lastUpdateStartDate).LastUpdateEndDate(lastUpdateEndDate).ShipbobChannelId(shipbobChannelId).Execute()

Get Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page of orders to get (optional)
    limit := int32(56) // int32 | Amount of orders per page to request (optional)
    iDs := []int32{int32(123)} // []int32 | order ids to filter by (optional)
    referenceIds := []string{"Inner_example"} // []string | Reference ids to filter by (optional)
    startDate := time.Now() // time.Time | Start date to filter orders inserted later than (optional)
    endDate := time.Now() // time.Time | End date to filter orders inserted earlier than (optional)
    sortOrder := "sortOrder_example" // string | Order to sort results in (optional)
    hasTracking := true // bool | Has any portion of this order been assigned a tracking number (optional)
    lastUpdateStartDate := time.Now() // time.Time | Start date to filter orders updated later than (optional)
    lastUpdateEndDate := time.Now() // time.Time | End date to filter orders updated later than (optional)
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrders(context.Background()).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).StartDate(startDate).EndDate(endDate).SortOrder(sortOrder).HasTracking(hasTracking).LastUpdateStartDate(lastUpdateStartDate).LastUpdateEndDate(lastUpdateEndDate).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: []Order
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of orders to get | 
 **limit** | **int32** | Amount of orders per page to request | 
 **iDs** | **[]int32** | order ids to filter by | 
 **referenceIds** | **[]string** | Reference ids to filter by | 
 **startDate** | **time.Time** | Start date to filter orders inserted later than | 
 **endDate** | **time.Time** | End date to filter orders inserted earlier than | 
 **sortOrder** | **string** | Order to sort results in | 
 **hasTracking** | **bool** | Has any portion of this order been assigned a tracking number | 
 **lastUpdateStartDate** | **time.Time** | Start date to filter orders updated later than | 
 **lastUpdateEndDate** | **time.Time** | End date to filter orders updated later than | 
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]Order**](Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipment

> Shipment GetShipment(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get one Shipment by Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetShipment(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Shipment**](Shipment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentLogs

> []ShipmentLog GetShipmentLogs(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get logs for one Shipment by Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetShipmentLogs(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetShipmentLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipmentLogs`: []ShipmentLog
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetShipmentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipmentLog**](ShipmentLog.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentTimeline

> []ShipmentLog GetShipmentTimeline(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

Get one Shipment's status timeline by Shipment Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shipmentId := int32(56) // int32 | The shipment id to get
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetShipmentTimeline(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetShipmentTimeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipmentTimeline`: []ShipmentLog
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetShipmentTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipmentLog**](ShipmentLog.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingMethodCollection

> []ShipMethodDetail GetShippingMethodCollection(ctx).Page(page).Limit(limit).Execute()

Get shipping methods



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page of orders to get (optional)
    limit := int32(56) // int32 | Amount of records per page to request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetShippingMethodCollection(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetShippingMethodCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShippingMethodCollection`: []ShipMethodDetail
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetShippingMethodCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingMethodCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of orders to get | 
 **limit** | **int32** | Amount of records per page to request | 

### Return type

[**[]ShipMethodDetail**](ShipMethodDetail.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

