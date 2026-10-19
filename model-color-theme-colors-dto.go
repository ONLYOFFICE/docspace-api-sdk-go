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

// checks if the ColorThemeColorsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ColorThemeColorsDto{}

// ColorThemeColorsDto A pair of colours of a theme.
type ColorThemeColorsDto struct {
	// The accent colour, as a hex code.
	Accent NullableString `json:"accent,omitempty"`
	// The button colour, as a hex code.
	Buttons NullableString `json:"buttons,omitempty"`
}

// NewColorThemeColorsDto instantiates a new ColorThemeColorsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColorThemeColorsDto() *ColorThemeColorsDto {
	this := ColorThemeColorsDto{}
	return &this
}

// NewColorThemeColorsDtoWithDefaults instantiates a new ColorThemeColorsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColorThemeColorsDtoWithDefaults() *ColorThemeColorsDto {
	this := ColorThemeColorsDto{}
	return &this
}

// GetAccent returns the Accent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ColorThemeColorsDto) GetAccent() string {
	if o == nil || IsNil(o.Accent.Get()) {
		var ret string
		return ret
	}
	return *o.Accent.Get()
}

// GetAccentOk returns a tuple with the Accent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ColorThemeColorsDto) GetAccentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accent.Get(), o.Accent.IsSet()
}

// HasAccent returns a boolean if a field has been set.
func (o *ColorThemeColorsDto) IsAccentSet() bool {
	if o != nil && o.Accent.IsSet() {
		return true
	}

	return false
}

// SetAccent gets a reference to the given NullableString and assigns it to the Accent field.
func (o *ColorThemeColorsDto) SetAccent(v string) {
	o.Accent.Set(&v)
}
// SetAccentNil sets the value for Accent to be an explicit nil
func (o *ColorThemeColorsDto) SetAccentNil() {
	o.Accent.Set(nil)
}

// UnsetAccent ensures that no value is present for Accent, not even an explicit nil
func (o *ColorThemeColorsDto) UnsetAccent() {
	o.Accent.Unset()
}

// GetButtons returns the Buttons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ColorThemeColorsDto) GetButtons() string {
	if o == nil || IsNil(o.Buttons.Get()) {
		var ret string
		return ret
	}
	return *o.Buttons.Get()
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ColorThemeColorsDto) GetButtonsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buttons.Get(), o.Buttons.IsSet()
}

// HasButtons returns a boolean if a field has been set.
func (o *ColorThemeColorsDto) IsButtonsSet() bool {
	if o != nil && o.Buttons.IsSet() {
		return true
	}

	return false
}

// SetButtons gets a reference to the given NullableString and assigns it to the Buttons field.
func (o *ColorThemeColorsDto) SetButtons(v string) {
	o.Buttons.Set(&v)
}
// SetButtonsNil sets the value for Buttons to be an explicit nil
func (o *ColorThemeColorsDto) SetButtonsNil() {
	o.Buttons.Set(nil)
}

// UnsetButtons ensures that no value is present for Buttons, not even an explicit nil
func (o *ColorThemeColorsDto) UnsetButtons() {
	o.Buttons.Unset()
}

func (o ColorThemeColorsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ColorThemeColorsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Accent.IsSet() {
		toSerialize["accent"] = o.Accent.Get()
	}
	if o.Buttons.IsSet() {
		toSerialize["buttons"] = o.Buttons.Get()
	}
	return toSerialize, nil
}

type NullableColorThemeColorsDto struct {
	value *ColorThemeColorsDto
	isSet bool
}

func (v NullableColorThemeColorsDto) Get() *ColorThemeColorsDto {
	return v.value
}

func (v *NullableColorThemeColorsDto) Set(val *ColorThemeColorsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableColorThemeColorsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableColorThemeColorsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorThemeColorsDto(val *ColorThemeColorsDto) *NullableColorThemeColorsDto {
	return &NullableColorThemeColorsDto{value: val, isSet: true}
}

func (v NullableColorThemeColorsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColorThemeColorsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

