# ShipBobOrdersPresentationViewModelsShipmentLogViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogTypeId** | Pointer to **int32** | Log type id of the shipment | [optional] 
**LogTypeName** | Pointer to **string** | Name of the log type | [optional] 
**LogTypeText** | Pointer to **string** | Summary of log type meaning | [optional] 
**Metadata** | Pointer to **map[string]string** | Specifics data for the event | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp of event | [optional] 

## Methods

### NewShipBobOrdersPresentationViewModelsShipmentLogViewModel

`func NewShipBobOrdersPresentationViewModelsShipmentLogViewModel() *ShipBobOrdersPresentationViewModelsShipmentLogViewModel`

NewShipBobOrdersPresentationViewModelsShipmentLogViewModel instantiates a new ShipBobOrdersPresentationViewModelsShipmentLogViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobOrdersPresentationViewModelsShipmentLogViewModelWithDefaults

`func NewShipBobOrdersPresentationViewModelsShipmentLogViewModelWithDefaults() *ShipBobOrdersPresentationViewModelsShipmentLogViewModel`

NewShipBobOrdersPresentationViewModelsShipmentLogViewModelWithDefaults instantiates a new ShipBobOrdersPresentationViewModelsShipmentLogViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogTypeId

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeId() int32`

GetLogTypeId returns the LogTypeId field if non-nil, zero value otherwise.

### GetLogTypeIdOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeIdOk() (*int32, bool)`

GetLogTypeIdOk returns a tuple with the LogTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeId

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) SetLogTypeId(v int32)`

SetLogTypeId sets LogTypeId field to given value.

### HasLogTypeId

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) HasLogTypeId() bool`

HasLogTypeId returns a boolean if a field has been set.

### GetLogTypeName

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeName() string`

GetLogTypeName returns the LogTypeName field if non-nil, zero value otherwise.

### GetLogTypeNameOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeNameOk() (*string, bool)`

GetLogTypeNameOk returns a tuple with the LogTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeName

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) SetLogTypeName(v string)`

SetLogTypeName sets LogTypeName field to given value.

### HasLogTypeName

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) HasLogTypeName() bool`

HasLogTypeName returns a boolean if a field has been set.

### GetLogTypeText

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeText() string`

GetLogTypeText returns the LogTypeText field if non-nil, zero value otherwise.

### GetLogTypeTextOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetLogTypeTextOk() (*string, bool)`

GetLogTypeTextOk returns a tuple with the LogTypeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeText

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) SetLogTypeText(v string)`

SetLogTypeText sets LogTypeText field to given value.

### HasLogTypeText

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) HasLogTypeText() bool`

HasLogTypeText returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShipBobOrdersPresentationViewModelsShipmentLogViewModel) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


