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


// AiProfilesListProviderModels400Response struct for AiProfilesListProviderModels400Response
type AiProfilesListProviderModels400Response struct {
	AiErrorResponse *AiErrorResponse
	AiProfilesListProviderModels400ResponseAnyOf *AiProfilesListProviderModels400ResponseAnyOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiProfilesListProviderModels400Response) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into AiProfilesListProviderModels400ResponseAnyOf
	err = json.Unmarshal(data, &dst.AiProfilesListProviderModels400ResponseAnyOf);
	if err == nil {
		jsonAiProfilesListProviderModels400ResponseAnyOf, _ := json.Marshal(dst.AiProfilesListProviderModels400ResponseAnyOf)
		if string(jsonAiProfilesListProviderModels400ResponseAnyOf) == "{}" { // empty struct
			dst.AiProfilesListProviderModels400ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.AiProfilesListProviderModels400ResponseAnyOf, return on the first match
		}
	} else {
		dst.AiProfilesListProviderModels400ResponseAnyOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiProfilesListProviderModels400Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiProfilesListProviderModels400Response) MarshalJSON() ([]byte, error) {
	if src.AiErrorResponse != nil {
		return json.Marshal(&src.AiErrorResponse)
	}

	if src.AiProfilesListProviderModels400ResponseAnyOf != nil {
		return json.Marshal(&src.AiProfilesListProviderModels400ResponseAnyOf)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiProfilesListProviderModels400Response struct {
	value *AiProfilesListProviderModels400Response
	isSet bool
}

func (v NullableAiProfilesListProviderModels400Response) Get() *AiProfilesListProviderModels400Response {
	return v.value
}

func (v *NullableAiProfilesListProviderModels400Response) Set(val *AiProfilesListProviderModels400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesListProviderModels400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesListProviderModels400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesListProviderModels400Response(val *AiProfilesListProviderModels400Response) *NullableAiProfilesListProviderModels400Response {
	return &NullableAiProfilesListProviderModels400Response{value: val, isSet: true}
}

func (v NullableAiProfilesListProviderModels400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesListProviderModels400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


