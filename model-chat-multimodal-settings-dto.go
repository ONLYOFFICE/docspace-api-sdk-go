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

// checks if the ChatMultimodalSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMultimodalSettingsDto{}

// ChatMultimodalSettingsDto The multimodal settings for the chat model.
type ChatMultimodalSettingsDto struct {
	Image *ChatImageMultimodalSettingsDto `json:"image,omitempty"`
}

// NewChatMultimodalSettingsDto instantiates a new ChatMultimodalSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMultimodalSettingsDto() *ChatMultimodalSettingsDto {
	this := ChatMultimodalSettingsDto{}
	return &this
}

// NewChatMultimodalSettingsDtoWithDefaults instantiates a new ChatMultimodalSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMultimodalSettingsDtoWithDefaults() *ChatMultimodalSettingsDto {
	this := ChatMultimodalSettingsDto{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ChatMultimodalSettingsDto) GetImage() ChatImageMultimodalSettingsDto {
	if o == nil || IsNil(o.Image) {
		var ret ChatImageMultimodalSettingsDto
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMultimodalSettingsDto) GetImageOk() (*ChatImageMultimodalSettingsDto, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ChatMultimodalSettingsDto) IsImageSet() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ChatImageMultimodalSettingsDto and assigns it to the Image field.
func (o *ChatMultimodalSettingsDto) SetImage(v ChatImageMultimodalSettingsDto) {
	o.Image = &v
}

func (o ChatMultimodalSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMultimodalSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableChatMultimodalSettingsDto struct {
	value *ChatMultimodalSettingsDto
	isSet bool
}

func (v NullableChatMultimodalSettingsDto) Get() *ChatMultimodalSettingsDto {
	return v.value
}

func (v *NullableChatMultimodalSettingsDto) Set(val *ChatMultimodalSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMultimodalSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMultimodalSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMultimodalSettingsDto(val *ChatMultimodalSettingsDto) *NullableChatMultimodalSettingsDto {
	return &NullableChatMultimodalSettingsDto{value: val, isSet: true}
}

func (v NullableChatMultimodalSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMultimodalSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

