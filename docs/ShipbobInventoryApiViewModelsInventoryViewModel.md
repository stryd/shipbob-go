# ShipbobInventoryApiViewModelsInventoryViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to [**ShipbobInventoryApiViewModelsDimensionViewModel**](Shipbob.Inventory.Api.ViewModels.DimensionViewModel.md) |  | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel**](ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel.md) | Fulfillable quantity of this inventory item broken down by fulfillment center location | [optional] 
**FulfillableQuantityByLot** | Pointer to [**[]ShipbobInventoryApiViewModelsLotQuantityViewModel**](ShipbobInventoryApiViewModelsLotQuantityViewModel.md) | Fulfillable quantity of this inventory item broken down by lot | [optional] 
**Id** | Pointer to **int32** | Unique id of the inventory item | [optional] 
**IsActive** | Pointer to **bool** | Whether the inventory is active or not | [optional] 
**IsCasePick** | Pointer to **bool** | True if the inventory item is marked as case pick | [optional] 
**IsDigital** | Pointer to **bool** | True if the inventory item is marked as a digital item | [optional] 
**IsLot** | Pointer to **bool** | True if this inventory item is organized into lots | [optional] 
**Name** | Pointer to **string** | Name of the inventory item | [optional] 
**PackagingAttribute** | Pointer to **string** | Attribute influencing the packaging requirements of this inventory item | [optional] 
**TotalAwaitingQuantity** | Pointer to **int32** | Total quantity in unreceived receiving orders for this inventory item | [optional] 
**TotalBackorderedQuantity** | Pointer to **int32** | The amount of the item you need to send to ShipBob to fulfill all exception orders containing \\r\\nthe item. This is the exception_quantity less the fulfillable_quantity of the item. | [optional] 
**TotalCommittedQuantity** | Pointer to **int32** | Total committed quantity of this inventory item | [optional] 
**TotalExceptionQuantity** | Pointer to **int32** | The total quantity of all items that are contained within orders that\\r\\nare in exception/out of stock status. Out of stock orders have not been\\r\\nprocessed and therefore do not have lot or fulfillment centers assigned. | [optional] 
**TotalFulfillableQuantity** | Pointer to **int32** | Total fulfillable quantity of this inventory item | [optional] 
**TotalInternalTransferQuantity** | Pointer to **int32** | The total quantity of all items that are in the process of internal transit \\r\\nbetween ShipBob fulfillment centers. These items are not pickable or fulfillable\\r\\nuntil they have been received and moved to the \&quot;On Hand\&quot; quantity of the destination FC.\\r\\nInternal transit quantities for each FC represent the incoming transfer stock\\r\\nfor that specific location. | [optional] 
**TotalOnhandQuantity** | Pointer to **int32** | Total onhand quantity of this inventory item | [optional] 

## Methods

### NewShipbobInventoryApiViewModelsInventoryViewModel

`func NewShipbobInventoryApiViewModelsInventoryViewModel() *ShipbobInventoryApiViewModelsInventoryViewModel`

NewShipbobInventoryApiViewModelsInventoryViewModel instantiates a new ShipbobInventoryApiViewModelsInventoryViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobInventoryApiViewModelsInventoryViewModelWithDefaults

`func NewShipbobInventoryApiViewModelsInventoryViewModelWithDefaults() *ShipbobInventoryApiViewModelsInventoryViewModel`

NewShipbobInventoryApiViewModelsInventoryViewModelWithDefaults instantiates a new ShipbobInventoryApiViewModelsInventoryViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetDimensions() ShipbobInventoryApiViewModelsDimensionViewModel`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetDimensionsOk() (*ShipbobInventoryApiViewModelsDimensionViewModel, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetDimensions(v ShipbobInventoryApiViewModelsDimensionViewModel)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByFulfillmentCenter() []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByFulfillmentCenterOk() (*[]ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetFulfillableQuantityByFulfillmentCenter(v []ShipbobInventoryApiViewModelsFulfillmentCenterQuantityViewModel)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### GetFulfillableQuantityByLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByLot() []ShipbobInventoryApiViewModelsLotQuantityViewModel`

GetFulfillableQuantityByLot returns the FulfillableQuantityByLot field if non-nil, zero value otherwise.

### GetFulfillableQuantityByLotOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetFulfillableQuantityByLotOk() (*[]ShipbobInventoryApiViewModelsLotQuantityViewModel, bool)`

GetFulfillableQuantityByLotOk returns a tuple with the FulfillableQuantityByLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetFulfillableQuantityByLot(v []ShipbobInventoryApiViewModelsLotQuantityViewModel)`

SetFulfillableQuantityByLot sets FulfillableQuantityByLot field to given value.

### HasFulfillableQuantityByLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasFulfillableQuantityByLot() bool`

HasFulfillableQuantityByLot returns a boolean if a field has been set.

### GetId

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsCasePick

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsCasePick() bool`

GetIsCasePick returns the IsCasePick field if non-nil, zero value otherwise.

### GetIsCasePickOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsCasePickOk() (*bool, bool)`

GetIsCasePickOk returns a tuple with the IsCasePick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCasePick

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsCasePick(v bool)`

SetIsCasePick sets IsCasePick field to given value.

### HasIsCasePick

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsCasePick() bool`

HasIsCasePick returns a boolean if a field has been set.

### GetIsDigital

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsDigital() bool`

GetIsDigital returns the IsDigital field if non-nil, zero value otherwise.

### GetIsDigitalOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsDigitalOk() (*bool, bool)`

GetIsDigitalOk returns a tuple with the IsDigital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigital

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsDigital(v bool)`

SetIsDigital sets IsDigital field to given value.

### HasIsDigital

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsDigital() bool`

HasIsDigital returns a boolean if a field has been set.

### GetIsLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsLot() bool`

GetIsLot returns the IsLot field if non-nil, zero value otherwise.

### GetIsLotOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetIsLotOk() (*bool, bool)`

GetIsLotOk returns a tuple with the IsLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetIsLot(v bool)`

SetIsLot sets IsLot field to given value.

### HasIsLot

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasIsLot() bool`

HasIsLot returns a boolean if a field has been set.

### GetName

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackagingAttribute

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetPackagingAttribute() string`

GetPackagingAttribute returns the PackagingAttribute field if non-nil, zero value otherwise.

### GetPackagingAttributeOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetPackagingAttributeOk() (*string, bool)`

GetPackagingAttributeOk returns a tuple with the PackagingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingAttribute

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetPackagingAttribute(v string)`

SetPackagingAttribute sets PackagingAttribute field to given value.

### HasPackagingAttribute

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasPackagingAttribute() bool`

HasPackagingAttribute returns a boolean if a field has been set.

### GetTotalAwaitingQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalAwaitingQuantity() int32`

GetTotalAwaitingQuantity returns the TotalAwaitingQuantity field if non-nil, zero value otherwise.

### GetTotalAwaitingQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalAwaitingQuantityOk() (*int32, bool)`

GetTotalAwaitingQuantityOk returns a tuple with the TotalAwaitingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAwaitingQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalAwaitingQuantity(v int32)`

SetTotalAwaitingQuantity sets TotalAwaitingQuantity field to given value.

### HasTotalAwaitingQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalAwaitingQuantity() bool`

HasTotalAwaitingQuantity returns a boolean if a field has been set.

### GetTotalBackorderedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalBackorderedQuantity() int32`

GetTotalBackorderedQuantity returns the TotalBackorderedQuantity field if non-nil, zero value otherwise.

### GetTotalBackorderedQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalBackorderedQuantityOk() (*int32, bool)`

GetTotalBackorderedQuantityOk returns a tuple with the TotalBackorderedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackorderedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalBackorderedQuantity(v int32)`

SetTotalBackorderedQuantity sets TotalBackorderedQuantity field to given value.

### HasTotalBackorderedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalBackorderedQuantity() bool`

HasTotalBackorderedQuantity returns a boolean if a field has been set.

### GetTotalCommittedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalCommittedQuantity() int32`

GetTotalCommittedQuantity returns the TotalCommittedQuantity field if non-nil, zero value otherwise.

### GetTotalCommittedQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalCommittedQuantityOk() (*int32, bool)`

GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommittedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalCommittedQuantity(v int32)`

SetTotalCommittedQuantity sets TotalCommittedQuantity field to given value.

### HasTotalCommittedQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalCommittedQuantity() bool`

HasTotalCommittedQuantity returns a boolean if a field has been set.

### GetTotalExceptionQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalExceptionQuantity() int32`

GetTotalExceptionQuantity returns the TotalExceptionQuantity field if non-nil, zero value otherwise.

### GetTotalExceptionQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalExceptionQuantityOk() (*int32, bool)`

GetTotalExceptionQuantityOk returns a tuple with the TotalExceptionQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExceptionQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalExceptionQuantity(v int32)`

SetTotalExceptionQuantity sets TotalExceptionQuantity field to given value.

### HasTotalExceptionQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalExceptionQuantity() bool`

HasTotalExceptionQuantity returns a boolean if a field has been set.

### GetTotalFulfillableQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalFulfillableQuantity() int32`

GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalFulfillableQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalFulfillableQuantityOk() (*int32, bool)`

GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFulfillableQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalFulfillableQuantity(v int32)`

SetTotalFulfillableQuantity sets TotalFulfillableQuantity field to given value.

### HasTotalFulfillableQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalFulfillableQuantity() bool`

HasTotalFulfillableQuantity returns a boolean if a field has been set.

### GetTotalInternalTransferQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalInternalTransferQuantity() int32`

GetTotalInternalTransferQuantity returns the TotalInternalTransferQuantity field if non-nil, zero value otherwise.

### GetTotalInternalTransferQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalInternalTransferQuantityOk() (*int32, bool)`

GetTotalInternalTransferQuantityOk returns a tuple with the TotalInternalTransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInternalTransferQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalInternalTransferQuantity(v int32)`

SetTotalInternalTransferQuantity sets TotalInternalTransferQuantity field to given value.

### HasTotalInternalTransferQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalInternalTransferQuantity() bool`

HasTotalInternalTransferQuantity returns a boolean if a field has been set.

### GetTotalOnhandQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalOnhandQuantity() int32`

GetTotalOnhandQuantity returns the TotalOnhandQuantity field if non-nil, zero value otherwise.

### GetTotalOnhandQuantityOk

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) GetTotalOnhandQuantityOk() (*int32, bool)`

GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnhandQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) SetTotalOnhandQuantity(v int32)`

SetTotalOnhandQuantity sets TotalOnhandQuantity field to given value.

### HasTotalOnhandQuantity

`func (o *ShipbobInventoryApiViewModelsInventoryViewModel) HasTotalOnhandQuantity() bool`

HasTotalOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


