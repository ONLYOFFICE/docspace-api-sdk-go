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

// checks if the AiThreadsAppendUserMessage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadsAppendUserMessage200Response{}

// AiThreadsAppendUserMessage200Response struct for AiThreadsAppendUserMessage200Response
type AiThreadsAppendUserMessage200Response struct {
	// Identifier of the message that was appended to the thread.
	MessageId string `json:"messageId"`
}

type _AiThreadsAppendUserMessage200Response AiThreadsAppendUserMessage200Response

// NewAiThreadsAppendUserMessage200Response instantiates a new AiThreadsAppendUserMessage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadsAppendUserMessage200Response(messageId string) *AiThreadsAppendUserMessage200Response {
	this := AiThreadsAppendUserMessage200Response{}
	this.MessageId = messageId
	return &this
}

// NewAiThreadsAppendUserMessage200ResponseWithDefaults instantiates a new AiThreadsAppendUserMessage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadsAppendUserMessage200ResponseWithDefaults() *AiThreadsAppendUserMessage200Response {
	this := AiThreadsAppendUserMessage200Response{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *AiThreadsAppendUserMessage200Response) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiThreadsAppendUserMessage200Response) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *AiThreadsAppendUserMessage200Response) SetMessageId(v string) {
	o.MessageId = v
}

func (o AiThreadsAppendUserMessage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadsAppendUserMessage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageId"] = o.MessageId
	return toSerialize, nil
}

func (o *AiThreadsAppendUserMessage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messageId",
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

	varAiThreadsAppendUserMessage200Response := _AiThreadsAppendUserMessage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadsAppendUserMessage200Response)

	if err != nil {
		return err
	}

	*o = AiThreadsAppendUserMessage200Response(varAiThreadsAppendUserMessage200Response)

	return err
}

type NullableAiThreadsAppendUserMessage200Response struct {
	value *AiThreadsAppendUserMessage200Response
	isSet bool
}

func (v NullableAiThreadsAppendUserMessage200Response) Get() *AiThreadsAppendUserMessage200Response {
	return v.value
}

func (v *NullableAiThreadsAppendUserMessage200Response) Set(val *AiThreadsAppendUserMessage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadsAppendUserMessage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadsAppendUserMessage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadsAppendUserMessage200Response(val *AiThreadsAppendUserMessage200Response) *NullableAiThreadsAppendUserMessage200Response {
	return &NullableAiThreadsAppendUserMessage200Response{value: val, isSet: true}
}

func (v NullableAiThreadsAppendUserMessage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadsAppendUserMessage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

