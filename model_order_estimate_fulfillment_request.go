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

// checks if the OrderEstimateFulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderEstimateFulfillmentRequest{}

// OrderEstimateFulfillmentRequest struct for OrderEstimateFulfillmentRequest
type OrderEstimateFulfillmentRequest struct {
	Address EstimationAddress `json:"address"`
	// Products to be included in the order. Each product must include one of reference_id or id
	Products []OrderEstimateProductInfo `json:"products"`
	// Array of strings specifying shipping methods for which to fetch estimates.  If this field is omitted we will return estimates for all shipping methods defined in ShipBob
	ShippingMethods []string `json:"shipping_methods,omitempty"`
}

// NewOrderEstimateFulfillmentRequest instantiates a new OrderEstimateFulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderEstimateFulfillmentRequest(address EstimationAddress, products []OrderEstimateProductInfo) *OrderEstimateFulfillmentRequest {
	this := OrderEstimateFulfillmentRequest{}
	this.Address = address
	this.Products = products
	return &this
}

// NewOrderEstimateFulfillmentRequestWithDefaults instantiates a new OrderEstimateFulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderEstimateFulfillmentRequestWithDefaults() *OrderEstimateFulfillmentRequest {
	this := OrderEstimateFulfillmentRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *OrderEstimateFulfillmentRequest) GetAddress() EstimationAddress {
	if o == nil {
		var ret EstimationAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *OrderEstimateFulfillmentRequest) GetAddressOk() (*EstimationAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *OrderEstimateFulfillmentRequest) SetAddress(v EstimationAddress) {
	o.Address = v
}

// GetProducts returns the Products field value
func (o *OrderEstimateFulfillmentRequest) GetProducts() []OrderEstimateProductInfo {
	if o == nil {
		var ret []OrderEstimateProductInfo
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *OrderEstimateFulfillmentRequest) GetProductsOk() ([]OrderEstimateProductInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *OrderEstimateFulfillmentRequest) SetProducts(v []OrderEstimateProductInfo) {
	o.Products = v
}

// GetShippingMethods returns the ShippingMethods field value if set, zero value otherwise.
func (o *OrderEstimateFulfillmentRequest) GetShippingMethods() []string {
	if o == nil || IsNil(o.ShippingMethods) {
		var ret []string
		return ret
	}
	return o.ShippingMethods
}

// GetShippingMethodsOk returns a tuple with the ShippingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderEstimateFulfillmentRequest) GetShippingMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShippingMethods) {
		return nil, false
	}
	return o.ShippingMethods, true
}

// HasShippingMethods returns a boolean if a field has been set.
func (o *OrderEstimateFulfillmentRequest) HasShippingMethods() bool {
	if o != nil && !IsNil(o.ShippingMethods) {
		return true
	}

	return false
}

// SetShippingMethods gets a reference to the given []string and assigns it to the ShippingMethods field.
func (o *OrderEstimateFulfillmentRequest) SetShippingMethods(v []string) {
	o.ShippingMethods = v
}

func (o OrderEstimateFulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderEstimateFulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["products"] = o.Products
	if !IsNil(o.ShippingMethods) {
		toSerialize["shipping_methods"] = o.ShippingMethods
	}
	return toSerialize, nil
}

type NullableOrderEstimateFulfillmentRequest struct {
	value *OrderEstimateFulfillmentRequest
	isSet bool
}

func (v NullableOrderEstimateFulfillmentRequest) Get() *OrderEstimateFulfillmentRequest {
	return v.value
}

func (v *NullableOrderEstimateFulfillmentRequest) Set(val *OrderEstimateFulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderEstimateFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderEstimateFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderEstimateFulfillmentRequest(val *OrderEstimateFulfillmentRequest) *NullableOrderEstimateFulfillmentRequest {
	return &NullableOrderEstimateFulfillmentRequest{value: val, isSet: true}
}

func (v NullableOrderEstimateFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderEstimateFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
