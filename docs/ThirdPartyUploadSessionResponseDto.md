# ThirdPartyUploadSessionResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The file the parts are being written into. An upload that took over a file of the same title carries it from  the start, while an upload that creates a new file has nothing to name yet and reports 0 until the answer that  sets `uploaded` to true. | [optional] 
**FolderId** | Pointer to **NullableString** | The folder receiving the file. It is the folder the upload was reserved against, or the sub-folder created for  it when the reservation declared a relative path. | [optional] 
**Version** | Pointer to **int32** | The revision the content is being written as: 1 for a file that did not exist, the next number when the upload  took over a file of the same title, and the unchanged current number for an upload opened over an existing  file, which replaces its content in place. | [optional] 
**Title** | Pointer to **NullableString** | The title the file is stored under, after characters a title cannot hold were replaced and, where a second  copy was asked for, a numeric suffix was added - so it can differ from the name that was sent. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The third-party service holding the destination, such as `GoogleDrive` or `OneDrive`, and null for a folder  stored on the portal itself. | [optional] 
**Uploaded** | Pointer to **bool** | False while bytes are still missing, when the answer only reports progress; true in the answer that reports  the stored file, which is also the answer that arrives with 201. | [optional] 
**File** | Pointer to [**ThirdPartyFileDto**](ThirdPartyFileDto.md) | The file as it stands. It is filled in both answers, but while `uploaded` is false it describes a file that  has not been written yet, so its identifier, size and links are only worth reading once that flag turns true. | [optional] 

## Methods

### NewThirdPartyUploadSessionResponseDto

`func NewThirdPartyUploadSessionResponseDto() *ThirdPartyUploadSessionResponseDto`

NewThirdPartyUploadSessionResponseDto instantiates a new ThirdPartyUploadSessionResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyUploadSessionResponseDtoWithDefaults

`func NewThirdPartyUploadSessionResponseDtoWithDefaults() *ThirdPartyUploadSessionResponseDto`

NewThirdPartyUploadSessionResponseDtoWithDefaults instantiates a new ThirdPartyUploadSessionResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyUploadSessionResponseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyUploadSessionResponseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyUploadSessionResponseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartyUploadSessionResponseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ThirdPartyUploadSessionResponseDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ThirdPartyUploadSessionResponseDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFolderId

`func (o *ThirdPartyUploadSessionResponseDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ThirdPartyUploadSessionResponseDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ThirdPartyUploadSessionResponseDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ThirdPartyUploadSessionResponseDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ThirdPartyUploadSessionResponseDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ThirdPartyUploadSessionResponseDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetVersion

`func (o *ThirdPartyUploadSessionResponseDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdPartyUploadSessionResponseDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdPartyUploadSessionResponseDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdPartyUploadSessionResponseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *ThirdPartyUploadSessionResponseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ThirdPartyUploadSessionResponseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ThirdPartyUploadSessionResponseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ThirdPartyUploadSessionResponseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ThirdPartyUploadSessionResponseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ThirdPartyUploadSessionResponseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetProviderKey

`func (o *ThirdPartyUploadSessionResponseDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *ThirdPartyUploadSessionResponseDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *ThirdPartyUploadSessionResponseDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *ThirdPartyUploadSessionResponseDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *ThirdPartyUploadSessionResponseDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *ThirdPartyUploadSessionResponseDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetUploaded

`func (o *ThirdPartyUploadSessionResponseDto) GetUploaded() bool`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *ThirdPartyUploadSessionResponseDto) GetUploadedOk() (*bool, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *ThirdPartyUploadSessionResponseDto) SetUploaded(v bool)`

SetUploaded sets Uploaded field to given value.

### HasUploaded

`func (o *ThirdPartyUploadSessionResponseDto) HasUploaded() bool`

HasUploaded returns a boolean if a field has been set.

### GetFile

`func (o *ThirdPartyUploadSessionResponseDto) GetFile() ThirdPartyFileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ThirdPartyUploadSessionResponseDto) GetFileOk() (*ThirdPartyFileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ThirdPartyUploadSessionResponseDto) SetFile(v ThirdPartyFileDto)`

SetFile sets File field to given value.

### HasFile

`func (o *ThirdPartyUploadSessionResponseDto) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


