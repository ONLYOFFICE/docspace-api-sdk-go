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

// checks if the AiAiApproveToolCallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiApproveToolCallRequest{}

// AiAiApproveToolCallRequest struct for AiAiApproveToolCallRequest
type AiAiApproveToolCallRequest struct {
	Result interface{} `json:"result"`
	// Persist auto-approve for this tool's name.
	AllowAlways *bool `json:"allowAlways,omitempty"`
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

type _AiAiApproveToolCallRequest AiAiApproveToolCallRequest

// NewAiAiApproveToolCallRequest instantiates a new AiAiApproveToolCallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiApproveToolCallRequest(result interface{}, threadId string, messageId string, idx float32, message AiThreadMessageLike) *AiAiApproveToolCallRequest {
	this := AiAiApproveToolCallRequest{}
	this.Result = result
	this.ThreadId = threadId
	this.MessageId = messageId
	this.Idx = idx
	this.Message = message
	return &this
}

// NewAiAiApproveToolCallRequestWithDefaults instantiates a new AiAiApproveToolCallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiApproveToolCallRequestWithDefaults() *AiAiApproveToolCallRequest {
	this := AiAiApproveToolCallRequest{}
	return &this
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AiAiApproveToolCallRequest) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiAiApproveToolCallRequest) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AiAiApproveToolCallRequest) SetResult(v interface{}) {
	o.Result = v
}

// GetAllowAlways returns the AllowAlways field value if set, zero value otherwise.
func (o *AiAiApproveToolCallRequest) GetAllowAlways() bool {
	if o == nil || IsNil(o.AllowAlways) {
		var ret bool
		return ret
	}
	return *o.AllowAlways
}

// GetAllowAlwaysOk returns a tuple with the AllowAlways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetAllowAlwaysOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAlways) {
		return nil, false
	}
	return o.AllowAlways, true
}

// HasAllowAlways returns a boolean if a field has been set.
func (o *AiAiApproveToolCallRequest) IsAllowAlwaysSet() bool {
	if o != nil && !IsNil(o.AllowAlways) {
		return true
	}

	return false
}

// SetAllowAlways gets a reference to the given bool and assigns it to the AllowAlways field.
func (o *AiAiApproveToolCallRequest) SetAllowAlways(v bool) {
	o.AllowAlways = &v
}

// GetThreadId returns the ThreadId field value
func (o *AiAiApproveToolCallRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiAiApproveToolCallRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetMessageId returns the MessageId field value
func (o *AiAiApproveToolCallRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiAiApproveToolCallRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetIdx returns the Idx field value
func (o *AiAiApproveToolCallRequest) GetIdx() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Idx
}

// GetIdxOk returns a tuple with the Idx field value
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetIdxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idx, true
}

// SetIdx sets field value
func (o *AiAiApproveToolCallRequest) SetIdx(v float32) {
	o.Idx = v
}

// GetMessage returns the Message field value
func (o *AiAiApproveToolCallRequest) GetMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiAiApproveToolCallRequest) SetMessage(v AiThreadMessageLike) {
	o.Message = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiAiApproveToolCallRequest) GetActionArgs() AiAiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiAiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetActionArgsOk() (*AiAiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiAiApproveToolCallRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiAiActionArgs and assigns it to the ActionArgs field.
func (o *AiAiApproveToolCallRequest) SetActionArgs(v AiAiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiAiApproveToolCallRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiAiApproveToolCallRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiAiApproveToolCallRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiAiApproveToolCallRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiApproveToolCallRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiAiApproveToolCallRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiAiApproveToolCallRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiAiApproveToolCallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiApproveToolCallRequest) ToMap() (map[string]interface{}, error) {
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

func (o *AiAiApproveToolCallRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varAiAiApproveToolCallRequest := _AiAiApproveToolCallRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAiApproveToolCallRequest)

	if err != nil {
		return err
	}

	*o = AiAiApproveToolCallRequest(varAiAiApproveToolCallRequest)

	return err
}

type NullableAiAiApproveToolCallRequest struct {
	value *AiAiApproveToolCallRequest
	isSet bool
}

func (v NullableAiAiApproveToolCallRequest) Get() *AiAiApproveToolCallRequest {
	return v.value
}

func (v *NullableAiAiApproveToolCallRequest) Set(val *AiAiApproveToolCallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiApproveToolCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiApproveToolCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiApproveToolCallRequest(val *AiAiApproveToolCallRequest) *NullableAiAiApproveToolCallRequest {
	return &NullableAiAiApproveToolCallRequest{value: val, isSet: true}
}

func (v NullableAiAiApproveToolCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiApproveToolCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

