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


// AiProviderType Provider type identifier. Accepts all 17 built-in types with autocomplete, plus any custom `string` for dynamically registered providers.
type AiProviderType struct {
	AiBuiltinProviderType *AiBuiltinProviderType
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiProviderType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AiBuiltinProviderType
	err = json.Unmarshal(data, &dst.AiBuiltinProviderType);
	if err == nil {
		jsonAiBuiltinProviderType, _ := json.Marshal(dst.AiBuiltinProviderType)
		if string(jsonAiBuiltinProviderType) == "{}" { // empty struct
			dst.AiBuiltinProviderType = nil
		} else {
			return nil // data stored in dst.AiBuiltinProviderType, return on the first match
		}
	} else {
		dst.AiBuiltinProviderType = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AiProviderType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiProviderType) MarshalJSON() ([]byte, error) {
	if src.AiBuiltinProviderType != nil {
		return json.Marshal(&src.AiBuiltinProviderType)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiProviderType struct {
	value *AiProviderType
	isSet bool
}

func (v NullableAiProviderType) Get() *AiProviderType {
	return v.value
}

func (v *NullableAiProviderType) Set(val *AiProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProviderType(val *AiProviderType) *NullableAiProviderType {
	return &NullableAiProviderType{value: val, isSet: true}
}

func (v NullableAiProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


