# Inventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to [**Dimension**](Dimension.md) |  | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]InventoryFulfillmentCenterQuantity**](InventoryFulfillmentCenterQuantity.md) | Fulfillable quantity of this inventory item broken down by fulfillment center location | [optional] 
**FulfillableQuantityByLot** | Pointer to [**[]InventoryLotQuantity**](InventoryLotQuantity.md) | Fulfillable quantity of this inventory item broken down by lot | [optional] 
**InventoryId** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**IsActive** | Pointer to **bool** | Whether the inventory is active or not | [optional] 
**IsCasePick** | Pointer to **bool** | True if the inventory item is marked as case pick | [optional] 
**IsDigital** | Pointer to **bool** | True if the inventory item is marked as a digital item | [optional] 
**IsLot** | Pointer to **bool** | True if this inventory item is organized into lots | [optional] 
**Name** | Pointer to **NullableString** | Name of the inventory item | [optional] 
**PackagingAttribute** | Pointer to [**PackagingAttribute**](PackagingAttribute.md) |  | [optional] 
**TotalAwaitingQuantity** | Pointer to **int32** | Total quantity in unreceived receiving orders for this inventory item | [optional] 
**TotalBackorderedQuantity** | Pointer to **int32** | The amount of the item you need to send to ShipBob to fulfill all exception orders containing  the item. This is the exception_quantity less the fulfillable_quantity of the item. | [optional] 
**TotalCommitedQuantity** | Pointer to **int32** | Total committed quantity of this inventory item | [optional] 
**TotalExceptionQuantity** | Pointer to **int32** | The total quantity of all items that are contained within orders that are in exception/out of stock status. Out of stock orders have not been processed and therefore do not have lot or fulfillment centers assigned. | [optional] 
**TotalFulfillableQuantity** | Pointer to **int32** | Total fulfillable quantity of this inventory item | [optional] 
**TotalInternalTransferQuantity** | Pointer to **int32** | The total quantity of all items that are in the process of internal transit  between ShipBob fulfillment centers. These items are not pickable or fulfillable until they have been received and moved to the \&quot;On Hand\&quot; quantity of the destination FC. Internal transit quantities for each FC represent the incoming transfer stock for that specific location. | [optional] 
**TotalOnHandQuantity** | Pointer to **int32** | Total onhand quantity of this inventory item | [optional] 
**TotalSellableQuantity** | Pointer to **int32** | Total quantity that can be sold without overselling the inventory item. This is calculated by subtracting the total exception quantity from the fulfillable quantity. | [optional] 

## Methods

### NewInventory

`func NewInventory() *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *Inventory) GetDimensions() Dimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Inventory) GetDimensionsOk() (*Dimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Inventory) SetDimensions(v Dimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *Inventory) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetFulfillableQuantityByFulfillmentCenter

`func (o *Inventory) GetFulfillableQuantityByFulfillmentCenter() []InventoryFulfillmentCenterQuantity`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *Inventory) GetFulfillableQuantityByFulfillmentCenterOk() (*[]InventoryFulfillmentCenterQuantity, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *Inventory) SetFulfillableQuantityByFulfillmentCenter(v []InventoryFulfillmentCenterQuantity)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *Inventory) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### SetFulfillableQuantityByFulfillmentCenterNil

`func (o *Inventory) SetFulfillableQuantityByFulfillmentCenterNil(b bool)`

 SetFulfillableQuantityByFulfillmentCenterNil sets the value for FulfillableQuantityByFulfillmentCenter to be an explicit nil

### UnsetFulfillableQuantityByFulfillmentCenter
`func (o *Inventory) UnsetFulfillableQuantityByFulfillmentCenter()`

UnsetFulfillableQuantityByFulfillmentCenter ensures that no value is present for FulfillableQuantityByFulfillmentCenter, not even an explicit nil
### GetFulfillableQuantityByLot

`func (o *Inventory) GetFulfillableQuantityByLot() []InventoryLotQuantity`

GetFulfillableQuantityByLot returns the FulfillableQuantityByLot field if non-nil, zero value otherwise.

### GetFulfillableQuantityByLotOk

`func (o *Inventory) GetFulfillableQuantityByLotOk() (*[]InventoryLotQuantity, bool)`

GetFulfillableQuantityByLotOk returns a tuple with the FulfillableQuantityByLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByLot

`func (o *Inventory) SetFulfillableQuantityByLot(v []InventoryLotQuantity)`

SetFulfillableQuantityByLot sets FulfillableQuantityByLot field to given value.

### HasFulfillableQuantityByLot

`func (o *Inventory) HasFulfillableQuantityByLot() bool`

HasFulfillableQuantityByLot returns a boolean if a field has been set.

### SetFulfillableQuantityByLotNil

`func (o *Inventory) SetFulfillableQuantityByLotNil(b bool)`

 SetFulfillableQuantityByLotNil sets the value for FulfillableQuantityByLot to be an explicit nil

### UnsetFulfillableQuantityByLot
`func (o *Inventory) UnsetFulfillableQuantityByLot()`

UnsetFulfillableQuantityByLot ensures that no value is present for FulfillableQuantityByLot, not even an explicit nil
### GetInventoryId

`func (o *Inventory) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *Inventory) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *Inventory) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *Inventory) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetIsActive

`func (o *Inventory) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Inventory) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Inventory) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Inventory) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsCasePick

`func (o *Inventory) GetIsCasePick() bool`

GetIsCasePick returns the IsCasePick field if non-nil, zero value otherwise.

### GetIsCasePickOk

`func (o *Inventory) GetIsCasePickOk() (*bool, bool)`

GetIsCasePickOk returns a tuple with the IsCasePick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCasePick

`func (o *Inventory) SetIsCasePick(v bool)`

SetIsCasePick sets IsCasePick field to given value.

### HasIsCasePick

`func (o *Inventory) HasIsCasePick() bool`

HasIsCasePick returns a boolean if a field has been set.

### GetIsDigital

`func (o *Inventory) GetIsDigital() bool`

GetIsDigital returns the IsDigital field if non-nil, zero value otherwise.

### GetIsDigitalOk

`func (o *Inventory) GetIsDigitalOk() (*bool, bool)`

GetIsDigitalOk returns a tuple with the IsDigital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigital

`func (o *Inventory) SetIsDigital(v bool)`

SetIsDigital sets IsDigital field to given value.

### HasIsDigital

`func (o *Inventory) HasIsDigital() bool`

HasIsDigital returns a boolean if a field has been set.

### GetIsLot

`func (o *Inventory) GetIsLot() bool`

GetIsLot returns the IsLot field if non-nil, zero value otherwise.

### GetIsLotOk

`func (o *Inventory) GetIsLotOk() (*bool, bool)`

GetIsLotOk returns a tuple with the IsLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLot

`func (o *Inventory) SetIsLot(v bool)`

SetIsLot sets IsLot field to given value.

### HasIsLot

`func (o *Inventory) HasIsLot() bool`

HasIsLot returns a boolean if a field has been set.

### GetName

`func (o *Inventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Inventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Inventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Inventory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Inventory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Inventory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackagingAttribute

`func (o *Inventory) GetPackagingAttribute() PackagingAttribute`

GetPackagingAttribute returns the PackagingAttribute field if non-nil, zero value otherwise.

### GetPackagingAttributeOk

`func (o *Inventory) GetPackagingAttributeOk() (*PackagingAttribute, bool)`

GetPackagingAttributeOk returns a tuple with the PackagingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingAttribute

`func (o *Inventory) SetPackagingAttribute(v PackagingAttribute)`

SetPackagingAttribute sets PackagingAttribute field to given value.

### HasPackagingAttribute

`func (o *Inventory) HasPackagingAttribute() bool`

HasPackagingAttribute returns a boolean if a field has been set.

### GetTotalAwaitingQuantity

`func (o *Inventory) GetTotalAwaitingQuantity() int32`

GetTotalAwaitingQuantity returns the TotalAwaitingQuantity field if non-nil, zero value otherwise.

### GetTotalAwaitingQuantityOk

`func (o *Inventory) GetTotalAwaitingQuantityOk() (*int32, bool)`

GetTotalAwaitingQuantityOk returns a tuple with the TotalAwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAwaitingQuantity

`func (o *Inventory) SetTotalAwaitingQuantity(v int32)`

SetTotalAwaitingQuantity sets TotalAwaitingQuantity field to given value.

### HasTotalAwaitingQuantity

`func (o *Inventory) HasTotalAwaitingQuantity() bool`

HasTotalAwaitingQuantity returns a boolean if a field has been set.

### GetTotalBackorderedQuantity

`func (o *Inventory) GetTotalBackorderedQuantity() int32`

GetTotalBackorderedQuantity returns the TotalBackorderedQuantity field if non-nil, zero value otherwise.

### GetTotalBackorderedQuantityOk

`func (o *Inventory) GetTotalBackorderedQuantityOk() (*int32, bool)`

GetTotalBackorderedQuantityOk returns a tuple with the TotalBackorderedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackorderedQuantity

`func (o *Inventory) SetTotalBackorderedQuantity(v int32)`

SetTotalBackorderedQuantity sets TotalBackorderedQuantity field to given value.

### HasTotalBackorderedQuantity

`func (o *Inventory) HasTotalBackorderedQuantity() bool`

HasTotalBackorderedQuantity returns a boolean if a field has been set.

### GetTotalCommitedQuantity

`func (o *Inventory) GetTotalCommitedQuantity() int32`

GetTotalCommitedQuantity returns the TotalCommitedQuantity field if non-nil, zero value otherwise.

### GetTotalCommitedQuantityOk

`func (o *Inventory) GetTotalCommitedQuantityOk() (*int32, bool)`

GetTotalCommitedQuantityOk returns a tuple with the TotalCommitedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitedQuantity

`func (o *Inventory) SetTotalCommitedQuantity(v int32)`

SetTotalCommitedQuantity sets TotalCommitedQuantity field to given value.

### HasTotalCommitedQuantity

`func (o *Inventory) HasTotalCommitedQuantity() bool`

HasTotalCommitedQuantity returns a boolean if a field has been set.

### GetTotalExceptionQuantity

`func (o *Inventory) GetTotalExceptionQuantity() int32`

GetTotalExceptionQuantity returns the TotalExceptionQuantity field if non-nil, zero value otherwise.

### GetTotalExceptionQuantityOk

`func (o *Inventory) GetTotalExceptionQuantityOk() (*int32, bool)`

GetTotalExceptionQuantityOk returns a tuple with the TotalExceptionQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExceptionQuantity

`func (o *Inventory) SetTotalExceptionQuantity(v int32)`

SetTotalExceptionQuantity sets TotalExceptionQuantity field to given value.

### HasTotalExceptionQuantity

`func (o *Inventory) HasTotalExceptionQuantity() bool`

HasTotalExceptionQuantity returns a boolean if a field has been set.

### GetTotalFulfillableQuantity

`func (o *Inventory) GetTotalFulfillableQuantity() int32`

GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalFulfillableQuantityOk

`func (o *Inventory) GetTotalFulfillableQuantityOk() (*int32, bool)`

GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFulfillableQuantity

`func (o *Inventory) SetTotalFulfillableQuantity(v int32)`

SetTotalFulfillableQuantity sets TotalFulfillableQuantity field to given value.

### HasTotalFulfillableQuantity

`func (o *Inventory) HasTotalFulfillableQuantity() bool`

HasTotalFulfillableQuantity returns a boolean if a field has been set.

### GetTotalInternalTransferQuantity

`func (o *Inventory) GetTotalInternalTransferQuantity() int32`

GetTotalInternalTransferQuantity returns the TotalInternalTransferQuantity field if non-nil, zero value otherwise.

### GetTotalInternalTransferQuantityOk

`func (o *Inventory) GetTotalInternalTransferQuantityOk() (*int32, bool)`

GetTotalInternalTransferQuantityOk returns a tuple with the TotalInternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInternalTransferQuantity

`func (o *Inventory) SetTotalInternalTransferQuantity(v int32)`

SetTotalInternalTransferQuantity sets TotalInternalTransferQuantity field to given value.

### HasTotalInternalTransferQuantity

`func (o *Inventory) HasTotalInternalTransferQuantity() bool`

HasTotalInternalTransferQuantity returns a boolean if a field has been set.

### GetTotalOnHandQuantity

`func (o *Inventory) GetTotalOnHandQuantity() int32`

GetTotalOnHandQuantity returns the TotalOnHandQuantity field if non-nil, zero value otherwise.

### GetTotalOnHandQuantityOk

`func (o *Inventory) GetTotalOnHandQuantityOk() (*int32, bool)`

GetTotalOnHandQuantityOk returns a tuple with the TotalOnHandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnHandQuantity

`func (o *Inventory) SetTotalOnHandQuantity(v int32)`

SetTotalOnHandQuantity sets TotalOnHandQuantity field to given value.

### HasTotalOnHandQuantity

`func (o *Inventory) HasTotalOnHandQuantity() bool`

HasTotalOnHandQuantity returns a boolean if a field has been set.

### GetTotalSellableQuantity

`func (o *Inventory) GetTotalSellableQuantity() int32`

GetTotalSellableQuantity returns the TotalSellableQuantity field if non-nil, zero value otherwise.

### GetTotalSellableQuantityOk

`func (o *Inventory) GetTotalSellableQuantityOk() (*int32, bool)`

GetTotalSellableQuantityOk returns a tuple with the TotalSellableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSellableQuantity

`func (o *Inventory) SetTotalSellableQuantity(v int32)`

SetTotalSellableQuantity sets TotalSellableQuantity field to given value.

### HasTotalSellableQuantity

`func (o *Inventory) HasTotalSellableQuantity() bool`

HasTotalSellableQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


