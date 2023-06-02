# \ReceivingApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReceiving**](ReceivingApi.md#CancelReceiving) | **Post** /receiving/{id}/cancel | Cancel Warehouse Receiving Order
[**CreateReceiving**](ReceivingApi.md#CreateReceiving) | **Post** /receiving | Create Warehouse Receiving Order
[**GetFulfillmentCenters**](ReceivingApi.md#GetFulfillmentCenters) | **Get** /fulfillmentCenter | Get Fulfillment Centers
[**GetReceiving**](ReceivingApi.md#GetReceiving) | **Get** /receiving/{id} | Get Warehouse Receiving Order
[**GetReceivingLabels**](ReceivingApi.md#GetReceivingLabels) | **Get** /receiving/{id}/labels | Get Warehouse Receiving Order Box Labels



## CancelReceiving

> CancelReceiving(ctx, id).Execute()

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


## CreateReceiving

> ReceivingOrder CreateReceiving(ctx).CreateReceivingOrder(createReceivingOrder).Execute()

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

