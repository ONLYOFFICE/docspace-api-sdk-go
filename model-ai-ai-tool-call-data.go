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

// checks if the AiAiToolCallData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiToolCallData{}

// AiAiToolCallData Identifies a pending tool call to resume — mirrors the library `ToolCallData` (its serializable fields).
type AiAiToolCallData struct {
	// Thread the assistant message belongs to.
	ThreadId string `json:"threadId"`
	// Storage id of the assistant message holding the tool call.
	MessageId string `json:"messageId"`
	// Index of the tool-call content part inside `message.content`.
	Idx float32 `json:"idx"`
	// Snapshot of the assistant message at the time the tool call surfaced.
	Message AiThreadMessageLike `json:"message"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiAiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
	// Session-level profile override for this request only.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiAiToolCallData AiAiToolCallData

// NewAiAiToolCallData instantiates a new AiAiToolCallData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiToolCallData(threadId string, messageId string, idx float32, message AiThreadMessageLike) *AiAiToolCallData {
	this := AiAiToolCallData{}
	this.ThreadId = threadId
	this.MessageId = messageId
	this.Idx = idx
	this.Message = message
	return &this
}

// NewAiAiToolCallDataWithDefaults instantiates a new AiAiToolCallData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiToolCallDataWithDefaults() *AiAiToolCallData {
	this := AiAiToolCallData{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiAiToolCallData) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiAiToolCallData) SetThreadId(v string) {
	o.ThreadId = v
}

// GetMessageId returns the MessageId field value
func (o *AiAiToolCallData) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiAiToolCallData) SetMessageId(v string) {
	o.MessageId = v
}

// GetIdx returns the Idx field value
func (o *AiAiToolCallData) GetIdx() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Idx
}

// GetIdxOk returns a tuple with the Idx field value
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetIdxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idx, true
}

// SetIdx sets field value
func (o *AiAiToolCallData) SetIdx(v float32) {
	o.Idx = v
}

// GetMessage returns the Message field value
func (o *AiAiToolCallData) GetMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiAiToolCallData) SetMessage(v AiThreadMessageLike) {
	o.Message = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiAiToolCallData) GetActionArgs() AiAiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiAiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetActionArgsOk() (*AiAiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiAiToolCallData) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiAiActionArgs and assigns it to the ActionArgs field.
func (o *AiAiToolCallData) SetActionArgs(v AiAiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiAiToolCallData) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiAiToolCallData) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiAiToolCallData) SetEntityId(v string) {
	o.EntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiAiToolCallData) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiToolCallData) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiAiToolCallData) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiAiToolCallData) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiAiToolCallData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiToolCallData) ToMap() (map[string]interface{}, error) {
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

func (o *AiAiToolCallData) UnmarshalJSON(data []byte) (err error) {
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

	varAiAiToolCallData := _AiAiToolCallData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAiToolCallData)

	if err != nil {
		return err
	}

	*o = AiAiToolCallData(varAiAiToolCallData)

	return err
}

type NullableAiAiToolCallData struct {
	value *AiAiToolCallData
	isSet bool
}

func (v NullableAiAiToolCallData) Get() *AiAiToolCallData {
	return v.value
}

func (v *NullableAiAiToolCallData) Set(val *AiAiToolCallData) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiToolCallData) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiToolCallData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiToolCallData(val *AiAiToolCallData) *NullableAiAiToolCallData {
	return &NullableAiAiToolCallData{value: val, isSet: true}
}

func (v NullableAiAiToolCallData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiToolCallData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

