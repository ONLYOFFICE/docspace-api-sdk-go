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

// checks if the AiToolPermissionModeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolPermissionModeDto{}

// AiToolPermissionModeDto The current user's tool permission mode.
type AiToolPermissionModeDto struct {
	// How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing.
	Mode *AiToolPermissionMode `json:"mode,omitempty"`
}

// NewAiToolPermissionModeDto instantiates a new AiToolPermissionModeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolPermissionModeDto() *AiToolPermissionModeDto {
	this := AiToolPermissionModeDto{}
	return &this
}

// NewAiToolPermissionModeDtoWithDefaults instantiates a new AiToolPermissionModeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolPermissionModeDtoWithDefaults() *AiToolPermissionModeDto {
	this := AiToolPermissionModeDto{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AiToolPermissionModeDto) GetMode() AiToolPermissionMode {
	if o == nil || IsNil(o.Mode) {
		var ret AiToolPermissionMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolPermissionModeDto) GetModeOk() (*AiToolPermissionMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AiToolPermissionModeDto) IsModeSet() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given AiToolPermissionMode and assigns it to the Mode field.
func (o *AiToolPermissionModeDto) SetMode(v AiToolPermissionMode) {
	o.Mode = &v
}

func (o AiToolPermissionModeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolPermissionModeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableAiToolPermissionModeDto struct {
	value *AiToolPermissionModeDto
	isSet bool
}

func (v NullableAiToolPermissionModeDto) Get() *AiToolPermissionModeDto {
	return v.value
}

func (v *NullableAiToolPermissionModeDto) Set(val *AiToolPermissionModeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolPermissionModeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolPermissionModeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolPermissionModeDto(val *AiToolPermissionModeDto) *NullableAiToolPermissionModeDto {
	return &NullableAiToolPermissionModeDto{value: val, isSet: true}
}

func (v NullableAiToolPermissionModeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolPermissionModeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

