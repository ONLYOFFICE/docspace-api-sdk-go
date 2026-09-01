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

// checks if the AiTErrorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiTErrorData{}

// AiTErrorData A field-scoped validation error: which form field was rejected, and why.
type AiTErrorData struct {
	// The rejected field.
	Field string `json:"field"`
	// The human-readable reason the field was rejected.
	Message string `json:"message"`
}

type _AiTErrorData AiTErrorData

// NewAiTErrorData instantiates a new AiTErrorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiTErrorData(field string, message string) *AiTErrorData {
	this := AiTErrorData{}
	this.Field = field
	this.Message = message
	return &this
}

// NewAiTErrorDataWithDefaults instantiates a new AiTErrorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiTErrorDataWithDefaults() *AiTErrorData {
	this := AiTErrorData{}
	return &this
}

// GetField returns the Field field value
func (o *AiTErrorData) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AiTErrorData) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AiTErrorData) SetField(v string) {
	o.Field = v
}

// GetMessage returns the Message field value
func (o *AiTErrorData) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiTErrorData) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiTErrorData) SetMessage(v string) {
	o.Message = v
}

func (o AiTErrorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiTErrorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AiTErrorData) UnmarshalJSON(data []byte) (err error) {
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

	varAiTErrorData := _AiTErrorData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiTErrorData)

	if err != nil {
		return err
	}

	*o = AiTErrorData(varAiTErrorData)

	return err
}

type NullableAiTErrorData struct {
	value *AiTErrorData
	isSet bool
}

func (v NullableAiTErrorData) Get() *AiTErrorData {
	return v.value
}

func (v *NullableAiTErrorData) Set(val *AiTErrorData) {
	v.value = val
	v.isSet = true
}

func (v NullableAiTErrorData) IsSet() bool {
	return v.isSet
}

func (v *NullableAiTErrorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiTErrorData(val *AiTErrorData) *NullableAiTErrorData {
	return &NullableAiTErrorData{value: val, isSet: true}
}

func (v NullableAiTErrorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiTErrorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

