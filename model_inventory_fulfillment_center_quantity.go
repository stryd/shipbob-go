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

// checks if the InventoryFulfillmentCenterQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryFulfillmentCenterQuantity{}

// InventoryFulfillmentCenterQuantity Break down of fulfillable quantity by fulfillment center
type InventoryFulfillmentCenterQuantity struct {
	// Amount of quantity awaiting arrival of a receiving order at this fulfillment center
	AwaitingQuantity *int32 `json:"awaitingQuantity,omitempty"`
	// Amount of committed quantity at this fulfillment center
	CommittedQuantity *int32 `json:"committedQuantity,omitempty"`
	// Amount of fulfillable quantity at this fulfillment center
	FulfillableQuantity *int32 `json:"fulfillableQuantity,omitempty"`
	// Unique id of the fulfillment center
	FulfillmentCenterId *int32 `json:"fulfillmentCenterId,omitempty"`
	// The quantity of items that are in the process of internal transit  between ShipBob fulfillment centers, with a destination of this fulfillment center. These items are not pickable or fulfillable until they have been received and moved  to the \"On Hand\" quantity of the destination FC.
	InternalTransferQuantity *int32 `json:"internalTransferQuantity,omitempty"`
	// Name of the fulfillment center
	Name NullableString `json:"name,omitempty"`
	// Amount of onhand quantity at this fulfillment center
	OnHandQuantity *int32 `json:"onHandQuantity,omitempty"`
}

// NewInventoryFulfillmentCenterQuantity instantiates a new InventoryFulfillmentCenterQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryFulfillmentCenterQuantity() *InventoryFulfillmentCenterQuantity {
	this := InventoryFulfillmentCenterQuantity{}
	return &this
}

// NewInventoryFulfillmentCenterQuantityWithDefaults instantiates a new InventoryFulfillmentCenterQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryFulfillmentCenterQuantityWithDefaults() *InventoryFulfillmentCenterQuantity {
	this := InventoryFulfillmentCenterQuantity{}
	return &this
}

// GetAwaitingQuantity returns the AwaitingQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetAwaitingQuantity() int32 {
	if o == nil || IsNil(o.AwaitingQuantity) {
		var ret int32
		return ret
	}
	return *o.AwaitingQuantity
}

// GetAwaitingQuantityOk returns a tuple with the AwaitingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetAwaitingQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.AwaitingQuantity) {
		return nil, false
	}
	return o.AwaitingQuantity, true
}

// HasAwaitingQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasAwaitingQuantity() bool {
	if o != nil && !IsNil(o.AwaitingQuantity) {
		return true
	}

	return false
}

// SetAwaitingQuantity gets a reference to the given int32 and assigns it to the AwaitingQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetAwaitingQuantity(v int32) {
	o.AwaitingQuantity = &v
}

// GetCommittedQuantity returns the CommittedQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetCommittedQuantity() int32 {
	if o == nil || IsNil(o.CommittedQuantity) {
		var ret int32
		return ret
	}
	return *o.CommittedQuantity
}

// GetCommittedQuantityOk returns a tuple with the CommittedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetCommittedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.CommittedQuantity) {
		return nil, false
	}
	return o.CommittedQuantity, true
}

// HasCommittedQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasCommittedQuantity() bool {
	if o != nil && !IsNil(o.CommittedQuantity) {
		return true
	}

	return false
}

// SetCommittedQuantity gets a reference to the given int32 and assigns it to the CommittedQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetCommittedQuantity(v int32) {
	o.CommittedQuantity = &v
}

// GetFulfillableQuantity returns the FulfillableQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetFulfillableQuantity() int32 {
	if o == nil || IsNil(o.FulfillableQuantity) {
		var ret int32
		return ret
	}
	return *o.FulfillableQuantity
}

// GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetFulfillableQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.FulfillableQuantity) {
		return nil, false
	}
	return o.FulfillableQuantity, true
}

// HasFulfillableQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasFulfillableQuantity() bool {
	if o != nil && !IsNil(o.FulfillableQuantity) {
		return true
	}

	return false
}

// SetFulfillableQuantity gets a reference to the given int32 and assigns it to the FulfillableQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetFulfillableQuantity(v int32) {
	o.FulfillableQuantity = &v
}

// GetFulfillmentCenterId returns the FulfillmentCenterId field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetFulfillmentCenterId() int32 {
	if o == nil || IsNil(o.FulfillmentCenterId) {
		var ret int32
		return ret
	}
	return *o.FulfillmentCenterId
}

// GetFulfillmentCenterIdOk returns a tuple with the FulfillmentCenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetFulfillmentCenterIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FulfillmentCenterId) {
		return nil, false
	}
	return o.FulfillmentCenterId, true
}

// HasFulfillmentCenterId returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasFulfillmentCenterId() bool {
	if o != nil && !IsNil(o.FulfillmentCenterId) {
		return true
	}

	return false
}

// SetFulfillmentCenterId gets a reference to the given int32 and assigns it to the FulfillmentCenterId field.
func (o *InventoryFulfillmentCenterQuantity) SetFulfillmentCenterId(v int32) {
	o.FulfillmentCenterId = &v
}

// GetInternalTransferQuantity returns the InternalTransferQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetInternalTransferQuantity() int32 {
	if o == nil || IsNil(o.InternalTransferQuantity) {
		var ret int32
		return ret
	}
	return *o.InternalTransferQuantity
}

// GetInternalTransferQuantityOk returns a tuple with the InternalTransferQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetInternalTransferQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.InternalTransferQuantity) {
		return nil, false
	}
	return o.InternalTransferQuantity, true
}

// HasInternalTransferQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasInternalTransferQuantity() bool {
	if o != nil && !IsNil(o.InternalTransferQuantity) {
		return true
	}

	return false
}

// SetInternalTransferQuantity gets a reference to the given int32 and assigns it to the InternalTransferQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetInternalTransferQuantity(v int32) {
	o.InternalTransferQuantity = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryFulfillmentCenterQuantity) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryFulfillmentCenterQuantity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InventoryFulfillmentCenterQuantity) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *InventoryFulfillmentCenterQuantity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InventoryFulfillmentCenterQuantity) UnsetName() {
	o.Name.Unset()
}

// GetOnHandQuantity returns the OnHandQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetOnHandQuantity() int32 {
	if o == nil || IsNil(o.OnHandQuantity) {
		var ret int32
		return ret
	}
	return *o.OnHandQuantity
}

// GetOnHandQuantityOk returns a tuple with the OnHandQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetOnHandQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.OnHandQuantity) {
		return nil, false
	}
	return o.OnHandQuantity, true
}

// HasOnHandQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasOnHandQuantity() bool {
	if o != nil && !IsNil(o.OnHandQuantity) {
		return true
	}

	return false
}

// SetOnHandQuantity gets a reference to the given int32 and assigns it to the OnHandQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetOnHandQuantity(v int32) {
	o.OnHandQuantity = &v
}

func (o InventoryFulfillmentCenterQuantity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryFulfillmentCenterQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwaitingQuantity) {
		toSerialize["awaitingQuantity"] = o.AwaitingQuantity
	}
	if !IsNil(o.CommittedQuantity) {
		toSerialize["committedQuantity"] = o.CommittedQuantity
	}
	if !IsNil(o.FulfillableQuantity) {
		toSerialize["fulfillableQuantity"] = o.FulfillableQuantity
	}
	if !IsNil(o.FulfillmentCenterId) {
		toSerialize["fulfillmentCenterId"] = o.FulfillmentCenterId
	}
	if !IsNil(o.InternalTransferQuantity) {
		toSerialize["internalTransferQuantity"] = o.InternalTransferQuantity
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.OnHandQuantity) {
		toSerialize["onHandQuantity"] = o.OnHandQuantity
	}
	return toSerialize, nil
}

type NullableInventoryFulfillmentCenterQuantity struct {
	value *InventoryFulfillmentCenterQuantity
	isSet bool
}

func (v NullableInventoryFulfillmentCenterQuantity) Get() *InventoryFulfillmentCenterQuantity {
	return v.value
}

func (v *NullableInventoryFulfillmentCenterQuantity) Set(val *InventoryFulfillmentCenterQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryFulfillmentCenterQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryFulfillmentCenterQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryFulfillmentCenterQuantity(val *InventoryFulfillmentCenterQuantity) *NullableInventoryFulfillmentCenterQuantity {
	return &NullableInventoryFulfillmentCenterQuantity{value: val, isSet: true}
}

func (v NullableInventoryFulfillmentCenterQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryFulfillmentCenterQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
