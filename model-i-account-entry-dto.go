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
	"gopkg.in/validator.v2"
)

// IAccountEntryDto - One entry of an account search: either a user or a group.
type IAccountEntryDto struct {
	EmployeeFullDto *EmployeeFullDto
	GroupDto *GroupDto
}

// EmployeeFullDtoAsIAccountEntryDto is a convenience function that returns EmployeeFullDto wrapped in IAccountEntryDto
func EmployeeFullDtoAsIAccountEntryDto(v *EmployeeFullDto) IAccountEntryDto {
	return IAccountEntryDto{
		EmployeeFullDto: v,
	}
}

// GroupDtoAsIAccountEntryDto is a convenience function that returns GroupDto wrapped in IAccountEntryDto
func GroupDtoAsIAccountEntryDto(v *GroupDto) IAccountEntryDto {
	return IAccountEntryDto{
		GroupDto: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IAccountEntryDto) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EmployeeFullDto
	err = newStrictDecoder(data).Decode(&dst.EmployeeFullDto)
	if err == nil {
		jsonEmployeeFullDto, _ := json.Marshal(dst.EmployeeFullDto)
		if string(jsonEmployeeFullDto) == "{}" { // empty struct
			dst.EmployeeFullDto = nil
		} else {
			if err = validator.Validate(dst.EmployeeFullDto); err != nil {
				dst.EmployeeFullDto = nil
			} else {
				match++
			}
		}
	} else {
		dst.EmployeeFullDto = nil
	}

	// try to unmarshal data into GroupDto
	err = newStrictDecoder(data).Decode(&dst.GroupDto)
	if err == nil {
		jsonGroupDto, _ := json.Marshal(dst.GroupDto)
		if string(jsonGroupDto) == "{}" { // empty struct
			dst.GroupDto = nil
		} else {
			if err = validator.Validate(dst.GroupDto); err != nil {
				dst.GroupDto = nil
			} else {
				match++
			}
		}
	} else {
		dst.GroupDto = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EmployeeFullDto = nil
		dst.GroupDto = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IAccountEntryDto)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(IAccountEntryDto): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(IAccountEntryDto)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(IAccountEntryDto): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(IAccountEntryDto)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IAccountEntryDto) MarshalJSON() ([]byte, error) {
	if src.EmployeeFullDto != nil {
		return json.Marshal(&src.EmployeeFullDto)
	}

	if src.GroupDto != nil {
		return json.Marshal(&src.GroupDto)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IAccountEntryDto) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EmployeeFullDto != nil {
		return obj.EmployeeFullDto
	}

	if obj.GroupDto != nil {
		return obj.GroupDto
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj IAccountEntryDto) GetActualInstanceValue() (interface{}) {
	if obj.EmployeeFullDto != nil {
		return *obj.EmployeeFullDto
	}

	if obj.GroupDto != nil {
		return *obj.GroupDto
	}

	// all schemas are nil
	return nil
}

type NullableIAccountEntryDto struct {
	value *IAccountEntryDto
	isSet bool
}

func (v NullableIAccountEntryDto) Get() *IAccountEntryDto {
	return v.value
}

func (v *NullableIAccountEntryDto) Set(val *IAccountEntryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIAccountEntryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIAccountEntryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIAccountEntryDto(val *IAccountEntryDto) *NullableIAccountEntryDto {
	return &NullableIAccountEntryDto{value: val, isSet: true}
}

func (v NullableIAccountEntryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIAccountEntryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


