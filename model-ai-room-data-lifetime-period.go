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

// AiRoomDataLifetimePeriod [0 - Day, 1 - Month, 2 - Year]
type AiRoomDataLifetimePeriod int32

// List of AiRoomDataLifetimePeriod
const (
	AIROOMDATALIFETIMEPERIOD_Day AiRoomDataLifetimePeriod = 0
	AIROOMDATALIFETIMEPERIOD_Month AiRoomDataLifetimePeriod = 1
	AIROOMDATALIFETIMEPERIOD_Year AiRoomDataLifetimePeriod = 2
)

// All allowed values of AiRoomDataLifetimePeriod enum
var AllowedAiRoomDataLifetimePeriodEnumValues = []AiRoomDataLifetimePeriod{
	0,
	1,
	2,
}

func (v *AiRoomDataLifetimePeriod) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiRoomDataLifetimePeriod(value)
	for _, existing := range AllowedAiRoomDataLifetimePeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiRoomDataLifetimePeriod", value)
}

// NewAiRoomDataLifetimePeriodFromValue returns a pointer to a valid AiRoomDataLifetimePeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiRoomDataLifetimePeriodFromValue(v int32) (*AiRoomDataLifetimePeriod, error) {
	ev := AiRoomDataLifetimePeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiRoomDataLifetimePeriod: valid values are %v", v, AllowedAiRoomDataLifetimePeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiRoomDataLifetimePeriod) IsValid() bool {
	for _, existing := range AllowedAiRoomDataLifetimePeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiRoomDataLifetimePeriod value
func (v AiRoomDataLifetimePeriod) Ptr() *AiRoomDataLifetimePeriod {
	return &v
}

type NullableAiRoomDataLifetimePeriod struct {
	value *AiRoomDataLifetimePeriod
	isSet bool
}

func (v NullableAiRoomDataLifetimePeriod) Get() *AiRoomDataLifetimePeriod {
	return v.value
}

func (v *NullableAiRoomDataLifetimePeriod) Set(val *AiRoomDataLifetimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableAiRoomDataLifetimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableAiRoomDataLifetimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiRoomDataLifetimePeriod(val *AiRoomDataLifetimePeriod) *NullableAiRoomDataLifetimePeriod {
	return &NullableAiRoomDataLifetimePeriod{value: val, isSet: true}
}

func (v NullableAiRoomDataLifetimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiRoomDataLifetimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

