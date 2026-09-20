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
	"bytes"
	"fmt"
)

// checks if the ExternalShareDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalShareDto{}

// ExternalShareDto The outcome of validating an external share link and the entry it points at.
type ExternalShareDto struct {
	// How validating the link went. It is the first field to read: a refused link is reported here with the answer  still arriving as a success. A link that resolved describes both the entry and the link, one that is waiting  for its password describes only the entry, and one that failed outright leaves the rest of the object empty.
	Status Status `json:"status"`
	// The identifier of the room, folder or file the link points at, always rendered as a string even where the  portal stores it as a number. It is null when the link could not be resolved.
	Id NullableString `json:"id,omitempty"`
	// The title of the entry the link points at, suitable for showing to the visitor before they are let in. It is  null when the link could not be resolved.
	Title NullableString `json:"title,omitempty"`
	// Whether the link points at a folder - a room counts as one - or at a single file. It is null when the link  could not be resolved.
	Type *FileEntryType `json:"type,omitempty"`
	// The portal the link belongs to, which matters for a client that works with more than one. It stays 0 for a  link that did not resolve.
	TenantId int32 `json:"tenantId"`
	// The identifier of the entry that was asked about through the request's file or folder parameter, echoed back  once it was found under the link's target. It is null when nothing was asked about, or when the entry lies  outside what the link opens.
	EntityId NullableString `json:"entityId,omitempty"`
	// The title of that entry, null under the same conditions as its identifier.
	EntityTitle NullableString `json:"entityTitle,omitempty"`
	// Whether that entry is a folder or a file, null under the same conditions as its identifier.
	EntityType *FileEntryType `json:"entityType,omitempty"`
	// True when the link opens a whole room rather than one entry inside it. It is null for a link to a file and for  a link that did not resolve.
	IsRoom NullableBool `json:"isRoom,omitempty"`
	// True when the entry now sits in the calling account's own lists - it was already shared with that account, or  resolving the link has just put it there. It stays false for a visitor browsing without an account, who  reaches the entry through the link alone.
	Shared bool `json:"shared"`
	// The link the token belongs to, which is also the subject under which the link appears among the sharing rights  of the entry. It is an empty identifier when the link did not resolve.
	LinkId string `json:"linkId"`
	// Whether the request carried a signed-in account. It says nothing about that account's rights on the entry, so  it must not be read as permission - it is false for every anonymous visitor and true for any member, even one  who is a stranger to the room.
	IsAuthenticated bool `json:"isAuthenticated"`
	// Whether the signed-in caller already has rights of their own on the room that holds the entry, as opposed to  reaching it through this link. It is false for an anonymous visitor and for a member who has never been  invited.
	IsRoomMember *bool `json:"isRoomMember,omitempty"`
}

type _ExternalShareDto ExternalShareDto

// NewExternalShareDto instantiates a new ExternalShareDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalShareDto(status Status, tenantId int32, shared bool, linkId string, isAuthenticated bool) *ExternalShareDto {
	this := ExternalShareDto{}
	this.Status = status
	this.TenantId = tenantId
	this.Shared = shared
	this.LinkId = linkId
	this.IsAuthenticated = isAuthenticated
	return &this
}

// NewExternalShareDtoWithDefaults instantiates a new ExternalShareDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalShareDtoWithDefaults() *ExternalShareDto {
	this := ExternalShareDto{}
	return &this
}

// GetStatus returns the Status field value
func (o *ExternalShareDto) GetStatus() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetStatusOk() (*Status, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalShareDto) SetStatus(v Status) {
	o.Status = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ExternalShareDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ExternalShareDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ExternalShareDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ExternalShareDto) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ExternalShareDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ExternalShareDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ExternalShareDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ExternalShareDto) UnsetTitle() {
	o.Title.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExternalShareDto) GetType() FileEntryType {
	if o == nil || IsNil(o.Type) {
		var ret FileEntryType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExternalShareDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FileEntryType and assigns it to the Type field.
func (o *ExternalShareDto) SetType(v FileEntryType) {
	o.Type = &v
}

// GetTenantId returns the TenantId field value
func (o *ExternalShareDto) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ExternalShareDto) SetTenantId(v int32) {
	o.TenantId = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareDto) GetEntityId() string {
	if o == nil || IsNil(o.EntityId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareDto) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *ExternalShareDto) IsEntityIdSet() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *ExternalShareDto) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *ExternalShareDto) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *ExternalShareDto) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetEntityTitle returns the EntityTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareDto) GetEntityTitle() string {
	if o == nil || IsNil(o.EntityTitle.Get()) {
		var ret string
		return ret
	}
	return *o.EntityTitle.Get()
}

// GetEntityTitleOk returns a tuple with the EntityTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareDto) GetEntityTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityTitle.Get(), o.EntityTitle.IsSet()
}

// HasEntityTitle returns a boolean if a field has been set.
func (o *ExternalShareDto) IsEntityTitleSet() bool {
	if o != nil && o.EntityTitle.IsSet() {
		return true
	}

	return false
}

// SetEntityTitle gets a reference to the given NullableString and assigns it to the EntityTitle field.
func (o *ExternalShareDto) SetEntityTitle(v string) {
	o.EntityTitle.Set(&v)
}
// SetEntityTitleNil sets the value for EntityTitle to be an explicit nil
func (o *ExternalShareDto) SetEntityTitleNil() {
	o.EntityTitle.Set(nil)
}

// UnsetEntityTitle ensures that no value is present for EntityTitle, not even an explicit nil
func (o *ExternalShareDto) UnsetEntityTitle() {
	o.EntityTitle.Unset()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ExternalShareDto) GetEntityType() FileEntryType {
	if o == nil || IsNil(o.EntityType) {
		var ret FileEntryType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetEntityTypeOk() (*FileEntryType, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ExternalShareDto) IsEntityTypeSet() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given FileEntryType and assigns it to the EntityType field.
func (o *ExternalShareDto) SetEntityType(v FileEntryType) {
	o.EntityType = &v
}

// GetIsRoom returns the IsRoom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalShareDto) GetIsRoom() bool {
	if o == nil || IsNil(o.IsRoom.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRoom.Get()
}

// GetIsRoomOk returns a tuple with the IsRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalShareDto) GetIsRoomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRoom.Get(), o.IsRoom.IsSet()
}

// HasIsRoom returns a boolean if a field has been set.
func (o *ExternalShareDto) IsIsRoomSet() bool {
	if o != nil && o.IsRoom.IsSet() {
		return true
	}

	return false
}

// SetIsRoom gets a reference to the given NullableBool and assigns it to the IsRoom field.
func (o *ExternalShareDto) SetIsRoom(v bool) {
	o.IsRoom.Set(&v)
}
// SetIsRoomNil sets the value for IsRoom to be an explicit nil
func (o *ExternalShareDto) SetIsRoomNil() {
	o.IsRoom.Set(nil)
}

// UnsetIsRoom ensures that no value is present for IsRoom, not even an explicit nil
func (o *ExternalShareDto) UnsetIsRoom() {
	o.IsRoom.Unset()
}

// GetShared returns the Shared field value
func (o *ExternalShareDto) GetShared() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Shared
}

// GetSharedOk returns a tuple with the Shared field value
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetSharedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shared, true
}

// SetShared sets field value
func (o *ExternalShareDto) SetShared(v bool) {
	o.Shared = v
}

// GetLinkId returns the LinkId field value
func (o *ExternalShareDto) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *ExternalShareDto) SetLinkId(v string) {
	o.LinkId = v
}

// GetIsAuthenticated returns the IsAuthenticated field value
func (o *ExternalShareDto) GetIsAuthenticated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAuthenticated
}

// GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field value
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetIsAuthenticatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAuthenticated, true
}

// SetIsAuthenticated sets field value
func (o *ExternalShareDto) SetIsAuthenticated(v bool) {
	o.IsAuthenticated = v
}

// GetIsRoomMember returns the IsRoomMember field value if set, zero value otherwise.
func (o *ExternalShareDto) GetIsRoomMember() bool {
	if o == nil || IsNil(o.IsRoomMember) {
		var ret bool
		return ret
	}
	return *o.IsRoomMember
}

// GetIsRoomMemberOk returns a tuple with the IsRoomMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalShareDto) GetIsRoomMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoomMember) {
		return nil, false
	}
	return o.IsRoomMember, true
}

// HasIsRoomMember returns a boolean if a field has been set.
func (o *ExternalShareDto) IsIsRoomMemberSet() bool {
	if o != nil && !IsNil(o.IsRoomMember) {
		return true
	}

	return false
}

// SetIsRoomMember gets a reference to the given bool and assigns it to the IsRoomMember field.
func (o *ExternalShareDto) SetIsRoomMember(v bool) {
	o.IsRoomMember = &v
}

func (o ExternalShareDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalShareDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["tenantId"] = o.TenantId
	if o.EntityId.IsSet() {
		toSerialize["entityId"] = o.EntityId.Get()
	}
	if o.EntityTitle.IsSet() {
		toSerialize["entityTitle"] = o.EntityTitle.Get()
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if o.IsRoom.IsSet() {
		toSerialize["isRoom"] = o.IsRoom.Get()
	}
	toSerialize["shared"] = o.Shared
	toSerialize["linkId"] = o.LinkId
	toSerialize["isAuthenticated"] = o.IsAuthenticated
	if !IsNil(o.IsRoomMember) {
		toSerialize["isRoomMember"] = o.IsRoomMember
	}
	return toSerialize, nil
}

func (o *ExternalShareDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"tenantId",
		"shared",
		"linkId",
		"isAuthenticated",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExternalShareDto := _ExternalShareDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExternalShareDto)

	if err != nil {
		return err
	}

	*o = ExternalShareDto(varExternalShareDto)

	return err
}

type NullableExternalShareDto struct {
	value *ExternalShareDto
	isSet bool
}

func (v NullableExternalShareDto) Get() *ExternalShareDto {
	return v.value
}

func (v *NullableExternalShareDto) Set(val *ExternalShareDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalShareDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalShareDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalShareDto(val *ExternalShareDto) *NullableExternalShareDto {
	return &NullableExternalShareDto{value: val, isSet: true}
}

func (v NullableExternalShareDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalShareDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

