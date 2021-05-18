# ReceivingCreateReceivingOrderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxPackagingType** | [**ReceivingPackingType**](Receiving.PackingType.md) |  | 
**Boxes** | [**[]ReceivingAddBoxToOrderModel**](ReceivingAddBoxToOrderModel.md) | Box shipments to be added to this receiving order | 
**ExpectedArrivalDate** | **time.Time** | Expected arrival date of all the box shipments in this receiving order | 
**FulfillmentCenter** | [**ReceivingAssignOrderToFulfillmentCenterModel**](Receiving.AssignOrderToFulfillmentCenterModel.md) |  | 
**PackageType** | [**ReceivingPackageType**](Receiving.PackageType.md) |  | 

## Methods

### NewReceivingCreateReceivingOrderModel

`func NewReceivingCreateReceivingOrderModel(boxPackagingType ReceivingPackingType, boxes []ReceivingAddBoxToOrderModel, expectedArrivalDate time.Time, fulfillmentCenter ReceivingAssignOrderToFulfillmentCenterModel, packageType ReceivingPackageType, ) *ReceivingCreateReceivingOrderModel`

NewReceivingCreateReceivingOrderModel instantiates a new ReceivingCreateReceivingOrderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingCreateReceivingOrderModelWithDefaults

`func NewReceivingCreateReceivingOrderModelWithDefaults() *ReceivingCreateReceivingOrderModel`

NewReceivingCreateReceivingOrderModelWithDefaults instantiates a new ReceivingCreateReceivingOrderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxPackagingType

`func (o *ReceivingCreateReceivingOrderModel) GetBoxPackagingType() ReceivingPackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *ReceivingCreateReceivingOrderModel) GetBoxPackagingTypeOk() (*ReceivingPackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *ReceivingCreateReceivingOrderModel) SetBoxPackagingType(v ReceivingPackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.


### GetBoxes

`func (o *ReceivingCreateReceivingOrderModel) GetBoxes() []ReceivingAddBoxToOrderModel`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ReceivingCreateReceivingOrderModel) GetBoxesOk() (*[]ReceivingAddBoxToOrderModel, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ReceivingCreateReceivingOrderModel) SetBoxes(v []ReceivingAddBoxToOrderModel)`

SetBoxes sets Boxes field to given value.


### SetBoxesNil

`func (o *ReceivingCreateReceivingOrderModel) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *ReceivingCreateReceivingOrderModel) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ReceivingCreateReceivingOrderModel) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ReceivingCreateReceivingOrderModel) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ReceivingCreateReceivingOrderModel) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.


### GetFulfillmentCenter

`func (o *ReceivingCreateReceivingOrderModel) GetFulfillmentCenter() ReceivingAssignOrderToFulfillmentCenterModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ReceivingCreateReceivingOrderModel) GetFulfillmentCenterOk() (*ReceivingAssignOrderToFulfillmentCenterModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ReceivingCreateReceivingOrderModel) SetFulfillmentCenter(v ReceivingAssignOrderToFulfillmentCenterModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.


### GetPackageType

`func (o *ReceivingCreateReceivingOrderModel) GetPackageType() ReceivingPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ReceivingCreateReceivingOrderModel) GetPackageTypeOk() (*ReceivingPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ReceivingCreateReceivingOrderModel) SetPackageType(v ReceivingPackageType)`

SetPackageType sets PackageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


