# WarehouseReceivingOrderV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxLabelsUri** | Pointer to **NullableString** | URL to the packing slip to be included in each box shipment for this receiving order | [optional] 
**BoxPackagingType** | Pointer to [**PackingType**](PackingType.md) |  | [optional] 
**ExpectedArrivalDate** | Pointer to **NullableTime** | Expected date that all packages will have arrived | [optional] 
**FulfillmentCenter** | Pointer to [**ReceivingFulfillmentCenter**](ReceivingFulfillmentCenter.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the warehouse receiving order | [optional] 
**InsertDate** | Pointer to **NullableTime** | Insert date of the receiving order | [optional] 
**InventoryQuantities** | Pointer to [**[]InventoryQuantityV2**](InventoryQuantityV2.md) | Inventory items and quantities within the WRO | [optional] 
**LastUpdatedDate** | Pointer to **NullableTime** | Last date the receiving order was updated | [optional] 
**PackageType** | Pointer to [**PackageType**](PackageType.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **NullableString** | Purchase order number for a receiving order | [optional] 
**Status** | Pointer to [**ReceivingStatus**](ReceivingStatus.md) |  | [optional] 

## Methods

### NewWarehouseReceivingOrderV2

`func NewWarehouseReceivingOrderV2() *WarehouseReceivingOrderV2`

NewWarehouseReceivingOrderV2 instantiates a new WarehouseReceivingOrderV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseReceivingOrderV2WithDefaults

`func NewWarehouseReceivingOrderV2WithDefaults() *WarehouseReceivingOrderV2`

NewWarehouseReceivingOrderV2WithDefaults instantiates a new WarehouseReceivingOrderV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxLabelsUri

`func (o *WarehouseReceivingOrderV2) GetBoxLabelsUri() string`

GetBoxLabelsUri returns the BoxLabelsUri field if non-nil, zero value otherwise.

### GetBoxLabelsUriOk

`func (o *WarehouseReceivingOrderV2) GetBoxLabelsUriOk() (*string, bool)`

GetBoxLabelsUriOk returns a tuple with the BoxLabelsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxLabelsUri

`func (o *WarehouseReceivingOrderV2) SetBoxLabelsUri(v string)`

SetBoxLabelsUri sets BoxLabelsUri field to given value.

### HasBoxLabelsUri

`func (o *WarehouseReceivingOrderV2) HasBoxLabelsUri() bool`

HasBoxLabelsUri returns a boolean if a field has been set.

### SetBoxLabelsUriNil

`func (o *WarehouseReceivingOrderV2) SetBoxLabelsUriNil(b bool)`

 SetBoxLabelsUriNil sets the value for BoxLabelsUri to be an explicit nil

### UnsetBoxLabelsUri
`func (o *WarehouseReceivingOrderV2) UnsetBoxLabelsUri()`

UnsetBoxLabelsUri ensures that no value is present for BoxLabelsUri, not even an explicit nil
### GetBoxPackagingType

`func (o *WarehouseReceivingOrderV2) GetBoxPackagingType() PackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *WarehouseReceivingOrderV2) GetBoxPackagingTypeOk() (*PackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *WarehouseReceivingOrderV2) SetBoxPackagingType(v PackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.

### HasBoxPackagingType

`func (o *WarehouseReceivingOrderV2) HasBoxPackagingType() bool`

HasBoxPackagingType returns a boolean if a field has been set.

### GetExpectedArrivalDate

`func (o *WarehouseReceivingOrderV2) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *WarehouseReceivingOrderV2) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *WarehouseReceivingOrderV2) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *WarehouseReceivingOrderV2) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### SetExpectedArrivalDateNil

`func (o *WarehouseReceivingOrderV2) SetExpectedArrivalDateNil(b bool)`

 SetExpectedArrivalDateNil sets the value for ExpectedArrivalDate to be an explicit nil

### UnsetExpectedArrivalDate
`func (o *WarehouseReceivingOrderV2) UnsetExpectedArrivalDate()`

UnsetExpectedArrivalDate ensures that no value is present for ExpectedArrivalDate, not even an explicit nil
### GetFulfillmentCenter

`func (o *WarehouseReceivingOrderV2) GetFulfillmentCenter() ReceivingFulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *WarehouseReceivingOrderV2) GetFulfillmentCenterOk() (*ReceivingFulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *WarehouseReceivingOrderV2) SetFulfillmentCenter(v ReceivingFulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *WarehouseReceivingOrderV2) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *WarehouseReceivingOrderV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseReceivingOrderV2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseReceivingOrderV2) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseReceivingOrderV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *WarehouseReceivingOrderV2) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *WarehouseReceivingOrderV2) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *WarehouseReceivingOrderV2) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *WarehouseReceivingOrderV2) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### SetInsertDateNil

`func (o *WarehouseReceivingOrderV2) SetInsertDateNil(b bool)`

 SetInsertDateNil sets the value for InsertDate to be an explicit nil

### UnsetInsertDate
`func (o *WarehouseReceivingOrderV2) UnsetInsertDate()`

UnsetInsertDate ensures that no value is present for InsertDate, not even an explicit nil
### GetInventoryQuantities

`func (o *WarehouseReceivingOrderV2) GetInventoryQuantities() []InventoryQuantityV2`

GetInventoryQuantities returns the InventoryQuantities field if non-nil, zero value otherwise.

### GetInventoryQuantitiesOk

`func (o *WarehouseReceivingOrderV2) GetInventoryQuantitiesOk() (*[]InventoryQuantityV2, bool)`

GetInventoryQuantitiesOk returns a tuple with the InventoryQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryQuantities

`func (o *WarehouseReceivingOrderV2) SetInventoryQuantities(v []InventoryQuantityV2)`

SetInventoryQuantities sets InventoryQuantities field to given value.

### HasInventoryQuantities

`func (o *WarehouseReceivingOrderV2) HasInventoryQuantities() bool`

HasInventoryQuantities returns a boolean if a field has been set.

### SetInventoryQuantitiesNil

`func (o *WarehouseReceivingOrderV2) SetInventoryQuantitiesNil(b bool)`

 SetInventoryQuantitiesNil sets the value for InventoryQuantities to be an explicit nil

### UnsetInventoryQuantities
`func (o *WarehouseReceivingOrderV2) UnsetInventoryQuantities()`

UnsetInventoryQuantities ensures that no value is present for InventoryQuantities, not even an explicit nil
### GetLastUpdatedDate

`func (o *WarehouseReceivingOrderV2) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *WarehouseReceivingOrderV2) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *WarehouseReceivingOrderV2) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *WarehouseReceivingOrderV2) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### SetLastUpdatedDateNil

`func (o *WarehouseReceivingOrderV2) SetLastUpdatedDateNil(b bool)`

 SetLastUpdatedDateNil sets the value for LastUpdatedDate to be an explicit nil

### UnsetLastUpdatedDate
`func (o *WarehouseReceivingOrderV2) UnsetLastUpdatedDate()`

UnsetLastUpdatedDate ensures that no value is present for LastUpdatedDate, not even an explicit nil
### GetPackageType

`func (o *WarehouseReceivingOrderV2) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *WarehouseReceivingOrderV2) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *WarehouseReceivingOrderV2) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *WarehouseReceivingOrderV2) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *WarehouseReceivingOrderV2) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *WarehouseReceivingOrderV2) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *WarehouseReceivingOrderV2) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *WarehouseReceivingOrderV2) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### SetPurchaseOrderNumberNil

`func (o *WarehouseReceivingOrderV2) SetPurchaseOrderNumberNil(b bool)`

 SetPurchaseOrderNumberNil sets the value for PurchaseOrderNumber to be an explicit nil

### UnsetPurchaseOrderNumber
`func (o *WarehouseReceivingOrderV2) UnsetPurchaseOrderNumber()`

UnsetPurchaseOrderNumber ensures that no value is present for PurchaseOrderNumber, not even an explicit nil
### GetStatus

`func (o *WarehouseReceivingOrderV2) GetStatus() ReceivingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WarehouseReceivingOrderV2) GetStatusOk() (*ReceivingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WarehouseReceivingOrderV2) SetStatus(v ReceivingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WarehouseReceivingOrderV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


