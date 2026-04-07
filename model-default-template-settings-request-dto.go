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

// checks if the DefaultTemplateSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultTemplateSettingsRequestDto{}

// DefaultTemplateSettingsRequestDto Default templates settings request parameters.
type DefaultTemplateSettingsRequestDto struct {
	SelectedFile DefaultTemplateSettingsRequestDtoSelectedFile `json:"selectedFile"`
	// File extension of a template to replace
	FileExtension NullableString `json:"fileExtension"`
}

type _DefaultTemplateSettingsRequestDto DefaultTemplateSettingsRequestDto

// NewDefaultTemplateSettingsRequestDto instantiates a new DefaultTemplateSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultTemplateSettingsRequestDto(selectedFile DefaultTemplateSettingsRequestDtoSelectedFile, fileExtension NullableString) *DefaultTemplateSettingsRequestDto {
	this := DefaultTemplateSettingsRequestDto{}
	this.SelectedFile = selectedFile
	this.FileExtension = fileExtension
	return &this
}

// NewDefaultTemplateSettingsRequestDtoWithDefaults instantiates a new DefaultTemplateSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultTemplateSettingsRequestDtoWithDefaults() *DefaultTemplateSettingsRequestDto {
	this := DefaultTemplateSettingsRequestDto{}
	return &this
}

// GetSelectedFile returns the SelectedFile field value
func (o *DefaultTemplateSettingsRequestDto) GetSelectedFile() DefaultTemplateSettingsRequestDtoSelectedFile {
	if o == nil {
		var ret DefaultTemplateSettingsRequestDtoSelectedFile
		return ret
	}

	return o.SelectedFile
}

// GetSelectedFileOk returns a tuple with the SelectedFile field value
// and a boolean to check if the value has been set.
func (o *DefaultTemplateSettingsRequestDto) GetSelectedFileOk() (*DefaultTemplateSettingsRequestDtoSelectedFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedFile, true
}

// SetSelectedFile sets field value
func (o *DefaultTemplateSettingsRequestDto) SetSelectedFile(v DefaultTemplateSettingsRequestDtoSelectedFile) {
	o.SelectedFile = v
}

// GetFileExtension returns the FileExtension field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DefaultTemplateSettingsRequestDto) GetFileExtension() string {
	if o == nil || o.FileExtension.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileExtension.Get()
}

// GetFileExtensionOk returns a tuple with the FileExtension field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateSettingsRequestDto) GetFileExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileExtension.Get(), o.FileExtension.IsSet()
}

// SetFileExtension sets field value
func (o *DefaultTemplateSettingsRequestDto) SetFileExtension(v string) {
	o.FileExtension.Set(&v)
}

func (o DefaultTemplateSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultTemplateSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selectedFile"] = o.SelectedFile
	toSerialize["fileExtension"] = o.FileExtension.Get()
	return toSerialize, nil
}

func (o *DefaultTemplateSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"selectedFile",
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

	varDefaultTemplateSettingsRequestDto := _DefaultTemplateSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultTemplateSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = DefaultTemplateSettingsRequestDto(varDefaultTemplateSettingsRequestDto)

	return err
}

type NullableDefaultTemplateSettingsRequestDto struct {
	value *DefaultTemplateSettingsRequestDto
	isSet bool
}

func (v NullableDefaultTemplateSettingsRequestDto) Get() *DefaultTemplateSettingsRequestDto {
	return v.value
}

func (v *NullableDefaultTemplateSettingsRequestDto) Set(val *DefaultTemplateSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultTemplateSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultTemplateSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultTemplateSettingsRequestDto(val *DefaultTemplateSettingsRequestDto) *NullableDefaultTemplateSettingsRequestDto {
	return &NullableDefaultTemplateSettingsRequestDto{value: val, isSet: true}
}

func (v NullableDefaultTemplateSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultTemplateSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

