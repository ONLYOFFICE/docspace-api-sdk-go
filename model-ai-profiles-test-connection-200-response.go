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


// AiProfilesTestConnection200Response struct for AiProfilesTestConnection200Response
type AiProfilesTestConnection200Response struct {
	AiProfilesTestConnection200ResponseAnyOf *AiProfilesTestConnection200ResponseAnyOf
	Bool *bool
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiProfilesTestConnection200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AiProfilesTestConnection200ResponseAnyOf
	err = json.Unmarshal(data, &dst.AiProfilesTestConnection200ResponseAnyOf);
	if err == nil {
		jsonAiProfilesTestConnection200ResponseAnyOf, _ := json.Marshal(dst.AiProfilesTestConnection200ResponseAnyOf)
		if string(jsonAiProfilesTestConnection200ResponseAnyOf) == "{}" { // empty struct
			dst.AiProfilesTestConnection200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.AiProfilesTestConnection200ResponseAnyOf, return on the first match
		}
	} else {
		dst.AiProfilesTestConnection200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into Bool
	err = json.Unmarshal(data, &dst.Bool);
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			return nil // data stored in dst.Bool, return on the first match
		}
	} else {
		dst.Bool = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiProfilesTestConnection200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiProfilesTestConnection200Response) MarshalJSON() ([]byte, error) {
	if src.AiProfilesTestConnection200ResponseAnyOf != nil {
		return json.Marshal(&src.AiProfilesTestConnection200ResponseAnyOf)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiProfilesTestConnection200Response struct {
	value *AiProfilesTestConnection200Response
	isSet bool
}

func (v NullableAiProfilesTestConnection200Response) Get() *AiProfilesTestConnection200Response {
	return v.value
}

func (v *NullableAiProfilesTestConnection200Response) Set(val *AiProfilesTestConnection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesTestConnection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesTestConnection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesTestConnection200Response(val *AiProfilesTestConnection200Response) *NullableAiProfilesTestConnection200Response {
	return &NullableAiProfilesTestConnection200Response{value: val, isSet: true}
}

func (v NullableAiProfilesTestConnection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesTestConnection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


