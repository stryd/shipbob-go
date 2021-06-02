# ReturnOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ReturnChannelInfo**](ReturnChannelInfo.md) |  | [optional] 
**FulfillmentCenter** | Pointer to [**ReturnFulfillmentCenter**](ReturnFulfillmentCenter.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the Return Order | [optional] 
**InsertDate** | Pointer to **time.Time** | Date this return order was created | [optional] 
**Inventory** | Pointer to [**[]ReturnInventoryItem**](ReturnInventoryItem.md) | List of inventory included in the return order | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Invoiced amount of return order (sum of transaction amounts) | [optional] 
**OriginalShipmentId** | Pointer to **NullableInt32** | Id of the corresponding shipment that is the souce of the return | [optional] 
**ReferenceId** | Pointer to **NullableString** | Client-defined external unique id of the return order | [optional] 
**ReturnType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ReturnStatus**](ReturnStatus.md) |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** | Tracking number of the return shipment | [optional] 
**Transactions** | Pointer to [**[]ReturnTransaction**](ReturnTransaction.md) | Array of transactions affiliated with the return order | [optional] 

## Methods

### NewReturnOrder

`func NewReturnOrder() *ReturnOrder`

NewReturnOrder instantiates a new ReturnOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnOrderWithDefaults

`func NewReturnOrderWithDefaults() *ReturnOrder`

NewReturnOrderWithDefaults instantiates a new ReturnOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ReturnOrder) GetChannel() ReturnChannelInfo`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ReturnOrder) GetChannelOk() (*ReturnChannelInfo, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ReturnOrder) SetChannel(v ReturnChannelInfo)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ReturnOrder) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ReturnOrder) GetFulfillmentCenter() ReturnFulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ReturnOrder) GetFulfillmentCenterOk() (*ReturnFulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ReturnOrder) SetFulfillmentCenter(v ReturnFulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ReturnOrder) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ReturnOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ReturnOrder) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ReturnOrder) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ReturnOrder) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ReturnOrder) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetInventory

`func (o *ReturnOrder) GetInventory() []ReturnInventoryItem`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ReturnOrder) GetInventoryOk() (*[]ReturnInventoryItem, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ReturnOrder) SetInventory(v []ReturnInventoryItem)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ReturnOrder) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### SetInventoryNil

`func (o *ReturnOrder) SetInventoryNil(b bool)`

 SetInventoryNil sets the value for Inventory to be an explicit nil

### UnsetInventory
`func (o *ReturnOrder) UnsetInventory()`

UnsetInventory ensures that no value is present for Inventory, not even an explicit nil
### GetInvoiceAmount

`func (o *ReturnOrder) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *ReturnOrder) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *ReturnOrder) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *ReturnOrder) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *ReturnOrder) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *ReturnOrder) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetOriginalShipmentId

`func (o *ReturnOrder) GetOriginalShipmentId() int32`

GetOriginalShipmentId returns the OriginalShipmentId field if non-nil, zero value otherwise.

### GetOriginalShipmentIdOk

`func (o *ReturnOrder) GetOriginalShipmentIdOk() (*int32, bool)`

GetOriginalShipmentIdOk returns a tuple with the OriginalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalShipmentId

`func (o *ReturnOrder) SetOriginalShipmentId(v int32)`

SetOriginalShipmentId sets OriginalShipmentId field to given value.

### HasOriginalShipmentId

`func (o *ReturnOrder) HasOriginalShipmentId() bool`

HasOriginalShipmentId returns a boolean if a field has been set.

### SetOriginalShipmentIdNil

`func (o *ReturnOrder) SetOriginalShipmentIdNil(b bool)`

 SetOriginalShipmentIdNil sets the value for OriginalShipmentId to be an explicit nil

### UnsetOriginalShipmentId
`func (o *ReturnOrder) UnsetOriginalShipmentId()`

UnsetOriginalShipmentId ensures that no value is present for OriginalShipmentId, not even an explicit nil
### GetReferenceId

`func (o *ReturnOrder) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ReturnOrder) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ReturnOrder) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ReturnOrder) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ReturnOrder) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ReturnOrder) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetReturnType

`func (o *ReturnOrder) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *ReturnOrder) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *ReturnOrder) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *ReturnOrder) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnOrder) GetStatus() ReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnOrder) GetStatusOk() (*ReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnOrder) SetStatus(v ReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ReturnOrder) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ReturnOrder) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ReturnOrder) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ReturnOrder) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ReturnOrder) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ReturnOrder) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTransactions

`func (o *ReturnOrder) GetTransactions() []ReturnTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ReturnOrder) GetTransactionsOk() (*[]ReturnTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ReturnOrder) SetTransactions(v []ReturnTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ReturnOrder) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *ReturnOrder) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *ReturnOrder) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


