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

// Area [0 - All, 1 - People, 2 - Guests]
type Area int32

// List of Area
const (
	AREA_All Area = 0
	AREA_People Area = 1
	AREA_Guests Area = 2
)

// All allowed values of Area enum
var AllowedAreaEnumValues = []Area{
	0,
	1,
	2,
}

func (v *Area) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Area(value)
	for _, existing := range AllowedAreaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Area", value)
}

// NewAreaFromValue returns a pointer to a valid Area
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAreaFromValue(v int32) (*Area, error) {
	ev := Area(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Area: valid values are %v", v, AllowedAreaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Area) IsValid() bool {
	for _, existing := range AllowedAreaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Area value
func (v Area) Ptr() *Area {
	return &v
}

type NullableArea struct {
	value *Area
	isSet bool
}

func (v NullableArea) Get() *Area {
	return v.value
}

func (v *NullableArea) Set(val *Area) {
	v.value = val
	v.isSet = true
}

func (v NullableArea) IsSet() bool {
	return v.isSet
}

func (v *NullableArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArea(val *Area) *NullableArea {
	return &NullableArea{value: val, isSet: true}
}

func (v NullableArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

