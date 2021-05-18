# ShipBobOrdersPresentationViewModelsEstimateDetailViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedPrice** | Pointer to **float64** | Estimated price in dollars for the provided shipping method | [optional] 
**FulfillmentCenter** | Pointer to [**ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel**](ShipBob.Orders.Presentation.ViewModels.FulfillmentCenterViewModel.md) |  | [optional] 
**ShippingMethod** | Pointer to **string** | Provided shipping method. Maps to ship option in ShipBob. | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsEstimateDetailViewModel

`func NewShipBobOrdersPresentationViewModelsEstimateDetailViewModel() *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel`

NewShipBobOrdersPresentationViewModelsEstimateDetailViewModel instantiates a new ShipBobOrdersPresentationViewModelsEstimateDetailViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsEstimateDetailViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsEstimateDetailViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel`

NewShipBobOrdersPresentationViewModelsEstimateDetailViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsEstimateDetailViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedPrice

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetEstimatedPrice() float64`

GetEstimatedPrice returns the EstimatedPrice field if non-nil, zero value otherwise.

### GetEstimatedPriceOk

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetEstimatedPriceOk() (*float64, bool)`

GetEstimatedPriceOk returns a tuple with the EstimatedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPrice

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) SetEstimatedPrice(v float64)`

SetEstimatedPrice sets EstimatedPrice field to given value.

### HasEstimatedPrice

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) HasEstimatedPrice() bool`

HasEstimatedPrice returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetFulfillmentCenter() ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetFulfillmentCenterOk() (*ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) SetFulfillmentCenter(v ShipBobOrdersPresentationViewModelsFulfillmentCenterViewModel)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *ShipBobOrdersPresentationViewModelsEstimateDetailViewModel) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


