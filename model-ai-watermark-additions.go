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

// AiWatermarkAdditions [1 - User name, 2 - User email, 4 - User ip adress, 8 - Current date, 16 - Room name]
type AiWatermarkAdditions int32

// List of AiWatermarkAdditions
const (
	AIWATERMARKADDITIONS_UserName AiWatermarkAdditions = 1
	AIWATERMARKADDITIONS_UserEmail AiWatermarkAdditions = 2
	AIWATERMARKADDITIONS_UserIpAdress AiWatermarkAdditions = 4
	AIWATERMARKADDITIONS_CurrentDate AiWatermarkAdditions = 8
	AIWATERMARKADDITIONS_RoomName AiWatermarkAdditions = 16
)

// All allowed values of AiWatermarkAdditions enum
var AllowedAiWatermarkAdditionsEnumValues = []AiWatermarkAdditions{
	1,
	2,
	4,
	8,
	16,
}

func (v *AiWatermarkAdditions) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiWatermarkAdditions(value)
	for _, existing := range AllowedAiWatermarkAdditionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiWatermarkAdditions", value)
}

// NewAiWatermarkAdditionsFromValue returns a pointer to a valid AiWatermarkAdditions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiWatermarkAdditionsFromValue(v int32) (*AiWatermarkAdditions, error) {
	ev := AiWatermarkAdditions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiWatermarkAdditions: valid values are %v", v, AllowedAiWatermarkAdditionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiWatermarkAdditions) IsValid() bool {
	for _, existing := range AllowedAiWatermarkAdditionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiWatermarkAdditions value
func (v AiWatermarkAdditions) Ptr() *AiWatermarkAdditions {
	return &v
}

type NullableAiWatermarkAdditions struct {
	value *AiWatermarkAdditions
	isSet bool
}

func (v NullableAiWatermarkAdditions) Get() *AiWatermarkAdditions {
	return v.value
}

func (v *NullableAiWatermarkAdditions) Set(val *AiWatermarkAdditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWatermarkAdditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWatermarkAdditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWatermarkAdditions(val *AiWatermarkAdditions) *NullableAiWatermarkAdditions {
	return &NullableAiWatermarkAdditions{value: val, isSet: true}
}

func (v NullableAiWatermarkAdditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWatermarkAdditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

