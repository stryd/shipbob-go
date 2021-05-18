# ShipBobOrdersPresentationViewModelsOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ShipBobOrdersPresentationViewModelsChannelInfoViewModel**](ShipBob.Orders.Presentation.ViewModels.ChannelInfoViewModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date this order was created | [optional] 
**Id** | Pointer to **int32** | Unique id of the order | [optional] 
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | Pointer to [**[]ShipBobOrdersPresentationViewModelsProductInfoViewModel**](ShipBobOrdersPresentationViewModelsProductInfoViewModel.md) | List of products included in the order | [optional] 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | Pointer to [**ShipBobOrdersPresentationViewModelsRecipientInfoViewModel**](ShipBob.Orders.Presentation.ViewModels.RecipientInfoViewModel.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order | [optional] 
**Shipments** | Pointer to [**[]ShipBobOrdersPresentationViewModelsShipmentViewModel**](ShipBobOrdersPresentationViewModelsShipmentViewModel.md) | Shipments affiliated with the order | [optional] 
**ShippingMethod** | Pointer to **string** | Client-defined shipping method | [optional] 
**Status** | Pointer to **string** | The order status | [optional] 
**Tags** | Pointer to [**[]ShipBobOrdersPresentationViewModelsTagViewModel**](ShipBobOrdersPresentationViewModelsTagViewModel.md) | Client-defined order tags | [optional] 
**Type** | Pointer to **string** | Shipment type of the order | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsOrderViewModel

`func NewShipBobOrdersPresentationViewModelsOrderViewModel() *ShipBobOrdersPresentationViewModelsOrderViewModel`

NewShipBobOrdersPresentationViewModelsOrderViewModel instantiates a new ShipBobOrdersPresentationViewModelsOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsOrderViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsOrderViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsOrderViewModel`

NewShipBobOrdersPresentationViewModelsOrderViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetChannel() ShipBobOrdersPresentationViewModelsChannelInfoViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetChannelOk() (*ShipBobOrdersPresentationViewModelsChannelInfoViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetChannel(v ShipBobOrdersPresentationViewModelsChannelInfoViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetProducts() []ShipBobOrdersPresentationViewModelsProductInfoViewModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetProductsOk() (*[]ShipBobOrdersPresentationViewModelsProductInfoViewModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetProducts(v []ShipBobOrdersPresentationViewModelsProductInfoViewModel)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetRecipient() ShipBobOrdersPresentationViewModelsRecipientInfoViewModel`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetRecipientOk() (*ShipBobOrdersPresentationViewModelsRecipientInfoViewModel, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetRecipient(v ShipBobOrdersPresentationViewModelsRecipientInfoViewModel)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetShipments

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetShipments() []ShipBobOrdersPresentationViewModelsShipmentViewModel`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetShipmentsOk() (*[]ShipBobOrdersPresentationViewModelsShipmentViewModel, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetShipments(v []ShipBobOrdersPresentationViewModelsShipmentViewModel)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetStatus

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetTags() []ShipBobOrdersPresentationViewModelsTagViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetTagsOk() (*[]ShipBobOrdersPresentationViewModelsTagViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetTags(v []ShipBobOrdersPresentationViewModelsTagViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShipBobOrdersPresentationViewModelsOrderViewModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


