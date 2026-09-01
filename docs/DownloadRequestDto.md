# DownloadRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Specifies whether to return only the current operation | [optional] 
**FolderIds** | Pointer to [**[]DownloadRequestDtoAllOfFolderIds**](DownloadRequestDtoAllOfFolderIds.md) | The list of folder IDs to be downloaded. | [optional] 
**FileIds** | Pointer to [**[]DownloadRequestDtoAllOfFileIds**](DownloadRequestDtoAllOfFileIds.md) | The list of file IDs to be downloaded. | [optional] 
**FileConvertIds** | Pointer to [**[]DownloadRequestItemDto**](DownloadRequestItemDto.md) | The list of file IDs which will be converted. | [optional] 

## Methods

### NewDownloadRequestDto

`func NewDownloadRequestDto() *DownloadRequestDto`

NewDownloadRequestDto instantiates a new DownloadRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadRequestDtoWithDefaults

`func NewDownloadRequestDtoWithDefaults() *DownloadRequestDto`

NewDownloadRequestDtoWithDefaults instantiates a new DownloadRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *DownloadRequestDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *DownloadRequestDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *DownloadRequestDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *DownloadRequestDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.

### GetFolderIds

`func (o *DownloadRequestDto) GetFolderIds() []DownloadRequestDtoAllOfFolderIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *DownloadRequestDto) GetFolderIdsOk() (*[]DownloadRequestDtoAllOfFolderIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *DownloadRequestDto) SetFolderIds(v []DownloadRequestDtoAllOfFolderIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *DownloadRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *DownloadRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *DownloadRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *DownloadRequestDto) GetFileIds() []DownloadRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *DownloadRequestDto) GetFileIdsOk() (*[]DownloadRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *DownloadRequestDto) SetFileIds(v []DownloadRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *DownloadRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *DownloadRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *DownloadRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetFileConvertIds

`func (o *DownloadRequestDto) GetFileConvertIds() []DownloadRequestItemDto`

GetFileConvertIds returns the FileConvertIds field if non-nil, zero value otherwise.

### GetFileConvertIdsOk

`func (o *DownloadRequestDto) GetFileConvertIdsOk() (*[]DownloadRequestItemDto, bool)`

GetFileConvertIdsOk returns a tuple with the FileConvertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileConvertIds

`func (o *DownloadRequestDto) SetFileConvertIds(v []DownloadRequestItemDto)`

SetFileConvertIds sets FileConvertIds field to given value.

### HasFileConvertIds

`func (o *DownloadRequestDto) HasFileConvertIds() bool`

HasFileConvertIds returns a boolean if a field has been set.

### SetFileConvertIdsNil

`func (o *DownloadRequestDto) SetFileConvertIdsNil(b bool)`

 SetFileConvertIdsNil sets the value for FileConvertIds to be an explicit nil

### UnsetFileConvertIds
`func (o *DownloadRequestDto) UnsetFileConvertIds()`

UnsetFileConvertIds ensures that no value is present for FileConvertIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


