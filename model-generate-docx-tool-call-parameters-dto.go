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

// checks if the GenerateDocxToolCallParametersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateDocxToolCallParametersDto{}

// GenerateDocxToolCallParametersDto The generate docx tool call parameters.
type GenerateDocxToolCallParametersDto struct {
	// What the generated text document should contain, in the words the request was made in.
	Description NullableString `json:"description"`
}

type _GenerateDocxToolCallParametersDto GenerateDocxToolCallParametersDto

// NewGenerateDocxToolCallParametersDto instantiates a new GenerateDocxToolCallParametersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDocxToolCallParametersDto(description NullableString) *GenerateDocxToolCallParametersDto {
	this := GenerateDocxToolCallParametersDto{}
	this.Description = description
	return &this
}

// NewGenerateDocxToolCallParametersDtoWithDefaults instantiates a new GenerateDocxToolCallParametersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDocxToolCallParametersDtoWithDefaults() *GenerateDocxToolCallParametersDto {
	this := GenerateDocxToolCallParametersDto{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GenerateDocxToolCallParametersDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateDocxToolCallParametersDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *GenerateDocxToolCallParametersDto) SetDescription(v string) {
	o.Description.Set(&v)
}

func (o GenerateDocxToolCallParametersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateDocxToolCallParametersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description.Get()
	return toSerialize, nil
}

func (o *GenerateDocxToolCallParametersDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varGenerateDocxToolCallParametersDto := _GenerateDocxToolCallParametersDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGenerateDocxToolCallParametersDto)

	if err != nil {
		return err
	}

	*o = GenerateDocxToolCallParametersDto(varGenerateDocxToolCallParametersDto)

	return err
}

type NullableGenerateDocxToolCallParametersDto struct {
	value *GenerateDocxToolCallParametersDto
	isSet bool
}

func (v NullableGenerateDocxToolCallParametersDto) Get() *GenerateDocxToolCallParametersDto {
	return v.value
}

func (v *NullableGenerateDocxToolCallParametersDto) Set(val *GenerateDocxToolCallParametersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDocxToolCallParametersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDocxToolCallParametersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDocxToolCallParametersDto(val *GenerateDocxToolCallParametersDto) *NullableGenerateDocxToolCallParametersDto {
	return &NullableGenerateDocxToolCallParametersDto{value: val, isSet: true}
}

func (v NullableGenerateDocxToolCallParametersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDocxToolCallParametersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

