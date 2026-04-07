# BatchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Specifies whether to return only the current operation | [optional] 
**FolderIds** | Pointer to [**[]BatchRequestDtoAllOfFolderIds**](BatchRequestDtoAllOfFolderIds.md) | The list of folder IDs to be copied/moved. | [optional] 
**FileIds** | Pointer to [**[]BatchRequestDtoAllOfFileIds**](BatchRequestDtoAllOfFileIds.md) | The list of file IDs to be copied/moved. | [optional] 
**DestFolderId** | Pointer to [**BatchRequestDtoAllOfDestFolderId**](BatchRequestDtoAllOfDestFolderId.md) |  | [optional] 
**ConflictResolveType** | Pointer to [**FileConflictResolveType**](FileConflictResolveType.md) |  | [optional] 
**DeleteAfter** | Pointer to **bool** | Specifies whether to delete the source files/folders after they are moved or copied to the destination folder. | [optional] 
**Content** | Pointer to **bool** | Specifies whether to copy or move the folder content or not. | [optional] 
**ToFillOut** | Pointer to **bool** | Specifies whether the file is copied for filling out | [optional] 

## Methods

### NewBatchRequestDto

`func NewBatchRequestDto() *BatchRequestDto`

NewBatchRequestDto instantiates a new BatchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestDtoWithDefaults

`func NewBatchRequestDtoWithDefaults() *BatchRequestDto`

NewBatchRequestDtoWithDefaults instantiates a new BatchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *BatchRequestDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *BatchRequestDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *BatchRequestDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *BatchRequestDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.

### GetFolderIds

`func (o *BatchRequestDto) GetFolderIds() []BatchRequestDtoAllOfFolderIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *BatchRequestDto) GetFolderIdsOk() (*[]BatchRequestDtoAllOfFolderIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *BatchRequestDto) SetFolderIds(v []BatchRequestDtoAllOfFolderIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *BatchRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *BatchRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *BatchRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *BatchRequestDto) GetFileIds() []BatchRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *BatchRequestDto) GetFileIdsOk() (*[]BatchRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *BatchRequestDto) SetFileIds(v []BatchRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *BatchRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *BatchRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *BatchRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetDestFolderId

`func (o *BatchRequestDto) GetDestFolderId() BatchRequestDtoAllOfDestFolderId`

GetDestFolderId returns the DestFolderId field if non-nil, zero value otherwise.

### GetDestFolderIdOk

`func (o *BatchRequestDto) GetDestFolderIdOk() (*BatchRequestDtoAllOfDestFolderId, bool)`

GetDestFolderIdOk returns a tuple with the DestFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestFolderId

`func (o *BatchRequestDto) SetDestFolderId(v BatchRequestDtoAllOfDestFolderId)`

SetDestFolderId sets DestFolderId field to given value.

### HasDestFolderId

`func (o *BatchRequestDto) HasDestFolderId() bool`

HasDestFolderId returns a boolean if a field has been set.

### GetConflictResolveType

`func (o *BatchRequestDto) GetConflictResolveType() FileConflictResolveType`

GetConflictResolveType returns the ConflictResolveType field if non-nil, zero value otherwise.

### GetConflictResolveTypeOk

`func (o *BatchRequestDto) GetConflictResolveTypeOk() (*FileConflictResolveType, bool)`

GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolveType

`func (o *BatchRequestDto) SetConflictResolveType(v FileConflictResolveType)`

SetConflictResolveType sets ConflictResolveType field to given value.

### HasConflictResolveType

`func (o *BatchRequestDto) HasConflictResolveType() bool`

HasConflictResolveType returns a boolean if a field has been set.

### GetDeleteAfter

`func (o *BatchRequestDto) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *BatchRequestDto) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *BatchRequestDto) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *BatchRequestDto) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.

### GetContent

`func (o *BatchRequestDto) GetContent() bool`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BatchRequestDto) GetContentOk() (*bool, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BatchRequestDto) SetContent(v bool)`

SetContent sets Content field to given value.

### HasContent

`func (o *BatchRequestDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetToFillOut

`func (o *BatchRequestDto) GetToFillOut() bool`

GetToFillOut returns the ToFillOut field if non-nil, zero value otherwise.

### GetToFillOutOk

`func (o *BatchRequestDto) GetToFillOutOk() (*bool, bool)`

GetToFillOutOk returns a tuple with the ToFillOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFillOut

`func (o *BatchRequestDto) SetToFillOut(v bool)`

SetToFillOut sets ToFillOut field to given value.

### HasToFillOut

`func (o *BatchRequestDto) HasToFillOut() bool`

HasToFillOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


