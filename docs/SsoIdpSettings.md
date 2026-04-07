# SsoIdpSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** | The entity ID. | [optional] 
**SsoUrl** | Pointer to **NullableString** | The SSO URL. | [optional] 
**SsoBinding** | Pointer to **NullableString** | The SSO binding. | [optional] 
**SloUrl** | Pointer to **NullableString** | The SLO URL. | [optional] 
**SloBinding** | Pointer to **NullableString** | The SLO binding. | [optional] 
**NameIdFormat** | Pointer to **NullableString** | The name ID format. | [optional] 

## Methods

### NewSsoIdpSettings

`func NewSsoIdpSettings() *SsoIdpSettings`

NewSsoIdpSettings instantiates a new SsoIdpSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoIdpSettingsWithDefaults

`func NewSsoIdpSettingsWithDefaults() *SsoIdpSettings`

NewSsoIdpSettingsWithDefaults instantiates a new SsoIdpSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SsoIdpSettings) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SsoIdpSettings) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SsoIdpSettings) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SsoIdpSettings) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *SsoIdpSettings) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *SsoIdpSettings) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetSsoUrl

`func (o *SsoIdpSettings) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SsoIdpSettings) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SsoIdpSettings) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *SsoIdpSettings) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### SetSsoUrlNil

`func (o *SsoIdpSettings) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *SsoIdpSettings) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
### GetSsoBinding

`func (o *SsoIdpSettings) GetSsoBinding() string`

GetSsoBinding returns the SsoBinding field if non-nil, zero value otherwise.

### GetSsoBindingOk

`func (o *SsoIdpSettings) GetSsoBindingOk() (*string, bool)`

GetSsoBindingOk returns a tuple with the SsoBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBinding

`func (o *SsoIdpSettings) SetSsoBinding(v string)`

SetSsoBinding sets SsoBinding field to given value.

### HasSsoBinding

`func (o *SsoIdpSettings) HasSsoBinding() bool`

HasSsoBinding returns a boolean if a field has been set.

### SetSsoBindingNil

`func (o *SsoIdpSettings) SetSsoBindingNil(b bool)`

 SetSsoBindingNil sets the value for SsoBinding to be an explicit nil

### UnsetSsoBinding
`func (o *SsoIdpSettings) UnsetSsoBinding()`

UnsetSsoBinding ensures that no value is present for SsoBinding, not even an explicit nil
### GetSloUrl

`func (o *SsoIdpSettings) GetSloUrl() string`

GetSloUrl returns the SloUrl field if non-nil, zero value otherwise.

### GetSloUrlOk

`func (o *SsoIdpSettings) GetSloUrlOk() (*string, bool)`

GetSloUrlOk returns a tuple with the SloUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloUrl

`func (o *SsoIdpSettings) SetSloUrl(v string)`

SetSloUrl sets SloUrl field to given value.

### HasSloUrl

`func (o *SsoIdpSettings) HasSloUrl() bool`

HasSloUrl returns a boolean if a field has been set.

### SetSloUrlNil

`func (o *SsoIdpSettings) SetSloUrlNil(b bool)`

 SetSloUrlNil sets the value for SloUrl to be an explicit nil

### UnsetSloUrl
`func (o *SsoIdpSettings) UnsetSloUrl()`

UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
### GetSloBinding

`func (o *SsoIdpSettings) GetSloBinding() string`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *SsoIdpSettings) GetSloBindingOk() (*string, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *SsoIdpSettings) SetSloBinding(v string)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *SsoIdpSettings) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### SetSloBindingNil

`func (o *SsoIdpSettings) SetSloBindingNil(b bool)`

 SetSloBindingNil sets the value for SloBinding to be an explicit nil

### UnsetSloBinding
`func (o *SsoIdpSettings) UnsetSloBinding()`

UnsetSloBinding ensures that no value is present for SloBinding, not even an explicit nil
### GetNameIdFormat

`func (o *SsoIdpSettings) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *SsoIdpSettings) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *SsoIdpSettings) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *SsoIdpSettings) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### SetNameIdFormatNil

`func (o *SsoIdpSettings) SetNameIdFormatNil(b bool)`

 SetNameIdFormatNil sets the value for NameIdFormat to be an explicit nil

### UnsetNameIdFormat
`func (o *SsoIdpSettings) UnsetNameIdFormat()`

UnsetNameIdFormat ensures that no value is present for NameIdFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


