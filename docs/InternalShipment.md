# InternalShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBob&#39;s completion of the fulfillment operation as promised. Currently, this means the shipment has been picked, packed, and label has been printed. | [optional] 
**CreatedDate** | Pointer to **NullableTime** | Date this shipment was created | [optional] 
**DeliveryDate** | Pointer to **NullableTime** | The datetime of Shipment delivered to customer. | [optional] 
**EstimatedFulfillmentDate** | Pointer to **NullableTime** | The datetime of ShipBob&#39;s commitment for completing the shipment and handing to the carrier for delivery. | [optional] 
**EstimatedFulfillmentDateStatus** | Pointer to **string** | Status of ShipBob&#39;s completion of the fulfillment operation. | [optional] 
**GiftMessage** | Pointer to **string** | Gift message associated with the shipment | [optional] 
**Id** | Pointer to **int32** | Unique id of the shipment | [optional] 
**InsuranceValue** | Pointer to **NullableFloat64** | Monetary amount that this shipment was insured for | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** | Monetary amount that was invoiced for this shipment | [optional] 
**InvoiceCurrencyCode** | Pointer to **string** |  | [optional] 
**IsTrackingUploaded** | Pointer to **bool** | Indicates whether the Shipment was marked with tracking information uploaded to a third-party system where the order originated. | [optional] 
**LastTrackingUpdateAt** | Pointer to **NullableTime** | Timestamp for the last time this shipment had a tracking update | [optional] 
**LastUpdateAt** | Pointer to **NullableTime** | Date this shipment was last updated | [optional] 
**Location** | Pointer to [**FulfillmentCenter**](FulfillmentCenter.md) |  | [optional] 
**Measurements** | Pointer to [**OrderMeasurements**](OrderMeasurements.md) |  | [optional] 
**OrderId** | Pointer to **int32** | Id of the order this shipment belongs to | [optional] 
**PackageMaterialType** | Pointer to **string** | Container type for the shipment | [optional] 
**ParentCartons** | Pointer to [**[]ParentCarton**](ParentCarton.md) | Carton information for this shipment | [optional] 
**Products** | Pointer to [**[]ShipmentProduct**](ShipmentProduct.md) | Information about the products contained in this shipment | [optional] 
**Recipient** | Pointer to [**Recipient**](Recipient.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | Client-defined external unique id of the order this shipment belongs to | [optional] 
**RequireSignature** | Pointer to **bool** | If a shipment requires signature | [optional] 
**ShipOption** | Pointer to **string** | Name of the shipping option used for this shipment | [optional] 
**Status** | Pointer to **string** | The shipment status | [optional] 
**StatusDetails** | Pointer to [**[]OrderStatusDetail**](OrderStatusDetail.md) | Additional details about the shipment status | [optional] 
**StoreOrderId** | Pointer to **string** | Unique store order id of the shipment | [optional] 
**Tracking** | Pointer to [**Tracking**](Tracking.md) |  | [optional] 

## Methods

### NewInternalShipment

`func NewInternalShipment() *InternalShipment`

NewInternalShipment instantiates a new InternalShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalShipmentWithDefaults

`func NewInternalShipmentWithDefaults() *InternalShipment`

NewInternalShipmentWithDefaults instantiates a new InternalShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualFulfillmentDate

`func (o *InternalShipment) GetActualFulfillmentDate() time.Time`

GetActualFulfillmentDate returns the ActualFulfillmentDate field if non-nil, zero value otherwise.

### GetActualFulfillmentDateOk

`func (o *InternalShipment) GetActualFulfillmentDateOk() (*time.Time, bool)`

GetActualFulfillmentDateOk returns a tuple with the ActualFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFulfillmentDate

`func (o *InternalShipment) SetActualFulfillmentDate(v time.Time)`

SetActualFulfillmentDate sets ActualFulfillmentDate field to given value.

### HasActualFulfillmentDate

`func (o *InternalShipment) HasActualFulfillmentDate() bool`

HasActualFulfillmentDate returns a boolean if a field has been set.

### SetActualFulfillmentDateNil

`func (o *InternalShipment) SetActualFulfillmentDateNil(b bool)`

 SetActualFulfillmentDateNil sets the value for ActualFulfillmentDate to be an explicit nil

### UnsetActualFulfillmentDate
`func (o *InternalShipment) UnsetActualFulfillmentDate()`

UnsetActualFulfillmentDate ensures that no value is present for ActualFulfillmentDate, not even an explicit nil
### GetCreatedDate

`func (o *InternalShipment) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *InternalShipment) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *InternalShipment) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *InternalShipment) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *InternalShipment) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *InternalShipment) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetDeliveryDate

`func (o *InternalShipment) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *InternalShipment) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *InternalShipment) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *InternalShipment) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *InternalShipment) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *InternalShipment) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetEstimatedFulfillmentDate

`func (o *InternalShipment) GetEstimatedFulfillmentDate() time.Time`

GetEstimatedFulfillmentDate returns the EstimatedFulfillmentDate field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateOk

`func (o *InternalShipment) GetEstimatedFulfillmentDateOk() (*time.Time, bool)`

GetEstimatedFulfillmentDateOk returns a tuple with the EstimatedFulfillmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDate

`func (o *InternalShipment) SetEstimatedFulfillmentDate(v time.Time)`

SetEstimatedFulfillmentDate sets EstimatedFulfillmentDate field to given value.

### HasEstimatedFulfillmentDate

`func (o *InternalShipment) HasEstimatedFulfillmentDate() bool`

HasEstimatedFulfillmentDate returns a boolean if a field has been set.

### SetEstimatedFulfillmentDateNil

`func (o *InternalShipment) SetEstimatedFulfillmentDateNil(b bool)`

 SetEstimatedFulfillmentDateNil sets the value for EstimatedFulfillmentDate to be an explicit nil

### UnsetEstimatedFulfillmentDate
`func (o *InternalShipment) UnsetEstimatedFulfillmentDate()`

UnsetEstimatedFulfillmentDate ensures that no value is present for EstimatedFulfillmentDate, not even an explicit nil
### GetEstimatedFulfillmentDateStatus

`func (o *InternalShipment) GetEstimatedFulfillmentDateStatus() string`

GetEstimatedFulfillmentDateStatus returns the EstimatedFulfillmentDateStatus field if non-nil, zero value otherwise.

### GetEstimatedFulfillmentDateStatusOk

`func (o *InternalShipment) GetEstimatedFulfillmentDateStatusOk() (*string, bool)`

GetEstimatedFulfillmentDateStatusOk returns a tuple with the EstimatedFulfillmentDateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFulfillmentDateStatus

`func (o *InternalShipment) SetEstimatedFulfillmentDateStatus(v string)`

SetEstimatedFulfillmentDateStatus sets EstimatedFulfillmentDateStatus field to given value.

### HasEstimatedFulfillmentDateStatus

`func (o *InternalShipment) HasEstimatedFulfillmentDateStatus() bool`

HasEstimatedFulfillmentDateStatus returns a boolean if a field has been set.

### GetGiftMessage

`func (o *InternalShipment) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *InternalShipment) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *InternalShipment) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *InternalShipment) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### GetId

`func (o *InternalShipment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalShipment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalShipment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalShipment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsuranceValue

`func (o *InternalShipment) GetInsuranceValue() float64`

GetInsuranceValue returns the InsuranceValue field if non-nil, zero value otherwise.

### GetInsuranceValueOk

`func (o *InternalShipment) GetInsuranceValueOk() (*float64, bool)`

GetInsuranceValueOk returns a tuple with the InsuranceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceValue

`func (o *InternalShipment) SetInsuranceValue(v float64)`

SetInsuranceValue sets InsuranceValue field to given value.

### HasInsuranceValue

`func (o *InternalShipment) HasInsuranceValue() bool`

HasInsuranceValue returns a boolean if a field has been set.

### SetInsuranceValueNil

`func (o *InternalShipment) SetInsuranceValueNil(b bool)`

 SetInsuranceValueNil sets the value for InsuranceValue to be an explicit nil

### UnsetInsuranceValue
`func (o *InternalShipment) UnsetInsuranceValue()`

UnsetInsuranceValue ensures that no value is present for InsuranceValue, not even an explicit nil
### GetInvoiceAmount

`func (o *InternalShipment) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *InternalShipment) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *InternalShipment) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *InternalShipment) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *InternalShipment) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *InternalShipment) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetInvoiceCurrencyCode

`func (o *InternalShipment) GetInvoiceCurrencyCode() string`

GetInvoiceCurrencyCode returns the InvoiceCurrencyCode field if non-nil, zero value otherwise.

### GetInvoiceCurrencyCodeOk

`func (o *InternalShipment) GetInvoiceCurrencyCodeOk() (*string, bool)`

GetInvoiceCurrencyCodeOk returns a tuple with the InvoiceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCurrencyCode

`func (o *InternalShipment) SetInvoiceCurrencyCode(v string)`

SetInvoiceCurrencyCode sets InvoiceCurrencyCode field to given value.

### HasInvoiceCurrencyCode

`func (o *InternalShipment) HasInvoiceCurrencyCode() bool`

HasInvoiceCurrencyCode returns a boolean if a field has been set.

### GetIsTrackingUploaded

`func (o *InternalShipment) GetIsTrackingUploaded() bool`

GetIsTrackingUploaded returns the IsTrackingUploaded field if non-nil, zero value otherwise.

### GetIsTrackingUploadedOk

`func (o *InternalShipment) GetIsTrackingUploadedOk() (*bool, bool)`

GetIsTrackingUploadedOk returns a tuple with the IsTrackingUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingUploaded

`func (o *InternalShipment) SetIsTrackingUploaded(v bool)`

SetIsTrackingUploaded sets IsTrackingUploaded field to given value.

### HasIsTrackingUploaded

`func (o *InternalShipment) HasIsTrackingUploaded() bool`

HasIsTrackingUploaded returns a boolean if a field has been set.

### GetLastTrackingUpdateAt

`func (o *InternalShipment) GetLastTrackingUpdateAt() time.Time`

GetLastTrackingUpdateAt returns the LastTrackingUpdateAt field if non-nil, zero value otherwise.

### GetLastTrackingUpdateAtOk

`func (o *InternalShipment) GetLastTrackingUpdateAtOk() (*time.Time, bool)`

GetLastTrackingUpdateAtOk returns a tuple with the LastTrackingUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrackingUpdateAt

`func (o *InternalShipment) SetLastTrackingUpdateAt(v time.Time)`

SetLastTrackingUpdateAt sets LastTrackingUpdateAt field to given value.

### HasLastTrackingUpdateAt

`func (o *InternalShipment) HasLastTrackingUpdateAt() bool`

HasLastTrackingUpdateAt returns a boolean if a field has been set.

### SetLastTrackingUpdateAtNil

`func (o *InternalShipment) SetLastTrackingUpdateAtNil(b bool)`

 SetLastTrackingUpdateAtNil sets the value for LastTrackingUpdateAt to be an explicit nil

### UnsetLastTrackingUpdateAt
`func (o *InternalShipment) UnsetLastTrackingUpdateAt()`

UnsetLastTrackingUpdateAt ensures that no value is present for LastTrackingUpdateAt, not even an explicit nil
### GetLastUpdateAt

`func (o *InternalShipment) GetLastUpdateAt() time.Time`

GetLastUpdateAt returns the LastUpdateAt field if non-nil, zero value otherwise.

### GetLastUpdateAtOk

`func (o *InternalShipment) GetLastUpdateAtOk() (*time.Time, bool)`

GetLastUpdateAtOk returns a tuple with the LastUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAt

`func (o *InternalShipment) SetLastUpdateAt(v time.Time)`

SetLastUpdateAt sets LastUpdateAt field to given value.

### HasLastUpdateAt

`func (o *InternalShipment) HasLastUpdateAt() bool`

HasLastUpdateAt returns a boolean if a field has been set.

### SetLastUpdateAtNil

`func (o *InternalShipment) SetLastUpdateAtNil(b bool)`

 SetLastUpdateAtNil sets the value for LastUpdateAt to be an explicit nil

### UnsetLastUpdateAt
`func (o *InternalShipment) UnsetLastUpdateAt()`

UnsetLastUpdateAt ensures that no value is present for LastUpdateAt, not even an explicit nil
### GetLocation

`func (o *InternalShipment) GetLocation() FulfillmentCenter`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InternalShipment) GetLocationOk() (*FulfillmentCenter, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InternalShipment) SetLocation(v FulfillmentCenter)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InternalShipment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeasurements

`func (o *InternalShipment) GetMeasurements() OrderMeasurements`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *InternalShipment) GetMeasurementsOk() (*OrderMeasurements, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *InternalShipment) SetMeasurements(v OrderMeasurements)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *InternalShipment) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.

### GetOrderId

`func (o *InternalShipment) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InternalShipment) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InternalShipment) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *InternalShipment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPackageMaterialType

`func (o *InternalShipment) GetPackageMaterialType() string`

GetPackageMaterialType returns the PackageMaterialType field if non-nil, zero value otherwise.

### GetPackageMaterialTypeOk

`func (o *InternalShipment) GetPackageMaterialTypeOk() (*string, bool)`

GetPackageMaterialTypeOk returns a tuple with the PackageMaterialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageMaterialType

`func (o *InternalShipment) SetPackageMaterialType(v string)`

SetPackageMaterialType sets PackageMaterialType field to given value.

### HasPackageMaterialType

`func (o *InternalShipment) HasPackageMaterialType() bool`

HasPackageMaterialType returns a boolean if a field has been set.

### GetParentCartons

`func (o *InternalShipment) GetParentCartons() []ParentCarton`

GetParentCartons returns the ParentCartons field if non-nil, zero value otherwise.

### GetParentCartonsOk

`func (o *InternalShipment) GetParentCartonsOk() (*[]ParentCarton, bool)`

GetParentCartonsOk returns a tuple with the ParentCartons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCartons

`func (o *InternalShipment) SetParentCartons(v []ParentCarton)`

SetParentCartons sets ParentCartons field to given value.

### HasParentCartons

`func (o *InternalShipment) HasParentCartons() bool`

HasParentCartons returns a boolean if a field has been set.

### GetProducts

`func (o *InternalShipment) GetProducts() []ShipmentProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InternalShipment) GetProductsOk() (*[]ShipmentProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InternalShipment) SetProducts(v []ShipmentProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InternalShipment) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRecipient

`func (o *InternalShipment) GetRecipient() Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *InternalShipment) GetRecipientOk() (*Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *InternalShipment) SetRecipient(v Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *InternalShipment) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetReferenceId

`func (o *InternalShipment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InternalShipment) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InternalShipment) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InternalShipment) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetRequireSignature

`func (o *InternalShipment) GetRequireSignature() bool`

GetRequireSignature returns the RequireSignature field if non-nil, zero value otherwise.

### GetRequireSignatureOk

`func (o *InternalShipment) GetRequireSignatureOk() (*bool, bool)`

GetRequireSignatureOk returns a tuple with the RequireSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignature

`func (o *InternalShipment) SetRequireSignature(v bool)`

SetRequireSignature sets RequireSignature field to given value.

### HasRequireSignature

`func (o *InternalShipment) HasRequireSignature() bool`

HasRequireSignature returns a boolean if a field has been set.

### GetShipOption

`func (o *InternalShipment) GetShipOption() string`

GetShipOption returns the ShipOption field if non-nil, zero value otherwise.

### GetShipOptionOk

`func (o *InternalShipment) GetShipOptionOk() (*string, bool)`

GetShipOptionOk returns a tuple with the ShipOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipOption

`func (o *InternalShipment) SetShipOption(v string)`

SetShipOption sets ShipOption field to given value.

### HasShipOption

`func (o *InternalShipment) HasShipOption() bool`

HasShipOption returns a boolean if a field has been set.

### GetStatus

`func (o *InternalShipment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalShipment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalShipment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalShipment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *InternalShipment) GetStatusDetails() []OrderStatusDetail`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *InternalShipment) GetStatusDetailsOk() (*[]OrderStatusDetail, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *InternalShipment) SetStatusDetails(v []OrderStatusDetail)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *InternalShipment) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetStoreOrderId

`func (o *InternalShipment) GetStoreOrderId() string`

GetStoreOrderId returns the StoreOrderId field if non-nil, zero value otherwise.

### GetStoreOrderIdOk

`func (o *InternalShipment) GetStoreOrderIdOk() (*string, bool)`

GetStoreOrderIdOk returns a tuple with the StoreOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOrderId

`func (o *InternalShipment) SetStoreOrderId(v string)`

SetStoreOrderId sets StoreOrderId field to given value.

### HasStoreOrderId

`func (o *InternalShipment) HasStoreOrderId() bool`

HasStoreOrderId returns a boolean if a field has been set.

### GetTracking

`func (o *InternalShipment) GetTracking() Tracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *InternalShipment) GetTrackingOk() (*Tracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *InternalShipment) SetTracking(v Tracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *InternalShipment) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


