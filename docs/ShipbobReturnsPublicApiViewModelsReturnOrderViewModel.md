# ShipbobReturnsPublicApiViewModelsReturnOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ShipbobReturnsPublicApiViewModelsChannelInfoViewModel**](Shipbob.Returns.Public.Api.ViewModels.ChannelInfoViewModel.md) |  | [optional] 
**FulfillmentCenter** | Pointer to [**ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel**](Shipbob.Returns.Public.Api.ViewModels.FulfillmentCenterViewModel.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the Return Order | [optional] 
**InsertDate** | Pointer to **time.Time** | Date this return order was created | [optional] 
**Inventory** | Pointer to [**[]ShipbobReturnsPublicApiViewModelsInventoryItemViewModel**](ShipbobReturnsPublicApiViewModelsInventoryItemViewModel.md) | List of inventory included in the return order | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Invoiced amount of return order (sum of transaction amounts) | [optional] 
**OriginalShipmentId** | Pointer to **NullableInt32** | Id of the corresponding shipment that is the souce of the return | [optional] 
**ReferenceId** | Pointer to **NullableString** | Client-defined external unique id of the return order | [optional] 
**ReturnType** | Pointer to [**ShipbobReturnsPublicCommonReturnType**](Shipbob.Returns.Public.Common.ReturnType.md) |  | [optional] 
**Status** | Pointer to [**ShipbobReturnsPublicCommonReturnStatus**](Shipbob.Returns.Public.Common.ReturnStatus.md) |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the return shipment | [optional] 
**Transactions** | Pointer to [**[]ShipbobReturnsPublicApiViewModelsTransactionViewModel**](ShipbobReturnsPublicApiViewModelsTransactionViewModel.md) | Array of transactions affiliated with the return order | [optional] 

## Methods

### NewShipbobReturnsPublicApiViewModelsReturnOrderViewModel

`func NewShipbobReturnsPublicApiViewModelsReturnOrderViewModel() *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel`

NewShipbobReturnsPublicApiViewModelsReturnOrderViewModel instantiates a new ShipbobReturnsPublicApiViewModelsReturnOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReturnsPublicApiViewModelsReturnOrderViewModelWithDefaults

`func NewShipbobReturnsPublicApiViewModelsReturnOrderViewModelWithDefaults() *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel`

NewShipbobReturnsPublicApiViewModelsReturnOrderViewModelWithDefaults instantiates a new ShipbobReturnsPublicApiViewModelsReturnOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetChannel() ShipbobReturnsPublicApiViewModelsChannelInfoViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetChannelOk() (*ShipbobReturnsPublicApiViewModelsChannelInfoViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetChannel(v ShipbobReturnsPublicApiViewModelsChannelInfoViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetFulfillmentCenter() ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetFulfillmentCenterOk() (*ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetFulfillmentCenter(v ShipbobReturnsPublicApiViewModelsFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetInventory

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInventory() []ShipbobReturnsPublicApiViewModelsInventoryItemViewModel`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInventoryOk() (*[]ShipbobReturnsPublicApiViewModelsInventoryItemViewModel, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetInventory(v []ShipbobReturnsPublicApiViewModelsInventoryItemViewModel)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetInvoiceAmount

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetOriginalShipmentId() int32`

GetOriginalShipmentId returns the OriginalShipmentId field if non-nil, zero value otherwise.

### GetOriginalShipmentIdOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetOriginalShipmentIdOk() (*int32, bool)`

GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetOriginalShipmentId(v int32)`

SetOriginalShipmentId sets OriginalShipmentId field to given value.

### HasOriginalShipmentId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasOriginalShipmentId() bool`

HasOriginalShipmentId returns a boolean if a field has been set.

### SetOriginalShipmentIdNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetOriginalShipmentIdNil(b bool)`

 SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil

### UnsetOriginalShipmentId
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetOriginalShipmentId()`

UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
### GetReferenceId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetReturnType

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetReturnType() ShipbobReturnsPublicCommonReturnType`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetReturnTypeOk() (*ShipbobReturnsPublicCommonReturnType, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetReturnType(v ShipbobReturnsPublicCommonReturnType)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetStatus() ShipbobReturnsPublicCommonReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetStatusOk() (*ShipbobReturnsPublicCommonReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetStatus(v ShipbobReturnsPublicCommonReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTransactions

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetTransactions() []ShipbobReturnsPublicApiViewModelsTransactionViewModel`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) GetTransactionsOk() (*[]ShipbobReturnsPublicApiViewModelsTransactionViewModel, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetTransactions(v []ShipbobReturnsPublicApiViewModelsTransactionViewModel)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *ShipbobReturnsPublicApiViewModelsReturnOrderViewModel) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


