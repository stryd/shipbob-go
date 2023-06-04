# ProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalLineId** | Pointer to **NullableInt32** | Numeric assignment per item. Used as a reference number for multiple purposes such as split orders, split containers, etc. | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number - unique and internationally recognized identifier assigned to item by company GS1 | [optional] 
**Id** | Pointer to **NullableInt32** | Unique id of the product | [optional] 
**Quantity** | Pointer to **int32** | The quantity of this product ordered | [optional] 
**QuantityUnitOfMeasureCode** | Pointer to **string** | Defined standard for measure for an item (each, inner pack, case, pallet).  Values: EA, INP, CS and PL | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit for the product | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | Price for one item | [optional] 
**Upc** | Pointer to **string** | Universal Product Code - Unique external identifier | [optional] 

## Methods

### NewProductInfo

`func NewProductInfo() *ProductInfo`

NewProductInfo instantiates a new ProductInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInfoWithDefaults

`func NewProductInfoWithDefaults() *ProductInfo`

NewProductInfoWithDefaults instantiates a new ProductInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalLineId

`func (o *ProductInfo) GetExternalLineId() int32`

GetExternalLineId returns the ExternalLineId field if non-nil, zero value otherwise.

### GetExternalLineIdOk

`func (o *ProductInfo) GetExternalLineIdOk() (*int32, bool)`

GetExternalLineIdOk returns a tuple with the ExternalLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLineId

`func (o *ProductInfo) SetExternalLineId(v int32)`

SetExternalLineId sets ExternalLineId field to given value.

### HasExternalLineId

`func (o *ProductInfo) HasExternalLineId() bool`

HasExternalLineId returns a boolean if a field has been set.

### SetExternalLineIdNil

`func (o *ProductInfo) SetExternalLineIdNil(b bool)`

 SetExternalLineIdNil sets the value for ExternalLineId to be an explicit nil

### UnsetExternalLineId
`func (o *ProductInfo) UnsetExternalLineId()`

UnsetExternalLineId ensures that no value is present for ExternalLineId, not even an explicit nil
### GetGtin

`func (o *ProductInfo) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductInfo) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductInfo) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductInfo) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetId

`func (o *ProductInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProductInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProductInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuantity

`func (o *ProductInfo) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductInfo) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductInfo) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityUnitOfMeasureCode

`func (o *ProductInfo) GetQuantityUnitOfMeasureCode() string`

GetQuantityUnitOfMeasureCode returns the QuantityUnitOfMeasureCode field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureCodeOk

`func (o *ProductInfo) GetQuantityUnitOfMeasureCodeOk() (*string, bool)`

GetQuantityUnitOfMeasureCodeOk returns a tuple with the QuantityUnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasureCode

`func (o *ProductInfo) SetQuantityUnitOfMeasureCode(v string)`

SetQuantityUnitOfMeasureCode sets QuantityUnitOfMeasureCode field to given value.

### HasQuantityUnitOfMeasureCode

`func (o *ProductInfo) HasQuantityUnitOfMeasureCode() bool`

HasQuantityUnitOfMeasureCode returns a boolean if a field has been set.

### GetReferenceId

`func (o *ProductInfo) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ProductInfo) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ProductInfo) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ProductInfo) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSku

`func (o *ProductInfo) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductInfo) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductInfo) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductInfo) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitPrice

`func (o *ProductInfo) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ProductInfo) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ProductInfo) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *ProductInfo) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *ProductInfo) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *ProductInfo) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetUpc

`func (o *ProductInfo) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductInfo) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductInfo) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductInfo) HasUpc() bool`

HasUpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


