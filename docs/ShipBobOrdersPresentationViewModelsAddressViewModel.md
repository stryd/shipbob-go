# ShipBobOrdersPresentationViewModelsAddressViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** | First line of the address | 
**Address2** | Pointer to **string** | Second line of the address | [optional] 
**City** | **string** | The city | 
**CompanyName** | Pointer to **string** | Name of the company receiving the shipment | [optional] 
**Country** | **string** | The country (Must be ISO Alpha-2 for estimates) | 
**State** | Pointer to **string** | The state or province | [optional] 
**ZipCode** | Pointer to **string** | The zip code or postal code | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsAddressViewModel

`func NewShipBobOrdersPresentationViewModelsAddressViewModel(address1 string, city string, country string, ) *ShipBobOrdersPresentationViewModelsAddressViewModel`

NewShipBobOrdersPresentationViewModelsAddressViewModel instantiates a new ShipBobOrdersPresentationViewModelsAddressViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsAddressViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsAddressViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsAddressViewModel`

NewShipBobOrdersPresentationViewModelsAddressViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsAddressViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetCity(v string)`

SetCity sets City field to given value.


### GetCompanyName

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ShipBobOrdersPresentationViewModelsAddressViewModel) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


