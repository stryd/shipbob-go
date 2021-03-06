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

// CreateReceivingOrderBoxes Information about a box shipment to be added to a receiving order
type CreateReceivingOrderBoxes struct {
	// Items contained in this box
	BoxItems []CreateReceivingOrderBoxItems `json:"box_items"`
	// Tracking number for the box shipment
	TrackingNumber NullableString `json:"tracking_number"`
}

// NewCreateReceivingOrderBoxes instantiates a new CreateReceivingOrderBoxes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceivingOrderBoxes(boxItems []CreateReceivingOrderBoxItems, trackingNumber NullableString) *CreateReceivingOrderBoxes {
	this := CreateReceivingOrderBoxes{}
	this.BoxItems = boxItems
	this.TrackingNumber = trackingNumber
	return &this
}

// NewCreateReceivingOrderBoxesWithDefaults instantiates a new CreateReceivingOrderBoxes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceivingOrderBoxesWithDefaults() *CreateReceivingOrderBoxes {
	this := CreateReceivingOrderBoxes{}
	return &this
}

// GetBoxItems returns the BoxItems field value
// If the value is explicit nil, the zero value for []CreateReceivingOrderBoxItems will be returned
func (o *CreateReceivingOrderBoxes) GetBoxItems() []CreateReceivingOrderBoxItems {
	if o == nil {
		var ret []CreateReceivingOrderBoxItems
		return ret
	}

	return o.BoxItems
}

// GetBoxItemsOk returns a tuple with the BoxItems field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrderBoxes) GetBoxItemsOk() (*[]CreateReceivingOrderBoxItems, bool) {
	if o == nil || o.BoxItems == nil {
		return nil, false
	}
	return &o.BoxItems, true
}

// SetBoxItems sets field value
func (o *CreateReceivingOrderBoxes) SetBoxItems(v []CreateReceivingOrderBoxItems) {
	o.BoxItems = v
}

// GetTrackingNumber returns the TrackingNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateReceivingOrderBoxes) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrderBoxes) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// SetTrackingNumber sets field value
func (o *CreateReceivingOrderBoxes) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}

func (o CreateReceivingOrderBoxes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoxItems != nil {
		toSerialize["box_items"] = o.BoxItems
	}
	if true {
		toSerialize["tracking_number"] = o.TrackingNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReceivingOrderBoxes struct {
	value *CreateReceivingOrderBoxes
	isSet bool
}

func (v NullableCreateReceivingOrderBoxes) Get() *CreateReceivingOrderBoxes {
	return v.value
}

func (v *NullableCreateReceivingOrderBoxes) Set(val *CreateReceivingOrderBoxes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceivingOrderBoxes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceivingOrderBoxes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceivingOrderBoxes(val *CreateReceivingOrderBoxes) *NullableCreateReceivingOrderBoxes {
	return &NullableCreateReceivingOrderBoxes{value: val, isSet: true}
}

func (v NullableCreateReceivingOrderBoxes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceivingOrderBoxes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
