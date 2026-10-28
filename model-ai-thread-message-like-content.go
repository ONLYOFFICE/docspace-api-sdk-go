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


// AiThreadMessageLikeContent Message content: either plain text or a list of typed content parts (text, image, tool-call, …). Parts are open-ended by content type.
type AiThreadMessageLikeContent struct {
	ArrayOfAiThreadMessageLikeContentAnyOfInner *[]AiThreadMessageLikeContentAnyOfInner
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiThreadMessageLikeContent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfAiThreadMessageLikeContentAnyOfInner
	err = json.Unmarshal(data, &dst.ArrayOfAiThreadMessageLikeContentAnyOfInner);
	if err == nil {
		jsonArrayOfAiThreadMessageLikeContentAnyOfInner, _ := json.Marshal(dst.ArrayOfAiThreadMessageLikeContentAnyOfInner)
		if string(jsonArrayOfAiThreadMessageLikeContentAnyOfInner) == "{}" { // empty struct
			dst.ArrayOfAiThreadMessageLikeContentAnyOfInner = nil
		} else {
			return nil // data stored in dst.ArrayOfAiThreadMessageLikeContentAnyOfInner, return on the first match
		}
	} else {
		dst.ArrayOfAiThreadMessageLikeContentAnyOfInner = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AiThreadMessageLikeContent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiThreadMessageLikeContent) MarshalJSON() ([]byte, error) {
	if src.ArrayOfAiThreadMessageLikeContentAnyOfInner != nil {
		return json.Marshal(&src.ArrayOfAiThreadMessageLikeContentAnyOfInner)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiThreadMessageLikeContent struct {
	value *AiThreadMessageLikeContent
	isSet bool
}

func (v NullableAiThreadMessageLikeContent) Get() *AiThreadMessageLikeContent {
	return v.value
}

func (v *NullableAiThreadMessageLikeContent) Set(val *AiThreadMessageLikeContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeContent(val *AiThreadMessageLikeContent) *NullableAiThreadMessageLikeContent {
	return &NullableAiThreadMessageLikeContent{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


