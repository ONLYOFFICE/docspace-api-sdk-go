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

// AiActionType The AI action a request or an assignment applies to. Each action has its own assignment slot; `Default` is the profile used when an action's own slot is empty.
type AiActionType string

// List of AiActionType
const (
	AIACTIONTYPE_DEFAULT AiActionType = "Default"
	AIACTIONTYPE_CHAT AiActionType = "Chat"
	AIACTIONTYPE_CODE AiActionType = "Code"
	AIACTIONTYPE_SUMMARIZATION AiActionType = "Summarization"
	AIACTIONTYPE_TRANSLATION AiActionType = "Translation"
	AIACTIONTYPE_TEXT_ANALYZE AiActionType = "TextAnalyze"
	AIACTIONTYPE_IMAGE_GENERATION AiActionType = "ImageGeneration"
	AIACTIONTYPE_OCR AiActionType = "OCR"
	AIACTIONTYPE_VISION AiActionType = "Vision"
)

// All allowed values of AiActionType enum
var AllowedAiActionTypeEnumValues = []AiActionType{
	"Default",
	"Chat",
	"Code",
	"Summarization",
	"Translation",
	"TextAnalyze",
	"ImageGeneration",
	"OCR",
	"Vision",
}

func (v *AiActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiActionType(value)
	for _, existing := range AllowedAiActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiActionType", value)
}

// NewAiActionTypeFromValue returns a pointer to a valid AiActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiActionTypeFromValue(v string) (*AiActionType, error) {
	ev := AiActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiActionType: valid values are %v", v, AllowedAiActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiActionType) IsValid() bool {
	for _, existing := range AllowedAiActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiActionType value
func (v AiActionType) Ptr() *AiActionType {
	return &v
}

type NullableAiActionType struct {
	value *AiActionType
	isSet bool
}

func (v NullableAiActionType) Get() *AiActionType {
	return v.value
}

func (v *NullableAiActionType) Set(val *AiActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiActionType(val *AiActionType) *NullableAiActionType {
	return &NullableAiActionType{value: val, isSet: true}
}

func (v NullableAiActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

