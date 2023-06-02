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

// ReturnTransactionLogSource the model 'ReturnTransactionLogSource'
type ReturnTransactionLogSource string

// List of ReturnTransactionLogSource
const (
	RETURN_LABEL_INVOICE  ReturnTransactionLogSource = "ReturnLabelInvoice"
	RETURN_PROCESSING_FEE ReturnTransactionLogSource = "ReturnProcessingFee"
	RETURN_TO_SENDER_FEE  ReturnTransactionLogSource = "ReturnToSenderFee"
)

// All allowed values of ReturnTransactionLogSource enum
var AllowedReturnTransactionLogSourceEnumValues = []ReturnTransactionLogSource{
	"ReturnLabelInvoice",
	"ReturnProcessingFee",
	"ReturnToSenderFee",
}

func (v *ReturnTransactionLogSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnTransactionLogSource(value)
	for _, existing := range AllowedReturnTransactionLogSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnTransactionLogSource", value)
}

// NewReturnTransactionLogSourceFromValue returns a pointer to a valid ReturnTransactionLogSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnTransactionLogSourceFromValue(v string) (*ReturnTransactionLogSource, error) {
	ev := ReturnTransactionLogSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnTransactionLogSource: valid values are %v", v, AllowedReturnTransactionLogSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnTransactionLogSource) IsValid() bool {
	for _, existing := range AllowedReturnTransactionLogSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnTransactionLogSource value
func (v ReturnTransactionLogSource) Ptr() *ReturnTransactionLogSource {
	return &v
}

type NullableReturnTransactionLogSource struct {
	value *ReturnTransactionLogSource
	isSet bool
}

func (v NullableReturnTransactionLogSource) Get() *ReturnTransactionLogSource {
	return v.value
}

func (v *NullableReturnTransactionLogSource) Set(val *ReturnTransactionLogSource) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnTransactionLogSource) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnTransactionLogSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnTransactionLogSource(val *ReturnTransactionLogSource) *NullableReturnTransactionLogSource {
	return &NullableReturnTransactionLogSource{value: val, isSet: true}
}

func (v NullableReturnTransactionLogSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnTransactionLogSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
