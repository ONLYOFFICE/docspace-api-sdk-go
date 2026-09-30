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

// checks if the AiChatPriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiChatPriceDto{}

// AiChatPriceDto What a chat model charges, split by the direction the tokens flow in.
type AiChatPriceDto struct {
	// The cost of one million tokens sent to the model, which includes the conversation history resent with  every turn and not just the newest message.
	Prompt *float64 `json:"prompt,omitempty"`
	// The cost of one million tokens the model writes back. It is normally the dearer of the two directions.
	Completion *float64 `json:"completion,omitempty"`
	// The cost of one million prompt tokens served from the prompt cache. It is absent when the model does not  support prompt caching.
	PromptCacheRead NullableFloat64 `json:"promptCacheRead,omitempty"`
	// The cost of one million prompt tokens written to the prompt cache with the default lifetime. It is absent  when the model does not support prompt caching.
	PromptCacheWrite NullableFloat64 `json:"promptCacheWrite,omitempty"`
	// The cost of one million prompt tokens written to the prompt cache with a one-hour lifetime. It is absent  when the model offers no such option.
	PromptCacheWrite1H NullableFloat64 `json:"promptCacheWrite1H,omitempty"`
}

// NewAiChatPriceDto instantiates a new AiChatPriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiChatPriceDto() *AiChatPriceDto {
	this := AiChatPriceDto{}
	return &this
}

// NewAiChatPriceDtoWithDefaults instantiates a new AiChatPriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiChatPriceDtoWithDefaults() *AiChatPriceDto {
	this := AiChatPriceDto{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiChatPriceDto) GetPrompt() float64 {
	if o == nil || IsNil(o.Prompt) {
		var ret float64
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatPriceDto) GetPromptOk() (*float64, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiChatPriceDto) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given float64 and assigns it to the Prompt field.
func (o *AiChatPriceDto) SetPrompt(v float64) {
	o.Prompt = &v
}

// GetCompletion returns the Completion field value if set, zero value otherwise.
func (o *AiChatPriceDto) GetCompletion() float64 {
	if o == nil || IsNil(o.Completion) {
		var ret float64
		return ret
	}
	return *o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatPriceDto) GetCompletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Completion) {
		return nil, false
	}
	return o.Completion, true
}

// HasCompletion returns a boolean if a field has been set.
func (o *AiChatPriceDto) IsCompletionSet() bool {
	if o != nil && !IsNil(o.Completion) {
		return true
	}

	return false
}

// SetCompletion gets a reference to the given float64 and assigns it to the Completion field.
func (o *AiChatPriceDto) SetCompletion(v float64) {
	o.Completion = &v
}

// GetPromptCacheRead returns the PromptCacheRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiChatPriceDto) GetPromptCacheRead() float64 {
	if o == nil || IsNil(o.PromptCacheRead.Get()) {
		var ret float64
		return ret
	}
	return *o.PromptCacheRead.Get()
}

// GetPromptCacheReadOk returns a tuple with the PromptCacheRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiChatPriceDto) GetPromptCacheReadOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromptCacheRead.Get(), o.PromptCacheRead.IsSet()
}

// HasPromptCacheRead returns a boolean if a field has been set.
func (o *AiChatPriceDto) IsPromptCacheReadSet() bool {
	if o != nil && o.PromptCacheRead.IsSet() {
		return true
	}

	return false
}

// SetPromptCacheRead gets a reference to the given NullableFloat64 and assigns it to the PromptCacheRead field.
func (o *AiChatPriceDto) SetPromptCacheRead(v float64) {
	o.PromptCacheRead.Set(&v)
}
// SetPromptCacheReadNil sets the value for PromptCacheRead to be an explicit nil
func (o *AiChatPriceDto) SetPromptCacheReadNil() {
	o.PromptCacheRead.Set(nil)
}

// UnsetPromptCacheRead ensures that no value is present for PromptCacheRead, not even an explicit nil
func (o *AiChatPriceDto) UnsetPromptCacheRead() {
	o.PromptCacheRead.Unset()
}

// GetPromptCacheWrite returns the PromptCacheWrite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiChatPriceDto) GetPromptCacheWrite() float64 {
	if o == nil || IsNil(o.PromptCacheWrite.Get()) {
		var ret float64
		return ret
	}
	return *o.PromptCacheWrite.Get()
}

// GetPromptCacheWriteOk returns a tuple with the PromptCacheWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiChatPriceDto) GetPromptCacheWriteOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromptCacheWrite.Get(), o.PromptCacheWrite.IsSet()
}

// HasPromptCacheWrite returns a boolean if a field has been set.
func (o *AiChatPriceDto) IsPromptCacheWriteSet() bool {
	if o != nil && o.PromptCacheWrite.IsSet() {
		return true
	}

	return false
}

// SetPromptCacheWrite gets a reference to the given NullableFloat64 and assigns it to the PromptCacheWrite field.
func (o *AiChatPriceDto) SetPromptCacheWrite(v float64) {
	o.PromptCacheWrite.Set(&v)
}
// SetPromptCacheWriteNil sets the value for PromptCacheWrite to be an explicit nil
func (o *AiChatPriceDto) SetPromptCacheWriteNil() {
	o.PromptCacheWrite.Set(nil)
}

// UnsetPromptCacheWrite ensures that no value is present for PromptCacheWrite, not even an explicit nil
func (o *AiChatPriceDto) UnsetPromptCacheWrite() {
	o.PromptCacheWrite.Unset()
}

// GetPromptCacheWrite1H returns the PromptCacheWrite1H field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiChatPriceDto) GetPromptCacheWrite1H() float64 {
	if o == nil || IsNil(o.PromptCacheWrite1H.Get()) {
		var ret float64
		return ret
	}
	return *o.PromptCacheWrite1H.Get()
}

// GetPromptCacheWrite1HOk returns a tuple with the PromptCacheWrite1H field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiChatPriceDto) GetPromptCacheWrite1HOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromptCacheWrite1H.Get(), o.PromptCacheWrite1H.IsSet()
}

// HasPromptCacheWrite1H returns a boolean if a field has been set.
func (o *AiChatPriceDto) IsPromptCacheWrite1HSet() bool {
	if o != nil && o.PromptCacheWrite1H.IsSet() {
		return true
	}

	return false
}

// SetPromptCacheWrite1H gets a reference to the given NullableFloat64 and assigns it to the PromptCacheWrite1H field.
func (o *AiChatPriceDto) SetPromptCacheWrite1H(v float64) {
	o.PromptCacheWrite1H.Set(&v)
}
// SetPromptCacheWrite1HNil sets the value for PromptCacheWrite1H to be an explicit nil
func (o *AiChatPriceDto) SetPromptCacheWrite1HNil() {
	o.PromptCacheWrite1H.Set(nil)
}

// UnsetPromptCacheWrite1H ensures that no value is present for PromptCacheWrite1H, not even an explicit nil
func (o *AiChatPriceDto) UnsetPromptCacheWrite1H() {
	o.PromptCacheWrite1H.Unset()
}

func (o AiChatPriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiChatPriceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Completion) {
		toSerialize["completion"] = o.Completion
	}
	if o.PromptCacheRead.IsSet() {
		toSerialize["promptCacheRead"] = o.PromptCacheRead.Get()
	}
	if o.PromptCacheWrite.IsSet() {
		toSerialize["promptCacheWrite"] = o.PromptCacheWrite.Get()
	}
	if o.PromptCacheWrite1H.IsSet() {
		toSerialize["promptCacheWrite1H"] = o.PromptCacheWrite1H.Get()
	}
	return toSerialize, nil
}

type NullableAiChatPriceDto struct {
	value *AiChatPriceDto
	isSet bool
}

func (v NullableAiChatPriceDto) Get() *AiChatPriceDto {
	return v.value
}

func (v *NullableAiChatPriceDto) Set(val *AiChatPriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatPriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatPriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatPriceDto(val *AiChatPriceDto) *NullableAiChatPriceDto {
	return &NullableAiChatPriceDto{value: val, isSet: true}
}

func (v NullableAiChatPriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatPriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

