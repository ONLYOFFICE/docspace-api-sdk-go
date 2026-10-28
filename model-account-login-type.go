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

// AccountLoginType [0 - SSO, 1 - LDAP, 2 - Standart]
type AccountLoginType int32

// List of AccountLoginType
const (
	ACCOUNTLOGINTYPE_SSO AccountLoginType = 0
	ACCOUNTLOGINTYPE_LDAP AccountLoginType = 1
	ACCOUNTLOGINTYPE_Standart AccountLoginType = 2
)

// All allowed values of AccountLoginType enum
var AllowedAccountLoginTypeEnumValues = []AccountLoginType{
	0,
	1,
	2,
}

func (v *AccountLoginType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountLoginType(value)
	for _, existing := range AllowedAccountLoginTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountLoginType", value)
}

// NewAccountLoginTypeFromValue returns a pointer to a valid AccountLoginType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountLoginTypeFromValue(v int32) (*AccountLoginType, error) {
	ev := AccountLoginType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountLoginType: valid values are %v", v, AllowedAccountLoginTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountLoginType) IsValid() bool {
	for _, existing := range AllowedAccountLoginTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountLoginType value
func (v AccountLoginType) Ptr() *AccountLoginType {
	return &v
}

type NullableAccountLoginType struct {
	value *AccountLoginType
	isSet bool
}

func (v NullableAccountLoginType) Get() *AccountLoginType {
	return v.value
}

func (v *NullableAccountLoginType) Set(val *AccountLoginType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLoginType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLoginType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLoginType(val *AccountLoginType) *NullableAccountLoginType {
	return &NullableAccountLoginType{value: val, isSet: true}
}

func (v NullableAccountLoginType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLoginType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

