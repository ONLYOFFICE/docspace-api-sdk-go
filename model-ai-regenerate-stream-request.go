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

// checks if the AiRegenerateStreamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiRegenerateStreamRequest{}

// AiRegenerateStreamRequest struct for AiRegenerateStreamRequest
type AiRegenerateStreamRequest struct {
	// Target thread (must already exist).
	ThreadId string `json:"threadId"`
	// The user message to re-roll from: everything stored after it is deleted. When omitted, or no longer in the thread, the last user message is used.
	MessageId *string `json:"messageId,omitempty"`
	// How many messages from the start of the thread to keep, used when `messageId` does not resolve. The message at `keepCount - 1` has to be a user message; otherwise the value is ignored.
	KeepCount *int32 `json:"keepCount,omitempty"`
	// The thread length the caller saw when it computed `keepCount`. When it differs from the stored length, `keepCount` is discarded and the last user message is used instead.
	ExpectedLength *int32 `json:"expectedLength,omitempty"`
	// Per-request options of the round: extra tools, reasoning, a prompt override and the connected context.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
	// The folder or room the user is in. Tools run against it and the model is told it is the user's location; falls back to `entityId` when omitted.
	ContextEntityId *string `json:"contextEntityId,omitempty"`
	// Session-level profile override for this request only.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiRegenerateStreamRequest AiRegenerateStreamRequest

// NewAiRegenerateStreamRequest instantiates a new AiRegenerateStreamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiRegenerateStreamRequest(threadId string) *AiRegenerateStreamRequest {
	this := AiRegenerateStreamRequest{}
	this.ThreadId = threadId
	return &this
}

// NewAiRegenerateStreamRequestWithDefaults instantiates a new AiRegenerateStreamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiRegenerateStreamRequestWithDefaults() *AiRegenerateStreamRequest {
	this := AiRegenerateStreamRequest{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiRegenerateStreamRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiRegenerateStreamRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsMessageIdSet() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *AiRegenerateStreamRequest) SetMessageId(v string) {
	o.MessageId = &v
}

// GetKeepCount returns the KeepCount field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetKeepCount() int32 {
	if o == nil || IsNil(o.KeepCount) {
		var ret int32
		return ret
	}
	return *o.KeepCount
}

// GetKeepCountOk returns a tuple with the KeepCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetKeepCountOk() (*int32, bool) {
	if o == nil || IsNil(o.KeepCount) {
		return nil, false
	}
	return o.KeepCount, true
}

// HasKeepCount returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsKeepCountSet() bool {
	if o != nil && !IsNil(o.KeepCount) {
		return true
	}

	return false
}

// SetKeepCount gets a reference to the given int32 and assigns it to the KeepCount field.
func (o *AiRegenerateStreamRequest) SetKeepCount(v int32) {
	o.KeepCount = &v
}

// GetExpectedLength returns the ExpectedLength field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetExpectedLength() int32 {
	if o == nil || IsNil(o.ExpectedLength) {
		var ret int32
		return ret
	}
	return *o.ExpectedLength
}

// GetExpectedLengthOk returns a tuple with the ExpectedLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetExpectedLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpectedLength) {
		return nil, false
	}
	return o.ExpectedLength, true
}

// HasExpectedLength returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsExpectedLengthSet() bool {
	if o != nil && !IsNil(o.ExpectedLength) {
		return true
	}

	return false
}

// SetExpectedLength gets a reference to the given int32 and assigns it to the ExpectedLength field.
func (o *AiRegenerateStreamRequest) SetExpectedLength(v int32) {
	o.ExpectedLength = &v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiRegenerateStreamRequest) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiRegenerateStreamRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetContextEntityId returns the ContextEntityId field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetContextEntityId() string {
	if o == nil || IsNil(o.ContextEntityId) {
		var ret string
		return ret
	}
	return *o.ContextEntityId
}

// GetContextEntityIdOk returns a tuple with the ContextEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetContextEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContextEntityId) {
		return nil, false
	}
	return o.ContextEntityId, true
}

// HasContextEntityId returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsContextEntityIdSet() bool {
	if o != nil && !IsNil(o.ContextEntityId) {
		return true
	}

	return false
}

// SetContextEntityId gets a reference to the given string and assigns it to the ContextEntityId field.
func (o *AiRegenerateStreamRequest) SetContextEntityId(v string) {
	o.ContextEntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiRegenerateStreamRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiRegenerateStreamRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiRegenerateStreamRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiRegenerateStreamRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiRegenerateStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiRegenerateStreamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.KeepCount) {
		toSerialize["keepCount"] = o.KeepCount
	}
	if !IsNil(o.ExpectedLength) {
		toSerialize["expectedLength"] = o.ExpectedLength
	}
	if !IsNil(o.ActionArgs) {
		toSerialize["actionArgs"] = o.ActionArgs
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.ContextEntityId) {
		toSerialize["contextEntityId"] = o.ContextEntityId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

func (o *AiRegenerateStreamRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
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

	varAiRegenerateStreamRequest := _AiRegenerateStreamRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiRegenerateStreamRequest)

	if err != nil {
		return err
	}

	*o = AiRegenerateStreamRequest(varAiRegenerateStreamRequest)

	return err
}

type NullableAiRegenerateStreamRequest struct {
	value *AiRegenerateStreamRequest
	isSet bool
}

func (v NullableAiRegenerateStreamRequest) Get() *AiRegenerateStreamRequest {
	return v.value
}

func (v *NullableAiRegenerateStreamRequest) Set(val *AiRegenerateStreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiRegenerateStreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiRegenerateStreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiRegenerateStreamRequest(val *AiRegenerateStreamRequest) *NullableAiRegenerateStreamRequest {
	return &NullableAiRegenerateStreamRequest{value: val, isSet: true}
}

func (v NullableAiRegenerateStreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiRegenerateStreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

