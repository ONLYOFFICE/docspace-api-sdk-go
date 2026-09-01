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

// checks if the CompanyWhiteLabelSettingsWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyWhiteLabelSettingsWrapper{}

// CompanyWhiteLabelSettingsWrapper The company white label settings wrapper.
type CompanyWhiteLabelSettingsWrapper struct {
	// The company white label settings.
	Settings *CompanyWhiteLabelSettings `json:"settings,omitempty"`
}

// NewCompanyWhiteLabelSettingsWrapper instantiates a new CompanyWhiteLabelSettingsWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyWhiteLabelSettingsWrapper() *CompanyWhiteLabelSettingsWrapper {
	this := CompanyWhiteLabelSettingsWrapper{}
	return &this
}

// NewCompanyWhiteLabelSettingsWrapperWithDefaults instantiates a new CompanyWhiteLabelSettingsWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWhiteLabelSettingsWrapperWithDefaults() *CompanyWhiteLabelSettingsWrapper {
	this := CompanyWhiteLabelSettingsWrapper{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CompanyWhiteLabelSettingsWrapper) GetSettings() CompanyWhiteLabelSettings {
	if o == nil || IsNil(o.Settings) {
		var ret CompanyWhiteLabelSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettingsWrapper) GetSettingsOk() (*CompanyWhiteLabelSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettingsWrapper) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CompanyWhiteLabelSettings and assigns it to the Settings field.
func (o *CompanyWhiteLabelSettingsWrapper) SetSettings(v CompanyWhiteLabelSettings) {
	o.Settings = &v
}

func (o CompanyWhiteLabelSettingsWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyWhiteLabelSettingsWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableCompanyWhiteLabelSettingsWrapper struct {
	value *CompanyWhiteLabelSettingsWrapper
	isSet bool
}

func (v NullableCompanyWhiteLabelSettingsWrapper) Get() *CompanyWhiteLabelSettingsWrapper {
	return v.value
}

func (v *NullableCompanyWhiteLabelSettingsWrapper) Set(val *CompanyWhiteLabelSettingsWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyWhiteLabelSettingsWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyWhiteLabelSettingsWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyWhiteLabelSettingsWrapper(val *CompanyWhiteLabelSettingsWrapper) *NullableCompanyWhiteLabelSettingsWrapper {
	return &NullableCompanyWhiteLabelSettingsWrapper{value: val, isSet: true}
}

func (v NullableCompanyWhiteLabelSettingsWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyWhiteLabelSettingsWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

