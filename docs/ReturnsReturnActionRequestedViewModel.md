# ReturnsReturnActionRequestedViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ReturnsReturnAction**](Returns.ReturnAction.md) |  | [optional] 
**ActionType** | Pointer to [**ReturnsReturnActionSource**](Returns.ReturnActionSource.md) |  | [optional] 
**Instructions** | Pointer to **NullableString** | Specific instructions to be taken for the inventory when processing the return | [optional] 

## Methods

### NewReturnsReturnActionRequestedViewModel

`func NewReturnsReturnActionRequestedViewModel() *ReturnsReturnActionRequestedViewModel`

NewReturnsReturnActionRequestedViewModel instantiates a new ReturnsReturnActionRequestedViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsReturnActionRequestedViewModelWithDefaults

`func NewReturnsReturnActionRequestedViewModelWithDefaults() *ReturnsReturnActionRequestedViewModel`

NewReturnsReturnActionRequestedViewModelWithDefaults instantiates a new ReturnsReturnActionRequestedViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ReturnsReturnActionRequestedViewModel) GetAction() ReturnsReturnAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReturnsReturnActionRequestedViewModel) GetActionOk() (*ReturnsReturnAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReturnsReturnActionRequestedViewModel) SetAction(v ReturnsReturnAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReturnsReturnActionRequestedViewModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionType

`func (o *ReturnsReturnActionRequestedViewModel) GetActionType() ReturnsReturnActionSource`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ReturnsReturnActionRequestedViewModel) GetActionTypeOk() (*ReturnsReturnActionSource, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ReturnsReturnActionRequestedViewModel) SetActionType(v ReturnsReturnActionSource)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ReturnsReturnActionRequestedViewModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetInstructions

`func (o *ReturnsReturnActionRequestedViewModel) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ReturnsReturnActionRequestedViewModel) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ReturnsReturnActionRequestedViewModel) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ReturnsReturnActionRequestedViewModel) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ReturnsReturnActionRequestedViewModel) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ReturnsReturnActionRequestedViewModel) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


