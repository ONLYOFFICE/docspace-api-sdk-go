# CreateWebhooksConfigRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The human-readable name of the webhook configuration. | 
**Uri** | **string** | The destination URL where the webhook events will be sent. | 
**SecretKey** | Pointer to **NullableString** | The webhook secret key used to sign the webhook payloads for the security verification. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the webhook configuration is active or not. | [optional] 
**Ssl** | Pointer to **bool** | Specifies whether the SSL certificate verification is required or not. | [optional] 
**Triggers** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) |  | [optional] 
**TargetId** | Pointer to **NullableString** | Target ID | [optional] 

## Methods

### NewCreateWebhooksConfigRequestsDto

`func NewCreateWebhooksConfigRequestsDto(name string, uri string, ) *CreateWebhooksConfigRequestsDto`

NewCreateWebhooksConfigRequestsDto instantiates a new CreateWebhooksConfigRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhooksConfigRequestsDtoWithDefaults

`func NewCreateWebhooksConfigRequestsDtoWithDefaults() *CreateWebhooksConfigRequestsDto`

NewCreateWebhooksConfigRequestsDtoWithDefaults instantiates a new CreateWebhooksConfigRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhooksConfigRequestsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhooksConfigRequestsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhooksConfigRequestsDto) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *CreateWebhooksConfigRequestsDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CreateWebhooksConfigRequestsDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CreateWebhooksConfigRequestsDto) SetUri(v string)`

SetUri sets Uri field to given value.


### GetSecretKey

`func (o *CreateWebhooksConfigRequestsDto) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CreateWebhooksConfigRequestsDto) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CreateWebhooksConfigRequestsDto) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CreateWebhooksConfigRequestsDto) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *CreateWebhooksConfigRequestsDto) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *CreateWebhooksConfigRequestsDto) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetEnabled

`func (o *CreateWebhooksConfigRequestsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWebhooksConfigRequestsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWebhooksConfigRequestsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWebhooksConfigRequestsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsl

`func (o *CreateWebhooksConfigRequestsDto) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *CreateWebhooksConfigRequestsDto) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *CreateWebhooksConfigRequestsDto) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *CreateWebhooksConfigRequestsDto) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTriggers

`func (o *CreateWebhooksConfigRequestsDto) GetTriggers() WebhookTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CreateWebhooksConfigRequestsDto) GetTriggersOk() (*WebhookTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CreateWebhooksConfigRequestsDto) SetTriggers(v WebhookTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CreateWebhooksConfigRequestsDto) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargetId

`func (o *CreateWebhooksConfigRequestsDto) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CreateWebhooksConfigRequestsDto) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CreateWebhooksConfigRequestsDto) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CreateWebhooksConfigRequestsDto) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *CreateWebhooksConfigRequestsDto) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *CreateWebhooksConfigRequestsDto) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


