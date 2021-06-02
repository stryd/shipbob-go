# ChannelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the channel | [optional] 
**Name** | Pointer to **string** | Name of the channel | [optional] 

## Methods

### NewChannelInfo

`func NewChannelInfo() *ChannelInfo`

NewChannelInfo instantiates a new ChannelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInfoWithDefaults

`func NewChannelInfoWithDefaults() *ChannelInfo`

NewChannelInfoWithDefaults instantiates a new ChannelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChannelInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


