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

// checks if the AiErrorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiErrorData{}

// AiErrorData A field-scoped validation error: which form field was rejected, and why.
type AiErrorData struct {
	// The rejected field.
	Field string `json:"field"`
	// The human-readable reason the field was rejected.
	Message string `json:"message"`
	// Optional: producers that know the reason set it (the prompts engine does).
	Code *AiErrorCode `json:"code,omitempty"`
}

type _AiErrorData AiErrorData

// NewAiErrorData instantiates a new AiErrorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiErrorData(field string, message string) *AiErrorData {
	this := AiErrorData{}
	this.Field = field
	this.Message = message
	return &this
}

// NewAiErrorDataWithDefaults instantiates a new AiErrorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiErrorDataWithDefaults() *AiErrorData {
	this := AiErrorData{}
	return &this
}

// GetField returns the Field field value
func (o *AiErrorData) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AiErrorData) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AiErrorData) SetField(v string) {
	o.Field = v
}

// GetMessage returns the Message field value
func (o *AiErrorData) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiErrorData) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiErrorData) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AiErrorData) GetCode() AiErrorCode {
	if o == nil || IsNil(o.Code) {
		var ret AiErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiErrorData) GetCodeOk() (*AiErrorCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AiErrorData) IsCodeSet() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given AiErrorCode and assigns it to the Code field.
func (o *AiErrorData) SetCode(v AiErrorCode) {
	o.Code = &v
}

func (o AiErrorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiErrorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["message"] = o.Message
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

func (o *AiErrorData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
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

	varAiErrorData := _AiErrorData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiErrorData)

	if err != nil {
		return err
	}

	*o = AiErrorData(varAiErrorData)

	return err
}

type NullableAiErrorData struct {
	value *AiErrorData
	isSet bool
}

func (v NullableAiErrorData) Get() *AiErrorData {
	return v.value
}

func (v *NullableAiErrorData) Set(val *AiErrorData) {
	v.value = val
	v.isSet = true
}

func (v NullableAiErrorData) IsSet() bool {
	return v.isSet
}

func (v *NullableAiErrorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiErrorData(val *AiErrorData) *NullableAiErrorData {
	return &NullableAiErrorData{value: val, isSet: true}
}

func (v NullableAiErrorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiErrorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

