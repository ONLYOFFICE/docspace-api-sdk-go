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

// checks if the AiChatSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiChatSettingsDto{}

// AiChatSettingsDto The chat settings parameters.
type AiChatSettingsDto struct {
	// The system prompt for the chat.
	Prompt NullableString `json:"prompt,omitempty"`
}

// NewAiChatSettingsDto instantiates a new AiChatSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiChatSettingsDto() *AiChatSettingsDto {
	this := AiChatSettingsDto{}
	return &this
}

// NewAiChatSettingsDtoWithDefaults instantiates a new AiChatSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiChatSettingsDtoWithDefaults() *AiChatSettingsDto {
	this := AiChatSettingsDto{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiChatSettingsDto) GetPrompt() string {
	if o == nil || IsNil(o.Prompt.Get()) {
		var ret string
		return ret
	}
	return *o.Prompt.Get()
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiChatSettingsDto) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompt.Get(), o.Prompt.IsSet()
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiChatSettingsDto) IsPromptSet() bool {
	if o != nil && o.Prompt.IsSet() {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given NullableString and assigns it to the Prompt field.
func (o *AiChatSettingsDto) SetPrompt(v string) {
	o.Prompt.Set(&v)
}
// SetPromptNil sets the value for Prompt to be an explicit nil
func (o *AiChatSettingsDto) SetPromptNil() {
	o.Prompt.Set(nil)
}

// UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
func (o *AiChatSettingsDto) UnsetPrompt() {
	o.Prompt.Unset()
}

func (o AiChatSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiChatSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Prompt.IsSet() {
		toSerialize["prompt"] = o.Prompt.Get()
	}
	return toSerialize, nil
}

type NullableAiChatSettingsDto struct {
	value *AiChatSettingsDto
	isSet bool
}

func (v NullableAiChatSettingsDto) Get() *AiChatSettingsDto {
	return v.value
}

func (v *NullableAiChatSettingsDto) Set(val *AiChatSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatSettingsDto(val *AiChatSettingsDto) *NullableAiChatSettingsDto {
	return &NullableAiChatSettingsDto{value: val, isSet: true}
}

func (v NullableAiChatSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

