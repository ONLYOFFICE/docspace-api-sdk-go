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

// checks if the AiSendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSendRequest{}

// AiSendRequest struct for AiSendRequest
type AiSendRequest struct {
	// Which AI action to run — selects the assignment slot and action.
	ActionType AiActionType `json:"actionType"`
	// The user turn to send.
	UserMessage AiThreadMessageLike `json:"userMessage"`
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
}

type _AiSendRequest AiSendRequest

// NewAiSendRequest instantiates a new AiSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSendRequest(actionType AiActionType, userMessage AiThreadMessageLike) *AiSendRequest {
	this := AiSendRequest{}
	this.ActionType = actionType
	this.UserMessage = userMessage
	return &this
}

// NewAiSendRequestWithDefaults instantiates a new AiSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSendRequestWithDefaults() *AiSendRequest {
	this := AiSendRequest{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *AiSendRequest) GetActionType() AiActionType {
	if o == nil {
		var ret AiActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AiSendRequest) GetActionTypeOk() (*AiActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *AiSendRequest) SetActionType(v AiActionType) {
	o.ActionType = v
}

// GetUserMessage returns the UserMessage field value
func (o *AiSendRequest) GetUserMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiSendRequest) GetUserMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value
func (o *AiSendRequest) SetUserMessage(v AiThreadMessageLike) {
	o.UserMessage = v
}

// GetActionArgs returns the ActionArgs field value if set, zero value otherwise.
func (o *AiSendRequest) GetActionArgs() AiActionArgs {
	if o == nil || IsNil(o.ActionArgs) {
		var ret AiActionArgs
		return ret
	}
	return *o.ActionArgs
}

// GetActionArgsOk returns a tuple with the ActionArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendRequest) GetActionArgsOk() (*AiActionArgs, bool) {
	if o == nil || IsNil(o.ActionArgs) {
		return nil, false
	}
	return o.ActionArgs, true
}

// HasActionArgs returns a boolean if a field has been set.
func (o *AiSendRequest) IsActionArgsSet() bool {
	if o != nil && !IsNil(o.ActionArgs) {
		return true
	}

	return false
}

// SetActionArgs gets a reference to the given AiActionArgs and assigns it to the ActionArgs field.
func (o *AiSendRequest) SetActionArgs(v AiActionArgs) {
	o.ActionArgs = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiSendRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSendRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiSendRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiSendRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSendRequest) ToMap() (map[string]interface{}, error) {
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

func (o *AiSendRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAiSendRequest := _AiSendRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSendRequest)

	if err != nil {
		return err
	}

	*o = AiSendRequest(varAiSendRequest)

	return err
}

type NullableAiSendRequest struct {
	value *AiSendRequest
	isSet bool
}

func (v NullableAiSendRequest) Get() *AiSendRequest {
	return v.value
}

func (v *NullableAiSendRequest) Set(val *AiSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSendRequest(val *AiSendRequest) *NullableAiSendRequest {
	return &NullableAiSendRequest{value: val, isSet: true}
}

func (v NullableAiSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

