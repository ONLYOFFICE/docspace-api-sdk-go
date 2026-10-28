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

// checks if the IpRestrictionsSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpRestrictionsSettingsDto{}

// IpRestrictionsSettingsDto Whether the portal limits sign-in to its list of allowed IP addresses.
type IpRestrictionsSettingsDto struct {
	// Specifies if the IP restrictions are enabled or not.
	Enable *bool `json:"enable,omitempty"`
	// The date and time when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewIpRestrictionsSettingsDto instantiates a new IpRestrictionsSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRestrictionsSettingsDto() *IpRestrictionsSettingsDto {
	this := IpRestrictionsSettingsDto{}
	return &this
}

// NewIpRestrictionsSettingsDtoWithDefaults instantiates a new IpRestrictionsSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRestrictionsSettingsDtoWithDefaults() *IpRestrictionsSettingsDto {
	this := IpRestrictionsSettingsDto{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IpRestrictionsSettingsDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionsSettingsDto) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IpRestrictionsSettingsDto) IsEnableSet() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IpRestrictionsSettingsDto) SetEnable(v bool) {
	o.Enable = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *IpRestrictionsSettingsDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionsSettingsDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *IpRestrictionsSettingsDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *IpRestrictionsSettingsDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o IpRestrictionsSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpRestrictionsSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableIpRestrictionsSettingsDto struct {
	value *IpRestrictionsSettingsDto
	isSet bool
}

func (v NullableIpRestrictionsSettingsDto) Get() *IpRestrictionsSettingsDto {
	return v.value
}

func (v *NullableIpRestrictionsSettingsDto) Set(val *IpRestrictionsSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRestrictionsSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRestrictionsSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRestrictionsSettingsDto(val *IpRestrictionsSettingsDto) *NullableIpRestrictionsSettingsDto {
	return &NullableIpRestrictionsSettingsDto{value: val, isSet: true}
}

func (v NullableIpRestrictionsSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRestrictionsSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

