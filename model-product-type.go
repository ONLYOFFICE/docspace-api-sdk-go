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

// ProductType [2 - Documents, 3 - Login, 7 - Settings, 8 - Contacts]
type ProductType int32

// List of ProductType
const (
	PRODUCTTYPE_Documents ProductType = 2
	PRODUCTTYPE_Login ProductType = 3
	PRODUCTTYPE_Settings ProductType = 7
	PRODUCTTYPE_Contacts ProductType = 8
)

// All allowed values of ProductType enum
var AllowedProductTypeEnumValues = []ProductType{
	2,
	3,
	7,
	8,
}

func (v *ProductType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductType(value)
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductType", value)
}

// NewProductTypeFromValue returns a pointer to a valid ProductType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductTypeFromValue(v int32) (*ProductType, error) {
	ev := ProductType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductType: valid values are %v", v, AllowedProductTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductType) IsValid() bool {
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductType value
func (v ProductType) Ptr() *ProductType {
	return &v
}

type NullableProductType struct {
	value *ProductType
	isSet bool
}

func (v NullableProductType) Get() *ProductType {
	return v.value
}

func (v *NullableProductType) Set(val *ProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductType(val *ProductType) *NullableProductType {
	return &NullableProductType{value: val, isSet: true}
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

