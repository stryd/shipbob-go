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

// checks if the LocationAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationAddress{}

// LocationAddress The service-specific address of the location. Each object contains address type, address1, address2, city, state, country, zip code, phone number, and email
type LocationAddress struct {
	// First part of the address of the location for this service
	Address1 NullableString `json:"address1,omitempty"`
	// Second part of the address of the location for this service
	Address2 NullableString `json:"address2,omitempty"`
	// City of the location
	City NullableString `json:"city,omitempty"`
	// Country of the location
	Country NullableString `json:"country,omitempty"`
	// Email of the location for this service
	Email NullableString `json:"email,omitempty"`
	// Name to use in the address of the location for this service
	Name NullableString `json:"name,omitempty"`
	// Phone Number of the location for this service
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// State of the location
	State NullableString `json:"state,omitempty"`
	// Zip code of the location
	ZipCode NullableString `json:"zip_code,omitempty"`
}

// NewLocationAddress instantiates a new LocationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAddress() *LocationAddress {
	this := LocationAddress{}
	return &this
}

// NewLocationAddressWithDefaults instantiates a new LocationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAddressWithDefaults() *LocationAddress {
	this := LocationAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetAddress1() string {
	if o == nil || IsNil(o.Address1.Get()) {
		var ret string
		return ret
	}
	return *o.Address1.Get()
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address1.Get(), o.Address1.IsSet()
}

// HasAddress1 returns a boolean if a field has been set.
func (o *LocationAddress) HasAddress1() bool {
	if o != nil && o.Address1.IsSet() {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given NullableString and assigns it to the Address1 field.
func (o *LocationAddress) SetAddress1(v string) {
	o.Address1.Set(&v)
}

// SetAddress1Nil sets the value for Address1 to be an explicit nil
func (o *LocationAddress) SetAddress1Nil() {
	o.Address1.Set(nil)
}

// UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
func (o *LocationAddress) UnsetAddress1() {
	o.Address1.Unset()
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2.Get()) {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetAddress2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *LocationAddress) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *LocationAddress) SetAddress2(v string) {
	o.Address2.Set(&v)
}

// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *LocationAddress) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *LocationAddress) UnsetAddress2() {
	o.Address2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *LocationAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *LocationAddress) SetCity(v string) {
	o.City.Set(&v)
}

// SetCityNil sets the value for City to be an explicit nil
func (o *LocationAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *LocationAddress) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *LocationAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *LocationAddress) SetCountry(v string) {
	o.Country.Set(&v)
}

// SetCountryNil sets the value for Country to be an explicit nil
func (o *LocationAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *LocationAddress) UnsetCountry() {
	o.Country.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LocationAddress) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LocationAddress) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *LocationAddress) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LocationAddress) UnsetEmail() {
	o.Email.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LocationAddress) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LocationAddress) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *LocationAddress) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LocationAddress) UnsetName() {
	o.Name.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *LocationAddress) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *LocationAddress) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}

// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *LocationAddress) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *LocationAddress) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *LocationAddress) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *LocationAddress) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *LocationAddress) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *LocationAddress) UnsetState() {
	o.State.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationAddress) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationAddress) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *LocationAddress) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *LocationAddress) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}

// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *LocationAddress) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *LocationAddress) UnsetZipCode() {
	o.ZipCode.Unset()
}

func (o LocationAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address1.IsSet() {
		toSerialize["address1"] = o.Address1.Get()
	}
	if o.Address2.IsSet() {
		toSerialize["address2"] = o.Address2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zip_code"] = o.ZipCode.Get()
	}
	return toSerialize, nil
}

type NullableLocationAddress struct {
	value *LocationAddress
	isSet bool
}

func (v NullableLocationAddress) Get() *LocationAddress {
	return v.value
}

func (v *NullableLocationAddress) Set(val *LocationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAddress(val *LocationAddress) *NullableLocationAddress {
	return &NullableLocationAddress{value: val, isSet: true}
}

func (v NullableLocationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
