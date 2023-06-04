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

// checks if the CreateReceivingOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReceivingOrder{}

// CreateReceivingOrder Information to create a new receiving order
type CreateReceivingOrder struct {
	BoxPackagingType PackingType `json:"box_packaging_type"`
	// Box shipments to be added to this receiving order
	Boxes []AddBoxToOrder `json:"boxes"`
	// Expected arrival date of all the box shipments in this receiving order
	ExpectedArrivalDate NullableTime `json:"expected_arrival_date"`
	FulfillmentCenter AssignOrderToFulfillmentCenter `json:"fulfillment_center"`
	PackageType PackageType `json:"package_type"`
	// Purchase order number for this receiving order
	PurchaseOrderNumber NullableString `json:"purchase_order_number,omitempty"`
}

// NewCreateReceivingOrder instantiates a new CreateReceivingOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceivingOrder(boxPackagingType PackingType, boxes []AddBoxToOrder, expectedArrivalDate NullableTime, fulfillmentCenter AssignOrderToFulfillmentCenter, packageType PackageType) *CreateReceivingOrder {
	this := CreateReceivingOrder{}
	this.BoxPackagingType = boxPackagingType
	this.Boxes = boxes
	this.ExpectedArrivalDate = expectedArrivalDate
	this.FulfillmentCenter = fulfillmentCenter
	this.PackageType = packageType
	return &this
}

// NewCreateReceivingOrderWithDefaults instantiates a new CreateReceivingOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceivingOrderWithDefaults() *CreateReceivingOrder {
	this := CreateReceivingOrder{}
	return &this
}

// GetBoxPackagingType returns the BoxPackagingType field value
func (o *CreateReceivingOrder) GetBoxPackagingType() PackingType {
	if o == nil {
		var ret PackingType
		return ret
	}

	return o.BoxPackagingType
}

// GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetBoxPackagingTypeOk() (*PackingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoxPackagingType, true
}

// SetBoxPackagingType sets field value
func (o *CreateReceivingOrder) SetBoxPackagingType(v PackingType) {
	o.BoxPackagingType = v
}

// GetBoxes returns the Boxes field value
// If the value is explicit nil, the zero value for []AddBoxToOrder will be returned
func (o *CreateReceivingOrder) GetBoxes() []AddBoxToOrder {
	if o == nil {
		var ret []AddBoxToOrder
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrder) GetBoxesOk() ([]AddBoxToOrder, bool) {
	if o == nil || IsNil(o.Boxes) {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *CreateReceivingOrder) SetBoxes(v []AddBoxToOrder) {
	o.Boxes = v
}

// GetExpectedArrivalDate returns the ExpectedArrivalDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CreateReceivingOrder) GetExpectedArrivalDate() time.Time {
	if o == nil || o.ExpectedArrivalDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpectedArrivalDate.Get()
}

// GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrder) GetExpectedArrivalDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedArrivalDate.Get(), o.ExpectedArrivalDate.IsSet()
}

// SetExpectedArrivalDate sets field value
func (o *CreateReceivingOrder) SetExpectedArrivalDate(v time.Time) {
	o.ExpectedArrivalDate.Set(&v)
}

// GetFulfillmentCenter returns the FulfillmentCenter field value
func (o *CreateReceivingOrder) GetFulfillmentCenter() AssignOrderToFulfillmentCenter {
	if o == nil {
		var ret AssignOrderToFulfillmentCenter
		return ret
	}

	return o.FulfillmentCenter
}

// GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetFulfillmentCenterOk() (*AssignOrderToFulfillmentCenter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentCenter, true
}

// SetFulfillmentCenter sets field value
func (o *CreateReceivingOrder) SetFulfillmentCenter(v AssignOrderToFulfillmentCenter) {
	o.FulfillmentCenter = v
}

// GetPackageType returns the PackageType field value
func (o *CreateReceivingOrder) GetPackageType() PackageType {
	if o == nil {
		var ret PackageType
		return ret
	}

	return o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value
// and a boolean to check if the value has been set.
func (o *CreateReceivingOrder) GetPackageTypeOk() (*PackageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageType, true
}

// SetPackageType sets field value
func (o *CreateReceivingOrder) SetPackageType(v PackageType) {
	o.PackageType = v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateReceivingOrder) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber.Get()
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReceivingOrder) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchaseOrderNumber.Get(), o.PurchaseOrderNumber.IsSet()
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *CreateReceivingOrder) HasPurchaseOrderNumber() bool {
	if o != nil && o.PurchaseOrderNumber.IsSet() {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given NullableString and assigns it to the PurchaseOrderNumber field.
func (o *CreateReceivingOrder) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber.Set(&v)
}
// SetPurchaseOrderNumberNil sets the value for PurchaseOrderNumber to be an explicit nil
func (o *CreateReceivingOrder) SetPurchaseOrderNumberNil() {
	o.PurchaseOrderNumber.Set(nil)
}

// UnsetPurchaseOrderNumber ensures that no value is present for PurchaseOrderNumber, not even an explicit nil
func (o *CreateReceivingOrder) UnsetPurchaseOrderNumber() {
	o.PurchaseOrderNumber.Unset()
}

func (o CreateReceivingOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReceivingOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["box_packaging_type"] = o.BoxPackagingType
	if o.Boxes != nil {
		toSerialize["boxes"] = o.Boxes
	}
	toSerialize["expected_arrival_date"] = o.ExpectedArrivalDate.Get()
	toSerialize["fulfillment_center"] = o.FulfillmentCenter
	toSerialize["package_type"] = o.PackageType
	if o.PurchaseOrderNumber.IsSet() {
		toSerialize["purchase_order_number"] = o.PurchaseOrderNumber.Get()
	}
	return toSerialize, nil
}

type NullableCreateReceivingOrder struct {
	value *CreateReceivingOrder
	isSet bool
}

func (v NullableCreateReceivingOrder) Get() *CreateReceivingOrder {
	return v.value
}

func (v *NullableCreateReceivingOrder) Set(val *CreateReceivingOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceivingOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceivingOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceivingOrder(val *CreateReceivingOrder) *NullableCreateReceivingOrder {
	return &NullableCreateReceivingOrder{value: val, isSet: true}
}

func (v NullableCreateReceivingOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceivingOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


