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

// ShipbobProductsApiModelsPublicUpdateProductModel Updates to an existing product product
type ShipbobProductsApiModelsPublicUpdateProductModel struct {
	// Barcode for the product
	Barcode NullableString `json:"barcode,omitempty"`
	// The name of the product
	Name NullableString `json:"name"`
	// The stock keeping unit of the product
	Sku NullableString `json:"sku,omitempty"`
}

// NewShipbobProductsApiModelsPublicUpdateProductModel instantiates a new ShipbobProductsApiModelsPublicUpdateProductModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipbobProductsApiModelsPublicUpdateProductModel(name NullableString, ) *ShipbobProductsApiModelsPublicUpdateProductModel {
	this := ShipbobProductsApiModelsPublicUpdateProductModel{}
	this.Name = name
	return &this
}

// NewShipbobProductsApiModelsPublicUpdateProductModelWithDefaults instantiates a new ShipbobProductsApiModelsPublicUpdateProductModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipbobProductsApiModelsPublicUpdateProductModelWithDefaults() *ShipbobProductsApiModelsPublicUpdateProductModel {
	this := ShipbobProductsApiModelsPublicUpdateProductModel{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetBarcode() string {
	if o == nil || o.Barcode.Get() == nil {
		var ret string
		return ret
	}
	return *o.Barcode.Get()
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetBarcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Barcode.Get(), o.Barcode.IsSet()
}

// HasBarcode returns a boolean if a field has been set.
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) HasBarcode() bool {
	if o != nil && o.Barcode.IsSet() {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given NullableString and assigns it to the Barcode field.
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) SetBarcode(v string) {
	o.Barcode.Set(&v)
}
// SetBarcodeNil sets the value for Barcode to be an explicit nil
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) SetBarcodeNil() {
	o.Barcode.Set(nil)
}

// UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) UnsetBarcode() {
	o.Barcode.Unset()
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) SetName(v string) {
	o.Name.Set(&v)
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetSku() string {
	if o == nil || o.Sku.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) GetSkuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) SetSku(v string) {
	o.Sku.Set(&v)
}
// SetSkuNil sets the value for Sku to be an explicit nil
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *ShipbobProductsApiModelsPublicUpdateProductModel) UnsetSku() {
	o.Sku.Unset()
}

func (o ShipbobProductsApiModelsPublicUpdateProductModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Barcode.IsSet() {
		toSerialize["barcode"] = o.Barcode.Get()
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableShipbobProductsApiModelsPublicUpdateProductModel struct {
	value *ShipbobProductsApiModelsPublicUpdateProductModel
	isSet bool
}

func (v NullableShipbobProductsApiModelsPublicUpdateProductModel) Get() *ShipbobProductsApiModelsPublicUpdateProductModel {
	return v.value
}

func (v *NullableShipbobProductsApiModelsPublicUpdateProductModel) Set(val *ShipbobProductsApiModelsPublicUpdateProductModel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipbobProductsApiModelsPublicUpdateProductModel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipbobProductsApiModelsPublicUpdateProductModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipbobProductsApiModelsPublicUpdateProductModel(val *ShipbobProductsApiModelsPublicUpdateProductModel) *NullableShipbobProductsApiModelsPublicUpdateProductModel {
	return &NullableShipbobProductsApiModelsPublicUpdateProductModel{value: val, isSet: true}
}

func (v NullableShipbobProductsApiModelsPublicUpdateProductModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipbobProductsApiModelsPublicUpdateProductModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


