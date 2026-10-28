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

// StorageFilter [0 - None, 1 - Internal, 2 - ThirdParty]
type StorageFilter int32

// List of StorageFilter
const (
	STORAGEFILTER_None StorageFilter = 0
	STORAGEFILTER_Internal StorageFilter = 1
	STORAGEFILTER_ThirdParty StorageFilter = 2
)

// All allowed values of StorageFilter enum
var AllowedStorageFilterEnumValues = []StorageFilter{
	0,
	1,
	2,
}

func (v *StorageFilter) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageFilter(value)
	for _, existing := range AllowedStorageFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageFilter", value)
}

// NewStorageFilterFromValue returns a pointer to a valid StorageFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageFilterFromValue(v int32) (*StorageFilter, error) {
	ev := StorageFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageFilter: valid values are %v", v, AllowedStorageFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageFilter) IsValid() bool {
	for _, existing := range AllowedStorageFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageFilter value
func (v StorageFilter) Ptr() *StorageFilter {
	return &v
}

type NullableStorageFilter struct {
	value *StorageFilter
	isSet bool
}

func (v NullableStorageFilter) Get() *StorageFilter {
	return v.value
}

func (v *NullableStorageFilter) Set(val *StorageFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFilter(val *StorageFilter) *NullableStorageFilter {
	return &NullableStorageFilter{value: val, isSet: true}
}

func (v NullableStorageFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

