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

// checks if the GroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupDto{}

// GroupDto The group parameters.
type GroupDto struct {
	// The group name.
	Name NullableString `json:"name"`
	// The parent group ID.
	Parent NullableString `json:"parent,omitempty"`
	// The group category ID.
	Category string `json:"category"`
	// The group ID.
	Id string `json:"id"`
	// Specifies if the LDAP settings are enabled for the group or not.
	IsLDAP bool `json:"isLDAP"`
	// Indicates whether the group is a system group.
	IsSystem NullableBool `json:"isSystem,omitempty"`
	Manager *EmployeeFullDto `json:"manager,omitempty"`
	// The list of group members.
	Members []EmployeeFullDto `json:"members,omitempty"`
	// Specifies whether the group can be shared or not.
	Shared NullableBool `json:"shared,omitempty"`
	// The number of group members.
	MembersCount *int32 `json:"membersCount,omitempty"`
}

type _GroupDto GroupDto

// NewGroupDto instantiates a new GroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDto(name NullableString, category string, id string, isLDAP bool) *GroupDto {
	this := GroupDto{}
	this.Name = name
	this.Category = category
	this.Id = id
	this.IsLDAP = isLDAP
	return &this
}

// NewGroupDtoWithDefaults instantiates a new GroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDtoWithDefaults() *GroupDto {
	this := GroupDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *GroupDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDto) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDto) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *GroupDto) IsParentSet() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *GroupDto) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *GroupDto) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *GroupDto) UnsetParent() {
	o.Parent.Unset()
}

// GetCategory returns the Category field value
func (o *GroupDto) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GroupDto) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *GroupDto) SetCategory(v string) {
	o.Category = v
}

// GetId returns the Id field value
func (o *GroupDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupDto) SetId(v string) {
	o.Id = v
}

// GetIsLDAP returns the IsLDAP field value
func (o *GroupDto) GetIsLDAP() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLDAP
}

// GetIsLDAPOk returns a tuple with the IsLDAP field value
// and a boolean to check if the value has been set.
func (o *GroupDto) GetIsLDAPOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLDAP, true
}

// SetIsLDAP sets field value
func (o *GroupDto) SetIsLDAP(v bool) {
	o.IsLDAP = v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDto) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSystem.Get()
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDto) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSystem.Get(), o.IsSystem.IsSet()
}

// HasIsSystem returns a boolean if a field has been set.
func (o *GroupDto) IsIsSystemSet() bool {
	if o != nil && o.IsSystem.IsSet() {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given NullableBool and assigns it to the IsSystem field.
func (o *GroupDto) SetIsSystem(v bool) {
	o.IsSystem.Set(&v)
}
// SetIsSystemNil sets the value for IsSystem to be an explicit nil
func (o *GroupDto) SetIsSystemNil() {
	o.IsSystem.Set(nil)
}

// UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
func (o *GroupDto) UnsetIsSystem() {
	o.IsSystem.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *GroupDto) GetManager() EmployeeFullDto {
	if o == nil || IsNil(o.Manager) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDto) GetManagerOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *GroupDto) IsManagerSet() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given EmployeeFullDto and assigns it to the Manager field.
func (o *GroupDto) SetManager(v EmployeeFullDto) {
	o.Manager = &v
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDto) GetMembers() []EmployeeFullDto {
	if o == nil {
		var ret []EmployeeFullDto
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDto) GetMembersOk() ([]EmployeeFullDto, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupDto) IsMembersSet() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []EmployeeFullDto and assigns it to the Members field.
func (o *GroupDto) SetMembers(v []EmployeeFullDto) {
	o.Members = v
}

// GetShared returns the Shared field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupDto) GetShared() bool {
	if o == nil || IsNil(o.Shared.Get()) {
		var ret bool
		return ret
	}
	return *o.Shared.Get()
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupDto) GetSharedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shared.Get(), o.Shared.IsSet()
}

// HasShared returns a boolean if a field has been set.
func (o *GroupDto) IsSharedSet() bool {
	if o != nil && o.Shared.IsSet() {
		return true
	}

	return false
}

// SetShared gets a reference to the given NullableBool and assigns it to the Shared field.
func (o *GroupDto) SetShared(v bool) {
	o.Shared.Set(&v)
}
// SetSharedNil sets the value for Shared to be an explicit nil
func (o *GroupDto) SetSharedNil() {
	o.Shared.Set(nil)
}

// UnsetShared ensures that no value is present for Shared, not even an explicit nil
func (o *GroupDto) UnsetShared() {
	o.Shared.Unset()
}

// GetMembersCount returns the MembersCount field value if set, zero value otherwise.
func (o *GroupDto) GetMembersCount() int32 {
	if o == nil || IsNil(o.MembersCount) {
		var ret int32
		return ret
	}
	return *o.MembersCount
}

// GetMembersCountOk returns a tuple with the MembersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDto) GetMembersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MembersCount) {
		return nil, false
	}
	return o.MembersCount, true
}

// HasMembersCount returns a boolean if a field has been set.
func (o *GroupDto) IsMembersCountSet() bool {
	if o != nil && !IsNil(o.MembersCount) {
		return true
	}

	return false
}

// SetMembersCount gets a reference to the given int32 and assigns it to the MembersCount field.
func (o *GroupDto) SetMembersCount(v int32) {
	o.MembersCount = &v
}

func (o GroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	toSerialize["category"] = o.Category
	toSerialize["id"] = o.Id
	toSerialize["isLDAP"] = o.IsLDAP
	if o.IsSystem.IsSet() {
		toSerialize["isSystem"] = o.IsSystem.Get()
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Shared.IsSet() {
		toSerialize["shared"] = o.Shared.Get()
	}
	if !IsNil(o.MembersCount) {
		toSerialize["membersCount"] = o.MembersCount
	}
	return toSerialize, nil
}

func (o *GroupDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"category",
		"id",
		"isLDAP",
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

	varGroupDto := _GroupDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGroupDto)

	if err != nil {
		return err
	}

	*o = GroupDto(varGroupDto)

	return err
}

type NullableGroupDto struct {
	value *GroupDto
	isSet bool
}

func (v NullableGroupDto) Get() *GroupDto {
	return v.value
}

func (v *NullableGroupDto) Set(val *GroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDto(val *GroupDto) *NullableGroupDto {
	return &NullableGroupDto{value: val, isSet: true}
}

func (v NullableGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

