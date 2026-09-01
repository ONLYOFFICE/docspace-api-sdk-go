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

// checks if the SecurityDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityDto{}

// SecurityDto The security information.
type SecurityDto struct {
	// The module ID.
	WebItemId NullableString `json:"webItemId,omitempty"`
	// The list of users with the access to the module.
	Users []EmployeeDto `json:"users,omitempty"`
	// The list of groups with the access to the module.
	Groups []GroupSummaryDto `json:"groups,omitempty"`
	// Specifies if the security settings are enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies if the module is a subitem or not.
	IsSubItem *bool `json:"isSubItem,omitempty"`
}

// NewSecurityDto instantiates a new SecurityDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityDto() *SecurityDto {
	this := SecurityDto{}
	return &this
}

// NewSecurityDtoWithDefaults instantiates a new SecurityDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityDtoWithDefaults() *SecurityDto {
	this := SecurityDto{}
	return &this
}

// GetWebItemId returns the WebItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityDto) GetWebItemId() string {
	if o == nil || IsNil(o.WebItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WebItemId.Get()
}

// GetWebItemIdOk returns a tuple with the WebItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityDto) GetWebItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebItemId.Get(), o.WebItemId.IsSet()
}

// HasWebItemId returns a boolean if a field has been set.
func (o *SecurityDto) IsWebItemIdSet() bool {
	if o != nil && o.WebItemId.IsSet() {
		return true
	}

	return false
}

// SetWebItemId gets a reference to the given NullableString and assigns it to the WebItemId field.
func (o *SecurityDto) SetWebItemId(v string) {
	o.WebItemId.Set(&v)
}
// SetWebItemIdNil sets the value for WebItemId to be an explicit nil
func (o *SecurityDto) SetWebItemIdNil() {
	o.WebItemId.Set(nil)
}

// UnsetWebItemId ensures that no value is present for WebItemId, not even an explicit nil
func (o *SecurityDto) UnsetWebItemId() {
	o.WebItemId.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityDto) GetUsers() []EmployeeDto {
	if o == nil {
		var ret []EmployeeDto
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityDto) GetUsersOk() ([]EmployeeDto, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SecurityDto) IsUsersSet() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []EmployeeDto and assigns it to the Users field.
func (o *SecurityDto) SetUsers(v []EmployeeDto) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityDto) GetGroups() []GroupSummaryDto {
	if o == nil {
		var ret []GroupSummaryDto
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityDto) GetGroupsOk() ([]GroupSummaryDto, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SecurityDto) IsGroupsSet() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSummaryDto and assigns it to the Groups field.
func (o *SecurityDto) SetGroups(v []GroupSummaryDto) {
	o.Groups = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityDto) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsSubItem returns the IsSubItem field value if set, zero value otherwise.
func (o *SecurityDto) GetIsSubItem() bool {
	if o == nil || IsNil(o.IsSubItem) {
		var ret bool
		return ret
	}
	return *o.IsSubItem
}

// GetIsSubItemOk returns a tuple with the IsSubItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityDto) GetIsSubItemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubItem) {
		return nil, false
	}
	return o.IsSubItem, true
}

// HasIsSubItem returns a boolean if a field has been set.
func (o *SecurityDto) IsIsSubItemSet() bool {
	if o != nil && !IsNil(o.IsSubItem) {
		return true
	}

	return false
}

// SetIsSubItem gets a reference to the given bool and assigns it to the IsSubItem field.
func (o *SecurityDto) SetIsSubItem(v bool) {
	o.IsSubItem = &v
}

func (o SecurityDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebItemId.IsSet() {
		toSerialize["webItemId"] = o.WebItemId.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IsSubItem) {
		toSerialize["isSubItem"] = o.IsSubItem
	}
	return toSerialize, nil
}

type NullableSecurityDto struct {
	value *SecurityDto
	isSet bool
}

func (v NullableSecurityDto) Get() *SecurityDto {
	return v.value
}

func (v *NullableSecurityDto) Set(val *SecurityDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityDto(val *SecurityDto) *NullableSecurityDto {
	return &NullableSecurityDto{value: val, isSet: true}
}

func (v NullableSecurityDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

