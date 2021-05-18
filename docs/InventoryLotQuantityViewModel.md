# InventoryLotQuantityViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwaitingQuantity** | Pointer to **int32** | Quantity of inventory items belonging to this lot awaiting arrival of a receiving order | [optional] 
**CommittedQuantity** | Pointer to **int32** | Committed quantity of inventory items belonging to this lot | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | Expiration date for this lot | [optional] 
**FulfillableQuantity** | Pointer to **int32** | Fulfillable quantity of inventory items belonging to this lot | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]InventoryFulfillmentCenterQuantityViewModel**](InventoryFulfillmentCenterQuantityViewModel.md) | Fulfillable quantity of inventory items belonging to this lot broken down by fulfillment center location | [optional] 
**InternalTransferQuantity** | Pointer to **int32** | The quantity of all items belonging to this lot that are in the process of internal transit \\r\\nbetween ShipBob fulfillment centers. These items are not pickable or fulfillable\\r\\nuntil they have been received and moved to the \&quot;On Hand\&quot; quantity of the destination FC.\\r\\nInternal transit quantities for each FC represent the incoming transfer stock\\r\\nfor that specific location. | [optional] 
**LotNumber** | Pointer to **string** | Identification number of this lot | [optional] 
**OnhandQuantity** | Pointer to **int32** | OnHand quantity of inventory items belonging to this lot | [optional] 

## Methods

### NewInventoryLotQuantityViewModel

`func NewInventoryLotQuantityViewModel() *InventoryLotQuantityViewModel`

NewInventoryLotQuantityViewModel instantiates a new InventoryLotQuantityViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryLotQuantityViewModelWithDefaults

`func NewInventoryLotQuantityViewModelWithDefaults() *InventoryLotQuantityViewModel`

NewInventoryLotQuantityViewModelWithDefaults instantiates a new InventoryLotQuantityViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwaitingQuantity

`func (o *InventoryLotQuantityViewModel) GetAwaitingQuantity() int32`

GetAwaitingQuantity returns the AwaitingQuantity field if non-nil, zero value otherwise.

### GetAwaitingQuantityOk

`func (o *InventoryLotQuantityViewModel) GetAwaitingQuantityOk() (*int32, bool)`

GetAwaitingQuantityOk returns a tuple with the AwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitingQuantity

`func (o *InventoryLotQuantityViewModel) SetAwaitingQuantity(v int32)`

SetAwaitingQuantity sets AwaitingQuantity field to given value.

### HasAwaitingQuantity

`func (o *InventoryLotQuantityViewModel) HasAwaitingQuantity() bool`

HasAwaitingQuantity returns a boolean if a field has been set.

### GetCommittedQuantity

`func (o *InventoryLotQuantityViewModel) GetCommittedQuantity() int32`

GetCommittedQuantity returns the CommittedQuantity field if non-nil, zero value otherwise.

### GetCommittedQuantityOk

`func (o *InventoryLotQuantityViewModel) GetCommittedQuantityOk() (*int32, bool)`

GetCommittedQuantityOk returns a tuple with the CommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedQuantity

`func (o *InventoryLotQuantityViewModel) SetCommittedQuantity(v int32)`

SetCommittedQuantity sets CommittedQuantity field to given value.

### HasCommittedQuantity

`func (o *InventoryLotQuantityViewModel) HasCommittedQuantity() bool`

HasCommittedQuantity returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InventoryLotQuantityViewModel) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InventoryLotQuantityViewModel) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InventoryLotQuantityViewModel) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InventoryLotQuantityViewModel) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *InventoryLotQuantityViewModel) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *InventoryLotQuantityViewModel) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetFulfillableQuantity

`func (o *InventoryLotQuantityViewModel) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *InventoryLotQuantityViewModel) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *InventoryLotQuantityViewModel) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *InventoryLotQuantityViewModel) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantityViewModel) GetFulfillableQuantityByFulfillmentCenter() []InventoryFulfillmentCenterQuantityViewModel`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *InventoryLotQuantityViewModel) GetFulfillableQuantityByFulfillmentCenterOk() (*[]InventoryFulfillmentCenterQuantityViewModel, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantityViewModel) SetFulfillableQuantityByFulfillmentCenter(v []InventoryFulfillmentCenterQuantityViewModel)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *InventoryLotQuantityViewModel) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### GetInternalTransferQuantity

`func (o *InventoryLotQuantityViewModel) GetInternalTransferQuantity() int32`

GetInternalTransferQuantity returns the InternalTransferQuantity field if non-nil, zero value otherwise.

### GetInternalTransferQuantityOk

`func (o *InventoryLotQuantityViewModel) GetInternalTransferQuantityOk() (*int32, bool)`

GetInternalTransferQuantityOk returns a tuple with the InternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransferQuantity

`func (o *InventoryLotQuantityViewModel) SetInternalTransferQuantity(v int32)`

SetInternalTransferQuantity sets InternalTransferQuantity field to given value.

### HasInternalTransferQuantity

`func (o *InventoryLotQuantityViewModel) HasInternalTransferQuantity() bool`

HasInternalTransferQuantity returns a boolean if a field has been set.

### GetLotNumber

`func (o *InventoryLotQuantityViewModel) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *InventoryLotQuantityViewModel) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *InventoryLotQuantityViewModel) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *InventoryLotQuantityViewModel) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### GetOnhandQuantity

`func (o *InventoryLotQuantityViewModel) GetOnhandQuantity() int32`

GetOnhandQuantity returns the OnhandQuantity field if non-nil, zero value otherwise.

### GetOnhandQuantityOk

`func (o *InventoryLotQuantityViewModel) GetOnhandQuantityOk() (*int32, bool)`

GetOnhandQuantityOk returns a tuple with the OnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnhandQuantity

`func (o *InventoryLotQuantityViewModel) SetOnhandQuantity(v int32)`

SetOnhandQuantity sets OnhandQuantity field to given value.

### HasOnhandQuantity

`func (o *InventoryLotQuantityViewModel) HasOnhandQuantity() bool`

HasOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


