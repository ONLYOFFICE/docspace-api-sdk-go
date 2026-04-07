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

// checks if the MigratingApiGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigratingApiGroup{}

// MigratingApiGroup struct for MigratingApiGroup
type MigratingApiGroup struct {
	ShouldImport *bool `json:"shouldImport,omitempty"`
	GroupName NullableString `json:"groupName,omitempty"`
	ModuleName NullableString `json:"moduleName,omitempty"`
	UserUidList []string `json:"userUidList,omitempty"`
}

// NewMigratingApiGroup instantiates a new MigratingApiGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigratingApiGroup() *MigratingApiGroup {
	this := MigratingApiGroup{}
	return &this
}

// NewMigratingApiGroupWithDefaults instantiates a new MigratingApiGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigratingApiGroupWithDefaults() *MigratingApiGroup {
	this := MigratingApiGroup{}
	return &this
}

// GetShouldImport returns the ShouldImport field value if set, zero value otherwise.
func (o *MigratingApiGroup) GetShouldImport() bool {
	if o == nil || IsNil(o.ShouldImport) {
		var ret bool
		return ret
	}
	return *o.ShouldImport
}

// GetShouldImportOk returns a tuple with the ShouldImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiGroup) GetShouldImportOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldImport) {
		return nil, false
	}
	return o.ShouldImport, true
}

// HasShouldImport returns a boolean if a field has been set.
func (o *MigratingApiGroup) IsShouldImportSet() bool {
	if o != nil && !IsNil(o.ShouldImport) {
		return true
	}

	return false
}

// SetShouldImport gets a reference to the given bool and assigns it to the ShouldImport field.
func (o *MigratingApiGroup) SetShouldImport(v bool) {
	o.ShouldImport = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiGroup) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiGroup) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *MigratingApiGroup) IsGroupNameSet() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *MigratingApiGroup) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *MigratingApiGroup) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *MigratingApiGroup) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetModuleName returns the ModuleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiGroup) GetModuleName() string {
	if o == nil || IsNil(o.ModuleName.Get()) {
		var ret string
		return ret
	}
	return *o.ModuleName.Get()
}

// GetModuleNameOk returns a tuple with the ModuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiGroup) GetModuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleName.Get(), o.ModuleName.IsSet()
}

// HasModuleName returns a boolean if a field has been set.
func (o *MigratingApiGroup) IsModuleNameSet() bool {
	if o != nil && o.ModuleName.IsSet() {
		return true
	}

	return false
}

// SetModuleName gets a reference to the given NullableString and assigns it to the ModuleName field.
func (o *MigratingApiGroup) SetModuleName(v string) {
	o.ModuleName.Set(&v)
}
// SetModuleNameNil sets the value for ModuleName to be an explicit nil
func (o *MigratingApiGroup) SetModuleNameNil() {
	o.ModuleName.Set(nil)
}

// UnsetModuleName ensures that no value is present for ModuleName, not even an explicit nil
func (o *MigratingApiGroup) UnsetModuleName() {
	o.ModuleName.Unset()
}

// GetUserUidList returns the UserUidList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MigratingApiGroup) GetUserUidList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserUidList
}

// GetUserUidListOk returns a tuple with the UserUidList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MigratingApiGroup) GetUserUidListOk() ([]string, bool) {
	if o == nil || IsNil(o.UserUidList) {
		return nil, false
	}
	return o.UserUidList, true
}

// HasUserUidList returns a boolean if a field has been set.
func (o *MigratingApiGroup) IsUserUidListSet() bool {
	if o != nil && !IsNil(o.UserUidList) {
		return true
	}

	return false
}

// SetUserUidList gets a reference to the given []string and assigns it to the UserUidList field.
func (o *MigratingApiGroup) SetUserUidList(v []string) {
	o.UserUidList = v
}

func (o MigratingApiGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigratingApiGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShouldImport) {
		toSerialize["shouldImport"] = o.ShouldImport
	}
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	if o.ModuleName.IsSet() {
		toSerialize["moduleName"] = o.ModuleName.Get()
	}
	if o.UserUidList != nil {
		toSerialize["userUidList"] = o.UserUidList
	}
	return toSerialize, nil
}

type NullableMigratingApiGroup struct {
	value *MigratingApiGroup
	isSet bool
}

func (v NullableMigratingApiGroup) Get() *MigratingApiGroup {
	return v.value
}

func (v *NullableMigratingApiGroup) Set(val *MigratingApiGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMigratingApiGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMigratingApiGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigratingApiGroup(val *MigratingApiGroup) *NullableMigratingApiGroup {
	return &NullableMigratingApiGroup{value: val, isSet: true}
}

func (v NullableMigratingApiGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigratingApiGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

