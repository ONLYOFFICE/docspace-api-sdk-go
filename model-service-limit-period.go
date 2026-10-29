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

// ServiceLimitPeriod [1 - Day, 2 - Week, 3 - Month, 4 - Lifetime]
type ServiceLimitPeriod int32

// List of ServiceLimitPeriod
const (
	SERVICELIMITPERIOD_Day ServiceLimitPeriod = 1
	SERVICELIMITPERIOD_Week ServiceLimitPeriod = 2
	SERVICELIMITPERIOD_Month ServiceLimitPeriod = 3
	SERVICELIMITPERIOD_Lifetime ServiceLimitPeriod = 4
)

// All allowed values of ServiceLimitPeriod enum
var AllowedServiceLimitPeriodEnumValues = []ServiceLimitPeriod{
	1,
	2,
	3,
	4,
}

func (v *ServiceLimitPeriod) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceLimitPeriod(value)
	for _, existing := range AllowedServiceLimitPeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceLimitPeriod", value)
}

// NewServiceLimitPeriodFromValue returns a pointer to a valid ServiceLimitPeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceLimitPeriodFromValue(v int32) (*ServiceLimitPeriod, error) {
	ev := ServiceLimitPeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceLimitPeriod: valid values are %v", v, AllowedServiceLimitPeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceLimitPeriod) IsValid() bool {
	for _, existing := range AllowedServiceLimitPeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceLimitPeriod value
func (v ServiceLimitPeriod) Ptr() *ServiceLimitPeriod {
	return &v
}

type NullableServiceLimitPeriod struct {
	value *ServiceLimitPeriod
	isSet bool
}

func (v NullableServiceLimitPeriod) Get() *ServiceLimitPeriod {
	return v.value
}

func (v *NullableServiceLimitPeriod) Set(val *ServiceLimitPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitPeriod(val *ServiceLimitPeriod) *NullableServiceLimitPeriod {
	return &NullableServiceLimitPeriod{value: val, isSet: true}
}

func (v NullableServiceLimitPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

