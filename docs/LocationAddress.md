# LocationAddress

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

### NewLocationAddress

`func NewLocationAddress() *LocationAddress`

NewLocationAddress instantiates a new LocationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAddressWithDefaults

`func NewLocationAddressWithDefaults() *LocationAddress`

NewLocationAddressWithDefaults instantiates a new LocationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *LocationAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *LocationAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *LocationAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *LocationAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *LocationAddress) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *LocationAddress) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *LocationAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *LocationAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *LocationAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *LocationAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *LocationAddress) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *LocationAddress) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetCity

`func (o *LocationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LocationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LocationAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LocationAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *LocationAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *LocationAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *LocationAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LocationAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LocationAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LocationAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *LocationAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LocationAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetEmail

`func (o *LocationAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LocationAddress) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LocationAddress) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *LocationAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LocationAddress) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LocationAddress) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhoneNumber

`func (o *LocationAddress) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *LocationAddress) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *LocationAddress) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *LocationAddress) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *LocationAddress) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *LocationAddress) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetState

`func (o *LocationAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LocationAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LocationAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LocationAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *LocationAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *LocationAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetZipCode

`func (o *LocationAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *LocationAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *LocationAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *LocationAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *LocationAddress) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *LocationAddress) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


