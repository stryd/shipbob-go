# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **NullableString** | Barcode for the product | [optional] 
**BundleRootInformation** | Pointer to [**ProductBundleRootInformation**](ProductBundleRootInformation.md) |  | [optional] 
**Channel** | Pointer to [**ProductChannel**](ProductChannel.md) |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** | Date the product was created | [optional] 
**FulfillableInventoryItems** | Pointer to [**[]ProductInventoryItem**](ProductInventoryItem.md) | The inventory that this product will resolve to when packing a shipment | [optional] 
**FulfillableQuantityByFulfillmentCenter** | Pointer to [**[]ProductFulfillmentCenterQuantity**](ProductFulfillmentCenterQuantity.md) | Fulfillable quantity of this product broken down by fulfillment center location | [optional] 
**Id** | Pointer to **int32** | Unique identifier of the product | [optional] 
**Name** | Pointer to **NullableString** | The name of the product | [optional] 
**ReferenceId** | Pointer to **NullableString** | Unique reference identifier of the product | [optional] 
**Sku** | Pointer to **NullableString** | Stock keeping unit for the product | [optional] 
**TotalCommittedQuantity** | Pointer to **int32** | Total committed quantity of this product | [optional] 
**TotalFulfillableQuantity** | Pointer to **int32** | Total fulfillable quantity of this product | [optional] 
**TotalOnhandQuantity** | Pointer to **int32** | Total on hand quantity of this product | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *Product) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Product) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Product) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Product) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *Product) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *Product) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetBundleRootInformation

`func (o *Product) GetBundleRootInformation() ProductBundleRootInformation`

GetBundleRootInformation returns the BundleRootInformation field if non-nil, zero value otherwise.

### GetBundleRootInformationOk

`func (o *Product) GetBundleRootInformationOk() (*ProductBundleRootInformation, bool)`

GetBundleRootInformationOk returns a tuple with the BundleRootInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleRootInformation

`func (o *Product) SetBundleRootInformation(v ProductBundleRootInformation)`

SetBundleRootInformation sets BundleRootInformation field to given value.

### HasBundleRootInformation

`func (o *Product) HasBundleRootInformation() bool`

HasBundleRootInformation returns a boolean if a field has been set.

### GetChannel

`func (o *Product) GetChannel() ProductChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Product) GetChannelOk() (*ProductChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Product) SetChannel(v ProductChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Product) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Product) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Product) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Product) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Product) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *Product) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *Product) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetFulfillableInventoryItems

`func (o *Product) GetFulfillableInventoryItems() []ProductInventoryItem`

GetFulfillableInventoryItems returns the FulfillableInventoryItems field if non-nil, zero value otherwise.

### GetFulfillableInventoryItemsOk

`func (o *Product) GetFulfillableInventoryItemsOk() (*[]ProductInventoryItem, bool)`

GetFulfillableInventoryItemsOk returns a tuple with the FulfillableInventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableInventoryItems

`func (o *Product) SetFulfillableInventoryItems(v []ProductInventoryItem)`

SetFulfillableInventoryItems sets FulfillableInventoryItems field to given value.

### HasFulfillableInventoryItems

`func (o *Product) HasFulfillableInventoryItems() bool`

HasFulfillableInventoryItems returns a boolean if a field has been set.

### SetFulfillableInventoryItemsNil

`func (o *Product) SetFulfillableInventoryItemsNil(b bool)`

 SetFulfillableInventoryItemsNil sets the value for FulfillableInventoryItems to be an explicit nil

### UnsetFulfillableInventoryItems
`func (o *Product) UnsetFulfillableInventoryItems()`

UnsetFulfillableInventoryItems ensures that no value is present for FulfillableInventoryItems, not even an explicit nil
### GetFulfillableQuantityByFulfillmentCenter

`func (o *Product) GetFulfillableQuantityByFulfillmentCenter() []ProductFulfillmentCenterQuantity`

GetFulfillableQuantityByFulfillmentCenter returns the FulfillableQuantityByFulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillableQuantityByFulfillmentCenterOk

`func (o *Product) GetFulfillableQuantityByFulfillmentCenterOk() (*[]ProductFulfillmentCenterQuantity, bool)`

GetFulfillableQuantityByFulfillmentCenterOk returns a tuple with the FulfillableQuantityByFulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantityByFulfillmentCenter

`func (o *Product) SetFulfillableQuantityByFulfillmentCenter(v []ProductFulfillmentCenterQuantity)`

SetFulfillableQuantityByFulfillmentCenter sets FulfillableQuantityByFulfillmentCenter field to given value.

### HasFulfillableQuantityByFulfillmentCenter

`func (o *Product) HasFulfillableQuantityByFulfillmentCenter() bool`

HasFulfillableQuantityByFulfillmentCenter returns a boolean if a field has been set.

### SetFulfillableQuantityByFulfillmentCenterNil

`func (o *Product) SetFulfillableQuantityByFulfillmentCenterNil(b bool)`

 SetFulfillableQuantityByFulfillmentCenterNil sets the value for FulfillableQuantityByFulfillmentCenter to be an explicit nil

### UnsetFulfillableQuantityByFulfillmentCenter
`func (o *Product) UnsetFulfillableQuantityByFulfillmentCenter()`

UnsetFulfillableQuantityByFulfillmentCenter ensures that no value is present for FulfillableQuantityByFulfillmentCenter, not even an explicit nil
### GetId

`func (o *Product) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Product) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Product) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReferenceId

`func (o *Product) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Product) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Product) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Product) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *Product) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *Product) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetSku

`func (o *Product) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Product) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Product) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Product) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *Product) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *Product) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetTotalCommittedQuantity

`func (o *Product) GetTotalCommittedQuantity() int32`

GetTotalCommittedQuantity returns the TotalCommittedQuantity field if non-nil, zero value otherwise.

### GetTotalCommittedQuantityOk

`func (o *Product) GetTotalCommittedQuantityOk() (*int32, bool)`

GetTotalCommittedQuantityOk returns a tuple with the TotalCommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommittedQuantity

`func (o *Product) SetTotalCommittedQuantity(v int32)`

SetTotalCommittedQuantity sets TotalCommittedQuantity field to given value.

### HasTotalCommittedQuantity

`func (o *Product) HasTotalCommittedQuantity() bool`

HasTotalCommittedQuantity returns a boolean if a field has been set.

### GetTotalFulfillableQuantity

`func (o *Product) GetTotalFulfillableQuantity() int32`

GetTotalFulfillableQuantity returns the TotalFulfillableQuantity field if non-nil, zero value otherwise.

### GetTotalFulfillableQuantityOk

`func (o *Product) GetTotalFulfillableQuantityOk() (*int32, bool)`

GetTotalFulfillableQuantityOk returns a tuple with the TotalFulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFulfillableQuantity

`func (o *Product) SetTotalFulfillableQuantity(v int32)`

SetTotalFulfillableQuantity sets TotalFulfillableQuantity field to given value.

### HasTotalFulfillableQuantity

`func (o *Product) HasTotalFulfillableQuantity() bool`

HasTotalFulfillableQuantity returns a boolean if a field has been set.

### GetTotalOnhandQuantity

`func (o *Product) GetTotalOnhandQuantity() int32`

GetTotalOnhandQuantity returns the TotalOnhandQuantity field if non-nil, zero value otherwise.

### GetTotalOnhandQuantityOk

`func (o *Product) GetTotalOnhandQuantityOk() (*int32, bool)`

GetTotalOnhandQuantityOk returns a tuple with the TotalOnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOnhandQuantity

`func (o *Product) SetTotalOnhandQuantity(v int32)`

SetTotalOnhandQuantity sets TotalOnhandQuantity field to given value.

### HasTotalOnhandQuantity

`func (o *Product) HasTotalOnhandQuantity() bool`

HasTotalOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


