# ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ShipbobReturnsPublicCommonReturnAction**](Shipbob.Returns.Public.Common.ReturnAction.md) |  | [optional] 
**ActionType** | Pointer to [**ShipbobReturnsPublicCommonReturnActionSource**](Shipbob.Returns.Public.Common.ReturnActionSource.md) |  | [optional] 
**Instructions** | Pointer to **NullableString** | Specific instructions to be taken for the inventory when processing the return | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel

`func NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel() *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel`

NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel instantiates a new ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel`

NewShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetAction() ShipbobReturnsPublicCommonReturnAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetActionOk() (*ShipbobReturnsPublicCommonReturnAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) SetAction(v ShipbobReturnsPublicCommonReturnAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionType

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetActionType() ShipbobReturnsPublicCommonReturnActionSource`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetActionTypeOk() (*ShipbobReturnsPublicCommonReturnActionSource, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) SetActionType(v ShipbobReturnsPublicCommonReturnActionSource)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetInstructions

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ShipbobReturnsPublicApiViewModelsReturnActionRequestedViewModel) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


