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

// LocationServiceViewModel struct for LocationServiceViewModel
type LocationServiceViewModel struct {
	Address *LocationAddressViewModel `json:"address,omitempty"`
	// Indicates if the user is authorized to access this service at the location
	Enabled *bool `json:"enabled,omitempty"`
	ServiceType *LocationServiceTypeEnum `json:"service_type,omitempty"`
}

// NewLocationServiceViewModel instantiates a new LocationServiceViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationServiceViewModel() *LocationServiceViewModel {
	this := LocationServiceViewModel{}
	return &this
}

// NewLocationServiceViewModelWithDefaults instantiates a new LocationServiceViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationServiceViewModelWithDefaults() *LocationServiceViewModel {
	this := LocationServiceViewModel{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LocationServiceViewModel) GetAddress() LocationAddressViewModel {
	if o == nil || o.Address == nil {
		var ret LocationAddressViewModel
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationServiceViewModel) GetAddressOk() (*LocationAddressViewModel, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LocationServiceViewModel) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given LocationAddressViewModel and assigns it to the Address field.
func (o *LocationServiceViewModel) SetAddress(v LocationAddressViewModel) {
	o.Address = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LocationServiceViewModel) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationServiceViewModel) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LocationServiceViewModel) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LocationServiceViewModel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *LocationServiceViewModel) GetServiceType() LocationServiceTypeEnum {
	if o == nil || o.ServiceType == nil {
		var ret LocationServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationServiceViewModel) GetServiceTypeOk() (*LocationServiceTypeEnum, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *LocationServiceViewModel) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LocationServiceTypeEnum and assigns it to the ServiceType field.
func (o *LocationServiceViewModel) SetServiceType(v LocationServiceTypeEnum) {
	o.ServiceType = &v
}

func (o LocationServiceViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableLocationServiceViewModel struct {
	value *LocationServiceViewModel
	isSet bool
}

func (v NullableLocationServiceViewModel) Get() *LocationServiceViewModel {
	return v.value
}

func (v *NullableLocationServiceViewModel) Set(val *LocationServiceViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationServiceViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationServiceViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationServiceViewModel(val *LocationServiceViewModel) *NullableLocationServiceViewModel {
	return &NullableLocationServiceViewModel{value: val, isSet: true}
}

func (v NullableLocationServiceViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationServiceViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

