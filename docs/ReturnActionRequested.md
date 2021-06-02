# ReturnActionRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ReturnAction**](ReturnAction.md) |  | [optional] 
**ActionType** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **NullableString** | Specific instructions to be taken for the inventory when processing the return | [optional] 

## Methods

### NewReturnActionRequested

`func NewReturnActionRequested() *ReturnActionRequested`

NewReturnActionRequested instantiates a new ReturnActionRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnActionRequestedWithDefaults

`func NewReturnActionRequestedWithDefaults() *ReturnActionRequested`

NewReturnActionRequestedWithDefaults instantiates a new ReturnActionRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ReturnActionRequested) GetAction() ReturnAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReturnActionRequested) GetActionOk() (*ReturnAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReturnActionRequested) SetAction(v ReturnAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReturnActionRequested) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionType

`func (o *ReturnActionRequested) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ReturnActionRequested) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ReturnActionRequested) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ReturnActionRequested) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetInstructions

`func (o *ReturnActionRequested) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ReturnActionRequested) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ReturnActionRequested) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ReturnActionRequested) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ReturnActionRequested) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ReturnActionRequested) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


