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
	"bytes"
	"fmt"
)

// checks if the AiSuggestedQuestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSuggestedQuestion{}

// AiSuggestedQuestion A starter question proposed for an attached form.
type AiSuggestedQuestion struct {
	// The question as it is shown to the user.
	Question string `json:"question"`
	// The prompt sent to the chat when the user picks the question.
	Prompt string `json:"prompt"`
}

type _AiSuggestedQuestion AiSuggestedQuestion

// NewAiSuggestedQuestion instantiates a new AiSuggestedQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSuggestedQuestion(question string, prompt string) *AiSuggestedQuestion {
	this := AiSuggestedQuestion{}
	this.Question = question
	this.Prompt = prompt
	return &this
}

// NewAiSuggestedQuestionWithDefaults instantiates a new AiSuggestedQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSuggestedQuestionWithDefaults() *AiSuggestedQuestion {
	this := AiSuggestedQuestion{}
	return &this
}

// GetQuestion returns the Question field value
func (o *AiSuggestedQuestion) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *AiSuggestedQuestion) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *AiSuggestedQuestion) SetQuestion(v string) {
	o.Question = v
}

// GetPrompt returns the Prompt field value
func (o *AiSuggestedQuestion) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *AiSuggestedQuestion) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *AiSuggestedQuestion) SetPrompt(v string) {
	o.Prompt = v
}

func (o AiSuggestedQuestion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSuggestedQuestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["question"] = o.Question
	toSerialize["prompt"] = o.Prompt
	return toSerialize, nil
}

func (o *AiSuggestedQuestion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"question",
		"prompt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAiSuggestedQuestion := _AiSuggestedQuestion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSuggestedQuestion)

	if err != nil {
		return err
	}

	*o = AiSuggestedQuestion(varAiSuggestedQuestion)

	return err
}

type NullableAiSuggestedQuestion struct {
	value *AiSuggestedQuestion
	isSet bool
}

func (v NullableAiSuggestedQuestion) Get() *AiSuggestedQuestion {
	return v.value
}

func (v *NullableAiSuggestedQuestion) Set(val *AiSuggestedQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSuggestedQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSuggestedQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSuggestedQuestion(val *AiSuggestedQuestion) *NullableAiSuggestedQuestion {
	return &NullableAiSuggestedQuestion{value: val, isSet: true}
}

func (v NullableAiSuggestedQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSuggestedQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

