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

// checks if the ReturnOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnOrder{}

// ReturnOrder struct for ReturnOrder
type ReturnOrder struct {
	Channel *ReturnChannelInfo `json:"channel,omitempty"`
	// Completed date for a return order.
	CompletedDate NullableTime `json:"completed_date,omitempty"`
	// Customer name from the related shipment.
	CustomerName      NullableString           `json:"customer_name,omitempty"`
	FulfillmentCenter *ReturnFulfillmentCenter `json:"fulfillment_center,omitempty"`
	// Unique id of the Return Order
	Id *int32 `json:"id,omitempty"`
	// Date this return order was created
	InsertDate NullableTime `json:"insert_date,omitempty"`
	// List of inventory included in the return order
	Inventory []ReturnInventoryItem `json:"inventory,omitempty"`
	// Invoiced amount of return order (sum of transaction amounts)
	InvoiceAmount NullableFloat64 `json:"invoice_amount,omitempty"`
	// Id of the corresponding shipment that is the souce of the return
	OriginalShipmentId NullableInt32 `json:"original_shipment_id,omitempty"`
	// Client-defined external unique id of the return order
	ReferenceId NullableString `json:"reference_id,omitempty"`
	ReturnType  *ReturnType    `json:"return_type,omitempty"`
	Status      *ReturnStatus  `json:"status,omitempty"`
	// Store order for the related shipment.
	StoreOrderId NullableString `json:"store_order_id,omitempty"`
	// Tracking number of the return shipment
	TrackingNumber NullableString `json:"tracking_number,omitempty"`
	// Array of transactions affiliated with the return order
	Transactions []ReturnTransaction `json:"transactions,omitempty"`
}

// NewReturnOrder instantiates a new ReturnOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnOrder() *ReturnOrder {
	this := ReturnOrder{}
	return &this
}

// NewReturnOrderWithDefaults instantiates a new ReturnOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnOrderWithDefaults() *ReturnOrder {
	this := ReturnOrder{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ReturnOrder) GetChannel() ReturnChannelInfo {
	if o == nil || IsNil(o.Channel) {
		var ret ReturnChannelInfo
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrder) GetChannelOk() (*ReturnChannelInfo, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ReturnOrder) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ReturnChannelInfo and assigns it to the Channel field.
func (o *ReturnOrder) SetChannel(v ReturnChannelInfo) {
	o.Channel = &v
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetCompletedDate() time.Time {
	if o == nil || IsNil(o.CompletedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedDate.Get()
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedDate.Get(), o.CompletedDate.IsSet()
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *ReturnOrder) HasCompletedDate() bool {
	if o != nil && o.CompletedDate.IsSet() {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given NullableTime and assigns it to the CompletedDate field.
func (o *ReturnOrder) SetCompletedDate(v time.Time) {
	o.CompletedDate.Set(&v)
}

// SetCompletedDateNil sets the value for CompletedDate to be an explicit nil
func (o *ReturnOrder) SetCompletedDateNil() {
	o.CompletedDate.Set(nil)
}

// UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
func (o *ReturnOrder) UnsetCompletedDate() {
	o.CompletedDate.Unset()
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetCustomerName() string {
	if o == nil || IsNil(o.CustomerName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerName.Get()
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerName.Get(), o.CustomerName.IsSet()
}

// HasCustomerName returns a boolean if a field has been set.
func (o *ReturnOrder) HasCustomerName() bool {
	if o != nil && o.CustomerName.IsSet() {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given NullableString and assigns it to the CustomerName field.
func (o *ReturnOrder) SetCustomerName(v string) {
	o.CustomerName.Set(&v)
}

// SetCustomerNameNil sets the value for CustomerName to be an explicit nil
func (o *ReturnOrder) SetCustomerNameNil() {
	o.CustomerName.Set(nil)
}

// UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
func (o *ReturnOrder) UnsetCustomerName() {
	o.CustomerName.Unset()
}

// GetFulfillmentCenter returns the FulfillmentCenter field value if set, zero value otherwise.
func (o *ReturnOrder) GetFulfillmentCenter() ReturnFulfillmentCenter {
	if o == nil || IsNil(o.FulfillmentCenter) {
		var ret ReturnFulfillmentCenter
		return ret
	}
	return *o.FulfillmentCenter
}

// GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrder) GetFulfillmentCenterOk() (*ReturnFulfillmentCenter, bool) {
	if o == nil || IsNil(o.FulfillmentCenter) {
		return nil, false
	}
	return o.FulfillmentCenter, true
}

// HasFulfillmentCenter returns a boolean if a field has been set.
func (o *ReturnOrder) HasFulfillmentCenter() bool {
	if o != nil && !IsNil(o.FulfillmentCenter) {
		return true
	}

	return false
}

// SetFulfillmentCenter gets a reference to the given ReturnFulfillmentCenter and assigns it to the FulfillmentCenter field.
func (o *ReturnOrder) SetFulfillmentCenter(v ReturnFulfillmentCenter) {
	o.FulfillmentCenter = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReturnOrder) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrder) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReturnOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReturnOrder) SetId(v int32) {
	o.Id = &v
}

// GetInsertDate returns the InsertDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetInsertDate() time.Time {
	if o == nil || IsNil(o.InsertDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.InsertDate.Get()
}

// GetInsertDateOk returns a tuple with the InsertDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetInsertDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsertDate.Get(), o.InsertDate.IsSet()
}

// HasInsertDate returns a boolean if a field has been set.
func (o *ReturnOrder) HasInsertDate() bool {
	if o != nil && o.InsertDate.IsSet() {
		return true
	}

	return false
}

// SetInsertDate gets a reference to the given NullableTime and assigns it to the InsertDate field.
func (o *ReturnOrder) SetInsertDate(v time.Time) {
	o.InsertDate.Set(&v)
}

// SetInsertDateNil sets the value for InsertDate to be an explicit nil
func (o *ReturnOrder) SetInsertDateNil() {
	o.InsertDate.Set(nil)
}

// UnsetInsertDate ensures that no value is present for InsertDate, not even an explicit nil
func (o *ReturnOrder) UnsetInsertDate() {
	o.InsertDate.Unset()
}

// GetInventory returns the Inventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetInventory() []ReturnInventoryItem {
	if o == nil {
		var ret []ReturnInventoryItem
		return ret
	}
	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetInventoryOk() ([]ReturnInventoryItem, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *ReturnOrder) HasInventory() bool {
	if o != nil && IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given []ReturnInventoryItem and assigns it to the Inventory field.
func (o *ReturnOrder) SetInventory(v []ReturnInventoryItem) {
	o.Inventory = v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetInvoiceAmount() float64 {
	if o == nil || IsNil(o.InvoiceAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.InvoiceAmount.Get()
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetInvoiceAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceAmount.Get(), o.InvoiceAmount.IsSet()
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *ReturnOrder) HasInvoiceAmount() bool {
	if o != nil && o.InvoiceAmount.IsSet() {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given NullableFloat64 and assigns it to the InvoiceAmount field.
func (o *ReturnOrder) SetInvoiceAmount(v float64) {
	o.InvoiceAmount.Set(&v)
}

// SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil
func (o *ReturnOrder) SetInvoiceAmountNil() {
	o.InvoiceAmount.Set(nil)
}

// UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
func (o *ReturnOrder) UnsetInvoiceAmount() {
	o.InvoiceAmount.Unset()
}

// GetOriginalShipmentId returns the OriginalShipmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetOriginalShipmentId() int32 {
	if o == nil || IsNil(o.OriginalShipmentId.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalShipmentId.Get()
}

// GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetOriginalShipmentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalShipmentId.Get(), o.OriginalShipmentId.IsSet()
}

// HasOriginalShipmentId returns a boolean if a field has been set.
func (o *ReturnOrder) HasOriginalShipmentId() bool {
	if o != nil && o.OriginalShipmentId.IsSet() {
		return true
	}

	return false
}

// SetOriginalShipmentId gets a reference to the given NullableInt32 and assigns it to the OriginalShipmentId field.
func (o *ReturnOrder) SetOriginalShipmentId(v int32) {
	o.OriginalShipmentId.Set(&v)
}

// SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil
func (o *ReturnOrder) SetOriginalShipmentIdNil() {
	o.OriginalShipmentId.Set(nil)
}

// UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
func (o *ReturnOrder) UnsetOriginalShipmentId() {
	o.OriginalShipmentId.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ReturnOrder) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *ReturnOrder) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *ReturnOrder) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *ReturnOrder) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *ReturnOrder) GetReturnType() ReturnType {
	if o == nil || IsNil(o.ReturnType) {
		var ret ReturnType
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrder) GetReturnTypeOk() (*ReturnType, bool) {
	if o == nil || IsNil(o.ReturnType) {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *ReturnOrder) HasReturnType() bool {
	if o != nil && !IsNil(o.ReturnType) {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given ReturnType and assigns it to the ReturnType field.
func (o *ReturnOrder) SetReturnType(v ReturnType) {
	o.ReturnType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReturnOrder) GetStatus() ReturnStatus {
	if o == nil || IsNil(o.Status) {
		var ret ReturnStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrder) GetStatusOk() (*ReturnStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReturnOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReturnStatus and assigns it to the Status field.
func (o *ReturnOrder) SetStatus(v ReturnStatus) {
	o.Status = &v
}

// GetStoreOrderId returns the StoreOrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetStoreOrderId() string {
	if o == nil || IsNil(o.StoreOrderId.Get()) {
		var ret string
		return ret
	}
	return *o.StoreOrderId.Get()
}

// GetStoreOrderIdOk returns a tuple with the StoreOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetStoreOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreOrderId.Get(), o.StoreOrderId.IsSet()
}

// HasStoreOrderId returns a boolean if a field has been set.
func (o *ReturnOrder) HasStoreOrderId() bool {
	if o != nil && o.StoreOrderId.IsSet() {
		return true
	}

	return false
}

// SetStoreOrderId gets a reference to the given NullableString and assigns it to the StoreOrderId field.
func (o *ReturnOrder) SetStoreOrderId(v string) {
	o.StoreOrderId.Set(&v)
}

// SetStoreOrderIdNil sets the value for StoreOrderId to be an explicit nil
func (o *ReturnOrder) SetStoreOrderIdNil() {
	o.StoreOrderId.Set(nil)
}

// UnsetStoreOrderId ensures that no value is present for StoreOrderId, not even an explicit nil
func (o *ReturnOrder) UnsetStoreOrderId() {
	o.StoreOrderId.Unset()
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ReturnOrder) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber.IsSet() {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given NullableString and assigns it to the TrackingNumber field.
func (o *ReturnOrder) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}

// SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil
func (o *ReturnOrder) SetTrackingNumberNil() {
	o.TrackingNumber.Set(nil)
}

// UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
func (o *ReturnOrder) UnsetTrackingNumber() {
	o.TrackingNumber.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnOrder) GetTransactions() []ReturnTransaction {
	if o == nil {
		var ret []ReturnTransaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnOrder) GetTransactionsOk() ([]ReturnTransaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *ReturnOrder) HasTransactions() bool {
	if o != nil && IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []ReturnTransaction and assigns it to the Transactions field.
func (o *ReturnOrder) SetTransactions(v []ReturnTransaction) {
	o.Transactions = v
}

func (o ReturnOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if o.CompletedDate.IsSet() {
		toSerialize["completed_date"] = o.CompletedDate.Get()
	}
	if o.CustomerName.IsSet() {
		toSerialize["customer_name"] = o.CustomerName.Get()
	}
	if !IsNil(o.FulfillmentCenter) {
		toSerialize["fulfillment_center"] = o.FulfillmentCenter
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.InsertDate.IsSet() {
		toSerialize["insert_date"] = o.InsertDate.Get()
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.InvoiceAmount.IsSet() {
		toSerialize["invoice_amount"] = o.InvoiceAmount.Get()
	}
	if o.OriginalShipmentId.IsSet() {
		toSerialize["original_shipment_id"] = o.OriginalShipmentId.Get()
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if !IsNil(o.ReturnType) {
		toSerialize["return_type"] = o.ReturnType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.StoreOrderId.IsSet() {
		toSerialize["store_order_id"] = o.StoreOrderId.Get()
	}
	if o.TrackingNumber.IsSet() {
		toSerialize["tracking_number"] = o.TrackingNumber.Get()
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return toSerialize, nil
}

type NullableReturnOrder struct {
	value *ReturnOrder
	isSet bool
}

func (v NullableReturnOrder) Get() *ReturnOrder {
	return v.value
}

func (v *NullableReturnOrder) Set(val *ReturnOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnOrder(val *ReturnOrder) *NullableReturnOrder {
	return &NullableReturnOrder{value: val, isSet: true}
}

func (v NullableReturnOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
