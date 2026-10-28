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

// FormFillingStatus [0 - None, 1 - Draft, 2 - You turn, 3 - In progress, 4 - Complete, 5 - Stoped]
type FormFillingStatus int32

// List of FormFillingStatus
const (
	FORMFILLINGSTATUS_None FormFillingStatus = 0
	FORMFILLINGSTATUS_Draft FormFillingStatus = 1
	FORMFILLINGSTATUS_YouTurn FormFillingStatus = 2
	FORMFILLINGSTATUS_InProgress FormFillingStatus = 3
	FORMFILLINGSTATUS_Complete FormFillingStatus = 4
	FORMFILLINGSTATUS_Stoped FormFillingStatus = 5
)

// All allowed values of FormFillingStatus enum
var AllowedFormFillingStatusEnumValues = []FormFillingStatus{
	0,
	1,
	2,
	3,
	4,
	5,
}

func (v *FormFillingStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FormFillingStatus(value)
	for _, existing := range AllowedFormFillingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FormFillingStatus", value)
}

// NewFormFillingStatusFromValue returns a pointer to a valid FormFillingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormFillingStatusFromValue(v int32) (*FormFillingStatus, error) {
	ev := FormFillingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormFillingStatus: valid values are %v", v, AllowedFormFillingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormFillingStatus) IsValid() bool {
	for _, existing := range AllowedFormFillingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormFillingStatus value
func (v FormFillingStatus) Ptr() *FormFillingStatus {
	return &v
}

type NullableFormFillingStatus struct {
	value *FormFillingStatus
	isSet bool
}

func (v NullableFormFillingStatus) Get() *FormFillingStatus {
	return v.value
}

func (v *NullableFormFillingStatus) Set(val *FormFillingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFillingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFillingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFillingStatus(val *FormFillingStatus) *NullableFormFillingStatus {
	return &NullableFormFillingStatus{value: val, isSet: true}
}

func (v NullableFormFillingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFillingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

