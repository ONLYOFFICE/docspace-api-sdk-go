# StartBackupRestoreRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **NullableString** | The ID of the backup to restore from, as listed by `GET api/2.0/backup/getbackuphistory`. Send  anything that is not a GUID to restore from a file given by `storageParams` instead; an all-zero GUID  selects neither, because it parses as a GUID and then matches no record. | 
**StorageType** | Pointer to [**BackupStorageType**](BackupStorageType.md) | The storage the archive is read from. It defaults to `Documents` and is only used when `backupId` is  not a GUID, because a known backup carries the storage of its own record. | [optional] 
**StorageParams** | Pointer to [**[]ItemKeyValuePairObjectObject**](ItemKeyValuePairObjectObject.md) | The location of the archive, as an array of key and value pairs. The key read here is `filePath` -  not the `folderId` a backup is started with - and it holds a file ID for `Documents`, a  provider-specific file ID for `ThridpartyDocuments` and a path on the server for `Local`. It is only  used when `backupId` is not a GUID. | [optional] 
**Notify** | Pointer to **bool** | Chooses who is emailed when the restoring starts and when it finishes: every active user of the  portal when true, and its owner alone when false. Mail goes only to accounts that have been  activated, so this decides the audience rather than whether anybody is notified at all. | [optional] 
**Dump** | Pointer to **bool** | Restores the whole server rather than this one portal. It requires the space access permission. | [optional] 

## Methods

### NewStartBackupRestoreRequestDto

`func NewStartBackupRestoreRequestDto(backupId NullableString, ) *StartBackupRestoreRequestDto`

NewStartBackupRestoreRequestDto instantiates a new StartBackupRestoreRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartBackupRestoreRequestDtoWithDefaults

`func NewStartBackupRestoreRequestDtoWithDefaults() *StartBackupRestoreRequestDto`

NewStartBackupRestoreRequestDtoWithDefaults instantiates a new StartBackupRestoreRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *StartBackupRestoreRequestDto) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *StartBackupRestoreRequestDto) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *StartBackupRestoreRequestDto) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### SetBackupIdNil

`func (o *StartBackupRestoreRequestDto) SetBackupIdNil(b bool)`

 SetBackupIdNil sets the value for BackupId to be an explicit nil

### UnsetBackupId
`func (o *StartBackupRestoreRequestDto) UnsetBackupId()`

UnsetBackupId ensures that no value is present for BackupId, not even an explicit nil
### GetStorageType

`func (o *StartBackupRestoreRequestDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *StartBackupRestoreRequestDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *StartBackupRestoreRequestDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *StartBackupRestoreRequestDto) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageParams

`func (o *StartBackupRestoreRequestDto) GetStorageParams() []ItemKeyValuePairObjectObject`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *StartBackupRestoreRequestDto) GetStorageParamsOk() (*[]ItemKeyValuePairObjectObject, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *StartBackupRestoreRequestDto) SetStorageParams(v []ItemKeyValuePairObjectObject)`

SetStorageParams sets StorageParams field to given value.

### HasStorageParams

`func (o *StartBackupRestoreRequestDto) HasStorageParams() bool`

HasStorageParams returns a boolean if a field has been set.

### SetStorageParamsNil

`func (o *StartBackupRestoreRequestDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *StartBackupRestoreRequestDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetNotify

`func (o *StartBackupRestoreRequestDto) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *StartBackupRestoreRequestDto) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *StartBackupRestoreRequestDto) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *StartBackupRestoreRequestDto) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetDump

`func (o *StartBackupRestoreRequestDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *StartBackupRestoreRequestDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *StartBackupRestoreRequestDto) SetDump(v bool)`

SetDump sets Dump field to given value.

### HasDump

`func (o *StartBackupRestoreRequestDto) HasDump() bool`

HasDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


