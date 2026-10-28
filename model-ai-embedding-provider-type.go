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

// AiEmbeddingProviderType [0 - None, 1 - OpenAi, 2 - OpenRouter, 3 - PortalAi]
type AiEmbeddingProviderType int32

// List of AiEmbeddingProviderType
const (
	AIEMBEDDINGPROVIDERTYPE_None AiEmbeddingProviderType = 0
	AIEMBEDDINGPROVIDERTYPE_OpenAi AiEmbeddingProviderType = 1
	AIEMBEDDINGPROVIDERTYPE_OpenRouter AiEmbeddingProviderType = 2
	AIEMBEDDINGPROVIDERTYPE_PortalAi AiEmbeddingProviderType = 3
)

// All allowed values of AiEmbeddingProviderType enum
var AllowedAiEmbeddingProviderTypeEnumValues = []AiEmbeddingProviderType{
	0,
	1,
	2,
	3,
}

func (v *AiEmbeddingProviderType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiEmbeddingProviderType(value)
	for _, existing := range AllowedAiEmbeddingProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiEmbeddingProviderType", value)
}

// NewAiEmbeddingProviderTypeFromValue returns a pointer to a valid AiEmbeddingProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiEmbeddingProviderTypeFromValue(v int32) (*AiEmbeddingProviderType, error) {
	ev := AiEmbeddingProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiEmbeddingProviderType: valid values are %v", v, AllowedAiEmbeddingProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiEmbeddingProviderType) IsValid() bool {
	for _, existing := range AllowedAiEmbeddingProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiEmbeddingProviderType value
func (v AiEmbeddingProviderType) Ptr() *AiEmbeddingProviderType {
	return &v
}

type NullableAiEmbeddingProviderType struct {
	value *AiEmbeddingProviderType
	isSet bool
}

func (v NullableAiEmbeddingProviderType) Get() *AiEmbeddingProviderType {
	return v.value
}

func (v *NullableAiEmbeddingProviderType) Set(val *AiEmbeddingProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEmbeddingProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEmbeddingProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEmbeddingProviderType(val *AiEmbeddingProviderType) *NullableAiEmbeddingProviderType {
	return &NullableAiEmbeddingProviderType{value: val, isSet: true}
}

func (v NullableAiEmbeddingProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEmbeddingProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

