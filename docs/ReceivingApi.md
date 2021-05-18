# \ReceivingApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FulfillmentCenterGet**](ReceivingApi.md#FulfillmentCenterGet) | **Get** /fulfillmentCenter | Get Fulfillment Centers
[**ReceivingIdCancelPost**](ReceivingApi.md#ReceivingIdCancelPost) | **Post** /receiving/{id}/cancel | Cancel Warehouse Receiving Order
[**ReceivingIdGet**](ReceivingApi.md#ReceivingIdGet) | **Get** /receiving/{id} | Get Warehouse Receiving Order
[**ReceivingIdLabelsGet**](ReceivingApi.md#ReceivingIdLabelsGet) | **Get** /receiving/{id}/labels | Get Warehouse Receiving Order Box Labels
[**ReceivingPost**](ReceivingApi.md#ReceivingPost) | **Post** /receiving | Create Warehouse Receiving Order



## FulfillmentCenterGet

> []ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel FulfillmentCenterGet(ctx).Execute()

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
    resp, r, err := api_client.ReceivingApi.FulfillmentCenterGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.FulfillmentCenterGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FulfillmentCenterGet`: []ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.FulfillmentCenterGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFulfillmentCenterGetRequest struct via the builder pattern


### Return type

[**[]ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel**](Shipbob.Receiving.Public.Api.Models.FulfillmentCenterViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceivingIdCancelPost

> ReceivingIdCancelPost(ctx, id).Execute()

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
    resp, r, err := api_client.ReceivingApi.ReceivingIdCancelPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.ReceivingIdCancelPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReceivingIdCancelPostRequest struct via the builder pattern


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


## ReceivingIdGet

> ShipbobReceivingPublicApiModelsReceivingOrderViewModel ReceivingIdGet(ctx, id).Execute()

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
    resp, r, err := api_client.ReceivingApi.ReceivingIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.ReceivingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReceivingIdGet`: ShipbobReceivingPublicApiModelsReceivingOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.ReceivingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShipbobReceivingPublicApiModelsReceivingOrderViewModel**](Shipbob.Receiving.Public.Api.Models.ReceivingOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceivingIdLabelsGet

> string ReceivingIdLabelsGet(ctx, id).Execute()

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
    resp, r, err := api_client.ReceivingApi.ReceivingIdLabelsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.ReceivingIdLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReceivingIdLabelsGet`: string
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.ReceivingIdLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Id of the receiving order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingIdLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceivingPost

> ShipbobReceivingPublicApiModelsReceivingOrderViewModel ReceivingPost(ctx).ShipbobReceivingPublicApiModelsCreateReceivingOrderModel(shipbobReceivingPublicApiModelsCreateReceivingOrderModel).Execute()

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
    shipbobReceivingPublicApiModelsCreateReceivingOrderModel := *openapiclient.NewShipbobReceivingPublicApiModelsCreateReceivingOrderModel(openapiclient.Shipbob.Receiving.Public.Common.Models.PackingType("EverythingInOneBox"), []openapiclient.ShipbobReceivingPublicApiModelsAddBoxToOrderModel{*openapiclient.NewShipbobReceivingPublicApiModelsAddBoxToOrderModel([]openapiclient.ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel{*openapiclient.NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModel(int32(123), int32(123))}, "860C8CDC8F0B4FC7AB69AC86C20539EC")}, time.Now(), *openapiclient.NewShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel(int32(123)), openapiclient.Shipbob.Receiving.Public.Common.Models.PackageType("Package")) // ShipbobReceivingPublicApiModelsCreateReceivingOrderModel | The receiving order to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReceivingApi.ReceivingPost(context.Background()).ShipbobReceivingPublicApiModelsCreateReceivingOrderModel(shipbobReceivingPublicApiModelsCreateReceivingOrderModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReceivingApi.ReceivingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReceivingPost`: ShipbobReceivingPublicApiModelsReceivingOrderViewModel
    fmt.Fprintf(os.Stdout, "Response from `ReceivingApi.ReceivingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceivingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobReceivingPublicApiModelsCreateReceivingOrderModel** | [**ShipbobReceivingPublicApiModelsCreateReceivingOrderModel**](ShipbobReceivingPublicApiModelsCreateReceivingOrderModel.md) | The receiving order to create | 

### Return type

[**ShipbobReceivingPublicApiModelsReceivingOrderViewModel**](Shipbob.Receiving.Public.Api.Models.ReceivingOrderViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

