# ShipmentProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the product | [optional] 
**InventoryItems** | Pointer to [**[]OrderInventory**](OrderInventory.md) | Information about fulfillable inventory items belonging to this product | [optional] 
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit for the product | [optional] 

## Methods

### NewShipmentProduct

`func NewShipmentProduct() *ShipmentProduct`

NewShipmentProduct instantiates a new ShipmentProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentProductWithDefaults

`func NewShipmentProductWithDefaults() *ShipmentProduct`

NewShipmentProductWithDefaults instantiates a new ShipmentProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentProduct) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryItems

`func (o *ShipmentProduct) GetInventoryItems() []OrderInventory`

GetInventoryItems returns the InventoryItems field if non-nil, zero value otherwise.

### GetInventoryItemsOk

`func (o *ShipmentProduct) GetInventoryItemsOk() (*[]OrderInventory, bool)`

GetInventoryItemsOk returns a tuple with the InventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItems

`func (o *ShipmentProduct) SetInventoryItems(v []OrderInventory)`

SetInventoryItems sets InventoryItems field to given value.

### HasInventoryItems

`func (o *ShipmentProduct) HasInventoryItems() bool`

HasInventoryItems returns a boolean if a field has been set.

### GetName

`func (o *ShipmentProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceId

`func (o *ShipmentProduct) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipmentProduct) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipmentProduct) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipmentProduct) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSku

`func (o *ShipmentProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ShipmentProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ShipmentProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ShipmentProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


