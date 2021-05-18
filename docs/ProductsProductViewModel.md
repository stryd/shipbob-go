# ProductsProductViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**BundleRootInformation** | Pointer to [**ProductsBundleRootInformationViewModel**](Products.BundleRootInformationViewModel.md) |  | [optional] 
**Channel** | Pointer to [**ProductsChannelViewModel**](Products.ChannelViewModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date the product was created | [optional] 
**FulfillableInventoryItems** | Pointer to [**[]ProductsInventoryItemViewModel**](ProductsInventoryItemViewModel.md) | The inventory that this product will resolve to when packing a shipment | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]ProductsFulfillmentCenterQuantityViewModel**](ProductsFulfillmentCenterQuantityViewModel.md) | Fulfillable quantity of this product broken down by fulfillment center location | [optional] 
**Id** | Pointer to **int32** | Unique identifier of the product | [optional] 
**Name** | Pointer to **NullableString** | The name of the product | [optional] 
**ReferenceId** | Pointer to **NullableString** | Unique reference identifier of the product | [optional] 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 
**TotalCommittedQuantity** | Pointer to **int32** | Total committed quantity of this product | [optional] 
**TotalFulfillableQuantity** | Pointer to **int32** | Total fulfillable quantity of this product | [optional] 
**TotalOnhandQuantity** | Pointer to **int32** | Total on hand quantity of this product | [optional] 

## Methods

### NewProductsProductViewModel

`func NewProductsProductViewModel() *ProductsProductViewModel`

NewProductsProductViewModel instantiates a new ProductsProductViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsProductViewModelWithDefaults

`func NewProductsProductViewModelWithDefaults() *ProductsProductViewModel`

NewProductsProductViewModelWithDefaults instantiates a new ProductsProductViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ProductsProductViewModel) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductsProductViewModel) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductsProductViewModel) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductsProductViewModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ProductsProductViewModel) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ProductsProductViewModel) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetBundleRootInformation

`func (o *ProductsProductViewModel) GetBundleRootInformation() ProductsBundleRootInformationViewModel`

GetBundleRootInformation returns the BundleRootInformation field if non-nil, zero value otherwise.

### GetBundleRootInformationOk

`func (o *ProductsProductViewModel) GetBundleRootInformationOk() (*ProductsBundleRootInformationViewModel, bool)`

GetBundleRootInformationOk returns a tuple with the BundleRootInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRootInformation

`func (o *ProductsProductViewModel) SetBundleRootInformation(v ProductsBundleRootInformationViewModel)`

SetBundleRootInformation sets BundleRootInformation field to given value.

### HasBundleRootInformation

`func (o *ProductsProductViewModel) HasBundleRootInformation() bool`

HasBundleRootInformation returns a boolean if a field has been set.

### GetChannel

`func (o *ProductsProductViewModel) GetChannel() ProductsChannelViewModel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ProductsProductViewModel) GetChannelOk() (*ProductsChannelViewModel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ProductsProductViewModel) SetChannel(v ProductsChannelViewModel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ProductsProductViewModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ProductsProductViewModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProductsProductViewModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProductsProductViewModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProductsProductViewModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFulfillableInventoryItems

`func (o *ProductsProductViewModel) GetFulfillableInventoryItems() []ProductsInventoryItemViewModel`

GetFulfillableInventoryItems returns the FulfillableInventoryItems field if non-nil, zero value otherwise.

### GetFulfillableInventoryItemsOk

`func (o *ProductsProductViewModel) GetFulfillableInventoryItemsOk() (*[]ProductsInventoryItemViewModel, bool)`

GetFulfillableInventoryItemsOk returns a tuple with the FulfillableInventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableInventoryItems

`func (o *ProductsProductViewModel) SetFulfillableInventoryItems(v []ProductsInventoryItemViewModel)`

SetFulfillableInventoryItems sets FulfillableInventoryItems field to given value.

### HasFulfillableInventoryItems

`func (o *ProductsProductViewModel) HasFulfillableInventoryItems() bool`

HasFulfillableInventoryItems returns a boolean if a field has been set.

### SetFulfillableInventoryItemsNil

`func (o *ProductsProductViewModel) SetFulfillableInventoryItemsNil(b bool)`

 SetFulfillableInventoryItemsNil sets the value for FulfillableInventoryItems to be an explicit nil

### UnsetFulfillableInventoryItems
`func (o *ProductsProductViewModel) UnsetFulfillableInventoryItems()`

UnsetFulfillableInventoryItems ensures that no value is present for FulfillableInventoryItems, not even an explicit nil
### GetFulfillableQuantityByFulfillmentCenter

`func (o *ProductsProductViewModel) GetFulfillableQuantityByFulfillmentCenter() []ProductsFulfillmentCenterQuantityViewModel`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *ProductsProductViewModel) GetFulfillableQuantityByFulfillmentCenterOk() (*[]ProductsFulfillmentCenterQuantityViewModel, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *ProductsProductViewModel) SetFulfillableQuantityByFulfillmentCenter(v []ProductsFulfillmentCenterQuantityViewModel)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *ProductsProductViewModel) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### SetFulfillableQuantityByFulfillmentCenterNil

`func (o *ProductsProductViewModel) SetFulfillableQuantityByFulfillmentCenterNil(b bool)`

 SetFulfillableQuantityByFulfillmentCenterNil sets the value for FulfillableQuantityByFulfillmentCenter to be an explicit nil

### UnsetFulfillableQuantityByFulfillmentCenter
`func (o *ProductsProductViewModel) UnsetFulfillableQuantityByFulfillmentCenter()`

UnsetFulfillableQuantityByFulfillmentCenter ensures that no value is present for FulfillableQuantityByFulfillmentCenter, not even an explicit nil
### GetId

`func (o *ProductsProductViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductsProductViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductsProductViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductsProductViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductsProductViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsProductViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsProductViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductsProductViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductsProductViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductsProductViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReferenceId

`func (o *ProductsProductViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ProductsProductViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ProductsProductViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ProductsProductViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ProductsProductViewModel) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ProductsProductViewModel) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetSku

`func (o *ProductsProductViewModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductsProductViewModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductsProductViewModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductsProductViewModel) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ProductsProductViewModel) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ProductsProductViewModel) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetTotalCommittedQuantity

`func (o *ProductsProductViewModel) GetTotalCommittedQuantity() int32`

GetTotalCommittedQuantity returns the TotalCommittedQuantity field if non-nil, zero value otherwise.

### GetTotalCommittedQuantityOk

`func (o *ProductsProductViewModel) GetTotalCommittedQuantityOk() (*int32, bool)`

GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommittedQuantity

`func (o *ProductsProductViewModel) SetTotalCommittedQuantity(v int32)`

SetTotalCommittedQuantity sets TotalCommittedQuantity field to given value.

### HasTotalCommittedQuantity

`func (o *ProductsProductViewModel) HasTotalCommittedQuantity() bool`

HasTotalCommittedQuantity returns a boolean if a field has been set.

### GetTotalFulfillableQuantity

`func (o *ProductsProductViewModel) GetTotalFulfillableQuantity() int32`

GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalFulfillableQuantityOk

`func (o *ProductsProductViewModel) GetTotalFulfillableQuantityOk() (*int32, bool)`

GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFulfillableQuantity

`func (o *ProductsProductViewModel) SetTotalFulfillableQuantity(v int32)`

SetTotalFulfillableQuantity sets TotalFulfillableQuantity field to given value.

### HasTotalFulfillableQuantity

`func (o *ProductsProductViewModel) HasTotalFulfillableQuantity() bool`

HasTotalFulfillableQuantity returns a boolean if a field has been set.

### GetTotalOnhandQuantity

`func (o *ProductsProductViewModel) GetTotalOnhandQuantity() int32`

GetTotalOnhandQuantity returns the TotalOnhandQuantity field if non-nil, zero value otherwise.

### GetTotalOnhandQuantityOk

`func (o *ProductsProductViewModel) GetTotalOnhandQuantityOk() (*int32, bool)`

GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnhandQuantity

`func (o *ProductsProductViewModel) SetTotalOnhandQuantity(v int32)`

SetTotalOnhandQuantity sets TotalOnhandQuantity field to given value.

### HasTotalOnhandQuantity

`func (o *ProductsProductViewModel) HasTotalOnhandQuantity() bool`

HasTotalOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


