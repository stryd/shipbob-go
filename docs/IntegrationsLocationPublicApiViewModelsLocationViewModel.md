# IntegrationsLocationPublicApiViewModelsLocationViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **NullableString** | Abbreviation of the location. Combination of nearest Airport Code and the sequence number. | [optional] 
**AccessGranted** | Pointer to **bool** | Indicates whether or not the user is authorized to interact at all with the location | [optional] 
**Attributes** | Pointer to **[]string** | Available attributes for the location | [optional] 
**Id** | Pointer to **int32** | Id of the location in ShipBobâs database | [optional] 
**IsActive** | Pointer to **bool** | Indicates if the location is operationally active or inactive | [optional] 
**IsReceivingEnabled** | Pointer to **bool** | Indicates if the receiving is enabled for FC | [optional] 
**IsShippingEnabled** | Pointer to **bool** | Indicates if the shipping is enabled for FC | [optional] 
**Name** | Pointer to **NullableString** | Name of the location. Follows the naming convention City (State Code)\\r\\nfor domestic FCs and City (Country Code) for international FCs | [optional] 
**Region** | Pointer to [**IntegrationsLocationPublicApiViewModelsRegionViewModel**](Integrations.Location.Public.Api.ViewModels.RegionViewModel.md) |  | [optional] 
**Services** | Pointer to [**[]IntegrationsLocationPublicApiViewModelsServiceViewModel**](IntegrationsLocationPublicApiViewModelsServiceViewModel.md) | Services provided by the location | [optional] 
**Timezone** | Pointer to **NullableString** | Time zone of the location | [optional] 

## Methods

### NewIntegrationsLocationPublicApiViewModelsLocationViewModel

`func NewIntegrationsLocationPublicApiViewModelsLocationViewModel() *IntegrationsLocationPublicApiViewModelsLocationViewModel`

NewIntegrationsLocationPublicApiViewModelsLocationViewModel instantiates a new IntegrationsLocationPublicApiViewModelsLocationViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsLocationPublicApiViewModelsLocationViewModelWithDefaults

`func NewIntegrationsLocationPublicApiViewModelsLocationViewModelWithDefaults() *IntegrationsLocationPublicApiViewModelsLocationViewModel`

NewIntegrationsLocationPublicApiViewModelsLocationViewModelWithDefaults instantiates a new IntegrationsLocationPublicApiViewModelsLocationViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetAccessGranted

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAccessGranted() bool`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAccessGrantedOk() (*bool, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetAccessGranted(v bool)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAttributes

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetId

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsReceivingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsReceivingEnabled() bool`

GetIsReceivingEnabled returns the IsReceivingEnabled field if non-nil, zero value otherwise.

### GetIsReceivingEnabledOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsReceivingEnabledOk() (*bool, bool)`

GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceivingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetIsReceivingEnabled(v bool)`

SetIsReceivingEnabled sets IsReceivingEnabled field to given value.

### HasIsReceivingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasIsReceivingEnabled() bool`

HasIsReceivingEnabled returns a boolean if a field has been set.

### GetIsShippingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsShippingEnabled() bool`

GetIsShippingEnabled returns the IsShippingEnabled field if non-nil, zero value otherwise.

### GetIsShippingEnabledOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetIsShippingEnabledOk() (*bool, bool)`

GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetIsShippingEnabled(v bool)`

SetIsShippingEnabled sets IsShippingEnabled field to given value.

### HasIsShippingEnabled

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasIsShippingEnabled() bool`

HasIsShippingEnabled returns a boolean if a field has been set.

### GetName

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetRegion() IntegrationsLocationPublicApiViewModelsRegionViewModel`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetRegionOk() (*IntegrationsLocationPublicApiViewModelsRegionViewModel, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetRegion(v IntegrationsLocationPublicApiViewModelsRegionViewModel)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetServices() []IntegrationsLocationPublicApiViewModelsServiceViewModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetServicesOk() (*[]IntegrationsLocationPublicApiViewModelsServiceViewModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetServices(v []IntegrationsLocationPublicApiViewModelsServiceViewModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetTimezone

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *IntegrationsLocationPublicApiViewModelsLocationViewModel) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


