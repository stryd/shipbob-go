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

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product struct for Product
type Product struct {
	// Barcode for the product
	Barcode               NullableString         `json:"barcode,omitempty"`
	BundleRootInformation *BundleRootInformation `json:"bundle_root_information,omitempty"`
	Channel               *ProductChannel        `json:"channel,omitempty"`
	// Date the product was created
	CreatedDate NullableTime `json:"created_date,omitempty"`
	// The inventory that this product will resolve to when packing a shipment
	FulfillableInventoryItems []ProductInventoryItem `json:"fulfillable_inventory_items,omitempty"`
	// Fulfillable quantity of this product broken down by fulfillment center location
	FulfillableQuantityByFulfillmentCenter []ProductFulfillmentCenterQuantity `json:"fulfillable_quantity_by_fulfillment_center,omitempty"`
	// Global Trade Item Number - unique and internationally recognized identifier assigned to item by company GS1.
	Gtin NullableString `json:"gtin,omitempty"`
	// Unique identifier of the product
	Id *int32 `json:"id,omitempty"`
	// The name of the product
	Name NullableString `json:"name,omitempty"`
	// Unique reference identifier of the product
	ReferenceId NullableString `json:"reference_id,omitempty"`
	// Stock keeping unit for the product
	Sku NullableString `json:"sku,omitempty"`
	// Total committed quantity of this product
	TotalCommittedQuantity *int32 `json:"total_committed_quantity,omitempty"`
	// Total fulfillable quantity of this product
	TotalFulfillableQuantity *int32 `json:"total_fulfillable_quantity,omitempty"`
	// Total on hand quantity of this product
	TotalOnhandQuantity *int32 `json:"total_onhand_quantity,omitempty"`
	// The price of one unit
	UnitPrice NullableFloat64 `json:"unit_price,omitempty"`
	// Universal Product Code - Unique external identifier
	Upc NullableString `json:"upc,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetBarcode() string {
	if o == nil || IsNil(o.Barcode.Get()) {
		var ret string
		return ret
	}
	return *o.Barcode.Get()
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetBarcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Barcode.Get(), o.Barcode.IsSet()
}

// HasBarcode returns a boolean if a field has been set.
func (o *Product) HasBarcode() bool {
	if o != nil && o.Barcode.IsSet() {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given NullableString and assigns it to the Barcode field.
func (o *Product) SetBarcode(v string) {
	o.Barcode.Set(&v)
}

// SetBarcodeNil sets the value for Barcode to be an explicit nil
func (o *Product) SetBarcodeNil() {
	o.Barcode.Set(nil)
}

// UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
func (o *Product) UnsetBarcode() {
	o.Barcode.Unset()
}

// GetBundleRootInformation returns the BundleRootInformation field value if set, zero value otherwise.
func (o *Product) GetBundleRootInformation() BundleRootInformation {
	if o == nil || IsNil(o.BundleRootInformation) {
		var ret BundleRootInformation
		return ret
	}
	return *o.BundleRootInformation
}

// GetBundleRootInformationOk returns a tuple with the BundleRootInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetBundleRootInformationOk() (*BundleRootInformation, bool) {
	if o == nil || IsNil(o.BundleRootInformation) {
		return nil, false
	}
	return o.BundleRootInformation, true
}

// HasBundleRootInformation returns a boolean if a field has been set.
func (o *Product) HasBundleRootInformation() bool {
	if o != nil && !IsNil(o.BundleRootInformation) {
		return true
	}

	return false
}

// SetBundleRootInformation gets a reference to the given BundleRootInformation and assigns it to the BundleRootInformation field.
func (o *Product) SetBundleRootInformation(v BundleRootInformation) {
	o.BundleRootInformation = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *Product) GetChannel() ProductChannel {
	if o == nil || IsNil(o.Channel) {
		var ret ProductChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetChannelOk() (*ProductChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *Product) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ProductChannel and assigns it to the Channel field.
func (o *Product) SetChannel(v ProductChannel) {
	o.Channel = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Product) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *Product) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}

// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *Product) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *Product) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetFulfillableInventoryItems returns the FulfillableInventoryItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetFulfillableInventoryItems() []ProductInventoryItem {
	if o == nil {
		var ret []ProductInventoryItem
		return ret
	}
	return o.FulfillableInventoryItems
}

// GetFulfillableInventoryItemsOk returns a tuple with the FulfillableInventoryItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetFulfillableInventoryItemsOk() ([]ProductInventoryItem, bool) {
	if o == nil || IsNil(o.FulfillableInventoryItems) {
		return nil, false
	}
	return o.FulfillableInventoryItems, true
}

// HasFulfillableInventoryItems returns a boolean if a field has been set.
func (o *Product) HasFulfillableInventoryItems() bool {
	if o != nil && IsNil(o.FulfillableInventoryItems) {
		return true
	}

	return false
}

// SetFulfillableInventoryItems gets a reference to the given []ProductInventoryItem and assigns it to the FulfillableInventoryItems field.
func (o *Product) SetFulfillableInventoryItems(v []ProductInventoryItem) {
	o.FulfillableInventoryItems = v
}

// GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetFulfillableQuantityByFulfillmentCenter() []ProductFulfillmentCenterQuantity {
	if o == nil {
		var ret []ProductFulfillmentCenterQuantity
		return ret
	}
	return o.FulfillableQuantityByFulfillmentCenter
}

// GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetFulfillableQuantityByFulfillmentCenterOk() ([]ProductFulfillmentCenterQuantity, bool) {
	if o == nil || IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		return nil, false
	}
	return o.FulfillableQuantityByFulfillmentCenter, true
}

// HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.
func (o *Product) HasFulfillableQuantityByFulfillmentCenter() bool {
	if o != nil && IsNil(o.FulfillableQuantityByFulfillmentCenter) {
		return true
	}

	return false
}

// SetFulfillableQuantityByFulfillmentCenter gets a reference to the given []ProductFulfillmentCenterQuantity and assigns it to the FulfillableQuantityByFulfillmentCenter field.
func (o *Product) SetFulfillableQuantityByFulfillmentCenter(v []ProductFulfillmentCenterQuantity) {
	o.FulfillableQuantityByFulfillmentCenter = v
}

// GetGtin returns the Gtin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetGtin() string {
	if o == nil || IsNil(o.Gtin.Get()) {
		var ret string
		return ret
	}
	return *o.Gtin.Get()
}

// GetGtinOk returns a tuple with the Gtin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetGtinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gtin.Get(), o.Gtin.IsSet()
}

// HasGtin returns a boolean if a field has been set.
func (o *Product) HasGtin() bool {
	if o != nil && o.Gtin.IsSet() {
		return true
	}

	return false
}

// SetGtin gets a reference to the given NullableString and assigns it to the Gtin field.
func (o *Product) SetGtin(v string) {
	o.Gtin.Set(&v)
}

// SetGtinNil sets the value for Gtin to be an explicit nil
func (o *Product) SetGtinNil() {
	o.Gtin.Set(nil)
}

// UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
func (o *Product) UnsetGtin() {
	o.Gtin.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Product) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Product) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Product) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Product) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Product) UnsetName() {
	o.Name.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *Product) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *Product) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *Product) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *Product) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *Product) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *Product) SetSku(v string) {
	o.Sku.Set(&v)
}

// SetSkuNil sets the value for Sku to be an explicit nil
func (o *Product) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *Product) UnsetSku() {
	o.Sku.Unset()
}

// GetTotalCommittedQuantity returns the TotalCommittedQuantity field value if set, zero value otherwise.
func (o *Product) GetTotalCommittedQuantity() int32 {
	if o == nil || IsNil(o.TotalCommittedQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalCommittedQuantity
}

// GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTotalCommittedQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCommittedQuantity) {
		return nil, false
	}
	return o.TotalCommittedQuantity, true
}

// HasTotalCommittedQuantity returns a boolean if a field has been set.
func (o *Product) HasTotalCommittedQuantity() bool {
	if o != nil && !IsNil(o.TotalCommittedQuantity) {
		return true
	}

	return false
}

// SetTotalCommittedQuantity gets a reference to the given int32 and assigns it to the TotalCommittedQuantity field.
func (o *Product) SetTotalCommittedQuantity(v int32) {
	o.TotalCommittedQuantity = &v
}

// GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field value if set, zero value otherwise.
func (o *Product) GetTotalFulfillableQuantity() int32 {
	if o == nil || IsNil(o.TotalFulfillableQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalFulfillableQuantity
}

// GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTotalFulfillableQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalFulfillableQuantity) {
		return nil, false
	}
	return o.TotalFulfillableQuantity, true
}

// HasTotalFulfillableQuantity returns a boolean if a field has been set.
func (o *Product) HasTotalFulfillableQuantity() bool {
	if o != nil && !IsNil(o.TotalFulfillableQuantity) {
		return true
	}

	return false
}

// SetTotalFulfillableQuantity gets a reference to the given int32 and assigns it to the TotalFulfillableQuantity field.
func (o *Product) SetTotalFulfillableQuantity(v int32) {
	o.TotalFulfillableQuantity = &v
}

// GetTotalOnhandQuantity returns the TotalOnhandQuantity field value if set, zero value otherwise.
func (o *Product) GetTotalOnhandQuantity() int32 {
	if o == nil || IsNil(o.TotalOnhandQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalOnhandQuantity
}

// GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTotalOnhandQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOnhandQuantity) {
		return nil, false
	}
	return o.TotalOnhandQuantity, true
}

// HasTotalOnhandQuantity returns a boolean if a field has been set.
func (o *Product) HasTotalOnhandQuantity() bool {
	if o != nil && !IsNil(o.TotalOnhandQuantity) {
		return true
	}

	return false
}

// SetTotalOnhandQuantity gets a reference to the given int32 and assigns it to the TotalOnhandQuantity field.
func (o *Product) SetTotalOnhandQuantity(v int32) {
	o.TotalOnhandQuantity = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetUnitPrice() float64 {
	if o == nil || IsNil(o.UnitPrice.Get()) {
		var ret float64
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetUnitPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *Product) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat64 and assigns it to the UnitPrice field.
func (o *Product) SetUnitPrice(v float64) {
	o.UnitPrice.Set(&v)
}

// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *Product) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *Product) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetUpc returns the Upc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetUpc() string {
	if o == nil || IsNil(o.Upc.Get()) {
		var ret string
		return ret
	}
	return *o.Upc.Get()
}

// GetUpcOk returns a tuple with the Upc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetUpcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Upc.Get(), o.Upc.IsSet()
}

// HasUpc returns a boolean if a field has been set.
func (o *Product) HasUpc() bool {
	if o != nil && o.Upc.IsSet() {
		return true
	}

	return false
}

// SetUpc gets a reference to the given NullableString and assigns it to the Upc field.
func (o *Product) SetUpc(v string) {
	o.Upc.Set(&v)
}

// SetUpcNil sets the value for Upc to be an explicit nil
func (o *Product) SetUpcNil() {
	o.Upc.Set(nil)
}

// UnsetUpc ensures that no value is present for Upc, not even an explicit nil
func (o *Product) UnsetUpc() {
	o.Upc.Unset()
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Barcode.IsSet() {
		toSerialize["barcode"] = o.Barcode.Get()
	}
	if !IsNil(o.BundleRootInformation) {
		toSerialize["bundle_root_information"] = o.BundleRootInformation
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if o.CreatedDate.IsSet() {
		toSerialize["created_date"] = o.CreatedDate.Get()
	}
	if o.FulfillableInventoryItems != nil {
		toSerialize["fulfillable_inventory_items"] = o.FulfillableInventoryItems
	}
	if o.FulfillableQuantityByFulfillmentCenter != nil {
		toSerialize["fulfillable_quantity_by_fulfillment_center"] = o.FulfillableQuantityByFulfillmentCenter
	}
	if o.Gtin.IsSet() {
		toSerialize["gtin"] = o.Gtin.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	if !IsNil(o.TotalCommittedQuantity) {
		toSerialize["total_committed_quantity"] = o.TotalCommittedQuantity
	}
	if !IsNil(o.TotalFulfillableQuantity) {
		toSerialize["total_fulfillable_quantity"] = o.TotalFulfillableQuantity
	}
	if !IsNil(o.TotalOnhandQuantity) {
		toSerialize["total_onhand_quantity"] = o.TotalOnhandQuantity
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unit_price"] = o.UnitPrice.Get()
	}
	if o.Upc.IsSet() {
		toSerialize["upc"] = o.Upc.Get()
	}
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
