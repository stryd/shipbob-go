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
	"time"
)

// OrdersShipmentLogViewModel struct for OrdersShipmentLogViewModel
type OrdersShipmentLogViewModel struct {
	// Log type id of the shipment
	LogTypeId *int32 `json:"log_type_id,omitempty"`
	// Name of the log type
	LogTypeName *string `json:"log_type_name,omitempty"`
	// Summary of log type meaning
	LogTypeText *string `json:"log_type_text,omitempty"`
	// Specifics data for the event
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Timestamp of event
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewOrdersShipmentLogViewModel instantiates a new OrdersShipmentLogViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersShipmentLogViewModel() *OrdersShipmentLogViewModel {
	this := OrdersShipmentLogViewModel{}
	return &this
}

// NewOrdersShipmentLogViewModelWithDefaults instantiates a new OrdersShipmentLogViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersShipmentLogViewModelWithDefaults() *OrdersShipmentLogViewModel {
	this := OrdersShipmentLogViewModel{}
	return &this
}

// GetLogTypeId returns the LogTypeId field value if set, zero value otherwise.
func (o *OrdersShipmentLogViewModel) GetLogTypeId() int32 {
	if o == nil || o.LogTypeId == nil {
		var ret int32
		return ret
	}
	return *o.LogTypeId
}

// GetLogTypeIdOk returns a tuple with the LogTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipmentLogViewModel) GetLogTypeIdOk() (*int32, bool) {
	if o == nil || o.LogTypeId == nil {
		return nil, false
	}
	return o.LogTypeId, true
}

// HasLogTypeId returns a boolean if a field has been set.
func (o *OrdersShipmentLogViewModel) HasLogTypeId() bool {
	if o != nil && o.LogTypeId != nil {
		return true
	}

	return false
}

// SetLogTypeId gets a reference to the given int32 and assigns it to the LogTypeId field.
func (o *OrdersShipmentLogViewModel) SetLogTypeId(v int32) {
	o.LogTypeId = &v
}

// GetLogTypeName returns the LogTypeName field value if set, zero value otherwise.
func (o *OrdersShipmentLogViewModel) GetLogTypeName() string {
	if o == nil || o.LogTypeName == nil {
		var ret string
		return ret
	}
	return *o.LogTypeName
}

// GetLogTypeNameOk returns a tuple with the LogTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipmentLogViewModel) GetLogTypeNameOk() (*string, bool) {
	if o == nil || o.LogTypeName == nil {
		return nil, false
	}
	return o.LogTypeName, true
}

// HasLogTypeName returns a boolean if a field has been set.
func (o *OrdersShipmentLogViewModel) HasLogTypeName() bool {
	if o != nil && o.LogTypeName != nil {
		return true
	}

	return false
}

// SetLogTypeName gets a reference to the given string and assigns it to the LogTypeName field.
func (o *OrdersShipmentLogViewModel) SetLogTypeName(v string) {
	o.LogTypeName = &v
}

// GetLogTypeText returns the LogTypeText field value if set, zero value otherwise.
func (o *OrdersShipmentLogViewModel) GetLogTypeText() string {
	if o == nil || o.LogTypeText == nil {
		var ret string
		return ret
	}
	return *o.LogTypeText
}

// GetLogTypeTextOk returns a tuple with the LogTypeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipmentLogViewModel) GetLogTypeTextOk() (*string, bool) {
	if o == nil || o.LogTypeText == nil {
		return nil, false
	}
	return o.LogTypeText, true
}

// HasLogTypeText returns a boolean if a field has been set.
func (o *OrdersShipmentLogViewModel) HasLogTypeText() bool {
	if o != nil && o.LogTypeText != nil {
		return true
	}

	return false
}

// SetLogTypeText gets a reference to the given string and assigns it to the LogTypeText field.
func (o *OrdersShipmentLogViewModel) SetLogTypeText(v string) {
	o.LogTypeText = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrdersShipmentLogViewModel) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipmentLogViewModel) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrdersShipmentLogViewModel) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OrdersShipmentLogViewModel) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OrdersShipmentLogViewModel) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipmentLogViewModel) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OrdersShipmentLogViewModel) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *OrdersShipmentLogViewModel) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o OrdersShipmentLogViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogTypeId != nil {
		toSerialize["log_type_id"] = o.LogTypeId
	}
	if o.LogTypeName != nil {
		toSerialize["log_type_name"] = o.LogTypeName
	}
	if o.LogTypeText != nil {
		toSerialize["log_type_text"] = o.LogTypeText
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableOrdersShipmentLogViewModel struct {
	value *OrdersShipmentLogViewModel
	isSet bool
}

func (v NullableOrdersShipmentLogViewModel) Get() *OrdersShipmentLogViewModel {
	return v.value
}

func (v *NullableOrdersShipmentLogViewModel) Set(val *OrdersShipmentLogViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersShipmentLogViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersShipmentLogViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersShipmentLogViewModel(val *OrdersShipmentLogViewModel) *NullableOrdersShipmentLogViewModel {
	return &NullableOrdersShipmentLogViewModel{value: val, isSet: true}
}

func (v NullableOrdersShipmentLogViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersShipmentLogViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

