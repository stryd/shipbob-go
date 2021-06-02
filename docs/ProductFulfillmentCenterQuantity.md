# ProductFulfillmentCenterQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommittedQuantity** | Pointer to **int32** | Amount of committed quantity at this fulfillment center | [optional] 
**FulfillableQuantity** | Pointer to **int32** | Amount of fulfillable quantity at this fulfillment center | [optional] 
**Id** | Pointer to **int32** | Unique id of the fulfillment center | [optional] 
**Name** | Pointer to **NullableString** | Name of the fulfillment center | [optional] 
**OnhandQuantity** | Pointer to **int32** | Amount of onhand quantity at this fulfillment center | [optional] 

## Methods

### NewProductFulfillmentCenterQuantity

`func NewProductFulfillmentCenterQuantity() *ProductFulfillmentCenterQuantity`

NewProductFulfillmentCenterQuantity instantiates a new ProductFulfillmentCenterQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductFulfillmentCenterQuantityWithDefaults

`func NewProductFulfillmentCenterQuantityWithDefaults() *ProductFulfillmentCenterQuantity`

NewProductFulfillmentCenterQuantityWithDefaults instantiates a new ProductFulfillmentCenterQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommittedQuantity

`func (o *ProductFulfillmentCenterQuantity) GetCommittedQuantity() int32`

GetCommittedQuantity returns the CommittedQuantity field if non-nil, zero value otherwise.

### GetCommittedQuantityOk

`func (o *ProductFulfillmentCenterQuantity) GetCommittedQuantityOk() (*int32, bool)`

GetCommittedQuantityOk returns a tuple with the CommittedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedQuantity

`func (o *ProductFulfillmentCenterQuantity) SetCommittedQuantity(v int32)`

SetCommittedQuantity sets CommittedQuantity field to given value.

### HasCommittedQuantity

`func (o *ProductFulfillmentCenterQuantity) HasCommittedQuantity() bool`

HasCommittedQuantity returns a boolean if a field has been set.

### GetFulfillableQuantity

`func (o *ProductFulfillmentCenterQuantity) GetFulfillableQuantity() int32`

GetFulfillableQuantity returns the FulfillableQuantity field if non-nil, zero value otherwise.

### GetFulfillableQuantityOk

`func (o *ProductFulfillmentCenterQuantity) GetFulfillableQuantityOk() (*int32, bool)`

GetFulfillableQuantityOk returns a tuple with the FulfillableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillableQuantity

`func (o *ProductFulfillmentCenterQuantity) SetFulfillableQuantity(v int32)`

SetFulfillableQuantity sets FulfillableQuantity field to given value.

### HasFulfillableQuantity

`func (o *ProductFulfillmentCenterQuantity) HasFulfillableQuantity() bool`

HasFulfillableQuantity returns a boolean if a field has been set.

### GetId

`func (o *ProductFulfillmentCenterQuantity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductFulfillmentCenterQuantity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductFulfillmentCenterQuantity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductFulfillmentCenterQuantity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductFulfillmentCenterQuantity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductFulfillmentCenterQuantity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductFulfillmentCenterQuantity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductFulfillmentCenterQuantity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductFulfillmentCenterQuantity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductFulfillmentCenterQuantity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOnhandQuantity

`func (o *ProductFulfillmentCenterQuantity) GetOnhandQuantity() int32`

GetOnhandQuantity returns the OnhandQuantity field if non-nil, zero value otherwise.

### GetOnhandQuantityOk

`func (o *ProductFulfillmentCenterQuantity) GetOnhandQuantityOk() (*int32, bool)`

GetOnhandQuantityOk returns a tuple with the OnhandQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnhandQuantity

`func (o *ProductFulfillmentCenterQuantity) SetOnhandQuantity(v int32)`

SetOnhandQuantity sets OnhandQuantity field to given value.

### HasOnhandQuantity

`func (o *ProductFulfillmentCenterQuantity) HasOnhandQuantity() bool`

HasOnhandQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


