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

// checks if the AiThreadsTouchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsTouchRequest{}

// AiThreadsTouchRequest struct for AiThreadsTouchRequest
type AiThreadsTouchRequest struct {
	ThreadId string `json:"threadId"`
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiThreadsTouchRequest AiThreadsTouchRequest

// NewAiThreadsTouchRequest instantiates a new AiThreadsTouchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsTouchRequest(threadId string) *AiThreadsTouchRequest {
	this := AiThreadsTouchRequest{}
	this.ThreadId = threadId
	return &this
}

// NewAiThreadsTouchRequestWithDefaults instantiates a new AiThreadsTouchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsTouchRequestWithDefaults() *AiThreadsTouchRequest {
	this := AiThreadsTouchRequest{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiThreadsTouchRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsTouchRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiThreadsTouchRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiThreadsTouchRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadsTouchRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiThreadsTouchRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiThreadsTouchRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiThreadsTouchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsTouchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

func (o *AiThreadsTouchRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAiThreadsTouchRequest := _AiThreadsTouchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsTouchRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsTouchRequest(varAiThreadsTouchRequest)

	return err
}

type NullableAiThreadsTouchRequest struct {
	value *AiThreadsTouchRequest
	isSet bool
}

func (v NullableAiThreadsTouchRequest) Get() *AiThreadsTouchRequest {
	return v.value
}

func (v *NullableAiThreadsTouchRequest) Set(val *AiThreadsTouchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsTouchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsTouchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsTouchRequest(val *AiThreadsTouchRequest) *NullableAiThreadsTouchRequest {
	return &NullableAiThreadsTouchRequest{value: val, isSet: true}
}

func (v NullableAiThreadsTouchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsTouchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

