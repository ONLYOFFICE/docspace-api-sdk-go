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

// SearchArea [0 - Active, 1 - Archive, 2 - Any, 3 - Recent by links, 4 - Template, 5 - Knowledge, 6 - Result storage, 7 - AiAgents]
type SearchArea int32

// List of SearchArea
const (
	SEARCHAREA_Active SearchArea = 0
	SEARCHAREA_Archive SearchArea = 1
	SEARCHAREA_Any SearchArea = 2
	SEARCHAREA_RecentByLinks SearchArea = 3
	SEARCHAREA_Templates SearchArea = 4
	SEARCHAREA_Knowledge SearchArea = 5
	SEARCHAREA_ResultStorage SearchArea = 6
	SEARCHAREA_AiAgents SearchArea = 7
)

// All allowed values of SearchArea enum
var AllowedSearchAreaEnumValues = []SearchArea{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

func (v *SearchArea) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchArea(value)
	for _, existing := range AllowedSearchAreaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchArea", value)
}

// NewSearchAreaFromValue returns a pointer to a valid SearchArea
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchAreaFromValue(v int32) (*SearchArea, error) {
	ev := SearchArea(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchArea: valid values are %v", v, AllowedSearchAreaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchArea) IsValid() bool {
	for _, existing := range AllowedSearchAreaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchArea value
func (v SearchArea) Ptr() *SearchArea {
	return &v
}

type NullableSearchArea struct {
	value *SearchArea
	isSet bool
}

func (v NullableSearchArea) Get() *SearchArea {
	return v.value
}

func (v *NullableSearchArea) Set(val *SearchArea) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchArea) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchArea(val *SearchArea) *NullableSearchArea {
	return &NullableSearchArea{value: val, isSet: true}
}

func (v NullableSearchArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

