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

// OperationType [0 - Any, 1 - Unknown, 2 - ServicePayment, 4 - PackagePayment, 8 - ServiceUsage, 16 - Deposit, 32 - ReceiveProviderInvoice, 64 - ProcessProviderInvoice, 128 - WriteOffServiceProfit, 256 - Profit, 512 - PartnerAccrual, 1024 - ProviderPayment, 2048 - PartnerPayment, 4096 - Refund, 8192 - BankDeposit, 16384 - BankWithdrawal, 32768 - GoodwillCredit, 65536 - WriteOffProfit, 131072 - WriteOffDifferenceCurrency]
type OperationType int32

// List of OperationType
const (
	OPERATIONTYPE_Any OperationType = 0
	OPERATIONTYPE_Unknown OperationType = 1
	OPERATIONTYPE_ServicePayment OperationType = 2
	OPERATIONTYPE_PackagePayment OperationType = 4
	OPERATIONTYPE_ServiceUsage OperationType = 8
	OPERATIONTYPE_Deposit OperationType = 16
	OPERATIONTYPE_ReceiveProviderInvoice OperationType = 32
	OPERATIONTYPE_ProcessProviderInvoice OperationType = 64
	OPERATIONTYPE_WriteOffServiceProfit OperationType = 128
	OPERATIONTYPE_Profit OperationType = 256
	OPERATIONTYPE_PartnerAccrual OperationType = 512
	OPERATIONTYPE_ProviderPayment OperationType = 1024
	OPERATIONTYPE_PartnerPayment OperationType = 2048
	OPERATIONTYPE_Refund OperationType = 4096
	OPERATIONTYPE_BankDeposit OperationType = 8192
	OPERATIONTYPE_BankWithdrawal OperationType = 16384
	OPERATIONTYPE_GoodwillCredit OperationType = 32768
	OPERATIONTYPE_WriteOffProfit OperationType = 65536
	OPERATIONTYPE_WriteOffDifferenceCurrency OperationType = 131072
)

// All allowed values of OperationType enum
var AllowedOperationTypeEnumValues = []OperationType{
	0,
	1,
	2,
	4,
	8,
	16,
	32,
	64,
	128,
	256,
	512,
	1024,
	2048,
	4096,
	8192,
	16384,
	32768,
	65536,
	131072,
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

