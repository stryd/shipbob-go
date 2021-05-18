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

// ReturnsReturnOrderStatusHistoryViewModel struct for ReturnsReturnOrderStatusHistoryViewModel
type ReturnsReturnOrderStatusHistoryViewModel struct {
	Status *ReturnsReturnStatus `json:"status,omitempty"`
	// Date this corresponding return order status was created
	TimeStamp *time.Time `json:"time_stamp,omitempty"`
}

// NewReturnsReturnOrderStatusHistoryViewModel instantiates a new ReturnsReturnOrderStatusHistoryViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsReturnOrderStatusHistoryViewModel() *ReturnsReturnOrderStatusHistoryViewModel {
	this := ReturnsReturnOrderStatusHistoryViewModel{}
	return &this
}

// NewReturnsReturnOrderStatusHistoryViewModelWithDefaults instantiates a new ReturnsReturnOrderStatusHistoryViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsReturnOrderStatusHistoryViewModelWithDefaults() *ReturnsReturnOrderStatusHistoryViewModel {
	this := ReturnsReturnOrderStatusHistoryViewModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReturnsReturnOrderStatusHistoryViewModel) GetStatus() ReturnsReturnStatus {
	if o == nil || o.Status == nil {
		var ret ReturnsReturnStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsReturnOrderStatusHistoryViewModel) GetStatusOk() (*ReturnsReturnStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReturnsReturnOrderStatusHistoryViewModel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReturnsReturnStatus and assigns it to the Status field.
func (o *ReturnsReturnOrderStatusHistoryViewModel) SetStatus(v ReturnsReturnStatus) {
	o.Status = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *ReturnsReturnOrderStatusHistoryViewModel) GetTimeStamp() time.Time {
	if o == nil || o.TimeStamp == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsReturnOrderStatusHistoryViewModel) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || o.TimeStamp == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *ReturnsReturnOrderStatusHistoryViewModel) HasTimeStamp() bool {
	if o != nil && o.TimeStamp != nil {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *ReturnsReturnOrderStatusHistoryViewModel) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

func (o ReturnsReturnOrderStatusHistoryViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TimeStamp != nil {
		toSerialize["time_stamp"] = o.TimeStamp
	}
	return json.Marshal(toSerialize)
}

type NullableReturnsReturnOrderStatusHistoryViewModel struct {
	value *ReturnsReturnOrderStatusHistoryViewModel
	isSet bool
}

func (v NullableReturnsReturnOrderStatusHistoryViewModel) Get() *ReturnsReturnOrderStatusHistoryViewModel {
	return v.value
}

func (v *NullableReturnsReturnOrderStatusHistoryViewModel) Set(val *ReturnsReturnOrderStatusHistoryViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsReturnOrderStatusHistoryViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsReturnOrderStatusHistoryViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsReturnOrderStatusHistoryViewModel(val *ReturnsReturnOrderStatusHistoryViewModel) *NullableReturnsReturnOrderStatusHistoryViewModel {
	return &NullableReturnsReturnOrderStatusHistoryViewModel{value: val, isSet: true}
}

func (v NullableReturnsReturnOrderStatusHistoryViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsReturnOrderStatusHistoryViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

