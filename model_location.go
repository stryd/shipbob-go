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

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location struct for Location
type Location struct {
	// Abbreviation of the location. Combination of nearest Airport Code and the sequence number.
	Abbreviation NullableString `json:"abbreviation,omitempty"`
	// Indicates whether or not the user is authorized to interact at all with the location
	AccessGranted *bool `json:"access_granted,omitempty"`
	// Available attributes for the location
	Attributes []string `json:"attributes,omitempty"`
	// Id of the location in ShipBobâs database
	Id *int32 `json:"id,omitempty"`
	// Indicates if the location is operationally active or inactive
	IsActive *bool `json:"is_active,omitempty"`
	// Indicates if the receiving is enabled for FC
	IsReceivingEnabled *bool `json:"is_receiving_enabled,omitempty"`
	// Indicates if the shipping is enabled for FC
	IsShippingEnabled *bool `json:"is_shipping_enabled,omitempty"`
	// Name of the location. Follows the naming convention City (State Code)\\r\\nfor domestic FCs and City (Country Code) for international FCs
	Name NullableString `json:"name,omitempty"`
	Region *LocationRegion `json:"region,omitempty"`
	// Services provided by the location
	Services []LocationService `json:"services,omitempty"`
	// Time zone of the location
	Timezone NullableString `json:"timezone,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetAbbreviation returns the Abbreviation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetAbbreviation() string {
	if o == nil || IsNil(o.Abbreviation.Get()) {
		var ret string
		return ret
	}
	return *o.Abbreviation.Get()
}

// GetAbbreviationOk returns a tuple with the Abbreviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetAbbreviationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Abbreviation.Get(), o.Abbreviation.IsSet()
}

// HasAbbreviation returns a boolean if a field has been set.
func (o *Location) HasAbbreviation() bool {
	if o != nil && o.Abbreviation.IsSet() {
		return true
	}

	return false
}

// SetAbbreviation gets a reference to the given NullableString and assigns it to the Abbreviation field.
func (o *Location) SetAbbreviation(v string) {
	o.Abbreviation.Set(&v)
}
// SetAbbreviationNil sets the value for Abbreviation to be an explicit nil
func (o *Location) SetAbbreviationNil() {
	o.Abbreviation.Set(nil)
}

// UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
func (o *Location) UnsetAbbreviation() {
	o.Abbreviation.Unset()
}

// GetAccessGranted returns the AccessGranted field value if set, zero value otherwise.
func (o *Location) GetAccessGranted() bool {
	if o == nil || IsNil(o.AccessGranted) {
		var ret bool
		return ret
	}
	return *o.AccessGranted
}

// GetAccessGrantedOk returns a tuple with the AccessGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAccessGrantedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessGranted) {
		return nil, false
	}
	return o.AccessGranted, true
}

// HasAccessGranted returns a boolean if a field has been set.
func (o *Location) HasAccessGranted() bool {
	if o != nil && !IsNil(o.AccessGranted) {
		return true
	}

	return false
}

// SetAccessGranted gets a reference to the given bool and assigns it to the AccessGranted field.
func (o *Location) SetAccessGranted(v bool) {
	o.AccessGranted = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Location) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *Location) SetAttributes(v []string) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Location) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Location) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Location) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Location) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Location) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Location) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsReceivingEnabled returns the IsReceivingEnabled field value if set, zero value otherwise.
func (o *Location) GetIsReceivingEnabled() bool {
	if o == nil || IsNil(o.IsReceivingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsReceivingEnabled
}

// GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIsReceivingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReceivingEnabled) {
		return nil, false
	}
	return o.IsReceivingEnabled, true
}

// HasIsReceivingEnabled returns a boolean if a field has been set.
func (o *Location) HasIsReceivingEnabled() bool {
	if o != nil && !IsNil(o.IsReceivingEnabled) {
		return true
	}

	return false
}

// SetIsReceivingEnabled gets a reference to the given bool and assigns it to the IsReceivingEnabled field.
func (o *Location) SetIsReceivingEnabled(v bool) {
	o.IsReceivingEnabled = &v
}

// GetIsShippingEnabled returns the IsShippingEnabled field value if set, zero value otherwise.
func (o *Location) GetIsShippingEnabled() bool {
	if o == nil || IsNil(o.IsShippingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsShippingEnabled
}

// GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIsShippingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShippingEnabled) {
		return nil, false
	}
	return o.IsShippingEnabled, true
}

// HasIsShippingEnabled returns a boolean if a field has been set.
func (o *Location) HasIsShippingEnabled() bool {
	if o != nil && !IsNil(o.IsShippingEnabled) {
		return true
	}

	return false
}

// SetIsShippingEnabled gets a reference to the given bool and assigns it to the IsShippingEnabled field.
func (o *Location) SetIsShippingEnabled(v bool) {
	o.IsShippingEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Location) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Location) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Location) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Location) UnsetName() {
	o.Name.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Location) GetRegion() LocationRegion {
	if o == nil || IsNil(o.Region) {
		var ret LocationRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetRegionOk() (*LocationRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Location) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given LocationRegion and assigns it to the Region field.
func (o *Location) SetRegion(v LocationRegion) {
	o.Region = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetServices() []LocationService {
	if o == nil {
		var ret []LocationService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetServicesOk() ([]LocationService, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Location) HasServices() bool {
	if o != nil && IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []LocationService and assigns it to the Services field.
func (o *Location) SetServices(v []LocationService) {
	o.Services = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetTimezone() string {
	if o == nil || IsNil(o.Timezone.Get()) {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *Location) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *Location) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *Location) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *Location) UnsetTimezone() {
	o.Timezone.Unset()
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Abbreviation.IsSet() {
		toSerialize["abbreviation"] = o.Abbreviation.Get()
	}
	if !IsNil(o.AccessGranted) {
		toSerialize["access_granted"] = o.AccessGranted
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsReceivingEnabled) {
		toSerialize["is_receiving_enabled"] = o.IsReceivingEnabled
	}
	if !IsNil(o.IsShippingEnabled) {
		toSerialize["is_shipping_enabled"] = o.IsShippingEnabled
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	return toSerialize, nil
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


