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

// CoEditingConfigMode [0 - Fast, 1 - Strict]
type CoEditingConfigMode int32

// List of CoEditingConfigMode
const (
	COEDITINGCONFIGMODE_Fast CoEditingConfigMode = 0
	COEDITINGCONFIGMODE_Strict CoEditingConfigMode = 1
)

// All allowed values of CoEditingConfigMode enum
var AllowedCoEditingConfigModeEnumValues = []CoEditingConfigMode{
	0,
	1,
}

func (v *CoEditingConfigMode) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CoEditingConfigMode(value)
	for _, existing := range AllowedCoEditingConfigModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CoEditingConfigMode", value)
}

// NewCoEditingConfigModeFromValue returns a pointer to a valid CoEditingConfigMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoEditingConfigModeFromValue(v int32) (*CoEditingConfigMode, error) {
	ev := CoEditingConfigMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CoEditingConfigMode: valid values are %v", v, AllowedCoEditingConfigModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoEditingConfigMode) IsValid() bool {
	for _, existing := range AllowedCoEditingConfigModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CoEditingConfigMode value
func (v CoEditingConfigMode) Ptr() *CoEditingConfigMode {
	return &v
}

type NullableCoEditingConfigMode struct {
	value *CoEditingConfigMode
	isSet bool
}

func (v NullableCoEditingConfigMode) Get() *CoEditingConfigMode {
	return v.value
}

func (v *NullableCoEditingConfigMode) Set(val *CoEditingConfigMode) {
	v.value = val
	v.isSet = true
}

func (v NullableCoEditingConfigMode) IsSet() bool {
	return v.isSet
}

func (v *NullableCoEditingConfigMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoEditingConfigMode(val *CoEditingConfigMode) *NullableCoEditingConfigMode {
	return &NullableCoEditingConfigMode{value: val, isSet: true}
}

func (v NullableCoEditingConfigMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoEditingConfigMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

