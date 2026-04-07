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

// checks if the GroupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRequestDto{}

// GroupRequestDto The group request parameters.
type GroupRequestDto struct {
	// The list of group member IDs.
	Members []string `json:"members,omitempty"`
	// The group manager ID.
	GroupManager string `json:"groupManager"`
	// The group name.
	GroupName NullableString `json:"groupName,omitempty"`
}

type _GroupRequestDto GroupRequestDto

// NewGroupRequestDto instantiates a new GroupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRequestDto(groupManager string) *GroupRequestDto {
	this := GroupRequestDto{}
	this.GroupManager = groupManager
	return &this
}

// NewGroupRequestDtoWithDefaults instantiates a new GroupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRequestDtoWithDefaults() *GroupRequestDto {
	this := GroupRequestDto{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupRequestDto) GetMembers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRequestDto) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupRequestDto) IsMembersSet() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *GroupRequestDto) SetMembers(v []string) {
	o.Members = v
}

// GetGroupManager returns the GroupManager field value
func (o *GroupRequestDto) GetGroupManager() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupManager
}

// GetGroupManagerOk returns a tuple with the GroupManager field value
// and a boolean to check if the value has been set.
func (o *GroupRequestDto) GetGroupManagerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupManager, true
}

// SetGroupManager sets field value
func (o *GroupRequestDto) SetGroupManager(v string) {
	o.GroupManager = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupRequestDto) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRequestDto) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *GroupRequestDto) IsGroupNameSet() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *GroupRequestDto) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *GroupRequestDto) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *GroupRequestDto) UnsetGroupName() {
	o.GroupName.Unset()
}

func (o GroupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	toSerialize["groupManager"] = o.GroupManager
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	return toSerialize, nil
}

func (o *GroupRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupManager",
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

	varGroupRequestDto := _GroupRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGroupRequestDto)

	if err != nil {
		return err
	}

	*o = GroupRequestDto(varGroupRequestDto)

	return err
}

type NullableGroupRequestDto struct {
	value *GroupRequestDto
	isSet bool
}

func (v NullableGroupRequestDto) Get() *GroupRequestDto {
	return v.value
}

func (v *NullableGroupRequestDto) Set(val *GroupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRequestDto(val *GroupRequestDto) *NullableGroupRequestDto {
	return &NullableGroupRequestDto{value: val, isSet: true}
}

func (v NullableGroupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

