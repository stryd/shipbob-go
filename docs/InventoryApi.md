# \InventoryApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryGet**](InventoryApi.md#InventoryGet) | **Get** /inventory | List inventory items
[**InventoryInventoryIdGet**](InventoryApi.md#InventoryInventoryIdGet) | **Get** /inventory/{inventoryId} | Get an inventory item
[**ProductProductIdInventoryGet**](InventoryApi.md#ProductProductIdInventoryGet) | **Get** /product/{productId}/inventory | Get a list of inventory items by product id



## InventoryGet

> []ShipbobInventoryApiViewModelsInventoryViewModel InventoryGet(ctx).Page(page).Limit(limit).IsActive(isActive).IsDigital(isDigital).IDs(iDs).Sort(sort).Search(search).ShipbobChannelId(shipbobChannelId).Execute()

List inventory items

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
    page := int32(56) // int32 | Page of inventory items to get (optional)
    limit := int32(56) // int32 | Amount of inventory items per page to request (optional)
    isActive := true // bool | Whether the inventory should be active or not (optional)
    isDigital := true // bool | Whether the inventory is digital or not (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated inventory ids to filter by (optional)
    sort := "sort_example" // string | Sort will default to ascending order for each field. \\r\\nTo sort in descending order please pass a \"-\" in front of the field name. \\r\\nFor example, Sort=-onHand,name will sort by onHand descending (optional)
    search := "search_example" // string | Search is available for 2 fields, Inventory ID and Name -\\r\\n1. Expected behavior for search by Inventory ID is exact match\\r\\n2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word, \\r\\nbut must be consecutive characters. This is not case sensitive. (optional)
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.InventoryGet(context.Background()).Page(page).Limit(limit).IsActive(isActive).IsDigital(isDigital).IDs(iDs).Sort(sort).Search(search).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryGet`: []ShipbobInventoryApiViewModelsInventoryViewModel
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of inventory items to get | 
 **limit** | **int32** | Amount of inventory items per page to request | 
 **isActive** | **bool** | Whether the inventory should be active or not | 
 **isDigital** | **bool** | Whether the inventory is digital or not | 
 **iDs** | **[]int32** | Comma separated inventory ids to filter by | 
 **sort** | **string** | Sort will default to ascending order for each field. \\r\\nTo sort in descending order please pass a \&quot;-\&quot; in front of the field name. \\r\\nFor example, Sort&#x3D;-onHand,name will sort by onHand descending | 
 **search** | **string** | Search is available for 2 fields, Inventory ID and Name -\\r\\n1. Expected behavior for search by Inventory ID is exact match\\r\\n2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word, \\r\\nbut must be consecutive characters. This is not case sensitive. | 
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipbobInventoryApiViewModelsInventoryViewModel**](Shipbob.Inventory.Api.ViewModels.InventoryViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryInventoryIdGet

> ShipbobInventoryApiViewModelsInventoryViewModel InventoryInventoryIdGet(ctx, inventoryId).Execute()

Get an inventory item

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
    inventoryId := int32(56) // int32 | The inventory id to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.InventoryInventoryIdGet(context.Background(), inventoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryInventoryIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryInventoryIdGet`: ShipbobInventoryApiViewModelsInventoryViewModel
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryInventoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryId** | **int32** | The inventory id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiInventoryInventoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShipbobInventoryApiViewModelsInventoryViewModel**](Shipbob.Inventory.Api.ViewModels.InventoryViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductProductIdInventoryGet

> []ShipbobInventoryApiViewModelsInventoryViewModel ProductProductIdInventoryGet(ctx, productId).ShipbobChannelId(shipbobChannelId).Execute()

Get a list of inventory items by product id

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
    productId := int32(56) // int32 | The product id to get inventory for
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ProductProductIdInventoryGet(context.Background(), productId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ProductProductIdInventoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductProductIdInventoryGet`: []ShipbobInventoryApiViewModelsInventoryViewModel
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ProductProductIdInventoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | The product id to get inventory for | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductProductIdInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ShipbobInventoryApiViewModelsInventoryViewModel**](Shipbob.Inventory.Api.ViewModels.InventoryViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

