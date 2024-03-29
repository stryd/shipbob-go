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

// checks if the InternalLocationV2AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalLocationV2AllOf{}

// InternalLocationV2AllOf struct for InternalLocationV2AllOf
type InternalLocationV2AllOf struct {
	// Abbreviation of the location. Combination of nearest Airport Code and the sequence number.
	Abbreviation NullableString `json:"abbreviation,omitempty"`
	// Indicates whether or not the user is authorized to interact at all with the location
	AccessGranted *bool `json:"access_granted,omitempty"`
	// Available attributes for the location
	Attributes                  []string                                           `json:"attributes,omitempty"`
	FulfillmentCenterAttributes []FcAttribute                                      `json:"fulfillment_center_attributes,omitempty"`
	FulfillmentCenterType       NullableInternalLocationAllOfFulfillmentCenterType `json:"fulfillment_center_type,omitempty"`
	// Id of the location in ShipBob's database
	Id *int32 `json:"id,omitempty"`
	// Indicates if the location is operationally active or inactive
	IsActive            *bool `json:"is_active,omitempty"`
	IsEnabledForNewUser *bool `json:"is_enabled_for_new_user,omitempty"`
	// Indicates if the receiving is enabled for FC
	IsReceivingEnabled *bool `json:"is_receiving_enabled,omitempty"`
	// Indicates if the shipping is enabled for FC
	IsShippingEnabled *bool `json:"is_shipping_enabled,omitempty"`
	// Name of the location. Follows the naming convention City (State Code) for domestic FCs and City (Country Code) for international FCs
	Name           NullableString           `json:"name,omitempty"`
	OrganizationId *string                  `json:"organization_id,omitempty"`
	OwnerId        NullableString           `json:"owner_id,omitempty"`
	ParentId       NullableString           `json:"parent_id,omitempty"`
	Region         *FulfillmentCenterRegion `json:"region,omitempty"`
	// Services provided by the location
	Services []LocationService `json:"services,omitempty"`
	// Time zone of the location
	Timezone             NullableString `json:"timezone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InternalLocationV2AllOf InternalLocationV2AllOf

// NewInternalLocationV2AllOf instantiates a new InternalLocationV2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalLocationV2AllOf() *InternalLocationV2AllOf {
	this := InternalLocationV2AllOf{}
	return &this
}

// NewInternalLocationV2AllOfWithDefaults instantiates a new InternalLocationV2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalLocationV2AllOfWithDefaults() *InternalLocationV2AllOf {
	this := InternalLocationV2AllOf{}
	return &this
}

// GetAbbreviation returns the Abbreviation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetAbbreviation() string {
	if o == nil || IsNil(o.Abbreviation.Get()) {
		var ret string
		return ret
	}
	return *o.Abbreviation.Get()
}

// GetAbbreviationOk returns a tuple with the Abbreviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetAbbreviationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Abbreviation.Get(), o.Abbreviation.IsSet()
}

// HasAbbreviation returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasAbbreviation() bool {
	if o != nil && o.Abbreviation.IsSet() {
		return true
	}

	return false
}

// SetAbbreviation gets a reference to the given NullableString and assigns it to the Abbreviation field.
func (o *InternalLocationV2AllOf) SetAbbreviation(v string) {
	o.Abbreviation.Set(&v)
}

// SetAbbreviationNil sets the value for Abbreviation to be an explicit nil
func (o *InternalLocationV2AllOf) SetAbbreviationNil() {
	o.Abbreviation.Set(nil)
}

// UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetAbbreviation() {
	o.Abbreviation.Unset()
}

// GetAccessGranted returns the AccessGranted field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetAccessGranted() bool {
	if o == nil || IsNil(o.AccessGranted) {
		var ret bool
		return ret
	}
	return *o.AccessGranted
}

// GetAccessGrantedOk returns a tuple with the AccessGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetAccessGrantedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessGranted) {
		return nil, false
	}
	return o.AccessGranted, true
}

// HasAccessGranted returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasAccessGranted() bool {
	if o != nil && !IsNil(o.AccessGranted) {
		return true
	}

	return false
}

// SetAccessGranted gets a reference to the given bool and assigns it to the AccessGranted field.
func (o *InternalLocationV2AllOf) SetAccessGranted(v bool) {
	o.AccessGranted = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *InternalLocationV2AllOf) SetAttributes(v []string) {
	o.Attributes = v
}

// GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetFulfillmentCenterAttributes() []FcAttribute {
	if o == nil {
		var ret []FcAttribute
		return ret
	}
	return o.FulfillmentCenterAttributes
}

// GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetFulfillmentCenterAttributesOk() ([]FcAttribute, bool) {
	if o == nil || IsNil(o.FulfillmentCenterAttributes) {
		return nil, false
	}
	return o.FulfillmentCenterAttributes, true
}

// HasFulfillmentCenterAttributes returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasFulfillmentCenterAttributes() bool {
	if o != nil && IsNil(o.FulfillmentCenterAttributes) {
		return true
	}

	return false
}

// SetFulfillmentCenterAttributes gets a reference to the given []FcAttribute and assigns it to the FulfillmentCenterAttributes field.
func (o *InternalLocationV2AllOf) SetFulfillmentCenterAttributes(v []FcAttribute) {
	o.FulfillmentCenterAttributes = v
}

// GetFulfillmentCenterType returns the FulfillmentCenterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType {
	if o == nil || IsNil(o.FulfillmentCenterType.Get()) {
		var ret InternalLocationAllOfFulfillmentCenterType
		return ret
	}
	return *o.FulfillmentCenterType.Get()
}

// GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfillmentCenterType.Get(), o.FulfillmentCenterType.IsSet()
}

// HasFulfillmentCenterType returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasFulfillmentCenterType() bool {
	if o != nil && o.FulfillmentCenterType.IsSet() {
		return true
	}

	return false
}

// SetFulfillmentCenterType gets a reference to the given NullableInternalLocationAllOfFulfillmentCenterType and assigns it to the FulfillmentCenterType field.
func (o *InternalLocationV2AllOf) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType) {
	o.FulfillmentCenterType.Set(&v)
}

// SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil
func (o *InternalLocationV2AllOf) SetFulfillmentCenterTypeNil() {
	o.FulfillmentCenterType.Set(nil)
}

// UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetFulfillmentCenterType() {
	o.FulfillmentCenterType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InternalLocationV2AllOf) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *InternalLocationV2AllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsEnabledForNewUser returns the IsEnabledForNewUser field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetIsEnabledForNewUser() bool {
	if o == nil || IsNil(o.IsEnabledForNewUser) {
		var ret bool
		return ret
	}
	return *o.IsEnabledForNewUser
}

// GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetIsEnabledForNewUserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabledForNewUser) {
		return nil, false
	}
	return o.IsEnabledForNewUser, true
}

// HasIsEnabledForNewUser returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasIsEnabledForNewUser() bool {
	if o != nil && !IsNil(o.IsEnabledForNewUser) {
		return true
	}

	return false
}

// SetIsEnabledForNewUser gets a reference to the given bool and assigns it to the IsEnabledForNewUser field.
func (o *InternalLocationV2AllOf) SetIsEnabledForNewUser(v bool) {
	o.IsEnabledForNewUser = &v
}

// GetIsReceivingEnabled returns the IsReceivingEnabled field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetIsReceivingEnabled() bool {
	if o == nil || IsNil(o.IsReceivingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsReceivingEnabled
}

// GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetIsReceivingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReceivingEnabled) {
		return nil, false
	}
	return o.IsReceivingEnabled, true
}

// HasIsReceivingEnabled returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasIsReceivingEnabled() bool {
	if o != nil && !IsNil(o.IsReceivingEnabled) {
		return true
	}

	return false
}

// SetIsReceivingEnabled gets a reference to the given bool and assigns it to the IsReceivingEnabled field.
func (o *InternalLocationV2AllOf) SetIsReceivingEnabled(v bool) {
	o.IsReceivingEnabled = &v
}

// GetIsShippingEnabled returns the IsShippingEnabled field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetIsShippingEnabled() bool {
	if o == nil || IsNil(o.IsShippingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsShippingEnabled
}

// GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetIsShippingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShippingEnabled) {
		return nil, false
	}
	return o.IsShippingEnabled, true
}

// HasIsShippingEnabled returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasIsShippingEnabled() bool {
	if o != nil && !IsNil(o.IsShippingEnabled) {
		return true
	}

	return false
}

// SetIsShippingEnabled gets a reference to the given bool and assigns it to the IsShippingEnabled field.
func (o *InternalLocationV2AllOf) SetIsShippingEnabled(v bool) {
	o.IsShippingEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InternalLocationV2AllOf) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *InternalLocationV2AllOf) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InternalLocationV2AllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *InternalLocationV2AllOf) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}

// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *InternalLocationV2AllOf) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *InternalLocationV2AllOf) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *InternalLocationV2AllOf) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetParentId() {
	o.ParentId.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *InternalLocationV2AllOf) GetRegion() FulfillmentCenterRegion {
	if o == nil || IsNil(o.Region) {
		var ret FulfillmentCenterRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalLocationV2AllOf) GetRegionOk() (*FulfillmentCenterRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given FulfillmentCenterRegion and assigns it to the Region field.
func (o *InternalLocationV2AllOf) SetRegion(v FulfillmentCenterRegion) {
	o.Region = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetServices() []LocationService {
	if o == nil {
		var ret []LocationService
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetServicesOk() ([]LocationService, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasServices() bool {
	if o != nil && IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []LocationService and assigns it to the Services field.
func (o *InternalLocationV2AllOf) SetServices(v []LocationService) {
	o.Services = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InternalLocationV2AllOf) GetTimezone() string {
	if o == nil || IsNil(o.Timezone.Get()) {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InternalLocationV2AllOf) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *InternalLocationV2AllOf) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *InternalLocationV2AllOf) SetTimezone(v string) {
	o.Timezone.Set(&v)
}

// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *InternalLocationV2AllOf) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *InternalLocationV2AllOf) UnsetTimezone() {
	o.Timezone.Unset()
}

func (o InternalLocationV2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalLocationV2AllOf) ToMap() (map[string]interface{}, error) {
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
	if o.FulfillmentCenterAttributes != nil {
		toSerialize["fulfillment_center_attributes"] = o.FulfillmentCenterAttributes
	}
	if o.FulfillmentCenterType.IsSet() {
		toSerialize["fulfillment_center_type"] = o.FulfillmentCenterType.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsEnabledForNewUser) {
		toSerialize["is_enabled_for_new_user"] = o.IsEnabledForNewUser
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
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.OwnerId.IsSet() {
		toSerialize["owner_id"] = o.OwnerId.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InternalLocationV2AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInternalLocationV2AllOf := _InternalLocationV2AllOf{}

	if err = json.Unmarshal(bytes, &varInternalLocationV2AllOf); err == nil {
		*o = InternalLocationV2AllOf(varInternalLocationV2AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "abbreviation")
		delete(additionalProperties, "access_granted")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "fulfillment_center_attributes")
		delete(additionalProperties, "fulfillment_center_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_enabled_for_new_user")
		delete(additionalProperties, "is_receiving_enabled")
		delete(additionalProperties, "is_shipping_enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "region")
		delete(additionalProperties, "services")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalLocationV2AllOf struct {
	value *InternalLocationV2AllOf
	isSet bool
}

func (v NullableInternalLocationV2AllOf) Get() *InternalLocationV2AllOf {
	return v.value
}

func (v *NullableInternalLocationV2AllOf) Set(val *InternalLocationV2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalLocationV2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalLocationV2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalLocationV2AllOf(val *InternalLocationV2AllOf) *NullableInternalLocationV2AllOf {
	return &NullableInternalLocationV2AllOf{value: val, isSet: true}
}

func (v NullableInternalLocationV2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalLocationV2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
