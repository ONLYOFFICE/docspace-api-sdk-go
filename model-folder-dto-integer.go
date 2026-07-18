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

// checks if the FolderDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderDtoInteger{}

// FolderDtoInteger The folder parameters.
type FolderDtoInteger struct {
	// The file entry title.
	Title NullableString `json:"title,omitempty"`
	Access *FileShare `json:"access,omitempty"`
	SharedBy *EmployeeDto `json:"sharedBy,omitempty"`
	OwnedBy *EmployeeDto `json:"ownedBy,omitempty"`
	// Specifies if the file entry is shared via link or not.
	Shared *bool `json:"shared,omitempty"`
	// Specifies if the file entry is shared for user or not.
	SharedForUser *bool `json:"sharedForUser,omitempty"`
	// Specifies if the file entry is shared via a public (non-internal) external link.
	SharedExternal *bool `json:"sharedExternal,omitempty"`
	// Indicates whether the parent entity is shared.
	ParentShared *bool `json:"parentShared,omitempty"`
	// The short Web URL.
	ShortWebUrl NullableString `json:"shortWebUrl,omitempty"`
	Created *ApiDateTime `json:"created,omitempty"`
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
	Updated *ApiDateTime `json:"updated,omitempty"`
	AutoDelete *ApiDateTime `json:"autoDelete,omitempty"`
	RootFolderType *FolderType `json:"rootFolderType,omitempty"`
	ParentRoomType *FolderType `json:"parentRoomType,omitempty"`
	UpdatedBy *EmployeeDto `json:"updatedBy,omitempty"`
	// Specifies if the file entry provider is specified or not.
	ProviderItem NullableBool `json:"providerItem,omitempty"`
	// The provider key of the file entry.
	ProviderKey NullableString `json:"providerKey,omitempty"`
	// The provider ID of the file entry.
	ProviderId NullableInt32 `json:"providerId,omitempty"`
	// The order of the file entry.
	Order NullableString `json:"order,omitempty"`
	// Specifies if the file is a favorite or not.
	IsFavorite NullableBool `json:"isFavorite,omitempty"`
	FileEntryType *FileEntryType `json:"fileEntryType,omitempty"`
	// The file entry ID.
	Id *int32 `json:"id,omitempty"`
	// The root folder ID of the file entry.
	RootFolderId *int32 `json:"rootFolderId,omitempty"`
	// The origin ID of the file entry.
	OriginId *int32 `json:"originId,omitempty"`
	// The origin room ID of the file entry.
	OriginRoomId *int32 `json:"originRoomId,omitempty"`
	// The origin title of the file entry.
	OriginTitle NullableString `json:"originTitle,omitempty"`
	// The origin room title of the file entry.
	OriginRoomTitle NullableString `json:"originRoomTitle,omitempty"`
	// Specifies if the file entry can be shared or not.
	CanShare *bool `json:"canShare,omitempty"`
	ShareSettings NullableFileEntryDtoIntegerAllOfShareSettings `json:"shareSettings,omitempty"`
	Security NullableFileEntryDtoIntegerAllOfSecurity `json:"security,omitempty"`
	AvailableShareRights NullableFileEntryDtoIntegerAllOfAvailableShareRights `json:"availableShareRights,omitempty"`
	// The request token of the file entry.
	RequestToken NullableString `json:"requestToken,omitempty"`
	// Specifies if the folder can be accessed via an external link or not.
	External NullableBool `json:"external,omitempty"`
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// Indicates whether the shareable link associated with the file or folder has expired.
	IsLinkExpired NullableBool `json:"isLinkExpired,omitempty"`
	// The parent folder ID of the folder.
	ParentId *int32 `json:"parentId,omitempty"`
	// The number of files that the folder contains.
	FilesCount *int32 `json:"filesCount,omitempty"`
	// The number of folders that the folder contains.
	FoldersCount *int32 `json:"foldersCount,omitempty"`
	// Specifies if the folder can be shared or not.
	IsShareable NullableBool `json:"isShareable,omitempty"`
	// The new element index in the folder.
	New *int32 `json:"new,omitempty"`
	// Specifies if the folder notifications are enabled or not.
	Mute *bool `json:"mute,omitempty"`
	// The list of tags of the folder.
	Tags []string `json:"tags,omitempty"`
	Logo *Logo `json:"logo,omitempty"`
	// Specifies if the folder is pinned or not.
	Pinned *bool `json:"pinned,omitempty"`
	RoomType *RoomType `json:"roomType,omitempty"`
	// Specifies if the folder is private or not.
	Private *bool `json:"private,omitempty"`
	// Specifies if the folder is indexed or not.
	Indexing *bool `json:"indexing,omitempty"`
	// Specifies if the folder can be downloaded or not.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	Lifetime *RoomDataLifetimeDto `json:"lifetime,omitempty"`
	Watermark *WatermarkDto `json:"watermark,omitempty"`
	Type *FolderType `json:"type,omitempty"`
	// Specifies if the folder is placed in the room or not.
	InRoom NullableBool `json:"inRoom,omitempty"`
	// The folder quota limit.
	QuotaLimit NullableInt64 `json:"quotaLimit,omitempty"`
	// Specifies if the folder room has a custom quota or not.
	IsCustomQuota NullableBool `json:"isCustomQuota,omitempty"`
	// How much folder space is used (counter).
	UsedSpace NullableInt64 `json:"usedSpace,omitempty"`
	// Specifies if the folder is password protected or not.
	PasswordProtected NullableBool `json:"passwordProtected,omitempty"`
	// Specifies if an external link to the folder is expired or not.
	// Deprecated
	Expired NullableBool `json:"expired,omitempty"`
	ChatSettings *ChatSettingsDto `json:"chatSettings,omitempty"`
	RootRoomType *RoomType `json:"rootRoomType,omitempty"`
	// Specifies whether to save form data as XLSX file.
	SaveFormAsXLSX NullableBool `json:"saveFormAsXLSX,omitempty"`
	// Specifies whether to send form data to external database.
	SendFormToExternalDB NullableBool `json:"sendFormToExternalDB,omitempty"`
	// The original form ID that corresponds to this FormFillingFolderDone folder.
	OriginalFormId NullableInt32 `json:"originalFormId,omitempty"`
}

// NewFolderDtoInteger instantiates a new FolderDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderDtoInteger() *FolderDtoInteger {
	this := FolderDtoInteger{}
	return &this
}

// NewFolderDtoIntegerWithDefaults instantiates a new FolderDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderDtoIntegerWithDefaults() *FolderDtoInteger {
	this := FolderDtoInteger{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FolderDtoInteger) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FolderDtoInteger) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FolderDtoInteger) UnsetTitle() {
	o.Title.Unset()
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FolderDtoInteger) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *FolderDtoInteger) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *FolderDtoInteger) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *FolderDtoInteger) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *FolderDtoInteger) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *FolderDtoInteger) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *FolderDtoInteger) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl.Get()
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetShortWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortWebUrl.Get(), o.ShortWebUrl.IsSet()
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsShortWebUrlSet() bool {
	if o != nil && o.ShortWebUrl.IsSet() {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given NullableString and assigns it to the ShortWebUrl field.
func (o *FolderDtoInteger) SetShortWebUrl(v string) {
	o.ShortWebUrl.Set(&v)
}
// SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil
func (o *FolderDtoInteger) SetShortWebUrlNil() {
	o.ShortWebUrl.Set(nil)
}

// UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
func (o *FolderDtoInteger) UnsetShortWebUrl() {
	o.ShortWebUrl.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *FolderDtoInteger) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *FolderDtoInteger) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *FolderDtoInteger) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetAutoDelete() ApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret ApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetAutoDeleteOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given ApiDateTime and assigns it to the AutoDelete field.
func (o *FolderDtoInteger) SetAutoDelete(v ApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *FolderDtoInteger) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *FolderDtoInteger) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *FolderDtoInteger) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem.Get()) {
		var ret bool
		return ret
	}
	return *o.ProviderItem.Get()
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetProviderItemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderItem.Get(), o.ProviderItem.IsSet()
}

// HasProviderItem returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsProviderItemSet() bool {
	if o != nil && o.ProviderItem.IsSet() {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given NullableBool and assigns it to the ProviderItem field.
func (o *FolderDtoInteger) SetProviderItem(v bool) {
	o.ProviderItem.Set(&v)
}
// SetProviderItemNil sets the value for ProviderItem to be an explicit nil
func (o *FolderDtoInteger) SetProviderItemNil() {
	o.ProviderItem.Set(nil)
}

// UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
func (o *FolderDtoInteger) UnsetProviderItem() {
	o.ProviderItem.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *FolderDtoInteger) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *FolderDtoInteger) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *FolderDtoInteger) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsProviderIdSet() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableInt32 and assigns it to the ProviderId field.
func (o *FolderDtoInteger) SetProviderId(v int32) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *FolderDtoInteger) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *FolderDtoInteger) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetOrder() string {
	if o == nil || IsNil(o.Order.Get()) {
		var ret string
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOrderSet() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableString and assigns it to the Order field.
func (o *FolderDtoInteger) SetOrder(v string) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *FolderDtoInteger) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *FolderDtoInteger) UnsetOrder() {
	o.Order.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIsFavoriteSet() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *FolderDtoInteger) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *FolderDtoInteger) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *FolderDtoInteger) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *FolderDtoInteger) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FolderDtoInteger) SetId(v int32) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetRootFolderId() int32 {
	if o == nil || IsNil(o.RootFolderId) {
		var ret int32
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetRootFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given int32 and assigns it to the RootFolderId field.
func (o *FolderDtoInteger) SetRootFolderId(v int32) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetOriginId() int32 {
	if o == nil || IsNil(o.OriginId) {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetOriginIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *FolderDtoInteger) SetOriginId(v int32) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetOriginRoomId() int32 {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret int32
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetOriginRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given int32 and assigns it to the OriginRoomId field.
func (o *FolderDtoInteger) SetOriginRoomId(v int32) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginTitle.Get()
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetOriginTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginTitle.Get(), o.OriginTitle.IsSet()
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOriginTitleSet() bool {
	if o != nil && o.OriginTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given NullableString and assigns it to the OriginTitle field.
func (o *FolderDtoInteger) SetOriginTitle(v string) {
	o.OriginTitle.Set(&v)
}
// SetOriginTitleNil sets the value for OriginTitle to be an explicit nil
func (o *FolderDtoInteger) SetOriginTitleNil() {
	o.OriginTitle.Set(nil)
}

// UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
func (o *FolderDtoInteger) UnsetOriginTitle() {
	o.OriginTitle.Unset()
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle.Get()
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomTitle.Get(), o.OriginRoomTitle.IsSet()
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOriginRoomTitleSet() bool {
	if o != nil && o.OriginRoomTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given NullableString and assigns it to the OriginRoomTitle field.
func (o *FolderDtoInteger) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle.Set(&v)
}
// SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil
func (o *FolderDtoInteger) SetOriginRoomTitleNil() {
	o.OriginRoomTitle.Set(nil)
}

// UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
func (o *FolderDtoInteger) UnsetOriginRoomTitle() {
	o.OriginRoomTitle.Unset()
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *FolderDtoInteger) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret FileEntryDtoIntegerAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableFileEntryDtoIntegerAllOfShareSettings and assigns it to the ShareSettings field.
func (o *FolderDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *FolderDtoInteger) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *FolderDtoInteger) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret FileEntryDtoIntegerAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableFileEntryDtoIntegerAllOfSecurity and assigns it to the Security field.
func (o *FolderDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *FolderDtoInteger) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *FolderDtoInteger) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret FileEntryDtoIntegerAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableFileEntryDtoIntegerAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *FolderDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *FolderDtoInteger) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *FolderDtoInteger) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *FolderDtoInteger) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *FolderDtoInteger) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *FolderDtoInteger) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetExternal() bool {
	if o == nil || IsNil(o.External.Get()) {
		var ret bool
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsExternalSet() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableBool and assigns it to the External field.
func (o *FolderDtoInteger) SetExternal(v bool) {
	o.External.Set(&v)
}
// SetExternalNil sets the value for External to be an explicit nil
func (o *FolderDtoInteger) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *FolderDtoInteger) UnsetExternal() {
	o.External.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *FolderDtoInteger) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired.Get()
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLinkExpired.Get(), o.IsLinkExpired.IsSet()
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIsLinkExpiredSet() bool {
	if o != nil && o.IsLinkExpired.IsSet() {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given NullableBool and assigns it to the IsLinkExpired field.
func (o *FolderDtoInteger) SetIsLinkExpired(v bool) {
	o.IsLinkExpired.Set(&v)
}
// SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil
func (o *FolderDtoInteger) SetIsLinkExpiredNil() {
	o.IsLinkExpired.Set(nil)
}

// UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
func (o *FolderDtoInteger) UnsetIsLinkExpired() {
	o.IsLinkExpired.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsParentIdSet() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *FolderDtoInteger) SetParentId(v int32) {
	o.ParentId = &v
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsFilesCountSet() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *FolderDtoInteger) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetFoldersCount returns the FoldersCount field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetFoldersCount() int32 {
	if o == nil || IsNil(o.FoldersCount) {
		var ret int32
		return ret
	}
	return *o.FoldersCount
}

// GetFoldersCountOk returns a tuple with the FoldersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetFoldersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FoldersCount) {
		return nil, false
	}
	return o.FoldersCount, true
}

// HasFoldersCount returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsFoldersCountSet() bool {
	if o != nil && !IsNil(o.FoldersCount) {
		return true
	}

	return false
}

// SetFoldersCount gets a reference to the given int32 and assigns it to the FoldersCount field.
func (o *FolderDtoInteger) SetFoldersCount(v int32) {
	o.FoldersCount = &v
}

// GetIsShareable returns the IsShareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetIsShareable() bool {
	if o == nil || IsNil(o.IsShareable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsShareable.Get()
}

// GetIsShareableOk returns a tuple with the IsShareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetIsShareableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsShareable.Get(), o.IsShareable.IsSet()
}

// HasIsShareable returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIsShareableSet() bool {
	if o != nil && o.IsShareable.IsSet() {
		return true
	}

	return false
}

// SetIsShareable gets a reference to the given NullableBool and assigns it to the IsShareable field.
func (o *FolderDtoInteger) SetIsShareable(v bool) {
	o.IsShareable.Set(&v)
}
// SetIsShareableNil sets the value for IsShareable to be an explicit nil
func (o *FolderDtoInteger) SetIsShareableNil() {
	o.IsShareable.Set(nil)
}

// UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
func (o *FolderDtoInteger) UnsetIsShareable() {
	o.IsShareable.Unset()
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *FolderDtoInteger) SetNew(v int32) {
	o.New = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *FolderDtoInteger) SetMute(v bool) {
	o.Mute = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FolderDtoInteger) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetLogo() Logo {
	if o == nil || IsNil(o.Logo) {
		var ret Logo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetLogoOk() (*Logo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given Logo and assigns it to the Logo field.
func (o *FolderDtoInteger) SetLogo(v Logo) {
	o.Logo = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsPinnedSet() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *FolderDtoInteger) SetPinned(v bool) {
	o.Pinned = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetRoomType() RoomType {
	if o == nil || IsNil(o.RoomType) {
		var ret RoomType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsRoomTypeSet() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given RoomType and assigns it to the RoomType field.
func (o *FolderDtoInteger) SetRoomType(v RoomType) {
	o.RoomType = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *FolderDtoInteger) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *FolderDtoInteger) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *FolderDtoInteger) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *FolderDtoInteger) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetWatermark() WatermarkDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetWatermarkOk() (*WatermarkDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkDto and assigns it to the Watermark field.
func (o *FolderDtoInteger) SetWatermark(v WatermarkDto) {
	o.Watermark = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetType() FolderType {
	if o == nil || IsNil(o.Type) {
		var ret FolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FolderType and assigns it to the Type field.
func (o *FolderDtoInteger) SetType(v FolderType) {
	o.Type = &v
}

// GetInRoom returns the InRoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetInRoom() bool {
	if o == nil || IsNil(o.InRoom.Get()) {
		var ret bool
		return ret
	}
	return *o.InRoom.Get()
}

// GetInRoomOk returns a tuple with the InRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetInRoomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InRoom.Get(), o.InRoom.IsSet()
}

// HasInRoom returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsInRoomSet() bool {
	if o != nil && o.InRoom.IsSet() {
		return true
	}

	return false
}

// SetInRoom gets a reference to the given NullableBool and assigns it to the InRoom field.
func (o *FolderDtoInteger) SetInRoom(v bool) {
	o.InRoom.Set(&v)
}
// SetInRoomNil sets the value for InRoom to be an explicit nil
func (o *FolderDtoInteger) SetInRoomNil() {
	o.InRoom.Set(nil)
}

// UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
func (o *FolderDtoInteger) UnsetInRoom() {
	o.InRoom.Unset()
}

// GetQuotaLimit returns the QuotaLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetQuotaLimit() int64 {
	if o == nil || IsNil(o.QuotaLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.QuotaLimit.Get()
}

// GetQuotaLimitOk returns a tuple with the QuotaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetQuotaLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaLimit.Get(), o.QuotaLimit.IsSet()
}

// HasQuotaLimit returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsQuotaLimitSet() bool {
	if o != nil && o.QuotaLimit.IsSet() {
		return true
	}

	return false
}

// SetQuotaLimit gets a reference to the given NullableInt64 and assigns it to the QuotaLimit field.
func (o *FolderDtoInteger) SetQuotaLimit(v int64) {
	o.QuotaLimit.Set(&v)
}
// SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil
func (o *FolderDtoInteger) SetQuotaLimitNil() {
	o.QuotaLimit.Set(nil)
}

// UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
func (o *FolderDtoInteger) UnsetQuotaLimit() {
	o.QuotaLimit.Unset()
}

// GetIsCustomQuota returns the IsCustomQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetIsCustomQuota() bool {
	if o == nil || IsNil(o.IsCustomQuota.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCustomQuota.Get()
}

// GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetIsCustomQuotaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCustomQuota.Get(), o.IsCustomQuota.IsSet()
}

// HasIsCustomQuota returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsIsCustomQuotaSet() bool {
	if o != nil && o.IsCustomQuota.IsSet() {
		return true
	}

	return false
}

// SetIsCustomQuota gets a reference to the given NullableBool and assigns it to the IsCustomQuota field.
func (o *FolderDtoInteger) SetIsCustomQuota(v bool) {
	o.IsCustomQuota.Set(&v)
}
// SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil
func (o *FolderDtoInteger) SetIsCustomQuotaNil() {
	o.IsCustomQuota.Set(nil)
}

// UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
func (o *FolderDtoInteger) UnsetIsCustomQuota() {
	o.IsCustomQuota.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.UsedSpace.Get()
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetUsedSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedSpace.Get(), o.UsedSpace.IsSet()
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsUsedSpaceSet() bool {
	if o != nil && o.UsedSpace.IsSet() {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given NullableInt64 and assigns it to the UsedSpace field.
func (o *FolderDtoInteger) SetUsedSpace(v int64) {
	o.UsedSpace.Set(&v)
}
// SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil
func (o *FolderDtoInteger) SetUsedSpaceNil() {
	o.UsedSpace.Set(nil)
}

// UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
func (o *FolderDtoInteger) UnsetUsedSpace() {
	o.UsedSpace.Unset()
}

// GetPasswordProtected returns the PasswordProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetPasswordProtected() bool {
	if o == nil || IsNil(o.PasswordProtected.Get()) {
		var ret bool
		return ret
	}
	return *o.PasswordProtected.Get()
}

// GetPasswordProtectedOk returns a tuple with the PasswordProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetPasswordProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProtected.Get(), o.PasswordProtected.IsSet()
}

// HasPasswordProtected returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsPasswordProtectedSet() bool {
	if o != nil && o.PasswordProtected.IsSet() {
		return true
	}

	return false
}

// SetPasswordProtected gets a reference to the given NullableBool and assigns it to the PasswordProtected field.
func (o *FolderDtoInteger) SetPasswordProtected(v bool) {
	o.PasswordProtected.Set(&v)
}
// SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil
func (o *FolderDtoInteger) SetPasswordProtectedNil() {
	o.PasswordProtected.Set(nil)
}

// UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
func (o *FolderDtoInteger) UnsetPasswordProtected() {
	o.PasswordProtected.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *FolderDtoInteger) GetExpired() bool {
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
func (o *FolderDtoInteger) GetExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expired.Get(), o.Expired.IsSet()
}

// HasExpired returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsExpiredSet() bool {
	if o != nil && o.Expired.IsSet() {
		return true
	}

	return false
}

// SetExpired gets a reference to the given NullableBool and assigns it to the Expired field.
// Deprecated
func (o *FolderDtoInteger) SetExpired(v bool) {
	o.Expired.Set(&v)
}
// SetExpiredNil sets the value for Expired to be an explicit nil
func (o *FolderDtoInteger) SetExpiredNil() {
	o.Expired.Set(nil)
}

// UnsetExpired ensures that no value is present for Expired, not even an explicit nil
func (o *FolderDtoInteger) UnsetExpired() {
	o.Expired.Unset()
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetChatSettings() ChatSettingsDto {
	if o == nil || IsNil(o.ChatSettings) {
		var ret ChatSettingsDto
		return ret
	}
	return *o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetChatSettingsOk() (*ChatSettingsDto, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return nil, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given ChatSettingsDto and assigns it to the ChatSettings field.
func (o *FolderDtoInteger) SetChatSettings(v ChatSettingsDto) {
	o.ChatSettings = &v
}

// GetRootRoomType returns the RootRoomType field value if set, zero value otherwise.
func (o *FolderDtoInteger) GetRootRoomType() RoomType {
	if o == nil || IsNil(o.RootRoomType) {
		var ret RoomType
		return ret
	}
	return *o.RootRoomType
}

// GetRootRoomTypeOk returns a tuple with the RootRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoInteger) GetRootRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RootRoomType) {
		return nil, false
	}
	return o.RootRoomType, true
}

// HasRootRoomType returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsRootRoomTypeSet() bool {
	if o != nil && !IsNil(o.RootRoomType) {
		return true
	}

	return false
}

// SetRootRoomType gets a reference to the given RoomType and assigns it to the RootRoomType field.
func (o *FolderDtoInteger) SetRootRoomType(v RoomType) {
	o.RootRoomType = &v
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX.Get()
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveFormAsXLSX.Get(), o.SaveFormAsXLSX.IsSet()
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSaveFormAsXLSXSet() bool {
	if o != nil && o.SaveFormAsXLSX.IsSet() {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given NullableBool and assigns it to the SaveFormAsXLSX field.
func (o *FolderDtoInteger) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX.Set(&v)
}
// SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil
func (o *FolderDtoInteger) SetSaveFormAsXLSXNil() {
	o.SaveFormAsXLSX.Set(nil)
}

// UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
func (o *FolderDtoInteger) UnsetSaveFormAsXLSX() {
	o.SaveFormAsXLSX.Unset()
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB.Get()) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB.Get()
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendFormToExternalDB.Get(), o.SendFormToExternalDB.IsSet()
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsSendFormToExternalDBSet() bool {
	if o != nil && o.SendFormToExternalDB.IsSet() {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given NullableBool and assigns it to the SendFormToExternalDB field.
func (o *FolderDtoInteger) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB.Set(&v)
}
// SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil
func (o *FolderDtoInteger) SetSendFormToExternalDBNil() {
	o.SendFormToExternalDB.Set(nil)
}

// UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
func (o *FolderDtoInteger) UnsetSendFormToExternalDB() {
	o.SendFormToExternalDB.Unset()
}

// GetOriginalFormId returns the OriginalFormId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoInteger) GetOriginalFormId() int32 {
	if o == nil || IsNil(o.OriginalFormId.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalFormId.Get()
}

// GetOriginalFormIdOk returns a tuple with the OriginalFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoInteger) GetOriginalFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFormId.Get(), o.OriginalFormId.IsSet()
}

// HasOriginalFormId returns a boolean if a field has been set.
func (o *FolderDtoInteger) IsOriginalFormIdSet() bool {
	if o != nil && o.OriginalFormId.IsSet() {
		return true
	}

	return false
}

// SetOriginalFormId gets a reference to the given NullableInt32 and assigns it to the OriginalFormId field.
func (o *FolderDtoInteger) SetOriginalFormId(v int32) {
	o.OriginalFormId.Set(&v)
}
// SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil
func (o *FolderDtoInteger) SetOriginalFormIdNil() {
	o.OriginalFormId.Set(nil)
}

// UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil
func (o *FolderDtoInteger) UnsetOriginalFormId() {
	o.OriginalFormId.Unset()
}

func (o FolderDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderDtoInteger) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
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

type NullableFolderDtoInteger struct {
	value *FolderDtoInteger
	isSet bool
}

func (v NullableFolderDtoInteger) Get() *FolderDtoInteger {
	return v.value
}

func (v *NullableFolderDtoInteger) Set(val *FolderDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderDtoInteger(val *FolderDtoInteger) *NullableFolderDtoInteger {
	return &NullableFolderDtoInteger{value: val, isSet: true}
}

func (v NullableFolderDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

