# LocationServiceViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**LocationAddressViewModel**](Location.AddressViewModel.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the user is authorized to access this service at the location | [optional] 
**ServiceType** | Pointer to [**LocationServiceTypeEnum**](Location.ServiceTypeEnum.md) |  | [optional] 

## Methods

### NewLocationServiceViewModel

`func NewLocationServiceViewModel() *LocationServiceViewModel`

NewLocationServiceViewModel instantiates a new LocationServiceViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationServiceViewModelWithDefaults

`func NewLocationServiceViewModelWithDefaults() *LocationServiceViewModel`

NewLocationServiceViewModelWithDefaults instantiates a new LocationServiceViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LocationServiceViewModel) GetAddress() LocationAddressViewModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationServiceViewModel) GetAddressOk() (*LocationAddressViewModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationServiceViewModel) SetAddress(v LocationAddressViewModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LocationServiceViewModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *LocationServiceViewModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LocationServiceViewModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LocationServiceViewModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LocationServiceViewModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceType

`func (o *LocationServiceViewModel) GetServiceType() LocationServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *LocationServiceViewModel) GetServiceTypeOk() (*LocationServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *LocationServiceViewModel) SetServiceType(v LocationServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *LocationServiceViewModel) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


