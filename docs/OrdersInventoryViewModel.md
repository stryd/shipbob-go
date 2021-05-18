# OrdersInventoryViewModel

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

### NewOrdersInventoryViewModel

`func NewOrdersInventoryViewModel() *OrdersInventoryViewModel`

NewOrdersInventoryViewModel instantiates a new OrdersInventoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersInventoryViewModelWithDefaults

`func NewOrdersInventoryViewModelWithDefaults() *OrdersInventoryViewModel`

NewOrdersInventoryViewModelWithDefaults instantiates a new OrdersInventoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *OrdersInventoryViewModel) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *OrdersInventoryViewModel) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *OrdersInventoryViewModel) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *OrdersInventoryViewModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *OrdersInventoryViewModel) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *OrdersInventoryViewModel) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetId

`func (o *OrdersInventoryViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersInventoryViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersInventoryViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersInventoryViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDangerousGoods

`func (o *OrdersInventoryViewModel) GetIsDangerousGoods() bool`

GetIsDangerousGoods returns the IsDangerousGoods field if non-nil, zero value otherwise.

### GetIsDangerousGoodsOk

`func (o *OrdersInventoryViewModel) GetIsDangerousGoodsOk() (*bool, bool)`

GetIsDangerousGoodsOk returns a tuple with the IsDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDangerousGoods

`func (o *OrdersInventoryViewModel) SetIsDangerousGoods(v bool)`

SetIsDangerousGoods sets IsDangerousGoods field to given value.

### HasIsDangerousGoods

`func (o *OrdersInventoryViewModel) HasIsDangerousGoods() bool`

HasIsDangerousGoods returns a boolean if a field has been set.

### GetLot

`func (o *OrdersInventoryViewModel) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *OrdersInventoryViewModel) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *OrdersInventoryViewModel) SetLot(v string)`

SetLot sets Lot field to given value.

### HasLot

`func (o *OrdersInventoryViewModel) HasLot() bool`

HasLot returns a boolean if a field has been set.

### GetName

`func (o *OrdersInventoryViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrdersInventoryViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrdersInventoryViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrdersInventoryViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrdersInventoryViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrdersInventoryViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrdersInventoryViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrdersInventoryViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityCommitted

`func (o *OrdersInventoryViewModel) GetQuantityCommitted() int32`

GetQuantityCommitted returns the QuantityCommitted field if non-nil, zero value otherwise.

### GetQuantityCommittedOk

`func (o *OrdersInventoryViewModel) GetQuantityCommittedOk() (*int32, bool)`

GetQuantityCommittedOk returns a tuple with the QuantityCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityCommitted

`func (o *OrdersInventoryViewModel) SetQuantityCommitted(v int32)`

SetQuantityCommitted sets QuantityCommitted field to given value.

### HasQuantityCommitted

`func (o *OrdersInventoryViewModel) HasQuantityCommitted() bool`

HasQuantityCommitted returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *OrdersInventoryViewModel) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *OrdersInventoryViewModel) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *OrdersInventoryViewModel) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *OrdersInventoryViewModel) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


