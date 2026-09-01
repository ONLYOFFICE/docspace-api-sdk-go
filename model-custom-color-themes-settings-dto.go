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

// checks if the CustomColorThemesSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomColorThemesSettingsDto{}

// CustomColorThemesSettingsDto The custom color themes settings.
type CustomColorThemesSettingsDto struct {
	// The list of the custom color themes.
	Themes []CustomColorThemesSettingsItem `json:"themes,omitempty"`
	// Specifies whether the custom color theme is selected.
	Selected *int32 `json:"selected,omitempty"`
	// The maximum number of the custom color themes.
	Limit *int32 `json:"limit,omitempty"`
}

// NewCustomColorThemesSettingsDto instantiates a new CustomColorThemesSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorThemesSettingsDto() *CustomColorThemesSettingsDto {
	this := CustomColorThemesSettingsDto{}
	return &this
}

// NewCustomColorThemesSettingsDtoWithDefaults instantiates a new CustomColorThemesSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorThemesSettingsDtoWithDefaults() *CustomColorThemesSettingsDto {
	this := CustomColorThemesSettingsDto{}
	return &this
}

// GetThemes returns the Themes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorThemesSettingsDto) GetThemes() []CustomColorThemesSettingsItem {
	if o == nil {
		var ret []CustomColorThemesSettingsItem
		return ret
	}
	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorThemesSettingsDto) GetThemesOk() ([]CustomColorThemesSettingsItem, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsDto) IsThemesSet() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given []CustomColorThemesSettingsItem and assigns it to the Themes field.
func (o *CustomColorThemesSettingsDto) SetThemes(v []CustomColorThemesSettingsItem) {
	o.Themes = v
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsDto) GetSelected() int32 {
	if o == nil || IsNil(o.Selected) {
		var ret int32
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsDto) GetSelectedOk() (*int32, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsDto) IsSelectedSet() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given int32 and assigns it to the Selected field.
func (o *CustomColorThemesSettingsDto) SetSelected(v int32) {
	o.Selected = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsDto) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsDto) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CustomColorThemesSettingsDto) SetLimit(v int32) {
	o.Limit = &v
}

func (o CustomColorThemesSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomColorThemesSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Themes != nil {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.Selected) {
		toSerialize["selected"] = o.Selected
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableCustomColorThemesSettingsDto struct {
	value *CustomColorThemesSettingsDto
	isSet bool
}

func (v NullableCustomColorThemesSettingsDto) Get() *CustomColorThemesSettingsDto {
	return v.value
}

func (v *NullableCustomColorThemesSettingsDto) Set(val *CustomColorThemesSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorThemesSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorThemesSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorThemesSettingsDto(val *CustomColorThemesSettingsDto) *NullableCustomColorThemesSettingsDto {
	return &NullableCustomColorThemesSettingsDto{value: val, isSet: true}
}

func (v NullableCustomColorThemesSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorThemesSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

