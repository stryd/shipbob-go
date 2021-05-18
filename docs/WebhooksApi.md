# \WebhooksApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookGet**](WebhooksApi.md#WebhookGet) | **Get** /webhook | Get Webhooks
[**WebhookIdDelete**](WebhooksApi.md#WebhookIdDelete) | **Delete** /webhook/{id} | Delete an existing webhook subscription
[**WebhookPost**](WebhooksApi.md#WebhookPost) | **Post** /webhook | Create a new webhook subscription



## WebhookGet

> []WebhooksWebhookViewModel WebhookGet(ctx).Topic(topic).Page(page).Limit(limit).Execute()

Get Webhooks



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
    topic := openapiclient.Webhooks.Topics("order_shipped") // WebhooksTopics | Topic of the webhooks requested (optional)
    page := int32(56) // int32 | Page of Webhooks to get (optional)
    limit := int32(56) // int32 | Amount of Webhooks per page to request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WebhookGet(context.Background()).Topic(topic).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookGet`: []WebhooksWebhookViewModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | [**WebhooksTopics**](WebhooksTopics.md) | Topic of the webhooks requested | 
 **page** | **int32** | Page of Webhooks to get | 
 **limit** | **int32** | Amount of Webhooks per page to request | 

### Return type

[**[]WebhooksWebhookViewModel**](Webhooks.WebhookViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookIdDelete

> WebhookIdDelete(ctx, id).Execute()

Delete an existing webhook subscription

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WebhookIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookIdDeleteRequest struct via the builder pattern


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


## WebhookPost

> WebhooksWebhookViewModel WebhookPost(ctx).ShipbobChannelId(shipbobChannelId).WebhooksCreateWebhookSubscriptionModel(webhooksCreateWebhookSubscriptionModel).Execute()

Create a new webhook subscription

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
    shipbobChannelId := int32(56) // int32 |  (optional)
    webhooksCreateWebhookSubscriptionModel := *openapiclient.NewWebhooksCreateWebhookSubscriptionModel("https://mywebsite.com/shipbob/handler", openapiclient.Webhooks.Topics("order_shipped")) // WebhooksCreateWebhookSubscriptionModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WebhookPost(context.Background()).ShipbobChannelId(shipbobChannelId).WebhooksCreateWebhookSubscriptionModel(webhooksCreateWebhookSubscriptionModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookPost`: WebhooksWebhookViewModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** |  | 
 **webhooksCreateWebhookSubscriptionModel** | [**WebhooksCreateWebhookSubscriptionModel**](WebhooksCreateWebhookSubscriptionModel.md) |  | 

### Return type

[**WebhooksWebhookViewModel**](Webhooks.WebhookViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

