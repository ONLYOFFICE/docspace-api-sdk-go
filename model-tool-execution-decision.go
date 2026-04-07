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

// ToolExecutionDecision [0 - Allow, 1 - AlwaysAllow, 2 - Deny]
type ToolExecutionDecision int32

// List of ToolExecutionDecision
const (
	TOOLEXECUTIONDECISION_Allow ToolExecutionDecision = 0
	TOOLEXECUTIONDECISION_AlwaysAllow ToolExecutionDecision = 1
	TOOLEXECUTIONDECISION_Deny ToolExecutionDecision = 2
)

// All allowed values of ToolExecutionDecision enum
var AllowedToolExecutionDecisionEnumValues = []ToolExecutionDecision{
	0,
	1,
	2,
}

func (v *ToolExecutionDecision) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ToolExecutionDecision(value)
	for _, existing := range AllowedToolExecutionDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ToolExecutionDecision", value)
}

// NewToolExecutionDecisionFromValue returns a pointer to a valid ToolExecutionDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolExecutionDecisionFromValue(v int32) (*ToolExecutionDecision, error) {
	ev := ToolExecutionDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolExecutionDecision: valid values are %v", v, AllowedToolExecutionDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolExecutionDecision) IsValid() bool {
	for _, existing := range AllowedToolExecutionDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToolExecutionDecision value
func (v ToolExecutionDecision) Ptr() *ToolExecutionDecision {
	return &v
}

type NullableToolExecutionDecision struct {
	value *ToolExecutionDecision
	isSet bool
}

func (v NullableToolExecutionDecision) Get() *ToolExecutionDecision {
	return v.value
}

func (v *NullableToolExecutionDecision) Set(val *ToolExecutionDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableToolExecutionDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableToolExecutionDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolExecutionDecision(val *ToolExecutionDecision) *NullableToolExecutionDecision {
	return &NullableToolExecutionDecision{value: val, isSet: true}
}

func (v NullableToolExecutionDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolExecutionDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

