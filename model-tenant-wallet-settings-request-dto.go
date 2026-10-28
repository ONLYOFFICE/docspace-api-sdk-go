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

// checks if the TenantWalletSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantWalletSettingsRequestDto{}

// TenantWalletSettingsRequestDto The body of an automatic top-up settings change.
type TenantWalletSettingsRequestDto struct {
	// The automatic top-up settings to store.
	Settings *SetWalletTopUpSettingsRequest `json:"settings,omitempty"`
}

// NewTenantWalletSettingsRequestDto instantiates a new TenantWalletSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantWalletSettingsRequestDto() *TenantWalletSettingsRequestDto {
	this := TenantWalletSettingsRequestDto{}
	return &this
}

// NewTenantWalletSettingsRequestDtoWithDefaults instantiates a new TenantWalletSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWalletSettingsRequestDtoWithDefaults() *TenantWalletSettingsRequestDto {
	this := TenantWalletSettingsRequestDto{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TenantWalletSettingsRequestDto) GetSettings() SetWalletTopUpSettingsRequest {
	if o == nil || IsNil(o.Settings) {
		var ret SetWalletTopUpSettingsRequest
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletSettingsRequestDto) GetSettingsOk() (*SetWalletTopUpSettingsRequest, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TenantWalletSettingsRequestDto) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SetWalletTopUpSettingsRequest and assigns it to the Settings field.
func (o *TenantWalletSettingsRequestDto) SetSettings(v SetWalletTopUpSettingsRequest) {
	o.Settings = &v
}

func (o TenantWalletSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantWalletSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableTenantWalletSettingsRequestDto struct {
	value *TenantWalletSettingsRequestDto
	isSet bool
}

func (v NullableTenantWalletSettingsRequestDto) Get() *TenantWalletSettingsRequestDto {
	return v.value
}

func (v *NullableTenantWalletSettingsRequestDto) Set(val *TenantWalletSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWalletSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWalletSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWalletSettingsRequestDto(val *TenantWalletSettingsRequestDto) *NullableTenantWalletSettingsRequestDto {
	return &NullableTenantWalletSettingsRequestDto{value: val, isSet: true}
}

func (v NullableTenantWalletSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWalletSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

