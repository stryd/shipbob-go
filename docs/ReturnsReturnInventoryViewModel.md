# ReturnsReturnInventoryViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the inventory item to return | 
**Quantity** | **int32** | Quantity of the returned inventory item in the return | 
**RequestedAction** | Pointer to [**ReturnsReturnAction**](Returns.ReturnAction.md) |  | [optional] 

## Methods

### NewReturnsReturnInventoryViewModel

`func NewReturnsReturnInventoryViewModel(id int32, quantity int32, ) *ReturnsReturnInventoryViewModel`

NewReturnsReturnInventoryViewModel instantiates a new ReturnsReturnInventoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsReturnInventoryViewModelWithDefaults

`func NewReturnsReturnInventoryViewModelWithDefaults() *ReturnsReturnInventoryViewModel`

NewReturnsReturnInventoryViewModelWithDefaults instantiates a new ReturnsReturnInventoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReturnsReturnInventoryViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnsReturnInventoryViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnsReturnInventoryViewModel) SetId(v int32)`

SetId sets Id field to given value.


### GetQuantity

`func (o *ReturnsReturnInventoryViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnsReturnInventoryViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnsReturnInventoryViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetRequestedAction

`func (o *ReturnsReturnInventoryViewModel) GetRequestedAction() ReturnsReturnAction`

GetRequestedAction returns the RequestedAction field if non-nil, zero value otherwise.

### GetRequestedActionOk

`func (o *ReturnsReturnInventoryViewModel) GetRequestedActionOk() (*ReturnsReturnAction, bool)`

GetRequestedActionOk returns a tuple with the RequestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAction

`func (o *ReturnsReturnInventoryViewModel) SetRequestedAction(v ReturnsReturnAction)`

SetRequestedAction sets RequestedAction field to given value.

### HasRequestedAction

`func (o *ReturnsReturnInventoryViewModel) HasRequestedAction() bool`

HasRequestedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


