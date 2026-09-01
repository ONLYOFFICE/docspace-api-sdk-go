# FileEncryptionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserKeys** | Pointer to [**[]EncryptionKeyDto**](EncryptionKeyDto.md) | The key pairs of the users who have access to the file. | [optional] 
**FileKeys** | Pointer to [**[]FileKeys**](FileKeys.md) | The file keys issued to those users. | [optional] 

## Methods

### NewFileEncryptionInfoDto

`func NewFileEncryptionInfoDto() *FileEncryptionInfoDto`

NewFileEncryptionInfoDto instantiates a new FileEncryptionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileEncryptionInfoDtoWithDefaults

`func NewFileEncryptionInfoDtoWithDefaults() *FileEncryptionInfoDto`

NewFileEncryptionInfoDtoWithDefaults instantiates a new FileEncryptionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserKeys

`func (o *FileEncryptionInfoDto) GetUserKeys() []EncryptionKeyDto`

GetUserKeys returns the UserKeys field if non-nil, zero value otherwise.

### GetUserKeysOk

`func (o *FileEncryptionInfoDto) GetUserKeysOk() (*[]EncryptionKeyDto, bool)`

GetUserKeysOk returns a tuple with the UserKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeys

`func (o *FileEncryptionInfoDto) SetUserKeys(v []EncryptionKeyDto)`

SetUserKeys sets UserKeys field to given value.

### HasUserKeys

`func (o *FileEncryptionInfoDto) HasUserKeys() bool`

HasUserKeys returns a boolean if a field has been set.

### SetUserKeysNil

`func (o *FileEncryptionInfoDto) SetUserKeysNil(b bool)`

 SetUserKeysNil sets the value for UserKeys to be an explicit nil

### UnsetUserKeys
`func (o *FileEncryptionInfoDto) UnsetUserKeys()`

UnsetUserKeys ensures that no value is present for UserKeys, not even an explicit nil
### GetFileKeys

`func (o *FileEncryptionInfoDto) GetFileKeys() []FileKeys`

GetFileKeys returns the FileKeys field if non-nil, zero value otherwise.

### GetFileKeysOk

`func (o *FileEncryptionInfoDto) GetFileKeysOk() (*[]FileKeys, bool)`

GetFileKeysOk returns a tuple with the FileKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKeys

`func (o *FileEncryptionInfoDto) SetFileKeys(v []FileKeys)`

SetFileKeys sets FileKeys field to given value.

### HasFileKeys

`func (o *FileEncryptionInfoDto) HasFileKeys() bool`

HasFileKeys returns a boolean if a field has been set.

### SetFileKeysNil

`func (o *FileEncryptionInfoDto) SetFileKeysNil(b bool)`

 SetFileKeysNil sets the value for FileKeys to be an explicit nil

### UnsetFileKeys
`func (o *FileEncryptionInfoDto) UnsetFileKeys()`

UnsetFileKeys ensures that no value is present for FileKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


