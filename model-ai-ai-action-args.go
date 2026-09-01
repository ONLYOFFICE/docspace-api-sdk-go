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

// checks if the AiAiActionArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiActionArgs{}

// AiAiActionArgs Wire-serializable subset of the engine's `ActionArgs` — drops the engine-injected `signal`/`fetch`; `profile`/`messages` are owned by the engine and never sent by the caller.
type AiAiActionArgs struct {
	// Extra tools offered to the model for this request.
	Tools []AiTMCPItem `json:"tools,omitempty"`
	// Enable extended thinking / reasoning for this request.
	IsReasoning *bool `json:"isReasoning,omitempty"`
	Prompt *AiAiActionArgsPrompt `json:"prompt,omitempty"`
}

// NewAiAiActionArgs instantiates a new AiAiActionArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiActionArgs() *AiAiActionArgs {
	this := AiAiActionArgs{}
	return &this
}

// NewAiAiActionArgsWithDefaults instantiates a new AiAiActionArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiActionArgsWithDefaults() *AiAiActionArgs {
	this := AiAiActionArgs{}
	return &this
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *AiAiActionArgs) GetTools() []AiTMCPItem {
	if o == nil || IsNil(o.Tools) {
		var ret []AiTMCPItem
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiActionArgs) GetToolsOk() ([]AiTMCPItem, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *AiAiActionArgs) IsToolsSet() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []AiTMCPItem and assigns it to the Tools field.
func (o *AiAiActionArgs) SetTools(v []AiTMCPItem) {
	o.Tools = v
}

// GetIsReasoning returns the IsReasoning field value if set, zero value otherwise.
func (o *AiAiActionArgs) GetIsReasoning() bool {
	if o == nil || IsNil(o.IsReasoning) {
		var ret bool
		return ret
	}
	return *o.IsReasoning
}

// GetIsReasoningOk returns a tuple with the IsReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiActionArgs) GetIsReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReasoning) {
		return nil, false
	}
	return o.IsReasoning, true
}

// HasIsReasoning returns a boolean if a field has been set.
func (o *AiAiActionArgs) IsIsReasoningSet() bool {
	if o != nil && !IsNil(o.IsReasoning) {
		return true
	}

	return false
}

// SetIsReasoning gets a reference to the given bool and assigns it to the IsReasoning field.
func (o *AiAiActionArgs) SetIsReasoning(v bool) {
	o.IsReasoning = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiAiActionArgs) GetPrompt() AiAiActionArgsPrompt {
	if o == nil || IsNil(o.Prompt) {
		var ret AiAiActionArgsPrompt
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiActionArgs) GetPromptOk() (*AiAiActionArgsPrompt, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiAiActionArgs) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given AiAiActionArgsPrompt and assigns it to the Prompt field.
func (o *AiAiActionArgs) SetPrompt(v AiAiActionArgsPrompt) {
	o.Prompt = &v
}

func (o AiAiActionArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiActionArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.IsReasoning) {
		toSerialize["isReasoning"] = o.IsReasoning
	}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	return toSerialize, nil
}

type NullableAiAiActionArgs struct {
	value *AiAiActionArgs
	isSet bool
}

func (v NullableAiAiActionArgs) Get() *AiAiActionArgs {
	return v.value
}

func (v *NullableAiAiActionArgs) Set(val *AiAiActionArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiActionArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiActionArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiActionArgs(val *AiAiActionArgs) *NullableAiAiActionArgs {
	return &NullableAiAiActionArgs{value: val, isSet: true}
}

func (v NullableAiAiActionArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiActionArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

