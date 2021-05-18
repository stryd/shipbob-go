# InventoryFulfillmentCenterQuantityViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwaitingQuantity** | Pointer to **int32** | Amount of quantity awaiting arrival of a receiving order at this fulfillment center | [optional] 
**CommittedQuantity** | Pointer to **int32** | Amount of committed quantity at this fulfillment center | [optional] 
**FulfillableQuantity** | Pointer to **int32** | Amount of fulfillable quantity at this fulfillment center | [optional] 
**Id** | Pointer to **int32** | Unique id of the fulfillment center | [optional] 
**InternalTransferQuantity** | Pointer to **int32** | The quantity of items that are in the process of internal transit \\r\\nbetween ShipBob fulfillment centers, with a destination of this fulfillment center.\\r\\nThese items are not pickable or fulfillable until they have been received and moved \\r\\nto the \&quot;On Hand\&quot; quantity of the destination FC. | [optional] 
**Name** | Pointer to **string** | Name of the fulfillment center | [optional] 
**OnhandQuantity** | Pointer to **int32** | Amount of onhand quantity at this fulfillment center | [optional] 

## Methods

### NewInventoryFulfillmentCenterQuantityViewModel

`func NewInventoryFulfillmentCenterQuantityViewModel() *InventoryFulfillmentCenterQuantityViewModel`

NewInventoryFulfillmentCenterQuantityViewModel instantiates a new InventoryFulfillmentCenterQuantityViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryFulfillmentCenterQuantityViewModelWithDefaults

`func NewInventoryFulfillmentCenterQuantityViewModelWithDefaults() *InventoryFulfillmentCenterQuantityViewModel`

NewInventoryFulfillmentCenterQuantityViewModelWithDefaults instantiates a new InventoryFulfillmentCenterQuantityViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetAwaitingQuantity() int32`

GetAwaitingQuantity returns the AwaitingQuantity field if non-nil, zero value otherwise.

### GetAwaitingQuantityOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetAwaitingQuantityOk() (*int32, bool)`

GetAwaitingQuantityOk returns a tuple with the AwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetAwaitingQuantity(v int32)`

SetAwaitingQuantity sets AwaitingQuantity field to given value.

### HasAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasAwaitingQuantity() bool`

HasAwaitingQuantity returns a boolean if a field has been set.

### GetCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetCommittedQuantity() int32`

GetCommittedQuantity returns the CommittedQuantity field if non-nil, zero value otherwise.

### GetCommittedQuantityOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetCommittedQuantityOk() (*int32, bool)`

GetCommittedQuantityOk returns a tuple with the CommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetCommittedQuantity(v int32)`

SetCommittedQuantity sets CommittedQuantity field to given value.

### HasCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasCommittedQuantity() bool`

HasCommittedQuantity returns a boolean if a field has been set.

### GetFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetId

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetInternalTransferQuantity() int32`

GetInternalTransferQuantity returns the InternalTransferQuantity field if non-nil, zero value otherwise.

### GetInternalTransferQuantityOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetInternalTransferQuantityOk() (*int32, bool)`

GetInternalTransferQuantityOk returns a tuple with the InternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetInternalTransferQuantity(v int32)`

SetInternalTransferQuantity sets InternalTransferQuantity field to given value.

### HasInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasInternalTransferQuantity() bool`

HasInternalTransferQuantity returns a boolean if a field has been set.

### GetName

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnhandQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetOnhandQuantity() int32`

GetOnhandQuantity returns the OnhandQuantity field if non-nil, zero value otherwise.

### GetOnhandQuantityOk

`func (o *InventoryFulfillmentCenterQuantityViewModel) GetOnhandQuantityOk() (*int32, bool)`

GetOnhandQuantityOk returns a tuple with the OnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnhandQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) SetOnhandQuantity(v int32)`

SetOnhandQuantity sets OnhandQuantity field to given value.

### HasOnhandQuantity

`func (o *InventoryFulfillmentCenterQuantityViewModel) HasOnhandQuantity() bool`

HasOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


