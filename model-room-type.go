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

// RoomType [1 - Form filling room, 2 - Collaboration room, 5 - Custom room, 6 - Public room, 8 - Virtual data room, 9 - AI Room]
type RoomType int32

// List of RoomType
const (
	ROOMTYPE_FillingFormsRoom RoomType = 1
	ROOMTYPE_EditingRoom RoomType = 2
	ROOMTYPE_CustomRoom RoomType = 5
	ROOMTYPE_PublicRoom RoomType = 6
	ROOMTYPE_VirtualDataRoom RoomType = 8
	ROOMTYPE_AiRoom RoomType = 9
)

// All allowed values of RoomType enum
var AllowedRoomTypeEnumValues = []RoomType{
	1,
	2,
	5,
	6,
	8,
	9,
}

func (v *RoomType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomType(value)
	for _, existing := range AllowedRoomTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomType", value)
}

// NewRoomTypeFromValue returns a pointer to a valid RoomType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomTypeFromValue(v int32) (*RoomType, error) {
	ev := RoomType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomType: valid values are %v", v, AllowedRoomTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomType) IsValid() bool {
	for _, existing := range AllowedRoomTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoomType value
func (v RoomType) Ptr() *RoomType {
	return &v
}

type NullableRoomType struct {
	value *RoomType
	isSet bool
}

func (v NullableRoomType) Get() *RoomType {
	return v.value
}

func (v *NullableRoomType) Set(val *RoomType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomType(val *RoomType) *NullableRoomType {
	return &NullableRoomType{value: val, isSet: true}
}

func (v NullableRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

