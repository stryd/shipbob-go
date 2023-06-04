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

// checks if the AddStoreOrderJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddStoreOrderJson{}

// AddStoreOrderJson Model for adding a Store Order Json to a ShipBob Order.
type AddStoreOrderJson struct {
	// Json String that represent the order on a store front system
	OrderJson string `json:"order_json"`
}

// NewAddStoreOrderJson instantiates a new AddStoreOrderJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddStoreOrderJson(orderJson string) *AddStoreOrderJson {
	this := AddStoreOrderJson{}
	this.OrderJson = orderJson
	return &this
}

// NewAddStoreOrderJsonWithDefaults instantiates a new AddStoreOrderJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddStoreOrderJsonWithDefaults() *AddStoreOrderJson {
	this := AddStoreOrderJson{}
	return &this
}

// GetOrderJson returns the OrderJson field value
func (o *AddStoreOrderJson) GetOrderJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderJson
}

// GetOrderJsonOk returns a tuple with the OrderJson field value
// and a boolean to check if the value has been set.
func (o *AddStoreOrderJson) GetOrderJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderJson, true
}

// SetOrderJson sets field value
func (o *AddStoreOrderJson) SetOrderJson(v string) {
	o.OrderJson = v
}

func (o AddStoreOrderJson) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddStoreOrderJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_json"] = o.OrderJson
	return toSerialize, nil
}

type NullableAddStoreOrderJson struct {
	value *AddStoreOrderJson
	isSet bool
}

func (v NullableAddStoreOrderJson) Get() *AddStoreOrderJson {
	return v.value
}

func (v *NullableAddStoreOrderJson) Set(val *AddStoreOrderJson) {
	v.value = val
	v.isSet = true
}

func (v NullableAddStoreOrderJson) IsSet() bool {
	return v.isSet
}

func (v *NullableAddStoreOrderJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddStoreOrderJson(val *AddStoreOrderJson) *NullableAddStoreOrderJson {
	return &NullableAddStoreOrderJson{value: val, isSet: true}
}

func (v NullableAddStoreOrderJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddStoreOrderJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
