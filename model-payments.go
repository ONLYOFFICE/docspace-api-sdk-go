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

// Payments [0 - Paid, 1 - Free]
type Payments int32

// List of Payments
const (
	PAYMENTS_Paid Payments = 0
	PAYMENTS_Free Payments = 1
)

// All allowed values of Payments enum
var AllowedPaymentsEnumValues = []Payments{
	0,
	1,
}

func (v *Payments) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Payments(value)
	for _, existing := range AllowedPaymentsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Payments", value)
}

// NewPaymentsFromValue returns a pointer to a valid Payments
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentsFromValue(v int32) (*Payments, error) {
	ev := Payments(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Payments: valid values are %v", v, AllowedPaymentsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Payments) IsValid() bool {
	for _, existing := range AllowedPaymentsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Payments value
func (v Payments) Ptr() *Payments {
	return &v
}

type NullablePayments struct {
	value *Payments
	isSet bool
}

func (v NullablePayments) Get() *Payments {
	return v.value
}

func (v *NullablePayments) Set(val *Payments) {
	v.value = val
	v.isSet = true
}

func (v NullablePayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayments(val *Payments) *NullablePayments {
	return &NullablePayments{value: val, isSet: true}
}

func (v NullablePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

