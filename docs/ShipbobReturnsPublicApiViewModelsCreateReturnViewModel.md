# ShipbobReturnsPublicApiViewModelsCreateReturnViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentCenter** | [**ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel**](Shipbob.Returns.Public.Api.ViewModels.FulfillmentCenterViewModel.md) |  | 
**Inventory** | [**[]ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel**](ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel.md) | Array of inventory items being returned | 
**OriginalShipmentId** | Pointer to **NullableInt32** | Shipment from which the items in the return originated | [optional] 
**ReferenceId** | **NullableString** | Client-defined external unique identifier for the return order | 
**TrackingNumber** | Pointer to **NullableString** | Tracking number for the return shipment | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsCreateReturnViewModel

`func NewShipbobReturnsPublicApiViewModelsCreateReturnViewModel(fulfillmentCenter ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel, inventory []ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel, referenceId NullableString, ) *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel`

NewShipbobReturnsPublicApiViewModelsCreateReturnViewModel instantiates a new ShipbobReturnsPublicApiViewModelsCreateReturnViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsCreateReturnViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsCreateReturnViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel`

NewShipbobReturnsPublicApiViewModelsCreateReturnViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsCreateReturnViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentCenter

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetFulfillmentCenter() ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetFulfillmentCenterOk() (*ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetFulfillmentCenter(v ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.


### GetInventory

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetInventory() []ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetInventoryOk() (*[]ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetInventory(v []ShipbobReturnsPublicApiViewModelsReturnInventoryViewModel)`

SetInventory sets Inventory field to given value.


### SetInventoryNil

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetOriginalShipmentId() int32`

GetOriginalShipmentId returns the OriginalShipmentId field if non-nil, zero value otherwise.

### GetOriginalShipmentIdOk

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetOriginalShipmentIdOk() (*int32, bool)`

GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetOriginalShipmentId(v int32)`

SetOriginalShipmentId sets OriginalShipmentId field to given value.

### HasOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) HasOriginalShipmentId() bool`

HasOriginalShipmentId returns a boolean if a field has been set.

### SetOriginalShipmentIdNil

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetOriginalShipmentIdNil(b bool)`

 SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil

### UnsetOriginalShipmentId
`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) UnsetOriginalShipmentId()`

UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
### GetReferenceId

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ShipbobReturnsPublicApiViewModelsCreateReturnViewModel) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


