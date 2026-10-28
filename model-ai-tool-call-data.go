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

// checks if the AiToolCallData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolCallData{}

// AiToolCallData Identifies a pending tool call to resume — mirrors the library `ToolCallData` (its serializable fields).
type AiToolCallData struct {
	// Thread the assistant message belongs to.
	ThreadId string `json:"threadId"`
	// Storage id of the assistant message holding the tool call.
	MessageId string `json:"messageId"`
	// Index of the tool-call content part inside `message.content`.
	Idx float32 `json:"idx"`
	// Snapshot of the assistant message at the time the tool call surfaced.
	Message AiThreadMessageLike `json:"message"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
	// Session-level profile override for this request only.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiToolCallData AiToolCallData

// NewAiToolCallData instantiates a new AiToolCallData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolCallData(threadId string, messageId string, idx float32, message AiThreadMessageLike) *AiToolCallData {
	this := AiToolCallData{}
	this.ThreadId = threadId
	this.MessageId = messageId
	this.Idx = idx
	this.Message = message
	return &this
}

// NewAiToolCallDataWithDefaults instantiates a new AiToolCallData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolCallDataWithDefaults() *AiToolCallData {
	this := AiToolCallData{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiToolCallData) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiToolCallData) SetThreadId(v string) {
	o.ThreadId = v
}

// GetMessageId returns the MessageId field value
func (o *AiToolCallData) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiToolCallData) SetMessageId(v string) {
	o.MessageId = v
}

// GetIdx returns the Idx field value
func (o *AiToolCallData) GetIdx() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Idx
}

// GetIdxOk returns a tuple with the Idx field value
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetIdxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idx, true
}

// SetIdx sets field value
func (o *AiToolCallData) SetIdx(v float32) {
	o.Idx = v
}

// GetMessage returns the Message field value
func (o *AiToolCallData) GetMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiToolCallData) SetMessage(v AiThreadMessageLike) {
	o.Message = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiToolCallData) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiToolCallData) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiToolCallData) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolCallData) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolCallData) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolCallData) SetEntityId(v string) {
	o.EntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiToolCallData) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolCallData) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiToolCallData) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiToolCallData) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiToolCallData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolCallData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["messageId"] = o.MessageId
	toSerialize["idx"] = o.Idx
	toSerialize["message"] = o.Message
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

func (o *AiToolCallData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
		"messageId",
		"idx",
		"message",
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

	varAiToolCallData := _AiToolCallData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolCallData)

	if err != nil {
		return err
	}

	*o = AiToolCallData(varAiToolCallData)

	return err
}

type NullableAiToolCallData struct {
	value *AiToolCallData
	isSet bool
}

func (v NullableAiToolCallData) Get() *AiToolCallData {
	return v.value
}

func (v *NullableAiToolCallData) Set(val *AiToolCallData) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolCallData) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolCallData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolCallData(val *AiToolCallData) *NullableAiToolCallData {
	return &NullableAiToolCallData{value: val, isSet: true}
}

func (v NullableAiToolCallData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolCallData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

