# LocationRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique Id for the location region | [optional] 
**Name** | Pointer to **NullableString** | Name of the region the location is in. | [optional] 

## Methods

### NewLocationRegion

`func NewLocationRegion() *LocationRegion`

NewLocationRegion instantiates a new LocationRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationRegionWithDefaults

`func NewLocationRegionWithDefaults() *LocationRegion`

NewLocationRegionWithDefaults instantiates a new LocationRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationRegion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationRegion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationRegion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LocationRegion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LocationRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationRegion) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LocationRegion) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocationRegion) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


