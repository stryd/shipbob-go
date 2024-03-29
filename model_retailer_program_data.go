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

// checks if the RetailerProgramData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetailerProgramData{}

// RetailerProgramData Contains properties that needs to be used for fulfilling B2B/Dropship orders.
type RetailerProgramData struct {
	// Ship From - Certain retailers want to display the ship from address as their return facility, not Shipbob's warehouse address        ///  Mark For Address - Final destination address
	Addresses []RetailerProgramDataAddress `json:"addresses,omitempty"`
	// Customer Ticket Number
	CustomerTicketNumber *string `json:"customer_ticket_number,omitempty"`
	// Expected delivery date
	DeliveryDate NullableTime `json:"delivery_date,omitempty"`
	// Identifies a merchant's store department
	Department *string `json:"department,omitempty"`
	// Store Number
	MarkForStore *string `json:"mark_for_store,omitempty"`
	// First initial documentation sent from buyer to seller with item(s) and quantities.
	PurchaseOrderNumber string `json:"purchase_order_number"`
	// Identifies retailer-merchant combination
	RetailerProgramType string `json:"retailer_program_type"`
}

// NewRetailerProgramData instantiates a new RetailerProgramData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetailerProgramData(purchaseOrderNumber string, retailerProgramType string) *RetailerProgramData {
	this := RetailerProgramData{}
	this.PurchaseOrderNumber = purchaseOrderNumber
	this.RetailerProgramType = retailerProgramType
	return &this
}

// NewRetailerProgramDataWithDefaults instantiates a new RetailerProgramData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetailerProgramDataWithDefaults() *RetailerProgramData {
	this := RetailerProgramData{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *RetailerProgramData) GetAddresses() []RetailerProgramDataAddress {
	if o == nil || IsNil(o.Addresses) {
		var ret []RetailerProgramDataAddress
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetAddressesOk() ([]RetailerProgramDataAddress, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *RetailerProgramData) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []RetailerProgramDataAddress and assigns it to the Addresses field.
func (o *RetailerProgramData) SetAddresses(v []RetailerProgramDataAddress) {
	o.Addresses = v
}

// GetCustomerTicketNumber returns the CustomerTicketNumber field value if set, zero value otherwise.
func (o *RetailerProgramData) GetCustomerTicketNumber() string {
	if o == nil || IsNil(o.CustomerTicketNumber) {
		var ret string
		return ret
	}
	return *o.CustomerTicketNumber
}

// GetCustomerTicketNumberOk returns a tuple with the CustomerTicketNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetCustomerTicketNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerTicketNumber) {
		return nil, false
	}
	return o.CustomerTicketNumber, true
}

// HasCustomerTicketNumber returns a boolean if a field has been set.
func (o *RetailerProgramData) HasCustomerTicketNumber() bool {
	if o != nil && !IsNil(o.CustomerTicketNumber) {
		return true
	}

	return false
}

// SetCustomerTicketNumber gets a reference to the given string and assigns it to the CustomerTicketNumber field.
func (o *RetailerProgramData) SetCustomerTicketNumber(v string) {
	o.CustomerTicketNumber = &v
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetailerProgramData) GetDeliveryDate() time.Time {
	if o == nil || IsNil(o.DeliveryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate.Get()
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetailerProgramData) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryDate.Get(), o.DeliveryDate.IsSet()
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *RetailerProgramData) HasDeliveryDate() bool {
	if o != nil && o.DeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given NullableTime and assigns it to the DeliveryDate field.
func (o *RetailerProgramData) SetDeliveryDate(v time.Time) {
	o.DeliveryDate.Set(&v)
}

// SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil
func (o *RetailerProgramData) SetDeliveryDateNil() {
	o.DeliveryDate.Set(nil)
}

// UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
func (o *RetailerProgramData) UnsetDeliveryDate() {
	o.DeliveryDate.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *RetailerProgramData) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *RetailerProgramData) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *RetailerProgramData) SetDepartment(v string) {
	o.Department = &v
}

// GetMarkForStore returns the MarkForStore field value if set, zero value otherwise.
func (o *RetailerProgramData) GetMarkForStore() string {
	if o == nil || IsNil(o.MarkForStore) {
		var ret string
		return ret
	}
	return *o.MarkForStore
}

// GetMarkForStoreOk returns a tuple with the MarkForStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetMarkForStoreOk() (*string, bool) {
	if o == nil || IsNil(o.MarkForStore) {
		return nil, false
	}
	return o.MarkForStore, true
}

// HasMarkForStore returns a boolean if a field has been set.
func (o *RetailerProgramData) HasMarkForStore() bool {
	if o != nil && !IsNil(o.MarkForStore) {
		return true
	}

	return false
}

// SetMarkForStore gets a reference to the given string and assigns it to the MarkForStore field.
func (o *RetailerProgramData) SetMarkForStore(v string) {
	o.MarkForStore = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value
func (o *RetailerProgramData) GetPurchaseOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOrderNumber, true
}

// SetPurchaseOrderNumber sets field value
func (o *RetailerProgramData) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = v
}

// GetRetailerProgramType returns the RetailerProgramType field value
func (o *RetailerProgramData) GetRetailerProgramType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetailerProgramType
}

// GetRetailerProgramTypeOk returns a tuple with the RetailerProgramType field value
// and a boolean to check if the value has been set.
func (o *RetailerProgramData) GetRetailerProgramTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetailerProgramType, true
}

// SetRetailerProgramType sets field value
func (o *RetailerProgramData) SetRetailerProgramType(v string) {
	o.RetailerProgramType = v
}

func (o RetailerProgramData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetailerProgramData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.CustomerTicketNumber) {
		toSerialize["customer_ticket_number"] = o.CustomerTicketNumber
	}
	if o.DeliveryDate.IsSet() {
		toSerialize["delivery_date"] = o.DeliveryDate.Get()
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.MarkForStore) {
		toSerialize["mark_for_store"] = o.MarkForStore
	}
	toSerialize["purchase_order_number"] = o.PurchaseOrderNumber
	toSerialize["retailer_program_type"] = o.RetailerProgramType
	return toSerialize, nil
}

type NullableRetailerProgramData struct {
	value *RetailerProgramData
	isSet bool
}

func (v NullableRetailerProgramData) Get() *RetailerProgramData {
	return v.value
}

func (v *NullableRetailerProgramData) Set(val *RetailerProgramData) {
	v.value = val
	v.isSet = true
}

func (v NullableRetailerProgramData) IsSet() bool {
	return v.isSet
}

func (v *NullableRetailerProgramData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetailerProgramData(val *RetailerProgramData) *NullableRetailerProgramData {
	return &NullableRetailerProgramData{value: val, isSet: true}
}

func (v NullableRetailerProgramData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetailerProgramData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
