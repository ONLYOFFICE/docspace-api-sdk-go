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

// CheckDestFolderResult [0 - All allowed, 1 - Part allowed, 2 - None allowed]
type CheckDestFolderResult int32

// List of CheckDestFolderResult
const (
	CHECKDESTFOLDERRESULT_AllAllowed CheckDestFolderResult = 0
	CHECKDESTFOLDERRESULT_PartAllowed CheckDestFolderResult = 1
	CHECKDESTFOLDERRESULT_NoneAllowed CheckDestFolderResult = 2
)

// All allowed values of CheckDestFolderResult enum
var AllowedCheckDestFolderResultEnumValues = []CheckDestFolderResult{
	0,
	1,
	2,
}

func (v *CheckDestFolderResult) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CheckDestFolderResult(value)
	for _, existing := range AllowedCheckDestFolderResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CheckDestFolderResult", value)
}

// NewCheckDestFolderResultFromValue returns a pointer to a valid CheckDestFolderResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCheckDestFolderResultFromValue(v int32) (*CheckDestFolderResult, error) {
	ev := CheckDestFolderResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CheckDestFolderResult: valid values are %v", v, AllowedCheckDestFolderResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CheckDestFolderResult) IsValid() bool {
	for _, existing := range AllowedCheckDestFolderResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckDestFolderResult value
func (v CheckDestFolderResult) Ptr() *CheckDestFolderResult {
	return &v
}

type NullableCheckDestFolderResult struct {
	value *CheckDestFolderResult
	isSet bool
}

func (v NullableCheckDestFolderResult) Get() *CheckDestFolderResult {
	return v.value
}

func (v *NullableCheckDestFolderResult) Set(val *CheckDestFolderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDestFolderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDestFolderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDestFolderResult(val *CheckDestFolderResult) *NullableCheckDestFolderResult {
	return &NullableCheckDestFolderResult{value: val, isSet: true}
}

func (v NullableCheckDestFolderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDestFolderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

