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

// checks if the OrderInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderInventory{}

// OrderInventory Information about inventory belonging to a store product
type OrderInventory struct {
	// Expiration date of the inventory
	ExpirationDate NullableTime `json:"expiration_date,omitempty"`
	// Unique id of the inventory
	Id *int32 `json:"id,omitempty"`
	// Is inventory Dangerous Good
	IsDangerousGoods *bool `json:"is_dangerous_goods,omitempty"`
	// Lot number of the inventory
	Lot *string `json:"lot,omitempty"`
	// Name of the inventory item
	Name *string `json:"name,omitempty"`
	// Quantity of the inventory item to be included in the fulfillment
	Quantity *int32 `json:"quantity,omitempty"`
	// The quantity of the inventory item allocated from the assigned fulfillment center and committed to the order. If quantity committed is less than order quantity, then the inventory item is out of stock at the assigned fulfillment center.
	QuantityCommitted *int32 `json:"quantity_committed,omitempty"`
	// Serial number of the inventory
	SerialNumbers []string `json:"serial_numbers,omitempty"`
}

// NewOrderInventory instantiates a new OrderInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderInventory() *OrderInventory {
	this := OrderInventory{}
	return &this
}

// NewOrderInventoryWithDefaults instantiates a new OrderInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderInventoryWithDefaults() *OrderInventory {
	this := OrderInventory{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderInventory) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderInventory) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *OrderInventory) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *OrderInventory) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}

// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *OrderInventory) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *OrderInventory) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderInventory) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderInventory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrderInventory) SetId(v int32) {
	o.Id = &v
}

// GetIsDangerousGoods returns the IsDangerousGoods field value if set, zero value otherwise.
func (o *OrderInventory) GetIsDangerousGoods() bool {
	if o == nil || IsNil(o.IsDangerousGoods) {
		var ret bool
		return ret
	}
	return *o.IsDangerousGoods
}

// GetIsDangerousGoodsOk returns a tuple with the IsDangerousGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetIsDangerousGoodsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDangerousGoods) {
		return nil, false
	}
	return o.IsDangerousGoods, true
}

// HasIsDangerousGoods returns a boolean if a field has been set.
func (o *OrderInventory) HasIsDangerousGoods() bool {
	if o != nil && !IsNil(o.IsDangerousGoods) {
		return true
	}

	return false
}

// SetIsDangerousGoods gets a reference to the given bool and assigns it to the IsDangerousGoods field.
func (o *OrderInventory) SetIsDangerousGoods(v bool) {
	o.IsDangerousGoods = &v
}

// GetLot returns the Lot field value if set, zero value otherwise.
func (o *OrderInventory) GetLot() string {
	if o == nil || IsNil(o.Lot) {
		var ret string
		return ret
	}
	return *o.Lot
}

// GetLotOk returns a tuple with the Lot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetLotOk() (*string, bool) {
	if o == nil || IsNil(o.Lot) {
		return nil, false
	}
	return o.Lot, true
}

// HasLot returns a boolean if a field has been set.
func (o *OrderInventory) HasLot() bool {
	if o != nil && !IsNil(o.Lot) {
		return true
	}

	return false
}

// SetLot gets a reference to the given string and assigns it to the Lot field.
func (o *OrderInventory) SetLot(v string) {
	o.Lot = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderInventory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderInventory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderInventory) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderInventory) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderInventory) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderInventory) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetQuantityCommitted returns the QuantityCommitted field value if set, zero value otherwise.
func (o *OrderInventory) GetQuantityCommitted() int32 {
	if o == nil || IsNil(o.QuantityCommitted) {
		var ret int32
		return ret
	}
	return *o.QuantityCommitted
}

// GetQuantityCommittedOk returns a tuple with the QuantityCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetQuantityCommittedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityCommitted) {
		return nil, false
	}
	return o.QuantityCommitted, true
}

// HasQuantityCommitted returns a boolean if a field has been set.
func (o *OrderInventory) HasQuantityCommitted() bool {
	if o != nil && !IsNil(o.QuantityCommitted) {
		return true
	}

	return false
}

// SetQuantityCommitted gets a reference to the given int32 and assigns it to the QuantityCommitted field.
func (o *OrderInventory) SetQuantityCommitted(v int32) {
	o.QuantityCommitted = &v
}

// GetSerialNumbers returns the SerialNumbers field value if set, zero value otherwise.
func (o *OrderInventory) GetSerialNumbers() []string {
	if o == nil || IsNil(o.SerialNumbers) {
		var ret []string
		return ret
	}
	return o.SerialNumbers
}

// GetSerialNumbersOk returns a tuple with the SerialNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderInventory) GetSerialNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.SerialNumbers) {
		return nil, false
	}
	return o.SerialNumbers, true
}

// HasSerialNumbers returns a boolean if a field has been set.
func (o *OrderInventory) HasSerialNumbers() bool {
	if o != nil && !IsNil(o.SerialNumbers) {
		return true
	}

	return false
}

// SetSerialNumbers gets a reference to the given []string and assigns it to the SerialNumbers field.
func (o *OrderInventory) SetSerialNumbers(v []string) {
	o.SerialNumbers = v
}

func (o OrderInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDangerousGoods) {
		toSerialize["is_dangerous_goods"] = o.IsDangerousGoods
	}
	if !IsNil(o.Lot) {
		toSerialize["lot"] = o.Lot
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.QuantityCommitted) {
		toSerialize["quantity_committed"] = o.QuantityCommitted
	}
	if !IsNil(o.SerialNumbers) {
		toSerialize["serial_numbers"] = o.SerialNumbers
	}
	return toSerialize, nil
}

type NullableOrderInventory struct {
	value *OrderInventory
	isSet bool
}

func (v NullableOrderInventory) Get() *OrderInventory {
	return v.value
}

func (v *NullableOrderInventory) Set(val *OrderInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderInventory(val *OrderInventory) *NullableOrderInventory {
	return &NullableOrderInventory{value: val, isSet: true}
}

func (v NullableOrderInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
