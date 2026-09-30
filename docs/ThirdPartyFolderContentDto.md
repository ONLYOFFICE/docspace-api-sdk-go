# ThirdPartyFolderContentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The file entries of this page. It is empty when the folder holds no files, when the filters matched none of  them, and in the sections that list rooms only. | [optional] 
**Folders** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The folder entries of this page. In a section of rooms these entries are the rooms themselves, which is where  their type, tags, logo and quota are read from. | [optional] 
**Current** | Pointer to [**ThirdPartyFolderDto**](ThirdPartyFolderDto.md) | The folder or section the page was read from, with its own title, type and access rights. It describes the  container, not the entries, and is filled in even when the page is empty. | [optional] 
**PathParts** | **interface{}** |  | 
**StartIndex** | Pointer to **int32** | The position of the first entry of this page in the whole result, echoing the requested start index. Add the  number of entries received to it to ask for the next page. | [optional] 
**Count** | Pointer to **int32** | How many entries this page carries, files and folders together. A page shorter than the requested size means  the result is exhausted. | [optional] 
**Total** | **int32** | How many entries matched before paging was applied, across the whole folder. Page until the start index plus  the entries received reaches it. | 
**New** | Pointer to **int32** | How many entries of this folder are marked as new for the caller. It is 0 for every listing when the account  has switched the new-item badges off, so a zero here does not prove that nothing has changed. | [optional] 

## Methods

### NewThirdPartyFolderContentDto

`func NewThirdPartyFolderContentDto(pathParts interface{}, total int32, ) *ThirdPartyFolderContentDto`

NewThirdPartyFolderContentDto instantiates a new ThirdPartyFolderContentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyFolderContentDtoWithDefaults

`func NewThirdPartyFolderContentDtoWithDefaults() *ThirdPartyFolderContentDto`

NewThirdPartyFolderContentDtoWithDefaults instantiates a new ThirdPartyFolderContentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ThirdPartyFolderContentDto) GetFiles() []FileEntryBaseDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ThirdPartyFolderContentDto) GetFilesOk() (*[]FileEntryBaseDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ThirdPartyFolderContentDto) SetFiles(v []FileEntryBaseDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ThirdPartyFolderContentDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ThirdPartyFolderContentDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ThirdPartyFolderContentDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFolders

`func (o *ThirdPartyFolderContentDto) GetFolders() []FileEntryBaseDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *ThirdPartyFolderContentDto) GetFoldersOk() (*[]FileEntryBaseDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *ThirdPartyFolderContentDto) SetFolders(v []FileEntryBaseDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *ThirdPartyFolderContentDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *ThirdPartyFolderContentDto) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *ThirdPartyFolderContentDto) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetCurrent

`func (o *ThirdPartyFolderContentDto) GetCurrent() ThirdPartyFolderDto`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ThirdPartyFolderContentDto) GetCurrentOk() (*ThirdPartyFolderDto, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ThirdPartyFolderContentDto) SetCurrent(v ThirdPartyFolderDto)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ThirdPartyFolderContentDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetPathParts

`func (o *ThirdPartyFolderContentDto) GetPathParts() interface{}`

GetPathParts returns the PathParts field if non-nil, zero value otherwise.

### GetPathPartsOk

`func (o *ThirdPartyFolderContentDto) GetPathPartsOk() (*interface{}, bool)`

GetPathPartsOk returns a tuple with the PathParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParts

`func (o *ThirdPartyFolderContentDto) SetPathParts(v interface{})`

SetPathParts sets PathParts field to given value.


### SetPathPartsNil

`func (o *ThirdPartyFolderContentDto) SetPathPartsNil(b bool)`

 SetPathPartsNil sets the value for PathParts to be an explicit nil

### UnsetPathParts
`func (o *ThirdPartyFolderContentDto) UnsetPathParts()`

UnsetPathParts ensures that no value is present for PathParts, not even an explicit nil
### GetStartIndex

`func (o *ThirdPartyFolderContentDto) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ThirdPartyFolderContentDto) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ThirdPartyFolderContentDto) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ThirdPartyFolderContentDto) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetCount

`func (o *ThirdPartyFolderContentDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ThirdPartyFolderContentDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ThirdPartyFolderContentDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ThirdPartyFolderContentDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *ThirdPartyFolderContentDto) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ThirdPartyFolderContentDto) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ThirdPartyFolderContentDto) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNew

`func (o *ThirdPartyFolderContentDto) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *ThirdPartyFolderContentDto) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *ThirdPartyFolderContentDto) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *ThirdPartyFolderContentDto) HasNew() bool`

HasNew returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


