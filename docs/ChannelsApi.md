# \ChannelsApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelGet**](ChannelsApi.md#ChannelGet) | **Get** /channel | Get user-authorized channel info



## ChannelGet

> []ShipBobChannelsApiViewModelsChannelViewModel ChannelGet(ctx).Execute()

Get user-authorized channel info

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
    resp, r, err := api_client.ChannelsApi.ChannelGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ChannelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChannelGet`: []ShipBobChannelsApiViewModelsChannelViewModel
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ChannelGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChannelGetRequest struct via the builder pattern


### Return type

[**[]ShipBobChannelsApiViewModelsChannelViewModel**](ShipBob.Channels.Api.ViewModels.ChannelViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

