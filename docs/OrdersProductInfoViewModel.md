# OrdersProductInfoViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique id of the product | [optional] 
**Quantity** | Pointer to **int32** | The quantity of this product ordered | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit for the product | [optional] 

## Methods

### NewOrdersProductInfoViewModel

`func NewOrdersProductInfoViewModel() *OrdersProductInfoViewModel`

NewOrdersProductInfoViewModel instantiates a new OrdersProductInfoViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersProductInfoViewModelWithDefaults

`func NewOrdersProductInfoViewModelWithDefaults() *OrdersProductInfoViewModel`

NewOrdersProductInfoViewModelWithDefaults instantiates a new OrdersProductInfoViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrdersProductInfoViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersProductInfoViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersProductInfoViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersProductInfoViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OrdersProductInfoViewModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrdersProductInfoViewModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuantity

`func (o *OrdersProductInfoViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrdersProductInfoViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrdersProductInfoViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrdersProductInfoViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReferenceId

`func (o *OrdersProductInfoViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrdersProductInfoViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrdersProductInfoViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *OrdersProductInfoViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSku

`func (o *OrdersProductInfoViewModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *OrdersProductInfoViewModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *OrdersProductInfoViewModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *OrdersProductInfoViewModel) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


