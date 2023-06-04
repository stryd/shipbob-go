# InventoryFulfillmentCenterQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwaitingQuantity** | Pointer to **int32** | Amount of quantity awaiting arrival of a receiving order at this fulfillment center | [optional] 
**CommittedQuantity** | Pointer to **int32** | Amount of committed quantity at this fulfillment center | [optional] 
**FulfillableQuantity** | Pointer to **int32** | Amount of fulfillable quantity at this fulfillment center | [optional] 
**FulfillmentCenterId** | Pointer to **int32** | Unique id of the fulfillment center | [optional] 
**InternalTransferQuantity** | Pointer to **int32** | The quantity of items that are in the process of internal transit  between ShipBob fulfillment centers, with a destination of this fulfillment center. These items are not pickable or fulfillable until they have been received and moved  to the \&quot;On Hand\&quot; quantity of the destination FC. | [optional] 
**Name** | Pointer to **NullableString** | Name of the fulfillment center | [optional] 
**OnHandQuantity** | Pointer to **int32** | Amount of onhand quantity at this fulfillment center | [optional] 

## Methods

### NewInventoryFulfillmentCenterQuantity

`func NewInventoryFulfillmentCenterQuantity() *InventoryFulfillmentCenterQuantity`

NewInventoryFulfillmentCenterQuantity instantiates a new InventoryFulfillmentCenterQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryFulfillmentCenterQuantityWithDefaults

`func NewInventoryFulfillmentCenterQuantityWithDefaults() *InventoryFulfillmentCenterQuantity`

NewInventoryFulfillmentCenterQuantityWithDefaults instantiates a new InventoryFulfillmentCenterQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantity) GetAwaitingQuantity() int32`

GetAwaitingQuantity returns the AwaitingQuantity field if non-nil, zero value otherwise.

### GetAwaitingQuantityOk

`func (o *InventoryFulfillmentCenterQuantity) GetAwaitingQuantityOk() (*int32, bool)`

GetAwaitingQuantityOk returns a tuple with the AwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantity) SetAwaitingQuantity(v int32)`

SetAwaitingQuantity sets AwaitingQuantity field to given value.

### HasAwaitingQuantity

`func (o *InventoryFulfillmentCenterQuantity) HasAwaitingQuantity() bool`

HasAwaitingQuantity returns a boolean if a field has been set.

### GetCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantity) GetCommittedQuantity() int32`

GetCommittedQuantity returns the CommittedQuantity field if non-nil, zero value otherwise.

### GetCommittedQuantityOk

`func (o *InventoryFulfillmentCenterQuantity) GetCommittedQuantityOk() (*int32, bool)`

GetCommittedQuantityOk returns a tuple with the CommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantity) SetCommittedQuantity(v int32)`

SetCommittedQuantity sets CommittedQuantity field to given value.

### HasCommittedQuantity

`func (o *InventoryFulfillmentCenterQuantity) HasCommittedQuantity() bool`

HasCommittedQuantity returns a boolean if a field has been set.

### GetFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantity) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *InventoryFulfillmentCenterQuantity) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantity) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *InventoryFulfillmentCenterQuantity) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetFulfillmentCenterId

`func (o *InventoryFulfillmentCenterQuantity) GetFulfillmentCenterId() int32`

GetFulfillmentCenterId returns the FulfillmentCenterId field if non-nil, zero value otherwise.

### GetFulfillmentCenterIdOk

`func (o *InventoryFulfillmentCenterQuantity) GetFulfillmentCenterIdOk() (*int32, bool)`

GetFulfillmentCenterIdOk returns a tuple with the FulfillmentCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenterId

`func (o *InventoryFulfillmentCenterQuantity) SetFulfillmentCenterId(v int32)`

SetFulfillmentCenterId sets FulfillmentCenterId field to given value.

### HasFulfillmentCenterId

`func (o *InventoryFulfillmentCenterQuantity) HasFulfillmentCenterId() bool`

HasFulfillmentCenterId returns a boolean if a field has been set.

### GetInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantity) GetInternalTransferQuantity() int32`

GetInternalTransferQuantity returns the InternalTransferQuantity field if non-nil, zero value otherwise.

### GetInternalTransferQuantityOk

`func (o *InventoryFulfillmentCenterQuantity) GetInternalTransferQuantityOk() (*int32, bool)`

GetInternalTransferQuantityOk returns a tuple with the InternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantity) SetInternalTransferQuantity(v int32)`

SetInternalTransferQuantity sets InternalTransferQuantity field to given value.

### HasInternalTransferQuantity

`func (o *InventoryFulfillmentCenterQuantity) HasInternalTransferQuantity() bool`

HasInternalTransferQuantity returns a boolean if a field has been set.

### GetName

`func (o *InventoryFulfillmentCenterQuantity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryFulfillmentCenterQuantity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryFulfillmentCenterQuantity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InventoryFulfillmentCenterQuantity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InventoryFulfillmentCenterQuantity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InventoryFulfillmentCenterQuantity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOnHandQuantity

`func (o *InventoryFulfillmentCenterQuantity) GetOnHandQuantity() int32`

GetOnHandQuantity returns the OnHandQuantity field if non-nil, zero value otherwise.

### GetOnHandQuantityOk

`func (o *InventoryFulfillmentCenterQuantity) GetOnHandQuantityOk() (*int32, bool)`

GetOnHandQuantityOk returns a tuple with the OnHandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHandQuantity

`func (o *InventoryFulfillmentCenterQuantity) SetOnHandQuantity(v int32)`

SetOnHandQuantity sets OnHandQuantity field to given value.

### HasOnHandQuantity

`func (o *InventoryFulfillmentCenterQuantity) HasOnHandQuantity() bool`

HasOnHandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


