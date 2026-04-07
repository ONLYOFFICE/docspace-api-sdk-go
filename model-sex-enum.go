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

// SexEnum [0 - Female, 1 - Male]
type SexEnum int32

// List of SexEnum
const (
	SEXENUM_Female SexEnum = 0
	SEXENUM_Male SexEnum = 1
)

// All allowed values of SexEnum enum
var AllowedSexEnumEnumValues = []SexEnum{
	0,
	1,
}

func (v *SexEnum) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SexEnum(value)
	for _, existing := range AllowedSexEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SexEnum", value)
}

// NewSexEnumFromValue returns a pointer to a valid SexEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSexEnumFromValue(v int32) (*SexEnum, error) {
	ev := SexEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SexEnum: valid values are %v", v, AllowedSexEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SexEnum) IsValid() bool {
	for _, existing := range AllowedSexEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SexEnum value
func (v SexEnum) Ptr() *SexEnum {
	return &v
}

type NullableSexEnum struct {
	value *SexEnum
	isSet bool
}

func (v NullableSexEnum) Get() *SexEnum {
	return v.value
}

func (v *NullableSexEnum) Set(val *SexEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSexEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSexEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSexEnum(val *SexEnum) *NullableSexEnum {
	return &NullableSexEnum{value: val, isSet: true}
}

func (v NullableSexEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSexEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

