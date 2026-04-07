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

// ApplyFilterOption [0 - All, 1 - Files, 2 - Folders]
type ApplyFilterOption int32

// List of ApplyFilterOption
const (
	APPLYFILTEROPTION_All ApplyFilterOption = 0
	APPLYFILTEROPTION_Files ApplyFilterOption = 1
	APPLYFILTEROPTION_Folders ApplyFilterOption = 2
)

// All allowed values of ApplyFilterOption enum
var AllowedApplyFilterOptionEnumValues = []ApplyFilterOption{
	0,
	1,
	2,
}

func (v *ApplyFilterOption) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplyFilterOption(value)
	for _, existing := range AllowedApplyFilterOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplyFilterOption", value)
}

// NewApplyFilterOptionFromValue returns a pointer to a valid ApplyFilterOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplyFilterOptionFromValue(v int32) (*ApplyFilterOption, error) {
	ev := ApplyFilterOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplyFilterOption: valid values are %v", v, AllowedApplyFilterOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplyFilterOption) IsValid() bool {
	for _, existing := range AllowedApplyFilterOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplyFilterOption value
func (v ApplyFilterOption) Ptr() *ApplyFilterOption {
	return &v
}

type NullableApplyFilterOption struct {
	value *ApplyFilterOption
	isSet bool
}

func (v NullableApplyFilterOption) Get() *ApplyFilterOption {
	return v.value
}

func (v *NullableApplyFilterOption) Set(val *ApplyFilterOption) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyFilterOption) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyFilterOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyFilterOption(val *ApplyFilterOption) *NullableApplyFilterOption {
	return &NullableApplyFilterOption{value: val, isSet: true}
}

func (v NullableApplyFilterOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyFilterOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

