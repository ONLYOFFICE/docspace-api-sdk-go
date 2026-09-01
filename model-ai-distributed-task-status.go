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

// AiDistributedTaskStatus [0 - Created, 1 - Running, 2 - Completed, 3 - Canceled, 4 - Failted]
type AiDistributedTaskStatus int32

// List of AiDistributedTaskStatus
const (
	AIDISTRIBUTEDTASKSTATUS_Created AiDistributedTaskStatus = 0
	AIDISTRIBUTEDTASKSTATUS_Running AiDistributedTaskStatus = 1
	AIDISTRIBUTEDTASKSTATUS_Completed AiDistributedTaskStatus = 2
	AIDISTRIBUTEDTASKSTATUS_Canceled AiDistributedTaskStatus = 3
	AIDISTRIBUTEDTASKSTATUS_Failted AiDistributedTaskStatus = 4
)

// All allowed values of AiDistributedTaskStatus enum
var AllowedAiDistributedTaskStatusEnumValues = []AiDistributedTaskStatus{
	0,
	1,
	2,
	3,
	4,
}

func (v *AiDistributedTaskStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiDistributedTaskStatus(value)
	for _, existing := range AllowedAiDistributedTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiDistributedTaskStatus", value)
}

// NewAiDistributedTaskStatusFromValue returns a pointer to a valid AiDistributedTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiDistributedTaskStatusFromValue(v int32) (*AiDistributedTaskStatus, error) {
	ev := AiDistributedTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiDistributedTaskStatus: valid values are %v", v, AllowedAiDistributedTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiDistributedTaskStatus) IsValid() bool {
	for _, existing := range AllowedAiDistributedTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiDistributedTaskStatus value
func (v AiDistributedTaskStatus) Ptr() *AiDistributedTaskStatus {
	return &v
}

type NullableAiDistributedTaskStatus struct {
	value *AiDistributedTaskStatus
	isSet bool
}

func (v NullableAiDistributedTaskStatus) Get() *AiDistributedTaskStatus {
	return v.value
}

func (v *NullableAiDistributedTaskStatus) Set(val *AiDistributedTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAiDistributedTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAiDistributedTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiDistributedTaskStatus(val *AiDistributedTaskStatus) *NullableAiDistributedTaskStatus {
	return &NullableAiDistributedTaskStatus{value: val, isSet: true}
}

func (v NullableAiDistributedTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiDistributedTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

