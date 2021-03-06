# OrderStatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Describes what the status detail means | [optional] 
**ExceptionFulfillmentCenterId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **int32** | The id of the status detail | [optional] 
**InventoryId** | Pointer to **NullableInt32** | Inventory Id the detail applies to (if applicable) | [optional] 
**Name** | Pointer to **string** | Short name of the status detail | [optional] 

## Methods

### NewOrderStatusDetails

`func NewOrderStatusDetails() *OrderStatusDetails`

NewOrderStatusDetails instantiates a new OrderStatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusDetailsWithDefaults

`func NewOrderStatusDetailsWithDefaults() *OrderStatusDetails`

NewOrderStatusDetailsWithDefaults instantiates a new OrderStatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrderStatusDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderStatusDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderStatusDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderStatusDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExceptionFulfillmentCenterId

`func (o *OrderStatusDetails) GetExceptionFulfillmentCenterId() int32`

GetExceptionFulfillmentCenterId returns the ExceptionFulfillmentCenterId field if non-nil, zero value otherwise.

### GetExceptionFulfillmentCenterIdOk

`func (o *OrderStatusDetails) GetExceptionFulfillmentCenterIdOk() (*int32, bool)`

GetExceptionFulfillmentCenterIdOk returns a tuple with the ExceptionFulfillmentCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionFulfillmentCenterId

`func (o *OrderStatusDetails) SetExceptionFulfillmentCenterId(v int32)`

SetExceptionFulfillmentCenterId sets ExceptionFulfillmentCenterId field to given value.

### HasExceptionFulfillmentCenterId

`func (o *OrderStatusDetails) HasExceptionFulfillmentCenterId() bool`

HasExceptionFulfillmentCenterId returns a boolean if a field has been set.

### SetExceptionFulfillmentCenterIdNil

`func (o *OrderStatusDetails) SetExceptionFulfillmentCenterIdNil(b bool)`

 SetExceptionFulfillmentCenterIdNil sets the value for ExceptionFulfillmentCenterId to be an explicit nil

### UnsetExceptionFulfillmentCenterId
`func (o *OrderStatusDetails) UnsetExceptionFulfillmentCenterId()`

UnsetExceptionFulfillmentCenterId ensures that no value is present for ExceptionFulfillmentCenterId, not even an explicit nil
### GetId

`func (o *OrderStatusDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatusDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatusDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatusDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryId

`func (o *OrderStatusDetails) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *OrderStatusDetails) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *OrderStatusDetails) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *OrderStatusDetails) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### SetInventoryIdNil

`func (o *OrderStatusDetails) SetInventoryIdNil(b bool)`

 SetInventoryIdNil sets the value for InventoryId to be an explicit nil

### UnsetInventoryId
`func (o *OrderStatusDetails) UnsetInventoryId()`

UnsetInventoryId ensures that no value is present for InventoryId, not even an explicit nil
### GetName

`func (o *OrderStatusDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderStatusDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderStatusDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderStatusDetails) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


