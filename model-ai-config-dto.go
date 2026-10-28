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

// checks if the AiConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiConfigDto{}

// AiConfigDto The AI settings the editor opens with. They follow the portal-wide AI access setting, so every document of the  portal gets the same value.
type AiConfigDto struct {
	// Indicates whether the AI feature is disabled.
	Disabled *bool `json:"disabled,omitempty"`
}

// NewAiConfigDto instantiates a new AiConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiConfigDto() *AiConfigDto {
	this := AiConfigDto{}
	return &this
}

// NewAiConfigDtoWithDefaults instantiates a new AiConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiConfigDtoWithDefaults() *AiConfigDto {
	this := AiConfigDto{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AiConfigDto) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiConfigDto) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AiConfigDto) IsDisabledSet() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AiConfigDto) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o AiConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	return toSerialize, nil
}

type NullableAiConfigDto struct {
	value *AiConfigDto
	isSet bool
}

func (v NullableAiConfigDto) Get() *AiConfigDto {
	return v.value
}

func (v *NullableAiConfigDto) Set(val *AiConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiConfigDto(val *AiConfigDto) *NullableAiConfigDto {
	return &NullableAiConfigDto{value: val, isSet: true}
}

func (v NullableAiConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

