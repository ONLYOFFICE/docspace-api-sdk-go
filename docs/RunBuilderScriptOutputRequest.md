# RunBuilderScriptOutputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **NullableInt32** | The id of a portal file to store the result in as a new version, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`. | [optional] 
**FolderId** | Pointer to **NullableInt32** | The id of a folder to save the result in as a new file, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`. | [optional] 
**Title** | Pointer to **NullableString** | The title of the new file, when it has to differ from the name the script saved the result under. | [optional] 

## Methods

### NewRunBuilderScriptOutputRequest

`func NewRunBuilderScriptOutputRequest() *RunBuilderScriptOutputRequest`

NewRunBuilderScriptOutputRequest instantiates a new RunBuilderScriptOutputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBuilderScriptOutputRequestWithDefaults

`func NewRunBuilderScriptOutputRequestWithDefaults() *RunBuilderScriptOutputRequest`

NewRunBuilderScriptOutputRequestWithDefaults instantiates a new RunBuilderScriptOutputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *RunBuilderScriptOutputRequest) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *RunBuilderScriptOutputRequest) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *RunBuilderScriptOutputRequest) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *RunBuilderScriptOutputRequest) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *RunBuilderScriptOutputRequest) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *RunBuilderScriptOutputRequest) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetFolderId

`func (o *RunBuilderScriptOutputRequest) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RunBuilderScriptOutputRequest) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RunBuilderScriptOutputRequest) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RunBuilderScriptOutputRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RunBuilderScriptOutputRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RunBuilderScriptOutputRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetTitle

`func (o *RunBuilderScriptOutputRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RunBuilderScriptOutputRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RunBuilderScriptOutputRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RunBuilderScriptOutputRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RunBuilderScriptOutputRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RunBuilderScriptOutputRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


