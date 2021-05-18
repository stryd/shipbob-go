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

// ReceivingAddBoxToOrderModel Information about a box shipment to be added to a receiving order
type ReceivingAddBoxToOrderModel struct {
	// Items contained in this box
	BoxItems []ReceivingAddBoxItemToBoxModel `json:"box_items"`
	// Tracking number for the box shipment
	TrackingNumber NullableString `json:"tracking_number"`
}

// NewReceivingAddBoxToOrderModel instantiates a new ReceivingAddBoxToOrderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivingAddBoxToOrderModel(boxItems []ReceivingAddBoxItemToBoxModel, trackingNumber NullableString, ) *ReceivingAddBoxToOrderModel {
	this := ReceivingAddBoxToOrderModel{}
	this.BoxItems = boxItems
	this.TrackingNumber = trackingNumber
	return &this
}

// NewReceivingAddBoxToOrderModelWithDefaults instantiates a new ReceivingAddBoxToOrderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivingAddBoxToOrderModelWithDefaults() *ReceivingAddBoxToOrderModel {
	this := ReceivingAddBoxToOrderModel{}
	return &this
}

// GetBoxItems returns the BoxItems field value
// If the value is explicit nil, the zero value for []ReceivingAddBoxItemToBoxModel will be returned
func (o *ReceivingAddBoxToOrderModel) GetBoxItems() []ReceivingAddBoxItemToBoxModel {
	if o == nil  {
		var ret []ReceivingAddBoxItemToBoxModel
		return ret
	}

	return o.BoxItems
}

// GetBoxItemsOk returns a tuple with the BoxItems field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingAddBoxToOrderModel) GetBoxItemsOk() (*[]ReceivingAddBoxItemToBoxModel, bool) {
	if o == nil || o.BoxItems == nil {
		return nil, false
	}
	return &o.BoxItems, true
}

// SetBoxItems sets field value
func (o *ReceivingAddBoxToOrderModel) SetBoxItems(v []ReceivingAddBoxItemToBoxModel) {
	o.BoxItems = v
}

// GetTrackingNumber returns the TrackingNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReceivingAddBoxToOrderModel) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingAddBoxToOrderModel) GetTrackingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// SetTrackingNumber sets field value
func (o *ReceivingAddBoxToOrderModel) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}

func (o ReceivingAddBoxToOrderModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoxItems != nil {
		toSerialize["box_items"] = o.BoxItems
	}
	if true {
		toSerialize["tracking_number"] = o.TrackingNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReceivingAddBoxToOrderModel struct {
	value *ReceivingAddBoxToOrderModel
	isSet bool
}

func (v NullableReceivingAddBoxToOrderModel) Get() *ReceivingAddBoxToOrderModel {
	return v.value
}

func (v *NullableReceivingAddBoxToOrderModel) Set(val *ReceivingAddBoxToOrderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingAddBoxToOrderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingAddBoxToOrderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingAddBoxToOrderModel(val *ReceivingAddBoxToOrderModel) *NullableReceivingAddBoxToOrderModel {
	return &NullableReceivingAddBoxToOrderModel{value: val, isSet: true}
}

func (v NullableReceivingAddBoxToOrderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingAddBoxToOrderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

