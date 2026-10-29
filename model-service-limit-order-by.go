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

// ServiceLimitOrderBy [0 - Participant, 1 - Created]
type ServiceLimitOrderBy int32

// List of ServiceLimitOrderBy
const (
	SERVICELIMITORDERBY_Participant ServiceLimitOrderBy = 0
	SERVICELIMITORDERBY_Created ServiceLimitOrderBy = 1
)

// All allowed values of ServiceLimitOrderBy enum
var AllowedServiceLimitOrderByEnumValues = []ServiceLimitOrderBy{
	0,
	1,
}

func (v *ServiceLimitOrderBy) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceLimitOrderBy(value)
	for _, existing := range AllowedServiceLimitOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceLimitOrderBy", value)
}

// NewServiceLimitOrderByFromValue returns a pointer to a valid ServiceLimitOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceLimitOrderByFromValue(v int32) (*ServiceLimitOrderBy, error) {
	ev := ServiceLimitOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceLimitOrderBy: valid values are %v", v, AllowedServiceLimitOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceLimitOrderBy) IsValid() bool {
	for _, existing := range AllowedServiceLimitOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceLimitOrderBy value
func (v ServiceLimitOrderBy) Ptr() *ServiceLimitOrderBy {
	return &v
}

type NullableServiceLimitOrderBy struct {
	value *ServiceLimitOrderBy
	isSet bool
}

func (v NullableServiceLimitOrderBy) Get() *ServiceLimitOrderBy {
	return v.value
}

func (v *NullableServiceLimitOrderBy) Set(val *ServiceLimitOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitOrderBy(val *ServiceLimitOrderBy) *NullableServiceLimitOrderBy {
	return &NullableServiceLimitOrderBy{value: val, isSet: true}
}

func (v NullableServiceLimitOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

