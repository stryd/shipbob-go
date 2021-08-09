# ReturnOrderStatusHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReturnStatus**](ReturnStatus.md) |  | [optional] 
**TimeStamp** | Pointer to **NullableTime** | Date this corresponding return order status was created | [optional] 

## Methods

### NewReturnOrderStatusHistory

`func NewReturnOrderStatusHistory() *ReturnOrderStatusHistory`

NewReturnOrderStatusHistory instantiates a new ReturnOrderStatusHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnOrderStatusHistoryWithDefaults

`func NewReturnOrderStatusHistoryWithDefaults() *ReturnOrderStatusHistory`

NewReturnOrderStatusHistoryWithDefaults instantiates a new ReturnOrderStatusHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReturnOrderStatusHistory) GetStatus() ReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnOrderStatusHistory) GetStatusOk() (*ReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnOrderStatusHistory) SetStatus(v ReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnOrderStatusHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ReturnOrderStatusHistory) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ReturnOrderStatusHistory) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ReturnOrderStatusHistory) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ReturnOrderStatusHistory) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### SetTimeStampNil

`func (o *ReturnOrderStatusHistory) SetTimeStampNil(b bool)`

 SetTimeStampNil sets the value for TimeStamp to be an explicit nil

### UnsetTimeStamp
`func (o *ReturnOrderStatusHistory) UnsetTimeStamp()`

UnsetTimeStamp ensures that no value is present for TimeStamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


