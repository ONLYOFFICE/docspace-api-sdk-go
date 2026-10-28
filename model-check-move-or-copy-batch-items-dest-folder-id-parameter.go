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

// CheckMoveOrCopyBatchItemsDestFolderIdParameter - The folder the items go to, by id — a number for a folder stored in the portal itself, a string for a folder  on a connected third-party account. Take it from a folder listing such as `GET api/2.0/files/@root`; the  caller has to be allowed to create items in it, and the id of a room addresses the root of that room.
type CheckMoveOrCopyBatchItemsDestFolderIdParameter struct {
	Int32 *int32
	String *string
}

// int32AsCheckMoveOrCopyBatchItemsDestFolderIdParameter is a convenience function that returns int32 wrapped in CheckMoveOrCopyBatchItemsDestFolderIdParameter
func Int32AsCheckMoveOrCopyBatchItemsDestFolderIdParameter(v *int32) CheckMoveOrCopyBatchItemsDestFolderIdParameter {
	return CheckMoveOrCopyBatchItemsDestFolderIdParameter{
		Int32: v,
	}
}

// stringAsCheckMoveOrCopyBatchItemsDestFolderIdParameter is a convenience function that returns string wrapped in CheckMoveOrCopyBatchItemsDestFolderIdParameter
func StringAsCheckMoveOrCopyBatchItemsDestFolderIdParameter(v *string) CheckMoveOrCopyBatchItemsDestFolderIdParameter {
	return CheckMoveOrCopyBatchItemsDestFolderIdParameter{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CheckMoveOrCopyBatchItemsDestFolderIdParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CheckMoveOrCopyBatchItemsDestFolderIdParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(CheckMoveOrCopyBatchItemsDestFolderIdParameter): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(CheckMoveOrCopyBatchItemsDestFolderIdParameter)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(CheckMoveOrCopyBatchItemsDestFolderIdParameter): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(CheckMoveOrCopyBatchItemsDestFolderIdParameter)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CheckMoveOrCopyBatchItemsDestFolderIdParameter) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CheckMoveOrCopyBatchItemsDestFolderIdParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CheckMoveOrCopyBatchItemsDestFolderIdParameter) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter struct {
	value *CheckMoveOrCopyBatchItemsDestFolderIdParameter
	isSet bool
}

func (v NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) Get() *CheckMoveOrCopyBatchItemsDestFolderIdParameter {
	return v.value
}

func (v *NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) Set(val *CheckMoveOrCopyBatchItemsDestFolderIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckMoveOrCopyBatchItemsDestFolderIdParameter(val *CheckMoveOrCopyBatchItemsDestFolderIdParameter) *NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter {
	return &NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter{value: val, isSet: true}
}

func (v NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckMoveOrCopyBatchItemsDestFolderIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


