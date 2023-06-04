# CartonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ShipmentProduct**](ShipmentProduct.md) | List of what is packed in this carton | [optional] 

## Methods

### NewCartonDetails

`func NewCartonDetails() *CartonDetails`

NewCartonDetails instantiates a new CartonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartonDetailsWithDefaults

`func NewCartonDetailsWithDefaults() *CartonDetails`

NewCartonDetailsWithDefaults instantiates a new CartonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CartonDetails) GetProducts() []ShipmentProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CartonDetails) GetProductsOk() (*[]ShipmentProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CartonDetails) SetProducts(v []ShipmentProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CartonDetails) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


