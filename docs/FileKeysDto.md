# FileKeysDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | The identifier of the user the file key was issued to. | [optional] 
**PublicKeyId** | Pointer to **string** | The identifier of the key pair the file key is encrypted for. | [optional] 
**PrivateKeyEnc** | Pointer to **NullableString** | The file key, encrypted with the public key of the pair. | [optional] 
**TenantId** | Pointer to **int32** | The identifier of the portal the file belongs to. | [optional] 
**FileId** | Pointer to **int32** | The identifier of the file the key unlocks. | [optional] 
**CreateOn** | Pointer to **time.Time** | The date and time when the file key was issued. | [optional] 

## Methods

### NewFileKeysDto

`func NewFileKeysDto() *FileKeysDto`

NewFileKeysDto instantiates a new FileKeysDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileKeysDtoWithDefaults

`func NewFileKeysDtoWithDefaults() *FileKeysDto`

NewFileKeysDtoWithDefaults instantiates a new FileKeysDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *FileKeysDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileKeysDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileKeysDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FileKeysDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPublicKeyId

`func (o *FileKeysDto) GetPublicKeyId() string`

GetPublicKeyId returns the PublicKeyId field if non-nil, zero value otherwise.

### GetPublicKeyIdOk

`func (o *FileKeysDto) GetPublicKeyIdOk() (*string, bool)`

GetPublicKeyIdOk returns a tuple with the PublicKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyId

`func (o *FileKeysDto) SetPublicKeyId(v string)`

SetPublicKeyId sets PublicKeyId field to given value.

### HasPublicKeyId

`func (o *FileKeysDto) HasPublicKeyId() bool`

HasPublicKeyId returns a boolean if a field has been set.

### GetPrivateKeyEnc

`func (o *FileKeysDto) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *FileKeysDto) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *FileKeysDto) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *FileKeysDto) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *FileKeysDto) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *FileKeysDto) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
### GetTenantId

`func (o *FileKeysDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FileKeysDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FileKeysDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FileKeysDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetFileId

`func (o *FileKeysDto) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileKeysDto) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileKeysDto) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileKeysDto) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetCreateOn

`func (o *FileKeysDto) GetCreateOn() time.Time`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *FileKeysDto) GetCreateOnOk() (*time.Time, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *FileKeysDto) SetCreateOn(v time.Time)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *FileKeysDto) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


