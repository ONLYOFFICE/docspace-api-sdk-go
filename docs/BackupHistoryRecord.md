# BackupHistoryRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The backup ID. | 
**FileName** | **NullableString** | The backup file name. | 
**StorageType** | [**BackupStorageType**](BackupStorageType.md) |  | 
**CreatedOn** | **time.Time** | The backup creation date. | 
**ExpiresOn** | **time.Time** | The backup expiration date. | 

## Methods

### NewBackupHistoryRecord

`func NewBackupHistoryRecord(id string, fileName NullableString, storageType BackupStorageType, createdOn time.Time, expiresOn time.Time, ) *BackupHistoryRecord`

NewBackupHistoryRecord instantiates a new BackupHistoryRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHistoryRecordWithDefaults

`func NewBackupHistoryRecordWithDefaults() *BackupHistoryRecord`

NewBackupHistoryRecordWithDefaults instantiates a new BackupHistoryRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupHistoryRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupHistoryRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupHistoryRecord) SetId(v string)`

SetId sets Id field to given value.


### GetFileName

`func (o *BackupHistoryRecord) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BackupHistoryRecord) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BackupHistoryRecord) SetFileName(v string)`

SetFileName sets FileName field to given value.


### SetFileNameNil

`func (o *BackupHistoryRecord) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *BackupHistoryRecord) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetStorageType

`func (o *BackupHistoryRecord) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *BackupHistoryRecord) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *BackupHistoryRecord) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedOn

`func (o *BackupHistoryRecord) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BackupHistoryRecord) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BackupHistoryRecord) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetExpiresOn

`func (o *BackupHistoryRecord) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *BackupHistoryRecord) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *BackupHistoryRecord) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


