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

// checks if the DefaultTemplateSettingsResetRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultTemplateSettingsResetRequestDto{}

// DefaultTemplateSettingsResetRequestDto Default templates settings reset request parameters.
type DefaultTemplateSettingsResetRequestDto struct {
	// File extension of a template to reset
	FileExtension NullableString `json:"fileExtension"`
}

type _DefaultTemplateSettingsResetRequestDto DefaultTemplateSettingsResetRequestDto

// NewDefaultTemplateSettingsResetRequestDto instantiates a new DefaultTemplateSettingsResetRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultTemplateSettingsResetRequestDto(fileExtension NullableString) *DefaultTemplateSettingsResetRequestDto {
	this := DefaultTemplateSettingsResetRequestDto{}
	this.FileExtension = fileExtension
	return &this
}

// NewDefaultTemplateSettingsResetRequestDtoWithDefaults instantiates a new DefaultTemplateSettingsResetRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultTemplateSettingsResetRequestDtoWithDefaults() *DefaultTemplateSettingsResetRequestDto {
	this := DefaultTemplateSettingsResetRequestDto{}
	return &this
}

// GetFileExtension returns the FileExtension field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DefaultTemplateSettingsResetRequestDto) GetFileExtension() string {
	if o == nil || o.FileExtension.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileExtension.Get()
}

// GetFileExtensionOk returns a tuple with the FileExtension field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateSettingsResetRequestDto) GetFileExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileExtension.Get(), o.FileExtension.IsSet()
}

// SetFileExtension sets field value
func (o *DefaultTemplateSettingsResetRequestDto) SetFileExtension(v string) {
	o.FileExtension.Set(&v)
}

func (o DefaultTemplateSettingsResetRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultTemplateSettingsResetRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileExtension"] = o.FileExtension.Get()
	return toSerialize, nil
}

func (o *DefaultTemplateSettingsResetRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileExtension",
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

	varDefaultTemplateSettingsResetRequestDto := _DefaultTemplateSettingsResetRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultTemplateSettingsResetRequestDto)

	if err != nil {
		return err
	}

	*o = DefaultTemplateSettingsResetRequestDto(varDefaultTemplateSettingsResetRequestDto)

	return err
}

type NullableDefaultTemplateSettingsResetRequestDto struct {
	value *DefaultTemplateSettingsResetRequestDto
	isSet bool
}

func (v NullableDefaultTemplateSettingsResetRequestDto) Get() *DefaultTemplateSettingsResetRequestDto {
	return v.value
}

func (v *NullableDefaultTemplateSettingsResetRequestDto) Set(val *DefaultTemplateSettingsResetRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultTemplateSettingsResetRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultTemplateSettingsResetRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultTemplateSettingsResetRequestDto(val *DefaultTemplateSettingsResetRequestDto) *NullableDefaultTemplateSettingsResetRequestDto {
	return &NullableDefaultTemplateSettingsResetRequestDto{value: val, isSet: true}
}

func (v NullableDefaultTemplateSettingsResetRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultTemplateSettingsResetRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

