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

// checks if the AiAiSendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiSendRequest{}

// AiAiSendRequest struct for AiAiSendRequest
type AiAiSendRequest struct {
	// Which AI action to run — selects the assignment slot and action.
	ActionType AiActionType `json:"actionType"`
	// The user turn to send.
	UserMessage AiThreadMessageLike `json:"userMessage"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiAiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
}

type _AiAiSendRequest AiAiSendRequest

// NewAiAiSendRequest instantiates a new AiAiSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiSendRequest(actionType AiActionType, userMessage AiThreadMessageLike) *AiAiSendRequest {
	this := AiAiSendRequest{}
	this.ActionType = actionType
	this.UserMessage = userMessage
	return &this
}

// NewAiAiSendRequestWithDefaults instantiates a new AiAiSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiSendRequestWithDefaults() *AiAiSendRequest {
	this := AiAiSendRequest{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *AiAiSendRequest) GetActionType() AiActionType {
	if o == nil {
		var ret AiActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AiAiSendRequest) GetActionTypeOk() (*AiActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *AiAiSendRequest) SetActionType(v AiActionType) {
	o.ActionType = v
}

// GetUserMessage returns the UserMessage field value
func (o *AiAiSendRequest) GetUserMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiAiSendRequest) GetUserMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value
func (o *AiAiSendRequest) SetUserMessage(v AiThreadMessageLike) {
	o.UserMessage = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiAiSendRequest) GetActionArgs() AiAiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiAiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSendRequest) GetActionArgsOk() (*AiAiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiAiSendRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiAiActionArgs and assigns it to the ActionArgs field.
func (o *AiAiSendRequest) SetActionArgs(v AiAiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiAiSendRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSendRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiAiSendRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiAiSendRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiAiSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiSendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	toSerialize["userMessage"] = o.UserMessage
	if !IsNil(o.ActionArgs) {
		toSerialize["actionArgs"] = o.ActionArgs
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiAiSendRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionType",
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

	varAiAiSendRequest := _AiAiSendRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAiSendRequest)

	if err != nil {
		return err
	}

	*o = AiAiSendRequest(varAiAiSendRequest)

	return err
}

type NullableAiAiSendRequest struct {
	value *AiAiSendRequest
	isSet bool
}

func (v NullableAiAiSendRequest) Get() *AiAiSendRequest {
	return v.value
}

func (v *NullableAiAiSendRequest) Set(val *AiAiSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiSendRequest(val *AiAiSendRequest) *NullableAiAiSendRequest {
	return &NullableAiAiSendRequest{value: val, isSet: true}
}

func (v NullableAiAiSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

