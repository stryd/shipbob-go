# CreateReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentCenter** | [**ReturnFulfillmentCenter**](ReturnFulfillmentCenter.md) |  | 
**Inventory** | [**[]ReturnInventory**](ReturnInventory.md) | Array of inventory items being returned | 
**OriginalShipmentId** | Pointer to **NullableInt32** | Shipment from which the items in the return originated | [optional] 
**ReferenceId** | **NullableString** | Client-defined external unique identifier for the return order | 
**TrackingNumber** | Pointer to **NullableString** | Tracking number for the return shipment | [optional] 

## Methods

### NewCreateReturn

`func NewCreateReturn(fulfillmentCenter ReturnFulfillmentCenter, inventory []ReturnInventory, referenceId NullableString, ) *CreateReturn`

NewCreateReturn instantiates a new CreateReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnWithDefaults

`func NewCreateReturnWithDefaults() *CreateReturn`

NewCreateReturnWithDefaults instantiates a new CreateReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentCenter

`func (o *CreateReturn) GetFulfillmentCenter() ReturnFulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *CreateReturn) GetFulfillmentCenterOk() (*ReturnFulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *CreateReturn) SetFulfillmentCenter(v ReturnFulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.


### GetInventory

`func (o *CreateReturn) GetInventory() []ReturnInventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CreateReturn) GetInventoryOk() (*[]ReturnInventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CreateReturn) SetInventory(v []ReturnInventory)`

SetInventory sets Inventory field to given value.


### SetInventoryNil

`func (o *CreateReturn) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *CreateReturn) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetOriginalShipmentId

`func (o *CreateReturn) GetOriginalShipmentId() int32`

GetOriginalShipmentId returns the OriginalShipmentId field if non-nil, zero value otherwise.

### GetOriginalShipmentIdOk

`func (o *CreateReturn) GetOriginalShipmentIdOk() (*int32, bool)`

GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShipmentId

`func (o *CreateReturn) SetOriginalShipmentId(v int32)`

SetOriginalShipmentId sets OriginalShipmentId field to given value.

### HasOriginalShipmentId

`func (o *CreateReturn) HasOriginalShipmentId() bool`

HasOriginalShipmentId returns a boolean if a field has been set.

### SetOriginalShipmentIdNil

`func (o *CreateReturn) SetOriginalShipmentIdNil(b bool)`

 SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil

### UnsetOriginalShipmentId
`func (o *CreateReturn) UnsetOriginalShipmentId()`

UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
### GetReferenceId

`func (o *CreateReturn) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateReturn) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateReturn) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *CreateReturn) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *CreateReturn) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetTrackingNumber

`func (o *CreateReturn) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *CreateReturn) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *CreateReturn) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *CreateReturn) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *CreateReturn) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *CreateReturn) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


