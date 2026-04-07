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

// checks if the ChatImageMultimodalSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatImageMultimodalSettingsDto{}

// ChatImageMultimodalSettingsDto The image multimodal settings for the chat model.
type ChatImageMultimodalSettingsDto struct {
	// The supported image formats.
	Formats []string `json:"formats,omitempty"`
}

// NewChatImageMultimodalSettingsDto instantiates a new ChatImageMultimodalSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatImageMultimodalSettingsDto() *ChatImageMultimodalSettingsDto {
	this := ChatImageMultimodalSettingsDto{}
	return &this
}

// NewChatImageMultimodalSettingsDtoWithDefaults instantiates a new ChatImageMultimodalSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatImageMultimodalSettingsDtoWithDefaults() *ChatImageMultimodalSettingsDto {
	this := ChatImageMultimodalSettingsDto{}
	return &this
}

// GetFormats returns the Formats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatImageMultimodalSettingsDto) GetFormats() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Formats
}

// GetFormatsOk returns a tuple with the Formats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatImageMultimodalSettingsDto) GetFormatsOk() ([]string, bool) {
	if o == nil || IsNil(o.Formats) {
		return nil, false
	}
	return o.Formats, true
}

// HasFormats returns a boolean if a field has been set.
func (o *ChatImageMultimodalSettingsDto) IsFormatsSet() bool {
	if o != nil && !IsNil(o.Formats) {
		return true
	}

	return false
}

// SetFormats gets a reference to the given []string and assigns it to the Formats field.
func (o *ChatImageMultimodalSettingsDto) SetFormats(v []string) {
	o.Formats = v
}

func (o ChatImageMultimodalSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatImageMultimodalSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Formats != nil {
		toSerialize["formats"] = o.Formats
	}
	return toSerialize, nil
}

type NullableChatImageMultimodalSettingsDto struct {
	value *ChatImageMultimodalSettingsDto
	isSet bool
}

func (v NullableChatImageMultimodalSettingsDto) Get() *ChatImageMultimodalSettingsDto {
	return v.value
}

func (v *NullableChatImageMultimodalSettingsDto) Set(val *ChatImageMultimodalSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatImageMultimodalSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatImageMultimodalSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatImageMultimodalSettingsDto(val *ChatImageMultimodalSettingsDto) *NullableChatImageMultimodalSettingsDto {
	return &NullableChatImageMultimodalSettingsDto{value: val, isSet: true}
}

func (v NullableChatImageMultimodalSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatImageMultimodalSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

