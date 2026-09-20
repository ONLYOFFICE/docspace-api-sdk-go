# BackupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | Pointer to [**BackupStorageType**](BackupStorageType.md) | The storage the archive is written to. It defaults to `Documents`, and it decides which keys  `storageParams` has to carry. | [optional] 
**StorageParams** | Pointer to [**[]ItemKeyValuePairObjectObject**](ItemKeyValuePairObjectObject.md) | The settings of the chosen storage, as an array of key and value pairs. `Documents` needs an integer  `folderId`, `ThridpartyDocuments` a provider-specific non-integer `folderId`, `Local` a `filePath`,  `ThirdPartyConsumer` a `module` plus the settings of that consumer, and `DataStore` none. The  `subdir` key is added by the operation itself and must not be sent. | [optional] 
**Dump** | Pointer to **bool** | Backs up the whole server rather than this one portal. It requires the space access permission and  works on a standalone installation only. | [optional] 

## Methods

### NewBackupDto

`func NewBackupDto() *BackupDto`

NewBackupDto instantiates a new BackupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDtoWithDefaults

`func NewBackupDtoWithDefaults() *BackupDto`

NewBackupDtoWithDefaults instantiates a new BackupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *BackupDto) GetStorageType() BackupStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *BackupDto) GetStorageTypeOk() (*BackupStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *BackupDto) SetStorageType(v BackupStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *BackupDto) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageParams

`func (o *BackupDto) GetStorageParams() []ItemKeyValuePairObjectObject`

GetStorageParams returns the StorageParams field if non-nil, zero value otherwise.

### GetStorageParamsOk

`func (o *BackupDto) GetStorageParamsOk() (*[]ItemKeyValuePairObjectObject, bool)`

GetStorageParamsOk returns a tuple with the StorageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageParams

`func (o *BackupDto) SetStorageParams(v []ItemKeyValuePairObjectObject)`

SetStorageParams sets StorageParams field to given value.

### HasStorageParams

`func (o *BackupDto) HasStorageParams() bool`

HasStorageParams returns a boolean if a field has been set.

### SetStorageParamsNil

`func (o *BackupDto) SetStorageParamsNil(b bool)`

 SetStorageParamsNil sets the value for StorageParams to be an explicit nil

### UnsetStorageParams
`func (o *BackupDto) UnsetStorageParams()`

UnsetStorageParams ensures that no value is present for StorageParams, not even an explicit nil
### GetDump

`func (o *BackupDto) GetDump() bool`

GetDump returns the Dump field if non-nil, zero value otherwise.

### GetDumpOk

`func (o *BackupDto) GetDumpOk() (*bool, bool)`

GetDumpOk returns a tuple with the Dump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDump

`func (o *BackupDto) SetDump(v bool)`

SetDump sets Dump field to given value.

### HasDump

`func (o *BackupDto) HasDump() bool`

HasDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


