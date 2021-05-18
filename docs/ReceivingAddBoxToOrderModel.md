# ReceivingAddBoxToOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxItems** | [**[]ReceivingAddBoxItemToBoxModel**](ReceivingAddBoxItemToBoxModel.md) | Items contained in this box | 
**TrackingNumber** | **NullableString** | Tracking number for the box shipment | 

## Methods

### NewReceivingAddBoxToOrderModel

`func NewReceivingAddBoxToOrderModel(boxItems []ReceivingAddBoxItemToBoxModel, trackingNumber NullableString, ) *ReceivingAddBoxToOrderModel`

NewReceivingAddBoxToOrderModel instantiates a new ReceivingAddBoxToOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingAddBoxToOrderModelWithDefaults

`func NewReceivingAddBoxToOrderModelWithDefaults() *ReceivingAddBoxToOrderModel`

NewReceivingAddBoxToOrderModelWithDefaults instantiates a new ReceivingAddBoxToOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxItems

`func (o *ReceivingAddBoxToOrderModel) GetBoxItems() []ReceivingAddBoxItemToBoxModel`

GetBoxItems returns the BoxItems field if non-nil, zero value otherwise.

### GetBoxItemsOk

`func (o *ReceivingAddBoxToOrderModel) GetBoxItemsOk() (*[]ReceivingAddBoxItemToBoxModel, bool)`

GetBoxItemsOk returns a tuple with the BoxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxItems

`func (o *ReceivingAddBoxToOrderModel) SetBoxItems(v []ReceivingAddBoxItemToBoxModel)`

SetBoxItems sets BoxItems field to given value.


### SetBoxItemsNil

`func (o *ReceivingAddBoxToOrderModel) SetBoxItemsNil(b bool)`

 SetBoxItemsNil sets the value for BoxItems to be an explicit nil

### UnsetBoxItems
`func (o *ReceivingAddBoxToOrderModel) UnsetBoxItems()`

UnsetBoxItems ensures that no value is present for BoxItems, not even an explicit nil
### GetTrackingNumber

`func (o *ReceivingAddBoxToOrderModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReceivingAddBoxToOrderModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReceivingAddBoxToOrderModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### SetTrackingNumberNil

`func (o *ReceivingAddBoxToOrderModel) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReceivingAddBoxToOrderModel) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


