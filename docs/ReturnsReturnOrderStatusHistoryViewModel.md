# ReturnsReturnOrderStatusHistoryViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReturnsReturnStatus**](Returns.ReturnStatus.md) |  | [optional] 
**TimeStamp** | Pointer to **time.Time** | Date this corresponding return order status was created | [optional] 

## Methods

### NewReturnsReturnOrderStatusHistoryViewModel

`func NewReturnsReturnOrderStatusHistoryViewModel() *ReturnsReturnOrderStatusHistoryViewModel`

NewReturnsReturnOrderStatusHistoryViewModel instantiates a new ReturnsReturnOrderStatusHistoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsReturnOrderStatusHistoryViewModelWithDefaults

`func NewReturnsReturnOrderStatusHistoryViewModelWithDefaults() *ReturnsReturnOrderStatusHistoryViewModel`

NewReturnsReturnOrderStatusHistoryViewModelWithDefaults instantiates a new ReturnsReturnOrderStatusHistoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReturnsReturnOrderStatusHistoryViewModel) GetStatus() ReturnsReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsReturnOrderStatusHistoryViewModel) GetStatusOk() (*ReturnsReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsReturnOrderStatusHistoryViewModel) SetStatus(v ReturnsReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsReturnOrderStatusHistoryViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ReturnsReturnOrderStatusHistoryViewModel) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ReturnsReturnOrderStatusHistoryViewModel) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ReturnsReturnOrderStatusHistoryViewModel) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ReturnsReturnOrderStatusHistoryViewModel) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


