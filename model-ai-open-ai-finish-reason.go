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
	"fmt"
)

// AiOpenAIFinishReason OpenAI Chat Completions streaming shapes.   `toOpenAIChatCompletionStream` maps the engine's transport-agnostic `ChatEvent` stream onto these chunks so a host can expose an OpenAI-compatible `POST /v1/chat/completions` (`stream: true`) endpoint backed by the same chat pipeline as the in-app widget. Only the subset of fields the engine can populate is emitted; everything else an OpenAI client tolerates as absent.
type AiOpenAIFinishReason string

// List of AiOpenAIFinishReason
const (
	AIOPENAIFINISHREASON_STOP AiOpenAIFinishReason = "stop"
	AIOPENAIFINISHREASON_LENGTH AiOpenAIFinishReason = "length"
	AIOPENAIFINISHREASON_TOOL_CALLS AiOpenAIFinishReason = "tool_calls"
	AIOPENAIFINISHREASON_CONTENT_FILTER AiOpenAIFinishReason = "content_filter"
)

// All allowed values of AiOpenAIFinishReason enum
var AllowedAiOpenAIFinishReasonEnumValues = []AiOpenAIFinishReason{
	"stop",
	"length",
	"tool_calls",
	"content_filter",
}

func (v *AiOpenAIFinishReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiOpenAIFinishReason(value)
	for _, existing := range AllowedAiOpenAIFinishReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiOpenAIFinishReason", value)
}

// NewAiOpenAIFinishReasonFromValue returns a pointer to a valid AiOpenAIFinishReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiOpenAIFinishReasonFromValue(v string) (*AiOpenAIFinishReason, error) {
	ev := AiOpenAIFinishReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiOpenAIFinishReason: valid values are %v", v, AllowedAiOpenAIFinishReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiOpenAIFinishReason) IsValid() bool {
	for _, existing := range AllowedAiOpenAIFinishReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiOpenAIFinishReason value
func (v AiOpenAIFinishReason) Ptr() *AiOpenAIFinishReason {
	return &v
}

type NullableAiOpenAIFinishReason struct {
	value *AiOpenAIFinishReason
	isSet bool
}

func (v NullableAiOpenAIFinishReason) Get() *AiOpenAIFinishReason {
	return v.value
}

func (v *NullableAiOpenAIFinishReason) Set(val *AiOpenAIFinishReason) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIFinishReason) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIFinishReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIFinishReason(val *AiOpenAIFinishReason) *NullableAiOpenAIFinishReason {
	return &NullableAiOpenAIFinishReason{value: val, isSet: true}
}

func (v NullableAiOpenAIFinishReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIFinishReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

