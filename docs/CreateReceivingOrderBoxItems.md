# CreateReceivingOrderBoxItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryId** | **int32** | Unique inventory id of the items in the box | 
**LotDate** | Pointer to **NullableTime** | Lot expiration date for the items in the box | [optional] 
**LotNumber** | Pointer to **NullableString** | Lot number of the items in the box | [optional] 
**Quantity** | **int32** | Quantity of the items in the box | 

## Methods

### NewCreateReceivingOrderBoxItems

`func NewCreateReceivingOrderBoxItems(inventoryId int32, quantity int32, ) *CreateReceivingOrderBoxItems`

NewCreateReceivingOrderBoxItems instantiates a new CreateReceivingOrderBoxItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceivingOrderBoxItemsWithDefaults

`func NewCreateReceivingOrderBoxItemsWithDefaults() *CreateReceivingOrderBoxItems`

NewCreateReceivingOrderBoxItemsWithDefaults instantiates a new CreateReceivingOrderBoxItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryId

`func (o *CreateReceivingOrderBoxItems) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *CreateReceivingOrderBoxItems) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *CreateReceivingOrderBoxItems) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.


### GetLotDate

`func (o *CreateReceivingOrderBoxItems) GetLotDate() time.Time`

GetLotDate returns the LotDate field if non-nil, zero value otherwise.

### GetLotDateOk

`func (o *CreateReceivingOrderBoxItems) GetLotDateOk() (*time.Time, bool)`

GetLotDateOk returns a tuple with the LotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotDate

`func (o *CreateReceivingOrderBoxItems) SetLotDate(v time.Time)`

SetLotDate sets LotDate field to given value.

### HasLotDate

`func (o *CreateReceivingOrderBoxItems) HasLotDate() bool`

HasLotDate returns a boolean if a field has been set.

### SetLotDateNil

`func (o *CreateReceivingOrderBoxItems) SetLotDateNil(b bool)`

 SetLotDateNil sets the value for LotDate to be an explicit nil

### UnsetLotDate
`func (o *CreateReceivingOrderBoxItems) UnsetLotDate()`

UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
### GetLotNumber

`func (o *CreateReceivingOrderBoxItems) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *CreateReceivingOrderBoxItems) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *CreateReceivingOrderBoxItems) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *CreateReceivingOrderBoxItems) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *CreateReceivingOrderBoxItems) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *CreateReceivingOrderBoxItems) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetQuantity

`func (o *CreateReceivingOrderBoxItems) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateReceivingOrderBoxItems) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateReceivingOrderBoxItems) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


