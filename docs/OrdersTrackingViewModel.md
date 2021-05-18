# OrdersTrackingViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | Carrier of the shipment | [optional] 
**CarrierService** | Pointer to **string** | The carrier&#39;s service which was used for this shipment | [optional] 
**TrackingNumber** | Pointer to **string** | Tracking number of the shipment | [optional] 
**TrackingUrl** | Pointer to **string** | URL to the website where a shipment can be tracked | [optional] 

## Methods

### NewOrdersTrackingViewModel

`func NewOrdersTrackingViewModel() *OrdersTrackingViewModel`

NewOrdersTrackingViewModel instantiates a new OrdersTrackingViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersTrackingViewModelWithDefaults

`func NewOrdersTrackingViewModelWithDefaults() *OrdersTrackingViewModel`

NewOrdersTrackingViewModelWithDefaults instantiates a new OrdersTrackingViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *OrdersTrackingViewModel) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *OrdersTrackingViewModel) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *OrdersTrackingViewModel) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *OrdersTrackingViewModel) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierService

`func (o *OrdersTrackingViewModel) GetCarrierService() string`

GetCarrierService returns the CarrierService field if non-nil, zero value otherwise.

### GetCarrierServiceOk

`func (o *OrdersTrackingViewModel) GetCarrierServiceOk() (*string, bool)`

GetCarrierServiceOk returns a tuple with the CarrierService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierService

`func (o *OrdersTrackingViewModel) SetCarrierService(v string)`

SetCarrierService sets CarrierService field to given value.

### HasCarrierService

`func (o *OrdersTrackingViewModel) HasCarrierService() bool`

HasCarrierService returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *OrdersTrackingViewModel) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrdersTrackingViewModel) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrdersTrackingViewModel) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *OrdersTrackingViewModel) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *OrdersTrackingViewModel) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *OrdersTrackingViewModel) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *OrdersTrackingViewModel) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *OrdersTrackingViewModel) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


