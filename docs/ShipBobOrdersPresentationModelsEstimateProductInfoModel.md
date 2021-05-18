# ShipBobOrdersPresentationModelsEstimateProductInfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique id of the product (Must be provided if reference_id is unknown) | [optional] 
**Quantity** | **int32** | The quantity of this product ordered | 
**ReferenceId** | Pointer to **string** | Unique reference id of the product (Must be provided if ID is unknown) | [optional] 

## Methods

### NewShipBobOrdersPresentationModelsEstimateProductInfoModel

`func NewShipBobOrdersPresentationModelsEstimateProductInfoModel(quantity int32, ) *ShipBobOrdersPresentationModelsEstimateProductInfoModel`

NewShipBobOrdersPresentationModelsEstimateProductInfoModel instantiates a new ShipBobOrdersPresentationModelsEstimateProductInfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationModelsEstimateProductInfoModelWithDefaults

`func NewShipBobOrdersPresentationModelsEstimateProductInfoModelWithDefaults() *ShipBobOrdersPresentationModelsEstimateProductInfoModel`

NewShipBobOrdersPresentationModelsEstimateProductInfoModelWithDefaults instantiates a new ShipBobOrdersPresentationModelsEstimateProductInfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuantity

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetReferenceId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipBobOrdersPresentationModelsEstimateProductInfoModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


