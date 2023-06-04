# InternalLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentCenterAttributes** | Pointer to [**[]FcAttribute**](FcAttribute.md) |  | [optional] 
**FulfillmentCenterType** | Pointer to [**NullableInternalLocationAllOfFulfillmentCenterType**](InternalLocationAllOfFulfillmentCenterType.md) |  | [optional] 
**IsEnabledForNewUser** | Pointer to **bool** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 

## Methods

### NewInternalLocation

`func NewInternalLocation() *InternalLocation`

NewInternalLocation instantiates a new InternalLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalLocationWithDefaults

`func NewInternalLocationWithDefaults() *InternalLocation`

NewInternalLocationWithDefaults instantiates a new InternalLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentCenterAttributes

`func (o *InternalLocation) GetFulfillmentCenterAttributes() []FcAttribute`

GetFulfillmentCenterAttributes returns the FulfillmentCenterAttributes field if non-nil, zero value otherwise.

### GetFulfillmentCenterAttributesOk

`func (o *InternalLocation) GetFulfillmentCenterAttributesOk() (*[]FcAttribute, bool)`

GetFulfillmentCenterAttributesOk returns a tuple with the FulfillmentCenterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterAttributes

`func (o *InternalLocation) SetFulfillmentCenterAttributes(v []FcAttribute)`

SetFulfillmentCenterAttributes sets FulfillmentCenterAttributes field to given value.

### HasFulfillmentCenterAttributes

`func (o *InternalLocation) HasFulfillmentCenterAttributes() bool`

HasFulfillmentCenterAttributes returns a boolean if a field has been set.

### SetFulfillmentCenterAttributesNil

`func (o *InternalLocation) SetFulfillmentCenterAttributesNil(b bool)`

 SetFulfillmentCenterAttributesNil sets the value for FulfillmentCenterAttributes to be an explicit nil

### UnsetFulfillmentCenterAttributes
`func (o *InternalLocation) UnsetFulfillmentCenterAttributes()`

UnsetFulfillmentCenterAttributes ensures that no value is present for FulfillmentCenterAttributes, not even an explicit nil
### GetFulfillmentCenterType

`func (o *InternalLocation) GetFulfillmentCenterType() InternalLocationAllOfFulfillmentCenterType`

GetFulfillmentCenterType returns the FulfillmentCenterType field if non-nil, zero value otherwise.

### GetFulfillmentCenterTypeOk

`func (o *InternalLocation) GetFulfillmentCenterTypeOk() (*InternalLocationAllOfFulfillmentCenterType, bool)`

GetFulfillmentCenterTypeOk returns a tuple with the FulfillmentCenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterType

`func (o *InternalLocation) SetFulfillmentCenterType(v InternalLocationAllOfFulfillmentCenterType)`

SetFulfillmentCenterType sets FulfillmentCenterType field to given value.

### HasFulfillmentCenterType

`func (o *InternalLocation) HasFulfillmentCenterType() bool`

HasFulfillmentCenterType returns a boolean if a field has been set.

### SetFulfillmentCenterTypeNil

`func (o *InternalLocation) SetFulfillmentCenterTypeNil(b bool)`

 SetFulfillmentCenterTypeNil sets the value for FulfillmentCenterType to be an explicit nil

### UnsetFulfillmentCenterType
`func (o *InternalLocation) UnsetFulfillmentCenterType()`

UnsetFulfillmentCenterType ensures that no value is present for FulfillmentCenterType, not even an explicit nil
### GetIsEnabledForNewUser

`func (o *InternalLocation) GetIsEnabledForNewUser() bool`

GetIsEnabledForNewUser returns the IsEnabledForNewUser field if non-nil, zero value otherwise.

### GetIsEnabledForNewUserOk

`func (o *InternalLocation) GetIsEnabledForNewUserOk() (*bool, bool)`

GetIsEnabledForNewUserOk returns a tuple with the IsEnabledForNewUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForNewUser

`func (o *InternalLocation) SetIsEnabledForNewUser(v bool)`

SetIsEnabledForNewUser sets IsEnabledForNewUser field to given value.

### HasIsEnabledForNewUser

`func (o *InternalLocation) HasIsEnabledForNewUser() bool`

HasIsEnabledForNewUser returns a boolean if a field has been set.

### GetIsExternal

`func (o *InternalLocation) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *InternalLocation) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *InternalLocation) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *InternalLocation) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


