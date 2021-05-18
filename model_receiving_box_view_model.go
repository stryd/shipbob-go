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

// ReceivingBoxViewModel Information about a box shipment included in a receiving order
type ReceivingBoxViewModel struct {
	// Date the box arrived
	ArrivedDate NullableTime `json:"arrived_date,omitempty"`
	// Information about the items included in the box
	BoxItems []ReceivingBoxItemViewModel `json:"box_items,omitempty"`
	// The number of the box in the receiving order
	BoxNumber *int32 `json:"box_number,omitempty"`
	BoxStatus *ReceivingBoxStatus `json:"box_status,omitempty"`
	// Date counting of the box's inventory items started
	CountingStartedDate NullableTime `json:"counting_started_date,omitempty"`
	// Date the box was received
	ReceivedDate NullableTime `json:"received_date,omitempty"`
	// Tracking number of the box shipment
	TrackingNumber NullableString `json:"tracking_number,omitempty"`
}

// NewReceivingBoxViewModel instantiates a new ReceivingBoxViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivingBoxViewModel() *ReceivingBoxViewModel {
	this := ReceivingBoxViewModel{}
	return &this
}

// NewReceivingBoxViewModelWithDefaults instantiates a new ReceivingBoxViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivingBoxViewModelWithDefaults() *ReceivingBoxViewModel {
	this := ReceivingBoxViewModel{}
	return &this
}

// GetArrivedDate returns the ArrivedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingBoxViewModel) GetArrivedDate() time.Time {
	if o == nil || o.ArrivedDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArrivedDate.Get()
}

// GetArrivedDateOk returns a tuple with the ArrivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingBoxViewModel) GetArrivedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArrivedDate.Get(), o.ArrivedDate.IsSet()
}

// HasArrivedDate returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasArrivedDate() bool {
	if o != nil && o.ArrivedDate.IsSet() {
		return true
	}

	return false
}

// SetArrivedDate gets a reference to the given NullableTime and assigns it to the ArrivedDate field.
func (o *ReceivingBoxViewModel) SetArrivedDate(v time.Time) {
	o.ArrivedDate.Set(&v)
}
// SetArrivedDateNil sets the value for ArrivedDate to be an explicit nil
func (o *ReceivingBoxViewModel) SetArrivedDateNil() {
	o.ArrivedDate.Set(nil)
}

// UnsetArrivedDate ensures that no value is present for ArrivedDate, not even an explicit nil
func (o *ReceivingBoxViewModel) UnsetArrivedDate() {
	o.ArrivedDate.Unset()
}

// GetBoxItems returns the BoxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingBoxViewModel) GetBoxItems() []ReceivingBoxItemViewModel {
	if o == nil  {
		var ret []ReceivingBoxItemViewModel
		return ret
	}
	return o.BoxItems
}

// GetBoxItemsOk returns a tuple with the BoxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingBoxViewModel) GetBoxItemsOk() (*[]ReceivingBoxItemViewModel, bool) {
	if o == nil || o.BoxItems == nil {
		return nil, false
	}
	return &o.BoxItems, true
}

// HasBoxItems returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasBoxItems() bool {
	if o != nil && o.BoxItems != nil {
		return true
	}

	return false
}

// SetBoxItems gets a reference to the given []ReceivingBoxItemViewModel and assigns it to the BoxItems field.
func (o *ReceivingBoxViewModel) SetBoxItems(v []ReceivingBoxItemViewModel) {
	o.BoxItems = v
}

// GetBoxNumber returns the BoxNumber field value if set, zero value otherwise.
func (o *ReceivingBoxViewModel) GetBoxNumber() int32 {
	if o == nil || o.BoxNumber == nil {
		var ret int32
		return ret
	}
	return *o.BoxNumber
}

// GetBoxNumberOk returns a tuple with the BoxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingBoxViewModel) GetBoxNumberOk() (*int32, bool) {
	if o == nil || o.BoxNumber == nil {
		return nil, false
	}
	return o.BoxNumber, true
}

// HasBoxNumber returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasBoxNumber() bool {
	if o != nil && o.BoxNumber != nil {
		return true
	}

	return false
}

// SetBoxNumber gets a reference to the given int32 and assigns it to the BoxNumber field.
func (o *ReceivingBoxViewModel) SetBoxNumber(v int32) {
	o.BoxNumber = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *ReceivingBoxViewModel) GetBoxStatus() ReceivingBoxStatus {
	if o == nil || o.BoxStatus == nil {
		var ret ReceivingBoxStatus
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceivingBoxViewModel) GetBoxStatusOk() (*ReceivingBoxStatus, bool) {
	if o == nil || o.BoxStatus == nil {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasBoxStatus() bool {
	if o != nil && o.BoxStatus != nil {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given ReceivingBoxStatus and assigns it to the BoxStatus field.
func (o *ReceivingBoxViewModel) SetBoxStatus(v ReceivingBoxStatus) {
	o.BoxStatus = &v
}

// GetCountingStartedDate returns the CountingStartedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingBoxViewModel) GetCountingStartedDate() time.Time {
	if o == nil || o.CountingStartedDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CountingStartedDate.Get()
}

// GetCountingStartedDateOk returns a tuple with the CountingStartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingBoxViewModel) GetCountingStartedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountingStartedDate.Get(), o.CountingStartedDate.IsSet()
}

// HasCountingStartedDate returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasCountingStartedDate() bool {
	if o != nil && o.CountingStartedDate.IsSet() {
		return true
	}

	return false
}

// SetCountingStartedDate gets a reference to the given NullableTime and assigns it to the CountingStartedDate field.
func (o *ReceivingBoxViewModel) SetCountingStartedDate(v time.Time) {
	o.CountingStartedDate.Set(&v)
}
// SetCountingStartedDateNil sets the value for CountingStartedDate to be an explicit nil
func (o *ReceivingBoxViewModel) SetCountingStartedDateNil() {
	o.CountingStartedDate.Set(nil)
}

// UnsetCountingStartedDate ensures that no value is present for CountingStartedDate, not even an explicit nil
func (o *ReceivingBoxViewModel) UnsetCountingStartedDate() {
	o.CountingStartedDate.Unset()
}

// GetReceivedDate returns the ReceivedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingBoxViewModel) GetReceivedDate() time.Time {
	if o == nil || o.ReceivedDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ReceivedDate.Get()
}

// GetReceivedDateOk returns a tuple with the ReceivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingBoxViewModel) GetReceivedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReceivedDate.Get(), o.ReceivedDate.IsSet()
}

// HasReceivedDate returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasReceivedDate() bool {
	if o != nil && o.ReceivedDate.IsSet() {
		return true
	}

	return false
}

// SetReceivedDate gets a reference to the given NullableTime and assigns it to the ReceivedDate field.
func (o *ReceivingBoxViewModel) SetReceivedDate(v time.Time) {
	o.ReceivedDate.Set(&v)
}
// SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil
func (o *ReceivingBoxViewModel) SetReceivedDateNil() {
	o.ReceivedDate.Set(nil)
}

// UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
func (o *ReceivingBoxViewModel) UnsetReceivedDate() {
	o.ReceivedDate.Unset()
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivingBoxViewModel) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber.Get()
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivingBoxViewModel) GetTrackingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingNumber.Get(), o.TrackingNumber.IsSet()
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ReceivingBoxViewModel) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber.IsSet() {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given NullableString and assigns it to the TrackingNumber field.
func (o *ReceivingBoxViewModel) SetTrackingNumber(v string) {
	o.TrackingNumber.Set(&v)
}
// SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil
func (o *ReceivingBoxViewModel) SetTrackingNumberNil() {
	o.TrackingNumber.Set(nil)
}

// UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
func (o *ReceivingBoxViewModel) UnsetTrackingNumber() {
	o.TrackingNumber.Unset()
}

func (o ReceivingBoxViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArrivedDate.IsSet() {
		toSerialize["arrived_date"] = o.ArrivedDate.Get()
	}
	if o.BoxItems != nil {
		toSerialize["box_items"] = o.BoxItems
	}
	if o.BoxNumber != nil {
		toSerialize["box_number"] = o.BoxNumber
	}
	if o.BoxStatus != nil {
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
	return json.Marshal(toSerialize)
}

type NullableReceivingBoxViewModel struct {
	value *ReceivingBoxViewModel
	isSet bool
}

func (v NullableReceivingBoxViewModel) Get() *ReceivingBoxViewModel {
	return v.value
}

func (v *NullableReceivingBoxViewModel) Set(val *ReceivingBoxViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivingBoxViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivingBoxViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivingBoxViewModel(val *ReceivingBoxViewModel) *NullableReceivingBoxViewModel {
	return &NullableReceivingBoxViewModel{value: val, isSet: true}
}

func (v NullableReceivingBoxViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivingBoxViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

