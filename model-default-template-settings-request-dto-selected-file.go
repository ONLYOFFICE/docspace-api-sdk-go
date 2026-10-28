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

// DefaultTemplateSettingsRequestDtoSelectedFile - The document to copy as the blank: a number for a file stored in the portal, a string for one in a connected  third-party storage. Take the identifier from a folder listing such as `GET api/2.0/files/{folderId}`; the  caller must be allowed to copy that file, and its extension must be the one named below.
type DefaultTemplateSettingsRequestDtoSelectedFile struct {
	Int32 *int32
	String *string
}

// int32AsDefaultTemplateSettingsRequestDtoSelectedFile is a convenience function that returns int32 wrapped in DefaultTemplateSettingsRequestDtoSelectedFile
func Int32AsDefaultTemplateSettingsRequestDtoSelectedFile(v *int32) DefaultTemplateSettingsRequestDtoSelectedFile {
	return DefaultTemplateSettingsRequestDtoSelectedFile{
		Int32: v,
	}
}

// stringAsDefaultTemplateSettingsRequestDtoSelectedFile is a convenience function that returns string wrapped in DefaultTemplateSettingsRequestDtoSelectedFile
func StringAsDefaultTemplateSettingsRequestDtoSelectedFile(v *string) DefaultTemplateSettingsRequestDtoSelectedFile {
	return DefaultTemplateSettingsRequestDtoSelectedFile{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultTemplateSettingsRequestDtoSelectedFile) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(DefaultTemplateSettingsRequestDtoSelectedFile)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(DefaultTemplateSettingsRequestDtoSelectedFile): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(DefaultTemplateSettingsRequestDtoSelectedFile)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(DefaultTemplateSettingsRequestDtoSelectedFile): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(DefaultTemplateSettingsRequestDtoSelectedFile)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultTemplateSettingsRequestDtoSelectedFile) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultTemplateSettingsRequestDtoSelectedFile) GetActualInstance() (interface{}) {
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
func (obj DefaultTemplateSettingsRequestDtoSelectedFile) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDefaultTemplateSettingsRequestDtoSelectedFile struct {
	value *DefaultTemplateSettingsRequestDtoSelectedFile
	isSet bool
}

func (v NullableDefaultTemplateSettingsRequestDtoSelectedFile) Get() *DefaultTemplateSettingsRequestDtoSelectedFile {
	return v.value
}

func (v *NullableDefaultTemplateSettingsRequestDtoSelectedFile) Set(val *DefaultTemplateSettingsRequestDtoSelectedFile) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultTemplateSettingsRequestDtoSelectedFile) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultTemplateSettingsRequestDtoSelectedFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultTemplateSettingsRequestDtoSelectedFile(val *DefaultTemplateSettingsRequestDtoSelectedFile) *NullableDefaultTemplateSettingsRequestDtoSelectedFile {
	return &NullableDefaultTemplateSettingsRequestDtoSelectedFile{value: val, isSet: true}
}

func (v NullableDefaultTemplateSettingsRequestDtoSelectedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultTemplateSettingsRequestDtoSelectedFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


