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

// checks if the UpdateGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupRequest{}

// UpdateGroupRequest The request for updating a group.
type UpdateGroupRequest struct {
	// The list of user IDs to add to the group.
	MembersToAdd []string `json:"membersToAdd,omitempty"`
	// The list of user IDs to remove from the group.
	MembersToRemove []string `json:"membersToRemove,omitempty"`
	// The group manager ID.
	GroupManager *string `json:"groupManager,omitempty"`
	// The group name.
	GroupName NullableString `json:"groupName,omitempty"`
}

// NewUpdateGroupRequest instantiates a new UpdateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupRequest() *UpdateGroupRequest {
	this := UpdateGroupRequest{}
	return &this
}

// NewUpdateGroupRequestWithDefaults instantiates a new UpdateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupRequestWithDefaults() *UpdateGroupRequest {
	this := UpdateGroupRequest{}
	return &this
}

// GetMembersToAdd returns the MembersToAdd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGroupRequest) GetMembersToAdd() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MembersToAdd
}

// GetMembersToAddOk returns a tuple with the MembersToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGroupRequest) GetMembersToAddOk() ([]string, bool) {
	if o == nil || IsNil(o.MembersToAdd) {
		return nil, false
	}
	return o.MembersToAdd, true
}

// HasMembersToAdd returns a boolean if a field has been set.
func (o *UpdateGroupRequest) IsMembersToAddSet() bool {
	if o != nil && !IsNil(o.MembersToAdd) {
		return true
	}

	return false
}

// SetMembersToAdd gets a reference to the given []string and assigns it to the MembersToAdd field.
func (o *UpdateGroupRequest) SetMembersToAdd(v []string) {
	o.MembersToAdd = v
}

// GetMembersToRemove returns the MembersToRemove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGroupRequest) GetMembersToRemove() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MembersToRemove
}

// GetMembersToRemoveOk returns a tuple with the MembersToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGroupRequest) GetMembersToRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.MembersToRemove) {
		return nil, false
	}
	return o.MembersToRemove, true
}

// HasMembersToRemove returns a boolean if a field has been set.
func (o *UpdateGroupRequest) IsMembersToRemoveSet() bool {
	if o != nil && !IsNil(o.MembersToRemove) {
		return true
	}

	return false
}

// SetMembersToRemove gets a reference to the given []string and assigns it to the MembersToRemove field.
func (o *UpdateGroupRequest) SetMembersToRemove(v []string) {
	o.MembersToRemove = v
}

// GetGroupManager returns the GroupManager field value if set, zero value otherwise.
func (o *UpdateGroupRequest) GetGroupManager() string {
	if o == nil || IsNil(o.GroupManager) {
		var ret string
		return ret
	}
	return *o.GroupManager
}

// GetGroupManagerOk returns a tuple with the GroupManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGroupRequest) GetGroupManagerOk() (*string, bool) {
	if o == nil || IsNil(o.GroupManager) {
		return nil, false
	}
	return o.GroupManager, true
}

// HasGroupManager returns a boolean if a field has been set.
func (o *UpdateGroupRequest) IsGroupManagerSet() bool {
	if o != nil && !IsNil(o.GroupManager) {
		return true
	}

	return false
}

// SetGroupManager gets a reference to the given string and assigns it to the GroupManager field.
func (o *UpdateGroupRequest) SetGroupManager(v string) {
	o.GroupManager = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGroupRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGroupRequest) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *UpdateGroupRequest) IsGroupNameSet() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *UpdateGroupRequest) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *UpdateGroupRequest) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *UpdateGroupRequest) UnsetGroupName() {
	o.GroupName.Unset()
}

func (o UpdateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MembersToAdd != nil {
		toSerialize["membersToAdd"] = o.MembersToAdd
	}
	if o.MembersToRemove != nil {
		toSerialize["membersToRemove"] = o.MembersToRemove
	}
	if !IsNil(o.GroupManager) {
		toSerialize["groupManager"] = o.GroupManager
	}
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	return toSerialize, nil
}

type NullableUpdateGroupRequest struct {
	value *UpdateGroupRequest
	isSet bool
}

func (v NullableUpdateGroupRequest) Get() *UpdateGroupRequest {
	return v.value
}

func (v *NullableUpdateGroupRequest) Set(val *UpdateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupRequest(val *UpdateGroupRequest) *NullableUpdateGroupRequest {
	return &NullableUpdateGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

