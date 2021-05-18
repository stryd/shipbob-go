# ReturnsTransactionViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | Transaction amount in dollars | [optional] 
**TransactionType** | Pointer to [**ReturnsTransactionLogSource**](Returns.TransactionLogSource.md) |  | [optional] 

## Methods

### NewReturnsTransactionViewModel

`func NewReturnsTransactionViewModel() *ReturnsTransactionViewModel`

NewReturnsTransactionViewModel instantiates a new ReturnsTransactionViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnsTransactionViewModelWithDefaults

`func NewReturnsTransactionViewModelWithDefaults() *ReturnsTransactionViewModel`

NewReturnsTransactionViewModelWithDefaults instantiates a new ReturnsTransactionViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ReturnsTransactionViewModel) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReturnsTransactionViewModel) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReturnsTransactionViewModel) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ReturnsTransactionViewModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *ReturnsTransactionViewModel) GetTransactionType() ReturnsTransactionLogSource`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ReturnsTransactionViewModel) GetTransactionTypeOk() (*ReturnsTransactionLogSource, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ReturnsTransactionViewModel) SetTransactionType(v ReturnsTransactionLogSource)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ReturnsTransactionViewModel) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


