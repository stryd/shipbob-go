# ReturnsInventoryItemViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRequested** | Pointer to [**ReturnsReturnActionRequestedViewModel**](Returns.ReturnActionRequestedViewModel.md) |  | [optional] 
**ActionTaken** | Pointer to [**[]ReturnsReturnActionTakenViewModel**](ReturnsReturnActionTakenViewModel.md) | Action(s) taken when processing the return | [optional] 
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity expected to be processed with the return | [optional] 

## Methods

### NewReturnsInventoryItemViewModel

`func NewReturnsInventoryItemViewModel() *ReturnsInventoryItemViewModel`

NewReturnsInventoryItemViewModel instantiates a new ReturnsInventoryItemViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsInventoryItemViewModelWithDefaults

`func NewReturnsInventoryItemViewModelWithDefaults() *ReturnsInventoryItemViewModel`

NewReturnsInventoryItemViewModelWithDefaults instantiates a new ReturnsInventoryItemViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRequested

`func (o *ReturnsInventoryItemViewModel) GetActionRequested() ReturnsReturnActionRequestedViewModel`

GetActionRequested returns the ActionRequested field if non-nil, zero value otherwise.

### GetActionRequestedOk

`func (o *ReturnsInventoryItemViewModel) GetActionRequestedOk() (*ReturnsReturnActionRequestedViewModel, bool)`

GetActionRequestedOk returns a tuple with the ActionRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequested

`func (o *ReturnsInventoryItemViewModel) SetActionRequested(v ReturnsReturnActionRequestedViewModel)`

SetActionRequested sets ActionRequested field to given value.

### HasActionRequested

`func (o *ReturnsInventoryItemViewModel) HasActionRequested() bool`

HasActionRequested returns a boolean if a field has been set.

### GetActionTaken

`func (o *ReturnsInventoryItemViewModel) GetActionTaken() []ReturnsReturnActionTakenViewModel`

GetActionTaken returns the ActionTaken field if non-nil, zero value otherwise.

### GetActionTakenOk

`func (o *ReturnsInventoryItemViewModel) GetActionTakenOk() (*[]ReturnsReturnActionTakenViewModel, bool)`

GetActionTakenOk returns a tuple with the ActionTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTaken

`func (o *ReturnsInventoryItemViewModel) SetActionTaken(v []ReturnsReturnActionTakenViewModel)`

SetActionTaken sets ActionTaken field to given value.

### HasActionTaken

`func (o *ReturnsInventoryItemViewModel) HasActionTaken() bool`

HasActionTaken returns a boolean if a field has been set.

### SetActionTakenNil

`func (o *ReturnsInventoryItemViewModel) SetActionTakenNil(b bool)`

 SetActionTakenNil sets the value for ActionTaken to be an explicit nil

### UnsetActionTaken
`func (o *ReturnsInventoryItemViewModel) UnsetActionTaken()`

UnsetActionTaken ensures that no value is present for ActionTaken, not even an explicit nil
### GetId

`func (o *ReturnsInventoryItemViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnsInventoryItemViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnsInventoryItemViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnsInventoryItemViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReturnsInventoryItemViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReturnsInventoryItemViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReturnsInventoryItemViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReturnsInventoryItemViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReturnsInventoryItemViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReturnsInventoryItemViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *ReturnsInventoryItemViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnsInventoryItemViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnsInventoryItemViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnsInventoryItemViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


