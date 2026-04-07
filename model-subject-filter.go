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

// SubjectFilter [0 - Owner, 1 - Member]
type SubjectFilter int32

// List of SubjectFilter
const (
	SUBJECTFILTER_Owner SubjectFilter = 0
	SUBJECTFILTER_Member SubjectFilter = 1
)

// All allowed values of SubjectFilter enum
var AllowedSubjectFilterEnumValues = []SubjectFilter{
	0,
	1,
}

func (v *SubjectFilter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubjectFilter(value)
	for _, existing := range AllowedSubjectFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubjectFilter", value)
}

// NewSubjectFilterFromValue returns a pointer to a valid SubjectFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubjectFilterFromValue(v int32) (*SubjectFilter, error) {
	ev := SubjectFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubjectFilter: valid values are %v", v, AllowedSubjectFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubjectFilter) IsValid() bool {
	for _, existing := range AllowedSubjectFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubjectFilter value
func (v SubjectFilter) Ptr() *SubjectFilter {
	return &v
}

type NullableSubjectFilter struct {
	value *SubjectFilter
	isSet bool
}

func (v NullableSubjectFilter) Get() *SubjectFilter {
	return v.value
}

func (v *NullableSubjectFilter) Set(val *SubjectFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectFilter(val *SubjectFilter) *NullableSubjectFilter {
	return &NullableSubjectFilter{value: val, isSet: true}
}

func (v NullableSubjectFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

