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

// checks if the DefaultTemplateSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultTemplateSettingsDto{}

// DefaultTemplateSettingsDto The blank document the portal creates for each extension it covers.
type DefaultTemplateSettingsDto struct {
	// One entry per extension the portal's built-in template set covers, whether or not a custom blank has been  chosen for it, so the list is never empty and its length follows the template set rather than the number of  custom blanks. Entries come in the order an interface shows them: text document, spreadsheet, presentation and  PDF first, everything else by extension.
	Items []DefaultTemplateItemDto `json:"items"`
}

type _DefaultTemplateSettingsDto DefaultTemplateSettingsDto

// NewDefaultTemplateSettingsDto instantiates a new DefaultTemplateSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultTemplateSettingsDto(items []DefaultTemplateItemDto) *DefaultTemplateSettingsDto {
	this := DefaultTemplateSettingsDto{}
	this.Items = items
	return &this
}

// NewDefaultTemplateSettingsDtoWithDefaults instantiates a new DefaultTemplateSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultTemplateSettingsDtoWithDefaults() *DefaultTemplateSettingsDto {
	this := DefaultTemplateSettingsDto{}
	return &this
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []DefaultTemplateItemDto will be returned
func (o *DefaultTemplateSettingsDto) GetItems() []DefaultTemplateItemDto {
	if o == nil {
		var ret []DefaultTemplateItemDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateSettingsDto) GetItemsOk() ([]DefaultTemplateItemDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DefaultTemplateSettingsDto) SetItems(v []DefaultTemplateItemDto) {
	o.Items = v
}

func (o DefaultTemplateSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultTemplateSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *DefaultTemplateSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varDefaultTemplateSettingsDto := _DefaultTemplateSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultTemplateSettingsDto)

	if err != nil {
		return err
	}

	*o = DefaultTemplateSettingsDto(varDefaultTemplateSettingsDto)

	return err
}

type NullableDefaultTemplateSettingsDto struct {
	value *DefaultTemplateSettingsDto
	isSet bool
}

func (v NullableDefaultTemplateSettingsDto) Get() *DefaultTemplateSettingsDto {
	return v.value
}

func (v *NullableDefaultTemplateSettingsDto) Set(val *DefaultTemplateSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultTemplateSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultTemplateSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultTemplateSettingsDto(val *DefaultTemplateSettingsDto) *NullableDefaultTemplateSettingsDto {
	return &NullableDefaultTemplateSettingsDto{value: val, isSet: true}
}

func (v NullableDefaultTemplateSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultTemplateSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

