# ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**ShipBobOrdersPresentationViewModelsEstimationAddressViewModel**](ShipBob.Orders.Presentation.ViewModels.EstimationAddressViewModel.md) |  | 
**Products** | [**[]ShipBobOrdersPresentationModelsEstimateProductInfoModel**](ShipBobOrdersPresentationModelsEstimateProductInfoModel.md) | Products to be included in the order. Each product must include one of reference_id or id | 
**ShippingMethods** | Pointer to **[]string** | Array of strings specifying shipping methods for which to fetch estimates.\\r\\n\\r\\nIf this field is omitted we will return estimates for all shipping methods defined in ShipBob | [optional] 

## Methods

### NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel

`func NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel(address ShipBobOrdersPresentationViewModelsEstimationAddressViewModel, products []ShipBobOrdersPresentationModelsEstimateProductInfoModel, ) *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel`

NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel instantiates a new ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModelWithDefaults

`func NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModelWithDefaults() *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel`

NewShipBobOrdersPresentationModelsEstimateFulfillmentRequestModelWithDefaults instantiates a new ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetAddress() ShipBobOrdersPresentationViewModelsEstimationAddressViewModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetAddressOk() (*ShipBobOrdersPresentationViewModelsEstimationAddressViewModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) SetAddress(v ShipBobOrdersPresentationViewModelsEstimationAddressViewModel)`

SetAddress sets Address field to given value.


### GetProducts

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetProducts() []ShipBobOrdersPresentationModelsEstimateProductInfoModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetProductsOk() (*[]ShipBobOrdersPresentationModelsEstimateProductInfoModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) SetProducts(v []ShipBobOrdersPresentationModelsEstimateProductInfoModel)`

SetProducts sets Products field to given value.


### GetShippingMethods

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetShippingMethods() []string`

GetShippingMethods returns the ShippingMethods field if non-nil, zero value otherwise.

### GetShippingMethodsOk

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) GetShippingMethodsOk() (*[]string, bool)`

GetShippingMethodsOk returns a tuple with the ShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethods

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) SetShippingMethods(v []string)`

SetShippingMethods sets ShippingMethods field to given value.

### HasShippingMethods

`func (o *ShipBobOrdersPresentationModelsEstimateFulfillmentRequestModel) HasShippingMethods() bool`

HasShippingMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


