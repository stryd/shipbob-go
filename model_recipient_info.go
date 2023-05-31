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

// checks if the RecipientInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientInfo{}

// RecipientInfo Information about the recipient of an order
type RecipientInfo struct {
	Address OrderAddress `json:"address"`
	// Email address of the recipient
	Email *string `json:"email,omitempty"`
	// Name of the recipient
	Name string `json:"name"`
	// Phone number of the recipient
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// NewRecipientInfo instantiates a new RecipientInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientInfo(address OrderAddress, name string) *RecipientInfo {
	this := RecipientInfo{}
	this.Address = address
	this.Name = name
	return &this
}

// NewRecipientInfoWithDefaults instantiates a new RecipientInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientInfoWithDefaults() *RecipientInfo {
	this := RecipientInfo{}
	return &this
}

// GetAddress returns the Address field value
func (o *RecipientInfo) GetAddress() OrderAddress {
	if o == nil {
		var ret OrderAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetAddressOk() (*OrderAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RecipientInfo) SetAddress(v OrderAddress) {
	o.Address = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RecipientInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RecipientInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RecipientInfo) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value
func (o *RecipientInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RecipientInfo) SetName(v string) {
	o.Name = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *RecipientInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *RecipientInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *RecipientInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o RecipientInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableRecipientInfo struct {
	value *RecipientInfo
	isSet bool
}

func (v NullableRecipientInfo) Get() *RecipientInfo {
	return v.value
}

func (v *NullableRecipientInfo) Set(val *RecipientInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientInfo(val *RecipientInfo) *NullableRecipientInfo {
	return &NullableRecipientInfo{value: val, isSet: true}
}

func (v NullableRecipientInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
