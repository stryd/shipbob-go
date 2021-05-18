# OrdersShipmentViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBobâs completion of the fulfillment operation as promised. \\r\\nCurrently, this means the shipment has been picked, packed, and label has been printed. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date this shipment was created | [optional] 
**EstimatedFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBobâs commitment for completing \\r\\nthe shipment and handing to the carrier for delivery. | [optional] 
**EstimatedFulfillmentDateStatus** | Pointer to **string** | Status of ShipBobâs completion of the fulfillment operation. | [optional] 
**Id** | Pointer to **int32** | Unique id of the shipment | [optional] 
**InsuranceValue** | Pointer to **NullableFloat64** | Monetary amount that this shipment was insured for | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Monetary amount that was invoiced for this shipment | [optional] 
**LastUpdateAt** | Pointer to **NullableTime** | Date this shipment was last updated | [optional] 
**Location** | Pointer to [**OrdersFulfillmentCenterViewModel**](Orders.FulfillmentCenterViewModel.md) |  | [optional] 
**Measurements** | Pointer to [**OrdersMeasurementsViewModel**](Orders.MeasurementsViewModel.md) |  | [optional] 
**OrderId** | Pointer to **int32** | Id of the order this shipment belongs to | [optional] 
**PackageMaterialType** | Pointer to **string** | Container type for the shipment | [optional] 
**Products** | Pointer to [**[]OrdersShipmentProductViewModel**](OrdersShipmentProductViewModel.md) | Information about the products contained in this shipment | [optional] 
**Recipient** | Pointer to [**OrdersRecipientViewModel**](Orders.RecipientViewModel.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order this shipment belongs to | [optional] 
**RequireSignature** | Pointer to **bool** | If a shipment requires signature | [optional] 
**ShipOption** | Pointer to **string** | Name of the shipping option used for this shipment | [optional] 
**Status** | Pointer to **string** | The shipment status | [optional] 
**StatusDetails** | Pointer to [**[]OrdersStatusDetailViewModel**](OrdersStatusDetailViewModel.md) | Additional details about the shipment status | [optional] 
**Tracking** | Pointer to [**OrdersTrackingViewModel**](Orders.TrackingViewModel.md) |  | [optional] 

## Methods

### NewOrdersShipmentViewModel

`func NewOrdersShipmentViewModel() *OrdersShipmentViewModel`

NewOrdersShipmentViewModel instantiates a new OrdersShipmentViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersShipmentViewModelWithDefaults

`func NewOrdersShipmentViewModelWithDefaults() *OrdersShipmentViewModel`

NewOrdersShipmentViewModelWithDefaults instantiates a new OrdersShipmentViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualFulfillmentDate

`func (o *OrdersShipmentViewModel) GetActualFulfillmentDate() time.Time`

GetActualFulfillmentDate returns the ActualFulfillmentDate field if non-nil, zero value otherwise.

### GetActualFulfillmentDateOk

`func (o *OrdersShipmentViewModel) GetActualFulfillmentDateOk() (*time.Time, bool)`

GetActualFulfillmentDateOk returns a tuple with the ActualFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFulfillmentDate

`func (o *OrdersShipmentViewModel) SetActualFulfillmentDate(v time.Time)`

SetActualFulfillmentDate sets ActualFulfillmentDate field to given value.

### HasActualFulfillmentDate

`func (o *OrdersShipmentViewModel) HasActualFulfillmentDate() bool`

HasActualFulfillmentDate returns a boolean if a field has been set.

### SetActualFulfillmentDateNil

`func (o *OrdersShipmentViewModel) SetActualFulfillmentDateNil(b bool)`

 SetActualFulfillmentDateNil sets the value for ActualFulfillmentDate to be an explicit nil

### UnsetActualFulfillmentDate
`func (o *OrdersShipmentViewModel) UnsetActualFulfillmentDate()`

UnsetActualFulfillmentDate ensures that no value is present for ActualFulfillmentDate, not even an explicit nil
### GetCreatedDate

`func (o *OrdersShipmentViewModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *OrdersShipmentViewModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *OrdersShipmentViewModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *OrdersShipmentViewModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEstimatedFulfillmentDate

`func (o *OrdersShipmentViewModel) GetEstimatedFulfillmentDate() time.Time`

GetEstimatedFulfillmentDate returns the EstimatedFulfillmentDate field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateOk

`func (o *OrdersShipmentViewModel) GetEstimatedFulfillmentDateOk() (*time.Time, bool)`

GetEstimatedFulfillmentDateOk returns a tuple with the EstimatedFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDate

`func (o *OrdersShipmentViewModel) SetEstimatedFulfillmentDate(v time.Time)`

SetEstimatedFulfillmentDate sets EstimatedFulfillmentDate field to given value.

### HasEstimatedFulfillmentDate

`func (o *OrdersShipmentViewModel) HasEstimatedFulfillmentDate() bool`

HasEstimatedFulfillmentDate returns a boolean if a field has been set.

### SetEstimatedFulfillmentDateNil

`func (o *OrdersShipmentViewModel) SetEstimatedFulfillmentDateNil(b bool)`

 SetEstimatedFulfillmentDateNil sets the value for EstimatedFulfillmentDate to be an explicit nil

### UnsetEstimatedFulfillmentDate
`func (o *OrdersShipmentViewModel) UnsetEstimatedFulfillmentDate()`

UnsetEstimatedFulfillmentDate ensures that no value is present for EstimatedFulfillmentDate, not even an explicit nil
### GetEstimatedFulfillmentDateStatus

`func (o *OrdersShipmentViewModel) GetEstimatedFulfillmentDateStatus() string`

GetEstimatedFulfillmentDateStatus returns the EstimatedFulfillmentDateStatus field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateStatusOk

`func (o *OrdersShipmentViewModel) GetEstimatedFulfillmentDateStatusOk() (*string, bool)`

GetEstimatedFulfillmentDateStatusOk returns a tuple with the EstimatedFulfillmentDateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDateStatus

`func (o *OrdersShipmentViewModel) SetEstimatedFulfillmentDateStatus(v string)`

SetEstimatedFulfillmentDateStatus sets EstimatedFulfillmentDateStatus field to given value.

### HasEstimatedFulfillmentDateStatus

`func (o *OrdersShipmentViewModel) HasEstimatedFulfillmentDateStatus() bool`

HasEstimatedFulfillmentDateStatus returns a boolean if a field has been set.

### GetId

`func (o *OrdersShipmentViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdersShipmentViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdersShipmentViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrdersShipmentViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsuranceValue

`func (o *OrdersShipmentViewModel) GetInsuranceValue() float64`

GetInsuranceValue returns the InsuranceValue field if non-nil, zero value otherwise.

### GetInsuranceValueOk

`func (o *OrdersShipmentViewModel) GetInsuranceValueOk() (*float64, bool)`

GetInsuranceValueOk returns a tuple with the InsuranceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceValue

`func (o *OrdersShipmentViewModel) SetInsuranceValue(v float64)`

SetInsuranceValue sets InsuranceValue field to given value.

### HasInsuranceValue

`func (o *OrdersShipmentViewModel) HasInsuranceValue() bool`

HasInsuranceValue returns a boolean if a field has been set.

### SetInsuranceValueNil

`func (o *OrdersShipmentViewModel) SetInsuranceValueNil(b bool)`

 SetInsuranceValueNil sets the value for InsuranceValue to be an explicit nil

### UnsetInsuranceValue
`func (o *OrdersShipmentViewModel) UnsetInsuranceValue()`

UnsetInsuranceValue ensures that no value is present for InsuranceValue, not even an explicit nil
### GetInvoiceAmount

`func (o *OrdersShipmentViewModel) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *OrdersShipmentViewModel) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *OrdersShipmentViewModel) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *OrdersShipmentViewModel) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *OrdersShipmentViewModel) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *OrdersShipmentViewModel) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetLastUpdateAt

`func (o *OrdersShipmentViewModel) GetLastUpdateAt() time.Time`

GetLastUpdateAt returns the LastUpdateAt field if non-nil, zero value otherwise.

### GetLastUpdateAtOk

`func (o *OrdersShipmentViewModel) GetLastUpdateAtOk() (*time.Time, bool)`

GetLastUpdateAtOk returns a tuple with the LastUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAt

`func (o *OrdersShipmentViewModel) SetLastUpdateAt(v time.Time)`

SetLastUpdateAt sets LastUpdateAt field to given value.

### HasLastUpdateAt

`func (o *OrdersShipmentViewModel) HasLastUpdateAt() bool`

HasLastUpdateAt returns a boolean if a field has been set.

### SetLastUpdateAtNil

`func (o *OrdersShipmentViewModel) SetLastUpdateAtNil(b bool)`

 SetLastUpdateAtNil sets the value for LastUpdateAt to be an explicit nil

### UnsetLastUpdateAt
`func (o *OrdersShipmentViewModel) UnsetLastUpdateAt()`

UnsetLastUpdateAt ensures that no value is present for LastUpdateAt, not even an explicit nil
### GetLocation

`func (o *OrdersShipmentViewModel) GetLocation() OrdersFulfillmentCenterViewModel`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrdersShipmentViewModel) GetLocationOk() (*OrdersFulfillmentCenterViewModel, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrdersShipmentViewModel) SetLocation(v OrdersFulfillmentCenterViewModel)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OrdersShipmentViewModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeasurements

`func (o *OrdersShipmentViewModel) GetMeasurements() OrdersMeasurementsViewModel`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *OrdersShipmentViewModel) GetMeasurementsOk() (*OrdersMeasurementsViewModel, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *OrdersShipmentViewModel) SetMeasurements(v OrdersMeasurementsViewModel)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *OrdersShipmentViewModel) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetOrderId

`func (o *OrdersShipmentViewModel) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrdersShipmentViewModel) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrdersShipmentViewModel) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrdersShipmentViewModel) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPackageMaterialType

`func (o *OrdersShipmentViewModel) GetPackageMaterialType() string`

GetPackageMaterialType returns the PackageMaterialType field if non-nil, zero value otherwise.

### GetPackageMaterialTypeOk

`func (o *OrdersShipmentViewModel) GetPackageMaterialTypeOk() (*string, bool)`

GetPackageMaterialTypeOk returns a tuple with the PackageMaterialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageMaterialType

`func (o *OrdersShipmentViewModel) SetPackageMaterialType(v string)`

SetPackageMaterialType sets PackageMaterialType field to given value.

### HasPackageMaterialType

`func (o *OrdersShipmentViewModel) HasPackageMaterialType() bool`

HasPackageMaterialType returns a boolean if a field has been set.

### GetProducts

`func (o *OrdersShipmentViewModel) GetProducts() []OrdersShipmentProductViewModel`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *OrdersShipmentViewModel) GetProductsOk() (*[]OrdersShipmentProductViewModel, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *OrdersShipmentViewModel) SetProducts(v []OrdersShipmentProductViewModel)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *OrdersShipmentViewModel) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRecipient

`func (o *OrdersShipmentViewModel) GetRecipient() OrdersRecipientViewModel`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *OrdersShipmentViewModel) GetRecipientOk() (*OrdersRecipientViewModel, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *OrdersShipmentViewModel) SetRecipient(v OrdersRecipientViewModel)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *OrdersShipmentViewModel) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *OrdersShipmentViewModel) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrdersShipmentViewModel) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrdersShipmentViewModel) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *OrdersShipmentViewModel) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetRequireSignature

`func (o *OrdersShipmentViewModel) GetRequireSignature() bool`

GetRequireSignature returns the RequireSignature field if non-nil, zero value otherwise.

### GetRequireSignatureOk

`func (o *OrdersShipmentViewModel) GetRequireSignatureOk() (*bool, bool)`

GetRequireSignatureOk returns a tuple with the RequireSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignature

`func (o *OrdersShipmentViewModel) SetRequireSignature(v bool)`

SetRequireSignature sets RequireSignature field to given value.

### HasRequireSignature

`func (o *OrdersShipmentViewModel) HasRequireSignature() bool`

HasRequireSignature returns a boolean if a field has been set.

### GetShipOption

`func (o *OrdersShipmentViewModel) GetShipOption() string`

GetShipOption returns the ShipOption field if non-nil, zero value otherwise.

### GetShipOptionOk

`func (o *OrdersShipmentViewModel) GetShipOptionOk() (*string, bool)`

GetShipOptionOk returns a tuple with the ShipOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipOption

`func (o *OrdersShipmentViewModel) SetShipOption(v string)`

SetShipOption sets ShipOption field to given value.

### HasShipOption

`func (o *OrdersShipmentViewModel) HasShipOption() bool`

HasShipOption returns a boolean if a field has been set.

### GetStatus

`func (o *OrdersShipmentViewModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrdersShipmentViewModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrdersShipmentViewModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrdersShipmentViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *OrdersShipmentViewModel) GetStatusDetails() []OrdersStatusDetailViewModel`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *OrdersShipmentViewModel) GetStatusDetailsOk() (*[]OrdersStatusDetailViewModel, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *OrdersShipmentViewModel) SetStatusDetails(v []OrdersStatusDetailViewModel)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *OrdersShipmentViewModel) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetTracking

`func (o *OrdersShipmentViewModel) GetTracking() OrdersTrackingViewModel`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *OrdersShipmentViewModel) GetTrackingOk() (*OrdersTrackingViewModel, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *OrdersShipmentViewModel) SetTracking(v OrdersTrackingViewModel)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *OrdersShipmentViewModel) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


