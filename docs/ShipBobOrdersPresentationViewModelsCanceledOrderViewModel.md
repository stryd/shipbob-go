# ShipBobOrdersPresentationViewModelsCanceledOrderViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledShipmentResults** | Pointer to [**[]ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel**](ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel.md) | Results of camceling the shipments associated with the order | [optional] 
**Order** | Pointer to [**ShipBobOrdersPresentationViewModelsOrderViewModel**](ShipBob.Orders.Presentation.ViewModels.OrderViewModel.md) |  | [optional] 
**OrderId** | Pointer to **int32** | The ID of the canceled order | [optional] 
**Status** | Pointer to **string** | The overall result of canceling the shipments associated with the order | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsCanceledOrderViewModel

`func NewShipBobOrdersPresentationViewModelsCanceledOrderViewModel() *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel`

NewShipBobOrdersPresentationViewModelsCanceledOrderViewModel instantiates a new ShipBobOrdersPresentationViewModelsCanceledOrderViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsCanceledOrderViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsCanceledOrderViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel`

NewShipBobOrdersPresentationViewModelsCanceledOrderViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsCanceledOrderViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledShipmentResults

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetCanceledShipmentResults() []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel`

GetCanceledShipmentResults returns the CanceledShipmentResults field if non-nil, zero value otherwise.

### GetCanceledShipmentResultsOk

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetCanceledShipmentResultsOk() (*[]ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel, bool)`

GetCanceledShipmentResultsOk returns a tuple with the CanceledShipmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledShipmentResults

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetCanceledShipmentResults(v []ShipBobOrdersPresentationViewModelsCanceledShipmentViewModel)`

SetCanceledShipmentResults sets CanceledShipmentResults field to given value.

### HasCanceledShipmentResults

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasCanceledShipmentResults() bool`

HasCanceledShipmentResults returns a boolean if a field has been set.

### GetOrder

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrder() ShipBobOrdersPresentationViewModelsOrderViewModel`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderOk() (*ShipBobOrdersPresentationViewModelsOrderViewModel, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetOrder(v ShipBobOrdersPresentationViewModelsOrderViewModel)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderId

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipBobOrdersPresentationViewModelsCanceledOrderViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


