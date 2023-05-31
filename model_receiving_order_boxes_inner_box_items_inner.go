/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"time"
)

// checks if the ReceivingOrderBoxesInnerBoxItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivingOrderBoxesInnerBoxItemsInner{}

// ReceivingOrderBoxesInnerBoxItemsInner Information about an item contained inside a box as part of a receiving order
type ReceivingOrderBoxesInnerBoxItemsInner struct {
	// Unique identifier of the inventory item
	InventoryId *int32 `json:"inventory_id,omitempty"`
	// Expiration date of the item's lot
	LotDate NullableTime `json:"lot_date,omitempty"`
	// Lot number the item belongs to
	LotNumber NullableString `json:"lot_number,omitempty"`
	// Quantity of the item included
	Quantity *int32 `json:"quantity,omitempty"`
	// Quantity of the item that was received after processing the receiving order
	ReceivedQuantity *int32 `json:"received_quantity,omitempty"`
}

// NewReceivingOrderBoxesInnerBoxItemsInner instantiates a new ReceivingOrderBoxesInnerBoxItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivingOrderBoxesInnerBoxItemsInner() *ReceivingOrderBoxesInnerBoxItemsInner {
	this := ReceivingOrderBoxesInnerBoxItemsInner{}
	return &this
}

// NewReceivingOrderBoxesInnerBoxItemsInnerWithDefaults instantiates a new ReceivingOrderBoxesInnerBoxItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivingOrderBoxesInnerBoxItemsInnerWithDefaults() *ReceivingOrderBoxesInnerBoxItemsInner {
	this := ReceivingOrderBoxesInnerBoxItemsInner{}
	return &this
}

// GetInventoryId returns the InventoryId field value if set, zero value otherwise.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetInventoryId() int32 {
	if o == nil || IsNil(o.InventoryId) {
		var ret int32
		return ret
	}
	return *o.InventoryId
}

// GetInventoryIdOk returns a tuple with the InventoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetInventoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InventoryId) {
		return nil, false
	}
	return o.InventoryId, true
}

// HasInventoryId returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasInventoryId() bool {
	if o != nil && !IsNil(o.InventoryId) {
		return true
	}

	return false
}

// SetInventoryId gets a reference to the given int32 and assigns it to the InventoryId field.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetInventoryId(v int32) {
	o.InventoryId = &v
}

// GetLotDate returns the LotDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotDate() time.Time {
	if o == nil || IsNil(o.LotDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LotDate.Get()
}

// GetLotDateOk returns a tuple with the LotDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LotDate.Get(), o.LotDate.IsSet()
}

// HasLotDate returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasLotDate() bool {
	if o != nil && o.LotDate.IsSet() {
		return true
	}

	return false
}

// SetLotDate gets a reference to the given NullableTime and assigns it to the LotDate field.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotDate(v time.Time) {
	o.LotDate.Set(&v)
}
// SetLotDateNil sets the value for LotDate to be an explicit nil
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotDateNil() {
	o.LotDate.Set(nil)
}

// UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
func (o *ReceivingOrderBoxesInnerBoxItemsInner) UnsetLotDate() {
	o.LotDate.Unset()
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber.Get()) {
		var ret string
		return ret
	}
	return *o.LotNumber.Get()
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LotNumber.Get(), o.LotNumber.IsSet()
}

// HasLotNumber returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasLotNumber() bool {
	if o != nil && o.LotNumber.IsSet() {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given NullableString and assigns it to the LotNumber field.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotNumber(v string) {
	o.LotNumber.Set(&v)
}
// SetLotNumberNil sets the value for LotNumber to be an explicit nil
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotNumberNil() {
	o.LotNumber.Set(nil)
}

// UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
func (o *ReceivingOrderBoxesInnerBoxItemsInner) UnsetLotNumber() {
	o.LotNumber.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetReceivedQuantity returns the ReceivedQuantity field value if set, zero value otherwise.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetReceivedQuantity() int32 {
	if o == nil || IsNil(o.ReceivedQuantity) {
		var ret int32
		return ret
	}
	return *o.ReceivedQuantity
}

// GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetReceivedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceivedQuantity) {
		return nil, false
	}
	return o.ReceivedQuantity, true
}

// HasReceivedQuantity returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasReceivedQuantity() bool {
	if o != nil && !IsNil(o.ReceivedQuantity) {
		return true
	}

	return false
}

// SetReceivedQuantity gets a reference to the given int32 and assigns it to the ReceivedQuantity field.
func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetReceivedQuantity(v int32) {
	o.ReceivedQuantity = &v
}

func (o ReceivingOrderBoxesInnerBoxItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivingOrderBoxesInnerBoxItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InventoryId) {
		toSerialize["inventory_id"] = o.InventoryId
	}
	if o.LotDate.IsSet() {
		toSerialize["lot_date"] = o.LotDate.Get()
	}
	if o.LotNumber.IsSet() {
		toSerialize["lot_number"] = o.LotNumber.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ReceivedQuantity) {
		toSerialize["received_quantity"] = o.ReceivedQuantity
	}
	return toSerialize, nil
}

type NullableReceivingOrderBoxesInnerBoxItemsInner struct {
	value *ReceivingOrderBoxesInnerBoxItemsInner
	isSet bool
}

func (v NullableReceivingOrderBoxesInnerBoxItemsInner) Get() *ReceivingOrderBoxesInnerBoxItemsInner {
	return v.value
}

func (v *NullableReceivingOrderBoxesInnerBoxItemsInner) Set(val *ReceivingOrderBoxesInnerBoxItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingOrderBoxesInnerBoxItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingOrderBoxesInnerBoxItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingOrderBoxesInnerBoxItemsInner(val *ReceivingOrderBoxesInnerBoxItemsInner) *NullableReceivingOrderBoxesInnerBoxItemsInner {
	return &NullableReceivingOrderBoxesInnerBoxItemsInner{value: val, isSet: true}
}

func (v NullableReceivingOrderBoxesInnerBoxItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingOrderBoxesInnerBoxItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


