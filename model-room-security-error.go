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

// RoomSecurityError [0 - None, 1 - Form role blocking deletion]
type RoomSecurityError int32

// List of RoomSecurityError
const (
	ROOMSECURITYERROR_None RoomSecurityError = 0
	ROOMSECURITYERROR_FormRoleBlockingDeletion RoomSecurityError = 1
)

// All allowed values of RoomSecurityError enum
var AllowedRoomSecurityErrorEnumValues = []RoomSecurityError{
	0,
	1,
}

func (v *RoomSecurityError) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomSecurityError(value)
	for _, existing := range AllowedRoomSecurityErrorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomSecurityError", value)
}

// NewRoomSecurityErrorFromValue returns a pointer to a valid RoomSecurityError
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomSecurityErrorFromValue(v int32) (*RoomSecurityError, error) {
	ev := RoomSecurityError(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomSecurityError: valid values are %v", v, AllowedRoomSecurityErrorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomSecurityError) IsValid() bool {
	for _, existing := range AllowedRoomSecurityErrorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoomSecurityError value
func (v RoomSecurityError) Ptr() *RoomSecurityError {
	return &v
}

type NullableRoomSecurityError struct {
	value *RoomSecurityError
	isSet bool
}

func (v NullableRoomSecurityError) Get() *RoomSecurityError {
	return v.value
}

func (v *NullableRoomSecurityError) Set(val *RoomSecurityError) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomSecurityError) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomSecurityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomSecurityError(val *RoomSecurityError) *NullableRoomSecurityError {
	return &NullableRoomSecurityError{value: val, isSet: true}
}

func (v NullableRoomSecurityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomSecurityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

