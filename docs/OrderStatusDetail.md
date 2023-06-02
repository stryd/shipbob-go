# OrderStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Describes what the status detail means | [optional] 
**ExceptionFulfillmentCenterId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **int32** | The id of the status detail | [optional] 
**InventoryId** | Pointer to **NullableInt32** | Inventory Id the detail applies to (if applicable) | [optional] 
**Name** | Pointer to **string** | Short name of the status detail | [optional] 

## Methods

### NewOrderStatusDetail

`func NewOrderStatusDetail() *OrderStatusDetail`

NewOrderStatusDetail instantiates a new OrderStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusDetailWithDefaults

`func NewOrderStatusDetailWithDefaults() *OrderStatusDetail`

NewOrderStatusDetailWithDefaults instantiates a new OrderStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrderStatusDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderStatusDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderStatusDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderStatusDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExceptionFulfillmentCenterId

`func (o *OrderStatusDetail) GetExceptionFulfillmentCenterId() int32`

GetExceptionFulfillmentCenterId returns the ExceptionFulfillmentCenterId field if non-nil, zero value otherwise.

### GetExceptionFulfillmentCenterIdOk

`func (o *OrderStatusDetail) GetExceptionFulfillmentCenterIdOk() (*int32, bool)`

GetExceptionFulfillmentCenterIdOk returns a tuple with the ExceptionFulfillmentCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionFulfillmentCenterId

`func (o *OrderStatusDetail) SetExceptionFulfillmentCenterId(v int32)`

SetExceptionFulfillmentCenterId sets ExceptionFulfillmentCenterId field to given value.

### HasExceptionFulfillmentCenterId

`func (o *OrderStatusDetail) HasExceptionFulfillmentCenterId() bool`

HasExceptionFulfillmentCenterId returns a boolean if a field has been set.

### SetExceptionFulfillmentCenterIdNil

`func (o *OrderStatusDetail) SetExceptionFulfillmentCenterIdNil(b bool)`

 SetExceptionFulfillmentCenterIdNil sets the value for ExceptionFulfillmentCenterId to be an explicit nil

### UnsetExceptionFulfillmentCenterId
`func (o *OrderStatusDetail) UnsetExceptionFulfillmentCenterId()`

UnsetExceptionFulfillmentCenterId ensures that no value is present for ExceptionFulfillmentCenterId, not even an explicit nil
### GetId

`func (o *OrderStatusDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatusDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatusDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatusDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryId

`func (o *OrderStatusDetail) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *OrderStatusDetail) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *OrderStatusDetail) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *OrderStatusDetail) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### SetInventoryIdNil

`func (o *OrderStatusDetail) SetInventoryIdNil(b bool)`

 SetInventoryIdNil sets the value for InventoryId to be an explicit nil

### UnsetInventoryId
`func (o *OrderStatusDetail) UnsetInventoryId()`

UnsetInventoryId ensures that no value is present for InventoryId, not even an explicit nil
### GetName

`func (o *OrderStatusDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderStatusDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderStatusDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderStatusDetail) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


