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
	AwaitingQuantity *int32 `json:"awaiting_quantity,omitempty"`
	// Amount of committed quantity at this fulfillment center
	CommittedQuantity *int32 `json:"committed_quantity,omitempty"`
	// Amount of fulfillable quantity at this fulfillment center
	FulfillableQuantity *int32 `json:"fulfillable_quantity,omitempty"`
	// Unique id of the fulfillment center
	Id *int32 `json:"id,omitempty"`
	// The quantity of items that are in the process of internal transit  between ShipBob fulfillment centers, with a destination of this fulfillment center. These items are not pickable or fulfillable until they have been received and moved  to the \"On Hand\" quantity of the destination FC.
	InternalTransferQuantity *int32 `json:"internal_transfer_quantity,omitempty"`
	// Name of the fulfillment center
	Name *string `json:"name,omitempty"`
	// Amount of onhand quantity at this fulfillment center
	OnhandQuantity *int32 `json:"onhand_quantity,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InventoryFulfillmentCenterQuantity) SetId(v int32) {
	o.Id = &v
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InventoryFulfillmentCenterQuantity) SetName(v string) {
	o.Name = &v
}

// GetOnhandQuantity returns the OnhandQuantity field value if set, zero value otherwise.
func (o *InventoryFulfillmentCenterQuantity) GetOnhandQuantity() int32 {
	if o == nil || IsNil(o.OnhandQuantity) {
		var ret int32
		return ret
	}
	return *o.OnhandQuantity
}

// GetOnhandQuantityOk returns a tuple with the OnhandQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryFulfillmentCenterQuantity) GetOnhandQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.OnhandQuantity) {
		return nil, false
	}
	return o.OnhandQuantity, true
}

// HasOnhandQuantity returns a boolean if a field has been set.
func (o *InventoryFulfillmentCenterQuantity) HasOnhandQuantity() bool {
	if o != nil && !IsNil(o.OnhandQuantity) {
		return true
	}

	return false
}

// SetOnhandQuantity gets a reference to the given int32 and assigns it to the OnhandQuantity field.
func (o *InventoryFulfillmentCenterQuantity) SetOnhandQuantity(v int32) {
	o.OnhandQuantity = &v
}

func (o InventoryFulfillmentCenterQuantity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryFulfillmentCenterQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwaitingQuantity) {
		toSerialize["awaiting_quantity"] = o.AwaitingQuantity
	}
	if !IsNil(o.CommittedQuantity) {
		toSerialize["committed_quantity"] = o.CommittedQuantity
	}
	if !IsNil(o.FulfillableQuantity) {
		toSerialize["fulfillable_quantity"] = o.FulfillableQuantity
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InternalTransferQuantity) {
		toSerialize["internal_transfer_quantity"] = o.InternalTransferQuantity
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OnhandQuantity) {
		toSerialize["onhand_quantity"] = o.OnhandQuantity
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


