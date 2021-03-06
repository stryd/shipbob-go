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

// ReturnAction the model 'ReturnAction'
type ReturnAction string

// List of ReturnAction
const (
	DEFAULT    ReturnAction = "Default"
	RESTOCK    ReturnAction = "Restock"
	QUARANTINE ReturnAction = "Quarantine"
	DISPOSE    ReturnAction = "Dispose"
)

func (v *ReturnAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnAction(value)
	for _, existing := range []ReturnAction{"Default", "Restock", "Quarantine", "Dispose"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnAction", value)
}

// Ptr returns reference to ReturnAction value
func (v ReturnAction) Ptr() *ReturnAction {
	return &v
}

type NullableReturnAction struct {
	value *ReturnAction
	isSet bool
}

func (v NullableReturnAction) Get() *ReturnAction {
	return v.value
}

func (v *NullableReturnAction) Set(val *ReturnAction) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnAction) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnAction(val *ReturnAction) *NullableReturnAction {
	return &NullableReturnAction{value: val, isSet: true}
}

func (v NullableReturnAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
