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

// EngineType [0 - None, 1 - Exa, 2 - PortalAi]
type EngineType int32

// List of EngineType
const (
	ENGINETYPE_None EngineType = 0
	ENGINETYPE_Exa EngineType = 1
	ENGINETYPE_PortalAi EngineType = 2
)

// All allowed values of EngineType enum
var AllowedEngineTypeEnumValues = []EngineType{
	0,
	1,
	2,
}

func (v *EngineType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EngineType(value)
	for _, existing := range AllowedEngineTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EngineType", value)
}

// NewEngineTypeFromValue returns a pointer to a valid EngineType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEngineTypeFromValue(v int32) (*EngineType, error) {
	ev := EngineType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EngineType: valid values are %v", v, AllowedEngineTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EngineType) IsValid() bool {
	for _, existing := range AllowedEngineTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineType value
func (v EngineType) Ptr() *EngineType {
	return &v
}

type NullableEngineType struct {
	value *EngineType
	isSet bool
}

func (v NullableEngineType) Get() *EngineType {
	return v.value
}

func (v *NullableEngineType) Set(val *EngineType) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineType) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineType(val *EngineType) *NullableEngineType {
	return &NullableEngineType{value: val, isSet: true}
}

func (v NullableEngineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

