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
)

// OrdersShipMethodDetailViewModel struct for OrdersShipMethodDetailViewModel
type OrdersShipMethodDetailViewModel struct {
	// Indicates if the shipping method is active
	Active *bool `json:"active,omitempty"`
	// Indicates the shipping method is a ShipBob default shipping method.
	Default *bool `json:"default,omitempty"`
	// Unique id for shipping method.
	Id *int32 `json:"id,omitempty"`
	// Name of the ship method as selected by the merchant and saved in ShipBobâs database (i.e. âgroundâ). Corresponds to the shipping_method field in the Orders API.
	Name *string `json:"name,omitempty"`
	ServiceLevel *OrdersServiceLevelDetailViewModel `json:"service_level,omitempty"`
}

// NewOrdersShipMethodDetailViewModel instantiates a new OrdersShipMethodDetailViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersShipMethodDetailViewModel() *OrdersShipMethodDetailViewModel {
	this := OrdersShipMethodDetailViewModel{}
	return &this
}

// NewOrdersShipMethodDetailViewModelWithDefaults instantiates a new OrdersShipMethodDetailViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersShipMethodDetailViewModelWithDefaults() *OrdersShipMethodDetailViewModel {
	this := OrdersShipMethodDetailViewModel{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *OrdersShipMethodDetailViewModel) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipMethodDetailViewModel) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *OrdersShipMethodDetailViewModel) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *OrdersShipMethodDetailViewModel) SetActive(v bool) {
	o.Active = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OrdersShipMethodDetailViewModel) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipMethodDetailViewModel) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OrdersShipMethodDetailViewModel) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *OrdersShipMethodDetailViewModel) SetDefault(v bool) {
	o.Default = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrdersShipMethodDetailViewModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipMethodDetailViewModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrdersShipMethodDetailViewModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrdersShipMethodDetailViewModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrdersShipMethodDetailViewModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipMethodDetailViewModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrdersShipMethodDetailViewModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrdersShipMethodDetailViewModel) SetName(v string) {
	o.Name = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *OrdersShipMethodDetailViewModel) GetServiceLevel() OrdersServiceLevelDetailViewModel {
	if o == nil || o.ServiceLevel == nil {
		var ret OrdersServiceLevelDetailViewModel
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersShipMethodDetailViewModel) GetServiceLevelOk() (*OrdersServiceLevelDetailViewModel, bool) {
	if o == nil || o.ServiceLevel == nil {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *OrdersShipMethodDetailViewModel) HasServiceLevel() bool {
	if o != nil && o.ServiceLevel != nil {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given OrdersServiceLevelDetailViewModel and assigns it to the ServiceLevel field.
func (o *OrdersShipMethodDetailViewModel) SetServiceLevel(v OrdersServiceLevelDetailViewModel) {
	o.ServiceLevel = &v
}

func (o OrdersShipMethodDetailViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServiceLevel != nil {
		toSerialize["service_level"] = o.ServiceLevel
	}
	return json.Marshal(toSerialize)
}

type NullableOrdersShipMethodDetailViewModel struct {
	value *OrdersShipMethodDetailViewModel
	isSet bool
}

func (v NullableOrdersShipMethodDetailViewModel) Get() *OrdersShipMethodDetailViewModel {
	return v.value
}

func (v *NullableOrdersShipMethodDetailViewModel) Set(val *OrdersShipMethodDetailViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersShipMethodDetailViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersShipMethodDetailViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersShipMethodDetailViewModel(val *OrdersShipMethodDetailViewModel) *NullableOrdersShipMethodDetailViewModel {
	return &NullableOrdersShipMethodDetailViewModel{value: val, isSet: true}
}

func (v NullableOrdersShipMethodDetailViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersShipMethodDetailViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

