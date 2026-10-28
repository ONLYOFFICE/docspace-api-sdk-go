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

// CreateFileRequestTemplateId - An existing file the new one copies its content from, as a number for a file in the portal and as a string for  one in a connected third-party storage; the caller has to be able to read it. Left out, a blank template for  the format and the language of the caller is used.
type CreateFileRequestTemplateId struct {
	Int32 *int32
	String *string
}

// int32AsCreateFileRequestTemplateId is a convenience function that returns int32 wrapped in CreateFileRequestTemplateId
func Int32AsCreateFileRequestTemplateId(v *int32) CreateFileRequestTemplateId {
	return CreateFileRequestTemplateId{
		Int32: v,
	}
}

// stringAsCreateFileRequestTemplateId is a convenience function that returns string wrapped in CreateFileRequestTemplateId
func StringAsCreateFileRequestTemplateId(v *string) CreateFileRequestTemplateId {
	return CreateFileRequestTemplateId{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFileRequestTemplateId) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(CreateFileRequestTemplateId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(CreateFileRequestTemplateId): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(CreateFileRequestTemplateId)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(CreateFileRequestTemplateId): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(CreateFileRequestTemplateId)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFileRequestTemplateId) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFileRequestTemplateId) GetActualInstance() (interface{}) {
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
func (obj CreateFileRequestTemplateId) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateFileRequestTemplateId struct {
	value *CreateFileRequestTemplateId
	isSet bool
}

func (v NullableCreateFileRequestTemplateId) Get() *CreateFileRequestTemplateId {
	return v.value
}

func (v *NullableCreateFileRequestTemplateId) Set(val *CreateFileRequestTemplateId) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFileRequestTemplateId) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFileRequestTemplateId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFileRequestTemplateId(val *CreateFileRequestTemplateId) *NullableCreateFileRequestTemplateId {
	return &NullableCreateFileRequestTemplateId{value: val, isSet: true}
}

func (v NullableCreateFileRequestTemplateId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFileRequestTemplateId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


