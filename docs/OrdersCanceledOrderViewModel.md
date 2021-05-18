# OrdersCanceledOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledShipmentResults** | Pointer to [**[]OrdersCanceledShipmentViewModel**](OrdersCanceledShipmentViewModel.md) | Results of camceling the shipments associated with the order | [optional] 
**Order** | Pointer to [**OrdersOrderViewModel**](Orders.OrderViewModel.md) |  | [optional] 
**OrderId** | Pointer to **int32** | The ID of the canceled order | [optional] 
**Status** | Pointer to **string** | The overall result of canceling the shipments associated with the order | [optional] 

## Methods

### NewOrdersCanceledOrderViewModel

`func NewOrdersCanceledOrderViewModel() *OrdersCanceledOrderViewModel`

NewOrdersCanceledOrderViewModel instantiates a new OrdersCanceledOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersCanceledOrderViewModelWithDefaults

`func NewOrdersCanceledOrderViewModelWithDefaults() *OrdersCanceledOrderViewModel`

NewOrdersCanceledOrderViewModelWithDefaults instantiates a new OrdersCanceledOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledShipmentResults

`func (o *OrdersCanceledOrderViewModel) GetCanceledShipmentResults() []OrdersCanceledShipmentViewModel`

GetCanceledShipmentResults returns the CanceledShipmentResults field if non-nil, zero value otherwise.

### GetCanceledShipmentResultsOk

`func (o *OrdersCanceledOrderViewModel) GetCanceledShipmentResultsOk() (*[]OrdersCanceledShipmentViewModel, bool)`

GetCanceledShipmentResultsOk returns a tuple with the CanceledShipmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledShipmentResults

`func (o *OrdersCanceledOrderViewModel) SetCanceledShipmentResults(v []OrdersCanceledShipmentViewModel)`

SetCanceledShipmentResults sets CanceledShipmentResults field to given value.

### HasCanceledShipmentResults

`func (o *OrdersCanceledOrderViewModel) HasCanceledShipmentResults() bool`

HasCanceledShipmentResults returns a boolean if a field has been set.

### GetOrder

`func (o *OrdersCanceledOrderViewModel) GetOrder() OrdersOrderViewModel`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrdersCanceledOrderViewModel) GetOrderOk() (*OrdersOrderViewModel, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrdersCanceledOrderViewModel) SetOrder(v OrdersOrderViewModel)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrdersCanceledOrderViewModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderId

`func (o *OrdersCanceledOrderViewModel) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrdersCanceledOrderViewModel) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrdersCanceledOrderViewModel) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrdersCanceledOrderViewModel) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *OrdersCanceledOrderViewModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrdersCanceledOrderViewModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrdersCanceledOrderViewModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrdersCanceledOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


