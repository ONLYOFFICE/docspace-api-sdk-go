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

// MetadataFieldType [0 - String, 1 - Date, 2 - Number, 3 - Single choice, 4 - Multiple choice]
type MetadataFieldType int32

// List of MetadataFieldType
const (
	METADATAFIELDTYPE_String MetadataFieldType = 0
	METADATAFIELDTYPE_Date MetadataFieldType = 1
	METADATAFIELDTYPE_Number MetadataFieldType = 2
	METADATAFIELDTYPE_SingleChoice MetadataFieldType = 3
	METADATAFIELDTYPE_MultiChoice MetadataFieldType = 4
)

// All allowed values of MetadataFieldType enum
var AllowedMetadataFieldTypeEnumValues = []MetadataFieldType{
	0,
	1,
	2,
	3,
	4,
}

func (v *MetadataFieldType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataFieldType(value)
	for _, existing := range AllowedMetadataFieldTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataFieldType", value)
}

// NewMetadataFieldTypeFromValue returns a pointer to a valid MetadataFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataFieldTypeFromValue(v int32) (*MetadataFieldType, error) {
	ev := MetadataFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataFieldType: valid values are %v", v, AllowedMetadataFieldTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataFieldType) IsValid() bool {
	for _, existing := range AllowedMetadataFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataFieldType value
func (v MetadataFieldType) Ptr() *MetadataFieldType {
	return &v
}

type NullableMetadataFieldType struct {
	value *MetadataFieldType
	isSet bool
}

func (v NullableMetadataFieldType) Get() *MetadataFieldType {
	return v.value
}

func (v *NullableMetadataFieldType) Set(val *MetadataFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFieldType(val *MetadataFieldType) *NullableMetadataFieldType {
	return &NullableMetadataFieldType{value: val, isSet: true}
}

func (v NullableMetadataFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

