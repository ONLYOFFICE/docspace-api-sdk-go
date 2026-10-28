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

// StartFillingMode [0 - None, 1 - Share to fill out, 2 - Start filling, 3 - Start filling form room]
type StartFillingMode int32

// List of StartFillingMode
const (
	STARTFILLINGMODE_None StartFillingMode = 0
	STARTFILLINGMODE_ShareToFillOut StartFillingMode = 1
	STARTFILLINGMODE_StartFilling StartFillingMode = 2
	STARTFILLINGMODE_StartFillingFormRoom StartFillingMode = 3
)

// All allowed values of StartFillingMode enum
var AllowedStartFillingModeEnumValues = []StartFillingMode{
	0,
	1,
	2,
	3,
}

func (v *StartFillingMode) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StartFillingMode(value)
	for _, existing := range AllowedStartFillingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StartFillingMode", value)
}

// NewStartFillingModeFromValue returns a pointer to a valid StartFillingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStartFillingModeFromValue(v int32) (*StartFillingMode, error) {
	ev := StartFillingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StartFillingMode: valid values are %v", v, AllowedStartFillingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StartFillingMode) IsValid() bool {
	for _, existing := range AllowedStartFillingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StartFillingMode value
func (v StartFillingMode) Ptr() *StartFillingMode {
	return &v
}

type NullableStartFillingMode struct {
	value *StartFillingMode
	isSet bool
}

func (v NullableStartFillingMode) Get() *StartFillingMode {
	return v.value
}

func (v *NullableStartFillingMode) Set(val *StartFillingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableStartFillingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableStartFillingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartFillingMode(val *StartFillingMode) *NullableStartFillingMode {
	return &NullableStartFillingMode{value: val, isSet: true}
}

func (v NullableStartFillingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartFillingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

