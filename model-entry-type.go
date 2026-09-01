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

// EntryType [0 - None, 1 - File, 2 - Folder, 23 - User, 24 - Group, 25 - Room, 26 - Tag, 27 - Agent]
type EntryType int32

// List of EntryType
const (
	ENTRYTYPE_None EntryType = 0
	ENTRYTYPE_File EntryType = 1
	ENTRYTYPE_Folder EntryType = 2
	ENTRYTYPE_User EntryType = 23
	ENTRYTYPE_Group EntryType = 24
	ENTRYTYPE_Room EntryType = 25
	ENTRYTYPE_Tag EntryType = 26
	ENTRYTYPE_Agent EntryType = 27
)

// All allowed values of EntryType enum
var AllowedEntryTypeEnumValues = []EntryType{
	0,
	1,
	2,
	23,
	24,
	25,
	26,
	27,
}

func (v *EntryType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntryType(value)
	for _, existing := range AllowedEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntryType", value)
}

// NewEntryTypeFromValue returns a pointer to a valid EntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntryTypeFromValue(v int32) (*EntryType, error) {
	ev := EntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntryType: valid values are %v", v, AllowedEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntryType) IsValid() bool {
	for _, existing := range AllowedEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntryType value
func (v EntryType) Ptr() *EntryType {
	return &v
}

type NullableEntryType struct {
	value *EntryType
	isSet bool
}

func (v NullableEntryType) Get() *EntryType {
	return v.value
}

func (v *NullableEntryType) Set(val *EntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryType(val *EntryType) *NullableEntryType {
	return &NullableEntryType{value: val, isSet: true}
}

func (v NullableEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

