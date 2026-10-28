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

// DateToAutoCleanUp [1 - One week, 2 - Two weeks, 3 - One month, 4 - Thirty days, 5 - Two months, 6 - Three months]
type DateToAutoCleanUp int32

// List of DateToAutoCleanUp
const (
	DATETOAUTOCLEANUP_OneWeek DateToAutoCleanUp = 1
	DATETOAUTOCLEANUP_TwoWeeks DateToAutoCleanUp = 2
	DATETOAUTOCLEANUP_OneMonth DateToAutoCleanUp = 3
	DATETOAUTOCLEANUP_ThirtyDays DateToAutoCleanUp = 4
	DATETOAUTOCLEANUP_TwoMonths DateToAutoCleanUp = 5
	DATETOAUTOCLEANUP_ThreeMonths DateToAutoCleanUp = 6
)

// All allowed values of DateToAutoCleanUp enum
var AllowedDateToAutoCleanUpEnumValues = []DateToAutoCleanUp{
	1,
	2,
	3,
	4,
	5,
	6,
}

func (v *DateToAutoCleanUp) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DateToAutoCleanUp(value)
	for _, existing := range AllowedDateToAutoCleanUpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DateToAutoCleanUp", value)
}

// NewDateToAutoCleanUpFromValue returns a pointer to a valid DateToAutoCleanUp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDateToAutoCleanUpFromValue(v int32) (*DateToAutoCleanUp, error) {
	ev := DateToAutoCleanUp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DateToAutoCleanUp: valid values are %v", v, AllowedDateToAutoCleanUpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DateToAutoCleanUp) IsValid() bool {
	for _, existing := range AllowedDateToAutoCleanUpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DateToAutoCleanUp value
func (v DateToAutoCleanUp) Ptr() *DateToAutoCleanUp {
	return &v
}

type NullableDateToAutoCleanUp struct {
	value *DateToAutoCleanUp
	isSet bool
}

func (v NullableDateToAutoCleanUp) Get() *DateToAutoCleanUp {
	return v.value
}

func (v *NullableDateToAutoCleanUp) Set(val *DateToAutoCleanUp) {
	v.value = val
	v.isSet = true
}

func (v NullableDateToAutoCleanUp) IsSet() bool {
	return v.isSet
}

func (v *NullableDateToAutoCleanUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateToAutoCleanUp(val *DateToAutoCleanUp) *NullableDateToAutoCleanUp {
	return &NullableDateToAutoCleanUp{value: val, isSet: true}
}

func (v NullableDateToAutoCleanUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateToAutoCleanUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

