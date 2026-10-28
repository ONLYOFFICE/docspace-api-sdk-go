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

// ExternalDatabaseType []
type ExternalDatabaseType int32

// List of ExternalDatabaseType
const (
	EXTERNALDATABASETYPE_MySql ExternalDatabaseType = 0
	EXTERNALDATABASETYPE_Sqlite ExternalDatabaseType = 1
)

// All allowed values of ExternalDatabaseType enum
var AllowedExternalDatabaseTypeEnumValues = []ExternalDatabaseType{
	0,
	1,
}

func (v *ExternalDatabaseType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalDatabaseType(value)
	for _, existing := range AllowedExternalDatabaseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalDatabaseType", value)
}

// NewExternalDatabaseTypeFromValue returns a pointer to a valid ExternalDatabaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalDatabaseTypeFromValue(v int32) (*ExternalDatabaseType, error) {
	ev := ExternalDatabaseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalDatabaseType: valid values are %v", v, AllowedExternalDatabaseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalDatabaseType) IsValid() bool {
	for _, existing := range AllowedExternalDatabaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalDatabaseType value
func (v ExternalDatabaseType) Ptr() *ExternalDatabaseType {
	return &v
}

type NullableExternalDatabaseType struct {
	value *ExternalDatabaseType
	isSet bool
}

func (v NullableExternalDatabaseType) Get() *ExternalDatabaseType {
	return v.value
}

func (v *NullableExternalDatabaseType) Set(val *ExternalDatabaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalDatabaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalDatabaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalDatabaseType(val *ExternalDatabaseType) *NullableExternalDatabaseType {
	return &NullableExternalDatabaseType{value: val, isSet: true}
}

func (v NullableExternalDatabaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalDatabaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

