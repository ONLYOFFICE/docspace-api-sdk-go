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

// AccountSearchArea [0 - All, 1 - People, 2 - Guests]
type AccountSearchArea int32

// List of AccountSearchArea
const (
	ACCOUNTSEARCHAREA_All AccountSearchArea = 0
	ACCOUNTSEARCHAREA_People AccountSearchArea = 1
	ACCOUNTSEARCHAREA_Guests AccountSearchArea = 2
)

// All allowed values of AccountSearchArea enum
var AllowedAccountSearchAreaEnumValues = []AccountSearchArea{
	0,
	1,
	2,
}

func (v *AccountSearchArea) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountSearchArea(value)
	for _, existing := range AllowedAccountSearchAreaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountSearchArea", value)
}

// NewAccountSearchAreaFromValue returns a pointer to a valid AccountSearchArea
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountSearchAreaFromValue(v int32) (*AccountSearchArea, error) {
	ev := AccountSearchArea(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountSearchArea: valid values are %v", v, AllowedAccountSearchAreaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountSearchArea) IsValid() bool {
	for _, existing := range AllowedAccountSearchAreaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountSearchArea value
func (v AccountSearchArea) Ptr() *AccountSearchArea {
	return &v
}

type NullableAccountSearchArea struct {
	value *AccountSearchArea
	isSet bool
}

func (v NullableAccountSearchArea) Get() *AccountSearchArea {
	return v.value
}

func (v *NullableAccountSearchArea) Set(val *AccountSearchArea) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSearchArea) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSearchArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSearchArea(val *AccountSearchArea) *NullableAccountSearchArea {
	return &NullableAccountSearchArea{value: val, isSet: true}
}

func (v NullableAccountSearchArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSearchArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

