# ShipMethodDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the shipping method is active | [optional] 
**Default** | Pointer to **bool** | Indicates the shipping method is a ShipBob default shipping method. | [optional] 
**Id** | Pointer to **int32** | Unique id for shipping method. | [optional] 
**Name** | Pointer to **string** | Name of the ship method as selected by the merchant and saved in ShipBob&#39;s database (i.e. \&quot;ground\&quot;). Corresponds to the shipping_method field in the Orders API. | [optional] 
**ServiceLevel** | Pointer to [**ServiceLevelDetail**](ServiceLevelDetail.md) |  | [optional] 

## Methods

### NewShipMethodDetail

`func NewShipMethodDetail() *ShipMethodDetail`

NewShipMethodDetail instantiates a new ShipMethodDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipMethodDetailWithDefaults

`func NewShipMethodDetailWithDefaults() *ShipMethodDetail`

NewShipMethodDetailWithDefaults instantiates a new ShipMethodDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ShipMethodDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ShipMethodDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ShipMethodDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ShipMethodDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefault

`func (o *ShipMethodDetail) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ShipMethodDetail) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ShipMethodDetail) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ShipMethodDetail) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *ShipMethodDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipMethodDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipMethodDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipMethodDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShipMethodDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipMethodDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipMethodDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipMethodDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceLevel

`func (o *ShipMethodDetail) GetServiceLevel() ServiceLevelDetail`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *ShipMethodDetail) GetServiceLevelOk() (*ServiceLevelDetail, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *ShipMethodDetail) SetServiceLevel(v ServiceLevelDetail)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *ShipMethodDetail) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


