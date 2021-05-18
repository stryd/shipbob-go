# OrdersShipMethodDetailViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the shipping method is active | [optional] 
**Default** | Pointer to **bool** | Indicates the shipping method is a ShipBob default shipping method. | [optional] 
**Id** | Pointer to **int32** | Unique id for shipping method. | [optional] 
**Name** | Pointer to **string** | Name of the ship method as selected by the merchant and saved in ShipBobâs database (i.e. âgroundâ). Corresponds to the shipping_method field in the Orders API. | [optional] 
**ServiceLevel** | Pointer to [**OrdersServiceLevelDetailViewModel**](Orders.ServiceLevelDetailViewModel.md) |  | [optional] 

## Methods

### NewOrdersShipMethodDetailViewModel

`func NewOrdersShipMethodDetailViewModel() *OrdersShipMethodDetailViewModel`

NewOrdersShipMethodDetailViewModel instantiates a new OrdersShipMethodDetailViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersShipMethodDetailViewModelWithDefaults

`func NewOrdersShipMethodDetailViewModelWithDefaults() *OrdersShipMethodDetailViewModel`

NewOrdersShipMethodDetailViewModelWithDefaults instantiates a new OrdersShipMethodDetailViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *OrdersShipMethodDetailViewModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrdersShipMethodDetailViewModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrdersShipMethodDetailViewModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrdersShipMethodDetailViewModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefault

`func (o *OrdersShipMethodDetailViewModel) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OrdersShipMethodDetailViewModel) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OrdersShipMethodDetailViewModel) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OrdersShipMethodDetailViewModel) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *OrdersShipMethodDetailViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersShipMethodDetailViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersShipMethodDetailViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersShipMethodDetailViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrdersShipMethodDetailViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrdersShipMethodDetailViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrdersShipMethodDetailViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrdersShipMethodDetailViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceLevel

`func (o *OrdersShipMethodDetailViewModel) GetServiceLevel() OrdersServiceLevelDetailViewModel`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *OrdersShipMethodDetailViewModel) GetServiceLevelOk() (*OrdersServiceLevelDetailViewModel, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *OrdersShipMethodDetailViewModel) SetServiceLevel(v OrdersServiceLevelDetailViewModel)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *OrdersShipMethodDetailViewModel) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


