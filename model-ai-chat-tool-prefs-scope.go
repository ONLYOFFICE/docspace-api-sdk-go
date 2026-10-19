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

// AiChatToolPrefsScope Which disabled tool list a read or write addresses: the caller's own (`user`, what an absent value means) or the one an administrator keeps for everyone (`tenant`), which every user's reads fold in.
type AiChatToolPrefsScope string

// List of AiChatToolPrefsScope
const (
	AICHATTOOLPREFSSCOPE_USER AiChatToolPrefsScope = "user"
	AICHATTOOLPREFSSCOPE_TENANT AiChatToolPrefsScope = "tenant"
)

// All allowed values of AiChatToolPrefsScope enum
var AllowedAiChatToolPrefsScopeEnumValues = []AiChatToolPrefsScope{
	"user",
	"tenant",
}

func (v *AiChatToolPrefsScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiChatToolPrefsScope(value)
	for _, existing := range AllowedAiChatToolPrefsScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiChatToolPrefsScope", value)
}

// NewAiChatToolPrefsScopeFromValue returns a pointer to a valid AiChatToolPrefsScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiChatToolPrefsScopeFromValue(v string) (*AiChatToolPrefsScope, error) {
	ev := AiChatToolPrefsScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiChatToolPrefsScope: valid values are %v", v, AllowedAiChatToolPrefsScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiChatToolPrefsScope) IsValid() bool {
	for _, existing := range AllowedAiChatToolPrefsScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiChatToolPrefsScope value
func (v AiChatToolPrefsScope) Ptr() *AiChatToolPrefsScope {
	return &v
}

type NullableAiChatToolPrefsScope struct {
	value *AiChatToolPrefsScope
	isSet bool
}

func (v NullableAiChatToolPrefsScope) Get() *AiChatToolPrefsScope {
	return v.value
}

func (v *NullableAiChatToolPrefsScope) Set(val *AiChatToolPrefsScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAiChatToolPrefsScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAiChatToolPrefsScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiChatToolPrefsScope(val *AiChatToolPrefsScope) *NullableAiChatToolPrefsScope {
	return &NullableAiChatToolPrefsScope{value: val, isSet: true}
}

func (v NullableAiChatToolPrefsScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiChatToolPrefsScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

