# \ReceivingApi

All URIs are relative to *https://api.shipbob.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReceiving**](ReceivingApi.md#CancelReceiving) | **Post** /1.0/receiving/{id}/cancel | Cancel Warehouse Receiving Order (DEPRECATED)
[**CancelReceivingV2**](ReceivingApi.md#CancelReceivingV2) | **Post** /2.0/receiving/{id}/cancel | Cancel Warehouse Receiving Order
[**CreateReceiving**](ReceivingApi.md#CreateReceiving) | **Post** /1.0/receiving | Create Warehouse Receiving Order (DEPRECATED)
[**CreateReceivingV2**](ReceivingApi.md#CreateReceivingV2) | **Post** /2.0/receiving | Create Warehouse Receiving Order
[**GetFulfillmentCenters**](ReceivingApi.md#GetFulfillmentCenters) | **Get** /1.0/fulfillmentCenter | Get Fulfillment Centers
[**GetReceiving**](ReceivingApi.md#GetReceiving) | **Get** /1.0/receiving/{id} | Get Warehouse Receiving Order (DEPRECATED)
[**GetReceivingLabels**](ReceivingApi.md#GetReceivingLabels) | **Get** /1.0/receiving/{id}/labels | Get Warehouse Receiving Order Box Labels (DEPRECATED)
[**GetReceivingLabelsV2**](ReceivingApi.md#GetReceivingLabelsV2) | **Get** /2.0/receiving/{id}/labels | Get Warehouse Receiving Order Box Labels
[**GetReceivingV2**](ReceivingApi.md#GetReceivingV2) | **Get** /2.0/receiving/{id} | Get Warehouse Receiving Order
[**GetReceivings**](ReceivingApi.md#GetReceivings) | **Get** /1.0/receiving | Get a Warehouse Receiving Order by Purchase Order Number (DEPRECATED)
[**GetReceivingsV2**](ReceivingApi.md#GetReceivingsV2) | **Get** /2.0/receiving | Get Multiple Warehouse Receiving Orders



## CancelReceiving

> CancelReceiving(ctx, id).Execute()

Cancel Warehouse Receiving Order (DEPRECATED)



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
    id := int32(56) // int32 | Id of the receiving order to cancel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReceivingApi.CancelReceiving(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CancelReceiving``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReceivingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelReceivingV2

> WarehouseReceivingOrderV2 CancelReceivingV2(ctx, id).Execute()

Cancel Warehouse Receiving Order

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
    id := int32(56) // int32 | Id of the receiving order to cancel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.CancelReceivingV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CancelReceivingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelReceivingV2`: WarehouseReceivingOrderV2
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.CancelReceivingV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReceivingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WarehouseReceivingOrderV2**](WarehouseReceivingOrderV2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceiving

> ReceivingOrder CreateReceiving(ctx).CreateReceivingOrder(createReceivingOrder).Execute()

Create Warehouse Receiving Order (DEPRECATED)



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
    createReceivingOrder := *openapiclient.NewCreateReceivingOrder(openapiclient.PackingType("EverythingInOneBox"), []openapiclient.AddBoxToOrder{*openapiclient.NewAddBoxToOrder([]openapiclient.AddBoxItemToBox{*openapiclient.NewAddBoxItemToBox(int32(123), int32(123))}, "860C8CDC8F0B4FC7AB69AC86C20539EC")}, time.Now(), *openapiclient.NewAssignOrderToFulfillmentCenter(int32(123)), openapiclient.PackageType("Package")) // CreateReceivingOrder | The receiving order to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.CreateReceiving(context.Background()).CreateReceivingOrder(createReceivingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CreateReceiving``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceiving`: ReceivingOrder
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.CreateReceiving`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceivingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReceivingOrder** | [**CreateReceivingOrder**](CreateReceivingOrder.md) | The receiving order to create | 

### Return type

[**ReceivingOrder**](ReceivingOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceivingV2

> WarehouseReceivingOrderV2 CreateReceivingV2(ctx).CreateReceivingOrder(createReceivingOrder).Execute()

Create Warehouse Receiving Order

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
    createReceivingOrder := *openapiclient.NewCreateReceivingOrder(openapiclient.PackingType("EverythingInOneBox"), []openapiclient.AddBoxToOrder{*openapiclient.NewAddBoxToOrder([]openapiclient.AddBoxItemToBox{*openapiclient.NewAddBoxItemToBox(int32(123), int32(123))}, "860C8CDC8F0B4FC7AB69AC86C20539EC")}, time.Now(), *openapiclient.NewAssignOrderToFulfillmentCenter(int32(123)), openapiclient.PackageType("Package")) // CreateReceivingOrder | The receiving order to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.CreateReceivingV2(context.Background()).CreateReceivingOrder(createReceivingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CreateReceivingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceivingV2`: WarehouseReceivingOrderV2
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.CreateReceivingV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceivingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReceivingOrder** | [**CreateReceivingOrder**](CreateReceivingOrder.md) | The receiving order to create | 

### Return type

[**WarehouseReceivingOrderV2**](WarehouseReceivingOrderV2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFulfillmentCenters

> []ReceivingFulfillmentCenter GetFulfillmentCenters(ctx).Execute()

Get Fulfillment Centers

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetFulfillmentCenters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetFulfillmentCenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFulfillmentCenters`: []ReceivingFulfillmentCenter
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetFulfillmentCenters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentCentersRequest struct via the builder pattern


### Return type

[**[]ReceivingFulfillmentCenter**](ReceivingFulfillmentCenter.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiving

> ReceivingOrder GetReceiving(ctx, id).Execute()

Get Warehouse Receiving Order (DEPRECATED)



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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceiving(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceiving``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceiving`: ReceivingOrder
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceiving`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReceivingOrder**](ReceivingOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivingLabels

> string GetReceivingLabels(ctx, id).Execute()

Get Warehouse Receiving Order Box Labels (DEPRECATED)



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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceivingLabels(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingLabels`: string
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivingLabelsV2

> string GetReceivingLabelsV2(ctx, id).Execute()

Get Warehouse Receiving Order Box Labels

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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceivingLabelsV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingLabelsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingLabelsV2`: string
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingLabelsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingLabelsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivingV2

> WarehouseReceivingOrderV2 GetReceivingV2(ctx, id).Execute()

Get Warehouse Receiving Order

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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceivingV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingV2`: WarehouseReceivingOrderV2
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WarehouseReceivingOrderV2**](WarehouseReceivingOrderV2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivings

> ReceivingOrder GetReceivings(ctx).PurchaseOrderNumber(purchaseOrderNumber).Execute()

Get a Warehouse Receiving Order by Purchase Order Number (DEPRECATED)



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
    purchaseOrderNumber := "purchaseOrderNumber_example" // string | Purchase order number of the receiving order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceivings(context.Background()).PurchaseOrderNumber(purchaseOrderNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivings`: ReceivingOrder
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseOrderNumber** | **string** | Purchase order number of the receiving order | 

### Return type

[**ReceivingOrder**](ReceivingOrder.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivingsV2

> []WarehouseReceivingOrderV2 GetReceivingsV2(ctx).Page(page).Limit(limit).IDs(iDs).Statuses(statuses).InsertStartDate(insertStartDate).InsertEndDate(insertEndDate).FulfillmentCenterIds(fulfillmentCenterIds).PurchaseOrderNumbers(purchaseOrderNumbers).Execute()

Get Multiple Warehouse Receiving Orders

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
    page := int32(56) // int32 | Page of WROs to get (optional)
    limit := int32(56) // int32 | Number of WROs per page to request (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated list of WRO IDs to filter by (optional)
    statuses := []openapiclient.ReceivingStatus{openapiclient.ReceivingStatus("Awaiting")} // []ReceivingStatus | Comma separated list of WRO statuses to filter by (optional)
    insertStartDate := time.Now() // time.Time | Earliest date that a WRO was created (optional)
    insertEndDate := time.Now() // time.Time | Latest date that a WRO was created (optional)
    fulfillmentCenterIds := []int32{int32(123)} // []int32 | Comma separated list of WRO fulfillment center IDs to filter by (optional)
    purchaseOrderNumbers := []string{"Inner_example"} // []string | Comma separated list of WRO PO numbers to filter by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReceivingApi.GetReceivingsV2(context.Background()).Page(page).Limit(limit).IDs(iDs).Statuses(statuses).InsertStartDate(insertStartDate).InsertEndDate(insertEndDate).FulfillmentCenterIds(fulfillmentCenterIds).PurchaseOrderNumbers(purchaseOrderNumbers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingsV2`: []WarehouseReceivingOrderV2
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of WROs to get | 
 **limit** | **int32** | Number of WROs per page to request | 
 **iDs** | **[]int32** | Comma separated list of WRO IDs to filter by | 
 **statuses** | [**[]ReceivingStatus**](ReceivingStatus.md) | Comma separated list of WRO statuses to filter by | 
 **insertStartDate** | **time.Time** | Earliest date that a WRO was created | 
 **insertEndDate** | **time.Time** | Latest date that a WRO was created | 
 **fulfillmentCenterIds** | **[]int32** | Comma separated list of WRO fulfillment center IDs to filter by | 
 **purchaseOrderNumbers** | **[]string** | Comma separated list of WRO PO numbers to filter by | 

### Return type

[**[]WarehouseReceivingOrderV2**](WarehouseReceivingOrderV2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

