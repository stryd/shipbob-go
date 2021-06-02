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

// CreateOrderProducts struct for CreateOrderProducts
type CreateOrderProducts struct {
	// Unique id of the product
	Id *int32 `json:"id,omitempty"`
	// The quantity of this product ordered
	Quantity int32 `json:"quantity"`
	// Name of the product
	Name *string `json:"name,omitempty"`
	// Unique reference id of the product
	ReferenceId *string `json:"reference_id,omitempty"`
}

// NewCreateOrderProducts instantiates a new CreateOrderProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderProducts(quantity int32) *CreateOrderProducts {
	this := CreateOrderProducts{}
	this.Quantity = quantity
	return &this
}

// NewCreateOrderProductsWithDefaults instantiates a new CreateOrderProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderProductsWithDefaults() *CreateOrderProducts {
	this := CreateOrderProducts{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrderProducts) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderProducts) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrderProducts) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateOrderProducts) SetId(v int32) {
	o.Id = &v
}

// GetQuantity returns the Quantity field value
func (o *CreateOrderProducts) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateOrderProducts) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateOrderProducts) SetQuantity(v int32) {
	o.Quantity = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrderProducts) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderProducts) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrderProducts) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrderProducts) SetName(v string) {
	o.Name = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *CreateOrderProducts) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderProducts) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *CreateOrderProducts) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *CreateOrderProducts) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o CreateOrderProducts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ReferenceId != nil {
		toSerialize["reference_id"] = o.ReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrderProducts struct {
	value *CreateOrderProducts
	isSet bool
}

func (v NullableCreateOrderProducts) Get() *CreateOrderProducts {
	return v.value
}

func (v *NullableCreateOrderProducts) Set(val *CreateOrderProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderProducts(val *CreateOrderProducts) *NullableCreateOrderProducts {
	return &NullableCreateOrderProducts{value: val, isSet: true}
}

func (v NullableCreateOrderProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
