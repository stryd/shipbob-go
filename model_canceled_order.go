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

// checks if the CanceledOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CanceledOrder{}

// CanceledOrder 
type CanceledOrder struct {
	// Results of canceling the shipments associated with the order
	CanceledShipmentResults []CanceledShipment `json:"canceled_shipment_results,omitempty"`
	Order *Order `json:"order,omitempty"`
	// The ID of the canceled order
	OrderId *int32 `json:"order_id,omitempty"`
	// The overall result of canceling the shipments associated with the order
	Status *string `json:"status,omitempty"`
}

// NewCanceledOrder instantiates a new CanceledOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanceledOrder() *CanceledOrder {
	this := CanceledOrder{}
	return &this
}

// NewCanceledOrderWithDefaults instantiates a new CanceledOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanceledOrderWithDefaults() *CanceledOrder {
	this := CanceledOrder{}
	return &this
}

// GetCanceledShipmentResults returns the CanceledShipmentResults field value if set, zero value otherwise.
func (o *CanceledOrder) GetCanceledShipmentResults() []CanceledShipment {
	if o == nil || IsNil(o.CanceledShipmentResults) {
		var ret []CanceledShipment
		return ret
	}
	return o.CanceledShipmentResults
}

// GetCanceledShipmentResultsOk returns a tuple with the CanceledShipmentResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanceledOrder) GetCanceledShipmentResultsOk() ([]CanceledShipment, bool) {
	if o == nil || IsNil(o.CanceledShipmentResults) {
		return nil, false
	}
	return o.CanceledShipmentResults, true
}

// HasCanceledShipmentResults returns a boolean if a field has been set.
func (o *CanceledOrder) HasCanceledShipmentResults() bool {
	if o != nil && !IsNil(o.CanceledShipmentResults) {
		return true
	}

	return false
}

// SetCanceledShipmentResults gets a reference to the given []CanceledShipment and assigns it to the CanceledShipmentResults field.
func (o *CanceledOrder) SetCanceledShipmentResults(v []CanceledShipment) {
	o.CanceledShipmentResults = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CanceledOrder) GetOrder() Order {
	if o == nil || IsNil(o.Order) {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanceledOrder) GetOrderOk() (*Order, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CanceledOrder) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *CanceledOrder) SetOrder(v Order) {
	o.Order = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CanceledOrder) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanceledOrder) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CanceledOrder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *CanceledOrder) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CanceledOrder) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanceledOrder) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CanceledOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CanceledOrder) SetStatus(v string) {
	o.Status = &v
}

func (o CanceledOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CanceledOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanceledShipmentResults) {
		toSerialize["canceled_shipment_results"] = o.CanceledShipmentResults
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCanceledOrder struct {
	value *CanceledOrder
	isSet bool
}

func (v NullableCanceledOrder) Get() *CanceledOrder {
	return v.value
}

func (v *NullableCanceledOrder) Set(val *CanceledOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCanceledOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCanceledOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanceledOrder(val *CanceledOrder) *NullableCanceledOrder {
	return &NullableCanceledOrder{value: val, isSet: true}
}

func (v NullableCanceledOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanceledOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


