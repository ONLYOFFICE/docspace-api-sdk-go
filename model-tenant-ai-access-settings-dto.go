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

// checks if the TenantAiAccessSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAiAccessSettingsDto{}

// TenantAiAccessSettingsDto Whether AI functionality is available on the portal.
type TenantAiAccessSettingsDto struct {
	// Whether AI is available on the portal at all - chat, agents and vectorization together. Switching it off  hides the AI Agents folder and makes every AI endpoint unreachable for all members at once, not only for the  caller, and the change is pushed to connected clients rather than waiting for their next request.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewTenantAiAccessSettingsDto instantiates a new TenantAiAccessSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAiAccessSettingsDto() *TenantAiAccessSettingsDto {
	this := TenantAiAccessSettingsDto{}
	return &this
}

// NewTenantAiAccessSettingsDtoWithDefaults instantiates a new TenantAiAccessSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAiAccessSettingsDtoWithDefaults() *TenantAiAccessSettingsDto {
	this := TenantAiAccessSettingsDto{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TenantAiAccessSettingsDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAccessSettingsDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenantAiAccessSettingsDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TenantAiAccessSettingsDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o TenantAiAccessSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAiAccessSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableTenantAiAccessSettingsDto struct {
	value *TenantAiAccessSettingsDto
	isSet bool
}

func (v NullableTenantAiAccessSettingsDto) Get() *TenantAiAccessSettingsDto {
	return v.value
}

func (v *NullableTenantAiAccessSettingsDto) Set(val *TenantAiAccessSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAiAccessSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAiAccessSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAiAccessSettingsDto(val *TenantAiAccessSettingsDto) *NullableTenantAiAccessSettingsDto {
	return &NullableTenantAiAccessSettingsDto{value: val, isSet: true}
}

func (v NullableTenantAiAccessSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAiAccessSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

