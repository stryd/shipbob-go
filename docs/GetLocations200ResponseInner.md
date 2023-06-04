# GetLocations200ResponseInner

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
**Ttype** | **string** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**FulfillmentCenterAttributes** | Pointer to [**[]FcAttribute**](FcAttribute.md) |  | [optional] 
**FulfillmentCenterType** | Pointer to [**NullableInternalLocationAllOfFulfillmentCenterType**](InternalLocationAllOfFulfillmentCenterType.md) |  | [optional] 
**IsEnabledForNewUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetLocations200ResponseInner

`func NewGetLocations200ResponseInner(ttype string, ) *GetLocations200ResponseInner`

NewGetLocations200ResponseInner instantiates a new GetLocations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLocations200ResponseInnerWithDefaults

`func NewGetLocations200ResponseInnerWithDefaults() *GetLocations200ResponseInner`

NewGetLocations200ResponseInnerWithDefaults instantiates a new GetLocations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *GetLocations200ResponseInner) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *GetLocations200ResponseInner) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *GetLocations200ResponseInner) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *GetLocations200ResponseInner) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### SetAbbreviationNil

`func (o *GetLocations200ResponseInner) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *GetLocations200ResponseInner) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetAccessGranted

`func (o *GetLocations200ResponseInner) GetAccessGranted() bool`

GetAccessGranted returns the AccessGranted field if non-nil, zero value otherwise.

### GetAccessGrantedOk

`func (o *GetLocations200ResponseInner) GetAccessGrantedOk() (*bool, bool)`

GetAccessGrantedOk returns a tuple with the AccessGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGranted

`func (o *GetLocations200ResponseInner) SetAccessGranted(v bool)`

SetAccessGranted sets AccessGranted field to given value.

### HasAccessGranted

`func (o *GetLocations200ResponseInner) HasAccessGranted() bool`

HasAccessGranted returns a boolean if a field has been set.

### GetAttributes

`func (o *GetLocations200ResponseInner) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetLocations200ResponseInner) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetLocations200ResponseInner) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetLocations200ResponseInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *GetLocations200ResponseInner) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *GetLocations200ResponseInner) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetId

`func (o *GetLocations200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLocations200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLocations200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLocations200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *GetLocations200ResponseInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetLocations200ResponseInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetLocations200ResponseInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *GetLocations200ResponseInner) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsReceivingEnabled

`func (o *GetLocations200ResponseInner) GetIsReceivingEnabled() bool`

GetIsReceivingEnabled returns the IsReceivingEnabled field if non-nil, zero value otherwise.

### GetIsReceivingEnabledOk

`func (o *GetLocations200ResponseInner) GetIsReceivingEnabledOk() (*bool, bool)`

GetIsReceivingEnabledOk returns a tuple with the IsReceivingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceivingEnabled

`func (o *GetLocations200ResponseInner) SetIsReceivingEnabled(v bool)`

SetIsReceivingEnabled sets IsReceivingEnabled field to given value.

### HasIsReceivingEnabled

`func (o *GetLocations200ResponseInner) HasIsReceivingEnabled() bool`

HasIsReceivingEnabled returns a boolean if a field has been set.

### GetIsShippingEnabled

`func (o *GetLocations200ResponseInner) GetIsShippingEnabled() bool`

GetIsShippingEnabled returns the IsShippingEnabled field if non-nil, zero value otherwise.

### GetIsShippingEnabledOk

`func (o *GetLocations200ResponseInner) GetIsShippingEnabledOk() (*bool, bool)`

GetIsShippingEnabledOk returns a tuple with the IsShippingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingEnabled

`func (o *GetLocations200ResponseInner) SetIsShippingEnabled(v bool)`

SetIsShippingEnabled sets IsShippingEnabled field to given value.

### HasIsShippingEnabled

`func (o *GetLocations200ResponseInner) HasIsShippingEnabled() bool`

HasIsShippingEnabled returns a boolean if a field has been set.

### GetName

`func (o *GetLocations200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLocations200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLocations200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLocations200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetLocations200ResponseInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetLocations200ResponseInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *GetLocations200ResponseInner) GetRegion() FulfillmentCenterRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetLocations200ResponseInner) GetRegionOk() (*FulfillmentCenterRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetLocations200ResponseInner) SetRegion(v FulfillmentCenterRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetLocations200ResponseInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServices

`func (o *GetLocations200ResponseInner) GetServices() []LocationService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetLocations200ResponseInner) GetServicesOk() (*[]LocationService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetLocations200ResponseInner) SetServices(v []LocationService)`

SetServices sets Services field to given value.

### HasServices

`func (o *GetLocations200ResponseInner) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *GetLocations200ResponseInner) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *GetLocations200ResponseInner) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetTimezone

`func (o *GetLocations200ResponseInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetLocations200ResponseInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetLocations200ResponseInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetLocations200ResponseInner) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *GetLocations200ResponseInner) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *GetLocations200ResponseInner) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTtype

`func (o *GetLocations200ResponseInner) GetTtype() string`

GetTtype returns the Ttype field if non-nil, zero value otherwise.

### GetTtypeOk

`func (o *GetLocations200ResponseInner) GetTtypeOk() (*string, bool)`

GetTtypeOk returns a tuple with the Ttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtype

`func (o *GetLocations200ResponseInner) SetTtype(v string)`

SetTtype sets Ttype field to given value.


### GetOrganizationId

`func (o *GetLocations200ResponseInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetLocations200ResponseInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetLocations200ResponseInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetLocations200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOwnerId

`func (o *GetLocations200ResponseInner) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetLocations200ResponseInner) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetLocations200ResponseInner) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GetLocations200ResponseInner) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *GetLocations200ResponseInner) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *GetLocations200ResponseInner) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetParentId

`func (o *GetLocations200ResponseInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetLocations200ResponseInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetLocations200ResponseInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetLocations200ResponseInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *GetLocations200ResponseInner) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *GetLocations200ResponseInner) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetFulfillmentCenterAttributes

`func (o *GetLocations200ResponseInner) GetFulfillmentCenterAttributes() []FcAttribute`

GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field if non-nil, zero value otherwise.

### GetFulfillmentCenterAttributesOk

`func (o *GetLocations200ResponseInner) GetFulfillmentCenterAttributesOk() (*[]FcAttribute, bool)`

GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterAttributes

`func (o *GetLocations200ResponseInner) SetFulfillmentCenterAttributes(v []FcAttribute)`

SetFulfillmentCenterAttributes sets FulfillmentCenterAttributes field to given value.

### HasFulfillmentCenterAttributes

`func (o *GetLocations200ResponseInner) HasFulfillmentCenterAttributes() bool`

HasFulfillmentCenterAttributes returns a boolean if a field has been set.

### SetFulfillmentCenterAttributesNil

`func (o *GetLocations200ResponseInner) SetFulfillmentCenterAttributesNil(b bool)`

 SetFulfillmentCenterAttributesNil sets the value for FulfillmentCenterAttributes to be an explicit nil

### UnsetFulfillmentCenterAttributes
`func (o *GetLocations200ResponseInner) UnsetFulfillmentCenterAttributes()`

UnsetFulfillmentCenterAttributes ensures that no value is present for FulfillmentCenterAttributes, not even an explicit nil
### GetFulfillmentCenterType

`func (o *GetLocations200ResponseInner) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType`

GetFulfillmentCenterType returns the FulfillmentCenterType field if non-nil, zero value otherwise.

### GetFulfillmentCenterTypeOk

`func (o *GetLocations200ResponseInner) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool)`

GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterType

`func (o *GetLocations200ResponseInner) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType)`

SetFulfillmentCenterType sets FulfillmentCenterType field to given value.

### HasFulfillmentCenterType

`func (o *GetLocations200ResponseInner) HasFulfillmentCenterType() bool`

HasFulfillmentCenterType returns a boolean if a field has been set.

### SetFulfillmentCenterTypeNil

`func (o *GetLocations200ResponseInner) SetFulfillmentCenterTypeNil(b bool)`

 SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil

### UnsetFulfillmentCenterType
`func (o *GetLocations200ResponseInner) UnsetFulfillmentCenterType()`

UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
### GetIsEnabledForNewUser

`func (o *GetLocations200ResponseInner) GetIsEnabledForNewUser() bool`

GetIsEnabledForNewUser returns the IsEnabledForNewUser field if non-nil, zero value otherwise.

### GetIsEnabledForNewUserOk

`func (o *GetLocations200ResponseInner) GetIsEnabledForNewUserOk() (*bool, bool)`

GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForNewUser

`func (o *GetLocations200ResponseInner) SetIsEnabledForNewUser(v bool)`

SetIsEnabledForNewUser sets IsEnabledForNewUser field to given value.

### HasIsEnabledForNewUser

`func (o *GetLocations200ResponseInner) HasIsEnabledForNewUser() bool`

HasIsEnabledForNewUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


