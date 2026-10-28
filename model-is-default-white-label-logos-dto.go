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

// checks if the IsDefaultWhiteLabelLogosDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsDefaultWhiteLabelLogosDto{}

// IsDefaultWhiteLabelLogosDto Whether one branding slot still holds the built-in image or wordmark.
type IsDefaultWhiteLabelLogosDto struct {
	// The stable name of the slot, matching the `name` of the same slot in  `GET api/2.0/settings/whitelabel/logos` - `LightSmall`, `LoginPage`, `Favicon`, `DocsEditor` and the rest,  plus `Notification`, which that list leaves out. The wordmark check reports the fixed name `logotext`  instead of a slot.
	Name NullableString `json:"name"`
	// Whether the slot has never been written for this portal, in which case the built-in image is what gets  rendered. It turns `false` once an image has been stored, for either the light or the dark theme, and back  to `true` after the matching restore operation. For `logotext` it stays `true` when the built-in wordmark  itself is saved, because saving that value counts as clearing the setting.
	Default bool `json:"default"`
}

type _IsDefaultWhiteLabelLogosDto IsDefaultWhiteLabelLogosDto

// NewIsDefaultWhiteLabelLogosDto instantiates a new IsDefaultWhiteLabelLogosDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsDefaultWhiteLabelLogosDto(name NullableString, default_ bool) *IsDefaultWhiteLabelLogosDto {
	this := IsDefaultWhiteLabelLogosDto{}
	this.Name = name
	this.Default = default_
	return &this
}

// NewIsDefaultWhiteLabelLogosDtoWithDefaults instantiates a new IsDefaultWhiteLabelLogosDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsDefaultWhiteLabelLogosDtoWithDefaults() *IsDefaultWhiteLabelLogosDto {
	this := IsDefaultWhiteLabelLogosDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IsDefaultWhiteLabelLogosDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IsDefaultWhiteLabelLogosDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *IsDefaultWhiteLabelLogosDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetDefault returns the Default field value
func (o *IsDefaultWhiteLabelLogosDto) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosDto) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *IsDefaultWhiteLabelLogosDto) SetDefault(v bool) {
	o.Default = v
}

func (o IsDefaultWhiteLabelLogosDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsDefaultWhiteLabelLogosDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

func (o *IsDefaultWhiteLabelLogosDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"default",
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

	varIsDefaultWhiteLabelLogosDto := _IsDefaultWhiteLabelLogosDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIsDefaultWhiteLabelLogosDto)

	if err != nil {
		return err
	}

	*o = IsDefaultWhiteLabelLogosDto(varIsDefaultWhiteLabelLogosDto)

	return err
}

type NullableIsDefaultWhiteLabelLogosDto struct {
	value *IsDefaultWhiteLabelLogosDto
	isSet bool
}

func (v NullableIsDefaultWhiteLabelLogosDto) Get() *IsDefaultWhiteLabelLogosDto {
	return v.value
}

func (v *NullableIsDefaultWhiteLabelLogosDto) Set(val *IsDefaultWhiteLabelLogosDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIsDefaultWhiteLabelLogosDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIsDefaultWhiteLabelLogosDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsDefaultWhiteLabelLogosDto(val *IsDefaultWhiteLabelLogosDto) *NullableIsDefaultWhiteLabelLogosDto {
	return &NullableIsDefaultWhiteLabelLogosDto{value: val, isSet: true}
}

func (v NullableIsDefaultWhiteLabelLogosDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsDefaultWhiteLabelLogosDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

