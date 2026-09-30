// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
)

// checks if the FilesSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSettingsDto{}

// FilesSettingsDto Everything a client needs to work with documents in this portal: the format tables, the address templates, the  upload limits, the portal-wide switches and the preferences of the calling account.
type FilesSettingsDto struct {
	// Images the portal can show in its own viewer. Anything outside the list has to be downloaded to be seen.
	ExtsImagePreviewed []string `json:"extsImagePreviewed,omitempty"`
	// Audio and video the portal can play in its own player.
	ExtsMediaPreviewed []string `json:"extsMediaPreviewed,omitempty"`
	// Documents the editor can open read-only. A format that is here but not in the edited list can be viewed and  not changed.
	ExtsWebPreviewed []string `json:"extsWebPreviewed,omitempty"`
	// Documents the editor can open for editing. Uploading a format outside this list and outside the convertible  list leaves a file that can only be downloaded.
	ExtsWebEdited []string `json:"extsWebEdited,omitempty"`
	// Documents that can be edited inside a private room, where the content is encrypted on the client.
	ExtsWebEncrypt []string `json:"extsWebEncrypt,omitempty"`
	// Documents that support the reviewing mode, so that granting review access to them is meaningful.
	ExtsWebReviewed []string `json:"extsWebReviewed,omitempty"`
	// Spreadsheets that support the custom filter mode, where a filter applied by one editor does not disturb the  others.
	ExtsWebCustomFilterEditing []string `json:"extsWebCustomFilterEditing,omitempty"`
	// Documents that can only be filled in or commented on rather than edited freely, whatever access the caller  holds.
	ExtsWebRestrictedEditing []string `json:"extsWebRestrictedEditing,omitempty"`
	// Documents that support comments, so that granting comment access to them is meaningful.
	ExtsWebCommented []string `json:"extsWebCommented,omitempty"`
	// Documents the portal treats as templates to create new files from.
	ExtsWebTemplate []string `json:"extsWebTemplate,omitempty"`
	// Formats that cannot be edited as they are and are converted on upload or on first opening. Which target each  one has is in the convertible table below.
	ExtsMustConvert []string `json:"extsMustConvert,omitempty"`
	// The conversion map of the portal: for each source extension, the extensions it can be converted into. Use it  to fill the target format of a conversion request instead of guessing one.
	ExtsConvertible map[string][]string `json:"extsConvertible,omitempty"`
	// Formats the portal offers to create and upload as documents. It is not an upload filter: files of other  formats are stored as they are.
	ExtsUploadable []string `json:"extsUploadable,omitempty"`
	// Formats recognised as archives, which is what decides the archive icon and the offer to unpack.
	ExtsArchive []string `json:"extsArchive,omitempty"`
	// Formats classified as video. The classification lists drive icons and the media filters of the listing  operations, and are wider than what the built-in player can show.
	ExtsVideo []string `json:"extsVideo,omitempty"`
	// Formats classified as audio.
	ExtsAudio []string `json:"extsAudio,omitempty"`
	// Formats classified as images.
	ExtsImage []string `json:"extsImage,omitempty"`
	// Formats classified as spreadsheets.
	ExtsSpreadsheet []string `json:"extsSpreadsheet,omitempty"`
	// Formats classified as presentations.
	ExtsPresentation []string `json:"extsPresentation,omitempty"`
	// Formats classified as text documents.
	ExtsDocument []string `json:"extsDocument,omitempty"`
	// Formats classified as diagrams.
	ExtsDiagram []string `json:"extsDiagram,omitempty"`
	InternalFormats NullableFilesSettingsDtoInternalFormats `json:"internalFormats,omitempty"`
	// The extension of a fillable form template in this portal. It is configurable, so read it rather than assuming  the product default.
	MasterFormExtension NullableString `json:"masterFormExtension,omitempty"`
	// The name of the query parameter that pins a document address to one version. Append it to the addresses below  instead of composing a version address by hand.
	ParamVersion NullableString `json:"paramVersion,omitempty"`
	// The name of the query parameter that asks a download address for a converted copy in another format.
	ParamOutType NullableString `json:"paramOutType,omitempty"`
	// The template of the address a file is downloaded from: substitute the file identifier for the `{0}`  placeholder. Add the version and output-type parameters named above for a particular version or format.
	FileDownloadUrlString NullableString `json:"fileDownloadUrlString,omitempty"`
	// The template of the address that opens a file in the viewer inside the portal, with `{0}` for the file  identifier. It is a portal-relative address, meant to be opened in a browser rather than called as an API.
	FileWebViewerUrlString NullableString `json:"fileWebViewerUrlString,omitempty"`
	// The same viewer address as an absolute one, for a message or a page outside the portal.
	FileWebViewerExternalUrlString NullableString `json:"fileWebViewerExternalUrlString,omitempty"`
	// The template of the address that opens a file for editing inside the portal, with `{0}` for the file  identifier. Whether the session really becomes editable still depends on the access the caller holds.
	FileWebEditorUrlString NullableString `json:"fileWebEditorUrlString,omitempty"`
	// The same editing address as an absolute one, for use outside the portal.
	FileWebEditorExternalUrlString NullableString `json:"fileWebEditorExternalUrlString,omitempty"`
	// The template of the address that sends the browser on to whichever viewer or editor suits the file, with `{0}`  for the file identifier. Use it when the kind of the file is not known in advance.
	FileRedirectPreviewUrlString NullableString `json:"fileRedirectPreviewUrlString,omitempty"`
	// The template of the address a file thumbnail is fetched from, with `{0}` for the file identifier. A thumbnail  is built in the background, so the address can answer with nothing for a while after the file appears.
	FileThumbnailUrlString NullableString `json:"fileThumbnailUrlString,omitempty"`
	// Whether the caller asked to be prompted before a deletion. Written by `PUT api/2.0/files/changedeleteconfrim`.
	ConfirmDelete *bool `json:"confirmDelete,omitempty"`
	// Whether this portal allows third-party storages to be connected at all. It is set portal-wide by an  administrator, so a member sees it as read-only.
	EnableThirdParty *bool `json:"enableThirdParty,omitempty"`
	// Whether links that open an entry without a portal account may be created in this portal. Set portal-wide by an  administrator.
	ExternalShare *bool `json:"externalShare,omitempty"`
	// Whether the share-to-network buttons are offered next to an external link. It is reported as false whenever  external sharing itself is off.
	ExternalShareSocialMedia *bool `json:"externalShareSocialMedia,omitempty"`
	// Whether the caller's uploads keep the original file when the portal converts them. With false the conversion  replaces the uploaded file with a new version of it.
	StoreOriginalFiles *bool `json:"storeOriginalFiles,omitempty"`
	// Whether the caller asked for new documents to be created with the default name instead of being prompted for  one.
	KeepNewFileName *bool `json:"keepNewFileName,omitempty"`
	// Whether the caller asked to see extensions in file titles. Stored titles always carry the extension whatever  this says.
	DisplayFileExtension *bool `json:"displayFileExtension,omitempty"`
	// Specifies whether to display the quick action buttons.
	ShowQuickActions *bool `json:"showQuickActions,omitempty"`
	// Whether the caller is told about the result of a conversion. There is no operation in this document that  writes it.
	ConvertNotify *bool `json:"convertNotify,omitempty"`
	// Whether the prompt shown before a running operation is abandoned is hidden for the caller.
	HideConfirmCancelOperation *bool `json:"hideConfirmCancelOperation,omitempty"`
	// Whether the prompt that offers to keep a copy in the original format on conversion is hidden for the caller.  Once true it cannot be turned back through the API.
	HideConfirmConvertSave *bool `json:"hideConfirmConvertSave,omitempty"`
	// Whether the prompt that offers to open the conversion result is hidden for the caller. Once true it cannot be  turned back through the API.
	HideConfirmConvertOpen *bool `json:"hideConfirmConvertOpen,omitempty"`
	// Whether the warning shown before the lifetime settings of a room are changed is hidden for the caller.
	HideConfirmRoomLifetime *bool `json:"hideConfirmRoomLifetime,omitempty"`
	// The ordering the listing operations fall back to when a request names none. It follows the last order the  caller asked a listing for, so it changes on its own as the account is used.
	DefaultOrder *OrderBy `json:"defaultOrder,omitempty"`
	// Whether the editor writes a document back to storage while the session is still open. It is on for every  portal and cannot be switched off.
	Forcesave *bool `json:"forcesave,omitempty"`
	// Whether those intermediate saves are kept as separate versions. They are not, in any portal: they update the  current version instead.
	StoreForcesave *bool `json:"storeForcesave,omitempty"`
	// Whether the Recent section is offered to the caller among the section roots.
	RecentSection *bool `json:"recentSection,omitempty"`
	// Whether the Favorites section is offered to the caller among the section roots.
	FavoritesSection *bool `json:"favoritesSection,omitempty"`
	// Whether the Templates section is offered to the caller among the section roots.
	TemplatesSection *bool `json:"templatesSection,omitempty"`
	// The archive format the caller's multi-item downloads are packed into: true for `.tar.gz`, false for `.zip`.
	DownloadTarGz *bool `json:"downloadTarGz,omitempty"`
	// The trash auto-clearing setting of the caller, the same pair `GET api/2.0/files/settings/autocleanup` returns.
	AutomaticallyCleanUp *AutoCleanUpData `json:"automaticallyCleanUp,omitempty"`
	// Whether documents in this portal can be searched by what is inside them and not only by title. It depends on  the full-text search service being configured and having indexed the portal.
	CanSearchByContent *bool `json:"canSearchByContent,omitempty"`
	// The access rights the sharing dialog offers the caller by default. The portal normalises the set it stores, so  this can be shorter than what was last sent.
	DefaultSharingAccessRights []int32 `json:"defaultSharingAccessRights,omitempty"`
	// How many upload requests the portal accepts from one account at a time. Sending more than this in parallel  gets the extra ones refused rather than queued.
	MaxUploadThreadCount *int32 `json:"maxUploadThreadCount,omitempty"`
	// The size in bytes of one chunk of a chunked upload. Split a large file exactly along this size: a chunk that  does not match is refused by the upload session.
	ChunkUploadSize *int64 `json:"chunkUploadSize,omitempty"`
	// Whether the caller asked for documents to open in the current browser tab.
	OpenEditorInSameTab *bool `json:"openEditorInSameTab,omitempty"`
	// Whether the caller asked to see rooms arranged by the groups they belong to.
	OrganizeRoomsGrouping *bool `json:"organizeRoomsGrouping,omitempty"`
	// The kind of external link this portal offers first: true for a link only its own accounts can open, false for  one anyone holding it can open.
	DefaultShareLinkInternal *bool `json:"defaultShareLinkInternal,omitempty"`
	// Whether the external sharing restriction covers personal documents. It matters only while external sharing is  off.
	ExternalShareApplyToDocuments *bool `json:"externalShareApplyToDocuments,omitempty"`
	// Whether the external sharing restriction covers rooms, including making a new one public. It matters only  while external sharing is off.
	ExternalShareApplyToRooms *bool `json:"externalShareApplyToRooms,omitempty"`
	// Whether links created before the restriction stop opening as well, rather than only new ones being refused.
	BlockExistingLinksOnRestrict *bool `json:"blockExistingLinksOnRestrict,omitempty"`
	// Formats whose content can be indexed for the AI features of the portal. A file outside the list is left out of  that index.
	ExtsFilesVectorized []string `json:"extsFilesVectorized,omitempty"`
	// The largest file size in bytes that is indexed for the AI features. A larger file is skipped even when its  format is listed above.
	MaxVectorizationFileSize *int64 `json:"maxVectorizationFileSize,omitempty"`
}

// NewFilesSettingsDto instantiates a new FilesSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSettingsDto() *FilesSettingsDto {
	this := FilesSettingsDto{}
	return &this
}

// NewFilesSettingsDtoWithDefaults instantiates a new FilesSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSettingsDtoWithDefaults() *FilesSettingsDto {
	this := FilesSettingsDto{}
	return &this
}

// GetExtsImagePreviewed returns the ExtsImagePreviewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsImagePreviewed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsImagePreviewed
}

// GetExtsImagePreviewedOk returns a tuple with the ExtsImagePreviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsImagePreviewedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsImagePreviewed) {
		return nil, false
	}
	return o.ExtsImagePreviewed, true
}

// HasExtsImagePreviewed returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsImagePreviewedSet() bool {
	if o != nil && !IsNil(o.ExtsImagePreviewed) {
		return true
	}

	return false
}

// SetExtsImagePreviewed gets a reference to the given []string and assigns it to the ExtsImagePreviewed field.
func (o *FilesSettingsDto) SetExtsImagePreviewed(v []string) {
	o.ExtsImagePreviewed = v
}

// GetExtsMediaPreviewed returns the ExtsMediaPreviewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsMediaPreviewed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsMediaPreviewed
}

// GetExtsMediaPreviewedOk returns a tuple with the ExtsMediaPreviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsMediaPreviewedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsMediaPreviewed) {
		return nil, false
	}
	return o.ExtsMediaPreviewed, true
}

// HasExtsMediaPreviewed returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsMediaPreviewedSet() bool {
	if o != nil && !IsNil(o.ExtsMediaPreviewed) {
		return true
	}

	return false
}

// SetExtsMediaPreviewed gets a reference to the given []string and assigns it to the ExtsMediaPreviewed field.
func (o *FilesSettingsDto) SetExtsMediaPreviewed(v []string) {
	o.ExtsMediaPreviewed = v
}

// GetExtsWebPreviewed returns the ExtsWebPreviewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebPreviewed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebPreviewed
}

// GetExtsWebPreviewedOk returns a tuple with the ExtsWebPreviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebPreviewedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebPreviewed) {
		return nil, false
	}
	return o.ExtsWebPreviewed, true
}

// HasExtsWebPreviewed returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebPreviewedSet() bool {
	if o != nil && !IsNil(o.ExtsWebPreviewed) {
		return true
	}

	return false
}

// SetExtsWebPreviewed gets a reference to the given []string and assigns it to the ExtsWebPreviewed field.
func (o *FilesSettingsDto) SetExtsWebPreviewed(v []string) {
	o.ExtsWebPreviewed = v
}

// GetExtsWebEdited returns the ExtsWebEdited field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebEdited() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebEdited
}

// GetExtsWebEditedOk returns a tuple with the ExtsWebEdited field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebEditedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebEdited) {
		return nil, false
	}
	return o.ExtsWebEdited, true
}

// HasExtsWebEdited returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebEditedSet() bool {
	if o != nil && !IsNil(o.ExtsWebEdited) {
		return true
	}

	return false
}

// SetExtsWebEdited gets a reference to the given []string and assigns it to the ExtsWebEdited field.
func (o *FilesSettingsDto) SetExtsWebEdited(v []string) {
	o.ExtsWebEdited = v
}

// GetExtsWebEncrypt returns the ExtsWebEncrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebEncrypt() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebEncrypt
}

// GetExtsWebEncryptOk returns a tuple with the ExtsWebEncrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebEncryptOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebEncrypt) {
		return nil, false
	}
	return o.ExtsWebEncrypt, true
}

// HasExtsWebEncrypt returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebEncryptSet() bool {
	if o != nil && !IsNil(o.ExtsWebEncrypt) {
		return true
	}

	return false
}

// SetExtsWebEncrypt gets a reference to the given []string and assigns it to the ExtsWebEncrypt field.
func (o *FilesSettingsDto) SetExtsWebEncrypt(v []string) {
	o.ExtsWebEncrypt = v
}

// GetExtsWebReviewed returns the ExtsWebReviewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebReviewed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebReviewed
}

// GetExtsWebReviewedOk returns a tuple with the ExtsWebReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebReviewedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebReviewed) {
		return nil, false
	}
	return o.ExtsWebReviewed, true
}

// HasExtsWebReviewed returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebReviewedSet() bool {
	if o != nil && !IsNil(o.ExtsWebReviewed) {
		return true
	}

	return false
}

// SetExtsWebReviewed gets a reference to the given []string and assigns it to the ExtsWebReviewed field.
func (o *FilesSettingsDto) SetExtsWebReviewed(v []string) {
	o.ExtsWebReviewed = v
}

// GetExtsWebCustomFilterEditing returns the ExtsWebCustomFilterEditing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebCustomFilterEditing() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebCustomFilterEditing
}

// GetExtsWebCustomFilterEditingOk returns a tuple with the ExtsWebCustomFilterEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebCustomFilterEditingOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebCustomFilterEditing) {
		return nil, false
	}
	return o.ExtsWebCustomFilterEditing, true
}

// HasExtsWebCustomFilterEditing returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebCustomFilterEditingSet() bool {
	if o != nil && !IsNil(o.ExtsWebCustomFilterEditing) {
		return true
	}

	return false
}

// SetExtsWebCustomFilterEditing gets a reference to the given []string and assigns it to the ExtsWebCustomFilterEditing field.
func (o *FilesSettingsDto) SetExtsWebCustomFilterEditing(v []string) {
	o.ExtsWebCustomFilterEditing = v
}

// GetExtsWebRestrictedEditing returns the ExtsWebRestrictedEditing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebRestrictedEditing() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebRestrictedEditing
}

// GetExtsWebRestrictedEditingOk returns a tuple with the ExtsWebRestrictedEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebRestrictedEditingOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebRestrictedEditing) {
		return nil, false
	}
	return o.ExtsWebRestrictedEditing, true
}

// HasExtsWebRestrictedEditing returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebRestrictedEditingSet() bool {
	if o != nil && !IsNil(o.ExtsWebRestrictedEditing) {
		return true
	}

	return false
}

// SetExtsWebRestrictedEditing gets a reference to the given []string and assigns it to the ExtsWebRestrictedEditing field.
func (o *FilesSettingsDto) SetExtsWebRestrictedEditing(v []string) {
	o.ExtsWebRestrictedEditing = v
}

// GetExtsWebCommented returns the ExtsWebCommented field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebCommented() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebCommented
}

// GetExtsWebCommentedOk returns a tuple with the ExtsWebCommented field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebCommentedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebCommented) {
		return nil, false
	}
	return o.ExtsWebCommented, true
}

// HasExtsWebCommented returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebCommentedSet() bool {
	if o != nil && !IsNil(o.ExtsWebCommented) {
		return true
	}

	return false
}

// SetExtsWebCommented gets a reference to the given []string and assigns it to the ExtsWebCommented field.
func (o *FilesSettingsDto) SetExtsWebCommented(v []string) {
	o.ExtsWebCommented = v
}

// GetExtsWebTemplate returns the ExtsWebTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsWebTemplate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsWebTemplate
}

// GetExtsWebTemplateOk returns a tuple with the ExtsWebTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsWebTemplateOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsWebTemplate) {
		return nil, false
	}
	return o.ExtsWebTemplate, true
}

// HasExtsWebTemplate returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsWebTemplateSet() bool {
	if o != nil && !IsNil(o.ExtsWebTemplate) {
		return true
	}

	return false
}

// SetExtsWebTemplate gets a reference to the given []string and assigns it to the ExtsWebTemplate field.
func (o *FilesSettingsDto) SetExtsWebTemplate(v []string) {
	o.ExtsWebTemplate = v
}

// GetExtsMustConvert returns the ExtsMustConvert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsMustConvert() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsMustConvert
}

// GetExtsMustConvertOk returns a tuple with the ExtsMustConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsMustConvertOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsMustConvert) {
		return nil, false
	}
	return o.ExtsMustConvert, true
}

// HasExtsMustConvert returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsMustConvertSet() bool {
	if o != nil && !IsNil(o.ExtsMustConvert) {
		return true
	}

	return false
}

// SetExtsMustConvert gets a reference to the given []string and assigns it to the ExtsMustConvert field.
func (o *FilesSettingsDto) SetExtsMustConvert(v []string) {
	o.ExtsMustConvert = v
}

// GetExtsConvertible returns the ExtsConvertible field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetExtsConvertible() map[string][]string {
	if o == nil || IsNil(o.ExtsConvertible) {
		var ret map[string][]string
		return ret
	}
	return o.ExtsConvertible
}

// GetExtsConvertibleOk returns a tuple with the ExtsConvertible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetExtsConvertibleOk() (map[string][]string, bool) {
	if o == nil || IsNil(o.ExtsConvertible) {
		return map[string][]string{}, false
	}
	return o.ExtsConvertible, true
}

// HasExtsConvertible returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsConvertibleSet() bool {
	if o != nil && !IsNil(o.ExtsConvertible) {
		return true
	}

	return false
}

// SetExtsConvertible gets a reference to the given map[string][]string and assigns it to the ExtsConvertible field.
func (o *FilesSettingsDto) SetExtsConvertible(v map[string][]string) {
	o.ExtsConvertible = v
}

// GetExtsUploadable returns the ExtsUploadable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsUploadable() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsUploadable
}

// GetExtsUploadableOk returns a tuple with the ExtsUploadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsUploadableOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsUploadable) {
		return nil, false
	}
	return o.ExtsUploadable, true
}

// HasExtsUploadable returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsUploadableSet() bool {
	if o != nil && !IsNil(o.ExtsUploadable) {
		return true
	}

	return false
}

// SetExtsUploadable gets a reference to the given []string and assigns it to the ExtsUploadable field.
func (o *FilesSettingsDto) SetExtsUploadable(v []string) {
	o.ExtsUploadable = v
}

// GetExtsArchive returns the ExtsArchive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsArchive() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsArchive
}

// GetExtsArchiveOk returns a tuple with the ExtsArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsArchiveOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsArchive) {
		return nil, false
	}
	return o.ExtsArchive, true
}

// HasExtsArchive returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsArchiveSet() bool {
	if o != nil && !IsNil(o.ExtsArchive) {
		return true
	}

	return false
}

// SetExtsArchive gets a reference to the given []string and assigns it to the ExtsArchive field.
func (o *FilesSettingsDto) SetExtsArchive(v []string) {
	o.ExtsArchive = v
}

// GetExtsVideo returns the ExtsVideo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsVideo() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsVideo
}

// GetExtsVideoOk returns a tuple with the ExtsVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsVideoOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsVideo) {
		return nil, false
	}
	return o.ExtsVideo, true
}

// HasExtsVideo returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsVideoSet() bool {
	if o != nil && !IsNil(o.ExtsVideo) {
		return true
	}

	return false
}

// SetExtsVideo gets a reference to the given []string and assigns it to the ExtsVideo field.
func (o *FilesSettingsDto) SetExtsVideo(v []string) {
	o.ExtsVideo = v
}

// GetExtsAudio returns the ExtsAudio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsAudio() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsAudio
}

// GetExtsAudioOk returns a tuple with the ExtsAudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsAudioOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsAudio) {
		return nil, false
	}
	return o.ExtsAudio, true
}

// HasExtsAudio returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsAudioSet() bool {
	if o != nil && !IsNil(o.ExtsAudio) {
		return true
	}

	return false
}

// SetExtsAudio gets a reference to the given []string and assigns it to the ExtsAudio field.
func (o *FilesSettingsDto) SetExtsAudio(v []string) {
	o.ExtsAudio = v
}

// GetExtsImage returns the ExtsImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsImage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsImage
}

// GetExtsImageOk returns a tuple with the ExtsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsImageOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsImage) {
		return nil, false
	}
	return o.ExtsImage, true
}

// HasExtsImage returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsImageSet() bool {
	if o != nil && !IsNil(o.ExtsImage) {
		return true
	}

	return false
}

// SetExtsImage gets a reference to the given []string and assigns it to the ExtsImage field.
func (o *FilesSettingsDto) SetExtsImage(v []string) {
	o.ExtsImage = v
}

// GetExtsSpreadsheet returns the ExtsSpreadsheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsSpreadsheet() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsSpreadsheet
}

// GetExtsSpreadsheetOk returns a tuple with the ExtsSpreadsheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsSpreadsheetOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsSpreadsheet) {
		return nil, false
	}
	return o.ExtsSpreadsheet, true
}

// HasExtsSpreadsheet returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsSpreadsheetSet() bool {
	if o != nil && !IsNil(o.ExtsSpreadsheet) {
		return true
	}

	return false
}

// SetExtsSpreadsheet gets a reference to the given []string and assigns it to the ExtsSpreadsheet field.
func (o *FilesSettingsDto) SetExtsSpreadsheet(v []string) {
	o.ExtsSpreadsheet = v
}

// GetExtsPresentation returns the ExtsPresentation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsPresentation() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsPresentation
}

// GetExtsPresentationOk returns a tuple with the ExtsPresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsPresentationOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsPresentation) {
		return nil, false
	}
	return o.ExtsPresentation, true
}

// HasExtsPresentation returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsPresentationSet() bool {
	if o != nil && !IsNil(o.ExtsPresentation) {
		return true
	}

	return false
}

// SetExtsPresentation gets a reference to the given []string and assigns it to the ExtsPresentation field.
func (o *FilesSettingsDto) SetExtsPresentation(v []string) {
	o.ExtsPresentation = v
}

// GetExtsDocument returns the ExtsDocument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsDocument() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsDocument
}

// GetExtsDocumentOk returns a tuple with the ExtsDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsDocumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsDocument) {
		return nil, false
	}
	return o.ExtsDocument, true
}

// HasExtsDocument returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsDocumentSet() bool {
	if o != nil && !IsNil(o.ExtsDocument) {
		return true
	}

	return false
}

// SetExtsDocument gets a reference to the given []string and assigns it to the ExtsDocument field.
func (o *FilesSettingsDto) SetExtsDocument(v []string) {
	o.ExtsDocument = v
}

// GetExtsDiagram returns the ExtsDiagram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsDiagram() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsDiagram
}

// GetExtsDiagramOk returns a tuple with the ExtsDiagram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsDiagramOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsDiagram) {
		return nil, false
	}
	return o.ExtsDiagram, true
}

// HasExtsDiagram returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsDiagramSet() bool {
	if o != nil && !IsNil(o.ExtsDiagram) {
		return true
	}

	return false
}

// SetExtsDiagram gets a reference to the given []string and assigns it to the ExtsDiagram field.
func (o *FilesSettingsDto) SetExtsDiagram(v []string) {
	o.ExtsDiagram = v
}

// GetInternalFormats returns the InternalFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetInternalFormats() FilesSettingsDtoInternalFormats {
	if o == nil || IsNil(o.InternalFormats.Get()) {
		var ret FilesSettingsDtoInternalFormats
		return ret
	}
	return *o.InternalFormats.Get()
}

// GetInternalFormatsOk returns a tuple with the InternalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetInternalFormatsOk() (*FilesSettingsDtoInternalFormats, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalFormats.Get(), o.InternalFormats.IsSet()
}

// HasInternalFormats returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsInternalFormatsSet() bool {
	if o != nil && o.InternalFormats.IsSet() {
		return true
	}

	return false
}

// SetInternalFormats gets a reference to the given NullableFilesSettingsDtoInternalFormats and assigns it to the InternalFormats field.
func (o *FilesSettingsDto) SetInternalFormats(v FilesSettingsDtoInternalFormats) {
	o.InternalFormats.Set(&v)
}
// SetInternalFormatsNil sets the value for InternalFormats to be an explicit nil
func (o *FilesSettingsDto) SetInternalFormatsNil() {
	o.InternalFormats.Set(nil)
}

// UnsetInternalFormats ensures that no value is present for InternalFormats, not even an explicit nil
func (o *FilesSettingsDto) UnsetInternalFormats() {
	o.InternalFormats.Unset()
}

// GetMasterFormExtension returns the MasterFormExtension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetMasterFormExtension() string {
	if o == nil || IsNil(o.MasterFormExtension.Get()) {
		var ret string
		return ret
	}
	return *o.MasterFormExtension.Get()
}

// GetMasterFormExtensionOk returns a tuple with the MasterFormExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetMasterFormExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MasterFormExtension.Get(), o.MasterFormExtension.IsSet()
}

// HasMasterFormExtension returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsMasterFormExtensionSet() bool {
	if o != nil && o.MasterFormExtension.IsSet() {
		return true
	}

	return false
}

// SetMasterFormExtension gets a reference to the given NullableString and assigns it to the MasterFormExtension field.
func (o *FilesSettingsDto) SetMasterFormExtension(v string) {
	o.MasterFormExtension.Set(&v)
}
// SetMasterFormExtensionNil sets the value for MasterFormExtension to be an explicit nil
func (o *FilesSettingsDto) SetMasterFormExtensionNil() {
	o.MasterFormExtension.Set(nil)
}

// UnsetMasterFormExtension ensures that no value is present for MasterFormExtension, not even an explicit nil
func (o *FilesSettingsDto) UnsetMasterFormExtension() {
	o.MasterFormExtension.Unset()
}

// GetParamVersion returns the ParamVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetParamVersion() string {
	if o == nil || IsNil(o.ParamVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ParamVersion.Get()
}

// GetParamVersionOk returns a tuple with the ParamVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetParamVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParamVersion.Get(), o.ParamVersion.IsSet()
}

// HasParamVersion returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsParamVersionSet() bool {
	if o != nil && o.ParamVersion.IsSet() {
		return true
	}

	return false
}

// SetParamVersion gets a reference to the given NullableString and assigns it to the ParamVersion field.
func (o *FilesSettingsDto) SetParamVersion(v string) {
	o.ParamVersion.Set(&v)
}
// SetParamVersionNil sets the value for ParamVersion to be an explicit nil
func (o *FilesSettingsDto) SetParamVersionNil() {
	o.ParamVersion.Set(nil)
}

// UnsetParamVersion ensures that no value is present for ParamVersion, not even an explicit nil
func (o *FilesSettingsDto) UnsetParamVersion() {
	o.ParamVersion.Unset()
}

// GetParamOutType returns the ParamOutType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetParamOutType() string {
	if o == nil || IsNil(o.ParamOutType.Get()) {
		var ret string
		return ret
	}
	return *o.ParamOutType.Get()
}

// GetParamOutTypeOk returns a tuple with the ParamOutType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetParamOutTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParamOutType.Get(), o.ParamOutType.IsSet()
}

// HasParamOutType returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsParamOutTypeSet() bool {
	if o != nil && o.ParamOutType.IsSet() {
		return true
	}

	return false
}

// SetParamOutType gets a reference to the given NullableString and assigns it to the ParamOutType field.
func (o *FilesSettingsDto) SetParamOutType(v string) {
	o.ParamOutType.Set(&v)
}
// SetParamOutTypeNil sets the value for ParamOutType to be an explicit nil
func (o *FilesSettingsDto) SetParamOutTypeNil() {
	o.ParamOutType.Set(nil)
}

// UnsetParamOutType ensures that no value is present for ParamOutType, not even an explicit nil
func (o *FilesSettingsDto) UnsetParamOutType() {
	o.ParamOutType.Unset()
}

// GetFileDownloadUrlString returns the FileDownloadUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileDownloadUrlString() string {
	if o == nil || IsNil(o.FileDownloadUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileDownloadUrlString.Get()
}

// GetFileDownloadUrlStringOk returns a tuple with the FileDownloadUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileDownloadUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileDownloadUrlString.Get(), o.FileDownloadUrlString.IsSet()
}

// HasFileDownloadUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileDownloadUrlStringSet() bool {
	if o != nil && o.FileDownloadUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileDownloadUrlString gets a reference to the given NullableString and assigns it to the FileDownloadUrlString field.
func (o *FilesSettingsDto) SetFileDownloadUrlString(v string) {
	o.FileDownloadUrlString.Set(&v)
}
// SetFileDownloadUrlStringNil sets the value for FileDownloadUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileDownloadUrlStringNil() {
	o.FileDownloadUrlString.Set(nil)
}

// UnsetFileDownloadUrlString ensures that no value is present for FileDownloadUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileDownloadUrlString() {
	o.FileDownloadUrlString.Unset()
}

// GetFileWebViewerUrlString returns the FileWebViewerUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileWebViewerUrlString() string {
	if o == nil || IsNil(o.FileWebViewerUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileWebViewerUrlString.Get()
}

// GetFileWebViewerUrlStringOk returns a tuple with the FileWebViewerUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileWebViewerUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileWebViewerUrlString.Get(), o.FileWebViewerUrlString.IsSet()
}

// HasFileWebViewerUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileWebViewerUrlStringSet() bool {
	if o != nil && o.FileWebViewerUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileWebViewerUrlString gets a reference to the given NullableString and assigns it to the FileWebViewerUrlString field.
func (o *FilesSettingsDto) SetFileWebViewerUrlString(v string) {
	o.FileWebViewerUrlString.Set(&v)
}
// SetFileWebViewerUrlStringNil sets the value for FileWebViewerUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileWebViewerUrlStringNil() {
	o.FileWebViewerUrlString.Set(nil)
}

// UnsetFileWebViewerUrlString ensures that no value is present for FileWebViewerUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileWebViewerUrlString() {
	o.FileWebViewerUrlString.Unset()
}

// GetFileWebViewerExternalUrlString returns the FileWebViewerExternalUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileWebViewerExternalUrlString() string {
	if o == nil || IsNil(o.FileWebViewerExternalUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileWebViewerExternalUrlString.Get()
}

// GetFileWebViewerExternalUrlStringOk returns a tuple with the FileWebViewerExternalUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileWebViewerExternalUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileWebViewerExternalUrlString.Get(), o.FileWebViewerExternalUrlString.IsSet()
}

// HasFileWebViewerExternalUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileWebViewerExternalUrlStringSet() bool {
	if o != nil && o.FileWebViewerExternalUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileWebViewerExternalUrlString gets a reference to the given NullableString and assigns it to the FileWebViewerExternalUrlString field.
func (o *FilesSettingsDto) SetFileWebViewerExternalUrlString(v string) {
	o.FileWebViewerExternalUrlString.Set(&v)
}
// SetFileWebViewerExternalUrlStringNil sets the value for FileWebViewerExternalUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileWebViewerExternalUrlStringNil() {
	o.FileWebViewerExternalUrlString.Set(nil)
}

// UnsetFileWebViewerExternalUrlString ensures that no value is present for FileWebViewerExternalUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileWebViewerExternalUrlString() {
	o.FileWebViewerExternalUrlString.Unset()
}

// GetFileWebEditorUrlString returns the FileWebEditorUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileWebEditorUrlString() string {
	if o == nil || IsNil(o.FileWebEditorUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileWebEditorUrlString.Get()
}

// GetFileWebEditorUrlStringOk returns a tuple with the FileWebEditorUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileWebEditorUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileWebEditorUrlString.Get(), o.FileWebEditorUrlString.IsSet()
}

// HasFileWebEditorUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileWebEditorUrlStringSet() bool {
	if o != nil && o.FileWebEditorUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileWebEditorUrlString gets a reference to the given NullableString and assigns it to the FileWebEditorUrlString field.
func (o *FilesSettingsDto) SetFileWebEditorUrlString(v string) {
	o.FileWebEditorUrlString.Set(&v)
}
// SetFileWebEditorUrlStringNil sets the value for FileWebEditorUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileWebEditorUrlStringNil() {
	o.FileWebEditorUrlString.Set(nil)
}

// UnsetFileWebEditorUrlString ensures that no value is present for FileWebEditorUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileWebEditorUrlString() {
	o.FileWebEditorUrlString.Unset()
}

// GetFileWebEditorExternalUrlString returns the FileWebEditorExternalUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileWebEditorExternalUrlString() string {
	if o == nil || IsNil(o.FileWebEditorExternalUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileWebEditorExternalUrlString.Get()
}

// GetFileWebEditorExternalUrlStringOk returns a tuple with the FileWebEditorExternalUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileWebEditorExternalUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileWebEditorExternalUrlString.Get(), o.FileWebEditorExternalUrlString.IsSet()
}

// HasFileWebEditorExternalUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileWebEditorExternalUrlStringSet() bool {
	if o != nil && o.FileWebEditorExternalUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileWebEditorExternalUrlString gets a reference to the given NullableString and assigns it to the FileWebEditorExternalUrlString field.
func (o *FilesSettingsDto) SetFileWebEditorExternalUrlString(v string) {
	o.FileWebEditorExternalUrlString.Set(&v)
}
// SetFileWebEditorExternalUrlStringNil sets the value for FileWebEditorExternalUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileWebEditorExternalUrlStringNil() {
	o.FileWebEditorExternalUrlString.Set(nil)
}

// UnsetFileWebEditorExternalUrlString ensures that no value is present for FileWebEditorExternalUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileWebEditorExternalUrlString() {
	o.FileWebEditorExternalUrlString.Unset()
}

// GetFileRedirectPreviewUrlString returns the FileRedirectPreviewUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileRedirectPreviewUrlString() string {
	if o == nil || IsNil(o.FileRedirectPreviewUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileRedirectPreviewUrlString.Get()
}

// GetFileRedirectPreviewUrlStringOk returns a tuple with the FileRedirectPreviewUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileRedirectPreviewUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileRedirectPreviewUrlString.Get(), o.FileRedirectPreviewUrlString.IsSet()
}

// HasFileRedirectPreviewUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileRedirectPreviewUrlStringSet() bool {
	if o != nil && o.FileRedirectPreviewUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileRedirectPreviewUrlString gets a reference to the given NullableString and assigns it to the FileRedirectPreviewUrlString field.
func (o *FilesSettingsDto) SetFileRedirectPreviewUrlString(v string) {
	o.FileRedirectPreviewUrlString.Set(&v)
}
// SetFileRedirectPreviewUrlStringNil sets the value for FileRedirectPreviewUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileRedirectPreviewUrlStringNil() {
	o.FileRedirectPreviewUrlString.Set(nil)
}

// UnsetFileRedirectPreviewUrlString ensures that no value is present for FileRedirectPreviewUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileRedirectPreviewUrlString() {
	o.FileRedirectPreviewUrlString.Unset()
}

// GetFileThumbnailUrlString returns the FileThumbnailUrlString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetFileThumbnailUrlString() string {
	if o == nil || IsNil(o.FileThumbnailUrlString.Get()) {
		var ret string
		return ret
	}
	return *o.FileThumbnailUrlString.Get()
}

// GetFileThumbnailUrlStringOk returns a tuple with the FileThumbnailUrlString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetFileThumbnailUrlStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileThumbnailUrlString.Get(), o.FileThumbnailUrlString.IsSet()
}

// HasFileThumbnailUrlString returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFileThumbnailUrlStringSet() bool {
	if o != nil && o.FileThumbnailUrlString.IsSet() {
		return true
	}

	return false
}

// SetFileThumbnailUrlString gets a reference to the given NullableString and assigns it to the FileThumbnailUrlString field.
func (o *FilesSettingsDto) SetFileThumbnailUrlString(v string) {
	o.FileThumbnailUrlString.Set(&v)
}
// SetFileThumbnailUrlStringNil sets the value for FileThumbnailUrlString to be an explicit nil
func (o *FilesSettingsDto) SetFileThumbnailUrlStringNil() {
	o.FileThumbnailUrlString.Set(nil)
}

// UnsetFileThumbnailUrlString ensures that no value is present for FileThumbnailUrlString, not even an explicit nil
func (o *FilesSettingsDto) UnsetFileThumbnailUrlString() {
	o.FileThumbnailUrlString.Unset()
}

// GetConfirmDelete returns the ConfirmDelete field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetConfirmDelete() bool {
	if o == nil || IsNil(o.ConfirmDelete) {
		var ret bool
		return ret
	}
	return *o.ConfirmDelete
}

// GetConfirmDeleteOk returns a tuple with the ConfirmDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetConfirmDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfirmDelete) {
		return nil, false
	}
	return o.ConfirmDelete, true
}

// HasConfirmDelete returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsConfirmDeleteSet() bool {
	if o != nil && !IsNil(o.ConfirmDelete) {
		return true
	}

	return false
}

// SetConfirmDelete gets a reference to the given bool and assigns it to the ConfirmDelete field.
func (o *FilesSettingsDto) SetConfirmDelete(v bool) {
	o.ConfirmDelete = &v
}

// GetEnableThirdParty returns the EnableThirdParty field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetEnableThirdParty() bool {
	if o == nil || IsNil(o.EnableThirdParty) {
		var ret bool
		return ret
	}
	return *o.EnableThirdParty
}

// GetEnableThirdPartyOk returns a tuple with the EnableThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetEnableThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableThirdParty) {
		return nil, false
	}
	return o.EnableThirdParty, true
}

// HasEnableThirdParty returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsEnableThirdPartySet() bool {
	if o != nil && !IsNil(o.EnableThirdParty) {
		return true
	}

	return false
}

// SetEnableThirdParty gets a reference to the given bool and assigns it to the EnableThirdParty field.
func (o *FilesSettingsDto) SetEnableThirdParty(v bool) {
	o.EnableThirdParty = &v
}

// GetExternalShare returns the ExternalShare field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetExternalShare() bool {
	if o == nil || IsNil(o.ExternalShare) {
		var ret bool
		return ret
	}
	return *o.ExternalShare
}

// GetExternalShareOk returns a tuple with the ExternalShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetExternalShareOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShare) {
		return nil, false
	}
	return o.ExternalShare, true
}

// HasExternalShare returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExternalShareSet() bool {
	if o != nil && !IsNil(o.ExternalShare) {
		return true
	}

	return false
}

// SetExternalShare gets a reference to the given bool and assigns it to the ExternalShare field.
func (o *FilesSettingsDto) SetExternalShare(v bool) {
	o.ExternalShare = &v
}

// GetExternalShareSocialMedia returns the ExternalShareSocialMedia field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetExternalShareSocialMedia() bool {
	if o == nil || IsNil(o.ExternalShareSocialMedia) {
		var ret bool
		return ret
	}
	return *o.ExternalShareSocialMedia
}

// GetExternalShareSocialMediaOk returns a tuple with the ExternalShareSocialMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetExternalShareSocialMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareSocialMedia) {
		return nil, false
	}
	return o.ExternalShareSocialMedia, true
}

// HasExternalShareSocialMedia returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExternalShareSocialMediaSet() bool {
	if o != nil && !IsNil(o.ExternalShareSocialMedia) {
		return true
	}

	return false
}

// SetExternalShareSocialMedia gets a reference to the given bool and assigns it to the ExternalShareSocialMedia field.
func (o *FilesSettingsDto) SetExternalShareSocialMedia(v bool) {
	o.ExternalShareSocialMedia = &v
}

// GetStoreOriginalFiles returns the StoreOriginalFiles field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetStoreOriginalFiles() bool {
	if o == nil || IsNil(o.StoreOriginalFiles) {
		var ret bool
		return ret
	}
	return *o.StoreOriginalFiles
}

// GetStoreOriginalFilesOk returns a tuple with the StoreOriginalFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetStoreOriginalFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreOriginalFiles) {
		return nil, false
	}
	return o.StoreOriginalFiles, true
}

// HasStoreOriginalFiles returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsStoreOriginalFilesSet() bool {
	if o != nil && !IsNil(o.StoreOriginalFiles) {
		return true
	}

	return false
}

// SetStoreOriginalFiles gets a reference to the given bool and assigns it to the StoreOriginalFiles field.
func (o *FilesSettingsDto) SetStoreOriginalFiles(v bool) {
	o.StoreOriginalFiles = &v
}

// GetKeepNewFileName returns the KeepNewFileName field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetKeepNewFileName() bool {
	if o == nil || IsNil(o.KeepNewFileName) {
		var ret bool
		return ret
	}
	return *o.KeepNewFileName
}

// GetKeepNewFileNameOk returns a tuple with the KeepNewFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetKeepNewFileNameOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepNewFileName) {
		return nil, false
	}
	return o.KeepNewFileName, true
}

// HasKeepNewFileName returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsKeepNewFileNameSet() bool {
	if o != nil && !IsNil(o.KeepNewFileName) {
		return true
	}

	return false
}

// SetKeepNewFileName gets a reference to the given bool and assigns it to the KeepNewFileName field.
func (o *FilesSettingsDto) SetKeepNewFileName(v bool) {
	o.KeepNewFileName = &v
}

// GetDisplayFileExtension returns the DisplayFileExtension field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetDisplayFileExtension() bool {
	if o == nil || IsNil(o.DisplayFileExtension) {
		var ret bool
		return ret
	}
	return *o.DisplayFileExtension
}

// GetDisplayFileExtensionOk returns a tuple with the DisplayFileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetDisplayFileExtensionOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayFileExtension) {
		return nil, false
	}
	return o.DisplayFileExtension, true
}

// HasDisplayFileExtension returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsDisplayFileExtensionSet() bool {
	if o != nil && !IsNil(o.DisplayFileExtension) {
		return true
	}

	return false
}

// SetDisplayFileExtension gets a reference to the given bool and assigns it to the DisplayFileExtension field.
func (o *FilesSettingsDto) SetDisplayFileExtension(v bool) {
	o.DisplayFileExtension = &v
}

// GetShowQuickActions returns the ShowQuickActions field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetShowQuickActions() bool {
	if o == nil || IsNil(o.ShowQuickActions) {
		var ret bool
		return ret
	}
	return *o.ShowQuickActions
}

// GetShowQuickActionsOk returns a tuple with the ShowQuickActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetShowQuickActionsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowQuickActions) {
		return nil, false
	}
	return o.ShowQuickActions, true
}

// HasShowQuickActions returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsShowQuickActionsSet() bool {
	if o != nil && !IsNil(o.ShowQuickActions) {
		return true
	}

	return false
}

// SetShowQuickActions gets a reference to the given bool and assigns it to the ShowQuickActions field.
func (o *FilesSettingsDto) SetShowQuickActions(v bool) {
	o.ShowQuickActions = &v
}

// GetConvertNotify returns the ConvertNotify field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetConvertNotify() bool {
	if o == nil || IsNil(o.ConvertNotify) {
		var ret bool
		return ret
	}
	return *o.ConvertNotify
}

// GetConvertNotifyOk returns a tuple with the ConvertNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetConvertNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.ConvertNotify) {
		return nil, false
	}
	return o.ConvertNotify, true
}

// HasConvertNotify returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsConvertNotifySet() bool {
	if o != nil && !IsNil(o.ConvertNotify) {
		return true
	}

	return false
}

// SetConvertNotify gets a reference to the given bool and assigns it to the ConvertNotify field.
func (o *FilesSettingsDto) SetConvertNotify(v bool) {
	o.ConvertNotify = &v
}

// GetHideConfirmCancelOperation returns the HideConfirmCancelOperation field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetHideConfirmCancelOperation() bool {
	if o == nil || IsNil(o.HideConfirmCancelOperation) {
		var ret bool
		return ret
	}
	return *o.HideConfirmCancelOperation
}

// GetHideConfirmCancelOperationOk returns a tuple with the HideConfirmCancelOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetHideConfirmCancelOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.HideConfirmCancelOperation) {
		return nil, false
	}
	return o.HideConfirmCancelOperation, true
}

// HasHideConfirmCancelOperation returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsHideConfirmCancelOperationSet() bool {
	if o != nil && !IsNil(o.HideConfirmCancelOperation) {
		return true
	}

	return false
}

// SetHideConfirmCancelOperation gets a reference to the given bool and assigns it to the HideConfirmCancelOperation field.
func (o *FilesSettingsDto) SetHideConfirmCancelOperation(v bool) {
	o.HideConfirmCancelOperation = &v
}

// GetHideConfirmConvertSave returns the HideConfirmConvertSave field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetHideConfirmConvertSave() bool {
	if o == nil || IsNil(o.HideConfirmConvertSave) {
		var ret bool
		return ret
	}
	return *o.HideConfirmConvertSave
}

// GetHideConfirmConvertSaveOk returns a tuple with the HideConfirmConvertSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetHideConfirmConvertSaveOk() (*bool, bool) {
	if o == nil || IsNil(o.HideConfirmConvertSave) {
		return nil, false
	}
	return o.HideConfirmConvertSave, true
}

// HasHideConfirmConvertSave returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsHideConfirmConvertSaveSet() bool {
	if o != nil && !IsNil(o.HideConfirmConvertSave) {
		return true
	}

	return false
}

// SetHideConfirmConvertSave gets a reference to the given bool and assigns it to the HideConfirmConvertSave field.
func (o *FilesSettingsDto) SetHideConfirmConvertSave(v bool) {
	o.HideConfirmConvertSave = &v
}

// GetHideConfirmConvertOpen returns the HideConfirmConvertOpen field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetHideConfirmConvertOpen() bool {
	if o == nil || IsNil(o.HideConfirmConvertOpen) {
		var ret bool
		return ret
	}
	return *o.HideConfirmConvertOpen
}

// GetHideConfirmConvertOpenOk returns a tuple with the HideConfirmConvertOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetHideConfirmConvertOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.HideConfirmConvertOpen) {
		return nil, false
	}
	return o.HideConfirmConvertOpen, true
}

// HasHideConfirmConvertOpen returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsHideConfirmConvertOpenSet() bool {
	if o != nil && !IsNil(o.HideConfirmConvertOpen) {
		return true
	}

	return false
}

// SetHideConfirmConvertOpen gets a reference to the given bool and assigns it to the HideConfirmConvertOpen field.
func (o *FilesSettingsDto) SetHideConfirmConvertOpen(v bool) {
	o.HideConfirmConvertOpen = &v
}

// GetHideConfirmRoomLifetime returns the HideConfirmRoomLifetime field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetHideConfirmRoomLifetime() bool {
	if o == nil || IsNil(o.HideConfirmRoomLifetime) {
		var ret bool
		return ret
	}
	return *o.HideConfirmRoomLifetime
}

// GetHideConfirmRoomLifetimeOk returns a tuple with the HideConfirmRoomLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetHideConfirmRoomLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.HideConfirmRoomLifetime) {
		return nil, false
	}
	return o.HideConfirmRoomLifetime, true
}

// HasHideConfirmRoomLifetime returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsHideConfirmRoomLifetimeSet() bool {
	if o != nil && !IsNil(o.HideConfirmRoomLifetime) {
		return true
	}

	return false
}

// SetHideConfirmRoomLifetime gets a reference to the given bool and assigns it to the HideConfirmRoomLifetime field.
func (o *FilesSettingsDto) SetHideConfirmRoomLifetime(v bool) {
	o.HideConfirmRoomLifetime = &v
}

// GetDefaultOrder returns the DefaultOrder field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetDefaultOrder() OrderBy {
	if o == nil || IsNil(o.DefaultOrder) {
		var ret OrderBy
		return ret
	}
	return *o.DefaultOrder
}

// GetDefaultOrderOk returns a tuple with the DefaultOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetDefaultOrderOk() (*OrderBy, bool) {
	if o == nil || IsNil(o.DefaultOrder) {
		return nil, false
	}
	return o.DefaultOrder, true
}

// HasDefaultOrder returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsDefaultOrderSet() bool {
	if o != nil && !IsNil(o.DefaultOrder) {
		return true
	}

	return false
}

// SetDefaultOrder gets a reference to the given OrderBy and assigns it to the DefaultOrder field.
func (o *FilesSettingsDto) SetDefaultOrder(v OrderBy) {
	o.DefaultOrder = &v
}

// GetForcesave returns the Forcesave field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetForcesave() bool {
	if o == nil || IsNil(o.Forcesave) {
		var ret bool
		return ret
	}
	return *o.Forcesave
}

// GetForcesaveOk returns a tuple with the Forcesave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetForcesaveOk() (*bool, bool) {
	if o == nil || IsNil(o.Forcesave) {
		return nil, false
	}
	return o.Forcesave, true
}

// HasForcesave returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsForcesaveSet() bool {
	if o != nil && !IsNil(o.Forcesave) {
		return true
	}

	return false
}

// SetForcesave gets a reference to the given bool and assigns it to the Forcesave field.
func (o *FilesSettingsDto) SetForcesave(v bool) {
	o.Forcesave = &v
}

// GetStoreForcesave returns the StoreForcesave field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetStoreForcesave() bool {
	if o == nil || IsNil(o.StoreForcesave) {
		var ret bool
		return ret
	}
	return *o.StoreForcesave
}

// GetStoreForcesaveOk returns a tuple with the StoreForcesave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetStoreForcesaveOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreForcesave) {
		return nil, false
	}
	return o.StoreForcesave, true
}

// HasStoreForcesave returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsStoreForcesaveSet() bool {
	if o != nil && !IsNil(o.StoreForcesave) {
		return true
	}

	return false
}

// SetStoreForcesave gets a reference to the given bool and assigns it to the StoreForcesave field.
func (o *FilesSettingsDto) SetStoreForcesave(v bool) {
	o.StoreForcesave = &v
}

// GetRecentSection returns the RecentSection field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetRecentSection() bool {
	if o == nil || IsNil(o.RecentSection) {
		var ret bool
		return ret
	}
	return *o.RecentSection
}

// GetRecentSectionOk returns a tuple with the RecentSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetRecentSectionOk() (*bool, bool) {
	if o == nil || IsNil(o.RecentSection) {
		return nil, false
	}
	return o.RecentSection, true
}

// HasRecentSection returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsRecentSectionSet() bool {
	if o != nil && !IsNil(o.RecentSection) {
		return true
	}

	return false
}

// SetRecentSection gets a reference to the given bool and assigns it to the RecentSection field.
func (o *FilesSettingsDto) SetRecentSection(v bool) {
	o.RecentSection = &v
}

// GetFavoritesSection returns the FavoritesSection field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetFavoritesSection() bool {
	if o == nil || IsNil(o.FavoritesSection) {
		var ret bool
		return ret
	}
	return *o.FavoritesSection
}

// GetFavoritesSectionOk returns a tuple with the FavoritesSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetFavoritesSectionOk() (*bool, bool) {
	if o == nil || IsNil(o.FavoritesSection) {
		return nil, false
	}
	return o.FavoritesSection, true
}

// HasFavoritesSection returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsFavoritesSectionSet() bool {
	if o != nil && !IsNil(o.FavoritesSection) {
		return true
	}

	return false
}

// SetFavoritesSection gets a reference to the given bool and assigns it to the FavoritesSection field.
func (o *FilesSettingsDto) SetFavoritesSection(v bool) {
	o.FavoritesSection = &v
}

// GetTemplatesSection returns the TemplatesSection field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetTemplatesSection() bool {
	if o == nil || IsNil(o.TemplatesSection) {
		var ret bool
		return ret
	}
	return *o.TemplatesSection
}

// GetTemplatesSectionOk returns a tuple with the TemplatesSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetTemplatesSectionOk() (*bool, bool) {
	if o == nil || IsNil(o.TemplatesSection) {
		return nil, false
	}
	return o.TemplatesSection, true
}

// HasTemplatesSection returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsTemplatesSectionSet() bool {
	if o != nil && !IsNil(o.TemplatesSection) {
		return true
	}

	return false
}

// SetTemplatesSection gets a reference to the given bool and assigns it to the TemplatesSection field.
func (o *FilesSettingsDto) SetTemplatesSection(v bool) {
	o.TemplatesSection = &v
}

// GetDownloadTarGz returns the DownloadTarGz field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetDownloadTarGz() bool {
	if o == nil || IsNil(o.DownloadTarGz) {
		var ret bool
		return ret
	}
	return *o.DownloadTarGz
}

// GetDownloadTarGzOk returns a tuple with the DownloadTarGz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetDownloadTarGzOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadTarGz) {
		return nil, false
	}
	return o.DownloadTarGz, true
}

// HasDownloadTarGz returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsDownloadTarGzSet() bool {
	if o != nil && !IsNil(o.DownloadTarGz) {
		return true
	}

	return false
}

// SetDownloadTarGz gets a reference to the given bool and assigns it to the DownloadTarGz field.
func (o *FilesSettingsDto) SetDownloadTarGz(v bool) {
	o.DownloadTarGz = &v
}

// GetAutomaticallyCleanUp returns the AutomaticallyCleanUp field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetAutomaticallyCleanUp() AutoCleanUpData {
	if o == nil || IsNil(o.AutomaticallyCleanUp) {
		var ret AutoCleanUpData
		return ret
	}
	return *o.AutomaticallyCleanUp
}

// GetAutomaticallyCleanUpOk returns a tuple with the AutomaticallyCleanUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetAutomaticallyCleanUpOk() (*AutoCleanUpData, bool) {
	if o == nil || IsNil(o.AutomaticallyCleanUp) {
		return nil, false
	}
	return o.AutomaticallyCleanUp, true
}

// HasAutomaticallyCleanUp returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsAutomaticallyCleanUpSet() bool {
	if o != nil && !IsNil(o.AutomaticallyCleanUp) {
		return true
	}

	return false
}

// SetAutomaticallyCleanUp gets a reference to the given AutoCleanUpData and assigns it to the AutomaticallyCleanUp field.
func (o *FilesSettingsDto) SetAutomaticallyCleanUp(v AutoCleanUpData) {
	o.AutomaticallyCleanUp = &v
}

// GetCanSearchByContent returns the CanSearchByContent field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetCanSearchByContent() bool {
	if o == nil || IsNil(o.CanSearchByContent) {
		var ret bool
		return ret
	}
	return *o.CanSearchByContent
}

// GetCanSearchByContentOk returns a tuple with the CanSearchByContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetCanSearchByContentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSearchByContent) {
		return nil, false
	}
	return o.CanSearchByContent, true
}

// HasCanSearchByContent returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsCanSearchByContentSet() bool {
	if o != nil && !IsNil(o.CanSearchByContent) {
		return true
	}

	return false
}

// SetCanSearchByContent gets a reference to the given bool and assigns it to the CanSearchByContent field.
func (o *FilesSettingsDto) SetCanSearchByContent(v bool) {
	o.CanSearchByContent = &v
}

// GetDefaultSharingAccessRights returns the DefaultSharingAccessRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetDefaultSharingAccessRights() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.DefaultSharingAccessRights
}

// GetDefaultSharingAccessRightsOk returns a tuple with the DefaultSharingAccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetDefaultSharingAccessRightsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DefaultSharingAccessRights) {
		return nil, false
	}
	return o.DefaultSharingAccessRights, true
}

// HasDefaultSharingAccessRights returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsDefaultSharingAccessRightsSet() bool {
	if o != nil && !IsNil(o.DefaultSharingAccessRights) {
		return true
	}

	return false
}

// SetDefaultSharingAccessRights gets a reference to the given []int32 and assigns it to the DefaultSharingAccessRights field.
func (o *FilesSettingsDto) SetDefaultSharingAccessRights(v []int32) {
	o.DefaultSharingAccessRights = v
}

// GetMaxUploadThreadCount returns the MaxUploadThreadCount field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetMaxUploadThreadCount() int32 {
	if o == nil || IsNil(o.MaxUploadThreadCount) {
		var ret int32
		return ret
	}
	return *o.MaxUploadThreadCount
}

// GetMaxUploadThreadCountOk returns a tuple with the MaxUploadThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetMaxUploadThreadCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUploadThreadCount) {
		return nil, false
	}
	return o.MaxUploadThreadCount, true
}

// HasMaxUploadThreadCount returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsMaxUploadThreadCountSet() bool {
	if o != nil && !IsNil(o.MaxUploadThreadCount) {
		return true
	}

	return false
}

// SetMaxUploadThreadCount gets a reference to the given int32 and assigns it to the MaxUploadThreadCount field.
func (o *FilesSettingsDto) SetMaxUploadThreadCount(v int32) {
	o.MaxUploadThreadCount = &v
}

// GetChunkUploadSize returns the ChunkUploadSize field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetChunkUploadSize() int64 {
	if o == nil || IsNil(o.ChunkUploadSize) {
		var ret int64
		return ret
	}
	return *o.ChunkUploadSize
}

// GetChunkUploadSizeOk returns a tuple with the ChunkUploadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetChunkUploadSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ChunkUploadSize) {
		return nil, false
	}
	return o.ChunkUploadSize, true
}

// HasChunkUploadSize returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsChunkUploadSizeSet() bool {
	if o != nil && !IsNil(o.ChunkUploadSize) {
		return true
	}

	return false
}

// SetChunkUploadSize gets a reference to the given int64 and assigns it to the ChunkUploadSize field.
func (o *FilesSettingsDto) SetChunkUploadSize(v int64) {
	o.ChunkUploadSize = &v
}

// GetOpenEditorInSameTab returns the OpenEditorInSameTab field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetOpenEditorInSameTab() bool {
	if o == nil || IsNil(o.OpenEditorInSameTab) {
		var ret bool
		return ret
	}
	return *o.OpenEditorInSameTab
}

// GetOpenEditorInSameTabOk returns a tuple with the OpenEditorInSameTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetOpenEditorInSameTabOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenEditorInSameTab) {
		return nil, false
	}
	return o.OpenEditorInSameTab, true
}

// HasOpenEditorInSameTab returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsOpenEditorInSameTabSet() bool {
	if o != nil && !IsNil(o.OpenEditorInSameTab) {
		return true
	}

	return false
}

// SetOpenEditorInSameTab gets a reference to the given bool and assigns it to the OpenEditorInSameTab field.
func (o *FilesSettingsDto) SetOpenEditorInSameTab(v bool) {
	o.OpenEditorInSameTab = &v
}

// GetOrganizeRoomsGrouping returns the OrganizeRoomsGrouping field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetOrganizeRoomsGrouping() bool {
	if o == nil || IsNil(o.OrganizeRoomsGrouping) {
		var ret bool
		return ret
	}
	return *o.OrganizeRoomsGrouping
}

// GetOrganizeRoomsGroupingOk returns a tuple with the OrganizeRoomsGrouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetOrganizeRoomsGroupingOk() (*bool, bool) {
	if o == nil || IsNil(o.OrganizeRoomsGrouping) {
		return nil, false
	}
	return o.OrganizeRoomsGrouping, true
}

// HasOrganizeRoomsGrouping returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsOrganizeRoomsGroupingSet() bool {
	if o != nil && !IsNil(o.OrganizeRoomsGrouping) {
		return true
	}

	return false
}

// SetOrganizeRoomsGrouping gets a reference to the given bool and assigns it to the OrganizeRoomsGrouping field.
func (o *FilesSettingsDto) SetOrganizeRoomsGrouping(v bool) {
	o.OrganizeRoomsGrouping = &v
}

// GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetDefaultShareLinkInternal() bool {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		var ret bool
		return ret
	}
	return *o.DefaultShareLinkInternal
}

// GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetDefaultShareLinkInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		return nil, false
	}
	return o.DefaultShareLinkInternal, true
}

// HasDefaultShareLinkInternal returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsDefaultShareLinkInternalSet() bool {
	if o != nil && !IsNil(o.DefaultShareLinkInternal) {
		return true
	}

	return false
}

// SetDefaultShareLinkInternal gets a reference to the given bool and assigns it to the DefaultShareLinkInternal field.
func (o *FilesSettingsDto) SetDefaultShareLinkInternal(v bool) {
	o.DefaultShareLinkInternal = &v
}

// GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetExternalShareApplyToDocuments() bool {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToDocuments
}

// GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetExternalShareApplyToDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		return nil, false
	}
	return o.ExternalShareApplyToDocuments, true
}

// HasExternalShareApplyToDocuments returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExternalShareApplyToDocumentsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToDocuments) {
		return true
	}

	return false
}

// SetExternalShareApplyToDocuments gets a reference to the given bool and assigns it to the ExternalShareApplyToDocuments field.
func (o *FilesSettingsDto) SetExternalShareApplyToDocuments(v bool) {
	o.ExternalShareApplyToDocuments = &v
}

// GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetExternalShareApplyToRooms() bool {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToRooms
}

// GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetExternalShareApplyToRoomsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		return nil, false
	}
	return o.ExternalShareApplyToRooms, true
}

// HasExternalShareApplyToRooms returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExternalShareApplyToRoomsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToRooms) {
		return true
	}

	return false
}

// SetExternalShareApplyToRooms gets a reference to the given bool and assigns it to the ExternalShareApplyToRooms field.
func (o *FilesSettingsDto) SetExternalShareApplyToRooms(v bool) {
	o.ExternalShareApplyToRooms = &v
}

// GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetBlockExistingLinksOnRestrict() bool {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		var ret bool
		return ret
	}
	return *o.BlockExistingLinksOnRestrict
}

// GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		return nil, false
	}
	return o.BlockExistingLinksOnRestrict, true
}

// HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsBlockExistingLinksOnRestrictSet() bool {
	if o != nil && !IsNil(o.BlockExistingLinksOnRestrict) {
		return true
	}

	return false
}

// SetBlockExistingLinksOnRestrict gets a reference to the given bool and assigns it to the BlockExistingLinksOnRestrict field.
func (o *FilesSettingsDto) SetBlockExistingLinksOnRestrict(v bool) {
	o.BlockExistingLinksOnRestrict = &v
}

// GetExtsFilesVectorized returns the ExtsFilesVectorized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesSettingsDto) GetExtsFilesVectorized() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtsFilesVectorized
}

// GetExtsFilesVectorizedOk returns a tuple with the ExtsFilesVectorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSettingsDto) GetExtsFilesVectorizedOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtsFilesVectorized) {
		return nil, false
	}
	return o.ExtsFilesVectorized, true
}

// HasExtsFilesVectorized returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsExtsFilesVectorizedSet() bool {
	if o != nil && !IsNil(o.ExtsFilesVectorized) {
		return true
	}

	return false
}

// SetExtsFilesVectorized gets a reference to the given []string and assigns it to the ExtsFilesVectorized field.
func (o *FilesSettingsDto) SetExtsFilesVectorized(v []string) {
	o.ExtsFilesVectorized = v
}

// GetMaxVectorizationFileSize returns the MaxVectorizationFileSize field value if set, zero value otherwise.
func (o *FilesSettingsDto) GetMaxVectorizationFileSize() int64 {
	if o == nil || IsNil(o.MaxVectorizationFileSize) {
		var ret int64
		return ret
	}
	return *o.MaxVectorizationFileSize
}

// GetMaxVectorizationFileSizeOk returns a tuple with the MaxVectorizationFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSettingsDto) GetMaxVectorizationFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxVectorizationFileSize) {
		return nil, false
	}
	return o.MaxVectorizationFileSize, true
}

// HasMaxVectorizationFileSize returns a boolean if a field has been set.
func (o *FilesSettingsDto) IsMaxVectorizationFileSizeSet() bool {
	if o != nil && !IsNil(o.MaxVectorizationFileSize) {
		return true
	}

	return false
}

// SetMaxVectorizationFileSize gets a reference to the given int64 and assigns it to the MaxVectorizationFileSize field.
func (o *FilesSettingsDto) SetMaxVectorizationFileSize(v int64) {
	o.MaxVectorizationFileSize = &v
}

func (o FilesSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtsImagePreviewed != nil {
		toSerialize["extsImagePreviewed"] = o.ExtsImagePreviewed
	}
	if o.ExtsMediaPreviewed != nil {
		toSerialize["extsMediaPreviewed"] = o.ExtsMediaPreviewed
	}
	if o.ExtsWebPreviewed != nil {
		toSerialize["extsWebPreviewed"] = o.ExtsWebPreviewed
	}
	if o.ExtsWebEdited != nil {
		toSerialize["extsWebEdited"] = o.ExtsWebEdited
	}
	if o.ExtsWebEncrypt != nil {
		toSerialize["extsWebEncrypt"] = o.ExtsWebEncrypt
	}
	if o.ExtsWebReviewed != nil {
		toSerialize["extsWebReviewed"] = o.ExtsWebReviewed
	}
	if o.ExtsWebCustomFilterEditing != nil {
		toSerialize["extsWebCustomFilterEditing"] = o.ExtsWebCustomFilterEditing
	}
	if o.ExtsWebRestrictedEditing != nil {
		toSerialize["extsWebRestrictedEditing"] = o.ExtsWebRestrictedEditing
	}
	if o.ExtsWebCommented != nil {
		toSerialize["extsWebCommented"] = o.ExtsWebCommented
	}
	if o.ExtsWebTemplate != nil {
		toSerialize["extsWebTemplate"] = o.ExtsWebTemplate
	}
	if o.ExtsMustConvert != nil {
		toSerialize["extsMustConvert"] = o.ExtsMustConvert
	}
	if !IsNil(o.ExtsConvertible) {
		toSerialize["extsConvertible"] = o.ExtsConvertible
	}
	if o.ExtsUploadable != nil {
		toSerialize["extsUploadable"] = o.ExtsUploadable
	}
	if o.ExtsArchive != nil {
		toSerialize["extsArchive"] = o.ExtsArchive
	}
	if o.ExtsVideo != nil {
		toSerialize["extsVideo"] = o.ExtsVideo
	}
	if o.ExtsAudio != nil {
		toSerialize["extsAudio"] = o.ExtsAudio
	}
	if o.ExtsImage != nil {
		toSerialize["extsImage"] = o.ExtsImage
	}
	if o.ExtsSpreadsheet != nil {
		toSerialize["extsSpreadsheet"] = o.ExtsSpreadsheet
	}
	if o.ExtsPresentation != nil {
		toSerialize["extsPresentation"] = o.ExtsPresentation
	}
	if o.ExtsDocument != nil {
		toSerialize["extsDocument"] = o.ExtsDocument
	}
	if o.ExtsDiagram != nil {
		toSerialize["extsDiagram"] = o.ExtsDiagram
	}
	if o.InternalFormats.IsSet() {
		toSerialize["internalFormats"] = o.InternalFormats.Get()
	}
	if o.MasterFormExtension.IsSet() {
		toSerialize["masterFormExtension"] = o.MasterFormExtension.Get()
	}
	if o.ParamVersion.IsSet() {
		toSerialize["paramVersion"] = o.ParamVersion.Get()
	}
	if o.ParamOutType.IsSet() {
		toSerialize["paramOutType"] = o.ParamOutType.Get()
	}
	if o.FileDownloadUrlString.IsSet() {
		toSerialize["fileDownloadUrlString"] = o.FileDownloadUrlString.Get()
	}
	if o.FileWebViewerUrlString.IsSet() {
		toSerialize["fileWebViewerUrlString"] = o.FileWebViewerUrlString.Get()
	}
	if o.FileWebViewerExternalUrlString.IsSet() {
		toSerialize["fileWebViewerExternalUrlString"] = o.FileWebViewerExternalUrlString.Get()
	}
	if o.FileWebEditorUrlString.IsSet() {
		toSerialize["fileWebEditorUrlString"] = o.FileWebEditorUrlString.Get()
	}
	if o.FileWebEditorExternalUrlString.IsSet() {
		toSerialize["fileWebEditorExternalUrlString"] = o.FileWebEditorExternalUrlString.Get()
	}
	if o.FileRedirectPreviewUrlString.IsSet() {
		toSerialize["fileRedirectPreviewUrlString"] = o.FileRedirectPreviewUrlString.Get()
	}
	if o.FileThumbnailUrlString.IsSet() {
		toSerialize["fileThumbnailUrlString"] = o.FileThumbnailUrlString.Get()
	}
	if !IsNil(o.ConfirmDelete) {
		toSerialize["confirmDelete"] = o.ConfirmDelete
	}
	if !IsNil(o.EnableThirdParty) {
		toSerialize["enableThirdParty"] = o.EnableThirdParty
	}
	if !IsNil(o.ExternalShare) {
		toSerialize["externalShare"] = o.ExternalShare
	}
	if !IsNil(o.ExternalShareSocialMedia) {
		toSerialize["externalShareSocialMedia"] = o.ExternalShareSocialMedia
	}
	if !IsNil(o.StoreOriginalFiles) {
		toSerialize["storeOriginalFiles"] = o.StoreOriginalFiles
	}
	if !IsNil(o.KeepNewFileName) {
		toSerialize["keepNewFileName"] = o.KeepNewFileName
	}
	if !IsNil(o.DisplayFileExtension) {
		toSerialize["displayFileExtension"] = o.DisplayFileExtension
	}
	if !IsNil(o.ShowQuickActions) {
		toSerialize["showQuickActions"] = o.ShowQuickActions
	}
	if !IsNil(o.ConvertNotify) {
		toSerialize["convertNotify"] = o.ConvertNotify
	}
	if !IsNil(o.HideConfirmCancelOperation) {
		toSerialize["hideConfirmCancelOperation"] = o.HideConfirmCancelOperation
	}
	if !IsNil(o.HideConfirmConvertSave) {
		toSerialize["hideConfirmConvertSave"] = o.HideConfirmConvertSave
	}
	if !IsNil(o.HideConfirmConvertOpen) {
		toSerialize["hideConfirmConvertOpen"] = o.HideConfirmConvertOpen
	}
	if !IsNil(o.HideConfirmRoomLifetime) {
		toSerialize["hideConfirmRoomLifetime"] = o.HideConfirmRoomLifetime
	}
	if !IsNil(o.DefaultOrder) {
		toSerialize["defaultOrder"] = o.DefaultOrder
	}
	if !IsNil(o.Forcesave) {
		toSerialize["forcesave"] = o.Forcesave
	}
	if !IsNil(o.StoreForcesave) {
		toSerialize["storeForcesave"] = o.StoreForcesave
	}
	if !IsNil(o.RecentSection) {
		toSerialize["recentSection"] = o.RecentSection
	}
	if !IsNil(o.FavoritesSection) {
		toSerialize["favoritesSection"] = o.FavoritesSection
	}
	if !IsNil(o.TemplatesSection) {
		toSerialize["templatesSection"] = o.TemplatesSection
	}
	if !IsNil(o.DownloadTarGz) {
		toSerialize["downloadTarGz"] = o.DownloadTarGz
	}
	if !IsNil(o.AutomaticallyCleanUp) {
		toSerialize["automaticallyCleanUp"] = o.AutomaticallyCleanUp
	}
	if !IsNil(o.CanSearchByContent) {
		toSerialize["canSearchByContent"] = o.CanSearchByContent
	}
	if o.DefaultSharingAccessRights != nil {
		toSerialize["defaultSharingAccessRights"] = o.DefaultSharingAccessRights
	}
	if !IsNil(o.MaxUploadThreadCount) {
		toSerialize["maxUploadThreadCount"] = o.MaxUploadThreadCount
	}
	if !IsNil(o.ChunkUploadSize) {
		toSerialize["chunkUploadSize"] = o.ChunkUploadSize
	}
	if !IsNil(o.OpenEditorInSameTab) {
		toSerialize["openEditorInSameTab"] = o.OpenEditorInSameTab
	}
	if !IsNil(o.OrganizeRoomsGrouping) {
		toSerialize["organizeRoomsGrouping"] = o.OrganizeRoomsGrouping
	}
	if !IsNil(o.DefaultShareLinkInternal) {
		toSerialize["defaultShareLinkInternal"] = o.DefaultShareLinkInternal
	}
	if !IsNil(o.ExternalShareApplyToDocuments) {
		toSerialize["externalShareApplyToDocuments"] = o.ExternalShareApplyToDocuments
	}
	if !IsNil(o.ExternalShareApplyToRooms) {
		toSerialize["externalShareApplyToRooms"] = o.ExternalShareApplyToRooms
	}
	if !IsNil(o.BlockExistingLinksOnRestrict) {
		toSerialize["blockExistingLinksOnRestrict"] = o.BlockExistingLinksOnRestrict
	}
	if o.ExtsFilesVectorized != nil {
		toSerialize["extsFilesVectorized"] = o.ExtsFilesVectorized
	}
	if !IsNil(o.MaxVectorizationFileSize) {
		toSerialize["maxVectorizationFileSize"] = o.MaxVectorizationFileSize
	}
	return toSerialize, nil
}

type NullableFilesSettingsDto struct {
	value *FilesSettingsDto
	isSet bool
}

func (v NullableFilesSettingsDto) Get() *FilesSettingsDto {
	return v.value
}

func (v *NullableFilesSettingsDto) Set(val *FilesSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSettingsDto(val *FilesSettingsDto) *NullableFilesSettingsDto {
	return &NullableFilesSettingsDto{value: val, isSet: true}
}

func (v NullableFilesSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

