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

// checks if the TenantBannerSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantBannerSettingsDto{}

// TenantBannerSettingsDto Whether the portal promotional banners are hidden.
type TenantBannerSettingsDto struct {
	// Whether the promotional banners are hidden from every user of the portal. The flag is only honoured on a  self-hosted installation; a SaaS portal keeps showing the banners whatever is stored here.
	Hidden *bool `json:"hidden,omitempty"`
}

// NewTenantBannerSettingsDto instantiates a new TenantBannerSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantBannerSettingsDto() *TenantBannerSettingsDto {
	this := TenantBannerSettingsDto{}
	return &this
}

// NewTenantBannerSettingsDtoWithDefaults instantiates a new TenantBannerSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantBannerSettingsDtoWithDefaults() *TenantBannerSettingsDto {
	this := TenantBannerSettingsDto{}
	return &this
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *TenantBannerSettingsDto) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantBannerSettingsDto) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *TenantBannerSettingsDto) IsHiddenSet() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *TenantBannerSettingsDto) SetHidden(v bool) {
	o.Hidden = &v
}

func (o TenantBannerSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantBannerSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	return toSerialize, nil
}

type NullableTenantBannerSettingsDto struct {
	value *TenantBannerSettingsDto
	isSet bool
}

func (v NullableTenantBannerSettingsDto) Get() *TenantBannerSettingsDto {
	return v.value
}

func (v *NullableTenantBannerSettingsDto) Set(val *TenantBannerSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantBannerSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantBannerSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantBannerSettingsDto(val *TenantBannerSettingsDto) *NullableTenantBannerSettingsDto {
	return &NullableTenantBannerSettingsDto{value: val, isSet: true}
}

func (v NullableTenantBannerSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantBannerSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

