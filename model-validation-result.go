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

// ValidationResult [0 - Ok, 1 - Invalid, 2 - Expired, 3 - Tariff limit, 4 - User existed, 5 - User excluded, 6 - Quota failed]
type ValidationResult int32

// List of ValidationResult
const (
	VALIDATIONRESULT_Ok ValidationResult = 0
	VALIDATIONRESULT_Invalid ValidationResult = 1
	VALIDATIONRESULT_Expired ValidationResult = 2
	VALIDATIONRESULT_TariffLimit ValidationResult = 3
	VALIDATIONRESULT_UserExisted ValidationResult = 4
	VALIDATIONRESULT_UserExcluded ValidationResult = 5
	VALIDATIONRESULT_QuotaFailed ValidationResult = 6
)

// All allowed values of ValidationResult enum
var AllowedValidationResultEnumValues = []ValidationResult{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

func (v *ValidationResult) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValidationResult(value)
	for _, existing := range AllowedValidationResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValidationResult", value)
}

// NewValidationResultFromValue returns a pointer to a valid ValidationResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValidationResultFromValue(v int32) (*ValidationResult, error) {
	ev := ValidationResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValidationResult: valid values are %v", v, AllowedValidationResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValidationResult) IsValid() bool {
	for _, existing := range AllowedValidationResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidationResult value
func (v ValidationResult) Ptr() *ValidationResult {
	return &v
}

type NullableValidationResult struct {
	value *ValidationResult
	isSet bool
}

func (v NullableValidationResult) Get() *ValidationResult {
	return v.value
}

func (v *NullableValidationResult) Set(val *ValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationResult(val *ValidationResult) *NullableValidationResult {
	return &NullableValidationResult{value: val, isSet: true}
}

func (v NullableValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

