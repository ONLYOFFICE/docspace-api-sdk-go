# DraftLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **int32** | The folder holding the draft: the sub-folder that the room for filling keeps for drafts of this particular  form. | [optional] 
**FolderTitle** | Pointer to **NullableString** | The title of that folder, which the portal takes from the form itself when the form is released for filling. | [optional] 
**FileId** | Pointer to **int32** | The draft itself - the copy the caller fills in, not the original form, and the identifier to pass to the file  operations while filling. | [optional] 
**FileTitle** | Pointer to **NullableString** | The title of the draft, which the portal builds from the name of the person filling it and the name of the  form. Null when the draft the record points at no longer exists. | [optional] 

## Methods

### NewDraftLocation

`func NewDraftLocation() *DraftLocation`

NewDraftLocation instantiates a new DraftLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftLocationWithDefaults

`func NewDraftLocationWithDefaults() *DraftLocation`

NewDraftLocationWithDefaults instantiates a new DraftLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *DraftLocation) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DraftLocation) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DraftLocation) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DraftLocation) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetFolderTitle

`func (o *DraftLocation) GetFolderTitle() string`

GetFolderTitle returns the FolderTitle field if non-nil, zero value otherwise.

### GetFolderTitleOk

`func (o *DraftLocation) GetFolderTitleOk() (*string, bool)`

GetFolderTitleOk returns a tuple with the FolderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderTitle

`func (o *DraftLocation) SetFolderTitle(v string)`

SetFolderTitle sets FolderTitle field to given value.

### HasFolderTitle

`func (o *DraftLocation) HasFolderTitle() bool`

HasFolderTitle returns a boolean if a field has been set.

### SetFolderTitleNil

`func (o *DraftLocation) SetFolderTitleNil(b bool)`

 SetFolderTitleNil sets the value for FolderTitle to be an explicit nil

### UnsetFolderTitle
`func (o *DraftLocation) UnsetFolderTitle()`

UnsetFolderTitle ensures that no value is present for FolderTitle, not even an explicit nil
### GetFileId

`func (o *DraftLocation) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DraftLocation) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DraftLocation) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *DraftLocation) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileTitle

`func (o *DraftLocation) GetFileTitle() string`

GetFileTitle returns the FileTitle field if non-nil, zero value otherwise.

### GetFileTitleOk

`func (o *DraftLocation) GetFileTitleOk() (*string, bool)`

GetFileTitleOk returns a tuple with the FileTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTitle

`func (o *DraftLocation) SetFileTitle(v string)`

SetFileTitle sets FileTitle field to given value.

### HasFileTitle

`func (o *DraftLocation) HasFileTitle() bool`

HasFileTitle returns a boolean if a field has been set.

### SetFileTitleNil

`func (o *DraftLocation) SetFileTitleNil(b bool)`

 SetFileTitleNil sets the value for FileTitle to be an explicit nil

### UnsetFileTitle
`func (o *DraftLocation) UnsetFileTitle()`

UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


