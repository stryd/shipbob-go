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
)

// ReturnsReturnActionTakenViewModel struct for ReturnsReturnActionTakenViewModel
type ReturnsReturnActionTakenViewModel struct {
	Action *ReturnsReturnAction `json:"action,omitempty"`
	// Reason the given action was taken
	ActionReason NullableString `json:"action_reason,omitempty"`
	// Quantity of inventory processed with the taken action
	QuantityProcessed *int32 `json:"quantity_processed,omitempty"`
}

// NewReturnsReturnActionTakenViewModel instantiates a new ReturnsReturnActionTakenViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsReturnActionTakenViewModel() *ReturnsReturnActionTakenViewModel {
	this := ReturnsReturnActionTakenViewModel{}
	return &this
}

// NewReturnsReturnActionTakenViewModelWithDefaults instantiates a new ReturnsReturnActionTakenViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsReturnActionTakenViewModelWithDefaults() *ReturnsReturnActionTakenViewModel {
	this := ReturnsReturnActionTakenViewModel{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ReturnsReturnActionTakenViewModel) GetAction() ReturnsReturnAction {
	if o == nil || o.Action == nil {
		var ret ReturnsReturnAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsReturnActionTakenViewModel) GetActionOk() (*ReturnsReturnAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ReturnsReturnActionTakenViewModel) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given ReturnsReturnAction and assigns it to the Action field.
func (o *ReturnsReturnActionTakenViewModel) SetAction(v ReturnsReturnAction) {
	o.Action = &v
}

// GetActionReason returns the ActionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReturnsReturnActionTakenViewModel) GetActionReason() string {
	if o == nil || o.ActionReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionReason.Get()
}

// GetActionReasonOk returns a tuple with the ActionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnsReturnActionTakenViewModel) GetActionReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionReason.Get(), o.ActionReason.IsSet()
}

// HasActionReason returns a boolean if a field has been set.
func (o *ReturnsReturnActionTakenViewModel) HasActionReason() bool {
	if o != nil && o.ActionReason.IsSet() {
		return true
	}

	return false
}

// SetActionReason gets a reference to the given NullableString and assigns it to the ActionReason field.
func (o *ReturnsReturnActionTakenViewModel) SetActionReason(v string) {
	o.ActionReason.Set(&v)
}
// SetActionReasonNil sets the value for ActionReason to be an explicit nil
func (o *ReturnsReturnActionTakenViewModel) SetActionReasonNil() {
	o.ActionReason.Set(nil)
}

// UnsetActionReason ensures that no value is present for ActionReason, not even an explicit nil
func (o *ReturnsReturnActionTakenViewModel) UnsetActionReason() {
	o.ActionReason.Unset()
}

// GetQuantityProcessed returns the QuantityProcessed field value if set, zero value otherwise.
func (o *ReturnsReturnActionTakenViewModel) GetQuantityProcessed() int32 {
	if o == nil || o.QuantityProcessed == nil {
		var ret int32
		return ret
	}
	return *o.QuantityProcessed
}

// GetQuantityProcessedOk returns a tuple with the QuantityProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsReturnActionTakenViewModel) GetQuantityProcessedOk() (*int32, bool) {
	if o == nil || o.QuantityProcessed == nil {
		return nil, false
	}
	return o.QuantityProcessed, true
}

// HasQuantityProcessed returns a boolean if a field has been set.
func (o *ReturnsReturnActionTakenViewModel) HasQuantityProcessed() bool {
	if o != nil && o.QuantityProcessed != nil {
		return true
	}

	return false
}

// SetQuantityProcessed gets a reference to the given int32 and assigns it to the QuantityProcessed field.
func (o *ReturnsReturnActionTakenViewModel) SetQuantityProcessed(v int32) {
	o.QuantityProcessed = &v
}

func (o ReturnsReturnActionTakenViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ActionReason.IsSet() {
		toSerialize["action_reason"] = o.ActionReason.Get()
	}
	if o.QuantityProcessed != nil {
		toSerialize["quantity_processed"] = o.QuantityProcessed
	}
	return json.Marshal(toSerialize)
}

type NullableReturnsReturnActionTakenViewModel struct {
	value *ReturnsReturnActionTakenViewModel
	isSet bool
}

func (v NullableReturnsReturnActionTakenViewModel) Get() *ReturnsReturnActionTakenViewModel {
	return v.value
}

func (v *NullableReturnsReturnActionTakenViewModel) Set(val *ReturnsReturnActionTakenViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsReturnActionTakenViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsReturnActionTakenViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsReturnActionTakenViewModel(val *ReturnsReturnActionTakenViewModel) *NullableReturnsReturnActionTakenViewModel {
	return &NullableReturnsReturnActionTakenViewModel{value: val, isSet: true}
}

func (v NullableReturnsReturnActionTakenViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsReturnActionTakenViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


