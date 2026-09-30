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

// PriceTimeUnit [0 - None, 1 - Hour, 2 - Day, 3 - Week, 4 - Month, 5 - Year, 6 - ThreeYears]
type PriceTimeUnit int32

// List of PriceTimeUnit
const (
	PRICETIMEUNIT_None PriceTimeUnit = 0
	PRICETIMEUNIT_Hour PriceTimeUnit = 1
	PRICETIMEUNIT_Day PriceTimeUnit = 2
	PRICETIMEUNIT_Week PriceTimeUnit = 3
	PRICETIMEUNIT_Month PriceTimeUnit = 4
	PRICETIMEUNIT_Year PriceTimeUnit = 5
	PRICETIMEUNIT_ThreeYears PriceTimeUnit = 6
)

// All allowed values of PriceTimeUnit enum
var AllowedPriceTimeUnitEnumValues = []PriceTimeUnit{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

func (v *PriceTimeUnit) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceTimeUnit(value)
	for _, existing := range AllowedPriceTimeUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceTimeUnit", value)
}

// NewPriceTimeUnitFromValue returns a pointer to a valid PriceTimeUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceTimeUnitFromValue(v int32) (*PriceTimeUnit, error) {
	ev := PriceTimeUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceTimeUnit: valid values are %v", v, AllowedPriceTimeUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceTimeUnit) IsValid() bool {
	for _, existing := range AllowedPriceTimeUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceTimeUnit value
func (v PriceTimeUnit) Ptr() *PriceTimeUnit {
	return &v
}

type NullablePriceTimeUnit struct {
	value *PriceTimeUnit
	isSet bool
}

func (v NullablePriceTimeUnit) Get() *PriceTimeUnit {
	return v.value
}

func (v *NullablePriceTimeUnit) Set(val *PriceTimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTimeUnit(val *PriceTimeUnit) *NullablePriceTimeUnit {
	return &NullablePriceTimeUnit{value: val, isSet: true}
}

func (v NullablePriceTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

