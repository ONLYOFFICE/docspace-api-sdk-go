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


// AiExportTextToDocxRequestFolderId Target folder id (int or string).
type AiExportTextToDocxRequestFolderId struct {
	Float32 *float32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiExportTextToDocxRequestFolderId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Float32
	err = json.Unmarshal(data, &dst.Float32);
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiExportTextToDocxRequestFolderId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiExportTextToDocxRequestFolderId) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiExportTextToDocxRequestFolderId struct {
	value *AiExportTextToDocxRequestFolderId
	isSet bool
}

func (v NullableAiExportTextToDocxRequestFolderId) Get() *AiExportTextToDocxRequestFolderId {
	return v.value
}

func (v *NullableAiExportTextToDocxRequestFolderId) Set(val *AiExportTextToDocxRequestFolderId) {
	v.value = val
	v.isSet = true
}

func (v NullableAiExportTextToDocxRequestFolderId) IsSet() bool {
	return v.isSet
}

func (v *NullableAiExportTextToDocxRequestFolderId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiExportTextToDocxRequestFolderId(val *AiExportTextToDocxRequestFolderId) *NullableAiExportTextToDocxRequestFolderId {
	return &NullableAiExportTextToDocxRequestFolderId{value: val, isSet: true}
}

func (v NullableAiExportTextToDocxRequestFolderId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiExportTextToDocxRequestFolderId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


