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

// AiToolPermissionMode [0 - Ask, 1 - Auto, 2 - Allow]
type AiToolPermissionMode int32

// List of AiToolPermissionMode
const (
	AITOOLPERMISSIONMODE_Ask AiToolPermissionMode = 0
	AITOOLPERMISSIONMODE_Auto AiToolPermissionMode = 1
	AITOOLPERMISSIONMODE_Allow AiToolPermissionMode = 2
)

// All allowed values of AiToolPermissionMode enum
var AllowedAiToolPermissionModeEnumValues = []AiToolPermissionMode{
	0,
	1,
	2,
}

func (v *AiToolPermissionMode) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiToolPermissionMode(value)
	for _, existing := range AllowedAiToolPermissionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiToolPermissionMode", value)
}

// NewAiToolPermissionModeFromValue returns a pointer to a valid AiToolPermissionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiToolPermissionModeFromValue(v int32) (*AiToolPermissionMode, error) {
	ev := AiToolPermissionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiToolPermissionMode: valid values are %v", v, AllowedAiToolPermissionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiToolPermissionMode) IsValid() bool {
	for _, existing := range AllowedAiToolPermissionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiToolPermissionMode value
func (v AiToolPermissionMode) Ptr() *AiToolPermissionMode {
	return &v
}

type NullableAiToolPermissionMode struct {
	value *AiToolPermissionMode
	isSet bool
}

func (v NullableAiToolPermissionMode) Get() *AiToolPermissionMode {
	return v.value
}

func (v *NullableAiToolPermissionMode) Set(val *AiToolPermissionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolPermissionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolPermissionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolPermissionMode(val *AiToolPermissionMode) *NullableAiToolPermissionMode {
	return &NullableAiToolPermissionMode{value: val, isSet: true}
}

func (v NullableAiToolPermissionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolPermissionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

