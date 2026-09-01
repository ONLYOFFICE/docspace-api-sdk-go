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

// checks if the CustomColorThemesSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomColorThemesSettingsRequestsDto{}

// CustomColorThemesSettingsRequestsDto The request parameters for managing the portal theme settings.
type CustomColorThemesSettingsRequestsDto struct {
	// The custom color theme configuration.
	Theme *CustomColorThemesSettingsItem `json:"theme,omitempty"`
	// Specifies the optional value indicating the selected custom color theme.
	Selected NullableInt32 `json:"selected,omitempty"`
}

// NewCustomColorThemesSettingsRequestsDto instantiates a new CustomColorThemesSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorThemesSettingsRequestsDto() *CustomColorThemesSettingsRequestsDto {
	this := CustomColorThemesSettingsRequestsDto{}
	return &this
}

// NewCustomColorThemesSettingsRequestsDtoWithDefaults instantiates a new CustomColorThemesSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorThemesSettingsRequestsDtoWithDefaults() *CustomColorThemesSettingsRequestsDto {
	this := CustomColorThemesSettingsRequestsDto{}
	return &this
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsRequestsDto) GetTheme() CustomColorThemesSettingsItem {
	if o == nil || IsNil(o.Theme) {
		var ret CustomColorThemesSettingsItem
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsRequestsDto) GetThemeOk() (*CustomColorThemesSettingsItem, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsRequestsDto) IsThemeSet() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given CustomColorThemesSettingsItem and assigns it to the Theme field.
func (o *CustomColorThemesSettingsRequestsDto) SetTheme(v CustomColorThemesSettingsItem) {
	o.Theme = &v
}

// GetSelected returns the Selected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorThemesSettingsRequestsDto) GetSelected() int32 {
	if o == nil || IsNil(o.Selected.Get()) {
		var ret int32
		return ret
	}
	return *o.Selected.Get()
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorThemesSettingsRequestsDto) GetSelectedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selected.Get(), o.Selected.IsSet()
}

// HasSelected returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsRequestsDto) IsSelectedSet() bool {
	if o != nil && o.Selected.IsSet() {
		return true
	}

	return false
}

// SetSelected gets a reference to the given NullableInt32 and assigns it to the Selected field.
func (o *CustomColorThemesSettingsRequestsDto) SetSelected(v int32) {
	o.Selected.Set(&v)
}
// SetSelectedNil sets the value for Selected to be an explicit nil
func (o *CustomColorThemesSettingsRequestsDto) SetSelectedNil() {
	o.Selected.Set(nil)
}

// UnsetSelected ensures that no value is present for Selected, not even an explicit nil
func (o *CustomColorThemesSettingsRequestsDto) UnsetSelected() {
	o.Selected.Unset()
}

func (o CustomColorThemesSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomColorThemesSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if o.Selected.IsSet() {
		toSerialize["selected"] = o.Selected.Get()
	}
	return toSerialize, nil
}

type NullableCustomColorThemesSettingsRequestsDto struct {
	value *CustomColorThemesSettingsRequestsDto
	isSet bool
}

func (v NullableCustomColorThemesSettingsRequestsDto) Get() *CustomColorThemesSettingsRequestsDto {
	return v.value
}

func (v *NullableCustomColorThemesSettingsRequestsDto) Set(val *CustomColorThemesSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorThemesSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorThemesSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorThemesSettingsRequestsDto(val *CustomColorThemesSettingsRequestsDto) *NullableCustomColorThemesSettingsRequestsDto {
	return &NullableCustomColorThemesSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableCustomColorThemesSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorThemesSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

