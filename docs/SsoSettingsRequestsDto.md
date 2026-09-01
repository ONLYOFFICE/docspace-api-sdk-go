# SsoSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerializeSettings** | **NullableString** | The JSON-serialized SSO configuration settings. | 

## Methods

### NewSsoSettingsRequestsDto

`func NewSsoSettingsRequestsDto(serializeSettings NullableString, ) *SsoSettingsRequestsDto`

NewSsoSettingsRequestsDto instantiates a new SsoSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsRequestsDtoWithDefaults

`func NewSsoSettingsRequestsDtoWithDefaults() *SsoSettingsRequestsDto`

NewSsoSettingsRequestsDtoWithDefaults instantiates a new SsoSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerializeSettings

`func (o *SsoSettingsRequestsDto) GetSerializeSettings() string`

GetSerializeSettings returns the SerializeSettings field if non-nil, zero value otherwise.

### GetSerializeSettingsOk

`func (o *SsoSettingsRequestsDto) GetSerializeSettingsOk() (*string, bool)`

GetSerializeSettingsOk returns a tuple with the SerializeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializeSettings

`func (o *SsoSettingsRequestsDto) SetSerializeSettings(v string)`

SetSerializeSettings sets SerializeSettings field to given value.


### SetSerializeSettingsNil

`func (o *SsoSettingsRequestsDto) SetSerializeSettingsNil(b bool)`

 SetSerializeSettingsNil sets the value for SerializeSettings to be an explicit nil

### UnsetSerializeSettings
`func (o *SsoSettingsRequestsDto) UnsetSerializeSettings()`

UnsetSerializeSettings ensures that no value is present for SerializeSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


