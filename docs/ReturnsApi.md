# \ReturnsApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnGet**](ReturnsApi.md#ReturnGet) | **Get** /return | Get Return Orders
[**ReturnIdCancelPost**](ReturnsApi.md#ReturnIdCancelPost) | **Post** /return/{id}/cancel | Cancel Return Order
[**ReturnIdGet**](ReturnsApi.md#ReturnIdGet) | **Get** /return/{id} | Get Return Order
[**ReturnIdPut**](ReturnsApi.md#ReturnIdPut) | **Put** /return/{id} | Modify Return Order
[**ReturnIdStatushistoryGet**](ReturnsApi.md#ReturnIdStatushistoryGet) | **Get** /return/{id}/statushistory | Get One Return&#39;s status history
[**ReturnPost**](ReturnsApi.md#ReturnPost) | **Post** /return | Create Return Order



## ReturnGet

> []ShipbobReturnsPublicApiViewModelsReturnOrderViewModel ReturnGet(ctx).Page(page).Limit(limit).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).IDs(iDs).ReferenceIds(referenceIds).Status(status).FulfillmentCenterIds(fulfillmentCenterIds).TrackingNumbers(trackingNumbers).OriginalShipmentIds(originalShipmentIds).InventoryIds(inventoryIds).ShipbobChannelId(shipbobChannelId).Execute()

Get Return Orders

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
    page := int32(56) // int32 | Page of return orders to get (optional)
    limit := int32(56) // int32 | Amount of return orders per page to request (optional)
    sortOrder := openapiclient.Shipbob.Returns.Public.Common.SortOrder("Newest") // ShipbobReturnsPublicCommonSortOrder | Order to sort results by (optional)
    startDate := time.Now() // time.Time | Start date to filter orders inserted later than (optional)
    endDate := time.Now() // time.Time | End date to filter orders inserted earlier than (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated list of return orders ids to filter by (optional)
    referenceIds := []string{"Inner_example"} // []string | Comma separated list of reference ids to filter by (optional)
    status := []openapiclient.ShipbobReturnsPublicCommonReturnStatus{openapiclient.Shipbob.Returns.Public.Common.ReturnStatus("AwaitingArrival")} // []ShipbobReturnsPublicCommonReturnStatus | Comma separated list of statuses to filter by (optional)
    fulfillmentCenterIds := []int32{int32(123)} // []int32 | Comma separated list of destination fulfillment center IDs to filter by (optional)
    trackingNumbers := []string{"Inner_example"} // []string | Comma separated list of tracking numbers to filter by (optional)
    originalShipmentIds := []int32{int32(123)} // []int32 | Comma separated list of original shipment IDs to filter by (optional)
    inventoryIds := []int32{int32(123)} // []int32 | Comma separated list of inventory IDs contained in return to filter by (optional)
    shipbobChannelId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnGet(context.Background()).Page(page).Limit(limit).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).IDs(iDs).ReferenceIds(referenceIds).Status(status).FulfillmentCenterIds(fulfillmentCenterIds).TrackingNumbers(trackingNumbers).OriginalShipmentIds(originalShipmentIds).InventoryIds(inventoryIds).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnGet`: []ShipbobReturnsPublicApiViewModelsReturnOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of return orders to get | 
 **limit** | **int32** | Amount of return orders per page to request | 
 **sortOrder** | [**ShipbobReturnsPublicCommonSortOrder**](ShipbobReturnsPublicCommonSortOrder.md) | Order to sort results by | 
 **startDate** | **time.Time** | Start date to filter orders inserted later than | 
 **endDate** | **time.Time** | End date to filter orders inserted earlier than | 
 **iDs** | **[]int32** | Comma separated list of return orders ids to filter by | 
 **referenceIds** | **[]string** | Comma separated list of reference ids to filter by | 
 **status** | [**[]ShipbobReturnsPublicCommonReturnStatus**](Shipbob.Returns.Public.Common.ReturnStatus.md) | Comma separated list of statuses to filter by | 
 **fulfillmentCenterIds** | **[]int32** | Comma separated list of destination fulfillment center IDs to filter by | 
 **trackingNumbers** | **[]string** | Comma separated list of tracking numbers to filter by | 
 **originalShipmentIds** | **[]int32** | Comma separated list of original shipment IDs to filter by | 
 **inventoryIds** | **[]int32** | Comma separated list of inventory IDs contained in return to filter by | 
 **shipbobChannelId** | **int32** |  | 

### Return type

[**[]ShipbobReturnsPublicApiViewModelsReturnOrderViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIdCancelPost

> ShipbobReturnsPublicApiViewModelsReturnOrderViewModel ReturnIdCancelPost(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Cancel Return Order

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
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnIdCancelPost(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnIdCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnIdCancelPost`: ShipbobReturnsPublicApiViewModelsReturnOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipbobReturnsPublicApiViewModelsReturnOrderViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIdGet

> ShipbobReturnsPublicApiViewModelsReturnOrderViewModel ReturnIdGet(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Get Return Order

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
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnIdGet(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnIdGet`: ShipbobReturnsPublicApiViewModelsReturnOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipbobReturnsPublicApiViewModelsReturnOrderViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIdPut

> ShipbobReturnsPublicApiViewModelsReturnOrderViewModel ReturnIdPut(ctx, id).ShipbobChannelId(shipbobChannelId).ShipbobReturnsPublicApiViewModelsCreateReturnViewModel(shipbobReturnsPublicApiViewModelsCreateReturnViewModel).Execute()

Modify Return Order

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
    id := int32(56) // int32 | Id of the return order
    shipbobReturnsPublicApiViewModelsCreateReturnViewModel := *openapiclient.NewShipbobReturnsPublicApiViewModelsCreateReturnViewModel(*openapiclient.NewShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel(int32(123)), []openapiclient.ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel{*openapiclient.NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModel(int32(111222), int32(1))}, "ShipBob_Return_123") // ShipbobReturnsPublicApiViewModelsCreateReturnViewModel | Model defining the return (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnIdPut(context.Background(), id).ShipbobChannelId(shipbobChannelId).ShipbobReturnsPublicApiViewModelsCreateReturnViewModel(shipbobReturnsPublicApiViewModelsCreateReturnViewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnIdPut`: ShipbobReturnsPublicApiViewModelsReturnOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

 **shipbobReturnsPublicApiViewModelsCreateReturnViewModel** | [**ShipbobReturnsPublicApiViewModelsCreateReturnViewModel**](ShipbobReturnsPublicApiViewModelsCreateReturnViewModel.md) | Model defining the return | 

### Return type

[**ShipbobReturnsPublicApiViewModelsReturnOrderViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnIdStatushistoryGet

> ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel ReturnIdStatushistoryGet(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Get One Return's status history

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
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnIdStatushistoryGet(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnIdStatushistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnIdStatushistoryGet`: ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnIdStatushistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnIdStatushistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderStatusHistoryViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnPost

> ShipbobReturnsPublicApiViewModelsReturnOrderViewModel ReturnPost(ctx).ShipbobChannelId(shipbobChannelId).ShipbobReturnsPublicApiViewModelsCreateReturnViewModel(shipbobReturnsPublicApiViewModelsCreateReturnViewModel).Execute()

Create Return Order

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
    shipbobReturnsPublicApiViewModelsCreateReturnViewModel := *openapiclient.NewShipbobReturnsPublicApiViewModelsCreateReturnViewModel(*openapiclient.NewShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel(int32(123)), []openapiclient.ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel{*openapiclient.NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModel(int32(111222), int32(1))}, "ShipBob_Return_123") // ShipbobReturnsPublicApiViewModelsCreateReturnViewModel | Model defining the return (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnsApi.ReturnPost(context.Background()).ShipbobChannelId(shipbobChannelId).ShipbobReturnsPublicApiViewModelsCreateReturnViewModel(shipbobReturnsPublicApiViewModelsCreateReturnViewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ReturnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnPost`: ShipbobReturnsPublicApiViewModelsReturnOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ReturnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **shipbobReturnsPublicApiViewModelsCreateReturnViewModel** | [**ShipbobReturnsPublicApiViewModelsCreateReturnViewModel**](ShipbobReturnsPublicApiViewModelsCreateReturnViewModel.md) | Model defining the return | 

### Return type

[**ShipbobReturnsPublicApiViewModelsReturnOrderViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

