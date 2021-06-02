# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ChannelInfo**](ChannelInfo.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date this order was created | [optional] 
**Id** | Pointer to **int32** | Unique id of the order | [optional] 
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | Pointer to [**[]ProductInfo**](ProductInfo.md) | List of products included in the order | [optional] 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | Pointer to [**RecipientInfo**](RecipientInfo.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order | [optional] 
**Shipments** | Pointer to [**[]Shipment**](Shipment.md) | Shipments affiliated with the order | [optional] 
**ShippingMethod** | Pointer to **string** | Client-defined shipping method | [optional] 
**Status** | Pointer to **string** | The order status | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Client-defined order tags | [optional] 
**Type** | Pointer to **string** | Shipment type of the order | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Order) GetChannel() ChannelInfo`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Order) GetChannelOk() (*ChannelInfo, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Order) SetChannel(v ChannelInfo)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Order) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Order) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Order) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Order) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Order) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *Order) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *Order) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *Order) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *Order) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *Order) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *Order) GetProducts() []ProductInfo`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Order) GetProductsOk() (*[]ProductInfo, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Order) SetProducts(v []ProductInfo)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Order) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *Order) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *Order) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *Order) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *Order) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *Order) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *Order) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *Order) GetRecipient() RecipientInfo`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Order) GetRecipientOk() (*RecipientInfo, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Order) SetRecipient(v RecipientInfo)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Order) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *Order) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Order) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Order) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Order) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetShipments

`func (o *Order) GetShipments() []Shipment`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *Order) GetShipmentsOk() (*[]Shipment, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *Order) SetShipments(v []Shipment)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *Order) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetShippingMethod

`func (o *Order) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *Order) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *Order) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *Order) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Order) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Order) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Order) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Order) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Order) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Order) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Order) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Order) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


