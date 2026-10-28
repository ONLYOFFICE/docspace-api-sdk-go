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

// checks if the AiUserSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiUserSettingsDto{}

// AiUserSettingsDto The per-user AI settings.
type AiUserSettingsDto struct {
	// Indicates whether the recommended model banner is visible in the AI chat for the current user.
	ChatRecommendedModelVisible *bool `json:"chatRecommendedModelVisible,omitempty"`
	// How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing.
	ToolPermissionMode *AiToolPermissionMode `json:"toolPermissionMode,omitempty"`
}

// NewAiUserSettingsDto instantiates a new AiUserSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiUserSettingsDto() *AiUserSettingsDto {
	this := AiUserSettingsDto{}
	return &this
}

// NewAiUserSettingsDtoWithDefaults instantiates a new AiUserSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiUserSettingsDtoWithDefaults() *AiUserSettingsDto {
	this := AiUserSettingsDto{}
	return &this
}

// GetChatRecommendedModelVisible returns the ChatRecommendedModelVisible field value if set, zero value otherwise.
func (o *AiUserSettingsDto) GetChatRecommendedModelVisible() bool {
	if o == nil || IsNil(o.ChatRecommendedModelVisible) {
		var ret bool
		return ret
	}
	return *o.ChatRecommendedModelVisible
}

// GetChatRecommendedModelVisibleOk returns a tuple with the ChatRecommendedModelVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiUserSettingsDto) GetChatRecommendedModelVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatRecommendedModelVisible) {
		return nil, false
	}
	return o.ChatRecommendedModelVisible, true
}

// HasChatRecommendedModelVisible returns a boolean if a field has been set.
func (o *AiUserSettingsDto) IsChatRecommendedModelVisibleSet() bool {
	if o != nil && !IsNil(o.ChatRecommendedModelVisible) {
		return true
	}

	return false
}

// SetChatRecommendedModelVisible gets a reference to the given bool and assigns it to the ChatRecommendedModelVisible field.
func (o *AiUserSettingsDto) SetChatRecommendedModelVisible(v bool) {
	o.ChatRecommendedModelVisible = &v
}

// GetToolPermissionMode returns the ToolPermissionMode field value if set, zero value otherwise.
func (o *AiUserSettingsDto) GetToolPermissionMode() AiToolPermissionMode {
	if o == nil || IsNil(o.ToolPermissionMode) {
		var ret AiToolPermissionMode
		return ret
	}
	return *o.ToolPermissionMode
}

// GetToolPermissionModeOk returns a tuple with the ToolPermissionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiUserSettingsDto) GetToolPermissionModeOk() (*AiToolPermissionMode, bool) {
	if o == nil || IsNil(o.ToolPermissionMode) {
		return nil, false
	}
	return o.ToolPermissionMode, true
}

// HasToolPermissionMode returns a boolean if a field has been set.
func (o *AiUserSettingsDto) IsToolPermissionModeSet() bool {
	if o != nil && !IsNil(o.ToolPermissionMode) {
		return true
	}

	return false
}

// SetToolPermissionMode gets a reference to the given AiToolPermissionMode and assigns it to the ToolPermissionMode field.
func (o *AiUserSettingsDto) SetToolPermissionMode(v AiToolPermissionMode) {
	o.ToolPermissionMode = &v
}

func (o AiUserSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiUserSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatRecommendedModelVisible) {
		toSerialize["chatRecommendedModelVisible"] = o.ChatRecommendedModelVisible
	}
	if !IsNil(o.ToolPermissionMode) {
		toSerialize["toolPermissionMode"] = o.ToolPermissionMode
	}
	return toSerialize, nil
}

type NullableAiUserSettingsDto struct {
	value *AiUserSettingsDto
	isSet bool
}

func (v NullableAiUserSettingsDto) Get() *AiUserSettingsDto {
	return v.value
}

func (v *NullableAiUserSettingsDto) Set(val *AiUserSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiUserSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiUserSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiUserSettingsDto(val *AiUserSettingsDto) *NullableAiUserSettingsDto {
	return &NullableAiUserSettingsDto{value: val, isSet: true}
}

func (v NullableAiUserSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiUserSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

