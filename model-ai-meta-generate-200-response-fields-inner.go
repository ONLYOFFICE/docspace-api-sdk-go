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

// checks if the AiMetaGenerate200ResponseFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerate200ResponseFieldsInner{}

// AiMetaGenerate200ResponseFieldsInner struct for AiMetaGenerate200ResponseFieldsInner
type AiMetaGenerate200ResponseFieldsInner struct {
	// Field name, at most 255 characters.
	Name string `json:"name"`
	// The value found in the documents, or null when the text does not state one.
	Value NullableString `json:"value"`
}

type _AiMetaGenerate200ResponseFieldsInner AiMetaGenerate200ResponseFieldsInner

// NewAiMetaGenerate200ResponseFieldsInner instantiates a new AiMetaGenerate200ResponseFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerate200ResponseFieldsInner(name string, value NullableString) *AiMetaGenerate200ResponseFieldsInner {
	this := AiMetaGenerate200ResponseFieldsInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewAiMetaGenerate200ResponseFieldsInnerWithDefaults instantiates a new AiMetaGenerate200ResponseFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerate200ResponseFieldsInnerWithDefaults() *AiMetaGenerate200ResponseFieldsInner {
	this := AiMetaGenerate200ResponseFieldsInner{}
	return &this
}

// GetName returns the Name field value
func (o *AiMetaGenerate200ResponseFieldsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200ResponseFieldsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiMetaGenerate200ResponseFieldsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiMetaGenerate200ResponseFieldsInner) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiMetaGenerate200ResponseFieldsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *AiMetaGenerate200ResponseFieldsInner) SetValue(v string) {
	o.Value.Set(&v)
}

func (o AiMetaGenerate200ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerate200ResponseFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *AiMetaGenerate200ResponseFieldsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varAiMetaGenerate200ResponseFieldsInner := _AiMetaGenerate200ResponseFieldsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerate200ResponseFieldsInner)

	if err != nil {
		return err
	}

	*o = AiMetaGenerate200ResponseFieldsInner(varAiMetaGenerate200ResponseFieldsInner)

	return err
}

type NullableAiMetaGenerate200ResponseFieldsInner struct {
	value *AiMetaGenerate200ResponseFieldsInner
	isSet bool
}

func (v NullableAiMetaGenerate200ResponseFieldsInner) Get() *AiMetaGenerate200ResponseFieldsInner {
	return v.value
}

func (v *NullableAiMetaGenerate200ResponseFieldsInner) Set(val *AiMetaGenerate200ResponseFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerate200ResponseFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerate200ResponseFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerate200ResponseFieldsInner(val *AiMetaGenerate200ResponseFieldsInner) *NullableAiMetaGenerate200ResponseFieldsInner {
	return &NullableAiMetaGenerate200ResponseFieldsInner{value: val, isSet: true}
}

func (v NullableAiMetaGenerate200ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerate200ResponseFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

