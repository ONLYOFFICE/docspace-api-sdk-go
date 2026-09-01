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

// PaymentMethodStatus [0 - None, 1 - Set, 2 - Expired]
type PaymentMethodStatus int32

// List of PaymentMethodStatus
const (
	PAYMENTMETHODSTATUS_None PaymentMethodStatus = 0
	PAYMENTMETHODSTATUS_Set PaymentMethodStatus = 1
	PAYMENTMETHODSTATUS_Expired PaymentMethodStatus = 2
)

// All allowed values of PaymentMethodStatus enum
var AllowedPaymentMethodStatusEnumValues = []PaymentMethodStatus{
	0,
	1,
	2,
}

func (v *PaymentMethodStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentMethodStatus(value)
	for _, existing := range AllowedPaymentMethodStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentMethodStatus", value)
}

// NewPaymentMethodStatusFromValue returns a pointer to a valid PaymentMethodStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentMethodStatusFromValue(v int32) (*PaymentMethodStatus, error) {
	ev := PaymentMethodStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentMethodStatus: valid values are %v", v, AllowedPaymentMethodStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentMethodStatus) IsValid() bool {
	for _, existing := range AllowedPaymentMethodStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentMethodStatus value
func (v PaymentMethodStatus) Ptr() *PaymentMethodStatus {
	return &v
}

type NullablePaymentMethodStatus struct {
	value *PaymentMethodStatus
	isSet bool
}

func (v NullablePaymentMethodStatus) Get() *PaymentMethodStatus {
	return v.value
}

func (v *NullablePaymentMethodStatus) Set(val *PaymentMethodStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodStatus(val *PaymentMethodStatus) *NullablePaymentMethodStatus {
	return &NullablePaymentMethodStatus{value: val, isSet: true}
}

func (v NullablePaymentMethodStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

