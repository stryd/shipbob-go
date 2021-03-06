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
	"time"
)

// CreateReceivingOrder Information to create a new receiving order
type CreateReceivingOrder struct {
	BoxPackagingType PackingType `json:"box_packaging_type"`
	// Box shipments to be added to this receiving order
	Boxes []CreateReceivingOrderBoxes `json:"boxes"`
	// Expected arrival date of all the box shipments in this receiving order
	ExpectedArrivalDate NullableTime                          `json:"expected_arrival_date"`
	FulfillmentCenter   CreateReceivingOrderFulfillmentCenter `json:"fulfillment_center"`
	PackageType         PackageType                           `json:"package_type"`
}

// NewCreateReceivingOrder instantiates a new CreateReceivingOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceivingOrder(boxPackagingType PackingType, boxes []CreateReceivingOrderBoxes, expectedArrivalDate NullableTime, fulfillmentCenter CreateReceivingOrderFulfillmentCenter, packageType PackageType) *CreateReceivingOrder {
	this := CreateReceivingOrder{}
	this.BoxPackagingType = boxPackagingType
	this.Boxes = boxes
	this.ExpectedArrivalDate = expectedArrivalDate
	this.FulfillmentCenter = fulfillmentCenter
	this.PackageType = packageType
	return &this
}

// NewCreateReceivingOrderWithDefaults instantiates a new CreateReceivingOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceivingOrderWithDefaults() *CreateReceivingOrder {
	this := CreateReceivingOrder{}
	return &this
}

// GetBoxPackagingType returns the BoxPackagingType field value
func (o *CreateReceivingOrder) GetBoxPackagingType() PackingType {
	if o == nil {
		var ret PackingType
		return ret
	}

	return o.BoxPackagingType
}

// GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetBoxPackagingTypeOk() (*PackingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoxPackagingType, true
}

// SetBoxPackagingType sets field value
func (o *CreateReceivingOrder) SetBoxPackagingType(v PackingType) {
	o.BoxPackagingType = v
}

// GetBoxes returns the Boxes field value
// If the value is explicit nil, the zero value for []CreateReceivingOrderBoxes will be returned
func (o *CreateReceivingOrder) GetBoxes() []CreateReceivingOrderBoxes {
	if o == nil {
		var ret []CreateReceivingOrderBoxes
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrder) GetBoxesOk() (*[]CreateReceivingOrderBoxes, bool) {
	if o == nil || o.Boxes == nil {
		return nil, false
	}
	return &o.Boxes, true
}

// SetBoxes sets field value
func (o *CreateReceivingOrder) SetBoxes(v []CreateReceivingOrderBoxes) {
	o.Boxes = v
}

// GetExpectedArrivalDate returns the ExpectedArrivalDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CreateReceivingOrder) GetExpectedArrivalDate() time.Time {
	if o == nil || o.ExpectedArrivalDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpectedArrivalDate.Get()
}

// GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrder) GetExpectedArrivalDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedArrivalDate.Get(), o.ExpectedArrivalDate.IsSet()
}

// SetExpectedArrivalDate sets field value
func (o *CreateReceivingOrder) SetExpectedArrivalDate(v time.Time) {
	o.ExpectedArrivalDate.Set(&v)
}

// GetFulfillmentCenter returns the FulfillmentCenter field value
func (o *CreateReceivingOrder) GetFulfillmentCenter() CreateReceivingOrderFulfillmentCenter {
	if o == nil {
		var ret CreateReceivingOrderFulfillmentCenter
		return ret
	}

	return o.FulfillmentCenter
}

// GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetFulfillmentCenterOk() (*CreateReceivingOrderFulfillmentCenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentCenter, true
}

// SetFulfillmentCenter sets field value
func (o *CreateReceivingOrder) SetFulfillmentCenter(v CreateReceivingOrderFulfillmentCenter) {
	o.FulfillmentCenter = v
}

// GetPackageType returns the PackageType field value
func (o *CreateReceivingOrder) GetPackageType() PackageType {
	if o == nil {
		var ret PackageType
		return ret
	}

	return o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetPackageTypeOk() (*PackageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageType, true
}

// SetPackageType sets field value
func (o *CreateReceivingOrder) SetPackageType(v PackageType) {
	o.PackageType = v
}

func (o CreateReceivingOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["box_packaging_type"] = o.BoxPackagingType
	}
	if o.Boxes != nil {
		toSerialize["boxes"] = o.Boxes
	}
	if true {
		toSerialize["expected_arrival_date"] = o.ExpectedArrivalDate.Get()
	}
	if true {
		toSerialize["fulfillment_center"] = o.FulfillmentCenter
	}
	if true {
		toSerialize["package_type"] = o.PackageType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReceivingOrder struct {
	value *CreateReceivingOrder
	isSet bool
}

func (v NullableCreateReceivingOrder) Get() *CreateReceivingOrder {
	return v.value
}

func (v *NullableCreateReceivingOrder) Set(val *CreateReceivingOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceivingOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceivingOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceivingOrder(val *CreateReceivingOrder) *NullableCreateReceivingOrder {
	return &NullableCreateReceivingOrder{value: val, isSet: true}
}

func (v NullableCreateReceivingOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceivingOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
