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

// TfaType [0 - None, 1 - Sms, 2 - App]
type TfaType int32

// List of TfaType
const (
	TFATYPE_None TfaType = 0
	TFATYPE_Sms TfaType = 1
	TFATYPE_App TfaType = 2
)

// All allowed values of TfaType enum
var AllowedTfaTypeEnumValues = []TfaType{
	0,
	1,
	2,
}

func (v *TfaType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TfaType(value)
	for _, existing := range AllowedTfaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TfaType", value)
}

// NewTfaTypeFromValue returns a pointer to a valid TfaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTfaTypeFromValue(v int32) (*TfaType, error) {
	ev := TfaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TfaType: valid values are %v", v, AllowedTfaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TfaType) IsValid() bool {
	for _, existing := range AllowedTfaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TfaType value
func (v TfaType) Ptr() *TfaType {
	return &v
}

type NullableTfaType struct {
	value *TfaType
	isSet bool
}

func (v NullableTfaType) Get() *TfaType {
	return v.value
}

func (v *NullableTfaType) Set(val *TfaType) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaType) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaType(val *TfaType) *NullableTfaType {
	return &NullableTfaType{value: val, isSet: true}
}

func (v NullableTfaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

