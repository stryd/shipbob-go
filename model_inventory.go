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

// checks if the Inventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Inventory{}

// Inventory Information about an inventory item
type Inventory struct {
	Dimensions *Dimension `json:"dimensions,omitempty"`
	// Fulfillable quantity of this inventory item broken down by fulfillment center location
	FulfillableQuantityByFulfillmentCenter []InventoryFulfillmentCenterQuantity `json:"fulfillable_quantity_by_fulfillment_center,omitempty"`
	// Fulfillable quantity of this inventory item broken down by lot
	FulfillableQuantityByLot []InventoryLotQuantity `json:"fulfillable_quantity_by_lot,omitempty"`
	// Unique id of the inventory item
	Id *int32 `json:"id,omitempty"`
	// Whether the inventory is active or not
	IsActive *bool `json:"is_active,omitempty"`
	// True if the inventory item is marked as case pick
	IsCasePick *bool `json:"is_case_pick,omitempty"`
	// True if the inventory item is marked as a digital item
	IsDigital *bool `json:"is_digital,omitempty"`
	// True if this inventory item is organized into lots
	IsLot *bool `json:"is_lot,omitempty"`
	// Name of the inventory item
	Name *string `json:"name,omitempty"`
	// Attribute influencing the packaging requirements of this inventory item
	PackagingAttribute *string `json:"packaging_attribute,omitempty"`
	// Total quantity in unreceived receiving orders for this inventory item
	TotalAwaitingQuantity *int32 `json:"total_awaiting_quantity,omitempty"`
	// The amount of the item you need to send to ShipBob to fulfill all exception orders containing  the item. This is the exception_quantity less the fulfillable_quantity of the item.
	TotalBackorderedQuantity *int32 `json:"total_backordered_quantity,omitempty"`
	// Total committed quantity of this inventory item
	TotalCommittedQuantity *int32 `json:"total_committed_quantity,omitempty"`
	// The total quantity of all items that are contained within orders that are in exception/out of stock status. Out of stock orders have not been processed and therefore do not have lot or fulfillment centers assigned.
	TotalExceptionQuantity *int32 `json:"total_exception_quantity,omitempty"`
	// Total fulfillable quantity of this inventory item
	TotalFulfillableQuantity *int32 `json:"total_fulfillable_quantity,omitempty"`
	// The total quantity of all items that are in the process of internal transit  between ShipBob fulfillment centers. These items are not pickable or fulfillable until they have been received and moved to the \"On Hand\" quantity of the destination FC. Internal transit quantities for each FC represent the incoming transfer stock for that specific location.
	TotalInternalTransferQuantity *int32 `json:"total_internal_transfer_quantity,omitempty"`
	// Total onhand quantity of this inventory item
	TotalOnhandQuantity *int32 `json:"total_onhand_quantity,omitempty"`
}

// NewInventory instantiates a new Inventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventory() *Inventory {
	this := Inventory{}
	return &this
}

// NewInventoryWithDefaults instantiates a new Inventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryWithDefaults() *Inventory {
	this := Inventory{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *Inventory) GetDimensions() Dimension {
	if o == nil || IsNil(o.Dimensions) {
		var ret Dimension
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetDimensionsOk() (*Dimension, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *Inventory) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimension and assigns it to the Dimensions field.
func (o *Inventory) SetDimensions(v Dimension) {
	o.Dimensions = &v
}

// GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field value if set, zero value otherwise.
func (o *Inventory) GetFulfillableQuantityByFulfillmentCenter() []InventoryFulfillmentCenterQuantity {
	if o == nil || IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		var ret []InventoryFulfillmentCenterQuantity
		return ret
	}
	return o.FulfillableQuantityByFulfillmentCenter
}

// GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetFulfillableQuantityByFulfillmentCenterOk() ([]InventoryFulfillmentCenterQuantity, bool) {
	if o == nil || IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		return nil, false
	}
	return o.FulfillableQuantityByFulfillmentCenter, true
}

// HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.
func (o *Inventory) HasFulfillableQuantityByFulfillmentCenter() bool {
	if o != nil && !IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		return true
	}

	return false
}

// SetFulfillableQuantityByFulfillmentCenter gets a reference to the given []InventoryFulfillmentCenterQuantity and assigns it to the FulfillableQuantityByFulfillmentCenter field.
func (o *Inventory) SetFulfillableQuantityByFulfillmentCenter(v []InventoryFulfillmentCenterQuantity) {
	o.FulfillableQuantityByFulfillmentCenter = v
}

// GetFulfillableQuantityByLot returns the FulfillableQuantityByLot field value if set, zero value otherwise.
func (o *Inventory) GetFulfillableQuantityByLot() []InventoryLotQuantity {
	if o == nil || IsNil(o.FulfillableQuantityByLot) {
		var ret []InventoryLotQuantity
		return ret
	}
	return o.FulfillableQuantityByLot
}

// GetFulfillableQuantityByLotOk returns a tuple with the FulfillableQuantityByLot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetFulfillableQuantityByLotOk() ([]InventoryLotQuantity, bool) {
	if o == nil || IsNil(o.FulfillableQuantityByLot) {
		return nil, false
	}
	return o.FulfillableQuantityByLot, true
}

// HasFulfillableQuantityByLot returns a boolean if a field has been set.
func (o *Inventory) HasFulfillableQuantityByLot() bool {
	if o != nil && !IsNil(o.FulfillableQuantityByLot) {
		return true
	}

	return false
}

// SetFulfillableQuantityByLot gets a reference to the given []InventoryLotQuantity and assigns it to the FulfillableQuantityByLot field.
func (o *Inventory) SetFulfillableQuantityByLot(v []InventoryLotQuantity) {
	o.FulfillableQuantityByLot = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Inventory) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Inventory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Inventory) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Inventory) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Inventory) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Inventory) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsCasePick returns the IsCasePick field value if set, zero value otherwise.
func (o *Inventory) GetIsCasePick() bool {
	if o == nil || IsNil(o.IsCasePick) {
		var ret bool
		return ret
	}
	return *o.IsCasePick
}

// GetIsCasePickOk returns a tuple with the IsCasePick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetIsCasePickOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCasePick) {
		return nil, false
	}
	return o.IsCasePick, true
}

// HasIsCasePick returns a boolean if a field has been set.
func (o *Inventory) HasIsCasePick() bool {
	if o != nil && !IsNil(o.IsCasePick) {
		return true
	}

	return false
}

// SetIsCasePick gets a reference to the given bool and assigns it to the IsCasePick field.
func (o *Inventory) SetIsCasePick(v bool) {
	o.IsCasePick = &v
}

// GetIsDigital returns the IsDigital field value if set, zero value otherwise.
func (o *Inventory) GetIsDigital() bool {
	if o == nil || IsNil(o.IsDigital) {
		var ret bool
		return ret
	}
	return *o.IsDigital
}

// GetIsDigitalOk returns a tuple with the IsDigital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetIsDigitalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDigital) {
		return nil, false
	}
	return o.IsDigital, true
}

// HasIsDigital returns a boolean if a field has been set.
func (o *Inventory) HasIsDigital() bool {
	if o != nil && !IsNil(o.IsDigital) {
		return true
	}

	return false
}

// SetIsDigital gets a reference to the given bool and assigns it to the IsDigital field.
func (o *Inventory) SetIsDigital(v bool) {
	o.IsDigital = &v
}

// GetIsLot returns the IsLot field value if set, zero value otherwise.
func (o *Inventory) GetIsLot() bool {
	if o == nil || IsNil(o.IsLot) {
		var ret bool
		return ret
	}
	return *o.IsLot
}

// GetIsLotOk returns a tuple with the IsLot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetIsLotOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLot) {
		return nil, false
	}
	return o.IsLot, true
}

// HasIsLot returns a boolean if a field has been set.
func (o *Inventory) HasIsLot() bool {
	if o != nil && !IsNil(o.IsLot) {
		return true
	}

	return false
}

// SetIsLot gets a reference to the given bool and assigns it to the IsLot field.
func (o *Inventory) SetIsLot(v bool) {
	o.IsLot = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Inventory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Inventory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Inventory) SetName(v string) {
	o.Name = &v
}

// GetPackagingAttribute returns the PackagingAttribute field value if set, zero value otherwise.
func (o *Inventory) GetPackagingAttribute() string {
	if o == nil || IsNil(o.PackagingAttribute) {
		var ret string
		return ret
	}
	return *o.PackagingAttribute
}

// GetPackagingAttributeOk returns a tuple with the PackagingAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetPackagingAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.PackagingAttribute) {
		return nil, false
	}
	return o.PackagingAttribute, true
}

// HasPackagingAttribute returns a boolean if a field has been set.
func (o *Inventory) HasPackagingAttribute() bool {
	if o != nil && !IsNil(o.PackagingAttribute) {
		return true
	}

	return false
}

// SetPackagingAttribute gets a reference to the given string and assigns it to the PackagingAttribute field.
func (o *Inventory) SetPackagingAttribute(v string) {
	o.PackagingAttribute = &v
}

// GetTotalAwaitingQuantity returns the TotalAwaitingQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalAwaitingQuantity() int32 {
	if o == nil || IsNil(o.TotalAwaitingQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalAwaitingQuantity
}

// GetTotalAwaitingQuantityOk returns a tuple with the TotalAwaitingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalAwaitingQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalAwaitingQuantity) {
		return nil, false
	}
	return o.TotalAwaitingQuantity, true
}

// HasTotalAwaitingQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalAwaitingQuantity() bool {
	if o != nil && !IsNil(o.TotalAwaitingQuantity) {
		return true
	}

	return false
}

// SetTotalAwaitingQuantity gets a reference to the given int32 and assigns it to the TotalAwaitingQuantity field.
func (o *Inventory) SetTotalAwaitingQuantity(v int32) {
	o.TotalAwaitingQuantity = &v
}

// GetTotalBackorderedQuantity returns the TotalBackorderedQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalBackorderedQuantity() int32 {
	if o == nil || IsNil(o.TotalBackorderedQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalBackorderedQuantity
}

// GetTotalBackorderedQuantityOk returns a tuple with the TotalBackorderedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalBackorderedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalBackorderedQuantity) {
		return nil, false
	}
	return o.TotalBackorderedQuantity, true
}

// HasTotalBackorderedQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalBackorderedQuantity() bool {
	if o != nil && !IsNil(o.TotalBackorderedQuantity) {
		return true
	}

	return false
}

// SetTotalBackorderedQuantity gets a reference to the given int32 and assigns it to the TotalBackorderedQuantity field.
func (o *Inventory) SetTotalBackorderedQuantity(v int32) {
	o.TotalBackorderedQuantity = &v
}

// GetTotalCommittedQuantity returns the TotalCommittedQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalCommittedQuantity() int32 {
	if o == nil || IsNil(o.TotalCommittedQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalCommittedQuantity
}

// GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalCommittedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCommittedQuantity) {
		return nil, false
	}
	return o.TotalCommittedQuantity, true
}

// HasTotalCommittedQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalCommittedQuantity() bool {
	if o != nil && !IsNil(o.TotalCommittedQuantity) {
		return true
	}

	return false
}

// SetTotalCommittedQuantity gets a reference to the given int32 and assigns it to the TotalCommittedQuantity field.
func (o *Inventory) SetTotalCommittedQuantity(v int32) {
	o.TotalCommittedQuantity = &v
}

// GetTotalExceptionQuantity returns the TotalExceptionQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalExceptionQuantity() int32 {
	if o == nil || IsNil(o.TotalExceptionQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalExceptionQuantity
}

// GetTotalExceptionQuantityOk returns a tuple with the TotalExceptionQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalExceptionQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalExceptionQuantity) {
		return nil, false
	}
	return o.TotalExceptionQuantity, true
}

// HasTotalExceptionQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalExceptionQuantity() bool {
	if o != nil && !IsNil(o.TotalExceptionQuantity) {
		return true
	}

	return false
}

// SetTotalExceptionQuantity gets a reference to the given int32 and assigns it to the TotalExceptionQuantity field.
func (o *Inventory) SetTotalExceptionQuantity(v int32) {
	o.TotalExceptionQuantity = &v
}

// GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalFulfillableQuantity() int32 {
	if o == nil || IsNil(o.TotalFulfillableQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalFulfillableQuantity
}

// GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalFulfillableQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalFulfillableQuantity) {
		return nil, false
	}
	return o.TotalFulfillableQuantity, true
}

// HasTotalFulfillableQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalFulfillableQuantity() bool {
	if o != nil && !IsNil(o.TotalFulfillableQuantity) {
		return true
	}

	return false
}

// SetTotalFulfillableQuantity gets a reference to the given int32 and assigns it to the TotalFulfillableQuantity field.
func (o *Inventory) SetTotalFulfillableQuantity(v int32) {
	o.TotalFulfillableQuantity = &v
}

// GetTotalInternalTransferQuantity returns the TotalInternalTransferQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalInternalTransferQuantity() int32 {
	if o == nil || IsNil(o.TotalInternalTransferQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalInternalTransferQuantity
}

// GetTotalInternalTransferQuantityOk returns a tuple with the TotalInternalTransferQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalInternalTransferQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalInternalTransferQuantity) {
		return nil, false
	}
	return o.TotalInternalTransferQuantity, true
}

// HasTotalInternalTransferQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalInternalTransferQuantity() bool {
	if o != nil && !IsNil(o.TotalInternalTransferQuantity) {
		return true
	}

	return false
}

// SetTotalInternalTransferQuantity gets a reference to the given int32 and assigns it to the TotalInternalTransferQuantity field.
func (o *Inventory) SetTotalInternalTransferQuantity(v int32) {
	o.TotalInternalTransferQuantity = &v
}

// GetTotalOnhandQuantity returns the TotalOnhandQuantity field value if set, zero value otherwise.
func (o *Inventory) GetTotalOnhandQuantity() int32 {
	if o == nil || IsNil(o.TotalOnhandQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalOnhandQuantity
}

// GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetTotalOnhandQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOnhandQuantity) {
		return nil, false
	}
	return o.TotalOnhandQuantity, true
}

// HasTotalOnhandQuantity returns a boolean if a field has been set.
func (o *Inventory) HasTotalOnhandQuantity() bool {
	if o != nil && !IsNil(o.TotalOnhandQuantity) {
		return true
	}

	return false
}

// SetTotalOnhandQuantity gets a reference to the given int32 and assigns it to the TotalOnhandQuantity field.
func (o *Inventory) SetTotalOnhandQuantity(v int32) {
	o.TotalOnhandQuantity = &v
}

func (o Inventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Inventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		toSerialize["fulfillable_quantity_by_fulfillment_center"] = o.FulfillableQuantityByFulfillmentCenter
	}
	if !IsNil(o.FulfillableQuantityByLot) {
		toSerialize["fulfillable_quantity_by_lot"] = o.FulfillableQuantityByLot
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsCasePick) {
		toSerialize["is_case_pick"] = o.IsCasePick
	}
	if !IsNil(o.IsDigital) {
		toSerialize["is_digital"] = o.IsDigital
	}
	if !IsNil(o.IsLot) {
		toSerialize["is_lot"] = o.IsLot
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PackagingAttribute) {
		toSerialize["packaging_attribute"] = o.PackagingAttribute
	}
	if !IsNil(o.TotalAwaitingQuantity) {
		toSerialize["total_awaiting_quantity"] = o.TotalAwaitingQuantity
	}
	if !IsNil(o.TotalBackorderedQuantity) {
		toSerialize["total_backordered_quantity"] = o.TotalBackorderedQuantity
	}
	if !IsNil(o.TotalCommittedQuantity) {
		toSerialize["total_committed_quantity"] = o.TotalCommittedQuantity
	}
	if !IsNil(o.TotalExceptionQuantity) {
		toSerialize["total_exception_quantity"] = o.TotalExceptionQuantity
	}
	if !IsNil(o.TotalFulfillableQuantity) {
		toSerialize["total_fulfillable_quantity"] = o.TotalFulfillableQuantity
	}
	if !IsNil(o.TotalInternalTransferQuantity) {
		toSerialize["total_internal_transfer_quantity"] = o.TotalInternalTransferQuantity
	}
	if !IsNil(o.TotalOnhandQuantity) {
		toSerialize["total_onhand_quantity"] = o.TotalOnhandQuantity
	}
	return toSerialize, nil
}

type NullableInventory struct {
	value *Inventory
	isSet bool
}

func (v NullableInventory) Get() *Inventory {
	return v.value
}

func (v *NullableInventory) Set(val *Inventory) {
	v.value = val
	v.isSet = true
}

func (v NullableInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventory(val *Inventory) *NullableInventory {
	return &NullableInventory{value: val, isSet: true}
}

func (v NullableInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


