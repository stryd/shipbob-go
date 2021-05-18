# ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the inventory item to return | 
**Quantity** | **int32** | Quantity of the returned inventory item in the return | 
**RequestedAction** | Pointer to [**ShipbobReturnsPublicCommonReturnAction**](Shipbob.Returns.Public.Common.ReturnAction.md) |  | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModel

`func NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModel(id int32, quantity int32, ) *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel`

NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModel instantiates a new ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel`

NewShipbobReturnsPublicApiViewModelsReturnInventoryViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) SetId(v int32)`

SetId sets Id field to given value.


### GetQuantity

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetRequestedAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetRequestedAction() ShipbobReturnsPublicCommonReturnAction`

GetRequestedAction returns the RequestedAction field if non-nil, zero value otherwise.

### GetRequestedActionOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) GetRequestedActionOk() (*ShipbobReturnsPublicCommonReturnAction, bool)`

GetRequestedActionOk returns a tuple with the RequestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) SetRequestedAction(v ShipbobReturnsPublicCommonReturnAction)`

SetRequestedAction sets RequestedAction field to given value.

### HasRequestedAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel) HasRequestedAction() bool`

HasRequestedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


