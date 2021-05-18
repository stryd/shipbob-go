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

// LocationRegionViewModel struct for LocationRegionViewModel
type LocationRegionViewModel struct {
	// Unique Id for the location region
	Id *int32 `json:"id,omitempty"`
	// Name of the region the location is in.
	Name NullableString `json:"name,omitempty"`
}

// NewLocationRegionViewModel instantiates a new LocationRegionViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRegionViewModel() *LocationRegionViewModel {
	this := LocationRegionViewModel{}
	return &this
}

// NewLocationRegionViewModelWithDefaults instantiates a new LocationRegionViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRegionViewModelWithDefaults() *LocationRegionViewModel {
	this := LocationRegionViewModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationRegionViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRegionViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationRegionViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LocationRegionViewModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRegionViewModel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRegionViewModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LocationRegionViewModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LocationRegionViewModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LocationRegionViewModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LocationRegionViewModel) UnsetName() {
	o.Name.Unset()
}

func (o LocationRegionViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLocationRegionViewModel struct {
	value *LocationRegionViewModel
	isSet bool
}

func (v NullableLocationRegionViewModel) Get() *LocationRegionViewModel {
	return v.value
}

func (v *NullableLocationRegionViewModel) Set(val *LocationRegionViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRegionViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRegionViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRegionViewModel(val *LocationRegionViewModel) *NullableLocationRegionViewModel {
	return &NullableLocationRegionViewModel{value: val, isSet: true}
}

func (v NullableLocationRegionViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRegionViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

