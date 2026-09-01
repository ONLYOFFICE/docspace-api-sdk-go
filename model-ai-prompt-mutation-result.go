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

// checks if the AiPromptMutationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptMutationResult{}

// AiPromptMutationResult Outcome of `create` / `update` / `move` on a prompt — either the persisted prompt or a field-scoped error.
type AiPromptMutationResult struct {
	// True when the prompt was persisted.
	Success bool `json:"success"`
	// The persisted prompt. Present on success.
	Prompt *AiPrompt `json:"prompt,omitempty"`
	// Why the prompt was rejected. Present on failure.
	Error *AiTErrorData `json:"error,omitempty"`
}

type _AiPromptMutationResult AiPromptMutationResult

// NewAiPromptMutationResult instantiates a new AiPromptMutationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptMutationResult(success bool) *AiPromptMutationResult {
	this := AiPromptMutationResult{}
	this.Success = success
	return &this
}

// NewAiPromptMutationResultWithDefaults instantiates a new AiPromptMutationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptMutationResultWithDefaults() *AiPromptMutationResult {
	this := AiPromptMutationResult{}
	return &this
}

// GetSuccess returns the Success field value
func (o *AiPromptMutationResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *AiPromptMutationResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *AiPromptMutationResult) SetSuccess(v bool) {
	o.Success = v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiPromptMutationResult) GetPrompt() AiPrompt {
	if o == nil || IsNil(o.Prompt) {
		var ret AiPrompt
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptMutationResult) GetPromptOk() (*AiPrompt, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiPromptMutationResult) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given AiPrompt and assigns it to the Prompt field.
func (o *AiPromptMutationResult) SetPrompt(v AiPrompt) {
	o.Prompt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AiPromptMutationResult) GetError() AiTErrorData {
	if o == nil || IsNil(o.Error) {
		var ret AiTErrorData
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptMutationResult) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AiPromptMutationResult) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AiTErrorData and assigns it to the Error field.
func (o *AiPromptMutationResult) SetError(v AiTErrorData) {
	o.Error = &v
}

func (o AiPromptMutationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptMutationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *AiPromptMutationResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
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

	varAiPromptMutationResult := _AiPromptMutationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptMutationResult)

	if err != nil {
		return err
	}

	*o = AiPromptMutationResult(varAiPromptMutationResult)

	return err
}

type NullableAiPromptMutationResult struct {
	value *AiPromptMutationResult
	isSet bool
}

func (v NullableAiPromptMutationResult) Get() *AiPromptMutationResult {
	return v.value
}

func (v *NullableAiPromptMutationResult) Set(val *AiPromptMutationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptMutationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptMutationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptMutationResult(val *AiPromptMutationResult) *NullableAiPromptMutationResult {
	return &NullableAiPromptMutationResult{value: val, isSet: true}
}

func (v NullableAiPromptMutationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptMutationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

