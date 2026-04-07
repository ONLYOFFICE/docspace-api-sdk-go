# BackupScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | Pointer to [**BackupStorageType**](BackupStorageType.md) |  | [optional] 
**StorageParams** | Pointer to [**[]ItemKeyValuePairObjectObject**](ItemKeyValuePairObjectObject.md) | The backup storage parameters. | [optional] 
**BackupsStored** | Pointer to **NullableInt32** | The maximum number of the stored backup copies. | [optional] 
**CronParams** | Pointer to [**Cron**](Cron.md) |  | [optional] 
**Dump** | Pointer to **bool** | Specifies if a dump will be created or not. | [optional] 

## Methods

### NewBackupScheduleDto

`func NewBackupScheduleDto() *BackupScheduleDto`

NewBackupScheduleDto instantiates a new BackupScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleDtoWithDefaults

`func NewBackupScheduleDtoWithDefaults() *BackupScheduleDto`

NewBackupScheduleDtoWithDefaults instantiates a new BackupScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *BackupScheduleDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *BackupScheduleDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *BackupScheduleDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *BackupScheduleDto) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageParams

`func (o *BackupScheduleDto) GetStorageParams() []ItemKeyValuePairObjectObject`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *BackupScheduleDto) GetStorageParamsOk() (*[]ItemKeyValuePairObjectObject, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *BackupScheduleDto) SetStorageParams(v []ItemKeyValuePairObjectObject)`

SetStorageParams sets StorageParams field to given value.

### HasStorageParams

`func (o *BackupScheduleDto) HasStorageParams() bool`

HasStorageParams returns a boolean if a field has been set.

### SetStorageParamsNil

`func (o *BackupScheduleDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *BackupScheduleDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetBackupsStored

`func (o *BackupScheduleDto) GetBackupsStored() int32`

GetBackupsStored returns the BackupsStored field if non-nil, zero value otherwise.

### GetBackupsStoredOk

`func (o *BackupScheduleDto) GetBackupsStoredOk() (*int32, bool)`

GetBackupsStoredOk returns a tuple with the BackupsStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsStored

`func (o *BackupScheduleDto) SetBackupsStored(v int32)`

SetBackupsStored sets BackupsStored field to given value.

### HasBackupsStored

`func (o *BackupScheduleDto) HasBackupsStored() bool`

HasBackupsStored returns a boolean if a field has been set.

### SetBackupsStoredNil

`func (o *BackupScheduleDto) SetBackupsStoredNil(b bool)`

 SetBackupsStoredNil sets the value for BackupsStored to be an explicit nil

### UnsetBackupsStored
`func (o *BackupScheduleDto) UnsetBackupsStored()`

UnsetBackupsStored ensures that no value is present for BackupsStored, not even an explicit nil
### GetCronParams

`func (o *BackupScheduleDto) GetCronParams() Cron`

GetCronParams returns the CronParams field if non-nil, zero value otherwise.

### GetCronParamsOk

`func (o *BackupScheduleDto) GetCronParamsOk() (*Cron, bool)`

GetCronParamsOk returns a tuple with the CronParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronParams

`func (o *BackupScheduleDto) SetCronParams(v Cron)`

SetCronParams sets CronParams field to given value.

### HasCronParams

`func (o *BackupScheduleDto) HasCronParams() bool`

HasCronParams returns a boolean if a field has been set.

### GetDump

`func (o *BackupScheduleDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *BackupScheduleDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *BackupScheduleDto) SetDump(v bool)`

SetDump sets Dump field to given value.

### HasDump

`func (o *BackupScheduleDto) HasDump() bool`

HasDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


