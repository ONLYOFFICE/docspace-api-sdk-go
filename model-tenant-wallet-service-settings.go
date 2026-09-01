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

// checks if the TenantWalletServiceSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantWalletServiceSettings{}

// TenantWalletServiceSettings The wallet services settings.
type TenantWalletServiceSettings struct {
	// The list of the enabled wallet services.
	EnabledServices []int32 `json:"enabledServices,omitempty"`
	// The date and time when the wallet services settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// NewTenantWalletServiceSettings instantiates a new TenantWalletServiceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantWalletServiceSettings() *TenantWalletServiceSettings {
	this := TenantWalletServiceSettings{}
	return &this
}

// NewTenantWalletServiceSettingsWithDefaults instantiates a new TenantWalletServiceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWalletServiceSettingsWithDefaults() *TenantWalletServiceSettings {
	this := TenantWalletServiceSettings{}
	return &this
}

// GetEnabledServices returns the EnabledServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantWalletServiceSettings) GetEnabledServices() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.EnabledServices
}

// GetEnabledServicesOk returns a tuple with the EnabledServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantWalletServiceSettings) GetEnabledServicesOk() ([]int32, bool) {
	if o == nil || IsNil(o.EnabledServices) {
		return nil, false
	}
	return o.EnabledServices, true
}

// HasEnabledServices returns a boolean if a field has been set.
func (o *TenantWalletServiceSettings) IsEnabledServicesSet() bool {
	if o != nil && !IsNil(o.EnabledServices) {
		return true
	}

	return false
}

// SetEnabledServices gets a reference to the given []int32 and assigns it to the EnabledServices field.
func (o *TenantWalletServiceSettings) SetEnabledServices(v []int32) {
	o.EnabledServices = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *TenantWalletServiceSettings) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWalletServiceSettings) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *TenantWalletServiceSettings) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *TenantWalletServiceSettings) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o TenantWalletServiceSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantWalletServiceSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnabledServices != nil {
		toSerialize["enabledServices"] = o.EnabledServices
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableTenantWalletServiceSettings struct {
	value *TenantWalletServiceSettings
	isSet bool
}

func (v NullableTenantWalletServiceSettings) Get() *TenantWalletServiceSettings {
	return v.value
}

func (v *NullableTenantWalletServiceSettings) Set(val *TenantWalletServiceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWalletServiceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWalletServiceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWalletServiceSettings(val *TenantWalletServiceSettings) *NullableTenantWalletServiceSettings {
	return &NullableTenantWalletServiceSettings{value: val, isSet: true}
}

func (v NullableTenantWalletServiceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWalletServiceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

