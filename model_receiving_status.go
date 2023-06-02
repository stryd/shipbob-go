/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"fmt"
)

// ReceivingStatus the model 'ReceivingStatus'
type ReceivingStatus string

// List of ReceivingStatus
const (
	AWAITING          ReceivingStatus = "Awaiting"
	PROCESSING        ReceivingStatus = "Processing"
	COMPLETED         ReceivingStatus = "Completed"
	CANCELLED         ReceivingStatus = "Cancelled"
	INCOMPLETE        ReceivingStatus = "Incomplete"
	ARRIVED           ReceivingStatus = "Arrived"
	PARTIALLY_ARRIVED ReceivingStatus = "PartiallyArrived"
)

// All allowed values of ReceivingStatus enum
var AllowedReceivingStatusEnumValues = []ReceivingStatus{
	"Awaiting",
	"Processing",
	"Completed",
	"Cancelled",
	"Incomplete",
	"Arrived",
	"PartiallyArrived",
}

func (v *ReceivingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReceivingStatus(value)
	for _, existing := range AllowedReceivingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReceivingStatus", value)
}

// NewReceivingStatusFromValue returns a pointer to a valid ReceivingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReceivingStatusFromValue(v string) (*ReceivingStatus, error) {
	ev := ReceivingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReceivingStatus: valid values are %v", v, AllowedReceivingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReceivingStatus) IsValid() bool {
	for _, existing := range AllowedReceivingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReceivingStatus value
func (v ReceivingStatus) Ptr() *ReceivingStatus {
	return &v
}

type NullableReceivingStatus struct {
	value *ReceivingStatus
	isSet bool
}

func (v NullableReceivingStatus) Get() *ReceivingStatus {
	return v.value
}

func (v *NullableReceivingStatus) Set(val *ReceivingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingStatus(val *ReceivingStatus) *NullableReceivingStatus {
	return &NullableReceivingStatus{value: val, isSet: true}
}

func (v NullableReceivingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}