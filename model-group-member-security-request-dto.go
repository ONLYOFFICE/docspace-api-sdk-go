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

// checks if the GroupMemberSecurityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMemberSecurityRequestDto{}

// GroupMemberSecurityRequestDto One member of a portal group together with the access that member has on the file or folder the group was granted  rights to. Every line of the answer describes the same file or folder and differs only in the member and in the  level that applies to them.
type GroupMemberSecurityRequestDto struct {
	// The member the line is about, as the portal reports the account: the display name, the avatar and the portal  role to show next to the access level.
	User EmployeeFullDto `json:"user"`
	// The level granted to the group as a whole on this file or folder. It belongs to the group record rather than  to the member, so the same value repeats on every line of the answer; a group whose record was set back to  none is answered with an empty list instead.
	GroupAccess FileShare `json:"groupAccess"`
	// The level granted to this member alone on the same file or folder, or `null` when the member has no record of  their own and the group level is what applies. The member who created the file or folder is always reported  here as a room manager, whatever their own record says.
	UserAccess *FileShare `json:"userAccess,omitempty"`
	// Whether `userAccess` is the level that decides what the member may do. When it is false the member inherits  `groupAccess`, and the creator of the file or folder is always reported as overridden because of the room  manager level forced onto them.
	Overridden bool `json:"overridden"`
	// Whether the caller may still change the level of this member. It comes back false on the line of the member  who created the file or folder, on the line of the caller themselves, and on every line at once when the  caller may read the file or folder but not manage access to it.
	CanEditAccess bool `json:"canEditAccess"`
	// Whether this member created the file or folder - the owner of the entry, not the owner of the group. Their  level is reported as a room manager one and cannot be taken away through this group.
	Owner bool `json:"owner"`
}

type _GroupMemberSecurityRequestDto GroupMemberSecurityRequestDto

// NewGroupMemberSecurityRequestDto instantiates a new GroupMemberSecurityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMemberSecurityRequestDto(user EmployeeFullDto, groupAccess FileShare, overridden bool, canEditAccess bool, owner bool) *GroupMemberSecurityRequestDto {
	this := GroupMemberSecurityRequestDto{}
	this.User = user
	this.GroupAccess = groupAccess
	this.Overridden = overridden
	this.CanEditAccess = canEditAccess
	this.Owner = owner
	return &this
}

// NewGroupMemberSecurityRequestDtoWithDefaults instantiates a new GroupMemberSecurityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMemberSecurityRequestDtoWithDefaults() *GroupMemberSecurityRequestDto {
	this := GroupMemberSecurityRequestDto{}
	return &this
}

// GetUser returns the User field value
func (o *GroupMemberSecurityRequestDto) GetUser() EmployeeFullDto {
	if o == nil {
		var ret EmployeeFullDto
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetUserOk() (*EmployeeFullDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GroupMemberSecurityRequestDto) SetUser(v EmployeeFullDto) {
	o.User = v
}

// GetGroupAccess returns the GroupAccess field value
func (o *GroupMemberSecurityRequestDto) GetGroupAccess() FileShare {
	if o == nil {
		var ret FileShare
		return ret
	}

	return o.GroupAccess
}

// GetGroupAccessOk returns a tuple with the GroupAccess field value
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetGroupAccessOk() (*FileShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupAccess, true
}

// SetGroupAccess sets field value
func (o *GroupMemberSecurityRequestDto) SetGroupAccess(v FileShare) {
	o.GroupAccess = v
}

// GetUserAccess returns the UserAccess field value if set, zero value otherwise.
func (o *GroupMemberSecurityRequestDto) GetUserAccess() FileShare {
	if o == nil || IsNil(o.UserAccess) {
		var ret FileShare
		return ret
	}
	return *o.UserAccess
}

// GetUserAccessOk returns a tuple with the UserAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetUserAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.UserAccess) {
		return nil, false
	}
	return o.UserAccess, true
}

// HasUserAccess returns a boolean if a field has been set.
func (o *GroupMemberSecurityRequestDto) IsUserAccessSet() bool {
	if o != nil && !IsNil(o.UserAccess) {
		return true
	}

	return false
}

// SetUserAccess gets a reference to the given FileShare and assigns it to the UserAccess field.
func (o *GroupMemberSecurityRequestDto) SetUserAccess(v FileShare) {
	o.UserAccess = &v
}

// GetOverridden returns the Overridden field value
func (o *GroupMemberSecurityRequestDto) GetOverridden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Overridden
}

// GetOverriddenOk returns a tuple with the Overridden field value
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetOverriddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overridden, true
}

// SetOverridden sets field value
func (o *GroupMemberSecurityRequestDto) SetOverridden(v bool) {
	o.Overridden = v
}

// GetCanEditAccess returns the CanEditAccess field value
func (o *GroupMemberSecurityRequestDto) GetCanEditAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditAccess
}

// GetCanEditAccessOk returns a tuple with the CanEditAccess field value
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetCanEditAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditAccess, true
}

// SetCanEditAccess sets field value
func (o *GroupMemberSecurityRequestDto) SetCanEditAccess(v bool) {
	o.CanEditAccess = v
}

// GetOwner returns the Owner field value
func (o *GroupMemberSecurityRequestDto) GetOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *GroupMemberSecurityRequestDto) GetOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *GroupMemberSecurityRequestDto) SetOwner(v bool) {
	o.Owner = v
}

func (o GroupMemberSecurityRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMemberSecurityRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["groupAccess"] = o.GroupAccess
	if !IsNil(o.UserAccess) {
		toSerialize["userAccess"] = o.UserAccess
	}
	toSerialize["overridden"] = o.Overridden
	toSerialize["canEditAccess"] = o.CanEditAccess
	toSerialize["owner"] = o.Owner
	return toSerialize, nil
}

func (o *GroupMemberSecurityRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"groupAccess",
		"overridden",
		"canEditAccess",
		"owner",
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

	varGroupMemberSecurityRequestDto := _GroupMemberSecurityRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGroupMemberSecurityRequestDto)

	if err != nil {
		return err
	}

	*o = GroupMemberSecurityRequestDto(varGroupMemberSecurityRequestDto)

	return err
}

type NullableGroupMemberSecurityRequestDto struct {
	value *GroupMemberSecurityRequestDto
	isSet bool
}

func (v NullableGroupMemberSecurityRequestDto) Get() *GroupMemberSecurityRequestDto {
	return v.value
}

func (v *NullableGroupMemberSecurityRequestDto) Set(val *GroupMemberSecurityRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMemberSecurityRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMemberSecurityRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMemberSecurityRequestDto(val *GroupMemberSecurityRequestDto) *NullableGroupMemberSecurityRequestDto {
	return &NullableGroupMemberSecurityRequestDto{value: val, isSet: true}
}

func (v NullableGroupMemberSecurityRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMemberSecurityRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

