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

// Role [0 - User, 1 - Assistant]
type Role int32

// List of Role
const (
	ROLE_User Role = 0
	ROLE_Assistant Role = 1
)

// All allowed values of Role enum
var AllowedRoleEnumValues = []Role{
	0,
	1,
}

func (v *Role) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Role(value)
	for _, existing := range AllowedRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Role", value)
}

// NewRoleFromValue returns a pointer to a valid Role
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleFromValue(v int32) (*Role, error) {
	ev := Role(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Role: valid values are %v", v, AllowedRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range AllowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Role value
func (v Role) Ptr() *Role {
	return &v
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

