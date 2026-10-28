# WebhooksConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier of the subscription, which is what `PUT api/2.0/settings/webhook`,  `DELETE api/2.0/settings/webhook/{id}` and the `configId` filter of the delivery log address it by. | 
**Name** | Pointer to **NullableString** | The label the subscription was given, free text with no meaning to the portal. | [optional] 
**Uri** | Pointer to **NullableString** | The address every delivery is posted to. The signing secret that lets the receiver verify a delivery is  never part of this answer, so it has to be kept from the moment the subscription was created. | [optional] 
**Enabled** | Pointer to **bool** | Whether the subscription is delivering. While it is `false` events are dropped rather than queued, so  nothing arrives late after it is switched back on. | [optional] 
**Ssl** | Pointer to **bool** | Whether the certificate of `uri` is verified before a delivery. While it is `false` a self-signed  certificate is accepted as well. | [optional] 
**Triggers** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) | The events the subscription covers, as the bits of `GET api/2.0/settings/webhook/triggers` added  together. `0` is the catch-all and means every event, not none. | [optional] 
**TargetId** | Pointer to **NullableString** | The single room or file the subscription is narrowed to, empty for a subscription that covers the whole  portal. It is kept as an opaque value, so both a numeric and a third-party identifier can appear. | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The member who created the subscription, which is also who a non-administrator is limited to seeing. It is  empty for a subscription created by a portal background job. | [optional] 
**CreatedOn** | Pointer to **NullableTime** | When the subscription was created, in the portal time zone. | [optional] 
**ModifiedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The member who last changed the subscription, empty while nobody has changed it since it was created. | [optional] 
**ModifiedOn** | Pointer to **NullableTime** | When it was last changed, in the portal time zone, and empty under the same condition as `modifiedBy`. | [optional] 
**LastFailureOn** | Pointer to **NullableTime** | When a delivery last failed, in the portal time zone. It is empty for a subscription that has never  failed, and it is not cleared by a later success - compare it with `lastSuccessOn` to see which came last. | [optional] 
**LastFailureContent** | Pointer to **NullableString** | What the target answered on that failure, truncated, for diagnosing without opening the delivery log. It  is empty when the failure produced no body at all, a timeout for instance. | [optional] 
**LastSuccessOn** | Pointer to **NullableTime** | When a delivery last succeeded, in the portal time zone, empty for a subscription that has never  delivered. Both this and `lastFailureOn` being empty means nothing has been attempted yet. | [optional] 

## Methods

### NewWebhooksConfigDto

`func NewWebhooksConfigDto(id int32, ) *WebhooksConfigDto`

NewWebhooksConfigDto instantiates a new WebhooksConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksConfigDtoWithDefaults

`func NewWebhooksConfigDtoWithDefaults() *WebhooksConfigDto`

NewWebhooksConfigDtoWithDefaults instantiates a new WebhooksConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhooksConfigDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhooksConfigDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhooksConfigDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *WebhooksConfigDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksConfigDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksConfigDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhooksConfigDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhooksConfigDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhooksConfigDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUri

`func (o *WebhooksConfigDto) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *WebhooksConfigDto) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *WebhooksConfigDto) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *WebhooksConfigDto) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *WebhooksConfigDto) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *WebhooksConfigDto) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetEnabled

`func (o *WebhooksConfigDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhooksConfigDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhooksConfigDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhooksConfigDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSsl

`func (o *WebhooksConfigDto) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *WebhooksConfigDto) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *WebhooksConfigDto) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *WebhooksConfigDto) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTriggers

`func (o *WebhooksConfigDto) GetTriggers() WebhookTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *WebhooksConfigDto) GetTriggersOk() (*WebhookTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *WebhooksConfigDto) SetTriggers(v WebhookTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *WebhooksConfigDto) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTargetId

`func (o *WebhooksConfigDto) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *WebhooksConfigDto) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *WebhooksConfigDto) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *WebhooksConfigDto) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *WebhooksConfigDto) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *WebhooksConfigDto) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetCreatedBy

`func (o *WebhooksConfigDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhooksConfigDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhooksConfigDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WebhooksConfigDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *WebhooksConfigDto) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WebhooksConfigDto) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WebhooksConfigDto) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *WebhooksConfigDto) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *WebhooksConfigDto) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *WebhooksConfigDto) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetModifiedBy

`func (o *WebhooksConfigDto) GetModifiedBy() EmployeeDto`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhooksConfigDto) GetModifiedByOk() (*EmployeeDto, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *WebhooksConfigDto) SetModifiedBy(v EmployeeDto)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *WebhooksConfigDto) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *WebhooksConfigDto) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *WebhooksConfigDto) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *WebhooksConfigDto) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *WebhooksConfigDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *WebhooksConfigDto) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *WebhooksConfigDto) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetLastFailureOn

`func (o *WebhooksConfigDto) GetLastFailureOn() time.Time`

GetLastFailureOn returns the LastFailureOn field if non-nil, zero value otherwise.

### GetLastFailureOnOk

`func (o *WebhooksConfigDto) GetLastFailureOnOk() (*time.Time, bool)`

GetLastFailureOnOk returns a tuple with the LastFailureOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureOn

`func (o *WebhooksConfigDto) SetLastFailureOn(v time.Time)`

SetLastFailureOn sets LastFailureOn field to given value.

### HasLastFailureOn

`func (o *WebhooksConfigDto) HasLastFailureOn() bool`

HasLastFailureOn returns a boolean if a field has been set.

### SetLastFailureOnNil

`func (o *WebhooksConfigDto) SetLastFailureOnNil(b bool)`

 SetLastFailureOnNil sets the value for LastFailureOn to be an explicit nil

### UnsetLastFailureOn
`func (o *WebhooksConfigDto) UnsetLastFailureOn()`

UnsetLastFailureOn ensures that no value is present for LastFailureOn, not even an explicit nil
### GetLastFailureContent

`func (o *WebhooksConfigDto) GetLastFailureContent() string`

GetLastFailureContent returns the LastFailureContent field if non-nil, zero value otherwise.

### GetLastFailureContentOk

`func (o *WebhooksConfigDto) GetLastFailureContentOk() (*string, bool)`

GetLastFailureContentOk returns a tuple with the LastFailureContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureContent

`func (o *WebhooksConfigDto) SetLastFailureContent(v string)`

SetLastFailureContent sets LastFailureContent field to given value.

### HasLastFailureContent

`func (o *WebhooksConfigDto) HasLastFailureContent() bool`

HasLastFailureContent returns a boolean if a field has been set.

### SetLastFailureContentNil

`func (o *WebhooksConfigDto) SetLastFailureContentNil(b bool)`

 SetLastFailureContentNil sets the value for LastFailureContent to be an explicit nil

### UnsetLastFailureContent
`func (o *WebhooksConfigDto) UnsetLastFailureContent()`

UnsetLastFailureContent ensures that no value is present for LastFailureContent, not even an explicit nil
### GetLastSuccessOn

`func (o *WebhooksConfigDto) GetLastSuccessOn() time.Time`

GetLastSuccessOn returns the LastSuccessOn field if non-nil, zero value otherwise.

### GetLastSuccessOnOk

`func (o *WebhooksConfigDto) GetLastSuccessOnOk() (*time.Time, bool)`

GetLastSuccessOnOk returns a tuple with the LastSuccessOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessOn

`func (o *WebhooksConfigDto) SetLastSuccessOn(v time.Time)`

SetLastSuccessOn sets LastSuccessOn field to given value.

### HasLastSuccessOn

`func (o *WebhooksConfigDto) HasLastSuccessOn() bool`

HasLastSuccessOn returns a boolean if a field has been set.

### SetLastSuccessOnNil

`func (o *WebhooksConfigDto) SetLastSuccessOnNil(b bool)`

 SetLastSuccessOnNil sets the value for LastSuccessOn to be an explicit nil

### UnsetLastSuccessOn
`func (o *WebhooksConfigDto) UnsetLastSuccessOn()`

UnsetLastSuccessOn ensures that no value is present for LastSuccessOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


