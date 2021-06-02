# WebhookSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionUrl** | **string** | URL we will call when an event matching the subscription topic is raised. Must have ssl enabled (https) and accept POST requests with content type of application/json | 
**Topic** | [**Topics**](Topics.md) |  | 

## Methods

### NewWebhookSubscription

`func NewWebhookSubscription(subscriptionUrl string, topic Topics, ) *WebhookSubscription`

NewWebhookSubscription instantiates a new WebhookSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionWithDefaults

`func NewWebhookSubscriptionWithDefaults() *WebhookSubscription`

NewWebhookSubscriptionWithDefaults instantiates a new WebhookSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionUrl

`func (o *WebhookSubscription) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *WebhookSubscription) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *WebhookSubscription) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTopic

`func (o *WebhookSubscription) GetTopic() Topics`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookSubscription) GetTopicOk() (*Topics, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookSubscription) SetTopic(v Topics)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


