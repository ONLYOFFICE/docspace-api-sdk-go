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

// TenantStatus [0 - Active, 1 - Suspended, 2 - Remove pending, 3 - Transfering, 4 - Restoring, 5 - Migrating, 6 - Encryption]
type TenantStatus int32

// List of TenantStatus
const (
	TENANTSTATUS_Active TenantStatus = 0
	TENANTSTATUS_Suspended TenantStatus = 1
	TENANTSTATUS_RemovePending TenantStatus = 2
	TENANTSTATUS_Transfering TenantStatus = 3
	TENANTSTATUS_Restoring TenantStatus = 4
	TENANTSTATUS_Migrating TenantStatus = 5
	TENANTSTATUS_Encryption TenantStatus = 6
)

// All allowed values of TenantStatus enum
var AllowedTenantStatusEnumValues = []TenantStatus{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

func (v *TenantStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenantStatus(value)
	for _, existing := range AllowedTenantStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenantStatus", value)
}

// NewTenantStatusFromValue returns a pointer to a valid TenantStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenantStatusFromValue(v int32) (*TenantStatus, error) {
	ev := TenantStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenantStatus: valid values are %v", v, AllowedTenantStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenantStatus) IsValid() bool {
	for _, existing := range AllowedTenantStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenantStatus value
func (v TenantStatus) Ptr() *TenantStatus {
	return &v
}

type NullableTenantStatus struct {
	value *TenantStatus
	isSet bool
}

func (v NullableTenantStatus) Get() *TenantStatus {
	return v.value
}

func (v *NullableTenantStatus) Set(val *TenantStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantStatus(val *TenantStatus) *NullableTenantStatus {
	return &NullableTenantStatus{value: val, isSet: true}
}

func (v NullableTenantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

