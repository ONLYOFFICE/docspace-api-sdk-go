# UploadSessionResponseDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The upload session ID. | [optional] 
**FolderId** | Pointer to **int32** | The folder ID where the file is being uploaded. | [optional] 
**Version** | Pointer to **int32** | The file version number. | [optional] 
**Title** | Pointer to **NullableString** | The file title. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The third-party provider key. | [optional] 
**Uploaded** | Pointer to **bool** | Specifies whether the file has been uploaded. | [optional] 
**File** | Pointer to [**FileDtoInteger**](FileDtoInteger.md) |  | [optional] 

## Methods

### NewUploadSessionResponseDtoInteger

`func NewUploadSessionResponseDtoInteger() *UploadSessionResponseDtoInteger`

NewUploadSessionResponseDtoInteger instantiates a new UploadSessionResponseDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSessionResponseDtoIntegerWithDefaults

`func NewUploadSessionResponseDtoIntegerWithDefaults() *UploadSessionResponseDtoInteger`

NewUploadSessionResponseDtoIntegerWithDefaults instantiates a new UploadSessionResponseDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadSessionResponseDtoInteger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadSessionResponseDtoInteger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadSessionResponseDtoInteger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UploadSessionResponseDtoInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFolderId

`func (o *UploadSessionResponseDtoInteger) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UploadSessionResponseDtoInteger) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UploadSessionResponseDtoInteger) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UploadSessionResponseDtoInteger) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVersion

`func (o *UploadSessionResponseDtoInteger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UploadSessionResponseDtoInteger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UploadSessionResponseDtoInteger) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UploadSessionResponseDtoInteger) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *UploadSessionResponseDtoInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UploadSessionResponseDtoInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UploadSessionResponseDtoInteger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UploadSessionResponseDtoInteger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UploadSessionResponseDtoInteger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UploadSessionResponseDtoInteger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetProviderKey

`func (o *UploadSessionResponseDtoInteger) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *UploadSessionResponseDtoInteger) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *UploadSessionResponseDtoInteger) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *UploadSessionResponseDtoInteger) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *UploadSessionResponseDtoInteger) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *UploadSessionResponseDtoInteger) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetUploaded

`func (o *UploadSessionResponseDtoInteger) GetUploaded() bool`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *UploadSessionResponseDtoInteger) GetUploadedOk() (*bool, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *UploadSessionResponseDtoInteger) SetUploaded(v bool)`

SetUploaded sets Uploaded field to given value.

### HasUploaded

`func (o *UploadSessionResponseDtoInteger) HasUploaded() bool`

HasUploaded returns a boolean if a field has been set.

### GetFile

`func (o *UploadSessionResponseDtoInteger) GetFile() FileDtoInteger`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadSessionResponseDtoInteger) GetFileOk() (*FileDtoInteger, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadSessionResponseDtoInteger) SetFile(v FileDtoInteger)`

SetFile sets File field to given value.

### HasFile

`func (o *UploadSessionResponseDtoInteger) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


