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

// DistributedTaskStatus [0 - Created, 1 - Running, 2 - Completed, 3 - Canceled, 4 - Failted]
type DistributedTaskStatus int32

// List of DistributedTaskStatus
const (
	DISTRIBUTEDTASKSTATUS_Created DistributedTaskStatus = 0
	DISTRIBUTEDTASKSTATUS_Running DistributedTaskStatus = 1
	DISTRIBUTEDTASKSTATUS_Completed DistributedTaskStatus = 2
	DISTRIBUTEDTASKSTATUS_Canceled DistributedTaskStatus = 3
	DISTRIBUTEDTASKSTATUS_Failted DistributedTaskStatus = 4
)

// All allowed values of DistributedTaskStatus enum
var AllowedDistributedTaskStatusEnumValues = []DistributedTaskStatus{
	0,
	1,
	2,
	3,
	4,
}

func (v *DistributedTaskStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DistributedTaskStatus(value)
	for _, existing := range AllowedDistributedTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DistributedTaskStatus", value)
}

// NewDistributedTaskStatusFromValue returns a pointer to a valid DistributedTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributedTaskStatusFromValue(v int32) (*DistributedTaskStatus, error) {
	ev := DistributedTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributedTaskStatus: valid values are %v", v, AllowedDistributedTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributedTaskStatus) IsValid() bool {
	for _, existing := range AllowedDistributedTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributedTaskStatus value
func (v DistributedTaskStatus) Ptr() *DistributedTaskStatus {
	return &v
}

type NullableDistributedTaskStatus struct {
	value *DistributedTaskStatus
	isSet bool
}

func (v NullableDistributedTaskStatus) Get() *DistributedTaskStatus {
	return v.value
}

func (v *NullableDistributedTaskStatus) Set(val *DistributedTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributedTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributedTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributedTaskStatus(val *DistributedTaskStatus) *NullableDistributedTaskStatus {
	return &NullableDistributedTaskStatus{value: val, isSet: true}
}

func (v NullableDistributedTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributedTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

