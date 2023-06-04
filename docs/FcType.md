# FcType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique Id for the fulfillment center type | [optional] 
**Name** | Pointer to **NullableString** | Name of the fc type | [optional] 
**Ttype** | **string** |  | 

## Methods

### NewFcType

`func NewFcType(ttype string, ) *FcType`

NewFcType instantiates a new FcType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcTypeWithDefaults

`func NewFcTypeWithDefaults() *FcType`

NewFcTypeWithDefaults instantiates a new FcType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FcType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FcType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FcType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FcType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FcType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FcType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FcType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FcType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTtype

`func (o *FcType) GetTtype() string`

GetTtype returns the Ttype field if non-nil, zero value otherwise.

### GetTtypeOk

`func (o *FcType) GetTtypeOk() (*string, bool)`

GetTtypeOk returns a tuple with the Ttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtype

`func (o *FcType) SetTtype(v string)`

SetTtype sets Ttype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

