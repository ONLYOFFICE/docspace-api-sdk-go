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

// checks if the CompanyWhiteLabelSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyWhiteLabelSettingsRequestDto{}

// CompanyWhiteLabelSettingsRequestDto The body of a company details change for the installation's branding.
type CompanyWhiteLabelSettingsRequestDto struct {
	// The company details to store.
	Settings *SaveCompanyInfoRequest `json:"settings,omitempty"`
}

// NewCompanyWhiteLabelSettingsRequestDto instantiates a new CompanyWhiteLabelSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyWhiteLabelSettingsRequestDto() *CompanyWhiteLabelSettingsRequestDto {
	this := CompanyWhiteLabelSettingsRequestDto{}
	return &this
}

// NewCompanyWhiteLabelSettingsRequestDtoWithDefaults instantiates a new CompanyWhiteLabelSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWhiteLabelSettingsRequestDtoWithDefaults() *CompanyWhiteLabelSettingsRequestDto {
	this := CompanyWhiteLabelSettingsRequestDto{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CompanyWhiteLabelSettingsRequestDto) GetSettings() SaveCompanyInfoRequest {
	if o == nil || IsNil(o.Settings) {
		var ret SaveCompanyInfoRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyWhiteLabelSettingsRequestDto) GetSettingsOk() (*SaveCompanyInfoRequest, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CompanyWhiteLabelSettingsRequestDto) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SaveCompanyInfoRequest and assigns it to the Settings field.
func (o *CompanyWhiteLabelSettingsRequestDto) SetSettings(v SaveCompanyInfoRequest) {
	o.Settings = &v
}

func (o CompanyWhiteLabelSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyWhiteLabelSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableCompanyWhiteLabelSettingsRequestDto struct {
	value *CompanyWhiteLabelSettingsRequestDto
	isSet bool
}

func (v NullableCompanyWhiteLabelSettingsRequestDto) Get() *CompanyWhiteLabelSettingsRequestDto {
	return v.value
}

func (v *NullableCompanyWhiteLabelSettingsRequestDto) Set(val *CompanyWhiteLabelSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyWhiteLabelSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyWhiteLabelSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyWhiteLabelSettingsRequestDto(val *CompanyWhiteLabelSettingsRequestDto) *NullableCompanyWhiteLabelSettingsRequestDto {
	return &NullableCompanyWhiteLabelSettingsRequestDto{value: val, isSet: true}
}

func (v NullableCompanyWhiteLabelSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyWhiteLabelSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

