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

// checks if the FolderDtoString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderDtoString{}

// FolderDtoString The folder parameters.
type FolderDtoString struct {
	// The file entry title.
	Title NullableString `json:"title,omitempty"`
	Access *FileShare `json:"access,omitempty"`
	SharedBy *EmployeeDto `json:"sharedBy,omitempty"`
	OwnedBy *EmployeeDto `json:"ownedBy,omitempty"`
	// Specifies if the file entry is shared via link or not.
	Shared *bool `json:"shared,omitempty"`
	// Specifies if the file entry is shared for user or not.
	SharedForUser *bool `json:"sharedForUser,omitempty"`
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
	Id NullableString `json:"id,omitempty"`
	// The root folder ID of the file entry.
	RootFolderId NullableString `json:"rootFolderId,omitempty"`
	// The origin ID of the file entry.
	OriginId NullableString `json:"originId,omitempty"`
	// The origin room ID of the file entry.
	OriginRoomId NullableString `json:"originRoomId,omitempty"`
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
	ParentId NullableString `json:"parentId,omitempty"`
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
}

// NewFolderDtoString instantiates a new FolderDtoString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderDtoString() *FolderDtoString {
	this := FolderDtoString{}
	return &this
}

// NewFolderDtoStringWithDefaults instantiates a new FolderDtoString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderDtoStringWithDefaults() *FolderDtoString {
	this := FolderDtoString{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FolderDtoString) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FolderDtoString) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FolderDtoString) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FolderDtoString) UnsetTitle() {
	o.Title.Unset()
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FolderDtoString) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FolderDtoString) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FolderDtoString) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *FolderDtoString) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *FolderDtoString) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *FolderDtoString) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *FolderDtoString) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *FolderDtoString) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *FolderDtoString) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *FolderDtoString) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *FolderDtoString) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *FolderDtoString) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *FolderDtoString) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *FolderDtoString) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *FolderDtoString) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *FolderDtoString) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *FolderDtoString) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *FolderDtoString) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl.Get()
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetShortWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortWebUrl.Get(), o.ShortWebUrl.IsSet()
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *FolderDtoString) IsShortWebUrlSet() bool {
	if o != nil && o.ShortWebUrl.IsSet() {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given NullableString and assigns it to the ShortWebUrl field.
func (o *FolderDtoString) SetShortWebUrl(v string) {
	o.ShortWebUrl.Set(&v)
}
// SetShortWebUrlNil sets the value for ShortWebUrl to be an explicit nil
func (o *FolderDtoString) SetShortWebUrlNil() {
	o.ShortWebUrl.Set(nil)
}

// UnsetShortWebUrl ensures that no value is present for ShortWebUrl, not even an explicit nil
func (o *FolderDtoString) UnsetShortWebUrl() {
	o.ShortWebUrl.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FolderDtoString) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FolderDtoString) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *FolderDtoString) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FolderDtoString) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FolderDtoString) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *FolderDtoString) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *FolderDtoString) GetUpdated() ApiDateTime {
	if o == nil || IsNil(o.Updated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetUpdatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *FolderDtoString) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given ApiDateTime and assigns it to the Updated field.
func (o *FolderDtoString) SetUpdated(v ApiDateTime) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *FolderDtoString) GetAutoDelete() ApiDateTime {
	if o == nil || IsNil(o.AutoDelete) {
		var ret ApiDateTime
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetAutoDeleteOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *FolderDtoString) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given ApiDateTime and assigns it to the AutoDelete field.
func (o *FolderDtoString) SetAutoDelete(v ApiDateTime) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *FolderDtoString) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *FolderDtoString) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *FolderDtoString) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *FolderDtoString) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *FolderDtoString) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *FolderDtoString) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FolderDtoString) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FolderDtoString) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *FolderDtoString) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem.Get()) {
		var ret bool
		return ret
	}
	return *o.ProviderItem.Get()
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetProviderItemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderItem.Get(), o.ProviderItem.IsSet()
}

// HasProviderItem returns a boolean if a field has been set.
func (o *FolderDtoString) IsProviderItemSet() bool {
	if o != nil && o.ProviderItem.IsSet() {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given NullableBool and assigns it to the ProviderItem field.
func (o *FolderDtoString) SetProviderItem(v bool) {
	o.ProviderItem.Set(&v)
}
// SetProviderItemNil sets the value for ProviderItem to be an explicit nil
func (o *FolderDtoString) SetProviderItemNil() {
	o.ProviderItem.Set(nil)
}

// UnsetProviderItem ensures that no value is present for ProviderItem, not even an explicit nil
func (o *FolderDtoString) UnsetProviderItem() {
	o.ProviderItem.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *FolderDtoString) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *FolderDtoString) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *FolderDtoString) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *FolderDtoString) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *FolderDtoString) IsProviderIdSet() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableInt32 and assigns it to the ProviderId field.
func (o *FolderDtoString) SetProviderId(v int32) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *FolderDtoString) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *FolderDtoString) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetOrder() string {
	if o == nil || IsNil(o.Order.Get()) {
		var ret string
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *FolderDtoString) IsOrderSet() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableString and assigns it to the Order field.
func (o *FolderDtoString) SetOrder(v string) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *FolderDtoString) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *FolderDtoString) UnsetOrder() {
	o.Order.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *FolderDtoString) IsIsFavoriteSet() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *FolderDtoString) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *FolderDtoString) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *FolderDtoString) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *FolderDtoString) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *FolderDtoString) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *FolderDtoString) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FolderDtoString) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *FolderDtoString) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *FolderDtoString) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FolderDtoString) UnsetId() {
	o.Id.Unset()
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetRootFolderId() string {
	if o == nil || IsNil(o.RootFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderId.Get()
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetRootFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderId.Get(), o.RootFolderId.IsSet()
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *FolderDtoString) IsRootFolderIdSet() bool {
	if o != nil && o.RootFolderId.IsSet() {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given NullableString and assigns it to the RootFolderId field.
func (o *FolderDtoString) SetRootFolderId(v string) {
	o.RootFolderId.Set(&v)
}
// SetRootFolderIdNil sets the value for RootFolderId to be an explicit nil
func (o *FolderDtoString) SetRootFolderIdNil() {
	o.RootFolderId.Set(nil)
}

// UnsetRootFolderId ensures that no value is present for RootFolderId, not even an explicit nil
func (o *FolderDtoString) UnsetRootFolderId() {
	o.RootFolderId.Unset()
}

// GetOriginId returns the OriginId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetOriginId() string {
	if o == nil || IsNil(o.OriginId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginId.Get()
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetOriginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginId.Get(), o.OriginId.IsSet()
}

// HasOriginId returns a boolean if a field has been set.
func (o *FolderDtoString) IsOriginIdSet() bool {
	if o != nil && o.OriginId.IsSet() {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given NullableString and assigns it to the OriginId field.
func (o *FolderDtoString) SetOriginId(v string) {
	o.OriginId.Set(&v)
}
// SetOriginIdNil sets the value for OriginId to be an explicit nil
func (o *FolderDtoString) SetOriginIdNil() {
	o.OriginId.Set(nil)
}

// UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
func (o *FolderDtoString) UnsetOriginId() {
	o.OriginId.Unset()
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetOriginRoomId() string {
	if o == nil || IsNil(o.OriginRoomId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomId.Get()
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetOriginRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomId.Get(), o.OriginRoomId.IsSet()
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *FolderDtoString) IsOriginRoomIdSet() bool {
	if o != nil && o.OriginRoomId.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given NullableString and assigns it to the OriginRoomId field.
func (o *FolderDtoString) SetOriginRoomId(v string) {
	o.OriginRoomId.Set(&v)
}
// SetOriginRoomIdNil sets the value for OriginRoomId to be an explicit nil
func (o *FolderDtoString) SetOriginRoomIdNil() {
	o.OriginRoomId.Set(nil)
}

// UnsetOriginRoomId ensures that no value is present for OriginRoomId, not even an explicit nil
func (o *FolderDtoString) UnsetOriginRoomId() {
	o.OriginRoomId.Unset()
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginTitle.Get()
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetOriginTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginTitle.Get(), o.OriginTitle.IsSet()
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *FolderDtoString) IsOriginTitleSet() bool {
	if o != nil && o.OriginTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given NullableString and assigns it to the OriginTitle field.
func (o *FolderDtoString) SetOriginTitle(v string) {
	o.OriginTitle.Set(&v)
}
// SetOriginTitleNil sets the value for OriginTitle to be an explicit nil
func (o *FolderDtoString) SetOriginTitleNil() {
	o.OriginTitle.Set(nil)
}

// UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
func (o *FolderDtoString) UnsetOriginTitle() {
	o.OriginTitle.Unset()
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle.Get()
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomTitle.Get(), o.OriginRoomTitle.IsSet()
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *FolderDtoString) IsOriginRoomTitleSet() bool {
	if o != nil && o.OriginRoomTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given NullableString and assigns it to the OriginRoomTitle field.
func (o *FolderDtoString) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle.Set(&v)
}
// SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil
func (o *FolderDtoString) SetOriginRoomTitleNil() {
	o.OriginRoomTitle.Set(nil)
}

// UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
func (o *FolderDtoString) UnsetOriginRoomTitle() {
	o.OriginRoomTitle.Unset()
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *FolderDtoString) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *FolderDtoString) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *FolderDtoString) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret FileEntryDtoIntegerAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *FolderDtoString) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableFileEntryDtoIntegerAllOfShareSettings and assigns it to the ShareSettings field.
func (o *FolderDtoString) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *FolderDtoString) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *FolderDtoString) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetSecurity() FileEntryDtoIntegerAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret FileEntryDtoIntegerAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *FolderDtoString) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableFileEntryDtoIntegerAllOfSecurity and assigns it to the Security field.
func (o *FolderDtoString) SetSecurity(v FileEntryDtoIntegerAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *FolderDtoString) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *FolderDtoString) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret FileEntryDtoIntegerAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *FolderDtoString) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableFileEntryDtoIntegerAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *FolderDtoString) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *FolderDtoString) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *FolderDtoString) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *FolderDtoString) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *FolderDtoString) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *FolderDtoString) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *FolderDtoString) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetExternal() bool {
	if o == nil || IsNil(o.External.Get()) {
		var ret bool
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *FolderDtoString) IsExternalSet() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableBool and assigns it to the External field.
func (o *FolderDtoString) SetExternal(v bool) {
	o.External.Set(&v)
}
// SetExternalNil sets the value for External to be an explicit nil
func (o *FolderDtoString) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *FolderDtoString) UnsetExternal() {
	o.External.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FolderDtoString) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FolderDtoString) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *FolderDtoString) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired.Get()
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLinkExpired.Get(), o.IsLinkExpired.IsSet()
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *FolderDtoString) IsIsLinkExpiredSet() bool {
	if o != nil && o.IsLinkExpired.IsSet() {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given NullableBool and assigns it to the IsLinkExpired field.
func (o *FolderDtoString) SetIsLinkExpired(v bool) {
	o.IsLinkExpired.Set(&v)
}
// SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil
func (o *FolderDtoString) SetIsLinkExpiredNil() {
	o.IsLinkExpired.Set(nil)
}

// UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
func (o *FolderDtoString) UnsetIsLinkExpired() {
	o.IsLinkExpired.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *FolderDtoString) IsParentIdSet() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *FolderDtoString) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *FolderDtoString) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *FolderDtoString) UnsetParentId() {
	o.ParentId.Unset()
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *FolderDtoString) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *FolderDtoString) IsFilesCountSet() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *FolderDtoString) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetFoldersCount returns the FoldersCount field value if set, zero value otherwise.
func (o *FolderDtoString) GetFoldersCount() int32 {
	if o == nil || IsNil(o.FoldersCount) {
		var ret int32
		return ret
	}
	return *o.FoldersCount
}

// GetFoldersCountOk returns a tuple with the FoldersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetFoldersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FoldersCount) {
		return nil, false
	}
	return o.FoldersCount, true
}

// HasFoldersCount returns a boolean if a field has been set.
func (o *FolderDtoString) IsFoldersCountSet() bool {
	if o != nil && !IsNil(o.FoldersCount) {
		return true
	}

	return false
}

// SetFoldersCount gets a reference to the given int32 and assigns it to the FoldersCount field.
func (o *FolderDtoString) SetFoldersCount(v int32) {
	o.FoldersCount = &v
}

// GetIsShareable returns the IsShareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetIsShareable() bool {
	if o == nil || IsNil(o.IsShareable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsShareable.Get()
}

// GetIsShareableOk returns a tuple with the IsShareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetIsShareableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsShareable.Get(), o.IsShareable.IsSet()
}

// HasIsShareable returns a boolean if a field has been set.
func (o *FolderDtoString) IsIsShareableSet() bool {
	if o != nil && o.IsShareable.IsSet() {
		return true
	}

	return false
}

// SetIsShareable gets a reference to the given NullableBool and assigns it to the IsShareable field.
func (o *FolderDtoString) SetIsShareable(v bool) {
	o.IsShareable.Set(&v)
}
// SetIsShareableNil sets the value for IsShareable to be an explicit nil
func (o *FolderDtoString) SetIsShareableNil() {
	o.IsShareable.Set(nil)
}

// UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
func (o *FolderDtoString) UnsetIsShareable() {
	o.IsShareable.Unset()
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *FolderDtoString) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *FolderDtoString) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *FolderDtoString) SetNew(v int32) {
	o.New = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *FolderDtoString) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *FolderDtoString) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *FolderDtoString) SetMute(v bool) {
	o.Mute = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FolderDtoString) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FolderDtoString) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *FolderDtoString) GetLogo() Logo {
	if o == nil || IsNil(o.Logo) {
		var ret Logo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetLogoOk() (*Logo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *FolderDtoString) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given Logo and assigns it to the Logo field.
func (o *FolderDtoString) SetLogo(v Logo) {
	o.Logo = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *FolderDtoString) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *FolderDtoString) IsPinnedSet() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *FolderDtoString) SetPinned(v bool) {
	o.Pinned = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *FolderDtoString) GetRoomType() RoomType {
	if o == nil || IsNil(o.RoomType) {
		var ret RoomType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *FolderDtoString) IsRoomTypeSet() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given RoomType and assigns it to the RoomType field.
func (o *FolderDtoString) SetRoomType(v RoomType) {
	o.RoomType = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *FolderDtoString) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *FolderDtoString) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *FolderDtoString) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *FolderDtoString) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *FolderDtoString) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *FolderDtoString) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *FolderDtoString) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *FolderDtoString) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *FolderDtoString) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *FolderDtoString) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *FolderDtoString) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *FolderDtoString) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *FolderDtoString) GetWatermark() WatermarkDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetWatermarkOk() (*WatermarkDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *FolderDtoString) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkDto and assigns it to the Watermark field.
func (o *FolderDtoString) SetWatermark(v WatermarkDto) {
	o.Watermark = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FolderDtoString) GetType() FolderType {
	if o == nil || IsNil(o.Type) {
		var ret FolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FolderDtoString) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FolderType and assigns it to the Type field.
func (o *FolderDtoString) SetType(v FolderType) {
	o.Type = &v
}

// GetInRoom returns the InRoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetInRoom() bool {
	if o == nil || IsNil(o.InRoom.Get()) {
		var ret bool
		return ret
	}
	return *o.InRoom.Get()
}

// GetInRoomOk returns a tuple with the InRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetInRoomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InRoom.Get(), o.InRoom.IsSet()
}

// HasInRoom returns a boolean if a field has been set.
func (o *FolderDtoString) IsInRoomSet() bool {
	if o != nil && o.InRoom.IsSet() {
		return true
	}

	return false
}

// SetInRoom gets a reference to the given NullableBool and assigns it to the InRoom field.
func (o *FolderDtoString) SetInRoom(v bool) {
	o.InRoom.Set(&v)
}
// SetInRoomNil sets the value for InRoom to be an explicit nil
func (o *FolderDtoString) SetInRoomNil() {
	o.InRoom.Set(nil)
}

// UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
func (o *FolderDtoString) UnsetInRoom() {
	o.InRoom.Unset()
}

// GetQuotaLimit returns the QuotaLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetQuotaLimit() int64 {
	if o == nil || IsNil(o.QuotaLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.QuotaLimit.Get()
}

// GetQuotaLimitOk returns a tuple with the QuotaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetQuotaLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaLimit.Get(), o.QuotaLimit.IsSet()
}

// HasQuotaLimit returns a boolean if a field has been set.
func (o *FolderDtoString) IsQuotaLimitSet() bool {
	if o != nil && o.QuotaLimit.IsSet() {
		return true
	}

	return false
}

// SetQuotaLimit gets a reference to the given NullableInt64 and assigns it to the QuotaLimit field.
func (o *FolderDtoString) SetQuotaLimit(v int64) {
	o.QuotaLimit.Set(&v)
}
// SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil
func (o *FolderDtoString) SetQuotaLimitNil() {
	o.QuotaLimit.Set(nil)
}

// UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
func (o *FolderDtoString) UnsetQuotaLimit() {
	o.QuotaLimit.Unset()
}

// GetIsCustomQuota returns the IsCustomQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetIsCustomQuota() bool {
	if o == nil || IsNil(o.IsCustomQuota.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCustomQuota.Get()
}

// GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetIsCustomQuotaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCustomQuota.Get(), o.IsCustomQuota.IsSet()
}

// HasIsCustomQuota returns a boolean if a field has been set.
func (o *FolderDtoString) IsIsCustomQuotaSet() bool {
	if o != nil && o.IsCustomQuota.IsSet() {
		return true
	}

	return false
}

// SetIsCustomQuota gets a reference to the given NullableBool and assigns it to the IsCustomQuota field.
func (o *FolderDtoString) SetIsCustomQuota(v bool) {
	o.IsCustomQuota.Set(&v)
}
// SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil
func (o *FolderDtoString) SetIsCustomQuotaNil() {
	o.IsCustomQuota.Set(nil)
}

// UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
func (o *FolderDtoString) UnsetIsCustomQuota() {
	o.IsCustomQuota.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.UsedSpace.Get()
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetUsedSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedSpace.Get(), o.UsedSpace.IsSet()
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *FolderDtoString) IsUsedSpaceSet() bool {
	if o != nil && o.UsedSpace.IsSet() {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given NullableInt64 and assigns it to the UsedSpace field.
func (o *FolderDtoString) SetUsedSpace(v int64) {
	o.UsedSpace.Set(&v)
}
// SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil
func (o *FolderDtoString) SetUsedSpaceNil() {
	o.UsedSpace.Set(nil)
}

// UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
func (o *FolderDtoString) UnsetUsedSpace() {
	o.UsedSpace.Unset()
}

// GetPasswordProtected returns the PasswordProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetPasswordProtected() bool {
	if o == nil || IsNil(o.PasswordProtected.Get()) {
		var ret bool
		return ret
	}
	return *o.PasswordProtected.Get()
}

// GetPasswordProtectedOk returns a tuple with the PasswordProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetPasswordProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProtected.Get(), o.PasswordProtected.IsSet()
}

// HasPasswordProtected returns a boolean if a field has been set.
func (o *FolderDtoString) IsPasswordProtectedSet() bool {
	if o != nil && o.PasswordProtected.IsSet() {
		return true
	}

	return false
}

// SetPasswordProtected gets a reference to the given NullableBool and assigns it to the PasswordProtected field.
func (o *FolderDtoString) SetPasswordProtected(v bool) {
	o.PasswordProtected.Set(&v)
}
// SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil
func (o *FolderDtoString) SetPasswordProtectedNil() {
	o.PasswordProtected.Set(nil)
}

// UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
func (o *FolderDtoString) UnsetPasswordProtected() {
	o.PasswordProtected.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *FolderDtoString) GetExpired() bool {
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
func (o *FolderDtoString) GetExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expired.Get(), o.Expired.IsSet()
}

// HasExpired returns a boolean if a field has been set.
func (o *FolderDtoString) IsExpiredSet() bool {
	if o != nil && o.Expired.IsSet() {
		return true
	}

	return false
}

// SetExpired gets a reference to the given NullableBool and assigns it to the Expired field.
// Deprecated
func (o *FolderDtoString) SetExpired(v bool) {
	o.Expired.Set(&v)
}
// SetExpiredNil sets the value for Expired to be an explicit nil
func (o *FolderDtoString) SetExpiredNil() {
	o.Expired.Set(nil)
}

// UnsetExpired ensures that no value is present for Expired, not even an explicit nil
func (o *FolderDtoString) UnsetExpired() {
	o.Expired.Unset()
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *FolderDtoString) GetChatSettings() ChatSettingsDto {
	if o == nil || IsNil(o.ChatSettings) {
		var ret ChatSettingsDto
		return ret
	}
	return *o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetChatSettingsOk() (*ChatSettingsDto, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return nil, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *FolderDtoString) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given ChatSettingsDto and assigns it to the ChatSettings field.
func (o *FolderDtoString) SetChatSettings(v ChatSettingsDto) {
	o.ChatSettings = &v
}

// GetRootRoomType returns the RootRoomType field value if set, zero value otherwise.
func (o *FolderDtoString) GetRootRoomType() RoomType {
	if o == nil || IsNil(o.RootRoomType) {
		var ret RoomType
		return ret
	}
	return *o.RootRoomType
}

// GetRootRoomTypeOk returns a tuple with the RootRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderDtoString) GetRootRoomTypeOk() (*RoomType, bool) {
	if o == nil || IsNil(o.RootRoomType) {
		return nil, false
	}
	return o.RootRoomType, true
}

// HasRootRoomType returns a boolean if a field has been set.
func (o *FolderDtoString) IsRootRoomTypeSet() bool {
	if o != nil && !IsNil(o.RootRoomType) {
		return true
	}

	return false
}

// SetRootRoomType gets a reference to the given RoomType and assigns it to the RootRoomType field.
func (o *FolderDtoString) SetRootRoomType(v RoomType) {
	o.RootRoomType = &v
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX.Get()
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveFormAsXLSX.Get(), o.SaveFormAsXLSX.IsSet()
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *FolderDtoString) IsSaveFormAsXLSXSet() bool {
	if o != nil && o.SaveFormAsXLSX.IsSet() {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given NullableBool and assigns it to the SaveFormAsXLSX field.
func (o *FolderDtoString) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX.Set(&v)
}
// SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil
func (o *FolderDtoString) SetSaveFormAsXLSXNil() {
	o.SaveFormAsXLSX.Set(nil)
}

// UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
func (o *FolderDtoString) UnsetSaveFormAsXLSX() {
	o.SaveFormAsXLSX.Unset()
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderDtoString) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB.Get()) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB.Get()
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderDtoString) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendFormToExternalDB.Get(), o.SendFormToExternalDB.IsSet()
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *FolderDtoString) IsSendFormToExternalDBSet() bool {
	if o != nil && o.SendFormToExternalDB.IsSet() {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given NullableBool and assigns it to the SendFormToExternalDB field.
func (o *FolderDtoString) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB.Set(&v)
}
// SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil
func (o *FolderDtoString) SetSendFormToExternalDBNil() {
	o.SendFormToExternalDB.Set(nil)
}

// UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
func (o *FolderDtoString) UnsetSendFormToExternalDB() {
	o.SendFormToExternalDB.Unset()
}

func (o FolderDtoString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderDtoString) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableFolderDtoString struct {
	value *FolderDtoString
	isSet bool
}

func (v NullableFolderDtoString) Get() *FolderDtoString {
	return v.value
}

func (v *NullableFolderDtoString) Set(val *FolderDtoString) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderDtoString) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderDtoString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderDtoString(val *FolderDtoString) *NullableFolderDtoString {
	return &NullableFolderDtoString{value: val, isSet: true}
}

func (v NullableFolderDtoString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderDtoString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

