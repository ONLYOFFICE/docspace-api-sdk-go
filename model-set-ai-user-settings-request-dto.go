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

// checks if the SetAiUserSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetAiUserSettingsRequestDto{}

// SetAiUserSettingsRequestDto Request to update per-user AI recommended model visibility settings.
type SetAiUserSettingsRequestDto struct {
	// Indicates whether the recommended model banner is visible in the AI chat.
	ChatRecommendedModelVisible *bool `json:"chatRecommendedModelVisible,omitempty"`
}

// NewSetAiUserSettingsRequestDto instantiates a new SetAiUserSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAiUserSettingsRequestDto() *SetAiUserSettingsRequestDto {
	this := SetAiUserSettingsRequestDto{}
	return &this
}

// NewSetAiUserSettingsRequestDtoWithDefaults instantiates a new SetAiUserSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAiUserSettingsRequestDtoWithDefaults() *SetAiUserSettingsRequestDto {
	this := SetAiUserSettingsRequestDto{}
	return &this
}

// GetChatRecommendedModelVisible returns the ChatRecommendedModelVisible field value if set, zero value otherwise.
func (o *SetAiUserSettingsRequestDto) GetChatRecommendedModelVisible() bool {
	if o == nil || IsNil(o.ChatRecommendedModelVisible) {
		var ret bool
		return ret
	}
	return *o.ChatRecommendedModelVisible
}

// GetChatRecommendedModelVisibleOk returns a tuple with the ChatRecommendedModelVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAiUserSettingsRequestDto) GetChatRecommendedModelVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatRecommendedModelVisible) {
		return nil, false
	}
	return o.ChatRecommendedModelVisible, true
}

// HasChatRecommendedModelVisible returns a boolean if a field has been set.
func (o *SetAiUserSettingsRequestDto) IsChatRecommendedModelVisibleSet() bool {
	if o != nil && !IsNil(o.ChatRecommendedModelVisible) {
		return true
	}

	return false
}

// SetChatRecommendedModelVisible gets a reference to the given bool and assigns it to the ChatRecommendedModelVisible field.
func (o *SetAiUserSettingsRequestDto) SetChatRecommendedModelVisible(v bool) {
	o.ChatRecommendedModelVisible = &v
}

func (o SetAiUserSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAiUserSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatRecommendedModelVisible) {
		toSerialize["chatRecommendedModelVisible"] = o.ChatRecommendedModelVisible
	}
	return toSerialize, nil
}

type NullableSetAiUserSettingsRequestDto struct {
	value *SetAiUserSettingsRequestDto
	isSet bool
}

func (v NullableSetAiUserSettingsRequestDto) Get() *SetAiUserSettingsRequestDto {
	return v.value
}

func (v *NullableSetAiUserSettingsRequestDto) Set(val *SetAiUserSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAiUserSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAiUserSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAiUserSettingsRequestDto(val *SetAiUserSettingsRequestDto) *NullableSetAiUserSettingsRequestDto {
	return &NullableSetAiUserSettingsRequestDto{value: val, isSet: true}
}

func (v NullableSetAiUserSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAiUserSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

