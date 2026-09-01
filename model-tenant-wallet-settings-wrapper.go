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

// checks if the TenantWalletSettingsWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantWalletSettingsWrapper{}

// TenantWalletSettingsWrapper The wrapper for the tenant wallet settings.
type TenantWalletSettingsWrapper struct {
	// The tenant wallet settings.
	Settings *TenantWalletSettings `json:"settings,omitempty"`
}

// NewTenantWalletSettingsWrapper instantiates a new TenantWalletSettingsWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantWalletSettingsWrapper() *TenantWalletSettingsWrapper {
	this := TenantWalletSettingsWrapper{}
	return &this
}

// NewTenantWalletSettingsWrapperWithDefaults instantiates a new TenantWalletSettingsWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWalletSettingsWrapperWithDefaults() *TenantWalletSettingsWrapper {
	this := TenantWalletSettingsWrapper{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TenantWalletSettingsWrapper) GetSettings() TenantWalletSettings {
	if o == nil || IsNil(o.Settings) {
		var ret TenantWalletSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettingsWrapper) GetSettingsOk() (*TenantWalletSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TenantWalletSettingsWrapper) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given TenantWalletSettings and assigns it to the Settings field.
func (o *TenantWalletSettingsWrapper) SetSettings(v TenantWalletSettings) {
	o.Settings = &v
}

func (o TenantWalletSettingsWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantWalletSettingsWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableTenantWalletSettingsWrapper struct {
	value *TenantWalletSettingsWrapper
	isSet bool
}

func (v NullableTenantWalletSettingsWrapper) Get() *TenantWalletSettingsWrapper {
	return v.value
}

func (v *NullableTenantWalletSettingsWrapper) Set(val *TenantWalletSettingsWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWalletSettingsWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWalletSettingsWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWalletSettingsWrapper(val *TenantWalletSettingsWrapper) *NullableTenantWalletSettingsWrapper {
	return &NullableTenantWalletSettingsWrapper{value: val, isSet: true}
}

func (v NullableTenantWalletSettingsWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWalletSettingsWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

