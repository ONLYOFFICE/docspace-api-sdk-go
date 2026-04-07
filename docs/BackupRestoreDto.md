# BackupRestoreDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **NullableString** | The backup ID. | 
**StorageType** | Pointer to [**BackupStorageType**](BackupStorageType.md) |  | [optional] 
**StorageParams** | Pointer to [**[]ItemKeyValuePairObjectObject**](ItemKeyValuePairObjectObject.md) | The backup storage parameters. | [optional] 
**Notify** | Pointer to **bool** | Notifies users about the portal restoring process or not. | [optional] 
**Dump** | Pointer to **bool** | Specifies if a dump will be created or not. | [optional] 

## Methods

### NewBackupRestoreDto

`func NewBackupRestoreDto(backupId NullableString, ) *BackupRestoreDto`

NewBackupRestoreDto instantiates a new BackupRestoreDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreDtoWithDefaults

`func NewBackupRestoreDtoWithDefaults() *BackupRestoreDto`

NewBackupRestoreDtoWithDefaults instantiates a new BackupRestoreDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *BackupRestoreDto) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupRestoreDto) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupRestoreDto) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### SetBackupIdNil

`func (o *BackupRestoreDto) SetBackupIdNil(b bool)`

 SetBackupIdNil sets the value for BackupId to be an explicit nil

### UnsetBackupId
`func (o *BackupRestoreDto) UnsetBackupId()`

UnsetBackupId ensures that no value is present for BackupId, not even an explicit nil
### GetStorageType

`func (o *BackupRestoreDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *BackupRestoreDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *BackupRestoreDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *BackupRestoreDto) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageParams

`func (o *BackupRestoreDto) GetStorageParams() []ItemKeyValuePairObjectObject`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *BackupRestoreDto) GetStorageParamsOk() (*[]ItemKeyValuePairObjectObject, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *BackupRestoreDto) SetStorageParams(v []ItemKeyValuePairObjectObject)`

SetStorageParams sets StorageParams field to given value.

### HasStorageParams

`func (o *BackupRestoreDto) HasStorageParams() bool`

HasStorageParams returns a boolean if a field has been set.

### SetStorageParamsNil

`func (o *BackupRestoreDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *BackupRestoreDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetNotify

`func (o *BackupRestoreDto) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *BackupRestoreDto) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *BackupRestoreDto) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *BackupRestoreDto) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetDump

`func (o *BackupRestoreDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *BackupRestoreDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *BackupRestoreDto) SetDump(v bool)`

SetDump sets Dump field to given value.

### HasDump

`func (o *BackupRestoreDto) HasDump() bool`

HasDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


