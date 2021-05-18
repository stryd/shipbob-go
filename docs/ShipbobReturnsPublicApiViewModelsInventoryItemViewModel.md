# ShipbobReturnsPublicApiViewModelsInventoryItemViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRequested** | Pointer to [**ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel**](Shipbob.Returns.Public.Api.ViewModels.ReturnActionRequestedViewModel.md) |  | [optional] 
**ActionTaken** | Pointer to [**[]ShipbobReturnsPublicApiViewModelsReturnActionTakenViewModel**](ShipbobReturnsPublicApiViewModelsReturnActionTakenViewModel.md) | Action(s) taken when processing the return | [optional] 
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity expected to be processed with the return | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsInventoryItemViewModel

`func NewShipbobReturnsPublicApiViewModelsInventoryItemViewModel() *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel`

NewShipbobReturnsPublicApiViewModelsInventoryItemViewModel instantiates a new ShipbobReturnsPublicApiViewModelsInventoryItemViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsInventoryItemViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsInventoryItemViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel`

NewShipbobReturnsPublicApiViewModelsInventoryItemViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsInventoryItemViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRequested

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetActionRequested() ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel`

GetActionRequested returns the ActionRequested field if non-nil, zero value otherwise.

### GetActionRequestedOk

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetActionRequestedOk() (*ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel, bool)`

GetActionRequestedOk returns a tuple with the ActionRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequested

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetActionRequested(v ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel)`

SetActionRequested sets ActionRequested field to given value.

### HasActionRequested

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) HasActionRequested() bool`

HasActionRequested returns a boolean if a field has been set.

### GetActionTaken

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetActionTaken() []ShipbobReturnsPublicApiViewModelsReturnActionTakenViewModel`

GetActionTaken returns the ActionTaken field if non-nil, zero value otherwise.

### GetActionTakenOk

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetActionTakenOk() (*[]ShipbobReturnsPublicApiViewModelsReturnActionTakenViewModel, bool)`

GetActionTakenOk returns a tuple with the ActionTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTaken

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetActionTaken(v []ShipbobReturnsPublicApiViewModelsReturnActionTakenViewModel)`

SetActionTaken sets ActionTaken field to given value.

### HasActionTaken

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) HasActionTaken() bool`

HasActionTaken returns a boolean if a field has been set.

### SetActionTakenNil

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetActionTakenNil(b bool)`

 SetActionTakenNil sets the value for ActionTaken to be an explicit nil

### UnsetActionTaken
`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) UnsetActionTaken()`

UnsetActionTaken ensures that no value is present for ActionTaken, not even an explicit nil
### GetId

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShipbobReturnsPublicApiViewModelsInventoryItemViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


