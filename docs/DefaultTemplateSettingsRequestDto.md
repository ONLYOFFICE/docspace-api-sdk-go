# DefaultTemplateSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedFile** | [**DefaultTemplateSettingsRequestDtoSelectedFile**](DefaultTemplateSettingsRequestDtoSelectedFile.md) |  | 
**FileExtension** | **NullableString** | File extension of a template to replace | 

## Methods

### NewDefaultTemplateSettingsRequestDto

`func NewDefaultTemplateSettingsRequestDto(selectedFile DefaultTemplateSettingsRequestDtoSelectedFile, fileExtension NullableString, ) *DefaultTemplateSettingsRequestDto`

NewDefaultTemplateSettingsRequestDto instantiates a new DefaultTemplateSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultTemplateSettingsRequestDtoWithDefaults

`func NewDefaultTemplateSettingsRequestDtoWithDefaults() *DefaultTemplateSettingsRequestDto`

NewDefaultTemplateSettingsRequestDtoWithDefaults instantiates a new DefaultTemplateSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedFile

`func (o *DefaultTemplateSettingsRequestDto) GetSelectedFile() DefaultTemplateSettingsRequestDtoSelectedFile`

GetSelectedFile returns the SelectedFile field if non-nil, zero value otherwise.

### GetSelectedFileOk

`func (o *DefaultTemplateSettingsRequestDto) GetSelectedFileOk() (*DefaultTemplateSettingsRequestDtoSelectedFile, bool)`

GetSelectedFileOk returns a tuple with the SelectedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedFile

`func (o *DefaultTemplateSettingsRequestDto) SetSelectedFile(v DefaultTemplateSettingsRequestDtoSelectedFile)`

SetSelectedFile sets SelectedFile field to given value.


### GetFileExtension

`func (o *DefaultTemplateSettingsRequestDto) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *DefaultTemplateSettingsRequestDto) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *DefaultTemplateSettingsRequestDto) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.


### SetFileExtensionNil

`func (o *DefaultTemplateSettingsRequestDto) SetFileExtensionNil(b bool)`

 SetFileExtensionNil sets the value for FileExtension to be an explicit nil

### UnsetFileExtension
`func (o *DefaultTemplateSettingsRequestDto) UnsetFileExtension()`

UnsetFileExtension ensures that no value is present for FileExtension, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


