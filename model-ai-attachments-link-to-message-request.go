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

// checks if the AiAttachmentsLinkToMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAttachmentsLinkToMessageRequest{}

// AiAttachmentsLinkToMessageRequest struct for AiAttachmentsLinkToMessageRequest
type AiAttachmentsLinkToMessageRequest struct {
	// Attachment ids to bind.
	Ids []string `json:"ids"`
	// Owning message id.
	MessageId string `json:"messageId"`
	// Owning thread id.
	ThreadId string `json:"threadId"`
}

type _AiAttachmentsLinkToMessageRequest AiAttachmentsLinkToMessageRequest

// NewAiAttachmentsLinkToMessageRequest instantiates a new AiAttachmentsLinkToMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAttachmentsLinkToMessageRequest(ids []string, messageId string, threadId string) *AiAttachmentsLinkToMessageRequest {
	this := AiAttachmentsLinkToMessageRequest{}
	this.Ids = ids
	this.MessageId = messageId
	this.ThreadId = threadId
	return &this
}

// NewAiAttachmentsLinkToMessageRequestWithDefaults instantiates a new AiAttachmentsLinkToMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAttachmentsLinkToMessageRequestWithDefaults() *AiAttachmentsLinkToMessageRequest {
	this := AiAttachmentsLinkToMessageRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *AiAttachmentsLinkToMessageRequest) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsLinkToMessageRequest) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *AiAttachmentsLinkToMessageRequest) SetIds(v []string) {
	o.Ids = v
}

// GetMessageId returns the MessageId field value
func (o *AiAttachmentsLinkToMessageRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsLinkToMessageRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiAttachmentsLinkToMessageRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetThreadId returns the ThreadId field value
func (o *AiAttachmentsLinkToMessageRequest) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsLinkToMessageRequest) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiAttachmentsLinkToMessageRequest) SetThreadId(v string) {
	o.ThreadId = v
}

func (o AiAttachmentsLinkToMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAttachmentsLinkToMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	toSerialize["messageId"] = o.MessageId
	toSerialize["threadId"] = o.ThreadId
	return toSerialize, nil
}

func (o *AiAttachmentsLinkToMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
		"messageId",
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

	varAiAttachmentsLinkToMessageRequest := _AiAttachmentsLinkToMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAttachmentsLinkToMessageRequest)

	if err != nil {
		return err
	}

	*o = AiAttachmentsLinkToMessageRequest(varAiAttachmentsLinkToMessageRequest)

	return err
}

type NullableAiAttachmentsLinkToMessageRequest struct {
	value *AiAttachmentsLinkToMessageRequest
	isSet bool
}

func (v NullableAiAttachmentsLinkToMessageRequest) Get() *AiAttachmentsLinkToMessageRequest {
	return v.value
}

func (v *NullableAiAttachmentsLinkToMessageRequest) Set(val *AiAttachmentsLinkToMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAttachmentsLinkToMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAttachmentsLinkToMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAttachmentsLinkToMessageRequest(val *AiAttachmentsLinkToMessageRequest) *NullableAiAttachmentsLinkToMessageRequest {
	return &NullableAiAttachmentsLinkToMessageRequest{value: val, isSet: true}
}

func (v NullableAiAttachmentsLinkToMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAttachmentsLinkToMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

