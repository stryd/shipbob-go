# ShipBobOrdersPresentationModelsCreateOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** | User friendly orderId or store order number that will be shown on the Orders Page. If not provided, referenceId will be used | [optional] 
**Products** | [**[]ShipBobOrdersPresentationModelsAddProductToOrderModel**](ShipBobOrdersPresentationModelsAddProductToOrderModel.md) |  | 
**PurchaseDate** | Pointer to **NullableTime** | Date this order was purchase by the end user | [optional] 
**Recipient** | [**ShipBobOrdersPresentationViewModelsRecipientInfoViewModel**](ShipBob.Orders.Presentation.ViewModels.RecipientInfoViewModel.md) |  | 
**ReferenceId** | **string** |  | 
**ShippingMethod** | **string** |  | 
**Tags** | Pointer to [**[]ShipBobOrdersPresentationViewModelsTagViewModel**](ShipBobOrdersPresentationViewModelsTagViewModel.md) |  | [optional] 
**Type** | Pointer to **string** | Defaults to Direct to Consumer (DTC) if not provided. Note: B2B is not supported at this time | [optional] 

## Methods

### NewShipBobOrdersPresentationModelsCreateOrderModel

`func NewShipBobOrdersPresentationModelsCreateOrderModel(products []ShipBobOrdersPresentationModelsAddProductToOrderModel, recipient ShipBobOrdersPresentationViewModelsRecipientInfoViewModel, referenceId string, shippingMethod string, ) *ShipBobOrdersPresentationModelsCreateOrderModel`

NewShipBobOrdersPresentationModelsCreateOrderModel instantiates a new ShipBobOrdersPresentationModelsCreateOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationModelsCreateOrderModelWithDefaults

`func NewShipBobOrdersPresentationModelsCreateOrderModelWithDefaults() *ShipBobOrdersPresentationModelsCreateOrderModel`

NewShipBobOrdersPresentationModelsCreateOrderModelWithDefaults instantiates a new ShipBobOrdersPresentationModelsCreateOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetProducts

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetProducts() []ShipBobOrdersPresentationModelsAddProductToOrderModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetProductsOk() (*[]ShipBobOrdersPresentationModelsAddProductToOrderModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetProducts(v []ShipBobOrdersPresentationModelsAddProductToOrderModel)`

SetProducts sets Products field to given value.


### GetPurchaseDate

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### SetPurchaseDateNil

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetPurchaseDateNil(b bool)`

 SetPurchaseDateNil sets the value for PurchaseDate to be an explicit nil

### UnsetPurchaseDate
`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) UnsetPurchaseDate()`

UnsetPurchaseDate ensures that no value is present for PurchaseDate, not even an explicit nil
### GetRecipient

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetRecipient() ShipBobOrdersPresentationViewModelsRecipientInfoViewModel`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetRecipientOk() (*ShipBobOrdersPresentationViewModelsRecipientInfoViewModel, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetRecipient(v ShipBobOrdersPresentationViewModelsRecipientInfoViewModel)`

SetRecipient sets Recipient field to given value.


### GetReferenceId

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetShippingMethod

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.


### GetTags

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetTags() []ShipBobOrdersPresentationViewModelsTagViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetTagsOk() (*[]ShipBobOrdersPresentationViewModelsTagViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetTags(v []ShipBobOrdersPresentationViewModelsTagViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShipBobOrdersPresentationModelsCreateOrderModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


