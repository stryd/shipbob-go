# Box

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivedDate** | Pointer to **NullableTime** | Date the box arrived | [optional] 
**BoxItems** | Pointer to [**[]BoxItem**](BoxItem.md) | Information about the items included in the box | [optional] 
**BoxNumber** | Pointer to **int32** | The number of the box in the receiving order | [optional] 
**BoxStatus** | Pointer to [**BoxStatus**](BoxStatus.md) |  | [optional] 
**CountingStartedDate** | Pointer to **NullableTime** | Date counting of the box&#39;s inventory items started | [optional] 
**ReceivedDate** | Pointer to **NullableTime** | Date the box was received | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the box shipment | [optional] 

## Methods

### NewBox

`func NewBox() *Box`

NewBox instantiates a new Box object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxWithDefaults

`func NewBoxWithDefaults() *Box`

NewBoxWithDefaults instantiates a new Box object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivedDate

`func (o *Box) GetArrivedDate() time.Time`

GetArrivedDate returns the ArrivedDate field if non-nil, zero value otherwise.

### GetArrivedDateOk

`func (o *Box) GetArrivedDateOk() (*time.Time, bool)`

GetArrivedDateOk returns a tuple with the ArrivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivedDate

`func (o *Box) SetArrivedDate(v time.Time)`

SetArrivedDate sets ArrivedDate field to given value.

### HasArrivedDate

`func (o *Box) HasArrivedDate() bool`

HasArrivedDate returns a boolean if a field has been set.

### SetArrivedDateNil

`func (o *Box) SetArrivedDateNil(b bool)`

 SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil

### UnsetArrivedDate
`func (o *Box) UnsetArrivedDate()`

UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
### GetBoxItems

`func (o *Box) GetBoxItems() []BoxItem`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *Box) GetBoxItemsOk() (*[]BoxItem, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *Box) SetBoxItems(v []BoxItem)`

SetBoxItems sets BoxItems field to given value.

### HasBoxItems

`func (o *Box) HasBoxItems() bool`

HasBoxItems returns a boolean if a field has been set.

### SetBoxItemsNil

`func (o *Box) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *Box) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetBoxNumber

`func (o *Box) GetBoxNumber() int32`

GetBoxNumber returns the BoxNumber field if non-nil, zero value otherwise.

### GetBoxNumberOk

`func (o *Box) GetBoxNumberOk() (*int32, bool)`

GetBoxNumberOk returns a tuple with the BoxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxNumber

`func (o *Box) SetBoxNumber(v int32)`

SetBoxNumber sets BoxNumber field to given value.

### HasBoxNumber

`func (o *Box) HasBoxNumber() bool`

HasBoxNumber returns a boolean if a field has been set.

### GetBoxStatus

`func (o *Box) GetBoxStatus() BoxStatus`

GetBoxStatus returns the BoxStatus field if non-nil, zero value otherwise.

### GetBoxStatusOk

`func (o *Box) GetBoxStatusOk() (*BoxStatus, bool)`

GetBoxStatusOk returns a tuple with the BoxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxStatus

`func (o *Box) SetBoxStatus(v BoxStatus)`

SetBoxStatus sets BoxStatus field to given value.

### HasBoxStatus

`func (o *Box) HasBoxStatus() bool`

HasBoxStatus returns a boolean if a field has been set.

### GetCountingStartedDate

`func (o *Box) GetCountingStartedDate() time.Time`

GetCountingStartedDate returns the CountingStartedDate field if non-nil, zero value otherwise.

### GetCountingStartedDateOk

`func (o *Box) GetCountingStartedDateOk() (*time.Time, bool)`

GetCountingStartedDateOk returns a tuple with the CountingStartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingStartedDate

`func (o *Box) SetCountingStartedDate(v time.Time)`

SetCountingStartedDate sets CountingStartedDate field to given value.

### HasCountingStartedDate

`func (o *Box) HasCountingStartedDate() bool`

HasCountingStartedDate returns a boolean if a field has been set.

### SetCountingStartedDateNil

`func (o *Box) SetCountingStartedDateNil(b bool)`

 SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil

### UnsetCountingStartedDate
`func (o *Box) UnsetCountingStartedDate()`

UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
### GetReceivedDate

`func (o *Box) GetReceivedDate() time.Time`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *Box) GetReceivedDateOk() (*time.Time, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDate

`func (o *Box) SetReceivedDate(v time.Time)`

SetReceivedDate sets ReceivedDate field to given value.

### HasReceivedDate

`func (o *Box) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### SetReceivedDateNil

`func (o *Box) SetReceivedDateNil(b bool)`

 SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil

### UnsetReceivedDate
`func (o *Box) UnsetReceivedDate()`

UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
### GetTrackingNumber

`func (o *Box) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *Box) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *Box) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *Box) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *Box) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *Box) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


