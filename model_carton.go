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

// checks if the Carton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Carton{}

// Carton struct for Carton
type Carton struct {
	// Barcode assigned to this carton
	Barcode *string `json:"barcode,omitempty"`
	// Details about the contents of this carton
	CartonDetails []CartonDetails `json:"carton_details,omitempty"`
	// ID assigned to this carton
	Id *int32 `json:"id,omitempty"`
	Measurements *CartonMeasurements `json:"measurements,omitempty"`
	// Type of this carton container
	Type *string `json:"type,omitempty"`
}

// NewCarton instantiates a new Carton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarton() *Carton {
	this := Carton{}
	return &this
}

// NewCartonWithDefaults instantiates a new Carton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartonWithDefaults() *Carton {
	this := Carton{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *Carton) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carton) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *Carton) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *Carton) SetBarcode(v string) {
	o.Barcode = &v
}

// GetCartonDetails returns the CartonDetails field value if set, zero value otherwise.
func (o *Carton) GetCartonDetails() []CartonDetails {
	if o == nil || IsNil(o.CartonDetails) {
		var ret []CartonDetails
		return ret
	}
	return o.CartonDetails
}

// GetCartonDetailsOk returns a tuple with the CartonDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carton) GetCartonDetailsOk() ([]CartonDetails, bool) {
	if o == nil || IsNil(o.CartonDetails) {
		return nil, false
	}
	return o.CartonDetails, true
}

// HasCartonDetails returns a boolean if a field has been set.
func (o *Carton) HasCartonDetails() bool {
	if o != nil && !IsNil(o.CartonDetails) {
		return true
	}

	return false
}

// SetCartonDetails gets a reference to the given []CartonDetails and assigns it to the CartonDetails field.
func (o *Carton) SetCartonDetails(v []CartonDetails) {
	o.CartonDetails = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Carton) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carton) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Carton) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Carton) SetId(v int32) {
	o.Id = &v
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *Carton) GetMeasurements() CartonMeasurements {
	if o == nil || IsNil(o.Measurements) {
		var ret CartonMeasurements
		return ret
	}
	return *o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carton) GetMeasurementsOk() (*CartonMeasurements, bool) {
	if o == nil || IsNil(o.Measurements) {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *Carton) HasMeasurements() bool {
	if o != nil && !IsNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given CartonMeasurements and assigns it to the Measurements field.
func (o *Carton) SetMeasurements(v CartonMeasurements) {
	o.Measurements = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Carton) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carton) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Carton) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Carton) SetType(v string) {
	o.Type = &v
}

func (o Carton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Carton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.CartonDetails) {
		toSerialize["carton_details"] = o.CartonDetails
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Measurements) {
		toSerialize["measurements"] = o.Measurements
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCarton struct {
	value *Carton
	isSet bool
}

func (v NullableCarton) Get() *Carton {
	return v.value
}

func (v *NullableCarton) Set(val *Carton) {
	v.value = val
	v.isSet = true
}

func (v NullableCarton) IsSet() bool {
	return v.isSet
}

func (v *NullableCarton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarton(val *Carton) *NullableCarton {
	return &NullableCarton{value: val, isSet: true}
}

func (v NullableCarton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


