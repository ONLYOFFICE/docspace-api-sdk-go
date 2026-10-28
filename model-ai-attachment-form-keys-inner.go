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

// checks if the AiAttachmentFormKeysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAttachmentFormKeysInner{}

// AiAttachmentFormKeysInner struct for AiAttachmentFormKeysInner
type AiAttachmentFormKeysInner struct {
	Key string `json:"key"`
	Text string `json:"text"`
}

type _AiAttachmentFormKeysInner AiAttachmentFormKeysInner

// NewAiAttachmentFormKeysInner instantiates a new AiAttachmentFormKeysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAttachmentFormKeysInner(key string, text string) *AiAttachmentFormKeysInner {
	this := AiAttachmentFormKeysInner{}
	this.Key = key
	this.Text = text
	return &this
}

// NewAiAttachmentFormKeysInnerWithDefaults instantiates a new AiAttachmentFormKeysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAttachmentFormKeysInnerWithDefaults() *AiAttachmentFormKeysInner {
	this := AiAttachmentFormKeysInner{}
	return &this
}

// GetKey returns the Key field value
func (o *AiAttachmentFormKeysInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentFormKeysInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AiAttachmentFormKeysInner) SetKey(v string) {
	o.Key = v
}

// GetText returns the Text field value
func (o *AiAttachmentFormKeysInner) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentFormKeysInner) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AiAttachmentFormKeysInner) SetText(v string) {
	o.Text = v
}

func (o AiAttachmentFormKeysInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAttachmentFormKeysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *AiAttachmentFormKeysInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"text",
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

	varAiAttachmentFormKeysInner := _AiAttachmentFormKeysInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAttachmentFormKeysInner)

	if err != nil {
		return err
	}

	*o = AiAttachmentFormKeysInner(varAiAttachmentFormKeysInner)

	return err
}

type NullableAiAttachmentFormKeysInner struct {
	value *AiAttachmentFormKeysInner
	isSet bool
}

func (v NullableAiAttachmentFormKeysInner) Get() *AiAttachmentFormKeysInner {
	return v.value
}

func (v *NullableAiAttachmentFormKeysInner) Set(val *AiAttachmentFormKeysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAttachmentFormKeysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAttachmentFormKeysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAttachmentFormKeysInner(val *AiAttachmentFormKeysInner) *NullableAiAttachmentFormKeysInner {
	return &NullableAiAttachmentFormKeysInner{value: val, isSet: true}
}

func (v NullableAiAttachmentFormKeysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAttachmentFormKeysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

