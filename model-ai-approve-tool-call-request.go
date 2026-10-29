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

// checks if the AiApproveToolCallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiApproveToolCallRequest{}

// AiApproveToolCallRequest struct for AiApproveToolCallRequest
type AiApproveToolCallRequest struct {
	Result interface{} `json:"result"`
	// Persist auto-approve for this tool's name.
	AllowAlways *bool `json:"allowAlways,omitempty"`
	// Thread the assistant message belongs to.
	ThreadId string `json:"threadId"`
	// Storage id of the assistant message holding the tool call.
	MessageId string `json:"messageId"`
	// Index of the tool-call content part inside `message.content`.
	Idx float32 `json:"idx"`
	// Snapshot of the assistant message at the time the tool call surfaced. Optional and ignored: the engine resumes from the stored message, and the copy inflated the body past the parser limit (Bug 84608).
	Message *AiThreadMessageLike `json:"message,omitempty"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
	// Session-level profile override for this request only.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiApproveToolCallRequest AiApproveToolCallRequest

// NewAiApproveToolCallRequest instantiates a new AiApproveToolCallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiApproveToolCallRequest(result interface{}, threadId string, messageId string, idx float32) *AiApproveToolCallRequest {
	this := AiApproveToolCallRequest{}
	this.Result = result
	this.ThreadId = threadId
	this.MessageId = messageId
	this.Idx = idx
	return &this
}

// NewAiApproveToolCallRequestWithDefaults instantiates a new AiApproveToolCallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiApproveToolCallRequestWithDefaults() *AiApproveToolCallRequest {
	this := AiApproveToolCallRequest{}
	return &this
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AiApproveToolCallRequest) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiApproveToolCallRequest) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AiApproveToolCallRequest) SetResult(v interface{}) {
	o.Result = v
}

// GetAllowAlways returns the AllowAlways field value if set, zero value otherwise.
func (o *AiApproveToolCallRequest) GetAllowAlways() bool {
	if o == nil || IsNil(o.AllowAlways) {
		var ret bool
		return ret
	}
	return *o.AllowAlways
}

// GetAllowAlwaysOk returns a tuple with the AllowAlways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetAllowAlwaysOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAlways) {
		return nil, false
	}
	return o.AllowAlways, true
}

// HasAllowAlways returns a boolean if a field has been set.
func (o *AiApproveToolCallRequest) IsAllowAlwaysSet() bool {
	if o != nil && !IsNil(o.AllowAlways) {
		return true
	}

	return false
}

// SetAllowAlways gets a reference to the given bool and assigns it to the AllowAlways field.
func (o *AiApproveToolCallRequest) SetAllowAlways(v bool) {
	o.AllowAlways = &v
}

// GetThreadId returns the ThreadId field value
func (o *AiApproveToolCallRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiApproveToolCallRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetMessageId returns the MessageId field value
func (o *AiApproveToolCallRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiApproveToolCallRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetIdx returns the Idx field value
func (o *AiApproveToolCallRequest) GetIdx() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Idx
}

// GetIdxOk returns a tuple with the Idx field value
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetIdxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idx, true
}

// SetIdx sets field value
func (o *AiApproveToolCallRequest) SetIdx(v float32) {
	o.Idx = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AiApproveToolCallRequest) GetMessage() AiThreadMessageLike {
	if o == nil || IsNil(o.Message) {
		var ret AiThreadMessageLike
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AiApproveToolCallRequest) IsMessageSet() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AiThreadMessageLike and assigns it to the Message field.
func (o *AiApproveToolCallRequest) SetMessage(v AiThreadMessageLike) {
	o.Message = &v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiApproveToolCallRequest) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiApproveToolCallRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiApproveToolCallRequest) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiApproveToolCallRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiApproveToolCallRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiApproveToolCallRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiApproveToolCallRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiApproveToolCallRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiApproveToolCallRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiApproveToolCallRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiApproveToolCallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiApproveToolCallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.AllowAlways) {
		toSerialize["allowAlways"] = o.AllowAlways
	}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["messageId"] = o.MessageId
	toSerialize["idx"] = o.Idx
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
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

func (o *AiApproveToolCallRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
		"threadId",
		"messageId",
		"idx",
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

	varAiApproveToolCallRequest := _AiApproveToolCallRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiApproveToolCallRequest)

	if err != nil {
		return err
	}

	*o = AiApproveToolCallRequest(varAiApproveToolCallRequest)

	return err
}

type NullableAiApproveToolCallRequest struct {
	value *AiApproveToolCallRequest
	isSet bool
}

func (v NullableAiApproveToolCallRequest) Get() *AiApproveToolCallRequest {
	return v.value
}

func (v *NullableAiApproveToolCallRequest) Set(val *AiApproveToolCallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiApproveToolCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiApproveToolCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiApproveToolCallRequest(val *AiApproveToolCallRequest) *NullableAiApproveToolCallRequest {
	return &NullableAiApproveToolCallRequest{value: val, isSet: true}
}

func (v NullableAiApproveToolCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiApproveToolCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

