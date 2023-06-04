# AddProductToOrderByProductId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalLineId** | Pointer to **NullableInt32** | Numeric assignment per item. Used as a reference number for multiple purposes such as split orders, split containers, etc. | [optional] 
**Id** | **int32** | Unique id of the product | 
**Quantity** | **int32** | The quantity of this product ordered | 
**QuantityUnitOfMeasureCode** | Pointer to **string** | Defined standard for measure for an item (each, inner pack, case, pallet).  Values: EA, INP, CS and PL | [optional] 

## Methods

### NewAddProductToOrderByProductId

`func NewAddProductToOrderByProductId(id int32, quantity int32, ) *AddProductToOrderByProductId`

NewAddProductToOrderByProductId instantiates a new AddProductToOrderByProductId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddProductToOrderByProductIdWithDefaults

`func NewAddProductToOrderByProductIdWithDefaults() *AddProductToOrderByProductId`

NewAddProductToOrderByProductIdWithDefaults instantiates a new AddProductToOrderByProductId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalLineId

`func (o *AddProductToOrderByProductId) GetExternalLineId() int32`

GetExternalLineId returns the ExternalLineId field if non-nil, zero value otherwise.

### GetExternalLineIdOk

`func (o *AddProductToOrderByProductId) GetExternalLineIdOk() (*int32, bool)`

GetExternalLineIdOk returns a tuple with the ExternalLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLineId

`func (o *AddProductToOrderByProductId) SetExternalLineId(v int32)`

SetExternalLineId sets ExternalLineId field to given value.

### HasExternalLineId

`func (o *AddProductToOrderByProductId) HasExternalLineId() bool`

HasExternalLineId returns a boolean if a field has been set.

### SetExternalLineIdNil

`func (o *AddProductToOrderByProductId) SetExternalLineIdNil(b bool)`

 SetExternalLineIdNil sets the value for ExternalLineId to be an explicit nil

### UnsetExternalLineId
`func (o *AddProductToOrderByProductId) UnsetExternalLineId()`

UnsetExternalLineId ensures that no value is present for ExternalLineId, not even an explicit nil
### GetId

`func (o *AddProductToOrderByProductId) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddProductToOrderByProductId) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddProductToOrderByProductId) SetId(v int32)`

SetId sets Id field to given value.


### GetQuantity

`func (o *AddProductToOrderByProductId) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddProductToOrderByProductId) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddProductToOrderByProductId) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByProductId) GetQuantityUnitOfMeasureCode() string`

GetQuantityUnitOfMeasureCode returns the QuantityUnitOfMeasureCode field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureCodeOk

`func (o *AddProductToOrderByProductId) GetQuantityUnitOfMeasureCodeOk() (*string, bool)`

GetQuantityUnitOfMeasureCodeOk returns a tuple with the QuantityUnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByProductId) SetQuantityUnitOfMeasureCode(v string)`

SetQuantityUnitOfMeasureCode sets QuantityUnitOfMeasureCode field to given value.

### HasQuantityUnitOfMeasureCode

`func (o *AddProductToOrderByProductId) HasQuantityUnitOfMeasureCode() bool`

HasQuantityUnitOfMeasureCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


