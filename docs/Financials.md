# Financials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPrice** | Pointer to **NullableFloat64** | Sum of all line item prices, discounts, and taxes in USD | [optional] 

## Methods

### NewFinancials

`func NewFinancials() *Financials`

NewFinancials instantiates a new Financials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialsWithDefaults

`func NewFinancialsWithDefaults() *Financials`

NewFinancialsWithDefaults instantiates a new Financials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPrice

`func (o *Financials) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Financials) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Financials) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *Financials) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### SetTotalPriceNil

`func (o *Financials) SetTotalPriceNil(b bool)`

 SetTotalPriceNil sets the value for TotalPrice to be an explicit nil

### UnsetTotalPrice
`func (o *Financials) UnsetTotalPrice()`

UnsetTotalPrice ensures that no value is present for TotalPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


