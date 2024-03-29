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

// checks if the CreateOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrder{}

// CreateOrder struct for CreateOrder
type CreateOrder struct {
	Financials *Financials `json:"financials,omitempty"`
	// Gift message associated with the order
	GiftMessage *string `json:"gift_message,omitempty"`
	// Desired Fulfillment Center Location ID. If not specified, ShipBob will determine the location that fulfills this order.
	LocationId NullableInt32 `json:"location_id,omitempty"`
	// User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used
	OrderNumber *string `json:"order_number,omitempty"`
	// Products included in the order. Products identified by reference_id must also include the product name if there is no matching ShipBob product.
	Products []AddProductToOrder `json:"products"`
	// Date this order was purchase by the end user
	PurchaseDate NullableTime  `json:"purchase_date,omitempty"`
	Recipient    RecipientInfo `json:"recipient"`
	// Unique and immutable order identifier from your upstream system
	ReferenceId         string               `json:"reference_id"`
	RetailerProgramData *RetailerProgramData `json:"retailer_program_data,omitempty"`
	// Client-defined shipping method matching what the user has listed as the shipping method on the Ship Option Mapping setup page in the ShipBob Merchant Portal. If they don't match, we will create a new one and default it to Standard
	ShippingMethod string         `json:"shipping_method"`
	ShippingTerms  *ShippingTerms `json:"shipping_terms,omitempty"`
	// Key value pair array to store extra information at the order level for API purposes. ShipBob won't display the info in the ShipBob Merchant Portal or react based on this data.
	Tags []Tag `json:"tags,omitempty"`
	// Defaults to Direct to Consumer (DTC) if not provided. Note: B2B is not supported at this time
	Type *string `json:"type,omitempty"`
}

// NewCreateOrder instantiates a new CreateOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrder(products []AddProductToOrder, recipient RecipientInfo, referenceId string, shippingMethod string) *CreateOrder {
	this := CreateOrder{}
	this.Products = products
	this.Recipient = recipient
	this.ReferenceId = referenceId
	this.ShippingMethod = shippingMethod
	return &this
}

// NewCreateOrderWithDefaults instantiates a new CreateOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderWithDefaults() *CreateOrder {
	this := CreateOrder{}
	return &this
}

// GetFinancials returns the Financials field value if set, zero value otherwise.
func (o *CreateOrder) GetFinancials() Financials {
	if o == nil || IsNil(o.Financials) {
		var ret Financials
		return ret
	}
	return *o.Financials
}

// GetFinancialsOk returns a tuple with the Financials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetFinancialsOk() (*Financials, bool) {
	if o == nil || IsNil(o.Financials) {
		return nil, false
	}
	return o.Financials, true
}

// HasFinancials returns a boolean if a field has been set.
func (o *CreateOrder) HasFinancials() bool {
	if o != nil && !IsNil(o.Financials) {
		return true
	}

	return false
}

// SetFinancials gets a reference to the given Financials and assigns it to the Financials field.
func (o *CreateOrder) SetFinancials(v Financials) {
	o.Financials = &v
}

// GetGiftMessage returns the GiftMessage field value if set, zero value otherwise.
func (o *CreateOrder) GetGiftMessage() string {
	if o == nil || IsNil(o.GiftMessage) {
		var ret string
		return ret
	}
	return *o.GiftMessage
}

// GetGiftMessageOk returns a tuple with the GiftMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetGiftMessageOk() (*string, bool) {
	if o == nil || IsNil(o.GiftMessage) {
		return nil, false
	}
	return o.GiftMessage, true
}

// HasGiftMessage returns a boolean if a field has been set.
func (o *CreateOrder) HasGiftMessage() bool {
	if o != nil && !IsNil(o.GiftMessage) {
		return true
	}

	return false
}

// SetGiftMessage gets a reference to the given string and assigns it to the GiftMessage field.
func (o *CreateOrder) SetGiftMessage(v string) {
	o.GiftMessage = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrder) GetLocationId() int32 {
	if o == nil || IsNil(o.LocationId.Get()) {
		var ret int32
		return ret
	}
	return *o.LocationId.Get()
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrder) GetLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationId.Get(), o.LocationId.IsSet()
}

// HasLocationId returns a boolean if a field has been set.
func (o *CreateOrder) HasLocationId() bool {
	if o != nil && o.LocationId.IsSet() {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given NullableInt32 and assigns it to the LocationId field.
func (o *CreateOrder) SetLocationId(v int32) {
	o.LocationId.Set(&v)
}

// SetLocationIdNil sets the value for LocationId to be an explicit nil
func (o *CreateOrder) SetLocationIdNil() {
	o.LocationId.Set(nil)
}

// UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
func (o *CreateOrder) UnsetLocationId() {
	o.LocationId.Unset()
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *CreateOrder) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *CreateOrder) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *CreateOrder) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetProducts returns the Products field value
func (o *CreateOrder) GetProducts() []AddProductToOrder {
	if o == nil {
		var ret []AddProductToOrder
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetProductsOk() ([]AddProductToOrder, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *CreateOrder) SetProducts(v []AddProductToOrder) {
	o.Products = v
}

// GetPurchaseDate returns the PurchaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrder) GetPurchaseDate() time.Time {
	if o == nil || IsNil(o.PurchaseDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PurchaseDate.Get()
}

// GetPurchaseDateOk returns a tuple with the PurchaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrder) GetPurchaseDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchaseDate.Get(), o.PurchaseDate.IsSet()
}

// HasPurchaseDate returns a boolean if a field has been set.
func (o *CreateOrder) HasPurchaseDate() bool {
	if o != nil && o.PurchaseDate.IsSet() {
		return true
	}

	return false
}

// SetPurchaseDate gets a reference to the given NullableTime and assigns it to the PurchaseDate field.
func (o *CreateOrder) SetPurchaseDate(v time.Time) {
	o.PurchaseDate.Set(&v)
}

// SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil
func (o *CreateOrder) SetPurchaseDateNil() {
	o.PurchaseDate.Set(nil)
}

// UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
func (o *CreateOrder) UnsetPurchaseDate() {
	o.PurchaseDate.Unset()
}

// GetRecipient returns the Recipient field value
func (o *CreateOrder) GetRecipient() RecipientInfo {
	if o == nil {
		var ret RecipientInfo
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetRecipientOk() (*RecipientInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *CreateOrder) SetRecipient(v RecipientInfo) {
	o.Recipient = v
}

// GetReferenceId returns the ReferenceId field value
func (o *CreateOrder) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CreateOrder) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetRetailerProgramData returns the RetailerProgramData field value if set, zero value otherwise.
func (o *CreateOrder) GetRetailerProgramData() RetailerProgramData {
	if o == nil || IsNil(o.RetailerProgramData) {
		var ret RetailerProgramData
		return ret
	}
	return *o.RetailerProgramData
}

// GetRetailerProgramDataOk returns a tuple with the RetailerProgramData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetRetailerProgramDataOk() (*RetailerProgramData, bool) {
	if o == nil || IsNil(o.RetailerProgramData) {
		return nil, false
	}
	return o.RetailerProgramData, true
}

// HasRetailerProgramData returns a boolean if a field has been set.
func (o *CreateOrder) HasRetailerProgramData() bool {
	if o != nil && !IsNil(o.RetailerProgramData) {
		return true
	}

	return false
}

// SetRetailerProgramData gets a reference to the given RetailerProgramData and assigns it to the RetailerProgramData field.
func (o *CreateOrder) SetRetailerProgramData(v RetailerProgramData) {
	o.RetailerProgramData = &v
}

// GetShippingMethod returns the ShippingMethod field value
func (o *CreateOrder) GetShippingMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetShippingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethod, true
}

// SetShippingMethod sets field value
func (o *CreateOrder) SetShippingMethod(v string) {
	o.ShippingMethod = v
}

// GetShippingTerms returns the ShippingTerms field value if set, zero value otherwise.
func (o *CreateOrder) GetShippingTerms() ShippingTerms {
	if o == nil || IsNil(o.ShippingTerms) {
		var ret ShippingTerms
		return ret
	}
	return *o.ShippingTerms
}

// GetShippingTermsOk returns a tuple with the ShippingTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetShippingTermsOk() (*ShippingTerms, bool) {
	if o == nil || IsNil(o.ShippingTerms) {
		return nil, false
	}
	return o.ShippingTerms, true
}

// HasShippingTerms returns a boolean if a field has been set.
func (o *CreateOrder) HasShippingTerms() bool {
	if o != nil && !IsNil(o.ShippingTerms) {
		return true
	}

	return false
}

// SetShippingTerms gets a reference to the given ShippingTerms and assigns it to the ShippingTerms field.
func (o *CreateOrder) SetShippingTerms(v ShippingTerms) {
	o.ShippingTerms = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateOrder) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOrder) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *CreateOrder) SetTags(v []Tag) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateOrder) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrder) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateOrder) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateOrder) SetType(v string) {
	o.Type = &v
}

func (o CreateOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Financials) {
		toSerialize["financials"] = o.Financials
	}
	if !IsNil(o.GiftMessage) {
		toSerialize["gift_message"] = o.GiftMessage
	}
	if o.LocationId.IsSet() {
		toSerialize["location_id"] = o.LocationId.Get()
	}
	if !IsNil(o.OrderNumber) {
		toSerialize["order_number"] = o.OrderNumber
	}
	toSerialize["products"] = o.Products
	if o.PurchaseDate.IsSet() {
		toSerialize["purchase_date"] = o.PurchaseDate.Get()
	}
	toSerialize["recipient"] = o.Recipient
	toSerialize["reference_id"] = o.ReferenceId
	if !IsNil(o.RetailerProgramData) {
		toSerialize["retailer_program_data"] = o.RetailerProgramData
	}
	toSerialize["shipping_method"] = o.ShippingMethod
	if !IsNil(o.ShippingTerms) {
		toSerialize["shipping_terms"] = o.ShippingTerms
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateOrder struct {
	value *CreateOrder
	isSet bool
}

func (v NullableCreateOrder) Get() *CreateOrder {
	return v.value
}

func (v *NullableCreateOrder) Set(val *CreateOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrder(val *CreateOrder) *NullableCreateOrder {
	return &NullableCreateOrder{value: val, isSet: true}
}

func (v NullableCreateOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
