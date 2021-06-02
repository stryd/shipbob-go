# ProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique id of the product | [optional] 
**Quantity** | Pointer to **int32** | The quantity of this product ordered | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit for the product | [optional] 

## Methods

### NewProductInfo

`func NewProductInfo() *ProductInfo`

NewProductInfo instantiates a new ProductInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInfoWithDefaults

`func NewProductInfoWithDefaults() *ProductInfo`

NewProductInfoWithDefaults instantiates a new ProductInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProductInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProductInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuantity

`func (o *ProductInfo) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductInfo) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductInfo) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReferenceId

`func (o *ProductInfo) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ProductInfo) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ProductInfo) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ProductInfo) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSku

`func (o *ProductInfo) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductInfo) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductInfo) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductInfo) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


