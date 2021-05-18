# ShipbobReceivingPublicApiModelsReceivingOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxLabelsUri** | Pointer to **NullableString** | URL to the packing slip to be included in each box shipment for this receiving order | [optional] 
**BoxPackagingType** | Pointer to [**ShipbobReceivingPublicCommonModelsPackingType**](Shipbob.Receiving.Public.Common.Models.PackingType.md) |  | [optional] 
**Boxes** | Pointer to [**[]ShipbobReceivingPublicApiModelsBoxViewModel**](ShipbobReceivingPublicApiModelsBoxViewModel.md) | Information about the boxes being shipped in this receiving order | [optional] 
**ExpectedArrivalDate** | Pointer to **time.Time** | Expected date that all packages will have arrived | [optional] 
**FulfillmentCenter** | Pointer to [**ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel**](Shipbob.Receiving.Public.Api.Models.FulfillmentCenterViewModel.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique id of the warehouse receiving order | [optional] 
**InsertDate** | Pointer to **time.Time** | Insert date of the receiving order | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Last date the receiving order was updated | [optional] 
**PackageType** | Pointer to [**ShipbobReceivingPublicCommonModelsPackageType**](Shipbob.Receiving.Public.Common.Models.PackageType.md) |  | [optional] 
**Status** | Pointer to [**ShipbobReceivingPublicCommonModelsReceivingStatus**](Shipbob.Receiving.Public.Common.Models.ReceivingStatus.md) |  | [optional] 

## Methods

### NewShipbobReceivingPublicApiModelsReceivingOrderViewModel

`func NewShipbobReceivingPublicApiModelsReceivingOrderViewModel() *ShipbobReceivingPublicApiModelsReceivingOrderViewModel`

NewShipbobReceivingPublicApiModelsReceivingOrderViewModel instantiates a new ShipbobReceivingPublicApiModelsReceivingOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReceivingPublicApiModelsReceivingOrderViewModelWithDefaults

`func NewShipbobReceivingPublicApiModelsReceivingOrderViewModelWithDefaults() *ShipbobReceivingPublicApiModelsReceivingOrderViewModel`

NewShipbobReceivingPublicApiModelsReceivingOrderViewModelWithDefaults instantiates a new ShipbobReceivingPublicApiModelsReceivingOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxLabelsUri

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxLabelsUri() string`

GetBoxLabelsUri returns the BoxLabelsUri field if non-nil, zero value otherwise.

### GetBoxLabelsUriOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxLabelsUriOk() (*string, bool)`

GetBoxLabelsUriOk returns a tuple with the BoxLabelsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxLabelsUri

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxLabelsUri(v string)`

SetBoxLabelsUri sets BoxLabelsUri field to given value.

### HasBoxLabelsUri

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxLabelsUri() bool`

HasBoxLabelsUri returns a boolean if a field has been set.

### SetBoxLabelsUriNil

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxLabelsUriNil(b bool)`

 SetBoxLabelsUriNil sets the value for BoxLabelsUri to be an explicit nil

### UnsetBoxLabelsUri
`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) UnsetBoxLabelsUri()`

UnsetBoxLabelsUri ensures that no value is present for BoxLabelsUri, not even an explicit nil
### GetBoxPackagingType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxPackagingType() ShipbobReceivingPublicCommonModelsPackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxPackagingTypeOk() (*ShipbobReceivingPublicCommonModelsPackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxPackagingType(v ShipbobReceivingPublicCommonModelsPackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.

### HasBoxPackagingType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxPackagingType() bool`

HasBoxPackagingType returns a boolean if a field has been set.

### GetBoxes

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxes() []ShipbobReceivingPublicApiModelsBoxViewModel`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxesOk() (*[]ShipbobReceivingPublicApiModelsBoxViewModel, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxes(v []ShipbobReceivingPublicApiModelsBoxViewModel)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### SetBoxesNil

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetFulfillmentCenter() ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetFulfillmentCenterOk() (*ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetFulfillmentCenter(v ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetId

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetInsertDate() time.Time`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetInsertDateOk() (*time.Time, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetInsertDate(v time.Time)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetPackageType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetPackageType() ShipbobReceivingPublicCommonModelsPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetPackageTypeOk() (*ShipbobReceivingPublicCommonModelsPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetPackageType(v ShipbobReceivingPublicCommonModelsPackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetStatus

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetStatus() ShipbobReceivingPublicCommonModelsReceivingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetStatusOk() (*ShipbobReceivingPublicCommonModelsReceivingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetStatus(v ShipbobReceivingPublicCommonModelsReceivingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


