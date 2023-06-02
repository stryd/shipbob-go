/*
ShipBob Developer API

ShipBob Developer API Documentation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"time"
)

// checks if the Shipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Shipment{}

// Shipment Information about a shipment
type Shipment struct {
	// The datetime of ShipBob's completion of the fulfillment operation as promised.  Currently, this means the shipment has been picked, packed, and label has been printed.
	ActualFulfillmentDate NullableTime `json:"actual_fulfillment_date,omitempty"`
	// Date this shipment was created
	CreatedDate NullableTime `json:"created_date,omitempty"`
	// The datetime of ShipBob's commitment for completing  the shipment and handing to the carrier for delivery.
	EstimatedFulfillmentDate NullableTime `json:"estimated_fulfillment_date,omitempty"`
	// Status of ShipBob's completion of the fulfillment operation.
	EstimatedFulfillmentDateStatus *string `json:"estimated_fulfillment_date_status,omitempty"`
	// Unique id of the shipment
	Id *int32 `json:"id,omitempty"`
	// Monetary amount that this shipment was insured for
	InsuranceValue NullableFloat64 `json:"insurance_value,omitempty"`
	// Monetary amount that was invoiced for this shipment
	InvoiceAmount NullableFloat64 `json:"invoice_amount,omitempty"`
	// Date this shipment was last updated
	LastUpdateAt NullableTime       `json:"last_update_at,omitempty"`
	Location     *FulfillmentCenter `json:"location,omitempty"`
	Measurements *OrderMeasurements `json:"measurements,omitempty"`
	// Id of the order this shipment belongs to
	OrderId *int32 `json:"order_id,omitempty"`
	// Container type for the shipment
	PackageMaterialType *string `json:"package_material_type,omitempty"`
	// Information about the products contained in this shipment
	Products  []ShipmentProduct `json:"products,omitempty"`
	Recipient *Recipient        `json:"recipient,omitempty"`
	// Client-defined external unique id of the order this shipment belongs to
	ReferenceId *string `json:"reference_id,omitempty"`
	// If a shipment requires signature
	RequireSignature *bool `json:"require_signature,omitempty"`
	// Name of the shipping option used for this shipment
	ShipOption *string `json:"ship_option,omitempty"`
	// The shipment status
	Status *string `json:"status,omitempty"`
	// Additional details about the shipment status
	StatusDetails []OrderStatusDetail `json:"status_details,omitempty"`
	Tracking      *Tracking           `json:"tracking,omitempty"`
}

// NewShipment instantiates a new Shipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipment() *Shipment {
	this := Shipment{}
	return &this
}

// NewShipmentWithDefaults instantiates a new Shipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentWithDefaults() *Shipment {
	this := Shipment{}
	return &this
}

// GetActualFulfillmentDate returns the ActualFulfillmentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetActualFulfillmentDate() time.Time {
	if o == nil || IsNil(o.ActualFulfillmentDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ActualFulfillmentDate.Get()
}

// GetActualFulfillmentDateOk returns a tuple with the ActualFulfillmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetActualFulfillmentDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualFulfillmentDate.Get(), o.ActualFulfillmentDate.IsSet()
}

// HasActualFulfillmentDate returns a boolean if a field has been set.
func (o *Shipment) HasActualFulfillmentDate() bool {
	if o != nil && o.ActualFulfillmentDate.IsSet() {
		return true
	}

	return false
}

// SetActualFulfillmentDate gets a reference to the given NullableTime and assigns it to the ActualFulfillmentDate field.
func (o *Shipment) SetActualFulfillmentDate(v time.Time) {
	o.ActualFulfillmentDate.Set(&v)
}

// SetActualFulfillmentDateNil sets the value for ActualFulfillmentDate to be an explicit nil
func (o *Shipment) SetActualFulfillmentDateNil() {
	o.ActualFulfillmentDate.Set(nil)
}

// UnsetActualFulfillmentDate ensures that no value is present for ActualFulfillmentDate, not even an explicit nil
func (o *Shipment) UnsetActualFulfillmentDate() {
	o.ActualFulfillmentDate.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Shipment) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *Shipment) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}

// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *Shipment) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *Shipment) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetEstimatedFulfillmentDate returns the EstimatedFulfillmentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetEstimatedFulfillmentDate() time.Time {
	if o == nil || IsNil(o.EstimatedFulfillmentDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedFulfillmentDate.Get()
}

// GetEstimatedFulfillmentDateOk returns a tuple with the EstimatedFulfillmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetEstimatedFulfillmentDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedFulfillmentDate.Get(), o.EstimatedFulfillmentDate.IsSet()
}

// HasEstimatedFulfillmentDate returns a boolean if a field has been set.
func (o *Shipment) HasEstimatedFulfillmentDate() bool {
	if o != nil && o.EstimatedFulfillmentDate.IsSet() {
		return true
	}

	return false
}

// SetEstimatedFulfillmentDate gets a reference to the given NullableTime and assigns it to the EstimatedFulfillmentDate field.
func (o *Shipment) SetEstimatedFulfillmentDate(v time.Time) {
	o.EstimatedFulfillmentDate.Set(&v)
}

// SetEstimatedFulfillmentDateNil sets the value for EstimatedFulfillmentDate to be an explicit nil
func (o *Shipment) SetEstimatedFulfillmentDateNil() {
	o.EstimatedFulfillmentDate.Set(nil)
}

// UnsetEstimatedFulfillmentDate ensures that no value is present for EstimatedFulfillmentDate, not even an explicit nil
func (o *Shipment) UnsetEstimatedFulfillmentDate() {
	o.EstimatedFulfillmentDate.Unset()
}

// GetEstimatedFulfillmentDateStatus returns the EstimatedFulfillmentDateStatus field value if set, zero value otherwise.
func (o *Shipment) GetEstimatedFulfillmentDateStatus() string {
	if o == nil || IsNil(o.EstimatedFulfillmentDateStatus) {
		var ret string
		return ret
	}
	return *o.EstimatedFulfillmentDateStatus
}

// GetEstimatedFulfillmentDateStatusOk returns a tuple with the EstimatedFulfillmentDateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetEstimatedFulfillmentDateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedFulfillmentDateStatus) {
		return nil, false
	}
	return o.EstimatedFulfillmentDateStatus, true
}

// HasEstimatedFulfillmentDateStatus returns a boolean if a field has been set.
func (o *Shipment) HasEstimatedFulfillmentDateStatus() bool {
	if o != nil && !IsNil(o.EstimatedFulfillmentDateStatus) {
		return true
	}

	return false
}

// SetEstimatedFulfillmentDateStatus gets a reference to the given string and assigns it to the EstimatedFulfillmentDateStatus field.
func (o *Shipment) SetEstimatedFulfillmentDateStatus(v string) {
	o.EstimatedFulfillmentDateStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Shipment) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Shipment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Shipment) SetId(v int32) {
	o.Id = &v
}

// GetInsuranceValue returns the InsuranceValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetInsuranceValue() float64 {
	if o == nil || IsNil(o.InsuranceValue.Get()) {
		var ret float64
		return ret
	}
	return *o.InsuranceValue.Get()
}

// GetInsuranceValueOk returns a tuple with the InsuranceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetInsuranceValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsuranceValue.Get(), o.InsuranceValue.IsSet()
}

// HasInsuranceValue returns a boolean if a field has been set.
func (o *Shipment) HasInsuranceValue() bool {
	if o != nil && o.InsuranceValue.IsSet() {
		return true
	}

	return false
}

// SetInsuranceValue gets a reference to the given NullableFloat64 and assigns it to the InsuranceValue field.
func (o *Shipment) SetInsuranceValue(v float64) {
	o.InsuranceValue.Set(&v)
}

// SetInsuranceValueNil sets the value for InsuranceValue to be an explicit nil
func (o *Shipment) SetInsuranceValueNil() {
	o.InsuranceValue.Set(nil)
}

// UnsetInsuranceValue ensures that no value is present for InsuranceValue, not even an explicit nil
func (o *Shipment) UnsetInsuranceValue() {
	o.InsuranceValue.Unset()
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetInvoiceAmount() float64 {
	if o == nil || IsNil(o.InvoiceAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.InvoiceAmount.Get()
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetInvoiceAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceAmount.Get(), o.InvoiceAmount.IsSet()
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *Shipment) HasInvoiceAmount() bool {
	if o != nil && o.InvoiceAmount.IsSet() {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given NullableFloat64 and assigns it to the InvoiceAmount field.
func (o *Shipment) SetInvoiceAmount(v float64) {
	o.InvoiceAmount.Set(&v)
}

// SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil
func (o *Shipment) SetInvoiceAmountNil() {
	o.InvoiceAmount.Set(nil)
}

// UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
func (o *Shipment) UnsetInvoiceAmount() {
	o.InvoiceAmount.Unset()
}

// GetLastUpdateAt returns the LastUpdateAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shipment) GetLastUpdateAt() time.Time {
	if o == nil || IsNil(o.LastUpdateAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateAt.Get()
}

// GetLastUpdateAtOk returns a tuple with the LastUpdateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shipment) GetLastUpdateAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdateAt.Get(), o.LastUpdateAt.IsSet()
}

// HasLastUpdateAt returns a boolean if a field has been set.
func (o *Shipment) HasLastUpdateAt() bool {
	if o != nil && o.LastUpdateAt.IsSet() {
		return true
	}

	return false
}

// SetLastUpdateAt gets a reference to the given NullableTime and assigns it to the LastUpdateAt field.
func (o *Shipment) SetLastUpdateAt(v time.Time) {
	o.LastUpdateAt.Set(&v)
}

// SetLastUpdateAtNil sets the value for LastUpdateAt to be an explicit nil
func (o *Shipment) SetLastUpdateAtNil() {
	o.LastUpdateAt.Set(nil)
}

// UnsetLastUpdateAt ensures that no value is present for LastUpdateAt, not even an explicit nil
func (o *Shipment) UnsetLastUpdateAt() {
	o.LastUpdateAt.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Shipment) GetLocation() FulfillmentCenter {
	if o == nil || IsNil(o.Location) {
		var ret FulfillmentCenter
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetLocationOk() (*FulfillmentCenter, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Shipment) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given FulfillmentCenter and assigns it to the Location field.
func (o *Shipment) SetLocation(v FulfillmentCenter) {
	o.Location = &v
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *Shipment) GetMeasurements() OrderMeasurements {
	if o == nil || IsNil(o.Measurements) {
		var ret OrderMeasurements
		return ret
	}
	return *o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetMeasurementsOk() (*OrderMeasurements, bool) {
	if o == nil || IsNil(o.Measurements) {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *Shipment) HasMeasurements() bool {
	if o != nil && !IsNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given OrderMeasurements and assigns it to the Measurements field.
func (o *Shipment) SetMeasurements(v OrderMeasurements) {
	o.Measurements = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Shipment) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Shipment) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *Shipment) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetPackageMaterialType returns the PackageMaterialType field value if set, zero value otherwise.
func (o *Shipment) GetPackageMaterialType() string {
	if o == nil || IsNil(o.PackageMaterialType) {
		var ret string
		return ret
	}
	return *o.PackageMaterialType
}

// GetPackageMaterialTypeOk returns a tuple with the PackageMaterialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetPackageMaterialTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageMaterialType) {
		return nil, false
	}
	return o.PackageMaterialType, true
}

// HasPackageMaterialType returns a boolean if a field has been set.
func (o *Shipment) HasPackageMaterialType() bool {
	if o != nil && !IsNil(o.PackageMaterialType) {
		return true
	}

	return false
}

// SetPackageMaterialType gets a reference to the given string and assigns it to the PackageMaterialType field.
func (o *Shipment) SetPackageMaterialType(v string) {
	o.PackageMaterialType = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Shipment) GetProducts() []ShipmentProduct {
	if o == nil || IsNil(o.Products) {
		var ret []ShipmentProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetProductsOk() ([]ShipmentProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Shipment) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ShipmentProduct and assigns it to the Products field.
func (o *Shipment) SetProducts(v []ShipmentProduct) {
	o.Products = v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *Shipment) GetRecipient() Recipient {
	if o == nil || IsNil(o.Recipient) {
		var ret Recipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetRecipientOk() (*Recipient, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *Shipment) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given Recipient and assigns it to the Recipient field.
func (o *Shipment) SetRecipient(v Recipient) {
	o.Recipient = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *Shipment) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *Shipment) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *Shipment) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetRequireSignature returns the RequireSignature field value if set, zero value otherwise.
func (o *Shipment) GetRequireSignature() bool {
	if o == nil || IsNil(o.RequireSignature) {
		var ret bool
		return ret
	}
	return *o.RequireSignature
}

// GetRequireSignatureOk returns a tuple with the RequireSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetRequireSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignature) {
		return nil, false
	}
	return o.RequireSignature, true
}

// HasRequireSignature returns a boolean if a field has been set.
func (o *Shipment) HasRequireSignature() bool {
	if o != nil && !IsNil(o.RequireSignature) {
		return true
	}

	return false
}

// SetRequireSignature gets a reference to the given bool and assigns it to the RequireSignature field.
func (o *Shipment) SetRequireSignature(v bool) {
	o.RequireSignature = &v
}

// GetShipOption returns the ShipOption field value if set, zero value otherwise.
func (o *Shipment) GetShipOption() string {
	if o == nil || IsNil(o.ShipOption) {
		var ret string
		return ret
	}
	return *o.ShipOption
}

// GetShipOptionOk returns a tuple with the ShipOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetShipOptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShipOption) {
		return nil, false
	}
	return o.ShipOption, true
}

// HasShipOption returns a boolean if a field has been set.
func (o *Shipment) HasShipOption() bool {
	if o != nil && !IsNil(o.ShipOption) {
		return true
	}

	return false
}

// SetShipOption gets a reference to the given string and assigns it to the ShipOption field.
func (o *Shipment) SetShipOption(v string) {
	o.ShipOption = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Shipment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Shipment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Shipment) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *Shipment) GetStatusDetails() []OrderStatusDetail {
	if o == nil || IsNil(o.StatusDetails) {
		var ret []OrderStatusDetail
		return ret
	}
	return o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetStatusDetailsOk() ([]OrderStatusDetail, bool) {
	if o == nil || IsNil(o.StatusDetails) {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *Shipment) HasStatusDetails() bool {
	if o != nil && !IsNil(o.StatusDetails) {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given []OrderStatusDetail and assigns it to the StatusDetails field.
func (o *Shipment) SetStatusDetails(v []OrderStatusDetail) {
	o.StatusDetails = v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *Shipment) GetTracking() Tracking {
	if o == nil || IsNil(o.Tracking) {
		var ret Tracking
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetTrackingOk() (*Tracking, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *Shipment) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given Tracking and assigns it to the Tracking field.
func (o *Shipment) SetTracking(v Tracking) {
	o.Tracking = &v
}

func (o Shipment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Shipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualFulfillmentDate.IsSet() {
		toSerialize["actual_fulfillment_date"] = o.ActualFulfillmentDate.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["created_date"] = o.CreatedDate.Get()
	}
	if o.EstimatedFulfillmentDate.IsSet() {
		toSerialize["estimated_fulfillment_date"] = o.EstimatedFulfillmentDate.Get()
	}
	if !IsNil(o.EstimatedFulfillmentDateStatus) {
		toSerialize["estimated_fulfillment_date_status"] = o.EstimatedFulfillmentDateStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.InsuranceValue.IsSet() {
		toSerialize["insurance_value"] = o.InsuranceValue.Get()
	}
	if o.InvoiceAmount.IsSet() {
		toSerialize["invoice_amount"] = o.InvoiceAmount.Get()
	}
	if o.LastUpdateAt.IsSet() {
		toSerialize["last_update_at"] = o.LastUpdateAt.Get()
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Measurements) {
		toSerialize["measurements"] = o.Measurements
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.PackageMaterialType) {
		toSerialize["package_material_type"] = o.PackageMaterialType
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.RequireSignature) {
		toSerialize["require_signature"] = o.RequireSignature
	}
	if !IsNil(o.ShipOption) {
		toSerialize["ship_option"] = o.ShipOption
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetails) {
		toSerialize["status_details"] = o.StatusDetails
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	return toSerialize, nil
}

type NullableShipment struct {
	value *Shipment
	isSet bool
}

func (v NullableShipment) Get() *Shipment {
	return v.value
}

func (v *NullableShipment) Set(val *Shipment) {
	v.value = val
	v.isSet = true
}

func (v NullableShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipment(val *Shipment) *NullableShipment {
	return &NullableShipment{value: val, isSet: true}
}

func (v NullableShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
