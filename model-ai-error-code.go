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

// AiErrorCode Why a field-scoped operation was refused, independent of the wording of `message`, so a host can pick an HTTP status without parsing text:  - `required` — the field is empty. - `conflict` — the value collides with an existing entry. - `not_found` — the referenced entry does not exist. - `invalid` — the value has the wrong shape or an unsupported value.
type AiErrorCode string

// List of AiErrorCode
const (
	AIERRORCODE_REQUIRED AiErrorCode = "required"
	AIERRORCODE_CONFLICT AiErrorCode = "conflict"
	AIERRORCODE_NOT_FOUND AiErrorCode = "not_found"
	AIERRORCODE_INVALID AiErrorCode = "invalid"
)

// All allowed values of AiErrorCode enum
var AllowedAiErrorCodeEnumValues = []AiErrorCode{
	"required",
	"conflict",
	"not_found",
	"invalid",
}

func (v *AiErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiErrorCode(value)
	for _, existing := range AllowedAiErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiErrorCode", value)
}

// NewAiErrorCodeFromValue returns a pointer to a valid AiErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiErrorCodeFromValue(v string) (*AiErrorCode, error) {
	ev := AiErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiErrorCode: valid values are %v", v, AllowedAiErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiErrorCode) IsValid() bool {
	for _, existing := range AllowedAiErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiErrorCode value
func (v AiErrorCode) Ptr() *AiErrorCode {
	return &v
}

type NullableAiErrorCode struct {
	value *AiErrorCode
	isSet bool
}

func (v NullableAiErrorCode) Get() *AiErrorCode {
	return v.value
}

func (v *NullableAiErrorCode) Set(val *AiErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAiErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAiErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiErrorCode(val *AiErrorCode) *NullableAiErrorCode {
	return &NullableAiErrorCode{value: val, isSet: true}
}

func (v NullableAiErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

