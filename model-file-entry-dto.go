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

// checks if the FileEntryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileEntryDto{}

// FileEntryDto The part of a file or folder that depends on how the entry is identified: by a number on the portal, or by a  string on a connected third-party account.
type FileEntryDto struct {
	FileEntryBaseDto
	// The identifier to pass back to the other operations of this entry. It is a number for storage on the portal  and a string for a connected third-party account, and it is unique only within its own kind, so files and  folders may carry the same value.
	Id *int32 `json:"id,omitempty"`
	// The section the entry ultimately lies in, as an identifier that can be listed like any other folder. For an  entry inside a room this is the rooms section, not the room.
	RootFolderId *int32 `json:"rootFolderId,omitempty"`
	// The folder the entry was deleted from, which is where restoring it puts it back. It is left out of the answer  unless the entry is in the trash.
	OriginId *int32 `json:"originId,omitempty"`
	// The room the entry was deleted from, left out of the answer for anything that was not deleted out of a room.
	OriginRoomId *int32 `json:"originRoomId,omitempty"`
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
	// The IDs of the metadata templates assigned to the file entry.
	AssignedMetadataTemplates []int32 `json:"assignedMetadataTemplates,omitempty"`
	// Where the search text of the listing was met in the metadata of the entry: one item per string value holding  the text, in the order of the fields. Present only in a listing filtered by text, and only for the entries whose  metadata contain it; an entry found by its title alone carries nothing here.
	MetadataMatches []MetadataMatchDto `json:"metadataMatches,omitempty"`
}

// NewFileEntryDto instantiates a new FileEntryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileEntryDto() *FileEntryDto {
	this := FileEntryDto{}
	return &this
}

// NewFileEntryDtoWithDefaults instantiates a new FileEntryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileEntryDtoWithDefaults() *FileEntryDto {
	this := FileEntryDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileEntryDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileEntryDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileEntryDto) SetId(v int32) {
	o.Id = &v
}

// GetRootFolderId returns the RootFolderId field value if set, zero value otherwise.
func (o *FileEntryDto) GetRootFolderId() int32 {
	if o == nil || IsNil(o.RootFolderId) {
		var ret int32
		return ret
	}
	return *o.RootFolderId
}

// GetRootFolderIdOk returns a tuple with the RootFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetRootFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RootFolderId) {
		return nil, false
	}
	return o.RootFolderId, true
}

// HasRootFolderId returns a boolean if a field has been set.
func (o *FileEntryDto) IsRootFolderIdSet() bool {
	if o != nil && !IsNil(o.RootFolderId) {
		return true
	}

	return false
}

// SetRootFolderId gets a reference to the given int32 and assigns it to the RootFolderId field.
func (o *FileEntryDto) SetRootFolderId(v int32) {
	o.RootFolderId = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *FileEntryDto) GetOriginId() int32 {
	if o == nil || IsNil(o.OriginId) {
		var ret int32
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetOriginIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginId) {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *FileEntryDto) IsOriginIdSet() bool {
	if o != nil && !IsNil(o.OriginId) {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given int32 and assigns it to the OriginId field.
func (o *FileEntryDto) SetOriginId(v int32) {
	o.OriginId = &v
}

// GetOriginRoomId returns the OriginRoomId field value if set, zero value otherwise.
func (o *FileEntryDto) GetOriginRoomId() int32 {
	if o == nil || IsNil(o.OriginRoomId) {
		var ret int32
		return ret
	}
	return *o.OriginRoomId
}

// GetOriginRoomIdOk returns a tuple with the OriginRoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetOriginRoomIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginRoomId) {
		return nil, false
	}
	return o.OriginRoomId, true
}

// HasOriginRoomId returns a boolean if a field has been set.
func (o *FileEntryDto) IsOriginRoomIdSet() bool {
	if o != nil && !IsNil(o.OriginRoomId) {
		return true
	}

	return false
}

// SetOriginRoomId gets a reference to the given int32 and assigns it to the OriginRoomId field.
func (o *FileEntryDto) SetOriginRoomId(v int32) {
	o.OriginRoomId = &v
}

// GetOriginTitle returns the OriginTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetOriginTitle() string {
	if o == nil || IsNil(o.OriginTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginTitle.Get()
}

// GetOriginTitleOk returns a tuple with the OriginTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetOriginTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginTitle.Get(), o.OriginTitle.IsSet()
}

// HasOriginTitle returns a boolean if a field has been set.
func (o *FileEntryDto) IsOriginTitleSet() bool {
	if o != nil && o.OriginTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginTitle gets a reference to the given NullableString and assigns it to the OriginTitle field.
func (o *FileEntryDto) SetOriginTitle(v string) {
	o.OriginTitle.Set(&v)
}
// SetOriginTitleNil sets the value for OriginTitle to be an explicit nil
func (o *FileEntryDto) SetOriginTitleNil() {
	o.OriginTitle.Set(nil)
}

// UnsetOriginTitle ensures that no value is present for OriginTitle, not even an explicit nil
func (o *FileEntryDto) UnsetOriginTitle() {
	o.OriginTitle.Unset()
}

// GetOriginRoomTitle returns the OriginRoomTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetOriginRoomTitle() string {
	if o == nil || IsNil(o.OriginRoomTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginRoomTitle.Get()
}

// GetOriginRoomTitleOk returns a tuple with the OriginRoomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetOriginRoomTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginRoomTitle.Get(), o.OriginRoomTitle.IsSet()
}

// HasOriginRoomTitle returns a boolean if a field has been set.
func (o *FileEntryDto) IsOriginRoomTitleSet() bool {
	if o != nil && o.OriginRoomTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginRoomTitle gets a reference to the given NullableString and assigns it to the OriginRoomTitle field.
func (o *FileEntryDto) SetOriginRoomTitle(v string) {
	o.OriginRoomTitle.Set(&v)
}
// SetOriginRoomTitleNil sets the value for OriginRoomTitle to be an explicit nil
func (o *FileEntryDto) SetOriginRoomTitleNil() {
	o.OriginRoomTitle.Set(nil)
}

// UnsetOriginRoomTitle ensures that no value is present for OriginRoomTitle, not even an explicit nil
func (o *FileEntryDto) UnsetOriginRoomTitle() {
	o.OriginRoomTitle.Unset()
}

// GetCanShare returns the CanShare field value if set, zero value otherwise.
func (o *FileEntryDto) GetCanShare() bool {
	if o == nil || IsNil(o.CanShare) {
		var ret bool
		return ret
	}
	return *o.CanShare
}

// GetCanShareOk returns a tuple with the CanShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetCanShareOk() (*bool, bool) {
	if o == nil || IsNil(o.CanShare) {
		return nil, false
	}
	return o.CanShare, true
}

// HasCanShare returns a boolean if a field has been set.
func (o *FileEntryDto) IsCanShareSet() bool {
	if o != nil && !IsNil(o.CanShare) {
		return true
	}

	return false
}

// SetCanShare gets a reference to the given bool and assigns it to the CanShare field.
func (o *FileEntryDto) SetCanShare(v bool) {
	o.CanShare = &v
}

// GetShareSettings returns the ShareSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetShareSettings() AiFileEntryDtoAllOfShareSettings {
	if o == nil || IsNil(o.ShareSettings.Get()) {
		var ret AiFileEntryDtoAllOfShareSettings
		return ret
	}
	return *o.ShareSettings.Get()
}

// GetShareSettingsOk returns a tuple with the ShareSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetShareSettingsOk() (*AiFileEntryDtoAllOfShareSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareSettings.Get(), o.ShareSettings.IsSet()
}

// HasShareSettings returns a boolean if a field has been set.
func (o *FileEntryDto) IsShareSettingsSet() bool {
	if o != nil && o.ShareSettings.IsSet() {
		return true
	}

	return false
}

// SetShareSettings gets a reference to the given NullableAiFileEntryDtoAllOfShareSettings and assigns it to the ShareSettings field.
func (o *FileEntryDto) SetShareSettings(v AiFileEntryDtoAllOfShareSettings) {
	o.ShareSettings.Set(&v)
}
// SetShareSettingsNil sets the value for ShareSettings to be an explicit nil
func (o *FileEntryDto) SetShareSettingsNil() {
	o.ShareSettings.Set(nil)
}

// UnsetShareSettings ensures that no value is present for ShareSettings, not even an explicit nil
func (o *FileEntryDto) UnsetShareSettings() {
	o.ShareSettings.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetSecurity() AiFileEntryDtoAllOfSecurity {
	if o == nil || IsNil(o.Security.Get()) {
		var ret AiFileEntryDtoAllOfSecurity
		return ret
	}
	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetSecurityOk() (*AiFileEntryDtoAllOfSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// HasSecurity returns a boolean if a field has been set.
func (o *FileEntryDto) IsSecuritySet() bool {
	if o != nil && o.Security.IsSet() {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NullableAiFileEntryDtoAllOfSecurity and assigns it to the Security field.
func (o *FileEntryDto) SetSecurity(v AiFileEntryDtoAllOfSecurity) {
	o.Security.Set(&v)
}
// SetSecurityNil sets the value for Security to be an explicit nil
func (o *FileEntryDto) SetSecurityNil() {
	o.Security.Set(nil)
}

// UnsetSecurity ensures that no value is present for Security, not even an explicit nil
func (o *FileEntryDto) UnsetSecurity() {
	o.Security.Unset()
}

// GetAvailableShareRights returns the AvailableShareRights field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetAvailableShareRights() AiFileEntryDtoAllOfAvailableShareRights {
	if o == nil || IsNil(o.AvailableShareRights.Get()) {
		var ret AiFileEntryDtoAllOfAvailableShareRights
		return ret
	}
	return *o.AvailableShareRights.Get()
}

// GetAvailableShareRightsOk returns a tuple with the AvailableShareRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetAvailableShareRightsOk() (*AiFileEntryDtoAllOfAvailableShareRights, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableShareRights.Get(), o.AvailableShareRights.IsSet()
}

// HasAvailableShareRights returns a boolean if a field has been set.
func (o *FileEntryDto) IsAvailableShareRightsSet() bool {
	if o != nil && o.AvailableShareRights.IsSet() {
		return true
	}

	return false
}

// SetAvailableShareRights gets a reference to the given NullableAiFileEntryDtoAllOfAvailableShareRights and assigns it to the AvailableShareRights field.
func (o *FileEntryDto) SetAvailableShareRights(v AiFileEntryDtoAllOfAvailableShareRights) {
	o.AvailableShareRights.Set(&v)
}
// SetAvailableShareRightsNil sets the value for AvailableShareRights to be an explicit nil
func (o *FileEntryDto) SetAvailableShareRightsNil() {
	o.AvailableShareRights.Set(nil)
}

// UnsetAvailableShareRights ensures that no value is present for AvailableShareRights, not even an explicit nil
func (o *FileEntryDto) UnsetAvailableShareRights() {
	o.AvailableShareRights.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *FileEntryDto) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *FileEntryDto) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *FileEntryDto) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *FileEntryDto) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetExternal() bool {
	if o == nil || IsNil(o.External.Get()) {
		var ret bool
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *FileEntryDto) IsExternalSet() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableBool and assigns it to the External field.
func (o *FileEntryDto) SetExternal(v bool) {
	o.External.Set(&v)
}
// SetExternalNil sets the value for External to be an explicit nil
func (o *FileEntryDto) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *FileEntryDto) UnsetExternal() {
	o.External.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FileEntryDto) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDto) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FileEntryDto) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *FileEntryDto) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetIsLinkExpired returns the IsLinkExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetIsLinkExpired() bool {
	if o == nil || IsNil(o.IsLinkExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLinkExpired.Get()
}

// GetIsLinkExpiredOk returns a tuple with the IsLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetIsLinkExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLinkExpired.Get(), o.IsLinkExpired.IsSet()
}

// HasIsLinkExpired returns a boolean if a field has been set.
func (o *FileEntryDto) IsIsLinkExpiredSet() bool {
	if o != nil && o.IsLinkExpired.IsSet() {
		return true
	}

	return false
}

// SetIsLinkExpired gets a reference to the given NullableBool and assigns it to the IsLinkExpired field.
func (o *FileEntryDto) SetIsLinkExpired(v bool) {
	o.IsLinkExpired.Set(&v)
}
// SetIsLinkExpiredNil sets the value for IsLinkExpired to be an explicit nil
func (o *FileEntryDto) SetIsLinkExpiredNil() {
	o.IsLinkExpired.Set(nil)
}

// UnsetIsLinkExpired ensures that no value is present for IsLinkExpired, not even an explicit nil
func (o *FileEntryDto) UnsetIsLinkExpired() {
	o.IsLinkExpired.Unset()
}

// GetAssignedMetadataTemplates returns the AssignedMetadataTemplates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetAssignedMetadataTemplates() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.AssignedMetadataTemplates
}

// GetAssignedMetadataTemplatesOk returns a tuple with the AssignedMetadataTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetAssignedMetadataTemplatesOk() ([]int32, bool) {
	if o == nil || IsNil(o.AssignedMetadataTemplates) {
		return nil, false
	}
	return o.AssignedMetadataTemplates, true
}

// HasAssignedMetadataTemplates returns a boolean if a field has been set.
func (o *FileEntryDto) IsAssignedMetadataTemplatesSet() bool {
	if o != nil && !IsNil(o.AssignedMetadataTemplates) {
		return true
	}

	return false
}

// SetAssignedMetadataTemplates gets a reference to the given []int32 and assigns it to the AssignedMetadataTemplates field.
func (o *FileEntryDto) SetAssignedMetadataTemplates(v []int32) {
	o.AssignedMetadataTemplates = v
}

// GetMetadataMatches returns the MetadataMatches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileEntryDto) GetMetadataMatches() []MetadataMatchDto {
	if o == nil {
		var ret []MetadataMatchDto
		return ret
	}
	return o.MetadataMatches
}

// GetMetadataMatchesOk returns a tuple with the MetadataMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileEntryDto) GetMetadataMatchesOk() ([]MetadataMatchDto, bool) {
	if o == nil || IsNil(o.MetadataMatches) {
		return nil, false
	}
	return o.MetadataMatches, true
}

// HasMetadataMatches returns a boolean if a field has been set.
func (o *FileEntryDto) IsMetadataMatchesSet() bool {
	if o != nil && !IsNil(o.MetadataMatches) {
		return true
	}

	return false
}

// SetMetadataMatches gets a reference to the given []MetadataMatchDto and assigns it to the MetadataMatches field.
func (o *FileEntryDto) SetMetadataMatches(v []MetadataMatchDto) {
	o.MetadataMatches = v
}

func (o FileEntryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileEntryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFileEntryBaseDto, errFileEntryBaseDto := json.Marshal(o.FileEntryBaseDto)
	if errFileEntryBaseDto != nil {
		return map[string]interface{}{}, errFileEntryBaseDto
	}
	errFileEntryBaseDto = json.Unmarshal([]byte(serializedFileEntryBaseDto), &toSerialize)
	if errFileEntryBaseDto != nil {
		return map[string]interface{}{}, errFileEntryBaseDto
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
	if o.AssignedMetadataTemplates != nil {
		toSerialize["assignedMetadataTemplates"] = o.AssignedMetadataTemplates
	}
	if o.MetadataMatches != nil {
		toSerialize["metadataMatches"] = o.MetadataMatches
	}
	return toSerialize, nil
}

type NullableFileEntryDto struct {
	value *FileEntryDto
	isSet bool
}

func (v NullableFileEntryDto) Get() *FileEntryDto {
	return v.value
}

func (v *NullableFileEntryDto) Set(val *FileEntryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileEntryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileEntryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileEntryDto(val *FileEntryDto) *NullableFileEntryDto {
	return &NullableFileEntryDto{value: val, isSet: true}
}

func (v NullableFileEntryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileEntryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

