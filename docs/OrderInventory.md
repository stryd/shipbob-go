# OrderInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **NullableTime** | Expiration date of the inventory | [optional] 
**Id** | Pointer to **int32** | Unique id of the inventory | [optional] 
**IsDangerousGoods** | Pointer to **bool** | Is inventory Dangerous Good | [optional] 
**Lot** | Pointer to **string** | Lot number of the inventory | [optional] 
**Name** | Pointer to **string** | Name of the inventory item | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the inventory item to be included in the fulfillment | [optional] 
**QuantityCommitted** | Pointer to **int32** | The quantity of the inventory item allocated from the assigned fulfillment center and committed to the order. If quantity committed is less than order quantity, then the inventory item is out of stock at the assigned fulfillment center. | [optional] 
**SerialNumbers** | Pointer to **[]string** | Serial number of the inventory | [optional] 

## Methods

### NewOrderInventory

`func NewOrderInventory() *OrderInventory`

NewOrderInventory instantiates a new OrderInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInventoryWithDefaults

`func NewOrderInventoryWithDefaults() *OrderInventory`

NewOrderInventoryWithDefaults instantiates a new OrderInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *OrderInventory) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *OrderInventory) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *OrderInventory) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *OrderInventory) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *OrderInventory) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *OrderInventory) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetId

`func (o *OrderInventory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderInventory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderInventory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderInventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDangerousGoods

`func (o *OrderInventory) GetIsDangerousGoods() bool`

GetIsDangerousGoods returns the IsDangerousGoods field if non-nil, zero value otherwise.

### GetIsDangerousGoodsOk

`func (o *OrderInventory) GetIsDangerousGoodsOk() (*bool, bool)`

GetIsDangerousGoodsOk returns a tuple with the IsDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDangerousGoods

`func (o *OrderInventory) SetIsDangerousGoods(v bool)`

SetIsDangerousGoods sets IsDangerousGoods field to given value.

### HasIsDangerousGoods

`func (o *OrderInventory) HasIsDangerousGoods() bool`

HasIsDangerousGoods returns a boolean if a field has been set.

### GetLot

`func (o *OrderInventory) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *OrderInventory) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *OrderInventory) SetLot(v string)`

SetLot sets Lot field to given value.

### HasLot

`func (o *OrderInventory) HasLot() bool`

HasLot returns a boolean if a field has been set.

### GetName

`func (o *OrderInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderInventory) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderInventory) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderInventory) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderInventory) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityCommitted

`func (o *OrderInventory) GetQuantityCommitted() int32`

GetQuantityCommitted returns the QuantityCommitted field if non-nil, zero value otherwise.

### GetQuantityCommittedOk

`func (o *OrderInventory) GetQuantityCommittedOk() (*int32, bool)`

GetQuantityCommittedOk returns a tuple with the QuantityCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityCommitted

`func (o *OrderInventory) SetQuantityCommitted(v int32)`

SetQuantityCommitted sets QuantityCommitted field to given value.

### HasQuantityCommitted

`func (o *OrderInventory) HasQuantityCommitted() bool`

HasQuantityCommitted returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *OrderInventory) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *OrderInventory) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *OrderInventory) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *OrderInventory) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


