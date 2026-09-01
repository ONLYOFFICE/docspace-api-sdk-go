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

// QuotaFilter [0 - All, 1 - Default, 2 - Custom]
type QuotaFilter int32

// List of QuotaFilter
const (
	QUOTAFILTER_All QuotaFilter = 0
	QUOTAFILTER_Default QuotaFilter = 1
	QUOTAFILTER_Custom QuotaFilter = 2
)

// All allowed values of QuotaFilter enum
var AllowedQuotaFilterEnumValues = []QuotaFilter{
	0,
	1,
	2,
}

func (v *QuotaFilter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuotaFilter(value)
	for _, existing := range AllowedQuotaFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuotaFilter", value)
}

// NewQuotaFilterFromValue returns a pointer to a valid QuotaFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuotaFilterFromValue(v int32) (*QuotaFilter, error) {
	ev := QuotaFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuotaFilter: valid values are %v", v, AllowedQuotaFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuotaFilter) IsValid() bool {
	for _, existing := range AllowedQuotaFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuotaFilter value
func (v QuotaFilter) Ptr() *QuotaFilter {
	return &v
}

type NullableQuotaFilter struct {
	value *QuotaFilter
	isSet bool
}

func (v NullableQuotaFilter) Get() *QuotaFilter {
	return v.value
}

func (v *NullableQuotaFilter) Set(val *QuotaFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaFilter(val *QuotaFilter) *NullableQuotaFilter {
	return &NullableQuotaFilter{value: val, isSet: true}
}

func (v NullableQuotaFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

