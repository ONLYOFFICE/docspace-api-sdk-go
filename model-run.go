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

// checks if the Run type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Run{}

// Run The text run parameters.
type Run struct {
	// The fill color of the text run in RGB format.
	Fill []int32 `json:"fill,omitempty"`
	// The run text.
	Text NullableString `json:"text,omitempty"`
	// The font size of the text run in points.
	FontSize NullableString `json:"font-size,omitempty"`
}

// NewRun instantiates a new Run object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRun() *Run {
	this := Run{}
	return &this
}

// NewRunWithDefaults instantiates a new Run object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithDefaults() *Run {
	this := Run{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetFill() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetFillOk() ([]int32, bool) {
	if o == nil || IsNil(o.Fill) {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *Run) IsFillSet() bool {
	if o != nil && !IsNil(o.Fill) {
		return true
	}

	return false
}

// SetFill gets a reference to the given []int32 and assigns it to the Fill field.
func (o *Run) SetFill(v []int32) {
	o.Fill = v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *Run) IsTextSet() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *Run) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *Run) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *Run) UnsetText() {
	o.Text.Unset()
}

// GetFontSize returns the FontSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetFontSize() string {
	if o == nil || IsNil(o.FontSize.Get()) {
		var ret string
		return ret
	}
	return *o.FontSize.Get()
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetFontSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FontSize.Get(), o.FontSize.IsSet()
}

// HasFontSize returns a boolean if a field has been set.
func (o *Run) IsFontSizeSet() bool {
	if o != nil && o.FontSize.IsSet() {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given NullableString and assigns it to the FontSize field.
func (o *Run) SetFontSize(v string) {
	o.FontSize.Set(&v)
}
// SetFontSizeNil sets the value for FontSize to be an explicit nil
func (o *Run) SetFontSizeNil() {
	o.FontSize.Set(nil)
}

// UnsetFontSize ensures that no value is present for FontSize, not even an explicit nil
func (o *Run) UnsetFontSize() {
	o.FontSize.Unset()
}

func (o Run) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Run) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.FontSize.IsSet() {
		toSerialize["font-size"] = o.FontSize.Get()
	}
	return toSerialize, nil
}

type NullableRun struct {
	value *Run
	isSet bool
}

func (v NullableRun) Get() *Run {
	return v.value
}

func (v *NullableRun) Set(val *Run) {
	v.value = val
	v.isSet = true
}

func (v NullableRun) IsSet() bool {
	return v.isSet
}

func (v *NullableRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRun(val *Run) *NullableRun {
	return &NullableRun{value: val, isSet: true}
}

func (v NullableRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

