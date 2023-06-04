# Tracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bol** | Pointer to **string** | Bill of lading (BOL) number. Document acknowledging receipt of cargo for shipment. | [optional] 
**Carrier** | Pointer to **string** | Carrier of the shipment | [optional] 
**CarrierService** | Pointer to **string** | The carrier&#39;s service which was used for this shipment | [optional] 
**ProNumber** | Pointer to **string** | Tracking number used for freight carriers | [optional] 
**Scac** | Pointer to **string** | Standard Carrier Alpha Code (SCAC). Unique 2-4 letter code used to identify transportation companies. | [optional] 
**ShippingDate** | Pointer to **NullableTime** | Date freight was shipped on | [optional] 
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

### GetBol

`func (o *Tracking) GetBol() string`

GetBol returns the Bol field if non-nil, zero value otherwise.

### GetBolOk

`func (o *Tracking) GetBolOk() (*string, bool)`

GetBolOk returns a tuple with the Bol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBol

`func (o *Tracking) SetBol(v string)`

SetBol sets Bol field to given value.

### HasBol

`func (o *Tracking) HasBol() bool`

HasBol returns a boolean if a field has been set.

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

### GetProNumber

`func (o *Tracking) GetProNumber() string`

GetProNumber returns the ProNumber field if non-nil, zero value otherwise.

### GetProNumberOk

`func (o *Tracking) GetProNumberOk() (*string, bool)`

GetProNumberOk returns a tuple with the ProNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProNumber

`func (o *Tracking) SetProNumber(v string)`

SetProNumber sets ProNumber field to given value.

### HasProNumber

`func (o *Tracking) HasProNumber() bool`

HasProNumber returns a boolean if a field has been set.

### GetScac

`func (o *Tracking) GetScac() string`

GetScac returns the Scac field if non-nil, zero value otherwise.

### GetScacOk

`func (o *Tracking) GetScacOk() (*string, bool)`

GetScacOk returns a tuple with the Scac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScac

`func (o *Tracking) SetScac(v string)`

SetScac sets Scac field to given value.

### HasScac

`func (o *Tracking) HasScac() bool`

HasScac returns a boolean if a field has been set.

### GetShippingDate

`func (o *Tracking) GetShippingDate() time.Time`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *Tracking) GetShippingDateOk() (*time.Time, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *Tracking) SetShippingDate(v time.Time)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *Tracking) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### SetShippingDateNil

`func (o *Tracking) SetShippingDateNil(b bool)`

 SetShippingDateNil sets the value for ShippingDate to be an explicit nil

### UnsetShippingDate
`func (o *Tracking) UnsetShippingDate()`

UnsetShippingDate ensures that no value is present for ShippingDate, not even an explicit nil
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


