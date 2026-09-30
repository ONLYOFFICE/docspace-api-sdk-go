# BackupProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCompleted** | Pointer to **bool** | Specifies whether the job has stopped running. This is the field to poll: true means the job will not  change any more, whether it succeeded, failed or was cancelled, and `status` tells which of the three  it is. | [optional] 
**Progress** | Pointer to **int32** | The share of the job that is already done, from 0 to 100. A job that has only been queued reports 0,  because the work starts when a separate worker service picks it up. | [optional] 
**Error** | Pointer to **NullableString** | The message of the error that stopped the job. It is an empty string, not null, while the job runs  and after a job that succeeded, so the sign of a failure is a non-empty value - and this is the only  place where the reason is reported. | [optional] 
**Warning** | Pointer to **NullableString** | A message about a job that stopped without failing: it names the entry inside the archive that lists  the files which could not be read, when a backup finished without some of them, and it says so when  the job was cancelled. It is an empty string otherwise, and it is only ever filled in for a backup  job - a cancelled restoring job leaves it empty. | [optional] 
**Link** | Pointer to **NullableString** | The link to download the stored archive. It is an empty string until the archive has been uploaded,  and it is only ever filled in for a backup job, never for a restoring one. | [optional] 
**TenantId** | Pointer to **int32** | The ID of the portal the job belongs to, or -1 for a job that covers the whole server. | [optional] 
**BackupProgressEnum** | Pointer to [**BackupProgressEnum**](BackupProgressEnum.md) | Whether this is a backup or a restoring job, reported as a number rather than as a name. | [optional] 
**Status** | Pointer to [**DistributedTaskStatus**](DistributedTaskStatus.md) | The state of the job: `Created` while it waits for a worker to pick it up, `Running` while it works,  `Completed` once it has finished on its own, `Canceled` after it was cancelled, and `Failted` when it  stopped on an error, in which case `error` carries the reason. Reported as a number rather than as a  name. | [optional] 
**TaskId** | Pointer to **NullableString** | The ID of the job. It is the handle to poll this operation with, and for a backup job it also becomes  the `id` of the record in `GET api/2.0/backup/getbackuphistory`. | [optional] 

## Methods

### NewBackupProgress

`func NewBackupProgress() *BackupProgress`

NewBackupProgress instantiates a new BackupProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupProgressWithDefaults

`func NewBackupProgressWithDefaults() *BackupProgress`

NewBackupProgressWithDefaults instantiates a new BackupProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCompleted

`func (o *BackupProgress) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *BackupProgress) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *BackupProgress) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *BackupProgress) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetProgress

`func (o *BackupProgress) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BackupProgress) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BackupProgress) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *BackupProgress) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetError

`func (o *BackupProgress) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackupProgress) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackupProgress) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BackupProgress) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BackupProgress) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BackupProgress) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetWarning

`func (o *BackupProgress) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *BackupProgress) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *BackupProgress) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *BackupProgress) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *BackupProgress) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *BackupProgress) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetLink

`func (o *BackupProgress) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BackupProgress) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BackupProgress) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *BackupProgress) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *BackupProgress) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *BackupProgress) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetTenantId

`func (o *BackupProgress) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BackupProgress) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BackupProgress) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BackupProgress) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetBackupProgressEnum

`func (o *BackupProgress) GetBackupProgressEnum() BackupProgressEnum`

GetBackupProgressEnum returns the BackupProgressEnum field if non-nil, zero value otherwise.

### GetBackupProgressEnumOk

`func (o *BackupProgress) GetBackupProgressEnumOk() (*BackupProgressEnum, bool)`

GetBackupProgressEnumOk returns a tuple with the BackupProgressEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProgressEnum

`func (o *BackupProgress) SetBackupProgressEnum(v BackupProgressEnum)`

SetBackupProgressEnum sets BackupProgressEnum field to given value.

### HasBackupProgressEnum

`func (o *BackupProgress) HasBackupProgressEnum() bool`

HasBackupProgressEnum returns a boolean if a field has been set.

### GetStatus

`func (o *BackupProgress) GetStatus() DistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupProgress) GetStatusOk() (*DistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupProgress) SetStatus(v DistributedTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *BackupProgress) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BackupProgress) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BackupProgress) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *BackupProgress) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *BackupProgress) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *BackupProgress) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


