# CreateReceivingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxPackagingType** | [**PackingType**](PackingType.md) |  | 
**Boxes** | [**[]AddBoxToOrder**](AddBoxToOrder.md) | Box shipments to be added to this receiving order | 
**ExpectedArrivalDate** | **NullableTime** | Expected arrival date of all the box shipments in this receiving order | 
**FulfillmentCenter** | [**AssignOrderToFulfillmentCenter**](AssignOrderToFulfillmentCenter.md) |  | 
**PackageType** | [**PackageType**](PackageType.md) |  | 
**PurchaseOrderNumber** | Pointer to **NullableString** | Purchase order number for this receiving order | [optional] 

## Methods

### NewCreateReceivingOrder

`func NewCreateReceivingOrder(boxPackagingType PackingType, boxes []AddBoxToOrder, expectedArrivalDate NullableTime, fulfillmentCenter AssignOrderToFulfillmentCenter, packageType PackageType, ) *CreateReceivingOrder`

NewCreateReceivingOrder instantiates a new CreateReceivingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceivingOrderWithDefaults

`func NewCreateReceivingOrderWithDefaults() *CreateReceivingOrder`

NewCreateReceivingOrderWithDefaults instantiates a new CreateReceivingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxPackagingType

`func (o *CreateReceivingOrder) GetBoxPackagingType() PackingType`

GetBoxPackagingType returns the BoxPackagingType field if non-nil, zero value otherwise.

### GetBoxPackagingTypeOk

`func (o *CreateReceivingOrder) GetBoxPackagingTypeOk() (*PackingType, bool)`

GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxPackagingType

`func (o *CreateReceivingOrder) SetBoxPackagingType(v PackingType)`

SetBoxPackagingType sets BoxPackagingType field to given value.


### GetBoxes

`func (o *CreateReceivingOrder) GetBoxes() []AddBoxToOrder`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *CreateReceivingOrder) GetBoxesOk() (*[]AddBoxToOrder, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *CreateReceivingOrder) SetBoxes(v []AddBoxToOrder)`

SetBoxes sets Boxes field to given value.


### SetBoxesNil

`func (o *CreateReceivingOrder) SetBoxesNil(b bool)`

 SetBoxesNil sets the value for Boxes to be an explicit nil

### UnsetBoxes
`func (o *CreateReceivingOrder) UnsetBoxes()`

UnsetBoxes ensures that no value is present for Boxes, not even an explicit nil
### GetExpectedArrivalDate

`func (o *CreateReceivingOrder) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *CreateReceivingOrder) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *CreateReceivingOrder) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.


### SetExpectedArrivalDateNil

`func (o *CreateReceivingOrder) SetExpectedArrivalDateNil(b bool)`

 SetExpectedArrivalDateNil sets the value for ExpectedArrivalDate to be an explicit nil

### UnsetExpectedArrivalDate
`func (o *CreateReceivingOrder) UnsetExpectedArrivalDate()`

UnsetExpectedArrivalDate ensures that no value is present for ExpectedArrivalDate, not even an explicit nil
### GetFulfillmentCenter

`func (o *CreateReceivingOrder) GetFulfillmentCenter() AssignOrderToFulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *CreateReceivingOrder) GetFulfillmentCenterOk() (*AssignOrderToFulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *CreateReceivingOrder) SetFulfillmentCenter(v AssignOrderToFulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.


### GetPackageType

`func (o *CreateReceivingOrder) GetPackageType() PackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *CreateReceivingOrder) GetPackageTypeOk() (*PackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *CreateReceivingOrder) SetPackageType(v PackageType)`

SetPackageType sets PackageType field to given value.


### GetPurchaseOrderNumber

`func (o *CreateReceivingOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *CreateReceivingOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *CreateReceivingOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *CreateReceivingOrder) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### SetPurchaseOrderNumberNil

`func (o *CreateReceivingOrder) SetPurchaseOrderNumberNil(b bool)`

 SetPurchaseOrderNumberNil sets the value for PurchaseOrderNumber to be an explicit nil

### UnsetPurchaseOrderNumber
`func (o *CreateReceivingOrder) UnsetPurchaseOrderNumber()`

UnsetPurchaseOrderNumber ensures that no value is present for PurchaseOrderNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


