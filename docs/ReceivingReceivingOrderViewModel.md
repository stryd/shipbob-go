# ReceivingReceivingOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxLabelsUri** | Pointer to **NullableString** | URL to the packing slip to be included in each box shipment for this receiving order | [optional] 
**BoxPackagingType** | Pointer to [**ReceivingPackingType**](Receiving.PackingType.md) |  | [optional] 
**Boxes** | Pointer to [**[]ReceivingBoxViewModel**](ReceivingBoxViewModel.md) | Information about the boxes being shipped in this receiving order | [optional] 
**ExpectedArrivalDate** | Pointer to **time.Time** | Expected date that all packages will have arrived | [optional] 
**FulfillmentCenter** | Pointer to [**ReceivingFulfillmentCenterViewModel**](Receiving.FulfillmentCenterViewModel.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the warehouse receiving order | [optional] 
**InsertDate** | Pointer to **time.Time** | Insert date of the receiving order | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Last date the receiving order was updated | [optional] 
**PackageType** | Pointer to [**ReceivingPackageType**](Receiving.PackageType.md) |  | [optional] 
**Status** | Pointer to [**ReceivingReceivingStatus**](Receiving.ReceivingStatus.md) |  | [optional] 

## Methods

### NewReceivingReceivingOrderViewModel

`func NewReceivingReceivingOrderViewModel() *ReceivingReceivingOrderViewModel`

NewReceivingReceivingOrderViewModel instantiates a new ReceivingReceivingOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingReceivingOrderViewModelWithDefaults

`func NewReceivingReceivingOrderViewModelWithDefaults() *ReceivingReceivingOrderViewModel`

NewReceivingReceivingOrderViewModelWithDefaults instantiates a new ReceivingReceivingOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxLabelsUri

`func (o *ReceivingReceivingOrderViewModel) GetBoxLabelsUri() string`

GetBoxLabelsUri returns the BoxLabelsUri field if non-nil, zero value otherwise.

### GetBoxLabelsUriOk

`func (o *ReceivingReceivingOrderViewModel) GetBoxLabelsUriOk() (*string, bool)`

GetBoxLabelsUriOk returns a tuple with the BoxLabelsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxLabelsUri

`func (o *ReceivingReceivingOrderViewModel) SetBoxLabelsUri(v string)`

SetBoxLabelsUri sets BoxLabelsUri field to given value.

### HasBoxLabelsUri

`func (o *ReceivingReceivingOrderViewModel) HasBoxLabelsUri() bool`

HasBoxLabelsUri returns a boolean if a field has been set.

### SetBoxLabelsUriNil

`func (o *ReceivingReceivingOrderViewModel) SetBoxLabelsUriNil(b bool)`

 SetBoxLabelsUriNil sets the value for BoxLabelsUri to be an explicit nil

### UnsetBoxLabelsUri
`func (o *ReceivingReceivingOrderViewModel) UnsetBoxLabelsUri()`

UnsetBoxLabelsUri ensures that no value is present for BoxLabelsUri, not even an explicit nil
### GetBoxPackagingType

`func (o *ReceivingReceivingOrderViewModel) GetBoxPackagingType() ReceivingPackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *ReceivingReceivingOrderViewModel) GetBoxPackagingTypeOk() (*ReceivingPackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *ReceivingReceivingOrderViewModel) SetBoxPackagingType(v ReceivingPackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.

### HasBoxPackagingType

`func (o *ReceivingReceivingOrderViewModel) HasBoxPackagingType() bool`

HasBoxPackagingType returns a boolean if a field has been set.

### GetBoxes

`func (o *ReceivingReceivingOrderViewModel) GetBoxes() []ReceivingBoxViewModel`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ReceivingReceivingOrderViewModel) GetBoxesOk() (*[]ReceivingBoxViewModel, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ReceivingReceivingOrderViewModel) SetBoxes(v []ReceivingBoxViewModel)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *ReceivingReceivingOrderViewModel) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### SetBoxesNil

`func (o *ReceivingReceivingOrderViewModel) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *ReceivingReceivingOrderViewModel) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ReceivingReceivingOrderViewModel) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ReceivingReceivingOrderViewModel) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ReceivingReceivingOrderViewModel) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *ReceivingReceivingOrderViewModel) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ReceivingReceivingOrderViewModel) GetFulfillmentCenter() ReceivingFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ReceivingReceivingOrderViewModel) GetFulfillmentCenterOk() (*ReceivingFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ReceivingReceivingOrderViewModel) SetFulfillmentCenter(v ReceivingFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ReceivingReceivingOrderViewModel) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ReceivingReceivingOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivingReceivingOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivingReceivingOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivingReceivingOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ReceivingReceivingOrderViewModel) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ReceivingReceivingOrderViewModel) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ReceivingReceivingOrderViewModel) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ReceivingReceivingOrderViewModel) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ReceivingReceivingOrderViewModel) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ReceivingReceivingOrderViewModel) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ReceivingReceivingOrderViewModel) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ReceivingReceivingOrderViewModel) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetPackageType

`func (o *ReceivingReceivingOrderViewModel) GetPackageType() ReceivingPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ReceivingReceivingOrderViewModel) GetPackageTypeOk() (*ReceivingPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ReceivingReceivingOrderViewModel) SetPackageType(v ReceivingPackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ReceivingReceivingOrderViewModel) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetStatus

`func (o *ReceivingReceivingOrderViewModel) GetStatus() ReceivingReceivingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReceivingReceivingOrderViewModel) GetStatusOk() (*ReceivingReceivingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReceivingReceivingOrderViewModel) SetStatus(v ReceivingReceivingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReceivingReceivingOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


