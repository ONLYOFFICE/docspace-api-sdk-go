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

// checks if the SubmitFormDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitFormDto{}

// SubmitFormDto The Complete & Submit button settings.
type SubmitFormDto struct {
	// Specifies whether the Complete  & Submit button will be displayed or hidden on the top toolbar.
	Visible *bool `json:"visible,omitempty"`
	// A message displayed after forms are submitted.
	ResultMessage NullableString `json:"resultMessage,omitempty"`
}

// NewSubmitFormDto instantiates a new SubmitFormDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitFormDto() *SubmitFormDto {
	this := SubmitFormDto{}
	return &this
}

// NewSubmitFormDtoWithDefaults instantiates a new SubmitFormDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitFormDtoWithDefaults() *SubmitFormDto {
	this := SubmitFormDto{}
	return &this
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *SubmitFormDto) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitFormDto) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *SubmitFormDto) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *SubmitFormDto) SetVisible(v bool) {
	o.Visible = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubmitFormDto) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ResultMessage.Get()
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubmitFormDto) GetResultMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultMessage.Get(), o.ResultMessage.IsSet()
}

// HasResultMessage returns a boolean if a field has been set.
func (o *SubmitFormDto) IsResultMessageSet() bool {
	if o != nil && o.ResultMessage.IsSet() {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given NullableString and assigns it to the ResultMessage field.
func (o *SubmitFormDto) SetResultMessage(v string) {
	o.ResultMessage.Set(&v)
}
// SetResultMessageNil sets the value for ResultMessage to be an explicit nil
func (o *SubmitFormDto) SetResultMessageNil() {
	o.ResultMessage.Set(nil)
}

// UnsetResultMessage ensures that no value is present for ResultMessage, not even an explicit nil
func (o *SubmitFormDto) UnsetResultMessage() {
	o.ResultMessage.Unset()
}

func (o SubmitFormDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitFormDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if o.ResultMessage.IsSet() {
		toSerialize["resultMessage"] = o.ResultMessage.Get()
	}
	return toSerialize, nil
}

type NullableSubmitFormDto struct {
	value *SubmitFormDto
	isSet bool
}

func (v NullableSubmitFormDto) Get() *SubmitFormDto {
	return v.value
}

func (v *NullableSubmitFormDto) Set(val *SubmitFormDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitFormDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitFormDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitFormDto(val *SubmitFormDto) *NullableSubmitFormDto {
	return &NullableSubmitFormDto{value: val, isSet: true}
}

func (v NullableSubmitFormDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitFormDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

