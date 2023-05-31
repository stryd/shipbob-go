# ReceivingOrderBoxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivedDate** | Pointer to **NullableTime** | Date the box arrived | [optional] 
**BoxItems** | Pointer to [**[]ReceivingOrderBoxesInnerBoxItemsInner**](ReceivingOrderBoxesInnerBoxItemsInner.md) | Information about the items included in the box | [optional] 
**BoxNumber** | Pointer to **int32** | The number of the box in the receiving order | [optional] 
**BoxStatus** | Pointer to **string** |  | [optional] 
**CountingStartedDate** | Pointer to **NullableTime** | Date counting of the box&#39;s inventory items started | [optional] 
**ReceivedDate** | Pointer to **NullableTime** | Date the box was received | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the box shipment | [optional] 

## Methods

### NewReceivingOrderBoxesInner

`func NewReceivingOrderBoxesInner() *ReceivingOrderBoxesInner`

NewReceivingOrderBoxesInner instantiates a new ReceivingOrderBoxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingOrderBoxesInnerWithDefaults

`func NewReceivingOrderBoxesInnerWithDefaults() *ReceivingOrderBoxesInner`

NewReceivingOrderBoxesInnerWithDefaults instantiates a new ReceivingOrderBoxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivedDate

`func (o *ReceivingOrderBoxesInner) GetArrivedDate() time.Time`

GetArrivedDate returns the ArrivedDate field if non-nil, zero value otherwise.

### GetArrivedDateOk

`func (o *ReceivingOrderBoxesInner) GetArrivedDateOk() (*time.Time, bool)`

GetArrivedDateOk returns a tuple with the ArrivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivedDate

`func (o *ReceivingOrderBoxesInner) SetArrivedDate(v time.Time)`

SetArrivedDate sets ArrivedDate field to given value.

### HasArrivedDate

`func (o *ReceivingOrderBoxesInner) HasArrivedDate() bool`

HasArrivedDate returns a boolean if a field has been set.

### SetArrivedDateNil

`func (o *ReceivingOrderBoxesInner) SetArrivedDateNil(b bool)`

 SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil

### UnsetArrivedDate
`func (o *ReceivingOrderBoxesInner) UnsetArrivedDate()`

UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
### GetBoxItems

`func (o *ReceivingOrderBoxesInner) GetBoxItems() []ReceivingOrderBoxesInnerBoxItemsInner`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *ReceivingOrderBoxesInner) GetBoxItemsOk() (*[]ReceivingOrderBoxesInnerBoxItemsInner, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *ReceivingOrderBoxesInner) SetBoxItems(v []ReceivingOrderBoxesInnerBoxItemsInner)`

SetBoxItems sets BoxItems field to given value.

### HasBoxItems

`func (o *ReceivingOrderBoxesInner) HasBoxItems() bool`

HasBoxItems returns a boolean if a field has been set.

### SetBoxItemsNil

`func (o *ReceivingOrderBoxesInner) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *ReceivingOrderBoxesInner) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetBoxNumber

`func (o *ReceivingOrderBoxesInner) GetBoxNumber() int32`

GetBoxNumber returns the BoxNumber field if non-nil, zero value otherwise.

### GetBoxNumberOk

`func (o *ReceivingOrderBoxesInner) GetBoxNumberOk() (*int32, bool)`

GetBoxNumberOk returns a tuple with the BoxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxNumber

`func (o *ReceivingOrderBoxesInner) SetBoxNumber(v int32)`

SetBoxNumber sets BoxNumber field to given value.

### HasBoxNumber

`func (o *ReceivingOrderBoxesInner) HasBoxNumber() bool`

HasBoxNumber returns a boolean if a field has been set.

### GetBoxStatus

`func (o *ReceivingOrderBoxesInner) GetBoxStatus() string`

GetBoxStatus returns the BoxStatus field if non-nil, zero value otherwise.

### GetBoxStatusOk

`func (o *ReceivingOrderBoxesInner) GetBoxStatusOk() (*string, bool)`

GetBoxStatusOk returns a tuple with the BoxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxStatus

`func (o *ReceivingOrderBoxesInner) SetBoxStatus(v string)`

SetBoxStatus sets BoxStatus field to given value.

### HasBoxStatus

`func (o *ReceivingOrderBoxesInner) HasBoxStatus() bool`

HasBoxStatus returns a boolean if a field has been set.

### GetCountingStartedDate

`func (o *ReceivingOrderBoxesInner) GetCountingStartedDate() time.Time`

GetCountingStartedDate returns the CountingStartedDate field if non-nil, zero value otherwise.

### GetCountingStartedDateOk

`func (o *ReceivingOrderBoxesInner) GetCountingStartedDateOk() (*time.Time, bool)`

GetCountingStartedDateOk returns a tuple with the CountingStartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountingStartedDate

`func (o *ReceivingOrderBoxesInner) SetCountingStartedDate(v time.Time)`

SetCountingStartedDate sets CountingStartedDate field to given value.

### HasCountingStartedDate

`func (o *ReceivingOrderBoxesInner) HasCountingStartedDate() bool`

HasCountingStartedDate returns a boolean if a field has been set.

### SetCountingStartedDateNil

`func (o *ReceivingOrderBoxesInner) SetCountingStartedDateNil(b bool)`

 SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil

### UnsetCountingStartedDate
`func (o *ReceivingOrderBoxesInner) UnsetCountingStartedDate()`

UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
### GetReceivedDate

`func (o *ReceivingOrderBoxesInner) GetReceivedDate() time.Time`

GetReceivedDate returns the ReceivedDate field if non-nil, zero value otherwise.

### GetReceivedDateOk

`func (o *ReceivingOrderBoxesInner) GetReceivedDateOk() (*time.Time, bool)`

GetReceivedDateOk returns a tuple with the ReceivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedDate

`func (o *ReceivingOrderBoxesInner) SetReceivedDate(v time.Time)`

SetReceivedDate sets ReceivedDate field to given value.

### HasReceivedDate

`func (o *ReceivingOrderBoxesInner) HasReceivedDate() bool`

HasReceivedDate returns a boolean if a field has been set.

### SetReceivedDateNil

`func (o *ReceivingOrderBoxesInner) SetReceivedDateNil(b bool)`

 SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil

### UnsetReceivedDate
`func (o *ReceivingOrderBoxesInner) UnsetReceivedDate()`

UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
### GetTrackingNumber

`func (o *ReceivingOrderBoxesInner) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReceivingOrderBoxesInner) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReceivingOrderBoxesInner) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ReceivingOrderBoxesInner) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ReceivingOrderBoxesInner) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReceivingOrderBoxesInner) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


