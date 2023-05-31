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

// checks if the ReturnTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnTransaction{}

// ReturnTransaction struct for ReturnTransaction
type ReturnTransaction struct {
	// Transaction amount in dollars
	Amount *float64 `json:"amount,omitempty"`
	TransactionType *string `json:"transaction_type,omitempty"`
}

// NewReturnTransaction instantiates a new ReturnTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnTransaction() *ReturnTransaction {
	this := ReturnTransaction{}
	return &this
}

// NewReturnTransactionWithDefaults instantiates a new ReturnTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnTransactionWithDefaults() *ReturnTransaction {
	this := ReturnTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ReturnTransaction) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransaction) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ReturnTransaction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *ReturnTransaction) SetAmount(v float64) {
	o.Amount = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *ReturnTransaction) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnTransaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *ReturnTransaction) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *ReturnTransaction) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o ReturnTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transaction_type"] = o.TransactionType
	}
	return toSerialize, nil
}

type NullableReturnTransaction struct {
	value *ReturnTransaction
	isSet bool
}

func (v NullableReturnTransaction) Get() *ReturnTransaction {
	return v.value
}

func (v *NullableReturnTransaction) Set(val *ReturnTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnTransaction(val *ReturnTransaction) *NullableReturnTransaction {
	return &NullableReturnTransaction{value: val, isSet: true}
}

func (v NullableReturnTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


