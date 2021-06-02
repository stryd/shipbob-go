# \ReceivingApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReceivingOrder**](ReceivingApi.md#CancelReceivingOrder) | **Post** /receiving/{id}/cancel | Cancel Warehouse Receiving Order
[**CreateReceivingOrder**](ReceivingApi.md#CreateReceivingOrder) | **Post** /receiving | Create Warehouse Receiving Order
[**GetFulfillmentCenters**](ReceivingApi.md#GetFulfillmentCenters) | **Get** /fulfillmentCenter | Get Fulfillment Centers
[**GetReceivingOrder**](ReceivingApi.md#GetReceivingOrder) | **Get** /receiving/{id} | Get Warehouse Receiving Order
[**GetReceivingOrderLabels**](ReceivingApi.md#GetReceivingOrderLabels) | **Get** /receiving/{id}/labels | Get Warehouse Receiving Order Box Labels



## CancelReceivingOrder

> CancelReceivingOrder(ctx, id).Execute()

Cancel Warehouse Receiving Order

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
    id := int32(56) // int32 | Id of the receiving order to cancel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.CancelReceivingOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CancelReceivingOrder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelReceivingOrderRequest struct via the builder pattern


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


## CreateReceivingOrder

> ReceivingOrder CreateReceivingOrder(ctx).CreateReceivingOrder(createReceivingOrder).Execute()

Create Warehouse Receiving Order

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
    createReceivingOrder := *openapiclient.NewCreateReceivingOrder(openapiclient.PackingType("EverythingInOneBox"), []openapiclient.CreateReceivingOrderBoxes{*openapiclient.NewCreateReceivingOrderBoxes([]openapiclient.CreateReceivingOrderBoxItems{*openapiclient.NewCreateReceivingOrderBoxItems(int32(123), int32(123))}, "860C8CDC8F0B4FC7AB69AC86C20539EC")}, time.Now(), *openapiclient.NewCreateReceivingOrderFulfillmentCenter(int32(123)), openapiclient.PackageType("Package")) // CreateReceivingOrder | The receiving order to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.CreateReceivingOrder(context.Background()).CreateReceivingOrder(createReceivingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.CreateReceivingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceivingOrder`: ReceivingOrder
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.CreateReceivingOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceivingOrderRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.GetFulfillmentCenters(context.Background()).Execute()
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


## GetReceivingOrder

> ReceivingOrder GetReceivingOrder(ctx, id).Execute()

Get Warehouse Receiving Order

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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.GetReceivingOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingOrder`: ReceivingOrder
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingOrderRequest struct via the builder pattern


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


## GetReceivingOrderLabels

> string GetReceivingOrderLabels(ctx, id).Execute()

Get Warehouse Receiving Order Box Labels

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
    id := int32(56) // int32 | Id of the receiving order

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.GetReceivingOrderLabels(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.GetReceivingOrderLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivingOrderLabels`: string
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.GetReceivingOrderLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivingOrderLabelsRequest struct via the builder pattern


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

