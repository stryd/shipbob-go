# \ProductsApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductBatchPost**](ProductsApi.md#ProductBatchPost) | **Post** /product/batch | Add multiple products to the store
[**ProductGet**](ProductsApi.md#ProductGet) | **Get** /product | Get multiple products
[**ProductPost**](ProductsApi.md#ProductPost) | **Post** /product | Add a single product to the store
[**ProductProductIdGet**](ProductsApi.md#ProductProductIdGet) | **Get** /product/{productId} | Get a single product
[**ProductProductIdPut**](ProductsApi.md#ProductProductIdPut) | **Put** /product/{productId} | Modify a single product



## ProductBatchPost

> []ProductsProductViewModel ProductBatchPost(ctx).ShipbobChannelId(shipbobChannelId).ProductsCreateProductModel(productsCreateProductModel).Execute()

Add multiple products to the store

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
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    productsCreateProductModel := []openapiclient.ProductsCreateProductModel{*openapiclient.NewProductsCreateProductModel("Medium Blue T-Shirt", "TShirtBlueM")} // []ProductsCreateProductModel | List of products to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProductBatchPost(context.Background()).ShipbobChannelId(shipbobChannelId).ProductsCreateProductModel(productsCreateProductModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductBatchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductBatchPost`: []ProductsProductViewModel
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **productsCreateProductModel** | [**[]ProductsCreateProductModel**](Products.CreateProductModel.md) | List of products to add | 

### Return type

[**[]ProductsProductViewModel**](Products.ProductViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGet

> []ProductsProductViewModel ProductGet(ctx).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).Search(search).ActiveStatus(activeStatus).BundleStatus(bundleStatus).ShipbobChannelId(shipbobChannelId).Execute()

Get multiple products

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
    page := int32(56) // int32 | Page of products to get (optional)
    limit := int32(56) // int32 | Amount of products per page to request (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated list of product ids to filter by (optional)
    referenceIds := []string{"Inner_example"} // []string | Comma separated list of reference ids to filter by (optional)
    search := "search_example" // string | Search is available for 2 fields of the inventory record related to the product: Inventory ID and Name -\\r\\n1. Expected behavior for search by Inventory ID is exact match\\r\\n2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word, \\r\\nbut must be consecutive characters. This is not case sensitive. (optional)
    activeStatus := openapiclient.Products.ProductActiveStatus("Any") // ProductsProductActiveStatus | Status filter for products:\\r\\n- Any: Include both active and inactive\\r\\n- Active: Filter products that are Active\\r\\n- Inactive: Filter products that are Inactive (optional)
    bundleStatus := openapiclient.Products.ProductBundleStatus("Any") // ProductsProductBundleStatus | Bundle filter for products:\\r\\n- Any: Don't filter and consider products that are bundles or not bundles\\r\\n- Bundle: Filter by products that are bundles\\r\\n- NotBundle: Filter by products that are not bundles (optional)
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProductGet(context.Background()).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).Search(search).ActiveStatus(activeStatus).BundleStatus(bundleStatus).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGet`: []ProductsProductViewModel
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of products to get | 
 **limit** | **int32** | Amount of products per page to request | 
 **iDs** | **[]int32** | Comma separated list of product ids to filter by | 
 **referenceIds** | **[]string** | Comma separated list of reference ids to filter by | 
 **search** | **string** | Search is available for 2 fields of the inventory record related to the product: Inventory ID and Name -\\r\\n1. Expected behavior for search by Inventory ID is exact match\\r\\n2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word, \\r\\nbut must be consecutive characters. This is not case sensitive. | 
 **activeStatus** | [**ProductsProductActiveStatus**](ProductsProductActiveStatus.md) | Status filter for products:\\r\\n- Any: Include both active and inactive\\r\\n- Active: Filter products that are Active\\r\\n- Inactive: Filter products that are Inactive | 
 **bundleStatus** | [**ProductsProductBundleStatus**](ProductsProductBundleStatus.md) | Bundle filter for products:\\r\\n- Any: Don&#39;t filter and consider products that are bundles or not bundles\\r\\n- Bundle: Filter by products that are bundles\\r\\n- NotBundle: Filter by products that are not bundles | 
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]ProductsProductViewModel**](Products.ProductViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductPost

> []ProductsProductViewModel ProductPost(ctx).ShipbobChannelId(shipbobChannelId).ProductsCreateProductModel(productsCreateProductModel).Execute()

Add a single product to the store

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
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    productsCreateProductModel := *openapiclient.NewProductsCreateProductModel("Medium Blue T-Shirt", "TShirtBlueM") // ProductsCreateProductModel | The product to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProductPost(context.Background()).ShipbobChannelId(shipbobChannelId).ProductsCreateProductModel(productsCreateProductModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductPost`: []ProductsProductViewModel
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **productsCreateProductModel** | [**ProductsCreateProductModel**](ProductsCreateProductModel.md) | The product to add | 

### Return type

[**[]ProductsProductViewModel**](Products.ProductViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductProductIdGet

> ProductsProductViewModel ProductProductIdGet(ctx, productId).ShipbobChannelId(shipbobChannelId).Execute()

Get a single product

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
    productId := int32(56) // int32 | Unique identifier of the product
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProductProductIdGet(context.Background(), productId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductProductIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductProductIdGet`: ProductsProductViewModel
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductProductIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | Unique identifier of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductProductIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**ProductsProductViewModel**](Products.ProductViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductProductIdPut

> []ProductsProductViewModel ProductProductIdPut(ctx, productId).ShipbobChannelId(shipbobChannelId).ProductsUpdateProductModel(productsUpdateProductModel).Execute()

Modify a single product

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
    productId := int32(56) // int32 | Unique identifier of the product to modify
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    productsUpdateProductModel := *openapiclient.NewProductsUpdateProductModel("Medium Blue T-Shirt") // ProductsUpdateProductModel | Updated fields to the product (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProductProductIdPut(context.Background(), productId).ShipbobChannelId(shipbobChannelId).ProductsUpdateProductModel(productsUpdateProductModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductProductIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductProductIdPut`: []ProductsProductViewModel
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductProductIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | Unique identifier of the product to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductProductIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **productsUpdateProductModel** | [**ProductsUpdateProductModel**](ProductsUpdateProductModel.md) | Updated fields to the product | 

### Return type

[**[]ProductsProductViewModel**](Products.ProductViewModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/_*+json, application/json, application/json-patch+json, text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

