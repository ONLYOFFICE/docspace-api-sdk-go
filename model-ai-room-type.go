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

// AiRoomType [1 - Form filling room, 2 - Collaboration room, 5 - Custom room, 6 - Public room, 8 - Virtual data room, 9 - AI Room]
type AiRoomType int32

// List of AiRoomType
const (
	AIROOMTYPE_FillingFormsRoom AiRoomType = 1
	AIROOMTYPE_EditingRoom AiRoomType = 2
	AIROOMTYPE_CustomRoom AiRoomType = 5
	AIROOMTYPE_PublicRoom AiRoomType = 6
	AIROOMTYPE_VirtualDataRoom AiRoomType = 8
	AIROOMTYPE_AiRoom AiRoomType = 9
)

// All allowed values of AiRoomType enum
var AllowedAiRoomTypeEnumValues = []AiRoomType{
	1,
	2,
	5,
	6,
	8,
	9,
}

func (v *AiRoomType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiRoomType(value)
	for _, existing := range AllowedAiRoomTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiRoomType", value)
}

// NewAiRoomTypeFromValue returns a pointer to a valid AiRoomType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiRoomTypeFromValue(v int32) (*AiRoomType, error) {
	ev := AiRoomType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiRoomType: valid values are %v", v, AllowedAiRoomTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiRoomType) IsValid() bool {
	for _, existing := range AllowedAiRoomTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiRoomType value
func (v AiRoomType) Ptr() *AiRoomType {
	return &v
}

type NullableAiRoomType struct {
	value *AiRoomType
	isSet bool
}

func (v NullableAiRoomType) Get() *AiRoomType {
	return v.value
}

func (v *NullableAiRoomType) Set(val *AiRoomType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiRoomType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiRoomType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiRoomType(val *AiRoomType) *NullableAiRoomType {
	return &NullableAiRoomType{value: val, isSet: true}
}

func (v NullableAiRoomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiRoomType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

