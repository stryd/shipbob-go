# OrdersRecipientInfoViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**OrdersAddressViewModel**](Orders.AddressViewModel.md) |  | 
**Email** | Pointer to **string** | Email address of the recipient | [optional] 
**Name** | **string** | Name of the recipient | 
**PhoneNumber** | Pointer to **string** | Phone number of the recipient | [optional] 

## Methods

### NewOrdersRecipientInfoViewModel

`func NewOrdersRecipientInfoViewModel(address OrdersAddressViewModel, name string, ) *OrdersRecipientInfoViewModel`

NewOrdersRecipientInfoViewModel instantiates a new OrdersRecipientInfoViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersRecipientInfoViewModelWithDefaults

`func NewOrdersRecipientInfoViewModelWithDefaults() *OrdersRecipientInfoViewModel`

NewOrdersRecipientInfoViewModelWithDefaults instantiates a new OrdersRecipientInfoViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrdersRecipientInfoViewModel) GetAddress() OrdersAddressViewModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrdersRecipientInfoViewModel) GetAddressOk() (*OrdersAddressViewModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrdersRecipientInfoViewModel) SetAddress(v OrdersAddressViewModel)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *OrdersRecipientInfoViewModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrdersRecipientInfoViewModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrdersRecipientInfoViewModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrdersRecipientInfoViewModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *OrdersRecipientInfoViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrdersRecipientInfoViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrdersRecipientInfoViewModel) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *OrdersRecipientInfoViewModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OrdersRecipientInfoViewModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OrdersRecipientInfoViewModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OrdersRecipientInfoViewModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


