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

// ShareFilterType [0 - User or group, 1 - Invitation link, 2 - External link, 4 - Additional external link, 8 - Primary external link, 16 - User, 32 - Group]
type ShareFilterType int32

// List of ShareFilterType
const (
	SHAREFILTERTYPE_UserOrGroup ShareFilterType = 0
	SHAREFILTERTYPE_InvitationLink ShareFilterType = 1
	SHAREFILTERTYPE_ExternalLink ShareFilterType = 2
	SHAREFILTERTYPE_AdditionalExternalLink ShareFilterType = 4
	SHAREFILTERTYPE_PrimaryExternalLink ShareFilterType = 8
	SHAREFILTERTYPE_Link ShareFilterType = 15
	SHAREFILTERTYPE_User ShareFilterType = 16
	SHAREFILTERTYPE_Group ShareFilterType = 32
)

// All allowed values of ShareFilterType enum
var AllowedShareFilterTypeEnumValues = []ShareFilterType{
	0,
	1,
	2,
	4,
	8,
	15,
	16,
	32,
}

func (v *ShareFilterType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShareFilterType(value)
	for _, existing := range AllowedShareFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShareFilterType", value)
}

// NewShareFilterTypeFromValue returns a pointer to a valid ShareFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShareFilterTypeFromValue(v int32) (*ShareFilterType, error) {
	ev := ShareFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShareFilterType: valid values are %v", v, AllowedShareFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShareFilterType) IsValid() bool {
	for _, existing := range AllowedShareFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShareFilterType value
func (v ShareFilterType) Ptr() *ShareFilterType {
	return &v
}

type NullableShareFilterType struct {
	value *ShareFilterType
	isSet bool
}

func (v NullableShareFilterType) Get() *ShareFilterType {
	return v.value
}

func (v *NullableShareFilterType) Set(val *ShareFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableShareFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableShareFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareFilterType(val *ShareFilterType) *NullableShareFilterType {
	return &NullableShareFilterType{value: val, isSet: true}
}

func (v NullableShareFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

