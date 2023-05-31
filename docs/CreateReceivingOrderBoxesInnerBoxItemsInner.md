# CreateReceivingOrderBoxesInnerBoxItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryId** | **int32** | Unique inventory id of the items in the box | 
**LotDate** | Pointer to **NullableTime** | Lot expiration date for the items in the box | [optional] 
**LotNumber** | Pointer to **NullableString** | Lot number of the items in the box | [optional] 
**Quantity** | **int32** | Quantity of the items in the box | 

## Methods

### NewCreateReceivingOrderBoxesInnerBoxItemsInner

`func NewCreateReceivingOrderBoxesInnerBoxItemsInner(inventoryId int32, quantity int32, ) *CreateReceivingOrderBoxesInnerBoxItemsInner`

NewCreateReceivingOrderBoxesInnerBoxItemsInner instantiates a new CreateReceivingOrderBoxesInnerBoxItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceivingOrderBoxesInnerBoxItemsInnerWithDefaults

`func NewCreateReceivingOrderBoxesInnerBoxItemsInnerWithDefaults() *CreateReceivingOrderBoxesInnerBoxItemsInner`

NewCreateReceivingOrderBoxesInnerBoxItemsInnerWithDefaults instantiates a new CreateReceivingOrderBoxesInnerBoxItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryId

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.


### GetLotDate

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetLotDate() time.Time`

GetLotDate returns the LotDate field if non-nil, zero value otherwise.

### GetLotDateOk

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetLotDateOk() (*time.Time, bool)`

GetLotDateOk returns a tuple with the LotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotDate

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetLotDate(v time.Time)`

SetLotDate sets LotDate field to given value.

### HasLotDate

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) HasLotDate() bool`

HasLotDate returns a boolean if a field has been set.

### SetLotDateNil

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetLotDateNil(b bool)`

 SetLotDateNil sets the value for LotDate to be an explicit nil

### UnsetLotDate
`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) UnsetLotDate()`

UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
### GetLotNumber

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetQuantity

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateReceivingOrderBoxesInnerBoxItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


