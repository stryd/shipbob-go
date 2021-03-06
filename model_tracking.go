/*
 * ShipBob Developer API
 *
 * ShipBob Developer API Documentation
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
)

// Tracking Tracking information for a shipment
type Tracking struct {
	// Carrier of the shipment
	Carrier *string `json:"carrier,omitempty"`
	// The carrier's service which was used for this shipment
	CarrierService *string `json:"carrier_service,omitempty"`
	// Tracking number of the shipment
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// URL to the website where a shipment can be tracked
	TrackingUrl *string `json:"tracking_url,omitempty"`
}

// NewTracking instantiates a new Tracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracking() *Tracking {
	this := Tracking{}
	return &this
}

// NewTrackingWithDefaults instantiates a new Tracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingWithDefaults() *Tracking {
	this := Tracking{}
	return &this
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *Tracking) GetCarrier() string {
	if o == nil || o.Carrier == nil {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetCarrierOk() (*string, bool) {
	if o == nil || o.Carrier == nil {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *Tracking) HasCarrier() bool {
	if o != nil && o.Carrier != nil {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *Tracking) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCarrierService returns the CarrierService field value if set, zero value otherwise.
func (o *Tracking) GetCarrierService() string {
	if o == nil || o.CarrierService == nil {
		var ret string
		return ret
	}
	return *o.CarrierService
}

// GetCarrierServiceOk returns a tuple with the CarrierService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetCarrierServiceOk() (*string, bool) {
	if o == nil || o.CarrierService == nil {
		return nil, false
	}
	return o.CarrierService, true
}

// HasCarrierService returns a boolean if a field has been set.
func (o *Tracking) HasCarrierService() bool {
	if o != nil && o.CarrierService != nil {
		return true
	}

	return false
}

// SetCarrierService gets a reference to the given string and assigns it to the CarrierService field.
func (o *Tracking) SetCarrierService(v string) {
	o.CarrierService = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *Tracking) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Tracking) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *Tracking) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *Tracking) GetTrackingUrl() string {
	if o == nil || o.TrackingUrl == nil {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetTrackingUrlOk() (*string, bool) {
	if o == nil || o.TrackingUrl == nil {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *Tracking) HasTrackingUrl() bool {
	if o != nil && o.TrackingUrl != nil {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *Tracking) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

func (o Tracking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Carrier != nil {
		toSerialize["carrier"] = o.Carrier
	}
	if o.CarrierService != nil {
		toSerialize["carrier_service"] = o.CarrierService
	}
	if o.TrackingNumber != nil {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if o.TrackingUrl != nil {
		toSerialize["tracking_url"] = o.TrackingUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTracking struct {
	value *Tracking
	isSet bool
}

func (v NullableTracking) Get() *Tracking {
	return v.value
}

func (v *NullableTracking) Set(val *Tracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracking(val *Tracking) *NullableTracking {
	return &NullableTracking{value: val, isSet: true}
}

func (v NullableTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
