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

// checks if the AiActionArgsPrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiActionArgsPrompt{}

// AiActionArgsPrompt Override the action's baked-in system prompt (replace or append).
type AiActionArgsPrompt struct {
	Mode string `json:"mode"`
	Text string `json:"text"`
}

type _AiActionArgsPrompt AiActionArgsPrompt

// NewAiActionArgsPrompt instantiates a new AiActionArgsPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiActionArgsPrompt(mode string, text string) *AiActionArgsPrompt {
	this := AiActionArgsPrompt{}
	this.Mode = mode
	this.Text = text
	return &this
}

// NewAiActionArgsPromptWithDefaults instantiates a new AiActionArgsPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiActionArgsPromptWithDefaults() *AiActionArgsPrompt {
	this := AiActionArgsPrompt{}
	return &this
}

// GetMode returns the Mode field value
func (o *AiActionArgsPrompt) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiActionArgsPrompt) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *AiActionArgsPrompt) SetMode(v string) {
	o.Mode = v
}

// GetText returns the Text field value
func (o *AiActionArgsPrompt) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AiActionArgsPrompt) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AiActionArgsPrompt) SetText(v string) {
	o.Text = v
}

func (o AiActionArgsPrompt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiActionArgsPrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

func (o *AiActionArgsPrompt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varAiActionArgsPrompt := _AiActionArgsPrompt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiActionArgsPrompt)

	if err != nil {
		return err
	}

	*o = AiActionArgsPrompt(varAiActionArgsPrompt)

	return err
}

type NullableAiActionArgsPrompt struct {
	value *AiActionArgsPrompt
	isSet bool
}

func (v NullableAiActionArgsPrompt) Get() *AiActionArgsPrompt {
	return v.value
}

func (v *NullableAiActionArgsPrompt) Set(val *AiActionArgsPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableAiActionArgsPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableAiActionArgsPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiActionArgsPrompt(val *AiActionArgsPrompt) *NullableAiActionArgsPrompt {
	return &NullableAiActionArgsPrompt{value: val, isSet: true}
}

func (v NullableAiActionArgsPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiActionArgsPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

