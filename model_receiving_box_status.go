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

// ReceivingBoxStatus the model 'ReceivingBoxStatus'
type ReceivingBoxStatus string

// List of Receiving.BoxStatus
const (
	AWAITING ReceivingBoxStatus = "Awaiting"
	ARRIVED ReceivingBoxStatus = "Arrived"
	COMPLETED ReceivingBoxStatus = "Completed"
	COUNTING ReceivingBoxStatus = "Counting"
	STOWING ReceivingBoxStatus = "Stowing"
	CANCELLED ReceivingBoxStatus = "Cancelled"
)

func (v *ReceivingBoxStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReceivingBoxStatus(value)
	for _, existing := range []ReceivingBoxStatus{ "Awaiting", "Arrived", "Completed", "Counting", "Stowing", "Cancelled",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReceivingBoxStatus", value)
}

// Ptr returns reference to Receiving.BoxStatus value
func (v ReceivingBoxStatus) Ptr() *ReceivingBoxStatus {
	return &v
}

type NullableReceivingBoxStatus struct {
	value *ReceivingBoxStatus
	isSet bool
}

func (v NullableReceivingBoxStatus) Get() *ReceivingBoxStatus {
	return v.value
}

func (v *NullableReceivingBoxStatus) Set(val *ReceivingBoxStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingBoxStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingBoxStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingBoxStatus(val *ReceivingBoxStatus) *NullableReceivingBoxStatus {
	return &NullableReceivingBoxStatus{value: val, isSet: true}
}

func (v NullableReceivingBoxStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingBoxStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
