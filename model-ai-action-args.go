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

// checks if the AiActionArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiActionArgs{}

// AiActionArgs struct for AiActionArgs
type AiActionArgs struct {
	// Extra tools offered to the model for this request.
	Tools []AiMCPItem `json:"tools,omitempty"`
	// Legacy extended-thinking switch; stands for `medium`. `reasoningLevel` wins when both are set.
	IsReasoning *bool `json:"isReasoning,omitempty"`
	// Depth of extended thinking for the round; providers clamp it to what the model accepts.
	ReasoningLevel *AiReasoningLevel `json:"reasoningLevel,omitempty"`
	Prompt *AiActionArgsPrompt `json:"prompt,omitempty"`
}

// NewAiActionArgs instantiates a new AiActionArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiActionArgs() *AiActionArgs {
	this := AiActionArgs{}
	return &this
}

// NewAiActionArgsWithDefaults instantiates a new AiActionArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiActionArgsWithDefaults() *AiActionArgs {
	this := AiActionArgs{}
	return &this
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *AiActionArgs) GetTools() []AiMCPItem {
	if o == nil || IsNil(o.Tools) {
		var ret []AiMCPItem
		return ret
	}
	return o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetToolsOk() ([]AiMCPItem, bool) {
	if o == nil || IsNil(o.Tools) {
		return nil, false
	}
	return o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *AiActionArgs) IsToolsSet() bool {
	if o != nil && !IsNil(o.Tools) {
		return true
	}

	return false
}

// SetTools gets a reference to the given []AiMCPItem and assigns it to the Tools field.
func (o *AiActionArgs) SetTools(v []AiMCPItem) {
	o.Tools = v
}

// GetIsReasoning returns the IsReasoning field value if set, zero value otherwise.
func (o *AiActionArgs) GetIsReasoning() bool {
	if o == nil || IsNil(o.IsReasoning) {
		var ret bool
		return ret
	}
	return *o.IsReasoning
}

// GetIsReasoningOk returns a tuple with the IsReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetIsReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReasoning) {
		return nil, false
	}
	return o.IsReasoning, true
}

// HasIsReasoning returns a boolean if a field has been set.
func (o *AiActionArgs) IsIsReasoningSet() bool {
	if o != nil && !IsNil(o.IsReasoning) {
		return true
	}

	return false
}

// SetIsReasoning gets a reference to the given bool and assigns it to the IsReasoning field.
func (o *AiActionArgs) SetIsReasoning(v bool) {
	o.IsReasoning = &v
}

// GetReasoningLevel returns the ReasoningLevel field value if set, zero value otherwise.
func (o *AiActionArgs) GetReasoningLevel() AiReasoningLevel {
	if o == nil || IsNil(o.ReasoningLevel) {
		var ret AiReasoningLevel
		return ret
	}
	return *o.ReasoningLevel
}

// GetReasoningLevelOk returns a tuple with the ReasoningLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetReasoningLevelOk() (*AiReasoningLevel, bool) {
	if o == nil || IsNil(o.ReasoningLevel) {
		return nil, false
	}
	return o.ReasoningLevel, true
}

// HasReasoningLevel returns a boolean if a field has been set.
func (o *AiActionArgs) IsReasoningLevelSet() bool {
	if o != nil && !IsNil(o.ReasoningLevel) {
		return true
	}

	return false
}

// SetReasoningLevel gets a reference to the given AiReasoningLevel and assigns it to the ReasoningLevel field.
func (o *AiActionArgs) SetReasoningLevel(v AiReasoningLevel) {
	o.ReasoningLevel = &v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiActionArgs) GetPrompt() AiActionArgsPrompt {
	if o == nil || IsNil(o.Prompt) {
		var ret AiActionArgsPrompt
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiActionArgs) GetPromptOk() (*AiActionArgsPrompt, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiActionArgs) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given AiActionArgsPrompt and assigns it to the Prompt field.
func (o *AiActionArgs) SetPrompt(v AiActionArgsPrompt) {
	o.Prompt = &v
}

func (o AiActionArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiActionArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tools) {
		toSerialize["tools"] = o.Tools
	}
	if !IsNil(o.IsReasoning) {
		toSerialize["isReasoning"] = o.IsReasoning
	}
	if !IsNil(o.ReasoningLevel) {
		toSerialize["reasoningLevel"] = o.ReasoningLevel
	}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	return toSerialize, nil
}

type NullableAiActionArgs struct {
	value *AiActionArgs
	isSet bool
}

func (v NullableAiActionArgs) Get() *AiActionArgs {
	return v.value
}

func (v *NullableAiActionArgs) Set(val *AiActionArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableAiActionArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableAiActionArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiActionArgs(val *AiActionArgs) *NullableAiActionArgs {
	return &NullableAiActionArgs{value: val, isSet: true}
}

func (v NullableAiActionArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiActionArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

