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

// FileEntryType [1 - Folder, 2 - File]
type FileEntryType int32

// List of FileEntryType
const (
	FILEENTRYTYPE_Folder FileEntryType = 1
	FILEENTRYTYPE_File FileEntryType = 2
)

// All allowed values of FileEntryType enum
var AllowedFileEntryTypeEnumValues = []FileEntryType{
	1,
	2,
}

func (v *FileEntryType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileEntryType(value)
	for _, existing := range AllowedFileEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileEntryType", value)
}

// NewFileEntryTypeFromValue returns a pointer to a valid FileEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileEntryTypeFromValue(v int32) (*FileEntryType, error) {
	ev := FileEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileEntryType: valid values are %v", v, AllowedFileEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileEntryType) IsValid() bool {
	for _, existing := range AllowedFileEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileEntryType value
func (v FileEntryType) Ptr() *FileEntryType {
	return &v
}

type NullableFileEntryType struct {
	value *FileEntryType
	isSet bool
}

func (v NullableFileEntryType) Get() *FileEntryType {
	return v.value
}

func (v *NullableFileEntryType) Set(val *FileEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileEntryType(val *FileEntryType) *NullableFileEntryType {
	return &NullableFileEntryType{value: val, isSet: true}
}

func (v NullableFileEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

