# UploadSessionResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The file the parts are being written into. An upload that took over a file of the same title carries it from  the start, while an upload that creates a new file has nothing to name yet and reports 0 until the answer that  sets `uploaded` to true. | [optional] 
**FolderId** | Pointer to **int32** | The folder receiving the file. It is the folder the upload was reserved against, or the sub-folder created for  it when the reservation declared a relative path. | [optional] 
**Version** | Pointer to **int32** | The revision the content is being written as: 1 for a file that did not exist, the next number when the upload  took over a file of the same title, and the unchanged current number for an upload opened over an existing  file, which replaces its content in place. | [optional] 
**Title** | Pointer to **NullableString** | The title the file is stored under, after characters a title cannot hold were replaced and, where a second  copy was asked for, a numeric suffix was added - so it can differ from the name that was sent. | [optional] 
**ProviderKey** | Pointer to **NullableString** | The third-party service holding the destination, such as `GoogleDrive` or `OneDrive`, and null for a folder  stored on the portal itself. | [optional] 
**Uploaded** | Pointer to **bool** | False while bytes are still missing, when the answer only reports progress; true in the answer that reports  the stored file, which is also the answer that arrives with 201. | [optional] 
**File** | Pointer to [**FileDto**](FileDto.md) | The file as it stands. It is filled in both answers, but while `uploaded` is false it describes a file that  has not been written yet, so its identifier, size and links are only worth reading once that flag turns true. | [optional] 

## Methods

### NewUploadSessionResponseDto

`func NewUploadSessionResponseDto() *UploadSessionResponseDto`

NewUploadSessionResponseDto instantiates a new UploadSessionResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSessionResponseDtoWithDefaults

`func NewUploadSessionResponseDtoWithDefaults() *UploadSessionResponseDto`

NewUploadSessionResponseDtoWithDefaults instantiates a new UploadSessionResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadSessionResponseDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadSessionResponseDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadSessionResponseDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UploadSessionResponseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFolderId

`func (o *UploadSessionResponseDto) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UploadSessionResponseDto) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UploadSessionResponseDto) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UploadSessionResponseDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetVersion

`func (o *UploadSessionResponseDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UploadSessionResponseDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UploadSessionResponseDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UploadSessionResponseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTitle

`func (o *UploadSessionResponseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UploadSessionResponseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UploadSessionResponseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UploadSessionResponseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UploadSessionResponseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UploadSessionResponseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetProviderKey

`func (o *UploadSessionResponseDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *UploadSessionResponseDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *UploadSessionResponseDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *UploadSessionResponseDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *UploadSessionResponseDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *UploadSessionResponseDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetUploaded

`func (o *UploadSessionResponseDto) GetUploaded() bool`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *UploadSessionResponseDto) GetUploadedOk() (*bool, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *UploadSessionResponseDto) SetUploaded(v bool)`

SetUploaded sets Uploaded field to given value.

### HasUploaded

`func (o *UploadSessionResponseDto) HasUploaded() bool`

HasUploaded returns a boolean if a field has been set.

### GetFile

`func (o *UploadSessionResponseDto) GetFile() FileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadSessionResponseDto) GetFileOk() (*FileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadSessionResponseDto) SetFile(v FileDto)`

SetFile sets File field to given value.

### HasFile

`func (o *UploadSessionResponseDto) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


