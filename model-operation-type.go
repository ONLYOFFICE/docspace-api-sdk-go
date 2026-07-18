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

// OperationType [0 - Unknown, 1 - ServicePayment, 2 - PackagePayment, 3 - AiServicePayment, 4 - Deposit, 5 - ReceiveProviderInvoice, 6 - ProcessProviderInvoice, 7 - WriteOffServiceProfit, 8 - Profit, 9 - PartnerAccrual, 10 - ProviderPayment, 11 - PartnerPayment, 12 - Refund, 13 - BankDeposit, 14 - BankWithdrawal, 15 - GoodwillCredit, 16 - WriteOffProfit, 17 - WriteOffDifferenceCurrency, 18 - AiDebit, 19 - AiCredit]
type OperationType int32

// List of OperationType
const (
	OPERATIONTYPE_Unknown OperationType = 0
	OPERATIONTYPE_ServicePayment OperationType = 1
	OPERATIONTYPE_PackagePayment OperationType = 2
	OPERATIONTYPE_AiServicePayment OperationType = 3
	OPERATIONTYPE_Deposit OperationType = 4
	OPERATIONTYPE_ReceiveProviderInvoice OperationType = 5
	OPERATIONTYPE_ProcessProviderInvoice OperationType = 6
	OPERATIONTYPE_WriteOffServiceProfit OperationType = 7
	OPERATIONTYPE_Profit OperationType = 8
	OPERATIONTYPE_PartnerAccrual OperationType = 9
	OPERATIONTYPE_ProviderPayment OperationType = 10
	OPERATIONTYPE_PartnerPayment OperationType = 11
	OPERATIONTYPE_Refund OperationType = 12
	OPERATIONTYPE_BankDeposit OperationType = 13
	OPERATIONTYPE_BankWithdrawal OperationType = 14
	OPERATIONTYPE_GoodwillCredit OperationType = 15
	OPERATIONTYPE_WriteOffProfit OperationType = 16
	OPERATIONTYPE_WriteOffDifferenceCurrency OperationType = 17
	OPERATIONTYPE_AiDebit OperationType = 18
	OPERATIONTYPE_AiCredit OperationType = 19
)

// All allowed values of OperationType enum
var AllowedOperationTypeEnumValues = []OperationType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
}

func (v *OperationType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationType(value)
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationType", value)
}

// NewOperationTypeFromValue returns a pointer to a valid OperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationTypeFromValue(v int32) (*OperationType, error) {
	ev := OperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationType: valid values are %v", v, AllowedOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationType) IsValid() bool {
	for _, existing := range AllowedOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationType value
func (v OperationType) Ptr() *OperationType {
	return &v
}

type NullableOperationType struct {
	value *OperationType
	isSet bool
}

func (v NullableOperationType) Get() *OperationType {
	return v.value
}

func (v *NullableOperationType) Set(val *OperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationType(val *OperationType) *NullableOperationType {
	return &NullableOperationType{value: val, isSet: true}
}

func (v NullableOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

