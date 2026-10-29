# UpdateWebhooksConfigRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The label the subscription is listed under. It is for the administrator reading the list and is never sent to  the target; it does not have to be unique. | 
**Uri** | **string** | The address the portal posts the event payload to. It has to be an absolute `http` or `https` address outside  the installation own network, and it is probed before anything is stored: it must answer a HEAD request with  a success code, and a redirect does not count as one. | 
**SecretKey** | Pointer to **string** | The shared secret the payload signature is computed with, so the receiver can tell a genuine call from a  forged one. It has to satisfy the portal password rules published by  `GET api/2.0/settings/security/password`, and it is never echoed back by any operation. On an update an empty  value keeps the secret already stored. | [optional] 
**Enabled** | Pointer to **bool** | Whether the subscription delivers at all. While it is off the matching events are dropped rather than queued,  so nothing from that period arrives once it is switched on again. | [optional] 
**Ssl** | Pointer to **bool** | Whether the target certificate is verified. `true` demands an `https` address with a valid certificate; `false`  delivers without checking the certificate, and is what a plain `http` target needs. | [optional] 
**Triggers** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) | The events the subscription listens for, as a bitmask adding together the `id` values of the entries of `GET  api/2.0/settings/webhook/triggers`. 0 is the catch-all, not none: it stands for every event the  subscription creator's role may receive, not for every event of the portal. | [optional] 
**TargetId** | Pointer to **string** | The single entity the subscription is narrowed to, by its identifier in the API that owns it. For user events  it is the GUID of a user (`id` of `GET api/2.0/people`), for group events the GUID of a group (`id` of `GET  api/2.0/group`), and either must be one the subscription's creator may see; for file, folder, room and form  events it is the entry ID from the Files API (`id` of `GET api/2.0/files/rooms` or of a folder listing), a  number for the portal's own storage and a string for third-party storage. Leaving it out delivers events  about every entity the subscribed triggers cover. | [optional] 
**Id** | **int32** | The subscription to act on, by the `id` that `GET api/2.0/settings/webhook` reports. It travels in the body  rather than in the path, and an id that exists in no portal subscription answers 404. | 

## Methods

### NewUpdateWebhooksConfigRequestDto

`func NewUpdateWebhooksConfigRequestDto(name string, uri string, id int32, ) *UpdateWebhooksConfigRequestDto`

NewUpdateWebhooksConfigRequestDto instantiates a new UpdateWebhooksConfigRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhooksConfigRequestDtoWithDefaults

`func NewUpdateWebhooksConfigRequestDtoWithDefaults() *UpdateWebhooksConfigRequestDto`

NewUpdateWebhooksConfigRequestDtoWithDefaults instantiates a new UpdateWebhooksConfigRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhooksConfigRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhooksConfigRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhooksConfigRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *UpdateWebhooksConfigRequestDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UpdateWebhooksConfigRequestDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UpdateWebhooksConfigRequestDto) SetUri(v string)`

SetUri sets Uri field to given value.


### GetSecretKey

`func (o *UpdateWebhooksConfigRequestDto) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UpdateWebhooksConfigRequestDto) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UpdateWebhooksConfigRequestDto) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *UpdateWebhooksConfigRequestDto) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateWebhooksConfigRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWebhooksConfigRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWebhooksConfigRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWebhooksConfigRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsl

`func (o *UpdateWebhooksConfigRequestDto) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *UpdateWebhooksConfigRequestDto) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *UpdateWebhooksConfigRequestDto) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *UpdateWebhooksConfigRequestDto) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTriggers

`func (o *UpdateWebhooksConfigRequestDto) GetTriggers() WebhookTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *UpdateWebhooksConfigRequestDto) GetTriggersOk() (*WebhookTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *UpdateWebhooksConfigRequestDto) SetTriggers(v WebhookTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *UpdateWebhooksConfigRequestDto) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargetId

`func (o *UpdateWebhooksConfigRequestDto) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *UpdateWebhooksConfigRequestDto) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *UpdateWebhooksConfigRequestDto) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *UpdateWebhooksConfigRequestDto) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetId

`func (o *UpdateWebhooksConfigRequestDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateWebhooksConfigRequestDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateWebhooksConfigRequestDto) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


