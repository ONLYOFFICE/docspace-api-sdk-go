# ScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**BackupStorageType**](BackupStorageType.md) |  | 
**StorageParams** | **map[string]string** | The backup storage parameters. | 
**CronParams** | [**CronParams**](CronParams.md) |  | 
**BackupsStored** | Pointer to **NullableInt32** | The maximum number of the stored backup copies. | [optional] 
**LastBackupTime** | **time.Time** | The date and time when the last backup was reated. | 
**Dump** | **bool** | Specifies if a dump will be created or not. | 

## Methods

### NewScheduleDto

`func NewScheduleDto(storageType BackupStorageType, storageParams map[string]string, cronParams CronParams, lastBackupTime time.Time, dump bool, ) *ScheduleDto`

NewScheduleDto instantiates a new ScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDtoWithDefaults

`func NewScheduleDtoWithDefaults() *ScheduleDto`

NewScheduleDtoWithDefaults instantiates a new ScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *ScheduleDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ScheduleDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ScheduleDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.


### GetStorageParams

`func (o *ScheduleDto) GetStorageParams() map[string]string`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *ScheduleDto) GetStorageParamsOk() (*map[string]string, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *ScheduleDto) SetStorageParams(v map[string]string)`

SetStorageParams sets StorageParams field to given value.


### SetStorageParamsNil

`func (o *ScheduleDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *ScheduleDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetCronParams

`func (o *ScheduleDto) GetCronParams() CronParams`

GetCronParams returns the CronParams field if non-nil, zero value otherwise.

### GetCronParamsOk

`func (o *ScheduleDto) GetCronParamsOk() (*CronParams, bool)`

GetCronParamsOk returns a tuple with the CronParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronParams

`func (o *ScheduleDto) SetCronParams(v CronParams)`

SetCronParams sets CronParams field to given value.


### GetBackupsStored

`func (o *ScheduleDto) GetBackupsStored() int32`

GetBackupsStored returns the BackupsStored field if non-nil, zero value otherwise.

### GetBackupsStoredOk

`func (o *ScheduleDto) GetBackupsStoredOk() (*int32, bool)`

GetBackupsStoredOk returns a tuple with the BackupsStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsStored

`func (o *ScheduleDto) SetBackupsStored(v int32)`

SetBackupsStored sets BackupsStored field to given value.

### HasBackupsStored

`func (o *ScheduleDto) HasBackupsStored() bool`

HasBackupsStored returns a boolean if a field has been set.

### SetBackupsStoredNil

`func (o *ScheduleDto) SetBackupsStoredNil(b bool)`

 SetBackupsStoredNil sets the value for BackupsStored to be an explicit nil

### UnsetBackupsStored
`func (o *ScheduleDto) UnsetBackupsStored()`

UnsetBackupsStored ensures that no value is present for BackupsStored, not even an explicit nil
### GetLastBackupTime

`func (o *ScheduleDto) GetLastBackupTime() time.Time`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *ScheduleDto) GetLastBackupTimeOk() (*time.Time, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *ScheduleDto) SetLastBackupTime(v time.Time)`

SetLastBackupTime sets LastBackupTime field to given value.


### GetDump

`func (o *ScheduleDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *ScheduleDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *ScheduleDto) SetDump(v bool)`

SetDump sets Dump field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


