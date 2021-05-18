# LocationAddressViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **NullableString** | First part of the address of the location for this service | [optional] 
**Address2** | Pointer to **NullableString** | Second part of the address of the location for this service | [optional] 
**City** | Pointer to **NullableString** | City of the location | [optional] 
**Country** | Pointer to **NullableString** | Country of the location | [optional] 
**Email** | Pointer to **NullableString** | Email of the location for this service | [optional] 
**Name** | Pointer to **NullableString** | Name to use in the address of the location for this service | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Phone Number of the location for this service | [optional] 
**State** | Pointer to **NullableString** | State of the location | [optional] 
**ZipCode** | Pointer to **NullableString** | Zip code of the location | [optional] 

## Methods

### NewLocationAddressViewModel

`func NewLocationAddressViewModel() *LocationAddressViewModel`

NewLocationAddressViewModel instantiates a new LocationAddressViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAddressViewModelWithDefaults

`func NewLocationAddressViewModelWithDefaults() *LocationAddressViewModel`

NewLocationAddressViewModelWithDefaults instantiates a new LocationAddressViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *LocationAddressViewModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationAddressViewModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationAddressViewModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationAddressViewModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *LocationAddressViewModel) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *LocationAddressViewModel) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *LocationAddressViewModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationAddressViewModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationAddressViewModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationAddressViewModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *LocationAddressViewModel) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *LocationAddressViewModel) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetCity

`func (o *LocationAddressViewModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationAddressViewModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationAddressViewModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LocationAddressViewModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *LocationAddressViewModel) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *LocationAddressViewModel) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *LocationAddressViewModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationAddressViewModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationAddressViewModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LocationAddressViewModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *LocationAddressViewModel) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LocationAddressViewModel) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetEmail

`func (o *LocationAddressViewModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationAddressViewModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationAddressViewModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationAddressViewModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocationAddressViewModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocationAddressViewModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *LocationAddressViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationAddressViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationAddressViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationAddressViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LocationAddressViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocationAddressViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhoneNumber

`func (o *LocationAddressViewModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *LocationAddressViewModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *LocationAddressViewModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *LocationAddressViewModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *LocationAddressViewModel) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *LocationAddressViewModel) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetState

`func (o *LocationAddressViewModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationAddressViewModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationAddressViewModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationAddressViewModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *LocationAddressViewModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *LocationAddressViewModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZipCode

`func (o *LocationAddressViewModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *LocationAddressViewModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *LocationAddressViewModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *LocationAddressViewModel) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *LocationAddressViewModel) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *LocationAddressViewModel) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


