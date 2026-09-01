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

// ExternalShareDto The external sharing information and validation data.
type ExternalShareDto struct {
	// The external data status.
	Status Status `json:"status"`
	// The external data ID.
	Id NullableString `json:"id,omitempty"`
	// The external data title.
	Title NullableString `json:"title,omitempty"`
	// The type of the external data.
	Type *FileEntryType `json:"type,omitempty"`
	// The tenant ID.
	TenantId int32 `json:"tenantId"`
	// The unique identifier of the shared entity.
	EntityId NullableString `json:"entityId,omitempty"`
	// The title of the shared entity.
	EntityTitle NullableString `json:"entityTitle,omitempty"`
	// The entry type of the external data.
	EntityType *FileEntryType `json:"entityType,omitempty"`
	// Indicates whether the entity represents a room.
	IsRoom NullableBool `json:"isRoom,omitempty"`
	// Specifies whether to share the external data or not.
	Shared bool `json:"shared"`
	// The link ID of the external data.
	LinkId string `json:"linkId"`
	// Specifies whether the user is authenticated or not.
	IsAuthenticated bool `json:"isAuthenticated"`
	// The room ID of the external data.
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

