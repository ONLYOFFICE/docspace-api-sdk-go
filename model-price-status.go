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

// PriceStatus [0 - Draft, 1 - Approved, 2 - Rejected]
type PriceStatus int32

// List of PriceStatus
const (
	PRICESTATUS_Draft PriceStatus = 0
	PRICESTATUS_Approved PriceStatus = 1
	PRICESTATUS_Rejected PriceStatus = 2
)

// All allowed values of PriceStatus enum
var AllowedPriceStatusEnumValues = []PriceStatus{
	0,
	1,
	2,
}

func (v *PriceStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceStatus(value)
	for _, existing := range AllowedPriceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceStatus", value)
}

// NewPriceStatusFromValue returns a pointer to a valid PriceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceStatusFromValue(v int32) (*PriceStatus, error) {
	ev := PriceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceStatus: valid values are %v", v, AllowedPriceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceStatus) IsValid() bool {
	for _, existing := range AllowedPriceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceStatus value
func (v PriceStatus) Ptr() *PriceStatus {
	return &v
}

type NullablePriceStatus struct {
	value *PriceStatus
	isSet bool
}

func (v NullablePriceStatus) Get() *PriceStatus {
	return v.value
}

func (v *NullablePriceStatus) Set(val *PriceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceStatus(val *PriceStatus) *NullablePriceStatus {
	return &NullablePriceStatus{value: val, isSet: true}
}

func (v NullablePriceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

