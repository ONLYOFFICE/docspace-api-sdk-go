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


// AiMetaFill200ResponseFilledInnerValue struct for AiMetaFill200ResponseFilledInnerValue
type AiMetaFill200ResponseFilledInnerValue struct {
	AiMetadataValue *AiMetadataValue
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AiMetaFill200ResponseFilledInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AiMetadataValue
	err = json.Unmarshal(data, &dst.AiMetadataValue);
	if err == nil {
		jsonAiMetadataValue, _ := json.Marshal(dst.AiMetadataValue)
		if string(jsonAiMetadataValue) == "{}" { // empty struct
			dst.AiMetadataValue = nil
		} else {
			return nil // data stored in dst.AiMetadataValue, return on the first match
		}
	} else {
		dst.AiMetadataValue = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AiMetaFill200ResponseFilledInnerValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AiMetaFill200ResponseFilledInnerValue) MarshalJSON() ([]byte, error) {
	if src.AiMetadataValue != nil {
		return json.Marshal(&src.AiMetadataValue)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAiMetaFill200ResponseFilledInnerValue struct {
	value *AiMetaFill200ResponseFilledInnerValue
	isSet bool
}

func (v NullableAiMetaFill200ResponseFilledInnerValue) Get() *AiMetaFill200ResponseFilledInnerValue {
	return v.value
}

func (v *NullableAiMetaFill200ResponseFilledInnerValue) Set(val *AiMetaFill200ResponseFilledInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaFill200ResponseFilledInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaFill200ResponseFilledInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaFill200ResponseFilledInnerValue(val *AiMetaFill200ResponseFilledInnerValue) *NullableAiMetaFill200ResponseFilledInnerValue {
	return &NullableAiMetaFill200ResponseFilledInnerValue{value: val, isSet: true}
}

func (v NullableAiMetaFill200ResponseFilledInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaFill200ResponseFilledInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


