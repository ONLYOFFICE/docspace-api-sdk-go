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

// EmployeeType [All - All, RoomAdmin - Room admin, Guest - Guest, DocSpaceAdmin - DocSpace admin, User - User]
type EmployeeType string

// List of EmployeeType
const (
	EMPLOYEETYPE_ALL EmployeeType = "All"
	EMPLOYEETYPE_ROOM_ADMIN EmployeeType = "RoomAdmin"
	EMPLOYEETYPE_GUEST EmployeeType = "Guest"
	EMPLOYEETYPE_DOC_SPACE_ADMIN EmployeeType = "DocSpaceAdmin"
	EMPLOYEETYPE_USER EmployeeType = "User"
)

// All allowed values of EmployeeType enum
var AllowedEmployeeTypeEnumValues = []EmployeeType{
	"All",
	"RoomAdmin",
	"Guest",
	"DocSpaceAdmin",
	"User",
}

func (v *EmployeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmployeeType(value)
	for _, existing := range AllowedEmployeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmployeeType", value)
}

// NewEmployeeTypeFromValue returns a pointer to a valid EmployeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmployeeTypeFromValue(v string) (*EmployeeType, error) {
	ev := EmployeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmployeeType: valid values are %v", v, AllowedEmployeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmployeeType) IsValid() bool {
	for _, existing := range AllowedEmployeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmployeeType value
func (v EmployeeType) Ptr() *EmployeeType {
	return &v
}

type NullableEmployeeType struct {
	value *EmployeeType
	isSet bool
}

func (v NullableEmployeeType) Get() *EmployeeType {
	return v.value
}

func (v *NullableEmployeeType) Set(val *EmployeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeType(val *EmployeeType) *NullableEmployeeType {
	return &NullableEmployeeType{value: val, isSet: true}
}

func (v NullableEmployeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

