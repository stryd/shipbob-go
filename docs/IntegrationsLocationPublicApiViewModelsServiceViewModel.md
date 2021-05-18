# IntegrationsLocationPublicApiViewModelsServiceViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**IntegrationsLocationPublicApiViewModelsAddressViewModel**](Integrations.Location.Public.Api.ViewModels.AddressViewModel.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the user is authorized to access this service at the location | [optional] 
**ServiceType** | Pointer to [**IntegrationsLocationPublicCommonServiceTypeEnum**](Integrations.Location.Public.Common.ServiceTypeEnum.md) |  | [optional] 

## Methods

### NewIntegrationsLocationPublicApiViewModelsServiceViewModel

`func NewIntegrationsLocationPublicApiViewModelsServiceViewModel() *IntegrationsLocationPublicApiViewModelsServiceViewModel`

NewIntegrationsLocationPublicApiViewModelsServiceViewModel instantiates a new IntegrationsLocationPublicApiViewModelsServiceViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsLocationPublicApiViewModelsServiceViewModelWithDefaults

`func NewIntegrationsLocationPublicApiViewModelsServiceViewModelWithDefaults() *IntegrationsLocationPublicApiViewModelsServiceViewModel`

NewIntegrationsLocationPublicApiViewModelsServiceViewModelWithDefaults instantiates a new IntegrationsLocationPublicApiViewModelsServiceViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetAddress() IntegrationsLocationPublicApiViewModelsAddressViewModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetAddressOk() (*IntegrationsLocationPublicApiViewModelsAddressViewModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) SetAddress(v IntegrationsLocationPublicApiViewModelsAddressViewModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceType

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetServiceType() IntegrationsLocationPublicCommonServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) GetServiceTypeOk() (*IntegrationsLocationPublicCommonServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) SetServiceType(v IntegrationsLocationPublicCommonServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *IntegrationsLocationPublicApiViewModelsServiceViewModel) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


