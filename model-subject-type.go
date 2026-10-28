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

// SubjectType [0 - User, 1 - External link, 2 - Group, 3 - Invitation link, 4 - Primary external link]
type SubjectType int32

// List of SubjectType
const (
	SUBJECTTYPE_User SubjectType = 0
	SUBJECTTYPE_ExternalLink SubjectType = 1
	SUBJECTTYPE_Group SubjectType = 2
	SUBJECTTYPE_InvitationLink SubjectType = 3
	SUBJECTTYPE_PrimaryExternalLink SubjectType = 4
)

// All allowed values of SubjectType enum
var AllowedSubjectTypeEnumValues = []SubjectType{
	0,
	1,
	2,
	3,
	4,
}

func (v *SubjectType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubjectType(value)
	for _, existing := range AllowedSubjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubjectType", value)
}

// NewSubjectTypeFromValue returns a pointer to a valid SubjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubjectTypeFromValue(v int32) (*SubjectType, error) {
	ev := SubjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubjectType: valid values are %v", v, AllowedSubjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubjectType) IsValid() bool {
	for _, existing := range AllowedSubjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubjectType value
func (v SubjectType) Ptr() *SubjectType {
	return &v
}

type NullableSubjectType struct {
	value *SubjectType
	isSet bool
}

func (v NullableSubjectType) Get() *SubjectType {
	return v.value
}

func (v *NullableSubjectType) Set(val *SubjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectType(val *SubjectType) *NullableSubjectType {
	return &NullableSubjectType{value: val, isSet: true}
}

func (v NullableSubjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

