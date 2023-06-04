# FcAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique Id for the fulfillment center attribute | [optional] 
**Name** | Pointer to **NullableString** | Name of the attribute. | [optional] 

## Methods

### NewFcAttribute

`func NewFcAttribute() *FcAttribute`

NewFcAttribute instantiates a new FcAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcAttributeWithDefaults

`func NewFcAttributeWithDefaults() *FcAttribute`

NewFcAttributeWithDefaults instantiates a new FcAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FcAttribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcAttribute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcAttribute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FcAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FcAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FcAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FcAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FcAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FcAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FcAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


