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

// QuotaScope [0 - User, 1 - Room, 2 - Tenant]
type QuotaScope int32

// List of QuotaScope
const (
	QUOTASCOPE_User QuotaScope = 0
	QUOTASCOPE_Room QuotaScope = 1
	QUOTASCOPE_Tenant QuotaScope = 2
)

// All allowed values of QuotaScope enum
var AllowedQuotaScopeEnumValues = []QuotaScope{
	0,
	1,
	2,
}

func (v *QuotaScope) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuotaScope(value)
	for _, existing := range AllowedQuotaScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuotaScope", value)
}

// NewQuotaScopeFromValue returns a pointer to a valid QuotaScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuotaScopeFromValue(v int32) (*QuotaScope, error) {
	ev := QuotaScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuotaScope: valid values are %v", v, AllowedQuotaScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuotaScope) IsValid() bool {
	for _, existing := range AllowedQuotaScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuotaScope value
func (v QuotaScope) Ptr() *QuotaScope {
	return &v
}

type NullableQuotaScope struct {
	value *QuotaScope
	isSet bool
}

func (v NullableQuotaScope) Get() *QuotaScope {
	return v.value
}

func (v *NullableQuotaScope) Set(val *QuotaScope) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaScope) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaScope(val *QuotaScope) *NullableQuotaScope {
	return &NullableQuotaScope{value: val, isSet: true}
}

func (v NullableQuotaScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

