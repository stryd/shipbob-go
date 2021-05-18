# ShipbobProductsApiModelsPublicCreateProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**Name** | **NullableString** | The name of the product | 
**ReferenceId** | **NullableString** | Unique reference identifier for the product. Any linked or generated inventory will also be uniquely identified by this value | 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 

## Methods

### NewShipbobProductsApiModelsPublicCreateProductModel

`func NewShipbobProductsApiModelsPublicCreateProductModel(name NullableString, referenceId NullableString, ) *ShipbobProductsApiModelsPublicCreateProductModel`

NewShipbobProductsApiModelsPublicCreateProductModel instantiates a new ShipbobProductsApiModelsPublicCreateProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobProductsApiModelsPublicCreateProductModelWithDefaults

`func NewShipbobProductsApiModelsPublicCreateProductModelWithDefaults() *ShipbobProductsApiModelsPublicCreateProductModel`

NewShipbobProductsApiModelsPublicCreateProductModelWithDefaults instantiates a new ShipbobProductsApiModelsPublicCreateProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ShipbobProductsApiModelsPublicCreateProductModel) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetName

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShipbobProductsApiModelsPublicCreateProductModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReferenceId

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ShipbobProductsApiModelsPublicCreateProductModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetSku

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ShipbobProductsApiModelsPublicCreateProductModel) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ShipbobProductsApiModelsPublicCreateProductModel) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


