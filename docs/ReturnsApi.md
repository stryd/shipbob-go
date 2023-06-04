# \ReturnsApi

All URIs are relative to *https://api.shipbob.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReturn**](ReturnsApi.md#CancelReturn) | **Post** /1.0/return/{id}/cancel | Cancel Return Order
[**CreateReturn**](ReturnsApi.md#CreateReturn) | **Post** /1.0/return | Create Return Order
[**GetReturn**](ReturnsApi.md#GetReturn) | **Get** /1.0/return/{id} | Get Return Order
[**GetReturnStatushistory**](ReturnsApi.md#GetReturnStatushistory) | **Get** /1.0/return/{id}/statushistory | Get One Return&#39;s status history
[**GetReturns**](ReturnsApi.md#GetReturns) | **Get** /1.0/return | Get Return Orders
[**UpdateReturn**](ReturnsApi.md#UpdateReturn) | **Put** /1.0/return/{id} | Modify Return Order



## CancelReturn

> ReturnOrder CancelReturn(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Cancel Return Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.CancelReturn(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CancelReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelReturn`: ReturnOrder
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CancelReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ReturnOrder**](ReturnOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReturn

> ReturnOrder CreateReturn(ctx).ShipbobChannelId(shipbobChannelId).CreateReturn(createReturn).Execute()

Create Return Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    createReturn := *openapiclient.NewCreateReturn(*openapiclient.NewReturnFulfillmentCenter(int32(123)), []openapiclient.ReturnInventory{*openapiclient.NewReturnInventory(int32(111222), int32(1))}, "ShipBob_Return_123") // CreateReturn | Model defining the return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.CreateReturn(context.Background()).ShipbobChannelId(shipbobChannelId).CreateReturn(createReturn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CreateReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReturn`: ReturnOrder
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CreateReturn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **createReturn** | [**CreateReturn**](CreateReturn.md) | Model defining the return | 

### Return type

[**ReturnOrder**](ReturnOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturn

> ReturnOrder GetReturn(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Get Return Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.GetReturn(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.GetReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturn`: ReturnOrder
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.GetReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ReturnOrder**](ReturnOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnStatushistory

> []ReturnOrderStatusHistory GetReturnStatushistory(ctx, id).ShipbobChannelId(shipbobChannelId).Execute()

Get One Return's status history

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    id := int32(56) // int32 | Id of the return order
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.GetReturnStatushistory(context.Background(), id).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.GetReturnStatushistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnStatushistory`: []ReturnOrderStatusHistory
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.GetReturnStatushistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnStatushistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ReturnOrderStatusHistory**](ReturnOrderStatusHistory.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturns

> []ReturnOrder GetReturns(ctx).Page(page).Limit(limit).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).IDs(iDs).ReferenceIds(referenceIds).Status(status).FulfillmentCenterIds(fulfillmentCenterIds).TrackingNumbers(trackingNumbers).OriginalShipmentIds(originalShipmentIds).InventoryIds(inventoryIds).ShipbobChannelId(shipbobChannelId).Execute()

Get Return Orders

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    page := int32(56) // int32 | Page of return orders to get (optional)
    limit := int32(56) // int32 | Amount of return orders per page to request (optional)
    sortOrder := openapiclient.SortOrder("Newest") // SortOrder | Order to sort results by (optional)
    startDate := time.Now() // time.Time | Start date to filter orders inserted later than (optional)
    endDate := time.Now() // time.Time | End date to filter orders inserted earlier than (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated list of return orders ids to filter by (optional)
    referenceIds := []string{"Inner_example"} // []string | Comma separated list of reference ids to filter by (optional)
    status := []openapiclient.ReturnStatus{openapiclient.ReturnStatus("AwaitingArrival")} // []ReturnStatus | Comma separated list of statuses to filter by (optional)
    fulfillmentCenterIds := []int32{int32(123)} // []int32 | Comma separated list of destination fulfillment center IDs to filter by (optional)
    trackingNumbers := []string{"Inner_example"} // []string | Comma separated list of tracking numbers to filter by (optional)
    originalShipmentIds := []int32{int32(123)} // []int32 | Comma separated list of original shipment IDs to filter by (optional)
    inventoryIds := []int32{int32(123)} // []int32 | Comma separated list of inventory IDs contained in return to filter by (optional)
    shipbobChannelId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.GetReturns(context.Background()).Page(page).Limit(limit).SortOrder(sortOrder).StartDate(startDate).EndDate(endDate).IDs(iDs).ReferenceIds(referenceIds).Status(status).FulfillmentCenterIds(fulfillmentCenterIds).TrackingNumbers(trackingNumbers).OriginalShipmentIds(originalShipmentIds).InventoryIds(inventoryIds).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.GetReturns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturns`: []ReturnOrder
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.GetReturns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of return orders to get | 
 **limit** | **int32** | Amount of return orders per page to request | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Order to sort results by | 
 **startDate** | **time.Time** | Start date to filter orders inserted later than | 
 **endDate** | **time.Time** | End date to filter orders inserted earlier than | 
 **iDs** | **[]int32** | Comma separated list of return orders ids to filter by | 
 **referenceIds** | **[]string** | Comma separated list of reference ids to filter by | 
 **status** | [**[]ReturnStatus**](ReturnStatus.md) | Comma separated list of statuses to filter by | 
 **fulfillmentCenterIds** | **[]int32** | Comma separated list of destination fulfillment center IDs to filter by | 
 **trackingNumbers** | **[]string** | Comma separated list of tracking numbers to filter by | 
 **originalShipmentIds** | **[]int32** | Comma separated list of original shipment IDs to filter by | 
 **inventoryIds** | **[]int32** | Comma separated list of inventory IDs contained in return to filter by | 
 **shipbobChannelId** | **int32** |  | 

### Return type

[**[]ReturnOrder**](ReturnOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReturn

> ReturnOrder UpdateReturn(ctx, id).ShipbobChannelId(shipbobChannelId).CreateReturn(createReturn).Execute()

Modify Return Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stryd/shipbob-go"
)

func main() {
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    id := int32(56) // int32 | Id of the return order
    createReturn := *openapiclient.NewCreateReturn(*openapiclient.NewReturnFulfillmentCenter(int32(123)), []openapiclient.ReturnInventory{*openapiclient.NewReturnInventory(int32(111222), int32(1))}, "ShipBob_Return_123") // CreateReturn | Model defining the return (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.UpdateReturn(context.Background(), id).ShipbobChannelId(shipbobChannelId).CreateReturn(createReturn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.UpdateReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReturn`: ReturnOrder
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.UpdateReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the return order | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

 **createReturn** | [**CreateReturn**](CreateReturn.md) | Model defining the return | 

### Return type

[**ReturnOrder**](ReturnOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

