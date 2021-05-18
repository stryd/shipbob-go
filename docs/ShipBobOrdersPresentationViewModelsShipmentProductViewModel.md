# ShipBobOrdersPresentationViewModelsShipmentProductViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique id of the product | [optional] 
**InventoryItems** | Pointer to [**[]ShipBobOrdersPresentationViewModelsInventoryViewModel**](ShipBobOrdersPresentationViewModelsInventoryViewModel.md) | Information about fulfillable inventory items belonging to this product | [optional] 
**Name** | Pointer to **string** | Name of the product | [optional] 
**ReferenceId** | Pointer to **string** | Unique reference id of the product | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit for the product | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsShipmentProductViewModel

`func NewShipBobOrdersPresentationViewModelsShipmentProductViewModel() *ShipBobOrdersPresentationViewModelsShipmentProductViewModel`

NewShipBobOrdersPresentationViewModelsShipmentProductViewModel instantiates a new ShipBobOrdersPresentationViewModelsShipmentProductViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsShipmentProductViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsShipmentProductViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsShipmentProductViewModel`

NewShipBobOrdersPresentationViewModelsShipmentProductViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsShipmentProductViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInventoryItems

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetInventoryItems() []ShipBobOrdersPresentationViewModelsInventoryViewModel`

GetInventoryItems returns the InventoryItems field if non-nil, zero value otherwise.

### GetInventoryItemsOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetInventoryItemsOk() (*[]ShipBobOrdersPresentationViewModelsInventoryViewModel, bool)`

GetInventoryItemsOk returns a tuple with the InventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItems

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) SetInventoryItems(v []ShipBobOrdersPresentationViewModelsInventoryViewModel)`

SetInventoryItems sets InventoryItems field to given value.

### HasInventoryItems

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) HasInventoryItems() bool`

HasInventoryItems returns a boolean if a field has been set.

### GetName

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSku

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ShipBobOrdersPresentationViewModelsShipmentProductViewModel) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


