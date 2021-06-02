# EstimateFulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**EstimationAddress**](EstimationAddress.md) |  | 
**Products** | [**[]EstimateFulfillmentRequestProducts**](EstimateFulfillmentRequestProducts.md) | Products to be included in the order. Each product must include one of reference_id or id | 
**ShippingMethods** | Pointer to **[]string** | Array of strings specifying shipping methods for which to fetch estimates.\\r\\nIf this field is omitted we will return estimates for all shipping methods defined in ShipBob | [optional] 

## Methods

### NewEstimateFulfillmentRequest

`func NewEstimateFulfillmentRequest(address EstimationAddress, products []EstimateFulfillmentRequestProducts, ) *EstimateFulfillmentRequest`

NewEstimateFulfillmentRequest instantiates a new EstimateFulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateFulfillmentRequestWithDefaults

`func NewEstimateFulfillmentRequestWithDefaults() *EstimateFulfillmentRequest`

NewEstimateFulfillmentRequestWithDefaults instantiates a new EstimateFulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EstimateFulfillmentRequest) GetAddress() EstimationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EstimateFulfillmentRequest) GetAddressOk() (*EstimationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EstimateFulfillmentRequest) SetAddress(v EstimationAddress)`

SetAddress sets Address field to given value.


### GetProducts

`func (o *EstimateFulfillmentRequest) GetProducts() []EstimateFulfillmentRequestProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EstimateFulfillmentRequest) GetProductsOk() (*[]EstimateFulfillmentRequestProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EstimateFulfillmentRequest) SetProducts(v []EstimateFulfillmentRequestProducts)`

SetProducts sets Products field to given value.


### GetShippingMethods

`func (o *EstimateFulfillmentRequest) GetShippingMethods() []string`

GetShippingMethods returns the ShippingMethods field if non-nil, zero value otherwise.

### GetShippingMethodsOk

`func (o *EstimateFulfillmentRequest) GetShippingMethodsOk() (*[]string, bool)`

GetShippingMethodsOk returns a tuple with the ShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethods

`func (o *EstimateFulfillmentRequest) SetShippingMethods(v []string)`

SetShippingMethods sets ShippingMethods field to given value.

### HasShippingMethods

`func (o *EstimateFulfillmentRequest) HasShippingMethods() bool`

HasShippingMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


