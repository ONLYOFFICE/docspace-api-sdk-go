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

// checks if the AiOpenAIStreamError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIStreamError{}

// AiOpenAIStreamError OpenAI streaming error envelope. When the upstream request fails mid-stream the OpenAI API emits a single `data:` line carrying an `error` object (no `choices`), then closes the stream — the official SDK turns this into a thrown `APIError`. Mirrors that shape so a host exposing an OpenAI-compatible endpoint stays wire-compatible.
type AiOpenAIStreamError struct {
	Error AiOpenAIStreamErrorError `json:"error"`
}

type _AiOpenAIStreamError AiOpenAIStreamError

// NewAiOpenAIStreamError instantiates a new AiOpenAIStreamError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIStreamError(error_ AiOpenAIStreamErrorError) *AiOpenAIStreamError {
	this := AiOpenAIStreamError{}
	this.Error = error_
	return &this
}

// NewAiOpenAIStreamErrorWithDefaults instantiates a new AiOpenAIStreamError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIStreamErrorWithDefaults() *AiOpenAIStreamError {
	this := AiOpenAIStreamError{}
	return &this
}

// GetError returns the Error field value
func (o *AiOpenAIStreamError) GetError() AiOpenAIStreamErrorError {
	if o == nil {
		var ret AiOpenAIStreamErrorError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIStreamError) GetErrorOk() (*AiOpenAIStreamErrorError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiOpenAIStreamError) SetError(v AiOpenAIStreamErrorError) {
	o.Error = v
}

func (o AiOpenAIStreamError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIStreamError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiOpenAIStreamError) UnmarshalJSON(data []byte) (err error) {
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

	varAiOpenAIStreamError := _AiOpenAIStreamError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenAIStreamError)

	if err != nil {
		return err
	}

	*o = AiOpenAIStreamError(varAiOpenAIStreamError)

	return err
}

type NullableAiOpenAIStreamError struct {
	value *AiOpenAIStreamError
	isSet bool
}

func (v NullableAiOpenAIStreamError) Get() *AiOpenAIStreamError {
	return v.value
}

func (v *NullableAiOpenAIStreamError) Set(val *AiOpenAIStreamError) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIStreamError) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIStreamError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIStreamError(val *AiOpenAIStreamError) *NullableAiOpenAIStreamError {
	return &NullableAiOpenAIStreamError{value: val, isSet: true}
}

func (v NullableAiOpenAIStreamError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIStreamError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

