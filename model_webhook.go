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

// Webhook struct for Webhook
type Webhook struct {
	// Timestamp the webhook subscription was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ID of the webhook subscription
	Id *int32 `json:"id,omitempty"`
	// URL subscription events will be posted to
	SubscriptionUrl NullableString `json:"subscription_url,omitempty"`
	Topic           *Topics        `json:"topic,omitempty"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook() *Webhook {
	this := Webhook{}
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Webhook) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Webhook) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Webhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Webhook) SetId(v int32) {
	o.Id = &v
}

// GetSubscriptionUrl returns the SubscriptionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Webhook) GetSubscriptionUrl() string {
	if o == nil || o.SubscriptionUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionUrl.Get()
}

// GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webhook) GetSubscriptionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionUrl.Get(), o.SubscriptionUrl.IsSet()
}

// HasSubscriptionUrl returns a boolean if a field has been set.
func (o *Webhook) HasSubscriptionUrl() bool {
	if o != nil && o.SubscriptionUrl.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionUrl gets a reference to the given NullableString and assigns it to the SubscriptionUrl field.
func (o *Webhook) SetSubscriptionUrl(v string) {
	o.SubscriptionUrl.Set(&v)
}

// SetSubscriptionUrlNil sets the value for SubscriptionUrl to be an explicit nil
func (o *Webhook) SetSubscriptionUrlNil() {
	o.SubscriptionUrl.Set(nil)
}

// UnsetSubscriptionUrl ensures that no value is present for SubscriptionUrl, not even an explicit nil
func (o *Webhook) UnsetSubscriptionUrl() {
	o.SubscriptionUrl.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *Webhook) GetTopic() Topics {
	if o == nil || o.Topic == nil {
		var ret Topics
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTopicOk() (*Topics, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *Webhook) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given Topics and assigns it to the Topic field.
func (o *Webhook) SetTopic(v Topics) {
	o.Topic = &v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SubscriptionUrl.IsSet() {
		toSerialize["subscription_url"] = o.SubscriptionUrl.Get()
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	return json.Marshal(toSerialize)
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}