# DeleteVersionBatchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | [optional] 
**DeleteAfter** | Pointer to **bool** | Whether the finished operation is still reported: `false` keeps its final record readable through  `GET api/2.0/files/fileops` until it has been read once, `true` drops the record as soon as the work is done.  It does not postpone the deletion and does not delete anything of its own. | [optional] 
**FileId** | **int32** | The file whose history the versions are taken from; only files stored in the portal itself are addressed here. | 
**Versions** | **[]int32** | The version numbers to remove, as reported by `GET api/2.0/files/file/{fileId}/history`. At least one number  has to be sent: an empty list removes the file itself instead of one of its versions. The number of the  current version is refused outright, while a number that no longer exists is passed over without a complaint. | 

## Methods

### NewDeleteVersionBatchRequestDto

`func NewDeleteVersionBatchRequestDto(fileId int32, versions []int32, ) *DeleteVersionBatchRequestDto`

NewDeleteVersionBatchRequestDto instantiates a new DeleteVersionBatchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVersionBatchRequestDtoWithDefaults

`func NewDeleteVersionBatchRequestDtoWithDefaults() *DeleteVersionBatchRequestDto`

NewDeleteVersionBatchRequestDtoWithDefaults instantiates a new DeleteVersionBatchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *DeleteVersionBatchRequestDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *DeleteVersionBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *DeleteVersionBatchRequestDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *DeleteVersionBatchRequestDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.

### GetDeleteAfter

`func (o *DeleteVersionBatchRequestDto) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *DeleteVersionBatchRequestDto) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *DeleteVersionBatchRequestDto) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *DeleteVersionBatchRequestDto) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.

### GetFileId

`func (o *DeleteVersionBatchRequestDto) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DeleteVersionBatchRequestDto) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DeleteVersionBatchRequestDto) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetVersions

`func (o *DeleteVersionBatchRequestDto) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DeleteVersionBatchRequestDto) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DeleteVersionBatchRequestDto) SetVersions(v []int32)`

SetVersions sets Versions field to given value.


### SetVersionsNil

`func (o *DeleteVersionBatchRequestDto) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *DeleteVersionBatchRequestDto) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


