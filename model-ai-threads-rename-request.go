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

// checks if the AiThreadsRenameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsRenameRequest{}

// AiThreadsRenameRequest struct for AiThreadsRenameRequest
type AiThreadsRenameRequest struct {
	ThreadId string `json:"threadId"`
	// New thread title.
	Title string `json:"title"`
}

type _AiThreadsRenameRequest AiThreadsRenameRequest

// NewAiThreadsRenameRequest instantiates a new AiThreadsRenameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsRenameRequest(threadId string, title string) *AiThreadsRenameRequest {
	this := AiThreadsRenameRequest{}
	this.ThreadId = threadId
	this.Title = title
	return &this
}

// NewAiThreadsRenameRequestWithDefaults instantiates a new AiThreadsRenameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsRenameRequestWithDefaults() *AiThreadsRenameRequest {
	this := AiThreadsRenameRequest{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiThreadsRenameRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsRenameRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiThreadsRenameRequest) SetThreadId(v string) {
	o.ThreadId = v
}

// GetTitle returns the Title field value
func (o *AiThreadsRenameRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiThreadsRenameRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiThreadsRenameRequest) SetTitle(v string) {
	o.Title = v
}

func (o AiThreadsRenameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsRenameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *AiThreadsRenameRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
		"title",
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

	varAiThreadsRenameRequest := _AiThreadsRenameRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsRenameRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsRenameRequest(varAiThreadsRenameRequest)

	return err
}

type NullableAiThreadsRenameRequest struct {
	value *AiThreadsRenameRequest
	isSet bool
}

func (v NullableAiThreadsRenameRequest) Get() *AiThreadsRenameRequest {
	return v.value
}

func (v *NullableAiThreadsRenameRequest) Set(val *AiThreadsRenameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsRenameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsRenameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsRenameRequest(val *AiThreadsRenameRequest) *NullableAiThreadsRenameRequest {
	return &NullableAiThreadsRenameRequest{value: val, isSet: true}
}

func (v NullableAiThreadsRenameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsRenameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

