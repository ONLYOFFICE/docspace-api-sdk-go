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

// MetadataConflictResolveType [0 - Skip existing values, 1 - Overwrite existing values]
type MetadataConflictResolveType int32

// List of MetadataConflictResolveType
const (
	METADATACONFLICTRESOLVETYPE_Skip MetadataConflictResolveType = 0
	METADATACONFLICTRESOLVETYPE_Overwrite MetadataConflictResolveType = 1
)

// All allowed values of MetadataConflictResolveType enum
var AllowedMetadataConflictResolveTypeEnumValues = []MetadataConflictResolveType{
	0,
	1,
}

func (v *MetadataConflictResolveType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataConflictResolveType(value)
	for _, existing := range AllowedMetadataConflictResolveTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataConflictResolveType", value)
}

// NewMetadataConflictResolveTypeFromValue returns a pointer to a valid MetadataConflictResolveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataConflictResolveTypeFromValue(v int32) (*MetadataConflictResolveType, error) {
	ev := MetadataConflictResolveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataConflictResolveType: valid values are %v", v, AllowedMetadataConflictResolveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataConflictResolveType) IsValid() bool {
	for _, existing := range AllowedMetadataConflictResolveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataConflictResolveType value
func (v MetadataConflictResolveType) Ptr() *MetadataConflictResolveType {
	return &v
}

type NullableMetadataConflictResolveType struct {
	value *MetadataConflictResolveType
	isSet bool
}

func (v NullableMetadataConflictResolveType) Get() *MetadataConflictResolveType {
	return v.value
}

func (v *NullableMetadataConflictResolveType) Set(val *MetadataConflictResolveType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataConflictResolveType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataConflictResolveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataConflictResolveType(val *MetadataConflictResolveType) *NullableMetadataConflictResolveType {
	return &NullableMetadataConflictResolveType{value: val, isSet: true}
}

func (v NullableMetadataConflictResolveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataConflictResolveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

