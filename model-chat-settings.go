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

// checks if the ChatSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSettings{}

// ChatSettings The chat settings.
type ChatSettings struct {
	// The provider ID.
	ProviderId *int32 `json:"providerId,omitempty"`
	// The model ID.
	ModelId NullableString `json:"modelId,omitempty"`
	// The prompt.
	Prompt NullableString `json:"prompt,omitempty"`
	// Specifies whether the provider is internal or not.
	Internal *bool `json:"internal,omitempty"`
}

// NewChatSettings instantiates a new ChatSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSettings() *ChatSettings {
	this := ChatSettings{}
	return &this
}

// NewChatSettingsWithDefaults instantiates a new ChatSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSettingsWithDefaults() *ChatSettings {
	this := ChatSettings{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *ChatSettings) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId) {
		var ret int32
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettings) GetProviderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *ChatSettings) IsProviderIdSet() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given int32 and assigns it to the ProviderId field.
func (o *ChatSettings) SetProviderId(v int32) {
	o.ProviderId = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatSettings) GetModelId() string {
	if o == nil || IsNil(o.ModelId.Get()) {
		var ret string
		return ret
	}
	return *o.ModelId.Get()
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatSettings) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModelId.Get(), o.ModelId.IsSet()
}

// HasModelId returns a boolean if a field has been set.
func (o *ChatSettings) IsModelIdSet() bool {
	if o != nil && o.ModelId.IsSet() {
		return true
	}

	return false
}

// SetModelId gets a reference to the given NullableString and assigns it to the ModelId field.
func (o *ChatSettings) SetModelId(v string) {
	o.ModelId.Set(&v)
}
// SetModelIdNil sets the value for ModelId to be an explicit nil
func (o *ChatSettings) SetModelIdNil() {
	o.ModelId.Set(nil)
}

// UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
func (o *ChatSettings) UnsetModelId() {
	o.ModelId.Unset()
}

// GetPrompt returns the Prompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatSettings) GetPrompt() string {
	if o == nil || IsNil(o.Prompt.Get()) {
		var ret string
		return ret
	}
	return *o.Prompt.Get()
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatSettings) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompt.Get(), o.Prompt.IsSet()
}

// HasPrompt returns a boolean if a field has been set.
func (o *ChatSettings) IsPromptSet() bool {
	if o != nil && o.Prompt.IsSet() {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given NullableString and assigns it to the Prompt field.
func (o *ChatSettings) SetPrompt(v string) {
	o.Prompt.Set(&v)
}
// SetPromptNil sets the value for Prompt to be an explicit nil
func (o *ChatSettings) SetPromptNil() {
	o.Prompt.Set(nil)
}

// UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
func (o *ChatSettings) UnsetPrompt() {
	o.Prompt.Unset()
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ChatSettings) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSettings) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ChatSettings) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ChatSettings) SetInternal(v bool) {
	o.Internal = &v
}

func (o ChatSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if o.ModelId.IsSet() {
		toSerialize["modelId"] = o.ModelId.Get()
	}
	if o.Prompt.IsSet() {
		toSerialize["prompt"] = o.Prompt.Get()
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	return toSerialize, nil
}

type NullableChatSettings struct {
	value *ChatSettings
	isSet bool
}

func (v NullableChatSettings) Get() *ChatSettings {
	return v.value
}

func (v *NullableChatSettings) Set(val *ChatSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSettings(val *ChatSettings) *NullableChatSettings {
	return &NullableChatSettings{value: val, isSet: true}
}

func (v NullableChatSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

