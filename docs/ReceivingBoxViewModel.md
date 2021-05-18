# ReceivingBoxViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivedDate** | Pointer to **NullableTime** | Date the box arrived | [optional] 
**BoxItems** | Pointer to [**[]ReceivingBoxItemViewModel**](ReceivingBoxItemViewModel.md) | Information about the items included in the box | [optional] 
**BoxNumber** | Pointer to **int32** | The number of the box in the receiving order | [optional] 
**BoxStatus** | Pointer to [**ReceivingBoxStatus**](Receiving.BoxStatus.md) |  | [optional] 
**CountingStartedDate** | Pointer to **NullableTime** | Date counting of the box&#39;s inventory items started | [optional] 
**ReceivedDate** | Pointer to **NullableTime** | Date the box was received | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the box shipment | [optional] 

## Methods

### NewReceivingBoxViewModel

`func NewReceivingBoxViewModel() *ReceivingBoxViewModel`

NewReceivingBoxViewModel instantiates a new ReceivingBoxViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingBoxViewModelWithDefaults

`func NewReceivingBoxViewModelWithDefaults() *ReceivingBoxViewModel`

NewReceivingBoxViewModelWithDefaults instantiates a new ReceivingBoxViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivedDate

`func (o *ReceivingBoxViewModel) GetArrivedDate() time.Time`

GetArrivedDate returns the ArrivedDate field if non-nil, zero value otherwise.

### GetArrivedDateOk

`func (o *ReceivingBoxViewModel) GetArrivedDateOk() (*time.Time, bool)`

GetArrivedDateOk returns a tuple with the ArrivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivedDate

`func (o *ReceivingBoxViewModel) SetArrivedDate(v time.Time)`

SetArrivedDate sets ArrivedDate field to given value.

### HasArrivedDate

`func (o *ReceivingBoxViewModel) HasArrivedDate() bool`

HasArrivedDate returns a boolean if a field has been set.

### SetArrivedDateNil

`func (o *ReceivingBoxViewModel) SetArrivedDateNil(b bool)`

 SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil

### UnsetArrivedDate
`func (o *ReceivingBoxViewModel) UnsetArrivedDate()`

UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
### GetBoxItems

`func (o *ReceivingBoxViewModel) GetBoxItems() []ReceivingBoxItemViewModel`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *ReceivingBoxViewModel) GetBoxItemsOk() (*[]ReceivingBoxItemViewModel, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *ReceivingBoxViewModel) SetBoxItems(v []ReceivingBoxItemViewModel)`

SetBoxItems sets BoxItems field to given value.

### HasBoxItems

`func (o *ReceivingBoxViewModel) HasBoxItems() bool`

HasBoxItems returns a boolean if a field has been set.

### SetBoxItemsNil

`func (o *ReceivingBoxViewModel) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *ReceivingBoxViewModel) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetBoxNumber

`func (o *ReceivingBoxViewModel) GetBoxNumber() int32`

GetBoxNumber returns the BoxNumber field if non-nil, zero value otherwise.

### GetBoxNumberOk

`func (o *ReceivingBoxViewModel) GetBoxNumberOk() (*int32, bool)`

GetBoxNumberOk returns a tuple with the BoxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxNumber

`func (o *ReceivingBoxViewModel) SetBoxNumber(v int32)`

SetBoxNumber sets BoxNumber field to given value.

### HasBoxNumber

`func (o *ReceivingBoxViewModel) HasBoxNumber() bool`

HasBoxNumber returns a boolean if a field has been set.

### GetBoxStatus

`func (o *ReceivingBoxViewModel) GetBoxStatus() ReceivingBoxStatus`

GetBoxStatus returns the BoxStatus field if non-nil, zero value otherwise.

### GetBoxStatusOk

`func (o *ReceivingBoxViewModel) GetBoxStatusOk() (*ReceivingBoxStatus, bool)`

GetBoxStatusOk returns a tuple with the BoxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxStatus

`func (o *ReceivingBoxViewModel) SetBoxStatus(v ReceivingBoxStatus)`

SetBoxStatus sets BoxStatus field to given value.

### HasBoxStatus

`func (o *ReceivingBoxViewModel) HasBoxStatus() bool`

HasBoxStatus returns a boolean if a field has been set.

### GetCountingStartedDate

`func (o *ReceivingBoxViewModel) GetCountingStartedDate() time.Time`

GetCountingStartedDate returns the CountingStartedDate field if non-nil, zero value otherwise.

### GetCountingStartedDateOk

`func (o *ReceivingBoxViewModel) GetCountingStartedDateOk() (*time.Time, bool)`

GetCountingStartedDateOk returns a tuple with the CountingStartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingStartedDate

`func (o *ReceivingBoxViewModel) SetCountingStartedDate(v time.Time)`

SetCountingStartedDate sets CountingStartedDate field to given value.

### HasCountingStartedDate

`func (o *ReceivingBoxViewModel) HasCountingStartedDate() bool`

HasCountingStartedDate returns a boolean if a field has been set.

### SetCountingStartedDateNil

`func (o *ReceivingBoxViewModel) SetCountingStartedDateNil(b bool)`

 SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil

### UnsetCountingStartedDate
`func (o *ReceivingBoxViewModel) UnsetCountingStartedDate()`

UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
### GetReceivedDate

`func (o *ReceivingBoxViewModel) GetReceivedDate() time.Time`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *ReceivingBoxViewModel) GetReceivedDateOk() (*time.Time, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDate

`func (o *ReceivingBoxViewModel) SetReceivedDate(v time.Time)`

SetReceivedDate sets ReceivedDate field to given value.

### HasReceivedDate

`func (o *ReceivingBoxViewModel) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### SetReceivedDateNil

`func (o *ReceivingBoxViewModel) SetReceivedDateNil(b bool)`

 SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil

### UnsetReceivedDate
`func (o *ReceivingBoxViewModel) UnsetReceivedDate()`

UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
### GetTrackingNumber

`func (o *ReceivingBoxViewModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReceivingBoxViewModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReceivingBoxViewModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ReceivingBoxViewModel) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ReceivingBoxViewModel) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReceivingBoxViewModel) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


