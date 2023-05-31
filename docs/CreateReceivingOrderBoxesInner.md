# CreateReceivingOrderBoxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxItems** | [**[]CreateReceivingOrderBoxesInnerBoxItemsInner**](CreateReceivingOrderBoxesInnerBoxItemsInner.md) | Items contained in this box | 
**TrackingNumber** | **NullableString** | Tracking number for the box shipment | 

## Methods

### NewCreateReceivingOrderBoxesInner

`func NewCreateReceivingOrderBoxesInner(boxItems []CreateReceivingOrderBoxesInnerBoxItemsInner, trackingNumber NullableString, ) *CreateReceivingOrderBoxesInner`

NewCreateReceivingOrderBoxesInner instantiates a new CreateReceivingOrderBoxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceivingOrderBoxesInnerWithDefaults

`func NewCreateReceivingOrderBoxesInnerWithDefaults() *CreateReceivingOrderBoxesInner`

NewCreateReceivingOrderBoxesInnerWithDefaults instantiates a new CreateReceivingOrderBoxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxItems

`func (o *CreateReceivingOrderBoxesInner) GetBoxItems() []CreateReceivingOrderBoxesInnerBoxItemsInner`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *CreateReceivingOrderBoxesInner) GetBoxItemsOk() (*[]CreateReceivingOrderBoxesInnerBoxItemsInner, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *CreateReceivingOrderBoxesInner) SetBoxItems(v []CreateReceivingOrderBoxesInnerBoxItemsInner)`

SetBoxItems sets BoxItems field to given value.


### SetBoxItemsNil

`func (o *CreateReceivingOrderBoxesInner) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *CreateReceivingOrderBoxesInner) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetTrackingNumber

`func (o *CreateReceivingOrderBoxesInner) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *CreateReceivingOrderBoxesInner) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *CreateReceivingOrderBoxesInner) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### SetTrackingNumberNil

`func (o *CreateReceivingOrderBoxesInner) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *CreateReceivingOrderBoxesInner) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


