# ShipBobWebhooksPublicApiModelsWebhookViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Timestamp the webhook subscription was created | [optional] 
**Id** | Pointer to **int32** | ID of the webhook subscription | [optional] 
**SubscriptionUrl** | Pointer to **NullableString** | URL subscription events will be posted to | [optional] 
**Topic** | Pointer to [**ShipBobWebhooksPublicCommonTopics**](ShipBob.Webhooks.Public.Common.Topics.md) |  | [optional] 

## Methods

### NewShipBobWebhooksPublicApiModelsWebhookViewModel

`func NewShipBobWebhooksPublicApiModelsWebhookViewModel() *ShipBobWebhooksPublicApiModelsWebhookViewModel`

NewShipBobWebhooksPublicApiModelsWebhookViewModel instantiates a new ShipBobWebhooksPublicApiModelsWebhookViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipBobWebhooksPublicApiModelsWebhookViewModelWithDefaults

`func NewShipBobWebhooksPublicApiModelsWebhookViewModelWithDefaults() *ShipBobWebhooksPublicApiModelsWebhookViewModel`

NewShipBobWebhooksPublicApiModelsWebhookViewModelWithDefaults instantiates a new ShipBobWebhooksPublicApiModelsWebhookViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubscriptionUrl

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetSubscriptionUrl() string`

GetSubscriptionUrl returns the SubscriptionUrl field if non-nil, zero value otherwise.

### GetSubscriptionUrlOk

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetSubscriptionUrlOk() (*string, bool)`

GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionUrl

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) SetSubscriptionUrl(v string)`

SetSubscriptionUrl sets SubscriptionUrl field to given value.

### HasSubscriptionUrl

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) HasSubscriptionUrl() bool`

HasSubscriptionUrl returns a boolean if a field has been set.

### SetSubscriptionUrlNil

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) SetSubscriptionUrlNil(b bool)`

 SetSubscriptionUrlNil sets the value for SubscriptionUrl to be an explicit nil

### UnsetSubscriptionUrl
`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) UnsetSubscriptionUrl()`

UnsetSubscriptionUrl ensures that no value is present for SubscriptionUrl, not even an explicit nil
### GetTopic

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetTopic() ShipBobWebhooksPublicCommonTopics`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) GetTopicOk() (*ShipBobWebhooksPublicCommonTopics, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) SetTopic(v ShipBobWebhooksPublicCommonTopics)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ShipBobWebhooksPublicApiModelsWebhookViewModel) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


