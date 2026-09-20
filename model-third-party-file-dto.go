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
	// Whether the PDF is a fillable form rather than a plain document. When the stored classification does not say,  the portal opens the file to find out, so the answer is reliable for a PDF and null for anything else.
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
	Dimensions *Size `json:"dimensions,omitempty"`
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

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ThirdPartyFileDto) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *ThirdPartyFileDto) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *ThirdPartyFileDto) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *ThirdPartyFileDto) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *ThirdPartyFileDto) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *ThirdPartyFileDto) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *ThirdPartyFileDto) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *ThirdPartyFileDto) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *ThirdPartyFileDto) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *ThirdPartyFileDto) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *ThirdPartyFileDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *ThirdPartyFileDto) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetAutoDelete() ApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret ApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetAutoDeleteOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given ApiDateTime and assigns it to the AutoDelete field.
func (o *ThirdPartyFileDto) SetAutoDelete(v ApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *ThirdPartyFileDto) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *ThirdPartyFileDto) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *ThirdPartyFileDto) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *ThirdPartyFileDto) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *ThirdPartyFileDto) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ThirdPartyFileDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ThirdPartyFileDto) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *ThirdPartyFileDto) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *ThirdPartyFileDto) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThirdPartyFileDto) SetId(v string) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetRootFolderId() string {
	if o == nil || IsNil(o.RootFolderId) {
		var ret string
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetRootFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given string and assigns it to the RootFolderId field.
func (o *ThirdPartyFileDto) SetRootFolderId(v string) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOriginId() string {
	if o == nil || IsNil(o.OriginId) {
		var ret string
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOriginIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given string and assigns it to the OriginId field.
func (o *ThirdPartyFileDto) SetOriginId(v string) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOriginRoomId() string {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret string
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOriginRoomIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given string and assigns it to the OriginRoomId field.
func (o *ThirdPartyFileDto) SetOriginRoomId(v string) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle) {
		var ret string
		return ret
	}
	return *o.OriginTitle
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOriginTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginTitle) {
		return nil, false
	}
	return o.OriginTitle, true
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOriginTitleSet() bool {
	if o != nil && !IsNil(o.OriginTitle) {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given string and assigns it to the OriginTitle field.
func (o *ThirdPartyFileDto) SetOriginTitle(v string) {
	o.OriginTitle = &v
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomTitle) {
		return nil, false
	}
	return o.OriginRoomTitle, true
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsOriginRoomTitleSet() bool {
	if o != nil && !IsNil(o.OriginRoomTitle) {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given string and assigns it to the OriginRoomTitle field.
func (o *ThirdPartyFileDto) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle = &v
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *ThirdPartyFileDto) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret AiFileEntryDtoAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableAiFileEntryDtoAllOfShareSettings and assigns it to the ShareSettings field.
func (o *ThirdPartyFileDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *ThirdPartyFileDto) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetSecurity() AiFileEntryDtoAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret AiFileEntryDtoAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableAiFileEntryDtoAllOfSecurity and assigns it to the Security field.
func (o *ThirdPartyFileDto) SetSecurity(v AiFileEntryDtoAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *ThirdPartyFileDto) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret AiFileEntryDtoAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableAiFileEntryDtoAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *ThirdPartyFileDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *ThirdPartyFileDto) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *ThirdPartyFileDto) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsRequestTokenSet() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *ThirdPartyFileDto) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *ThirdPartyFileDto) SetExternal(v bool) {
	o.External = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *ThirdPartyFileDto) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise.
func (o *ThirdPartyFileDto) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkExpired) {
		return nil, false
	}
	return o.IsLinkExpired, true
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *ThirdPartyFileDto) IsIsLinkExpiredSet() bool {
	if o != nil && !IsNil(o.IsLinkExpired) {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given bool and assigns it to the IsLinkExpired field.
func (o *ThirdPartyFileDto) SetIsLinkExpired(v bool) {
	o.IsLinkExpired = &v
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
func (o *ThirdPartyFileDto) GetDimensions() Size {
	if o == nil || IsNil(o.Dimensions) {
		var ret Size
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileDto) GetDimensionsOk() (*Size, bool) {
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

// SetDimensions gets a reference to the given Size and assigns it to the Dimensions field.
func (o *ThirdPartyFileDto) SetDimensions(v Size) {
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

