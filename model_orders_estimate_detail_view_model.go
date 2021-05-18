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

// OrdersEstimateDetailViewModel struct for OrdersEstimateDetailViewModel
type OrdersEstimateDetailViewModel struct {
	// Estimated price in dollars for the provided shipping method
	EstimatedPrice *float64 `json:"estimated_price,omitempty"`
	FulfillmentCenter *OrdersFulfillmentCenterViewModel `json:"fulfillment_center,omitempty"`
	// Provided shipping method. Maps to ship option in ShipBob.
	ShippingMethod *string `json:"shipping_method,omitempty"`
}

// NewOrdersEstimateDetailViewModel instantiates a new OrdersEstimateDetailViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersEstimateDetailViewModel() *OrdersEstimateDetailViewModel {
	this := OrdersEstimateDetailViewModel{}
	return &this
}

// NewOrdersEstimateDetailViewModelWithDefaults instantiates a new OrdersEstimateDetailViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersEstimateDetailViewModelWithDefaults() *OrdersEstimateDetailViewModel {
	this := OrdersEstimateDetailViewModel{}
	return &this
}

// GetEstimatedPrice returns the EstimatedPrice field value if set, zero value otherwise.
func (o *OrdersEstimateDetailViewModel) GetEstimatedPrice() float64 {
	if o == nil || o.EstimatedPrice == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedPrice
}

// GetEstimatedPriceOk returns a tuple with the EstimatedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersEstimateDetailViewModel) GetEstimatedPriceOk() (*float64, bool) {
	if o == nil || o.EstimatedPrice == nil {
		return nil, false
	}
	return o.EstimatedPrice, true
}

// HasEstimatedPrice returns a boolean if a field has been set.
func (o *OrdersEstimateDetailViewModel) HasEstimatedPrice() bool {
	if o != nil && o.EstimatedPrice != nil {
		return true
	}

	return false
}

// SetEstimatedPrice gets a reference to the given float64 and assigns it to the EstimatedPrice field.
func (o *OrdersEstimateDetailViewModel) SetEstimatedPrice(v float64) {
	o.EstimatedPrice = &v
}

// GetFulfillmentCenter returns the FulfillmentCenter field value if set, zero value otherwise.
func (o *OrdersEstimateDetailViewModel) GetFulfillmentCenter() OrdersFulfillmentCenterViewModel {
	if o == nil || o.FulfillmentCenter == nil {
		var ret OrdersFulfillmentCenterViewModel
		return ret
	}
	return *o.FulfillmentCenter
}

// GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersEstimateDetailViewModel) GetFulfillmentCenterOk() (*OrdersFulfillmentCenterViewModel, bool) {
	if o == nil || o.FulfillmentCenter == nil {
		return nil, false
	}
	return o.FulfillmentCenter, true
}

// HasFulfillmentCenter returns a boolean if a field has been set.
func (o *OrdersEstimateDetailViewModel) HasFulfillmentCenter() bool {
	if o != nil && o.FulfillmentCenter != nil {
		return true
	}

	return false
}

// SetFulfillmentCenter gets a reference to the given OrdersFulfillmentCenterViewModel and assigns it to the FulfillmentCenter field.
func (o *OrdersEstimateDetailViewModel) SetFulfillmentCenter(v OrdersFulfillmentCenterViewModel) {
	o.FulfillmentCenter = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *OrdersEstimateDetailViewModel) GetShippingMethod() string {
	if o == nil || o.ShippingMethod == nil {
		var ret string
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersEstimateDetailViewModel) GetShippingMethodOk() (*string, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *OrdersEstimateDetailViewModel) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given string and assigns it to the ShippingMethod field.
func (o *OrdersEstimateDetailViewModel) SetShippingMethod(v string) {
	o.ShippingMethod = &v
}

func (o OrdersEstimateDetailViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EstimatedPrice != nil {
		toSerialize["estimated_price"] = o.EstimatedPrice
	}
	if o.FulfillmentCenter != nil {
		toSerialize["fulfillment_center"] = o.FulfillmentCenter
	}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullableOrdersEstimateDetailViewModel struct {
	value *OrdersEstimateDetailViewModel
	isSet bool
}

func (v NullableOrdersEstimateDetailViewModel) Get() *OrdersEstimateDetailViewModel {
	return v.value
}

func (v *NullableOrdersEstimateDetailViewModel) Set(val *OrdersEstimateDetailViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersEstimateDetailViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersEstimateDetailViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersEstimateDetailViewModel(val *OrdersEstimateDetailViewModel) *NullableOrdersEstimateDetailViewModel {
	return &NullableOrdersEstimateDetailViewModel{value: val, isSet: true}
}

func (v NullableOrdersEstimateDetailViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersEstimateDetailViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


