# \ChannelsApi

All URIs are relative to *https://api.shipbob.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannels**](ChannelsApi.md#GetChannels) | **Get** /1.0/channel | Get user-authorized channel info



## GetChannels

> []Channel GetChannels(ctx).Execute()

Get user-authorized channel info

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
    resp, r, err := apiClient.ChannelsApi.GetChannels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.GetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannels`: []Channel
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.GetChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


### Return type

[**[]Channel**](Channel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

