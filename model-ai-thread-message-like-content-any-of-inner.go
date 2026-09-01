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

// checks if the AiThreadMessageLikeContentAnyOfInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThreadMessageLikeContentAnyOfInner{}

// AiThreadMessageLikeContentAnyOfInner struct for AiThreadMessageLikeContentAnyOfInner
type AiThreadMessageLikeContentAnyOfInner struct {
	Type string `json:"type"`
	Text *string `json:"text,omitempty"`
}

type _AiThreadMessageLikeContentAnyOfInner AiThreadMessageLikeContentAnyOfInner

// NewAiThreadMessageLikeContentAnyOfInner instantiates a new AiThreadMessageLikeContentAnyOfInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThreadMessageLikeContentAnyOfInner(type_ string) *AiThreadMessageLikeContentAnyOfInner {
	this := AiThreadMessageLikeContentAnyOfInner{}
	this.Type = type_
	return &this
}

// NewAiThreadMessageLikeContentAnyOfInnerWithDefaults instantiates a new AiThreadMessageLikeContentAnyOfInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadMessageLikeContentAnyOfInnerWithDefaults() *AiThreadMessageLikeContentAnyOfInner {
	this := AiThreadMessageLikeContentAnyOfInner{}
	return &this
}

// GetType returns the Type field value
func (o *AiThreadMessageLikeContentAnyOfInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiThreadMessageLikeContentAnyOfInner) SetType(v string) {
	o.Type = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiThreadMessageLikeContentAnyOfInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiThreadMessageLikeContentAnyOfInner) IsTextSet() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiThreadMessageLikeContentAnyOfInner) SetText(v string) {
	o.Text = &v
}

func (o AiThreadMessageLikeContentAnyOfInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThreadMessageLikeContentAnyOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

func (o *AiThreadMessageLikeContentAnyOfInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiThreadMessageLikeContentAnyOfInner := _AiThreadMessageLikeContentAnyOfInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThreadMessageLikeContentAnyOfInner)

	if err != nil {
		return err
	}

	*o = AiThreadMessageLikeContentAnyOfInner(varAiThreadMessageLikeContentAnyOfInner)

	return err
}

type NullableAiThreadMessageLikeContentAnyOfInner struct {
	value *AiThreadMessageLikeContentAnyOfInner
	isSet bool
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) Get() *AiThreadMessageLikeContentAnyOfInner {
	return v.value
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) Set(val *AiThreadMessageLikeContentAnyOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThreadMessageLikeContentAnyOfInner(val *AiThreadMessageLikeContentAnyOfInner) *NullableAiThreadMessageLikeContentAnyOfInner {
	return &NullableAiThreadMessageLikeContentAnyOfInner{value: val, isSet: true}
}

func (v NullableAiThreadMessageLikeContentAnyOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThreadMessageLikeContentAnyOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

