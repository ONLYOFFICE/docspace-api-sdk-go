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

// checks if the AiOpenAIChoiceDelta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIChoiceDelta{}

// AiOpenAIChoiceDelta The incremental part of one choice - what this chunk adds to the assistant message.
type AiOpenAIChoiceDelta struct {
	// Sent on the first chunk only, always `assistant`.
	Role *string `json:"role,omitempty"`
	// The text this chunk appends. Null when the chunk carries no text.
	Content NullableString `json:"content,omitempty"`
	// The tool calls the model requested, emitted in place of text.
	ToolCalls []AiOpenAIToolCallDelta `json:"tool_calls,omitempty"`
}

// NewAiOpenAIChoiceDelta instantiates a new AiOpenAIChoiceDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIChoiceDelta() *AiOpenAIChoiceDelta {
	this := AiOpenAIChoiceDelta{}
	return &this
}

// NewAiOpenAIChoiceDeltaWithDefaults instantiates a new AiOpenAIChoiceDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIChoiceDeltaWithDefaults() *AiOpenAIChoiceDelta {
	this := AiOpenAIChoiceDelta{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AiOpenAIChoiceDelta) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIChoiceDelta) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AiOpenAIChoiceDelta) IsRoleSet() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AiOpenAIChoiceDelta) SetRole(v string) {
	o.Role = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiOpenAIChoiceDelta) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenAIChoiceDelta) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *AiOpenAIChoiceDelta) IsContentSet() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *AiOpenAIChoiceDelta) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *AiOpenAIChoiceDelta) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *AiOpenAIChoiceDelta) UnsetContent() {
	o.Content.Unset()
}

// GetToolCalls returns the ToolCalls field value if set, zero value otherwise.
func (o *AiOpenAIChoiceDelta) GetToolCalls() []AiOpenAIToolCallDelta {
	if o == nil || IsNil(o.ToolCalls) {
		var ret []AiOpenAIToolCallDelta
		return ret
	}
	return o.ToolCalls
}

// GetToolCallsOk returns a tuple with the ToolCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiOpenAIChoiceDelta) GetToolCallsOk() ([]AiOpenAIToolCallDelta, bool) {
	if o == nil || IsNil(o.ToolCalls) {
		return nil, false
	}
	return o.ToolCalls, true
}

// HasToolCalls returns a boolean if a field has been set.
func (o *AiOpenAIChoiceDelta) IsToolCallsSet() bool {
	if o != nil && !IsNil(o.ToolCalls) {
		return true
	}

	return false
}

// SetToolCalls gets a reference to the given []AiOpenAIToolCallDelta and assigns it to the ToolCalls field.
func (o *AiOpenAIChoiceDelta) SetToolCalls(v []AiOpenAIToolCallDelta) {
	o.ToolCalls = v
}

func (o AiOpenAIChoiceDelta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIChoiceDelta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if !IsNil(o.ToolCalls) {
		toSerialize["tool_calls"] = o.ToolCalls
	}
	return toSerialize, nil
}

type NullableAiOpenAIChoiceDelta struct {
	value *AiOpenAIChoiceDelta
	isSet bool
}

func (v NullableAiOpenAIChoiceDelta) Get() *AiOpenAIChoiceDelta {
	return v.value
}

func (v *NullableAiOpenAIChoiceDelta) Set(val *AiOpenAIChoiceDelta) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIChoiceDelta) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIChoiceDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIChoiceDelta(val *AiOpenAIChoiceDelta) *NullableAiOpenAIChoiceDelta {
	return &NullableAiOpenAIChoiceDelta{value: val, isSet: true}
}

func (v NullableAiOpenAIChoiceDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIChoiceDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

