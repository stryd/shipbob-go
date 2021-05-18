# OrdersOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**OrdersChannelInfoViewModel**](Orders.ChannelInfoViewModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date this order was created | [optional] 
**Id** | Pointer to **int32** | Unique id of the order | [optional] 
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | Pointer to [**[]OrdersProductInfoViewModel**](OrdersProductInfoViewModel.md) | List of products included in the order | [optional] 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | Pointer to [**OrdersRecipientInfoViewModel**](Orders.RecipientInfoViewModel.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order | [optional] 
**Shipments** | Pointer to [**[]OrdersShipmentViewModel**](OrdersShipmentViewModel.md) | Shipments affiliated with the order | [optional] 
**ShippingMethod** | Pointer to **string** | Client-defined shipping method | [optional] 
**Status** | Pointer to **string** | The order status | [optional] 
**Tags** | Pointer to [**[]OrdersTagViewModel**](OrdersTagViewModel.md) | Client-defined order tags | [optional] 
**Type** | Pointer to **string** | Shipment type of the order | [optional] 

## Methods

### NewOrdersOrderViewModel

`func NewOrdersOrderViewModel() *OrdersOrderViewModel`

NewOrdersOrderViewModel instantiates a new OrdersOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersOrderViewModelWithDefaults

`func NewOrdersOrderViewModelWithDefaults() *OrdersOrderViewModel`

NewOrdersOrderViewModelWithDefaults instantiates a new OrdersOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *OrdersOrderViewModel) GetChannel() OrdersChannelInfoViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OrdersOrderViewModel) GetChannelOk() (*OrdersChannelInfoViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OrdersOrderViewModel) SetChannel(v OrdersChannelInfoViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *OrdersOrderViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *OrdersOrderViewModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *OrdersOrderViewModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *OrdersOrderViewModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *OrdersOrderViewModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *OrdersOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *OrdersOrderViewModel) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrdersOrderViewModel) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrdersOrderViewModel) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *OrdersOrderViewModel) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *OrdersOrderViewModel) GetProducts() []OrdersProductInfoViewModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrdersOrderViewModel) GetProductsOk() (*[]OrdersProductInfoViewModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrdersOrderViewModel) SetProducts(v []OrdersProductInfoViewModel)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OrdersOrderViewModel) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *OrdersOrderViewModel) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *OrdersOrderViewModel) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *OrdersOrderViewModel) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *OrdersOrderViewModel) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *OrdersOrderViewModel) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *OrdersOrderViewModel) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *OrdersOrderViewModel) GetRecipient() OrdersRecipientInfoViewModel`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *OrdersOrderViewModel) GetRecipientOk() (*OrdersRecipientInfoViewModel, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *OrdersOrderViewModel) SetRecipient(v OrdersRecipientInfoViewModel)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *OrdersOrderViewModel) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *OrdersOrderViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrdersOrderViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrdersOrderViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *OrdersOrderViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetShipments

`func (o *OrdersOrderViewModel) GetShipments() []OrdersShipmentViewModel`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *OrdersOrderViewModel) GetShipmentsOk() (*[]OrdersShipmentViewModel, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *OrdersOrderViewModel) SetShipments(v []OrdersShipmentViewModel)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *OrdersOrderViewModel) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetShippingMethod

`func (o *OrdersOrderViewModel) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *OrdersOrderViewModel) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *OrdersOrderViewModel) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *OrdersOrderViewModel) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetStatus

`func (o *OrdersOrderViewModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrdersOrderViewModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrdersOrderViewModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrdersOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *OrdersOrderViewModel) GetTags() []OrdersTagViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrdersOrderViewModel) GetTagsOk() (*[]OrdersTagViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrdersOrderViewModel) SetTags(v []OrdersTagViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrdersOrderViewModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *OrdersOrderViewModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrdersOrderViewModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrdersOrderViewModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrdersOrderViewModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


