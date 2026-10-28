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

// FileType [0 - Unknown, 1 - Archive, 2 - Video, 3 - Audio, 4 - Image, 5 - Spreadsheet, 6 - Presentation, 7 - Document, 10 - Pdf, 11 - Diagram]
type FileType int32

// List of FileType
const (
	FILETYPE_Unknown FileType = 0
	FILETYPE_Archive FileType = 1
	FILETYPE_Video FileType = 2
	FILETYPE_Audio FileType = 3
	FILETYPE_Image FileType = 4
	FILETYPE_Spreadsheet FileType = 5
	FILETYPE_Presentation FileType = 6
	FILETYPE_Document FileType = 7
	FILETYPE_Pdf FileType = 10
	FILETYPE_Diagram FileType = 11
)

// All allowed values of FileType enum
var AllowedFileTypeEnumValues = []FileType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	10,
	11,
}

func (v *FileType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileType(value)
	for _, existing := range AllowedFileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileType", value)
}

// NewFileTypeFromValue returns a pointer to a valid FileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileTypeFromValue(v int32) (*FileType, error) {
	ev := FileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileType: valid values are %v", v, AllowedFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileType) IsValid() bool {
	for _, existing := range AllowedFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileType value
func (v FileType) Ptr() *FileType {
	return &v
}

type NullableFileType struct {
	value *FileType
	isSet bool
}

func (v NullableFileType) Get() *FileType {
	return v.value
}

func (v *NullableFileType) Set(val *FileType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileType(val *FileType) *NullableFileType {
	return &NullableFileType{value: val, isSet: true}
}

func (v NullableFileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

