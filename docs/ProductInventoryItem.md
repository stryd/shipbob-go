# ProductInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the inventory item included in a store product | [optional] 

## Methods

### NewProductInventoryItem

`func NewProductInventoryItem() *ProductInventoryItem`

NewProductInventoryItem instantiates a new ProductInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInventoryItemWithDefaults

`func NewProductInventoryItemWithDefaults() *ProductInventoryItem`

NewProductInventoryItemWithDefaults instantiates a new ProductInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductInventoryItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductInventoryItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductInventoryItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductInventoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductInventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductInventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductInventoryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductInventoryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductInventoryItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductInventoryItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuantity

`func (o *ProductInventoryItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductInventoryItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductInventoryItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductInventoryItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


