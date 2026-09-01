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

// EditorType [0 - Desktop, 1 - Mobile, 2 - Embedded]
type EditorType int32

// List of EditorType
const (
	EDITORTYPE_Desktop EditorType = 0
	EDITORTYPE_Mobile EditorType = 1
	EDITORTYPE_Embedded EditorType = 2
)

// All allowed values of EditorType enum
var AllowedEditorTypeEnumValues = []EditorType{
	0,
	1,
	2,
}

func (v *EditorType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EditorType(value)
	for _, existing := range AllowedEditorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EditorType", value)
}

// NewEditorTypeFromValue returns a pointer to a valid EditorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEditorTypeFromValue(v int32) (*EditorType, error) {
	ev := EditorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EditorType: valid values are %v", v, AllowedEditorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EditorType) IsValid() bool {
	for _, existing := range AllowedEditorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EditorType value
func (v EditorType) Ptr() *EditorType {
	return &v
}

type NullableEditorType struct {
	value *EditorType
	isSet bool
}

func (v NullableEditorType) Get() *EditorType {
	return v.value
}

func (v *NullableEditorType) Set(val *EditorType) {
	v.value = val
	v.isSet = true
}

func (v NullableEditorType) IsSet() bool {
	return v.isSet
}

func (v *NullableEditorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditorType(val *EditorType) *NullableEditorType {
	return &NullableEditorType{value: val, isSet: true}
}

func (v NullableEditorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

