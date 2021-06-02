# ReceivingOrderBoxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivedDate** | Pointer to **NullableTime** | Date the box arrived | [optional] 
**BoxItems** | Pointer to [**[]ReceivingOrderBoxItems**](ReceivingOrderBoxItems.md) | Information about the items included in the box | [optional] 
**BoxNumber** | Pointer to **int32** | The number of the box in the receiving order | [optional] 
**BoxStatus** | Pointer to **string** |  | [optional] 
**CountingStartedDate** | Pointer to **NullableTime** | Date counting of the box&#39;s inventory items started | [optional] 
**ReceivedDate** | Pointer to **NullableTime** | Date the box was received | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the box shipment | [optional] 

## Methods

### NewReceivingOrderBoxes

`func NewReceivingOrderBoxes() *ReceivingOrderBoxes`

NewReceivingOrderBoxes instantiates a new ReceivingOrderBoxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingOrderBoxesWithDefaults

`func NewReceivingOrderBoxesWithDefaults() *ReceivingOrderBoxes`

NewReceivingOrderBoxesWithDefaults instantiates a new ReceivingOrderBoxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivedDate

`func (o *ReceivingOrderBoxes) GetArrivedDate() time.Time`

GetArrivedDate returns the ArrivedDate field if non-nil, zero value otherwise.

### GetArrivedDateOk

`func (o *ReceivingOrderBoxes) GetArrivedDateOk() (*time.Time, bool)`

GetArrivedDateOk returns a tuple with the ArrivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivedDate

`func (o *ReceivingOrderBoxes) SetArrivedDate(v time.Time)`

SetArrivedDate sets ArrivedDate field to given value.

### HasArrivedDate

`func (o *ReceivingOrderBoxes) HasArrivedDate() bool`

HasArrivedDate returns a boolean if a field has been set.

### SetArrivedDateNil

`func (o *ReceivingOrderBoxes) SetArrivedDateNil(b bool)`

 SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil

### UnsetArrivedDate
`func (o *ReceivingOrderBoxes) UnsetArrivedDate()`

UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
### GetBoxItems

`func (o *ReceivingOrderBoxes) GetBoxItems() []ReceivingOrderBoxItems`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *ReceivingOrderBoxes) GetBoxItemsOk() (*[]ReceivingOrderBoxItems, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *ReceivingOrderBoxes) SetBoxItems(v []ReceivingOrderBoxItems)`

SetBoxItems sets BoxItems field to given value.

### HasBoxItems

`func (o *ReceivingOrderBoxes) HasBoxItems() bool`

HasBoxItems returns a boolean if a field has been set.

### SetBoxItemsNil

`func (o *ReceivingOrderBoxes) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *ReceivingOrderBoxes) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetBoxNumber

`func (o *ReceivingOrderBoxes) GetBoxNumber() int32`

GetBoxNumber returns the BoxNumber field if non-nil, zero value otherwise.

### GetBoxNumberOk

`func (o *ReceivingOrderBoxes) GetBoxNumberOk() (*int32, bool)`

GetBoxNumberOk returns a tuple with the BoxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxNumber

`func (o *ReceivingOrderBoxes) SetBoxNumber(v int32)`

SetBoxNumber sets BoxNumber field to given value.

### HasBoxNumber

`func (o *ReceivingOrderBoxes) HasBoxNumber() bool`

HasBoxNumber returns a boolean if a field has been set.

### GetBoxStatus

`func (o *ReceivingOrderBoxes) GetBoxStatus() string`

GetBoxStatus returns the BoxStatus field if non-nil, zero value otherwise.

### GetBoxStatusOk

`func (o *ReceivingOrderBoxes) GetBoxStatusOk() (*string, bool)`

GetBoxStatusOk returns a tuple with the BoxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxStatus

`func (o *ReceivingOrderBoxes) SetBoxStatus(v string)`

SetBoxStatus sets BoxStatus field to given value.

### HasBoxStatus

`func (o *ReceivingOrderBoxes) HasBoxStatus() bool`

HasBoxStatus returns a boolean if a field has been set.

### GetCountingStartedDate

`func (o *ReceivingOrderBoxes) GetCountingStartedDate() time.Time`

GetCountingStartedDate returns the CountingStartedDate field if non-nil, zero value otherwise.

### GetCountingStartedDateOk

`func (o *ReceivingOrderBoxes) GetCountingStartedDateOk() (*time.Time, bool)`

GetCountingStartedDateOk returns a tuple with the CountingStartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingStartedDate

`func (o *ReceivingOrderBoxes) SetCountingStartedDate(v time.Time)`

SetCountingStartedDate sets CountingStartedDate field to given value.

### HasCountingStartedDate

`func (o *ReceivingOrderBoxes) HasCountingStartedDate() bool`

HasCountingStartedDate returns a boolean if a field has been set.

### SetCountingStartedDateNil

`func (o *ReceivingOrderBoxes) SetCountingStartedDateNil(b bool)`

 SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil

### UnsetCountingStartedDate
`func (o *ReceivingOrderBoxes) UnsetCountingStartedDate()`

UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
### GetReceivedDate

`func (o *ReceivingOrderBoxes) GetReceivedDate() time.Time`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *ReceivingOrderBoxes) GetReceivedDateOk() (*time.Time, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDate

`func (o *ReceivingOrderBoxes) SetReceivedDate(v time.Time)`

SetReceivedDate sets ReceivedDate field to given value.

### HasReceivedDate

`func (o *ReceivingOrderBoxes) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### SetReceivedDateNil

`func (o *ReceivingOrderBoxes) SetReceivedDateNil(b bool)`

 SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil

### UnsetReceivedDate
`func (o *ReceivingOrderBoxes) UnsetReceivedDate()`

UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
### GetTrackingNumber

`func (o *ReceivingOrderBoxes) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReceivingOrderBoxes) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReceivingOrderBoxes) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ReceivingOrderBoxes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ReceivingOrderBoxes) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReceivingOrderBoxes) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


