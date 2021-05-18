# OrdersCreateOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | [**[]OrdersAddProductToOrderModel**](OrdersAddProductToOrderModel.md) |  | 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | [**OrdersRecipientInfoViewModel**](Orders.RecipientInfoViewModel.md) |  | 
**ReferenceId** | **string** |  | 
**ShippingMethod** | **string** |  | 
**Tags** | Pointer to [**[]OrdersTagViewModel**](OrdersTagViewModel.md) |  | [optional] 
**Type** | Pointer to **string** | Defaults to Direct to Consumer (DTC) if not provided. Note: B2B is not supported at this time | [optional] 

## Methods

### NewOrdersCreateOrderModel

`func NewOrdersCreateOrderModel(products []OrdersAddProductToOrderModel, recipient OrdersRecipientInfoViewModel, referenceId string, shippingMethod string, ) *OrdersCreateOrderModel`

NewOrdersCreateOrderModel instantiates a new OrdersCreateOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersCreateOrderModelWithDefaults

`func NewOrdersCreateOrderModelWithDefaults() *OrdersCreateOrderModel`

NewOrdersCreateOrderModelWithDefaults instantiates a new OrdersCreateOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *OrdersCreateOrderModel) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrdersCreateOrderModel) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrdersCreateOrderModel) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *OrdersCreateOrderModel) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *OrdersCreateOrderModel) GetProducts() []OrdersAddProductToOrderModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrdersCreateOrderModel) GetProductsOk() (*[]OrdersAddProductToOrderModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrdersCreateOrderModel) SetProducts(v []OrdersAddProductToOrderModel)`

SetProducts sets Products field to given value.


### GetPurchaseDate

`func (o *OrdersCreateOrderModel) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *OrdersCreateOrderModel) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *OrdersCreateOrderModel) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *OrdersCreateOrderModel) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *OrdersCreateOrderModel) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *OrdersCreateOrderModel) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *OrdersCreateOrderModel) GetRecipient() OrdersRecipientInfoViewModel`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *OrdersCreateOrderModel) GetRecipientOk() (*OrdersRecipientInfoViewModel, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *OrdersCreateOrderModel) SetRecipient(v OrdersRecipientInfoViewModel)`

SetRecipient sets Recipient field to given value.


### GetReferenceId

`func (o *OrdersCreateOrderModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrdersCreateOrderModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrdersCreateOrderModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetShippingMethod

`func (o *OrdersCreateOrderModel) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *OrdersCreateOrderModel) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *OrdersCreateOrderModel) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.


### GetTags

`func (o *OrdersCreateOrderModel) GetTags() []OrdersTagViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrdersCreateOrderModel) GetTagsOk() (*[]OrdersTagViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrdersCreateOrderModel) SetTags(v []OrdersTagViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrdersCreateOrderModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *OrdersCreateOrderModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrdersCreateOrderModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrdersCreateOrderModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrdersCreateOrderModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


