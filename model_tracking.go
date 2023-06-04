/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"time"
)

// checks if the Tracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tracking{}

// Tracking Tracking information for a shipment
type Tracking struct {
	// Bill of lading (BOL) number. Document acknowledging receipt of cargo for shipment.
	Bol *string `json:"bol,omitempty"`
	// Carrier of the shipment
	Carrier *string `json:"carrier,omitempty"`
	// The carrier's service which was used for this shipment
	CarrierService *string `json:"carrier_service,omitempty"`
	// Tracking number used for freight carriers
	ProNumber *string `json:"pro_number,omitempty"`
	// Standard Carrier Alpha Code (SCAC). Unique 2-4 letter code used to identify transportation companies.
	Scac *string `json:"scac,omitempty"`
	// Date freight was shipped on
	ShippingDate NullableTime `json:"shipping_date,omitempty"`
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

// GetBol returns the Bol field value if set, zero value otherwise.
func (o *Tracking) GetBol() string {
	if o == nil || IsNil(o.Bol) {
		var ret string
		return ret
	}
	return *o.Bol
}

// GetBolOk returns a tuple with the Bol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetBolOk() (*string, bool) {
	if o == nil || IsNil(o.Bol) {
		return nil, false
	}
	return o.Bol, true
}

// HasBol returns a boolean if a field has been set.
func (o *Tracking) HasBol() bool {
	if o != nil && !IsNil(o.Bol) {
		return true
	}

	return false
}

// SetBol gets a reference to the given string and assigns it to the Bol field.
func (o *Tracking) SetBol(v string) {
	o.Bol = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *Tracking) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *Tracking) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
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
	if o == nil || IsNil(o.CarrierService) {
		var ret string
		return ret
	}
	return *o.CarrierService
}

// GetCarrierServiceOk returns a tuple with the CarrierService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetCarrierServiceOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierService) {
		return nil, false
	}
	return o.CarrierService, true
}

// HasCarrierService returns a boolean if a field has been set.
func (o *Tracking) HasCarrierService() bool {
	if o != nil && !IsNil(o.CarrierService) {
		return true
	}

	return false
}

// SetCarrierService gets a reference to the given string and assigns it to the CarrierService field.
func (o *Tracking) SetCarrierService(v string) {
	o.CarrierService = &v
}

// GetProNumber returns the ProNumber field value if set, zero value otherwise.
func (o *Tracking) GetProNumber() string {
	if o == nil || IsNil(o.ProNumber) {
		var ret string
		return ret
	}
	return *o.ProNumber
}

// GetProNumberOk returns a tuple with the ProNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetProNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ProNumber) {
		return nil, false
	}
	return o.ProNumber, true
}

// HasProNumber returns a boolean if a field has been set.
func (o *Tracking) HasProNumber() bool {
	if o != nil && !IsNil(o.ProNumber) {
		return true
	}

	return false
}

// SetProNumber gets a reference to the given string and assigns it to the ProNumber field.
func (o *Tracking) SetProNumber(v string) {
	o.ProNumber = &v
}

// GetScac returns the Scac field value if set, zero value otherwise.
func (o *Tracking) GetScac() string {
	if o == nil || IsNil(o.Scac) {
		var ret string
		return ret
	}
	return *o.Scac
}

// GetScacOk returns a tuple with the Scac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetScacOk() (*string, bool) {
	if o == nil || IsNil(o.Scac) {
		return nil, false
	}
	return o.Scac, true
}

// HasScac returns a boolean if a field has been set.
func (o *Tracking) HasScac() bool {
	if o != nil && !IsNil(o.Scac) {
		return true
	}

	return false
}

// SetScac gets a reference to the given string and assigns it to the Scac field.
func (o *Tracking) SetScac(v string) {
	o.Scac = &v
}

// GetShippingDate returns the ShippingDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tracking) GetShippingDate() time.Time {
	if o == nil || IsNil(o.ShippingDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ShippingDate.Get()
}

// GetShippingDateOk returns a tuple with the ShippingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tracking) GetShippingDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingDate.Get(), o.ShippingDate.IsSet()
}

// HasShippingDate returns a boolean if a field has been set.
func (o *Tracking) HasShippingDate() bool {
	if o != nil && o.ShippingDate.IsSet() {
		return true
	}

	return false
}

// SetShippingDate gets a reference to the given NullableTime and assigns it to the ShippingDate field.
func (o *Tracking) SetShippingDate(v time.Time) {
	o.ShippingDate.Set(&v)
}
// SetShippingDateNil sets the value for ShippingDate to be an explicit nil
func (o *Tracking) SetShippingDateNil() {
	o.ShippingDate.Set(nil)
}

// UnsetShippingDate ensures that no value is present for ShippingDate, not even an explicit nil
func (o *Tracking) UnsetShippingDate() {
	o.ShippingDate.Unset()
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *Tracking) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *Tracking) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
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
	if o == nil || IsNil(o.TrackingUrl) {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingUrl) {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *Tracking) HasTrackingUrl() bool {
	if o != nil && !IsNil(o.TrackingUrl) {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *Tracking) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

func (o Tracking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bol) {
		toSerialize["bol"] = o.Bol
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.CarrierService) {
		toSerialize["carrier_service"] = o.CarrierService
	}
	if !IsNil(o.ProNumber) {
		toSerialize["pro_number"] = o.ProNumber
	}
	if !IsNil(o.Scac) {
		toSerialize["scac"] = o.Scac
	}
	if o.ShippingDate.IsSet() {
		toSerialize["shipping_date"] = o.ShippingDate.Get()
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.TrackingUrl) {
		toSerialize["tracking_url"] = o.TrackingUrl
	}
	return toSerialize, nil
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


