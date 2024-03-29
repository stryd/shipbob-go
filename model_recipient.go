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

// checks if the Recipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recipient{}

// Recipient Information about the recipient of a shipment
type Recipient struct {
	Address *RetailerProgramDataAddress `json:"address,omitempty"`
	// Email address of the recipient
	Email *string `json:"email,omitempty"`
	// Name of the recipient
	Name *string `json:"name,omitempty"`
	// Phone number of the recipient
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// NewRecipient instantiates a new Recipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipient() *Recipient {
	this := Recipient{}
	return &this
}

// NewRecipientWithDefaults instantiates a new Recipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientWithDefaults() *Recipient {
	this := Recipient{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Recipient) GetAddress() RetailerProgramDataAddress {
	if o == nil || IsNil(o.Address) {
		var ret RetailerProgramDataAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetAddressOk() (*RetailerProgramDataAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Recipient) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given RetailerProgramDataAddress and assigns it to the Address field.
func (o *Recipient) SetAddress(v RetailerProgramDataAddress) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Recipient) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Recipient) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Recipient) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Recipient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Recipient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Recipient) SetName(v string) {
	o.Name = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Recipient) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipient) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Recipient) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Recipient) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o Recipient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableRecipient struct {
	value *Recipient
	isSet bool
}

func (v NullableRecipient) Get() *Recipient {
	return v.value
}

func (v *NullableRecipient) Set(val *Recipient) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipient(val *Recipient) *NullableRecipient {
	return &NullableRecipient{value: val, isSet: true}
}

func (v NullableRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
