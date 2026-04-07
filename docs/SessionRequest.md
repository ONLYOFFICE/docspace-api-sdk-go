# SessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **NullableString** | The file name. | 
**FileSize** | Pointer to **int64** | The file size. | [optional] 
**RelativePath** | Pointer to **NullableString** | The relative path to the file. | [optional] 
**CreateOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Encrypted** | Pointer to **bool** | Specifies whether the file is encrypted or not. | [optional] 
**CreateNewIfExist** | Pointer to **bool** | Specifies whether to create a new file if it already exists. | [optional] 

## Methods

### NewSessionRequest

`func NewSessionRequest(fileName NullableString, ) *SessionRequest`

NewSessionRequest instantiates a new SessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRequestWithDefaults

`func NewSessionRequestWithDefaults() *SessionRequest`

NewSessionRequestWithDefaults instantiates a new SessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *SessionRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SessionRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SessionRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### SetFileNameNil

`func (o *SessionRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SessionRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileSize

`func (o *SessionRequest) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *SessionRequest) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *SessionRequest) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *SessionRequest) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetRelativePath

`func (o *SessionRequest) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *SessionRequest) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *SessionRequest) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *SessionRequest) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *SessionRequest) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *SessionRequest) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetCreateOn

`func (o *SessionRequest) GetCreateOn() ApiDateTime`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *SessionRequest) GetCreateOnOk() (*ApiDateTime, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *SessionRequest) SetCreateOn(v ApiDateTime)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *SessionRequest) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.

### GetEncrypted

`func (o *SessionRequest) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *SessionRequest) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *SessionRequest) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *SessionRequest) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetCreateNewIfExist

`func (o *SessionRequest) GetCreateNewIfExist() bool`

GetCreateNewIfExist returns the CreateNewIfExist field if non-nil, zero value otherwise.

### GetCreateNewIfExistOk

`func (o *SessionRequest) GetCreateNewIfExistOk() (*bool, bool)`

GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewIfExist

`func (o *SessionRequest) SetCreateNewIfExist(v bool)`

SetCreateNewIfExist sets CreateNewIfExist field to given value.

### HasCreateNewIfExist

`func (o *SessionRequest) HasCreateNewIfExist() bool`

HasCreateNewIfExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


