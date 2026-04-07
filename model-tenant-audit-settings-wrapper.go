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

// checks if the TenantAuditSettingsWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAuditSettingsWrapper{}

// TenantAuditSettingsWrapper The tenant audit settings wrapper.
type TenantAuditSettingsWrapper struct {
	Settings *TenantAuditSettings `json:"settings,omitempty"`
}

// NewTenantAuditSettingsWrapper instantiates a new TenantAuditSettingsWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAuditSettingsWrapper() *TenantAuditSettingsWrapper {
	this := TenantAuditSettingsWrapper{}
	return &this
}

// NewTenantAuditSettingsWrapperWithDefaults instantiates a new TenantAuditSettingsWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAuditSettingsWrapperWithDefaults() *TenantAuditSettingsWrapper {
	this := TenantAuditSettingsWrapper{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TenantAuditSettingsWrapper) GetSettings() TenantAuditSettings {
	if o == nil || IsNil(o.Settings) {
		var ret TenantAuditSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAuditSettingsWrapper) GetSettingsOk() (*TenantAuditSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TenantAuditSettingsWrapper) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given TenantAuditSettings and assigns it to the Settings field.
func (o *TenantAuditSettingsWrapper) SetSettings(v TenantAuditSettings) {
	o.Settings = &v
}

func (o TenantAuditSettingsWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAuditSettingsWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableTenantAuditSettingsWrapper struct {
	value *TenantAuditSettingsWrapper
	isSet bool
}

func (v NullableTenantAuditSettingsWrapper) Get() *TenantAuditSettingsWrapper {
	return v.value
}

func (v *NullableTenantAuditSettingsWrapper) Set(val *TenantAuditSettingsWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAuditSettingsWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAuditSettingsWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAuditSettingsWrapper(val *TenantAuditSettingsWrapper) *NullableTenantAuditSettingsWrapper {
	return &NullableTenantAuditSettingsWrapper{value: val, isSet: true}
}

func (v NullableTenantAuditSettingsWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAuditSettingsWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

