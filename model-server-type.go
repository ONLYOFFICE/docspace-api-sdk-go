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

// ServerType [0 - Custom, 1 - DocSpace, 2 - Github, 3 - Box]
type ServerType int32

// List of ServerType
const (
	SERVERTYPE_Custom ServerType = 0
	SERVERTYPE_DocSpace ServerType = 1
	SERVERTYPE_Github ServerType = 2
	SERVERTYPE_Box ServerType = 3
)

// All allowed values of ServerType enum
var AllowedServerTypeEnumValues = []ServerType{
	0,
	1,
	2,
	3,
}

func (v *ServerType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServerType(value)
	for _, existing := range AllowedServerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServerType", value)
}

// NewServerTypeFromValue returns a pointer to a valid ServerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServerTypeFromValue(v int32) (*ServerType, error) {
	ev := ServerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServerType: valid values are %v", v, AllowedServerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServerType) IsValid() bool {
	for _, existing := range AllowedServerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServerType value
func (v ServerType) Ptr() *ServerType {
	return &v
}

type NullableServerType struct {
	value *ServerType
	isSet bool
}

func (v NullableServerType) Get() *ServerType {
	return v.value
}

func (v *NullableServerType) Set(val *ServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerType(val *ServerType) *NullableServerType {
	return &NullableServerType{value: val, isSet: true}
}

func (v NullableServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

