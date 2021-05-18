# WebhooksWebhookViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Timestamp the webhook subscription was created | [optional] 
**Id** | Pointer to **int32** | ID of the webhook subscription | [optional] 
**SubscriptionUrl** | Pointer to **NullableString** | URL subscription events will be posted to | [optional] 
**Topic** | Pointer to [**WebhooksTopics**](Webhooks.Topics.md) |  | [optional] 

## Methods

### NewWebhooksWebhookViewModel

`func NewWebhooksWebhookViewModel() *WebhooksWebhookViewModel`

NewWebhooksWebhookViewModel instantiates a new WebhooksWebhookViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksWebhookViewModelWithDefaults

`func NewWebhooksWebhookViewModelWithDefaults() *WebhooksWebhookViewModel`

NewWebhooksWebhookViewModelWithDefaults instantiates a new WebhooksWebhookViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WebhooksWebhookViewModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhooksWebhookViewModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhooksWebhookViewModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhooksWebhookViewModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WebhooksWebhookViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhooksWebhookViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhooksWebhookViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebhooksWebhookViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubscriptionUrl

`func (o *WebhooksWebhookViewModel) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *WebhooksWebhookViewModel) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *WebhooksWebhookViewModel) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.

### HasSubscriptionUrl

`func (o *WebhooksWebhookViewModel) HasSubscriptionUrl() bool`

HasSubscriptionUrl returns a boolean if a field has been set.

### SetSubscriptionUrlNil

`func (o *WebhooksWebhookViewModel) SetSubscriptionUrlNil(b bool)`

 SetSubscriptionUrlNil sets the value for SubscriptionUrl to be an explicit nil

### UnsetSubscriptionUrl
`func (o *WebhooksWebhookViewModel) UnsetSubscriptionUrl()`

UnsetSubscriptionUrl ensures that no value is present for SubscriptionUrl, not even an explicit nil
### GetTopic

`func (o *WebhooksWebhookViewModel) GetTopic() WebhooksTopics`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhooksWebhookViewModel) GetTopicOk() (*WebhooksTopics, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhooksWebhookViewModel) SetTopic(v WebhooksTopics)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhooksWebhookViewModel) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


