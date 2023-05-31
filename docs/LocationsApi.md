# \LocationsApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocations**](LocationsApi.md#GetLocations) | **Get** /location | Get locations



## GetLocations

> []Location GetLocations(ctx).IncludeInactive(includeInactive).ReceivingEnabled(receivingEnabled).AccessGranted(accessGranted).Execute()

Get locations

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
    includeInactive := true // bool | Whether the inactive locations should be included or not (optional)
    receivingEnabled := true // bool | Return all the receiving enabled locations (optional)
    accessGranted := true // bool | Return all the access granted locations (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.GetLocations(context.Background()).IncludeInactive(includeInactive).ReceivingEnabled(receivingEnabled).AccessGranted(accessGranted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.GetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocations`: []Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeInactive** | **bool** | Whether the inactive locations should be included or not | 
 **receivingEnabled** | **bool** | Return all the receiving enabled locations | 
 **accessGranted** | **bool** | Return all the access granted locations | 

### Return type

[**[]Location**](Location.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

