# \OrdersApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderEstimatePost**](OrdersApi.md#OrderEstimatePost) | **Post** /order/estimate | Estimate Fulfillment Cost For Order
[**OrderGet**](OrdersApi.md#OrderGet) | **Get** /order | Get Orders
[**OrderOrderIdCancelPost**](OrdersApi.md#OrderOrderIdCancelPost) | **Post** /order/{orderId}/cancel | Cancel single Order by Order ID
[**OrderOrderIdGet**](OrdersApi.md#OrderOrderIdGet) | **Get** /order/{orderId} | Get Order
[**OrderOrderIdShipmentGet**](OrdersApi.md#OrderOrderIdShipmentGet) | **Get** /order/{orderId}/shipment | Get all Shipments for Order
[**OrderOrderIdShipmentShipmentIdCancelPost**](OrdersApi.md#OrderOrderIdShipmentShipmentIdCancelPost) | **Post** /order/{orderId}/shipment/{shipmentId}/cancel | Cancel one Shipment by Order Id and Shipment Id
[**OrderOrderIdShipmentShipmentIdGet**](OrdersApi.md#OrderOrderIdShipmentShipmentIdGet) | **Get** /order/{orderId}/shipment/{shipmentId} | Get one Shipment by Order Id and Shipment Id
[**OrderOrderIdShipmentShipmentIdLogsGet**](OrdersApi.md#OrderOrderIdShipmentShipmentIdLogsGet) | **Get** /order/{orderId}/shipment/{shipmentId}/logs | Get logs for one Shipment by Order Id and Shipment Id
[**OrderOrderIdShipmentShipmentIdTimelineGet**](OrdersApi.md#OrderOrderIdShipmentShipmentIdTimelineGet) | **Get** /order/{orderId}/shipment/{shipmentId}/timeline | Get one Shipment&#39;s status timeline by Order Id and Shipment Id
[**OrderPost**](OrdersApi.md#OrderPost) | **Post** /order | Create Order
[**ShipmentCancelbulkPost**](OrdersApi.md#ShipmentCancelbulkPost) | **Post** /shipment/cancelbulk | Cancel multiple Shipments by Shipment Id
[**ShipmentShipmentIdCancelPost**](OrdersApi.md#ShipmentShipmentIdCancelPost) | **Post** /shipment/{shipmentId}/cancel | Cancel one Shipment by Shipment Id
[**ShipmentShipmentIdGet**](OrdersApi.md#ShipmentShipmentIdGet) | **Get** /shipment/{shipmentId} | Get one Shipment by Shipment Id
[**ShipmentShipmentIdLogsGet**](OrdersApi.md#ShipmentShipmentIdLogsGet) | **Get** /shipment/{shipmentId}/logs | Get logs for one Shipment by Shipment Id
[**ShipmentShipmentIdTimelineGet**](OrdersApi.md#ShipmentShipmentIdTimelineGet) | **Get** /shipment/{shipmentId}/timeline | Get one Shipment&#39;s status timeline by Shipment Id
[**ShippingmethodGet**](OrdersApi.md#ShippingmethodGet) | **Get** /shippingmethod | Get shipping methods



## OrderEstimatePost

> ShipBobOrdersPresentationViewModelsEstimateViewModel OrderEstimatePost(ctx).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel(shipBobOrdersPresentationModelsEstimateFulfillmentRequestModel).Execute()

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
    shipBobOrdersPresentationModelsEstimateFulfillmentRequestModel := *openapiclient.NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel(*openapiclient.NewShipBobOrdersPresentationViewModelsEstimationAddressViewModel("US"), []openapiclient.ShipBobOrdersPresentationModelsEstimateProductInfoModel{*openapiclient.NewShipBobOrdersPresentationModelsEstimateProductInfoModel(int32(123))}) // ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.OrderEstimatePost(context.Background()).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel(shipBobOrdersPresentationModelsEstimateFulfillmentRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderEstimatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderEstimatePost`: ShipBobOrdersPresentationViewModelsEstimateViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderEstimatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderEstimatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **shipBobOrdersPresentationModelsEstimateFulfillmentRequestModel** | [**ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel**](ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel.md) |  | 

### Return type

[**ShipBobOrdersPresentationViewModelsEstimateViewModel**](ShipBob.Orders.Presentation.ViewModels.EstimateViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderGet

> []ShipBobOrdersPresentationViewModelsOrderViewModel OrderGet(ctx).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).StartDate(startDate).EndDate(endDate).SortOrder(sortOrder).HasTracking(hasTracking).LastUpdateStartDate(lastUpdateStartDate).LastUpdateEndDate(lastUpdateEndDate).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderGet(context.Background()).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).StartDate(startDate).EndDate(endDate).SortOrder(sortOrder).HasTracking(hasTracking).LastUpdateStartDate(lastUpdateStartDate).LastUpdateEndDate(lastUpdateEndDate).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderGet`: []ShipBobOrdersPresentationViewModelsOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderGetRequest struct via the builder pattern


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

[**[]ShipBobOrdersPresentationViewModelsOrderViewModel**](ShipBob.Orders.Presentation.ViewModels.OrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdCancelPost

> ShipBobOrdersPresentationViewModelsCanceledOrderViewModel OrderOrderIdCancelPost(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdCancelPost(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdCancelPost`: ShipBobOrdersPresentationViewModelsCanceledOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order ID to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel ID for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsCanceledOrderViewModel**](ShipBob.Orders.Presentation.ViewModels.CanceledOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdGet

> ShipBobOrdersPresentationViewModelsOrderViewModel OrderOrderIdGet(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdGet(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdGet`: ShipBobOrdersPresentationViewModelsOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsOrderViewModel**](ShipBob.Orders.Presentation.ViewModels.OrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdShipmentGet

> []ShipBobOrdersPresentationViewModelsShipmentViewModel OrderOrderIdShipmentGet(ctx, orderId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdShipmentGet(context.Background(), orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdShipmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdShipmentGet`: []ShipBobOrdersPresentationViewModelsShipmentViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdShipmentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get shipments for | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdShipmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdShipmentShipmentIdCancelPost

> ShipBobOrdersPresentationViewModelsShipmentViewModel OrderOrderIdShipmentShipmentIdCancelPost(ctx, shipmentId, orderId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdShipmentShipmentIdCancelPost(context.Background(), shipmentId, orderId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdShipmentShipmentIdCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdShipmentShipmentIdCancelPost`: ShipBobOrdersPresentationViewModelsShipmentViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdShipmentShipmentIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdShipmentShipmentIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdShipmentShipmentIdGet

> ShipBobOrdersPresentationViewModelsShipmentViewModel OrderOrderIdShipmentShipmentIdGet(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdShipmentShipmentIdGet(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdShipmentShipmentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdShipmentShipmentIdGet`: ShipBobOrdersPresentationViewModelsShipmentViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdShipmentShipmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdShipmentShipmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdShipmentShipmentIdLogsGet

> []ShipBobOrdersPresentationViewModelsShipmentLogViewModel OrderOrderIdShipmentShipmentIdLogsGet(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdShipmentShipmentIdLogsGet(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdShipmentShipmentIdLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdShipmentShipmentIdLogsGet`: []ShipBobOrdersPresentationViewModelsShipmentLogViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdShipmentShipmentIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdShipmentShipmentIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipmentLogViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentLogViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdShipmentShipmentIdTimelineGet

> []ShipBobOrdersPresentationViewModelsShipmentLogViewModel OrderOrderIdShipmentShipmentIdTimelineGet(ctx, orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.OrderOrderIdShipmentShipmentIdTimelineGet(context.Background(), orderId, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdShipmentShipmentIdTimelineGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdShipmentShipmentIdTimelineGet`: []ShipBobOrdersPresentationViewModelsShipmentLogViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdShipmentShipmentIdTimelineGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order id to get the shipment for | 
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdShipmentShipmentIdTimelineGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipmentLogViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentLogViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPost

> ShipBobOrdersPresentationViewModelsOrderViewModel OrderPost(ctx).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsCreateOrderModel(shipBobOrdersPresentationModelsCreateOrderModel).Execute()

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
    shipBobOrdersPresentationModelsCreateOrderModel := *openapiclient.NewShipBobOrdersPresentationModelsCreateOrderModel([]openapiclient.ShipBobOrdersPresentationModelsAddProductToOrderModel{openapiclient.ShipBob.Orders.Presentation.Models.AddProductToOrderModel{ShipBobOrdersPresentationModelsAddProductToOrderByProductIdModel: openapiclient.NewShipBobOrdersPresentationModelsAddProductToOrderByProductIdModel(int32(123), int32(123))}}, *openapiclient.NewShipBobOrdersPresentationViewModelsRecipientInfoViewModel(*openapiclient.NewShipBobOrdersPresentationViewModelsAddressViewModel("100 Nowhere Blvd", "Gotham City", "US"), "John Doe"), "ReferenceId_example", "ShippingMethod_example") // ShipBobOrdersPresentationModelsCreateOrderModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.OrderPost(context.Background()).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsCreateOrderModel(shipBobOrdersPresentationModelsCreateOrderModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderPost`: ShipBobOrdersPresentationViewModelsOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **shipBobOrdersPresentationModelsCreateOrderModel** | [**ShipBobOrdersPresentationModelsCreateOrderModel**](ShipBobOrdersPresentationModelsCreateOrderModel.md) |  | 

### Return type

[**ShipBobOrdersPresentationViewModelsOrderViewModel**](ShipBob.Orders.Presentation.ViewModels.OrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentCancelbulkPost

> ShipBobOrdersPresentationViewModelsCanceledShipmentsViewModel ShipmentCancelbulkPost(ctx).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsCancelShipmentsModel(shipBobOrdersPresentationModelsCancelShipmentsModel).Execute()

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
    shipBobOrdersPresentationModelsCancelShipmentsModel := *openapiclient.NewShipBobOrdersPresentationModelsCancelShipmentsModel() // ShipBobOrdersPresentationModelsCancelShipmentsModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ShipmentCancelbulkPost(context.Background()).ShipbobChannelId(shipbobChannelId).ShipBobOrdersPresentationModelsCancelShipmentsModel(shipBobOrdersPresentationModelsCancelShipmentsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShipmentCancelbulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentCancelbulkPost`: ShipBobOrdersPresentationViewModelsCanceledShipmentsViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShipmentCancelbulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShipmentCancelbulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel ID for Operation | 
 **shipBobOrdersPresentationModelsCancelShipmentsModel** | [**ShipBobOrdersPresentationModelsCancelShipmentsModel**](ShipBobOrdersPresentationModelsCancelShipmentsModel.md) |  | 

### Return type

[**ShipBobOrdersPresentationViewModelsCanceledShipmentsViewModel**](ShipBob.Orders.Presentation.ViewModels.CanceledShipmentsViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentShipmentIdCancelPost

> ShipBobOrdersPresentationViewModelsShipmentViewModel ShipmentShipmentIdCancelPost(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.ShipmentShipmentIdCancelPost(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShipmentShipmentIdCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentShipmentIdCancelPost`: ShipBobOrdersPresentationViewModelsShipmentViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShipmentShipmentIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentShipmentIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentShipmentIdGet

> ShipBobOrdersPresentationViewModelsShipmentViewModel ShipmentShipmentIdGet(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.ShipmentShipmentIdGet(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShipmentShipmentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentShipmentIdGet`: ShipBobOrdersPresentationViewModelsShipmentViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShipmentShipmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentShipmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentShipmentIdLogsGet

> []ShipBobOrdersPresentationViewModelsShipmentLogViewModel ShipmentShipmentIdLogsGet(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.ShipmentShipmentIdLogsGet(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShipmentShipmentIdLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentShipmentIdLogsGet`: []ShipBobOrdersPresentationViewModelsShipmentLogViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShipmentShipmentIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentShipmentIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipmentLogViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentLogViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipmentShipmentIdTimelineGet

> []ShipBobOrdersPresentationViewModelsShipmentLogViewModel ShipmentShipmentIdTimelineGet(ctx, shipmentId).ShipbobChannelId(shipbobChannelId).Execute()

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
    resp, r, err := api_client.OrdersApi.ShipmentShipmentIdTimelineGet(context.Background(), shipmentId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShipmentShipmentIdTimelineGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipmentShipmentIdTimelineGet`: []ShipBobOrdersPresentationViewModelsShipmentLogViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShipmentShipmentIdTimelineGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **int32** | The shipment id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipmentShipmentIdTimelineGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipmentLogViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipmentLogViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShippingmethodGet

> []ShipBobOrdersPresentationViewModelsShipMethodDetailViewModel ShippingmethodGet(ctx).Page(page).Limit(limit).Execute()

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
    resp, r, err := api_client.OrdersApi.ShippingmethodGet(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ShippingmethodGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShippingmethodGet`: []ShipBobOrdersPresentationViewModelsShipMethodDetailViewModel
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ShippingmethodGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShippingmethodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of orders to get | 
 **limit** | **int32** | Amount of records per page to request | 

### Return type

[**[]ShipBobOrdersPresentationViewModelsShipMethodDetailViewModel**](ShipBob.Orders.Presentation.ViewModels.ShipMethodDetailViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

