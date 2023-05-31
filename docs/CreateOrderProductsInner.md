# CreateOrderProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the product | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 

## Methods

### NewCreateOrderProductsInner

`func NewCreateOrderProductsInner(quantity int32, ) *CreateOrderProductsInner`

NewCreateOrderProductsInner instantiates a new CreateOrderProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderProductsInnerWithDefaults

`func NewCreateOrderProductsInnerWithDefaults() *CreateOrderProductsInner`

NewCreateOrderProductsInnerWithDefaults instantiates a new CreateOrderProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrderProductsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrderProductsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrderProductsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrderProductsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *CreateOrderProductsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateOrderProductsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateOrderProductsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetName

`func (o *CreateOrderProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrderProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrderProductsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrderProductsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceId

`func (o *CreateOrderProductsInner) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateOrderProductsInner) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateOrderProductsInner) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CreateOrderProductsInner) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


