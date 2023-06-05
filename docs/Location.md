# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **NullableString** | Abbreviation of the location. Combination of nearest Airport Code and the sequence number. | [optional] 
**AccessGranted** | Pointer to **bool** | Indicates whether or not the user is authorized to interact at all with the location | [optional] 
**Attributes** | Pointer to **[]string** | Available attributes for the location | [optional] 
**Id** | Pointer to **int32** | Id of the location in ShipBob&#39;s database | [optional] 
**IsActive** | Pointer to **bool** | Indicates if the location is operationally active or inactive | [optional] 
**IsReceivingEnabled** | Pointer to **bool** | Indicates if the receiving is enabled for FC | [optional] 
**IsShippingEnabled** | Pointer to **bool** | Indicates if the shipping is enabled for FC | [optional] 
**Name** | Pointer to **NullableString** | Name of the location. Follows the naming convention City (State Code) for domestic FCs and City (Country Code) for international FCs | [optional] 
**Region** | Pointer to [**FulfillmentCenterRegion**](FulfillmentCenterRegion.md) |  | [optional] 
**Services** | Pointer to [**[]LocationService**](LocationService.md) | Services provided by the location | [optional] 
**Timezone** | Pointer to **NullableString** | Time zone of the location | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *Location) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *Location) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *Location) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *Location) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *Location) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *Location) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetAccessGranted

`func (o *Location) GetAccessGranted() bool`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *Location) GetAccessGrantedOk() (*bool, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *Location) SetAccessGranted(v bool)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *Location) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAttributes

`func (o *Location) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Location) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Location) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Location) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *Location) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *Location) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetId

`func (o *Location) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *Location) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Location) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Location) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Location) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsReceivingEnabled

`func (o *Location) GetIsReceivingEnabled() bool`

GetIsReceivingEnabled returns the IsReceivingEnabled field if non-nil, zero value otherwise.

### GetIsReceivingEnabledOk

`func (o *Location) GetIsReceivingEnabledOk() (*bool, bool)`

GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceivingEnabled

`func (o *Location) SetIsReceivingEnabled(v bool)`

SetIsReceivingEnabled sets IsReceivingEnabled field to given value.

### HasIsReceivingEnabled

`func (o *Location) HasIsReceivingEnabled() bool`

HasIsReceivingEnabled returns a boolean if a field has been set.

### GetIsShippingEnabled

`func (o *Location) GetIsShippingEnabled() bool`

GetIsShippingEnabled returns the IsShippingEnabled field if non-nil, zero value otherwise.

### GetIsShippingEnabledOk

`func (o *Location) GetIsShippingEnabledOk() (*bool, bool)`

GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingEnabled

`func (o *Location) SetIsShippingEnabled(v bool)`

SetIsShippingEnabled sets IsShippingEnabled field to given value.

### HasIsShippingEnabled

`func (o *Location) HasIsShippingEnabled() bool`

HasIsShippingEnabled returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Location) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Location) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *Location) GetRegion() FulfillmentCenterRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Location) GetRegionOk() (*FulfillmentCenterRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Location) SetRegion(v FulfillmentCenterRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Location) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *Location) GetServices() []LocationService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Location) GetServicesOk() (*[]LocationService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Location) SetServices(v []LocationService)`

SetServices sets Services field to given value.

### HasServices

`func (o *Location) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *Location) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *Location) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetTimezone

`func (o *Location) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Location) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Location) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Location) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *Location) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *Location) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


