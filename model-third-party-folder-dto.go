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
	// The name shown for the entry. For a file it carries the extension, which is how the format is recognised, and  for a room it is the room name.
	Title *string `json:"title,omitempty"`
	// The level the calling account holds on this entry, resolved from its own rights, the groups it belongs to and  any link it came in through. It is the level itself, not what the account may do with it - the action flags  below answer that.
	Access *FileShare `json:"access,omitempty"`
	// Who gave the calling account the access it is using. It is filled in only while the entry is being read  through a share, and never for a caller without an account.
	SharedBy *EmployeeDto `json:"sharedBy,omitempty"`
	// Who owns the place the entry is shared from - the creator of the room it lies in, or of the personal section  that holds it. It is filled in only while the entry is being read through a share, and never for a caller  without an account.
	OwnedBy *EmployeeDto `json:"ownedBy,omitempty"`
	// Whether at least one external link exists for the entry, whichever kind. It says nothing about accounts and  groups - those are counted by the flag for members below.
	Shared *bool `json:"shared,omitempty"`
	// Whether at least one account or group has been given rights on the entry directly, as opposed to reaching it  through a link or through the room around it.
	SharedForUser *bool `json:"sharedForUser,omitempty"`
	// Whether one of the entry's links is open to people outside the portal, as opposed to a link that only its own  members can follow. This is the flag to watch when the concern is who can reach the content from outside.
	SharedExternal *bool `json:"sharedExternal,omitempty"`
	// Whether the entry is reachable because the room or folder around it is shared, rather than through rights of  its own. A copy or a move takes the entry out of that scope.
	ParentShared *bool `json:"parentShared,omitempty"`
	// A shortened address that opens the entry through the link it is being read with. It is an empty string  whenever no link applies, which is the usual case for a member browsing their own rooms.
	ShortWebUrl *string `json:"shortWebUrl,omitempty"`
	// When the entry was created, written with the offset of the portal's time zone. For a file restored from an  older version this is still the moment the file first appeared.
	Created *ApiDateTime `json:"created,omitempty"`
	// Who created the entry. It is null for a caller without an account, who is told nothing about the portal's  members.
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
	// When the entry last changed, written with the offset of the portal's time zone. It is never reported as  earlier than the creation moment, so the two can be compared safely.
	Updated *ApiDateTime `json:"updated,omitempty"`
	// When the entry will disappear on its own, written with the offset of the portal's time zone. It is filled in  only where a removal is actually scheduled - something in the trash while the portal cleans it up  automatically, or a guest's own documents - so a null means nothing is scheduled rather than that the entry is  permanent.
	AutoDelete *ApiDateTime `json:"autoDelete,omitempty"`
	// The section the entry ultimately belongs to, which is what tells a personal document from one inside a room,  from a template and from something in the trash or the archive.
	RootFolderType *FolderType `json:"rootFolderType,omitempty"`
	// The kind of room the entry lies in, which decides what the room allows - filling forms, public links,  indexing. It is null for an entry that is not inside a room at all.
	ParentRoomType *FolderType `json:"parentRoomType,omitempty"`
	// Who changed the entry last. It is null for a caller without an account.
	UpdatedBy *EmployeeDto `json:"updatedBy,omitempty"`
	// Set when the entry is stored on a connected third-party account rather than on the portal, and null when it is  stored on the portal. Such an entry is identified by a string rather than a number, and some operations skip  it.
	ProviderItem *bool `json:"providerItem,omitempty"`
	// Which third-party service holds the entry, matching the keys accepted by the third-party operations. It is  null for an entry stored on the portal.
	ProviderKey *string `json:"providerKey,omitempty"`
	// The connected account the entry comes from, for telling apart two connections to the same service. It is null  for an entry stored on the portal.
	ProviderId *int32 `json:"providerId,omitempty"`
	// The place of the entry in a room where the members arrange the content themselves, given as the position of  the entry preceded by the positions of the folders leading to it, separated by dots. It is empty when nothing  has been arranged.
	Order *string `json:"order,omitempty"`
	// Set when the calling account has marked the entry as a favorite, which is what puts it into the favorites  listing. For a file that is not marked it is null rather than false.
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// Tells a folder from a file, and so which of the two shapes the rest of the object has. A room is reported as a  folder here.
	FileEntryType *FileEntryType `json:"fileEntryType,omitempty"`
	// The identifier to pass back to the other operations of this entry. It is a number for storage on the portal  and a string for a connected third-party account, and it is unique only within its own kind, so files and  folders may carry the same value.
	Id *string `json:"id,omitempty"`
	// The section the entry ultimately lies in, as an identifier that can be listed like any other folder. For an  entry inside a room this is the rooms section, not the room.
	RootFolderId *string `json:"rootFolderId,omitempty"`
	// The folder the entry was deleted from, which is where restoring it puts it back. It is left out of the answer  unless the entry is in the trash.
	OriginId *string `json:"originId,omitempty"`
	// The room the entry was deleted from, left out of the answer for anything that was not deleted out of a room.
	OriginRoomId *string `json:"originRoomId,omitempty"`
	// The name of the folder the entry was deleted from, for showing where it would be restored to. It is null for  an entry that is not in the trash.
	OriginTitle *string `json:"originTitle,omitempty"`
	// The name of the room the entry was deleted from, null for anything that was not deleted out of a room.
	OriginRoomTitle *string `json:"originRoomTitle,omitempty"`
	// Whether the calling account may change who has access to the entry, and so whether offering a sharing dialog  for it makes sense. It is false in rooms whose access is fixed by the room itself, such as a private one, even  for its manager.
	CanShare *bool `json:"canShare,omitempty"`
	ShareSettings NullableAiFileEntryDtoAllOfShareSettings `json:"shareSettings,omitempty"`
	Security NullableAiFileEntryDtoAllOfSecurity `json:"security,omitempty"`
	AvailableShareRights NullableAiFileEntryDtoAllOfAvailableShareRights `json:"availableShareRights,omitempty"`
	// The token of the link the entry is being read through, which is the value the external-share operations expect  and which also has to be carried by the download and preview addresses. It is null whenever the entry is not  being read through a link.
	RequestToken *string `json:"requestToken,omitempty"`
	// Set when the link being used was made for this very entry, and false when the entry is reached through a link  to the room around it. It is null when no link is involved.
	External *bool `json:"external,omitempty"`
	// When the link being used stops working, written with the offset of the portal's time zone. It is null for a  link that never expires and whenever no link is involved.
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// Set when the link being used has already passed its expiration date, which is why the entry cannot be opened  even though it is described here. It is null when no link is involved.
	IsLinkExpired *bool `json:"isLinkExpired,omitempty"`
	// The folder this one is listed in. For a room it is the root of the section the room lives in, and for an entry  opened through a sharing link whose real parent the caller may not read it is the root of the section with the  entries shared with them.
	ParentId NullableString `json:"parentId,omitempty"`
	// How many files lie directly in the folder, without counting the subfolders. The roots of the `Rooms`, room  templates and default templates sections always report 0, because the number is not collected for them.
	FilesCount *int32 `json:"filesCount,omitempty"`
	// How many subfolders lie directly in the folder. For an AI room the two service subfolders it always holds are  subtracted, so the number matches what a listing of it shows, and the roots of the `Rooms` and templates  sections report 0.
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
	Logo *Logo `json:"logo,omitempty"`
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
	// How much space the files of the room may take, in bytes. It is the limit set on this room, or the portal  default for rooms when none was set. Null when the tariff of the portal does not count room statistics, when  room quotas are switched off, when the room lies in the archive or the trash, or when the caller may only read  it.
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

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ThirdPartyFolderDto) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *ThirdPartyFolderDto) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *ThirdPartyFolderDto) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *ThirdPartyFolderDto) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *ThirdPartyFolderDto) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *ThirdPartyFolderDto) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *ThirdPartyFolderDto) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *ThirdPartyFolderDto) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *ThirdPartyFolderDto) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *ThirdPartyFolderDto) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *ThirdPartyFolderDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *ThirdPartyFolderDto) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetAutoDelete() ApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret ApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetAutoDeleteOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given ApiDateTime and assigns it to the AutoDelete field.
func (o *ThirdPartyFolderDto) SetAutoDelete(v ApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *ThirdPartyFolderDto) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *ThirdPartyFolderDto) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *ThirdPartyFolderDto) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *ThirdPartyFolderDto) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *ThirdPartyFolderDto) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ThirdPartyFolderDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ThirdPartyFolderDto) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *ThirdPartyFolderDto) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *ThirdPartyFolderDto) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThirdPartyFolderDto) SetId(v string) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetRootFolderId() string {
	if o == nil || IsNil(o.RootFolderId) {
		var ret string
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetRootFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given string and assigns it to the RootFolderId field.
func (o *ThirdPartyFolderDto) SetRootFolderId(v string) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOriginId() string {
	if o == nil || IsNil(o.OriginId) {
		var ret string
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOriginIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given string and assigns it to the OriginId field.
func (o *ThirdPartyFolderDto) SetOriginId(v string) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOriginRoomId() string {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret string
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOriginRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given string and assigns it to the OriginRoomId field.
func (o *ThirdPartyFolderDto) SetOriginRoomId(v string) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle) {
		var ret string
		return ret
	}
	return *o.OriginTitle
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOriginTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginTitle) {
		return nil, false
	}
	return o.OriginTitle, true
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOriginTitleSet() bool {
	if o != nil && !IsNil(o.OriginTitle) {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given string and assigns it to the OriginTitle field.
func (o *ThirdPartyFolderDto) SetOriginTitle(v string) {
	o.OriginTitle = &v
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomTitle) {
		return nil, false
	}
	return o.OriginRoomTitle, true
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsOriginRoomTitleSet() bool {
	if o != nil && !IsNil(o.OriginRoomTitle) {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given string and assigns it to the OriginRoomTitle field.
func (o *ThirdPartyFolderDto) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle = &v
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *ThirdPartyFolderDto) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret AiFileEntryDtoAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableAiFileEntryDtoAllOfShareSettings and assigns it to the ShareSettings field.
func (o *ThirdPartyFolderDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *ThirdPartyFolderDto) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetSecurity() AiFileEntryDtoAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret AiFileEntryDtoAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableAiFileEntryDtoAllOfSecurity and assigns it to the Security field.
func (o *ThirdPartyFolderDto) SetSecurity(v AiFileEntryDtoAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *ThirdPartyFolderDto) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFolderDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret AiFileEntryDtoAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFolderDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableAiFileEntryDtoAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *ThirdPartyFolderDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *ThirdPartyFolderDto) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *ThirdPartyFolderDto) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsRequestTokenSet() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *ThirdPartyFolderDto) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *ThirdPartyFolderDto) SetExternal(v bool) {
	o.External = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *ThirdPartyFolderDto) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise.
func (o *ThirdPartyFolderDto) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkExpired) {
		return nil, false
	}
	return o.IsLinkExpired, true
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *ThirdPartyFolderDto) IsIsLinkExpiredSet() bool {
	if o != nil && !IsNil(o.IsLinkExpired) {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given bool and assigns it to the IsLinkExpired field.
func (o *ThirdPartyFolderDto) SetIsLinkExpired(v bool) {
	o.IsLinkExpired = &v
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
func (o *ThirdPartyFolderDto) GetLogo() Logo {
	if o == nil || IsNil(o.Logo) {
		var ret Logo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFolderDto) GetLogoOk() (*Logo, bool) {
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

// SetLogo gets a reference to the given Logo and assigns it to the Logo field.
func (o *ThirdPartyFolderDto) SetLogo(v Logo) {
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
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.SharedBy) {
		toSerialize["sharedBy"] = o.SharedBy
	}
	if !IsNil(o.OwnedBy) {
		toSerialize["ownedBy"] = o.OwnedBy
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.SharedForUser) {
		toSerialize["sharedForUser"] = o.SharedForUser
	}
	if !IsNil(o.SharedExternal) {
		toSerialize["sharedExternal"] = o.SharedExternal
	}
	if !IsNil(o.ParentShared) {
		toSerialize["parentShared"] = o.ParentShared
	}
	if !IsNil(o.ShortWebUrl) {
		toSerialize["shortWebUrl"] = o.ShortWebUrl
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.AutoDelete) {
		toSerialize["autoDelete"] = o.AutoDelete
	}
	if !IsNil(o.RootFolderType) {
		toSerialize["rootFolderType"] = o.RootFolderType
	}
	if !IsNil(o.ParentRoomType) {
		toSerialize["parentRoomType"] = o.ParentRoomType
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.ProviderItem) {
		toSerialize["providerItem"] = o.ProviderItem
	}
	if !IsNil(o.ProviderKey) {
		toSerialize["providerKey"] = o.ProviderKey
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.IsFavorite) {
		toSerialize["isFavorite"] = o.IsFavorite
	}
	if !IsNil(o.FileEntryType) {
		toSerialize["fileEntryType"] = o.FileEntryType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RootFolderId) {
		toSerialize["rootFolderId"] = o.RootFolderId
	}
	if !IsNil(o.OriginId) {
		toSerialize["originId"] = o.OriginId
	}
	if !IsNil(o.OriginRoomId) {
		toSerialize["originRoomId"] = o.OriginRoomId
	}
	if !IsNil(o.OriginTitle) {
		toSerialize["originTitle"] = o.OriginTitle
	}
	if !IsNil(o.OriginRoomTitle) {
		toSerialize["originRoomTitle"] = o.OriginRoomTitle
	}
	if !IsNil(o.CanShare) {
		toSerialize["canShare"] = o.CanShare
	}
	if o.ShareSettings.IsSet() {
		toSerialize["shareSettings"] = o.ShareSettings.Get()
	}
	if o.Security.IsSet() {
		toSerialize["security"] = o.Security.Get()
	}
	if o.AvailableShareRights.IsSet() {
		toSerialize["availableShareRights"] = o.AvailableShareRights.Get()
	}
	if !IsNil(o.RequestToken) {
		toSerialize["requestToken"] = o.RequestToken
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.IsLinkExpired) {
		toSerialize["isLinkExpired"] = o.IsLinkExpired
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

