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
	// Per-request engine options: extra tools, reasoning, prompt override.
	ActionArgs *AiActionArgs `json:"actionArgs,omitempty"`
	// Optional entity (room) scope for profile resolution.
	EntityId *string `json:"entityId,omitempty"`
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

