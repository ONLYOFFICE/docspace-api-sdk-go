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

// AiReasoningDepth A `ReasoningLevel` above off — a depth the model can think at.
type AiReasoningDepth string

// List of AiReasoningDepth
const (
	AIREASONINGDEPTH_LOW AiReasoningDepth = "low"
	AIREASONINGDEPTH_MEDIUM AiReasoningDepth = "medium"
	AIREASONINGDEPTH_HIGH AiReasoningDepth = "high"
	AIREASONINGDEPTH_MAX AiReasoningDepth = "max"
)

// All allowed values of AiReasoningDepth enum
var AllowedAiReasoningDepthEnumValues = []AiReasoningDepth{
	"low",
	"medium",
	"high",
	"max",
}

func (v *AiReasoningDepth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiReasoningDepth(value)
	for _, existing := range AllowedAiReasoningDepthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiReasoningDepth", value)
}

// NewAiReasoningDepthFromValue returns a pointer to a valid AiReasoningDepth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiReasoningDepthFromValue(v string) (*AiReasoningDepth, error) {
	ev := AiReasoningDepth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiReasoningDepth: valid values are %v", v, AllowedAiReasoningDepthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiReasoningDepth) IsValid() bool {
	for _, existing := range AllowedAiReasoningDepthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiReasoningDepth value
func (v AiReasoningDepth) Ptr() *AiReasoningDepth {
	return &v
}

type NullableAiReasoningDepth struct {
	value *AiReasoningDepth
	isSet bool
}

func (v NullableAiReasoningDepth) Get() *AiReasoningDepth {
	return v.value
}

func (v *NullableAiReasoningDepth) Set(val *AiReasoningDepth) {
	v.value = val
	v.isSet = true
}

func (v NullableAiReasoningDepth) IsSet() bool {
	return v.isSet
}

func (v *NullableAiReasoningDepth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiReasoningDepth(val *AiReasoningDepth) *NullableAiReasoningDepth {
	return &NullableAiReasoningDepth{value: val, isSet: true}
}

func (v NullableAiReasoningDepth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiReasoningDepth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

