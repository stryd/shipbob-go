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

// PackageType the model 'PackageType'
type PackageType string

// List of PackageType
const (
	PACKAGE                PackageType = "Package"
	PALLET                 PackageType = "Pallet"
	FLOOR_LOADED_CONTAINER PackageType = "FloorLoadedContainer"
)

func (v *PackageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageType(value)
	for _, existing := range []PackageType{"Package", "Pallet", "FloorLoadedContainer"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageType", value)
}

// Ptr returns reference to PackageType value
func (v PackageType) Ptr() *PackageType {
	return &v
}

type NullablePackageType struct {
	value *PackageType
	isSet bool
}

func (v NullablePackageType) Get() *PackageType {
	return v.value
}

func (v *NullablePackageType) Set(val *PackageType) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageType) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageType(val *PackageType) *NullablePackageType {
	return &NullablePackageType{value: val, isSet: true}
}

func (v NullablePackageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
