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

// checks if the AiChatPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiChatPrice{}

// AiChatPrice The price of a chat model, per token.
type AiChatPrice struct {
	// The price of a single prompt token.
	Prompt *float64 `json:"prompt,omitempty"`
	// The price of a single completion token.
	Completion *float64 `json:"completion,omitempty"`
}

// NewAiChatPrice instantiates a new AiChatPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiChatPrice() *AiChatPrice {
	this := AiChatPrice{}
	return &this
}

// NewAiChatPriceWithDefaults instantiates a new AiChatPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiChatPriceWithDefaults() *AiChatPrice {
	this := AiChatPrice{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiChatPrice) GetPrompt() float64 {
	if o == nil || IsNil(o.Prompt) {
		var ret float64
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatPrice) GetPromptOk() (*float64, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiChatPrice) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given float64 and assigns it to the Prompt field.
func (o *AiChatPrice) SetPrompt(v float64) {
	o.Prompt = &v
}

// GetCompletion returns the Completion field value if set, zero value otherwise.
func (o *AiChatPrice) GetCompletion() float64 {
	if o == nil || IsNil(o.Completion) {
		var ret float64
		return ret
	}
	return *o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiChatPrice) GetCompletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Completion) {
		return nil, false
	}
	return o.Completion, true
}

// HasCompletion returns a boolean if a field has been set.
func (o *AiChatPrice) IsCompletionSet() bool {
	if o != nil && !IsNil(o.Completion) {
		return true
	}

	return false
}

// SetCompletion gets a reference to the given float64 and assigns it to the Completion field.
func (o *AiChatPrice) SetCompletion(v float64) {
	o.Completion = &v
}

func (o AiChatPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiChatPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Completion) {
		toSerialize["completion"] = o.Completion
	}
	return toSerialize, nil
}

type NullableAiChatPrice struct {
	value *AiChatPrice
	isSet bool
}

func (v NullableAiChatPrice) Get() *AiChatPrice {
	return v.value
}

func (v *NullableAiChatPrice) Set(val *AiChatPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatPrice(val *AiChatPrice) *NullableAiChatPrice {
	return &NullableAiChatPrice{value: val, isSet: true}
}

func (v NullableAiChatPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

