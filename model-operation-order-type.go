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

// OperationOrderType [0 - Descending, 1 - Ascending]
type OperationOrderType int32

// List of OperationOrderType
const (
	OPERATIONORDERTYPE_Descending OperationOrderType = 0
	OPERATIONORDERTYPE_Ascending OperationOrderType = 1
)

// All allowed values of OperationOrderType enum
var AllowedOperationOrderTypeEnumValues = []OperationOrderType{
	0,
	1,
}

func (v *OperationOrderType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationOrderType(value)
	for _, existing := range AllowedOperationOrderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationOrderType", value)
}

// NewOperationOrderTypeFromValue returns a pointer to a valid OperationOrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationOrderTypeFromValue(v int32) (*OperationOrderType, error) {
	ev := OperationOrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationOrderType: valid values are %v", v, AllowedOperationOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationOrderType) IsValid() bool {
	for _, existing := range AllowedOperationOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationOrderType value
func (v OperationOrderType) Ptr() *OperationOrderType {
	return &v
}

type NullableOperationOrderType struct {
	value *OperationOrderType
	isSet bool
}

func (v NullableOperationOrderType) Get() *OperationOrderType {
	return v.value
}

func (v *NullableOperationOrderType) Set(val *OperationOrderType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationOrderType(val *OperationOrderType) *NullableOperationOrderType {
	return &NullableOperationOrderType{value: val, isSet: true}
}

func (v NullableOperationOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

