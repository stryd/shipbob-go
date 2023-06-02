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

// checks if the ReturnFulfillmentCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnFulfillmentCenter{}

// ReturnFulfillmentCenter Information about a fulfillment center
type ReturnFulfillmentCenter struct {
	// Unique identifier of the fulfillment center
	Id int32 `json:"id"`
	// Name of the fulfillment center
	Name NullableString `json:"name,omitempty"`
}

// NewReturnFulfillmentCenter instantiates a new ReturnFulfillmentCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnFulfillmentCenter(id int32) *ReturnFulfillmentCenter {
	this := ReturnFulfillmentCenter{}
	this.Id = id
	return &this
}

// NewReturnFulfillmentCenterWithDefaults instantiates a new ReturnFulfillmentCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnFulfillmentCenterWithDefaults() *ReturnFulfillmentCenter {
	this := ReturnFulfillmentCenter{}
	return &this
}

// GetId returns the Id field value
func (o *ReturnFulfillmentCenter) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReturnFulfillmentCenter) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReturnFulfillmentCenter) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnFulfillmentCenter) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnFulfillmentCenter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReturnFulfillmentCenter) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReturnFulfillmentCenter) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ReturnFulfillmentCenter) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReturnFulfillmentCenter) UnsetName() {
	o.Name.Unset()
}

func (o ReturnFulfillmentCenter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnFulfillmentCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableReturnFulfillmentCenter struct {
	value *ReturnFulfillmentCenter
	isSet bool
}

func (v NullableReturnFulfillmentCenter) Get() *ReturnFulfillmentCenter {
	return v.value
}

func (v *NullableReturnFulfillmentCenter) Set(val *ReturnFulfillmentCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnFulfillmentCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnFulfillmentCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnFulfillmentCenter(val *ReturnFulfillmentCenter) *NullableReturnFulfillmentCenter {
	return &NullableReturnFulfillmentCenter{value: val, isSet: true}
}

func (v NullableReturnFulfillmentCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnFulfillmentCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
