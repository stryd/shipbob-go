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

// checks if the CreateReceivingOrderFulfillmentCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReceivingOrderFulfillmentCenter{}

// CreateReceivingOrderFulfillmentCenter Model containing information that assigns a receiving order to a fulfillment center
type CreateReceivingOrderFulfillmentCenter struct {
	// ID of the fulfillment center to assign this receiving order to
	Id int32 `json:"id"`
}

// NewCreateReceivingOrderFulfillmentCenter instantiates a new CreateReceivingOrderFulfillmentCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceivingOrderFulfillmentCenter(id int32) *CreateReceivingOrderFulfillmentCenter {
	this := CreateReceivingOrderFulfillmentCenter{}
	this.Id = id
	return &this
}

// NewCreateReceivingOrderFulfillmentCenterWithDefaults instantiates a new CreateReceivingOrderFulfillmentCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceivingOrderFulfillmentCenterWithDefaults() *CreateReceivingOrderFulfillmentCenter {
	this := CreateReceivingOrderFulfillmentCenter{}
	return &this
}

// GetId returns the Id field value
func (o *CreateReceivingOrderFulfillmentCenter) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrderFulfillmentCenter) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateReceivingOrderFulfillmentCenter) SetId(v int32) {
	o.Id = v
}

func (o CreateReceivingOrderFulfillmentCenter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReceivingOrderFulfillmentCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateReceivingOrderFulfillmentCenter struct {
	value *CreateReceivingOrderFulfillmentCenter
	isSet bool
}

func (v NullableCreateReceivingOrderFulfillmentCenter) Get() *CreateReceivingOrderFulfillmentCenter {
	return v.value
}

func (v *NullableCreateReceivingOrderFulfillmentCenter) Set(val *CreateReceivingOrderFulfillmentCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceivingOrderFulfillmentCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceivingOrderFulfillmentCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceivingOrderFulfillmentCenter(val *CreateReceivingOrderFulfillmentCenter) *NullableCreateReceivingOrderFulfillmentCenter {
	return &NullableCreateReceivingOrderFulfillmentCenter{value: val, isSet: true}
}

func (v NullableCreateReceivingOrderFulfillmentCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceivingOrderFulfillmentCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
