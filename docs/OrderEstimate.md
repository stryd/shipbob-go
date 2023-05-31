# OrderEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Estimates** | Pointer to [**[]OrderEstimateEstimatesInner**](OrderEstimateEstimatesInner.md) | Array of estimates for each shipping method | [optional] 

## Methods

### NewOrderEstimate

`func NewOrderEstimate() *OrderEstimate`

NewOrderEstimate instantiates a new OrderEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderEstimateWithDefaults

`func NewOrderEstimateWithDefaults() *OrderEstimate`

NewOrderEstimateWithDefaults instantiates a new OrderEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimates

`func (o *OrderEstimate) GetEstimates() []OrderEstimateEstimatesInner`

GetEstimates returns the Estimates field if non-nil, zero value otherwise.

### GetEstimatesOk

`func (o *OrderEstimate) GetEstimatesOk() (*[]OrderEstimateEstimatesInner, bool)`

GetEstimatesOk returns a tuple with the Estimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimates

`func (o *OrderEstimate) SetEstimates(v []OrderEstimateEstimatesInner)`

SetEstimates sets Estimates field to given value.

### HasEstimates

`func (o *OrderEstimate) HasEstimates() bool`

HasEstimates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


