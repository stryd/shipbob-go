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

// ShipbobInventoryApiViewModelsInventoryViewModel Information about an inventory item
type ShipbobInventoryApiViewModelsInventoryViewModel struct {
	Dimensions *ShipbobInventoryApiViewModelsDimensionViewModel `json:"dimensions,omitempty"`
	// Fulfillable quantity of this inventory item broken down by fulfillment center location
	FulfillableQuantityByFulfillmentCenter *[]ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel `json:"fulfillable_quantity_by_fulfillment_center,omitempty"`
	// Fulfillable quantity of this inventory item broken down by lot
	FulfillableQuantityByLot *[]ShipbobInventoryApiViewModelsLotQuantityViewModel `json:"fulfillable_quantity_by_lot,omitempty"`
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
	// The amount of the item you need to send to ShipBob to fulfill all exception orders containing \\r\\nthe item. This is the exception_quantity less the fulfillable_quantity of the item.
	TotalBackorderedQuantity *int32 `json:"total_backordered_quantity,omitempty"`
	// Total committed quantity of this inventory item
	TotalCommittedQuantity *int32 `json:"total_committed_quantity,omitempty"`
	// The total quantity of all items that are contained within orders that\\r\\nare in exception/out of stock status. Out of stock orders have not been\\r\\nprocessed and therefore do not have lot or fulfillment centers assigned.
	TotalExceptionQuantity *int32 `json:"total_exception_quantity,omitempty"`
	// Total fulfillable quantity of this inventory item
	TotalFulfillableQuantity *int32 `json:"total_fulfillable_quantity,omitempty"`
	// The total quantity of all items that are in the process of internal transit \\r\\nbetween ShipBob fulfillment centers. These items are not pickable or fulfillable\\r\\nuntil they have been received and moved to the \"On Hand\" quantity of the destination FC.\\r\\nInternal transit quantities for each FC represent the incoming transfer stock\\r\\nfor that specific location.
	TotalInternalTransferQuantity *int32 `json:"total_internal_transfer_quantity,omitempty"`
	// Total onhand quantity of this inventory item
	TotalOnhandQuantity *int32 `json:"total_onhand_quantity,omitempty"`
}

// NewShipbobInventoryApiViewModelsInventoryViewModel instantiates a new ShipbobInventoryApiViewModelsInventoryViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipbobInventoryApiViewModelsInventoryViewModel() *ShipbobInventoryApiViewModelsInventoryViewModel {
	this := ShipbobInventoryApiViewModelsInventoryViewModel{}
	return &this
}

// NewShipbobInventoryApiViewModelsInventoryViewModelWithDefaults instantiates a new ShipbobInventoryApiViewModelsInventoryViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipbobInventoryApiViewModelsInventoryViewModelWithDefaults() *ShipbobInventoryApiViewModelsInventoryViewModel {
	this := ShipbobInventoryApiViewModelsInventoryViewModel{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetDimensions() ShipbobInventoryApiViewModelsDimensionViewModel {
	if o == nil || o.Dimensions == nil {
		var ret ShipbobInventoryApiViewModelsDimensionViewModel
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetDimensionsOk() (*ShipbobInventoryApiViewModelsDimensionViewModel, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given ShipbobInventoryApiViewModelsDimensionViewModel and assigns it to the Dimensions field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetDimensions(v ShipbobInventoryApiViewModelsDimensionViewModel) {
	o.Dimensions = &v
}

// GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByFulfillmentCenter() []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel {
	if o == nil || o.FulfillableQuantityByFulfillmentCenter == nil {
		var ret []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel
		return ret
	}
	return *o.FulfillableQuantityByFulfillmentCenter
}

// GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByFulfillmentCenterOk() (*[]ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel, bool) {
	if o == nil || o.FulfillableQuantityByFulfillmentCenter == nil {
		return nil, false
	}
	return o.FulfillableQuantityByFulfillmentCenter, true
}

// HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasFulfillableQuantityByFulfillmentCenter() bool {
	if o != nil && o.FulfillableQuantityByFulfillmentCenter != nil {
		return true
	}

	return false
}

// SetFulfillableQuantityByFulfillmentCenter gets a reference to the given []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel and assigns it to the FulfillableQuantityByFulfillmentCenter field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetFulfillableQuantityByFulfillmentCenter(v []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel) {
	o.FulfillableQuantityByFulfillmentCenter = &v
}

// GetFulfillableQuantityByLot returns the FulfillableQuantityByLot field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByLot() []ShipbobInventoryApiViewModelsLotQuantityViewModel {
	if o == nil || o.FulfillableQuantityByLot == nil {
		var ret []ShipbobInventoryApiViewModelsLotQuantityViewModel
		return ret
	}
	return *o.FulfillableQuantityByLot
}

// GetFulfillableQuantityByLotOk returns a tuple with the FulfillableQuantityByLot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByLotOk() (*[]ShipbobInventoryApiViewModelsLotQuantityViewModel, bool) {
	if o == nil || o.FulfillableQuantityByLot == nil {
		return nil, false
	}
	return o.FulfillableQuantityByLot, true
}

// HasFulfillableQuantityByLot returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasFulfillableQuantityByLot() bool {
	if o != nil && o.FulfillableQuantityByLot != nil {
		return true
	}

	return false
}

// SetFulfillableQuantityByLot gets a reference to the given []ShipbobInventoryApiViewModelsLotQuantityViewModel and assigns it to the FulfillableQuantityByLot field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetFulfillableQuantityByLot(v []ShipbobInventoryApiViewModelsLotQuantityViewModel) {
	o.FulfillableQuantityByLot = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsCasePick returns the IsCasePick field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsCasePick() bool {
	if o == nil || o.IsCasePick == nil {
		var ret bool
		return ret
	}
	return *o.IsCasePick
}

// GetIsCasePickOk returns a tuple with the IsCasePick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsCasePickOk() (*bool, bool) {
	if o == nil || o.IsCasePick == nil {
		return nil, false
	}
	return o.IsCasePick, true
}

// HasIsCasePick returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsCasePick() bool {
	if o != nil && o.IsCasePick != nil {
		return true
	}

	return false
}

// SetIsCasePick gets a reference to the given bool and assigns it to the IsCasePick field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsCasePick(v bool) {
	o.IsCasePick = &v
}

// GetIsDigital returns the IsDigital field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsDigital() bool {
	if o == nil || o.IsDigital == nil {
		var ret bool
		return ret
	}
	return *o.IsDigital
}

// GetIsDigitalOk returns a tuple with the IsDigital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsDigitalOk() (*bool, bool) {
	if o == nil || o.IsDigital == nil {
		return nil, false
	}
	return o.IsDigital, true
}

// HasIsDigital returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsDigital() bool {
	if o != nil && o.IsDigital != nil {
		return true
	}

	return false
}

// SetIsDigital gets a reference to the given bool and assigns it to the IsDigital field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsDigital(v bool) {
	o.IsDigital = &v
}

// GetIsLot returns the IsLot field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsLot() bool {
	if o == nil || o.IsLot == nil {
		var ret bool
		return ret
	}
	return *o.IsLot
}

// GetIsLotOk returns a tuple with the IsLot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsLotOk() (*bool, bool) {
	if o == nil || o.IsLot == nil {
		return nil, false
	}
	return o.IsLot, true
}

// HasIsLot returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsLot() bool {
	if o != nil && o.IsLot != nil {
		return true
	}

	return false
}

// SetIsLot gets a reference to the given bool and assigns it to the IsLot field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsLot(v bool) {
	o.IsLot = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetName(v string) {
	o.Name = &v
}

// GetPackagingAttribute returns the PackagingAttribute field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetPackagingAttribute() string {
	if o == nil || o.PackagingAttribute == nil {
		var ret string
		return ret
	}
	return *o.PackagingAttribute
}

// GetPackagingAttributeOk returns a tuple with the PackagingAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetPackagingAttributeOk() (*string, bool) {
	if o == nil || o.PackagingAttribute == nil {
		return nil, false
	}
	return o.PackagingAttribute, true
}

// HasPackagingAttribute returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasPackagingAttribute() bool {
	if o != nil && o.PackagingAttribute != nil {
		return true
	}

	return false
}

// SetPackagingAttribute gets a reference to the given string and assigns it to the PackagingAttribute field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetPackagingAttribute(v string) {
	o.PackagingAttribute = &v
}

// GetTotalAwaitingQuantity returns the TotalAwaitingQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalAwaitingQuantity() int32 {
	if o == nil || o.TotalAwaitingQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalAwaitingQuantity
}

// GetTotalAwaitingQuantityOk returns a tuple with the TotalAwaitingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalAwaitingQuantityOk() (*int32, bool) {
	if o == nil || o.TotalAwaitingQuantity == nil {
		return nil, false
	}
	return o.TotalAwaitingQuantity, true
}

// HasTotalAwaitingQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalAwaitingQuantity() bool {
	if o != nil && o.TotalAwaitingQuantity != nil {
		return true
	}

	return false
}

// SetTotalAwaitingQuantity gets a reference to the given int32 and assigns it to the TotalAwaitingQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalAwaitingQuantity(v int32) {
	o.TotalAwaitingQuantity = &v
}

// GetTotalBackorderedQuantity returns the TotalBackorderedQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalBackorderedQuantity() int32 {
	if o == nil || o.TotalBackorderedQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalBackorderedQuantity
}

// GetTotalBackorderedQuantityOk returns a tuple with the TotalBackorderedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalBackorderedQuantityOk() (*int32, bool) {
	if o == nil || o.TotalBackorderedQuantity == nil {
		return nil, false
	}
	return o.TotalBackorderedQuantity, true
}

// HasTotalBackorderedQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalBackorderedQuantity() bool {
	if o != nil && o.TotalBackorderedQuantity != nil {
		return true
	}

	return false
}

// SetTotalBackorderedQuantity gets a reference to the given int32 and assigns it to the TotalBackorderedQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalBackorderedQuantity(v int32) {
	o.TotalBackorderedQuantity = &v
}

// GetTotalCommittedQuantity returns the TotalCommittedQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalCommittedQuantity() int32 {
	if o == nil || o.TotalCommittedQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalCommittedQuantity
}

// GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalCommittedQuantityOk() (*int32, bool) {
	if o == nil || o.TotalCommittedQuantity == nil {
		return nil, false
	}
	return o.TotalCommittedQuantity, true
}

// HasTotalCommittedQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalCommittedQuantity() bool {
	if o != nil && o.TotalCommittedQuantity != nil {
		return true
	}

	return false
}

// SetTotalCommittedQuantity gets a reference to the given int32 and assigns it to the TotalCommittedQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalCommittedQuantity(v int32) {
	o.TotalCommittedQuantity = &v
}

// GetTotalExceptionQuantity returns the TotalExceptionQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalExceptionQuantity() int32 {
	if o == nil || o.TotalExceptionQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalExceptionQuantity
}

// GetTotalExceptionQuantityOk returns a tuple with the TotalExceptionQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalExceptionQuantityOk() (*int32, bool) {
	if o == nil || o.TotalExceptionQuantity == nil {
		return nil, false
	}
	return o.TotalExceptionQuantity, true
}

// HasTotalExceptionQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalExceptionQuantity() bool {
	if o != nil && o.TotalExceptionQuantity != nil {
		return true
	}

	return false
}

// SetTotalExceptionQuantity gets a reference to the given int32 and assigns it to the TotalExceptionQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalExceptionQuantity(v int32) {
	o.TotalExceptionQuantity = &v
}

// GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalFulfillableQuantity() int32 {
	if o == nil || o.TotalFulfillableQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalFulfillableQuantity
}

// GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalFulfillableQuantityOk() (*int32, bool) {
	if o == nil || o.TotalFulfillableQuantity == nil {
		return nil, false
	}
	return o.TotalFulfillableQuantity, true
}

// HasTotalFulfillableQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalFulfillableQuantity() bool {
	if o != nil && o.TotalFulfillableQuantity != nil {
		return true
	}

	return false
}

// SetTotalFulfillableQuantity gets a reference to the given int32 and assigns it to the TotalFulfillableQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalFulfillableQuantity(v int32) {
	o.TotalFulfillableQuantity = &v
}

// GetTotalInternalTransferQuantity returns the TotalInternalTransferQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalInternalTransferQuantity() int32 {
	if o == nil || o.TotalInternalTransferQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalInternalTransferQuantity
}

// GetTotalInternalTransferQuantityOk returns a tuple with the TotalInternalTransferQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalInternalTransferQuantityOk() (*int32, bool) {
	if o == nil || o.TotalInternalTransferQuantity == nil {
		return nil, false
	}
	return o.TotalInternalTransferQuantity, true
}

// HasTotalInternalTransferQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalInternalTransferQuantity() bool {
	if o != nil && o.TotalInternalTransferQuantity != nil {
		return true
	}

	return false
}

// SetTotalInternalTransferQuantity gets a reference to the given int32 and assigns it to the TotalInternalTransferQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalInternalTransferQuantity(v int32) {
	o.TotalInternalTransferQuantity = &v
}

// GetTotalOnhandQuantity returns the TotalOnhandQuantity field value if set, zero value otherwise.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalOnhandQuantity() int32 {
	if o == nil || o.TotalOnhandQuantity == nil {
		var ret int32
		return ret
	}
	return *o.TotalOnhandQuantity
}

// GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalOnhandQuantityOk() (*int32, bool) {
	if o == nil || o.TotalOnhandQuantity == nil {
		return nil, false
	}
	return o.TotalOnhandQuantity, true
}

// HasTotalOnhandQuantity returns a boolean if a field has been set.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalOnhandQuantity() bool {
	if o != nil && o.TotalOnhandQuantity != nil {
		return true
	}

	return false
}

// SetTotalOnhandQuantity gets a reference to the given int32 and assigns it to the TotalOnhandQuantity field.
func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalOnhandQuantity(v int32) {
	o.TotalOnhandQuantity = &v
}

func (o ShipbobInventoryApiViewModelsInventoryViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dimensions != nil {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.FulfillableQuantityByFulfillmentCenter != nil {
		toSerialize["fulfillable_quantity_by_fulfillment_center"] = o.FulfillableQuantityByFulfillmentCenter
	}
	if o.FulfillableQuantityByLot != nil {
		toSerialize["fulfillable_quantity_by_lot"] = o.FulfillableQuantityByLot
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.IsCasePick != nil {
		toSerialize["is_case_pick"] = o.IsCasePick
	}
	if o.IsDigital != nil {
		toSerialize["is_digital"] = o.IsDigital
	}
	if o.IsLot != nil {
		toSerialize["is_lot"] = o.IsLot
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PackagingAttribute != nil {
		toSerialize["packaging_attribute"] = o.PackagingAttribute
	}
	if o.TotalAwaitingQuantity != nil {
		toSerialize["total_awaiting_quantity"] = o.TotalAwaitingQuantity
	}
	if o.TotalBackorderedQuantity != nil {
		toSerialize["total_backordered_quantity"] = o.TotalBackorderedQuantity
	}
	if o.TotalCommittedQuantity != nil {
		toSerialize["total_committed_quantity"] = o.TotalCommittedQuantity
	}
	if o.TotalExceptionQuantity != nil {
		toSerialize["total_exception_quantity"] = o.TotalExceptionQuantity
	}
	if o.TotalFulfillableQuantity != nil {
		toSerialize["total_fulfillable_quantity"] = o.TotalFulfillableQuantity
	}
	if o.TotalInternalTransferQuantity != nil {
		toSerialize["total_internal_transfer_quantity"] = o.TotalInternalTransferQuantity
	}
	if o.TotalOnhandQuantity != nil {
		toSerialize["total_onhand_quantity"] = o.TotalOnhandQuantity
	}
	return json.Marshal(toSerialize)
}

type NullableShipbobInventoryApiViewModelsInventoryViewModel struct {
	value *ShipbobInventoryApiViewModelsInventoryViewModel
	isSet bool
}

func (v NullableShipbobInventoryApiViewModelsInventoryViewModel) Get() *ShipbobInventoryApiViewModelsInventoryViewModel {
	return v.value
}

func (v *NullableShipbobInventoryApiViewModelsInventoryViewModel) Set(val *ShipbobInventoryApiViewModelsInventoryViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipbobInventoryApiViewModelsInventoryViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipbobInventoryApiViewModelsInventoryViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipbobInventoryApiViewModelsInventoryViewModel(val *ShipbobInventoryApiViewModelsInventoryViewModel) *NullableShipbobInventoryApiViewModelsInventoryViewModel {
	return &NullableShipbobInventoryApiViewModelsInventoryViewModel{value: val, isSet: true}
}

func (v NullableShipbobInventoryApiViewModelsInventoryViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipbobInventoryApiViewModelsInventoryViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


