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

// checks if the SsoSigningAlgorithmTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSigningAlgorithmTypeDto{}

// SsoSigningAlgorithmTypeDto The signing algorithms the SSO settings accept.
type SsoSigningAlgorithmTypeDto struct {
	// The RSA-SHA1 signing algorithm, which the built-in configuration uses. SHA-1 is the weakest of the three  and some identity providers no longer accept it.
	RsaSha1 NullableString `json:"rsaSha1,omitempty"`
	// The RSA-SHA256 signing algorithm.
	RsaSha256 NullableString `json:"rsaSha256,omitempty"`
	// The RSA-SHA512 signing algorithm.
	RsaSha512 NullableString `json:"rsaSha512,omitempty"`
}

// NewSsoSigningAlgorithmTypeDto instantiates a new SsoSigningAlgorithmTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSigningAlgorithmTypeDto() *SsoSigningAlgorithmTypeDto {
	this := SsoSigningAlgorithmTypeDto{}
	return &this
}

// NewSsoSigningAlgorithmTypeDtoWithDefaults instantiates a new SsoSigningAlgorithmTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSigningAlgorithmTypeDtoWithDefaults() *SsoSigningAlgorithmTypeDto {
	this := SsoSigningAlgorithmTypeDto{}
	return &this
}

// GetRsaSha1 returns the RsaSha1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha1() string {
	if o == nil || IsNil(o.RsaSha1.Get()) {
		var ret string
		return ret
	}
	return *o.RsaSha1.Get()
}

// GetRsaSha1Ok returns a tuple with the RsaSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RsaSha1.Get(), o.RsaSha1.IsSet()
}

// HasRsaSha1 returns a boolean if a field has been set.
func (o *SsoSigningAlgorithmTypeDto) IsRsaSha1Set() bool {
	if o != nil && o.RsaSha1.IsSet() {
		return true
	}

	return false
}

// SetRsaSha1 gets a reference to the given NullableString and assigns it to the RsaSha1 field.
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha1(v string) {
	o.RsaSha1.Set(&v)
}
// SetRsaSha1Nil sets the value for RsaSha1 to be an explicit nil
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha1Nil() {
	o.RsaSha1.Set(nil)
}

// UnsetRsaSha1 ensures that no value is present for RsaSha1, not even an explicit nil
func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha1() {
	o.RsaSha1.Unset()
}

// GetRsaSha256 returns the RsaSha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha256() string {
	if o == nil || IsNil(o.RsaSha256.Get()) {
		var ret string
		return ret
	}
	return *o.RsaSha256.Get()
}

// GetRsaSha256Ok returns a tuple with the RsaSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RsaSha256.Get(), o.RsaSha256.IsSet()
}

// HasRsaSha256 returns a boolean if a field has been set.
func (o *SsoSigningAlgorithmTypeDto) IsRsaSha256Set() bool {
	if o != nil && o.RsaSha256.IsSet() {
		return true
	}

	return false
}

// SetRsaSha256 gets a reference to the given NullableString and assigns it to the RsaSha256 field.
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha256(v string) {
	o.RsaSha256.Set(&v)
}
// SetRsaSha256Nil sets the value for RsaSha256 to be an explicit nil
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha256Nil() {
	o.RsaSha256.Set(nil)
}

// UnsetRsaSha256 ensures that no value is present for RsaSha256, not even an explicit nil
func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha256() {
	o.RsaSha256.Unset()
}

// GetRsaSha512 returns the RsaSha512 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha512() string {
	if o == nil || IsNil(o.RsaSha512.Get()) {
		var ret string
		return ret
	}
	return *o.RsaSha512.Get()
}

// GetRsaSha512Ok returns a tuple with the RsaSha512 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSigningAlgorithmTypeDto) GetRsaSha512Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RsaSha512.Get(), o.RsaSha512.IsSet()
}

// HasRsaSha512 returns a boolean if a field has been set.
func (o *SsoSigningAlgorithmTypeDto) IsRsaSha512Set() bool {
	if o != nil && o.RsaSha512.IsSet() {
		return true
	}

	return false
}

// SetRsaSha512 gets a reference to the given NullableString and assigns it to the RsaSha512 field.
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha512(v string) {
	o.RsaSha512.Set(&v)
}
// SetRsaSha512Nil sets the value for RsaSha512 to be an explicit nil
func (o *SsoSigningAlgorithmTypeDto) SetRsaSha512Nil() {
	o.RsaSha512.Set(nil)
}

// UnsetRsaSha512 ensures that no value is present for RsaSha512, not even an explicit nil
func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha512() {
	o.RsaSha512.Unset()
}

func (o SsoSigningAlgorithmTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSigningAlgorithmTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RsaSha1.IsSet() {
		toSerialize["rsaSha1"] = o.RsaSha1.Get()
	}
	if o.RsaSha256.IsSet() {
		toSerialize["rsaSha256"] = o.RsaSha256.Get()
	}
	if o.RsaSha512.IsSet() {
		toSerialize["rsaSha512"] = o.RsaSha512.Get()
	}
	return toSerialize, nil
}

type NullableSsoSigningAlgorithmTypeDto struct {
	value *SsoSigningAlgorithmTypeDto
	isSet bool
}

func (v NullableSsoSigningAlgorithmTypeDto) Get() *SsoSigningAlgorithmTypeDto {
	return v.value
}

func (v *NullableSsoSigningAlgorithmTypeDto) Set(val *SsoSigningAlgorithmTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSigningAlgorithmTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSigningAlgorithmTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSigningAlgorithmTypeDto(val *SsoSigningAlgorithmTypeDto) *NullableSsoSigningAlgorithmTypeDto {
	return &NullableSsoSigningAlgorithmTypeDto{value: val, isSet: true}
}

func (v NullableSsoSigningAlgorithmTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSigningAlgorithmTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

