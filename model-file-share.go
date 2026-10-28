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

// FileShare [0 - None, 1 - Read and write, 2 - Read, 3 - Restrict, 4 - Varies, 5 - Review, 6 - Comment, 7 - Fill forms, 8 - Custom filter, 9 - Room manager, 10 - Editing, 11 - Content creator]
type FileShare int32

// List of FileShare
const (
	FILESHARE_None FileShare = 0
	FILESHARE_ReadWrite FileShare = 1
	FILESHARE_Read FileShare = 2
	FILESHARE_Restrict FileShare = 3
	FILESHARE_Varies FileShare = 4
	FILESHARE_Review FileShare = 5
	FILESHARE_Comment FileShare = 6
	FILESHARE_FillForms FileShare = 7
	FILESHARE_CustomFilter FileShare = 8
	FILESHARE_RoomManager FileShare = 9
	FILESHARE_Editing FileShare = 10
	FILESHARE_ContentCreator FileShare = 11
)

// All allowed values of FileShare enum
var AllowedFileShareEnumValues = []FileShare{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
}

func (v *FileShare) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileShare(value)
	for _, existing := range AllowedFileShareEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileShare", value)
}

// NewFileShareFromValue returns a pointer to a valid FileShare
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileShareFromValue(v int32) (*FileShare, error) {
	ev := FileShare(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileShare: valid values are %v", v, AllowedFileShareEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileShare) IsValid() bool {
	for _, existing := range AllowedFileShareEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileShare value
func (v FileShare) Ptr() *FileShare {
	return &v
}

type NullableFileShare struct {
	value *FileShare
	isSet bool
}

func (v NullableFileShare) Get() *FileShare {
	return v.value
}

func (v *NullableFileShare) Set(val *FileShare) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShare) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShare(val *FileShare) *NullableFileShare {
	return &NullableFileShare{value: val, isSet: true}
}

func (v NullableFileShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

