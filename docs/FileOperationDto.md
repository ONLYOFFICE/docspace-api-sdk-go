# FileOperationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The file operation ID. | 
**Operation** | [**FileOperationType**](FileOperationType.md) | The file operation type. | 
**Progress** | **int32** | The file operation progress in percentage. | 
**Error** | **NullableString** | The file operation error message. | 
**Processed** | **NullableString** | The file operation processing status. | 
**Finished** | **bool** | Specifies if the file operation is finished or not. | 
**Url** | Pointer to **NullableString** | The file operation URL. | [optional] 
**Files** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of files of the file operation. | [optional] 
**Folders** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of folders of the file operation. | [optional] 
**Status** | Pointer to [**DistributedTaskStatus**](DistributedTaskStatus.md) | The status of the distributed task related to the file operation. | [optional] 

## Methods

### NewFileOperationDto

`func NewFileOperationDto(id NullableString, operation FileOperationType, progress int32, error_ NullableString, processed NullableString, finished bool, ) *FileOperationDto`

NewFileOperationDto instantiates a new FileOperationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOperationDtoWithDefaults

`func NewFileOperationDtoWithDefaults() *FileOperationDto`

NewFileOperationDtoWithDefaults instantiates a new FileOperationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileOperationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileOperationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileOperationDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FileOperationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FileOperationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOperation

`func (o *FileOperationDto) GetOperation() FileOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *FileOperationDto) GetOperationOk() (*FileOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *FileOperationDto) SetOperation(v FileOperationType)`

SetOperation sets Operation field to given value.


### GetProgress

`func (o *FileOperationDto) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FileOperationDto) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FileOperationDto) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetError

`func (o *FileOperationDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FileOperationDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FileOperationDto) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *FileOperationDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *FileOperationDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProcessed

`func (o *FileOperationDto) GetProcessed() string`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *FileOperationDto) GetProcessedOk() (*string, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *FileOperationDto) SetProcessed(v string)`

SetProcessed sets Processed field to given value.


### SetProcessedNil

`func (o *FileOperationDto) SetProcessedNil(b bool)`

 SetProcessedNil sets the value for Processed to be an explicit nil

### UnsetProcessed
`func (o *FileOperationDto) UnsetProcessed()`

UnsetProcessed ensures that no value is present for Processed, not even an explicit nil
### GetFinished

`func (o *FileOperationDto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *FileOperationDto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *FileOperationDto) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetUrl

`func (o *FileOperationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileOperationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileOperationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileOperationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FileOperationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FileOperationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFiles

`func (o *FileOperationDto) GetFiles() []FileEntryBaseDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileOperationDto) GetFilesOk() (*[]FileEntryBaseDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileOperationDto) SetFiles(v []FileEntryBaseDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileOperationDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FileOperationDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FileOperationDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFolders

`func (o *FileOperationDto) GetFolders() []FileEntryBaseDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *FileOperationDto) GetFoldersOk() (*[]FileEntryBaseDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *FileOperationDto) SetFolders(v []FileEntryBaseDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *FileOperationDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *FileOperationDto) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *FileOperationDto) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetStatus

`func (o *FileOperationDto) GetStatus() DistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileOperationDto) GetStatusOk() (*DistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileOperationDto) SetStatus(v DistributedTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileOperationDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


