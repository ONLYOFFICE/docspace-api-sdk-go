# DefaultTemplateSettingsResetRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileExtension** | **NullableString** | The extension whose custom blank is dropped, written in lower case with the leading dot. Only the extensions  the portal's built-in template set covers are accepted, and `GET api/2.0/files/settings/defaulttemplate`  returns exactly that list; an extension outside it leaves the settings unchanged instead of failing. | 

## Methods

### NewDefaultTemplateSettingsResetRequestDto

`func NewDefaultTemplateSettingsResetRequestDto(fileExtension NullableString, ) *DefaultTemplateSettingsResetRequestDto`

NewDefaultTemplateSettingsResetRequestDto instantiates a new DefaultTemplateSettingsResetRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultTemplateSettingsResetRequestDtoWithDefaults

`func NewDefaultTemplateSettingsResetRequestDtoWithDefaults() *DefaultTemplateSettingsResetRequestDto`

NewDefaultTemplateSettingsResetRequestDtoWithDefaults instantiates a new DefaultTemplateSettingsResetRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileExtension

`func (o *DefaultTemplateSettingsResetRequestDto) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *DefaultTemplateSettingsResetRequestDto) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *DefaultTemplateSettingsResetRequestDto) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.


### SetFileExtensionNil

`func (o *DefaultTemplateSettingsResetRequestDto) SetFileExtensionNil(b bool)`

 SetFileExtensionNil sets the value for FileExtension to be an explicit nil

### UnsetFileExtension
`func (o *DefaultTemplateSettingsResetRequestDto) UnsetFileExtension()`

UnsetFileExtension ensures that no value is present for FileExtension, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


