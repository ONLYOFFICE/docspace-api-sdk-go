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

// RoomPrivacyFilter [0 - None, 1 - Private, 2 - NotPrivate]
type RoomPrivacyFilter int32

// List of RoomPrivacyFilter
const (
	ROOMPRIVACYFILTER_None RoomPrivacyFilter = 0
	ROOMPRIVACYFILTER_Private RoomPrivacyFilter = 1
	ROOMPRIVACYFILTER_NotPrivate RoomPrivacyFilter = 2
)

// All allowed values of RoomPrivacyFilter enum
var AllowedRoomPrivacyFilterEnumValues = []RoomPrivacyFilter{
	0,
	1,
	2,
}

func (v *RoomPrivacyFilter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomPrivacyFilter(value)
	for _, existing := range AllowedRoomPrivacyFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomPrivacyFilter", value)
}

// NewRoomPrivacyFilterFromValue returns a pointer to a valid RoomPrivacyFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomPrivacyFilterFromValue(v int32) (*RoomPrivacyFilter, error) {
	ev := RoomPrivacyFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomPrivacyFilter: valid values are %v", v, AllowedRoomPrivacyFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomPrivacyFilter) IsValid() bool {
	for _, existing := range AllowedRoomPrivacyFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoomPrivacyFilter value
func (v RoomPrivacyFilter) Ptr() *RoomPrivacyFilter {
	return &v
}

type NullableRoomPrivacyFilter struct {
	value *RoomPrivacyFilter
	isSet bool
}

func (v NullableRoomPrivacyFilter) Get() *RoomPrivacyFilter {
	return v.value
}

func (v *NullableRoomPrivacyFilter) Set(val *RoomPrivacyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomPrivacyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomPrivacyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomPrivacyFilter(val *RoomPrivacyFilter) *NullableRoomPrivacyFilter {
	return &NullableRoomPrivacyFilter{value: val, isSet: true}
}

func (v NullableRoomPrivacyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomPrivacyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

