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

// RegStatus []
type RegStatus int32

// List of RegStatus
const (
	REGSTATUS_unlinked RegStatus = 0
	REGSTATUS_linked RegStatus = 1
	REGSTATUS_linking RegStatus = 2
)

// All allowed values of RegStatus enum
var AllowedRegStatusEnumValues = []RegStatus{
	0,
	1,
	2,
}

func (v *RegStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegStatus(value)
	for _, existing := range AllowedRegStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegStatus", value)
}

// NewRegStatusFromValue returns a pointer to a valid RegStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegStatusFromValue(v int32) (*RegStatus, error) {
	ev := RegStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegStatus: valid values are %v", v, AllowedRegStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegStatus) IsValid() bool {
	for _, existing := range AllowedRegStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegStatus value
func (v RegStatus) Ptr() *RegStatus {
	return &v
}

type NullableRegStatus struct {
	value *RegStatus
	isSet bool
}

func (v NullableRegStatus) Get() *RegStatus {
	return v.value
}

func (v *NullableRegStatus) Set(val *RegStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegStatus(val *RegStatus) *NullableRegStatus {
	return &NullableRegStatus{value: val, isSet: true}
}

func (v NullableRegStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

