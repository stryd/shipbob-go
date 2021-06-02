# ReturnActionTaken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ReturnAction**](ReturnAction.md) |  | [optional] 
**ActionReason** | Pointer to **NullableString** | Reason the given action was taken | [optional] 
**QuantityProcessed** | Pointer to **int32** | Quantity of inventory processed with the taken action | [optional] 

## Methods

### NewReturnActionTaken

`func NewReturnActionTaken() *ReturnActionTaken`

NewReturnActionTaken instantiates a new ReturnActionTaken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnActionTakenWithDefaults

`func NewReturnActionTakenWithDefaults() *ReturnActionTaken`

NewReturnActionTakenWithDefaults instantiates a new ReturnActionTaken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ReturnActionTaken) GetAction() ReturnAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReturnActionTaken) GetActionOk() (*ReturnAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReturnActionTaken) SetAction(v ReturnAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReturnActionTaken) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionReason

`func (o *ReturnActionTaken) GetActionReason() string`

GetActionReason returns the ActionReason field if non-nil, zero value otherwise.

### GetActionReasonOk

`func (o *ReturnActionTaken) GetActionReasonOk() (*string, bool)`

GetActionReasonOk returns a tuple with the ActionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionReason

`func (o *ReturnActionTaken) SetActionReason(v string)`

SetActionReason sets ActionReason field to given value.

### HasActionReason

`func (o *ReturnActionTaken) HasActionReason() bool`

HasActionReason returns a boolean if a field has been set.

### SetActionReasonNil

`func (o *ReturnActionTaken) SetActionReasonNil(b bool)`

 SetActionReasonNil sets the value for ActionReason to be an explicit nil

### UnsetActionReason
`func (o *ReturnActionTaken) UnsetActionReason()`

UnsetActionReason ensures that no value is present for ActionReason, not even an explicit nil
### GetQuantityProcessed

`func (o *ReturnActionTaken) GetQuantityProcessed() int32`

GetQuantityProcessed returns the QuantityProcessed field if non-nil, zero value otherwise.

### GetQuantityProcessedOk

`func (o *ReturnActionTaken) GetQuantityProcessedOk() (*int32, bool)`

GetQuantityProcessedOk returns a tuple with the QuantityProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityProcessed

`func (o *ReturnActionTaken) SetQuantityProcessed(v int32)`

SetQuantityProcessed sets QuantityProcessed field to given value.

### HasQuantityProcessed

`func (o *ReturnActionTaken) HasQuantityProcessed() bool`

HasQuantityProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


