# InternalLocationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **NullableString** | Abbreviation of the location. Combination of nearest Airport Code and the sequence number. | [optional] 
**AccessGranted** | Pointer to **bool** | Indicates whether or not the user is authorized to interact at all with the location | [optional] 
**Attributes** | Pointer to **[]string** | Available attributes for the location | [optional] 
**FulfillmentCenterAttributes** | Pointer to [**[]FcAttribute**](FcAttribute.md) |  | [optional] 
**FulfillmentCenterType** | Pointer to [**NullableInternalLocationAllOfFulfillmentCenterType**](InternalLocationAllOfFulfillmentCenterType.md) |  | [optional] 
**Id** | Pointer to **int32** | Id of the location in ShipBob&#39;s database | [optional] 
**IsActive** | Pointer to **bool** | Indicates if the location is operationally active or inactive | [optional] 
**IsEnabledForNewUser** | Pointer to **bool** |  | [optional] 
**IsReceivingEnabled** | Pointer to **bool** | Indicates if the receiving is enabled for FC | [optional] 
**IsShippingEnabled** | Pointer to **bool** | Indicates if the shipping is enabled for FC | [optional] 
**Name** | Pointer to **NullableString** | Name of the location. Follows the naming convention City (State Code) for domestic FCs and City (Country Code) for international FCs | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**FulfillmentCenterRegion**](FulfillmentCenterRegion.md) |  | [optional] 
**Services** | Pointer to [**[]LocationService**](LocationService.md) | Services provided by the location | [optional] 
**Timezone** | Pointer to **NullableString** | Time zone of the location | [optional] 

## Methods

### NewInternalLocationV2

`func NewInternalLocationV2() *InternalLocationV2`

NewInternalLocationV2 instantiates a new InternalLocationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalLocationV2WithDefaults

`func NewInternalLocationV2WithDefaults() *InternalLocationV2`

NewInternalLocationV2WithDefaults instantiates a new InternalLocationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *InternalLocationV2) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *InternalLocationV2) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *InternalLocationV2) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *InternalLocationV2) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *InternalLocationV2) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *InternalLocationV2) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetAccessGranted

`func (o *InternalLocationV2) GetAccessGranted() bool`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *InternalLocationV2) GetAccessGrantedOk() (*bool, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *InternalLocationV2) SetAccessGranted(v bool)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *InternalLocationV2) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAttributes

`func (o *InternalLocationV2) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InternalLocationV2) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InternalLocationV2) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InternalLocationV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *InternalLocationV2) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *InternalLocationV2) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetFulfillmentCenterAttributes

`func (o *InternalLocationV2) GetFulfillmentCenterAttributes() []FcAttribute`

GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field if non-nil, zero value otherwise.

### GetFulfillmentCenterAttributesOk

`func (o *InternalLocationV2) GetFulfillmentCenterAttributesOk() (*[]FcAttribute, bool)`

GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterAttributes

`func (o *InternalLocationV2) SetFulfillmentCenterAttributes(v []FcAttribute)`

SetFulfillmentCenterAttributes sets FulfillmentCenterAttributes field to given value.

### HasFulfillmentCenterAttributes

`func (o *InternalLocationV2) HasFulfillmentCenterAttributes() bool`

HasFulfillmentCenterAttributes returns a boolean if a field has been set.

### SetFulfillmentCenterAttributesNil

`func (o *InternalLocationV2) SetFulfillmentCenterAttributesNil(b bool)`

 SetFulfillmentCenterAttributesNil sets the value for FulfillmentCenterAttributes to be an explicit nil

### UnsetFulfillmentCenterAttributes
`func (o *InternalLocationV2) UnsetFulfillmentCenterAttributes()`

UnsetFulfillmentCenterAttributes ensures that no value is present for FulfillmentCenterAttributes, not even an explicit nil
### GetFulfillmentCenterType

`func (o *InternalLocationV2) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType`

GetFulfillmentCenterType returns the FulfillmentCenterType field if non-nil, zero value otherwise.

### GetFulfillmentCenterTypeOk

`func (o *InternalLocationV2) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool)`

GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterType

`func (o *InternalLocationV2) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType)`

SetFulfillmentCenterType sets FulfillmentCenterType field to given value.

### HasFulfillmentCenterType

`func (o *InternalLocationV2) HasFulfillmentCenterType() bool`

HasFulfillmentCenterType returns a boolean if a field has been set.

### SetFulfillmentCenterTypeNil

`func (o *InternalLocationV2) SetFulfillmentCenterTypeNil(b bool)`

 SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil

### UnsetFulfillmentCenterType
`func (o *InternalLocationV2) UnsetFulfillmentCenterType()`

UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
### GetId

`func (o *InternalLocationV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalLocationV2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalLocationV2) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalLocationV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *InternalLocationV2) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *InternalLocationV2) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *InternalLocationV2) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *InternalLocationV2) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsEnabledForNewUser

`func (o *InternalLocationV2) GetIsEnabledForNewUser() bool`

GetIsEnabledForNewUser returns the IsEnabledForNewUser field if non-nil, zero value otherwise.

### GetIsEnabledForNewUserOk

`func (o *InternalLocationV2) GetIsEnabledForNewUserOk() (*bool, bool)`

GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForNewUser

`func (o *InternalLocationV2) SetIsEnabledForNewUser(v bool)`

SetIsEnabledForNewUser sets IsEnabledForNewUser field to given value.

### HasIsEnabledForNewUser

`func (o *InternalLocationV2) HasIsEnabledForNewUser() bool`

HasIsEnabledForNewUser returns a boolean if a field has been set.

### GetIsReceivingEnabled

`func (o *InternalLocationV2) GetIsReceivingEnabled() bool`

GetIsReceivingEnabled returns the IsReceivingEnabled field if non-nil, zero value otherwise.

### GetIsReceivingEnabledOk

`func (o *InternalLocationV2) GetIsReceivingEnabledOk() (*bool, bool)`

GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceivingEnabled

`func (o *InternalLocationV2) SetIsReceivingEnabled(v bool)`

SetIsReceivingEnabled sets IsReceivingEnabled field to given value.

### HasIsReceivingEnabled

`func (o *InternalLocationV2) HasIsReceivingEnabled() bool`

HasIsReceivingEnabled returns a boolean if a field has been set.

### GetIsShippingEnabled

`func (o *InternalLocationV2) GetIsShippingEnabled() bool`

GetIsShippingEnabled returns the IsShippingEnabled field if non-nil, zero value otherwise.

### GetIsShippingEnabledOk

`func (o *InternalLocationV2) GetIsShippingEnabledOk() (*bool, bool)`

GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingEnabled

`func (o *InternalLocationV2) SetIsShippingEnabled(v bool)`

SetIsShippingEnabled sets IsShippingEnabled field to given value.

### HasIsShippingEnabled

`func (o *InternalLocationV2) HasIsShippingEnabled() bool`

HasIsShippingEnabled returns a boolean if a field has been set.

### GetName

`func (o *InternalLocationV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalLocationV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalLocationV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalLocationV2) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InternalLocationV2) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternalLocationV2) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *InternalLocationV2) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InternalLocationV2) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InternalLocationV2) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InternalLocationV2) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOwnerId

`func (o *InternalLocationV2) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InternalLocationV2) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InternalLocationV2) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *InternalLocationV2) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *InternalLocationV2) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *InternalLocationV2) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetParentId

`func (o *InternalLocationV2) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *InternalLocationV2) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *InternalLocationV2) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *InternalLocationV2) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *InternalLocationV2) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *InternalLocationV2) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRegion

`func (o *InternalLocationV2) GetRegion() FulfillmentCenterRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InternalLocationV2) GetRegionOk() (*FulfillmentCenterRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InternalLocationV2) SetRegion(v FulfillmentCenterRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InternalLocationV2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *InternalLocationV2) GetServices() []LocationService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InternalLocationV2) GetServicesOk() (*[]LocationService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InternalLocationV2) SetServices(v []LocationService)`

SetServices sets Services field to given value.

### HasServices

`func (o *InternalLocationV2) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *InternalLocationV2) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *InternalLocationV2) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetTimezone

`func (o *InternalLocationV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InternalLocationV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InternalLocationV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InternalLocationV2) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *InternalLocationV2) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *InternalLocationV2) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


