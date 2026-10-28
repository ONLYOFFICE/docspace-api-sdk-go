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

// checks if the StartFillingForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartFillingForm{}

// StartFillingForm The button the editor shows to begin filling out a form.
type StartFillingForm struct {
	// The caption to put on the button, already translated into the language of the caller.
	Text NullableString `json:"text,omitempty"`
}

// NewStartFillingForm instantiates a new StartFillingForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartFillingForm() *StartFillingForm {
	this := StartFillingForm{}
	return &this
}

// NewStartFillingFormWithDefaults instantiates a new StartFillingForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartFillingFormWithDefaults() *StartFillingForm {
	this := StartFillingForm{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartFillingForm) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartFillingForm) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *StartFillingForm) IsTextSet() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *StartFillingForm) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *StartFillingForm) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *StartFillingForm) UnsetText() {
	o.Text.Unset()
}

func (o StartFillingForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartFillingForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	return toSerialize, nil
}

type NullableStartFillingForm struct {
	value *StartFillingForm
	isSet bool
}

func (v NullableStartFillingForm) Get() *StartFillingForm {
	return v.value
}

func (v *NullableStartFillingForm) Set(val *StartFillingForm) {
	v.value = val
	v.isSet = true
}

func (v NullableStartFillingForm) IsSet() bool {
	return v.isSet
}

func (v *NullableStartFillingForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartFillingForm(val *StartFillingForm) *NullableStartFillingForm {
	return &NullableStartFillingForm{value: val, isSet: true}
}

func (v NullableStartFillingForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartFillingForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

