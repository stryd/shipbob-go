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

// checks if the LocationRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRegion{}

// LocationRegion struct for LocationRegion
type LocationRegion struct {
	// Unique Id for the location region
	Id *int32 `json:"id,omitempty"`
	// Name of the region the location is in.
	Name NullableString `json:"name,omitempty"`
}

// NewLocationRegion instantiates a new LocationRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRegion() *LocationRegion {
	this := LocationRegion{}
	return &this
}

// NewLocationRegionWithDefaults instantiates a new LocationRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRegionWithDefaults() *LocationRegion {
	this := LocationRegion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationRegion) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegion) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationRegion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LocationRegion) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRegion) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LocationRegion) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LocationRegion) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *LocationRegion) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LocationRegion) UnsetName() {
	o.Name.Unset()
}

func (o LocationRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableLocationRegion struct {
	value *LocationRegion
	isSet bool
}

func (v NullableLocationRegion) Get() *LocationRegion {
	return v.value
}

func (v *NullableLocationRegion) Set(val *LocationRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRegion(val *LocationRegion) *NullableLocationRegion {
	return &NullableLocationRegion{value: val, isSet: true}
}

func (v NullableLocationRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
