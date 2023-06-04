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

// checks if the InventoryQuantityV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryQuantityV2{}

// InventoryQuantityV2 struct for InventoryQuantityV2
type InventoryQuantityV2 struct {
	// Quantity of the inventory item submitted in the WRO
	ExpectedQuantity *int32 `json:"expected_quantity,omitempty"`
	// ID of the inventory item
	InventoryId *int32 `json:"inventory_id,omitempty"`
	// Quantity of the inventory item received by the warehouse
	ReceivedQuantity *int32 `json:"received_quantity,omitempty"`
	// Sku of the inventory item
	Sku NullableString `json:"sku,omitempty"`
}

// NewInventoryQuantityV2 instantiates a new InventoryQuantityV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryQuantityV2() *InventoryQuantityV2 {
	this := InventoryQuantityV2{}
	return &this
}

// NewInventoryQuantityV2WithDefaults instantiates a new InventoryQuantityV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryQuantityV2WithDefaults() *InventoryQuantityV2 {
	this := InventoryQuantityV2{}
	return &this
}

// GetExpectedQuantity returns the ExpectedQuantity field value if set, zero value otherwise.
func (o *InventoryQuantityV2) GetExpectedQuantity() int32 {
	if o == nil || IsNil(o.ExpectedQuantity) {
		var ret int32
		return ret
	}
	return *o.ExpectedQuantity
}

// GetExpectedQuantityOk returns a tuple with the ExpectedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryQuantityV2) GetExpectedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpectedQuantity) {
		return nil, false
	}
	return o.ExpectedQuantity, true
}

// HasExpectedQuantity returns a boolean if a field has been set.
func (o *InventoryQuantityV2) HasExpectedQuantity() bool {
	if o != nil && !IsNil(o.ExpectedQuantity) {
		return true
	}

	return false
}

// SetExpectedQuantity gets a reference to the given int32 and assigns it to the ExpectedQuantity field.
func (o *InventoryQuantityV2) SetExpectedQuantity(v int32) {
	o.ExpectedQuantity = &v
}

// GetInventoryId returns the InventoryId field value if set, zero value otherwise.
func (o *InventoryQuantityV2) GetInventoryId() int32 {
	if o == nil || IsNil(o.InventoryId) {
		var ret int32
		return ret
	}
	return *o.InventoryId
}

// GetInventoryIdOk returns a tuple with the InventoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryQuantityV2) GetInventoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InventoryId) {
		return nil, false
	}
	return o.InventoryId, true
}

// HasInventoryId returns a boolean if a field has been set.
func (o *InventoryQuantityV2) HasInventoryId() bool {
	if o != nil && !IsNil(o.InventoryId) {
		return true
	}

	return false
}

// SetInventoryId gets a reference to the given int32 and assigns it to the InventoryId field.
func (o *InventoryQuantityV2) SetInventoryId(v int32) {
	o.InventoryId = &v
}

// GetReceivedQuantity returns the ReceivedQuantity field value if set, zero value otherwise.
func (o *InventoryQuantityV2) GetReceivedQuantity() int32 {
	if o == nil || IsNil(o.ReceivedQuantity) {
		var ret int32
		return ret
	}
	return *o.ReceivedQuantity
}

// GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryQuantityV2) GetReceivedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceivedQuantity) {
		return nil, false
	}
	return o.ReceivedQuantity, true
}

// HasReceivedQuantity returns a boolean if a field has been set.
func (o *InventoryQuantityV2) HasReceivedQuantity() bool {
	if o != nil && !IsNil(o.ReceivedQuantity) {
		return true
	}

	return false
}

// SetReceivedQuantity gets a reference to the given int32 and assigns it to the ReceivedQuantity field.
func (o *InventoryQuantityV2) SetReceivedQuantity(v int32) {
	o.ReceivedQuantity = &v
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryQuantityV2) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryQuantityV2) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryQuantityV2) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *InventoryQuantityV2) SetSku(v string) {
	o.Sku.Set(&v)
}

// SetSkuNil sets the value for Sku to be an explicit nil
func (o *InventoryQuantityV2) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *InventoryQuantityV2) UnsetSku() {
	o.Sku.Unset()
}

func (o InventoryQuantityV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryQuantityV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpectedQuantity) {
		toSerialize["expected_quantity"] = o.ExpectedQuantity
	}
	if !IsNil(o.InventoryId) {
		toSerialize["inventory_id"] = o.InventoryId
	}
	if !IsNil(o.ReceivedQuantity) {
		toSerialize["received_quantity"] = o.ReceivedQuantity
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	return toSerialize, nil
}

type NullableInventoryQuantityV2 struct {
	value *InventoryQuantityV2
	isSet bool
}

func (v NullableInventoryQuantityV2) Get() *InventoryQuantityV2 {
	return v.value
}

func (v *NullableInventoryQuantityV2) Set(val *InventoryQuantityV2) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryQuantityV2) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryQuantityV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryQuantityV2(val *InventoryQuantityV2) *NullableInventoryQuantityV2 {
	return &NullableInventoryQuantityV2{value: val, isSet: true}
}

func (v NullableInventoryQuantityV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryQuantityV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
