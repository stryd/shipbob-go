# CreateOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | [**[]CreateOrderProductsInner**](CreateOrderProductsInner.md) |  | 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | [**RecipientInfo**](RecipientInfo.md) |  | 
**ReferenceId** | **string** |  | 
**ShippingMethod** | **string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Type** | Pointer to **string** | Defaults to Direct to Consumer (DTC) if not provided. Note: B2B is not supported at this time | [optional] 

## Methods

### NewCreateOrder

`func NewCreateOrder(products []CreateOrderProductsInner, recipient RecipientInfo, referenceId string, shippingMethod string, ) *CreateOrder`

NewCreateOrder instantiates a new CreateOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderWithDefaults

`func NewCreateOrderWithDefaults() *CreateOrder`

NewCreateOrderWithDefaults instantiates a new CreateOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *CreateOrder) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *CreateOrder) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *CreateOrder) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *CreateOrder) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *CreateOrder) GetProducts() []CreateOrderProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateOrder) GetProductsOk() (*[]CreateOrderProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateOrder) SetProducts(v []CreateOrderProductsInner)`

SetProducts sets Products field to given value.


### GetPurchaseDate

`func (o *CreateOrder) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *CreateOrder) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *CreateOrder) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *CreateOrder) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *CreateOrder) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *CreateOrder) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *CreateOrder) GetRecipient() RecipientInfo`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CreateOrder) GetRecipientOk() (*RecipientInfo, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CreateOrder) SetRecipient(v RecipientInfo)`

SetRecipient sets Recipient field to given value.


### GetReferenceId

`func (o *CreateOrder) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateOrder) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateOrder) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetShippingMethod

`func (o *CreateOrder) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *CreateOrder) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *CreateOrder) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.


### GetTags

`func (o *CreateOrder) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOrder) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOrder) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOrder) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CreateOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOrder) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


