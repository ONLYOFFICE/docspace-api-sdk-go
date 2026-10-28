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

// AiImportMode Mode passed to `PromptsEngine.importBundle`.
type AiImportMode string

// List of AiImportMode
const (
	AIIMPORTMODE_REPLACE AiImportMode = "replace"
	AIIMPORTMODE_MERGE AiImportMode = "merge"
)

// All allowed values of AiImportMode enum
var AllowedAiImportModeEnumValues = []AiImportMode{
	"replace",
	"merge",
}

func (v *AiImportMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiImportMode(value)
	for _, existing := range AllowedAiImportModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiImportMode", value)
}

// NewAiImportModeFromValue returns a pointer to a valid AiImportMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiImportModeFromValue(v string) (*AiImportMode, error) {
	ev := AiImportMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiImportMode: valid values are %v", v, AllowedAiImportModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiImportMode) IsValid() bool {
	for _, existing := range AllowedAiImportModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiImportMode value
func (v AiImportMode) Ptr() *AiImportMode {
	return &v
}

type NullableAiImportMode struct {
	value *AiImportMode
	isSet bool
}

func (v NullableAiImportMode) Get() *AiImportMode {
	return v.value
}

func (v *NullableAiImportMode) Set(val *AiImportMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImportMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImportMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImportMode(val *AiImportMode) *NullableAiImportMode {
	return &NullableAiImportMode{value: val, isSet: true}
}

func (v NullableAiImportMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImportMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

