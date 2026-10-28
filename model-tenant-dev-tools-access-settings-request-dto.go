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

// checks if the TenantDevToolsAccessSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantDevToolsAccessSettingsRequestDto{}

// TenantDevToolsAccessSettingsRequestDto Whether the `User` role is barred from the portal developer tools.
type TenantDevToolsAccessSettingsRequestDto struct {
	// Whether members holding the `User` role are barred from the developer tools - API keys, OAuth applications  and webhooks. Room administrators and DocSpace administrators keep their access either way.
	LimitedAccessForUsers *bool `json:"limitedAccessForUsers,omitempty"`
}

// NewTenantDevToolsAccessSettingsRequestDto instantiates a new TenantDevToolsAccessSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantDevToolsAccessSettingsRequestDto() *TenantDevToolsAccessSettingsRequestDto {
	this := TenantDevToolsAccessSettingsRequestDto{}
	return &this
}

// NewTenantDevToolsAccessSettingsRequestDtoWithDefaults instantiates a new TenantDevToolsAccessSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDevToolsAccessSettingsRequestDtoWithDefaults() *TenantDevToolsAccessSettingsRequestDto {
	this := TenantDevToolsAccessSettingsRequestDto{}
	return &this
}

// GetLimitedAccessForUsers returns the LimitedAccessForUsers field value if set, zero value otherwise.
func (o *TenantDevToolsAccessSettingsRequestDto) GetLimitedAccessForUsers() bool {
	if o == nil || IsNil(o.LimitedAccessForUsers) {
		var ret bool
		return ret
	}
	return *o.LimitedAccessForUsers
}

// GetLimitedAccessForUsersOk returns a tuple with the LimitedAccessForUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDevToolsAccessSettingsRequestDto) GetLimitedAccessForUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitedAccessForUsers) {
		return nil, false
	}
	return o.LimitedAccessForUsers, true
}

// HasLimitedAccessForUsers returns a boolean if a field has been set.
func (o *TenantDevToolsAccessSettingsRequestDto) IsLimitedAccessForUsersSet() bool {
	if o != nil && !IsNil(o.LimitedAccessForUsers) {
		return true
	}

	return false
}

// SetLimitedAccessForUsers gets a reference to the given bool and assigns it to the LimitedAccessForUsers field.
func (o *TenantDevToolsAccessSettingsRequestDto) SetLimitedAccessForUsers(v bool) {
	o.LimitedAccessForUsers = &v
}

func (o TenantDevToolsAccessSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantDevToolsAccessSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitedAccessForUsers) {
		toSerialize["limitedAccessForUsers"] = o.LimitedAccessForUsers
	}
	return toSerialize, nil
}

type NullableTenantDevToolsAccessSettingsRequestDto struct {
	value *TenantDevToolsAccessSettingsRequestDto
	isSet bool
}

func (v NullableTenantDevToolsAccessSettingsRequestDto) Get() *TenantDevToolsAccessSettingsRequestDto {
	return v.value
}

func (v *NullableTenantDevToolsAccessSettingsRequestDto) Set(val *TenantDevToolsAccessSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantDevToolsAccessSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantDevToolsAccessSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantDevToolsAccessSettingsRequestDto(val *TenantDevToolsAccessSettingsRequestDto) *NullableTenantDevToolsAccessSettingsRequestDto {
	return &NullableTenantDevToolsAccessSettingsRequestDto{value: val, isSet: true}
}

func (v NullableTenantDevToolsAccessSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantDevToolsAccessSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

