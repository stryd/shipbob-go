# OrdersAddProductToOrderByReferenceIdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the product | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**ReferenceId** | **string** | Unique reference id of the product | 

## Methods

### NewOrdersAddProductToOrderByReferenceIdModel

`func NewOrdersAddProductToOrderByReferenceIdModel(quantity int32, referenceId string, ) *OrdersAddProductToOrderByReferenceIdModel`

NewOrdersAddProductToOrderByReferenceIdModel instantiates a new OrdersAddProductToOrderByReferenceIdModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersAddProductToOrderByReferenceIdModelWithDefaults

`func NewOrdersAddProductToOrderByReferenceIdModelWithDefaults() *OrdersAddProductToOrderByReferenceIdModel`

NewOrdersAddProductToOrderByReferenceIdModelWithDefaults instantiates a new OrdersAddProductToOrderByReferenceIdModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrdersAddProductToOrderByReferenceIdModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrdersAddProductToOrderByReferenceIdModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrdersAddProductToOrderByReferenceIdModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetReferenceId

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrdersAddProductToOrderByReferenceIdModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrdersAddProductToOrderByReferenceIdModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


