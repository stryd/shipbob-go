# CreateProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**Gtin** | Pointer to **NullableString** | Global Trade Item Number - unique and internationally recognized identifier assigned to item by company GS1. | [optional] 
**Name** | **string** | The name of the product | 
**ReferenceId** | **string** | Unique reference identifier for the product. Any linked or generated inventory will also be uniquely identified by this value | 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | The price of one unit | [optional] 
**Upc** | Pointer to **NullableString** | Universal Product Code - Unique external identifier | [optional] 

## Methods

### NewCreateProduct

`func NewCreateProduct(name string, referenceId string, ) *CreateProduct`

NewCreateProduct instantiates a new CreateProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductWithDefaults

`func NewCreateProductWithDefaults() *CreateProduct`

NewCreateProductWithDefaults instantiates a new CreateProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *CreateProduct) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *CreateProduct) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *CreateProduct) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *CreateProduct) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *CreateProduct) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *CreateProduct) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetGtin

`func (o *CreateProduct) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *CreateProduct) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *CreateProduct) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *CreateProduct) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### SetGtinNil

`func (o *CreateProduct) SetGtinNil(b bool)`

 SetGtinNil sets the value for Gtin to be an explicit nil

### UnsetGtin
`func (o *CreateProduct) UnsetGtin()`

UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
### GetName

`func (o *CreateProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProduct) SetName(v string)`

SetName sets Name field to given value.


### GetReferenceId

`func (o *CreateProduct) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateProduct) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateProduct) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetSku

`func (o *CreateProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CreateProduct) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CreateProduct) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetUnitPrice

`func (o *CreateProduct) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CreateProduct) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CreateProduct) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CreateProduct) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *CreateProduct) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *CreateProduct) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetUpc

`func (o *CreateProduct) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *CreateProduct) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *CreateProduct) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *CreateProduct) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *CreateProduct) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *CreateProduct) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


