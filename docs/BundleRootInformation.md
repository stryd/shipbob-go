# BundleRootInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Id of the bundle root product | [optional] 
**Name** | Pointer to **NullableString** | Name of the bundle root product | [optional] 

## Methods

### NewBundleRootInformation

`func NewBundleRootInformation() *BundleRootInformation`

NewBundleRootInformation instantiates a new BundleRootInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleRootInformationWithDefaults

`func NewBundleRootInformationWithDefaults() *BundleRootInformation`

NewBundleRootInformationWithDefaults instantiates a new BundleRootInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BundleRootInformation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleRootInformation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleRootInformation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BundleRootInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BundleRootInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleRootInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleRootInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BundleRootInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BundleRootInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BundleRootInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


