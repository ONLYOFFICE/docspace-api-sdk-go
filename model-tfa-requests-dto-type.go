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

// TfaRequestsDtoType [0 - None, 1 - Sms, 2 - App]
type TfaRequestsDtoType int32

// List of TfaRequestsDtoType
const (
	TFAREQUESTSDTOTYPE_None TfaRequestsDtoType = 0
	TFAREQUESTSDTOTYPE_Sms TfaRequestsDtoType = 1
	TFAREQUESTSDTOTYPE_App TfaRequestsDtoType = 2
)

// All allowed values of TfaRequestsDtoType enum
var AllowedTfaRequestsDtoTypeEnumValues = []TfaRequestsDtoType{
	0,
	1,
	2,
}

func (v *TfaRequestsDtoType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TfaRequestsDtoType(value)
	for _, existing := range AllowedTfaRequestsDtoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TfaRequestsDtoType", value)
}

// NewTfaRequestsDtoTypeFromValue returns a pointer to a valid TfaRequestsDtoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTfaRequestsDtoTypeFromValue(v int32) (*TfaRequestsDtoType, error) {
	ev := TfaRequestsDtoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TfaRequestsDtoType: valid values are %v", v, AllowedTfaRequestsDtoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TfaRequestsDtoType) IsValid() bool {
	for _, existing := range AllowedTfaRequestsDtoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TfaRequestsDtoType value
func (v TfaRequestsDtoType) Ptr() *TfaRequestsDtoType {
	return &v
}

type NullableTfaRequestsDtoType struct {
	value *TfaRequestsDtoType
	isSet bool
}

func (v NullableTfaRequestsDtoType) Get() *TfaRequestsDtoType {
	return v.value
}

func (v *NullableTfaRequestsDtoType) Set(val *TfaRequestsDtoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaRequestsDtoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaRequestsDtoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaRequestsDtoType(val *TfaRequestsDtoType) *NullableTfaRequestsDtoType {
	return &NullableTfaRequestsDtoType{value: val, isSet: true}
}

func (v NullableTfaRequestsDtoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaRequestsDtoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

