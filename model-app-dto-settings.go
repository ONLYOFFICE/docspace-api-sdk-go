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

// AppDtoSettings - The settings document saved for this portal, stored and returned verbatim - the portal never looks inside  it, and only the application knows its shape. It is empty while the portal has saved none, which means the  application falls back to its own defaults, and it also survives the application being switched off.
type AppDtoSettings struct {
	Int32 *int32
	String *string
}

// int32AsAppDtoSettings is a convenience function that returns int32 wrapped in AppDtoSettings
func Int32AsAppDtoSettings(v *int32) AppDtoSettings {
	return AppDtoSettings{
		Int32: v,
	}
}

// stringAsAppDtoSettings is a convenience function that returns string wrapped in AppDtoSettings
func StringAsAppDtoSettings(v *string) AppDtoSettings {
	return AppDtoSettings{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppDtoSettings) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(AppDtoSettings)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(AppDtoSettings): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(AppDtoSettings)")
        }
        if err != nil {
            return fmt.Errorf("data failed to match schemas in oneOf(AppDtoSettings): %v", err)
        } else {
            return fmt.Errorf("data failed to match schemas in oneOf(AppDtoSettings)")
        }
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppDtoSettings) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppDtoSettings) GetActualInstance() (interface{}) {
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
func (obj AppDtoSettings) GetActualInstanceValue() (interface{}) {
	if obj.Int32 != nil {
		return *obj.Int32
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAppDtoSettings struct {
	value *AppDtoSettings
	isSet bool
}

func (v NullableAppDtoSettings) Get() *AppDtoSettings {
	return v.value
}

func (v *NullableAppDtoSettings) Set(val *AppDtoSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDtoSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDtoSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDtoSettings(val *AppDtoSettings) *NullableAppDtoSettings {
	return &NullableAppDtoSettings{value: val, isSet: true}
}

func (v NullableAppDtoSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDtoSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


