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

// ActionType [0 - None, 1 - Create, 2 - Update, 3 - Delete, 4 - Link, 5 - Unlink, 6 - Attach, 7 - Detach, 8 - Send, 9 - Import, 10 - Export, 11 - Update access, 12 - Download, 13 - Upload, 14 - Copy, 15 - Move, 16 - Reassigns, 17 - Follow, 18 - Unfollow, 19 - Logout]
type ActionType int32

// List of ActionType
const (
	ACTIONTYPE_None ActionType = 0
	ACTIONTYPE_Create ActionType = 1
	ACTIONTYPE_Update ActionType = 2
	ACTIONTYPE_Delete ActionType = 3
	ACTIONTYPE_Link ActionType = 4
	ACTIONTYPE_Unlink ActionType = 5
	ACTIONTYPE_Attach ActionType = 6
	ACTIONTYPE_Detach ActionType = 7
	ACTIONTYPE_Send ActionType = 8
	ACTIONTYPE_Import ActionType = 9
	ACTIONTYPE_Export ActionType = 10
	ACTIONTYPE_UpdateAccess ActionType = 11
	ACTIONTYPE_Download ActionType = 12
	ACTIONTYPE_Upload ActionType = 13
	ACTIONTYPE_Copy ActionType = 14
	ACTIONTYPE_Move ActionType = 15
	ACTIONTYPE_Reassigns ActionType = 16
	ACTIONTYPE_Follow ActionType = 17
	ACTIONTYPE_Unfollow ActionType = 18
	ACTIONTYPE_Logout ActionType = 19
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
}

func (v *ActionType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionType(value)
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionType", value)
}

// NewActionTypeFromValue returns a pointer to a valid ActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionTypeFromValue(v int32) (*ActionType, error) {
	ev := ActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionType: valid values are %v", v, AllowedActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionType value
func (v ActionType) Ptr() *ActionType {
	return &v
}

type NullableActionType struct {
	value *ActionType
	isSet bool
}

func (v NullableActionType) Get() *ActionType {
	return v.value
}

func (v *NullableActionType) Set(val *ActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionType(val *ActionType) *NullableActionType {
	return &NullableActionType{value: val, isSet: true}
}

func (v NullableActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

