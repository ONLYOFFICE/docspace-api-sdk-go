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

// TenantIndustry [0 - Other, 1 - Accounting, 2 - Advertising marketing PR, 3 - Banking, 4 - Consulting, 5 - Design, 6 - Education, 7 - Environment, 8 - Financial services, 9 - Health care, 10 - IT, 11 - Legal, 12 - Manufacturing, 13 - Public sector, 14 - Publishing, 15 - Retail sales, 16 - Telecommunications]
type TenantIndustry int32

// List of TenantIndustry
const (
	TENANTINDUSTRY_Other TenantIndustry = 0
	TENANTINDUSTRY_Accounting TenantIndustry = 1
	TENANTINDUSTRY_AdvertisingMarketingPR TenantIndustry = 2
	TENANTINDUSTRY_Banking TenantIndustry = 3
	TENANTINDUSTRY_Consulting TenantIndustry = 4
	TENANTINDUSTRY_Design TenantIndustry = 5
	TENANTINDUSTRY_Education TenantIndustry = 6
	TENANTINDUSTRY_Environment TenantIndustry = 7
	TENANTINDUSTRY_FinancialServices TenantIndustry = 8
	TENANTINDUSTRY_HealthCare TenantIndustry = 9
	TENANTINDUSTRY_IT TenantIndustry = 10
	TENANTINDUSTRY_Legal TenantIndustry = 11
	TENANTINDUSTRY_Manufacturing TenantIndustry = 12
	TENANTINDUSTRY_PublicSector TenantIndustry = 13
	TENANTINDUSTRY_Publishing TenantIndustry = 14
	TENANTINDUSTRY_RetailSales TenantIndustry = 15
	TENANTINDUSTRY_Telecommunications TenantIndustry = 16
)

// All allowed values of TenantIndustry enum
var AllowedTenantIndustryEnumValues = []TenantIndustry{
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
}

func (v *TenantIndustry) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenantIndustry(value)
	for _, existing := range AllowedTenantIndustryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenantIndustry", value)
}

// NewTenantIndustryFromValue returns a pointer to a valid TenantIndustry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenantIndustryFromValue(v int32) (*TenantIndustry, error) {
	ev := TenantIndustry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenantIndustry: valid values are %v", v, AllowedTenantIndustryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenantIndustry) IsValid() bool {
	for _, existing := range AllowedTenantIndustryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenantIndustry value
func (v TenantIndustry) Ptr() *TenantIndustry {
	return &v
}

type NullableTenantIndustry struct {
	value *TenantIndustry
	isSet bool
}

func (v NullableTenantIndustry) Get() *TenantIndustry {
	return v.value
}

func (v *NullableTenantIndustry) Set(val *TenantIndustry) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantIndustry) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantIndustry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantIndustry(val *TenantIndustry) *NullableTenantIndustry {
	return &NullableTenantIndustry{value: val, isSet: true}
}

func (v NullableTenantIndustry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantIndustry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

