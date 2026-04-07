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

// checks if the SettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsRequestDto{}

// SettingsRequestDto The settings request parameters.
type SettingsRequestDto struct {
	// Specifies whether to set the specified settings or not.
	Set *bool `json:"set,omitempty"`
}

// NewSettingsRequestDto instantiates a new SettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRequestDto() *SettingsRequestDto {
	this := SettingsRequestDto{}
	return &this
}

// NewSettingsRequestDtoWithDefaults instantiates a new SettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRequestDtoWithDefaults() *SettingsRequestDto {
	this := SettingsRequestDto{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SettingsRequestDto) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequestDto) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SettingsRequestDto) IsSetSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *SettingsRequestDto) SetSet(v bool) {
	o.Set = &v
}

func (o SettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	return toSerialize, nil
}

type NullableSettingsRequestDto struct {
	value *SettingsRequestDto
	isSet bool
}

func (v NullableSettingsRequestDto) Get() *SettingsRequestDto {
	return v.value
}

func (v *NullableSettingsRequestDto) Set(val *SettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRequestDto(val *SettingsRequestDto) *NullableSettingsRequestDto {
	return &NullableSettingsRequestDto{value: val, isSet: true}
}

func (v NullableSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

