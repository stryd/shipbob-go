# AddProductToOrderByReferenceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the product | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**ReferenceId** | **string** | Unique reference id of the product | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


