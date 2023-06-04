# InternalLocationV2AllOf

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

### NewInternalLocationV2AllOf

`func NewInternalLocationV2AllOf() *InternalLocationV2AllOf`

NewInternalLocationV2AllOf instantiates a new InternalLocationV2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalLocationV2AllOfWithDefaults

`func NewInternalLocationV2AllOfWithDefaults() *InternalLocationV2AllOf`

NewInternalLocationV2AllOfWithDefaults instantiates a new InternalLocationV2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *InternalLocationV2AllOf) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *InternalLocationV2AllOf) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *InternalLocationV2AllOf) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *InternalLocationV2AllOf) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *InternalLocationV2AllOf) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *InternalLocationV2AllOf) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetAccessGranted

`func (o *InternalLocationV2AllOf) GetAccessGranted() bool`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *InternalLocationV2AllOf) GetAccessGrantedOk() (*bool, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *InternalLocationV2AllOf) SetAccessGranted(v bool)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *InternalLocationV2AllOf) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAttributes

`func (o *InternalLocationV2AllOf) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InternalLocationV2AllOf) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InternalLocationV2AllOf) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InternalLocationV2AllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *InternalLocationV2AllOf) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *InternalLocationV2AllOf) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetFulfillmentCenterAttributes

`func (o *InternalLocationV2AllOf) GetFulfillmentCenterAttributes() []FcAttribute`

GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field if non-nil, zero value otherwise.

### GetFulfillmentCenterAttributesOk

`func (o *InternalLocationV2AllOf) GetFulfillmentCenterAttributesOk() (*[]FcAttribute, bool)`

GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterAttributes

`func (o *InternalLocationV2AllOf) SetFulfillmentCenterAttributes(v []FcAttribute)`

SetFulfillmentCenterAttributes sets FulfillmentCenterAttributes field to given value.

### HasFulfillmentCenterAttributes

`func (o *InternalLocationV2AllOf) HasFulfillmentCenterAttributes() bool`

HasFulfillmentCenterAttributes returns a boolean if a field has been set.

### SetFulfillmentCenterAttributesNil

`func (o *InternalLocationV2AllOf) SetFulfillmentCenterAttributesNil(b bool)`

 SetFulfillmentCenterAttributesNil sets the value for FulfillmentCenterAttributes to be an explicit nil

### UnsetFulfillmentCenterAttributes
`func (o *InternalLocationV2AllOf) UnsetFulfillmentCenterAttributes()`

UnsetFulfillmentCenterAttributes ensures that no value is present for FulfillmentCenterAttributes, not even an explicit nil
### GetFulfillmentCenterType

`func (o *InternalLocationV2AllOf) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType`

GetFulfillmentCenterType returns the FulfillmentCenterType field if non-nil, zero value otherwise.

### GetFulfillmentCenterTypeOk

`func (o *InternalLocationV2AllOf) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool)`

GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterType

`func (o *InternalLocationV2AllOf) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType)`

SetFulfillmentCenterType sets FulfillmentCenterType field to given value.

### HasFulfillmentCenterType

`func (o *InternalLocationV2AllOf) HasFulfillmentCenterType() bool`

HasFulfillmentCenterType returns a boolean if a field has been set.

### SetFulfillmentCenterTypeNil

`func (o *InternalLocationV2AllOf) SetFulfillmentCenterTypeNil(b bool)`

 SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil

### UnsetFulfillmentCenterType
`func (o *InternalLocationV2AllOf) UnsetFulfillmentCenterType()`

UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
### GetId

`func (o *InternalLocationV2AllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalLocationV2AllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalLocationV2AllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalLocationV2AllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *InternalLocationV2AllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *InternalLocationV2AllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *InternalLocationV2AllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *InternalLocationV2AllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsEnabledForNewUser

`func (o *InternalLocationV2AllOf) GetIsEnabledForNewUser() bool`

GetIsEnabledForNewUser returns the IsEnabledForNewUser field if non-nil, zero value otherwise.

### GetIsEnabledForNewUserOk

`func (o *InternalLocationV2AllOf) GetIsEnabledForNewUserOk() (*bool, bool)`

GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForNewUser

`func (o *InternalLocationV2AllOf) SetIsEnabledForNewUser(v bool)`

SetIsEnabledForNewUser sets IsEnabledForNewUser field to given value.

### HasIsEnabledForNewUser

`func (o *InternalLocationV2AllOf) HasIsEnabledForNewUser() bool`

HasIsEnabledForNewUser returns a boolean if a field has been set.

### GetIsReceivingEnabled

`func (o *InternalLocationV2AllOf) GetIsReceivingEnabled() bool`

GetIsReceivingEnabled returns the IsReceivingEnabled field if non-nil, zero value otherwise.

### GetIsReceivingEnabledOk

`func (o *InternalLocationV2AllOf) GetIsReceivingEnabledOk() (*bool, bool)`

GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceivingEnabled

`func (o *InternalLocationV2AllOf) SetIsReceivingEnabled(v bool)`

SetIsReceivingEnabled sets IsReceivingEnabled field to given value.

### HasIsReceivingEnabled

`func (o *InternalLocationV2AllOf) HasIsReceivingEnabled() bool`

HasIsReceivingEnabled returns a boolean if a field has been set.

### GetIsShippingEnabled

`func (o *InternalLocationV2AllOf) GetIsShippingEnabled() bool`

GetIsShippingEnabled returns the IsShippingEnabled field if non-nil, zero value otherwise.

### GetIsShippingEnabledOk

`func (o *InternalLocationV2AllOf) GetIsShippingEnabledOk() (*bool, bool)`

GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingEnabled

`func (o *InternalLocationV2AllOf) SetIsShippingEnabled(v bool)`

SetIsShippingEnabled sets IsShippingEnabled field to given value.

### HasIsShippingEnabled

`func (o *InternalLocationV2AllOf) HasIsShippingEnabled() bool`

HasIsShippingEnabled returns a boolean if a field has been set.

### GetName

`func (o *InternalLocationV2AllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalLocationV2AllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalLocationV2AllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalLocationV2AllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InternalLocationV2AllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternalLocationV2AllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *InternalLocationV2AllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InternalLocationV2AllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InternalLocationV2AllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InternalLocationV2AllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOwnerId

`func (o *InternalLocationV2AllOf) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InternalLocationV2AllOf) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InternalLocationV2AllOf) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *InternalLocationV2AllOf) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *InternalLocationV2AllOf) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *InternalLocationV2AllOf) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetParentId

`func (o *InternalLocationV2AllOf) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *InternalLocationV2AllOf) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *InternalLocationV2AllOf) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *InternalLocationV2AllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *InternalLocationV2AllOf) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *InternalLocationV2AllOf) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRegion

`func (o *InternalLocationV2AllOf) GetRegion() FulfillmentCenterRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InternalLocationV2AllOf) GetRegionOk() (*FulfillmentCenterRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InternalLocationV2AllOf) SetRegion(v FulfillmentCenterRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InternalLocationV2AllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *InternalLocationV2AllOf) GetServices() []LocationService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InternalLocationV2AllOf) GetServicesOk() (*[]LocationService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InternalLocationV2AllOf) SetServices(v []LocationService)`

SetServices sets Services field to given value.

### HasServices

`func (o *InternalLocationV2AllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *InternalLocationV2AllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *InternalLocationV2AllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetTimezone

`func (o *InternalLocationV2AllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *InternalLocationV2AllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *InternalLocationV2AllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *InternalLocationV2AllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *InternalLocationV2AllOf) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *InternalLocationV2AllOf) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


