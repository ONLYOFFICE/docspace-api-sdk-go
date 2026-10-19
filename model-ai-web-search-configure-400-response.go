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


// AiWebSearchConfigure400Response struct for AiWebSearchConfigure400Response
type AiWebSearchConfigure400Response struct {
	AiErrorResponse *AiErrorResponse
	AiWebSearchMutationResult *AiWebSearchMutationResult
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiWebSearchConfigure400Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into AiWebSearchMutationResult
	err = json.Unmarshal(data, &dst.AiWebSearchMutationResult);
	if err == nil {
		jsonAiWebSearchMutationResult, _ := json.Marshal(dst.AiWebSearchMutationResult)
		if string(jsonAiWebSearchMutationResult) == "{}" { // empty struct
			dst.AiWebSearchMutationResult = nil
		} else {
			return nil // data stored in dst.AiWebSearchMutationResult, return on the first match
		}
	} else {
		dst.AiWebSearchMutationResult = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiWebSearchConfigure400Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiWebSearchConfigure400Response) MarshalJSON() ([]byte, error) {
	if src.AiErrorResponse != nil {
		return json.Marshal(&src.AiErrorResponse)
	}

	if src.AiWebSearchMutationResult != nil {
		return json.Marshal(&src.AiWebSearchMutationResult)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiWebSearchConfigure400Response struct {
	value *AiWebSearchConfigure400Response
	isSet bool
}

func (v NullableAiWebSearchConfigure400Response) Get() *AiWebSearchConfigure400Response {
	return v.value
}

func (v *NullableAiWebSearchConfigure400Response) Set(val *AiWebSearchConfigure400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWebSearchConfigure400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWebSearchConfigure400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWebSearchConfigure400Response(val *AiWebSearchConfigure400Response) *NullableAiWebSearchConfigure400Response {
	return &NullableAiWebSearchConfigure400Response{value: val, isSet: true}
}

func (v NullableAiWebSearchConfigure400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWebSearchConfigure400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


