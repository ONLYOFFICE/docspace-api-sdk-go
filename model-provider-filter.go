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

// ProviderFilter [0 - None, 1 - Box, 2 - DropBox, 3 - GoogleDrive, 4 - kDrive, 5 - OneDrive, 6 - SharePoint, 7 - WebDav, 8 - Yandex, 9 - Storage]
type ProviderFilter int32

// List of ProviderFilter
const (
	PROVIDERFILTER_None ProviderFilter = 0
	PROVIDERFILTER_Box ProviderFilter = 1
	PROVIDERFILTER_DropBox ProviderFilter = 2
	PROVIDERFILTER_GoogleDrive ProviderFilter = 3
	PROVIDERFILTER_kDrive ProviderFilter = 4
	PROVIDERFILTER_OneDrive ProviderFilter = 5
	PROVIDERFILTER_SharePoint ProviderFilter = 6
	PROVIDERFILTER_WebDav ProviderFilter = 7
	PROVIDERFILTER_Yandex ProviderFilter = 8
	PROVIDERFILTER_Storage ProviderFilter = 9
)

// All allowed values of ProviderFilter enum
var AllowedProviderFilterEnumValues = []ProviderFilter{
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
}

func (v *ProviderFilter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderFilter(value)
	for _, existing := range AllowedProviderFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderFilter", value)
}

// NewProviderFilterFromValue returns a pointer to a valid ProviderFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderFilterFromValue(v int32) (*ProviderFilter, error) {
	ev := ProviderFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderFilter: valid values are %v", v, AllowedProviderFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderFilter) IsValid() bool {
	for _, existing := range AllowedProviderFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderFilter value
func (v ProviderFilter) Ptr() *ProviderFilter {
	return &v
}

type NullableProviderFilter struct {
	value *ProviderFilter
	isSet bool
}

func (v NullableProviderFilter) Get() *ProviderFilter {
	return v.value
}

func (v *NullableProviderFilter) Set(val *ProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderFilter(val *ProviderFilter) *NullableProviderFilter {
	return &NullableProviderFilter{value: val, isSet: true}
}

func (v NullableProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

