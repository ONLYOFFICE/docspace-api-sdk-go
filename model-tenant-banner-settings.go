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

// checks if the TenantBannerSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantBannerSettings{}

// TenantBannerSettings The visibility settings of the promotional banners.
type TenantBannerSettings struct {
	// The banners visibility flag.
	Hidden *bool `json:"hidden,omitempty"`
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantBannerSettings instantiates a new TenantBannerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantBannerSettings() *TenantBannerSettings {
	this := TenantBannerSettings{}
	return &this
}

// NewTenantBannerSettingsWithDefaults instantiates a new TenantBannerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantBannerSettingsWithDefaults() *TenantBannerSettings {
	this := TenantBannerSettings{}
	return &this
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *TenantBannerSettings) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantBannerSettings) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *TenantBannerSettings) IsHiddenSet() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *TenantBannerSettings) SetHidden(v bool) {
	o.Hidden = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantBannerSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantBannerSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantBannerSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantBannerSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantBannerSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantBannerSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantBannerSettings struct {
	value *TenantBannerSettings
	isSet bool
}

func (v NullableTenantBannerSettings) Get() *TenantBannerSettings {
	return v.value
}

func (v *NullableTenantBannerSettings) Set(val *TenantBannerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantBannerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantBannerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantBannerSettings(val *TenantBannerSettings) *NullableTenantBannerSettings {
	return &NullableTenantBannerSettings{value: val, isSet: true}
}

func (v NullableTenantBannerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantBannerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

