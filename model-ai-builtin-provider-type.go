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

// AiBuiltinProviderType Union of all 17 built-in AI provider type identifiers.  The `external` provider has no built-in transport — it delegates every HTTP request to `PlatformAdapter.externalFetch` and parses the response with the inner provider selected by `Profile.basedOn`.
type AiBuiltinProviderType string

// List of AiBuiltinProviderType
const (
	AIBUILTINPROVIDERTYPE_ANTHROPIC AiBuiltinProviderType = "anthropic"
	AIBUILTINPROVIDERTYPE_OLLAMA AiBuiltinProviderType = "ollama"
	AIBUILTINPROVIDERTYPE_OPENAI AiBuiltinProviderType = "openai"
	AIBUILTINPROVIDERTYPE_OPENAICOMPATIBLE AiBuiltinProviderType = "openaicompatible"
	AIBUILTINPROVIDERTYPE_TOGETHER AiBuiltinProviderType = "together"
	AIBUILTINPROVIDERTYPE_OPENROUTER AiBuiltinProviderType = "openrouter"
	AIBUILTINPROVIDERTYPE_GENAI AiBuiltinProviderType = "genai"
	AIBUILTINPROVIDERTYPE_DEEPSEEK AiBuiltinProviderType = "deepseek"
	AIBUILTINPROVIDERTYPE_XAI AiBuiltinProviderType = "xai"
	AIBUILTINPROVIDERTYPE_LM_STUDIO AiBuiltinProviderType = "lm-studio"
	AIBUILTINPROVIDERTYPE_MISTRAL AiBuiltinProviderType = "mistral"
	AIBUILTINPROVIDERTYPE_GROQ AiBuiltinProviderType = "groq"
	AIBUILTINPROVIDERTYPE_ZHIPU AiBuiltinProviderType = "zhipu"
	AIBUILTINPROVIDERTYPE_STABILITYAI AiBuiltinProviderType = "stabilityai"
	AIBUILTINPROVIDERTYPE_GPT4ALL AiBuiltinProviderType = "gpt4all"
	AIBUILTINPROVIDERTYPE_ONLYOFFICE AiBuiltinProviderType = "onlyoffice"
	AIBUILTINPROVIDERTYPE_EXTERNAL AiBuiltinProviderType = "external"
)

// All allowed values of AiBuiltinProviderType enum
var AllowedAiBuiltinProviderTypeEnumValues = []AiBuiltinProviderType{
	"anthropic",
	"ollama",
	"openai",
	"openaicompatible",
	"together",
	"openrouter",
	"genai",
	"deepseek",
	"xai",
	"lm-studio",
	"mistral",
	"groq",
	"zhipu",
	"stabilityai",
	"gpt4all",
	"onlyoffice",
	"external",
}

func (v *AiBuiltinProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiBuiltinProviderType(value)
	for _, existing := range AllowedAiBuiltinProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiBuiltinProviderType", value)
}

// NewAiBuiltinProviderTypeFromValue returns a pointer to a valid AiBuiltinProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiBuiltinProviderTypeFromValue(v string) (*AiBuiltinProviderType, error) {
	ev := AiBuiltinProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiBuiltinProviderType: valid values are %v", v, AllowedAiBuiltinProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiBuiltinProviderType) IsValid() bool {
	for _, existing := range AllowedAiBuiltinProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiBuiltinProviderType value
func (v AiBuiltinProviderType) Ptr() *AiBuiltinProviderType {
	return &v
}

type NullableAiBuiltinProviderType struct {
	value *AiBuiltinProviderType
	isSet bool
}

func (v NullableAiBuiltinProviderType) Get() *AiBuiltinProviderType {
	return v.value
}

func (v *NullableAiBuiltinProviderType) Set(val *AiBuiltinProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiBuiltinProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiBuiltinProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiBuiltinProviderType(val *AiBuiltinProviderType) *NullableAiBuiltinProviderType {
	return &NullableAiBuiltinProviderType{value: val, isSet: true}
}

func (v NullableAiBuiltinProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiBuiltinProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

