# Carton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** | Barcode assigned to this carton | [optional] 
**CartonDetails** | Pointer to [**[]CartonDetails**](CartonDetails.md) | Details about the contents of this carton | [optional] 
**Id** | Pointer to **int32** | ID assigned to this carton | [optional] 
**Measurements** | Pointer to [**CartonMeasurements**](CartonMeasurements.md) |  | [optional] 
**Type** | Pointer to **string** | Type of this carton container | [optional] 

## Methods

### NewCarton

`func NewCarton() *Carton`

NewCarton instantiates a new Carton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartonWithDefaults

`func NewCartonWithDefaults() *Carton`

NewCartonWithDefaults instantiates a new Carton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *Carton) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Carton) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Carton) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Carton) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetCartonDetails

`func (o *Carton) GetCartonDetails() []CartonDetails`

GetCartonDetails returns the CartonDetails field if non-nil, zero value otherwise.

### GetCartonDetailsOk

`func (o *Carton) GetCartonDetailsOk() (*[]CartonDetails, bool)`

GetCartonDetailsOk returns a tuple with the CartonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartonDetails

`func (o *Carton) SetCartonDetails(v []CartonDetails)`

SetCartonDetails sets CartonDetails field to given value.

### HasCartonDetails

`func (o *Carton) HasCartonDetails() bool`

HasCartonDetails returns a boolean if a field has been set.

### GetId

`func (o *Carton) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Carton) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Carton) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Carton) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeasurements

`func (o *Carton) GetMeasurements() CartonMeasurements`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *Carton) GetMeasurementsOk() (*CartonMeasurements, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *Carton) SetMeasurements(v CartonMeasurements)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *Carton) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetType

`func (o *Carton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Carton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Carton) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Carton) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


