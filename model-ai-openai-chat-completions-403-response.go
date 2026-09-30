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

// checks if the AiOpenaiChatCompletions403Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenaiChatCompletions403Response{}

// AiOpenaiChatCompletions403Response struct for AiOpenaiChatCompletions403Response
type AiOpenaiChatCompletions403Response struct {
	Error AiOpenaiChatCompletions403ResponseError `json:"error"`
}

type _AiOpenaiChatCompletions403Response AiOpenaiChatCompletions403Response

// NewAiOpenaiChatCompletions403Response instantiates a new AiOpenaiChatCompletions403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenaiChatCompletions403Response(error_ AiOpenaiChatCompletions403ResponseError) *AiOpenaiChatCompletions403Response {
	this := AiOpenaiChatCompletions403Response{}
	this.Error = error_
	return &this
}

// NewAiOpenaiChatCompletions403ResponseWithDefaults instantiates a new AiOpenaiChatCompletions403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenaiChatCompletions403ResponseWithDefaults() *AiOpenaiChatCompletions403Response {
	this := AiOpenaiChatCompletions403Response{}
	return &this
}

// GetError returns the Error field value
func (o *AiOpenaiChatCompletions403Response) GetError() AiOpenaiChatCompletions403ResponseError {
	if o == nil {
		var ret AiOpenaiChatCompletions403ResponseError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiOpenaiChatCompletions403Response) GetErrorOk() (*AiOpenaiChatCompletions403ResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiOpenaiChatCompletions403Response) SetError(v AiOpenaiChatCompletions403ResponseError) {
	o.Error = v
}

func (o AiOpenaiChatCompletions403Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenaiChatCompletions403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiOpenaiChatCompletions403Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varAiOpenaiChatCompletions403Response := _AiOpenaiChatCompletions403Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenaiChatCompletions403Response)

	if err != nil {
		return err
	}

	*o = AiOpenaiChatCompletions403Response(varAiOpenaiChatCompletions403Response)

	return err
}

type NullableAiOpenaiChatCompletions403Response struct {
	value *AiOpenaiChatCompletions403Response
	isSet bool
}

func (v NullableAiOpenaiChatCompletions403Response) Get() *AiOpenaiChatCompletions403Response {
	return v.value
}

func (v *NullableAiOpenaiChatCompletions403Response) Set(val *AiOpenaiChatCompletions403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenaiChatCompletions403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenaiChatCompletions403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenaiChatCompletions403Response(val *AiOpenaiChatCompletions403Response) *NullableAiOpenaiChatCompletions403Response {
	return &NullableAiOpenaiChatCompletions403Response{value: val, isSet: true}
}

func (v NullableAiOpenaiChatCompletions403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenaiChatCompletions403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

