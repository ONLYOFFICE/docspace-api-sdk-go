# AiFileOperationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The file operation ID. | 
**Operation** | [**AiFileOperationType**](AiFileOperationType.md) | The file operation type. | 
**Progress** | **int32** | The file operation progress in percentage. | 
**Error** | **NullableString** | The file operation error message. | 
**Processed** | **NullableString** | The file operation processing status. | 
**Finished** | **bool** | Specifies if the file operation is finished or not. | 
**Url** | Pointer to **NullableString** | The file operation URL. | [optional] 
**Files** | Pointer to [**[]AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The list of files of the file operation. | [optional] 
**Folders** | Pointer to [**[]AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The list of folders of the file operation. | [optional] 
**Status** | Pointer to [**AiDistributedTaskStatus**](AiDistributedTaskStatus.md) | The status of the distributed task related to the file operation. | [optional] 

## Methods

### NewAiFileOperationDto

`func NewAiFileOperationDto(id NullableString, operation AiFileOperationType, progress int32, error_ NullableString, processed NullableString, finished bool, ) *AiFileOperationDto`

NewAiFileOperationDto instantiates a new AiFileOperationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFileOperationDtoWithDefaults

`func NewAiFileOperationDtoWithDefaults() *AiFileOperationDto`

NewAiFileOperationDtoWithDefaults instantiates a new AiFileOperationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiFileOperationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiFileOperationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiFileOperationDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiFileOperationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiFileOperationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOperation

`func (o *AiFileOperationDto) GetOperation() AiFileOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AiFileOperationDto) GetOperationOk() (*AiFileOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AiFileOperationDto) SetOperation(v AiFileOperationType)`

SetOperation sets Operation field to given value.


### GetProgress

`func (o *AiFileOperationDto) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AiFileOperationDto) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AiFileOperationDto) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetError

`func (o *AiFileOperationDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiFileOperationDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiFileOperationDto) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *AiFileOperationDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AiFileOperationDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProcessed

`func (o *AiFileOperationDto) GetProcessed() string`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *AiFileOperationDto) GetProcessedOk() (*string, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *AiFileOperationDto) SetProcessed(v string)`

SetProcessed sets Processed field to given value.


### SetProcessedNil

`func (o *AiFileOperationDto) SetProcessedNil(b bool)`

 SetProcessedNil sets the value for Processed to be an explicit nil

### UnsetProcessed
`func (o *AiFileOperationDto) UnsetProcessed()`

UnsetProcessed ensures that no value is present for Processed, not even an explicit nil
### GetFinished

`func (o *AiFileOperationDto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *AiFileOperationDto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *AiFileOperationDto) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetUrl

`func (o *AiFileOperationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiFileOperationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiFileOperationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiFileOperationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AiFileOperationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AiFileOperationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFiles

`func (o *AiFileOperationDto) GetFiles() []AiFileEntryBaseDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AiFileOperationDto) GetFilesOk() (*[]AiFileEntryBaseDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AiFileOperationDto) SetFiles(v []AiFileEntryBaseDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AiFileOperationDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *AiFileOperationDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *AiFileOperationDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFolders

`func (o *AiFileOperationDto) GetFolders() []AiFileEntryBaseDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *AiFileOperationDto) GetFoldersOk() (*[]AiFileEntryBaseDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *AiFileOperationDto) SetFolders(v []AiFileEntryBaseDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *AiFileOperationDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *AiFileOperationDto) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *AiFileOperationDto) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetStatus

`func (o *AiFileOperationDto) GetStatus() AiDistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiFileOperationDto) GetStatusOk() (*AiDistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiFileOperationDto) SetStatus(v AiDistributedTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiFileOperationDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


