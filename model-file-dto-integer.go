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

// checks if the FileDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDtoInteger{}

// FileDtoInteger The file parameters.
type FileDtoInteger struct {
	// The file entry title.
	Title *string `json:"title,omitempty"`
	// The access rights to the file entry.
	Access *FileShare `json:"access,omitempty"`
	// Provides information about the employee who shared the file or folder.
	SharedBy *EmployeeDto `json:"sharedBy,omitempty"`
	// The information about the employee who owns the file entry.
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
	ShortWebUrl *string `json:"shortWebUrl,omitempty"`
	// The creation date and time of the file entry.
	Created *time.Time `json:"created,omitempty"`
	// The file entry author.
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
	// The last date and time when the file entry was updated.
	Updated *time.Time `json:"updated,omitempty"`
	// The date and time when the file entry will be automatically deleted.
	AutoDelete *time.Time `json:"autoDelete,omitempty"`
	// The root folder type of the file entry.
	RootFolderType *FolderType `json:"rootFolderType,omitempty"`
	// The parent room type of the file entry.
	ParentRoomType *FolderType `json:"parentRoomType,omitempty"`
	// The user who updated the file entry.
	UpdatedBy *EmployeeDto `json:"updatedBy,omitempty"`
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
	// The folder ID where the file is located.
	FolderId *int32 `json:"folderId,omitempty"`
	// The file version.
	Version *int32 `json:"version,omitempty"`
	// The version group of the file.
	VersionGroup *int32 `json:"versionGroup,omitempty"`
	// The content length of the file.
	ContentLength NullableString `json:"contentLength,omitempty"`
	// The pure content length of the file.
	PureContentLength NullableInt64 `json:"pureContentLength,omitempty"`
	// The current status of the file.
	FileStatus *FileStatus `json:"fileStatus,omitempty"`
	// The list of users editing the file.
	EditingBy map[string]*string `json:"editingBy,omitempty"`
	// Specifies if the file is muted or not.
	Mute *bool `json:"mute,omitempty"`
	// The URL link to view the file.
	ViewUrl NullableString `json:"viewUrl,omitempty"`
	// The Web URL link to the file.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// The file type.
	FileType *FileType `json:"fileType,omitempty"`
	// The file extension.
	FileExst NullableString `json:"fileExst,omitempty"`
	// The comment to the file.
	Comment NullableString `json:"comment,omitempty"`
	// Specifies if the file is encrypted or not.
	Encrypted NullableBool `json:"encrypted,omitempty"`
	// The thumbnail URL of the file.
	ThumbnailUrl NullableString `json:"thumbnailUrl,omitempty"`
	// The current thumbnail status of the file.
	ThumbnailStatus *Thumbnail `json:"thumbnailStatus,omitempty"`
	// Specifies if the file is locked or not.
	Locked NullableBool `json:"locked,omitempty"`
	// The user ID of the person who locked the file.
	LockedBy NullableString `json:"lockedBy,omitempty"`
	// Specifies if the file has a draft or not.
	HasDraft NullableBool `json:"hasDraft,omitempty"`
	// The status of the form filling process.
	FormFillingStatus *FormFillingStatus `json:"formFillingStatus,omitempty"`
	// Specifies if the file is a form or not.
	IsForm NullableBool `json:"isForm,omitempty"`
	// Specifies if the Custom Filter editing mode is enabled for a file or not.
	CustomFilterEnabled NullableBool `json:"customFilterEnabled,omitempty"`
	// The name of the user who enabled a Custom Filter editing mode for a file.
	CustomFilterEnabledBy NullableString `json:"customFilterEnabledBy,omitempty"`
	// Specifies if the filling has started or not.
	StartFilling NullableBool `json:"startFilling,omitempty"`
	// Specifies if the form filling has started but the file is still being saved by the document editor. Filling and editing are not allowed.
	IsFillingPreparing NullableBool `json:"isFillingPreparing,omitempty"`
	// The InProcess folder ID of the file.
	InProcessFolderId NullableInt32 `json:"inProcessFolderId,omitempty"`
	// The InProcess folder title of the file.
	InProcessFolderTitle NullableString `json:"inProcessFolderTitle,omitempty"`
	// The ID of the FormFillingFolderDone folder that corresponds to this original form.
	ResultsFolderId NullableInt32 `json:"resultsFolderId,omitempty"`
	// The file draft information with its location.
	DraftLocation *DraftLocationInteger `json:"draftLocation,omitempty"`
	ViewAccessibility NullableFileDtoIntegerAllOfViewAccessibility `json:"viewAccessibility,omitempty"`
	// The time when the file was last opened.
	LastOpened NullableTime `json:"lastOpened,omitempty"`
	// The date when the file will be expired.
	Expired NullableTime `json:"expired,omitempty"`
	// The vectorization status of the file.
	VectorizationStatus *VectorizationStatus `json:"vectorizationStatus,omitempty"`
	// The name of the table in the external database that corresponds to this form.
	ExternalDbTableName NullableString `json:"externalDbTableName,omitempty"`
	// The dimensions (width and height) of the image file in pixels.  This property is populated only for image files that can be viewed (supported formats like PNG, JPEG, GIF, BMP, etc.).  For non-image files, this property remains null.
	Dimensions *Size `json:"dimensions,omitempty"`
}

// NewFileDtoInteger instantiates a new FileDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDtoInteger() *FileDtoInteger {
	this := FileDtoInteger{}
	return &this
}

// NewFileDtoIntegerWithDefaults instantiates a new FileDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDtoIntegerWithDefaults() *FileDtoInteger {
	this := FileDtoInteger{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FileDtoInteger) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FileDtoInteger) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FileDtoInteger) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FileDtoInteger) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FileDtoInteger) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FileDtoInteger) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *FileDtoInteger) GetSharedBy() EmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetSharedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given EmployeeDto and assigns it to the SharedBy field.
func (o *FileDtoInteger) SetSharedBy(v EmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOwnedBy() EmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOwnedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given EmployeeDto and assigns it to the OwnedBy field.
func (o *FileDtoInteger) SetOwnedBy(v EmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *FileDtoInteger) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *FileDtoInteger) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *FileDtoInteger) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *FileDtoInteger) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *FileDtoInteger) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *FileDtoInteger) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *FileDtoInteger) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *FileDtoInteger) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *FileDtoInteger) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *FileDtoInteger) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *FileDtoInteger) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *FileDtoInteger) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *FileDtoInteger) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *FileDtoInteger) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *FileDtoInteger) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FileDtoInteger) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FileDtoInteger) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FileDtoInteger) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *FileDtoInteger) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *FileDtoInteger) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *FileDtoInteger) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *FileDtoInteger) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *FileDtoInteger) GetAutoDelete() time.Time {
	if o == nil || IsNil(o.AutoDelete) {
		var ret time.Time
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetAutoDeleteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *FileDtoInteger) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given time.Time and assigns it to the AutoDelete field.
func (o *FileDtoInteger) SetAutoDelete(v time.Time) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *FileDtoInteger) GetRootFolderType() FolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret FolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetRootFolderTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *FileDtoInteger) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given FolderType and assigns it to the RootFolderType field.
func (o *FileDtoInteger) SetRootFolderType(v FolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *FileDtoInteger) GetParentRoomType() FolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret FolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetParentRoomTypeOk() (*FolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *FileDtoInteger) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given FolderType and assigns it to the ParentRoomType field.
func (o *FileDtoInteger) SetParentRoomType(v FolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FileDtoInteger) GetUpdatedBy() EmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetUpdatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given EmployeeDto and assigns it to the UpdatedBy field.
func (o *FileDtoInteger) SetUpdatedBy(v EmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *FileDtoInteger) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *FileDtoInteger) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *FileDtoInteger) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *FileDtoInteger) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *FileDtoInteger) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *FileDtoInteger) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *FileDtoInteger) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *FileDtoInteger) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *FileDtoInteger) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *FileDtoInteger) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *FileDtoInteger) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *FileDtoInteger) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *FileDtoInteger) GetFileEntryType() FileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret FileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetFileEntryTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given FileEntryType and assigns it to the FileEntryType field.
func (o *FileDtoInteger) SetFileEntryType(v FileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileDtoInteger) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileDtoInteger) SetId(v int32) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *FileDtoInteger) GetRootFolderId() int32 {
	if o == nil || IsNil(o.RootFolderId) {
		var ret int32
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetRootFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given int32 and assigns it to the RootFolderId field.
func (o *FileDtoInteger) SetRootFolderId(v int32) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOriginId() int32 {
	if o == nil || IsNil(o.OriginId) {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOriginIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *FileDtoInteger) SetOriginId(v int32) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOriginRoomId() int32 {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret int32
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOriginRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given int32 and assigns it to the OriginRoomId field.
func (o *FileDtoInteger) SetOriginRoomId(v int32) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle) {
		var ret string
		return ret
	}
	return *o.OriginTitle
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOriginTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginTitle) {
		return nil, false
	}
	return o.OriginTitle, true
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOriginTitleSet() bool {
	if o != nil && !IsNil(o.OriginTitle) {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given string and assigns it to the OriginTitle field.
func (o *FileDtoInteger) SetOriginTitle(v string) {
	o.OriginTitle = &v
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise.
func (o *FileDtoInteger) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginRoomTitle) {
		return nil, false
	}
	return o.OriginRoomTitle, true
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *FileDtoInteger) IsOriginRoomTitleSet() bool {
	if o != nil && !IsNil(o.OriginRoomTitle) {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given string and assigns it to the OriginRoomTitle field.
func (o *FileDtoInteger) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle = &v
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *FileDtoInteger) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *FileDtoInteger) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret FileEntryDtoIntegerAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *FileDtoInteger) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableFileEntryDtoIntegerAllOfShareSettings and assigns it to the ShareSettings field.
func (o *FileDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *FileDtoInteger) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *FileDtoInteger) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret FileEntryDtoIntegerAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *FileDtoInteger) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableFileEntryDtoIntegerAllOfSecurity and assigns it to the Security field.
func (o *FileDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *FileDtoInteger) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *FileDtoInteger) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret FileEntryDtoIntegerAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *FileDtoInteger) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableFileEntryDtoIntegerAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *FileDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *FileDtoInteger) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *FileDtoInteger) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *FileDtoInteger) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *FileDtoInteger) IsRequestTokenSet() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *FileDtoInteger) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *FileDtoInteger) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *FileDtoInteger) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *FileDtoInteger) SetExternal(v bool) {
	o.External = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FileDtoInteger) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FileDtoInteger) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *FileDtoInteger) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise.
func (o *FileDtoInteger) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkExpired) {
		return nil, false
	}
	return o.IsLinkExpired, true
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *FileDtoInteger) IsIsLinkExpiredSet() bool {
	if o != nil && !IsNil(o.IsLinkExpired) {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given bool and assigns it to the IsLinkExpired field.
func (o *FileDtoInteger) SetIsLinkExpired(v bool) {
	o.IsLinkExpired = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *FileDtoInteger) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFolderIdSet() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *FileDtoInteger) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FileDtoInteger) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FileDtoInteger) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *FileDtoInteger) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *FileDtoInteger) GetVersionGroup() int32 {
	if o == nil || IsNil(o.VersionGroup) {
		var ret int32
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetVersionGroupOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *FileDtoInteger) IsVersionGroupSet() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given int32 and assigns it to the VersionGroup field.
func (o *FileDtoInteger) SetVersionGroup(v int32) {
	o.VersionGroup = &v
}

// GetContentLength returns the ContentLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetContentLength() string {
	if o == nil || IsNil(o.ContentLength.Get()) {
		var ret string
		return ret
	}
	return *o.ContentLength.Get()
}

// GetContentLengthOk returns a tuple with the ContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetContentLengthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentLength.Get(), o.ContentLength.IsSet()
}

// HasContentLength returns a boolean if a field has been set.
func (o *FileDtoInteger) IsContentLengthSet() bool {
	if o != nil && o.ContentLength.IsSet() {
		return true
	}

	return false
}

// SetContentLength gets a reference to the given NullableString and assigns it to the ContentLength field.
func (o *FileDtoInteger) SetContentLength(v string) {
	o.ContentLength.Set(&v)
}
// SetContentLengthNil sets the value for ContentLength to be an explicit nil
func (o *FileDtoInteger) SetContentLengthNil() {
	o.ContentLength.Set(nil)
}

// UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
func (o *FileDtoInteger) UnsetContentLength() {
	o.ContentLength.Unset()
}

// GetPureContentLength returns the PureContentLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetPureContentLength() int64 {
	if o == nil || IsNil(o.PureContentLength.Get()) {
		var ret int64
		return ret
	}
	return *o.PureContentLength.Get()
}

// GetPureContentLengthOk returns a tuple with the PureContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetPureContentLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PureContentLength.Get(), o.PureContentLength.IsSet()
}

// HasPureContentLength returns a boolean if a field has been set.
func (o *FileDtoInteger) IsPureContentLengthSet() bool {
	if o != nil && o.PureContentLength.IsSet() {
		return true
	}

	return false
}

// SetPureContentLength gets a reference to the given NullableInt64 and assigns it to the PureContentLength field.
func (o *FileDtoInteger) SetPureContentLength(v int64) {
	o.PureContentLength.Set(&v)
}
// SetPureContentLengthNil sets the value for PureContentLength to be an explicit nil
func (o *FileDtoInteger) SetPureContentLengthNil() {
	o.PureContentLength.Set(nil)
}

// UnsetPureContentLength ensures that no value is present for PureContentLength, not even an explicit nil
func (o *FileDtoInteger) UnsetPureContentLength() {
	o.PureContentLength.Unset()
}

// GetFileStatus returns the FileStatus field value if set, zero value otherwise.
func (o *FileDtoInteger) GetFileStatus() FileStatus {
	if o == nil || IsNil(o.FileStatus) {
		var ret FileStatus
		return ret
	}
	return *o.FileStatus
}

// GetFileStatusOk returns a tuple with the FileStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetFileStatusOk() (*FileStatus, bool) {
	if o == nil || IsNil(o.FileStatus) {
		return nil, false
	}
	return o.FileStatus, true
}

// HasFileStatus returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFileStatusSet() bool {
	if o != nil && !IsNil(o.FileStatus) {
		return true
	}

	return false
}

// SetFileStatus gets a reference to the given FileStatus and assigns it to the FileStatus field.
func (o *FileDtoInteger) SetFileStatus(v FileStatus) {
	o.FileStatus = &v
}

// GetEditingBy returns the EditingBy field value if set, zero value otherwise.
func (o *FileDtoInteger) GetEditingBy() map[string]*string {
	if o == nil || IsNil(o.EditingBy) {
		var ret map[string]*string
		return ret
	}
	return o.EditingBy
}

// GetEditingByOk returns a tuple with the EditingBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetEditingByOk() (map[string]*string, bool) {
	if o == nil || IsNil(o.EditingBy) {
		return map[string]*string{}, false
	}
	return o.EditingBy, true
}

// HasEditingBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsEditingBySet() bool {
	if o != nil && !IsNil(o.EditingBy) {
		return true
	}

	return false
}

// SetEditingBy gets a reference to the given map[string]*string and assigns it to the EditingBy field.
func (o *FileDtoInteger) SetEditingBy(v map[string]*string) {
	o.EditingBy = v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *FileDtoInteger) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *FileDtoInteger) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *FileDtoInteger) SetMute(v bool) {
	o.Mute = &v
}

// GetViewUrl returns the ViewUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetViewUrl() string {
	if o == nil || IsNil(o.ViewUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ViewUrl.Get()
}

// GetViewUrlOk returns a tuple with the ViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetViewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewUrl.Get(), o.ViewUrl.IsSet()
}

// HasViewUrl returns a boolean if a field has been set.
func (o *FileDtoInteger) IsViewUrlSet() bool {
	if o != nil && o.ViewUrl.IsSet() {
		return true
	}

	return false
}

// SetViewUrl gets a reference to the given NullableString and assigns it to the ViewUrl field.
func (o *FileDtoInteger) SetViewUrl(v string) {
	o.ViewUrl.Set(&v)
}
// SetViewUrlNil sets the value for ViewUrl to be an explicit nil
func (o *FileDtoInteger) SetViewUrlNil() {
	o.ViewUrl.Set(nil)
}

// UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
func (o *FileDtoInteger) UnsetViewUrl() {
	o.ViewUrl.Unset()
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *FileDtoInteger) IsWebUrlSet() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *FileDtoInteger) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *FileDtoInteger) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *FileDtoInteger) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *FileDtoInteger) GetFileType() FileType {
	if o == nil || IsNil(o.FileType) {
		var ret FileType
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetFileTypeOk() (*FileType, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFileTypeSet() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given FileType and assigns it to the FileType field.
func (o *FileDtoInteger) SetFileType(v FileType) {
	o.FileType = &v
}

// GetFileExst returns the FileExst field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetFileExst() string {
	if o == nil || IsNil(o.FileExst.Get()) {
		var ret string
		return ret
	}
	return *o.FileExst.Get()
}

// GetFileExstOk returns a tuple with the FileExst field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetFileExstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileExst.Get(), o.FileExst.IsSet()
}

// HasFileExst returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFileExstSet() bool {
	if o != nil && o.FileExst.IsSet() {
		return true
	}

	return false
}

// SetFileExst gets a reference to the given NullableString and assigns it to the FileExst field.
func (o *FileDtoInteger) SetFileExst(v string) {
	o.FileExst.Set(&v)
}
// SetFileExstNil sets the value for FileExst to be an explicit nil
func (o *FileDtoInteger) SetFileExstNil() {
	o.FileExst.Set(nil)
}

// UnsetFileExst ensures that no value is present for FileExst, not even an explicit nil
func (o *FileDtoInteger) UnsetFileExst() {
	o.FileExst.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCommentSet() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *FileDtoInteger) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *FileDtoInteger) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *FileDtoInteger) UnsetComment() {
	o.Comment.Unset()
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.Encrypted.Get()
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encrypted.Get(), o.Encrypted.IsSet()
}

// HasEncrypted returns a boolean if a field has been set.
func (o *FileDtoInteger) IsEncryptedSet() bool {
	if o != nil && o.Encrypted.IsSet() {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given NullableBool and assigns it to the Encrypted field.
func (o *FileDtoInteger) SetEncrypted(v bool) {
	o.Encrypted.Set(&v)
}
// SetEncryptedNil sets the value for Encrypted to be an explicit nil
func (o *FileDtoInteger) SetEncryptedNil() {
	o.Encrypted.Set(nil)
}

// UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
func (o *FileDtoInteger) UnsetEncrypted() {
	o.Encrypted.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *FileDtoInteger) IsThumbnailUrlSet() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *FileDtoInteger) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *FileDtoInteger) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *FileDtoInteger) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetThumbnailStatus returns the ThumbnailStatus field value if set, zero value otherwise.
func (o *FileDtoInteger) GetThumbnailStatus() Thumbnail {
	if o == nil || IsNil(o.ThumbnailStatus) {
		var ret Thumbnail
		return ret
	}
	return *o.ThumbnailStatus
}

// GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetThumbnailStatusOk() (*Thumbnail, bool) {
	if o == nil || IsNil(o.ThumbnailStatus) {
		return nil, false
	}
	return o.ThumbnailStatus, true
}

// HasThumbnailStatus returns a boolean if a field has been set.
func (o *FileDtoInteger) IsThumbnailStatusSet() bool {
	if o != nil && !IsNil(o.ThumbnailStatus) {
		return true
	}

	return false
}

// SetThumbnailStatus gets a reference to the given Thumbnail and assigns it to the ThumbnailStatus field.
func (o *FileDtoInteger) SetThumbnailStatus(v Thumbnail) {
	o.ThumbnailStatus = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetLocked() bool {
	if o == nil || IsNil(o.Locked.Get()) {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *FileDtoInteger) IsLockedSet() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *FileDtoInteger) SetLocked(v bool) {
	o.Locked.Set(&v)
}
// SetLockedNil sets the value for Locked to be an explicit nil
func (o *FileDtoInteger) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *FileDtoInteger) UnsetLocked() {
	o.Locked.Unset()
}

// GetLockedBy returns the LockedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetLockedBy() string {
	if o == nil || IsNil(o.LockedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LockedBy.Get()
}

// GetLockedByOk returns a tuple with the LockedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetLockedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedBy.Get(), o.LockedBy.IsSet()
}

// HasLockedBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsLockedBySet() bool {
	if o != nil && o.LockedBy.IsSet() {
		return true
	}

	return false
}

// SetLockedBy gets a reference to the given NullableString and assigns it to the LockedBy field.
func (o *FileDtoInteger) SetLockedBy(v string) {
	o.LockedBy.Set(&v)
}
// SetLockedByNil sets the value for LockedBy to be an explicit nil
func (o *FileDtoInteger) SetLockedByNil() {
	o.LockedBy.Set(nil)
}

// UnsetLockedBy ensures that no value is present for LockedBy, not even an explicit nil
func (o *FileDtoInteger) UnsetLockedBy() {
	o.LockedBy.Unset()
}

// GetHasDraft returns the HasDraft field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetHasDraft() bool {
	if o == nil || IsNil(o.HasDraft.Get()) {
		var ret bool
		return ret
	}
	return *o.HasDraft.Get()
}

// GetHasDraftOk returns a tuple with the HasDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetHasDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasDraft.Get(), o.HasDraft.IsSet()
}

// HasHasDraft returns a boolean if a field has been set.
func (o *FileDtoInteger) IsHasDraftSet() bool {
	if o != nil && o.HasDraft.IsSet() {
		return true
	}

	return false
}

// SetHasDraft gets a reference to the given NullableBool and assigns it to the HasDraft field.
func (o *FileDtoInteger) SetHasDraft(v bool) {
	o.HasDraft.Set(&v)
}
// SetHasDraftNil sets the value for HasDraft to be an explicit nil
func (o *FileDtoInteger) SetHasDraftNil() {
	o.HasDraft.Set(nil)
}

// UnsetHasDraft ensures that no value is present for HasDraft, not even an explicit nil
func (o *FileDtoInteger) UnsetHasDraft() {
	o.HasDraft.Unset()
}

// GetFormFillingStatus returns the FormFillingStatus field value if set, zero value otherwise.
func (o *FileDtoInteger) GetFormFillingStatus() FormFillingStatus {
	if o == nil || IsNil(o.FormFillingStatus) {
		var ret FormFillingStatus
		return ret
	}
	return *o.FormFillingStatus
}

// GetFormFillingStatusOk returns a tuple with the FormFillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetFormFillingStatusOk() (*FormFillingStatus, bool) {
	if o == nil || IsNil(o.FormFillingStatus) {
		return nil, false
	}
	return o.FormFillingStatus, true
}

// HasFormFillingStatus returns a boolean if a field has been set.
func (o *FileDtoInteger) IsFormFillingStatusSet() bool {
	if o != nil && !IsNil(o.FormFillingStatus) {
		return true
	}

	return false
}

// SetFormFillingStatus gets a reference to the given FormFillingStatus and assigns it to the FormFillingStatus field.
func (o *FileDtoInteger) SetFormFillingStatus(v FormFillingStatus) {
	o.FormFillingStatus = &v
}

// GetIsForm returns the IsForm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetIsForm() bool {
	if o == nil || IsNil(o.IsForm.Get()) {
		var ret bool
		return ret
	}
	return *o.IsForm.Get()
}

// GetIsFormOk returns a tuple with the IsForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetIsFormOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsForm.Get(), o.IsForm.IsSet()
}

// HasIsForm returns a boolean if a field has been set.
func (o *FileDtoInteger) IsIsFormSet() bool {
	if o != nil && o.IsForm.IsSet() {
		return true
	}

	return false
}

// SetIsForm gets a reference to the given NullableBool and assigns it to the IsForm field.
func (o *FileDtoInteger) SetIsForm(v bool) {
	o.IsForm.Set(&v)
}
// SetIsFormNil sets the value for IsForm to be an explicit nil
func (o *FileDtoInteger) SetIsFormNil() {
	o.IsForm.Set(nil)
}

// UnsetIsForm ensures that no value is present for IsForm, not even an explicit nil
func (o *FileDtoInteger) UnsetIsForm() {
	o.IsForm.Unset()
}

// GetCustomFilterEnabled returns the CustomFilterEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetCustomFilterEnabled() bool {
	if o == nil || IsNil(o.CustomFilterEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.CustomFilterEnabled.Get()
}

// GetCustomFilterEnabledOk returns a tuple with the CustomFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetCustomFilterEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFilterEnabled.Get(), o.CustomFilterEnabled.IsSet()
}

// HasCustomFilterEnabled returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCustomFilterEnabledSet() bool {
	if o != nil && o.CustomFilterEnabled.IsSet() {
		return true
	}

	return false
}

// SetCustomFilterEnabled gets a reference to the given NullableBool and assigns it to the CustomFilterEnabled field.
func (o *FileDtoInteger) SetCustomFilterEnabled(v bool) {
	o.CustomFilterEnabled.Set(&v)
}
// SetCustomFilterEnabledNil sets the value for CustomFilterEnabled to be an explicit nil
func (o *FileDtoInteger) SetCustomFilterEnabledNil() {
	o.CustomFilterEnabled.Set(nil)
}

// UnsetCustomFilterEnabled ensures that no value is present for CustomFilterEnabled, not even an explicit nil
func (o *FileDtoInteger) UnsetCustomFilterEnabled() {
	o.CustomFilterEnabled.Unset()
}

// GetCustomFilterEnabledBy returns the CustomFilterEnabledBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetCustomFilterEnabledBy() string {
	if o == nil || IsNil(o.CustomFilterEnabledBy.Get()) {
		var ret string
		return ret
	}
	return *o.CustomFilterEnabledBy.Get()
}

// GetCustomFilterEnabledByOk returns a tuple with the CustomFilterEnabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetCustomFilterEnabledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFilterEnabledBy.Get(), o.CustomFilterEnabledBy.IsSet()
}

// HasCustomFilterEnabledBy returns a boolean if a field has been set.
func (o *FileDtoInteger) IsCustomFilterEnabledBySet() bool {
	if o != nil && o.CustomFilterEnabledBy.IsSet() {
		return true
	}

	return false
}

// SetCustomFilterEnabledBy gets a reference to the given NullableString and assigns it to the CustomFilterEnabledBy field.
func (o *FileDtoInteger) SetCustomFilterEnabledBy(v string) {
	o.CustomFilterEnabledBy.Set(&v)
}
// SetCustomFilterEnabledByNil sets the value for CustomFilterEnabledBy to be an explicit nil
func (o *FileDtoInteger) SetCustomFilterEnabledByNil() {
	o.CustomFilterEnabledBy.Set(nil)
}

// UnsetCustomFilterEnabledBy ensures that no value is present for CustomFilterEnabledBy, not even an explicit nil
func (o *FileDtoInteger) UnsetCustomFilterEnabledBy() {
	o.CustomFilterEnabledBy.Unset()
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling.Get()) {
		var ret bool
		return ret
	}
	return *o.StartFilling.Get()
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetStartFillingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartFilling.Get(), o.StartFilling.IsSet()
}

// HasStartFilling returns a boolean if a field has been set.
func (o *FileDtoInteger) IsStartFillingSet() bool {
	if o != nil && o.StartFilling.IsSet() {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given NullableBool and assigns it to the StartFilling field.
func (o *FileDtoInteger) SetStartFilling(v bool) {
	o.StartFilling.Set(&v)
}
// SetStartFillingNil sets the value for StartFilling to be an explicit nil
func (o *FileDtoInteger) SetStartFillingNil() {
	o.StartFilling.Set(nil)
}

// UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
func (o *FileDtoInteger) UnsetStartFilling() {
	o.StartFilling.Unset()
}

// GetIsFillingPreparing returns the IsFillingPreparing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetIsFillingPreparing() bool {
	if o == nil || IsNil(o.IsFillingPreparing.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFillingPreparing.Get()
}

// GetIsFillingPreparingOk returns a tuple with the IsFillingPreparing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetIsFillingPreparingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFillingPreparing.Get(), o.IsFillingPreparing.IsSet()
}

// HasIsFillingPreparing returns a boolean if a field has been set.
func (o *FileDtoInteger) IsIsFillingPreparingSet() bool {
	if o != nil && o.IsFillingPreparing.IsSet() {
		return true
	}

	return false
}

// SetIsFillingPreparing gets a reference to the given NullableBool and assigns it to the IsFillingPreparing field.
func (o *FileDtoInteger) SetIsFillingPreparing(v bool) {
	o.IsFillingPreparing.Set(&v)
}
// SetIsFillingPreparingNil sets the value for IsFillingPreparing to be an explicit nil
func (o *FileDtoInteger) SetIsFillingPreparingNil() {
	o.IsFillingPreparing.Set(nil)
}

// UnsetIsFillingPreparing ensures that no value is present for IsFillingPreparing, not even an explicit nil
func (o *FileDtoInteger) UnsetIsFillingPreparing() {
	o.IsFillingPreparing.Unset()
}

// GetInProcessFolderId returns the InProcessFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetInProcessFolderId() int32 {
	if o == nil || IsNil(o.InProcessFolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.InProcessFolderId.Get()
}

// GetInProcessFolderIdOk returns a tuple with the InProcessFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetInProcessFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InProcessFolderId.Get(), o.InProcessFolderId.IsSet()
}

// HasInProcessFolderId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsInProcessFolderIdSet() bool {
	if o != nil && o.InProcessFolderId.IsSet() {
		return true
	}

	return false
}

// SetInProcessFolderId gets a reference to the given NullableInt32 and assigns it to the InProcessFolderId field.
func (o *FileDtoInteger) SetInProcessFolderId(v int32) {
	o.InProcessFolderId.Set(&v)
}
// SetInProcessFolderIdNil sets the value for InProcessFolderId to be an explicit nil
func (o *FileDtoInteger) SetInProcessFolderIdNil() {
	o.InProcessFolderId.Set(nil)
}

// UnsetInProcessFolderId ensures that no value is present for InProcessFolderId, not even an explicit nil
func (o *FileDtoInteger) UnsetInProcessFolderId() {
	o.InProcessFolderId.Unset()
}

// GetInProcessFolderTitle returns the InProcessFolderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetInProcessFolderTitle() string {
	if o == nil || IsNil(o.InProcessFolderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.InProcessFolderTitle.Get()
}

// GetInProcessFolderTitleOk returns a tuple with the InProcessFolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetInProcessFolderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InProcessFolderTitle.Get(), o.InProcessFolderTitle.IsSet()
}

// HasInProcessFolderTitle returns a boolean if a field has been set.
func (o *FileDtoInteger) IsInProcessFolderTitleSet() bool {
	if o != nil && o.InProcessFolderTitle.IsSet() {
		return true
	}

	return false
}

// SetInProcessFolderTitle gets a reference to the given NullableString and assigns it to the InProcessFolderTitle field.
func (o *FileDtoInteger) SetInProcessFolderTitle(v string) {
	o.InProcessFolderTitle.Set(&v)
}
// SetInProcessFolderTitleNil sets the value for InProcessFolderTitle to be an explicit nil
func (o *FileDtoInteger) SetInProcessFolderTitleNil() {
	o.InProcessFolderTitle.Set(nil)
}

// UnsetInProcessFolderTitle ensures that no value is present for InProcessFolderTitle, not even an explicit nil
func (o *FileDtoInteger) UnsetInProcessFolderTitle() {
	o.InProcessFolderTitle.Unset()
}

// GetResultsFolderId returns the ResultsFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetResultsFolderId() int32 {
	if o == nil || IsNil(o.ResultsFolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ResultsFolderId.Get()
}

// GetResultsFolderIdOk returns a tuple with the ResultsFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetResultsFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultsFolderId.Get(), o.ResultsFolderId.IsSet()
}

// HasResultsFolderId returns a boolean if a field has been set.
func (o *FileDtoInteger) IsResultsFolderIdSet() bool {
	if o != nil && o.ResultsFolderId.IsSet() {
		return true
	}

	return false
}

// SetResultsFolderId gets a reference to the given NullableInt32 and assigns it to the ResultsFolderId field.
func (o *FileDtoInteger) SetResultsFolderId(v int32) {
	o.ResultsFolderId.Set(&v)
}
// SetResultsFolderIdNil sets the value for ResultsFolderId to be an explicit nil
func (o *FileDtoInteger) SetResultsFolderIdNil() {
	o.ResultsFolderId.Set(nil)
}

// UnsetResultsFolderId ensures that no value is present for ResultsFolderId, not even an explicit nil
func (o *FileDtoInteger) UnsetResultsFolderId() {
	o.ResultsFolderId.Unset()
}

// GetDraftLocation returns the DraftLocation field value if set, zero value otherwise.
func (o *FileDtoInteger) GetDraftLocation() DraftLocationInteger {
	if o == nil || IsNil(o.DraftLocation) {
		var ret DraftLocationInteger
		return ret
	}
	return *o.DraftLocation
}

// GetDraftLocationOk returns a tuple with the DraftLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetDraftLocationOk() (*DraftLocationInteger, bool) {
	if o == nil || IsNil(o.DraftLocation) {
		return nil, false
	}
	return o.DraftLocation, true
}

// HasDraftLocation returns a boolean if a field has been set.
func (o *FileDtoInteger) IsDraftLocationSet() bool {
	if o != nil && !IsNil(o.DraftLocation) {
		return true
	}

	return false
}

// SetDraftLocation gets a reference to the given DraftLocationInteger and assigns it to the DraftLocation field.
func (o *FileDtoInteger) SetDraftLocation(v DraftLocationInteger) {
	o.DraftLocation = &v
}

// GetViewAccessibility returns the ViewAccessibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetViewAccessibility() FileDtoIntegerAllOfViewAccessibility {
	if o == nil || IsNil(o.ViewAccessibility.Get()) {
		var ret FileDtoIntegerAllOfViewAccessibility
		return ret
	}
	return *o.ViewAccessibility.Get()
}

// GetViewAccessibilityOk returns a tuple with the ViewAccessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetViewAccessibilityOk() (*FileDtoIntegerAllOfViewAccessibility, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewAccessibility.Get(), o.ViewAccessibility.IsSet()
}

// HasViewAccessibility returns a boolean if a field has been set.
func (o *FileDtoInteger) IsViewAccessibilitySet() bool {
	if o != nil && o.ViewAccessibility.IsSet() {
		return true
	}

	return false
}

// SetViewAccessibility gets a reference to the given NullableFileDtoIntegerAllOfViewAccessibility and assigns it to the ViewAccessibility field.
func (o *FileDtoInteger) SetViewAccessibility(v FileDtoIntegerAllOfViewAccessibility) {
	o.ViewAccessibility.Set(&v)
}
// SetViewAccessibilityNil sets the value for ViewAccessibility to be an explicit nil
func (o *FileDtoInteger) SetViewAccessibilityNil() {
	o.ViewAccessibility.Set(nil)
}

// UnsetViewAccessibility ensures that no value is present for ViewAccessibility, not even an explicit nil
func (o *FileDtoInteger) UnsetViewAccessibility() {
	o.ViewAccessibility.Unset()
}

// GetLastOpened returns the LastOpened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetLastOpened() time.Time {
	if o == nil || IsNil(o.LastOpened.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastOpened.Get()
}

// GetLastOpenedOk returns a tuple with the LastOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetLastOpenedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOpened.Get(), o.LastOpened.IsSet()
}

// HasLastOpened returns a boolean if a field has been set.
func (o *FileDtoInteger) IsLastOpenedSet() bool {
	if o != nil && o.LastOpened.IsSet() {
		return true
	}

	return false
}

// SetLastOpened gets a reference to the given NullableTime and assigns it to the LastOpened field.
func (o *FileDtoInteger) SetLastOpened(v time.Time) {
	o.LastOpened.Set(&v)
}
// SetLastOpenedNil sets the value for LastOpened to be an explicit nil
func (o *FileDtoInteger) SetLastOpenedNil() {
	o.LastOpened.Set(nil)
}

// UnsetLastOpened ensures that no value is present for LastOpened, not even an explicit nil
func (o *FileDtoInteger) UnsetLastOpened() {
	o.LastOpened.Unset()
}

// GetExpired returns the Expired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetExpired() time.Time {
	if o == nil || IsNil(o.Expired.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expired.Get()
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetExpiredOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expired.Get(), o.Expired.IsSet()
}

// HasExpired returns a boolean if a field has been set.
func (o *FileDtoInteger) IsExpiredSet() bool {
	if o != nil && o.Expired.IsSet() {
		return true
	}

	return false
}

// SetExpired gets a reference to the given NullableTime and assigns it to the Expired field.
func (o *FileDtoInteger) SetExpired(v time.Time) {
	o.Expired.Set(&v)
}
// SetExpiredNil sets the value for Expired to be an explicit nil
func (o *FileDtoInteger) SetExpiredNil() {
	o.Expired.Set(nil)
}

// UnsetExpired ensures that no value is present for Expired, not even an explicit nil
func (o *FileDtoInteger) UnsetExpired() {
	o.Expired.Unset()
}

// GetVectorizationStatus returns the VectorizationStatus field value if set, zero value otherwise.
func (o *FileDtoInteger) GetVectorizationStatus() VectorizationStatus {
	if o == nil || IsNil(o.VectorizationStatus) {
		var ret VectorizationStatus
		return ret
	}
	return *o.VectorizationStatus
}

// GetVectorizationStatusOk returns a tuple with the VectorizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetVectorizationStatusOk() (*VectorizationStatus, bool) {
	if o == nil || IsNil(o.VectorizationStatus) {
		return nil, false
	}
	return o.VectorizationStatus, true
}

// HasVectorizationStatus returns a boolean if a field has been set.
func (o *FileDtoInteger) IsVectorizationStatusSet() bool {
	if o != nil && !IsNil(o.VectorizationStatus) {
		return true
	}

	return false
}

// SetVectorizationStatus gets a reference to the given VectorizationStatus and assigns it to the VectorizationStatus field.
func (o *FileDtoInteger) SetVectorizationStatus(v VectorizationStatus) {
	o.VectorizationStatus = &v
}

// GetExternalDbTableName returns the ExternalDbTableName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileDtoInteger) GetExternalDbTableName() string {
	if o == nil || IsNil(o.ExternalDbTableName.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalDbTableName.Get()
}

// GetExternalDbTableNameOk returns a tuple with the ExternalDbTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileDtoInteger) GetExternalDbTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalDbTableName.Get(), o.ExternalDbTableName.IsSet()
}

// HasExternalDbTableName returns a boolean if a field has been set.
func (o *FileDtoInteger) IsExternalDbTableNameSet() bool {
	if o != nil && o.ExternalDbTableName.IsSet() {
		return true
	}

	return false
}

// SetExternalDbTableName gets a reference to the given NullableString and assigns it to the ExternalDbTableName field.
func (o *FileDtoInteger) SetExternalDbTableName(v string) {
	o.ExternalDbTableName.Set(&v)
}
// SetExternalDbTableNameNil sets the value for ExternalDbTableName to be an explicit nil
func (o *FileDtoInteger) SetExternalDbTableNameNil() {
	o.ExternalDbTableName.Set(nil)
}

// UnsetExternalDbTableName ensures that no value is present for ExternalDbTableName, not even an explicit nil
func (o *FileDtoInteger) UnsetExternalDbTableName() {
	o.ExternalDbTableName.Unset()
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *FileDtoInteger) GetDimensions() Size {
	if o == nil || IsNil(o.Dimensions) {
		var ret Size
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoInteger) GetDimensionsOk() (*Size, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *FileDtoInteger) IsDimensionsSet() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Size and assigns it to the Dimensions field.
func (o *FileDtoInteger) SetDimensions(v Size) {
	o.Dimensions = &v
}

func (o FileDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDtoInteger) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
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
	if o.LastOpened.IsSet() {
		toSerialize["lastOpened"] = o.LastOpened.Get()
	}
	if o.Expired.IsSet() {
		toSerialize["expired"] = o.Expired.Get()
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

type NullableFileDtoInteger struct {
	value *FileDtoInteger
	isSet bool
}

func (v NullableFileDtoInteger) Get() *FileDtoInteger {
	return v.value
}

func (v *NullableFileDtoInteger) Set(val *FileDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDtoInteger(val *FileDtoInteger) *NullableFileDtoInteger {
	return &NullableFileDtoInteger{value: val, isSet: true}
}

func (v NullableFileDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

