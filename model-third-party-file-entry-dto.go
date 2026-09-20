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

// checks if the ThirdPartyFileEntryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyFileEntryDto{}

// ThirdPartyFileEntryDto The part of a file or folder that depends on how the entry is identified: by a number on the portal, or by a  string on a connected third-party account.
type ThirdPartyFileEntryDto struct {
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
	Id NullableString `json:"id,omitempty"`
	// The section the entry ultimately lies in, as an identifier that can be listed like any other folder. For an  entry inside a room this is the rooms section, not the room.
	RootFolderId NullableString `json:"rootFolderId,omitempty"`
	// The folder the entry was deleted from, which is where restoring it puts it back. It is left out of the answer  unless the entry is in the trash.
	OriginId NullableString `json:"originId,omitempty"`
	// The room the entry was deleted from, left out of the answer for anything that was not deleted out of a room.
	OriginRoomId NullableString `json:"originRoomId,omitempty"`
	// The name of the folder the entry was deleted from, for showing where it would be restored to. It is null for  an entry that is not in the trash.
	OriginTitle NullableString `json:"originTitle,omitempty"`
	// The name of the room the entry was deleted from, null for anything that was not deleted out of a room.
	OriginRoomTitle NullableString `json:"originRoomTitle,omitempty"`
	// Whether the calling account may change who has access to the entry, and so whether offering a sharing dialog  for it makes sense. It is false in rooms whose access is fixed by the room itself, such as a private one, even  for its manager.
	CanShare *bool `json:"canShare,omitempty"`
	ShareSettings NullableAiFileEntryDtoAllOfShareSettings `json:"shareSettings,omitempty"`
	Security NullableAiFileEntryDtoAllOfSecurity `json:"security,omitempty"`
	AvailableShareRights NullableAiFileEntryDtoAllOfAvailableShareRights `json:"availableShareRights,omitempty"`
	// The token of the link the entry is being read through, which is the value the external-share operations expect  and which also has to be carried by the download and preview addresses. It is null whenever the entry is not  being read through a link.
	RequestToken NullableString `json:"requestToken,omitempty"`
	// Set when the link being used was made for this very entry, and false when the entry is reached through a link  to the room around it. It is null when no link is involved.
	External NullableBool `json:"external,omitempty"`
	// When the link being used stops working, written with the offset of the portal's time zone. It is null for a  link that never expires and whenever no link is involved.
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// Set when the link being used has already passed its expiration date, which is why the entry cannot be opened  even though it is described here. It is null when no link is involved.
	IsLinkExpired NullableBool `json:"isLinkExpired,omitempty"`
}

// NewThirdPartyFileEntryDto instantiates a new ThirdPartyFileEntryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyFileEntryDto() *ThirdPartyFileEntryDto {
	this := ThirdPartyFileEntryDto{}
	return &this
}

// NewThirdPartyFileEntryDtoWithDefaults instantiates a new ThirdPartyFileEntryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyFileEntryDtoWithDefaults() *ThirdPartyFileEntryDto {
	this := ThirdPartyFileEntryDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ThirdPartyFileEntryDto) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *ThirdPartyFileEntryDto) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *ThirdPartyFileEntryDto) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *ThirdPartyFileEntryDto) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *ThirdPartyFileEntryDto) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *ThirdPartyFileEntryDto) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *ThirdPartyFileEntryDto) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *ThirdPartyFileEntryDto) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *ThirdPartyFileEntryDto) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *ThirdPartyFileEntryDto) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *ThirdPartyFileEntryDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *ThirdPartyFileEntryDto) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetAutoDelete() ApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret ApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetAutoDeleteOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given ApiDateTime and assigns it to the AutoDelete field.
func (o *ThirdPartyFileEntryDto) SetAutoDelete(v ApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *ThirdPartyFileEntryDto) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *ThirdPartyFileEntryDto) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *ThirdPartyFileEntryDto) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *ThirdPartyFileEntryDto) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *ThirdPartyFileEntryDto) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ThirdPartyFileEntryDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ThirdPartyFileEntryDto) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *ThirdPartyFileEntryDto) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *ThirdPartyFileEntryDto) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ThirdPartyFileEntryDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetId() {
	o.Id.Unset()
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetRootFolderId() string {
	if o == nil || IsNil(o.RootFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderId.Get()
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetRootFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderId.Get(), o.RootFolderId.IsSet()
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsRootFolderIdSet() bool {
	if o != nil && o.RootFolderId.IsSet() {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given NullableString and assigns it to the RootFolderId field.
func (o *ThirdPartyFileEntryDto) SetRootFolderId(v string) {
	o.RootFolderId.Set(&v)
}
// SetRootFolderIdNil sets the value for RootFolderId to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetRootFolderIdNil() {
	o.RootFolderId.Set(nil)
}

// UnsetRootFolderId ensures that no value is present for RootFolderId, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetRootFolderId() {
	o.RootFolderId.Unset()
}

// GetOriginId returns the OriginId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetOriginId() string {
	if o == nil || IsNil(o.OriginId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginId.Get()
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetOriginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginId.Get(), o.OriginId.IsSet()
}

// HasOriginId returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOriginIdSet() bool {
	if o != nil && o.OriginId.IsSet() {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given NullableString and assigns it to the OriginId field.
func (o *ThirdPartyFileEntryDto) SetOriginId(v string) {
	o.OriginId.Set(&v)
}
// SetOriginIdNil sets the value for OriginId to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetOriginIdNil() {
	o.OriginId.Set(nil)
}

// UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetOriginId() {
	o.OriginId.Unset()
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetOriginRoomId() string {
	if o == nil || IsNil(o.OriginRoomId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomId.Get()
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetOriginRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomId.Get(), o.OriginRoomId.IsSet()
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOriginRoomIdSet() bool {
	if o != nil && o.OriginRoomId.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given NullableString and assigns it to the OriginRoomId field.
func (o *ThirdPartyFileEntryDto) SetOriginRoomId(v string) {
	o.OriginRoomId.Set(&v)
}
// SetOriginRoomIdNil sets the value for OriginRoomId to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetOriginRoomIdNil() {
	o.OriginRoomId.Set(nil)
}

// UnsetOriginRoomId ensures that no value is present for OriginRoomId, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetOriginRoomId() {
	o.OriginRoomId.Unset()
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginTitle.Get()
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetOriginTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginTitle.Get(), o.OriginTitle.IsSet()
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOriginTitleSet() bool {
	if o != nil && o.OriginTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given NullableString and assigns it to the OriginTitle field.
func (o *ThirdPartyFileEntryDto) SetOriginTitle(v string) {
	o.OriginTitle.Set(&v)
}
// SetOriginTitleNil sets the value for OriginTitle to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetOriginTitleNil() {
	o.OriginTitle.Set(nil)
}

// UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetOriginTitle() {
	o.OriginTitle.Unset()
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle.Get()
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomTitle.Get(), o.OriginRoomTitle.IsSet()
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsOriginRoomTitleSet() bool {
	if o != nil && o.OriginRoomTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given NullableString and assigns it to the OriginRoomTitle field.
func (o *ThirdPartyFileEntryDto) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle.Set(&v)
}
// SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetOriginRoomTitleNil() {
	o.OriginRoomTitle.Set(nil)
}

// UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetOriginRoomTitle() {
	o.OriginRoomTitle.Unset()
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *ThirdPartyFileEntryDto) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret AiFileEntryDtoAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableAiFileEntryDtoAllOfShareSettings and assigns it to the ShareSettings field.
func (o *ThirdPartyFileEntryDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetSecurity() AiFileEntryDtoAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret AiFileEntryDtoAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableAiFileEntryDtoAllOfSecurity and assigns it to the Security field.
func (o *ThirdPartyFileEntryDto) SetSecurity(v AiFileEntryDtoAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret AiFileEntryDtoAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableAiFileEntryDtoAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *ThirdPartyFileEntryDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *ThirdPartyFileEntryDto) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetExternal() bool {
	if o == nil || IsNil(o.External.Get()) {
		var ret bool
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsExternalSet() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableBool and assigns it to the External field.
func (o *ThirdPartyFileEntryDto) SetExternal(v bool) {
	o.External.Set(&v)
}
// SetExternalNil sets the value for External to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetExternal() {
	o.External.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ThirdPartyFileEntryDto) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyFileEntryDto) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *ThirdPartyFileEntryDto) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyFileEntryDto) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired.Get()
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyFileEntryDto) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLinkExpired.Get(), o.IsLinkExpired.IsSet()
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *ThirdPartyFileEntryDto) IsIsLinkExpiredSet() bool {
	if o != nil && o.IsLinkExpired.IsSet() {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given NullableBool and assigns it to the IsLinkExpired field.
func (o *ThirdPartyFileEntryDto) SetIsLinkExpired(v bool) {
	o.IsLinkExpired.Set(&v)
}
// SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil
func (o *ThirdPartyFileEntryDto) SetIsLinkExpiredNil() {
	o.IsLinkExpired.Set(nil)
}

// UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
func (o *ThirdPartyFileEntryDto) UnsetIsLinkExpired() {
	o.IsLinkExpired.Unset()
}

func (o ThirdPartyFileEntryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyFileEntryDto) ToMap() (map[string]interface{}, error) {
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
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.RootFolderId.IsSet() {
		toSerialize["rootFolderId"] = o.RootFolderId.Get()
	}
	if o.OriginId.IsSet() {
		toSerialize["originId"] = o.OriginId.Get()
	}
	if o.OriginRoomId.IsSet() {
		toSerialize["originRoomId"] = o.OriginRoomId.Get()
	}
	if o.OriginTitle.IsSet() {
		toSerialize["originTitle"] = o.OriginTitle.Get()
	}
	if o.OriginRoomTitle.IsSet() {
		toSerialize["originRoomTitle"] = o.OriginRoomTitle.Get()
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
	if o.RequestToken.IsSet() {
		toSerialize["requestToken"] = o.RequestToken.Get()
	}
	if o.External.IsSet() {
		toSerialize["external"] = o.External.Get()
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.IsLinkExpired.IsSet() {
		toSerialize["isLinkExpired"] = o.IsLinkExpired.Get()
	}
	return toSerialize, nil
}

type NullableThirdPartyFileEntryDto struct {
	value *ThirdPartyFileEntryDto
	isSet bool
}

func (v NullableThirdPartyFileEntryDto) Get() *ThirdPartyFileEntryDto {
	return v.value
}

func (v *NullableThirdPartyFileEntryDto) Set(val *ThirdPartyFileEntryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyFileEntryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyFileEntryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyFileEntryDto(val *ThirdPartyFileEntryDto) *NullableThirdPartyFileEntryDto {
	return &NullableThirdPartyFileEntryDto{value: val, isSet: true}
}

func (v NullableThirdPartyFileEntryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyFileEntryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

