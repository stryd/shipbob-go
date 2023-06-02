# OrderEstimateDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedPrice** | Pointer to **float64** | Estimated price in dollars for the provided shipping method | [optional] 
**FulfillmentCenter** | Pointer to [**FulfillmentCenter**](FulfillmentCenter.md) |  | [optional] 
**ShippingMethod** | Pointer to **string** | Provided shipping method. Maps to ship option in ShipBob. | [optional] 

## Methods

### NewOrderEstimateDetail

`func NewOrderEstimateDetail() *OrderEstimateDetail`

NewOrderEstimateDetail instantiates a new OrderEstimateDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderEstimateDetailWithDefaults

`func NewOrderEstimateDetailWithDefaults() *OrderEstimateDetail`

NewOrderEstimateDetailWithDefaults instantiates a new OrderEstimateDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedPrice

`func (o *OrderEstimateDetail) GetEstimatedPrice() float64`

GetEstimatedPrice returns the EstimatedPrice field if non-nil, zero value otherwise.

### GetEstimatedPriceOk

`func (o *OrderEstimateDetail) GetEstimatedPriceOk() (*float64, bool)`

GetEstimatedPriceOk returns a tuple with the EstimatedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPrice

`func (o *OrderEstimateDetail) SetEstimatedPrice(v float64)`

SetEstimatedPrice sets EstimatedPrice field to given value.

### HasEstimatedPrice

`func (o *OrderEstimateDetail) HasEstimatedPrice() bool`

HasEstimatedPrice returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *OrderEstimateDetail) GetFulfillmentCenter() FulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *OrderEstimateDetail) GetFulfillmentCenterOk() (*FulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *OrderEstimateDetail) SetFulfillmentCenter(v FulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *OrderEstimateDetail) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetShippingMethod

`func (o *OrderEstimateDetail) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *OrderEstimateDetail) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *OrderEstimateDetail) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *OrderEstimateDetail) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


