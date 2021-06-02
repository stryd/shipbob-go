# EstimationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | First line of the address | [optional] 
**Address2** | Pointer to **string** | Second line of the address | [optional] 
**City** | Pointer to **string** | The city | [optional] 
**CompanyName** | Pointer to **string** | Name of the company receiving the shipment | [optional] 
**Country** | **string** | The country (Must be ISO Alpha-2 for estimates) | 
**State** | Pointer to **string** | The state or province | [optional] 
**ZipCode** | Pointer to **string** | The zip code or postal code | [optional] 

## Methods

### NewEstimationAddress

`func NewEstimationAddress(country string, ) *EstimationAddress`

NewEstimationAddress instantiates a new EstimationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimationAddressWithDefaults

`func NewEstimationAddressWithDefaults() *EstimationAddress`

NewEstimationAddressWithDefaults instantiates a new EstimationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *EstimationAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *EstimationAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *EstimationAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *EstimationAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *EstimationAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *EstimationAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *EstimationAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *EstimationAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *EstimationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EstimationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EstimationAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *EstimationAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *EstimationAddress) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EstimationAddress) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EstimationAddress) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *EstimationAddress) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *EstimationAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EstimationAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EstimationAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *EstimationAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EstimationAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EstimationAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EstimationAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *EstimationAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *EstimationAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *EstimationAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *EstimationAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


