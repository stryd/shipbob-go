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

// checks if the CancelShipmentBulk200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelShipmentBulk200Response{}

// CancelShipmentBulk200Response
type CancelShipmentBulk200Response struct {
	// The results of all cancellation actions
	Results []CanceledShipment `json:"results,omitempty"`
}

// NewCancelShipmentBulk200Response instantiates a new CancelShipmentBulk200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelShipmentBulk200Response() *CancelShipmentBulk200Response {
	this := CancelShipmentBulk200Response{}
	return &this
}

// NewCancelShipmentBulk200ResponseWithDefaults instantiates a new CancelShipmentBulk200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelShipmentBulk200ResponseWithDefaults() *CancelShipmentBulk200Response {
	this := CancelShipmentBulk200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CancelShipmentBulk200Response) GetResults() []CanceledShipment {
	if o == nil || IsNil(o.Results) {
		var ret []CanceledShipment
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelShipmentBulk200Response) GetResultsOk() ([]CanceledShipment, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CancelShipmentBulk200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CanceledShipment and assigns it to the Results field.
func (o *CancelShipmentBulk200Response) SetResults(v []CanceledShipment) {
	o.Results = v
}

func (o CancelShipmentBulk200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelShipmentBulk200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableCancelShipmentBulk200Response struct {
	value *CancelShipmentBulk200Response
	isSet bool
}

func (v NullableCancelShipmentBulk200Response) Get() *CancelShipmentBulk200Response {
	return v.value
}

func (v *NullableCancelShipmentBulk200Response) Set(val *CancelShipmentBulk200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelShipmentBulk200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelShipmentBulk200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelShipmentBulk200Response(val *CancelShipmentBulk200Response) *NullableCancelShipmentBulk200Response {
	return &NullableCancelShipmentBulk200Response{value: val, isSet: true}
}

func (v NullableCancelShipmentBulk200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelShipmentBulk200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
