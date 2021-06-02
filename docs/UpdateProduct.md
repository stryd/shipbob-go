# UpdateProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**Name** | **NullableString** | The name of the product | 
**Sku** | Pointer to **NullableString** | The stock keeping unit of the product | [optional] 

## Methods

### NewUpdateProduct

`func NewUpdateProduct(name NullableString, ) *UpdateProduct`

NewUpdateProduct instantiates a new UpdateProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductWithDefaults

`func NewUpdateProductWithDefaults() *UpdateProduct`

NewUpdateProductWithDefaults instantiates a new UpdateProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *UpdateProduct) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *UpdateProduct) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *UpdateProduct) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *UpdateProduct) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *UpdateProduct) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *UpdateProduct) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetName

`func (o *UpdateProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProduct) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateProduct) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateProduct) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSku

`func (o *UpdateProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdateProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *UpdateProduct) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *UpdateProduct) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


