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

// checks if the DefaultProductRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultProductRequestDto{}

// DefaultProductRequestDto The section the calling user's account opens into after signing in.
type DefaultProductRequestDto struct {
	// The section to land on. Only the folder types the client offers as a landing page are accepted - the rooms  list, My documents, shared with me, favorites, recent, forms and the AI agents folder - and anything else is  refused. My documents is refused for a guest as well, since a guest has no personal storage.
	DefaultFolderType FolderType `json:"defaultFolderType"`
}

type _DefaultProductRequestDto DefaultProductRequestDto

// NewDefaultProductRequestDto instantiates a new DefaultProductRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultProductRequestDto(defaultFolderType FolderType) *DefaultProductRequestDto {
	this := DefaultProductRequestDto{}
	this.DefaultFolderType = defaultFolderType
	return &this
}

// NewDefaultProductRequestDtoWithDefaults instantiates a new DefaultProductRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultProductRequestDtoWithDefaults() *DefaultProductRequestDto {
	this := DefaultProductRequestDto{}
	return &this
}

// GetDefaultFolderType returns the DefaultFolderType field value
func (o *DefaultProductRequestDto) GetDefaultFolderType() FolderType {
	if o == nil {
		var ret FolderType
		return ret
	}

	return o.DefaultFolderType
}

// GetDefaultFolderTypeOk returns a tuple with the DefaultFolderType field value
// and a boolean to check if the value has been set.
func (o *DefaultProductRequestDto) GetDefaultFolderTypeOk() (*FolderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultFolderType, true
}

// SetDefaultFolderType sets field value
func (o *DefaultProductRequestDto) SetDefaultFolderType(v FolderType) {
	o.DefaultFolderType = v
}

func (o DefaultProductRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultProductRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultFolderType"] = o.DefaultFolderType
	return toSerialize, nil
}

func (o *DefaultProductRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultFolderType",
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

	varDefaultProductRequestDto := _DefaultProductRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultProductRequestDto)

	if err != nil {
		return err
	}

	*o = DefaultProductRequestDto(varDefaultProductRequestDto)

	return err
}

type NullableDefaultProductRequestDto struct {
	value *DefaultProductRequestDto
	isSet bool
}

func (v NullableDefaultProductRequestDto) Get() *DefaultProductRequestDto {
	return v.value
}

func (v *NullableDefaultProductRequestDto) Set(val *DefaultProductRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultProductRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultProductRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultProductRequestDto(val *DefaultProductRequestDto) *NullableDefaultProductRequestDto {
	return &NullableDefaultProductRequestDto{value: val, isSet: true}
}

func (v NullableDefaultProductRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultProductRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

