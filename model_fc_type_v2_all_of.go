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

// checks if the FcTypeV2AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FcTypeV2AllOf{}

// FcTypeV2AllOf struct for FcTypeV2AllOf
type FcTypeV2AllOf struct {
	IsChild              *bool              `json:"is_child,omitempty"`
	OrganizationRoles    []OrganizationRole `json:"organization_roles,omitempty"`
	OrganizationTypeId   *string            `json:"organization_type_id,omitempty"`
	OrganizationTypeName NullableString     `json:"organization_type_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcTypeV2AllOf FcTypeV2AllOf

// NewFcTypeV2AllOf instantiates a new FcTypeV2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcTypeV2AllOf() *FcTypeV2AllOf {
	this := FcTypeV2AllOf{}
	return &this
}

// NewFcTypeV2AllOfWithDefaults instantiates a new FcTypeV2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcTypeV2AllOfWithDefaults() *FcTypeV2AllOf {
	this := FcTypeV2AllOf{}
	return &this
}

// GetIsChild returns the IsChild field value if set, zero value otherwise.
func (o *FcTypeV2AllOf) GetIsChild() bool {
	if o == nil || IsNil(o.IsChild) {
		var ret bool
		return ret
	}
	return *o.IsChild
}

// GetIsChildOk returns a tuple with the IsChild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcTypeV2AllOf) GetIsChildOk() (*bool, bool) {
	if o == nil || IsNil(o.IsChild) {
		return nil, false
	}
	return o.IsChild, true
}

// HasIsChild returns a boolean if a field has been set.
func (o *FcTypeV2AllOf) HasIsChild() bool {
	if o != nil && !IsNil(o.IsChild) {
		return true
	}

	return false
}

// SetIsChild gets a reference to the given bool and assigns it to the IsChild field.
func (o *FcTypeV2AllOf) SetIsChild(v bool) {
	o.IsChild = &v
}

// GetOrganizationRoles returns the OrganizationRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcTypeV2AllOf) GetOrganizationRoles() []OrganizationRole {
	if o == nil {
		var ret []OrganizationRole
		return ret
	}
	return o.OrganizationRoles
}

// GetOrganizationRolesOk returns a tuple with the OrganizationRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcTypeV2AllOf) GetOrganizationRolesOk() ([]OrganizationRole, bool) {
	if o == nil || IsNil(o.OrganizationRoles) {
		return nil, false
	}
	return o.OrganizationRoles, true
}

// HasOrganizationRoles returns a boolean if a field has been set.
func (o *FcTypeV2AllOf) HasOrganizationRoles() bool {
	if o != nil && IsNil(o.OrganizationRoles) {
		return true
	}

	return false
}

// SetOrganizationRoles gets a reference to the given []OrganizationRole and assigns it to the OrganizationRoles field.
func (o *FcTypeV2AllOf) SetOrganizationRoles(v []OrganizationRole) {
	o.OrganizationRoles = v
}

// GetOrganizationTypeId returns the OrganizationTypeId field value if set, zero value otherwise.
func (o *FcTypeV2AllOf) GetOrganizationTypeId() string {
	if o == nil || IsNil(o.OrganizationTypeId) {
		var ret string
		return ret
	}
	return *o.OrganizationTypeId
}

// GetOrganizationTypeIdOk returns a tuple with the OrganizationTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcTypeV2AllOf) GetOrganizationTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationTypeId) {
		return nil, false
	}
	return o.OrganizationTypeId, true
}

// HasOrganizationTypeId returns a boolean if a field has been set.
func (o *FcTypeV2AllOf) HasOrganizationTypeId() bool {
	if o != nil && !IsNil(o.OrganizationTypeId) {
		return true
	}

	return false
}

// SetOrganizationTypeId gets a reference to the given string and assigns it to the OrganizationTypeId field.
func (o *FcTypeV2AllOf) SetOrganizationTypeId(v string) {
	o.OrganizationTypeId = &v
}

// GetOrganizationTypeName returns the OrganizationTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcTypeV2AllOf) GetOrganizationTypeName() string {
	if o == nil || IsNil(o.OrganizationTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationTypeName.Get()
}

// GetOrganizationTypeNameOk returns a tuple with the OrganizationTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcTypeV2AllOf) GetOrganizationTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationTypeName.Get(), o.OrganizationTypeName.IsSet()
}

// HasOrganizationTypeName returns a boolean if a field has been set.
func (o *FcTypeV2AllOf) HasOrganizationTypeName() bool {
	if o != nil && o.OrganizationTypeName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationTypeName gets a reference to the given NullableString and assigns it to the OrganizationTypeName field.
func (o *FcTypeV2AllOf) SetOrganizationTypeName(v string) {
	o.OrganizationTypeName.Set(&v)
}

// SetOrganizationTypeNameNil sets the value for OrganizationTypeName to be an explicit nil
func (o *FcTypeV2AllOf) SetOrganizationTypeNameNil() {
	o.OrganizationTypeName.Set(nil)
}

// UnsetOrganizationTypeName ensures that no value is present for OrganizationTypeName, not even an explicit nil
func (o *FcTypeV2AllOf) UnsetOrganizationTypeName() {
	o.OrganizationTypeName.Unset()
}

func (o FcTypeV2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FcTypeV2AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsChild) {
		toSerialize["is_child"] = o.IsChild
	}
	if o.OrganizationRoles != nil {
		toSerialize["organization_roles"] = o.OrganizationRoles
	}
	if !IsNil(o.OrganizationTypeId) {
		toSerialize["organization_type_id"] = o.OrganizationTypeId
	}
	if o.OrganizationTypeName.IsSet() {
		toSerialize["organization_type_name"] = o.OrganizationTypeName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FcTypeV2AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcTypeV2AllOf := _FcTypeV2AllOf{}

	if err = json.Unmarshal(bytes, &varFcTypeV2AllOf); err == nil {
		*o = FcTypeV2AllOf(varFcTypeV2AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "is_child")
		delete(additionalProperties, "organization_roles")
		delete(additionalProperties, "organization_type_id")
		delete(additionalProperties, "organization_type_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcTypeV2AllOf struct {
	value *FcTypeV2AllOf
	isSet bool
}

func (v NullableFcTypeV2AllOf) Get() *FcTypeV2AllOf {
	return v.value
}

func (v *NullableFcTypeV2AllOf) Set(val *FcTypeV2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcTypeV2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcTypeV2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcTypeV2AllOf(val *FcTypeV2AllOf) *NullableFcTypeV2AllOf {
	return &NullableFcTypeV2AllOf{value: val, isSet: true}
}

func (v NullableFcTypeV2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcTypeV2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
