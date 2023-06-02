# AddProductToOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique id of the product | 
**Quantity** | **int32** | The quantity of this product ordered | 
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferenceId** | **string** | Unique reference id of the product | 

## Methods

### NewAddProductToOrder

`func NewAddProductToOrder(id int32, quantity int32, referenceId string, ) *AddProductToOrder`

NewAddProductToOrder instantiates a new AddProductToOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddProductToOrderWithDefaults

`func NewAddProductToOrderWithDefaults() *AddProductToOrder`

NewAddProductToOrderWithDefaults instantiates a new AddProductToOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddProductToOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddProductToOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddProductToOrder) SetId(v int32)`

SetId sets Id field to given value.


### GetQuantity

`func (o *AddProductToOrder) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddProductToOrder) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddProductToOrder) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetName

`func (o *AddProductToOrder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddProductToOrder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddProductToOrder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddProductToOrder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceId

`func (o *AddProductToOrder) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AddProductToOrder) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AddProductToOrder) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

