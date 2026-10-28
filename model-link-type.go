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

// LinkType [0 - Invitation, 1 - External]
type LinkType int32

// List of LinkType
const (
	LINKTYPE_Invitation LinkType = 0
	LINKTYPE_External LinkType = 1
)

// All allowed values of LinkType enum
var AllowedLinkTypeEnumValues = []LinkType{
	0,
	1,
}

func (v *LinkType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkType(value)
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkType", value)
}

// NewLinkTypeFromValue returns a pointer to a valid LinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkTypeFromValue(v int32) (*LinkType, error) {
	ev := LinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkType: valid values are %v", v, AllowedLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkType) IsValid() bool {
	for _, existing := range AllowedLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkType value
func (v LinkType) Ptr() *LinkType {
	return &v
}

type NullableLinkType struct {
	value *LinkType
	isSet bool
}

func (v NullableLinkType) Get() *LinkType {
	return v.value
}

func (v *NullableLinkType) Set(val *LinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkType(val *LinkType) *NullableLinkType {
	return &NullableLinkType{value: val, isSet: true}
}

func (v NullableLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

