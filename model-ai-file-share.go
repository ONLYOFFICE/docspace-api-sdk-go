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

// AiFileShare [0 - None, 1 - Read and write, 2 - Read, 3 - Restrict, 4 - Varies, 5 - Review, 6 - Comment, 7 - Fill forms, 8 - Custom filter, 9 - Room manager, 10 - Editing, 11 - Content creator]
type AiFileShare int32

// List of AiFileShare
const (
	AIFILESHARE_None AiFileShare = 0
	AIFILESHARE_ReadWrite AiFileShare = 1
	AIFILESHARE_Read AiFileShare = 2
	AIFILESHARE_Restrict AiFileShare = 3
	AIFILESHARE_Varies AiFileShare = 4
	AIFILESHARE_Review AiFileShare = 5
	AIFILESHARE_Comment AiFileShare = 6
	AIFILESHARE_FillForms AiFileShare = 7
	AIFILESHARE_CustomFilter AiFileShare = 8
	AIFILESHARE_RoomManager AiFileShare = 9
	AIFILESHARE_Editing AiFileShare = 10
	AIFILESHARE_ContentCreator AiFileShare = 11
)

// All allowed values of AiFileShare enum
var AllowedAiFileShareEnumValues = []AiFileShare{
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

func (v *AiFileShare) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiFileShare(value)
	for _, existing := range AllowedAiFileShareEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiFileShare", value)
}

// NewAiFileShareFromValue returns a pointer to a valid AiFileShare
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiFileShareFromValue(v int32) (*AiFileShare, error) {
	ev := AiFileShare(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiFileShare: valid values are %v", v, AllowedAiFileShareEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiFileShare) IsValid() bool {
	for _, existing := range AllowedAiFileShareEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiFileShare value
func (v AiFileShare) Ptr() *AiFileShare {
	return &v
}

type NullableAiFileShare struct {
	value *AiFileShare
	isSet bool
}

func (v NullableAiFileShare) Get() *AiFileShare {
	return v.value
}

func (v *NullableAiFileShare) Set(val *AiFileShare) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileShare) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileShare(val *AiFileShare) *NullableAiFileShare {
	return &NullableAiFileShare{value: val, isSet: true}
}

func (v NullableAiFileShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

