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

// ProductQuantityType [0 - Set, 1 - Add, 2 - Sub, 3 - Renew]
type ProductQuantityType int32

// List of ProductQuantityType
const (
	PRODUCTQUANTITYTYPE_Set ProductQuantityType = 0
	PRODUCTQUANTITYTYPE_Add ProductQuantityType = 1
	PRODUCTQUANTITYTYPE_Sub ProductQuantityType = 2
	PRODUCTQUANTITYTYPE_Renew ProductQuantityType = 3
)

// All allowed values of ProductQuantityType enum
var AllowedProductQuantityTypeEnumValues = []ProductQuantityType{
	0,
	1,
	2,
	3,
}

func (v *ProductQuantityType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductQuantityType(value)
	for _, existing := range AllowedProductQuantityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductQuantityType", value)
}

// NewProductQuantityTypeFromValue returns a pointer to a valid ProductQuantityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductQuantityTypeFromValue(v int32) (*ProductQuantityType, error) {
	ev := ProductQuantityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductQuantityType: valid values are %v", v, AllowedProductQuantityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductQuantityType) IsValid() bool {
	for _, existing := range AllowedProductQuantityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductQuantityType value
func (v ProductQuantityType) Ptr() *ProductQuantityType {
	return &v
}

type NullableProductQuantityType struct {
	value *ProductQuantityType
	isSet bool
}

func (v NullableProductQuantityType) Get() *ProductQuantityType {
	return v.value
}

func (v *NullableProductQuantityType) Set(val *ProductQuantityType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductQuantityType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductQuantityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductQuantityType(val *ProductQuantityType) *NullableProductQuantityType {
	return &NullableProductQuantityType{value: val, isSet: true}
}

func (v NullableProductQuantityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductQuantityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

