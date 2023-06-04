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

// checks if the AddProductToOrderByReferenceId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProductToOrderByReferenceId{}

// AddProductToOrderByReferenceId struct for AddProductToOrderByReferenceId
type AddProductToOrderByReferenceId struct {
	// Numeric assignment per item. Used as a reference number for multiple purposes such as split orders, split containers, etc.
	ExternalLineId NullableInt32 `json:"external_line_id,omitempty"`
	// Global Trade Item Number - unique and internationally recognized identifier assigned to item by company GS1
	Gtin *string `json:"gtin,omitempty"`
	// Name of the product. Required if there is not an existing ShipBob product with a matching reference_id value.
	Name *string `json:"name,omitempty"`
	// The quantity of this product ordered
	Quantity int32 `json:"quantity"`
	// Defined standard for measure for an item (each, inner pack, case, pallet).  Values: EA, INP, CS and PL
	QuantityUnitOfMeasureCode *string `json:"quantity_unit_of_measure_code,omitempty"`
	// Unique reference id of the product
	ReferenceId string `json:"reference_id"`
	// Product SKU
	Sku *string `json:"sku,omitempty"`
	// Price for one item
	UnitPrice NullableFloat64 `json:"unit_price,omitempty"`
	// Universal Product Code - Unique external identifier
	Upc *string `json:"upc,omitempty"`
}

// NewAddProductToOrderByReferenceId instantiates a new AddProductToOrderByReferenceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProductToOrderByReferenceId(quantity int32, referenceId string) *AddProductToOrderByReferenceId {
	this := AddProductToOrderByReferenceId{}
	this.Quantity = quantity
	this.ReferenceId = referenceId
	return &this
}

// NewAddProductToOrderByReferenceIdWithDefaults instantiates a new AddProductToOrderByReferenceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProductToOrderByReferenceIdWithDefaults() *AddProductToOrderByReferenceId {
	this := AddProductToOrderByReferenceId{}
	return &this
}

// GetExternalLineId returns the ExternalLineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddProductToOrderByReferenceId) GetExternalLineId() int32 {
	if o == nil || IsNil(o.ExternalLineId.Get()) {
		var ret int32
		return ret
	}
	return *o.ExternalLineId.Get()
}

// GetExternalLineIdOk returns a tuple with the ExternalLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddProductToOrderByReferenceId) GetExternalLineIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalLineId.Get(), o.ExternalLineId.IsSet()
}

// HasExternalLineId returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasExternalLineId() bool {
	if o != nil && o.ExternalLineId.IsSet() {
		return true
	}

	return false
}

// SetExternalLineId gets a reference to the given NullableInt32 and assigns it to the ExternalLineId field.
func (o *AddProductToOrderByReferenceId) SetExternalLineId(v int32) {
	o.ExternalLineId.Set(&v)
}
// SetExternalLineIdNil sets the value for ExternalLineId to be an explicit nil
func (o *AddProductToOrderByReferenceId) SetExternalLineIdNil() {
	o.ExternalLineId.Set(nil)
}

// UnsetExternalLineId ensures that no value is present for ExternalLineId, not even an explicit nil
func (o *AddProductToOrderByReferenceId) UnsetExternalLineId() {
	o.ExternalLineId.Unset()
}

// GetGtin returns the Gtin field value if set, zero value otherwise.
func (o *AddProductToOrderByReferenceId) GetGtin() string {
	if o == nil || IsNil(o.Gtin) {
		var ret string
		return ret
	}
	return *o.Gtin
}

// GetGtinOk returns a tuple with the Gtin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetGtinOk() (*string, bool) {
	if o == nil || IsNil(o.Gtin) {
		return nil, false
	}
	return o.Gtin, true
}

// HasGtin returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasGtin() bool {
	if o != nil && !IsNil(o.Gtin) {
		return true
	}

	return false
}

// SetGtin gets a reference to the given string and assigns it to the Gtin field.
func (o *AddProductToOrderByReferenceId) SetGtin(v string) {
	o.Gtin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddProductToOrderByReferenceId) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddProductToOrderByReferenceId) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value
func (o *AddProductToOrderByReferenceId) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *AddProductToOrderByReferenceId) SetQuantity(v int32) {
	o.Quantity = v
}

// GetQuantityUnitOfMeasureCode returns the QuantityUnitOfMeasureCode field value if set, zero value otherwise.
func (o *AddProductToOrderByReferenceId) GetQuantityUnitOfMeasureCode() string {
	if o == nil || IsNil(o.QuantityUnitOfMeasureCode) {
		var ret string
		return ret
	}
	return *o.QuantityUnitOfMeasureCode
}

// GetQuantityUnitOfMeasureCodeOk returns a tuple with the QuantityUnitOfMeasureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetQuantityUnitOfMeasureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.QuantityUnitOfMeasureCode) {
		return nil, false
	}
	return o.QuantityUnitOfMeasureCode, true
}

// HasQuantityUnitOfMeasureCode returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasQuantityUnitOfMeasureCode() bool {
	if o != nil && !IsNil(o.QuantityUnitOfMeasureCode) {
		return true
	}

	return false
}

// SetQuantityUnitOfMeasureCode gets a reference to the given string and assigns it to the QuantityUnitOfMeasureCode field.
func (o *AddProductToOrderByReferenceId) SetQuantityUnitOfMeasureCode(v string) {
	o.QuantityUnitOfMeasureCode = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *AddProductToOrderByReferenceId) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *AddProductToOrderByReferenceId) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AddProductToOrderByReferenceId) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AddProductToOrderByReferenceId) SetSku(v string) {
	o.Sku = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddProductToOrderByReferenceId) GetUnitPrice() float64 {
	if o == nil || IsNil(o.UnitPrice.Get()) {
		var ret float64
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddProductToOrderByReferenceId) GetUnitPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat64 and assigns it to the UnitPrice field.
func (o *AddProductToOrderByReferenceId) SetUnitPrice(v float64) {
	o.UnitPrice.Set(&v)
}
// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *AddProductToOrderByReferenceId) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *AddProductToOrderByReferenceId) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetUpc returns the Upc field value if set, zero value otherwise.
func (o *AddProductToOrderByReferenceId) GetUpc() string {
	if o == nil || IsNil(o.Upc) {
		var ret string
		return ret
	}
	return *o.Upc
}

// GetUpcOk returns a tuple with the Upc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductToOrderByReferenceId) GetUpcOk() (*string, bool) {
	if o == nil || IsNil(o.Upc) {
		return nil, false
	}
	return o.Upc, true
}

// HasUpc returns a boolean if a field has been set.
func (o *AddProductToOrderByReferenceId) HasUpc() bool {
	if o != nil && !IsNil(o.Upc) {
		return true
	}

	return false
}

// SetUpc gets a reference to the given string and assigns it to the Upc field.
func (o *AddProductToOrderByReferenceId) SetUpc(v string) {
	o.Upc = &v
}

func (o AddProductToOrderByReferenceId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProductToOrderByReferenceId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalLineId.IsSet() {
		toSerialize["external_line_id"] = o.ExternalLineId.Get()
	}
	if !IsNil(o.Gtin) {
		toSerialize["gtin"] = o.Gtin
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.QuantityUnitOfMeasureCode) {
		toSerialize["quantity_unit_of_measure_code"] = o.QuantityUnitOfMeasureCode
	}
	toSerialize["reference_id"] = o.ReferenceId
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unit_price"] = o.UnitPrice.Get()
	}
	if !IsNil(o.Upc) {
		toSerialize["upc"] = o.Upc
	}
	return toSerialize, nil
}

type NullableAddProductToOrderByReferenceId struct {
	value *AddProductToOrderByReferenceId
	isSet bool
}

func (v NullableAddProductToOrderByReferenceId) Get() *AddProductToOrderByReferenceId {
	return v.value
}

func (v *NullableAddProductToOrderByReferenceId) Set(val *AddProductToOrderByReferenceId) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProductToOrderByReferenceId) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProductToOrderByReferenceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProductToOrderByReferenceId(val *AddProductToOrderByReferenceId) *NullableAddProductToOrderByReferenceId {
	return &NullableAddProductToOrderByReferenceId{value: val, isSet: true}
}

func (v NullableAddProductToOrderByReferenceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProductToOrderByReferenceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


