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
	"time"
)

// checks if the AiFolderDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFolderDtoInteger{}

// AiFolderDtoInteger The folder parameters.
type AiFolderDtoInteger struct {
	// The file entry title.
	Title *string `json:"title,omitempty"`
	// The access rights to the file entry.
	Access *AiFileShare `json:"access,omitempty"`
	// Provides information about the employee who shared the file or folder.
	SharedBy *AiEmployeeDto `json:"sharedBy,omitempty"`
	// The information about the employee who owns the file entry.
	OwnedBy *AiEmployeeDto `json:"ownedBy,omitempty"`
	// Specifies if the file entry is shared via link or not.
	Shared *bool `json:"shared,omitempty"`
	// Specifies if the file entry is shared for user or not.
	SharedForUser *bool `json:"sharedForUser,omitempty"`
	// Specifies if the file entry is shared via a public (non-internal) external link.
	SharedExternal *bool `json:"sharedExternal,omitempty"`
	// Indicates whether the parent entity is shared.
	ParentShared *bool `json:"parentShared,omitempty"`
	// The short Web URL.
	ShortWebUrl *string `json:"shortWebUrl,omitempty"`
	// The creation date and time of the file entry.
	Created *time.Time `json:"created,omitempty"`
	// The file entry author.
	CreatedBy *AiEmployeeDto `json:"createdBy,omitempty"`
	// The last date and time when the file entry was updated.
	Updated *time.Time `json:"updated,omitempty"`
	// The date and time when the file entry will be automatically deleted.
	AutoDelete *time.Time `json:"autoDelete,omitempty"`
	// The root folder type of the file entry.
	RootFolderType *AiFolderType `json:"rootFolderType,omitempty"`
	// The parent room type of the file entry.
	ParentRoomType *AiFolderType `json:"parentRoomType,omitempty"`
	// The user who updated the file entry.
	UpdatedBy *AiEmployeeDto `json:"updatedBy,omitempty"`
	// Specifies if the file entry provider is specified or not.
	ProviderItem *bool `json:"providerItem,omitempty"`
	// The provider key of the file entry.
	ProviderKey *string `json:"providerKey,omitempty"`
	// The provider ID of the file entry.
	ProviderId *int32 `json:"providerId,omitempty"`
	// The order of the file entry.
	Order *string `json:"order,omitempty"`
	// Specifies if the file is a favorite or not.
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// The file entry type.
	FileEntryType *AiFileEntryType `json:"fileEntryType,omitempty"`
	// The file entry ID.
	Id *int32 `json:"id,omitempty"`
	// The root folder ID of the file entry.
	RootFolderId *int32 `json:"rootFolderId,omitempty"`
	// The origin ID of the file entry.
	OriginId *int32 `json:"originId,omitempty"`
	// The origin room ID of the file entry.
	OriginRoomId *int32 `json:"originRoomId,omitempty"`
	// The origin title of the file entry.
	OriginTitle *string `json:"originTitle,omitempty"`
	// The origin room title of the file entry.
	OriginRoomTitle *string `json:"originRoomTitle,omitempty"`
	// Specifies if the file entry can be shared or not.
	CanShare *bool `json:"canShare,omitempty"`
	ShareSettings NullableFileEntryDtoIntegerAllOfShareSettings `json:"shareSettings,omitempty"`
	Security NullableFileEntryDtoIntegerAllOfSecurity `json:"security,omitempty"`
	AvailableShareRights NullableFileEntryDtoIntegerAllOfAvailableShareRights `json:"availableShareRights,omitempty"`
	// The request token of the file entry.
	RequestToken *string `json:"requestToken,omitempty"`
	// Specifies if the folder can be accessed via an external link or not.
	External *bool `json:"external,omitempty"`
	// Represents the expiration date of the file entry.
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	// Indicates whether the shareable link associated with the file or folder has expired.
	IsLinkExpired *bool `json:"isLinkExpired,omitempty"`
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
	// The folder logo.
	Logo *AiLogo `json:"logo,omitempty"`
	// Specifies if the folder is pinned or not.
	Pinned *bool `json:"pinned,omitempty"`
	// The room type of the folder.
	RoomType *AiRoomType `json:"roomType,omitempty"`
	// Specifies if the folder is private or not.
	Private *bool `json:"private,omitempty"`
	// Specifies if the folder is indexed or not.
	Indexing *bool `json:"indexing,omitempty"`
	// Specifies if the folder can be downloaded or not.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The room data lifetime settings of the folder.
	Lifetime *AiRoomDataLifetimeDto `json:"lifetime,omitempty"`
	// The watermark settings of the folder.
	Watermark *AiWatermarkDto `json:"watermark,omitempty"`
	// The folder type.
	Type *AiFolderType `json:"type,omitempty"`
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
	// The AI chat settings for the folder room. Contains configuration for AI provider, model selection, and custom prompts.  Only applicable to rooms with AI chat functionality enabled. Null if the room does not have chat settings configured.
	ChatSettings *AiChatSettingsDto `json:"chatSettings,omitempty"`
	// The room type of the root folder. Indicates the type of the parent room if the current folder is nested within a room hierarchy.  This property helps identify the context in which a nested folder exists.
	RootRoomType *AiRoomType `json:"rootRoomType,omitempty"`
	// Specifies whether to save form data as XLSX file.
	SaveFormAsXLSX NullableBool `json:"saveFormAsXLSX,omitempty"`
	// Specifies whether to send form data to external database.
	SendFormToExternalDB NullableBool `json:"sendFormToExternalDB,omitempty"`
	// The original form ID that corresponds to this FormFillingFolderDone folder.
	OriginalFormId NullableInt32 `json:"originalFormId,omitempty"`
}

// NewAiFolderDtoInteger instantiates a new AiFolderDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFolderDtoInteger() *AiFolderDtoInteger {
	this := AiFolderDtoInteger{}
	return &this
}

// NewAiFolderDtoIntegerWithDefaults instantiates a new AiFolderDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFolderDtoIntegerWithDefaults() *AiFolderDtoInteger {
	this := AiFolderDtoInteger{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiFolderDtoInteger) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetAccess() AiFileShare {
	if o == nil || IsNil(o.Access) {
		var ret AiFileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetAccessOk() (*AiFileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AiFileShare and assigns it to the Access field.
func (o *AiFolderDtoInteger) SetAccess(v AiFileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetSharedBy() AiEmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetSharedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given AiEmployeeDto and assigns it to the SharedBy field.
func (o *AiFolderDtoInteger) SetSharedBy(v AiEmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOwnedBy() AiEmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOwnedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given AiEmployeeDto and assigns it to the OwnedBy field.
func (o *AiFolderDtoInteger) SetOwnedBy(v AiEmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *AiFolderDtoInteger) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *AiFolderDtoInteger) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *AiFolderDtoInteger) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *AiFolderDtoInteger) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *AiFolderDtoInteger) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AiFolderDtoInteger) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetCreatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetCreatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AiEmployeeDto and assigns it to the CreatedBy field.
func (o *AiFolderDtoInteger) SetCreatedBy(v AiEmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AiFolderDtoInteger) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetAutoDelete() time.Time {
	if o == nil || IsNil(o.AutoDelete) {
		var ret time.Time
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetAutoDeleteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given time.Time and assigns it to the AutoDelete field.
func (o *AiFolderDtoInteger) SetAutoDelete(v time.Time) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetRootFolderType() AiFolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret AiFolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetRootFolderTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given AiFolderType and assigns it to the RootFolderType field.
func (o *AiFolderDtoInteger) SetRootFolderType(v AiFolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetParentRoomType() AiFolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret AiFolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetParentRoomTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given AiFolderType and assigns it to the ParentRoomType field.
func (o *AiFolderDtoInteger) SetParentRoomType(v AiFolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetUpdatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetUpdatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given AiEmployeeDto and assigns it to the UpdatedBy field.
func (o *AiFolderDtoInteger) SetUpdatedBy(v AiEmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *AiFolderDtoInteger) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *AiFolderDtoInteger) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *AiFolderDtoInteger) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *AiFolderDtoInteger) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *AiFolderDtoInteger) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetFileEntryType() AiFileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret AiFileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetFileEntryTypeOk() (*AiFileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given AiFileEntryType and assigns it to the FileEntryType field.
func (o *AiFolderDtoInteger) SetFileEntryType(v AiFileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AiFolderDtoInteger) SetId(v int32) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetRootFolderId() int32 {
	if o == nil || IsNil(o.RootFolderId) {
		var ret int32
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetRootFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given int32 and assigns it to the RootFolderId field.
func (o *AiFolderDtoInteger) SetRootFolderId(v int32) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOriginId() int32 {
	if o == nil || IsNil(o.OriginId) {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOriginIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *AiFolderDtoInteger) SetOriginId(v int32) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOriginRoomId() int32 {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret int32
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOriginRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given int32 and assigns it to the OriginRoomId field.
func (o *AiFolderDtoInteger) SetOriginRoomId(v int32) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle) {
		var ret string
		return ret
	}
	return *o.OriginTitle
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOriginTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginTitle) {
		return nil, false
	}
	return o.OriginTitle, true
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOriginTitleSet() bool {
	if o != nil && !IsNil(o.OriginTitle) {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given string and assigns it to the OriginTitle field.
func (o *AiFolderDtoInteger) SetOriginTitle(v string) {
	o.OriginTitle = &v
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomTitle) {
		return nil, false
	}
	return o.OriginRoomTitle, true
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOriginRoomTitleSet() bool {
	if o != nil && !IsNil(o.OriginRoomTitle) {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given string and assigns it to the OriginRoomTitle field.
func (o *AiFolderDtoInteger) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle = &v
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *AiFolderDtoInteger) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret FileEntryDtoIntegerAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableFileEntryDtoIntegerAllOfShareSettings and assigns it to the ShareSettings field.
func (o *AiFolderDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *AiFolderDtoInteger) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret FileEntryDtoIntegerAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableFileEntryDtoIntegerAllOfSecurity and assigns it to the Security field.
func (o *AiFolderDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *AiFolderDtoInteger) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret FileEntryDtoIntegerAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableFileEntryDtoIntegerAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *AiFolderDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *AiFolderDtoInteger) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsRequestTokenSet() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *AiFolderDtoInteger) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *AiFolderDtoInteger) SetExternal(v bool) {
	o.External = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *AiFolderDtoInteger) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkExpired) {
		return nil, false
	}
	return o.IsLinkExpired, true
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIsLinkExpiredSet() bool {
	if o != nil && !IsNil(o.IsLinkExpired) {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given bool and assigns it to the IsLinkExpired field.
func (o *AiFolderDtoInteger) SetIsLinkExpired(v bool) {
	o.IsLinkExpired = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsParentIdSet() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *AiFolderDtoInteger) SetParentId(v int32) {
	o.ParentId = &v
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsFilesCountSet() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *AiFolderDtoInteger) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetFoldersCount returns the FoldersCount field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetFoldersCount() int32 {
	if o == nil || IsNil(o.FoldersCount) {
		var ret int32
		return ret
	}
	return *o.FoldersCount
}

// GetFoldersCountOk returns a tuple with the FoldersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetFoldersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FoldersCount) {
		return nil, false
	}
	return o.FoldersCount, true
}

// HasFoldersCount returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsFoldersCountSet() bool {
	if o != nil && !IsNil(o.FoldersCount) {
		return true
	}

	return false
}

// SetFoldersCount gets a reference to the given int32 and assigns it to the FoldersCount field.
func (o *AiFolderDtoInteger) SetFoldersCount(v int32) {
	o.FoldersCount = &v
}

// GetIsShareable returns the IsShareable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetIsShareable() bool {
	if o == nil || IsNil(o.IsShareable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsShareable.Get()
}

// GetIsShareableOk returns a tuple with the IsShareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetIsShareableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsShareable.Get(), o.IsShareable.IsSet()
}

// HasIsShareable returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIsShareableSet() bool {
	if o != nil && o.IsShareable.IsSet() {
		return true
	}

	return false
}

// SetIsShareable gets a reference to the given NullableBool and assigns it to the IsShareable field.
func (o *AiFolderDtoInteger) SetIsShareable(v bool) {
	o.IsShareable.Set(&v)
}
// SetIsShareableNil sets the value for IsShareable to be an explicit nil
func (o *AiFolderDtoInteger) SetIsShareableNil() {
	o.IsShareable.Set(nil)
}

// UnsetIsShareable ensures that no value is present for IsShareable, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetIsShareable() {
	o.IsShareable.Unset()
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *AiFolderDtoInteger) SetNew(v int32) {
	o.New = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *AiFolderDtoInteger) SetMute(v bool) {
	o.Mute = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AiFolderDtoInteger) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetLogo() AiLogo {
	if o == nil || IsNil(o.Logo) {
		var ret AiLogo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetLogoOk() (*AiLogo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given AiLogo and assigns it to the Logo field.
func (o *AiFolderDtoInteger) SetLogo(v AiLogo) {
	o.Logo = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetPinned() bool {
	if o == nil || IsNil(o.Pinned) {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetPinnedOk() (*bool, bool) {
	if o == nil || IsNil(o.Pinned) {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsPinnedSet() bool {
	if o != nil && !IsNil(o.Pinned) {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *AiFolderDtoInteger) SetPinned(v bool) {
	o.Pinned = &v
}

// GetRoomType returns the RoomType field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetRoomType() AiRoomType {
	if o == nil || IsNil(o.RoomType) {
		var ret AiRoomType
		return ret
	}
	return *o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetRoomTypeOk() (*AiRoomType, bool) {
	if o == nil || IsNil(o.RoomType) {
		return nil, false
	}
	return o.RoomType, true
}

// HasRoomType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsRoomTypeSet() bool {
	if o != nil && !IsNil(o.RoomType) {
		return true
	}

	return false
}

// SetRoomType gets a reference to the given AiRoomType and assigns it to the RoomType field.
func (o *AiFolderDtoInteger) SetRoomType(v AiRoomType) {
	o.RoomType = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *AiFolderDtoInteger) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *AiFolderDtoInteger) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *AiFolderDtoInteger) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetLifetime() AiRoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret AiRoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetLifetimeOk() (*AiRoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given AiRoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *AiFolderDtoInteger) SetLifetime(v AiRoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetWatermark() AiWatermarkDto {
	if o == nil || IsNil(o.Watermark) {
		var ret AiWatermarkDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetWatermarkOk() (*AiWatermarkDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given AiWatermarkDto and assigns it to the Watermark field.
func (o *AiFolderDtoInteger) SetWatermark(v AiWatermarkDto) {
	o.Watermark = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetType() AiFolderType {
	if o == nil || IsNil(o.Type) {
		var ret AiFolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AiFolderType and assigns it to the Type field.
func (o *AiFolderDtoInteger) SetType(v AiFolderType) {
	o.Type = &v
}

// GetInRoom returns the InRoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetInRoom() bool {
	if o == nil || IsNil(o.InRoom.Get()) {
		var ret bool
		return ret
	}
	return *o.InRoom.Get()
}

// GetInRoomOk returns a tuple with the InRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetInRoomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InRoom.Get(), o.InRoom.IsSet()
}

// HasInRoom returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsInRoomSet() bool {
	if o != nil && o.InRoom.IsSet() {
		return true
	}

	return false
}

// SetInRoom gets a reference to the given NullableBool and assigns it to the InRoom field.
func (o *AiFolderDtoInteger) SetInRoom(v bool) {
	o.InRoom.Set(&v)
}
// SetInRoomNil sets the value for InRoom to be an explicit nil
func (o *AiFolderDtoInteger) SetInRoomNil() {
	o.InRoom.Set(nil)
}

// UnsetInRoom ensures that no value is present for InRoom, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetInRoom() {
	o.InRoom.Unset()
}

// GetQuotaLimit returns the QuotaLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetQuotaLimit() int64 {
	if o == nil || IsNil(o.QuotaLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.QuotaLimit.Get()
}

// GetQuotaLimitOk returns a tuple with the QuotaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetQuotaLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaLimit.Get(), o.QuotaLimit.IsSet()
}

// HasQuotaLimit returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsQuotaLimitSet() bool {
	if o != nil && o.QuotaLimit.IsSet() {
		return true
	}

	return false
}

// SetQuotaLimit gets a reference to the given NullableInt64 and assigns it to the QuotaLimit field.
func (o *AiFolderDtoInteger) SetQuotaLimit(v int64) {
	o.QuotaLimit.Set(&v)
}
// SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil
func (o *AiFolderDtoInteger) SetQuotaLimitNil() {
	o.QuotaLimit.Set(nil)
}

// UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetQuotaLimit() {
	o.QuotaLimit.Unset()
}

// GetIsCustomQuota returns the IsCustomQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetIsCustomQuota() bool {
	if o == nil || IsNil(o.IsCustomQuota.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCustomQuota.Get()
}

// GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetIsCustomQuotaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCustomQuota.Get(), o.IsCustomQuota.IsSet()
}

// HasIsCustomQuota returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsIsCustomQuotaSet() bool {
	if o != nil && o.IsCustomQuota.IsSet() {
		return true
	}

	return false
}

// SetIsCustomQuota gets a reference to the given NullableBool and assigns it to the IsCustomQuota field.
func (o *AiFolderDtoInteger) SetIsCustomQuota(v bool) {
	o.IsCustomQuota.Set(&v)
}
// SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil
func (o *AiFolderDtoInteger) SetIsCustomQuotaNil() {
	o.IsCustomQuota.Set(nil)
}

// UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetIsCustomQuota() {
	o.IsCustomQuota.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.UsedSpace.Get()
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetUsedSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedSpace.Get(), o.UsedSpace.IsSet()
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsUsedSpaceSet() bool {
	if o != nil && o.UsedSpace.IsSet() {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given NullableInt64 and assigns it to the UsedSpace field.
func (o *AiFolderDtoInteger) SetUsedSpace(v int64) {
	o.UsedSpace.Set(&v)
}
// SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil
func (o *AiFolderDtoInteger) SetUsedSpaceNil() {
	o.UsedSpace.Set(nil)
}

// UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetUsedSpace() {
	o.UsedSpace.Unset()
}

// GetPasswordProtected returns the PasswordProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetPasswordProtected() bool {
	if o == nil || IsNil(o.PasswordProtected.Get()) {
		var ret bool
		return ret
	}
	return *o.PasswordProtected.Get()
}

// GetPasswordProtectedOk returns a tuple with the PasswordProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetPasswordProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProtected.Get(), o.PasswordProtected.IsSet()
}

// HasPasswordProtected returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsPasswordProtectedSet() bool {
	if o != nil && o.PasswordProtected.IsSet() {
		return true
	}

	return false
}

// SetPasswordProtected gets a reference to the given NullableBool and assigns it to the PasswordProtected field.
func (o *AiFolderDtoInteger) SetPasswordProtected(v bool) {
	o.PasswordProtected.Set(&v)
}
// SetPasswordProtectedNil sets the value for PasswordProtected to be an explicit nil
func (o *AiFolderDtoInteger) SetPasswordProtectedNil() {
	o.PasswordProtected.Set(nil)
}

// UnsetPasswordProtected ensures that no value is present for PasswordProtected, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetPasswordProtected() {
	o.PasswordProtected.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *AiFolderDtoInteger) GetExpired() bool {
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
func (o *AiFolderDtoInteger) GetExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expired.Get(), o.Expired.IsSet()
}

// HasExpired returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsExpiredSet() bool {
	if o != nil && o.Expired.IsSet() {
		return true
	}

	return false
}

// SetExpired gets a reference to the given NullableBool and assigns it to the Expired field.
// Deprecated
func (o *AiFolderDtoInteger) SetExpired(v bool) {
	o.Expired.Set(&v)
}
// SetExpiredNil sets the value for Expired to be an explicit nil
func (o *AiFolderDtoInteger) SetExpiredNil() {
	o.Expired.Set(nil)
}

// UnsetExpired ensures that no value is present for Expired, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetExpired() {
	o.Expired.Unset()
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetChatSettings() AiChatSettingsDto {
	if o == nil || IsNil(o.ChatSettings) {
		var ret AiChatSettingsDto
		return ret
	}
	return *o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetChatSettingsOk() (*AiChatSettingsDto, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return nil, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given AiChatSettingsDto and assigns it to the ChatSettings field.
func (o *AiFolderDtoInteger) SetChatSettings(v AiChatSettingsDto) {
	o.ChatSettings = &v
}

// GetRootRoomType returns the RootRoomType field value if set, zero value otherwise.
func (o *AiFolderDtoInteger) GetRootRoomType() AiRoomType {
	if o == nil || IsNil(o.RootRoomType) {
		var ret AiRoomType
		return ret
	}
	return *o.RootRoomType
}

// GetRootRoomTypeOk returns a tuple with the RootRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderDtoInteger) GetRootRoomTypeOk() (*AiRoomType, bool) {
	if o == nil || IsNil(o.RootRoomType) {
		return nil, false
	}
	return o.RootRoomType, true
}

// HasRootRoomType returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsRootRoomTypeSet() bool {
	if o != nil && !IsNil(o.RootRoomType) {
		return true
	}

	return false
}

// SetRootRoomType gets a reference to the given AiRoomType and assigns it to the RootRoomType field.
func (o *AiFolderDtoInteger) SetRootRoomType(v AiRoomType) {
	o.RootRoomType = &v
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX.Get()
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveFormAsXLSX.Get(), o.SaveFormAsXLSX.IsSet()
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSaveFormAsXLSXSet() bool {
	if o != nil && o.SaveFormAsXLSX.IsSet() {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given NullableBool and assigns it to the SaveFormAsXLSX field.
func (o *AiFolderDtoInteger) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX.Set(&v)
}
// SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil
func (o *AiFolderDtoInteger) SetSaveFormAsXLSXNil() {
	o.SaveFormAsXLSX.Set(nil)
}

// UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetSaveFormAsXLSX() {
	o.SaveFormAsXLSX.Unset()
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB.Get()) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB.Get()
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendFormToExternalDB.Get(), o.SendFormToExternalDB.IsSet()
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsSendFormToExternalDBSet() bool {
	if o != nil && o.SendFormToExternalDB.IsSet() {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given NullableBool and assigns it to the SendFormToExternalDB field.
func (o *AiFolderDtoInteger) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB.Set(&v)
}
// SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil
func (o *AiFolderDtoInteger) SetSendFormToExternalDBNil() {
	o.SendFormToExternalDB.Set(nil)
}

// UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetSendFormToExternalDB() {
	o.SendFormToExternalDB.Unset()
}

// GetOriginalFormId returns the OriginalFormId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderDtoInteger) GetOriginalFormId() int32 {
	if o == nil || IsNil(o.OriginalFormId.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalFormId.Get()
}

// GetOriginalFormIdOk returns a tuple with the OriginalFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderDtoInteger) GetOriginalFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalFormId.Get(), o.OriginalFormId.IsSet()
}

// HasOriginalFormId returns a boolean if a field has been set.
func (o *AiFolderDtoInteger) IsOriginalFormIdSet() bool {
	if o != nil && o.OriginalFormId.IsSet() {
		return true
	}

	return false
}

// SetOriginalFormId gets a reference to the given NullableInt32 and assigns it to the OriginalFormId field.
func (o *AiFolderDtoInteger) SetOriginalFormId(v int32) {
	o.OriginalFormId.Set(&v)
}
// SetOriginalFormIdNil sets the value for OriginalFormId to be an explicit nil
func (o *AiFolderDtoInteger) SetOriginalFormIdNil() {
	o.OriginalFormId.Set(nil)
}

// UnsetOriginalFormId ensures that no value is present for OriginalFormId, not even an explicit nil
func (o *AiFolderDtoInteger) UnsetOriginalFormId() {
	o.OriginalFormId.Unset()
}

func (o AiFolderDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFolderDtoInteger) ToMap() (map[string]interface{}, error) {
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

type NullableAiFolderDtoInteger struct {
	value *AiFolderDtoInteger
	isSet bool
}

func (v NullableAiFolderDtoInteger) Get() *AiFolderDtoInteger {
	return v.value
}

func (v *NullableAiFolderDtoInteger) Set(val *AiFolderDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderDtoInteger(val *AiFolderDtoInteger) *NullableAiFolderDtoInteger {
	return &NullableAiFolderDtoInteger{value: val, isSet: true}
}

func (v NullableAiFolderDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

