# ShipBobOrdersPresentationViewModelsInventoryViewModel

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

### NewShipBobOrdersPresentationViewModelsInventoryViewModel

`func NewShipBobOrdersPresentationViewModelsInventoryViewModel() *ShipBobOrdersPresentationViewModelsInventoryViewModel`

NewShipBobOrdersPresentationViewModelsInventoryViewModel instantiates a new ShipBobOrdersPresentationViewModelsInventoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsInventoryViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsInventoryViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsInventoryViewModel`

NewShipBobOrdersPresentationViewModelsInventoryViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsInventoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetId

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDangerousGoods

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetIsDangerousGoods() bool`

GetIsDangerousGoods returns the IsDangerousGoods field if non-nil, zero value otherwise.

### GetIsDangerousGoodsOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetIsDangerousGoodsOk() (*bool, bool)`

GetIsDangerousGoodsOk returns a tuple with the IsDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDangerousGoods

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetIsDangerousGoods(v bool)`

SetIsDangerousGoods sets IsDangerousGoods field to given value.

### HasIsDangerousGoods

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasIsDangerousGoods() bool`

HasIsDangerousGoods returns a boolean if a field has been set.

### GetLot

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetLot(v string)`

SetLot sets Lot field to given value.

### HasLot

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasLot() bool`

HasLot returns a boolean if a field has been set.

### GetName

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityCommitted

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetQuantityCommitted() int32`

GetQuantityCommitted returns the QuantityCommitted field if non-nil, zero value otherwise.

### GetQuantityCommittedOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetQuantityCommittedOk() (*int32, bool)`

GetQuantityCommittedOk returns a tuple with the QuantityCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityCommitted

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetQuantityCommitted(v int32)`

SetQuantityCommitted sets QuantityCommitted field to given value.

### HasQuantityCommitted

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasQuantityCommitted() bool`

HasQuantityCommitted returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetSerialNumbers() []string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) GetSerialNumbersOk() (*[]string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) SetSerialNumbers(v []string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *ShipBobOrdersPresentationViewModelsInventoryViewModel) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


