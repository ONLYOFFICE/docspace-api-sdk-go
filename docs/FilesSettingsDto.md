# FilesSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtsImagePreviewed** | Pointer to **[]string** | The list of extensions of the viewed images. | [optional] 
**ExtsMediaPreviewed** | Pointer to **[]string** | The list of extensions of the viewed media files. | [optional] 
**ExtsWebPreviewed** | Pointer to **[]string** | The list of extensions of the viewed files. | [optional] 
**ExtsWebEdited** | Pointer to **[]string** | The list of extensions of the edited files. | [optional] 
**ExtsWebEncrypt** | Pointer to **[]string** | The list of extensions of the encrypted files. | [optional] 
**ExtsWebReviewed** | Pointer to **[]string** | The list of extensions of the reviewed files. | [optional] 
**ExtsWebCustomFilterEditing** | Pointer to **[]string** | The list of extensions of the custom filter files. | [optional] 
**ExtsWebRestrictedEditing** | Pointer to **[]string** | The list of extensions of the files that are restricted for editing. | [optional] 
**ExtsWebCommented** | Pointer to **[]string** | The list of extensions of the commented files. | [optional] 
**ExtsWebTemplate** | Pointer to **[]string** | The list of extensions of the template files. | [optional] 
**ExtsMustConvert** | Pointer to **[]string** | The list of extensions of the files that must be converted. | [optional] 
**ExtsConvertible** | Pointer to **map[string][]string** | The list of the convertible extensions. | [optional] 
**ExtsUploadable** | Pointer to **[]string** | The list of the uploadable extensions. | [optional] 
**ExtsArchive** | Pointer to **[]string** | The list of extensions of the archive files. | [optional] 
**ExtsVideo** | Pointer to **[]string** | The list of the video extensions. | [optional] 
**ExtsAudio** | Pointer to **[]string** | The list of the audio extensions. | [optional] 
**ExtsImage** | Pointer to **[]string** | The list of the image extensions. | [optional] 
**ExtsSpreadsheet** | Pointer to **[]string** | The list of the spreadsheet extensions. | [optional] 
**ExtsPresentation** | Pointer to **[]string** | The list of the presentation extensions. | [optional] 
**ExtsDocument** | Pointer to **[]string** | The list of the text document extensions. | [optional] 
**ExtsDiagram** | Pointer to **[]string** | The list of the diagram extensions. | [optional] 
**InternalFormats** | Pointer to [**NullableFilesSettingsDtoInternalFormats**](FilesSettingsDtoInternalFormats.md) |  | [optional] 
**MasterFormExtension** | Pointer to **NullableString** | The master form extension. | [optional] 
**ParamVersion** | Pointer to **NullableString** | The URL parameter which specifies the file version. | [optional] 
**ParamOutType** | Pointer to **NullableString** | The URL parameter which specifies the output type of the converted file. | [optional] 
**FileDownloadUrlString** | Pointer to **NullableString** | The URL to download a file. | [optional] 
**FileWebViewerUrlString** | Pointer to **NullableString** | The URL to the file web viewer. | [optional] 
**FileWebViewerExternalUrlString** | Pointer to **NullableString** | The external URL to the file web viewer. | [optional] 
**FileWebEditorUrlString** | Pointer to **NullableString** | The URL to the file web editor. | [optional] 
**FileWebEditorExternalUrlString** | Pointer to **NullableString** | The external URL to the file web editor. | [optional] 
**FileRedirectPreviewUrlString** | Pointer to **NullableString** | The redirect URL to the file viewer. | [optional] 
**FileThumbnailUrlString** | Pointer to **NullableString** | The URL to the file thumbnail. | [optional] 
**ConfirmDelete** | Pointer to **bool** | Specifies whether to confirm the file deletion or not. | [optional] 
**EnableThirdParty** | Pointer to **bool** | Specifies whether to allow users to connect the third-party storages. | [optional] 
**ExternalShare** | Pointer to **bool** | Specifies whether to enable sharing external links to the files. | [optional] 
**ExternalShareSocialMedia** | Pointer to **bool** | Specifies whether to enable sharing files on social media. | [optional] 
**StoreOriginalFiles** | Pointer to **bool** | Specifies whether to enable storing original files. | [optional] 
**KeepNewFileName** | Pointer to **bool** | Specifies whether to keep the new file name. | [optional] 
**DisplayFileExtension** | Pointer to **bool** | Specifies whether to display the file extension. | [optional] 
**ConvertNotify** | Pointer to **bool** | Specifies whether to display the conversion notification. | [optional] 
**HideConfirmCancelOperation** | Pointer to **bool** | Specifies whether to hide the confirmation dialog for the cancel operation. | [optional] 
**HideConfirmConvertSave** | Pointer to **bool** | Specifies whether to hide the confirmation dialog  for saving the file copy in the original format when converting a file. | [optional] 
**HideConfirmConvertOpen** | Pointer to **bool** | Specifies whether to hide the confirmation dialog  for opening the conversion result. | [optional] 
**HideConfirmRoomLifetime** | Pointer to **bool** | Specifies whether to hide the confirmation dialog about the file lifetime in the room. | [optional] 
**DefaultOrder** | Pointer to [**OrderBy**](OrderBy.md) | The default order of files. | [optional] 
**Forcesave** | Pointer to **bool** | Specifies whether to forcesave the files or not. | [optional] 
**StoreForcesave** | Pointer to **bool** | Specifies whether to store the forcesaved file versions or not. | [optional] 
**RecentSection** | Pointer to **bool** | Specifies if the Recent section is displayed or not. | [optional] 
**FavoritesSection** | Pointer to **bool** | Specifies if the Favorites section is displayed or not. | [optional] 
**TemplatesSection** | Pointer to **bool** | Specifies if the Templates section is displayed or not. | [optional] 
**DownloadTarGz** | Pointer to **bool** | Specifies whether to download the .tar.gz files or not. | [optional] 
**AutomaticallyCleanUp** | Pointer to [**AutoCleanUpData**](AutoCleanUpData.md) | The auto-clearing setting parameters. | [optional] 
**CanSearchByContent** | Pointer to **bool** | Specifies whether the file can be searched by its content or not. | [optional] 
**DefaultSharingAccessRights** | Pointer to **[]int32** | The default access rights in sharing settings. | [optional] 
**MaxUploadThreadCount** | Pointer to **int32** | The maximum number of upload threads. | [optional] 
**ChunkUploadSize** | Pointer to **int64** | The size of a large file that is uploaded in chunks. | [optional] 
**OpenEditorInSameTab** | Pointer to **bool** | Specifies whether to open the editor in the same tab or not. | [optional] 
**OrganizeRoomsGrouping** | Pointer to **bool** | Specifies whether the grouping of rooms is enabled or not. | [optional] 
**DefaultShareLinkInternal** | Pointer to **bool** | Specifies the default sharing link type: true = DocSpace users only (internal), false = Anyone with the link. | [optional] 
**ExternalShareApplyToDocuments** | Pointer to **bool** | When external sharing is restricted, specifies whether the restriction applies to the My Documents section. | [optional] 
**ExternalShareApplyToRooms** | Pointer to **bool** | When external sharing is restricted, specifies whether the restriction applies to the Rooms section. | [optional] 
**BlockExistingLinksOnRestrict** | Pointer to **bool** | When external sharing is restricted, specifies whether existing public links are blocked immediately. | [optional] 
**ExtsFilesVectorized** | Pointer to **[]string** | List of extensions available for vectorization | [optional] 
**MaxVectorizationFileSize** | Pointer to **int64** | The maximum file size for vectorization | [optional] 

## Methods

### NewFilesSettingsDto

`func NewFilesSettingsDto() *FilesSettingsDto`

NewFilesSettingsDto instantiates a new FilesSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSettingsDtoWithDefaults

`func NewFilesSettingsDtoWithDefaults() *FilesSettingsDto`

NewFilesSettingsDtoWithDefaults instantiates a new FilesSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtsImagePreviewed

`func (o *FilesSettingsDto) GetExtsImagePreviewed() []string`

GetExtsImagePreviewed returns the ExtsImagePreviewed field if non-nil, zero value otherwise.

### GetExtsImagePreviewedOk

`func (o *FilesSettingsDto) GetExtsImagePreviewedOk() (*[]string, bool)`

GetExtsImagePreviewedOk returns a tuple with the ExtsImagePreviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsImagePreviewed

`func (o *FilesSettingsDto) SetExtsImagePreviewed(v []string)`

SetExtsImagePreviewed sets ExtsImagePreviewed field to given value.

### HasExtsImagePreviewed

`func (o *FilesSettingsDto) HasExtsImagePreviewed() bool`

HasExtsImagePreviewed returns a boolean if a field has been set.

### SetExtsImagePreviewedNil

`func (o *FilesSettingsDto) SetExtsImagePreviewedNil(b bool)`

 SetExtsImagePreviewedNil sets the value for ExtsImagePreviewed to be an explicit nil

### UnsetExtsImagePreviewed
`func (o *FilesSettingsDto) UnsetExtsImagePreviewed()`

UnsetExtsImagePreviewed ensures that no value is present for ExtsImagePreviewed, not even an explicit nil
### GetExtsMediaPreviewed

`func (o *FilesSettingsDto) GetExtsMediaPreviewed() []string`

GetExtsMediaPreviewed returns the ExtsMediaPreviewed field if non-nil, zero value otherwise.

### GetExtsMediaPreviewedOk

`func (o *FilesSettingsDto) GetExtsMediaPreviewedOk() (*[]string, bool)`

GetExtsMediaPreviewedOk returns a tuple with the ExtsMediaPreviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsMediaPreviewed

`func (o *FilesSettingsDto) SetExtsMediaPreviewed(v []string)`

SetExtsMediaPreviewed sets ExtsMediaPreviewed field to given value.

### HasExtsMediaPreviewed

`func (o *FilesSettingsDto) HasExtsMediaPreviewed() bool`

HasExtsMediaPreviewed returns a boolean if a field has been set.

### SetExtsMediaPreviewedNil

`func (o *FilesSettingsDto) SetExtsMediaPreviewedNil(b bool)`

 SetExtsMediaPreviewedNil sets the value for ExtsMediaPreviewed to be an explicit nil

### UnsetExtsMediaPreviewed
`func (o *FilesSettingsDto) UnsetExtsMediaPreviewed()`

UnsetExtsMediaPreviewed ensures that no value is present for ExtsMediaPreviewed, not even an explicit nil
### GetExtsWebPreviewed

`func (o *FilesSettingsDto) GetExtsWebPreviewed() []string`

GetExtsWebPreviewed returns the ExtsWebPreviewed field if non-nil, zero value otherwise.

### GetExtsWebPreviewedOk

`func (o *FilesSettingsDto) GetExtsWebPreviewedOk() (*[]string, bool)`

GetExtsWebPreviewedOk returns a tuple with the ExtsWebPreviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebPreviewed

`func (o *FilesSettingsDto) SetExtsWebPreviewed(v []string)`

SetExtsWebPreviewed sets ExtsWebPreviewed field to given value.

### HasExtsWebPreviewed

`func (o *FilesSettingsDto) HasExtsWebPreviewed() bool`

HasExtsWebPreviewed returns a boolean if a field has been set.

### SetExtsWebPreviewedNil

`func (o *FilesSettingsDto) SetExtsWebPreviewedNil(b bool)`

 SetExtsWebPreviewedNil sets the value for ExtsWebPreviewed to be an explicit nil

### UnsetExtsWebPreviewed
`func (o *FilesSettingsDto) UnsetExtsWebPreviewed()`

UnsetExtsWebPreviewed ensures that no value is present for ExtsWebPreviewed, not even an explicit nil
### GetExtsWebEdited

`func (o *FilesSettingsDto) GetExtsWebEdited() []string`

GetExtsWebEdited returns the ExtsWebEdited field if non-nil, zero value otherwise.

### GetExtsWebEditedOk

`func (o *FilesSettingsDto) GetExtsWebEditedOk() (*[]string, bool)`

GetExtsWebEditedOk returns a tuple with the ExtsWebEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebEdited

`func (o *FilesSettingsDto) SetExtsWebEdited(v []string)`

SetExtsWebEdited sets ExtsWebEdited field to given value.

### HasExtsWebEdited

`func (o *FilesSettingsDto) HasExtsWebEdited() bool`

HasExtsWebEdited returns a boolean if a field has been set.

### SetExtsWebEditedNil

`func (o *FilesSettingsDto) SetExtsWebEditedNil(b bool)`

 SetExtsWebEditedNil sets the value for ExtsWebEdited to be an explicit nil

### UnsetExtsWebEdited
`func (o *FilesSettingsDto) UnsetExtsWebEdited()`

UnsetExtsWebEdited ensures that no value is present for ExtsWebEdited, not even an explicit nil
### GetExtsWebEncrypt

`func (o *FilesSettingsDto) GetExtsWebEncrypt() []string`

GetExtsWebEncrypt returns the ExtsWebEncrypt field if non-nil, zero value otherwise.

### GetExtsWebEncryptOk

`func (o *FilesSettingsDto) GetExtsWebEncryptOk() (*[]string, bool)`

GetExtsWebEncryptOk returns a tuple with the ExtsWebEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebEncrypt

`func (o *FilesSettingsDto) SetExtsWebEncrypt(v []string)`

SetExtsWebEncrypt sets ExtsWebEncrypt field to given value.

### HasExtsWebEncrypt

`func (o *FilesSettingsDto) HasExtsWebEncrypt() bool`

HasExtsWebEncrypt returns a boolean if a field has been set.

### SetExtsWebEncryptNil

`func (o *FilesSettingsDto) SetExtsWebEncryptNil(b bool)`

 SetExtsWebEncryptNil sets the value for ExtsWebEncrypt to be an explicit nil

### UnsetExtsWebEncrypt
`func (o *FilesSettingsDto) UnsetExtsWebEncrypt()`

UnsetExtsWebEncrypt ensures that no value is present for ExtsWebEncrypt, not even an explicit nil
### GetExtsWebReviewed

`func (o *FilesSettingsDto) GetExtsWebReviewed() []string`

GetExtsWebReviewed returns the ExtsWebReviewed field if non-nil, zero value otherwise.

### GetExtsWebReviewedOk

`func (o *FilesSettingsDto) GetExtsWebReviewedOk() (*[]string, bool)`

GetExtsWebReviewedOk returns a tuple with the ExtsWebReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebReviewed

`func (o *FilesSettingsDto) SetExtsWebReviewed(v []string)`

SetExtsWebReviewed sets ExtsWebReviewed field to given value.

### HasExtsWebReviewed

`func (o *FilesSettingsDto) HasExtsWebReviewed() bool`

HasExtsWebReviewed returns a boolean if a field has been set.

### SetExtsWebReviewedNil

`func (o *FilesSettingsDto) SetExtsWebReviewedNil(b bool)`

 SetExtsWebReviewedNil sets the value for ExtsWebReviewed to be an explicit nil

### UnsetExtsWebReviewed
`func (o *FilesSettingsDto) UnsetExtsWebReviewed()`

UnsetExtsWebReviewed ensures that no value is present for ExtsWebReviewed, not even an explicit nil
### GetExtsWebCustomFilterEditing

`func (o *FilesSettingsDto) GetExtsWebCustomFilterEditing() []string`

GetExtsWebCustomFilterEditing returns the ExtsWebCustomFilterEditing field if non-nil, zero value otherwise.

### GetExtsWebCustomFilterEditingOk

`func (o *FilesSettingsDto) GetExtsWebCustomFilterEditingOk() (*[]string, bool)`

GetExtsWebCustomFilterEditingOk returns a tuple with the ExtsWebCustomFilterEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebCustomFilterEditing

`func (o *FilesSettingsDto) SetExtsWebCustomFilterEditing(v []string)`

SetExtsWebCustomFilterEditing sets ExtsWebCustomFilterEditing field to given value.

### HasExtsWebCustomFilterEditing

`func (o *FilesSettingsDto) HasExtsWebCustomFilterEditing() bool`

HasExtsWebCustomFilterEditing returns a boolean if a field has been set.

### SetExtsWebCustomFilterEditingNil

`func (o *FilesSettingsDto) SetExtsWebCustomFilterEditingNil(b bool)`

 SetExtsWebCustomFilterEditingNil sets the value for ExtsWebCustomFilterEditing to be an explicit nil

### UnsetExtsWebCustomFilterEditing
`func (o *FilesSettingsDto) UnsetExtsWebCustomFilterEditing()`

UnsetExtsWebCustomFilterEditing ensures that no value is present for ExtsWebCustomFilterEditing, not even an explicit nil
### GetExtsWebRestrictedEditing

`func (o *FilesSettingsDto) GetExtsWebRestrictedEditing() []string`

GetExtsWebRestrictedEditing returns the ExtsWebRestrictedEditing field if non-nil, zero value otherwise.

### GetExtsWebRestrictedEditingOk

`func (o *FilesSettingsDto) GetExtsWebRestrictedEditingOk() (*[]string, bool)`

GetExtsWebRestrictedEditingOk returns a tuple with the ExtsWebRestrictedEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebRestrictedEditing

`func (o *FilesSettingsDto) SetExtsWebRestrictedEditing(v []string)`

SetExtsWebRestrictedEditing sets ExtsWebRestrictedEditing field to given value.

### HasExtsWebRestrictedEditing

`func (o *FilesSettingsDto) HasExtsWebRestrictedEditing() bool`

HasExtsWebRestrictedEditing returns a boolean if a field has been set.

### SetExtsWebRestrictedEditingNil

`func (o *FilesSettingsDto) SetExtsWebRestrictedEditingNil(b bool)`

 SetExtsWebRestrictedEditingNil sets the value for ExtsWebRestrictedEditing to be an explicit nil

### UnsetExtsWebRestrictedEditing
`func (o *FilesSettingsDto) UnsetExtsWebRestrictedEditing()`

UnsetExtsWebRestrictedEditing ensures that no value is present for ExtsWebRestrictedEditing, not even an explicit nil
### GetExtsWebCommented

`func (o *FilesSettingsDto) GetExtsWebCommented() []string`

GetExtsWebCommented returns the ExtsWebCommented field if non-nil, zero value otherwise.

### GetExtsWebCommentedOk

`func (o *FilesSettingsDto) GetExtsWebCommentedOk() (*[]string, bool)`

GetExtsWebCommentedOk returns a tuple with the ExtsWebCommented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebCommented

`func (o *FilesSettingsDto) SetExtsWebCommented(v []string)`

SetExtsWebCommented sets ExtsWebCommented field to given value.

### HasExtsWebCommented

`func (o *FilesSettingsDto) HasExtsWebCommented() bool`

HasExtsWebCommented returns a boolean if a field has been set.

### SetExtsWebCommentedNil

`func (o *FilesSettingsDto) SetExtsWebCommentedNil(b bool)`

 SetExtsWebCommentedNil sets the value for ExtsWebCommented to be an explicit nil

### UnsetExtsWebCommented
`func (o *FilesSettingsDto) UnsetExtsWebCommented()`

UnsetExtsWebCommented ensures that no value is present for ExtsWebCommented, not even an explicit nil
### GetExtsWebTemplate

`func (o *FilesSettingsDto) GetExtsWebTemplate() []string`

GetExtsWebTemplate returns the ExtsWebTemplate field if non-nil, zero value otherwise.

### GetExtsWebTemplateOk

`func (o *FilesSettingsDto) GetExtsWebTemplateOk() (*[]string, bool)`

GetExtsWebTemplateOk returns a tuple with the ExtsWebTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsWebTemplate

`func (o *FilesSettingsDto) SetExtsWebTemplate(v []string)`

SetExtsWebTemplate sets ExtsWebTemplate field to given value.

### HasExtsWebTemplate

`func (o *FilesSettingsDto) HasExtsWebTemplate() bool`

HasExtsWebTemplate returns a boolean if a field has been set.

### SetExtsWebTemplateNil

`func (o *FilesSettingsDto) SetExtsWebTemplateNil(b bool)`

 SetExtsWebTemplateNil sets the value for ExtsWebTemplate to be an explicit nil

### UnsetExtsWebTemplate
`func (o *FilesSettingsDto) UnsetExtsWebTemplate()`

UnsetExtsWebTemplate ensures that no value is present for ExtsWebTemplate, not even an explicit nil
### GetExtsMustConvert

`func (o *FilesSettingsDto) GetExtsMustConvert() []string`

GetExtsMustConvert returns the ExtsMustConvert field if non-nil, zero value otherwise.

### GetExtsMustConvertOk

`func (o *FilesSettingsDto) GetExtsMustConvertOk() (*[]string, bool)`

GetExtsMustConvertOk returns a tuple with the ExtsMustConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsMustConvert

`func (o *FilesSettingsDto) SetExtsMustConvert(v []string)`

SetExtsMustConvert sets ExtsMustConvert field to given value.

### HasExtsMustConvert

`func (o *FilesSettingsDto) HasExtsMustConvert() bool`

HasExtsMustConvert returns a boolean if a field has been set.

### SetExtsMustConvertNil

`func (o *FilesSettingsDto) SetExtsMustConvertNil(b bool)`

 SetExtsMustConvertNil sets the value for ExtsMustConvert to be an explicit nil

### UnsetExtsMustConvert
`func (o *FilesSettingsDto) UnsetExtsMustConvert()`

UnsetExtsMustConvert ensures that no value is present for ExtsMustConvert, not even an explicit nil
### GetExtsConvertible

`func (o *FilesSettingsDto) GetExtsConvertible() map[string][]string`

GetExtsConvertible returns the ExtsConvertible field if non-nil, zero value otherwise.

### GetExtsConvertibleOk

`func (o *FilesSettingsDto) GetExtsConvertibleOk() (*map[string][]string, bool)`

GetExtsConvertibleOk returns a tuple with the ExtsConvertible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsConvertible

`func (o *FilesSettingsDto) SetExtsConvertible(v map[string][]string)`

SetExtsConvertible sets ExtsConvertible field to given value.

### HasExtsConvertible

`func (o *FilesSettingsDto) HasExtsConvertible() bool`

HasExtsConvertible returns a boolean if a field has been set.

### GetExtsUploadable

`func (o *FilesSettingsDto) GetExtsUploadable() []string`

GetExtsUploadable returns the ExtsUploadable field if non-nil, zero value otherwise.

### GetExtsUploadableOk

`func (o *FilesSettingsDto) GetExtsUploadableOk() (*[]string, bool)`

GetExtsUploadableOk returns a tuple with the ExtsUploadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsUploadable

`func (o *FilesSettingsDto) SetExtsUploadable(v []string)`

SetExtsUploadable sets ExtsUploadable field to given value.

### HasExtsUploadable

`func (o *FilesSettingsDto) HasExtsUploadable() bool`

HasExtsUploadable returns a boolean if a field has been set.

### SetExtsUploadableNil

`func (o *FilesSettingsDto) SetExtsUploadableNil(b bool)`

 SetExtsUploadableNil sets the value for ExtsUploadable to be an explicit nil

### UnsetExtsUploadable
`func (o *FilesSettingsDto) UnsetExtsUploadable()`

UnsetExtsUploadable ensures that no value is present for ExtsUploadable, not even an explicit nil
### GetExtsArchive

`func (o *FilesSettingsDto) GetExtsArchive() []string`

GetExtsArchive returns the ExtsArchive field if non-nil, zero value otherwise.

### GetExtsArchiveOk

`func (o *FilesSettingsDto) GetExtsArchiveOk() (*[]string, bool)`

GetExtsArchiveOk returns a tuple with the ExtsArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsArchive

`func (o *FilesSettingsDto) SetExtsArchive(v []string)`

SetExtsArchive sets ExtsArchive field to given value.

### HasExtsArchive

`func (o *FilesSettingsDto) HasExtsArchive() bool`

HasExtsArchive returns a boolean if a field has been set.

### SetExtsArchiveNil

`func (o *FilesSettingsDto) SetExtsArchiveNil(b bool)`

 SetExtsArchiveNil sets the value for ExtsArchive to be an explicit nil

### UnsetExtsArchive
`func (o *FilesSettingsDto) UnsetExtsArchive()`

UnsetExtsArchive ensures that no value is present for ExtsArchive, not even an explicit nil
### GetExtsVideo

`func (o *FilesSettingsDto) GetExtsVideo() []string`

GetExtsVideo returns the ExtsVideo field if non-nil, zero value otherwise.

### GetExtsVideoOk

`func (o *FilesSettingsDto) GetExtsVideoOk() (*[]string, bool)`

GetExtsVideoOk returns a tuple with the ExtsVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsVideo

`func (o *FilesSettingsDto) SetExtsVideo(v []string)`

SetExtsVideo sets ExtsVideo field to given value.

### HasExtsVideo

`func (o *FilesSettingsDto) HasExtsVideo() bool`

HasExtsVideo returns a boolean if a field has been set.

### SetExtsVideoNil

`func (o *FilesSettingsDto) SetExtsVideoNil(b bool)`

 SetExtsVideoNil sets the value for ExtsVideo to be an explicit nil

### UnsetExtsVideo
`func (o *FilesSettingsDto) UnsetExtsVideo()`

UnsetExtsVideo ensures that no value is present for ExtsVideo, not even an explicit nil
### GetExtsAudio

`func (o *FilesSettingsDto) GetExtsAudio() []string`

GetExtsAudio returns the ExtsAudio field if non-nil, zero value otherwise.

### GetExtsAudioOk

`func (o *FilesSettingsDto) GetExtsAudioOk() (*[]string, bool)`

GetExtsAudioOk returns a tuple with the ExtsAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsAudio

`func (o *FilesSettingsDto) SetExtsAudio(v []string)`

SetExtsAudio sets ExtsAudio field to given value.

### HasExtsAudio

`func (o *FilesSettingsDto) HasExtsAudio() bool`

HasExtsAudio returns a boolean if a field has been set.

### SetExtsAudioNil

`func (o *FilesSettingsDto) SetExtsAudioNil(b bool)`

 SetExtsAudioNil sets the value for ExtsAudio to be an explicit nil

### UnsetExtsAudio
`func (o *FilesSettingsDto) UnsetExtsAudio()`

UnsetExtsAudio ensures that no value is present for ExtsAudio, not even an explicit nil
### GetExtsImage

`func (o *FilesSettingsDto) GetExtsImage() []string`

GetExtsImage returns the ExtsImage field if non-nil, zero value otherwise.

### GetExtsImageOk

`func (o *FilesSettingsDto) GetExtsImageOk() (*[]string, bool)`

GetExtsImageOk returns a tuple with the ExtsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsImage

`func (o *FilesSettingsDto) SetExtsImage(v []string)`

SetExtsImage sets ExtsImage field to given value.

### HasExtsImage

`func (o *FilesSettingsDto) HasExtsImage() bool`

HasExtsImage returns a boolean if a field has been set.

### SetExtsImageNil

`func (o *FilesSettingsDto) SetExtsImageNil(b bool)`

 SetExtsImageNil sets the value for ExtsImage to be an explicit nil

### UnsetExtsImage
`func (o *FilesSettingsDto) UnsetExtsImage()`

UnsetExtsImage ensures that no value is present for ExtsImage, not even an explicit nil
### GetExtsSpreadsheet

`func (o *FilesSettingsDto) GetExtsSpreadsheet() []string`

GetExtsSpreadsheet returns the ExtsSpreadsheet field if non-nil, zero value otherwise.

### GetExtsSpreadsheetOk

`func (o *FilesSettingsDto) GetExtsSpreadsheetOk() (*[]string, bool)`

GetExtsSpreadsheetOk returns a tuple with the ExtsSpreadsheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsSpreadsheet

`func (o *FilesSettingsDto) SetExtsSpreadsheet(v []string)`

SetExtsSpreadsheet sets ExtsSpreadsheet field to given value.

### HasExtsSpreadsheet

`func (o *FilesSettingsDto) HasExtsSpreadsheet() bool`

HasExtsSpreadsheet returns a boolean if a field has been set.

### SetExtsSpreadsheetNil

`func (o *FilesSettingsDto) SetExtsSpreadsheetNil(b bool)`

 SetExtsSpreadsheetNil sets the value for ExtsSpreadsheet to be an explicit nil

### UnsetExtsSpreadsheet
`func (o *FilesSettingsDto) UnsetExtsSpreadsheet()`

UnsetExtsSpreadsheet ensures that no value is present for ExtsSpreadsheet, not even an explicit nil
### GetExtsPresentation

`func (o *FilesSettingsDto) GetExtsPresentation() []string`

GetExtsPresentation returns the ExtsPresentation field if non-nil, zero value otherwise.

### GetExtsPresentationOk

`func (o *FilesSettingsDto) GetExtsPresentationOk() (*[]string, bool)`

GetExtsPresentationOk returns a tuple with the ExtsPresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsPresentation

`func (o *FilesSettingsDto) SetExtsPresentation(v []string)`

SetExtsPresentation sets ExtsPresentation field to given value.

### HasExtsPresentation

`func (o *FilesSettingsDto) HasExtsPresentation() bool`

HasExtsPresentation returns a boolean if a field has been set.

### SetExtsPresentationNil

`func (o *FilesSettingsDto) SetExtsPresentationNil(b bool)`

 SetExtsPresentationNil sets the value for ExtsPresentation to be an explicit nil

### UnsetExtsPresentation
`func (o *FilesSettingsDto) UnsetExtsPresentation()`

UnsetExtsPresentation ensures that no value is present for ExtsPresentation, not even an explicit nil
### GetExtsDocument

`func (o *FilesSettingsDto) GetExtsDocument() []string`

GetExtsDocument returns the ExtsDocument field if non-nil, zero value otherwise.

### GetExtsDocumentOk

`func (o *FilesSettingsDto) GetExtsDocumentOk() (*[]string, bool)`

GetExtsDocumentOk returns a tuple with the ExtsDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsDocument

`func (o *FilesSettingsDto) SetExtsDocument(v []string)`

SetExtsDocument sets ExtsDocument field to given value.

### HasExtsDocument

`func (o *FilesSettingsDto) HasExtsDocument() bool`

HasExtsDocument returns a boolean if a field has been set.

### SetExtsDocumentNil

`func (o *FilesSettingsDto) SetExtsDocumentNil(b bool)`

 SetExtsDocumentNil sets the value for ExtsDocument to be an explicit nil

### UnsetExtsDocument
`func (o *FilesSettingsDto) UnsetExtsDocument()`

UnsetExtsDocument ensures that no value is present for ExtsDocument, not even an explicit nil
### GetExtsDiagram

`func (o *FilesSettingsDto) GetExtsDiagram() []string`

GetExtsDiagram returns the ExtsDiagram field if non-nil, zero value otherwise.

### GetExtsDiagramOk

`func (o *FilesSettingsDto) GetExtsDiagramOk() (*[]string, bool)`

GetExtsDiagramOk returns a tuple with the ExtsDiagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsDiagram

`func (o *FilesSettingsDto) SetExtsDiagram(v []string)`

SetExtsDiagram sets ExtsDiagram field to given value.

### HasExtsDiagram

`func (o *FilesSettingsDto) HasExtsDiagram() bool`

HasExtsDiagram returns a boolean if a field has been set.

### SetExtsDiagramNil

`func (o *FilesSettingsDto) SetExtsDiagramNil(b bool)`

 SetExtsDiagramNil sets the value for ExtsDiagram to be an explicit nil

### UnsetExtsDiagram
`func (o *FilesSettingsDto) UnsetExtsDiagram()`

UnsetExtsDiagram ensures that no value is present for ExtsDiagram, not even an explicit nil
### GetInternalFormats

`func (o *FilesSettingsDto) GetInternalFormats() FilesSettingsDtoInternalFormats`

GetInternalFormats returns the InternalFormats field if non-nil, zero value otherwise.

### GetInternalFormatsOk

`func (o *FilesSettingsDto) GetInternalFormatsOk() (*FilesSettingsDtoInternalFormats, bool)`

GetInternalFormatsOk returns a tuple with the InternalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFormats

`func (o *FilesSettingsDto) SetInternalFormats(v FilesSettingsDtoInternalFormats)`

SetInternalFormats sets InternalFormats field to given value.

### HasInternalFormats

`func (o *FilesSettingsDto) HasInternalFormats() bool`

HasInternalFormats returns a boolean if a field has been set.

### SetInternalFormatsNil

`func (o *FilesSettingsDto) SetInternalFormatsNil(b bool)`

 SetInternalFormatsNil sets the value for InternalFormats to be an explicit nil

### UnsetInternalFormats
`func (o *FilesSettingsDto) UnsetInternalFormats()`

UnsetInternalFormats ensures that no value is present for InternalFormats, not even an explicit nil
### GetMasterFormExtension

`func (o *FilesSettingsDto) GetMasterFormExtension() string`

GetMasterFormExtension returns the MasterFormExtension field if non-nil, zero value otherwise.

### GetMasterFormExtensionOk

`func (o *FilesSettingsDto) GetMasterFormExtensionOk() (*string, bool)`

GetMasterFormExtensionOk returns a tuple with the MasterFormExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterFormExtension

`func (o *FilesSettingsDto) SetMasterFormExtension(v string)`

SetMasterFormExtension sets MasterFormExtension field to given value.

### HasMasterFormExtension

`func (o *FilesSettingsDto) HasMasterFormExtension() bool`

HasMasterFormExtension returns a boolean if a field has been set.

### SetMasterFormExtensionNil

`func (o *FilesSettingsDto) SetMasterFormExtensionNil(b bool)`

 SetMasterFormExtensionNil sets the value for MasterFormExtension to be an explicit nil

### UnsetMasterFormExtension
`func (o *FilesSettingsDto) UnsetMasterFormExtension()`

UnsetMasterFormExtension ensures that no value is present for MasterFormExtension, not even an explicit nil
### GetParamVersion

`func (o *FilesSettingsDto) GetParamVersion() string`

GetParamVersion returns the ParamVersion field if non-nil, zero value otherwise.

### GetParamVersionOk

`func (o *FilesSettingsDto) GetParamVersionOk() (*string, bool)`

GetParamVersionOk returns a tuple with the ParamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamVersion

`func (o *FilesSettingsDto) SetParamVersion(v string)`

SetParamVersion sets ParamVersion field to given value.

### HasParamVersion

`func (o *FilesSettingsDto) HasParamVersion() bool`

HasParamVersion returns a boolean if a field has been set.

### SetParamVersionNil

`func (o *FilesSettingsDto) SetParamVersionNil(b bool)`

 SetParamVersionNil sets the value for ParamVersion to be an explicit nil

### UnsetParamVersion
`func (o *FilesSettingsDto) UnsetParamVersion()`

UnsetParamVersion ensures that no value is present for ParamVersion, not even an explicit nil
### GetParamOutType

`func (o *FilesSettingsDto) GetParamOutType() string`

GetParamOutType returns the ParamOutType field if non-nil, zero value otherwise.

### GetParamOutTypeOk

`func (o *FilesSettingsDto) GetParamOutTypeOk() (*string, bool)`

GetParamOutTypeOk returns a tuple with the ParamOutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOutType

`func (o *FilesSettingsDto) SetParamOutType(v string)`

SetParamOutType sets ParamOutType field to given value.

### HasParamOutType

`func (o *FilesSettingsDto) HasParamOutType() bool`

HasParamOutType returns a boolean if a field has been set.

### SetParamOutTypeNil

`func (o *FilesSettingsDto) SetParamOutTypeNil(b bool)`

 SetParamOutTypeNil sets the value for ParamOutType to be an explicit nil

### UnsetParamOutType
`func (o *FilesSettingsDto) UnsetParamOutType()`

UnsetParamOutType ensures that no value is present for ParamOutType, not even an explicit nil
### GetFileDownloadUrlString

`func (o *FilesSettingsDto) GetFileDownloadUrlString() string`

GetFileDownloadUrlString returns the FileDownloadUrlString field if non-nil, zero value otherwise.

### GetFileDownloadUrlStringOk

`func (o *FilesSettingsDto) GetFileDownloadUrlStringOk() (*string, bool)`

GetFileDownloadUrlStringOk returns a tuple with the FileDownloadUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadUrlString

`func (o *FilesSettingsDto) SetFileDownloadUrlString(v string)`

SetFileDownloadUrlString sets FileDownloadUrlString field to given value.

### HasFileDownloadUrlString

`func (o *FilesSettingsDto) HasFileDownloadUrlString() bool`

HasFileDownloadUrlString returns a boolean if a field has been set.

### SetFileDownloadUrlStringNil

`func (o *FilesSettingsDto) SetFileDownloadUrlStringNil(b bool)`

 SetFileDownloadUrlStringNil sets the value for FileDownloadUrlString to be an explicit nil

### UnsetFileDownloadUrlString
`func (o *FilesSettingsDto) UnsetFileDownloadUrlString()`

UnsetFileDownloadUrlString ensures that no value is present for FileDownloadUrlString, not even an explicit nil
### GetFileWebViewerUrlString

`func (o *FilesSettingsDto) GetFileWebViewerUrlString() string`

GetFileWebViewerUrlString returns the FileWebViewerUrlString field if non-nil, zero value otherwise.

### GetFileWebViewerUrlStringOk

`func (o *FilesSettingsDto) GetFileWebViewerUrlStringOk() (*string, bool)`

GetFileWebViewerUrlStringOk returns a tuple with the FileWebViewerUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWebViewerUrlString

`func (o *FilesSettingsDto) SetFileWebViewerUrlString(v string)`

SetFileWebViewerUrlString sets FileWebViewerUrlString field to given value.

### HasFileWebViewerUrlString

`func (o *FilesSettingsDto) HasFileWebViewerUrlString() bool`

HasFileWebViewerUrlString returns a boolean if a field has been set.

### SetFileWebViewerUrlStringNil

`func (o *FilesSettingsDto) SetFileWebViewerUrlStringNil(b bool)`

 SetFileWebViewerUrlStringNil sets the value for FileWebViewerUrlString to be an explicit nil

### UnsetFileWebViewerUrlString
`func (o *FilesSettingsDto) UnsetFileWebViewerUrlString()`

UnsetFileWebViewerUrlString ensures that no value is present for FileWebViewerUrlString, not even an explicit nil
### GetFileWebViewerExternalUrlString

`func (o *FilesSettingsDto) GetFileWebViewerExternalUrlString() string`

GetFileWebViewerExternalUrlString returns the FileWebViewerExternalUrlString field if non-nil, zero value otherwise.

### GetFileWebViewerExternalUrlStringOk

`func (o *FilesSettingsDto) GetFileWebViewerExternalUrlStringOk() (*string, bool)`

GetFileWebViewerExternalUrlStringOk returns a tuple with the FileWebViewerExternalUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWebViewerExternalUrlString

`func (o *FilesSettingsDto) SetFileWebViewerExternalUrlString(v string)`

SetFileWebViewerExternalUrlString sets FileWebViewerExternalUrlString field to given value.

### HasFileWebViewerExternalUrlString

`func (o *FilesSettingsDto) HasFileWebViewerExternalUrlString() bool`

HasFileWebViewerExternalUrlString returns a boolean if a field has been set.

### SetFileWebViewerExternalUrlStringNil

`func (o *FilesSettingsDto) SetFileWebViewerExternalUrlStringNil(b bool)`

 SetFileWebViewerExternalUrlStringNil sets the value for FileWebViewerExternalUrlString to be an explicit nil

### UnsetFileWebViewerExternalUrlString
`func (o *FilesSettingsDto) UnsetFileWebViewerExternalUrlString()`

UnsetFileWebViewerExternalUrlString ensures that no value is present for FileWebViewerExternalUrlString, not even an explicit nil
### GetFileWebEditorUrlString

`func (o *FilesSettingsDto) GetFileWebEditorUrlString() string`

GetFileWebEditorUrlString returns the FileWebEditorUrlString field if non-nil, zero value otherwise.

### GetFileWebEditorUrlStringOk

`func (o *FilesSettingsDto) GetFileWebEditorUrlStringOk() (*string, bool)`

GetFileWebEditorUrlStringOk returns a tuple with the FileWebEditorUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWebEditorUrlString

`func (o *FilesSettingsDto) SetFileWebEditorUrlString(v string)`

SetFileWebEditorUrlString sets FileWebEditorUrlString field to given value.

### HasFileWebEditorUrlString

`func (o *FilesSettingsDto) HasFileWebEditorUrlString() bool`

HasFileWebEditorUrlString returns a boolean if a field has been set.

### SetFileWebEditorUrlStringNil

`func (o *FilesSettingsDto) SetFileWebEditorUrlStringNil(b bool)`

 SetFileWebEditorUrlStringNil sets the value for FileWebEditorUrlString to be an explicit nil

### UnsetFileWebEditorUrlString
`func (o *FilesSettingsDto) UnsetFileWebEditorUrlString()`

UnsetFileWebEditorUrlString ensures that no value is present for FileWebEditorUrlString, not even an explicit nil
### GetFileWebEditorExternalUrlString

`func (o *FilesSettingsDto) GetFileWebEditorExternalUrlString() string`

GetFileWebEditorExternalUrlString returns the FileWebEditorExternalUrlString field if non-nil, zero value otherwise.

### GetFileWebEditorExternalUrlStringOk

`func (o *FilesSettingsDto) GetFileWebEditorExternalUrlStringOk() (*string, bool)`

GetFileWebEditorExternalUrlStringOk returns a tuple with the FileWebEditorExternalUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWebEditorExternalUrlString

`func (o *FilesSettingsDto) SetFileWebEditorExternalUrlString(v string)`

SetFileWebEditorExternalUrlString sets FileWebEditorExternalUrlString field to given value.

### HasFileWebEditorExternalUrlString

`func (o *FilesSettingsDto) HasFileWebEditorExternalUrlString() bool`

HasFileWebEditorExternalUrlString returns a boolean if a field has been set.

### SetFileWebEditorExternalUrlStringNil

`func (o *FilesSettingsDto) SetFileWebEditorExternalUrlStringNil(b bool)`

 SetFileWebEditorExternalUrlStringNil sets the value for FileWebEditorExternalUrlString to be an explicit nil

### UnsetFileWebEditorExternalUrlString
`func (o *FilesSettingsDto) UnsetFileWebEditorExternalUrlString()`

UnsetFileWebEditorExternalUrlString ensures that no value is present for FileWebEditorExternalUrlString, not even an explicit nil
### GetFileRedirectPreviewUrlString

`func (o *FilesSettingsDto) GetFileRedirectPreviewUrlString() string`

GetFileRedirectPreviewUrlString returns the FileRedirectPreviewUrlString field if non-nil, zero value otherwise.

### GetFileRedirectPreviewUrlStringOk

`func (o *FilesSettingsDto) GetFileRedirectPreviewUrlStringOk() (*string, bool)`

GetFileRedirectPreviewUrlStringOk returns a tuple with the FileRedirectPreviewUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRedirectPreviewUrlString

`func (o *FilesSettingsDto) SetFileRedirectPreviewUrlString(v string)`

SetFileRedirectPreviewUrlString sets FileRedirectPreviewUrlString field to given value.

### HasFileRedirectPreviewUrlString

`func (o *FilesSettingsDto) HasFileRedirectPreviewUrlString() bool`

HasFileRedirectPreviewUrlString returns a boolean if a field has been set.

### SetFileRedirectPreviewUrlStringNil

`func (o *FilesSettingsDto) SetFileRedirectPreviewUrlStringNil(b bool)`

 SetFileRedirectPreviewUrlStringNil sets the value for FileRedirectPreviewUrlString to be an explicit nil

### UnsetFileRedirectPreviewUrlString
`func (o *FilesSettingsDto) UnsetFileRedirectPreviewUrlString()`

UnsetFileRedirectPreviewUrlString ensures that no value is present for FileRedirectPreviewUrlString, not even an explicit nil
### GetFileThumbnailUrlString

`func (o *FilesSettingsDto) GetFileThumbnailUrlString() string`

GetFileThumbnailUrlString returns the FileThumbnailUrlString field if non-nil, zero value otherwise.

### GetFileThumbnailUrlStringOk

`func (o *FilesSettingsDto) GetFileThumbnailUrlStringOk() (*string, bool)`

GetFileThumbnailUrlStringOk returns a tuple with the FileThumbnailUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileThumbnailUrlString

`func (o *FilesSettingsDto) SetFileThumbnailUrlString(v string)`

SetFileThumbnailUrlString sets FileThumbnailUrlString field to given value.

### HasFileThumbnailUrlString

`func (o *FilesSettingsDto) HasFileThumbnailUrlString() bool`

HasFileThumbnailUrlString returns a boolean if a field has been set.

### SetFileThumbnailUrlStringNil

`func (o *FilesSettingsDto) SetFileThumbnailUrlStringNil(b bool)`

 SetFileThumbnailUrlStringNil sets the value for FileThumbnailUrlString to be an explicit nil

### UnsetFileThumbnailUrlString
`func (o *FilesSettingsDto) UnsetFileThumbnailUrlString()`

UnsetFileThumbnailUrlString ensures that no value is present for FileThumbnailUrlString, not even an explicit nil
### GetConfirmDelete

`func (o *FilesSettingsDto) GetConfirmDelete() bool`

GetConfirmDelete returns the ConfirmDelete field if non-nil, zero value otherwise.

### GetConfirmDeleteOk

`func (o *FilesSettingsDto) GetConfirmDeleteOk() (*bool, bool)`

GetConfirmDeleteOk returns a tuple with the ConfirmDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmDelete

`func (o *FilesSettingsDto) SetConfirmDelete(v bool)`

SetConfirmDelete sets ConfirmDelete field to given value.

### HasConfirmDelete

`func (o *FilesSettingsDto) HasConfirmDelete() bool`

HasConfirmDelete returns a boolean if a field has been set.

### GetEnableThirdParty

`func (o *FilesSettingsDto) GetEnableThirdParty() bool`

GetEnableThirdParty returns the EnableThirdParty field if non-nil, zero value otherwise.

### GetEnableThirdPartyOk

`func (o *FilesSettingsDto) GetEnableThirdPartyOk() (*bool, bool)`

GetEnableThirdPartyOk returns a tuple with the EnableThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThirdParty

`func (o *FilesSettingsDto) SetEnableThirdParty(v bool)`

SetEnableThirdParty sets EnableThirdParty field to given value.

### HasEnableThirdParty

`func (o *FilesSettingsDto) HasEnableThirdParty() bool`

HasEnableThirdParty returns a boolean if a field has been set.

### GetExternalShare

`func (o *FilesSettingsDto) GetExternalShare() bool`

GetExternalShare returns the ExternalShare field if non-nil, zero value otherwise.

### GetExternalShareOk

`func (o *FilesSettingsDto) GetExternalShareOk() (*bool, bool)`

GetExternalShareOk returns a tuple with the ExternalShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShare

`func (o *FilesSettingsDto) SetExternalShare(v bool)`

SetExternalShare sets ExternalShare field to given value.

### HasExternalShare

`func (o *FilesSettingsDto) HasExternalShare() bool`

HasExternalShare returns a boolean if a field has been set.

### GetExternalShareSocialMedia

`func (o *FilesSettingsDto) GetExternalShareSocialMedia() bool`

GetExternalShareSocialMedia returns the ExternalShareSocialMedia field if non-nil, zero value otherwise.

### GetExternalShareSocialMediaOk

`func (o *FilesSettingsDto) GetExternalShareSocialMediaOk() (*bool, bool)`

GetExternalShareSocialMediaOk returns a tuple with the ExternalShareSocialMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareSocialMedia

`func (o *FilesSettingsDto) SetExternalShareSocialMedia(v bool)`

SetExternalShareSocialMedia sets ExternalShareSocialMedia field to given value.

### HasExternalShareSocialMedia

`func (o *FilesSettingsDto) HasExternalShareSocialMedia() bool`

HasExternalShareSocialMedia returns a boolean if a field has been set.

### GetStoreOriginalFiles

`func (o *FilesSettingsDto) GetStoreOriginalFiles() bool`

GetStoreOriginalFiles returns the StoreOriginalFiles field if non-nil, zero value otherwise.

### GetStoreOriginalFilesOk

`func (o *FilesSettingsDto) GetStoreOriginalFilesOk() (*bool, bool)`

GetStoreOriginalFilesOk returns a tuple with the StoreOriginalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOriginalFiles

`func (o *FilesSettingsDto) SetStoreOriginalFiles(v bool)`

SetStoreOriginalFiles sets StoreOriginalFiles field to given value.

### HasStoreOriginalFiles

`func (o *FilesSettingsDto) HasStoreOriginalFiles() bool`

HasStoreOriginalFiles returns a boolean if a field has been set.

### GetKeepNewFileName

`func (o *FilesSettingsDto) GetKeepNewFileName() bool`

GetKeepNewFileName returns the KeepNewFileName field if non-nil, zero value otherwise.

### GetKeepNewFileNameOk

`func (o *FilesSettingsDto) GetKeepNewFileNameOk() (*bool, bool)`

GetKeepNewFileNameOk returns a tuple with the KeepNewFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepNewFileName

`func (o *FilesSettingsDto) SetKeepNewFileName(v bool)`

SetKeepNewFileName sets KeepNewFileName field to given value.

### HasKeepNewFileName

`func (o *FilesSettingsDto) HasKeepNewFileName() bool`

HasKeepNewFileName returns a boolean if a field has been set.

### GetDisplayFileExtension

`func (o *FilesSettingsDto) GetDisplayFileExtension() bool`

GetDisplayFileExtension returns the DisplayFileExtension field if non-nil, zero value otherwise.

### GetDisplayFileExtensionOk

`func (o *FilesSettingsDto) GetDisplayFileExtensionOk() (*bool, bool)`

GetDisplayFileExtensionOk returns a tuple with the DisplayFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFileExtension

`func (o *FilesSettingsDto) SetDisplayFileExtension(v bool)`

SetDisplayFileExtension sets DisplayFileExtension field to given value.

### HasDisplayFileExtension

`func (o *FilesSettingsDto) HasDisplayFileExtension() bool`

HasDisplayFileExtension returns a boolean if a field has been set.

### GetConvertNotify

`func (o *FilesSettingsDto) GetConvertNotify() bool`

GetConvertNotify returns the ConvertNotify field if non-nil, zero value otherwise.

### GetConvertNotifyOk

`func (o *FilesSettingsDto) GetConvertNotifyOk() (*bool, bool)`

GetConvertNotifyOk returns a tuple with the ConvertNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertNotify

`func (o *FilesSettingsDto) SetConvertNotify(v bool)`

SetConvertNotify sets ConvertNotify field to given value.

### HasConvertNotify

`func (o *FilesSettingsDto) HasConvertNotify() bool`

HasConvertNotify returns a boolean if a field has been set.

### GetHideConfirmCancelOperation

`func (o *FilesSettingsDto) GetHideConfirmCancelOperation() bool`

GetHideConfirmCancelOperation returns the HideConfirmCancelOperation field if non-nil, zero value otherwise.

### GetHideConfirmCancelOperationOk

`func (o *FilesSettingsDto) GetHideConfirmCancelOperationOk() (*bool, bool)`

GetHideConfirmCancelOperationOk returns a tuple with the HideConfirmCancelOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideConfirmCancelOperation

`func (o *FilesSettingsDto) SetHideConfirmCancelOperation(v bool)`

SetHideConfirmCancelOperation sets HideConfirmCancelOperation field to given value.

### HasHideConfirmCancelOperation

`func (o *FilesSettingsDto) HasHideConfirmCancelOperation() bool`

HasHideConfirmCancelOperation returns a boolean if a field has been set.

### GetHideConfirmConvertSave

`func (o *FilesSettingsDto) GetHideConfirmConvertSave() bool`

GetHideConfirmConvertSave returns the HideConfirmConvertSave field if non-nil, zero value otherwise.

### GetHideConfirmConvertSaveOk

`func (o *FilesSettingsDto) GetHideConfirmConvertSaveOk() (*bool, bool)`

GetHideConfirmConvertSaveOk returns a tuple with the HideConfirmConvertSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideConfirmConvertSave

`func (o *FilesSettingsDto) SetHideConfirmConvertSave(v bool)`

SetHideConfirmConvertSave sets HideConfirmConvertSave field to given value.

### HasHideConfirmConvertSave

`func (o *FilesSettingsDto) HasHideConfirmConvertSave() bool`

HasHideConfirmConvertSave returns a boolean if a field has been set.

### GetHideConfirmConvertOpen

`func (o *FilesSettingsDto) GetHideConfirmConvertOpen() bool`

GetHideConfirmConvertOpen returns the HideConfirmConvertOpen field if non-nil, zero value otherwise.

### GetHideConfirmConvertOpenOk

`func (o *FilesSettingsDto) GetHideConfirmConvertOpenOk() (*bool, bool)`

GetHideConfirmConvertOpenOk returns a tuple with the HideConfirmConvertOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideConfirmConvertOpen

`func (o *FilesSettingsDto) SetHideConfirmConvertOpen(v bool)`

SetHideConfirmConvertOpen sets HideConfirmConvertOpen field to given value.

### HasHideConfirmConvertOpen

`func (o *FilesSettingsDto) HasHideConfirmConvertOpen() bool`

HasHideConfirmConvertOpen returns a boolean if a field has been set.

### GetHideConfirmRoomLifetime

`func (o *FilesSettingsDto) GetHideConfirmRoomLifetime() bool`

GetHideConfirmRoomLifetime returns the HideConfirmRoomLifetime field if non-nil, zero value otherwise.

### GetHideConfirmRoomLifetimeOk

`func (o *FilesSettingsDto) GetHideConfirmRoomLifetimeOk() (*bool, bool)`

GetHideConfirmRoomLifetimeOk returns a tuple with the HideConfirmRoomLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideConfirmRoomLifetime

`func (o *FilesSettingsDto) SetHideConfirmRoomLifetime(v bool)`

SetHideConfirmRoomLifetime sets HideConfirmRoomLifetime field to given value.

### HasHideConfirmRoomLifetime

`func (o *FilesSettingsDto) HasHideConfirmRoomLifetime() bool`

HasHideConfirmRoomLifetime returns a boolean if a field has been set.

### GetDefaultOrder

`func (o *FilesSettingsDto) GetDefaultOrder() OrderBy`

GetDefaultOrder returns the DefaultOrder field if non-nil, zero value otherwise.

### GetDefaultOrderOk

`func (o *FilesSettingsDto) GetDefaultOrderOk() (*OrderBy, bool)`

GetDefaultOrderOk returns a tuple with the DefaultOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrder

`func (o *FilesSettingsDto) SetDefaultOrder(v OrderBy)`

SetDefaultOrder sets DefaultOrder field to given value.

### HasDefaultOrder

`func (o *FilesSettingsDto) HasDefaultOrder() bool`

HasDefaultOrder returns a boolean if a field has been set.

### GetForcesave

`func (o *FilesSettingsDto) GetForcesave() bool`

GetForcesave returns the Forcesave field if non-nil, zero value otherwise.

### GetForcesaveOk

`func (o *FilesSettingsDto) GetForcesaveOk() (*bool, bool)`

GetForcesaveOk returns a tuple with the Forcesave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcesave

`func (o *FilesSettingsDto) SetForcesave(v bool)`

SetForcesave sets Forcesave field to given value.

### HasForcesave

`func (o *FilesSettingsDto) HasForcesave() bool`

HasForcesave returns a boolean if a field has been set.

### GetStoreForcesave

`func (o *FilesSettingsDto) GetStoreForcesave() bool`

GetStoreForcesave returns the StoreForcesave field if non-nil, zero value otherwise.

### GetStoreForcesaveOk

`func (o *FilesSettingsDto) GetStoreForcesaveOk() (*bool, bool)`

GetStoreForcesaveOk returns a tuple with the StoreForcesave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreForcesave

`func (o *FilesSettingsDto) SetStoreForcesave(v bool)`

SetStoreForcesave sets StoreForcesave field to given value.

### HasStoreForcesave

`func (o *FilesSettingsDto) HasStoreForcesave() bool`

HasStoreForcesave returns a boolean if a field has been set.

### GetRecentSection

`func (o *FilesSettingsDto) GetRecentSection() bool`

GetRecentSection returns the RecentSection field if non-nil, zero value otherwise.

### GetRecentSectionOk

`func (o *FilesSettingsDto) GetRecentSectionOk() (*bool, bool)`

GetRecentSectionOk returns a tuple with the RecentSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSection

`func (o *FilesSettingsDto) SetRecentSection(v bool)`

SetRecentSection sets RecentSection field to given value.

### HasRecentSection

`func (o *FilesSettingsDto) HasRecentSection() bool`

HasRecentSection returns a boolean if a field has been set.

### GetFavoritesSection

`func (o *FilesSettingsDto) GetFavoritesSection() bool`

GetFavoritesSection returns the FavoritesSection field if non-nil, zero value otherwise.

### GetFavoritesSectionOk

`func (o *FilesSettingsDto) GetFavoritesSectionOk() (*bool, bool)`

GetFavoritesSectionOk returns a tuple with the FavoritesSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoritesSection

`func (o *FilesSettingsDto) SetFavoritesSection(v bool)`

SetFavoritesSection sets FavoritesSection field to given value.

### HasFavoritesSection

`func (o *FilesSettingsDto) HasFavoritesSection() bool`

HasFavoritesSection returns a boolean if a field has been set.

### GetTemplatesSection

`func (o *FilesSettingsDto) GetTemplatesSection() bool`

GetTemplatesSection returns the TemplatesSection field if non-nil, zero value otherwise.

### GetTemplatesSectionOk

`func (o *FilesSettingsDto) GetTemplatesSectionOk() (*bool, bool)`

GetTemplatesSectionOk returns a tuple with the TemplatesSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesSection

`func (o *FilesSettingsDto) SetTemplatesSection(v bool)`

SetTemplatesSection sets TemplatesSection field to given value.

### HasTemplatesSection

`func (o *FilesSettingsDto) HasTemplatesSection() bool`

HasTemplatesSection returns a boolean if a field has been set.

### GetDownloadTarGz

`func (o *FilesSettingsDto) GetDownloadTarGz() bool`

GetDownloadTarGz returns the DownloadTarGz field if non-nil, zero value otherwise.

### GetDownloadTarGzOk

`func (o *FilesSettingsDto) GetDownloadTarGzOk() (*bool, bool)`

GetDownloadTarGzOk returns a tuple with the DownloadTarGz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadTarGz

`func (o *FilesSettingsDto) SetDownloadTarGz(v bool)`

SetDownloadTarGz sets DownloadTarGz field to given value.

### HasDownloadTarGz

`func (o *FilesSettingsDto) HasDownloadTarGz() bool`

HasDownloadTarGz returns a boolean if a field has been set.

### GetAutomaticallyCleanUp

`func (o *FilesSettingsDto) GetAutomaticallyCleanUp() AutoCleanUpData`

GetAutomaticallyCleanUp returns the AutomaticallyCleanUp field if non-nil, zero value otherwise.

### GetAutomaticallyCleanUpOk

`func (o *FilesSettingsDto) GetAutomaticallyCleanUpOk() (*AutoCleanUpData, bool)`

GetAutomaticallyCleanUpOk returns a tuple with the AutomaticallyCleanUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyCleanUp

`func (o *FilesSettingsDto) SetAutomaticallyCleanUp(v AutoCleanUpData)`

SetAutomaticallyCleanUp sets AutomaticallyCleanUp field to given value.

### HasAutomaticallyCleanUp

`func (o *FilesSettingsDto) HasAutomaticallyCleanUp() bool`

HasAutomaticallyCleanUp returns a boolean if a field has been set.

### GetCanSearchByContent

`func (o *FilesSettingsDto) GetCanSearchByContent() bool`

GetCanSearchByContent returns the CanSearchByContent field if non-nil, zero value otherwise.

### GetCanSearchByContentOk

`func (o *FilesSettingsDto) GetCanSearchByContentOk() (*bool, bool)`

GetCanSearchByContentOk returns a tuple with the CanSearchByContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSearchByContent

`func (o *FilesSettingsDto) SetCanSearchByContent(v bool)`

SetCanSearchByContent sets CanSearchByContent field to given value.

### HasCanSearchByContent

`func (o *FilesSettingsDto) HasCanSearchByContent() bool`

HasCanSearchByContent returns a boolean if a field has been set.

### GetDefaultSharingAccessRights

`func (o *FilesSettingsDto) GetDefaultSharingAccessRights() []int32`

GetDefaultSharingAccessRights returns the DefaultSharingAccessRights field if non-nil, zero value otherwise.

### GetDefaultSharingAccessRightsOk

`func (o *FilesSettingsDto) GetDefaultSharingAccessRightsOk() (*[]int32, bool)`

GetDefaultSharingAccessRightsOk returns a tuple with the DefaultSharingAccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSharingAccessRights

`func (o *FilesSettingsDto) SetDefaultSharingAccessRights(v []int32)`

SetDefaultSharingAccessRights sets DefaultSharingAccessRights field to given value.

### HasDefaultSharingAccessRights

`func (o *FilesSettingsDto) HasDefaultSharingAccessRights() bool`

HasDefaultSharingAccessRights returns a boolean if a field has been set.

### SetDefaultSharingAccessRightsNil

`func (o *FilesSettingsDto) SetDefaultSharingAccessRightsNil(b bool)`

 SetDefaultSharingAccessRightsNil sets the value for DefaultSharingAccessRights to be an explicit nil

### UnsetDefaultSharingAccessRights
`func (o *FilesSettingsDto) UnsetDefaultSharingAccessRights()`

UnsetDefaultSharingAccessRights ensures that no value is present for DefaultSharingAccessRights, not even an explicit nil
### GetMaxUploadThreadCount

`func (o *FilesSettingsDto) GetMaxUploadThreadCount() int32`

GetMaxUploadThreadCount returns the MaxUploadThreadCount field if non-nil, zero value otherwise.

### GetMaxUploadThreadCountOk

`func (o *FilesSettingsDto) GetMaxUploadThreadCountOk() (*int32, bool)`

GetMaxUploadThreadCountOk returns a tuple with the MaxUploadThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUploadThreadCount

`func (o *FilesSettingsDto) SetMaxUploadThreadCount(v int32)`

SetMaxUploadThreadCount sets MaxUploadThreadCount field to given value.

### HasMaxUploadThreadCount

`func (o *FilesSettingsDto) HasMaxUploadThreadCount() bool`

HasMaxUploadThreadCount returns a boolean if a field has been set.

### GetChunkUploadSize

`func (o *FilesSettingsDto) GetChunkUploadSize() int64`

GetChunkUploadSize returns the ChunkUploadSize field if non-nil, zero value otherwise.

### GetChunkUploadSizeOk

`func (o *FilesSettingsDto) GetChunkUploadSizeOk() (*int64, bool)`

GetChunkUploadSizeOk returns a tuple with the ChunkUploadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkUploadSize

`func (o *FilesSettingsDto) SetChunkUploadSize(v int64)`

SetChunkUploadSize sets ChunkUploadSize field to given value.

### HasChunkUploadSize

`func (o *FilesSettingsDto) HasChunkUploadSize() bool`

HasChunkUploadSize returns a boolean if a field has been set.

### GetOpenEditorInSameTab

`func (o *FilesSettingsDto) GetOpenEditorInSameTab() bool`

GetOpenEditorInSameTab returns the OpenEditorInSameTab field if non-nil, zero value otherwise.

### GetOpenEditorInSameTabOk

`func (o *FilesSettingsDto) GetOpenEditorInSameTabOk() (*bool, bool)`

GetOpenEditorInSameTabOk returns a tuple with the OpenEditorInSameTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenEditorInSameTab

`func (o *FilesSettingsDto) SetOpenEditorInSameTab(v bool)`

SetOpenEditorInSameTab sets OpenEditorInSameTab field to given value.

### HasOpenEditorInSameTab

`func (o *FilesSettingsDto) HasOpenEditorInSameTab() bool`

HasOpenEditorInSameTab returns a boolean if a field has been set.

### GetOrganizeRoomsGrouping

`func (o *FilesSettingsDto) GetOrganizeRoomsGrouping() bool`

GetOrganizeRoomsGrouping returns the OrganizeRoomsGrouping field if non-nil, zero value otherwise.

### GetOrganizeRoomsGroupingOk

`func (o *FilesSettingsDto) GetOrganizeRoomsGroupingOk() (*bool, bool)`

GetOrganizeRoomsGroupingOk returns a tuple with the OrganizeRoomsGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizeRoomsGrouping

`func (o *FilesSettingsDto) SetOrganizeRoomsGrouping(v bool)`

SetOrganizeRoomsGrouping sets OrganizeRoomsGrouping field to given value.

### HasOrganizeRoomsGrouping

`func (o *FilesSettingsDto) HasOrganizeRoomsGrouping() bool`

HasOrganizeRoomsGrouping returns a boolean if a field has been set.

### GetDefaultShareLinkInternal

`func (o *FilesSettingsDto) GetDefaultShareLinkInternal() bool`

GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field if non-nil, zero value otherwise.

### GetDefaultShareLinkInternalOk

`func (o *FilesSettingsDto) GetDefaultShareLinkInternalOk() (*bool, bool)`

GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareLinkInternal

`func (o *FilesSettingsDto) SetDefaultShareLinkInternal(v bool)`

SetDefaultShareLinkInternal sets DefaultShareLinkInternal field to given value.

### HasDefaultShareLinkInternal

`func (o *FilesSettingsDto) HasDefaultShareLinkInternal() bool`

HasDefaultShareLinkInternal returns a boolean if a field has been set.

### GetExternalShareApplyToDocuments

`func (o *FilesSettingsDto) GetExternalShareApplyToDocuments() bool`

GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field if non-nil, zero value otherwise.

### GetExternalShareApplyToDocumentsOk

`func (o *FilesSettingsDto) GetExternalShareApplyToDocumentsOk() (*bool, bool)`

GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToDocuments

`func (o *FilesSettingsDto) SetExternalShareApplyToDocuments(v bool)`

SetExternalShareApplyToDocuments sets ExternalShareApplyToDocuments field to given value.

### HasExternalShareApplyToDocuments

`func (o *FilesSettingsDto) HasExternalShareApplyToDocuments() bool`

HasExternalShareApplyToDocuments returns a boolean if a field has been set.

### GetExternalShareApplyToRooms

`func (o *FilesSettingsDto) GetExternalShareApplyToRooms() bool`

GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field if non-nil, zero value otherwise.

### GetExternalShareApplyToRoomsOk

`func (o *FilesSettingsDto) GetExternalShareApplyToRoomsOk() (*bool, bool)`

GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShareApplyToRooms

`func (o *FilesSettingsDto) SetExternalShareApplyToRooms(v bool)`

SetExternalShareApplyToRooms sets ExternalShareApplyToRooms field to given value.

### HasExternalShareApplyToRooms

`func (o *FilesSettingsDto) HasExternalShareApplyToRooms() bool`

HasExternalShareApplyToRooms returns a boolean if a field has been set.

### GetBlockExistingLinksOnRestrict

`func (o *FilesSettingsDto) GetBlockExistingLinksOnRestrict() bool`

GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field if non-nil, zero value otherwise.

### GetBlockExistingLinksOnRestrictOk

`func (o *FilesSettingsDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool)`

GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExistingLinksOnRestrict

`func (o *FilesSettingsDto) SetBlockExistingLinksOnRestrict(v bool)`

SetBlockExistingLinksOnRestrict sets BlockExistingLinksOnRestrict field to given value.

### HasBlockExistingLinksOnRestrict

`func (o *FilesSettingsDto) HasBlockExistingLinksOnRestrict() bool`

HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.

### GetExtsFilesVectorized

`func (o *FilesSettingsDto) GetExtsFilesVectorized() []string`

GetExtsFilesVectorized returns the ExtsFilesVectorized field if non-nil, zero value otherwise.

### GetExtsFilesVectorizedOk

`func (o *FilesSettingsDto) GetExtsFilesVectorizedOk() (*[]string, bool)`

GetExtsFilesVectorizedOk returns a tuple with the ExtsFilesVectorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtsFilesVectorized

`func (o *FilesSettingsDto) SetExtsFilesVectorized(v []string)`

SetExtsFilesVectorized sets ExtsFilesVectorized field to given value.

### HasExtsFilesVectorized

`func (o *FilesSettingsDto) HasExtsFilesVectorized() bool`

HasExtsFilesVectorized returns a boolean if a field has been set.

### SetExtsFilesVectorizedNil

`func (o *FilesSettingsDto) SetExtsFilesVectorizedNil(b bool)`

 SetExtsFilesVectorizedNil sets the value for ExtsFilesVectorized to be an explicit nil

### UnsetExtsFilesVectorized
`func (o *FilesSettingsDto) UnsetExtsFilesVectorized()`

UnsetExtsFilesVectorized ensures that no value is present for ExtsFilesVectorized, not even an explicit nil
### GetMaxVectorizationFileSize

`func (o *FilesSettingsDto) GetMaxVectorizationFileSize() int64`

GetMaxVectorizationFileSize returns the MaxVectorizationFileSize field if non-nil, zero value otherwise.

### GetMaxVectorizationFileSizeOk

`func (o *FilesSettingsDto) GetMaxVectorizationFileSizeOk() (*int64, bool)`

GetMaxVectorizationFileSizeOk returns a tuple with the MaxVectorizationFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVectorizationFileSize

`func (o *FilesSettingsDto) SetMaxVectorizationFileSize(v int64)`

SetMaxVectorizationFileSize sets MaxVectorizationFileSize field to given value.

### HasMaxVectorizationFileSize

`func (o *FilesSettingsDto) HasMaxVectorizationFileSize() bool`

HasMaxVectorizationFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


