# ReceivingOrderBoxesInnerBoxItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryId** | Pointer to **int32** | Unique identifier of the inventory item | [optional] 
**LotDate** | Pointer to **NullableTime** | Expiration date of the item&#39;s lot | [optional] 
**LotNumber** | Pointer to **NullableString** | Lot number the item belongs to | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the item included | [optional] 
**ReceivedQuantity** | Pointer to **int32** | Quantity of the item that was received after processing the receiving order | [optional] 

## Methods

### NewReceivingOrderBoxesInnerBoxItemsInner

`func NewReceivingOrderBoxesInnerBoxItemsInner() *ReceivingOrderBoxesInnerBoxItemsInner`

NewReceivingOrderBoxesInnerBoxItemsInner instantiates a new ReceivingOrderBoxesInnerBoxItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingOrderBoxesInnerBoxItemsInnerWithDefaults

`func NewReceivingOrderBoxesInnerBoxItemsInnerWithDefaults() *ReceivingOrderBoxesInnerBoxItemsInner`

NewReceivingOrderBoxesInnerBoxItemsInnerWithDefaults instantiates a new ReceivingOrderBoxesInnerBoxItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryId

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetLotDate

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotDate() time.Time`

GetLotDate returns the LotDate field if non-nil, zero value otherwise.

### GetLotDateOk

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotDateOk() (*time.Time, bool)`

GetLotDateOk returns a tuple with the LotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotDate

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotDate(v time.Time)`

SetLotDate sets LotDate field to given value.

### HasLotDate

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasLotDate() bool`

HasLotDate returns a boolean if a field has been set.

### SetLotDateNil

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotDateNil(b bool)`

 SetLotDateNil sets the value for LotDate to be an explicit nil

### UnsetLotDate
`func (o *ReceivingOrderBoxesInnerBoxItemsInner) UnsetLotDate()`

UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
### GetLotNumber

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *ReceivingOrderBoxesInnerBoxItemsInner) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReceivedQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetReceivedQuantity() int32`

GetReceivedQuantity returns the ReceivedQuantity field if non-nil, zero value otherwise.

### GetReceivedQuantityOk

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) GetReceivedQuantityOk() (*int32, bool)`

GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) SetReceivedQuantity(v int32)`

SetReceivedQuantity sets ReceivedQuantity field to given value.

### HasReceivedQuantity

`func (o *ReceivingOrderBoxesInnerBoxItemsInner) HasReceivedQuantity() bool`

HasReceivedQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


