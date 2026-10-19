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

// checks if the AiSendCustomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSendCustomRequest{}

// AiSendCustomRequest struct for AiSendCustomRequest
type AiSendCustomRequest struct {
	// Stream the reply (ndjson) when true, else return a single message.
	IsStream bool `json:"isStream"`
	// Caller-supplied system prompt for this one-turn call.
	SystemPrompt string `json:"systemPrompt"`
	UserMessage AiThreadMessageLike `json:"userMessage"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
}

type _AiSendCustomRequest AiSendCustomRequest

// NewAiSendCustomRequest instantiates a new AiSendCustomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSendCustomRequest(isStream bool, systemPrompt string, userMessage AiThreadMessageLike) *AiSendCustomRequest {
	this := AiSendCustomRequest{}
	this.IsStream = isStream
	this.SystemPrompt = systemPrompt
	this.UserMessage = userMessage
	return &this
}

// NewAiSendCustomRequestWithDefaults instantiates a new AiSendCustomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSendCustomRequestWithDefaults() *AiSendCustomRequest {
	this := AiSendCustomRequest{}
	return &this
}

// GetIsStream returns the IsStream field value
func (o *AiSendCustomRequest) GetIsStream() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsStream
}

// GetIsStreamOk returns a tuple with the IsStream field value
// and a boolean to check if the value has been set.
func (o *AiSendCustomRequest) GetIsStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsStream, true
}

// SetIsStream sets field value
func (o *AiSendCustomRequest) SetIsStream(v bool) {
	o.IsStream = v
}

// GetSystemPrompt returns the SystemPrompt field value
func (o *AiSendCustomRequest) GetSystemPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemPrompt
}

// GetSystemPromptOk returns a tuple with the SystemPrompt field value
// and a boolean to check if the value has been set.
func (o *AiSendCustomRequest) GetSystemPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemPrompt, true
}

// SetSystemPrompt sets field value
func (o *AiSendCustomRequest) SetSystemPrompt(v string) {
	o.SystemPrompt = v
}

// GetUserMessage returns the UserMessage field value
func (o *AiSendCustomRequest) GetUserMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiSendCustomRequest) GetUserMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value
func (o *AiSendCustomRequest) SetUserMessage(v AiThreadMessageLike) {
	o.UserMessage = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiSendCustomRequest) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendCustomRequest) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiSendCustomRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiSendCustomRequest) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

func (o AiSendCustomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSendCustomRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isStream"] = o.IsStream
	toSerialize["systemPrompt"] = o.SystemPrompt
	toSerialize["userMessage"] = o.UserMessage
	if !IsNil(o.ActionArgs) {
		toSerialize["actionArgs"] = o.ActionArgs
	}
	return toSerialize, nil
}

func (o *AiSendCustomRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isStream",
		"systemPrompt",
		"userMessage",
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

	varAiSendCustomRequest := _AiSendCustomRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSendCustomRequest)

	if err != nil {
		return err
	}

	*o = AiSendCustomRequest(varAiSendCustomRequest)

	return err
}

type NullableAiSendCustomRequest struct {
	value *AiSendCustomRequest
	isSet bool
}

func (v NullableAiSendCustomRequest) Get() *AiSendCustomRequest {
	return v.value
}

func (v *NullableAiSendCustomRequest) Set(val *AiSendCustomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSendCustomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSendCustomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSendCustomRequest(val *AiSendCustomRequest) *NullableAiSendCustomRequest {
	return &NullableAiSendCustomRequest{value: val, isSet: true}
}

func (v NullableAiSendCustomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSendCustomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

