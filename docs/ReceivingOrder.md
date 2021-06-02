# ReceivingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxLabelsUri** | Pointer to **NullableString** | URL to the packing slip to be included in each box shipment for this receiving order | [optional] 
**BoxPackagingType** | Pointer to [**PackingType**](PackingType.md) |  | [optional] 
**Boxes** | Pointer to [**[]ReceivingOrderBoxes**](ReceivingOrderBoxes.md) | Information about the boxes being shipped in this receiving order | [optional] 
**ExpectedArrivalDate** | Pointer to **time.Time** | Expected date that all packages will have arrived | [optional] 
**FulfillmentCenter** | Pointer to [**ReceivingFulfillmentCenter**](ReceivingFulfillmentCenter.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the warehouse receiving order | [optional] 
**InsertDate** | Pointer to **time.Time** | Insert date of the receiving order | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Last date the receiving order was updated | [optional] 
**PackageType** | Pointer to [**PackageType**](PackageType.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewReceivingOrder

`func NewReceivingOrder() *ReceivingOrder`

NewReceivingOrder instantiates a new ReceivingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingOrderWithDefaults

`func NewReceivingOrderWithDefaults() *ReceivingOrder`

NewReceivingOrderWithDefaults instantiates a new ReceivingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxLabelsUri

`func (o *ReceivingOrder) GetBoxLabelsUri() string`

GetBoxLabelsUri returns the BoxLabelsUri field if non-nil, zero value otherwise.

### GetBoxLabelsUriOk

`func (o *ReceivingOrder) GetBoxLabelsUriOk() (*string, bool)`

GetBoxLabelsUriOk returns a tuple with the BoxLabelsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxLabelsUri

`func (o *ReceivingOrder) SetBoxLabelsUri(v string)`

SetBoxLabelsUri sets BoxLabelsUri field to given value.

### HasBoxLabelsUri

`func (o *ReceivingOrder) HasBoxLabelsUri() bool`

HasBoxLabelsUri returns a boolean if a field has been set.

### SetBoxLabelsUriNil

`func (o *ReceivingOrder) SetBoxLabelsUriNil(b bool)`

 SetBoxLabelsUriNil sets the value for BoxLabelsUri to be an explicit nil

### UnsetBoxLabelsUri
`func (o *ReceivingOrder) UnsetBoxLabelsUri()`

UnsetBoxLabelsUri ensures that no value is present for BoxLabelsUri, not even an explicit nil
### GetBoxPackagingType

`func (o *ReceivingOrder) GetBoxPackagingType() PackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *ReceivingOrder) GetBoxPackagingTypeOk() (*PackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *ReceivingOrder) SetBoxPackagingType(v PackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.

### HasBoxPackagingType

`func (o *ReceivingOrder) HasBoxPackagingType() bool`

HasBoxPackagingType returns a boolean if a field has been set.

### GetBoxes

`func (o *ReceivingOrder) GetBoxes() []ReceivingOrderBoxes`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ReceivingOrder) GetBoxesOk() (*[]ReceivingOrderBoxes, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ReceivingOrder) SetBoxes(v []ReceivingOrderBoxes)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *ReceivingOrder) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### SetBoxesNil

`func (o *ReceivingOrder) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *ReceivingOrder) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ReceivingOrder) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ReceivingOrder) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ReceivingOrder) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *ReceivingOrder) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ReceivingOrder) GetFulfillmentCenter() ReceivingFulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ReceivingOrder) GetFulfillmentCenterOk() (*ReceivingFulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ReceivingOrder) SetFulfillmentCenter(v ReceivingFulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ReceivingOrder) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ReceivingOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivingOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivingOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivingOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ReceivingOrder) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ReceivingOrder) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ReceivingOrder) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ReceivingOrder) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ReceivingOrder) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ReceivingOrder) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ReceivingOrder) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ReceivingOrder) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetPackageType

`func (o *ReceivingOrder) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ReceivingOrder) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ReceivingOrder) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ReceivingOrder) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetStatus

`func (o *ReceivingOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReceivingOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReceivingOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReceivingOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


