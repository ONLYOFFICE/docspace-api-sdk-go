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

// AiFileEntryType [1 - Folder, 2 - File]
type AiFileEntryType int32

// List of AiFileEntryType
const (
	AIFILEENTRYTYPE_Folder AiFileEntryType = 1
	AIFILEENTRYTYPE_File AiFileEntryType = 2
)

// All allowed values of AiFileEntryType enum
var AllowedAiFileEntryTypeEnumValues = []AiFileEntryType{
	1,
	2,
}

func (v *AiFileEntryType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiFileEntryType(value)
	for _, existing := range AllowedAiFileEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiFileEntryType", value)
}

// NewAiFileEntryTypeFromValue returns a pointer to a valid AiFileEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiFileEntryTypeFromValue(v int32) (*AiFileEntryType, error) {
	ev := AiFileEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiFileEntryType: valid values are %v", v, AllowedAiFileEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiFileEntryType) IsValid() bool {
	for _, existing := range AllowedAiFileEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiFileEntryType value
func (v AiFileEntryType) Ptr() *AiFileEntryType {
	return &v
}

type NullableAiFileEntryType struct {
	value *AiFileEntryType
	isSet bool
}

func (v NullableAiFileEntryType) Get() *AiFileEntryType {
	return v.value
}

func (v *NullableAiFileEntryType) Set(val *AiFileEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryType(val *AiFileEntryType) *NullableAiFileEntryType {
	return &NullableAiFileEntryType{value: val, isSet: true}
}

func (v NullableAiFileEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

