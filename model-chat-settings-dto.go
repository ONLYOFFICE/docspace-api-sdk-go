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

// checks if the ChatSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSettingsDto{}

// ChatSettingsDto The chat settings parameters.
type ChatSettingsDto struct {
	// The AI provider ID.
	ProviderId *int32 `json:"providerId,omitempty"`
	// The AI model ID used for chat completions.
	ModelId NullableString `json:"modelId,omitempty"`
	// The AI model display alias.
	ModelAlias NullableString `json:"modelAlias,omitempty"`
	// The system prompt for the chat.
	Prompt NullableString `json:"prompt,omitempty"`
	Multimodal *ChatMultimodalSettingsDto `json:"multimodal,omitempty"`
	// Indicates whether the model supports extended thinking mode.
	Thinking *bool `json:"thinking,omitempty"`
	// Indicates whether this is an internal AI gateway provider.
	Internal *bool `json:"internal,omitempty"`
}

// NewChatSettingsDto instantiates a new ChatSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSettingsDto() *ChatSettingsDto {
	this := ChatSettingsDto{}
	return &this
}

// NewChatSettingsDtoWithDefaults instantiates a new ChatSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSettingsDtoWithDefaults() *ChatSettingsDto {
	this := ChatSettingsDto{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ChatSettingsDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettingsDto) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ChatSettingsDto) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatSettingsDto) GetModelId() string {
	if o == nil || IsNil(o.ModelId.Get()) {
		var ret string
		return ret
	}
	return *o.ModelId.Get()
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatSettingsDto) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelId.Get(), o.ModelId.IsSet()
}

// HasModelId returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsModelIdSet() bool {
	if o != nil && o.ModelId.IsSet() {
		return true
	}

	return false
}

// SetModelId gets a reference to the given NullableString and assigns it to the ModelId field.
func (o *ChatSettingsDto) SetModelId(v string) {
	o.ModelId.Set(&v)
}
// SetModelIdNil sets the value for ModelId to be an explicit nil
func (o *ChatSettingsDto) SetModelIdNil() {
	o.ModelId.Set(nil)
}

// UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
func (o *ChatSettingsDto) UnsetModelId() {
	o.ModelId.Unset()
}

// GetModelAlias returns the ModelAlias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatSettingsDto) GetModelAlias() string {
	if o == nil || IsNil(o.ModelAlias.Get()) {
		var ret string
		return ret
	}
	return *o.ModelAlias.Get()
}

// GetModelAliasOk returns a tuple with the ModelAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatSettingsDto) GetModelAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelAlias.Get(), o.ModelAlias.IsSet()
}

// HasModelAlias returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsModelAliasSet() bool {
	if o != nil && o.ModelAlias.IsSet() {
		return true
	}

	return false
}

// SetModelAlias gets a reference to the given NullableString and assigns it to the ModelAlias field.
func (o *ChatSettingsDto) SetModelAlias(v string) {
	o.ModelAlias.Set(&v)
}
// SetModelAliasNil sets the value for ModelAlias to be an explicit nil
func (o *ChatSettingsDto) SetModelAliasNil() {
	o.ModelAlias.Set(nil)
}

// UnsetModelAlias ensures that no value is present for ModelAlias, not even an explicit nil
func (o *ChatSettingsDto) UnsetModelAlias() {
	o.ModelAlias.Unset()
}

// GetPrompt returns the Prompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatSettingsDto) GetPrompt() string {
	if o == nil || IsNil(o.Prompt.Get()) {
		var ret string
		return ret
	}
	return *o.Prompt.Get()
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatSettingsDto) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompt.Get(), o.Prompt.IsSet()
}

// HasPrompt returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsPromptSet() bool {
	if o != nil && o.Prompt.IsSet() {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given NullableString and assigns it to the Prompt field.
func (o *ChatSettingsDto) SetPrompt(v string) {
	o.Prompt.Set(&v)
}
// SetPromptNil sets the value for Prompt to be an explicit nil
func (o *ChatSettingsDto) SetPromptNil() {
	o.Prompt.Set(nil)
}

// UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
func (o *ChatSettingsDto) UnsetPrompt() {
	o.Prompt.Unset()
}

// GetMultimodal returns the Multimodal field value if set, zero value otherwise.
func (o *ChatSettingsDto) GetMultimodal() ChatMultimodalSettingsDto {
	if o == nil || IsNil(o.Multimodal) {
		var ret ChatMultimodalSettingsDto
		return ret
	}
	return *o.Multimodal
}

// GetMultimodalOk returns a tuple with the Multimodal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettingsDto) GetMultimodalOk() (*ChatMultimodalSettingsDto, bool) {
	if o == nil || IsNil(o.Multimodal) {
		return nil, false
	}
	return o.Multimodal, true
}

// HasMultimodal returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsMultimodalSet() bool {
	if o != nil && !IsNil(o.Multimodal) {
		return true
	}

	return false
}

// SetMultimodal gets a reference to the given ChatMultimodalSettingsDto and assigns it to the Multimodal field.
func (o *ChatSettingsDto) SetMultimodal(v ChatMultimodalSettingsDto) {
	o.Multimodal = &v
}

// GetThinking returns the Thinking field value if set, zero value otherwise.
func (o *ChatSettingsDto) GetThinking() bool {
	if o == nil || IsNil(o.Thinking) {
		var ret bool
		return ret
	}
	return *o.Thinking
}

// GetThinkingOk returns a tuple with the Thinking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettingsDto) GetThinkingOk() (*bool, bool) {
	if o == nil || IsNil(o.Thinking) {
		return nil, false
	}
	return o.Thinking, true
}

// HasThinking returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsThinkingSet() bool {
	if o != nil && !IsNil(o.Thinking) {
		return true
	}

	return false
}

// SetThinking gets a reference to the given bool and assigns it to the Thinking field.
func (o *ChatSettingsDto) SetThinking(v bool) {
	o.Thinking = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ChatSettingsDto) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettingsDto) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ChatSettingsDto) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ChatSettingsDto) SetInternal(v bool) {
	o.Internal = &v
}

func (o ChatSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if o.ModelId.IsSet() {
		toSerialize["modelId"] = o.ModelId.Get()
	}
	if o.ModelAlias.IsSet() {
		toSerialize["modelAlias"] = o.ModelAlias.Get()
	}
	if o.Prompt.IsSet() {
		toSerialize["prompt"] = o.Prompt.Get()
	}
	if !IsNil(o.Multimodal) {
		toSerialize["multimodal"] = o.Multimodal
	}
	if !IsNil(o.Thinking) {
		toSerialize["thinking"] = o.Thinking
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	return toSerialize, nil
}

type NullableChatSettingsDto struct {
	value *ChatSettingsDto
	isSet bool
}

func (v NullableChatSettingsDto) Get() *ChatSettingsDto {
	return v.value
}

func (v *NullableChatSettingsDto) Set(val *ChatSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSettingsDto(val *ChatSettingsDto) *NullableChatSettingsDto {
	return &NullableChatSettingsDto{value: val, isSet: true}
}

func (v NullableChatSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

