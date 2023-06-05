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

// BoxStatus the model 'BoxStatus'
type BoxStatus string

// List of BoxStatus
const (
	BOXSTATUS_AWAITING  BoxStatus = "Awaiting"
	BOXSTATUS_ARRIVED   BoxStatus = "Arrived"
	BOXSTATUS_COMPLETED BoxStatus = "Completed"
	BOXSTATUS_COUNTING  BoxStatus = "Counting"
	BOXSTATUS_STOWING   BoxStatus = "Stowing"
	BOXSTATUS_CANCELLED BoxStatus = "Cancelled"
)

// All allowed values of BoxStatus enum
var AllowedBoxStatusEnumValues = []BoxStatus{
	"Awaiting",
	"Arrived",
	"Completed",
	"Counting",
	"Stowing",
	"Cancelled",
}

func (v *BoxStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BoxStatus(value)
	for _, existing := range AllowedBoxStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BoxStatus", value)
}

// NewBoxStatusFromValue returns a pointer to a valid BoxStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBoxStatusFromValue(v string) (*BoxStatus, error) {
	ev := BoxStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BoxStatus: valid values are %v", v, AllowedBoxStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BoxStatus) IsValid() bool {
	for _, existing := range AllowedBoxStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BoxStatus value
func (v BoxStatus) Ptr() *BoxStatus {
	return &v
}

type NullableBoxStatus struct {
	value *BoxStatus
	isSet bool
}

func (v NullableBoxStatus) Get() *BoxStatus {
	return v.value
}

func (v *NullableBoxStatus) Set(val *BoxStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxStatus(val *BoxStatus) *NullableBoxStatus {
	return &NullableBoxStatus{value: val, isSet: true}
}

func (v NullableBoxStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoxStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
