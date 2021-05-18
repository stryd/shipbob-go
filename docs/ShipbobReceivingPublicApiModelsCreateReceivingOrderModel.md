# ShipbobReceivingPublicApiModelsCreateReceivingOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxPackagingType** | [**ShipbobReceivingPublicCommonModelsPackingType**](Shipbob.Receiving.Public.Common.Models.PackingType.md) |  | 
**Boxes** | [**[]ShipbobReceivingPublicApiModelsAddBoxToOrderModel**](ShipbobReceivingPublicApiModelsAddBoxToOrderModel.md) | Box shipments to be added to this receiving order | 
**ExpectedArrivalDate** | **time.Time** | Expected arrival date of all the box shipments in this receiving order | 
**FulfillmentCenter** | [**ShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel**](Shipbob.Receiving.Public.Api.Models.AssignOrderToFulfillmentCenterModel.md) |  | 
**PackageType** | [**ShipbobReceivingPublicCommonModelsPackageType**](Shipbob.Receiving.Public.Common.Models.PackageType.md) |  | 

## Methods

### NewShipbobReceivingPublicApiModelsCreateReceivingOrderModel

`func NewShipbobReceivingPublicApiModelsCreateReceivingOrderModel(boxPackagingType ShipbobReceivingPublicCommonModelsPackingType, boxes []ShipbobReceivingPublicApiModelsAddBoxToOrderModel, expectedArrivalDate time.Time, fulfillmentCenter ShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel, packageType ShipbobReceivingPublicCommonModelsPackageType, ) *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel`

NewShipbobReceivingPublicApiModelsCreateReceivingOrderModel instantiates a new ShipbobReceivingPublicApiModelsCreateReceivingOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipbobReceivingPublicApiModelsCreateReceivingOrderModelWithDefaults

`func NewShipbobReceivingPublicApiModelsCreateReceivingOrderModelWithDefaults() *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel`

NewShipbobReceivingPublicApiModelsCreateReceivingOrderModelWithDefaults instantiates a new ShipbobReceivingPublicApiModelsCreateReceivingOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxPackagingType

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetBoxPackagingType() ShipbobReceivingPublicCommonModelsPackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetBoxPackagingTypeOk() (*ShipbobReceivingPublicCommonModelsPackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetBoxPackagingType(v ShipbobReceivingPublicCommonModelsPackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.


### GetBoxes

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetBoxes() []ShipbobReceivingPublicApiModelsAddBoxToOrderModel`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetBoxesOk() (*[]ShipbobReceivingPublicApiModelsAddBoxToOrderModel, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetBoxes(v []ShipbobReceivingPublicApiModelsAddBoxToOrderModel)`

SetBoxes sets Boxes field to given value.


### SetBoxesNil

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.


### GetFulfillmentCenter

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetFulfillmentCenter() ShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetFulfillmentCenterOk() (*ShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetFulfillmentCenter(v ShipbobReceivingPublicApiModelsAssignOrderToFulfillmentCenterModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.


### GetPackageType

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetPackageType() ShipbobReceivingPublicCommonModelsPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) GetPackageTypeOk() (*ShipbobReceivingPublicCommonModelsPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ShipbobReceivingPublicApiModelsCreateReceivingOrderModel) SetPackageType(v ShipbobReceivingPublicCommonModelsPackageType)`

SetPackageType sets PackageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


