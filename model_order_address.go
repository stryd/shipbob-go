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

// checks if the OrderAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAddress{}

// OrderAddress struct for OrderAddress
type OrderAddress struct {
	// First line of the address
	Address1 string `json:"address1"`
	// Second line of the address
	Address2 *string `json:"address2,omitempty"`
	// The city
	City string `json:"city"`
	// Name of the company receiving the shipment
	CompanyName *string `json:"company_name,omitempty"`
	// The country (Must be ISO Alpha-2 for estimates)
	Country string `json:"country"`
	// The state or province
	State *string `json:"state,omitempty"`
	// The zip code or postal code
	ZipCode *string `json:"zip_code,omitempty"`
}

// NewOrderAddress instantiates a new OrderAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddress(address1 string, city string, country string) *OrderAddress {
	this := OrderAddress{}
	this.Address1 = address1
	this.City = city
	this.Country = country
	return &this
}

// NewOrderAddressWithDefaults instantiates a new OrderAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddressWithDefaults() *OrderAddress {
	this := OrderAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value
func (o *OrderAddress) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *OrderAddress) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *OrderAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *OrderAddress) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *OrderAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value
func (o *OrderAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *OrderAddress) SetCity(v string) {
	o.City = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderAddress) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderAddress) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderAddress) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value
func (o *OrderAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *OrderAddress) SetCountry(v string) {
	o.Country = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderAddress) SetState(v string) {
	o.State = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *OrderAddress) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *OrderAddress) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *OrderAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o OrderAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	toSerialize["city"] = o.City
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	toSerialize["country"] = o.Country
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	return toSerialize, nil
}

type NullableOrderAddress struct {
	value *OrderAddress
	isSet bool
}

func (v NullableOrderAddress) Get() *OrderAddress {
	return v.value
}

func (v *NullableOrderAddress) Set(val *OrderAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddress(val *OrderAddress) *NullableOrderAddress {
	return &NullableOrderAddress{value: val, isSet: true}
}

func (v NullableOrderAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
