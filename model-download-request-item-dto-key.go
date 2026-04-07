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

// DownloadRequestItemDtoKey - The unique identifier or reference key for the file to be downloaded.
type DownloadRequestItemDtoKey struct {
	Int32 *int32
	String *string
}

// int32AsDownloadRequestItemDtoKey is a convenience function that returns int32 wrapped in DownloadRequestItemDtoKey
func Int32AsDownloadRequestItemDtoKey(v *int32) DownloadRequestItemDtoKey {
	return DownloadRequestItemDtoKey{
		Int32: v,
	}
}

// stringAsDownloadRequestItemDtoKey is a convenience function that returns string wrapped in DownloadRequestItemDtoKey
func StringAsDownloadRequestItemDtoKey(v *string) DownloadRequestItemDtoKey {
	return DownloadRequestItemDtoKey{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DownloadRequestItemDtoKey) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(DownloadRequestItemDtoKey)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DownloadRequestItemDtoKey)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DownloadRequestItemDtoKey) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DownloadRequestItemDtoKey) GetActualInstance() (interface{}) {
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
func (obj DownloadRequestItemDtoKey) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDownloadRequestItemDtoKey struct {
	value *DownloadRequestItemDtoKey
	isSet bool
}

func (v NullableDownloadRequestItemDtoKey) Get() *DownloadRequestItemDtoKey {
	return v.value
}

func (v *NullableDownloadRequestItemDtoKey) Set(val *DownloadRequestItemDtoKey) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadRequestItemDtoKey) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadRequestItemDtoKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadRequestItemDtoKey(val *DownloadRequestItemDtoKey) *NullableDownloadRequestItemDtoKey {
	return &NullableDownloadRequestItemDtoKey{value: val, isSet: true}
}

func (v NullableDownloadRequestItemDtoKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadRequestItemDtoKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


