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

// checks if the ShipMethodDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipMethodDetail{}

// ShipMethodDetail struct for ShipMethodDetail
type ShipMethodDetail struct {
	// Indicates if the shipping method is active
	Active *bool `json:"active,omitempty"`
	// Indicates the shipping method is a ShipBob default shipping method.
	Default *bool `json:"default,omitempty"`
	// Unique id for shipping method.
	Id *int32 `json:"id,omitempty"`
	// Name of the ship method as selected by the merchant and saved in ShipBobâs database (i.e. âgroundâ). Corresponds to the shipping_method field in the Orders API.
	Name         *string             `json:"name,omitempty"`
	ServiceLevel *ServiceLevelDetail `json:"service_level,omitempty"`
}

// NewShipMethodDetail instantiates a new ShipMethodDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipMethodDetail() *ShipMethodDetail {
	this := ShipMethodDetail{}
	return &this
}

// NewShipMethodDetailWithDefaults instantiates a new ShipMethodDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipMethodDetailWithDefaults() *ShipMethodDetail {
	this := ShipMethodDetail{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ShipMethodDetail) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMethodDetail) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ShipMethodDetail) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ShipMethodDetail) SetActive(v bool) {
	o.Active = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ShipMethodDetail) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMethodDetail) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ShipMethodDetail) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *ShipMethodDetail) SetDefault(v bool) {
	o.Default = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipMethodDetail) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMethodDetail) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipMethodDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipMethodDetail) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipMethodDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMethodDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipMethodDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipMethodDetail) SetName(v string) {
	o.Name = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *ShipMethodDetail) GetServiceLevel() ServiceLevelDetail {
	if o == nil || IsNil(o.ServiceLevel) {
		var ret ServiceLevelDetail
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipMethodDetail) GetServiceLevelOk() (*ServiceLevelDetail, bool) {
	if o == nil || IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *ShipMethodDetail) HasServiceLevel() bool {
	if o != nil && !IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given ServiceLevelDetail and assigns it to the ServiceLevel field.
func (o *ShipMethodDetail) SetServiceLevel(v ServiceLevelDetail) {
	o.ServiceLevel = &v
}

func (o ShipMethodDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipMethodDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceLevel) {
		toSerialize["service_level"] = o.ServiceLevel
	}
	return toSerialize, nil
}

type NullableShipMethodDetail struct {
	value *ShipMethodDetail
	isSet bool
}

func (v NullableShipMethodDetail) Get() *ShipMethodDetail {
	return v.value
}

func (v *NullableShipMethodDetail) Set(val *ShipMethodDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableShipMethodDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableShipMethodDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipMethodDetail(val *ShipMethodDetail) *NullableShipMethodDetail {
	return &NullableShipMethodDetail{value: val, isSet: true}
}

func (v NullableShipMethodDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipMethodDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
