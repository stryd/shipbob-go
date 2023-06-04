# CreateOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Financials** | Pointer to [**Financials**](Financials.md) |  | [optional] 
**GiftMessage** | Pointer to **string** | Gift message associated with the order | [optional] 
**LocationId** | Pointer to **NullableInt32** | Desired Fulfillment Center Location ID. If not specified, ShipBob will determine the location that fulfills this order. | [optional] 
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | [**[]AddProductToOrder**](AddProductToOrder.md) | Products included in the order. Products identified by reference_id must also include the product name if there is no matching ShipBob product. | 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | [**RecipientInfo**](RecipientInfo.md) |  | 
**ReferenceId** | **string** | Unique and immutable order identifier from your upstream system | 
**RetailerProgramData** | Pointer to [**RetailerProgramData**](RetailerProgramData.md) |  | [optional] 
**ShippingMethod** | **string** | Client-defined shipping method matching what the user has listed as the shipping method on the Ship Option Mapping setup page in the ShipBob Merchant Portal. If they don&#39;t match, we will create a new one and default it to Standard | 
**ShippingTerms** | Pointer to [**ShippingTerms**](ShippingTerms.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Key value pair array to store extra information at the order level for API purposes. ShipBob won&#39;t display the info in the ShipBob Merchant Portal or react based on this data. | [optional] 
**Type** | Pointer to **string** | Defaults to Direct to Consumer (DTC) if not provided. Note: B2B is not supported at this time | [optional] 

## Methods

### NewCreateOrder

`func NewCreateOrder(products []AddProductToOrder, recipient RecipientInfo, referenceId string, shippingMethod string, ) *CreateOrder`

NewCreateOrder instantiates a new CreateOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderWithDefaults

`func NewCreateOrderWithDefaults() *CreateOrder`

NewCreateOrderWithDefaults instantiates a new CreateOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinancials

`func (o *CreateOrder) GetFinancials() Financials`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *CreateOrder) GetFinancialsOk() (*Financials, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *CreateOrder) SetFinancials(v Financials)`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *CreateOrder) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.

### GetGiftMessage

`func (o *CreateOrder) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *CreateOrder) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *CreateOrder) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *CreateOrder) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetLocationId

`func (o *CreateOrder) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CreateOrder) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CreateOrder) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *CreateOrder) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *CreateOrder) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *CreateOrder) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
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

`func (o *CreateOrder) GetProducts() []AddProductToOrder`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateOrder) GetProductsOk() (*[]AddProductToOrder, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateOrder) SetProducts(v []AddProductToOrder)`

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


### GetRetailerProgramData

`func (o *CreateOrder) GetRetailerProgramData() RetailerProgramData`

GetRetailerProgramData returns the RetailerProgramData field if non-nil, zero value otherwise.

### GetRetailerProgramDataOk

`func (o *CreateOrder) GetRetailerProgramDataOk() (*RetailerProgramData, bool)`

GetRetailerProgramDataOk returns a tuple with the RetailerProgramData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailerProgramData

`func (o *CreateOrder) SetRetailerProgramData(v RetailerProgramData)`

SetRetailerProgramData sets RetailerProgramData field to given value.

### HasRetailerProgramData

`func (o *CreateOrder) HasRetailerProgramData() bool`

HasRetailerProgramData returns a boolean if a field has been set.

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


### GetShippingTerms

`func (o *CreateOrder) GetShippingTerms() ShippingTerms`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *CreateOrder) GetShippingTermsOk() (*ShippingTerms, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *CreateOrder) SetShippingTerms(v ShippingTerms)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *CreateOrder) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

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


