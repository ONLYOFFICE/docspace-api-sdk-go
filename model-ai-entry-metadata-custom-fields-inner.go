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

// checks if the AiEntryMetadataCustomFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryMetadataCustomFieldsInner{}

// AiEntryMetadataCustomFieldsInner struct for AiEntryMetadataCustomFieldsInner
type AiEntryMetadataCustomFieldsInner struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _AiEntryMetadataCustomFieldsInner AiEntryMetadataCustomFieldsInner

// NewAiEntryMetadataCustomFieldsInner instantiates a new AiEntryMetadataCustomFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryMetadataCustomFieldsInner(name string, value string) *AiEntryMetadataCustomFieldsInner {
	this := AiEntryMetadataCustomFieldsInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewAiEntryMetadataCustomFieldsInnerWithDefaults instantiates a new AiEntryMetadataCustomFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryMetadataCustomFieldsInnerWithDefaults() *AiEntryMetadataCustomFieldsInner {
	this := AiEntryMetadataCustomFieldsInner{}
	return &this
}

// GetName returns the Name field value
func (o *AiEntryMetadataCustomFieldsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataCustomFieldsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiEntryMetadataCustomFieldsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *AiEntryMetadataCustomFieldsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataCustomFieldsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AiEntryMetadataCustomFieldsInner) SetValue(v string) {
	o.Value = v
}

func (o AiEntryMetadataCustomFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryMetadataCustomFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *AiEntryMetadataCustomFieldsInner) UnmarshalJSON(data []byte) (err error) {
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

	varAiEntryMetadataCustomFieldsInner := _AiEntryMetadataCustomFieldsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryMetadataCustomFieldsInner)

	if err != nil {
		return err
	}

	*o = AiEntryMetadataCustomFieldsInner(varAiEntryMetadataCustomFieldsInner)

	return err
}

type NullableAiEntryMetadataCustomFieldsInner struct {
	value *AiEntryMetadataCustomFieldsInner
	isSet bool
}

func (v NullableAiEntryMetadataCustomFieldsInner) Get() *AiEntryMetadataCustomFieldsInner {
	return v.value
}

func (v *NullableAiEntryMetadataCustomFieldsInner) Set(val *AiEntryMetadataCustomFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryMetadataCustomFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryMetadataCustomFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryMetadataCustomFieldsInner(val *AiEntryMetadataCustomFieldsInner) *NullableAiEntryMetadataCustomFieldsInner {
	return &NullableAiEntryMetadataCustomFieldsInner{value: val, isSet: true}
}

func (v NullableAiEntryMetadataCustomFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryMetadataCustomFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

