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

// checks if the ReceivingOrderBoxesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivingOrderBoxesInner{}

// ReceivingOrderBoxesInner Information about a box shipment included in a receiving order
type ReceivingOrderBoxesInner struct {
	// Date the box arrived
	ArrivedDate NullableTime `json:"arrived_date,omitempty"`
	// Information about the items included in the box
	BoxItems []ReceivingOrderBoxesInnerBoxItemsInner `json:"box_items,omitempty"`
	// The number of the box in the receiving order
	BoxNumber *int32 `json:"box_number,omitempty"`
	BoxStatus *string `json:"box_status,omitempty"`
	// Date counting of the box's inventory items started
	CountingStartedDate NullableTime `json:"counting_started_date,omitempty"`
	// Date the box was received
	ReceivedDate NullableTime `json:"received_date,omitempty"`
	// Tracking number of the box shipment
	TrackingNumber NullableString `json:"tracking_number,omitempty"`
}

// NewReceivingOrderBoxesInner instantiates a new ReceivingOrderBoxesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivingOrderBoxesInner() *ReceivingOrderBoxesInner {
	this := ReceivingOrderBoxesInner{}
	return &this
}

// NewReceivingOrderBoxesInnerWithDefaults instantiates a new ReceivingOrderBoxesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivingOrderBoxesInnerWithDefaults() *ReceivingOrderBoxesInner {
	this := ReceivingOrderBoxesInner{}
	return &this
}

// GetArrivedDate returns the ArrivedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInner) GetArrivedDate() time.Time {
	if o == nil || IsNil(o.ArrivedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ArrivedDate.Get()
}

// GetArrivedDateOk returns a tuple with the ArrivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInner) GetArrivedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrivedDate.Get(), o.ArrivedDate.IsSet()
}

// HasArrivedDate returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasArrivedDate() bool {
	if o != nil && o.ArrivedDate.IsSet() {
		return true
	}

	return false
}

// SetArrivedDate gets a reference to the given NullableTime and assigns it to the ArrivedDate field.
func (o *ReceivingOrderBoxesInner) SetArrivedDate(v time.Time) {
	o.ArrivedDate.Set(&v)
}
// SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil
func (o *ReceivingOrderBoxesInner) SetArrivedDateNil() {
	o.ArrivedDate.Set(nil)
}

// UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
func (o *ReceivingOrderBoxesInner) UnsetArrivedDate() {
	o.ArrivedDate.Unset()
}

// GetBoxItems returns the BoxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInner) GetBoxItems() []ReceivingOrderBoxesInnerBoxItemsInner {
	if o == nil {
		var ret []ReceivingOrderBoxesInnerBoxItemsInner
		return ret
	}
	return o.BoxItems
}

// GetBoxItemsOk returns a tuple with the BoxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInner) GetBoxItemsOk() ([]ReceivingOrderBoxesInnerBoxItemsInner, bool) {
	if o == nil || IsNil(o.BoxItems) {
		return nil, false
	}
	return o.BoxItems, true
}

// HasBoxItems returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasBoxItems() bool {
	if o != nil && IsNil(o.BoxItems) {
		return true
	}

	return false
}

// SetBoxItems gets a reference to the given []ReceivingOrderBoxesInnerBoxItemsInner and assigns it to the BoxItems field.
func (o *ReceivingOrderBoxesInner) SetBoxItems(v []ReceivingOrderBoxesInnerBoxItemsInner) {
	o.BoxItems = v
}

// GetBoxNumber returns the BoxNumber field value if set, zero value otherwise.
func (o *ReceivingOrderBoxesInner) GetBoxNumber() int32 {
	if o == nil || IsNil(o.BoxNumber) {
		var ret int32
		return ret
	}
	return *o.BoxNumber
}

// GetBoxNumberOk returns a tuple with the BoxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingOrderBoxesInner) GetBoxNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BoxNumber) {
		return nil, false
	}
	return o.BoxNumber, true
}

// HasBoxNumber returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasBoxNumber() bool {
	if o != nil && !IsNil(o.BoxNumber) {
		return true
	}

	return false
}

// SetBoxNumber gets a reference to the given int32 and assigns it to the BoxNumber field.
func (o *ReceivingOrderBoxesInner) SetBoxNumber(v int32) {
	o.BoxNumber = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *ReceivingOrderBoxesInner) GetBoxStatus() string {
	if o == nil || IsNil(o.BoxStatus) {
		var ret string
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingOrderBoxesInner) GetBoxStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BoxStatus) {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasBoxStatus() bool {
	if o != nil && !IsNil(o.BoxStatus) {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given string and assigns it to the BoxStatus field.
func (o *ReceivingOrderBoxesInner) SetBoxStatus(v string) {
	o.BoxStatus = &v
}

// GetCountingStartedDate returns the CountingStartedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInner) GetCountingStartedDate() time.Time {
	if o == nil || IsNil(o.CountingStartedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CountingStartedDate.Get()
}

// GetCountingStartedDateOk returns a tuple with the CountingStartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInner) GetCountingStartedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountingStartedDate.Get(), o.CountingStartedDate.IsSet()
}

// HasCountingStartedDate returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasCountingStartedDate() bool {
	if o != nil && o.CountingStartedDate.IsSet() {
		return true
	}

	return false
}

// SetCountingStartedDate gets a reference to the given NullableTime and assigns it to the CountingStartedDate field.
func (o *ReceivingOrderBoxesInner) SetCountingStartedDate(v time.Time) {
	o.CountingStartedDate.Set(&v)
}
// SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil
func (o *ReceivingOrderBoxesInner) SetCountingStartedDateNil() {
	o.CountingStartedDate.Set(nil)
}

// UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
func (o *ReceivingOrderBoxesInner) UnsetCountingStartedDate() {
	o.CountingStartedDate.Unset()
}

// GetReceivedDate returns the ReceivedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInner) GetReceivedDate() time.Time {
	if o == nil || IsNil(o.ReceivedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReceivedDate.Get()
}

// GetReceivedDateOk returns a tuple with the ReceivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInner) GetReceivedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedDate.Get(), o.ReceivedDate.IsSet()
}

// HasReceivedDate returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasReceivedDate() bool {
	if o != nil && o.ReceivedDate.IsSet() {
		return true
	}

	return false
}

// SetReceivedDate gets a reference to the given NullableTime and assigns it to the ReceivedDate field.
func (o *ReceivingOrderBoxesInner) SetReceivedDate(v time.Time) {
	o.ReceivedDate.Set(&v)
}
// SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil
func (o *ReceivingOrderBoxesInner) SetReceivedDateNil() {
	o.ReceivedDate.Set(nil)
}

// UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
func (o *ReceivingOrderBoxesInner) UnsetReceivedDate() {
	o.ReceivedDate.Unset()
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingOrderBoxesInner) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber.Get()) {
		var ret string
		return ret
	}
	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingOrderBoxesInner) GetTrackingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ReceivingOrderBoxesInner) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber.IsSet() {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given NullableString and assigns it to the TrackingNumber field.
func (o *ReceivingOrderBoxesInner) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}
// SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil
func (o *ReceivingOrderBoxesInner) SetTrackingNumberNil() {
	o.TrackingNumber.Set(nil)
}

// UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
func (o *ReceivingOrderBoxesInner) UnsetTrackingNumber() {
	o.TrackingNumber.Unset()
}

func (o ReceivingOrderBoxesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivingOrderBoxesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ArrivedDate.IsSet() {
		toSerialize["arrived_date"] = o.ArrivedDate.Get()
	}
	if o.BoxItems != nil {
		toSerialize["box_items"] = o.BoxItems
	}
	if !IsNil(o.BoxNumber) {
		toSerialize["box_number"] = o.BoxNumber
	}
	if !IsNil(o.BoxStatus) {
		toSerialize["box_status"] = o.BoxStatus
	}
	if o.CountingStartedDate.IsSet() {
		toSerialize["counting_started_date"] = o.CountingStartedDate.Get()
	}
	if o.ReceivedDate.IsSet() {
		toSerialize["received_date"] = o.ReceivedDate.Get()
	}
	if o.TrackingNumber.IsSet() {
		toSerialize["tracking_number"] = o.TrackingNumber.Get()
	}
	return toSerialize, nil
}

type NullableReceivingOrderBoxesInner struct {
	value *ReceivingOrderBoxesInner
	isSet bool
}

func (v NullableReceivingOrderBoxesInner) Get() *ReceivingOrderBoxesInner {
	return v.value
}

func (v *NullableReceivingOrderBoxesInner) Set(val *ReceivingOrderBoxesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingOrderBoxesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingOrderBoxesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingOrderBoxesInner(val *ReceivingOrderBoxesInner) *NullableReceivingOrderBoxesInner {
	return &NullableReceivingOrderBoxesInner{value: val, isSet: true}
}

func (v NullableReceivingOrderBoxesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingOrderBoxesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


