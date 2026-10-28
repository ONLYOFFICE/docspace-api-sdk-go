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

// EmployeeStatus [1 - Active, 2 - Terminated, 4 - Pending, 5 - Default, 7 - All]
type EmployeeStatus int32

// List of EmployeeStatus
const (
	EMPLOYEESTATUS_Active EmployeeStatus = 1
	EMPLOYEESTATUS_Terminated EmployeeStatus = 2
	EMPLOYEESTATUS_Pending EmployeeStatus = 4
	EMPLOYEESTATUS_Default EmployeeStatus = 5
	EMPLOYEESTATUS_All EmployeeStatus = 7
)

// All allowed values of EmployeeStatus enum
var AllowedEmployeeStatusEnumValues = []EmployeeStatus{
	1,
	2,
	4,
	5,
	7,
}

func (v *EmployeeStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmployeeStatus(value)
	for _, existing := range AllowedEmployeeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmployeeStatus", value)
}

// NewEmployeeStatusFromValue returns a pointer to a valid EmployeeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmployeeStatusFromValue(v int32) (*EmployeeStatus, error) {
	ev := EmployeeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmployeeStatus: valid values are %v", v, AllowedEmployeeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmployeeStatus) IsValid() bool {
	for _, existing := range AllowedEmployeeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmployeeStatus value
func (v EmployeeStatus) Ptr() *EmployeeStatus {
	return &v
}

type NullableEmployeeStatus struct {
	value *EmployeeStatus
	isSet bool
}

func (v NullableEmployeeStatus) Get() *EmployeeStatus {
	return v.value
}

func (v *NullableEmployeeStatus) Set(val *EmployeeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeStatus(val *EmployeeStatus) *NullableEmployeeStatus {
	return &NullableEmployeeStatus{value: val, isSet: true}
}

func (v NullableEmployeeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

