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

// checks if the FcAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FcAttribute{}

// FcAttribute struct for FcAttribute
type FcAttribute struct {
	// Unique Id for the fulfillment center attribute
	Id *int32 `json:"id,omitempty"`
	// Name of the attribute.
	Name NullableString `json:"name,omitempty"`
}

// NewFcAttribute instantiates a new FcAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcAttribute() *FcAttribute {
	this := FcAttribute{}
	return &this
}

// NewFcAttributeWithDefaults instantiates a new FcAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcAttributeWithDefaults() *FcAttribute {
	this := FcAttribute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FcAttribute) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcAttribute) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FcAttribute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FcAttribute) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcAttribute) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FcAttribute) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FcAttribute) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *FcAttribute) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FcAttribute) UnsetName() {
	o.Name.Unset()
}

func (o FcAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FcAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableFcAttribute struct {
	value *FcAttribute
	isSet bool
}

func (v NullableFcAttribute) Get() *FcAttribute {
	return v.value
}

func (v *NullableFcAttribute) Set(val *FcAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableFcAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableFcAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcAttribute(val *FcAttribute) *NullableFcAttribute {
	return &NullableFcAttribute{value: val, isSet: true}
}

func (v NullableFcAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
