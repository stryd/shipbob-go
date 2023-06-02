# LocationService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**LocationAddress**](LocationAddress.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the user is authorized to access this service at the location | [optional] 
**ServiceType** | Pointer to [**LocationServiceTypeEnum**](LocationServiceTypeEnum.md) |  | [optional] 

## Methods

### NewLocationService

`func NewLocationService() *LocationService`

NewLocationService instantiates a new LocationService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationServiceWithDefaults

`func NewLocationServiceWithDefaults() *LocationService`

NewLocationServiceWithDefaults instantiates a new LocationService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LocationService) GetAddress() LocationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationService) GetAddressOk() (*LocationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationService) SetAddress(v LocationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LocationService) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *LocationService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LocationService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LocationService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LocationService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceType

`func (o *LocationService) GetServiceType() LocationServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *LocationService) GetServiceTypeOk() (*LocationServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *LocationService) SetServiceType(v LocationServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *LocationService) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


