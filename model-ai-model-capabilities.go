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

// checks if the AiModelCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiModelCapabilities{}

// AiModelCapabilities The AI model capabilities.
type AiModelCapabilities struct {
	// Indicates whether the model supports image and vision input.
	Vision *bool `json:"vision,omitempty"`
	// Indicates whether the model supports tool (function) calling.
	ToolCalling *bool `json:"toolCalling,omitempty"`
	// Indicates whether the model supports extended thinking and reasoning.
	Thinking *bool `json:"thinking,omitempty"`
}

// NewAiModelCapabilities instantiates a new AiModelCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelCapabilities() *AiModelCapabilities {
	this := AiModelCapabilities{}
	return &this
}

// NewAiModelCapabilitiesWithDefaults instantiates a new AiModelCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelCapabilitiesWithDefaults() *AiModelCapabilities {
	this := AiModelCapabilities{}
	return &this
}

// GetVision returns the Vision field value if set, zero value otherwise.
func (o *AiModelCapabilities) GetVision() bool {
	if o == nil || IsNil(o.Vision) {
		var ret bool
		return ret
	}
	return *o.Vision
}

// GetVisionOk returns a tuple with the Vision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelCapabilities) GetVisionOk() (*bool, bool) {
	if o == nil || IsNil(o.Vision) {
		return nil, false
	}
	return o.Vision, true
}

// HasVision returns a boolean if a field has been set.
func (o *AiModelCapabilities) IsVisionSet() bool {
	if o != nil && !IsNil(o.Vision) {
		return true
	}

	return false
}

// SetVision gets a reference to the given bool and assigns it to the Vision field.
func (o *AiModelCapabilities) SetVision(v bool) {
	o.Vision = &v
}

// GetToolCalling returns the ToolCalling field value if set, zero value otherwise.
func (o *AiModelCapabilities) GetToolCalling() bool {
	if o == nil || IsNil(o.ToolCalling) {
		var ret bool
		return ret
	}
	return *o.ToolCalling
}

// GetToolCallingOk returns a tuple with the ToolCalling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelCapabilities) GetToolCallingOk() (*bool, bool) {
	if o == nil || IsNil(o.ToolCalling) {
		return nil, false
	}
	return o.ToolCalling, true
}

// HasToolCalling returns a boolean if a field has been set.
func (o *AiModelCapabilities) IsToolCallingSet() bool {
	if o != nil && !IsNil(o.ToolCalling) {
		return true
	}

	return false
}

// SetToolCalling gets a reference to the given bool and assigns it to the ToolCalling field.
func (o *AiModelCapabilities) SetToolCalling(v bool) {
	o.ToolCalling = &v
}

// GetThinking returns the Thinking field value if set, zero value otherwise.
func (o *AiModelCapabilities) GetThinking() bool {
	if o == nil || IsNil(o.Thinking) {
		var ret bool
		return ret
	}
	return *o.Thinking
}

// GetThinkingOk returns a tuple with the Thinking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelCapabilities) GetThinkingOk() (*bool, bool) {
	if o == nil || IsNil(o.Thinking) {
		return nil, false
	}
	return o.Thinking, true
}

// HasThinking returns a boolean if a field has been set.
func (o *AiModelCapabilities) IsThinkingSet() bool {
	if o != nil && !IsNil(o.Thinking) {
		return true
	}

	return false
}

// SetThinking gets a reference to the given bool and assigns it to the Thinking field.
func (o *AiModelCapabilities) SetThinking(v bool) {
	o.Thinking = &v
}

func (o AiModelCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiModelCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vision) {
		toSerialize["vision"] = o.Vision
	}
	if !IsNil(o.ToolCalling) {
		toSerialize["toolCalling"] = o.ToolCalling
	}
	if !IsNil(o.Thinking) {
		toSerialize["thinking"] = o.Thinking
	}
	return toSerialize, nil
}

type NullableAiModelCapabilities struct {
	value *AiModelCapabilities
	isSet bool
}

func (v NullableAiModelCapabilities) Get() *AiModelCapabilities {
	return v.value
}

func (v *NullableAiModelCapabilities) Set(val *AiModelCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAiModelCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAiModelCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiModelCapabilities(val *AiModelCapabilities) *NullableAiModelCapabilities {
	return &NullableAiModelCapabilities{value: val, isSet: true}
}

func (v NullableAiModelCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiModelCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

