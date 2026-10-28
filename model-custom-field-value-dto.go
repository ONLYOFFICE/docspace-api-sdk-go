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
)

// checks if the CustomFieldValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldValueDto{}

// CustomFieldValueDto The custom text field of an entry: a free-form name with its value. Custom fields belong to no template, need no  assignment and are addressed by name.
type CustomFieldValueDto struct {
	// The field name.
	Name NullableString `json:"name,omitempty"`
	// The field value on the entry.
	Value NullableString `json:"value,omitempty"`
}

// NewCustomFieldValueDto instantiates a new CustomFieldValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldValueDto() *CustomFieldValueDto {
	this := CustomFieldValueDto{}
	return &this
}

// NewCustomFieldValueDtoWithDefaults instantiates a new CustomFieldValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldValueDtoWithDefaults() *CustomFieldValueDto {
	this := CustomFieldValueDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldValueDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldValueDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomFieldValueDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomFieldValueDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomFieldValueDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomFieldValueDto) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldValueDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldValueDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldValueDto) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *CustomFieldValueDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *CustomFieldValueDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *CustomFieldValueDto) UnsetValue() {
	o.Value.Unset()
}

func (o CustomFieldValueDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableCustomFieldValueDto struct {
	value *CustomFieldValueDto
	isSet bool
}

func (v NullableCustomFieldValueDto) Get() *CustomFieldValueDto {
	return v.value
}

func (v *NullableCustomFieldValueDto) Set(val *CustomFieldValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldValueDto(val *CustomFieldValueDto) *NullableCustomFieldValueDto {
	return &NullableCustomFieldValueDto{value: val, isSet: true}
}

func (v NullableCustomFieldValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

