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

// SortedByType [0 - Date and time, 1 - AZ, 2 - Size, 3 - Author, 4 - Type, 5 - New, 6 - Date and time creation, 7 - Room type, 8 - Tags, 9 - Room, 10 - Custom order, 11 - Last opened, 12 - Used space]
type SortedByType int32

// List of SortedByType
const (
	SORTEDBYTYPE_DateAndTime SortedByType = 0
	SORTEDBYTYPE_AZ SortedByType = 1
	SORTEDBYTYPE_Size SortedByType = 2
	SORTEDBYTYPE_Author SortedByType = 3
	SORTEDBYTYPE_Type SortedByType = 4
	SORTEDBYTYPE_New SortedByType = 5
	SORTEDBYTYPE_DateAndTimeCreation SortedByType = 6
	SORTEDBYTYPE_RoomType SortedByType = 7
	SORTEDBYTYPE_Tags SortedByType = 8
	SORTEDBYTYPE_Room SortedByType = 9
	SORTEDBYTYPE_CustomOrder SortedByType = 10
	SORTEDBYTYPE_LastOpened SortedByType = 11
	SORTEDBYTYPE_UsedSpace SortedByType = 12
)

// All allowed values of SortedByType enum
var AllowedSortedByTypeEnumValues = []SortedByType{
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
	12,
}

func (v *SortedByType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortedByType(value)
	for _, existing := range AllowedSortedByTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortedByType", value)
}

// NewSortedByTypeFromValue returns a pointer to a valid SortedByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortedByTypeFromValue(v int32) (*SortedByType, error) {
	ev := SortedByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortedByType: valid values are %v", v, AllowedSortedByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortedByType) IsValid() bool {
	for _, existing := range AllowedSortedByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortedByType value
func (v SortedByType) Ptr() *SortedByType {
	return &v
}

type NullableSortedByType struct {
	value *SortedByType
	isSet bool
}

func (v NullableSortedByType) Get() *SortedByType {
	return v.value
}

func (v *NullableSortedByType) Set(val *SortedByType) {
	v.value = val
	v.isSet = true
}

func (v NullableSortedByType) IsSet() bool {
	return v.isSet
}

func (v *NullableSortedByType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortedByType(val *SortedByType) *NullableSortedByType {
	return &NullableSortedByType{value: val, isSet: true}
}

func (v NullableSortedByType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortedByType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

