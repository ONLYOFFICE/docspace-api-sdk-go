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

// WatermarkAdditions [1 - User name, 2 - User email, 4 - User ip adress, 8 - Current date, 16 - Room name]
type WatermarkAdditions int32

// List of WatermarkAdditions
const (
	WATERMARKADDITIONS_UserName WatermarkAdditions = 1
	WATERMARKADDITIONS_UserEmail WatermarkAdditions = 2
	WATERMARKADDITIONS_UserIpAdress WatermarkAdditions = 4
	WATERMARKADDITIONS_CurrentDate WatermarkAdditions = 8
	WATERMARKADDITIONS_RoomName WatermarkAdditions = 16
)

// All allowed values of WatermarkAdditions enum
var AllowedWatermarkAdditionsEnumValues = []WatermarkAdditions{
	1,
	2,
	4,
	8,
	16,
}

func (v *WatermarkAdditions) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WatermarkAdditions(value)
	for _, existing := range AllowedWatermarkAdditionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WatermarkAdditions", value)
}

// NewWatermarkAdditionsFromValue returns a pointer to a valid WatermarkAdditions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWatermarkAdditionsFromValue(v int32) (*WatermarkAdditions, error) {
	ev := WatermarkAdditions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WatermarkAdditions: valid values are %v", v, AllowedWatermarkAdditionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WatermarkAdditions) IsValid() bool {
	for _, existing := range AllowedWatermarkAdditionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WatermarkAdditions value
func (v WatermarkAdditions) Ptr() *WatermarkAdditions {
	return &v
}

type NullableWatermarkAdditions struct {
	value *WatermarkAdditions
	isSet bool
}

func (v NullableWatermarkAdditions) Get() *WatermarkAdditions {
	return v.value
}

func (v *NullableWatermarkAdditions) Set(val *WatermarkAdditions) {
	v.value = val
	v.isSet = true
}

func (v NullableWatermarkAdditions) IsSet() bool {
	return v.isSet
}

func (v *NullableWatermarkAdditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatermarkAdditions(val *WatermarkAdditions) *NullableWatermarkAdditions {
	return &NullableWatermarkAdditions{value: val, isSet: true}
}

func (v NullableWatermarkAdditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatermarkAdditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

