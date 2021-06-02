# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**OrderAddress**](OrderAddress.md) |  | [optional] 
**Email** | Pointer to **string** | Email address of the recipient | [optional] 
**Name** | Pointer to **string** | Name of the recipient | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number of the recipient | [optional] 

## Methods

### NewRecipient

`func NewRecipient() *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Recipient) GetAddress() OrderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Recipient) GetAddressOk() (*OrderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Recipient) SetAddress(v OrderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Recipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *Recipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Recipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Recipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Recipient) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Recipient) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Recipient) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Recipient) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


