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

// QuotaState [0 - Active, 1 - Overdue]
type QuotaState int32

// List of QuotaState
const (
	QUOTASTATE_Active QuotaState = 0
	QUOTASTATE_Overdue QuotaState = 1
)

// All allowed values of QuotaState enum
var AllowedQuotaStateEnumValues = []QuotaState{
	0,
	1,
}

func (v *QuotaState) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuotaState(value)
	for _, existing := range AllowedQuotaStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuotaState", value)
}

// NewQuotaStateFromValue returns a pointer to a valid QuotaState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuotaStateFromValue(v int32) (*QuotaState, error) {
	ev := QuotaState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuotaState: valid values are %v", v, AllowedQuotaStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuotaState) IsValid() bool {
	for _, existing := range AllowedQuotaStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuotaState value
func (v QuotaState) Ptr() *QuotaState {
	return &v
}

type NullableQuotaState struct {
	value *QuotaState
	isSet bool
}

func (v NullableQuotaState) Get() *QuotaState {
	return v.value
}

func (v *NullableQuotaState) Set(val *QuotaState) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaState) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaState(val *QuotaState) *NullableQuotaState {
	return &NullableQuotaState{value: val, isSet: true}
}

func (v NullableQuotaState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

