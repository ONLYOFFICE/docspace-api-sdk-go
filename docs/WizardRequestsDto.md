# WizardRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | The address the portal owner account is created with, which is also the address every administrative letter  goes to afterwards. It has to be a well-formed email address; a malformed one leaves the wizard unfinished. | 
**PasswordHash** | **NullableString** | The owner password, already hashed in the client rather than sent in the clear. Hash it with the `salt`,  iteration count and hash size that `GET api/2.0/settings?withpassword=true` publishes, so the portal can  recognise it later; an empty value leaves the wizard unfinished. | 
**Lng** | Pointer to **NullableString** | The portal interface language, as a culture name such as `en-US`. It has to be one of the cultures enabled  for the installation, and an unknown one leaves the shipped default in place instead of failing the wizard. | [optional] 
**TimeZone** | Pointer to **NullableString** | The time zone every portal date is rendered in, as an IANA identifier such as `Europe/Riga`. A value that  matches nothing falls back to UTC rather than failing the wizard. | [optional] 
**AmiId** | Pointer to **NullableString** | The identifier of the Amazon Machine Image the portal was launched from, for an installation started from an  AWS image. It is recorded for the installation record only and changes nothing about the portal; leave it out  anywhere else. | [optional] 
**SubscribeFromSite** | Pointer to **bool** | Whether the owner agrees to receive product news at the address in `email`. It is a mailing consent and has  no bearing on the portal notifications, which are subscribed separately. | [optional] 

## Methods

### NewWizardRequestsDto

`func NewWizardRequestsDto(email NullableString, passwordHash NullableString, ) *WizardRequestsDto`

NewWizardRequestsDto instantiates a new WizardRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardRequestsDtoWithDefaults

`func NewWizardRequestsDtoWithDefaults() *WizardRequestsDto`

NewWizardRequestsDtoWithDefaults instantiates a new WizardRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *WizardRequestsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WizardRequestsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WizardRequestsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *WizardRequestsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *WizardRequestsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPasswordHash

`func (o *WizardRequestsDto) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *WizardRequestsDto) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *WizardRequestsDto) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.


### SetPasswordHashNil

`func (o *WizardRequestsDto) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *WizardRequestsDto) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLng

`func (o *WizardRequestsDto) GetLng() string`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *WizardRequestsDto) GetLngOk() (*string, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *WizardRequestsDto) SetLng(v string)`

SetLng sets Lng field to given value.

### HasLng

`func (o *WizardRequestsDto) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *WizardRequestsDto) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *WizardRequestsDto) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetTimeZone

`func (o *WizardRequestsDto) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *WizardRequestsDto) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *WizardRequestsDto) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *WizardRequestsDto) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *WizardRequestsDto) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *WizardRequestsDto) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetAmiId

`func (o *WizardRequestsDto) GetAmiId() string`

GetAmiId returns the AmiId field if non-nil, zero value otherwise.

### GetAmiIdOk

`func (o *WizardRequestsDto) GetAmiIdOk() (*string, bool)`

GetAmiIdOk returns a tuple with the AmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiId

`func (o *WizardRequestsDto) SetAmiId(v string)`

SetAmiId sets AmiId field to given value.

### HasAmiId

`func (o *WizardRequestsDto) HasAmiId() bool`

HasAmiId returns a boolean if a field has been set.

### SetAmiIdNil

`func (o *WizardRequestsDto) SetAmiIdNil(b bool)`

 SetAmiIdNil sets the value for AmiId to be an explicit nil

### UnsetAmiId
`func (o *WizardRequestsDto) UnsetAmiId()`

UnsetAmiId ensures that no value is present for AmiId, not even an explicit nil
### GetSubscribeFromSite

`func (o *WizardRequestsDto) GetSubscribeFromSite() bool`

GetSubscribeFromSite returns the SubscribeFromSite field if non-nil, zero value otherwise.

### GetSubscribeFromSiteOk

`func (o *WizardRequestsDto) GetSubscribeFromSiteOk() (*bool, bool)`

GetSubscribeFromSiteOk returns a tuple with the SubscribeFromSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeFromSite

`func (o *WizardRequestsDto) SetSubscribeFromSite(v bool)`

SetSubscribeFromSite sets SubscribeFromSite field to given value.

### HasSubscribeFromSite

`func (o *WizardRequestsDto) HasSubscribeFromSite() bool`

HasSubscribeFromSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


