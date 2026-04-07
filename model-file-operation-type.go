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

// FileOperationType [0 - Move, 1 - Copy, 2 - Delete, 3 - Download, 4 - MarkAsRead, 5 - Import, 6 - Convert, 7 - Duplicate]
type FileOperationType int32

// List of FileOperationType
const (
	FILEOPERATIONTYPE_Move FileOperationType = 0
	FILEOPERATIONTYPE_Copy FileOperationType = 1
	FILEOPERATIONTYPE_Delete FileOperationType = 2
	FILEOPERATIONTYPE_Download FileOperationType = 3
	FILEOPERATIONTYPE_MarkAsRead FileOperationType = 4
	FILEOPERATIONTYPE_Import FileOperationType = 5
	FILEOPERATIONTYPE_Convert FileOperationType = 6
	FILEOPERATIONTYPE_Duplicate FileOperationType = 7
)

// All allowed values of FileOperationType enum
var AllowedFileOperationTypeEnumValues = []FileOperationType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

func (v *FileOperationType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileOperationType(value)
	for _, existing := range AllowedFileOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileOperationType", value)
}

// NewFileOperationTypeFromValue returns a pointer to a valid FileOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileOperationTypeFromValue(v int32) (*FileOperationType, error) {
	ev := FileOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileOperationType: valid values are %v", v, AllowedFileOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileOperationType) IsValid() bool {
	for _, existing := range AllowedFileOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileOperationType value
func (v FileOperationType) Ptr() *FileOperationType {
	return &v
}

type NullableFileOperationType struct {
	value *FileOperationType
	isSet bool
}

func (v NullableFileOperationType) Get() *FileOperationType {
	return v.value
}

func (v *NullableFileOperationType) Set(val *FileOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileOperationType(val *FileOperationType) *NullableFileOperationType {
	return &NullableFileOperationType{value: val, isSet: true}
}

func (v NullableFileOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

