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

// ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel Information about a fulfillment center
type ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel struct {
	// Address line one of the fulfillment center
	Address1 NullableString `json:"address1,omitempty"`
	// Address line two of the fulfillment center
	Address2 NullableString `json:"address2,omitempty"`
	// City the fulfillment center is located in
	City NullableString `json:"city,omitempty"`
	// Country the fulfillment center is located in
	Country NullableString `json:"country,omitempty"`
	// Email contact for the fulfillment center
	Email NullableString `json:"email,omitempty"`
	// Unique identifier of the fulfillment center
	Id *int32 `json:"id,omitempty"`
	// Name of the fulfillment center
	Name NullableString `json:"name,omitempty"`
	// Phone number contact for the fulfillment center
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// State the fulfillment center is located in
	State NullableString `json:"state,omitempty"`
	// Timezone the fulfillment center is located in
	Timezone NullableString `json:"timezone,omitempty"`
	// Postal code of the fulfillment center
	ZipCode NullableString `json:"zip_code,omitempty"`
}

// NewShipbobReceivingPublicApiModelsFulfillmentCenterViewModel instantiates a new ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipbobReceivingPublicApiModelsFulfillmentCenterViewModel() *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel {
	this := ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel{}
	return &this
}

// NewShipbobReceivingPublicApiModelsFulfillmentCenterViewModelWithDefaults instantiates a new ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipbobReceivingPublicApiModelsFulfillmentCenterViewModelWithDefaults() *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel {
	this := ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetAddress1() string {
	if o == nil || o.Address1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address1.Get()
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetAddress1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address1.Get(), o.Address1.IsSet()
}

// HasAddress1 returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasAddress1() bool {
	if o != nil && o.Address1.IsSet() {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given NullableString and assigns it to the Address1 field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetAddress1(v string) {
	o.Address1.Set(&v)
}
// SetAddress1Nil sets the value for Address1 to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetAddress1Nil() {
	o.Address1.Set(nil)
}

// UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetAddress1() {
	o.Address1.Unset()
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetAddress2() string {
	if o == nil || o.Address2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetAddress2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetAddress2(v string) {
	o.Address2.Set(&v)
}
// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetAddress2() {
	o.Address2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetCountry() {
	o.Country.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetEmail() {
	o.Email.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetName() {
	o.Name.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetState() {
	o.State.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetTimezone() string {
	if o == nil || o.Timezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetZipCode() string {
	if o == nil || o.ZipCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) GetZipCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnsetZipCode() {
	o.ZipCode.Unset()
}

func (o ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) MarshalJSON() ([]byte, error) {
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zip_code"] = o.ZipCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel struct {
	value *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel
	isSet bool
}

func (v NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) Get() *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel {
	return v.value
}

func (v *NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) Set(val *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel(val *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) *NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel {
	return &NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel{value: val, isSet: true}
}

func (v NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


