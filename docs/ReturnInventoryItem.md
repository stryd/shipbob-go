# ReturnInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRequested** | Pointer to [**ReturnActionRequested**](ReturnActionRequested.md) |  | [optional] 
**ActionTaken** | Pointer to [**[]ReturnActionTaken**](ReturnActionTaken.md) | Action(s) taken when processing the return | [optional] 
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity expected to be processed with the return | [optional] 

## Methods

### NewReturnInventoryItem

`func NewReturnInventoryItem() *ReturnInventoryItem`

NewReturnInventoryItem instantiates a new ReturnInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnInventoryItemWithDefaults

`func NewReturnInventoryItemWithDefaults() *ReturnInventoryItem`

NewReturnInventoryItemWithDefaults instantiates a new ReturnInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRequested

`func (o *ReturnInventoryItem) GetActionRequested() ReturnActionRequested`

GetActionRequested returns the ActionRequested field if non-nil, zero value otherwise.

### GetActionRequestedOk

`func (o *ReturnInventoryItem) GetActionRequestedOk() (*ReturnActionRequested, bool)`

GetActionRequestedOk returns a tuple with the ActionRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequested

`func (o *ReturnInventoryItem) SetActionRequested(v ReturnActionRequested)`

SetActionRequested sets ActionRequested field to given value.

### HasActionRequested

`func (o *ReturnInventoryItem) HasActionRequested() bool`

HasActionRequested returns a boolean if a field has been set.

### GetActionTaken

`func (o *ReturnInventoryItem) GetActionTaken() []ReturnActionTaken`

GetActionTaken returns the ActionTaken field if non-nil, zero value otherwise.

### GetActionTakenOk

`func (o *ReturnInventoryItem) GetActionTakenOk() (*[]ReturnActionTaken, bool)`

GetActionTakenOk returns a tuple with the ActionTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTaken

`func (o *ReturnInventoryItem) SetActionTaken(v []ReturnActionTaken)`

SetActionTaken sets ActionTaken field to given value.

### HasActionTaken

`func (o *ReturnInventoryItem) HasActionTaken() bool`

HasActionTaken returns a boolean if a field has been set.

### SetActionTakenNil

`func (o *ReturnInventoryItem) SetActionTakenNil(b bool)`

 SetActionTakenNil sets the value for ActionTaken to be an explicit nil

### UnsetActionTaken
`func (o *ReturnInventoryItem) UnsetActionTaken()`

UnsetActionTaken ensures that no value is present for ActionTaken, not even an explicit nil
### GetId

`func (o *ReturnInventoryItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnInventoryItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnInventoryItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnInventoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReturnInventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReturnInventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReturnInventoryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReturnInventoryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReturnInventoryItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReturnInventoryItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *ReturnInventoryItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnInventoryItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnInventoryItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnInventoryItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


