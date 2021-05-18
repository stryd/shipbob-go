# ProductsCreateProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**Name** | **NullableString** | The name of the product | 
**ReferenceId** | **NullableString** | Unique reference identifier for the product. Any linked or generated inventory will also be uniquely identified by this value | 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 

## Methods

### NewProductsCreateProductModel

`func NewProductsCreateProductModel(name NullableString, referenceId NullableString, ) *ProductsCreateProductModel`

NewProductsCreateProductModel instantiates a new ProductsCreateProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsCreateProductModelWithDefaults

`func NewProductsCreateProductModelWithDefaults() *ProductsCreateProductModel`

NewProductsCreateProductModelWithDefaults instantiates a new ProductsCreateProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ProductsCreateProductModel) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductsCreateProductModel) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductsCreateProductModel) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductsCreateProductModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ProductsCreateProductModel) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ProductsCreateProductModel) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetName

`func (o *ProductsCreateProductModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsCreateProductModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsCreateProductModel) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProductsCreateProductModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductsCreateProductModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReferenceId

`func (o *ProductsCreateProductModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ProductsCreateProductModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ProductsCreateProductModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *ProductsCreateProductModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ProductsCreateProductModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetSku

`func (o *ProductsCreateProductModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductsCreateProductModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductsCreateProductModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductsCreateProductModel) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ProductsCreateProductModel) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ProductsCreateProductModel) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


