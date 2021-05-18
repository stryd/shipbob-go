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
	"fmt"
)

// ReturnsReturnStatus the model 'ReturnsReturnStatus'
type ReturnsReturnStatus string

// List of Returns.ReturnStatus
const (
	AWAITING_ARRIVAL ReturnsReturnStatus = "AwaitingArrival"
	ARRIVED ReturnsReturnStatus = "Arrived"
	PROCESSING ReturnsReturnStatus = "Processing"
	COMPLETED ReturnsReturnStatus = "Completed"
	CANCELLED ReturnsReturnStatus = "Cancelled"
)

func (v *ReturnsReturnStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnsReturnStatus(value)
	for _, existing := range []ReturnsReturnStatus{ "AwaitingArrival", "Arrived", "Processing", "Completed", "Cancelled",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnsReturnStatus", value)
}

// Ptr returns reference to Returns.ReturnStatus value
func (v ReturnsReturnStatus) Ptr() *ReturnsReturnStatus {
	return &v
}

type NullableReturnsReturnStatus struct {
	value *ReturnsReturnStatus
	isSet bool
}

func (v NullableReturnsReturnStatus) Get() *ReturnsReturnStatus {
	return v.value
}

func (v *NullableReturnsReturnStatus) Set(val *ReturnsReturnStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsReturnStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsReturnStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsReturnStatus(val *ReturnsReturnStatus) *NullableReturnsReturnStatus {
	return &NullableReturnsReturnStatus{value: val, isSet: true}
}

func (v NullableReturnsReturnStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsReturnStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

