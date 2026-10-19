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

// checks if the DeepLinkConfigurationRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeepLinkConfigurationRequestDto{}

// DeepLinkConfigurationRequestDto How the portal opens its links on a mobile device.
type DeepLinkConfigurationRequestDto struct {
	// The deep link configuration to store. Only its `handlingMode` is read - whether a link always opens in the  browser, always in the native application, or asks the user each time - and a mode outside the defined set is  refused with 400 before anything is stored.
	DeepLinkSettings *DeepLinkSettingsRequestDto `json:"deepLinkSettings,omitempty"`
}

// NewDeepLinkConfigurationRequestDto instantiates a new DeepLinkConfigurationRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeepLinkConfigurationRequestDto() *DeepLinkConfigurationRequestDto {
	this := DeepLinkConfigurationRequestDto{}
	return &this
}

// NewDeepLinkConfigurationRequestDtoWithDefaults instantiates a new DeepLinkConfigurationRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeepLinkConfigurationRequestDtoWithDefaults() *DeepLinkConfigurationRequestDto {
	this := DeepLinkConfigurationRequestDto{}
	return &this
}

// GetDeepLinkSettings returns the DeepLinkSettings field value if set, zero value otherwise.
func (o *DeepLinkConfigurationRequestDto) GetDeepLinkSettings() DeepLinkSettingsRequestDto {
	if o == nil || IsNil(o.DeepLinkSettings) {
		var ret DeepLinkSettingsRequestDto
		return ret
	}
	return *o.DeepLinkSettings
}

// GetDeepLinkSettingsOk returns a tuple with the DeepLinkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeepLinkConfigurationRequestDto) GetDeepLinkSettingsOk() (*DeepLinkSettingsRequestDto, bool) {
	if o == nil || IsNil(o.DeepLinkSettings) {
		return nil, false
	}
	return o.DeepLinkSettings, true
}

// HasDeepLinkSettings returns a boolean if a field has been set.
func (o *DeepLinkConfigurationRequestDto) IsDeepLinkSettingsSet() bool {
	if o != nil && !IsNil(o.DeepLinkSettings) {
		return true
	}

	return false
}

// SetDeepLinkSettings gets a reference to the given DeepLinkSettingsRequestDto and assigns it to the DeepLinkSettings field.
func (o *DeepLinkConfigurationRequestDto) SetDeepLinkSettings(v DeepLinkSettingsRequestDto) {
	o.DeepLinkSettings = &v
}

func (o DeepLinkConfigurationRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeepLinkConfigurationRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeepLinkSettings) {
		toSerialize["deepLinkSettings"] = o.DeepLinkSettings
	}
	return toSerialize, nil
}

type NullableDeepLinkConfigurationRequestDto struct {
	value *DeepLinkConfigurationRequestDto
	isSet bool
}

func (v NullableDeepLinkConfigurationRequestDto) Get() *DeepLinkConfigurationRequestDto {
	return v.value
}

func (v *NullableDeepLinkConfigurationRequestDto) Set(val *DeepLinkConfigurationRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeepLinkConfigurationRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeepLinkConfigurationRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeepLinkConfigurationRequestDto(val *DeepLinkConfigurationRequestDto) *NullableDeepLinkConfigurationRequestDto {
	return &NullableDeepLinkConfigurationRequestDto{value: val, isSet: true}
}

func (v NullableDeepLinkConfigurationRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeepLinkConfigurationRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

