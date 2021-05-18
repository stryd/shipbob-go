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

// InventoryDimensionViewModel Information about an inventory item's dimensions
type InventoryDimensionViewModel struct {
	// Depth in inches of this inventory item
	Depth *float64 `json:"depth,omitempty"`
	// Length in inches of this inventory item
	Length *float64 `json:"length,omitempty"`
	// Weight in ounces of this inventory item
	Weight *float64 `json:"weight,omitempty"`
	// Width in inches of this inventory item
	Width *float64 `json:"width,omitempty"`
}

// NewInventoryDimensionViewModel instantiates a new InventoryDimensionViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDimensionViewModel() *InventoryDimensionViewModel {
	this := InventoryDimensionViewModel{}
	return &this
}

// NewInventoryDimensionViewModelWithDefaults instantiates a new InventoryDimensionViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDimensionViewModelWithDefaults() *InventoryDimensionViewModel {
	this := InventoryDimensionViewModel{}
	return &this
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *InventoryDimensionViewModel) GetDepth() float64 {
	if o == nil || o.Depth == nil {
		var ret float64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDimensionViewModel) GetDepthOk() (*float64, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *InventoryDimensionViewModel) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given float64 and assigns it to the Depth field.
func (o *InventoryDimensionViewModel) SetDepth(v float64) {
	o.Depth = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *InventoryDimensionViewModel) GetLength() float64 {
	if o == nil || o.Length == nil {
		var ret float64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDimensionViewModel) GetLengthOk() (*float64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *InventoryDimensionViewModel) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *InventoryDimensionViewModel) SetLength(v float64) {
	o.Length = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *InventoryDimensionViewModel) GetWeight() float64 {
	if o == nil || o.Weight == nil {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDimensionViewModel) GetWeightOk() (*float64, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *InventoryDimensionViewModel) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *InventoryDimensionViewModel) SetWeight(v float64) {
	o.Weight = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *InventoryDimensionViewModel) GetWidth() float64 {
	if o == nil || o.Width == nil {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryDimensionViewModel) GetWidthOk() (*float64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *InventoryDimensionViewModel) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *InventoryDimensionViewModel) SetWidth(v float64) {
	o.Width = &v
}

func (o InventoryDimensionViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depth != nil {
		toSerialize["depth"] = o.Depth
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryDimensionViewModel struct {
	value *InventoryDimensionViewModel
	isSet bool
}

func (v NullableInventoryDimensionViewModel) Get() *InventoryDimensionViewModel {
	return v.value
}

func (v *NullableInventoryDimensionViewModel) Set(val *InventoryDimensionViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDimensionViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDimensionViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDimensionViewModel(val *InventoryDimensionViewModel) *NullableInventoryDimensionViewModel {
	return &NullableInventoryDimensionViewModel{value: val, isSet: true}
}

func (v NullableInventoryDimensionViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDimensionViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

