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

// checks if the AiFileEntryDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileEntryDtoInteger{}

// AiFileEntryDtoInteger The generic file entry information.
type AiFileEntryDtoInteger struct {
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
	// Represents the expiration date of the file entry.
	ExpirationDate NullableTime `json:"expirationDate,omitempty"`
	// Indicates whether the shareable link associated with the file or folder has expired.
	IsLinkExpired NullableBool `json:"isLinkExpired,omitempty"`
}

// NewAiFileEntryDtoInteger instantiates a new AiFileEntryDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileEntryDtoInteger() *AiFileEntryDtoInteger {
	this := AiFileEntryDtoInteger{}
	return &this
}

// NewAiFileEntryDtoIntegerWithDefaults instantiates a new AiFileEntryDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileEntryDtoIntegerWithDefaults() *AiFileEntryDtoInteger {
	this := AiFileEntryDtoInteger{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiFileEntryDtoInteger) SetTitle(v string) {
	o.Title = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetAccess() AiFileShare {
	if o == nil || IsNil(o.Access) {
		var ret AiFileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetAccessOk() (*AiFileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AiFileShare and assigns it to the Access field.
func (o *AiFileEntryDtoInteger) SetAccess(v AiFileShare) {
	o.Access = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetSharedBy() AiEmployeeDto {
	if o == nil || IsNil(o.SharedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetSharedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsSharedBySet() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given AiEmployeeDto and assigns it to the SharedBy field.
func (o *AiFileEntryDtoInteger) SetSharedBy(v AiEmployeeDto) {
	o.SharedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetOwnedBy() AiEmployeeDto {
	if o == nil || IsNil(o.OwnedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetOwnedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.OwnedBy) {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOwnedBySet() bool {
	if o != nil && !IsNil(o.OwnedBy) {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given AiEmployeeDto and assigns it to the OwnedBy field.
func (o *AiFileEntryDtoInteger) SetOwnedBy(v AiEmployeeDto) {
	o.OwnedBy = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsSharedSet() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *AiFileEntryDtoInteger) SetShared(v bool) {
	o.Shared = &v
}

// GetSharedForUser returns the SharedForUser field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetSharedForUser() bool {
	if o == nil || IsNil(o.SharedForUser) {
		var ret bool
		return ret
	}
	return *o.SharedForUser
}

// GetSharedForUserOk returns a tuple with the SharedForUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetSharedForUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedForUser) {
		return nil, false
	}
	return o.SharedForUser, true
}

// HasSharedForUser returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsSharedForUserSet() bool {
	if o != nil && !IsNil(o.SharedForUser) {
		return true
	}

	return false
}

// SetSharedForUser gets a reference to the given bool and assigns it to the SharedForUser field.
func (o *AiFileEntryDtoInteger) SetSharedForUser(v bool) {
	o.SharedForUser = &v
}

// GetSharedExternal returns the SharedExternal field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetSharedExternal() bool {
	if o == nil || IsNil(o.SharedExternal) {
		var ret bool
		return ret
	}
	return *o.SharedExternal
}

// GetSharedExternalOk returns a tuple with the SharedExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetSharedExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedExternal) {
		return nil, false
	}
	return o.SharedExternal, true
}

// HasSharedExternal returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsSharedExternalSet() bool {
	if o != nil && !IsNil(o.SharedExternal) {
		return true
	}

	return false
}

// SetSharedExternal gets a reference to the given bool and assigns it to the SharedExternal field.
func (o *AiFileEntryDtoInteger) SetSharedExternal(v bool) {
	o.SharedExternal = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsParentSharedSet() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *AiFileEntryDtoInteger) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetShortWebUrl returns the ShortWebUrl field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetShortWebUrl() string {
	if o == nil || IsNil(o.ShortWebUrl) {
		var ret string
		return ret
	}
	return *o.ShortWebUrl
}

// GetShortWebUrlOk returns a tuple with the ShortWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetShortWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ShortWebUrl) {
		return nil, false
	}
	return o.ShortWebUrl, true
}

// HasShortWebUrl returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsShortWebUrlSet() bool {
	if o != nil && !IsNil(o.ShortWebUrl) {
		return true
	}

	return false
}

// SetShortWebUrl gets a reference to the given string and assigns it to the ShortWebUrl field.
func (o *AiFileEntryDtoInteger) SetShortWebUrl(v string) {
	o.ShortWebUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AiFileEntryDtoInteger) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetCreatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetCreatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AiEmployeeDto and assigns it to the CreatedBy field.
func (o *AiFileEntryDtoInteger) SetCreatedBy(v AiEmployeeDto) {
	o.CreatedBy = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsUpdatedSet() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AiFileEntryDtoInteger) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetAutoDelete() time.Time {
	if o == nil || IsNil(o.AutoDelete) {
		var ret time.Time
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetAutoDeleteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AutoDelete) {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsAutoDeleteSet() bool {
	if o != nil && !IsNil(o.AutoDelete) {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given time.Time and assigns it to the AutoDelete field.
func (o *AiFileEntryDtoInteger) SetAutoDelete(v time.Time) {
	o.AutoDelete = &v
}

// GetRootFolderType returns the RootFolderType field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetRootFolderType() AiFolderType {
	if o == nil || IsNil(o.RootFolderType) {
		var ret AiFolderType
		return ret
	}
	return *o.RootFolderType
}

// GetRootFolderTypeOk returns a tuple with the RootFolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetRootFolderTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.RootFolderType) {
		return nil, false
	}
	return o.RootFolderType, true
}

// HasRootFolderType returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsRootFolderTypeSet() bool {
	if o != nil && !IsNil(o.RootFolderType) {
		return true
	}

	return false
}

// SetRootFolderType gets a reference to the given AiFolderType and assigns it to the RootFolderType field.
func (o *AiFileEntryDtoInteger) SetRootFolderType(v AiFolderType) {
	o.RootFolderType = &v
}

// GetParentRoomType returns the ParentRoomType field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetParentRoomType() AiFolderType {
	if o == nil || IsNil(o.ParentRoomType) {
		var ret AiFolderType
		return ret
	}
	return *o.ParentRoomType
}

// GetParentRoomTypeOk returns a tuple with the ParentRoomType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetParentRoomTypeOk() (*AiFolderType, bool) {
	if o == nil || IsNil(o.ParentRoomType) {
		return nil, false
	}
	return o.ParentRoomType, true
}

// HasParentRoomType returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsParentRoomTypeSet() bool {
	if o != nil && !IsNil(o.ParentRoomType) {
		return true
	}

	return false
}

// SetParentRoomType gets a reference to the given AiFolderType and assigns it to the ParentRoomType field.
func (o *AiFileEntryDtoInteger) SetParentRoomType(v AiFolderType) {
	o.ParentRoomType = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetUpdatedBy() AiEmployeeDto {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret AiEmployeeDto
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetUpdatedByOk() (*AiEmployeeDto, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsUpdatedBySet() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given AiEmployeeDto and assigns it to the UpdatedBy field.
func (o *AiFileEntryDtoInteger) SetUpdatedBy(v AiEmployeeDto) {
	o.UpdatedBy = &v
}

// GetProviderItem returns the ProviderItem field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetProviderItem() bool {
	if o == nil || IsNil(o.ProviderItem) {
		var ret bool
		return ret
	}
	return *o.ProviderItem
}

// GetProviderItemOk returns a tuple with the ProviderItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetProviderItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderItem) {
		return nil, false
	}
	return o.ProviderItem, true
}

// HasProviderItem returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsProviderItemSet() bool {
	if o != nil && !IsNil(o.ProviderItem) {
		return true
	}

	return false
}

// SetProviderItem gets a reference to the given bool and assigns it to the ProviderItem field.
func (o *AiFileEntryDtoInteger) SetProviderItem(v bool) {
	o.ProviderItem = &v
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey) {
		var ret string
		return ret
	}
	return *o.ProviderKey
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetProviderKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderKey) {
		return nil, false
	}
	return o.ProviderKey, true
}

// HasProviderKey returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsProviderKeySet() bool {
	if o != nil && !IsNil(o.ProviderKey) {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given string and assigns it to the ProviderKey field.
func (o *AiFileEntryDtoInteger) SetProviderKey(v string) {
	o.ProviderKey = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *AiFileEntryDtoInteger) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOrderSet() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *AiFileEntryDtoInteger) SetOrder(v string) {
	o.Order = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsIsFavoriteSet() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *AiFileEntryDtoInteger) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetFileEntryType returns the FileEntryType field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetFileEntryType() AiFileEntryType {
	if o == nil || IsNil(o.FileEntryType) {
		var ret AiFileEntryType
		return ret
	}
	return *o.FileEntryType
}

// GetFileEntryTypeOk returns a tuple with the FileEntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetFileEntryTypeOk() (*AiFileEntryType, bool) {
	if o == nil || IsNil(o.FileEntryType) {
		return nil, false
	}
	return o.FileEntryType, true
}

// HasFileEntryType returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsFileEntryTypeSet() bool {
	if o != nil && !IsNil(o.FileEntryType) {
		return true
	}

	return false
}

// SetFileEntryType gets a reference to the given AiFileEntryType and assigns it to the FileEntryType field.
func (o *AiFileEntryDtoInteger) SetFileEntryType(v AiFileEntryType) {
	o.FileEntryType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AiFileEntryDtoInteger) SetId(v int32) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetRootFolderId() int32 {
	if o == nil || IsNil(o.RootFolderId) {
		var ret int32
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetRootFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given int32 and assigns it to the RootFolderId field.
func (o *AiFileEntryDtoInteger) SetRootFolderId(v int32) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetOriginId() int32 {
	if o == nil || IsNil(o.OriginId) {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetOriginIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *AiFileEntryDtoInteger) SetOriginId(v int32) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetOriginRoomId() int32 {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret int32
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetOriginRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given int32 and assigns it to the OriginRoomId field.
func (o *AiFileEntryDtoInteger) SetOriginRoomId(v int32) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginTitle.Get()
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetOriginTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginTitle.Get(), o.OriginTitle.IsSet()
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOriginTitleSet() bool {
	if o != nil && o.OriginTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given NullableString and assigns it to the OriginTitle field.
func (o *AiFileEntryDtoInteger) SetOriginTitle(v string) {
	o.OriginTitle.Set(&v)
}
// SetOriginTitleNil sets the value for OriginTitle to be an explicit nil
func (o *AiFileEntryDtoInteger) SetOriginTitleNil() {
	o.OriginTitle.Set(nil)
}

// UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetOriginTitle() {
	o.OriginTitle.Unset()
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle.Get()
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomTitle.Get(), o.OriginRoomTitle.IsSet()
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsOriginRoomTitleSet() bool {
	if o != nil && o.OriginRoomTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given NullableString and assigns it to the OriginRoomTitle field.
func (o *AiFileEntryDtoInteger) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle.Set(&v)
}
// SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil
func (o *AiFileEntryDtoInteger) SetOriginRoomTitleNil() {
	o.OriginRoomTitle.Set(nil)
}

// UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetOriginRoomTitle() {
	o.OriginRoomTitle.Unset()
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *AiFileEntryDtoInteger) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoInteger) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *AiFileEntryDtoInteger) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetShareSettings() FileEntryDtoIntegerAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret FileEntryDtoIntegerAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetShareSettingsOk() (*FileEntryDtoIntegerAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableFileEntryDtoIntegerAllOfShareSettings and assigns it to the ShareSettings field.
func (o *AiFileEntryDtoInteger) SetShareSettings(v FileEntryDtoIntegerAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *AiFileEntryDtoInteger) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetSecurity() FileEntryDtoIntegerAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret FileEntryDtoIntegerAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetSecurityOk() (*FileEntryDtoIntegerAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableFileEntryDtoIntegerAllOfSecurity and assigns it to the Security field.
func (o *AiFileEntryDtoInteger) SetSecurity(v FileEntryDtoIntegerAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *AiFileEntryDtoInteger) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetAvailableShareRights() FileEntryDtoIntegerAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret FileEntryDtoIntegerAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetAvailableShareRightsOk() (*FileEntryDtoIntegerAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableFileEntryDtoIntegerAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *AiFileEntryDtoInteger) SetAvailableShareRights(v FileEntryDtoIntegerAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *AiFileEntryDtoInteger) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *AiFileEntryDtoInteger) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *AiFileEntryDtoInteger) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetExternal() bool {
	if o == nil || IsNil(o.External.Get()) {
		var ret bool
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsExternalSet() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableBool and assigns it to the External field.
func (o *AiFileEntryDtoInteger) SetExternal(v bool) {
	o.External.Set(&v)
}
// SetExternalNil sets the value for External to be an explicit nil
func (o *AiFileEntryDtoInteger) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetExternal() {
	o.External.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsExpirationDateSet() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *AiFileEntryDtoInteger) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *AiFileEntryDtoInteger) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileEntryDtoInteger) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired.Get()
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileEntryDtoInteger) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLinkExpired.Get(), o.IsLinkExpired.IsSet()
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *AiFileEntryDtoInteger) IsIsLinkExpiredSet() bool {
	if o != nil && o.IsLinkExpired.IsSet() {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given NullableBool and assigns it to the IsLinkExpired field.
func (o *AiFileEntryDtoInteger) SetIsLinkExpired(v bool) {
	o.IsLinkExpired.Set(&v)
}
// SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil
func (o *AiFileEntryDtoInteger) SetIsLinkExpiredNil() {
	o.IsLinkExpired.Set(nil)
}

// UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
func (o *AiFileEntryDtoInteger) UnsetIsLinkExpired() {
	o.IsLinkExpired.Unset()
}

func (o AiFileEntryDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileEntryDtoInteger) ToMap() (map[string]interface{}, error) {
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
	if o.ExpirationDate.IsSet() {
		toSerialize["expirationDate"] = o.ExpirationDate.Get()
	}
	if o.IsLinkExpired.IsSet() {
		toSerialize["isLinkExpired"] = o.IsLinkExpired.Get()
	}
	return toSerialize, nil
}

type NullableAiFileEntryDtoInteger struct {
	value *AiFileEntryDtoInteger
	isSet bool
}

func (v NullableAiFileEntryDtoInteger) Get() *AiFileEntryDtoInteger {
	return v.value
}

func (v *NullableAiFileEntryDtoInteger) Set(val *AiFileEntryDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryDtoInteger(val *AiFileEntryDtoInteger) *NullableAiFileEntryDtoInteger {
	return &NullableAiFileEntryDtoInteger{value: val, isSet: true}
}

func (v NullableAiFileEntryDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

