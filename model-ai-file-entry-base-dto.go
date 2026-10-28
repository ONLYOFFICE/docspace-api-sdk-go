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

// checks if the AiFileEntryBaseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileEntryBaseDto{}

// AiFileEntryBaseDto What every file and folder in an answer has in common; the concrete shape is a file or a folder, told apart by the  entry type.
type AiFileEntryBaseDto struct {
	// The name shown for the entry. For a file it carries the extension, which is how the format is recognised, and  for a room it is the room name.
	Title NullableString `json:"title,omitempty"`
	// The level the calling account holds on this entry, resolved from its own rights, the groups it belongs to and  any link it came in through. It is the level itself, not what the account may do with it - the action flags  below answer that.
	Access *AiFileShare `json:"access,omitempty"`
	// Who gave the calling account the access it is using. It is filled in only while the entry is being read  through a share, and never for a caller without an account.
	SharedBy *AiEmployeeDto `json:"sharedBy,omitempty"`
	// Who owns the place the entry is shared from - the creator of the room it lies in, or of the personal section  that holds it. It is filled in only while the entry is being read through a share, and never for a caller  without an account.
	OwnedBy *AiEmployeeDto `json:"ownedBy,omitempty"`
	// Whether at least one external link exists for the entry, whichever kind. It says nothing about accounts and  groups - those are counted by the flag for members below.
	Shared *bool `json:"shared,omitempty"`
	// Whether at least one account or group has been given rights on the entry directly, as opposed to reaching it  through a link or through the room around it.
	SharedForUser *bool `json:"sharedForUser,omitempty"`
	// Whether one of the entry's links is open to people outside the portal, as opposed to a link that only its own  members can follow. This is the flag to watch when the concern is who can reach the content from outside.
	SharedExternal *bool `json:"sharedExternal,omitempty"`
	// Whether the entry is reachable because the room or folder around it is shared, rather than through rights of  its own. A copy or a move takes the entry out of that scope.
	ParentShared *bool `json:"parentShared,omitempty"`
	// A shortened address that opens the entry through the link it is being read with. It is an empty string  whenever no link applies, which is the usual case for a member browsing their own rooms.
	ShortWebUrl NullableString `json:"shortWebUrl,omitempty"`
	// When the entry was created, written with the offset of the portal's time zone. For a file restored from an  older version this is still the moment the file first appeared.
	Created *AiApiDateTime `json:"created,omitempty"`
	// Who created the entry. It is null for a caller without an account, who is told nothing about the portal's  members.
	CreatedBy *AiEmployeeDto `json:"createdBy,omitempty"`
	// When the entry last changed, written with the offset of the portal's time zone. It is never reported as  earlier than the creation moment, so the two can be compared safely.
	Updated *AiApiDateTime `json:"updated,omitempty"`
	// When the entry will disappear on its own, written with the offset of the portal's time zone. It is filled in  only where a removal is actually scheduled - something in the trash while the portal cleans it up  automatically, or a guest's own documents - so a null means nothing is scheduled rather than that the entry is  permanent.
	AutoDelete *AiApiDateTime `json:"autoDelete,omitempty"`
	// The section the entry ultimately belongs to, which is what tells a personal document from one inside a room,  from a template and from something in the trash or the archive.
	RootFolderType *AiFolderType `json:"rootFolderType,omitempty"`
	// The kind of room the entry lies in, which decides what the room allows - filling forms, public links,  indexing. It is null for an entry that is not inside a room at all.
	ParentRoomType *AiFolderType `json:"parentRoomType,omitempty"`
	// Who changed the entry last. It is null for a caller without an account.
	UpdatedBy *AiEmployeeDto `json:"updatedBy,omitempty"`
	// Set when the entry is stored on a connected third-party account rather than on the portal, and null when it is  stored on the portal. Such an entry is identified by a string rather than a number, and some operations skip  it.
	ProviderItem NullableBool `json:"providerItem,omitempty"`
	// Which third-party service holds the entry, matching the keys accepted by the third-party operations. It is  null for an entry stored on the portal.
	ProviderKey NullableString `json:"providerKey,omitempty"`
	// The connected account the entry comes from, for telling apart two connections to the same service. It is null  for an entry stored on the portal.
	ProviderId NullableInt32 `json:"providerId,omitempty"`
	// The place of the entry in a room where the members arrange the content themselves, given as the position of  the entry preceded by the positions of the folders leading to it, separated by dots. It is empty when nothing  has been arranged.
	Order NullableString `json:"order,omitempty"`
	// Set when the calling account has marked the entry as a favorite, which is what puts it into the favorites  listing. For a file that is not marked it is null rather than false.
	IsFavorite NullableBool `json:"isFavorite,omitempty"`
	// Tells a folder from a file, and so which of the two shapes the rest of the object has. A room is reported as a  folder here.
	FileEntryType *AiFileEntryType `json:"fileEntryType,omitempty"`
}

// NewAiFileEntryBaseDto instantiates a new AiFileEntryBaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileEntryBaseDto() *AiFileEntryBaseDto {
	this := AiFileEntryBaseDto{}
	return &this
}

// NewAiFileEntryBaseDtoWithDefaults instantiates a new AiFileEntryBaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileEntryBaseDtoWithDefaults() *AiFileEntryBaseDto {
	this := AiFileEntryBaseDto{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AiFileEntryBaseDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AiFileEntryBaseDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetTitle() {
	o.Title.Unset()
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetAccess() AiFileShare {
	if o == nil || IsNil(o.Access) {
		var ret AiFileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetAccessOk() (*AiFileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AiFileShare and assigns it to the Access field.
func (o *AiFileEntryBaseDto) SetAccess(v AiFileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetSharedBy() AiEmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetSharedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given AiEmployeeDto and assigns it to the SharedBy field.
func (o *AiFileEntryBaseDto) SetSharedBy(v AiEmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetOwnedBy() AiEmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetOwnedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given AiEmployeeDto and assigns it to the OwnedBy field.
func (o *AiFileEntryBaseDto) SetOwnedBy(v AiEmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *AiFileEntryBaseDto) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *AiFileEntryBaseDto) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *AiFileEntryBaseDto) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *AiFileEntryBaseDto) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl.Get()
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetShortWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortWebUrl.Get(), o.ShortWebUrl.IsSet()
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsShortWebUrlSet() bool {
	if o != nil && o.ShortWebUrl.IsSet() {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given NullableString and assigns it to the ShortWebUrl field.
func (o *AiFileEntryBaseDto) SetShortWebUrl(v string) {
	o.ShortWebUrl.Set(&v)
}
// SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil
func (o *AiFileEntryBaseDto) SetShortWebUrlNil() {
	o.ShortWebUrl.Set(nil)
}

// UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetShortWebUrl() {
	o.ShortWebUrl.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetCreated() AiApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret AiApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetCreatedOk() (*AiApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given AiApiDateTime and assigns it to the Created field.
func (o *AiFileEntryBaseDto) SetCreated(v AiApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetCreatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetCreatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AiEmployeeDto and assigns it to the CreatedBy field.
func (o *AiFileEntryBaseDto) SetCreatedBy(v AiEmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetUpdated() AiApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret AiApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetUpdatedOk() (*AiApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given AiApiDateTime and assigns it to the Updated field.
func (o *AiFileEntryBaseDto) SetUpdated(v AiApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetAutoDelete() AiApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret AiApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetAutoDeleteOk() (*AiApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given AiApiDateTime and assigns it to the AutoDelete field.
func (o *AiFileEntryBaseDto) SetAutoDelete(v AiApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetRootFolderType() AiFolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret AiFolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetRootFolderTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given AiFolderType and assigns it to the RootFolderType field.
func (o *AiFileEntryBaseDto) SetRootFolderType(v AiFolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetParentRoomType() AiFolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret AiFolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetParentRoomTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given AiFolderType and assigns it to the ParentRoomType field.
func (o *AiFileEntryBaseDto) SetParentRoomType(v AiFolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetUpdatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetUpdatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given AiEmployeeDto and assigns it to the UpdatedBy field.
func (o *AiFileEntryBaseDto) SetUpdatedBy(v AiEmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem.Get()) {
		var ret bool
		return ret
	}
	return *o.ProviderItem.Get()
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetProviderItemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderItem.Get(), o.ProviderItem.IsSet()
}

// HasProviderItem returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsProviderItemSet() bool {
	if o != nil && o.ProviderItem.IsSet() {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given NullableBool and assigns it to the ProviderItem field.
func (o *AiFileEntryBaseDto) SetProviderItem(v bool) {
	o.ProviderItem.Set(&v)
}
// SetProviderItemNil sets the value for ProviderItem to be an explicit nil
func (o *AiFileEntryBaseDto) SetProviderItemNil() {
	o.ProviderItem.Set(nil)
}

// UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetProviderItem() {
	o.ProviderItem.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *AiFileEntryBaseDto) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *AiFileEntryBaseDto) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsProviderIdSet() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableInt32 and assigns it to the ProviderId field.
func (o *AiFileEntryBaseDto) SetProviderId(v int32) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *AiFileEntryBaseDto) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetOrder() string {
	if o == nil || IsNil(o.Order.Get()) {
		var ret string
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsOrderSet() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableString and assigns it to the Order field.
func (o *AiFileEntryBaseDto) SetOrder(v string) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *AiFileEntryBaseDto) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetOrder() {
	o.Order.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryBaseDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryBaseDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsIsFavoriteSet() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *AiFileEntryBaseDto) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *AiFileEntryBaseDto) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *AiFileEntryBaseDto) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *AiFileEntryBaseDto) GetFileEntryType() AiFileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret AiFileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryBaseDto) GetFileEntryTypeOk() (*AiFileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *AiFileEntryBaseDto) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given AiFileEntryType and assigns it to the FileEntryType field.
func (o *AiFileEntryBaseDto) SetFileEntryType(v AiFileEntryType) {
	o.FileEntryType = &v
}

func (o AiFileEntryBaseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileEntryBaseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
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
	if o.ShortWebUrl.IsSet() {
		toSerialize["shortWebUrl"] = o.ShortWebUrl.Get()
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
	if o.ProviderItem.IsSet() {
		toSerialize["providerItem"] = o.ProviderItem.Get()
	}
	if o.ProviderKey.IsSet() {
		toSerialize["providerKey"] = o.ProviderKey.Get()
	}
	if o.ProviderId.IsSet() {
		toSerialize["providerId"] = o.ProviderId.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.IsFavorite.IsSet() {
		toSerialize["isFavorite"] = o.IsFavorite.Get()
	}
	if !IsNil(o.FileEntryType) {
		toSerialize["fileEntryType"] = o.FileEntryType
	}
	return toSerialize, nil
}

type NullableAiFileEntryBaseDto struct {
	value *AiFileEntryBaseDto
	isSet bool
}

func (v NullableAiFileEntryBaseDto) Get() *AiFileEntryBaseDto {
	return v.value
}

func (v *NullableAiFileEntryBaseDto) Set(val *AiFileEntryBaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryBaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryBaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryBaseDto(val *AiFileEntryBaseDto) *NullableAiFileEntryBaseDto {
	return &NullableAiFileEntryBaseDto{value: val, isSet: true}
}

func (v NullableAiFileEntryBaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryBaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

