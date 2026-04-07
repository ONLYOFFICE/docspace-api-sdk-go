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

// DarkThemeSettingsType [Base - Base, Dark - Dark, System - System]
type DarkThemeSettingsType string

// List of DarkThemeSettingsType
const (
	DARKTHEMESETTINGSTYPE_BASE DarkThemeSettingsType = "Base"
	DARKTHEMESETTINGSTYPE_DARK DarkThemeSettingsType = "Dark"
	DARKTHEMESETTINGSTYPE_SYSTEM DarkThemeSettingsType = "System"
)

// All allowed values of DarkThemeSettingsType enum
var AllowedDarkThemeSettingsTypeEnumValues = []DarkThemeSettingsType{
	"Base",
	"Dark",
	"System",
}

func (v *DarkThemeSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DarkThemeSettingsType(value)
	for _, existing := range AllowedDarkThemeSettingsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DarkThemeSettingsType", value)
}

// NewDarkThemeSettingsTypeFromValue returns a pointer to a valid DarkThemeSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDarkThemeSettingsTypeFromValue(v string) (*DarkThemeSettingsType, error) {
	ev := DarkThemeSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DarkThemeSettingsType: valid values are %v", v, AllowedDarkThemeSettingsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DarkThemeSettingsType) IsValid() bool {
	for _, existing := range AllowedDarkThemeSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DarkThemeSettingsType value
func (v DarkThemeSettingsType) Ptr() *DarkThemeSettingsType {
	return &v
}

type NullableDarkThemeSettingsType struct {
	value *DarkThemeSettingsType
	isSet bool
}

func (v NullableDarkThemeSettingsType) Get() *DarkThemeSettingsType {
	return v.value
}

func (v *NullableDarkThemeSettingsType) Set(val *DarkThemeSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableDarkThemeSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableDarkThemeSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDarkThemeSettingsType(val *DarkThemeSettingsType) *NullableDarkThemeSettingsType {
	return &NullableDarkThemeSettingsType{value: val, isSet: true}
}

func (v NullableDarkThemeSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDarkThemeSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

