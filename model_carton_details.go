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

// checks if the CartonDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartonDetails{}

// CartonDetails struct for CartonDetails
type CartonDetails struct {
	// List of what is packed in this carton
	Products []ShipmentProduct `json:"products,omitempty"`
}

// NewCartonDetails instantiates a new CartonDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartonDetails() *CartonDetails {
	this := CartonDetails{}
	return &this
}

// NewCartonDetailsWithDefaults instantiates a new CartonDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartonDetailsWithDefaults() *CartonDetails {
	this := CartonDetails{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *CartonDetails) GetProducts() []ShipmentProduct {
	if o == nil || IsNil(o.Products) {
		var ret []ShipmentProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartonDetails) GetProductsOk() ([]ShipmentProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *CartonDetails) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ShipmentProduct and assigns it to the Products field.
func (o *CartonDetails) SetProducts(v []ShipmentProduct) {
	o.Products = v
}

func (o CartonDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartonDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableCartonDetails struct {
	value *CartonDetails
	isSet bool
}

func (v NullableCartonDetails) Get() *CartonDetails {
	return v.value
}

func (v *NullableCartonDetails) Set(val *CartonDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCartonDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCartonDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartonDetails(val *CartonDetails) *NullableCartonDetails {
	return &NullableCartonDetails{value: val, isSet: true}
}

func (v NullableCartonDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartonDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


