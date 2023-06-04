# AddProductToOrderByReferenceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalLineId** | Pointer to **NullableInt32** | Numeric assignment per item. Used as a reference number for multiple purposes such as split orders, split containers, etc. | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number - unique and internationally recognized identifier assigned to item by company GS1 | [optional] 
**Name** | Pointer to **string** | Name of the product. Required if there is not an existing ShipBob product with a matching reference_id value. | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**QuantityUnitOfMeasureCode** | Pointer to **string** | Defined standard for measure for an item (each, inner pack, case, pallet).  Values: EA, INP, CS and PL | [optional] 
**ReferenceId** | **string** | Unique reference id of the product | 
**Sku** | Pointer to **string** | Product SKU | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | Price for one item | [optional] 
**Upc** | Pointer to **string** | Universal Product Code - Unique external identifier | [optional] 

## Methods

### NewAddProductToOrderByReferenceId

`func NewAddProductToOrderByReferenceId(quantity int32, referenceId string, ) *AddProductToOrderByReferenceId`

NewAddProductToOrderByReferenceId instantiates a new AddProductToOrderByReferenceId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddProductToOrderByReferenceIdWithDefaults

`func NewAddProductToOrderByReferenceIdWithDefaults() *AddProductToOrderByReferenceId`

NewAddProductToOrderByReferenceIdWithDefaults instantiates a new AddProductToOrderByReferenceId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalLineId

`func (o *AddProductToOrderByReferenceId) GetExternalLineId() int32`

GetExternalLineId returns the ExternalLineId field if non-nil, zero value otherwise.

### GetExternalLineIdOk

`func (o *AddProductToOrderByReferenceId) GetExternalLineIdOk() (*int32, bool)`

GetExternalLineIdOk returns a tuple with the ExternalLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLineId

`func (o *AddProductToOrderByReferenceId) SetExternalLineId(v int32)`

SetExternalLineId sets ExternalLineId field to given value.

### HasExternalLineId

`func (o *AddProductToOrderByReferenceId) HasExternalLineId() bool`

HasExternalLineId returns a boolean if a field has been set.

### SetExternalLineIdNil

`func (o *AddProductToOrderByReferenceId) SetExternalLineIdNil(b bool)`

 SetExternalLineIdNil sets the value for ExternalLineId to be an explicit nil

### UnsetExternalLineId
`func (o *AddProductToOrderByReferenceId) UnsetExternalLineId()`

UnsetExternalLineId ensures that no value is present for ExternalLineId, not even an explicit nil
### GetGtin

`func (o *AddProductToOrderByReferenceId) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *AddProductToOrderByReferenceId) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *AddProductToOrderByReferenceId) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *AddProductToOrderByReferenceId) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetName

`func (o *AddProductToOrderByReferenceId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddProductToOrderByReferenceId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddProductToOrderByReferenceId) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddProductToOrderByReferenceId) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *AddProductToOrderByReferenceId) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddProductToOrderByReferenceId) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddProductToOrderByReferenceId) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByReferenceId) GetQuantityUnitOfMeasureCode() string`

GetQuantityUnitOfMeasureCode returns the QuantityUnitOfMeasureCode field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureCodeOk

`func (o *AddProductToOrderByReferenceId) GetQuantityUnitOfMeasureCodeOk() (*string, bool)`

GetQuantityUnitOfMeasureCodeOk returns a tuple with the QuantityUnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByReferenceId) SetQuantityUnitOfMeasureCode(v string)`

SetQuantityUnitOfMeasureCode sets QuantityUnitOfMeasureCode field to given value.

### HasQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByReferenceId) HasQuantityUnitOfMeasureCode() bool`

HasQuantityUnitOfMeasureCode returns a boolean if a field has been set.

### GetReferenceId

`func (o *AddProductToOrderByReferenceId) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AddProductToOrderByReferenceId) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AddProductToOrderByReferenceId) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetSku

`func (o *AddProductToOrderByReferenceId) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddProductToOrderByReferenceId) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AddProductToOrderByReferenceId) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *AddProductToOrderByReferenceId) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitPrice

`func (o *AddProductToOrderByReferenceId) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *AddProductToOrderByReferenceId) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *AddProductToOrderByReferenceId) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *AddProductToOrderByReferenceId) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *AddProductToOrderByReferenceId) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *AddProductToOrderByReferenceId) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetUpc

`func (o *AddProductToOrderByReferenceId) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *AddProductToOrderByReferenceId) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *AddProductToOrderByReferenceId) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *AddProductToOrderByReferenceId) HasUpc() bool`

HasUpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


