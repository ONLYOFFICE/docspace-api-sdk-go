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

// checks if the CustomColorThemesSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomColorThemesSettingsRequestDto{}

// CustomColorThemesSettingsRequestDto The custom colour theme being saved, the theme being selected, or both.
type CustomColorThemesSettingsRequestDto struct {
	// The theme to store, with its accent and button colours for the interface and for the text on it. An `id` that  matches a stored custom theme replaces it, an unknown `id` appends a new one, and an `id` belonging to a  built-in theme is treated as a request for a new custom theme rather than overwriting the built-in one. Once  the plan limit on custom themes is reached a new theme is silently not added, so compare the returned themes  against `limit` instead of assuming it was saved. Leave it out to change only the selection.
	Theme *CustomColorThemeRequestDto `json:"theme,omitempty"`
	// The theme the whole portal switches to, by theme ID. An ID matching no stored theme is ignored rather than  refused, and leaving it out keeps the selection as it is.
	Selected NullableInt32 `json:"selected,omitempty"`
}

// NewCustomColorThemesSettingsRequestDto instantiates a new CustomColorThemesSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorThemesSettingsRequestDto() *CustomColorThemesSettingsRequestDto {
	this := CustomColorThemesSettingsRequestDto{}
	return &this
}

// NewCustomColorThemesSettingsRequestDtoWithDefaults instantiates a new CustomColorThemesSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorThemesSettingsRequestDtoWithDefaults() *CustomColorThemesSettingsRequestDto {
	this := CustomColorThemesSettingsRequestDto{}
	return &this
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsRequestDto) GetTheme() CustomColorThemeRequestDto {
	if o == nil || IsNil(o.Theme) {
		var ret CustomColorThemeRequestDto
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsRequestDto) GetThemeOk() (*CustomColorThemeRequestDto, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsRequestDto) IsThemeSet() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given CustomColorThemeRequestDto and assigns it to the Theme field.
func (o *CustomColorThemesSettingsRequestDto) SetTheme(v CustomColorThemeRequestDto) {
	o.Theme = &v
}

// GetSelected returns the Selected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorThemesSettingsRequestDto) GetSelected() int32 {
	if o == nil || IsNil(o.Selected.Get()) {
		var ret int32
		return ret
	}
	return *o.Selected.Get()
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorThemesSettingsRequestDto) GetSelectedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selected.Get(), o.Selected.IsSet()
}

// HasSelected returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsRequestDto) IsSelectedSet() bool {
	if o != nil && o.Selected.IsSet() {
		return true
	}

	return false
}

// SetSelected gets a reference to the given NullableInt32 and assigns it to the Selected field.
func (o *CustomColorThemesSettingsRequestDto) SetSelected(v int32) {
	o.Selected.Set(&v)
}
// SetSelectedNil sets the value for Selected to be an explicit nil
func (o *CustomColorThemesSettingsRequestDto) SetSelectedNil() {
	o.Selected.Set(nil)
}

// UnsetSelected ensures that no value is present for Selected, not even an explicit nil
func (o *CustomColorThemesSettingsRequestDto) UnsetSelected() {
	o.Selected.Unset()
}

func (o CustomColorThemesSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomColorThemesSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if o.Selected.IsSet() {
		toSerialize["selected"] = o.Selected.Get()
	}
	return toSerialize, nil
}

type NullableCustomColorThemesSettingsRequestDto struct {
	value *CustomColorThemesSettingsRequestDto
	isSet bool
}

func (v NullableCustomColorThemesSettingsRequestDto) Get() *CustomColorThemesSettingsRequestDto {
	return v.value
}

func (v *NullableCustomColorThemesSettingsRequestDto) Set(val *CustomColorThemesSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorThemesSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorThemesSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorThemesSettingsRequestDto(val *CustomColorThemesSettingsRequestDto) *NullableCustomColorThemesSettingsRequestDto {
	return &NullableCustomColorThemesSettingsRequestDto{value: val, isSet: true}
}

func (v NullableCustomColorThemesSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorThemesSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

