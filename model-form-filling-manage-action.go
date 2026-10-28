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

// FormFillingManageAction [0 - Stop, 1 - Resume, 2 - Start, 3 - Edit]
type FormFillingManageAction int32

// List of FormFillingManageAction
const (
	FORMFILLINGMANAGEACTION_Stop FormFillingManageAction = 0
	FORMFILLINGMANAGEACTION_Resume FormFillingManageAction = 1
	FORMFILLINGMANAGEACTION_Start FormFillingManageAction = 2
	FORMFILLINGMANAGEACTION_Edit FormFillingManageAction = 3
)

// All allowed values of FormFillingManageAction enum
var AllowedFormFillingManageActionEnumValues = []FormFillingManageAction{
	0,
	1,
	2,
	3,
}

func (v *FormFillingManageAction) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FormFillingManageAction(value)
	for _, existing := range AllowedFormFillingManageActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FormFillingManageAction", value)
}

// NewFormFillingManageActionFromValue returns a pointer to a valid FormFillingManageAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormFillingManageActionFromValue(v int32) (*FormFillingManageAction, error) {
	ev := FormFillingManageAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormFillingManageAction: valid values are %v", v, AllowedFormFillingManageActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormFillingManageAction) IsValid() bool {
	for _, existing := range AllowedFormFillingManageActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormFillingManageAction value
func (v FormFillingManageAction) Ptr() *FormFillingManageAction {
	return &v
}

type NullableFormFillingManageAction struct {
	value *FormFillingManageAction
	isSet bool
}

func (v NullableFormFillingManageAction) Get() *FormFillingManageAction {
	return v.value
}

func (v *NullableFormFillingManageAction) Set(val *FormFillingManageAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFillingManageAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFillingManageAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFillingManageAction(val *FormFillingManageAction) *NullableFormFillingManageAction {
	return &NullableFormFillingManageAction{value: val, isSet: true}
}

func (v NullableFormFillingManageAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFillingManageAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

