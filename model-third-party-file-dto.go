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

// checks if the ThirdPartyFileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyFileDto{}

// ThirdPartyFileDto A stored file as the calling account sees it: where it lives, which revision this is, how it can be opened and  what the portal is currently doing with it.
type ThirdPartyFileDto struct {
	ThirdPartyFileEntryDto
	// The folder the file is stored in. When the file was reached through a share and the caller cannot open its  real parent, the identifier of the Shared with me section is reported instead, so this is where the file is  visible rather than where it physically sits.
	FolderId NullableString `json:"folderId,omitempty"`
	// The revision this entry describes. It starts at 1 and moves to the next number each time new content is stored  over the file, except for an editing session opened against the file itself, which replaces the content and  keeps the number. `GET api/2.0/files/file/{fileId}/history` lists them all.
	Version *int32 `json:"version,omitempty"`
	// Groups revisions that belong together, which is how a history can fold a long editing session into one entry:  versions saved inside one session share this number, and an upload over the file starts a new group.
	VersionGroup *int32 `json:"versionGroup,omitempty"`
	// The size already formatted for display, with a unit and the separators of the caller's language. Read  `pureContentLength` for a number to calculate with.
	ContentLength NullableString `json:"contentLength,omitempty"`
	// The size of the stored content in bytes, and null for an empty file.
	PureContentLength NullableInt64 `json:"pureContentLength,omitempty"`
	// What the portal is currently doing with the file and how the caller stands towards it - open in the editor,  unread, being converted, and so on. The value is a bit mask that combines those states, so a file can report a  number that matches none of the published members on its own.
	FileStatus *FileStatus `json:"fileStatus,omitempty"`
	// The accounts that have the file open in the editor at this moment, as account identifier to display name, and  empty when nobody has. The all-zero identifier stands for people who came in through an external link without  signing in, and its name carries their number in brackets when there is more than one.
	EditingBy map[string]*string `json:"editingBy,omitempty"`
	// Not a property of the file at all: it repeats, inverted, the calling account's own switch for new-item badges,  so it is the same in every entry of one answer. True means that account has badges turned off.
	Mute *bool `json:"mute,omitempty"`
	// The address that returns the bytes of the file - a download, in spite of the name; `webUrl` is the address a  person opens. When the file was reached through an external link the address carries the key of that link, so  it keeps working without signing in.
	ViewUrl NullableString `json:"viewUrl,omitempty"`
	// The page that opens the file in a browser: the editor for a format the portal edits, the media viewer for  pictures, audio and video, and the download address for a format it cannot show at all.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// The broad kind of content, worked out from the extension, which is what a client uses to pick an icon or a  viewer without parsing `fileExst` itself.
	FileType *FileType `json:"fileType,omitempty"`
	// The extension of the stored file, leading dot included and always lower case. For a format the portal keeps in  a converted shape this is the extension it is served under, not the one it was uploaded with.
	FileExst NullableString `json:"fileExst,omitempty"`
	// The note kept with this revision. The portal writes it itself for revisions it creates, an upload over an  existing file among them, and an editor stores the note a person typed when saving a version.
	Comment NullableString `json:"comment,omitempty"`
	// True for a file in a private room, whose content the server never sees and which therefore cannot be converted  or taken over by an upload. Null, rather than false, for an ordinary file.
	Encrypted NullableBool `json:"encrypted,omitempty"`
	// The address of the generated preview image. It is filled in only while `thumbnailStatus` says the preview has  been created, and it carries a suffix that changes with the file, so an image cached for an earlier revision  is not reused.
	ThumbnailUrl NullableString `json:"thumbnailUrl,omitempty"`
	// How far the preview image has got. Only the created state means `thumbnailUrl` holds an address; the others  mean there is none, either because it is still being produced or because this format has no preview.
	ThumbnailStatus *Thumbnail `json:"thumbnailStatus,omitempty"`
	// True while the file is held under a lock that stops anyone but its holder from editing it, and null rather  than false when there is no lock. `lockedBy` names the holder unless the caller is the holder.
	Locked NullableBool `json:"locked,omitempty"`
	// The display name of the account holding the lock, and null when the caller holds it - so `locked` true  together with no name here means the lock is the caller's own.
	LockedBy NullableString `json:"lockedBy,omitempty"`
	// For a fillable PDF form, whether the caller already has a filling draft of it, in which case `draftLocation`  says where that draft lives. Null for anything that is not a form.
	HasDraft NullableBool `json:"hasDraft,omitempty"`
	// How far the filling of this form has got for the calling account, and whose turn it is now. It is worked out  only inside a virtual data room, where filling runs in steps; everywhere else it stays at the none value.
	FormFillingStatus *FormFillingStatus `json:"formFillingStatus,omitempty"`
	// Whether the file is a PDF, and so offered as a fillable form. It is null for any other file type.
	IsForm NullableBool `json:"isForm,omitempty"`
	// True while a spreadsheet is in the mode where each person sorts and filters their own view without changing  what the others see, and null rather than false when it is not.
	CustomFilterEnabled NullableBool `json:"customFilterEnabled,omitempty"`
	// The display name of the account that turned that mode on, and null when the caller turned it on themselves.
	CustomFilterEnabledBy NullableString `json:"customFilterEnabledBy,omitempty"`
	// For a form in a room for filling, whether it has been released for filling; until then it is still being  prepared and only the people running the room work with it. Null for a file this does not apply to.
	StartFilling NullableBool `json:"startFilling,omitempty"`
	// True during the short window in which a released form is still being written out by the editor. Neither  filling nor editing is accepted while it lasts, so a client should wait and read the file again.
	IsFillingPreparing NullableBool `json:"isFillingPreparing,omitempty"`
	// Left empty by the portal: the folder holding the caller's draft is reported in `draftLocation` instead.
	InProcessFolderId NullableInt32 `json:"inProcessFolderId,omitempty"`
	// Left empty by the portal, like the identifier beside it; the draft's folder is named in `draftLocation`.
	InProcessFolderTitle NullableString `json:"inProcessFolderTitle,omitempty"`
	// The folder that collects the completed copies of this form. It is filled in only for the original form of a  room for filling, and only for a caller allowed to work with that form; null everywhere else.
	ResultsFolderId NullableInt32 `json:"resultsFolderId,omitempty"`
	// Where the caller's own filling draft of this form is kept. Null when there is no draft yet, which is the same  thing `hasDraft` reports.
	DraftLocation *ThirdPartyDraftLocation `json:"draftLocation,omitempty"`
	ViewAccessibility NullableFileDtoAllOfViewAccessibility `json:"viewAccessibility,omitempty"`
	// The moment the caller last opened the file. It is kept per account and is what orders the Recent section, so  it is null for a file this account has never opened. Written with the offset of the portal's time zone.
	LastOpened *ApiDateTime `json:"lastOpened,omitempty"`
	// The moment the file falls under the lifetime rule of the room holding it and is removed. It is counted from  the first revision rather than the latest one, so editing a file does not postpone it, and it is null when the  room sets no lifetime. Written with the offset of the portal's time zone.
	Expired *ApiDateTime `json:"expired,omitempty"`
	// How far the indexing of the file's content for AI search has got. It is null for a file that has never been  queued for indexing, which is every file while the feature is off for the portal.
	VectorizationStatus *VectorizationStatus `json:"vectorizationStatus,omitempty"`
	// The table collecting the submitted values of this form in the external database configured for its room. The  field is left out of the answer entirely when the form has no such table.
	ExternalDbTableName NullableString `json:"externalDbTableName,omitempty"`
	// The pixel size of the picture, measured by reading the stored file rather than taken from any stored metadata.  Null for anything that is not a picture the portal can show, and also when the file could not be read.
	Dimensions *ImageSizeDto `json:"dimensions,omitempty"`
}

// NewThirdPartyFileDto instantiates a new ThirdPartyFileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyFileDto() *ThirdPartyFileDto {
	this := ThirdPartyFileDto{}
	return &this
}

// NewThirdPartyFileDtoWithDefaults instantiates a new ThirdPartyFileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyFileDtoWithDefaults() *ThirdPartyFileDto {
	this := ThirdPartyFileDto{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *ThirdPartyFileDto) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ThirdPartyFileDto) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ThirdPartyFileDto) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetVersionGroup() int32 {
	if o == nil || IsNil(o.VersionGroup) {
		var ret int32
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetVersionGroupOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsVersionGroupSet() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given int32 and assigns it to the VersionGroup field.
func (o *ThirdPartyFileDto) SetVersionGroup(v int32) {
	o.VersionGroup = &v
}

// GetContentLength returns the ContentLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetContentLength() string {
	if o == nil || IsNil(o.ContentLength.Get()) {
		var ret string
		return ret
	}
	return *o.ContentLength.Get()
}

// GetContentLengthOk returns a tuple with the ContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetContentLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentLength.Get(), o.ContentLength.IsSet()
}

// HasContentLength returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsContentLengthSet() bool {
	if o != nil && o.ContentLength.IsSet() {
		return true
	}

	return false
}

// SetContentLength gets a reference to the given NullableString and assigns it to the ContentLength field.
func (o *ThirdPartyFileDto) SetContentLength(v string) {
	o.ContentLength.Set(&v)
}
// SetContentLengthNil sets the value for ContentLength to be an explicit nil
func (o *ThirdPartyFileDto) SetContentLengthNil() {
	o.ContentLength.Set(nil)
}

// UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetContentLength() {
	o.ContentLength.Unset()
}

// GetPureContentLength returns the PureContentLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetPureContentLength() int64 {
	if o == nil || IsNil(o.PureContentLength.Get()) {
		var ret int64
		return ret
	}
	return *o.PureContentLength.Get()
}

// GetPureContentLengthOk returns a tuple with the PureContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetPureContentLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PureContentLength.Get(), o.PureContentLength.IsSet()
}

// HasPureContentLength returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsPureContentLengthSet() bool {
	if o != nil && o.PureContentLength.IsSet() {
		return true
	}

	return false
}

// SetPureContentLength gets a reference to the given NullableInt64 and assigns it to the PureContentLength field.
func (o *ThirdPartyFileDto) SetPureContentLength(v int64) {
	o.PureContentLength.Set(&v)
}
// SetPureContentLengthNil sets the value for PureContentLength to be an explicit nil
func (o *ThirdPartyFileDto) SetPureContentLengthNil() {
	o.PureContentLength.Set(nil)
}

// UnsetPureContentLength ensures that no value is present for PureContentLength, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetPureContentLength() {
	o.PureContentLength.Unset()
}

// GetFileStatus returns the FileStatus field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetFileStatus() FileStatus {
	if o == nil || IsNil(o.FileStatus) {
		var ret FileStatus
		return ret
	}
	return *o.FileStatus
}

// GetFileStatusOk returns a tuple with the FileStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetFileStatusOk() (*FileStatus, bool) {
	if o == nil || IsNil(o.FileStatus) {
		return nil, false
	}
	return o.FileStatus, true
}

// HasFileStatus returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFileStatusSet() bool {
	if o != nil && !IsNil(o.FileStatus) {
		return true
	}

	return false
}

// SetFileStatus gets a reference to the given FileStatus and assigns it to the FileStatus field.
func (o *ThirdPartyFileDto) SetFileStatus(v FileStatus) {
	o.FileStatus = &v
}

// GetEditingBy returns the EditingBy field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetEditingBy() map[string]*string {
	if o == nil || IsNil(o.EditingBy) {
		var ret map[string]*string
		return ret
	}
	return o.EditingBy
}

// GetEditingByOk returns a tuple with the EditingBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetEditingByOk() (map[string]*string, bool) {
	if o == nil || IsNil(o.EditingBy) {
		return map[string]*string{}, false
	}
	return o.EditingBy, true
}

// HasEditingBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsEditingBySet() bool {
	if o != nil && !IsNil(o.EditingBy) {
		return true
	}

	return false
}

// SetEditingBy gets a reference to the given map[string]*string and assigns it to the EditingBy field.
func (o *ThirdPartyFileDto) SetEditingBy(v map[string]*string) {
	o.EditingBy = v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *ThirdPartyFileDto) SetMute(v bool) {
	o.Mute = &v
}

// GetViewUrl returns the ViewUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetViewUrl() string {
	if o == nil || IsNil(o.ViewUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ViewUrl.Get()
}

// GetViewUrlOk returns a tuple with the ViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetViewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewUrl.Get(), o.ViewUrl.IsSet()
}

// HasViewUrl returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsViewUrlSet() bool {
	if o != nil && o.ViewUrl.IsSet() {
		return true
	}

	return false
}

// SetViewUrl gets a reference to the given NullableString and assigns it to the ViewUrl field.
func (o *ThirdPartyFileDto) SetViewUrl(v string) {
	o.ViewUrl.Set(&v)
}
// SetViewUrlNil sets the value for ViewUrl to be an explicit nil
func (o *ThirdPartyFileDto) SetViewUrlNil() {
	o.ViewUrl.Set(nil)
}

// UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetViewUrl() {
	o.ViewUrl.Unset()
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsWebUrlSet() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *ThirdPartyFileDto) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *ThirdPartyFileDto) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetFileType() FileType {
	if o == nil || IsNil(o.FileType) {
		var ret FileType
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetFileTypeOk() (*FileType, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFileTypeSet() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given FileType and assigns it to the FileType field.
func (o *ThirdPartyFileDto) SetFileType(v FileType) {
	o.FileType = &v
}

// GetFileExst returns the FileExst field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetFileExst() string {
	if o == nil || IsNil(o.FileExst.Get()) {
		var ret string
		return ret
	}
	return *o.FileExst.Get()
}

// GetFileExstOk returns a tuple with the FileExst field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetFileExstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileExst.Get(), o.FileExst.IsSet()
}

// HasFileExst returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFileExstSet() bool {
	if o != nil && o.FileExst.IsSet() {
		return true
	}

	return false
}

// SetFileExst gets a reference to the given NullableString and assigns it to the FileExst field.
func (o *ThirdPartyFileDto) SetFileExst(v string) {
	o.FileExst.Set(&v)
}
// SetFileExstNil sets the value for FileExst to be an explicit nil
func (o *ThirdPartyFileDto) SetFileExstNil() {
	o.FileExst.Set(nil)
}

// UnsetFileExst ensures that no value is present for FileExst, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetFileExst() {
	o.FileExst.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCommentSet() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ThirdPartyFileDto) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ThirdPartyFileDto) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetComment() {
	o.Comment.Unset()
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.Encrypted.Get()
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encrypted.Get(), o.Encrypted.IsSet()
}

// HasEncrypted returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsEncryptedSet() bool {
	if o != nil && o.Encrypted.IsSet() {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given NullableBool and assigns it to the Encrypted field.
func (o *ThirdPartyFileDto) SetEncrypted(v bool) {
	o.Encrypted.Set(&v)
}
// SetEncryptedNil sets the value for Encrypted to be an explicit nil
func (o *ThirdPartyFileDto) SetEncryptedNil() {
	o.Encrypted.Set(nil)
}

// UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetEncrypted() {
	o.Encrypted.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsThumbnailUrlSet() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *ThirdPartyFileDto) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *ThirdPartyFileDto) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetThumbnailStatus returns the ThumbnailStatus field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetThumbnailStatus() Thumbnail {
	if o == nil || IsNil(o.ThumbnailStatus) {
		var ret Thumbnail
		return ret
	}
	return *o.ThumbnailStatus
}

// GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetThumbnailStatusOk() (*Thumbnail, bool) {
	if o == nil || IsNil(o.ThumbnailStatus) {
		return nil, false
	}
	return o.ThumbnailStatus, true
}

// HasThumbnailStatus returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsThumbnailStatusSet() bool {
	if o != nil && !IsNil(o.ThumbnailStatus) {
		return true
	}

	return false
}

// SetThumbnailStatus gets a reference to the given Thumbnail and assigns it to the ThumbnailStatus field.
func (o *ThirdPartyFileDto) SetThumbnailStatus(v Thumbnail) {
	o.ThumbnailStatus = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetLocked() bool {
	if o == nil || IsNil(o.Locked.Get()) {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsLockedSet() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *ThirdPartyFileDto) SetLocked(v bool) {
	o.Locked.Set(&v)
}
// SetLockedNil sets the value for Locked to be an explicit nil
func (o *ThirdPartyFileDto) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetLocked() {
	o.Locked.Unset()
}

// GetLockedBy returns the LockedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetLockedBy() string {
	if o == nil || IsNil(o.LockedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LockedBy.Get()
}

// GetLockedByOk returns a tuple with the LockedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetLockedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedBy.Get(), o.LockedBy.IsSet()
}

// HasLockedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsLockedBySet() bool {
	if o != nil && o.LockedBy.IsSet() {
		return true
	}

	return false
}

// SetLockedBy gets a reference to the given NullableString and assigns it to the LockedBy field.
func (o *ThirdPartyFileDto) SetLockedBy(v string) {
	o.LockedBy.Set(&v)
}
// SetLockedByNil sets the value for LockedBy to be an explicit nil
func (o *ThirdPartyFileDto) SetLockedByNil() {
	o.LockedBy.Set(nil)
}

// UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetLockedBy() {
	o.LockedBy.Unset()
}

// GetHasDraft returns the HasDraft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetHasDraft() bool {
	if o == nil || IsNil(o.HasDraft.Get()) {
		var ret bool
		return ret
	}
	return *o.HasDraft.Get()
}

// GetHasDraftOk returns a tuple with the HasDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetHasDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasDraft.Get(), o.HasDraft.IsSet()
}

// HasHasDraft returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsHasDraftSet() bool {
	if o != nil && o.HasDraft.IsSet() {
		return true
	}

	return false
}

// SetHasDraft gets a reference to the given NullableBool and assigns it to the HasDraft field.
func (o *ThirdPartyFileDto) SetHasDraft(v bool) {
	o.HasDraft.Set(&v)
}
// SetHasDraftNil sets the value for HasDraft to be an explicit nil
func (o *ThirdPartyFileDto) SetHasDraftNil() {
	o.HasDraft.Set(nil)
}

// UnsetHasDraft ensures that no value is present for HasDraft, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetHasDraft() {
	o.HasDraft.Unset()
}

// GetFormFillingStatus returns the FormFillingStatus field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetFormFillingStatus() FormFillingStatus {
	if o == nil || IsNil(o.FormFillingStatus) {
		var ret FormFillingStatus
		return ret
	}
	return *o.FormFillingStatus
}

// GetFormFillingStatusOk returns a tuple with the FormFillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetFormFillingStatusOk() (*FormFillingStatus, bool) {
	if o == nil || IsNil(o.FormFillingStatus) {
		return nil, false
	}
	return o.FormFillingStatus, true
}

// HasFormFillingStatus returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFormFillingStatusSet() bool {
	if o != nil && !IsNil(o.FormFillingStatus) {
		return true
	}

	return false
}

// SetFormFillingStatus gets a reference to the given FormFillingStatus and assigns it to the FormFillingStatus field.
func (o *ThirdPartyFileDto) SetFormFillingStatus(v FormFillingStatus) {
	o.FormFillingStatus = &v
}

// GetIsForm returns the IsForm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetIsForm() bool {
	if o == nil || IsNil(o.IsForm.Get()) {
		var ret bool
		return ret
	}
	return *o.IsForm.Get()
}

// GetIsFormOk returns a tuple with the IsForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetIsFormOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsForm.Get(), o.IsForm.IsSet()
}

// HasIsForm returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsIsFormSet() bool {
	if o != nil && o.IsForm.IsSet() {
		return true
	}

	return false
}

// SetIsForm gets a reference to the given NullableBool and assigns it to the IsForm field.
func (o *ThirdPartyFileDto) SetIsForm(v bool) {
	o.IsForm.Set(&v)
}
// SetIsFormNil sets the value for IsForm to be an explicit nil
func (o *ThirdPartyFileDto) SetIsFormNil() {
	o.IsForm.Set(nil)
}

// UnsetIsForm ensures that no value is present for IsForm, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetIsForm() {
	o.IsForm.Unset()
}

// GetCustomFilterEnabled returns the CustomFilterEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetCustomFilterEnabled() bool {
	if o == nil || IsNil(o.CustomFilterEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.CustomFilterEnabled.Get()
}

// GetCustomFilterEnabledOk returns a tuple with the CustomFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetCustomFilterEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFilterEnabled.Get(), o.CustomFilterEnabled.IsSet()
}

// HasCustomFilterEnabled returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCustomFilterEnabledSet() bool {
	if o != nil && o.CustomFilterEnabled.IsSet() {
		return true
	}

	return false
}

// SetCustomFilterEnabled gets a reference to the given NullableBool and assigns it to the CustomFilterEnabled field.
func (o *ThirdPartyFileDto) SetCustomFilterEnabled(v bool) {
	o.CustomFilterEnabled.Set(&v)
}
// SetCustomFilterEnabledNil sets the value for CustomFilterEnabled to be an explicit nil
func (o *ThirdPartyFileDto) SetCustomFilterEnabledNil() {
	o.CustomFilterEnabled.Set(nil)
}

// UnsetCustomFilterEnabled ensures that no value is present for CustomFilterEnabled, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetCustomFilterEnabled() {
	o.CustomFilterEnabled.Unset()
}

// GetCustomFilterEnabledBy returns the CustomFilterEnabledBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetCustomFilterEnabledBy() string {
	if o == nil || IsNil(o.CustomFilterEnabledBy.Get()) {
		var ret string
		return ret
	}
	return *o.CustomFilterEnabledBy.Get()
}

// GetCustomFilterEnabledByOk returns a tuple with the CustomFilterEnabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetCustomFilterEnabledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFilterEnabledBy.Get(), o.CustomFilterEnabledBy.IsSet()
}

// HasCustomFilterEnabledBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCustomFilterEnabledBySet() bool {
	if o != nil && o.CustomFilterEnabledBy.IsSet() {
		return true
	}

	return false
}

// SetCustomFilterEnabledBy gets a reference to the given NullableString and assigns it to the CustomFilterEnabledBy field.
func (o *ThirdPartyFileDto) SetCustomFilterEnabledBy(v string) {
	o.CustomFilterEnabledBy.Set(&v)
}
// SetCustomFilterEnabledByNil sets the value for CustomFilterEnabledBy to be an explicit nil
func (o *ThirdPartyFileDto) SetCustomFilterEnabledByNil() {
	o.CustomFilterEnabledBy.Set(nil)
}

// UnsetCustomFilterEnabledBy ensures that no value is present for CustomFilterEnabledBy, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetCustomFilterEnabledBy() {
	o.CustomFilterEnabledBy.Unset()
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling.Get()) {
		var ret bool
		return ret
	}
	return *o.StartFilling.Get()
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetStartFillingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartFilling.Get(), o.StartFilling.IsSet()
}

// HasStartFilling returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsStartFillingSet() bool {
	if o != nil && o.StartFilling.IsSet() {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given NullableBool and assigns it to the StartFilling field.
func (o *ThirdPartyFileDto) SetStartFilling(v bool) {
	o.StartFilling.Set(&v)
}
// SetStartFillingNil sets the value for StartFilling to be an explicit nil
func (o *ThirdPartyFileDto) SetStartFillingNil() {
	o.StartFilling.Set(nil)
}

// UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetStartFilling() {
	o.StartFilling.Unset()
}

// GetIsFillingPreparing returns the IsFillingPreparing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetIsFillingPreparing() bool {
	if o == nil || IsNil(o.IsFillingPreparing.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFillingPreparing.Get()
}

// GetIsFillingPreparingOk returns a tuple with the IsFillingPreparing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetIsFillingPreparingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFillingPreparing.Get(), o.IsFillingPreparing.IsSet()
}

// HasIsFillingPreparing returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsIsFillingPreparingSet() bool {
	if o != nil && o.IsFillingPreparing.IsSet() {
		return true
	}

	return false
}

// SetIsFillingPreparing gets a reference to the given NullableBool and assigns it to the IsFillingPreparing field.
func (o *ThirdPartyFileDto) SetIsFillingPreparing(v bool) {
	o.IsFillingPreparing.Set(&v)
}
// SetIsFillingPreparingNil sets the value for IsFillingPreparing to be an explicit nil
func (o *ThirdPartyFileDto) SetIsFillingPreparingNil() {
	o.IsFillingPreparing.Set(nil)
}

// UnsetIsFillingPreparing ensures that no value is present for IsFillingPreparing, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetIsFillingPreparing() {
	o.IsFillingPreparing.Unset()
}

// GetInProcessFolderId returns the InProcessFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetInProcessFolderId() int32 {
	if o == nil || IsNil(o.InProcessFolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.InProcessFolderId.Get()
}

// GetInProcessFolderIdOk returns a tuple with the InProcessFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetInProcessFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InProcessFolderId.Get(), o.InProcessFolderId.IsSet()
}

// HasInProcessFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsInProcessFolderIdSet() bool {
	if o != nil && o.InProcessFolderId.IsSet() {
		return true
	}

	return false
}

// SetInProcessFolderId gets a reference to the given NullableInt32 and assigns it to the InProcessFolderId field.
func (o *ThirdPartyFileDto) SetInProcessFolderId(v int32) {
	o.InProcessFolderId.Set(&v)
}
// SetInProcessFolderIdNil sets the value for InProcessFolderId to be an explicit nil
func (o *ThirdPartyFileDto) SetInProcessFolderIdNil() {
	o.InProcessFolderId.Set(nil)
}

// UnsetInProcessFolderId ensures that no value is present for InProcessFolderId, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetInProcessFolderId() {
	o.InProcessFolderId.Unset()
}

// GetInProcessFolderTitle returns the InProcessFolderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetInProcessFolderTitle() string {
	if o == nil || IsNil(o.InProcessFolderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.InProcessFolderTitle.Get()
}

// GetInProcessFolderTitleOk returns a tuple with the InProcessFolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetInProcessFolderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InProcessFolderTitle.Get(), o.InProcessFolderTitle.IsSet()
}

// HasInProcessFolderTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsInProcessFolderTitleSet() bool {
	if o != nil && o.InProcessFolderTitle.IsSet() {
		return true
	}

	return false
}

// SetInProcessFolderTitle gets a reference to the given NullableString and assigns it to the InProcessFolderTitle field.
func (o *ThirdPartyFileDto) SetInProcessFolderTitle(v string) {
	o.InProcessFolderTitle.Set(&v)
}
// SetInProcessFolderTitleNil sets the value for InProcessFolderTitle to be an explicit nil
func (o *ThirdPartyFileDto) SetInProcessFolderTitleNil() {
	o.InProcessFolderTitle.Set(nil)
}

// UnsetInProcessFolderTitle ensures that no value is present for InProcessFolderTitle, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetInProcessFolderTitle() {
	o.InProcessFolderTitle.Unset()
}

// GetResultsFolderId returns the ResultsFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetResultsFolderId() int32 {
	if o == nil || IsNil(o.ResultsFolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ResultsFolderId.Get()
}

// GetResultsFolderIdOk returns a tuple with the ResultsFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetResultsFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsFolderId.Get(), o.ResultsFolderId.IsSet()
}

// HasResultsFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsResultsFolderIdSet() bool {
	if o != nil && o.ResultsFolderId.IsSet() {
		return true
	}

	return false
}

// SetResultsFolderId gets a reference to the given NullableInt32 and assigns it to the ResultsFolderId field.
func (o *ThirdPartyFileDto) SetResultsFolderId(v int32) {
	o.ResultsFolderId.Set(&v)
}
// SetResultsFolderIdNil sets the value for ResultsFolderId to be an explicit nil
func (o *ThirdPartyFileDto) SetResultsFolderIdNil() {
	o.ResultsFolderId.Set(nil)
}

// UnsetResultsFolderId ensures that no value is present for ResultsFolderId, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetResultsFolderId() {
	o.ResultsFolderId.Unset()
}

// GetDraftLocation returns the DraftLocation field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetDraftLocation() ThirdPartyDraftLocation {
	if o == nil || IsNil(o.DraftLocation) {
		var ret ThirdPartyDraftLocation
		return ret
	}
	return *o.DraftLocation
}

// GetDraftLocationOk returns a tuple with the DraftLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetDraftLocationOk() (*ThirdPartyDraftLocation, bool) {
	if o == nil || IsNil(o.DraftLocation) {
		return nil, false
	}
	return o.DraftLocation, true
}

// HasDraftLocation returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsDraftLocationSet() bool {
	if o != nil && !IsNil(o.DraftLocation) {
		return true
	}

	return false
}

// SetDraftLocation gets a reference to the given ThirdPartyDraftLocation and assigns it to the DraftLocation field.
func (o *ThirdPartyFileDto) SetDraftLocation(v ThirdPartyDraftLocation) {
	o.DraftLocation = &v
}

// GetViewAccessibility returns the ViewAccessibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetViewAccessibility() FileDtoAllOfViewAccessibility {
	if o == nil || IsNil(o.ViewAccessibility.Get()) {
		var ret FileDtoAllOfViewAccessibility
		return ret
	}
	return *o.ViewAccessibility.Get()
}

// GetViewAccessibilityOk returns a tuple with the ViewAccessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetViewAccessibilityOk() (*FileDtoAllOfViewAccessibility, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewAccessibility.Get(), o.ViewAccessibility.IsSet()
}

// HasViewAccessibility returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsViewAccessibilitySet() bool {
	if o != nil && o.ViewAccessibility.IsSet() {
		return true
	}

	return false
}

// SetViewAccessibility gets a reference to the given NullableFileDtoAllOfViewAccessibility and assigns it to the ViewAccessibility field.
func (o *ThirdPartyFileDto) SetViewAccessibility(v FileDtoAllOfViewAccessibility) {
	o.ViewAccessibility.Set(&v)
}
// SetViewAccessibilityNil sets the value for ViewAccessibility to be an explicit nil
func (o *ThirdPartyFileDto) SetViewAccessibilityNil() {
	o.ViewAccessibility.Set(nil)
}

// UnsetViewAccessibility ensures that no value is present for ViewAccessibility, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetViewAccessibility() {
	o.ViewAccessibility.Unset()
}

// GetLastOpened returns the LastOpened field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetLastOpened() ApiDateTime {
	if o == nil || IsNil(o.LastOpened) {
		var ret ApiDateTime
		return ret
	}
	return *o.LastOpened
}

// GetLastOpenedOk returns a tuple with the LastOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetLastOpenedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.LastOpened) {
		return nil, false
	}
	return o.LastOpened, true
}

// HasLastOpened returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsLastOpenedSet() bool {
	if o != nil && !IsNil(o.LastOpened) {
		return true
	}

	return false
}

// SetLastOpened gets a reference to the given ApiDateTime and assigns it to the LastOpened field.
func (o *ThirdPartyFileDto) SetLastOpened(v ApiDateTime) {
	o.LastOpened = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetExpired() ApiDateTime {
	if o == nil || IsNil(o.Expired) {
		var ret ApiDateTime
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetExpiredOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsExpiredSet() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given ApiDateTime and assigns it to the Expired field.
func (o *ThirdPartyFileDto) SetExpired(v ApiDateTime) {
	o.Expired = &v
}

// GetVectorizationStatus returns the VectorizationStatus field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetVectorizationStatus() VectorizationStatus {
	if o == nil || IsNil(o.VectorizationStatus) {
		var ret VectorizationStatus
		return ret
	}
	return *o.VectorizationStatus
}

// GetVectorizationStatusOk returns a tuple with the VectorizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetVectorizationStatusOk() (*VectorizationStatus, bool) {
	if o == nil || IsNil(o.VectorizationStatus) {
		return nil, false
	}
	return o.VectorizationStatus, true
}

// HasVectorizationStatus returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsVectorizationStatusSet() bool {
	if o != nil && !IsNil(o.VectorizationStatus) {
		return true
	}

	return false
}

// SetVectorizationStatus gets a reference to the given VectorizationStatus and assigns it to the VectorizationStatus field.
func (o *ThirdPartyFileDto) SetVectorizationStatus(v VectorizationStatus) {
	o.VectorizationStatus = &v
}

// GetExternalDbTableName returns the ExternalDbTableName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetExternalDbTableName() string {
	if o == nil || IsNil(o.ExternalDbTableName.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalDbTableName.Get()
}

// GetExternalDbTableNameOk returns a tuple with the ExternalDbTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetExternalDbTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalDbTableName.Get(), o.ExternalDbTableName.IsSet()
}

// HasExternalDbTableName returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsExternalDbTableNameSet() bool {
	if o != nil && o.ExternalDbTableName.IsSet() {
		return true
	}

	return false
}

// SetExternalDbTableName gets a reference to the given NullableString and assigns it to the ExternalDbTableName field.
func (o *ThirdPartyFileDto) SetExternalDbTableName(v string) {
	o.ExternalDbTableName.Set(&v)
}
// SetExternalDbTableNameNil sets the value for ExternalDbTableName to be an explicit nil
func (o *ThirdPartyFileDto) SetExternalDbTableNameNil() {
	o.ExternalDbTableName.Set(nil)
}

// UnsetExternalDbTableName ensures that no value is present for ExternalDbTableName, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetExternalDbTableName() {
	o.ExternalDbTableName.Unset()
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetDimensions() ImageSizeDto {
	if o == nil || IsNil(o.Dimensions) {
		var ret ImageSizeDto
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetDimensionsOk() (*ImageSizeDto, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsDimensionsSet() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given ImageSizeDto and assigns it to the Dimensions field.
func (o *ThirdPartyFileDto) SetDimensions(v ImageSizeDto) {
	o.Dimensions = &v
}

func (o ThirdPartyFileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyFileDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedThirdPartyFileEntryDto, errThirdPartyFileEntryDto := json.Marshal(o.ThirdPartyFileEntryDto)
	if errThirdPartyFileEntryDto != nil {
		return map[string]interface{}{}, errThirdPartyFileEntryDto
	}
	errThirdPartyFileEntryDto = json.Unmarshal([]byte(serializedThirdPartyFileEntryDto), &toSerialize)
	if errThirdPartyFileEntryDto != nil {
		return map[string]interface{}{}, errThirdPartyFileEntryDto
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionGroup) {
		toSerialize["versionGroup"] = o.VersionGroup
	}
	if o.ContentLength.IsSet() {
		toSerialize["contentLength"] = o.ContentLength.Get()
	}
	if o.PureContentLength.IsSet() {
		toSerialize["pureContentLength"] = o.PureContentLength.Get()
	}
	if !IsNil(o.FileStatus) {
		toSerialize["fileStatus"] = o.FileStatus
	}
	if !IsNil(o.EditingBy) {
		toSerialize["editingBy"] = o.EditingBy
	}
	if !IsNil(o.Mute) {
		toSerialize["mute"] = o.Mute
	}
	if o.ViewUrl.IsSet() {
		toSerialize["viewUrl"] = o.ViewUrl.Get()
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if !IsNil(o.FileType) {
		toSerialize["fileType"] = o.FileType
	}
	if o.FileExst.IsSet() {
		toSerialize["fileExst"] = o.FileExst.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Encrypted.IsSet() {
		toSerialize["encrypted"] = o.Encrypted.Get()
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnailUrl"] = o.ThumbnailUrl.Get()
	}
	if !IsNil(o.ThumbnailStatus) {
		toSerialize["thumbnailStatus"] = o.ThumbnailStatus
	}
	if o.Locked.IsSet() {
		toSerialize["locked"] = o.Locked.Get()
	}
	if o.LockedBy.IsSet() {
		toSerialize["lockedBy"] = o.LockedBy.Get()
	}
	if o.HasDraft.IsSet() {
		toSerialize["hasDraft"] = o.HasDraft.Get()
	}
	if !IsNil(o.FormFillingStatus) {
		toSerialize["formFillingStatus"] = o.FormFillingStatus
	}
	if o.IsForm.IsSet() {
		toSerialize["isForm"] = o.IsForm.Get()
	}
	if o.CustomFilterEnabled.IsSet() {
		toSerialize["customFilterEnabled"] = o.CustomFilterEnabled.Get()
	}
	if o.CustomFilterEnabledBy.IsSet() {
		toSerialize["customFilterEnabledBy"] = o.CustomFilterEnabledBy.Get()
	}
	if o.StartFilling.IsSet() {
		toSerialize["startFilling"] = o.StartFilling.Get()
	}
	if o.IsFillingPreparing.IsSet() {
		toSerialize["isFillingPreparing"] = o.IsFillingPreparing.Get()
	}
	if o.InProcessFolderId.IsSet() {
		toSerialize["inProcessFolderId"] = o.InProcessFolderId.Get()
	}
	if o.InProcessFolderTitle.IsSet() {
		toSerialize["inProcessFolderTitle"] = o.InProcessFolderTitle.Get()
	}
	if o.ResultsFolderId.IsSet() {
		toSerialize["resultsFolderId"] = o.ResultsFolderId.Get()
	}
	if !IsNil(o.DraftLocation) {
		toSerialize["draftLocation"] = o.DraftLocation
	}
	if o.ViewAccessibility.IsSet() {
		toSerialize["viewAccessibility"] = o.ViewAccessibility.Get()
	}
	if !IsNil(o.LastOpened) {
		toSerialize["lastOpened"] = o.LastOpened
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !IsNil(o.VectorizationStatus) {
		toSerialize["vectorizationStatus"] = o.VectorizationStatus
	}
	if o.ExternalDbTableName.IsSet() {
		toSerialize["externalDbTableName"] = o.ExternalDbTableName.Get()
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	return toSerialize, nil
}

type NullableThirdPartyFileDto struct {
	value *ThirdPartyFileDto
	isSet bool
}

func (v NullableThirdPartyFileDto) Get() *ThirdPartyFileDto {
	return v.value
}

func (v *NullableThirdPartyFileDto) Set(val *ThirdPartyFileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyFileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyFileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyFileDto(val *ThirdPartyFileDto) *NullableThirdPartyFileDto {
	return &NullableThirdPartyFileDto{value: val, isSet: true}
}

func (v NullableThirdPartyFileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyFileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

