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

// Location [1 - Room, 2 - Documents, 3 - Link]
type Location int32

// List of Location
const (
	LOCATION_Room Location = 1
	LOCATION_Documents Location = 2
	LOCATION_Link Location = 3
)

// All allowed values of Location enum
var AllowedLocationEnumValues = []Location{
	1,
	2,
	3,
}

func (v *Location) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Location(value)
	for _, existing := range AllowedLocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Location", value)
}

// NewLocationFromValue returns a pointer to a valid Location
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationFromValue(v int32) (*Location, error) {
	ev := Location(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Location: valid values are %v", v, AllowedLocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Location) IsValid() bool {
	for _, existing := range AllowedLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Location value
func (v Location) Ptr() *Location {
	return &v
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

