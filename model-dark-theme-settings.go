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
	"time"
)

// checks if the DarkThemeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DarkThemeSettings{}

// DarkThemeSettings The theme parameters.
type DarkThemeSettings struct {
	// The theme type.
	Theme *DarkThemeSettingsType `json:"theme,omitempty"`
	// The last modified date.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewDarkThemeSettings instantiates a new DarkThemeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDarkThemeSettings() *DarkThemeSettings {
	this := DarkThemeSettings{}
	return &this
}

// NewDarkThemeSettingsWithDefaults instantiates a new DarkThemeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDarkThemeSettingsWithDefaults() *DarkThemeSettings {
	this := DarkThemeSettings{}
	return &this
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *DarkThemeSettings) GetTheme() DarkThemeSettingsType {
	if o == nil || IsNil(o.Theme) {
		var ret DarkThemeSettingsType
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarkThemeSettings) GetThemeOk() (*DarkThemeSettingsType, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *DarkThemeSettings) IsThemeSet() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given DarkThemeSettingsType and assigns it to the Theme field.
func (o *DarkThemeSettings) SetTheme(v DarkThemeSettingsType) {
	o.Theme = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *DarkThemeSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DarkThemeSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *DarkThemeSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *DarkThemeSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o DarkThemeSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DarkThemeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableDarkThemeSettings struct {
	value *DarkThemeSettings
	isSet bool
}

func (v NullableDarkThemeSettings) Get() *DarkThemeSettings {
	return v.value
}

func (v *NullableDarkThemeSettings) Set(val *DarkThemeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDarkThemeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDarkThemeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDarkThemeSettings(val *DarkThemeSettings) *NullableDarkThemeSettings {
	return &NullableDarkThemeSettings{value: val, isSet: true}
}

func (v NullableDarkThemeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDarkThemeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

