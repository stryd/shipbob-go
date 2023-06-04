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

// checks if the ShippingTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingTerms{}

// ShippingTerms Contains shipping properties that need to be used for fulfilling an order.
type ShippingTerms struct {
	// Identifies whether to ship parcel or freight.              Parcel: Smaller, light weight boxes.              Freight: Larger boxes, usually transported by truckload.
	CarrierType NullableString `json:"carrier_type,omitempty"`
	// Identifies the party responsible for shipping charges.              Collect: The person/entity receiving the product pays the shipping charges.              ThirdParty: Another party pays for the shipping charges (not Shipbob) [parcel only].              Prepaid: The shipper pays the shipping charges (Shipbob or merchant).              MerchantResponsible: The merchant will be responsible for uploading shipping labels or booking freight transportation.
	PaymentTerm NullableString `json:"payment_term,omitempty"`
}

// NewShippingTerms instantiates a new ShippingTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingTerms() *ShippingTerms {
	this := ShippingTerms{}
	return &this
}

// NewShippingTermsWithDefaults instantiates a new ShippingTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingTermsWithDefaults() *ShippingTerms {
	this := ShippingTerms{}
	return &this
}

// GetCarrierType returns the CarrierType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingTerms) GetCarrierType() string {
	if o == nil || IsNil(o.CarrierType.Get()) {
		var ret string
		return ret
	}
	return *o.CarrierType.Get()
}

// GetCarrierTypeOk returns a tuple with the CarrierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingTerms) GetCarrierTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CarrierType.Get(), o.CarrierType.IsSet()
}

// HasCarrierType returns a boolean if a field has been set.
func (o *ShippingTerms) HasCarrierType() bool {
	if o != nil && o.CarrierType.IsSet() {
		return true
	}

	return false
}

// SetCarrierType gets a reference to the given NullableString and assigns it to the CarrierType field.
func (o *ShippingTerms) SetCarrierType(v string) {
	o.CarrierType.Set(&v)
}
// SetCarrierTypeNil sets the value for CarrierType to be an explicit nil
func (o *ShippingTerms) SetCarrierTypeNil() {
	o.CarrierType.Set(nil)
}

// UnsetCarrierType ensures that no value is present for CarrierType, not even an explicit nil
func (o *ShippingTerms) UnsetCarrierType() {
	o.CarrierType.Unset()
}

// GetPaymentTerm returns the PaymentTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingTerms) GetPaymentTerm() string {
	if o == nil || IsNil(o.PaymentTerm.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentTerm.Get()
}

// GetPaymentTermOk returns a tuple with the PaymentTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingTerms) GetPaymentTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentTerm.Get(), o.PaymentTerm.IsSet()
}

// HasPaymentTerm returns a boolean if a field has been set.
func (o *ShippingTerms) HasPaymentTerm() bool {
	if o != nil && o.PaymentTerm.IsSet() {
		return true
	}

	return false
}

// SetPaymentTerm gets a reference to the given NullableString and assigns it to the PaymentTerm field.
func (o *ShippingTerms) SetPaymentTerm(v string) {
	o.PaymentTerm.Set(&v)
}
// SetPaymentTermNil sets the value for PaymentTerm to be an explicit nil
func (o *ShippingTerms) SetPaymentTermNil() {
	o.PaymentTerm.Set(nil)
}

// UnsetPaymentTerm ensures that no value is present for PaymentTerm, not even an explicit nil
func (o *ShippingTerms) UnsetPaymentTerm() {
	o.PaymentTerm.Unset()
}

func (o ShippingTerms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CarrierType.IsSet() {
		toSerialize["carrier_type"] = o.CarrierType.Get()
	}
	if o.PaymentTerm.IsSet() {
		toSerialize["payment_term"] = o.PaymentTerm.Get()
	}
	return toSerialize, nil
}

type NullableShippingTerms struct {
	value *ShippingTerms
	isSet bool
}

func (v NullableShippingTerms) Get() *ShippingTerms {
	return v.value
}

func (v *NullableShippingTerms) Set(val *ShippingTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingTerms(val *ShippingTerms) *NullableShippingTerms {
	return &NullableShippingTerms{value: val, isSet: true}
}

func (v NullableShippingTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

