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

// checks if the ThirdPartyFolderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyFolderDto{}

// ThirdPartyFolderDto The folder, with the fields that only a room carries filled in when the folder is a room.
type ThirdPartyFolderDto struct {
	ThirdPartyFileEntryDto
	// The folder this one is listed in. For a room it is the root of the section the room lives in, and for an entry  opened through a sharing link whose real parent the caller may not read it is the root of the section with the  entries shared with them.
	ParentId NullableString `json:"parentId,omitempty"`
	// How many files lie directly in the folder, without counting the subfolders. The roots of the `Rooms`, room  templates, default templates and AI agents sections always report 0, because the number is not collected for  them.
	FilesCount *int32 `json:"filesCount,omitempty"`
	// How many subfolders lie directly in the folder. For an AI room the two service subfolders it always holds are  subtracted, so the number matches what a listing of it shows, and the roots of the `Rooms`, templates and AI  agents sections report 0.
	FoldersCount *int32 `json:"foldersCount,omitempty"`
	// Whether the caller may hand out access to the folder. It is filled in only for the folder a folder-contents  answer is about, and is null in every other answer, so null says nothing about the sharing rights.
	IsShareable NullableBool `json:"isShareable,omitempty"`
	// How many entries inside the folder the caller has not opened yet, the number drawn as the badge on it. An  account that turned the badges off in its own settings always reads 0 here, so 0 alone does not prove that  everything has been seen.
	New *int32 `json:"new,omitempty"`
	// Whether the caller silenced the notifications of this room: true means no message about its activity reaches  them. The choice belongs to the reading account rather than to the room, so two members of one room read  different values.
	Mute *bool `json:"mute,omitempty"`
	// The names of the tags attached to the room. Empty for a folder that is not a room, since only rooms carry  tags, and the names are the ones from the portal tag catalogue.
	Tags []string `json:"tags,omitempty"`
	// The addresses of the room logo in four sizes, together with the colour and the built-in cover that are drawn  when no logo was uploaded. A room without a logo answers with four empty addresses rather than with null, and  the field is null for a folder that is not a room.
	Logo *LogoDto `json:"logo,omitempty"`
	// Whether the caller pinned the room to the top of their own room list. Pinning is personal and is lost when the  room is archived.
	Pinned *bool `json:"pinned,omitempty"`
	// The kind of the room, which decides the default access rules of its members. Null for a folder that is not a  room.
	RoomType *RoomType `json:"roomType,omitempty"`
	// Whether the room is a private one, which limits it to the accounts invited into it and needs encryption keys  set up for each of them.
	Private *bool `json:"private,omitempty"`
	// Whether the contents of the room are kept in an explicit numbered order, the one reported as `order` on each  entry, instead of being left to the sorting the reader asks for.
	Indexing *bool `json:"indexing,omitempty"`
	// Whether downloading and printing the contents of the room is forbidden, which leaves its members with viewing  and editing in the editor.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The rule by which the files of the room are removed once they grow old. Null when the room has no such rule,  which is also what is reported after the rule is switched off, because switching it off erases it.
	Lifetime *RoomDataLifetimeDto `json:"lifetime,omitempty"`
	// The watermark stamped over the documents of the room while they are viewed and printed. Null when the room has  no watermark, and for every folder that is not a room.
	Watermark *WatermarkDto `json:"watermark,omitempty"`
	// The part the folder plays inside its room: one of the service folders of the form-filling flow, or the  knowledge and result storages of an AI room. It stays null for an ordinary folder and for the room itself, so  it does not describe folders in general.
	Type *FolderType `json:"type,omitempty"`
	// Whether the caller holds the room through an invitation of their own: true for the account that created it and  for a member invited personally, false when the access comes from a group they belong to, and null for a  folder that is not a room.
	InRoom NullableBool `json:"inRoom,omitempty"`
	// How much space the files of the room may take, in bytes. It is the limit set on this room, or the portal  default for rooms when none was set. Null when the tariff of the portal does not count room statistics, when  room quotas are switched off, when the room lies in the archive or the trash, or when the role of the caller  in the room carries no editing rights - a viewer, a commenter, a reviewer and a form filler do not see the  figure, a room manager, a content creator and an editor do.
	QuotaLimit NullableInt64 `json:"quotaLimit,omitempty"`
	// Whether `quotaLimit` is a limit set on this room (true) or the portal default for rooms (false). Null exactly  when `quotaLimit` is null.
	IsCustomQuota NullableBool `json:"isCustomQuota,omitempty"`
	// How much the files of the room take, in bytes, as of the last time the counter was recomputed. The counter is  refreshed when a file operation finishes, so a read right after an upload or a deletion can still report the  previous figure. Null for a folder that is not a room.
	UsedSpace NullableInt64 `json:"usedSpace,omitempty"`
	// Whether the sharing link the folder was opened through asks for a password that has not been entered yet.  While it is true the contents stay unreadable; send the password to `POST api/2.0/files/share/{key}/password`  first. Null when the folder was not reached through a link.
	PasswordProtected NullableBool `json:"passwordProtected,omitempty"`
	// Deprecated, read `isLinkExpired` instead: whether the sharing link the folder was opened through has run out  of its lifetime.
	// Deprecated
	Expired NullableBool `json:"expired,omitempty"`
	// The chat configuration of an AI room. Only the system prompt is reported here, whatever else the room stores,  and the field is null for every folder that is not an AI room.
	ChatSettings *ChatSettingsDto `json:"chatSettings,omitempty"`
	// The kind of the room the folder lies in. It is filled in only for the folder a folder-contents answer is  about, and only when that room is an AI room, so it is null in every other answer and for every other room  kind.
	RootRoomType *RoomType `json:"rootRoomType,omitempty"`
	// Whether the answers collected in this form-filling room are also gathered into a spreadsheet next to the  completed copies. Filled in for form-filling rooms only.
	SaveFormAsXLSX NullableBool `json:"saveFormAsXLSX,omitempty"`
	// Whether the answers collected in this form-filling room are also pushed into the external database configured  for the portal. Filled in for form-filling rooms only.
	SendFormToExternalDB NullableBool `json:"sendFormToExternalDB,omitempty"`
	// The form the completed copies in this folder were filled from, taken from the copy submitted last. Null while  the folder holds no completed copy, and for every folder that does not collect them.
	OriginalFormId NullableInt32 `json:"originalFormId,omitempty"`
}

// NewThirdPartyFolderDto instantiates a new ThirdPartyFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyFolderDto() *ThirdPartyFolderDto {
	this := ThirdPartyFolderDto{}
	return &this
}

// NewThirdPartyFolderDtoWithDefaults instantiates a new ThirdPartyFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyFolderDtoWithDefaults() *ThirdPartyFolderDto {
	this := ThirdPartyFolderDto{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsParentIdSet() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ThirdPartyFolderDto) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ThirdPartyFolderDto) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetParentId() {
	o.ParentId.Unset()
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsFilesCountSet() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *ThirdPartyFolderDto) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetFoldersCount returns the FoldersCount field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetFoldersCount() int32 {
	if o == nil || IsNil(o.FoldersCount) {
		var ret int32
		return ret
	}
	return *o.FoldersCount
}

// GetFoldersCountOk returns a tuple with the FoldersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetFoldersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FoldersCount) {
		return nil, false
	}
	return o.FoldersCount, true
}

// HasFoldersCount returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsFoldersCountSet() bool {
	if o != nil && !IsNil(o.FoldersCount) {
		return true
	}

	return false
}

// SetFoldersCount gets a reference to the given int32 and assigns it to the FoldersCount field.
func (o *ThirdPartyFolderDto) SetFoldersCount(v int32) {
	o.FoldersCount = &v
}

// GetIsShareable returns the IsShareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetIsShareable() bool {
	if o == nil || IsNil(o.IsShareable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsShareable.Get()
}

// GetIsShareableOk returns a tuple with the IsShareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetIsShareableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsShareable.Get(), o.IsShareable.IsSet()
}

// HasIsShareable returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIsShareableSet() bool {
	if o != nil && o.IsShareable.IsSet() {
		return true
	}

	return false
}

// SetIsShareable gets a reference to the given NullableBool and assigns it to the IsShareable field.
func (o *ThirdPartyFolderDto) SetIsShareable(v bool) {
	o.IsShareable.Set(&v)
}
// SetIsShareableNil sets the value for IsShareable to be an explicit nil
func (o *ThirdPartyFolderDto) SetIsShareableNil() {
	o.IsShareable.Set(nil)
}

// UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetIsShareable() {
	o.IsShareable.Unset()
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *ThirdPartyFolderDto) SetNew(v int32) {
	o.New = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *ThirdPartyFolderDto) SetMute(v bool) {
	o.Mute = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ThirdPartyFolderDto) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetLogo() LogoDto {
	if o == nil || IsNil(o.Logo) {
		var ret LogoDto
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetLogoOk() (*LogoDto, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoDto and assigns it to the Logo field.
func (o *ThirdPartyFolderDto) SetLogo(v LogoDto) {
	o.Logo = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsPinnedSet() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *ThirdPartyFolderDto) SetPinned(v bool) {
	o.Pinned = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetRoomType() RoomType {
	if o == nil || IsNil(o.RoomType) {
		var ret RoomType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsRoomTypeSet() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given RoomType and assigns it to the RoomType field.
func (o *ThirdPartyFolderDto) SetRoomType(v RoomType) {
	o.RoomType = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ThirdPartyFolderDto) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *ThirdPartyFolderDto) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *ThirdPartyFolderDto) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *ThirdPartyFolderDto) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetWatermark() WatermarkDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetWatermarkOk() (*WatermarkDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkDto and assigns it to the Watermark field.
func (o *ThirdPartyFolderDto) SetWatermark(v WatermarkDto) {
	o.Watermark = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetType() FolderType {
	if o == nil || IsNil(o.Type) {
		var ret FolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FolderType and assigns it to the Type field.
func (o *ThirdPartyFolderDto) SetType(v FolderType) {
	o.Type = &v
}

// GetInRoom returns the InRoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetInRoom() bool {
	if o == nil || IsNil(o.InRoom.Get()) {
		var ret bool
		return ret
	}
	return *o.InRoom.Get()
}

// GetInRoomOk returns a tuple with the InRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetInRoomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InRoom.Get(), o.InRoom.IsSet()
}

// HasInRoom returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsInRoomSet() bool {
	if o != nil && o.InRoom.IsSet() {
		return true
	}

	return false
}

// SetInRoom gets a reference to the given NullableBool and assigns it to the InRoom field.
func (o *ThirdPartyFolderDto) SetInRoom(v bool) {
	o.InRoom.Set(&v)
}
// SetInRoomNil sets the value for InRoom to be an explicit nil
func (o *ThirdPartyFolderDto) SetInRoomNil() {
	o.InRoom.Set(nil)
}

// UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetInRoom() {
	o.InRoom.Unset()
}

// GetQuotaLimit returns the QuotaLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetQuotaLimit() int64 {
	if o == nil || IsNil(o.QuotaLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.QuotaLimit.Get()
}

// GetQuotaLimitOk returns a tuple with the QuotaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetQuotaLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaLimit.Get(), o.QuotaLimit.IsSet()
}

// HasQuotaLimit returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsQuotaLimitSet() bool {
	if o != nil && o.QuotaLimit.IsSet() {
		return true
	}

	return false
}

// SetQuotaLimit gets a reference to the given NullableInt64 and assigns it to the QuotaLimit field.
func (o *ThirdPartyFolderDto) SetQuotaLimit(v int64) {
	o.QuotaLimit.Set(&v)
}
// SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil
func (o *ThirdPartyFolderDto) SetQuotaLimitNil() {
	o.QuotaLimit.Set(nil)
}

// UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetQuotaLimit() {
	o.QuotaLimit.Unset()
}

// GetIsCustomQuota returns the IsCustomQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetIsCustomQuota() bool {
	if o == nil || IsNil(o.IsCustomQuota.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCustomQuota.Get()
}

// GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetIsCustomQuotaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCustomQuota.Get(), o.IsCustomQuota.IsSet()
}

// HasIsCustomQuota returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIsCustomQuotaSet() bool {
	if o != nil && o.IsCustomQuota.IsSet() {
		return true
	}

	return false
}

// SetIsCustomQuota gets a reference to the given NullableBool and assigns it to the IsCustomQuota field.
func (o *ThirdPartyFolderDto) SetIsCustomQuota(v bool) {
	o.IsCustomQuota.Set(&v)
}
// SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil
func (o *ThirdPartyFolderDto) SetIsCustomQuotaNil() {
	o.IsCustomQuota.Set(nil)
}

// UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetIsCustomQuota() {
	o.IsCustomQuota.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.UsedSpace.Get()
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetUsedSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedSpace.Get(), o.UsedSpace.IsSet()
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsUsedSpaceSet() bool {
	if o != nil && o.UsedSpace.IsSet() {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given NullableInt64 and assigns it to the UsedSpace field.
func (o *ThirdPartyFolderDto) SetUsedSpace(v int64) {
	o.UsedSpace.Set(&v)
}
// SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil
func (o *ThirdPartyFolderDto) SetUsedSpaceNil() {
	o.UsedSpace.Set(nil)
}

// UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetUsedSpace() {
	o.UsedSpace.Unset()
}

// GetPasswordProtected returns the PasswordProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetPasswordProtected() bool {
	if o == nil || IsNil(o.PasswordProtected.Get()) {
		var ret bool
		return ret
	}
	return *o.PasswordProtected.Get()
}

// GetPasswordProtectedOk returns a tuple with the PasswordProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetPasswordProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProtected.Get(), o.PasswordProtected.IsSet()
}

// HasPasswordProtected returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsPasswordProtectedSet() bool {
	if o != nil && o.PasswordProtected.IsSet() {
		return true
	}

	return false
}

// SetPasswordProtected gets a reference to the given NullableBool and assigns it to the PasswordProtected field.
func (o *ThirdPartyFolderDto) SetPasswordProtected(v bool) {
	o.PasswordProtected.Set(&v)
}
// SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil
func (o *ThirdPartyFolderDto) SetPasswordProtectedNil() {
	o.PasswordProtected.Set(nil)
}

// UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetPasswordProtected() {
	o.PasswordProtected.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ThirdPartyFolderDto) GetExpired() bool {
	if o == nil || IsNil(o.Expired.Get()) {
		var ret bool
		return ret
	}
	return *o.Expired.Get()
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ThirdPartyFolderDto) GetExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expired.Get(), o.Expired.IsSet()
}

// HasExpired returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsExpiredSet() bool {
	if o != nil && o.Expired.IsSet() {
		return true
	}

	return false
}

// SetExpired gets a reference to the given NullableBool and assigns it to the Expired field.
// Deprecated
func (o *ThirdPartyFolderDto) SetExpired(v bool) {
	o.Expired.Set(&v)
}
// SetExpiredNil sets the value for Expired to be an explicit nil
func (o *ThirdPartyFolderDto) SetExpiredNil() {
	o.Expired.Set(nil)
}

// UnsetExpired ensures that no value is present for Expired, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetExpired() {
	o.Expired.Unset()
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetChatSettings() ChatSettingsDto {
	if o == nil || IsNil(o.ChatSettings) {
		var ret ChatSettingsDto
		return ret
	}
	return *o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetChatSettingsOk() (*ChatSettingsDto, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return nil, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given ChatSettingsDto and assigns it to the ChatSettings field.
func (o *ThirdPartyFolderDto) SetChatSettings(v ChatSettingsDto) {
	o.ChatSettings = &v
}

// GetRootRoomType returns the RootRoomType field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetRootRoomType() RoomType {
	if o == nil || IsNil(o.RootRoomType) {
		var ret RoomType
		return ret
	}
	return *o.RootRoomType
}

// GetRootRoomTypeOk returns a tuple with the RootRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetRootRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RootRoomType) {
		return nil, false
	}
	return o.RootRoomType, true
}

// HasRootRoomType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsRootRoomTypeSet() bool {
	if o != nil && !IsNil(o.RootRoomType) {
		return true
	}

	return false
}

// SetRootRoomType gets a reference to the given RoomType and assigns it to the RootRoomType field.
func (o *ThirdPartyFolderDto) SetRootRoomType(v RoomType) {
	o.RootRoomType = &v
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX.Get()
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveFormAsXLSX.Get(), o.SaveFormAsXLSX.IsSet()
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSaveFormAsXLSXSet() bool {
	if o != nil && o.SaveFormAsXLSX.IsSet() {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given NullableBool and assigns it to the SaveFormAsXLSX field.
func (o *ThirdPartyFolderDto) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX.Set(&v)
}
// SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil
func (o *ThirdPartyFolderDto) SetSaveFormAsXLSXNil() {
	o.SaveFormAsXLSX.Set(nil)
}

// UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetSaveFormAsXLSX() {
	o.SaveFormAsXLSX.Unset()
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB.Get()) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB.Get()
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendFormToExternalDB.Get(), o.SendFormToExternalDB.IsSet()
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSendFormToExternalDBSet() bool {
	if o != nil && o.SendFormToExternalDB.IsSet() {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given NullableBool and assigns it to the SendFormToExternalDB field.
func (o *ThirdPartyFolderDto) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB.Set(&v)
}
// SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil
func (o *ThirdPartyFolderDto) SetSendFormToExternalDBNil() {
	o.SendFormToExternalDB.Set(nil)
}

// UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetSendFormToExternalDB() {
	o.SendFormToExternalDB.Unset()
}

// GetOriginalFormId returns the OriginalFormId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetOriginalFormId() int32 {
	if o == nil || IsNil(o.OriginalFormId.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalFormId.Get()
}

// GetOriginalFormIdOk returns a tuple with the OriginalFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetOriginalFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFormId.Get(), o.OriginalFormId.IsSet()
}

// HasOriginalFormId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOriginalFormIdSet() bool {
	if o != nil && o.OriginalFormId.IsSet() {
		return true
	}

	return false
}

// SetOriginalFormId gets a reference to the given NullableInt32 and assigns it to the OriginalFormId field.
func (o *ThirdPartyFolderDto) SetOriginalFormId(v int32) {
	o.OriginalFormId.Set(&v)
}
// SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil
func (o *ThirdPartyFolderDto) SetOriginalFormIdNil() {
	o.OriginalFormId.Set(nil)
}

// UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetOriginalFormId() {
	o.OriginalFormId.Unset()
}

func (o ThirdPartyFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyFolderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedThirdPartyFileEntryDto, errThirdPartyFileEntryDto := json.Marshal(o.ThirdPartyFileEntryDto)
	if errThirdPartyFileEntryDto != nil {
		return map[string]interface{}{}, errThirdPartyFileEntryDto
	}
	errThirdPartyFileEntryDto = json.Unmarshal([]byte(serializedThirdPartyFileEntryDto), &toSerialize)
	if errThirdPartyFileEntryDto != nil {
		return map[string]interface{}{}, errThirdPartyFileEntryDto
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if !IsNil(o.FilesCount) {
		toSerialize["filesCount"] = o.FilesCount
	}
	if !IsNil(o.FoldersCount) {
		toSerialize["foldersCount"] = o.FoldersCount
	}
	if o.IsShareable.IsSet() {
		toSerialize["isShareable"] = o.IsShareable.Get()
	}
	if !IsNil(o.New) {
		toSerialize["new"] = o.New
	}
	if !IsNil(o.Mute) {
		toSerialize["mute"] = o.Mute
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Pinned) {
		toSerialize["pinned"] = o.Pinned
	}
	if !IsNil(o.RoomType) {
		toSerialize["roomType"] = o.RoomType
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Indexing) {
		toSerialize["indexing"] = o.Indexing
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
	}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}
	if !IsNil(o.Watermark) {
		toSerialize["watermark"] = o.Watermark
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.InRoom.IsSet() {
		toSerialize["inRoom"] = o.InRoom.Get()
	}
	if o.QuotaLimit.IsSet() {
		toSerialize["quotaLimit"] = o.QuotaLimit.Get()
	}
	if o.IsCustomQuota.IsSet() {
		toSerialize["isCustomQuota"] = o.IsCustomQuota.Get()
	}
	if o.UsedSpace.IsSet() {
		toSerialize["usedSpace"] = o.UsedSpace.Get()
	}
	if o.PasswordProtected.IsSet() {
		toSerialize["passwordProtected"] = o.PasswordProtected.Get()
	}
	if o.Expired.IsSet() {
		toSerialize["expired"] = o.Expired.Get()
	}
	if !IsNil(o.ChatSettings) {
		toSerialize["chatSettings"] = o.ChatSettings
	}
	if !IsNil(o.RootRoomType) {
		toSerialize["rootRoomType"] = o.RootRoomType
	}
	if o.SaveFormAsXLSX.IsSet() {
		toSerialize["saveFormAsXLSX"] = o.SaveFormAsXLSX.Get()
	}
	if o.SendFormToExternalDB.IsSet() {
		toSerialize["sendFormToExternalDB"] = o.SendFormToExternalDB.Get()
	}
	if o.OriginalFormId.IsSet() {
		toSerialize["originalFormId"] = o.OriginalFormId.Get()
	}
	return toSerialize, nil
}

type NullableThirdPartyFolderDto struct {
	value *ThirdPartyFolderDto
	isSet bool
}

func (v NullableThirdPartyFolderDto) Get() *ThirdPartyFolderDto {
	return v.value
}

func (v *NullableThirdPartyFolderDto) Set(val *ThirdPartyFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyFolderDto(val *ThirdPartyFolderDto) *NullableThirdPartyFolderDto {
	return &NullableThirdPartyFolderDto{value: val, isSet: true}
}

func (v NullableThirdPartyFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

