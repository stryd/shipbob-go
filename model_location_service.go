/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
)

// checks if the LocationService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationService{}

// LocationService struct for LocationService
type LocationService struct {
	Address *LocationAddress `json:"address,omitempty"`
	// Indicates if the user is authorized to access this service at the location
	Enabled *bool `json:"enabled,omitempty"`
	ServiceType *LocationServiceTypeEnum `json:"service_type,omitempty"`
}

// NewLocationService instantiates a new LocationService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationService() *LocationService {
	this := LocationService{}
	return &this
}

// NewLocationServiceWithDefaults instantiates a new LocationService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationServiceWithDefaults() *LocationService {
	this := LocationService{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LocationService) GetAddress() LocationAddress {
	if o == nil || IsNil(o.Address) {
		var ret LocationAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationService) GetAddressOk() (*LocationAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LocationService) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given LocationAddress and assigns it to the Address field.
func (o *LocationService) SetAddress(v LocationAddress) {
	o.Address = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LocationService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LocationService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LocationService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *LocationService) GetServiceType() LocationServiceTypeEnum {
	if o == nil || IsNil(o.ServiceType) {
		var ret LocationServiceTypeEnum
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationService) GetServiceTypeOk() (*LocationServiceTypeEnum, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *LocationService) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given LocationServiceTypeEnum and assigns it to the ServiceType field.
func (o *LocationService) SetServiceType(v LocationServiceTypeEnum) {
	o.ServiceType = &v
}

func (o LocationService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	return toSerialize, nil
}

type NullableLocationService struct {
	value *LocationService
	isSet bool
}

func (v NullableLocationService) Get() *LocationService {
	return v.value
}

func (v *NullableLocationService) Set(val *LocationService) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationService) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationService(val *LocationService) *NullableLocationService {
	return &NullableLocationService{value: val, isSet: true}
}

func (v NullableLocationService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


