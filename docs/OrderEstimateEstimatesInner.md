# OrderEstimateEstimatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedPrice** | Pointer to **float64** | Estimated price in dollars for the provided shipping method | [optional] 
**FulfillmentCenter** | Pointer to [**FulfillmentCenter**](FulfillmentCenter.md) |  | [optional] 
**ShippingMethod** | Pointer to **string** | Provided shipping method. Maps to ship option in ShipBob. | [optional] 

## Methods

### NewOrderEstimateEstimatesInner

`func NewOrderEstimateEstimatesInner() *OrderEstimateEstimatesInner`

NewOrderEstimateEstimatesInner instantiates a new OrderEstimateEstimatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderEstimateEstimatesInnerWithDefaults

`func NewOrderEstimateEstimatesInnerWithDefaults() *OrderEstimateEstimatesInner`

NewOrderEstimateEstimatesInnerWithDefaults instantiates a new OrderEstimateEstimatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedPrice

`func (o *OrderEstimateEstimatesInner) GetEstimatedPrice() float64`

GetEstimatedPrice returns the EstimatedPrice field if non-nil, zero value otherwise.

### GetEstimatedPriceOk

`func (o *OrderEstimateEstimatesInner) GetEstimatedPriceOk() (*float64, bool)`

GetEstimatedPriceOk returns a tuple with the EstimatedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPrice

`func (o *OrderEstimateEstimatesInner) SetEstimatedPrice(v float64)`

SetEstimatedPrice sets EstimatedPrice field to given value.

### HasEstimatedPrice

`func (o *OrderEstimateEstimatesInner) HasEstimatedPrice() bool`

HasEstimatedPrice returns a boolean if a field has been set.

### GetFulfillmentCenter

`func (o *OrderEstimateEstimatesInner) GetFulfillmentCenter() FulfillmentCenter`

GetFulfillmentCenter returns the FulfillmentCenter field if non-nil, zero value otherwise.

### GetFulfillmentCenterOk

`func (o *OrderEstimateEstimatesInner) GetFulfillmentCenterOk() (*FulfillmentCenter, bool)`

GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentCenter

`func (o *OrderEstimateEstimatesInner) SetFulfillmentCenter(v FulfillmentCenter)`

SetFulfillmentCenter sets FulfillmentCenter field to given value.

### HasFulfillmentCenter

`func (o *OrderEstimateEstimatesInner) HasFulfillmentCenter() bool`

HasFulfillmentCenter returns a boolean if a field has been set.

### GetShippingMethod

`func (o *OrderEstimateEstimatesInner) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *OrderEstimateEstimatesInner) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *OrderEstimateEstimatesInner) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *OrderEstimateEstimatesInner) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


