# BaseBatchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | [optional] 
**FolderIds** | Pointer to [**[]BaseBatchRequestDtoAllOfFolderIds**](BaseBatchRequestDtoAllOfFolderIds.md) | The folders to act on, by id, as reported by a folder listing such as `GET api/2.0/files/{folderId}`. A number  addresses a folder stored in the portal itself, a string addresses a folder on a connected third-party  account, and both kinds may be sent in one list. | [optional] 
**FileIds** | Pointer to [**[]BaseBatchRequestDtoAllOfFileIds**](BaseBatchRequestDtoAllOfFileIds.md) | The files to act on, by id, as reported by a folder listing such as `GET api/2.0/files/{folderId}`. A number  addresses a file stored in the portal itself, a string addresses a file on a connected third-party account,  and both kinds may be sent in one list. | [optional] 

## Methods

### NewBaseBatchRequestDto

`func NewBaseBatchRequestDto() *BaseBatchRequestDto`

NewBaseBatchRequestDto instantiates a new BaseBatchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseBatchRequestDtoWithDefaults

`func NewBaseBatchRequestDtoWithDefaults() *BaseBatchRequestDto`

NewBaseBatchRequestDtoWithDefaults instantiates a new BaseBatchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *BaseBatchRequestDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *BaseBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *BaseBatchRequestDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *BaseBatchRequestDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.

### GetFolderIds

`func (o *BaseBatchRequestDto) GetFolderIds() []BaseBatchRequestDtoAllOfFolderIds`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *BaseBatchRequestDto) GetFolderIdsOk() (*[]BaseBatchRequestDtoAllOfFolderIds, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *BaseBatchRequestDto) SetFolderIds(v []BaseBatchRequestDtoAllOfFolderIds)`

SetFolderIds sets FolderIds field to given value.

### HasFolderIds

`func (o *BaseBatchRequestDto) HasFolderIds() bool`

HasFolderIds returns a boolean if a field has been set.

### SetFolderIdsNil

`func (o *BaseBatchRequestDto) SetFolderIdsNil(b bool)`

 SetFolderIdsNil sets the value for FolderIds to be an explicit nil

### UnsetFolderIds
`func (o *BaseBatchRequestDto) UnsetFolderIds()`

UnsetFolderIds ensures that no value is present for FolderIds, not even an explicit nil
### GetFileIds

`func (o *BaseBatchRequestDto) GetFileIds() []BaseBatchRequestDtoAllOfFileIds`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *BaseBatchRequestDto) GetFileIdsOk() (*[]BaseBatchRequestDtoAllOfFileIds, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *BaseBatchRequestDto) SetFileIds(v []BaseBatchRequestDtoAllOfFileIds)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *BaseBatchRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *BaseBatchRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *BaseBatchRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


