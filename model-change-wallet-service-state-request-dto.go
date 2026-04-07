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

// checks if the ChangeWalletServiceStateRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeWalletServiceStateRequestDto{}

// ChangeWalletServiceStateRequestDto The request parameters for changing the tenant wallet service state.
type ChangeWalletServiceStateRequestDto struct {
	Service *TenantWalletService `json:"service,omitempty"`
	// Specifies whether the wallet service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewChangeWalletServiceStateRequestDto instantiates a new ChangeWalletServiceStateRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeWalletServiceStateRequestDto() *ChangeWalletServiceStateRequestDto {
	this := ChangeWalletServiceStateRequestDto{}
	return &this
}

// NewChangeWalletServiceStateRequestDtoWithDefaults instantiates a new ChangeWalletServiceStateRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWalletServiceStateRequestDtoWithDefaults() *ChangeWalletServiceStateRequestDto {
	this := ChangeWalletServiceStateRequestDto{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ChangeWalletServiceStateRequestDto) GetService() TenantWalletService {
	if o == nil || IsNil(o.Service) {
		var ret TenantWalletService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWalletServiceStateRequestDto) GetServiceOk() (*TenantWalletService, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ChangeWalletServiceStateRequestDto) IsServiceSet() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given TenantWalletService and assigns it to the Service field.
func (o *ChangeWalletServiceStateRequestDto) SetService(v TenantWalletService) {
	o.Service = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ChangeWalletServiceStateRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeWalletServiceStateRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ChangeWalletServiceStateRequestDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ChangeWalletServiceStateRequestDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ChangeWalletServiceStateRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeWalletServiceStateRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableChangeWalletServiceStateRequestDto struct {
	value *ChangeWalletServiceStateRequestDto
	isSet bool
}

func (v NullableChangeWalletServiceStateRequestDto) Get() *ChangeWalletServiceStateRequestDto {
	return v.value
}

func (v *NullableChangeWalletServiceStateRequestDto) Set(val *ChangeWalletServiceStateRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeWalletServiceStateRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeWalletServiceStateRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeWalletServiceStateRequestDto(val *ChangeWalletServiceStateRequestDto) *NullableChangeWalletServiceStateRequestDto {
	return &NullableChangeWalletServiceStateRequestDto{value: val, isSet: true}
}

func (v NullableChangeWalletServiceStateRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeWalletServiceStateRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

