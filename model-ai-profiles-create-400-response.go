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


// AiProfilesCreate400Response struct for AiProfilesCreate400Response
type AiProfilesCreate400Response struct {
	AiErrorResponse *AiErrorResponse
	AiProfileMutationResult *AiProfileMutationResult
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiProfilesCreate400Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AiErrorResponse
	err = json.Unmarshal(data, &dst.AiErrorResponse);
	if err == nil {
		jsonAiErrorResponse, _ := json.Marshal(dst.AiErrorResponse)
		if string(jsonAiErrorResponse) == "{}" { // empty struct
			dst.AiErrorResponse = nil
		} else {
			return nil // data stored in dst.AiErrorResponse, return on the first match
		}
	} else {
		dst.AiErrorResponse = nil
	}

	// try to unmarshal JSON data into AiProfileMutationResult
	err = json.Unmarshal(data, &dst.AiProfileMutationResult);
	if err == nil {
		jsonAiProfileMutationResult, _ := json.Marshal(dst.AiProfileMutationResult)
		if string(jsonAiProfileMutationResult) == "{}" { // empty struct
			dst.AiProfileMutationResult = nil
		} else {
			return nil // data stored in dst.AiProfileMutationResult, return on the first match
		}
	} else {
		dst.AiProfileMutationResult = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiProfilesCreate400Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiProfilesCreate400Response) MarshalJSON() ([]byte, error) {
	if src.AiErrorResponse != nil {
		return json.Marshal(&src.AiErrorResponse)
	}

	if src.AiProfileMutationResult != nil {
		return json.Marshal(&src.AiProfileMutationResult)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiProfilesCreate400Response struct {
	value *AiProfilesCreate400Response
	isSet bool
}

func (v NullableAiProfilesCreate400Response) Get() *AiProfilesCreate400Response {
	return v.value
}

func (v *NullableAiProfilesCreate400Response) Set(val *AiProfilesCreate400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesCreate400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesCreate400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesCreate400Response(val *AiProfilesCreate400Response) *NullableAiProfilesCreate400Response {
	return &NullableAiProfilesCreate400Response{value: val, isSet: true}
}

func (v NullableAiProfilesCreate400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesCreate400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


