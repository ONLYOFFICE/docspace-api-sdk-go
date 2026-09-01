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


// AiOpenAIStreamChunk A chunk or the terminal error envelope emitted on a failed stream.
type AiOpenAIStreamChunk struct {
	AiOpenAIChatCompletionChunk *AiOpenAIChatCompletionChunk
	AiOpenAIStreamError *AiOpenAIStreamError
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiOpenAIStreamChunk) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AiOpenAIChatCompletionChunk
	err = json.Unmarshal(data, &dst.AiOpenAIChatCompletionChunk);
	if err == nil {
		jsonAiOpenAIChatCompletionChunk, _ := json.Marshal(dst.AiOpenAIChatCompletionChunk)
		if string(jsonAiOpenAIChatCompletionChunk) == "{}" { // empty struct
			dst.AiOpenAIChatCompletionChunk = nil
		} else {
			return nil // data stored in dst.AiOpenAIChatCompletionChunk, return on the first match
		}
	} else {
		dst.AiOpenAIChatCompletionChunk = nil
	}

	// try to unmarshal JSON data into AiOpenAIStreamError
	err = json.Unmarshal(data, &dst.AiOpenAIStreamError);
	if err == nil {
		jsonAiOpenAIStreamError, _ := json.Marshal(dst.AiOpenAIStreamError)
		if string(jsonAiOpenAIStreamError) == "{}" { // empty struct
			dst.AiOpenAIStreamError = nil
		} else {
			return nil // data stored in dst.AiOpenAIStreamError, return on the first match
		}
	} else {
		dst.AiOpenAIStreamError = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AiOpenAIStreamChunk)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiOpenAIStreamChunk) MarshalJSON() ([]byte, error) {
	if src.AiOpenAIChatCompletionChunk != nil {
		return json.Marshal(&src.AiOpenAIChatCompletionChunk)
	}

	if src.AiOpenAIStreamError != nil {
		return json.Marshal(&src.AiOpenAIStreamError)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiOpenAIStreamChunk struct {
	value *AiOpenAIStreamChunk
	isSet bool
}

func (v NullableAiOpenAIStreamChunk) Get() *AiOpenAIStreamChunk {
	return v.value
}

func (v *NullableAiOpenAIStreamChunk) Set(val *AiOpenAIStreamChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIStreamChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIStreamChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIStreamChunk(val *AiOpenAIStreamChunk) *NullableAiOpenAIStreamChunk {
	return &NullableAiOpenAIStreamChunk{value: val, isSet: true}
}

func (v NullableAiOpenAIStreamChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIStreamChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


