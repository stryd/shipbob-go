# \ProductsApi

All URIs are relative to *https://api.shipbob.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductsApi.md#CreateProduct) | **Post** /product | Add a single product to the store
[**CreateProductBatch**](ProductsApi.md#CreateProductBatch) | **Post** /product/batch | Add multiple products to the store
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /product/{productId} | Get a single product
[**GetProducts**](ProductsApi.md#GetProducts) | **Get** /product | Get multiple products
[**UpdateProduct**](ProductsApi.md#UpdateProduct) | **Put** /product/{productId} | Modify a single product



## CreateProduct

> []Product CreateProduct(ctx).ShipbobChannelId(shipbobChannelId).CreateProduct(createProduct).Execute()

Add a single product to the store

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
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    createProduct := *openapiclient.NewCreateProduct("Medium Blue T-Shirt", "TShirtBlueM") // CreateProduct | The product to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.CreateProduct(context.Background()).ShipbobChannelId(shipbobChannelId).CreateProduct(createProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **createProduct** | [**CreateProduct**](CreateProduct.md) | The product to add | 

### Return type

[**[]Product**](Product.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductBatch

> []Product CreateProductBatch(ctx).ShipbobChannelId(shipbobChannelId).CreateProduct(createProduct).Execute()

Add multiple products to the store

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
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    createProduct := []openapiclient.CreateProduct{*openapiclient.NewCreateProduct("Medium Blue T-Shirt", "TShirtBlueM")} // []CreateProduct | List of products to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.CreateProductBatch(context.Background()).ShipbobChannelId(shipbobChannelId).CreateProduct(createProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.CreateProductBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductBatch`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.CreateProductBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **createProduct** | [**[]CreateProduct**](CreateProduct.md) | List of products to add | 

### Return type

[**[]Product**](Product.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> Product GetProduct(ctx, productId).ShipbobChannelId(shipbobChannelId).Execute()

Get a single product

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
    productId := int32(56) // int32 | Unique identifier of the product
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.GetProduct(context.Background(), productId).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | Unique identifier of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**Product**](Product.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> []Product GetProducts(ctx).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).Search(search).ActiveStatus(activeStatus).BundleStatus(bundleStatus).ShipbobChannelId(shipbobChannelId).Execute()

Get multiple products

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
    page := int32(56) // int32 | Page of products to get (optional)
    limit := int32(56) // int32 | Amount of products per page to request (optional)
    iDs := []int32{int32(123)} // []int32 | Comma separated list of product ids to filter by (optional)
    referenceIds := []string{"Inner_example"} // []string | Comma separated list of reference ids to filter by (optional)
    search := "search_example" // string | Search is available for 2 fields of the inventory record related to the product: Inventory ID and Name - 1. Expected behavior for search by Inventory ID is exact match 2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word,  but must be consecutive characters. This is not case sensitive. (optional)
    activeStatus := openapiclient.ProductActiveStatus("Any") // ProductActiveStatus | Status filter for products: - Any: Include both active and inactive - Active: Filter products that are Active - Inactive: Filter products that are Inactive (optional)
    bundleStatus := openapiclient.ProductBundleStatus("Any") // ProductBundleStatus | Bundle filter for products: - Any: Don't filter and consider products that are bundles or not bundles - Bundle: Filter by products that are bundles - NotBundle: Filter by products that are not bundles (optional)
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.GetProducts(context.Background()).Page(page).Limit(limit).IDs(iDs).ReferenceIds(referenceIds).Search(search).ActiveStatus(activeStatus).BundleStatus(bundleStatus).ShipbobChannelId(shipbobChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProducts`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page of products to get | 
 **limit** | **int32** | Amount of products per page to request | 
 **iDs** | **[]int32** | Comma separated list of product ids to filter by | 
 **referenceIds** | **[]string** | Comma separated list of reference ids to filter by | 
 **search** | **string** | Search is available for 2 fields of the inventory record related to the product: Inventory ID and Name - 1. Expected behavior for search by Inventory ID is exact match 2. Expected behavior for search by Inventory Name is partial match, i.e. does not have to be start of word,  but must be consecutive characters. This is not case sensitive. | 
 **activeStatus** | [**ProductActiveStatus**](ProductActiveStatus.md) | Status filter for products: - Any: Include both active and inactive - Active: Filter products that are Active - Inactive: Filter products that are Inactive | 
 **bundleStatus** | [**ProductBundleStatus**](ProductBundleStatus.md) | Bundle filter for products: - Any: Don&#39;t filter and consider products that are bundles or not bundles - Bundle: Filter by products that are bundles - NotBundle: Filter by products that are not bundles | 
 **shipbobChannelId** | **int32** | Channel Id for Operation | 

### Return type

[**[]Product**](Product.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> []Product UpdateProduct(ctx, productId).ShipbobChannelId(shipbobChannelId).UpdateProduct(updateProduct).Execute()

Modify a single product

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
    productId := int32(56) // int32 | Unique identifier of the product to modify
    shipbobChannelId := int32(56) // int32 | Channel Id for Operation
    updateProduct := *openapiclient.NewUpdateProduct("Medium Blue T-Shirt") // UpdateProduct | Updated fields to the product (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.UpdateProduct(context.Background(), productId).ShipbobChannelId(shipbobChannelId).UpdateProduct(updateProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.UpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProduct`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | Unique identifier of the product to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipbobChannelId** | **int32** | Channel Id for Operation | 
 **updateProduct** | [**UpdateProduct**](UpdateProduct.md) | Updated fields to the product | 

### Return type

[**[]Product**](Product.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

