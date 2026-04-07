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


// ContinueChatBodyFilesInner struct for ContinueChatBodyFilesInner
type ContinueChatBodyFilesInner struct {
	Int32 *int32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ContinueChatBodyFilesInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ContinueChatBodyFilesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContinueChatBodyFilesInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableContinueChatBodyFilesInner struct {
	value *ContinueChatBodyFilesInner
	isSet bool
}

func (v NullableContinueChatBodyFilesInner) Get() *ContinueChatBodyFilesInner {
	return v.value
}

func (v *NullableContinueChatBodyFilesInner) Set(val *ContinueChatBodyFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueChatBodyFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueChatBodyFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueChatBodyFilesInner(val *ContinueChatBodyFilesInner) *NullableContinueChatBodyFilesInner {
	return &NullableContinueChatBodyFilesInner{value: val, isSet: true}
}

func (v NullableContinueChatBodyFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueChatBodyFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


