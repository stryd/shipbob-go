# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **NullableString** | Name of the application that owns the channel | [optional] 
**Id** | Pointer to **int32** | Unique id of the channel | [optional] 
**Name** | Pointer to **NullableString** | Name of the channel | [optional] 
**Scopes** | Pointer to **[]string** | Array of permissions granted for the channel | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *Channel) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *Channel) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *Channel) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *Channel) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *Channel) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *Channel) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetId

`func (o *Channel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Channel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Channel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Channel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Channel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScopes

`func (o *Channel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Channel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Channel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Channel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *Channel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *Channel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


