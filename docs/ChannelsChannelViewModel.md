# ChannelsChannelViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **NullableString** | Name of the application that owns the channel | [optional] 
**Id** | Pointer to **int32** | Unique id of the channel | [optional] 
**Name** | Pointer to **NullableString** | Name of the channel | [optional] 
**Scopes** | Pointer to **[]string** | Array of permissions granted for the channel | [optional] 

## Methods

### NewChannelsChannelViewModel

`func NewChannelsChannelViewModel() *ChannelsChannelViewModel`

NewChannelsChannelViewModel instantiates a new ChannelsChannelViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelViewModelWithDefaults

`func NewChannelsChannelViewModelWithDefaults() *ChannelsChannelViewModel`

NewChannelsChannelViewModelWithDefaults instantiates a new ChannelsChannelViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *ChannelsChannelViewModel) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ChannelsChannelViewModel) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ChannelsChannelViewModel) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ChannelsChannelViewModel) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *ChannelsChannelViewModel) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *ChannelsChannelViewModel) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetId

`func (o *ChannelsChannelViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelsChannelViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelsChannelViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelsChannelViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChannelsChannelViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelsChannelViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelsChannelViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelsChannelViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChannelsChannelViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChannelsChannelViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScopes

`func (o *ChannelsChannelViewModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ChannelsChannelViewModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ChannelsChannelViewModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ChannelsChannelViewModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ChannelsChannelViewModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ChannelsChannelViewModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


