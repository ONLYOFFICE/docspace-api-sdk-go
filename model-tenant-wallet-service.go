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

// TenantWalletService [-13 - AITools, -12 - Backup, -11 - Storage]
type TenantWalletService int32

// List of TenantWalletService
const (
	TENANTWALLETSERVICE_AITools TenantWalletService = -13
	TENANTWALLETSERVICE_Backup TenantWalletService = -12
	TENANTWALLETSERVICE_Storage TenantWalletService = -11
)

// All allowed values of TenantWalletService enum
var AllowedTenantWalletServiceEnumValues = []TenantWalletService{
	-13,
	-12,
	-11,
}

func (v *TenantWalletService) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenantWalletService(value)
	for _, existing := range AllowedTenantWalletServiceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenantWalletService", value)
}

// NewTenantWalletServiceFromValue returns a pointer to a valid TenantWalletService
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenantWalletServiceFromValue(v int32) (*TenantWalletService, error) {
	ev := TenantWalletService(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenantWalletService: valid values are %v", v, AllowedTenantWalletServiceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenantWalletService) IsValid() bool {
	for _, existing := range AllowedTenantWalletServiceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenantWalletService value
func (v TenantWalletService) Ptr() *TenantWalletService {
	return &v
}

type NullableTenantWalletService struct {
	value *TenantWalletService
	isSet bool
}

func (v NullableTenantWalletService) Get() *TenantWalletService {
	return v.value
}

func (v *NullableTenantWalletService) Set(val *TenantWalletService) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWalletService) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWalletService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWalletService(val *TenantWalletService) *NullableTenantWalletService {
	return &NullableTenantWalletService{value: val, isSet: true}
}

func (v NullableTenantWalletService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWalletService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

