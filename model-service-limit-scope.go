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

// ServiceLimitScope [1 - Total, 2 - Global, 3 - Custom]
type ServiceLimitScope int32

// List of ServiceLimitScope
const (
	SERVICELIMITSCOPE_Total ServiceLimitScope = 1
	SERVICELIMITSCOPE_Global ServiceLimitScope = 2
	SERVICELIMITSCOPE_Custom ServiceLimitScope = 3
)

// All allowed values of ServiceLimitScope enum
var AllowedServiceLimitScopeEnumValues = []ServiceLimitScope{
	1,
	2,
	3,
}

func (v *ServiceLimitScope) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceLimitScope(value)
	for _, existing := range AllowedServiceLimitScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceLimitScope", value)
}

// NewServiceLimitScopeFromValue returns a pointer to a valid ServiceLimitScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceLimitScopeFromValue(v int32) (*ServiceLimitScope, error) {
	ev := ServiceLimitScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceLimitScope: valid values are %v", v, AllowedServiceLimitScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceLimitScope) IsValid() bool {
	for _, existing := range AllowedServiceLimitScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceLimitScope value
func (v ServiceLimitScope) Ptr() *ServiceLimitScope {
	return &v
}

type NullableServiceLimitScope struct {
	value *ServiceLimitScope
	isSet bool
}

func (v NullableServiceLimitScope) Get() *ServiceLimitScope {
	return v.value
}

func (v *NullableServiceLimitScope) Set(val *ServiceLimitScope) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitScope) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitScope(val *ServiceLimitScope) *NullableServiceLimitScope {
	return &NullableServiceLimitScope{value: val, isSet: true}
}

func (v NullableServiceLimitScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

