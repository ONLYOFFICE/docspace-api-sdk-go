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

// AiReasoningLevel Provider-neutral extended-thinking depth. `off` disables thinking where the model allows it.
type AiReasoningLevel string

// List of AiReasoningLevel
const (
	AIREASONINGLEVEL_OFF AiReasoningLevel = "off"
	AIREASONINGLEVEL_LOW AiReasoningLevel = "low"
	AIREASONINGLEVEL_MEDIUM AiReasoningLevel = "medium"
	AIREASONINGLEVEL_HIGH AiReasoningLevel = "high"
	AIREASONINGLEVEL_MAX AiReasoningLevel = "max"
)

// All allowed values of AiReasoningLevel enum
var AllowedAiReasoningLevelEnumValues = []AiReasoningLevel{
	"off",
	"low",
	"medium",
	"high",
	"max",
}

func (v *AiReasoningLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiReasoningLevel(value)
	for _, existing := range AllowedAiReasoningLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiReasoningLevel", value)
}

// NewAiReasoningLevelFromValue returns a pointer to a valid AiReasoningLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiReasoningLevelFromValue(v string) (*AiReasoningLevel, error) {
	ev := AiReasoningLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiReasoningLevel: valid values are %v", v, AllowedAiReasoningLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiReasoningLevel) IsValid() bool {
	for _, existing := range AllowedAiReasoningLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiReasoningLevel value
func (v AiReasoningLevel) Ptr() *AiReasoningLevel {
	return &v
}

type NullableAiReasoningLevel struct {
	value *AiReasoningLevel
	isSet bool
}

func (v NullableAiReasoningLevel) Get() *AiReasoningLevel {
	return v.value
}

func (v *NullableAiReasoningLevel) Set(val *AiReasoningLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAiReasoningLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAiReasoningLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiReasoningLevel(val *AiReasoningLevel) *NullableAiReasoningLevel {
	return &NullableAiReasoningLevel{value: val, isSet: true}
}

func (v NullableAiReasoningLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiReasoningLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

