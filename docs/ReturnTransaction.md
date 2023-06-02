# ReturnTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | Transaction amount in dollars | [optional] 
**TransactionType** | Pointer to [**ReturnTransactionLogSource**](ReturnTransactionLogSource.md) |  | [optional] 

## Methods

### NewReturnTransaction

`func NewReturnTransaction() *ReturnTransaction`

NewReturnTransaction instantiates a new ReturnTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnTransactionWithDefaults

`func NewReturnTransactionWithDefaults() *ReturnTransaction`

NewReturnTransactionWithDefaults instantiates a new ReturnTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ReturnTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReturnTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReturnTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ReturnTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *ReturnTransaction) GetTransactionType() ReturnTransactionLogSource`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ReturnTransaction) GetTransactionTypeOk() (*ReturnTransactionLogSource, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ReturnTransaction) SetTransactionType(v ReturnTransactionLogSource)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ReturnTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


