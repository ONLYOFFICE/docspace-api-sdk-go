# DeleteBatchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | [optional] 
**FolderIds** | Pointer to [**[]DeleteBatchRequestDtoAllOfFolderIds**](DeleteBatchRequestDtoAllOfFolderIds.md) | The folders to delete, by id, each with everything it contains. A number addresses a folder stored in the  portal itself, a string addresses a folder on a connected third-party account, and both kinds may be sent in  one list. | [optional] 
**FileIds** | Pointer to [**[]DeleteBatchRequestDtoAllOfFileIds**](DeleteBatchRequestDtoAllOfFileIds.md) | The files to delete, by id. A number addresses a file stored in the portal itself, a string addresses a file  on a connected third-party account, and both kinds may be sent in one list. | [optional] 
**DeleteAfter** | Pointer to **bool** | Whether the finished operation is still reported: `false` keeps its final record readable through  `GET api/2.0/files/fileops` until it has been read once, `true` drops the record as soon as the work is done.  It does not postpone the deletion and does not delete anything of its own. | [optional] 
**Immediately** | Pointer to **bool** | Where the deleted items go: `false` moves them to the Trash of the caller, from which they can be restored,  `true` removes them at once and for good. | [optional] 

## Methods

### NewDeleteBatchRequestDto

`func NewDeleteBatchRequestDto() *DeleteBatchRequestDto`

NewDeleteBatchRequestDto instantiates a new DeleteBatchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBatchRequestDtoWithDefaults

`func NewDeleteBatchRequestDtoWithDefaults() *DeleteBatchRequestDto`

NewDeleteBatchRequestDtoWithDefaults instantiates a new DeleteBatchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *DeleteBatchRequestDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *DeleteBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *DeleteBatchRequestDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *DeleteBatchRequestDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.

### GetFolderIds

`func (o *DeleteBatchRequestDto) GetFolderIds() []DeleteBatchRequestDtoAllOfFolderIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *DeleteBatchRequestDto) GetFolderIdsOk() (*[]DeleteBatchRequestDtoAllOfFolderIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *DeleteBatchRequestDto) SetFolderIds(v []DeleteBatchRequestDtoAllOfFolderIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *DeleteBatchRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *DeleteBatchRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *DeleteBatchRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *DeleteBatchRequestDto) GetFileIds() []DeleteBatchRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *DeleteBatchRequestDto) GetFileIdsOk() (*[]DeleteBatchRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *DeleteBatchRequestDto) SetFileIds(v []DeleteBatchRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *DeleteBatchRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *DeleteBatchRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *DeleteBatchRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetDeleteAfter

`func (o *DeleteBatchRequestDto) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *DeleteBatchRequestDto) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *DeleteBatchRequestDto) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *DeleteBatchRequestDto) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.

### GetImmediately

`func (o *DeleteBatchRequestDto) GetImmediately() bool`

GetImmediately returns the Immediately field if non-nil, zero value otherwise.

### GetImmediatelyOk

`func (o *DeleteBatchRequestDto) GetImmediatelyOk() (*bool, bool)`

GetImmediatelyOk returns a tuple with the Immediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediately

`func (o *DeleteBatchRequestDto) SetImmediately(v bool)`

SetImmediately sets Immediately field to given value.

### HasImmediately

`func (o *DeleteBatchRequestDto) HasImmediately() bool`

HasImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


