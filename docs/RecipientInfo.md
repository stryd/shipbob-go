# RecipientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**RetailerProgramDataAddress**](RetailerProgramDataAddress.md) |  | 
**Email** | Pointer to **string** | Email address of the recipient | [optional] 
**Name** | **string** | Name of the recipient | 
**PhoneNumber** | Pointer to **string** | Phone number of the recipient | [optional] 

## Methods

### NewRecipientInfo

`func NewRecipientInfo(address RetailerProgramDataAddress, name string, ) *RecipientInfo`

NewRecipientInfo instantiates a new RecipientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientInfoWithDefaults

`func NewRecipientInfoWithDefaults() *RecipientInfo`

NewRecipientInfoWithDefaults instantiates a new RecipientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RecipientInfo) GetAddress() RetailerProgramDataAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RecipientInfo) GetAddressOk() (*RetailerProgramDataAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RecipientInfo) SetAddress(v RetailerProgramDataAddress)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *RecipientInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RecipientInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RecipientInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RecipientInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *RecipientInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecipientInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecipientInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *RecipientInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RecipientInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RecipientInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RecipientInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


