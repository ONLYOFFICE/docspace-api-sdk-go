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

// DeepLinkHandlingMode [0 - Provide choice, 1 - Web, 2 - App]
type DeepLinkHandlingMode int32

// List of DeepLinkHandlingMode
const (
	DEEPLINKHANDLINGMODE_ProvideChoice DeepLinkHandlingMode = 0
	DEEPLINKHANDLINGMODE_Web DeepLinkHandlingMode = 1
	DEEPLINKHANDLINGMODE_App DeepLinkHandlingMode = 2
)

// All allowed values of DeepLinkHandlingMode enum
var AllowedDeepLinkHandlingModeEnumValues = []DeepLinkHandlingMode{
	0,
	1,
	2,
}

func (v *DeepLinkHandlingMode) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeepLinkHandlingMode(value)
	for _, existing := range AllowedDeepLinkHandlingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeepLinkHandlingMode", value)
}

// NewDeepLinkHandlingModeFromValue returns a pointer to a valid DeepLinkHandlingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeepLinkHandlingModeFromValue(v int32) (*DeepLinkHandlingMode, error) {
	ev := DeepLinkHandlingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeepLinkHandlingMode: valid values are %v", v, AllowedDeepLinkHandlingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeepLinkHandlingMode) IsValid() bool {
	for _, existing := range AllowedDeepLinkHandlingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeepLinkHandlingMode value
func (v DeepLinkHandlingMode) Ptr() *DeepLinkHandlingMode {
	return &v
}

type NullableDeepLinkHandlingMode struct {
	value *DeepLinkHandlingMode
	isSet bool
}

func (v NullableDeepLinkHandlingMode) Get() *DeepLinkHandlingMode {
	return v.value
}

func (v *NullableDeepLinkHandlingMode) Set(val *DeepLinkHandlingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDeepLinkHandlingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDeepLinkHandlingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeepLinkHandlingMode(val *DeepLinkHandlingMode) *NullableDeepLinkHandlingMode {
	return &NullableDeepLinkHandlingMode{value: val, isSet: true}
}

func (v NullableDeepLinkHandlingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeepLinkHandlingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

