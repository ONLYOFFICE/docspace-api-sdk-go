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

// TenantTrustedDomainsType [0 - None, 1 - Custom, 2 - All]
type TenantTrustedDomainsType int32

// List of TenantTrustedDomainsType
const (
	TENANTTRUSTEDDOMAINSTYPE_None TenantTrustedDomainsType = 0
	TENANTTRUSTEDDOMAINSTYPE_Custom TenantTrustedDomainsType = 1
	TENANTTRUSTEDDOMAINSTYPE_All TenantTrustedDomainsType = 2
)

// All allowed values of TenantTrustedDomainsType enum
var AllowedTenantTrustedDomainsTypeEnumValues = []TenantTrustedDomainsType{
	0,
	1,
	2,
}

func (v *TenantTrustedDomainsType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenantTrustedDomainsType(value)
	for _, existing := range AllowedTenantTrustedDomainsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenantTrustedDomainsType", value)
}

// NewTenantTrustedDomainsTypeFromValue returns a pointer to a valid TenantTrustedDomainsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenantTrustedDomainsTypeFromValue(v int32) (*TenantTrustedDomainsType, error) {
	ev := TenantTrustedDomainsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenantTrustedDomainsType: valid values are %v", v, AllowedTenantTrustedDomainsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenantTrustedDomainsType) IsValid() bool {
	for _, existing := range AllowedTenantTrustedDomainsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenantTrustedDomainsType value
func (v TenantTrustedDomainsType) Ptr() *TenantTrustedDomainsType {
	return &v
}

type NullableTenantTrustedDomainsType struct {
	value *TenantTrustedDomainsType
	isSet bool
}

func (v NullableTenantTrustedDomainsType) Get() *TenantTrustedDomainsType {
	return v.value
}

func (v *NullableTenantTrustedDomainsType) Set(val *TenantTrustedDomainsType) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantTrustedDomainsType) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantTrustedDomainsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantTrustedDomainsType(val *TenantTrustedDomainsType) *NullableTenantTrustedDomainsType {
	return &NullableTenantTrustedDomainsType{value: val, isSet: true}
}

func (v NullableTenantTrustedDomainsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantTrustedDomainsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

