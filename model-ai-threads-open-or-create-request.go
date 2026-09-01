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

// checks if the AiThreadsOpenOrCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsOpenOrCreateRequest{}

// AiThreadsOpenOrCreateRequest struct for AiThreadsOpenOrCreateRequest
type AiThreadsOpenOrCreateRequest struct {
	ThreadId *string `json:"threadId,omitempty"`
	// Profile the title generation runs on.
	Profile AiProfile `json:"profile"`
	ProfileId string `json:"profileId"`
	// First user message a fresh thread derives its title from.
	FirstMessage AiThreadMessageLike `json:"firstMessage"`
	// Opaque scope token persisted on a freshly created thread.
	EntityId *string `json:"entityId,omitempty"`
	EntityMeta *AiThreadsOpenOrCreateRequestEntityMeta `json:"entityMeta,omitempty"`
}

type _AiThreadsOpenOrCreateRequest AiThreadsOpenOrCreateRequest

// NewAiThreadsOpenOrCreateRequest instantiates a new AiThreadsOpenOrCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsOpenOrCreateRequest(profile AiProfile, profileId string, firstMessage AiThreadMessageLike) *AiThreadsOpenOrCreateRequest {
	this := AiThreadsOpenOrCreateRequest{}
	this.Profile = profile
	this.ProfileId = profileId
	this.FirstMessage = firstMessage
	return &this
}

// NewAiThreadsOpenOrCreateRequestWithDefaults instantiates a new AiThreadsOpenOrCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsOpenOrCreateRequestWithDefaults() *AiThreadsOpenOrCreateRequest {
	this := AiThreadsOpenOrCreateRequest{}
	return &this
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *AiThreadsOpenOrCreateRequest) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *AiThreadsOpenOrCreateRequest) IsThreadIdSet() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *AiThreadsOpenOrCreateRequest) SetThreadId(v string) {
	o.ThreadId = &v
}

// GetProfile returns the Profile field value
func (o *AiThreadsOpenOrCreateRequest) GetProfile() AiProfile {
	if o == nil {
		var ret AiProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetProfileOk() (*AiProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AiThreadsOpenOrCreateRequest) SetProfile(v AiProfile) {
	o.Profile = v
}

// GetProfileId returns the ProfileId field value
func (o *AiThreadsOpenOrCreateRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *AiThreadsOpenOrCreateRequest) SetProfileId(v string) {
	o.ProfileId = v
}

// GetFirstMessage returns the FirstMessage field value
func (o *AiThreadsOpenOrCreateRequest) GetFirstMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.FirstMessage
}

// GetFirstMessageOk returns a tuple with the FirstMessage field value
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetFirstMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstMessage, true
}

// SetFirstMessage sets field value
func (o *AiThreadsOpenOrCreateRequest) SetFirstMessage(v AiThreadMessageLike) {
	o.FirstMessage = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiThreadsOpenOrCreateRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiThreadsOpenOrCreateRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiThreadsOpenOrCreateRequest) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityMeta returns the EntityMeta field value if set, zero value otherwise.
func (o *AiThreadsOpenOrCreateRequest) GetEntityMeta() AiThreadsOpenOrCreateRequestEntityMeta {
	if o == nil || IsNil(o.EntityMeta) {
		var ret AiThreadsOpenOrCreateRequestEntityMeta
		return ret
	}
	return *o.EntityMeta
}

// GetEntityMetaOk returns a tuple with the EntityMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsOpenOrCreateRequest) GetEntityMetaOk() (*AiThreadsOpenOrCreateRequestEntityMeta, bool) {
	if o == nil || IsNil(o.EntityMeta) {
		return nil, false
	}
	return o.EntityMeta, true
}

// HasEntityMeta returns a boolean if a field has been set.
func (o *AiThreadsOpenOrCreateRequest) IsEntityMetaSet() bool {
	if o != nil && !IsNil(o.EntityMeta) {
		return true
	}

	return false
}

// SetEntityMeta gets a reference to the given AiThreadsOpenOrCreateRequestEntityMeta and assigns it to the EntityMeta field.
func (o *AiThreadsOpenOrCreateRequest) SetEntityMeta(v AiThreadsOpenOrCreateRequestEntityMeta) {
	o.EntityMeta = &v
}

func (o AiThreadsOpenOrCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsOpenOrCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	toSerialize["profile"] = o.Profile
	toSerialize["profileId"] = o.ProfileId
	toSerialize["firstMessage"] = o.FirstMessage
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityMeta) {
		toSerialize["entityMeta"] = o.EntityMeta
	}
	return toSerialize, nil
}

func (o *AiThreadsOpenOrCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
		"profileId",
		"firstMessage",
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

	varAiThreadsOpenOrCreateRequest := _AiThreadsOpenOrCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsOpenOrCreateRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsOpenOrCreateRequest(varAiThreadsOpenOrCreateRequest)

	return err
}

type NullableAiThreadsOpenOrCreateRequest struct {
	value *AiThreadsOpenOrCreateRequest
	isSet bool
}

func (v NullableAiThreadsOpenOrCreateRequest) Get() *AiThreadsOpenOrCreateRequest {
	return v.value
}

func (v *NullableAiThreadsOpenOrCreateRequest) Set(val *AiThreadsOpenOrCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsOpenOrCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsOpenOrCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsOpenOrCreateRequest(val *AiThreadsOpenOrCreateRequest) *NullableAiThreadsOpenOrCreateRequest {
	return &NullableAiThreadsOpenOrCreateRequest{value: val, isSet: true}
}

func (v NullableAiThreadsOpenOrCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsOpenOrCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

