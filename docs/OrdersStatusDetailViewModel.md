# OrdersStatusDetailViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Describes what the status detail means | [optional] 
**ExceptionFulfillmentCenterId** | Pointer to **NullableInt32** |  | [optional] 
**Id** | Pointer to **int32** | The id of the status detail | [optional] 
**InventoryId** | Pointer to **NullableInt32** | Inventory Id the detail applies to (if applicable) | [optional] 
**Name** | Pointer to **string** | Short name of the status detail | [optional] 

## Methods

### NewOrdersStatusDetailViewModel

`func NewOrdersStatusDetailViewModel() *OrdersStatusDetailViewModel`

NewOrdersStatusDetailViewModel instantiates a new OrdersStatusDetailViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersStatusDetailViewModelWithDefaults

`func NewOrdersStatusDetailViewModelWithDefaults() *OrdersStatusDetailViewModel`

NewOrdersStatusDetailViewModelWithDefaults instantiates a new OrdersStatusDetailViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrdersStatusDetailViewModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrdersStatusDetailViewModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrdersStatusDetailViewModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrdersStatusDetailViewModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExceptionFulfillmentCenterId

`func (o *OrdersStatusDetailViewModel) GetExceptionFulfillmentCenterId() int32`

GetExceptionFulfillmentCenterId returns the ExceptionFulfillmentCenterId field if non-nil, zero value otherwise.

### GetExceptionFulfillmentCenterIdOk

`func (o *OrdersStatusDetailViewModel) GetExceptionFulfillmentCenterIdOk() (*int32, bool)`

GetExceptionFulfillmentCenterIdOk returns a tuple with the ExceptionFulfillmentCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionFulfillmentCenterId

`func (o *OrdersStatusDetailViewModel) SetExceptionFulfillmentCenterId(v int32)`

SetExceptionFulfillmentCenterId sets ExceptionFulfillmentCenterId field to given value.

### HasExceptionFulfillmentCenterId

`func (o *OrdersStatusDetailViewModel) HasExceptionFulfillmentCenterId() bool`

HasExceptionFulfillmentCenterId returns a boolean if a field has been set.

### SetExceptionFulfillmentCenterIdNil

`func (o *OrdersStatusDetailViewModel) SetExceptionFulfillmentCenterIdNil(b bool)`

 SetExceptionFulfillmentCenterIdNil sets the value for ExceptionFulfillmentCenterId to be an explicit nil

### UnsetExceptionFulfillmentCenterId
`func (o *OrdersStatusDetailViewModel) UnsetExceptionFulfillmentCenterId()`

UnsetExceptionFulfillmentCenterId ensures that no value is present for ExceptionFulfillmentCenterId, not even an explicit nil
### GetId

`func (o *OrdersStatusDetailViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersStatusDetailViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersStatusDetailViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersStatusDetailViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryId

`func (o *OrdersStatusDetailViewModel) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *OrdersStatusDetailViewModel) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *OrdersStatusDetailViewModel) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *OrdersStatusDetailViewModel) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### SetInventoryIdNil

`func (o *OrdersStatusDetailViewModel) SetInventoryIdNil(b bool)`

 SetInventoryIdNil sets the value for InventoryId to be an explicit nil

### UnsetInventoryId
`func (o *OrdersStatusDetailViewModel) UnsetInventoryId()`

UnsetInventoryId ensures that no value is present for InventoryId, not even an explicit nil
### GetName

`func (o *OrdersStatusDetailViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrdersStatusDetailViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrdersStatusDetailViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrdersStatusDetailViewModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


