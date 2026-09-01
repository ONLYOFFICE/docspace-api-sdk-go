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

// checks if the AiThreadsRegenerateTitleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsRegenerateTitleRequest{}

// AiThreadsRegenerateTitleRequest struct for AiThreadsRegenerateTitleRequest
type AiThreadsRegenerateTitleRequest struct {
	ThreadId string `json:"threadId"`
	// Profile used to regenerate the title.
	Profile AiProfile `json:"profile"`
	EntityMeta *AiThreadsOpenOrCreateRequestEntityMeta `json:"entityMeta,omitempty"`
}

type _AiThreadsRegenerateTitleRequest AiThreadsRegenerateTitleRequest

// NewAiThreadsRegenerateTitleRequest instantiates a new AiThreadsRegenerateTitleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsRegenerateTitleRequest(threadId string, profile AiProfile) *AiThreadsRegenerateTitleRequest {
	this := AiThreadsRegenerateTitleRequest{}
	this.ThreadId = threadId
	this.Profile = profile
	return &this
}

// NewAiThreadsRegenerateTitleRequestWithDefaults instantiates a new AiThreadsRegenerateTitleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsRegenerateTitleRequestWithDefaults() *AiThreadsRegenerateTitleRequest {
	this := AiThreadsRegenerateTitleRequest{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiThreadsRegenerateTitleRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsRegenerateTitleRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiThreadsRegenerateTitleRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetProfile returns the Profile field value
func (o *AiThreadsRegenerateTitleRequest) GetProfile() AiProfile {
	if o == nil {
		var ret AiProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AiThreadsRegenerateTitleRequest) GetProfileOk() (*AiProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AiThreadsRegenerateTitleRequest) SetProfile(v AiProfile) {
	o.Profile = v
}

// GetEntityMeta returns the EntityMeta field value if set, zero value otherwise.
func (o *AiThreadsRegenerateTitleRequest) GetEntityMeta() AiThreadsOpenOrCreateRequestEntityMeta {
	if o == nil || IsNil(o.EntityMeta) {
		var ret AiThreadsOpenOrCreateRequestEntityMeta
		return ret
	}
	return *o.EntityMeta
}

// GetEntityMetaOk returns a tuple with the EntityMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsRegenerateTitleRequest) GetEntityMetaOk() (*AiThreadsOpenOrCreateRequestEntityMeta, bool) {
	if o == nil || IsNil(o.EntityMeta) {
		return nil, false
	}
	return o.EntityMeta, true
}

// HasEntityMeta returns a boolean if a field has been set.
func (o *AiThreadsRegenerateTitleRequest) IsEntityMetaSet() bool {
	if o != nil && !IsNil(o.EntityMeta) {
		return true
	}

	return false
}

// SetEntityMeta gets a reference to the given AiThreadsOpenOrCreateRequestEntityMeta and assigns it to the EntityMeta field.
func (o *AiThreadsRegenerateTitleRequest) SetEntityMeta(v AiThreadsOpenOrCreateRequestEntityMeta) {
	o.EntityMeta = &v
}

func (o AiThreadsRegenerateTitleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsRegenerateTitleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["profile"] = o.Profile
	if !IsNil(o.EntityMeta) {
		toSerialize["entityMeta"] = o.EntityMeta
	}
	return toSerialize, nil
}

func (o *AiThreadsRegenerateTitleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
		"profile",
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

	varAiThreadsRegenerateTitleRequest := _AiThreadsRegenerateTitleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsRegenerateTitleRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsRegenerateTitleRequest(varAiThreadsRegenerateTitleRequest)

	return err
}

type NullableAiThreadsRegenerateTitleRequest struct {
	value *AiThreadsRegenerateTitleRequest
	isSet bool
}

func (v NullableAiThreadsRegenerateTitleRequest) Get() *AiThreadsRegenerateTitleRequest {
	return v.value
}

func (v *NullableAiThreadsRegenerateTitleRequest) Set(val *AiThreadsRegenerateTitleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsRegenerateTitleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsRegenerateTitleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsRegenerateTitleRequest(val *AiThreadsRegenerateTitleRequest) *NullableAiThreadsRegenerateTitleRequest {
	return &NullableAiThreadsRegenerateTitleRequest{value: val, isSet: true}
}

func (v NullableAiThreadsRegenerateTitleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsRegenerateTitleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

