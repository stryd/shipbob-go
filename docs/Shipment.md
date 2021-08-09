# Shipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBob&#39;s completion of the fulfillment operation as promised. Currently, this means the shipment has been picked, packed, and label has been printed. | [optional] 
**CreatedDate** | Pointer to **NullableTime** | Date this shipment was created | [optional] 
**EstimatedFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBob&#39;s commitment for completing the shipment and handing to the carrier for delivery. | [optional] 
**EstimatedFulfillmentDateStatus** | Pointer to **string** | Status of ShipBob&#39;s completion of the fulfillment operation. | [optional] 
**Id** | Pointer to **int32** | Unique id of the shipment | [optional] 
**InsuranceValue** | Pointer to **NullableFloat64** | Monetary amount that this shipment was insured for | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Monetary amount that was invoiced for this shipment | [optional] 
**LastUpdateAt** | Pointer to **NullableTime** | Date this shipment was last updated | [optional] 
**Location** | Pointer to [**FulfillmentCenter**](FulfillmentCenter.md) |  | [optional] 
**Measurements** | Pointer to [**OrderMeasurements**](OrderMeasurements.md) |  | [optional] 
**OrderId** | Pointer to **int32** | Id of the order this shipment belongs to | [optional] 
**PackageMaterialType** | Pointer to **string** | Container type for the shipment | [optional] 
**Products** | Pointer to [**[]ShipmentProduct**](ShipmentProduct.md) | Information about the products contained in this shipment | [optional] 
**Recipient** | Pointer to [**Recipient**](Recipient.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order this shipment belongs to | [optional] 
**RequireSignature** | Pointer to **bool** | If a shipment requires signature | [optional] 
**ShipOption** | Pointer to **string** | Name of the shipping option used for this shipment | [optional] 
**Status** | Pointer to **string** | The shipment status | [optional] 
**StatusDetails** | Pointer to [**[]OrderStatusDetails**](OrderStatusDetails.md) | Additional details about the shipment status | [optional] 
**Tracking** | Pointer to [**Tracking**](Tracking.md) |  | [optional] 

## Methods

### NewShipment

`func NewShipment() *Shipment`

NewShipment instantiates a new Shipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentWithDefaults

`func NewShipmentWithDefaults() *Shipment`

NewShipmentWithDefaults instantiates a new Shipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualFulfillmentDate

`func (o *Shipment) GetActualFulfillmentDate() time.Time`

GetActualFulfillmentDate returns the ActualFulfillmentDate field if non-nil, zero value otherwise.

### GetActualFulfillmentDateOk

`func (o *Shipment) GetActualFulfillmentDateOk() (*time.Time, bool)`

GetActualFulfillmentDateOk returns a tuple with the ActualFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFulfillmentDate

`func (o *Shipment) SetActualFulfillmentDate(v time.Time)`

SetActualFulfillmentDate sets ActualFulfillmentDate field to given value.

### HasActualFulfillmentDate

`func (o *Shipment) HasActualFulfillmentDate() bool`

HasActualFulfillmentDate returns a boolean if a field has been set.

### SetActualFulfillmentDateNil

`func (o *Shipment) SetActualFulfillmentDateNil(b bool)`

 SetActualFulfillmentDateNil sets the value for ActualFulfillmentDate to be an explicit nil

### UnsetActualFulfillmentDate
`func (o *Shipment) UnsetActualFulfillmentDate()`

UnsetActualFulfillmentDate ensures that no value is present for ActualFulfillmentDate, not even an explicit nil
### GetCreatedDate

`func (o *Shipment) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Shipment) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Shipment) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Shipment) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *Shipment) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *Shipment) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetEstimatedFulfillmentDate

`func (o *Shipment) GetEstimatedFulfillmentDate() time.Time`

GetEstimatedFulfillmentDate returns the EstimatedFulfillmentDate field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateOk

`func (o *Shipment) GetEstimatedFulfillmentDateOk() (*time.Time, bool)`

GetEstimatedFulfillmentDateOk returns a tuple with the EstimatedFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDate

`func (o *Shipment) SetEstimatedFulfillmentDate(v time.Time)`

SetEstimatedFulfillmentDate sets EstimatedFulfillmentDate field to given value.

### HasEstimatedFulfillmentDate

`func (o *Shipment) HasEstimatedFulfillmentDate() bool`

HasEstimatedFulfillmentDate returns a boolean if a field has been set.

### SetEstimatedFulfillmentDateNil

`func (o *Shipment) SetEstimatedFulfillmentDateNil(b bool)`

 SetEstimatedFulfillmentDateNil sets the value for EstimatedFulfillmentDate to be an explicit nil

### UnsetEstimatedFulfillmentDate
`func (o *Shipment) UnsetEstimatedFulfillmentDate()`

UnsetEstimatedFulfillmentDate ensures that no value is present for EstimatedFulfillmentDate, not even an explicit nil
### GetEstimatedFulfillmentDateStatus

`func (o *Shipment) GetEstimatedFulfillmentDateStatus() string`

GetEstimatedFulfillmentDateStatus returns the EstimatedFulfillmentDateStatus field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateStatusOk

`func (o *Shipment) GetEstimatedFulfillmentDateStatusOk() (*string, bool)`

GetEstimatedFulfillmentDateStatusOk returns a tuple with the EstimatedFulfillmentDateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDateStatus

`func (o *Shipment) SetEstimatedFulfillmentDateStatus(v string)`

SetEstimatedFulfillmentDateStatus sets EstimatedFulfillmentDateStatus field to given value.

### HasEstimatedFulfillmentDateStatus

`func (o *Shipment) HasEstimatedFulfillmentDateStatus() bool`

HasEstimatedFulfillmentDateStatus returns a boolean if a field has been set.

### GetId

`func (o *Shipment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Shipment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Shipment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Shipment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsuranceValue

`func (o *Shipment) GetInsuranceValue() float64`

GetInsuranceValue returns the InsuranceValue field if non-nil, zero value otherwise.

### GetInsuranceValueOk

`func (o *Shipment) GetInsuranceValueOk() (*float64, bool)`

GetInsuranceValueOk returns a tuple with the InsuranceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceValue

`func (o *Shipment) SetInsuranceValue(v float64)`

SetInsuranceValue sets InsuranceValue field to given value.

### HasInsuranceValue

`func (o *Shipment) HasInsuranceValue() bool`

HasInsuranceValue returns a boolean if a field has been set.

### SetInsuranceValueNil

`func (o *Shipment) SetInsuranceValueNil(b bool)`

 SetInsuranceValueNil sets the value for InsuranceValue to be an explicit nil

### UnsetInsuranceValue
`func (o *Shipment) UnsetInsuranceValue()`

UnsetInsuranceValue ensures that no value is present for InsuranceValue, not even an explicit nil
### GetInvoiceAmount

`func (o *Shipment) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *Shipment) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *Shipment) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *Shipment) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *Shipment) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *Shipment) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetLastUpdateAt

`func (o *Shipment) GetLastUpdateAt() time.Time`

GetLastUpdateAt returns the LastUpdateAt field if non-nil, zero value otherwise.

### GetLastUpdateAtOk

`func (o *Shipment) GetLastUpdateAtOk() (*time.Time, bool)`

GetLastUpdateAtOk returns a tuple with the LastUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAt

`func (o *Shipment) SetLastUpdateAt(v time.Time)`

SetLastUpdateAt sets LastUpdateAt field to given value.

### HasLastUpdateAt

`func (o *Shipment) HasLastUpdateAt() bool`

HasLastUpdateAt returns a boolean if a field has been set.

### SetLastUpdateAtNil

`func (o *Shipment) SetLastUpdateAtNil(b bool)`

 SetLastUpdateAtNil sets the value for LastUpdateAt to be an explicit nil

### UnsetLastUpdateAt
`func (o *Shipment) UnsetLastUpdateAt()`

UnsetLastUpdateAt ensures that no value is present for LastUpdateAt, not even an explicit nil
### GetLocation

`func (o *Shipment) GetLocation() FulfillmentCenter`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Shipment) GetLocationOk() (*FulfillmentCenter, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Shipment) SetLocation(v FulfillmentCenter)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Shipment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeasurements

`func (o *Shipment) GetMeasurements() OrderMeasurements`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *Shipment) GetMeasurementsOk() (*OrderMeasurements, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *Shipment) SetMeasurements(v OrderMeasurements)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *Shipment) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetOrderId

`func (o *Shipment) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Shipment) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Shipment) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Shipment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPackageMaterialType

`func (o *Shipment) GetPackageMaterialType() string`

GetPackageMaterialType returns the PackageMaterialType field if non-nil, zero value otherwise.

### GetPackageMaterialTypeOk

`func (o *Shipment) GetPackageMaterialTypeOk() (*string, bool)`

GetPackageMaterialTypeOk returns a tuple with the PackageMaterialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageMaterialType

`func (o *Shipment) SetPackageMaterialType(v string)`

SetPackageMaterialType sets PackageMaterialType field to given value.

### HasPackageMaterialType

`func (o *Shipment) HasPackageMaterialType() bool`

HasPackageMaterialType returns a boolean if a field has been set.

### GetProducts

`func (o *Shipment) GetProducts() []ShipmentProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Shipment) GetProductsOk() (*[]ShipmentProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Shipment) SetProducts(v []ShipmentProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Shipment) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRecipient

`func (o *Shipment) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Shipment) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Shipment) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Shipment) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *Shipment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Shipment) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Shipment) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Shipment) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetRequireSignature

`func (o *Shipment) GetRequireSignature() bool`

GetRequireSignature returns the RequireSignature field if non-nil, zero value otherwise.

### GetRequireSignatureOk

`func (o *Shipment) GetRequireSignatureOk() (*bool, bool)`

GetRequireSignatureOk returns a tuple with the RequireSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignature

`func (o *Shipment) SetRequireSignature(v bool)`

SetRequireSignature sets RequireSignature field to given value.

### HasRequireSignature

`func (o *Shipment) HasRequireSignature() bool`

HasRequireSignature returns a boolean if a field has been set.

### GetShipOption

`func (o *Shipment) GetShipOption() string`

GetShipOption returns the ShipOption field if non-nil, zero value otherwise.

### GetShipOptionOk

`func (o *Shipment) GetShipOptionOk() (*string, bool)`

GetShipOptionOk returns a tuple with the ShipOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipOption

`func (o *Shipment) SetShipOption(v string)`

SetShipOption sets ShipOption field to given value.

### HasShipOption

`func (o *Shipment) HasShipOption() bool`

HasShipOption returns a boolean if a field has been set.

### GetStatus

`func (o *Shipment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Shipment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Shipment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Shipment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *Shipment) GetStatusDetails() []OrderStatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Shipment) GetStatusDetailsOk() (*[]OrderStatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Shipment) SetStatusDetails(v []OrderStatusDetails)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Shipment) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetTracking

`func (o *Shipment) GetTracking() Tracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *Shipment) GetTrackingOk() (*Tracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *Shipment) SetTracking(v Tracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *Shipment) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


