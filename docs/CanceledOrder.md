# CanceledOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledShipmentResults** | Pointer to [**[]CanceledShipment**](CanceledShipment.md) | Results of camceling the shipments associated with the order | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**OrderId** | Pointer to **int32** | The ID of the canceled order | [optional] 
**Status** | Pointer to **string** | The overall result of canceling the shipments associated with the order | [optional] 

## Methods

### NewCanceledOrder

`func NewCanceledOrder() *CanceledOrder`

NewCanceledOrder instantiates a new CanceledOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanceledOrderWithDefaults

`func NewCanceledOrderWithDefaults() *CanceledOrder`

NewCanceledOrderWithDefaults instantiates a new CanceledOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledShipmentResults

`func (o *CanceledOrder) GetCanceledShipmentResults() []CanceledShipment`

GetCanceledShipmentResults returns the CanceledShipmentResults field if non-nil, zero value otherwise.

### GetCanceledShipmentResultsOk

`func (o *CanceledOrder) GetCanceledShipmentResultsOk() (*[]CanceledShipment, bool)`

GetCanceledShipmentResultsOk returns a tuple with the CanceledShipmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledShipmentResults

`func (o *CanceledOrder) SetCanceledShipmentResults(v []CanceledShipment)`

SetCanceledShipmentResults sets CanceledShipmentResults field to given value.

### HasCanceledShipmentResults

`func (o *CanceledOrder) HasCanceledShipmentResults() bool`

HasCanceledShipmentResults returns a boolean if a field has been set.

### GetOrder

`func (o *CanceledOrder) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CanceledOrder) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CanceledOrder) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CanceledOrder) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderId

`func (o *CanceledOrder) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CanceledOrder) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CanceledOrder) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CanceledOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *CanceledOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CanceledOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CanceledOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CanceledOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


