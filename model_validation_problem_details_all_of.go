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

// checks if the ValidationProblemDetailsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationProblemDetailsAllOf{}

// ValidationProblemDetailsAllOf struct for ValidationProblemDetailsAllOf
type ValidationProblemDetailsAllOf struct {
	Errors               map[string][]string `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidationProblemDetailsAllOf ValidationProblemDetailsAllOf

// NewValidationProblemDetailsAllOf instantiates a new ValidationProblemDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationProblemDetailsAllOf() *ValidationProblemDetailsAllOf {
	this := ValidationProblemDetailsAllOf{}
	return &this
}

// NewValidationProblemDetailsAllOfWithDefaults instantiates a new ValidationProblemDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationProblemDetailsAllOfWithDefaults() *ValidationProblemDetailsAllOf {
	this := ValidationProblemDetailsAllOf{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidationProblemDetailsAllOf) GetErrors() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationProblemDetailsAllOf) GetErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidationProblemDetailsAllOf) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string][]string and assigns it to the Errors field.
func (o *ValidationProblemDetailsAllOf) SetErrors(v map[string][]string) {
	o.Errors = v
}

func (o ValidationProblemDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationProblemDetailsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidationProblemDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varValidationProblemDetailsAllOf := _ValidationProblemDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varValidationProblemDetailsAllOf); err == nil {
		*o = ValidationProblemDetailsAllOf(varValidationProblemDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidationProblemDetailsAllOf struct {
	value *ValidationProblemDetailsAllOf
	isSet bool
}

func (v NullableValidationProblemDetailsAllOf) Get() *ValidationProblemDetailsAllOf {
	return v.value
}

func (v *NullableValidationProblemDetailsAllOf) Set(val *ValidationProblemDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationProblemDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationProblemDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationProblemDetailsAllOf(val *ValidationProblemDetailsAllOf) *NullableValidationProblemDetailsAllOf {
	return &NullableValidationProblemDetailsAllOf{value: val, isSet: true}
}

func (v NullableValidationProblemDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationProblemDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
