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

// FileStatus [0 - None, 1 - Is editing, 2 - Is new, 4 - Is converting, 8 - Is original, 16 - Is editing alone, 32 - Is favorite, 64 - Is template, 128 - Is fill form draft, 256 - Is completed form]
type FileStatus int32

// List of FileStatus
const (
	FILESTATUS_None FileStatus = 0
	FILESTATUS_IsEditing FileStatus = 1
	FILESTATUS_IsNew FileStatus = 2
	FILESTATUS_IsConverting FileStatus = 4
	FILESTATUS_IsOriginal FileStatus = 8
	FILESTATUS_IsEditingAlone FileStatus = 16
	FILESTATUS_IsFavorite FileStatus = 32
	FILESTATUS_IsTemplate FileStatus = 64
	FILESTATUS_IsFillFormDraft FileStatus = 128
	FILESTATUS_IsCompletedForm FileStatus = 256
)

// All allowed values of FileStatus enum
var AllowedFileStatusEnumValues = []FileStatus{
	0,
	1,
	2,
	4,
	8,
	16,
	32,
	64,
	128,
	256,
}

func (v *FileStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileStatus(value)
	for _, existing := range AllowedFileStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileStatus", value)
}

// NewFileStatusFromValue returns a pointer to a valid FileStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileStatusFromValue(v int32) (*FileStatus, error) {
	ev := FileStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileStatus: valid values are %v", v, AllowedFileStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileStatus) IsValid() bool {
	for _, existing := range AllowedFileStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileStatus value
func (v FileStatus) Ptr() *FileStatus {
	return &v
}

type NullableFileStatus struct {
	value *FileStatus
	isSet bool
}

func (v NullableFileStatus) Get() *FileStatus {
	return v.value
}

func (v *NullableFileStatus) Set(val *FileStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFileStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFileStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileStatus(val *FileStatus) *NullableFileStatus {
	return &NullableFileStatus{value: val, isSet: true}
}

func (v NullableFileStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

