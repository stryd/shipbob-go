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

// checks if the InternalLocationAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalLocationAllOf{}

// InternalLocationAllOf struct for InternalLocationAllOf
type InternalLocationAllOf struct {
	FulfillmentCenterAttributes []FcAttribute                                      `json:"fulfillment_center_attributes,omitempty"`
	FulfillmentCenterType       NullableInternalLocationAllOfFulfillmentCenterType `json:"fulfillment_center_type,omitempty"`
	IsEnabledForNewUser         *bool                                              `json:"is_enabled_for_new_user,omitempty"`
	IsExternal                  *bool                                              `json:"is_external,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _InternalLocationAllOf InternalLocationAllOf

// NewInternalLocationAllOf instantiates a new InternalLocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalLocationAllOf() *InternalLocationAllOf {
	this := InternalLocationAllOf{}
	return &this
}

// NewInternalLocationAllOfWithDefaults instantiates a new InternalLocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalLocationAllOfWithDefaults() *InternalLocationAllOf {
	this := InternalLocationAllOf{}
	return &this
}

// GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationAllOf) GetFulfillmentCenterAttributes() []FcAttribute {
	if o == nil {
		var ret []FcAttribute
		return ret
	}
	return o.FulfillmentCenterAttributes
}

// GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationAllOf) GetFulfillmentCenterAttributesOk() ([]FcAttribute, bool) {
	if o == nil || IsNil(o.FulfillmentCenterAttributes) {
		return nil, false
	}
	return o.FulfillmentCenterAttributes, true
}

// HasFulfillmentCenterAttributes returns a boolean if a field has been set.
func (o *InternalLocationAllOf) HasFulfillmentCenterAttributes() bool {
	if o != nil && IsNil(o.FulfillmentCenterAttributes) {
		return true
	}

	return false
}

// SetFulfillmentCenterAttributes gets a reference to the given []FcAttribute and assigns it to the FulfillmentCenterAttributes field.
func (o *InternalLocationAllOf) SetFulfillmentCenterAttributes(v []FcAttribute) {
	o.FulfillmentCenterAttributes = v
}

// GetFulfillmentCenterType returns the FulfillmentCenterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationAllOf) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType {
	if o == nil || IsNil(o.FulfillmentCenterType.Get()) {
		var ret InternalLocationAllOfFulfillmentCenterType
		return ret
	}
	return *o.FulfillmentCenterType.Get()
}

// GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationAllOf) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfillmentCenterType.Get(), o.FulfillmentCenterType.IsSet()
}

// HasFulfillmentCenterType returns a boolean if a field has been set.
func (o *InternalLocationAllOf) HasFulfillmentCenterType() bool {
	if o != nil && o.FulfillmentCenterType.IsSet() {
		return true
	}

	return false
}

// SetFulfillmentCenterType gets a reference to the given NullableInternalLocationAllOfFulfillmentCenterType and assigns it to the FulfillmentCenterType field.
func (o *InternalLocationAllOf) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType) {
	o.FulfillmentCenterType.Set(&v)
}

// SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil
func (o *InternalLocationAllOf) SetFulfillmentCenterTypeNil() {
	o.FulfillmentCenterType.Set(nil)
}

// UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
func (o *InternalLocationAllOf) UnsetFulfillmentCenterType() {
	o.FulfillmentCenterType.Unset()
}

// GetIsEnabledForNewUser returns the IsEnabledForNewUser field value if set, zero value otherwise.
func (o *InternalLocationAllOf) GetIsEnabledForNewUser() bool {
	if o == nil || IsNil(o.IsEnabledForNewUser) {
		var ret bool
		return ret
	}
	return *o.IsEnabledForNewUser
}

// GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationAllOf) GetIsEnabledForNewUserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabledForNewUser) {
		return nil, false
	}
	return o.IsEnabledForNewUser, true
}

// HasIsEnabledForNewUser returns a boolean if a field has been set.
func (o *InternalLocationAllOf) HasIsEnabledForNewUser() bool {
	if o != nil && !IsNil(o.IsEnabledForNewUser) {
		return true
	}

	return false
}

// SetIsEnabledForNewUser gets a reference to the given bool and assigns it to the IsEnabledForNewUser field.
func (o *InternalLocationAllOf) SetIsEnabledForNewUser(v bool) {
	o.IsEnabledForNewUser = &v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *InternalLocationAllOf) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationAllOf) GetIsExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternal) {
		return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *InternalLocationAllOf) HasIsExternal() bool {
	if o != nil && !IsNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *InternalLocationAllOf) SetIsExternal(v bool) {
	o.IsExternal = &v
}

func (o InternalLocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalLocationAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FulfillmentCenterAttributes != nil {
		toSerialize["fulfillment_center_attributes"] = o.FulfillmentCenterAttributes
	}
	if o.FulfillmentCenterType.IsSet() {
		toSerialize["fulfillment_center_type"] = o.FulfillmentCenterType.Get()
	}
	if !IsNil(o.IsEnabledForNewUser) {
		toSerialize["is_enabled_for_new_user"] = o.IsEnabledForNewUser
	}
	if !IsNil(o.IsExternal) {
		toSerialize["is_external"] = o.IsExternal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternalLocationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInternalLocationAllOf := _InternalLocationAllOf{}

	if err = json.Unmarshal(bytes, &varInternalLocationAllOf); err == nil {
		*o = InternalLocationAllOf(varInternalLocationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fulfillment_center_attributes")
		delete(additionalProperties, "fulfillment_center_type")
		delete(additionalProperties, "is_enabled_for_new_user")
		delete(additionalProperties, "is_external")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalLocationAllOf struct {
	value *InternalLocationAllOf
	isSet bool
}

func (v NullableInternalLocationAllOf) Get() *InternalLocationAllOf {
	return v.value
}

func (v *NullableInternalLocationAllOf) Set(val *InternalLocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalLocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalLocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalLocationAllOf(val *InternalLocationAllOf) *NullableInternalLocationAllOf {
	return &NullableInternalLocationAllOf{value: val, isSet: true}
}

func (v NullableInternalLocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalLocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
