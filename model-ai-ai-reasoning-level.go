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

// AiAiReasoningLevel Provider-neutral extended-thinking depth. `off` disables thinking where the model allows it.
type AiAiReasoningLevel string

// List of AiAiReasoningLevel
const (
	AIAIREASONINGLEVEL_OFF AiAiReasoningLevel = "off"
	AIAIREASONINGLEVEL_LOW AiAiReasoningLevel = "low"
	AIAIREASONINGLEVEL_MEDIUM AiAiReasoningLevel = "medium"
	AIAIREASONINGLEVEL_HIGH AiAiReasoningLevel = "high"
	AIAIREASONINGLEVEL_MAX AiAiReasoningLevel = "max"
)

// All allowed values of AiAiReasoningLevel enum
var AllowedAiAiReasoningLevelEnumValues = []AiAiReasoningLevel{
	"off",
	"low",
	"medium",
	"high",
	"max",
}

func (v *AiAiReasoningLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiAiReasoningLevel(value)
	for _, existing := range AllowedAiAiReasoningLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiAiReasoningLevel", value)
}

// NewAiAiReasoningLevelFromValue returns a pointer to a valid AiAiReasoningLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiAiReasoningLevelFromValue(v string) (*AiAiReasoningLevel, error) {
	ev := AiAiReasoningLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiAiReasoningLevel: valid values are %v", v, AllowedAiAiReasoningLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiAiReasoningLevel) IsValid() bool {
	for _, existing := range AllowedAiAiReasoningLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiAiReasoningLevel value
func (v AiAiReasoningLevel) Ptr() *AiAiReasoningLevel {
	return &v
}

type NullableAiAiReasoningLevel struct {
	value *AiAiReasoningLevel
	isSet bool
}

func (v NullableAiAiReasoningLevel) Get() *AiAiReasoningLevel {
	return v.value
}

func (v *NullableAiAiReasoningLevel) Set(val *AiAiReasoningLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiReasoningLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiReasoningLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiReasoningLevel(val *AiAiReasoningLevel) *NullableAiAiReasoningLevel {
	return &NullableAiAiReasoningLevel{value: val, isSet: true}
}

func (v NullableAiAiReasoningLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiReasoningLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

