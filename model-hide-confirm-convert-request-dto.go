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

// checks if the HideConfirmConvertRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HideConfirmConvertRequestDto{}

// HideConfirmConvertRequestDto The body of the conversion prompt switch: which of the two prompts to hide.
type HideConfirmConvertRequestDto struct {
	// Chooses the prompt to hide rather than the state to store: true hides the prompt that offers to keep a copy in  the original format when a document is converted, false hides the prompt that offers to open the conversion  result. Each of the two flags is stored separately for the calling account, and both are one-way - the portal  can hide a prompt but has no way to show it again.
	Save *bool `json:"save,omitempty"`
}

// NewHideConfirmConvertRequestDto instantiates a new HideConfirmConvertRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHideConfirmConvertRequestDto() *HideConfirmConvertRequestDto {
	this := HideConfirmConvertRequestDto{}
	return &this
}

// NewHideConfirmConvertRequestDtoWithDefaults instantiates a new HideConfirmConvertRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHideConfirmConvertRequestDtoWithDefaults() *HideConfirmConvertRequestDto {
	this := HideConfirmConvertRequestDto{}
	return &this
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *HideConfirmConvertRequestDto) GetSave() bool {
	if o == nil || IsNil(o.Save) {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HideConfirmConvertRequestDto) GetSaveOk() (*bool, bool) {
	if o == nil || IsNil(o.Save) {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *HideConfirmConvertRequestDto) IsSaveSet() bool {
	if o != nil && !IsNil(o.Save) {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *HideConfirmConvertRequestDto) SetSave(v bool) {
	o.Save = &v
}

func (o HideConfirmConvertRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HideConfirmConvertRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Save) {
		toSerialize["save"] = o.Save
	}
	return toSerialize, nil
}

type NullableHideConfirmConvertRequestDto struct {
	value *HideConfirmConvertRequestDto
	isSet bool
}

func (v NullableHideConfirmConvertRequestDto) Get() *HideConfirmConvertRequestDto {
	return v.value
}

func (v *NullableHideConfirmConvertRequestDto) Set(val *HideConfirmConvertRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHideConfirmConvertRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHideConfirmConvertRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHideConfirmConvertRequestDto(val *HideConfirmConvertRequestDto) *NullableHideConfirmConvertRequestDto {
	return &NullableHideConfirmConvertRequestDto{value: val, isSet: true}
}

func (v NullableHideConfirmConvertRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHideConfirmConvertRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

