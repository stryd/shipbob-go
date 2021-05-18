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

// ShipBobOrdersPresentationViewModelsCanceledOrderViewModel struct for ShipBobOrdersPresentationViewModelsCanceledOrderViewModel
type ShipBobOrdersPresentationViewModelsCanceledOrderViewModel struct {
	// Results of camceling the shipments associated with the order
	CanceledShipmentResults *[]ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel `json:"canceled_shipment_results,omitempty"`
	Order *ShipBobOrdersPresentationViewModelsOrderViewModel `json:"order,omitempty"`
	// The ID of the canceled order
	OrderId *int32 `json:"order_id,omitempty"`
	// The overall result of canceling the shipments associated with the order
	Status *string `json:"status,omitempty"`
}

// NewShipBobOrdersPresentationViewModelsCanceledOrderViewModel instantiates a new ShipBobOrdersPresentationViewModelsCanceledOrderViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipBobOrdersPresentationViewModelsCanceledOrderViewModel() *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel {
	this := ShipBobOrdersPresentationViewModelsCanceledOrderViewModel{}
	return &this
}

// NewShipBobOrdersPresentationViewModelsCanceledOrderViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsCanceledOrderViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipBobOrdersPresentationViewModelsCanceledOrderViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel {
	this := ShipBobOrdersPresentationViewModelsCanceledOrderViewModel{}
	return &this
}

// GetCanceledShipmentResults returns the CanceledShipmentResults field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetCanceledShipmentResults() []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel {
	if o == nil || o.CanceledShipmentResults == nil {
		var ret []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel
		return ret
	}
	return *o.CanceledShipmentResults
}

// GetCanceledShipmentResultsOk returns a tuple with the CanceledShipmentResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetCanceledShipmentResultsOk() (*[]ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel, bool) {
	if o == nil || o.CanceledShipmentResults == nil {
		return nil, false
	}
	return o.CanceledShipmentResults, true
}

// HasCanceledShipmentResults returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasCanceledShipmentResults() bool {
	if o != nil && o.CanceledShipmentResults != nil {
		return true
	}

	return false
}

// SetCanceledShipmentResults gets a reference to the given []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel and assigns it to the CanceledShipmentResults field.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetCanceledShipmentResults(v []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel) {
	o.CanceledShipmentResults = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrder() ShipBobOrdersPresentationViewModelsOrderViewModel {
	if o == nil || o.Order == nil {
		var ret ShipBobOrdersPresentationViewModelsOrderViewModel
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderOk() (*ShipBobOrdersPresentationViewModelsOrderViewModel, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given ShipBobOrdersPresentationViewModelsOrderViewModel and assigns it to the Order field.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetOrder(v ShipBobOrdersPresentationViewModelsOrderViewModel) {
	o.Order = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderId() int32 {
	if o == nil || o.OrderId == nil {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderIdOk() (*int32, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetStatus(v string) {
	o.Status = &v
}

func (o ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanceledShipmentResults != nil {
		toSerialize["canceled_shipment_results"] = o.CanceledShipmentResults
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel struct {
	value *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel
	isSet bool
}

func (v NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) Get() *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel {
	return v.value
}

func (v *NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) Set(val *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel(val *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) *NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel {
	return &NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel{value: val, isSet: true}
}

func (v NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipBobOrdersPresentationViewModelsCanceledOrderViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


