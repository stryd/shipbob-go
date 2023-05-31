/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"time"
)

// checks if the ShipmentLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentLog{}

// ShipmentLog struct for ShipmentLog
type ShipmentLog struct {
	// Log type id of the shipment
	LogTypeId *int32 `json:"log_type_id,omitempty"`
	// Name of the log type
	LogTypeName *string `json:"log_type_name,omitempty"`
	// Summary of log type meaning
	LogTypeText *string `json:"log_type_text,omitempty"`
	// Specifics data for the event
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Timestamp of event
	Timestamp NullableTime `json:"timestamp,omitempty"`
}

// NewShipmentLog instantiates a new ShipmentLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentLog() *ShipmentLog {
	this := ShipmentLog{}
	return &this
}

// NewShipmentLogWithDefaults instantiates a new ShipmentLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentLogWithDefaults() *ShipmentLog {
	this := ShipmentLog{}
	return &this
}

// GetLogTypeId returns the LogTypeId field value if set, zero value otherwise.
func (o *ShipmentLog) GetLogTypeId() int32 {
	if o == nil || IsNil(o.LogTypeId) {
		var ret int32
		return ret
	}
	return *o.LogTypeId
}

// GetLogTypeIdOk returns a tuple with the LogTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentLog) GetLogTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LogTypeId) {
		return nil, false
	}
	return o.LogTypeId, true
}

// HasLogTypeId returns a boolean if a field has been set.
func (o *ShipmentLog) HasLogTypeId() bool {
	if o != nil && !IsNil(o.LogTypeId) {
		return true
	}

	return false
}

// SetLogTypeId gets a reference to the given int32 and assigns it to the LogTypeId field.
func (o *ShipmentLog) SetLogTypeId(v int32) {
	o.LogTypeId = &v
}

// GetLogTypeName returns the LogTypeName field value if set, zero value otherwise.
func (o *ShipmentLog) GetLogTypeName() string {
	if o == nil || IsNil(o.LogTypeName) {
		var ret string
		return ret
	}
	return *o.LogTypeName
}

// GetLogTypeNameOk returns a tuple with the LogTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentLog) GetLogTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogTypeName) {
		return nil, false
	}
	return o.LogTypeName, true
}

// HasLogTypeName returns a boolean if a field has been set.
func (o *ShipmentLog) HasLogTypeName() bool {
	if o != nil && !IsNil(o.LogTypeName) {
		return true
	}

	return false
}

// SetLogTypeName gets a reference to the given string and assigns it to the LogTypeName field.
func (o *ShipmentLog) SetLogTypeName(v string) {
	o.LogTypeName = &v
}

// GetLogTypeText returns the LogTypeText field value if set, zero value otherwise.
func (o *ShipmentLog) GetLogTypeText() string {
	if o == nil || IsNil(o.LogTypeText) {
		var ret string
		return ret
	}
	return *o.LogTypeText
}

// GetLogTypeTextOk returns a tuple with the LogTypeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentLog) GetLogTypeTextOk() (*string, bool) {
	if o == nil || IsNil(o.LogTypeText) {
		return nil, false
	}
	return o.LogTypeText, true
}

// HasLogTypeText returns a boolean if a field has been set.
func (o *ShipmentLog) HasLogTypeText() bool {
	if o != nil && !IsNil(o.LogTypeText) {
		return true
	}

	return false
}

// SetLogTypeText gets a reference to the given string and assigns it to the LogTypeText field.
func (o *ShipmentLog) SetLogTypeText(v string) {
	o.LogTypeText = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShipmentLog) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentLog) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShipmentLog) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ShipmentLog) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipmentLog) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipmentLog) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ShipmentLog) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *ShipmentLog) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ShipmentLog) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ShipmentLog) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o ShipmentLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogTypeId) {
		toSerialize["log_type_id"] = o.LogTypeId
	}
	if !IsNil(o.LogTypeName) {
		toSerialize["log_type_name"] = o.LogTypeName
	}
	if !IsNil(o.LogTypeText) {
		toSerialize["log_type_text"] = o.LogTypeText
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return toSerialize, nil
}

type NullableShipmentLog struct {
	value *ShipmentLog
	isSet bool
}

func (v NullableShipmentLog) Get() *ShipmentLog {
	return v.value
}

func (v *NullableShipmentLog) Set(val *ShipmentLog) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentLog) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentLog(val *ShipmentLog) *NullableShipmentLog {
	return &NullableShipmentLog{value: val, isSet: true}
}

func (v NullableShipmentLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
