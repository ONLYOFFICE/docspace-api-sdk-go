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
	"fmt"
)

// checks if the AiOpenaiChatCompletions403ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenaiChatCompletions403ResponseError{}

// AiOpenaiChatCompletions403ResponseError struct for AiOpenaiChatCompletions403ResponseError
type AiOpenaiChatCompletions403ResponseError struct {
	// Human-readable description of the failure.
	Message string `json:"message"`
	// OpenAI error class, for example `invalid_request_error`.
	Type string `json:"type"`
	// Machine-readable code, when the provider supplies one.
	Code NullableString `json:"code,omitempty"`
	// The request parameter at fault, when the failure names one.
	Param NullableString `json:"param,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiOpenaiChatCompletions403ResponseError AiOpenaiChatCompletions403ResponseError

// NewAiOpenaiChatCompletions403ResponseError instantiates a new AiOpenaiChatCompletions403ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenaiChatCompletions403ResponseError(message string, type_ string) *AiOpenaiChatCompletions403ResponseError {
	this := AiOpenaiChatCompletions403ResponseError{}
	this.Message = message
	this.Type = type_
	return &this
}

// NewAiOpenaiChatCompletions403ResponseErrorWithDefaults instantiates a new AiOpenaiChatCompletions403ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenaiChatCompletions403ResponseErrorWithDefaults() *AiOpenaiChatCompletions403ResponseError {
	this := AiOpenaiChatCompletions403ResponseError{}
	return &this
}

// GetMessage returns the Message field value
func (o *AiOpenaiChatCompletions403ResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiOpenaiChatCompletions403ResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiOpenaiChatCompletions403ResponseError) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *AiOpenaiChatCompletions403ResponseError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiOpenaiChatCompletions403ResponseError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiOpenaiChatCompletions403ResponseError) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiOpenaiChatCompletions403ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenaiChatCompletions403ResponseError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *AiOpenaiChatCompletions403ResponseError) IsCodeSet() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *AiOpenaiChatCompletions403ResponseError) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *AiOpenaiChatCompletions403ResponseError) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *AiOpenaiChatCompletions403ResponseError) UnsetCode() {
	o.Code.Unset()
}

// GetParam returns the Param field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiOpenaiChatCompletions403ResponseError) GetParam() string {
	if o == nil || IsNil(o.Param.Get()) {
		var ret string
		return ret
	}
	return *o.Param.Get()
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenaiChatCompletions403ResponseError) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Param.Get(), o.Param.IsSet()
}

// HasParam returns a boolean if a field has been set.
func (o *AiOpenaiChatCompletions403ResponseError) IsParamSet() bool {
	if o != nil && o.Param.IsSet() {
		return true
	}

	return false
}

// SetParam gets a reference to the given NullableString and assigns it to the Param field.
func (o *AiOpenaiChatCompletions403ResponseError) SetParam(v string) {
	o.Param.Set(&v)
}
// SetParamNil sets the value for Param to be an explicit nil
func (o *AiOpenaiChatCompletions403ResponseError) SetParamNil() {
	o.Param.Set(nil)
}

// UnsetParam ensures that no value is present for Param, not even an explicit nil
func (o *AiOpenaiChatCompletions403ResponseError) UnsetParam() {
	o.Param.Unset()
}

func (o AiOpenaiChatCompletions403ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenaiChatCompletions403ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["type"] = o.Type
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Param.IsSet() {
		toSerialize["param"] = o.Param.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiOpenaiChatCompletions403ResponseError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"type",
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

	varAiOpenaiChatCompletions403ResponseError := _AiOpenaiChatCompletions403ResponseError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenaiChatCompletions403ResponseError)

	if err != nil {
		return err
	}

	*o = AiOpenaiChatCompletions403ResponseError(varAiOpenaiChatCompletions403ResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "param")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiOpenaiChatCompletions403ResponseError struct {
	value *AiOpenaiChatCompletions403ResponseError
	isSet bool
}

func (v NullableAiOpenaiChatCompletions403ResponseError) Get() *AiOpenaiChatCompletions403ResponseError {
	return v.value
}

func (v *NullableAiOpenaiChatCompletions403ResponseError) Set(val *AiOpenaiChatCompletions403ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenaiChatCompletions403ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenaiChatCompletions403ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenaiChatCompletions403ResponseError(val *AiOpenaiChatCompletions403ResponseError) *NullableAiOpenaiChatCompletions403ResponseError {
	return &NullableAiOpenaiChatCompletions403ResponseError{value: val, isSet: true}
}

func (v NullableAiOpenaiChatCompletions403ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenaiChatCompletions403ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

