# ReceivingFulfillmentCenterViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **NullableString** | Address line one of the fulfillment center | [optional] 
**Address2** | Pointer to **NullableString** | Address line two of the fulfillment center | [optional] 
**City** | Pointer to **NullableString** | City the fulfillment center is located in | [optional] 
**Country** | Pointer to **NullableString** | Country the fulfillment center is located in | [optional] 
**Email** | Pointer to **NullableString** | Email contact for the fulfillment center | [optional] 
**Id** | Pointer to **int32** | Unique identifier of the fulfillment center | [optional] 
**Name** | Pointer to **NullableString** | Name of the fulfillment center | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Phone number contact for the fulfillment center | [optional] 
**State** | Pointer to **NullableString** | State the fulfillment center is located in | [optional] 
**Timezone** | Pointer to **NullableString** | Timezone the fulfillment center is located in | [optional] 
**ZipCode** | Pointer to **NullableString** | Postal code of the fulfillment center | [optional] 

## Methods

### NewReceivingFulfillmentCenterViewModel

`func NewReceivingFulfillmentCenterViewModel() *ReceivingFulfillmentCenterViewModel`

NewReceivingFulfillmentCenterViewModel instantiates a new ReceivingFulfillmentCenterViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceivingFulfillmentCenterViewModelWithDefaults

`func NewReceivingFulfillmentCenterViewModelWithDefaults() *ReceivingFulfillmentCenterViewModel`

NewReceivingFulfillmentCenterViewModelWithDefaults instantiates a new ReceivingFulfillmentCenterViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *ReceivingFulfillmentCenterViewModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ReceivingFulfillmentCenterViewModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ReceivingFulfillmentCenterViewModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ReceivingFulfillmentCenterViewModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *ReceivingFulfillmentCenterViewModel) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *ReceivingFulfillmentCenterViewModel) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *ReceivingFulfillmentCenterViewModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ReceivingFulfillmentCenterViewModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ReceivingFulfillmentCenterViewModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ReceivingFulfillmentCenterViewModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *ReceivingFulfillmentCenterViewModel) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *ReceivingFulfillmentCenterViewModel) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetCity

`func (o *ReceivingFulfillmentCenterViewModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ReceivingFulfillmentCenterViewModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ReceivingFulfillmentCenterViewModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ReceivingFulfillmentCenterViewModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ReceivingFulfillmentCenterViewModel) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ReceivingFulfillmentCenterViewModel) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ReceivingFulfillmentCenterViewModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ReceivingFulfillmentCenterViewModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ReceivingFulfillmentCenterViewModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ReceivingFulfillmentCenterViewModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ReceivingFulfillmentCenterViewModel) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ReceivingFulfillmentCenterViewModel) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetEmail

`func (o *ReceivingFulfillmentCenterViewModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReceivingFulfillmentCenterViewModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReceivingFulfillmentCenterViewModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ReceivingFulfillmentCenterViewModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ReceivingFulfillmentCenterViewModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ReceivingFulfillmentCenterViewModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *ReceivingFulfillmentCenterViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceivingFulfillmentCenterViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceivingFulfillmentCenterViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReceivingFulfillmentCenterViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReceivingFulfillmentCenterViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReceivingFulfillmentCenterViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReceivingFulfillmentCenterViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReceivingFulfillmentCenterViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReceivingFulfillmentCenterViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReceivingFulfillmentCenterViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhoneNumber

`func (o *ReceivingFulfillmentCenterViewModel) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ReceivingFulfillmentCenterViewModel) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ReceivingFulfillmentCenterViewModel) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ReceivingFulfillmentCenterViewModel) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ReceivingFulfillmentCenterViewModel) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ReceivingFulfillmentCenterViewModel) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetState

`func (o *ReceivingFulfillmentCenterViewModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReceivingFulfillmentCenterViewModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReceivingFulfillmentCenterViewModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReceivingFulfillmentCenterViewModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ReceivingFulfillmentCenterViewModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ReceivingFulfillmentCenterViewModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTimezone

`func (o *ReceivingFulfillmentCenterViewModel) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReceivingFulfillmentCenterViewModel) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReceivingFulfillmentCenterViewModel) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ReceivingFulfillmentCenterViewModel) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ReceivingFulfillmentCenterViewModel) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ReceivingFulfillmentCenterViewModel) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetZipCode

`func (o *ReceivingFulfillmentCenterViewModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ReceivingFulfillmentCenterViewModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ReceivingFulfillmentCenterViewModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ReceivingFulfillmentCenterViewModel) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ReceivingFulfillmentCenterViewModel) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ReceivingFulfillmentCenterViewModel) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


