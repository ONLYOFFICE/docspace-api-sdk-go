# FileKeys

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

### NewFileKeys

`func NewFileKeys() *FileKeys`

NewFileKeys instantiates a new FileKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileKeysWithDefaults

`func NewFileKeysWithDefaults() *FileKeys`

NewFileKeysWithDefaults instantiates a new FileKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *FileKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FileKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FileKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FileKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPublicKeyId

`func (o *FileKeys) GetPublicKeyId() string`

GetPublicKeyId returns the PublicKeyId field if non-nil, zero value otherwise.

### GetPublicKeyIdOk

`func (o *FileKeys) GetPublicKeyIdOk() (*string, bool)`

GetPublicKeyIdOk returns a tuple with the PublicKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyId

`func (o *FileKeys) SetPublicKeyId(v string)`

SetPublicKeyId sets PublicKeyId field to given value.

### HasPublicKeyId

`func (o *FileKeys) HasPublicKeyId() bool`

HasPublicKeyId returns a boolean if a field has been set.

### GetPrivateKeyEnc

`func (o *FileKeys) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *FileKeys) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *FileKeys) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *FileKeys) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *FileKeys) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *FileKeys) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
### GetTenantId

`func (o *FileKeys) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FileKeys) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FileKeys) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FileKeys) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetFileId

`func (o *FileKeys) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileKeys) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileKeys) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *FileKeys) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetCreateOn

`func (o *FileKeys) GetCreateOn() time.Time`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *FileKeys) GetCreateOnOk() (*time.Time, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *FileKeys) SetCreateOn(v time.Time)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *FileKeys) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


