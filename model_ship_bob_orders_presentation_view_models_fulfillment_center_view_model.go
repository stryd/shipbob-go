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

// ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel Information about a fulfillment center that a shipment can belong to
type ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel struct {
	// Id of the fulfillment center
	Id *int32 `json:"id,omitempty"`
	// Name of the fulfillment center
	Name *string `json:"name,omitempty"`
}

// NewShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel instantiates a new ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel() *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel {
	this := ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel{}
	return &this
}

// NewShipBobOrdersPresentationViewModelsFulfillmentCenterViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipBobOrdersPresentationViewModelsFulfillmentCenterViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel {
	this := ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) SetName(v string) {
	o.Name = &v
}

func (o ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel struct {
	value *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel
	isSet bool
}

func (v NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) Get() *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel {
	return v.value
}

func (v *NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) Set(val *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel(val *ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) *NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel {
	return &NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel{value: val, isSet: true}
}

func (v NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


