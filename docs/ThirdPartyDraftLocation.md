# ThirdPartyDraftLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **NullableString** | The folder holding the draft: the sub-folder that the room for filling keeps for drafts of this particular  form. | [optional] 
**FolderTitle** | Pointer to **NullableString** | The title of that folder, which the portal takes from the form itself when the form is released for filling. | [optional] 
**FileId** | Pointer to **NullableString** | The draft itself - the copy the caller fills in, not the original form, and the identifier to pass to the file  operations while filling. | [optional] 
**FileTitle** | Pointer to **NullableString** | The title of the draft, which the portal builds from the name of the person filling it and the name of the  form. Null when the draft the record points at no longer exists. | [optional] 

## Methods

### NewThirdPartyDraftLocation

`func NewThirdPartyDraftLocation() *ThirdPartyDraftLocation`

NewThirdPartyDraftLocation instantiates a new ThirdPartyDraftLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyDraftLocationWithDefaults

`func NewThirdPartyDraftLocationWithDefaults() *ThirdPartyDraftLocation`

NewThirdPartyDraftLocationWithDefaults instantiates a new ThirdPartyDraftLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *ThirdPartyDraftLocation) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ThirdPartyDraftLocation) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ThirdPartyDraftLocation) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ThirdPartyDraftLocation) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ThirdPartyDraftLocation) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ThirdPartyDraftLocation) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetFolderTitle

`func (o *ThirdPartyDraftLocation) GetFolderTitle() string`

GetFolderTitle returns the FolderTitle field if non-nil, zero value otherwise.

### GetFolderTitleOk

`func (o *ThirdPartyDraftLocation) GetFolderTitleOk() (*string, bool)`

GetFolderTitleOk returns a tuple with the FolderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderTitle

`func (o *ThirdPartyDraftLocation) SetFolderTitle(v string)`

SetFolderTitle sets FolderTitle field to given value.

### HasFolderTitle

`func (o *ThirdPartyDraftLocation) HasFolderTitle() bool`

HasFolderTitle returns a boolean if a field has been set.

### SetFolderTitleNil

`func (o *ThirdPartyDraftLocation) SetFolderTitleNil(b bool)`

 SetFolderTitleNil sets the value for FolderTitle to be an explicit nil

### UnsetFolderTitle
`func (o *ThirdPartyDraftLocation) UnsetFolderTitle()`

UnsetFolderTitle ensures that no value is present for FolderTitle, not even an explicit nil
### GetFileId

`func (o *ThirdPartyDraftLocation) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ThirdPartyDraftLocation) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ThirdPartyDraftLocation) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ThirdPartyDraftLocation) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *ThirdPartyDraftLocation) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *ThirdPartyDraftLocation) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetFileTitle

`func (o *ThirdPartyDraftLocation) GetFileTitle() string`

GetFileTitle returns the FileTitle field if non-nil, zero value otherwise.

### GetFileTitleOk

`func (o *ThirdPartyDraftLocation) GetFileTitleOk() (*string, bool)`

GetFileTitleOk returns a tuple with the FileTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTitle

`func (o *ThirdPartyDraftLocation) SetFileTitle(v string)`

SetFileTitle sets FileTitle field to given value.

### HasFileTitle

`func (o *ThirdPartyDraftLocation) HasFileTitle() bool`

HasFileTitle returns a boolean if a field has been set.

### SetFileTitleNil

`func (o *ThirdPartyDraftLocation) SetFileTitleNil(b bool)`

 SetFileTitleNil sets the value for FileTitle to be an explicit nil

### UnsetFileTitle
`func (o *ThirdPartyDraftLocation) UnsetFileTitle()`

UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


