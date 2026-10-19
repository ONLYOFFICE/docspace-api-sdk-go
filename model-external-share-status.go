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

// ExternalShareStatus [0 - Ok, 1 - Invalid, 2 - Expired, 3 - Required password, 4 - Invalid password, 5 - External access denied]
type ExternalShareStatus int32

// List of ExternalShareStatus
const (
	EXTERNALSHARESTATUS_Ok ExternalShareStatus = 0
	EXTERNALSHARESTATUS_Invalid ExternalShareStatus = 1
	EXTERNALSHARESTATUS_Expired ExternalShareStatus = 2
	EXTERNALSHARESTATUS_RequiredPassword ExternalShareStatus = 3
	EXTERNALSHARESTATUS_InvalidPassword ExternalShareStatus = 4
	EXTERNALSHARESTATUS_ExternalAccessDenied ExternalShareStatus = 5
)

// All allowed values of ExternalShareStatus enum
var AllowedExternalShareStatusEnumValues = []ExternalShareStatus{
	0,
	1,
	2,
	3,
	4,
	5,
}

func (v *ExternalShareStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalShareStatus(value)
	for _, existing := range AllowedExternalShareStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalShareStatus", value)
}

// NewExternalShareStatusFromValue returns a pointer to a valid ExternalShareStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalShareStatusFromValue(v int32) (*ExternalShareStatus, error) {
	ev := ExternalShareStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalShareStatus: valid values are %v", v, AllowedExternalShareStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalShareStatus) IsValid() bool {
	for _, existing := range AllowedExternalShareStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalShareStatus value
func (v ExternalShareStatus) Ptr() *ExternalShareStatus {
	return &v
}

type NullableExternalShareStatus struct {
	value *ExternalShareStatus
	isSet bool
}

func (v NullableExternalShareStatus) Get() *ExternalShareStatus {
	return v.value
}

func (v *NullableExternalShareStatus) Set(val *ExternalShareStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalShareStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalShareStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalShareStatus(val *ExternalShareStatus) *NullableExternalShareStatus {
	return &NullableExternalShareStatus{value: val, isSet: true}
}

func (v NullableExternalShareStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalShareStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

