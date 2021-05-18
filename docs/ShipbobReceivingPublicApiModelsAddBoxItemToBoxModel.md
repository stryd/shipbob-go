# ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryId** | **int32** | Unique inventory id of the items in the box | 
**LotDate** | Pointer to **NullableTime** | Lot expiration date for the items in the box | [optional] 
**LotNumber** | Pointer to **NullableString** | Lot number of the items in the box | [optional] 
**Quantity** | **int32** | Quantity of the items in the box | 

## Methods

### NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModel

`func NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModel(inventoryId int32, quantity int32, ) *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel`

NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModel instantiates a new ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModelWithDefaults

`func NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModelWithDefaults() *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel`

NewShipbobReceivingPublicApiModelsAddBoxItemToBoxModelWithDefaults instantiates a new ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryId

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.


### GetLotDate

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetLotDate() time.Time`

GetLotDate returns the LotDate field if non-nil, zero value otherwise.

### GetLotDateOk

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetLotDateOk() (*time.Time, bool)`

GetLotDateOk returns a tuple with the LotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotDate

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetLotDate(v time.Time)`

SetLotDate sets LotDate field to given value.

### HasLotDate

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) HasLotDate() bool`

HasLotDate returns a boolean if a field has been set.

### SetLotDateNil

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetLotDateNil(b bool)`

 SetLotDateNil sets the value for LotDate to be an explicit nil

### UnsetLotDate
`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) UnsetLotDate()`

UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
### GetLotNumber

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetQuantity

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipbobReceivingPublicApiModelsAddBoxItemToBoxModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


