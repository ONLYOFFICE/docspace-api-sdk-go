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

// checks if the AiOpenAIStreamErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIStreamErrorError{}

// AiOpenAIStreamErrorError The error that ended the stream: its message, type, code and the offending parameter.
type AiOpenAIStreamErrorError struct {
	Message string `json:"message"`
	Type string `json:"type"`
	Code NullableString `json:"code"`
	Param NullableString `json:"param"`
}

type _AiOpenAIStreamErrorError AiOpenAIStreamErrorError

// NewAiOpenAIStreamErrorError instantiates a new AiOpenAIStreamErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIStreamErrorError(message string, type_ string, code NullableString, param NullableString) *AiOpenAIStreamErrorError {
	this := AiOpenAIStreamErrorError{}
	this.Message = message
	this.Type = type_
	this.Code = code
	this.Param = param
	return &this
}

// NewAiOpenAIStreamErrorErrorWithDefaults instantiates a new AiOpenAIStreamErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIStreamErrorErrorWithDefaults() *AiOpenAIStreamErrorError {
	this := AiOpenAIStreamErrorError{}
	return &this
}

// GetMessage returns the Message field value
func (o *AiOpenAIStreamErrorError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIStreamErrorError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiOpenAIStreamErrorError) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *AiOpenAIStreamErrorError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIStreamErrorError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiOpenAIStreamErrorError) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiOpenAIStreamErrorError) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenAIStreamErrorError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *AiOpenAIStreamErrorError) SetCode(v string) {
	o.Code.Set(&v)
}

// GetParam returns the Param field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiOpenAIStreamErrorError) GetParam() string {
	if o == nil || o.Param.Get() == nil {
		var ret string
		return ret
	}

	return *o.Param.Get()
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenAIStreamErrorError) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Param.Get(), o.Param.IsSet()
}

// SetParam sets field value
func (o *AiOpenAIStreamErrorError) SetParam(v string) {
	o.Param.Set(&v)
}

func (o AiOpenAIStreamErrorError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIStreamErrorError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code.Get()
	toSerialize["param"] = o.Param.Get()
	return toSerialize, nil
}

func (o *AiOpenAIStreamErrorError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"type",
		"code",
		"param",
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

	varAiOpenAIStreamErrorError := _AiOpenAIStreamErrorError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenAIStreamErrorError)

	if err != nil {
		return err
	}

	*o = AiOpenAIStreamErrorError(varAiOpenAIStreamErrorError)

	return err
}

type NullableAiOpenAIStreamErrorError struct {
	value *AiOpenAIStreamErrorError
	isSet bool
}

func (v NullableAiOpenAIStreamErrorError) Get() *AiOpenAIStreamErrorError {
	return v.value
}

func (v *NullableAiOpenAIStreamErrorError) Set(val *AiOpenAIStreamErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIStreamErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIStreamErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIStreamErrorError(val *AiOpenAIStreamErrorError) *NullableAiOpenAIStreamErrorError {
	return &NullableAiOpenAIStreamErrorError{value: val, isSet: true}
}

func (v NullableAiOpenAIStreamErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIStreamErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

