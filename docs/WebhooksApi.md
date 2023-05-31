# \WebhooksApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /webhook | Create a new webhook subscription
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhook/{id} | Delete an existing webhook subscription
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhook | Get Webhooks



## CreateWebhook

> Webhook CreateWebhook(ctx).ShipbobChannelId(shipbobChannelId).WebhookSubscription(webhookSubscription).Execute()

Create a new webhook subscription

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
    shipbobChannelId := int32(56) // int32 |  (optional)
    webhookSubscription := *openapiclient.NewWebhookSubscription("https://mywebsite.com/shipbob/handler", openapiclient.Topics("order_shipped")) // WebhookSubscription |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.CreateWebhook(context.Background()).ShipbobChannelId(shipbobChannelId).WebhookSubscription(webhookSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** |  | 
 **webhookSubscription** | [**WebhookSubscription**](WebhookSubscription.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, id).Execute()

Delete an existing webhook subscription

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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.DeleteWebhook(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetWebhooks

> []Webhook GetWebhooks(ctx).Topic(topic).Page(page).Limit(limit).Execute()

Get Webhooks



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
    topic := openapiclient.Topics("order_shipped") // Topics | Topic of the webhooks requested (optional)
    page := int32(56) // int32 | Page of Webhooks to get (optional)
    limit := int32(56) // int32 | Amount of Webhooks per page to request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhooks(context.Background()).Topic(topic).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: []Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | [**Topics**](Topics.md) | Topic of the webhooks requested | 
 **page** | **int32** | Page of Webhooks to get | 
 **limit** | **int32** | Amount of Webhooks per page to request | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

