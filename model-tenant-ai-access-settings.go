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
	"time"
)

// checks if the TenantAiAccessSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantAiAccessSettings{}

// TenantAiAccessSettings The tenant-level settings for enabling or disabling all AI functionality in DocSpace.
type TenantAiAccessSettings struct {
	// Specifies whether AI functionality is enabled for the tenant.  When set to `false`, all AI features (chat, agents, vectorization) are disabled tenant-wide.
	Enabled *bool `json:"enabled,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantAiAccessSettings instantiates a new TenantAiAccessSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAiAccessSettings() *TenantAiAccessSettings {
	this := TenantAiAccessSettings{}
	return &this
}

// NewTenantAiAccessSettingsWithDefaults instantiates a new TenantAiAccessSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAiAccessSettingsWithDefaults() *TenantAiAccessSettings {
	this := TenantAiAccessSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TenantAiAccessSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAccessSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenantAiAccessSettings) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TenantAiAccessSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantAiAccessSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAiAccessSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantAiAccessSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantAiAccessSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantAiAccessSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantAiAccessSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantAiAccessSettings struct {
	value *TenantAiAccessSettings
	isSet bool
}

func (v NullableTenantAiAccessSettings) Get() *TenantAiAccessSettings {
	return v.value
}

func (v *NullableTenantAiAccessSettings) Set(val *TenantAiAccessSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAiAccessSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAiAccessSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAiAccessSettings(val *TenantAiAccessSettings) *NullableTenantAiAccessSettings {
	return &NullableTenantAiAccessSettings{value: val, isSet: true}
}

func (v NullableTenantAiAccessSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAiAccessSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

