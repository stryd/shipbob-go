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

// checks if the Dimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimension{}

// Dimension Information about an inventory item's dimensions
type Dimension struct {
	// Depth in inches of this inventory item
	Depth *float64 `json:"depth,omitempty"`
	// Length in inches of this inventory item
	Length *float64 `json:"length,omitempty"`
	// Weight in ounces of this inventory item
	Weight *float64 `json:"weight,omitempty"`
	// Width in inches of this inventory item
	Width *float64 `json:"width,omitempty"`
}

// NewDimension instantiates a new Dimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimension() *Dimension {
	this := Dimension{}
	return &this
}

// NewDimensionWithDefaults instantiates a new Dimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionWithDefaults() *Dimension {
	this := Dimension{}
	return &this
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *Dimension) GetDepth() float64 {
	if o == nil || IsNil(o.Depth) {
		var ret float64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetDepthOk() (*float64, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *Dimension) HasDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given float64 and assigns it to the Depth field.
func (o *Dimension) SetDepth(v float64) {
	o.Depth = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Dimension) GetLength() float64 {
	if o == nil || IsNil(o.Length) {
		var ret float64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Dimension) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *Dimension) SetLength(v float64) {
	o.Length = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Dimension) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Dimension) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *Dimension) SetWeight(v float64) {
	o.Weight = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Dimension) GetWidth() float64 {
	if o == nil || IsNil(o.Width) {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimension) GetWidthOk() (*float64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Dimension) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *Dimension) SetWidth(v float64) {
	o.Width = &v
}

func (o Dimension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableDimension struct {
	value *Dimension
	isSet bool
}

func (v NullableDimension) Get() *Dimension {
	return v.value
}

func (v *NullableDimension) Set(val *Dimension) {
	v.value = val
	v.isSet = true
}

func (v NullableDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimension(val *Dimension) *NullableDimension {
	return &NullableDimension{value: val, isSet: true}
}

func (v NullableDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
