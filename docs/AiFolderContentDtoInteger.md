# AiFolderContentDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The list of files in the folder. | [optional] 
**Folders** | Pointer to [**[]AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The list of folders in the folder. | [optional] 
**Current** | Pointer to [**AiFolderDtoInteger**](AiFolderDtoInteger.md) | The current folder information. | [optional] 
**PathParts** | **interface{}** |  | 
**StartIndex** | Pointer to **int32** | The folder start index. | [optional] 
**Count** | Pointer to **int32** | The number of folder elements. | [optional] 
**Total** | **int32** | The total number of elements in the folder. | 
**New** | Pointer to **int32** | The new element index in the folder. | [optional] 

## Methods

### NewAiFolderContentDtoInteger

`func NewAiFolderContentDtoInteger(pathParts interface{}, total int32, ) *AiFolderContentDtoInteger`

NewAiFolderContentDtoInteger instantiates a new AiFolderContentDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFolderContentDtoIntegerWithDefaults

`func NewAiFolderContentDtoIntegerWithDefaults() *AiFolderContentDtoInteger`

NewAiFolderContentDtoIntegerWithDefaults instantiates a new AiFolderContentDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *AiFolderContentDtoInteger) GetFiles() []AiFileEntryBaseDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AiFolderContentDtoInteger) GetFilesOk() (*[]AiFileEntryBaseDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AiFolderContentDtoInteger) SetFiles(v []AiFileEntryBaseDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AiFolderContentDtoInteger) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *AiFolderContentDtoInteger) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *AiFolderContentDtoInteger) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFolders

`func (o *AiFolderContentDtoInteger) GetFolders() []AiFileEntryBaseDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *AiFolderContentDtoInteger) GetFoldersOk() (*[]AiFileEntryBaseDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *AiFolderContentDtoInteger) SetFolders(v []AiFileEntryBaseDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *AiFolderContentDtoInteger) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *AiFolderContentDtoInteger) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *AiFolderContentDtoInteger) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetCurrent

`func (o *AiFolderContentDtoInteger) GetCurrent() AiFolderDtoInteger`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AiFolderContentDtoInteger) GetCurrentOk() (*AiFolderDtoInteger, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AiFolderContentDtoInteger) SetCurrent(v AiFolderDtoInteger)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AiFolderContentDtoInteger) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetPathParts

`func (o *AiFolderContentDtoInteger) GetPathParts() interface{}`

GetPathParts returns the PathParts field if non-nil, zero value otherwise.

### GetPathPartsOk

`func (o *AiFolderContentDtoInteger) GetPathPartsOk() (*interface{}, bool)`

GetPathPartsOk returns a tuple with the PathParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParts

`func (o *AiFolderContentDtoInteger) SetPathParts(v interface{})`

SetPathParts sets PathParts field to given value.


### SetPathPartsNil

`func (o *AiFolderContentDtoInteger) SetPathPartsNil(b bool)`

 SetPathPartsNil sets the value for PathParts to be an explicit nil

### UnsetPathParts
`func (o *AiFolderContentDtoInteger) UnsetPathParts()`

UnsetPathParts ensures that no value is present for PathParts, not even an explicit nil
### GetStartIndex

`func (o *AiFolderContentDtoInteger) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *AiFolderContentDtoInteger) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *AiFolderContentDtoInteger) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *AiFolderContentDtoInteger) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetCount

`func (o *AiFolderContentDtoInteger) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AiFolderContentDtoInteger) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AiFolderContentDtoInteger) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AiFolderContentDtoInteger) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *AiFolderContentDtoInteger) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AiFolderContentDtoInteger) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AiFolderContentDtoInteger) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNew

`func (o *AiFolderContentDtoInteger) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *AiFolderContentDtoInteger) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *AiFolderContentDtoInteger) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *AiFolderContentDtoInteger) HasNew() bool`

HasNew returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


