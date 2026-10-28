# CreateBackupScheduleRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | Pointer to [**BackupStorageType**](BackupStorageType.md) | The storage the scheduled archives are written to. It defaults to `Documents`, and it decides which  keys `storageParams` has to carry. | [optional] 
**StorageParams** | Pointer to [**[]ItemKeyValuePairObjectObject**](ItemKeyValuePairObjectObject.md) | The settings of the chosen storage, as an array of key and value pairs. `Documents` and  `ThridpartyDocuments` need `folderId`, `Local` needs `filePath`, `ThirdPartyConsumer` needs `module`  plus the settings of that consumer, and `DataStore` needs none. | [optional] 
**BackupsStored** | Pointer to **NullableInt32** | The number of scheduled copies to keep, from 1 to 30. It defaults to 1, and only the copies this  schedule creates are counted and removed - archives started by hand are left alone. | [optional] 
**CronParams** | Pointer to [**BackupCronRequest**](BackupCronRequest.md) | When the backup runs. It is required: a request without it is rejected rather than falling back to  a default. | [optional] 
**Dump** | Pointer to **bool** | Schedules a backup of the whole server rather than of this one portal. It requires the space access  permission and works on a standalone installation only. | [optional] 

## Methods

### NewCreateBackupScheduleRequestDto

`func NewCreateBackupScheduleRequestDto() *CreateBackupScheduleRequestDto`

NewCreateBackupScheduleRequestDto instantiates a new CreateBackupScheduleRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupScheduleRequestDtoWithDefaults

`func NewCreateBackupScheduleRequestDtoWithDefaults() *CreateBackupScheduleRequestDto`

NewCreateBackupScheduleRequestDtoWithDefaults instantiates a new CreateBackupScheduleRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *CreateBackupScheduleRequestDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CreateBackupScheduleRequestDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CreateBackupScheduleRequestDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *CreateBackupScheduleRequestDto) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageParams

`func (o *CreateBackupScheduleRequestDto) GetStorageParams() []ItemKeyValuePairObjectObject`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *CreateBackupScheduleRequestDto) GetStorageParamsOk() (*[]ItemKeyValuePairObjectObject, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *CreateBackupScheduleRequestDto) SetStorageParams(v []ItemKeyValuePairObjectObject)`

SetStorageParams sets StorageParams field to given value.

### HasStorageParams

`func (o *CreateBackupScheduleRequestDto) HasStorageParams() bool`

HasStorageParams returns a boolean if a field has been set.

### SetStorageParamsNil

`func (o *CreateBackupScheduleRequestDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *CreateBackupScheduleRequestDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetBackupsStored

`func (o *CreateBackupScheduleRequestDto) GetBackupsStored() int32`

GetBackupsStored returns the BackupsStored field if non-nil, zero value otherwise.

### GetBackupsStoredOk

`func (o *CreateBackupScheduleRequestDto) GetBackupsStoredOk() (*int32, bool)`

GetBackupsStoredOk returns a tuple with the BackupsStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsStored

`func (o *CreateBackupScheduleRequestDto) SetBackupsStored(v int32)`

SetBackupsStored sets BackupsStored field to given value.

### HasBackupsStored

`func (o *CreateBackupScheduleRequestDto) HasBackupsStored() bool`

HasBackupsStored returns a boolean if a field has been set.

### SetBackupsStoredNil

`func (o *CreateBackupScheduleRequestDto) SetBackupsStoredNil(b bool)`

 SetBackupsStoredNil sets the value for BackupsStored to be an explicit nil

### UnsetBackupsStored
`func (o *CreateBackupScheduleRequestDto) UnsetBackupsStored()`

UnsetBackupsStored ensures that no value is present for BackupsStored, not even an explicit nil
### GetCronParams

`func (o *CreateBackupScheduleRequestDto) GetCronParams() BackupCronRequest`

GetCronParams returns the CronParams field if non-nil, zero value otherwise.

### GetCronParamsOk

`func (o *CreateBackupScheduleRequestDto) GetCronParamsOk() (*BackupCronRequest, bool)`

GetCronParamsOk returns a tuple with the CronParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronParams

`func (o *CreateBackupScheduleRequestDto) SetCronParams(v BackupCronRequest)`

SetCronParams sets CronParams field to given value.

### HasCronParams

`func (o *CreateBackupScheduleRequestDto) HasCronParams() bool`

HasCronParams returns a boolean if a field has been set.

### GetDump

`func (o *CreateBackupScheduleRequestDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *CreateBackupScheduleRequestDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *CreateBackupScheduleRequestDto) SetDump(v bool)`

SetDump sets Dump field to given value.

### HasDump

`func (o *CreateBackupScheduleRequestDto) HasDump() bool`

HasDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


