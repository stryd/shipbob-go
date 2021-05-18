# ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionUrl** | **string** | URL we will call when an event matching the subscription topic is raised. Must have ssl enabled (https) and accept POST requests with content type of application/json | 
**Topic** | [**ShipBobWebhooksPublicCommonTopics**](ShipBob.Webhooks.Public.Common.Topics.md) |  | 

## Methods

### NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel

`func NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel(subscriptionUrl string, topic ShipBobWebhooksPublicCommonTopics, ) *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel`

NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel instantiates a new ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModelWithDefaults

`func NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModelWithDefaults() *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel`

NewShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModelWithDefaults instantiates a new ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionUrl

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.


### GetTopic

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) GetTopic() ShipBobWebhooksPublicCommonTopics`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) GetTopicOk() (*ShipBobWebhooksPublicCommonTopics, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ShipBobWebhooksPublicApiModelsCreateWebhookSubscriptionModel) SetTopic(v ShipBobWebhooksPublicCommonTopics)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


