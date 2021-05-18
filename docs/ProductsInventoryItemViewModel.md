# ProductsInventoryItemViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the inventory item included in a store product | [optional] 

## Methods

### NewProductsInventoryItemViewModel

`func NewProductsInventoryItemViewModel() *ProductsInventoryItemViewModel`

NewProductsInventoryItemViewModel instantiates a new ProductsInventoryItemViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsInventoryItemViewModelWithDefaults

`func NewProductsInventoryItemViewModelWithDefaults() *ProductsInventoryItemViewModel`

NewProductsInventoryItemViewModelWithDefaults instantiates a new ProductsInventoryItemViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductsInventoryItemViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductsInventoryItemViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductsInventoryItemViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductsInventoryItemViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductsInventoryItemViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsInventoryItemViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsInventoryItemViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductsInventoryItemViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductsInventoryItemViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductsInventoryItemViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *ProductsInventoryItemViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductsInventoryItemViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductsInventoryItemViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductsInventoryItemViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


