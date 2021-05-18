# OrdersCanceledShipmentViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**IsSuccess** | Pointer to **bool** | If the cancel action was successfull | [optional] 
**Reason** | Pointer to **string** | The reason the cancellation result | [optional] 
**ShipmentId** | Pointer to **int64** | The ID of the shipment | [optional] 

## Methods

### NewOrdersCanceledShipmentViewModel

`func NewOrdersCanceledShipmentViewModel() *OrdersCanceledShipmentViewModel`

NewOrdersCanceledShipmentViewModel instantiates a new OrdersCanceledShipmentViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersCanceledShipmentViewModelWithDefaults

`func NewOrdersCanceledShipmentViewModelWithDefaults() *OrdersCanceledShipmentViewModel`

NewOrdersCanceledShipmentViewModelWithDefaults instantiates a new OrdersCanceledShipmentViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrdersCanceledShipmentViewModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrdersCanceledShipmentViewModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrdersCanceledShipmentViewModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrdersCanceledShipmentViewModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsSuccess

`func (o *OrdersCanceledShipmentViewModel) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *OrdersCanceledShipmentViewModel) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *OrdersCanceledShipmentViewModel) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *OrdersCanceledShipmentViewModel) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetReason

`func (o *OrdersCanceledShipmentViewModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OrdersCanceledShipmentViewModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OrdersCanceledShipmentViewModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OrdersCanceledShipmentViewModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShipmentId

`func (o *OrdersCanceledShipmentViewModel) GetShipmentId() int64`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *OrdersCanceledShipmentViewModel) GetShipmentIdOk() (*int64, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *OrdersCanceledShipmentViewModel) SetShipmentId(v int64)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *OrdersCanceledShipmentViewModel) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


