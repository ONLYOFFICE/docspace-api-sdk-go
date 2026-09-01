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

// Status [0 - Ok, 1 - Invalid, 2 - Expired, 3 - Required password, 4 - Invalid password, 5 - External access denied]
type Status int32

// List of Status
const (
	STATUS_Ok Status = 0
	STATUS_Invalid Status = 1
	STATUS_Expired Status = 2
	STATUS_RequiredPassword Status = 3
	STATUS_InvalidPassword Status = 4
	STATUS_ExternalAccessDenied Status = 5
)

// All allowed values of Status enum
var AllowedStatusEnumValues = []Status{
	0,
	1,
	2,
	3,
	4,
	5,
}

func (v *Status) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status(value)
	for _, existing := range AllowedStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status", value)
}

// NewStatusFromValue returns a pointer to a valid Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusFromValue(v int32) (*Status, error) {
	ev := Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Status: valid values are %v", v, AllowedStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Status) IsValid() bool {
	for _, existing := range AllowedStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status value
func (v Status) Ptr() *Status {
	return &v
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

