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

// checks if the GeneratedFileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratedFileDto{}

// GeneratedFileDto Information about a file created by an AI editor generation tool.
type GeneratedFileDto struct {
	// The unique identifier of the created file.
	Id *int32 `json:"id,omitempty"`
	// The file title, including extension.
	Title NullableString `json:"title"`
	// The file extension.
	Extension NullableString `json:"extension"`
}

type _GeneratedFileDto GeneratedFileDto

// NewGeneratedFileDto instantiates a new GeneratedFileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratedFileDto(title NullableString, extension NullableString) *GeneratedFileDto {
	this := GeneratedFileDto{}
	this.Title = title
	this.Extension = extension
	return &this
}

// NewGeneratedFileDtoWithDefaults instantiates a new GeneratedFileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratedFileDtoWithDefaults() *GeneratedFileDto {
	this := GeneratedFileDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GeneratedFileDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratedFileDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GeneratedFileDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GeneratedFileDto) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeneratedFileDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeneratedFileDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *GeneratedFileDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetExtension returns the Extension field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GeneratedFileDto) GetExtension() string {
	if o == nil || o.Extension.Get() == nil {
		var ret string
		return ret
	}

	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeneratedFileDto) GetExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// SetExtension sets field value
func (o *GeneratedFileDto) SetExtension(v string) {
	o.Extension.Set(&v)
}

func (o GeneratedFileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneratedFileDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title.Get()
	toSerialize["extension"] = o.Extension.Get()
	return toSerialize, nil
}

func (o *GeneratedFileDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"extension",
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

	varGeneratedFileDto := _GeneratedFileDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGeneratedFileDto)

	if err != nil {
		return err
	}

	*o = GeneratedFileDto(varGeneratedFileDto)

	return err
}

type NullableGeneratedFileDto struct {
	value *GeneratedFileDto
	isSet bool
}

func (v NullableGeneratedFileDto) Get() *GeneratedFileDto {
	return v.value
}

func (v *NullableGeneratedFileDto) Set(val *GeneratedFileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratedFileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratedFileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratedFileDto(val *GeneratedFileDto) *NullableGeneratedFileDto {
	return &NullableGeneratedFileDto{value: val, isSet: true}
}

func (v NullableGeneratedFileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratedFileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

