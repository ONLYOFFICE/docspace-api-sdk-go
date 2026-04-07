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

// LocationType [0 - None, 1 - Files, 2 - Folders, 3 - Documents settings, 27 - Rooms, 29 - Settings, 30 - Contacts, 31 - Agents]
type LocationType int32

// List of LocationType
const (
	LOCATIONTYPE_None LocationType = 0
	LOCATIONTYPE_Files LocationType = 1
	LOCATIONTYPE_Folders LocationType = 2
	LOCATIONTYPE_DocumentsSettings LocationType = 3
	LOCATIONTYPE_Rooms LocationType = 27
	LOCATIONTYPE_Settings LocationType = 29
	LOCATIONTYPE_Contacts LocationType = 30
	LOCATIONTYPE_Agents LocationType = 31
)

// All allowed values of LocationType enum
var AllowedLocationTypeEnumValues = []LocationType{
	0,
	1,
	2,
	3,
	27,
	29,
	30,
	31,
}

func (v *LocationType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationType(value)
	for _, existing := range AllowedLocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationType", value)
}

// NewLocationTypeFromValue returns a pointer to a valid LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationTypeFromValue(v int32) (*LocationType, error) {
	ev := LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationType: valid values are %v", v, AllowedLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationType) IsValid() bool {
	for _, existing := range AllowedLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationType value
func (v LocationType) Ptr() *LocationType {
	return &v
}

type NullableLocationType struct {
	value *LocationType
	isSet bool
}

func (v NullableLocationType) Get() *LocationType {
	return v.value
}

func (v *NullableLocationType) Set(val *LocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationType(val *LocationType) *NullableLocationType {
	return &NullableLocationType{value: val, isSet: true}
}

func (v NullableLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

