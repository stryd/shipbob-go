# InventoryLotQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwaitingQuantity** | Pointer to **int32** | Quantity of inventory items belonging to this lot awaiting arrival of a receiving order | [optional] 
**CommittedQuantity** | Pointer to **int32** | Committed quantity of inventory items belonging to this lot | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | Expiration date for this lot | [optional] 
**FulfillableQuantity** | Pointer to **int32** | Fulfillable quantity of inventory items belonging to this lot | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]InventoryFulfillmentCenterQuantity**](InventoryFulfillmentCenterQuantity.md) | Fulfillable quantity of inventory items belonging to this lot broken down by fulfillment center location | [optional] 
**InternalTransferQuantity** | Pointer to **int32** | The quantity of all items belonging to this lot that are in the process of internal transit \\r\\nbetween ShipBob fulfillment centers. These items are not pickable or fulfillable\\r\\nuntil they have been received and moved to the \&quot;On Hand\&quot; quantity of the destination FC.\\r\\nInternal transit quantities for each FC represent the incoming transfer stock\\r\\nfor that specific location. | [optional] 
**LotNumber** | Pointer to **string** | Identification number of this lot | [optional] 
**OnhandQuantity** | Pointer to **int32** | OnHand quantity of inventory items belonging to this lot | [optional] 

## Methods

### NewInventoryLotQuantity

`func NewInventoryLotQuantity() *InventoryLotQuantity`

NewInventoryLotQuantity instantiates a new InventoryLotQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryLotQuantityWithDefaults

`func NewInventoryLotQuantityWithDefaults() *InventoryLotQuantity`

NewInventoryLotQuantityWithDefaults instantiates a new InventoryLotQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwaitingQuantity

`func (o *InventoryLotQuantity) GetAwaitingQuantity() int32`

GetAwaitingQuantity returns the AwaitingQuantity field if non-nil, zero value otherwise.

### GetAwaitingQuantityOk

`func (o *InventoryLotQuantity) GetAwaitingQuantityOk() (*int32, bool)`

GetAwaitingQuantityOk returns a tuple with the AwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitingQuantity

`func (o *InventoryLotQuantity) SetAwaitingQuantity(v int32)`

SetAwaitingQuantity sets AwaitingQuantity field to given value.

### HasAwaitingQuantity

`func (o *InventoryLotQuantity) HasAwaitingQuantity() bool`

HasAwaitingQuantity returns a boolean if a field has been set.

### GetCommittedQuantity

`func (o *InventoryLotQuantity) GetCommittedQuantity() int32`

GetCommittedQuantity returns the CommittedQuantity field if non-nil, zero value otherwise.

### GetCommittedQuantityOk

`func (o *InventoryLotQuantity) GetCommittedQuantityOk() (*int32, bool)`

GetCommittedQuantityOk returns a tuple with the CommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedQuantity

`func (o *InventoryLotQuantity) SetCommittedQuantity(v int32)`

SetCommittedQuantity sets CommittedQuantity field to given value.

### HasCommittedQuantity

`func (o *InventoryLotQuantity) HasCommittedQuantity() bool`

HasCommittedQuantity returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InventoryLotQuantity) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InventoryLotQuantity) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InventoryLotQuantity) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InventoryLotQuantity) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *InventoryLotQuantity) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *InventoryLotQuantity) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetFulfillableQuantity

`func (o *InventoryLotQuantity) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *InventoryLotQuantity) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *InventoryLotQuantity) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *InventoryLotQuantity) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantity) GetFulfillableQuantityByFulfillmentCenter() []InventoryFulfillmentCenterQuantity`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *InventoryLotQuantity) GetFulfillableQuantityByFulfillmentCenterOk() (*[]InventoryFulfillmentCenterQuantity, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantity) SetFulfillableQuantityByFulfillmentCenter(v []InventoryFulfillmentCenterQuantity)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantity) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### GetInternalTransferQuantity

`func (o *InventoryLotQuantity) GetInternalTransferQuantity() int32`

GetInternalTransferQuantity returns the InternalTransferQuantity field if non-nil, zero value otherwise.

### GetInternalTransferQuantityOk

`func (o *InventoryLotQuantity) GetInternalTransferQuantityOk() (*int32, bool)`

GetInternalTransferQuantityOk returns a tuple with the InternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransferQuantity

`func (o *InventoryLotQuantity) SetInternalTransferQuantity(v int32)`

SetInternalTransferQuantity sets InternalTransferQuantity field to given value.

### HasInternalTransferQuantity

`func (o *InventoryLotQuantity) HasInternalTransferQuantity() bool`

HasInternalTransferQuantity returns a boolean if a field has been set.

### GetLotNumber

`func (o *InventoryLotQuantity) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *InventoryLotQuantity) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *InventoryLotQuantity) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *InventoryLotQuantity) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetOnhandQuantity

`func (o *InventoryLotQuantity) GetOnhandQuantity() int32`

GetOnhandQuantity returns the OnhandQuantity field if non-nil, zero value otherwise.

### GetOnhandQuantityOk

`func (o *InventoryLotQuantity) GetOnhandQuantityOk() (*int32, bool)`

GetOnhandQuantityOk returns a tuple with the OnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnhandQuantity

`func (o *InventoryLotQuantity) SetOnhandQuantity(v int32)`

SetOnhandQuantity sets OnhandQuantity field to given value.

### HasOnhandQuantity

`func (o *InventoryLotQuantity) HasOnhandQuantity() bool`

HasOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


