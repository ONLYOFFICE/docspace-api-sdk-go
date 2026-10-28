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

// checks if the AiSendStreamBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSendStreamBody{}

// AiSendStreamBody Shared body of the two streaming send endpoints (`sendWithStream` and its OpenAI-framed twin) — the `Chat` action is implied, so there is no `actionType`.
type AiSendStreamBody struct {
	// Target thread; a new one is created (with an auto title) when omitted.
	ThreadId *string `json:"threadId,omitempty"`
	// The user turn to send.
	UserMessage AiThreadMessageLike `json:"userMessage"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
	// Session-level profile override for this request only.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiSendStreamBody AiSendStreamBody

// NewAiSendStreamBody instantiates a new AiSendStreamBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSendStreamBody(userMessage AiThreadMessageLike) *AiSendStreamBody {
	this := AiSendStreamBody{}
	this.UserMessage = userMessage
	return &this
}

// NewAiSendStreamBodyWithDefaults instantiates a new AiSendStreamBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSendStreamBodyWithDefaults() *AiSendStreamBody {
	this := AiSendStreamBody{}
	return &this
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *AiSendStreamBody) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendStreamBody) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *AiSendStreamBody) IsThreadIdSet() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *AiSendStreamBody) SetThreadId(v string) {
	o.ThreadId = &v
}

// GetUserMessage returns the UserMessage field value
func (o *AiSendStreamBody) GetUserMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiSendStreamBody) GetUserMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value
func (o *AiSendStreamBody) SetUserMessage(v AiThreadMessageLike) {
	o.UserMessage = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiSendStreamBody) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendStreamBody) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiSendStreamBody) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiSendStreamBody) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiSendStreamBody) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendStreamBody) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiSendStreamBody) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiSendStreamBody) SetEntityId(v string) {
	o.EntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiSendStreamBody) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendStreamBody) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiSendStreamBody) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiSendStreamBody) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiSendStreamBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSendStreamBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	toSerialize["userMessage"] = o.UserMessage
	if !IsNil(o.ActionArgs) {
		toSerialize["actionArgs"] = o.ActionArgs
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

func (o *AiSendStreamBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiSendStreamBody := _AiSendStreamBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSendStreamBody)

	if err != nil {
		return err
	}

	*o = AiSendStreamBody(varAiSendStreamBody)

	return err
}

type NullableAiSendStreamBody struct {
	value *AiSendStreamBody
	isSet bool
}

func (v NullableAiSendStreamBody) Get() *AiSendStreamBody {
	return v.value
}

func (v *NullableAiSendStreamBody) Set(val *AiSendStreamBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSendStreamBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSendStreamBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSendStreamBody(val *AiSendStreamBody) *NullableAiSendStreamBody {
	return &NullableAiSendStreamBody{value: val, isSet: true}
}

func (v NullableAiSendStreamBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSendStreamBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

