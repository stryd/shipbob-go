# ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ShipbobReturnsPublicCommonReturnStatus**](Shipbob.Returns.Public.Common.ReturnStatus.md) |  | [optional] 
**TimeStamp** | Pointer to **time.Time** | Date this corresponding return order status was created | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel

`func NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel() *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel`

NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel instantiates a new ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel`

NewShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) GetStatus() ShipbobReturnsPublicCommonReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) GetStatusOk() (*ShipbobReturnsPublicCommonReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) SetStatus(v ShipbobReturnsPublicCommonReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeStamp

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderStatusHistoryViewModel) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


