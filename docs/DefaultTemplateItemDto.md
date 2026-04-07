# DefaultTemplateItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedFile** | Pointer to **NullableInt32** | File id to use as a default template | [optional] 
**FileExtension** | **NullableString** | Extension of a default template | 
**FileTitle** | Pointer to **NullableString** | Title of a default template | [optional] 
**LastModified** | Pointer to **NullableTime** | Last modified date of a default template | [optional] 
**FileSize** | Pointer to **NullableInt64** | Filesize (in bytes) of a default template | [optional] 
**ViewUrl** | Pointer to **NullableString** | View url of a default template | [optional] 

## Methods

### NewDefaultTemplateItemDto

`func NewDefaultTemplateItemDto(fileExtension NullableString, ) *DefaultTemplateItemDto`

NewDefaultTemplateItemDto instantiates a new DefaultTemplateItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultTemplateItemDtoWithDefaults

`func NewDefaultTemplateItemDtoWithDefaults() *DefaultTemplateItemDto`

NewDefaultTemplateItemDtoWithDefaults instantiates a new DefaultTemplateItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedFile

`func (o *DefaultTemplateItemDto) GetSelectedFile() int32`

GetSelectedFile returns the SelectedFile field if non-nil, zero value otherwise.

### GetSelectedFileOk

`func (o *DefaultTemplateItemDto) GetSelectedFileOk() (*int32, bool)`

GetSelectedFileOk returns a tuple with the SelectedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedFile

`func (o *DefaultTemplateItemDto) SetSelectedFile(v int32)`

SetSelectedFile sets SelectedFile field to given value.

### HasSelectedFile

`func (o *DefaultTemplateItemDto) HasSelectedFile() bool`

HasSelectedFile returns a boolean if a field has been set.

### SetSelectedFileNil

`func (o *DefaultTemplateItemDto) SetSelectedFileNil(b bool)`

 SetSelectedFileNil sets the value for SelectedFile to be an explicit nil

### UnsetSelectedFile
`func (o *DefaultTemplateItemDto) UnsetSelectedFile()`

UnsetSelectedFile ensures that no value is present for SelectedFile, not even an explicit nil
### GetFileExtension

`func (o *DefaultTemplateItemDto) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *DefaultTemplateItemDto) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *DefaultTemplateItemDto) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.


### SetFileExtensionNil

`func (o *DefaultTemplateItemDto) SetFileExtensionNil(b bool)`

 SetFileExtensionNil sets the value for FileExtension to be an explicit nil

### UnsetFileExtension
`func (o *DefaultTemplateItemDto) UnsetFileExtension()`

UnsetFileExtension ensures that no value is present for FileExtension, not even an explicit nil
### GetFileTitle

`func (o *DefaultTemplateItemDto) GetFileTitle() string`

GetFileTitle returns the FileTitle field if non-nil, zero value otherwise.

### GetFileTitleOk

`func (o *DefaultTemplateItemDto) GetFileTitleOk() (*string, bool)`

GetFileTitleOk returns a tuple with the FileTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTitle

`func (o *DefaultTemplateItemDto) SetFileTitle(v string)`

SetFileTitle sets FileTitle field to given value.

### HasFileTitle

`func (o *DefaultTemplateItemDto) HasFileTitle() bool`

HasFileTitle returns a boolean if a field has been set.

### SetFileTitleNil

`func (o *DefaultTemplateItemDto) SetFileTitleNil(b bool)`

 SetFileTitleNil sets the value for FileTitle to be an explicit nil

### UnsetFileTitle
`func (o *DefaultTemplateItemDto) UnsetFileTitle()`

UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil
### GetLastModified

`func (o *DefaultTemplateItemDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DefaultTemplateItemDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DefaultTemplateItemDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DefaultTemplateItemDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *DefaultTemplateItemDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *DefaultTemplateItemDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetFileSize

`func (o *DefaultTemplateItemDto) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DefaultTemplateItemDto) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DefaultTemplateItemDto) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DefaultTemplateItemDto) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *DefaultTemplateItemDto) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *DefaultTemplateItemDto) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetViewUrl

`func (o *DefaultTemplateItemDto) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *DefaultTemplateItemDto) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *DefaultTemplateItemDto) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *DefaultTemplateItemDto) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### SetViewUrlNil

`func (o *DefaultTemplateItemDto) SetViewUrlNil(b bool)`

 SetViewUrlNil sets the value for ViewUrl to be an explicit nil

### UnsetViewUrl
`func (o *DefaultTemplateItemDto) UnsetViewUrl()`

UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


