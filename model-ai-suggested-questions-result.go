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

// checks if the AiSuggestedQuestionsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSuggestedQuestionsResult{}

// AiSuggestedQuestionsResult The starter questions of an attached form, and whether they are ready.
type AiSuggestedQuestionsResult struct {
	// `ready` - the questions are there; `pending` - they are being generated and arrive over the socket, ask again later; `unavailable` - the attachment is not a form that can have any.
	Status string `json:"status"`
	// The questions; empty unless `status` is `ready`.
	Questions []AiSuggestedQuestion `json:"questions"`
}

type _AiSuggestedQuestionsResult AiSuggestedQuestionsResult

// NewAiSuggestedQuestionsResult instantiates a new AiSuggestedQuestionsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSuggestedQuestionsResult(status string, questions []AiSuggestedQuestion) *AiSuggestedQuestionsResult {
	this := AiSuggestedQuestionsResult{}
	this.Status = status
	this.Questions = questions
	return &this
}

// NewAiSuggestedQuestionsResultWithDefaults instantiates a new AiSuggestedQuestionsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSuggestedQuestionsResultWithDefaults() *AiSuggestedQuestionsResult {
	this := AiSuggestedQuestionsResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *AiSuggestedQuestionsResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiSuggestedQuestionsResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AiSuggestedQuestionsResult) SetStatus(v string) {
	o.Status = v
}

// GetQuestions returns the Questions field value
func (o *AiSuggestedQuestionsResult) GetQuestions() []AiSuggestedQuestion {
	if o == nil {
		var ret []AiSuggestedQuestion
		return ret
	}

	return o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value
// and a boolean to check if the value has been set.
func (o *AiSuggestedQuestionsResult) GetQuestionsOk() ([]AiSuggestedQuestion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Questions, true
}

// SetQuestions sets field value
func (o *AiSuggestedQuestionsResult) SetQuestions(v []AiSuggestedQuestion) {
	o.Questions = v
}

func (o AiSuggestedQuestionsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSuggestedQuestionsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["questions"] = o.Questions
	return toSerialize, nil
}

func (o *AiSuggestedQuestionsResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"questions",
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

	varAiSuggestedQuestionsResult := _AiSuggestedQuestionsResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSuggestedQuestionsResult)

	if err != nil {
		return err
	}

	*o = AiSuggestedQuestionsResult(varAiSuggestedQuestionsResult)

	return err
}

type NullableAiSuggestedQuestionsResult struct {
	value *AiSuggestedQuestionsResult
	isSet bool
}

func (v NullableAiSuggestedQuestionsResult) Get() *AiSuggestedQuestionsResult {
	return v.value
}

func (v *NullableAiSuggestedQuestionsResult) Set(val *AiSuggestedQuestionsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSuggestedQuestionsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSuggestedQuestionsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSuggestedQuestionsResult(val *AiSuggestedQuestionsResult) *NullableAiSuggestedQuestionsResult {
	return &NullableAiSuggestedQuestionsResult{value: val, isSet: true}
}

func (v NullableAiSuggestedQuestionsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSuggestedQuestionsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

