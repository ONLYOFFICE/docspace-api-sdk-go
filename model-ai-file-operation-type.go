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

// AiFileOperationType [0 - Move, 1 - Copy, 2 - Delete, 3 - Download, 4 - MarkAsRead, 5 - Import, 6 - Convert, 7 - Duplicate]
type AiFileOperationType int32

// List of AiFileOperationType
const (
	AIFILEOPERATIONTYPE_Move AiFileOperationType = 0
	AIFILEOPERATIONTYPE_Copy AiFileOperationType = 1
	AIFILEOPERATIONTYPE_Delete AiFileOperationType = 2
	AIFILEOPERATIONTYPE_Download AiFileOperationType = 3
	AIFILEOPERATIONTYPE_MarkAsRead AiFileOperationType = 4
	AIFILEOPERATIONTYPE_Import AiFileOperationType = 5
	AIFILEOPERATIONTYPE_Convert AiFileOperationType = 6
	AIFILEOPERATIONTYPE_Duplicate AiFileOperationType = 7
)

// All allowed values of AiFileOperationType enum
var AllowedAiFileOperationTypeEnumValues = []AiFileOperationType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

func (v *AiFileOperationType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiFileOperationType(value)
	for _, existing := range AllowedAiFileOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiFileOperationType", value)
}

// NewAiFileOperationTypeFromValue returns a pointer to a valid AiFileOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiFileOperationTypeFromValue(v int32) (*AiFileOperationType, error) {
	ev := AiFileOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiFileOperationType: valid values are %v", v, AllowedAiFileOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiFileOperationType) IsValid() bool {
	for _, existing := range AllowedAiFileOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiFileOperationType value
func (v AiFileOperationType) Ptr() *AiFileOperationType {
	return &v
}

type NullableAiFileOperationType struct {
	value *AiFileOperationType
	isSet bool
}

func (v NullableAiFileOperationType) Get() *AiFileOperationType {
	return v.value
}

func (v *NullableAiFileOperationType) Set(val *AiFileOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileOperationType(val *AiFileOperationType) *NullableAiFileOperationType {
	return &NullableAiFileOperationType{value: val, isSet: true}
}

func (v NullableAiFileOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

