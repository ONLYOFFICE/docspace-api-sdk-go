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

// RoomDataLifetimePeriod [0 - Day, 1 - Month, 2 - Year]
type RoomDataLifetimePeriod int32

// List of RoomDataLifetimePeriod
const (
	ROOMDATALIFETIMEPERIOD_Day RoomDataLifetimePeriod = 0
	ROOMDATALIFETIMEPERIOD_Month RoomDataLifetimePeriod = 1
	ROOMDATALIFETIMEPERIOD_Year RoomDataLifetimePeriod = 2
)

// All allowed values of RoomDataLifetimePeriod enum
var AllowedRoomDataLifetimePeriodEnumValues = []RoomDataLifetimePeriod{
	0,
	1,
	2,
}

func (v *RoomDataLifetimePeriod) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoomDataLifetimePeriod(value)
	for _, existing := range AllowedRoomDataLifetimePeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoomDataLifetimePeriod", value)
}

// NewRoomDataLifetimePeriodFromValue returns a pointer to a valid RoomDataLifetimePeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoomDataLifetimePeriodFromValue(v int32) (*RoomDataLifetimePeriod, error) {
	ev := RoomDataLifetimePeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoomDataLifetimePeriod: valid values are %v", v, AllowedRoomDataLifetimePeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoomDataLifetimePeriod) IsValid() bool {
	for _, existing := range AllowedRoomDataLifetimePeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoomDataLifetimePeriod value
func (v RoomDataLifetimePeriod) Ptr() *RoomDataLifetimePeriod {
	return &v
}

type NullableRoomDataLifetimePeriod struct {
	value *RoomDataLifetimePeriod
	isSet bool
}

func (v NullableRoomDataLifetimePeriod) Get() *RoomDataLifetimePeriod {
	return v.value
}

func (v *NullableRoomDataLifetimePeriod) Set(val *RoomDataLifetimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomDataLifetimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomDataLifetimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomDataLifetimePeriod(val *RoomDataLifetimePeriod) *NullableRoomDataLifetimePeriod {
	return &NullableRoomDataLifetimePeriod{value: val, isSet: true}
}

func (v NullableRoomDataLifetimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomDataLifetimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

