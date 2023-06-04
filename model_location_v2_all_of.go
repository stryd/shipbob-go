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

// checks if the LocationV2AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationV2AllOf{}

// LocationV2AllOf struct for LocationV2AllOf
type LocationV2AllOf struct {
	OrganizationId *string `json:"organization_id,omitempty"`
	OwnerId NullableString `json:"owner_id,omitempty"`
	ParentId NullableString `json:"parent_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocationV2AllOf LocationV2AllOf

// NewLocationV2AllOf instantiates a new LocationV2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationV2AllOf() *LocationV2AllOf {
	this := LocationV2AllOf{}
	return &this
}

// NewLocationV2AllOfWithDefaults instantiates a new LocationV2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationV2AllOfWithDefaults() *LocationV2AllOf {
	this := LocationV2AllOf{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *LocationV2AllOf) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationV2AllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *LocationV2AllOf) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *LocationV2AllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationV2AllOf) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationV2AllOf) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *LocationV2AllOf) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *LocationV2AllOf) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *LocationV2AllOf) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *LocationV2AllOf) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationV2AllOf) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationV2AllOf) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *LocationV2AllOf) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *LocationV2AllOf) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *LocationV2AllOf) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *LocationV2AllOf) UnsetParentId() {
	o.ParentId.Unset()
}

func (o LocationV2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationV2AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.OwnerId.IsSet() {
		toSerialize["owner_id"] = o.OwnerId.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LocationV2AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLocationV2AllOf := _LocationV2AllOf{}

	if err = json.Unmarshal(bytes, &varLocationV2AllOf); err == nil {
		*o = LocationV2AllOf(varLocationV2AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "parent_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocationV2AllOf struct {
	value *LocationV2AllOf
	isSet bool
}

func (v NullableLocationV2AllOf) Get() *LocationV2AllOf {
	return v.value
}

func (v *NullableLocationV2AllOf) Set(val *LocationV2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationV2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationV2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationV2AllOf(val *LocationV2AllOf) *NullableLocationV2AllOf {
	return &NullableLocationV2AllOf{value: val, isSet: true}
}

func (v NullableLocationV2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationV2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


