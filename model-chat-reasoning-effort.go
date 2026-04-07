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

// ChatReasoningEffort []
type ChatReasoningEffort int32

// List of ChatReasoningEffort
const (
	CHATREASONINGEFFORT_None ChatReasoningEffort = 0
	CHATREASONINGEFFORT_Low ChatReasoningEffort = 1
	CHATREASONINGEFFORT_Medium ChatReasoningEffort = 2
	CHATREASONINGEFFORT_High ChatReasoningEffort = 3
	CHATREASONINGEFFORT_XHigh ChatReasoningEffort = 4
)

// All allowed values of ChatReasoningEffort enum
var AllowedChatReasoningEffortEnumValues = []ChatReasoningEffort{
	0,
	1,
	2,
	3,
	4,
}

func (v *ChatReasoningEffort) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChatReasoningEffort(value)
	for _, existing := range AllowedChatReasoningEffortEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChatReasoningEffort", value)
}

// NewChatReasoningEffortFromValue returns a pointer to a valid ChatReasoningEffort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChatReasoningEffortFromValue(v int32) (*ChatReasoningEffort, error) {
	ev := ChatReasoningEffort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChatReasoningEffort: valid values are %v", v, AllowedChatReasoningEffortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChatReasoningEffort) IsValid() bool {
	for _, existing := range AllowedChatReasoningEffortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChatReasoningEffort value
func (v ChatReasoningEffort) Ptr() *ChatReasoningEffort {
	return &v
}

type NullableChatReasoningEffort struct {
	value *ChatReasoningEffort
	isSet bool
}

func (v NullableChatReasoningEffort) Get() *ChatReasoningEffort {
	return v.value
}

func (v *NullableChatReasoningEffort) Set(val *ChatReasoningEffort) {
	v.value = val
	v.isSet = true
}

func (v NullableChatReasoningEffort) IsSet() bool {
	return v.isSet
}

func (v *NullableChatReasoningEffort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatReasoningEffort(val *ChatReasoningEffort) *NullableChatReasoningEffort {
	return &NullableChatReasoningEffort{value: val, isSet: true}
}

func (v NullableChatReasoningEffort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatReasoningEffort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

