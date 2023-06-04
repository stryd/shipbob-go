# ParentCarton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** | Barcode assigned to this carton | [optional] 
**Cartons** | Pointer to [**[]Carton**](Carton.md) | Cartons packed inside this parent container | [optional] 
**Measurements** | Pointer to [**CartonMeasurements**](CartonMeasurements.md) |  | [optional] 
**Type** | Pointer to **string** | Type of this carton container | [optional] 

## Methods

### NewParentCarton

`func NewParentCarton() *ParentCarton`

NewParentCarton instantiates a new ParentCarton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentCartonWithDefaults

`func NewParentCartonWithDefaults() *ParentCarton`

NewParentCartonWithDefaults instantiates a new ParentCarton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ParentCarton) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ParentCarton) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ParentCarton) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ParentCarton) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetCartons

`func (o *ParentCarton) GetCartons() []Carton`

GetCartons returns the Cartons field if non-nil, zero value otherwise.

### GetCartonsOk

`func (o *ParentCarton) GetCartonsOk() (*[]Carton, bool)`

GetCartonsOk returns a tuple with the Cartons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartons

`func (o *ParentCarton) SetCartons(v []Carton)`

SetCartons sets Cartons field to given value.

### HasCartons

`func (o *ParentCarton) HasCartons() bool`

HasCartons returns a boolean if a field has been set.

### GetMeasurements

`func (o *ParentCarton) GetMeasurements() CartonMeasurements`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *ParentCarton) GetMeasurementsOk() (*CartonMeasurements, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *ParentCarton) SetMeasurements(v CartonMeasurements)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *ParentCarton) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetType

`func (o *ParentCarton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParentCarton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParentCarton) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ParentCarton) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


