# WebhooksConfigWithStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**WebhooksConfigDto**](WebhooksConfigDto.md) | The webhook configuration. | [optional] 
**Status** | Pointer to **int32** | The webhook status. | [optional] 

## Methods

### NewWebhooksConfigWithStatusDto

`func NewWebhooksConfigWithStatusDto() *WebhooksConfigWithStatusDto`

NewWebhooksConfigWithStatusDto instantiates a new WebhooksConfigWithStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksConfigWithStatusDtoWithDefaults

`func NewWebhooksConfigWithStatusDtoWithDefaults() *WebhooksConfigWithStatusDto`

NewWebhooksConfigWithStatusDtoWithDefaults instantiates a new WebhooksConfigWithStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *WebhooksConfigWithStatusDto) GetConfigs() WebhooksConfigDto`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *WebhooksConfigWithStatusDto) GetConfigsOk() (*WebhooksConfigDto, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *WebhooksConfigWithStatusDto) SetConfigs(v WebhooksConfigDto)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *WebhooksConfigWithStatusDto) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetStatus

`func (o *WebhooksConfigWithStatusDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhooksConfigWithStatusDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhooksConfigWithStatusDto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhooksConfigWithStatusDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


