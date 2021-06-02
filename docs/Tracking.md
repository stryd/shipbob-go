# Tracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | Carrier of the shipment | [optional] 
**CarrierService** | Pointer to **string** | The carrier&#39;s service which was used for this shipment | [optional] 
**TrackingNumber** | Pointer to **string** | Tracking number of the shipment | [optional] 
**TrackingUrl** | Pointer to **string** | URL to the website where a shipment can be tracked | [optional] 

## Methods

### NewTracking

`func NewTracking() *Tracking`

NewTracking instantiates a new Tracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingWithDefaults

`func NewTrackingWithDefaults() *Tracking`

NewTrackingWithDefaults instantiates a new Tracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *Tracking) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *Tracking) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *Tracking) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *Tracking) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierService

`func (o *Tracking) GetCarrierService() string`

GetCarrierService returns the CarrierService field if non-nil, zero value otherwise.

### GetCarrierServiceOk

`func (o *Tracking) GetCarrierServiceOk() (*string, bool)`

GetCarrierServiceOk returns a tuple with the CarrierService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierService

`func (o *Tracking) SetCarrierService(v string)`

SetCarrierService sets CarrierService field to given value.

### HasCarrierService

`func (o *Tracking) HasCarrierService() bool`

HasCarrierService returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *Tracking) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *Tracking) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *Tracking) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *Tracking) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *Tracking) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *Tracking) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *Tracking) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *Tracking) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


