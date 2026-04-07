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

// checks if the DarkThemeSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DarkThemeSettingsRequestDto{}

// DarkThemeSettingsRequestDto The theme settings request parameters.
type DarkThemeSettingsRequestDto struct {
	Theme DarkThemeSettingsType `json:"theme"`
}

type _DarkThemeSettingsRequestDto DarkThemeSettingsRequestDto

// NewDarkThemeSettingsRequestDto instantiates a new DarkThemeSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDarkThemeSettingsRequestDto(theme DarkThemeSettingsType) *DarkThemeSettingsRequestDto {
	this := DarkThemeSettingsRequestDto{}
	this.Theme = theme
	return &this
}

// NewDarkThemeSettingsRequestDtoWithDefaults instantiates a new DarkThemeSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDarkThemeSettingsRequestDtoWithDefaults() *DarkThemeSettingsRequestDto {
	this := DarkThemeSettingsRequestDto{}
	return &this
}

// GetTheme returns the Theme field value
func (o *DarkThemeSettingsRequestDto) GetTheme() DarkThemeSettingsType {
	if o == nil {
		var ret DarkThemeSettingsType
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *DarkThemeSettingsRequestDto) GetThemeOk() (*DarkThemeSettingsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *DarkThemeSettingsRequestDto) SetTheme(v DarkThemeSettingsType) {
	o.Theme = v
}

func (o DarkThemeSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DarkThemeSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["theme"] = o.Theme
	return toSerialize, nil
}

func (o *DarkThemeSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"theme",
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

	varDarkThemeSettingsRequestDto := _DarkThemeSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDarkThemeSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = DarkThemeSettingsRequestDto(varDarkThemeSettingsRequestDto)

	return err
}

type NullableDarkThemeSettingsRequestDto struct {
	value *DarkThemeSettingsRequestDto
	isSet bool
}

func (v NullableDarkThemeSettingsRequestDto) Get() *DarkThemeSettingsRequestDto {
	return v.value
}

func (v *NullableDarkThemeSettingsRequestDto) Set(val *DarkThemeSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDarkThemeSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDarkThemeSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDarkThemeSettingsRequestDto(val *DarkThemeSettingsRequestDto) *NullableDarkThemeSettingsRequestDto {
	return &NullableDarkThemeSettingsRequestDto{value: val, isSet: true}
}

func (v NullableDarkThemeSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDarkThemeSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

