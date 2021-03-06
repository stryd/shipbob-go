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

// EstimationAddress struct for EstimationAddress
type EstimationAddress struct {
	// First line of the address
	Address1 *string `json:"address1,omitempty"`
	// Second line of the address
	Address2 *string `json:"address2,omitempty"`
	// The city
	City *string `json:"city,omitempty"`
	// Name of the company receiving the shipment
	CompanyName *string `json:"company_name,omitempty"`
	// The country (Must be ISO Alpha-2 for estimates)
	Country string `json:"country"`
	// The state or province
	State *string `json:"state,omitempty"`
	// The zip code or postal code
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewEstimationAddress instantiates a new EstimationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimationAddress(country string) *EstimationAddress {
	this := EstimationAddress{}
	this.Country = country
	return &this
}

// NewEstimationAddressWithDefaults instantiates a new EstimationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimationAddressWithDefaults() *EstimationAddress {
	this := EstimationAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *EstimationAddress) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *EstimationAddress) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *EstimationAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *EstimationAddress) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *EstimationAddress) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *EstimationAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *EstimationAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *EstimationAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *EstimationAddress) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *EstimationAddress) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *EstimationAddress) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *EstimationAddress) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value
func (o *EstimationAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *EstimationAddress) SetCountry(v string) {
	o.Country = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EstimationAddress) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EstimationAddress) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EstimationAddress) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *EstimationAddress) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimationAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *EstimationAddress) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *EstimationAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o EstimationAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address1 != nil {
		toSerialize["address1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["address2"] = o.Address2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["company_name"] = o.CompanyName
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	return json.Marshal(toSerialize)
}

type NullableEstimationAddress struct {
	value *EstimationAddress
	isSet bool
}

func (v NullableEstimationAddress) Get() *EstimationAddress {
	return v.value
}

func (v *NullableEstimationAddress) Set(val *EstimationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimationAddress(val *EstimationAddress) *NullableEstimationAddress {
	return &NullableEstimationAddress{value: val, isSet: true}
}

func (v NullableEstimationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
