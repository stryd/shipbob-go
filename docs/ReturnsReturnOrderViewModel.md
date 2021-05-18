# ReturnsReturnOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ReturnsChannelInfoViewModel**](Returns.ChannelInfoViewModel.md) |  | [optional] 
**FulfillmentCenter** | Pointer to [**ReturnsFulfillmentCenterViewModel**](Returns.FulfillmentCenterViewModel.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the Return Order | [optional] 
**InsertDate** | Pointer to **time.Time** | Date this return order was created | [optional] 
**Inventory** | Pointer to [**[]ReturnsInventoryItemViewModel**](ReturnsInventoryItemViewModel.md) | List of inventory included in the return order | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Invoiced amount of return order (sum of transaction amounts) | [optional] 
**OriginalShipmentId** | Pointer to **NullableInt32** | Id of the corresponding shipment that is the souce of the return | [optional] 
**ReferenceId** | Pointer to **NullableString** | Client-defined external unique id of the return order | [optional] 
**ReturnType** | Pointer to [**ReturnsReturnType**](Returns.ReturnType.md) |  | [optional] 
**Status** | Pointer to [**ReturnsReturnStatus**](Returns.ReturnStatus.md) |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the return shipment | [optional] 
**Transactions** | Pointer to [**[]ReturnsTransactionViewModel**](ReturnsTransactionViewModel.md) | Array of transactions affiliated with the return order | [optional] 

## Methods

### NewReturnsReturnOrderViewModel

`func NewReturnsReturnOrderViewModel() *ReturnsReturnOrderViewModel`

NewReturnsReturnOrderViewModel instantiates a new ReturnsReturnOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsReturnOrderViewModelWithDefaults

`func NewReturnsReturnOrderViewModelWithDefaults() *ReturnsReturnOrderViewModel`

NewReturnsReturnOrderViewModelWithDefaults instantiates a new ReturnsReturnOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ReturnsReturnOrderViewModel) GetChannel() ReturnsChannelInfoViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ReturnsReturnOrderViewModel) GetChannelOk() (*ReturnsChannelInfoViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ReturnsReturnOrderViewModel) SetChannel(v ReturnsChannelInfoViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ReturnsReturnOrderViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ReturnsReturnOrderViewModel) GetFulfillmentCenter() ReturnsFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ReturnsReturnOrderViewModel) GetFulfillmentCenterOk() (*ReturnsFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ReturnsReturnOrderViewModel) SetFulfillmentCenter(v ReturnsFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ReturnsReturnOrderViewModel) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ReturnsReturnOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnsReturnOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnsReturnOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnsReturnOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ReturnsReturnOrderViewModel) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ReturnsReturnOrderViewModel) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ReturnsReturnOrderViewModel) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ReturnsReturnOrderViewModel) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetInventory

`func (o *ReturnsReturnOrderViewModel) GetInventory() []ReturnsInventoryItemViewModel`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ReturnsReturnOrderViewModel) GetInventoryOk() (*[]ReturnsInventoryItemViewModel, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ReturnsReturnOrderViewModel) SetInventory(v []ReturnsInventoryItemViewModel)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ReturnsReturnOrderViewModel) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *ReturnsReturnOrderViewModel) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *ReturnsReturnOrderViewModel) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetInvoiceAmount

`func (o *ReturnsReturnOrderViewModel) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *ReturnsReturnOrderViewModel) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *ReturnsReturnOrderViewModel) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *ReturnsReturnOrderViewModel) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *ReturnsReturnOrderViewModel) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *ReturnsReturnOrderViewModel) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetOriginalShipmentId

`func (o *ReturnsReturnOrderViewModel) GetOriginalShipmentId() int32`

GetOriginalShipmentId returns the OriginalShipmentId field if non-nil, zero value otherwise.

### GetOriginalShipmentIdOk

`func (o *ReturnsReturnOrderViewModel) GetOriginalShipmentIdOk() (*int32, bool)`

GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShipmentId

`func (o *ReturnsReturnOrderViewModel) SetOriginalShipmentId(v int32)`

SetOriginalShipmentId sets OriginalShipmentId field to given value.

### HasOriginalShipmentId

`func (o *ReturnsReturnOrderViewModel) HasOriginalShipmentId() bool`

HasOriginalShipmentId returns a boolean if a field has been set.

### SetOriginalShipmentIdNil

`func (o *ReturnsReturnOrderViewModel) SetOriginalShipmentIdNil(b bool)`

 SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil

### UnsetOriginalShipmentId
`func (o *ReturnsReturnOrderViewModel) UnsetOriginalShipmentId()`

UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
### GetReferenceId

`func (o *ReturnsReturnOrderViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ReturnsReturnOrderViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ReturnsReturnOrderViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ReturnsReturnOrderViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ReturnsReturnOrderViewModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ReturnsReturnOrderViewModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetReturnType

`func (o *ReturnsReturnOrderViewModel) GetReturnType() ReturnsReturnType`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *ReturnsReturnOrderViewModel) GetReturnTypeOk() (*ReturnsReturnType, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *ReturnsReturnOrderViewModel) SetReturnType(v ReturnsReturnType)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *ReturnsReturnOrderViewModel) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnsReturnOrderViewModel) GetStatus() ReturnsReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnsReturnOrderViewModel) GetStatusOk() (*ReturnsReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnsReturnOrderViewModel) SetStatus(v ReturnsReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnsReturnOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ReturnsReturnOrderViewModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReturnsReturnOrderViewModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReturnsReturnOrderViewModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ReturnsReturnOrderViewModel) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ReturnsReturnOrderViewModel) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReturnsReturnOrderViewModel) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTransactions

`func (o *ReturnsReturnOrderViewModel) GetTransactions() []ReturnsTransactionViewModel`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ReturnsReturnOrderViewModel) GetTransactionsOk() (*[]ReturnsTransactionViewModel, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ReturnsReturnOrderViewModel) SetTransactions(v []ReturnsTransactionViewModel)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ReturnsReturnOrderViewModel) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *ReturnsReturnOrderViewModel) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *ReturnsReturnOrderViewModel) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


