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

// checks if the TenantDevToolsAccessSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantDevToolsAccessSettings{}

// TenantDevToolsAccessSettings The Developer Tools access settings.
type TenantDevToolsAccessSettings struct {
	// Specifies if the Developer Tools access are limited for users or not.
	LimitedAccessForUsers *bool `json:"limitedAccessForUsers,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantDevToolsAccessSettings instantiates a new TenantDevToolsAccessSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantDevToolsAccessSettings() *TenantDevToolsAccessSettings {
	this := TenantDevToolsAccessSettings{}
	return &this
}

// NewTenantDevToolsAccessSettingsWithDefaults instantiates a new TenantDevToolsAccessSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDevToolsAccessSettingsWithDefaults() *TenantDevToolsAccessSettings {
	this := TenantDevToolsAccessSettings{}
	return &this
}

// GetLimitedAccessForUsers returns the LimitedAccessForUsers field value if set, zero value otherwise.
func (o *TenantDevToolsAccessSettings) GetLimitedAccessForUsers() bool {
	if o == nil || IsNil(o.LimitedAccessForUsers) {
		var ret bool
		return ret
	}
	return *o.LimitedAccessForUsers
}

// GetLimitedAccessForUsersOk returns a tuple with the LimitedAccessForUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDevToolsAccessSettings) GetLimitedAccessForUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitedAccessForUsers) {
		return nil, false
	}
	return o.LimitedAccessForUsers, true
}

// HasLimitedAccessForUsers returns a boolean if a field has been set.
func (o *TenantDevToolsAccessSettings) IsLimitedAccessForUsersSet() bool {
	if o != nil && !IsNil(o.LimitedAccessForUsers) {
		return true
	}

	return false
}

// SetLimitedAccessForUsers gets a reference to the given bool and assigns it to the LimitedAccessForUsers field.
func (o *TenantDevToolsAccessSettings) SetLimitedAccessForUsers(v bool) {
	o.LimitedAccessForUsers = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantDevToolsAccessSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantDevToolsAccessSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantDevToolsAccessSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantDevToolsAccessSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantDevToolsAccessSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantDevToolsAccessSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitedAccessForUsers) {
		toSerialize["limitedAccessForUsers"] = o.LimitedAccessForUsers
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantDevToolsAccessSettings struct {
	value *TenantDevToolsAccessSettings
	isSet bool
}

func (v NullableTenantDevToolsAccessSettings) Get() *TenantDevToolsAccessSettings {
	return v.value
}

func (v *NullableTenantDevToolsAccessSettings) Set(val *TenantDevToolsAccessSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantDevToolsAccessSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantDevToolsAccessSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantDevToolsAccessSettings(val *TenantDevToolsAccessSettings) *NullableTenantDevToolsAccessSettings {
	return &NullableTenantDevToolsAccessSettings{value: val, isSet: true}
}

func (v NullableTenantDevToolsAccessSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantDevToolsAccessSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

