# CreateWebhooksConfigRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The label the subscription is listed under. It is for the administrator reading the list and is never sent to  the target; it does not have to be unique. | 
**Uri** | **string** | The address the portal posts the event payload to. It has to be an absolute `http` or `https` address outside  the installation own network, and it is probed before anything is stored: it must answer a HEAD request with  a success code, and a redirect does not count as one. | 
**SecretKey** | Pointer to **NullableString** | The shared secret the payload signature is computed with, so the receiver can tell a genuine call from a  forged one. It has to satisfy the portal password rules published by  `GET api/2.0/settings/security/password`, and it is never echoed back by any operation. On an update an empty  value keeps the secret already stored. | [optional] 
**Enabled** | Pointer to **bool** | Whether the subscription delivers at all. While it is off the matching events are dropped rather than queued,  so nothing from that period arrives once it is switched on again. | [optional] 
**Ssl** | Pointer to **bool** | Whether the target certificate is verified. Setting it demands an `https` target with a valid certificate;  leaving it off delivers without checking the certificate at all. | [optional] 
**Triggers** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) | The events the subscription listens for, as a bitmask combining the flags; 0 subscribes to all of them. Take  the flags the caller role is allowed to use from `GET api/2.0/settings/webhook/triggers`, since a flag beyond  that set is refused with 400. A subscription still only fires for events its creator may see. | [optional] 
**TargetId** | Pointer to **NullableString** | The single entity the subscription is narrowed to, by its identifier - a room or a file, for instance.  Leaving it out delivers events about every entity the subscribed triggers cover. | [optional] 

## Methods

### NewCreateWebhooksConfigRequestDto

`func NewCreateWebhooksConfigRequestDto(name string, uri string, ) *CreateWebhooksConfigRequestDto`

NewCreateWebhooksConfigRequestDto instantiates a new CreateWebhooksConfigRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhooksConfigRequestDtoWithDefaults

`func NewCreateWebhooksConfigRequestDtoWithDefaults() *CreateWebhooksConfigRequestDto`

NewCreateWebhooksConfigRequestDtoWithDefaults instantiates a new CreateWebhooksConfigRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhooksConfigRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhooksConfigRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhooksConfigRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *CreateWebhooksConfigRequestDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CreateWebhooksConfigRequestDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CreateWebhooksConfigRequestDto) SetUri(v string)`

SetUri sets Uri field to given value.


### GetSecretKey

`func (o *CreateWebhooksConfigRequestDto) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CreateWebhooksConfigRequestDto) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CreateWebhooksConfigRequestDto) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CreateWebhooksConfigRequestDto) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *CreateWebhooksConfigRequestDto) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *CreateWebhooksConfigRequestDto) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetEnabled

`func (o *CreateWebhooksConfigRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateWebhooksConfigRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateWebhooksConfigRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateWebhooksConfigRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsl

`func (o *CreateWebhooksConfigRequestDto) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *CreateWebhooksConfigRequestDto) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *CreateWebhooksConfigRequestDto) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *CreateWebhooksConfigRequestDto) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTriggers

`func (o *CreateWebhooksConfigRequestDto) GetTriggers() WebhookTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CreateWebhooksConfigRequestDto) GetTriggersOk() (*WebhookTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CreateWebhooksConfigRequestDto) SetTriggers(v WebhookTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CreateWebhooksConfigRequestDto) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargetId

`func (o *CreateWebhooksConfigRequestDto) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CreateWebhooksConfigRequestDto) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CreateWebhooksConfigRequestDto) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CreateWebhooksConfigRequestDto) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *CreateWebhooksConfigRequestDto) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *CreateWebhooksConfigRequestDto) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


