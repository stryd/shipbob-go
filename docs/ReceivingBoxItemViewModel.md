# ReceivingBoxItemViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryId** | Pointer to **int32** | Unique identifier of the inventory item | [optional] 
**LotDate** | Pointer to **NullableTime** | Expiration date of the item&#39;s lot | [optional] 
**LotNumber** | Pointer to **NullableString** | Lot number the item belongs to | [optional] 
**Quantity** | Pointer to **int32** | Quantity of the item included | [optional] 
**ReceivedQuantity** | Pointer to **int32** | Quantity of the item that was received after processing the receiving order | [optional] 

## Methods

### NewReceivingBoxItemViewModel

`func NewReceivingBoxItemViewModel() *ReceivingBoxItemViewModel`

NewReceivingBoxItemViewModel instantiates a new ReceivingBoxItemViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingBoxItemViewModelWithDefaults

`func NewReceivingBoxItemViewModelWithDefaults() *ReceivingBoxItemViewModel`

NewReceivingBoxItemViewModelWithDefaults instantiates a new ReceivingBoxItemViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryId

`func (o *ReceivingBoxItemViewModel) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *ReceivingBoxItemViewModel) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *ReceivingBoxItemViewModel) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *ReceivingBoxItemViewModel) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetLotDate

`func (o *ReceivingBoxItemViewModel) GetLotDate() time.Time`

GetLotDate returns the LotDate field if non-nil, zero value otherwise.

### GetLotDateOk

`func (o *ReceivingBoxItemViewModel) GetLotDateOk() (*time.Time, bool)`

GetLotDateOk returns a tuple with the LotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotDate

`func (o *ReceivingBoxItemViewModel) SetLotDate(v time.Time)`

SetLotDate sets LotDate field to given value.

### HasLotDate

`func (o *ReceivingBoxItemViewModel) HasLotDate() bool`

HasLotDate returns a boolean if a field has been set.

### SetLotDateNil

`func (o *ReceivingBoxItemViewModel) SetLotDateNil(b bool)`

 SetLotDateNil sets the value for LotDate to be an explicit nil

### UnsetLotDate
`func (o *ReceivingBoxItemViewModel) UnsetLotDate()`

UnsetLotDate ensures that no value is present for LotDate, not even an explicit nil
### GetLotNumber

`func (o *ReceivingBoxItemViewModel) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *ReceivingBoxItemViewModel) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *ReceivingBoxItemViewModel) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *ReceivingBoxItemViewModel) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *ReceivingBoxItemViewModel) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *ReceivingBoxItemViewModel) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetQuantity

`func (o *ReceivingBoxItemViewModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReceivingBoxItemViewModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReceivingBoxItemViewModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReceivingBoxItemViewModel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReceivedQuantity

`func (o *ReceivingBoxItemViewModel) GetReceivedQuantity() int32`

GetReceivedQuantity returns the ReceivedQuantity field if non-nil, zero value otherwise.

### GetReceivedQuantityOk

`func (o *ReceivingBoxItemViewModel) GetReceivedQuantityOk() (*int32, bool)`

GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQuantity

`func (o *ReceivingBoxItemViewModel) SetReceivedQuantity(v int32)`

SetReceivedQuantity sets ReceivedQuantity field to given value.

### HasReceivedQuantity

`func (o *ReceivingBoxItemViewModel) HasReceivedQuantity() bool`

HasReceivedQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


