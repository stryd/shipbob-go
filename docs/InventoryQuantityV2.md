# InventoryQuantityV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedQuantity** | Pointer to **int32** | Quantity of the inventory item submitted in the WRO | [optional] 
**InventoryId** | Pointer to **int32** | ID of the inventory item | [optional] 
**ReceivedQuantity** | Pointer to **int32** | Quantity of the inventory item received by the warehouse | [optional] 
**Sku** | Pointer to **NullableString** | Sku of the inventory item | [optional] 

## Methods

### NewInventoryQuantityV2

`func NewInventoryQuantityV2() *InventoryQuantityV2`

NewInventoryQuantityV2 instantiates a new InventoryQuantityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryQuantityV2WithDefaults

`func NewInventoryQuantityV2WithDefaults() *InventoryQuantityV2`

NewInventoryQuantityV2WithDefaults instantiates a new InventoryQuantityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedQuantity

`func (o *InventoryQuantityV2) GetExpectedQuantity() int32`

GetExpectedQuantity returns the ExpectedQuantity field if non-nil, zero value otherwise.

### GetExpectedQuantityOk

`func (o *InventoryQuantityV2) GetExpectedQuantityOk() (*int32, bool)`

GetExpectedQuantityOk returns a tuple with the ExpectedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedQuantity

`func (o *InventoryQuantityV2) SetExpectedQuantity(v int32)`

SetExpectedQuantity sets ExpectedQuantity field to given value.

### HasExpectedQuantity

`func (o *InventoryQuantityV2) HasExpectedQuantity() bool`

HasExpectedQuantity returns a boolean if a field has been set.

### GetInventoryId

`func (o *InventoryQuantityV2) GetInventoryId() int32`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *InventoryQuantityV2) GetInventoryIdOk() (*int32, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *InventoryQuantityV2) SetInventoryId(v int32)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *InventoryQuantityV2) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetReceivedQuantity

`func (o *InventoryQuantityV2) GetReceivedQuantity() int32`

GetReceivedQuantity returns the ReceivedQuantity field if non-nil, zero value otherwise.

### GetReceivedQuantityOk

`func (o *InventoryQuantityV2) GetReceivedQuantityOk() (*int32, bool)`

GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQuantity

`func (o *InventoryQuantityV2) SetReceivedQuantity(v int32)`

SetReceivedQuantity sets ReceivedQuantity field to given value.

### HasReceivedQuantity

`func (o *InventoryQuantityV2) HasReceivedQuantity() bool`

HasReceivedQuantity returns a boolean if a field has been set.

### GetSku

`func (o *InventoryQuantityV2) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *InventoryQuantityV2) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *InventoryQuantityV2) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *InventoryQuantityV2) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *InventoryQuantityV2) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *InventoryQuantityV2) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


