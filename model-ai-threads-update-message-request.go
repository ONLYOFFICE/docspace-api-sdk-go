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

// checks if the AiThreadsUpdateMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsUpdateMessageRequest{}

// AiThreadsUpdateMessageRequest struct for AiThreadsUpdateMessageRequest
type AiThreadsUpdateMessageRequest struct {
	MessageId string `json:"messageId"`
	// Replacement message content.
	Message AiThreadMessageLike `json:"message"`
}

type _AiThreadsUpdateMessageRequest AiThreadsUpdateMessageRequest

// NewAiThreadsUpdateMessageRequest instantiates a new AiThreadsUpdateMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsUpdateMessageRequest(messageId string, message AiThreadMessageLike) *AiThreadsUpdateMessageRequest {
	this := AiThreadsUpdateMessageRequest{}
	this.MessageId = messageId
	this.Message = message
	return &this
}

// NewAiThreadsUpdateMessageRequestWithDefaults instantiates a new AiThreadsUpdateMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsUpdateMessageRequestWithDefaults() *AiThreadsUpdateMessageRequest {
	this := AiThreadsUpdateMessageRequest{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *AiThreadsUpdateMessageRequest) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsUpdateMessageRequest) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiThreadsUpdateMessageRequest) SetMessageId(v string) {
	o.MessageId = v
}

// GetMessage returns the Message field value
func (o *AiThreadsUpdateMessageRequest) GetMessage() AiThreadMessageLike {
	if o == nil {
		var ret AiThreadMessageLike
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiThreadsUpdateMessageRequest) GetMessageOk() (*AiThreadMessageLike, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiThreadsUpdateMessageRequest) SetMessage(v AiThreadMessageLike) {
	o.Message = v
}

func (o AiThreadsUpdateMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsUpdateMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageId"] = o.MessageId
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AiThreadsUpdateMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messageId",
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

	varAiThreadsUpdateMessageRequest := _AiThreadsUpdateMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsUpdateMessageRequest)

	if err != nil {
		return err
	}

	*o = AiThreadsUpdateMessageRequest(varAiThreadsUpdateMessageRequest)

	return err
}

type NullableAiThreadsUpdateMessageRequest struct {
	value *AiThreadsUpdateMessageRequest
	isSet bool
}

func (v NullableAiThreadsUpdateMessageRequest) Get() *AiThreadsUpdateMessageRequest {
	return v.value
}

func (v *NullableAiThreadsUpdateMessageRequest) Set(val *AiThreadsUpdateMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsUpdateMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsUpdateMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsUpdateMessageRequest(val *AiThreadsUpdateMessageRequest) *NullableAiThreadsUpdateMessageRequest {
	return &NullableAiThreadsUpdateMessageRequest{value: val, isSet: true}
}

func (v NullableAiThreadsUpdateMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsUpdateMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

