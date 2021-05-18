# ProductsChannelViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the store channel | [optional] 
**Name** | Pointer to **NullableString** | Name of the store channel | [optional] 

## Methods

### NewProductsChannelViewModel

`func NewProductsChannelViewModel() *ProductsChannelViewModel`

NewProductsChannelViewModel instantiates a new ProductsChannelViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsChannelViewModelWithDefaults

`func NewProductsChannelViewModelWithDefaults() *ProductsChannelViewModel`

NewProductsChannelViewModelWithDefaults instantiates a new ProductsChannelViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductsChannelViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductsChannelViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductsChannelViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductsChannelViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductsChannelViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsChannelViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsChannelViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductsChannelViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductsChannelViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductsChannelViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


