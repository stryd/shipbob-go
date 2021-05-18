# ShipbobProductsApiViewModelsPublicProductViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**BundleRootInformation** | Pointer to [**ShipbobProductsApiViewModelsPublicBundleRootInformationViewModel**](Shipbob.Products.Api.ViewModels.Public.BundleRootInformationViewModel.md) |  | [optional] 
**Channel** | Pointer to [**ShipbobProductsApiViewModelsPublicChannelViewModel**](Shipbob.Products.Api.ViewModels.Public.ChannelViewModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date the product was created | [optional] 
**FulfillableInventoryItems** | Pointer to [**[]ShipbobProductsApiViewModelsPublicInventoryItemViewModel**](ShipbobProductsApiViewModelsPublicInventoryItemViewModel.md) | The inventory that this product will resolve to when packing a shipment | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]ShipbobProductsApiViewModelsPublicFulfillmentCenterQuantityViewModel**](ShipbobProductsApiViewModelsPublicFulfillmentCenterQuantityViewModel.md) | Fulfillable quantity of this product broken down by fulfillment center location | [optional] 
**Id** | Pointer to **int32** | Unique identifier of the product | [optional] 
**Name** | Pointer to **NullableString** | The name of the product | [optional] 
**ReferenceId** | Pointer to **NullableString** | Unique reference identifier of the product | [optional] 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 
**TotalCommittedQuantity** | Pointer to **int32** | Total committed quantity of this product | [optional] 
**TotalFulfillableQuantity** | Pointer to **int32** | Total fulfillable quantity of this product | [optional] 
**TotalOnhandQuantity** | Pointer to **int32** | Total on hand quantity of this product | [optional] 

## Methods

### NewShipbobProductsApiViewModelsPublicProductViewModel

`func NewShipbobProductsApiViewModelsPublicProductViewModel() *ShipbobProductsApiViewModelsPublicProductViewModel`

NewShipbobProductsApiViewModelsPublicProductViewModel instantiates a new ShipbobProductsApiViewModelsPublicProductViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobProductsApiViewModelsPublicProductViewModelWithDefaults

`func NewShipbobProductsApiViewModelsPublicProductViewModelWithDefaults() *ShipbobProductsApiViewModelsPublicProductViewModel`

NewShipbobProductsApiViewModelsPublicProductViewModelWithDefaults instantiates a new ShipbobProductsApiViewModelsPublicProductViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetBundleRootInformation

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetBundleRootInformation() ShipbobProductsApiViewModelsPublicBundleRootInformationViewModel`

GetBundleRootInformation returns the BundleRootInformation field if non-nil, zero value otherwise.

### GetBundleRootInformationOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetBundleRootInformationOk() (*ShipbobProductsApiViewModelsPublicBundleRootInformationViewModel, bool)`

GetBundleRootInformationOk returns a tuple with the BundleRootInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRootInformation

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetBundleRootInformation(v ShipbobProductsApiViewModelsPublicBundleRootInformationViewModel)`

SetBundleRootInformation sets BundleRootInformation field to given value.

### HasBundleRootInformation

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasBundleRootInformation() bool`

HasBundleRootInformation returns a boolean if a field has been set.

### GetChannel

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetChannel() ShipbobProductsApiViewModelsPublicChannelViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetChannelOk() (*ShipbobProductsApiViewModelsPublicChannelViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetChannel(v ShipbobProductsApiViewModelsPublicChannelViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFulfillableInventoryItems

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetFulfillableInventoryItems() []ShipbobProductsApiViewModelsPublicInventoryItemViewModel`

GetFulfillableInventoryItems returns the FulfillableInventoryItems field if non-nil, zero value otherwise.

### GetFulfillableInventoryItemsOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetFulfillableInventoryItemsOk() (*[]ShipbobProductsApiViewModelsPublicInventoryItemViewModel, bool)`

GetFulfillableInventoryItemsOk returns a tuple with the FulfillableInventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableInventoryItems

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetFulfillableInventoryItems(v []ShipbobProductsApiViewModelsPublicInventoryItemViewModel)`

SetFulfillableInventoryItems sets FulfillableInventoryItems field to given value.

### HasFulfillableInventoryItems

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasFulfillableInventoryItems() bool`

HasFulfillableInventoryItems returns a boolean if a field has been set.

### SetFulfillableInventoryItemsNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetFulfillableInventoryItemsNil(b bool)`

 SetFulfillableInventoryItemsNil sets the value for FulfillableInventoryItems to be an explicit nil

### UnsetFulfillableInventoryItems
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetFulfillableInventoryItems()`

UnsetFulfillableInventoryItems ensures that no value is present for FulfillableInventoryItems, not even an explicit nil
### GetFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetFulfillableQuantityByFulfillmentCenter() []ShipbobProductsApiViewModelsPublicFulfillmentCenterQuantityViewModel`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetFulfillableQuantityByFulfillmentCenterOk() (*[]ShipbobProductsApiViewModelsPublicFulfillmentCenterQuantityViewModel, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetFulfillableQuantityByFulfillmentCenter(v []ShipbobProductsApiViewModelsPublicFulfillmentCenterQuantityViewModel)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### SetFulfillableQuantityByFulfillmentCenterNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetFulfillableQuantityByFulfillmentCenterNil(b bool)`

 SetFulfillableQuantityByFulfillmentCenterNil sets the value for FulfillableQuantityByFulfillmentCenter to be an explicit nil

### UnsetFulfillableQuantityByFulfillmentCenter
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetFulfillableQuantityByFulfillmentCenter()`

UnsetFulfillableQuantityByFulfillmentCenter ensures that no value is present for FulfillableQuantityByFulfillmentCenter, not even an explicit nil
### GetId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReferenceId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetSku

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetTotalCommittedQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalCommittedQuantity() int32`

GetTotalCommittedQuantity returns the TotalCommittedQuantity field if non-nil, zero value otherwise.

### GetTotalCommittedQuantityOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalCommittedQuantityOk() (*int32, bool)`

GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommittedQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetTotalCommittedQuantity(v int32)`

SetTotalCommittedQuantity sets TotalCommittedQuantity field to given value.

### HasTotalCommittedQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasTotalCommittedQuantity() bool`

HasTotalCommittedQuantity returns a boolean if a field has been set.

### GetTotalFulfillableQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalFulfillableQuantity() int32`

GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalFulfillableQuantityOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalFulfillableQuantityOk() (*int32, bool)`

GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFulfillableQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetTotalFulfillableQuantity(v int32)`

SetTotalFulfillableQuantity sets TotalFulfillableQuantity field to given value.

### HasTotalFulfillableQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasTotalFulfillableQuantity() bool`

HasTotalFulfillableQuantity returns a boolean if a field has been set.

### GetTotalOnhandQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalOnhandQuantity() int32`

GetTotalOnhandQuantity returns the TotalOnhandQuantity field if non-nil, zero value otherwise.

### GetTotalOnhandQuantityOk

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) GetTotalOnhandQuantityOk() (*int32, bool)`

GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnhandQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) SetTotalOnhandQuantity(v int32)`

SetTotalOnhandQuantity sets TotalOnhandQuantity field to given value.

### HasTotalOnhandQuantity

`func (o *ShipbobProductsApiViewModelsPublicProductViewModel) HasTotalOnhandQuantity() bool`

HasTotalOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


