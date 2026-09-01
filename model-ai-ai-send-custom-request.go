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

// checks if the AiAiSendCustomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiSendCustomRequest{}

// AiAiSendCustomRequest struct for AiAiSendCustomRequest
type AiAiSendCustomRequest struct {
	// Stream the reply (ndjson) when true, else return a single message.
	IsStream bool `json:"isStream"`
	// Caller-supplied system prompt for this one-turn call.
	SystemPrompt string `json:"systemPrompt"`
	UserMessage AiThreadMessageLike `json:"userMessage"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiAiActionArgs `json:"actionArgs,omitempty"`
}

type _AiAiSendCustomRequest AiAiSendCustomRequest

// NewAiAiSendCustomRequest instantiates a new AiAiSendCustomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiSendCustomRequest(isStream bool, systemPrompt string, userMessage AiThreadMessageLike) *AiAiSendCustomRequest {
	this := AiAiSendCustomRequest{}
	this.IsStream = isStream
	this.SystemPrompt = systemPrompt
	this.UserMessage = userMessage
	return &this
}

// NewAiAiSendCustomRequestWithDefaults instantiates a new AiAiSendCustomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiSendCustomRequestWithDefaults() *AiAiSendCustomRequest {
	this := AiAiSendCustomRequest{}
	return &this
}

// GetIsStream returns the IsStream field value
func (o *AiAiSendCustomRequest) GetIsStream() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsStream
}

// GetIsStreamOk returns a tuple with the IsStream field value
// and a boolean to check if the value has been set.
func (o *AiAiSendCustomRequest) GetIsStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsStream, true
}

// SetIsStream sets field value
func (o *AiAiSendCustomRequest) SetIsStream(v bool) {
	o.IsStream = v
}

// GetSystemPrompt returns the SystemPrompt field value
func (o *AiAiSendCustomRequest) GetSystemPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemPrompt
}

// GetSystemPromptOk returns a tuple with the SystemPrompt field value
// and a boolean to check if the value has been set.
func (o *AiAiSendCustomRequest) GetSystemPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemPrompt, true
}

// SetSystemPrompt sets field value
func (o *AiAiSendCustomRequest) SetSystemPrompt(v string) {
	o.SystemPrompt = v
}

// GetUserMessage returns the UserMessage field value
func (o *AiAiSendCustomRequest) GetUserMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiAiSendCustomRequest) GetUserMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value
func (o *AiAiSendCustomRequest) SetUserMessage(v AiThreadMessageLike) {
	o.UserMessage = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiAiSendCustomRequest) GetActionArgs() AiAiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiAiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSendCustomRequest) GetActionArgsOk() (*AiAiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiAiSendCustomRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiAiActionArgs and assigns it to the ActionArgs field.
func (o *AiAiSendCustomRequest) SetActionArgs(v AiAiActionArgs) {
	o.ActionArgs = &v
}

func (o AiAiSendCustomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiSendCustomRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isStream"] = o.IsStream
	toSerialize["systemPrompt"] = o.SystemPrompt
	toSerialize["userMessage"] = o.UserMessage
	if !IsNil(o.ActionArgs) {
		toSerialize["actionArgs"] = o.ActionArgs
	}
	return toSerialize, nil
}

func (o *AiAiSendCustomRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAiAiSendCustomRequest := _AiAiSendCustomRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAiSendCustomRequest)

	if err != nil {
		return err
	}

	*o = AiAiSendCustomRequest(varAiAiSendCustomRequest)

	return err
}

type NullableAiAiSendCustomRequest struct {
	value *AiAiSendCustomRequest
	isSet bool
}

func (v NullableAiAiSendCustomRequest) Get() *AiAiSendCustomRequest {
	return v.value
}

func (v *NullableAiAiSendCustomRequest) Set(val *AiAiSendCustomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiSendCustomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiSendCustomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiSendCustomRequest(val *AiAiSendCustomRequest) *NullableAiAiSendCustomRequest {
	return &NullableAiAiSendCustomRequest{value: val, isSet: true}
}

func (v NullableAiAiSendCustomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiSendCustomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

