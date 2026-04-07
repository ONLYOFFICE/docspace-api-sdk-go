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

// ProviderType [0 - PortalAi, 1 - OpenAi, 2 - TogetherAi, 3 - OpenAiCompatible, 4 - Anthropic, 5 - OpenRouter, 6 - DeepSeek, 7 - XAi, 8 - GoogleAi]
type ProviderType int32

// List of ProviderType
const (
	PROVIDERTYPE_PortalAi ProviderType = 0
	PROVIDERTYPE_OpenAi ProviderType = 1
	PROVIDERTYPE_TogetherAi ProviderType = 2
	PROVIDERTYPE_OpenAiCompatible ProviderType = 3
	PROVIDERTYPE_Anthropic ProviderType = 4
	PROVIDERTYPE_OpenRouter ProviderType = 5
	PROVIDERTYPE_DeepSeek ProviderType = 6
	PROVIDERTYPE_XAi ProviderType = 7
	PROVIDERTYPE_GoogleAi ProviderType = 8
)

// All allowed values of ProviderType enum
var AllowedProviderTypeEnumValues = []ProviderType{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

func (v *ProviderType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderType(value)
	for _, existing := range AllowedProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderType", value)
}

// NewProviderTypeFromValue returns a pointer to a valid ProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderTypeFromValue(v int32) (*ProviderType, error) {
	ev := ProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderType: valid values are %v", v, AllowedProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderType) IsValid() bool {
	for _, existing := range AllowedProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderType value
func (v ProviderType) Ptr() *ProviderType {
	return &v
}

type NullableProviderType struct {
	value *ProviderType
	isSet bool
}

func (v NullableProviderType) Get() *ProviderType {
	return v.value
}

func (v *NullableProviderType) Set(val *ProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderType(val *ProviderType) *NullableProviderType {
	return &NullableProviderType{value: val, isSet: true}
}

func (v NullableProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

