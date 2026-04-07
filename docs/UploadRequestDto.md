# UploadRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **Nullable*os.File** | The file to be uploaded. | [optional] 
**ContentType** | Pointer to [**ContentType**](ContentType.md) |  | [optional] 
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**Files** | Pointer to **[]*os.File** | The list of files when specified as multipart/form-data. | [optional] 
**CreateNewIfExist** | Pointer to **bool** | Specifies whether to create the new file if it already exists or not. | [optional] 
**StoreOriginalFileFlag** | Pointer to **NullableBool** | Specifies whether to upload documents in the original formats as well or not. | [optional] 
**KeepConvertStatus** | Pointer to **bool** | Specifies whether to keep the file converting status or not. | [optional] 
**Stream** | Pointer to **Nullable*os.File** | The request input stream. | [optional] 

## Methods

### NewUploadRequestDto

`func NewUploadRequestDto() *UploadRequestDto`

NewUploadRequestDto instantiates a new UploadRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadRequestDtoWithDefaults

`func NewUploadRequestDtoWithDefaults() *UploadRequestDto`

NewUploadRequestDtoWithDefaults instantiates a new UploadRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *UploadRequestDto) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadRequestDto) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadRequestDto) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *UploadRequestDto) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *UploadRequestDto) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *UploadRequestDto) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetContentType

`func (o *UploadRequestDto) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadRequestDto) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadRequestDto) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *UploadRequestDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContentDisposition

`func (o *UploadRequestDto) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *UploadRequestDto) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *UploadRequestDto) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *UploadRequestDto) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetFiles

`func (o *UploadRequestDto) GetFiles() []*os.File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UploadRequestDto) GetFilesOk() (*[]*os.File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UploadRequestDto) SetFiles(v []*os.File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UploadRequestDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *UploadRequestDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *UploadRequestDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetCreateNewIfExist

`func (o *UploadRequestDto) GetCreateNewIfExist() bool`

GetCreateNewIfExist returns the CreateNewIfExist field if non-nil, zero value otherwise.

### GetCreateNewIfExistOk

`func (o *UploadRequestDto) GetCreateNewIfExistOk() (*bool, bool)`

GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewIfExist

`func (o *UploadRequestDto) SetCreateNewIfExist(v bool)`

SetCreateNewIfExist sets CreateNewIfExist field to given value.

### HasCreateNewIfExist

`func (o *UploadRequestDto) HasCreateNewIfExist() bool`

HasCreateNewIfExist returns a boolean if a field has been set.

### GetStoreOriginalFileFlag

`func (o *UploadRequestDto) GetStoreOriginalFileFlag() bool`

GetStoreOriginalFileFlag returns the StoreOriginalFileFlag field if non-nil, zero value otherwise.

### GetStoreOriginalFileFlagOk

`func (o *UploadRequestDto) GetStoreOriginalFileFlagOk() (*bool, bool)`

GetStoreOriginalFileFlagOk returns a tuple with the StoreOriginalFileFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOriginalFileFlag

`func (o *UploadRequestDto) SetStoreOriginalFileFlag(v bool)`

SetStoreOriginalFileFlag sets StoreOriginalFileFlag field to given value.

### HasStoreOriginalFileFlag

`func (o *UploadRequestDto) HasStoreOriginalFileFlag() bool`

HasStoreOriginalFileFlag returns a boolean if a field has been set.

### SetStoreOriginalFileFlagNil

`func (o *UploadRequestDto) SetStoreOriginalFileFlagNil(b bool)`

 SetStoreOriginalFileFlagNil sets the value for StoreOriginalFileFlag to be an explicit nil

### UnsetStoreOriginalFileFlag
`func (o *UploadRequestDto) UnsetStoreOriginalFileFlag()`

UnsetStoreOriginalFileFlag ensures that no value is present for StoreOriginalFileFlag, not even an explicit nil
### GetKeepConvertStatus

`func (o *UploadRequestDto) GetKeepConvertStatus() bool`

GetKeepConvertStatus returns the KeepConvertStatus field if non-nil, zero value otherwise.

### GetKeepConvertStatusOk

`func (o *UploadRequestDto) GetKeepConvertStatusOk() (*bool, bool)`

GetKeepConvertStatusOk returns a tuple with the KeepConvertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepConvertStatus

`func (o *UploadRequestDto) SetKeepConvertStatus(v bool)`

SetKeepConvertStatus sets KeepConvertStatus field to given value.

### HasKeepConvertStatus

`func (o *UploadRequestDto) HasKeepConvertStatus() bool`

HasKeepConvertStatus returns a boolean if a field has been set.

### GetStream

`func (o *UploadRequestDto) GetStream() *os.File`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *UploadRequestDto) GetStreamOk() (**os.File, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *UploadRequestDto) SetStream(v *os.File)`

SetStream sets Stream field to given value.

### HasStream

`func (o *UploadRequestDto) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *UploadRequestDto) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *UploadRequestDto) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


