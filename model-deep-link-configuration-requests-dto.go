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

// checks if the DeepLinkConfigurationRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeepLinkConfigurationRequestsDto{}

// DeepLinkConfigurationRequestsDto The request parameters for managing the deep link configuration.
type DeepLinkConfigurationRequestsDto struct {
	DeepLinkSettings *TenantDeepLinkSettings `json:"deepLinkSettings,omitempty"`
}

// NewDeepLinkConfigurationRequestsDto instantiates a new DeepLinkConfigurationRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeepLinkConfigurationRequestsDto() *DeepLinkConfigurationRequestsDto {
	this := DeepLinkConfigurationRequestsDto{}
	return &this
}

// NewDeepLinkConfigurationRequestsDtoWithDefaults instantiates a new DeepLinkConfigurationRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeepLinkConfigurationRequestsDtoWithDefaults() *DeepLinkConfigurationRequestsDto {
	this := DeepLinkConfigurationRequestsDto{}
	return &this
}

// GetDeepLinkSettings returns the DeepLinkSettings field value if set, zero value otherwise.
func (o *DeepLinkConfigurationRequestsDto) GetDeepLinkSettings() TenantDeepLinkSettings {
	if o == nil || IsNil(o.DeepLinkSettings) {
		var ret TenantDeepLinkSettings
		return ret
	}
	return *o.DeepLinkSettings
}

// GetDeepLinkSettingsOk returns a tuple with the DeepLinkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeepLinkConfigurationRequestsDto) GetDeepLinkSettingsOk() (*TenantDeepLinkSettings, bool) {
	if o == nil || IsNil(o.DeepLinkSettings) {
		return nil, false
	}
	return o.DeepLinkSettings, true
}

// HasDeepLinkSettings returns a boolean if a field has been set.
func (o *DeepLinkConfigurationRequestsDto) IsDeepLinkSettingsSet() bool {
	if o != nil && !IsNil(o.DeepLinkSettings) {
		return true
	}

	return false
}

// SetDeepLinkSettings gets a reference to the given TenantDeepLinkSettings and assigns it to the DeepLinkSettings field.
func (o *DeepLinkConfigurationRequestsDto) SetDeepLinkSettings(v TenantDeepLinkSettings) {
	o.DeepLinkSettings = &v
}

func (o DeepLinkConfigurationRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeepLinkConfigurationRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeepLinkSettings) {
		toSerialize["deepLinkSettings"] = o.DeepLinkSettings
	}
	return toSerialize, nil
}

type NullableDeepLinkConfigurationRequestsDto struct {
	value *DeepLinkConfigurationRequestsDto
	isSet bool
}

func (v NullableDeepLinkConfigurationRequestsDto) Get() *DeepLinkConfigurationRequestsDto {
	return v.value
}

func (v *NullableDeepLinkConfigurationRequestsDto) Set(val *DeepLinkConfigurationRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeepLinkConfigurationRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeepLinkConfigurationRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeepLinkConfigurationRequestsDto(val *DeepLinkConfigurationRequestsDto) *NullableDeepLinkConfigurationRequestsDto {
	return &NullableDeepLinkConfigurationRequestsDto{value: val, isSet: true}
}

func (v NullableDeepLinkConfigurationRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeepLinkConfigurationRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

