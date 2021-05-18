/*
 * ShipBob Developer API
 *
 * ShipBob Developer API Documentation
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipbob

import (
	"encoding/json"
	"time"
)

// ShipbobReceivingPublicApiModelsReceivingOrderViewModel Information about a receiving order
type ShipbobReceivingPublicApiModelsReceivingOrderViewModel struct {
	// URL to the packing slip to be included in each box shipment for this receiving order
	BoxLabelsUri NullableString `json:"box_labels_uri,omitempty"`
	BoxPackagingType *ShipbobReceivingPublicCommonModelsPackingType `json:"box_packaging_type,omitempty"`
	// Information about the boxes being shipped in this receiving order
	Boxes []ShipbobReceivingPublicApiModelsBoxViewModel `json:"boxes,omitempty"`
	// Expected date that all packages will have arrived
	ExpectedArrivalDate *time.Time `json:"expected_arrival_date,omitempty"`
	FulfillmentCenter *ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel `json:"fulfillment_center,omitempty"`
	// Unique id of the warehouse receiving order
	Id *int32 `json:"id,omitempty"`
	// Insert date of the receiving order
	InsertDate *time.Time `json:"insert_date,omitempty"`
	// Last date the receiving order was updated
	LastUpdatedDate *time.Time `json:"last_updated_date,omitempty"`
	PackageType *ShipbobReceivingPublicCommonModelsPackageType `json:"package_type,omitempty"`
	Status *ShipbobReceivingPublicCommonModelsReceivingStatus `json:"status,omitempty"`
}

// NewShipbobReceivingPublicApiModelsReceivingOrderViewModel instantiates a new ShipbobReceivingPublicApiModelsReceivingOrderViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipbobReceivingPublicApiModelsReceivingOrderViewModel() *ShipbobReceivingPublicApiModelsReceivingOrderViewModel {
	this := ShipbobReceivingPublicApiModelsReceivingOrderViewModel{}
	return &this
}

// NewShipbobReceivingPublicApiModelsReceivingOrderViewModelWithDefaults instantiates a new ShipbobReceivingPublicApiModelsReceivingOrderViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipbobReceivingPublicApiModelsReceivingOrderViewModelWithDefaults() *ShipbobReceivingPublicApiModelsReceivingOrderViewModel {
	this := ShipbobReceivingPublicApiModelsReceivingOrderViewModel{}
	return &this
}

// GetBoxLabelsUri returns the BoxLabelsUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxLabelsUri() string {
	if o == nil || o.BoxLabelsUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.BoxLabelsUri.Get()
}

// GetBoxLabelsUriOk returns a tuple with the BoxLabelsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxLabelsUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BoxLabelsUri.Get(), o.BoxLabelsUri.IsSet()
}

// HasBoxLabelsUri returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxLabelsUri() bool {
	if o != nil && o.BoxLabelsUri.IsSet() {
		return true
	}

	return false
}

// SetBoxLabelsUri gets a reference to the given NullableString and assigns it to the BoxLabelsUri field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxLabelsUri(v string) {
	o.BoxLabelsUri.Set(&v)
}
// SetBoxLabelsUriNil sets the value for BoxLabelsUri to be an explicit nil
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxLabelsUriNil() {
	o.BoxLabelsUri.Set(nil)
}

// UnsetBoxLabelsUri ensures that no value is present for BoxLabelsUri, not even an explicit nil
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) UnsetBoxLabelsUri() {
	o.BoxLabelsUri.Unset()
}

// GetBoxPackagingType returns the BoxPackagingType field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxPackagingType() ShipbobReceivingPublicCommonModelsPackingType {
	if o == nil || o.BoxPackagingType == nil {
		var ret ShipbobReceivingPublicCommonModelsPackingType
		return ret
	}
	return *o.BoxPackagingType
}

// GetBoxPackagingTypeOk returns a tuple with the BoxPackagingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxPackagingTypeOk() (*ShipbobReceivingPublicCommonModelsPackingType, bool) {
	if o == nil || o.BoxPackagingType == nil {
		return nil, false
	}
	return o.BoxPackagingType, true
}

// HasBoxPackagingType returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxPackagingType() bool {
	if o != nil && o.BoxPackagingType != nil {
		return true
	}

	return false
}

// SetBoxPackagingType gets a reference to the given ShipbobReceivingPublicCommonModelsPackingType and assigns it to the BoxPackagingType field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxPackagingType(v ShipbobReceivingPublicCommonModelsPackingType) {
	o.BoxPackagingType = &v
}

// GetBoxes returns the Boxes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxes() []ShipbobReceivingPublicApiModelsBoxViewModel {
	if o == nil  {
		var ret []ShipbobReceivingPublicApiModelsBoxViewModel
		return ret
	}
	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetBoxesOk() (*[]ShipbobReceivingPublicApiModelsBoxViewModel, bool) {
	if o == nil || o.Boxes == nil {
		return nil, false
	}
	return &o.Boxes, true
}

// HasBoxes returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasBoxes() bool {
	if o != nil && o.Boxes != nil {
		return true
	}

	return false
}

// SetBoxes gets a reference to the given []ShipbobReceivingPublicApiModelsBoxViewModel and assigns it to the Boxes field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetBoxes(v []ShipbobReceivingPublicApiModelsBoxViewModel) {
	o.Boxes = v
}

// GetExpectedArrivalDate returns the ExpectedArrivalDate field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetExpectedArrivalDate() time.Time {
	if o == nil || o.ExpectedArrivalDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpectedArrivalDate
}

// GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetExpectedArrivalDateOk() (*time.Time, bool) {
	if o == nil || o.ExpectedArrivalDate == nil {
		return nil, false
	}
	return o.ExpectedArrivalDate, true
}

// HasExpectedArrivalDate returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasExpectedArrivalDate() bool {
	if o != nil && o.ExpectedArrivalDate != nil {
		return true
	}

	return false
}

// SetExpectedArrivalDate gets a reference to the given time.Time and assigns it to the ExpectedArrivalDate field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetExpectedArrivalDate(v time.Time) {
	o.ExpectedArrivalDate = &v
}

// GetFulfillmentCenter returns the FulfillmentCenter field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetFulfillmentCenter() ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel {
	if o == nil || o.FulfillmentCenter == nil {
		var ret ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel
		return ret
	}
	return *o.FulfillmentCenter
}

// GetFulfillmentCenterOk returns a tuple with the FulfillmentCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetFulfillmentCenterOk() (*ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel, bool) {
	if o == nil || o.FulfillmentCenter == nil {
		return nil, false
	}
	return o.FulfillmentCenter, true
}

// HasFulfillmentCenter returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasFulfillmentCenter() bool {
	if o != nil && o.FulfillmentCenter != nil {
		return true
	}

	return false
}

// SetFulfillmentCenter gets a reference to the given ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel and assigns it to the FulfillmentCenter field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetFulfillmentCenter(v ShipbobReceivingPublicApiModelsFulfillmentCenterViewModel) {
	o.FulfillmentCenter = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetId(v int32) {
	o.Id = &v
}

// GetInsertDate returns the InsertDate field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetInsertDate() time.Time {
	if o == nil || o.InsertDate == nil {
		var ret time.Time
		return ret
	}
	return *o.InsertDate
}

// GetInsertDateOk returns a tuple with the InsertDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetInsertDateOk() (*time.Time, bool) {
	if o == nil || o.InsertDate == nil {
		return nil, false
	}
	return o.InsertDate, true
}

// HasInsertDate returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasInsertDate() bool {
	if o != nil && o.InsertDate != nil {
		return true
	}

	return false
}

// SetInsertDate gets a reference to the given time.Time and assigns it to the InsertDate field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetInsertDate(v time.Time) {
	o.InsertDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetLastUpdatedDate() time.Time {
	if o == nil || o.LastUpdatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedDate == nil {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasLastUpdatedDate() bool {
	if o != nil && o.LastUpdatedDate != nil {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetPackageType returns the PackageType field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetPackageType() ShipbobReceivingPublicCommonModelsPackageType {
	if o == nil || o.PackageType == nil {
		var ret ShipbobReceivingPublicCommonModelsPackageType
		return ret
	}
	return *o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetPackageTypeOk() (*ShipbobReceivingPublicCommonModelsPackageType, bool) {
	if o == nil || o.PackageType == nil {
		return nil, false
	}
	return o.PackageType, true
}

// HasPackageType returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasPackageType() bool {
	if o != nil && o.PackageType != nil {
		return true
	}

	return false
}

// SetPackageType gets a reference to the given ShipbobReceivingPublicCommonModelsPackageType and assigns it to the PackageType field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetPackageType(v ShipbobReceivingPublicCommonModelsPackageType) {
	o.PackageType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetStatus() ShipbobReceivingPublicCommonModelsReceivingStatus {
	if o == nil || o.Status == nil {
		var ret ShipbobReceivingPublicCommonModelsReceivingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) GetStatusOk() (*ShipbobReceivingPublicCommonModelsReceivingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ShipbobReceivingPublicCommonModelsReceivingStatus and assigns it to the Status field.
func (o *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) SetStatus(v ShipbobReceivingPublicCommonModelsReceivingStatus) {
	o.Status = &v
}

func (o ShipbobReceivingPublicApiModelsReceivingOrderViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoxLabelsUri.IsSet() {
		toSerialize["box_labels_uri"] = o.BoxLabelsUri.Get()
	}
	if o.BoxPackagingType != nil {
		toSerialize["box_packaging_type"] = o.BoxPackagingType
	}
	if o.Boxes != nil {
		toSerialize["boxes"] = o.Boxes
	}
	if o.ExpectedArrivalDate != nil {
		toSerialize["expected_arrival_date"] = o.ExpectedArrivalDate
	}
	if o.FulfillmentCenter != nil {
		toSerialize["fulfillment_center"] = o.FulfillmentCenter
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InsertDate != nil {
		toSerialize["insert_date"] = o.InsertDate
	}
	if o.LastUpdatedDate != nil {
		toSerialize["last_updated_date"] = o.LastUpdatedDate
	}
	if o.PackageType != nil {
		toSerialize["package_type"] = o.PackageType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel struct {
	value *ShipbobReceivingPublicApiModelsReceivingOrderViewModel
	isSet bool
}

func (v NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) Get() *ShipbobReceivingPublicApiModelsReceivingOrderViewModel {
	return v.value
}

func (v *NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) Set(val *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipbobReceivingPublicApiModelsReceivingOrderViewModel(val *ShipbobReceivingPublicApiModelsReceivingOrderViewModel) *NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel {
	return &NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel{value: val, isSet: true}
}

func (v NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipbobReceivingPublicApiModelsReceivingOrderViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


