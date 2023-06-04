# \InventoryApi

All URIs are relative to *https://api.shipbob.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInventories**](InventoryApi.md#GetInventories) | **Get** /1.0/inventory | List inventory items
[**GetInventory**](InventoryApi.md#GetInventory) | **Get** /1.0/inventory/{inventoryId} | Get an inventory item
[**GetProductInventories**](InventoryApi.md#GetProductInventories) | **Get** /1.0/product/{productId}/inventory | Get a list of inventory items by product id



## GetInventories

> []Inventory GetInventories(ctx).Page(page).Limit(limit).IsActive(isActive).IsDigital(isDigital).IDs(iDs).Sort(sort).Search(search).LocationType(locationType).ShipbobChannelId(shipbobChannelId).Execute()

List inventory items

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
    page := int32(56) // int32 | Page of inventory items to get (optional)
    limit := int32(56) // int32 | Amount of inventory items per page to request (optional)
    isActive := true // bool | Whether the inventory should be active or not (optional)
    isDigital := true // bool | Whether the inventory is digital or not (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated inventory ids to filter by (optional)
    sort := "sort_example" // string | Sort will default to ascending order for each field.  To sort in descending order please pass a \"-\" in front of the field name.  For example, Sort=-onHand,name will sort by onHand descending (optional)
    search := "search_example" // string | Search is available for 2 fields, Inventory ID and Name - 1. Expected behavior for search by Inventory ID is exact match 2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word,  but must be consecutive characters. This is not case sensitive. (optional)
    locationType := "locationType_example" // string | LocationType is valid for hub, spoke, or lts. LocationType will default to all locations. (optional)
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetInventories(context.Background()).Page(page).Limit(limit).IsActive(isActive).IsDigital(isDigital).IDs(iDs).Sort(sort).Search(search).LocationType(locationType).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetInventories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventories`: []Inventory
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetInventories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of inventory items to get | 
 **limit** | **int32** | Amount of inventory items per page to request | 
 **isActive** | **bool** | Whether the inventory should be active or not | 
 **isDigital** | **bool** | Whether the inventory is digital or not | 
 **iDs** | **[]int32** | Comma separated inventory ids to filter by | 
 **sort** | **string** | Sort will default to ascending order for each field.  To sort in descending order please pass a \&quot;-\&quot; in front of the field name.  For example, Sort&#x3D;-onHand,name will sort by onHand descending | 
 **search** | **string** | Search is available for 2 fields, Inventory ID and Name - 1. Expected behavior for search by Inventory ID is exact match 2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word,  but must be consecutive characters. This is not case sensitive. | 
 **locationType** | **string** | LocationType is valid for hub, spoke, or lts. LocationType will default to all locations. | 
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventory

> Inventory GetInventory(ctx, inventoryId).Execute()

Get an inventory item

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
    inventoryId := int32(56) // int32 | The inventory id to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetInventory(context.Background(), inventoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventory`: Inventory
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryId** | **int32** | The inventory id to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Inventory**](Inventory.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductInventories

> []Inventory GetProductInventories(ctx, productId).ShipbobChannelId(shipbobChannelId).Execute()

Get a list of inventory items by product id

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
    productId := int32(56) // int32 | The product id to get inventory for
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetProductInventories(context.Background(), productId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetProductInventories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductInventories`: []Inventory
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetProductInventories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | The product id to get inventory for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductInventoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

