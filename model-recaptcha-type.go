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

// RecaptchaType [0 - Default, 1 - AndroidV2, 2 - iOSV2, 3 - hCaptcha]
type RecaptchaType int32

// List of RecaptchaType
const (
	RECAPTCHATYPE_Default RecaptchaType = 0
	RECAPTCHATYPE_AndroidV2 RecaptchaType = 1
	RECAPTCHATYPE_iOSV2 RecaptchaType = 2
	RECAPTCHATYPE_hCaptcha RecaptchaType = 3
)

// All allowed values of RecaptchaType enum
var AllowedRecaptchaTypeEnumValues = []RecaptchaType{
	0,
	1,
	2,
	3,
}

func (v *RecaptchaType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecaptchaType(value)
	for _, existing := range AllowedRecaptchaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecaptchaType", value)
}

// NewRecaptchaTypeFromValue returns a pointer to a valid RecaptchaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecaptchaTypeFromValue(v int32) (*RecaptchaType, error) {
	ev := RecaptchaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecaptchaType: valid values are %v", v, AllowedRecaptchaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecaptchaType) IsValid() bool {
	for _, existing := range AllowedRecaptchaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecaptchaType value
func (v RecaptchaType) Ptr() *RecaptchaType {
	return &v
}

type NullableRecaptchaType struct {
	value *RecaptchaType
	isSet bool
}

func (v NullableRecaptchaType) Get() *RecaptchaType {
	return v.value
}

func (v *NullableRecaptchaType) Set(val *RecaptchaType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecaptchaType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecaptchaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecaptchaType(val *RecaptchaType) *NullableRecaptchaType {
	return &NullableRecaptchaType{value: val, isSet: true}
}

func (v NullableRecaptchaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecaptchaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

