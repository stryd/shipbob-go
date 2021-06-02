# CreateOrderProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the product | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 

## Methods

### NewCreateOrderProducts

`func NewCreateOrderProducts(quantity int32, ) *CreateOrderProducts`

NewCreateOrderProducts instantiates a new CreateOrderProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderProductsWithDefaults

`func NewCreateOrderProductsWithDefaults() *CreateOrderProducts`

NewCreateOrderProductsWithDefaults instantiates a new CreateOrderProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrderProducts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrderProducts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrderProducts) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrderProducts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *CreateOrderProducts) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateOrderProducts) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateOrderProducts) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetName

`func (o *CreateOrderProducts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrderProducts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrderProducts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrderProducts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceId

`func (o *CreateOrderProducts) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateOrderProducts) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateOrderProducts) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CreateOrderProducts) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


