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

// checks if the SsoEncryptAlgorithmTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoEncryptAlgorithmTypeDto{}

// SsoEncryptAlgorithmTypeDto The encryption algorithms the SSO settings accept.
type SsoEncryptAlgorithmTypeDto struct {
	// The AES-128-CBC encryption algorithm, which the built-in configuration uses.
	Aes128 NullableString `json:"aes128,omitempty"`
	// The AES-256-CBC encryption algorithm, the strongest of the three.
	Aes256 NullableString `json:"aes256,omitempty"`
	// The Triple DES CBC encryption algorithm, kept for identity providers that support nothing newer.
	TriDec NullableString `json:"triDec,omitempty"`
}

// NewSsoEncryptAlgorithmTypeDto instantiates a new SsoEncryptAlgorithmTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoEncryptAlgorithmTypeDto() *SsoEncryptAlgorithmTypeDto {
	this := SsoEncryptAlgorithmTypeDto{}
	return &this
}

// NewSsoEncryptAlgorithmTypeDtoWithDefaults instantiates a new SsoEncryptAlgorithmTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoEncryptAlgorithmTypeDtoWithDefaults() *SsoEncryptAlgorithmTypeDto {
	this := SsoEncryptAlgorithmTypeDto{}
	return &this
}

// GetAes128 returns the Aes128 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoEncryptAlgorithmTypeDto) GetAes128() string {
	if o == nil || IsNil(o.Aes128.Get()) {
		var ret string
		return ret
	}
	return *o.Aes128.Get()
}

// GetAes128Ok returns a tuple with the Aes128 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoEncryptAlgorithmTypeDto) GetAes128Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aes128.Get(), o.Aes128.IsSet()
}

// HasAes128 returns a boolean if a field has been set.
func (o *SsoEncryptAlgorithmTypeDto) IsAes128Set() bool {
	if o != nil && o.Aes128.IsSet() {
		return true
	}

	return false
}

// SetAes128 gets a reference to the given NullableString and assigns it to the Aes128 field.
func (o *SsoEncryptAlgorithmTypeDto) SetAes128(v string) {
	o.Aes128.Set(&v)
}
// SetAes128Nil sets the value for Aes128 to be an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) SetAes128Nil() {
	o.Aes128.Set(nil)
}

// UnsetAes128 ensures that no value is present for Aes128, not even an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) UnsetAes128() {
	o.Aes128.Unset()
}

// GetAes256 returns the Aes256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoEncryptAlgorithmTypeDto) GetAes256() string {
	if o == nil || IsNil(o.Aes256.Get()) {
		var ret string
		return ret
	}
	return *o.Aes256.Get()
}

// GetAes256Ok returns a tuple with the Aes256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoEncryptAlgorithmTypeDto) GetAes256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aes256.Get(), o.Aes256.IsSet()
}

// HasAes256 returns a boolean if a field has been set.
func (o *SsoEncryptAlgorithmTypeDto) IsAes256Set() bool {
	if o != nil && o.Aes256.IsSet() {
		return true
	}

	return false
}

// SetAes256 gets a reference to the given NullableString and assigns it to the Aes256 field.
func (o *SsoEncryptAlgorithmTypeDto) SetAes256(v string) {
	o.Aes256.Set(&v)
}
// SetAes256Nil sets the value for Aes256 to be an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) SetAes256Nil() {
	o.Aes256.Set(nil)
}

// UnsetAes256 ensures that no value is present for Aes256, not even an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) UnsetAes256() {
	o.Aes256.Unset()
}

// GetTriDec returns the TriDec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoEncryptAlgorithmTypeDto) GetTriDec() string {
	if o == nil || IsNil(o.TriDec.Get()) {
		var ret string
		return ret
	}
	return *o.TriDec.Get()
}

// GetTriDecOk returns a tuple with the TriDec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoEncryptAlgorithmTypeDto) GetTriDecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriDec.Get(), o.TriDec.IsSet()
}

// HasTriDec returns a boolean if a field has been set.
func (o *SsoEncryptAlgorithmTypeDto) IsTriDecSet() bool {
	if o != nil && o.TriDec.IsSet() {
		return true
	}

	return false
}

// SetTriDec gets a reference to the given NullableString and assigns it to the TriDec field.
func (o *SsoEncryptAlgorithmTypeDto) SetTriDec(v string) {
	o.TriDec.Set(&v)
}
// SetTriDecNil sets the value for TriDec to be an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) SetTriDecNil() {
	o.TriDec.Set(nil)
}

// UnsetTriDec ensures that no value is present for TriDec, not even an explicit nil
func (o *SsoEncryptAlgorithmTypeDto) UnsetTriDec() {
	o.TriDec.Unset()
}

func (o SsoEncryptAlgorithmTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoEncryptAlgorithmTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Aes128.IsSet() {
		toSerialize["aes128"] = o.Aes128.Get()
	}
	if o.Aes256.IsSet() {
		toSerialize["aes256"] = o.Aes256.Get()
	}
	if o.TriDec.IsSet() {
		toSerialize["triDec"] = o.TriDec.Get()
	}
	return toSerialize, nil
}

type NullableSsoEncryptAlgorithmTypeDto struct {
	value *SsoEncryptAlgorithmTypeDto
	isSet bool
}

func (v NullableSsoEncryptAlgorithmTypeDto) Get() *SsoEncryptAlgorithmTypeDto {
	return v.value
}

func (v *NullableSsoEncryptAlgorithmTypeDto) Set(val *SsoEncryptAlgorithmTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoEncryptAlgorithmTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoEncryptAlgorithmTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoEncryptAlgorithmTypeDto(val *SsoEncryptAlgorithmTypeDto) *NullableSsoEncryptAlgorithmTypeDto {
	return &NullableSsoEncryptAlgorithmTypeDto{value: val, isSet: true}
}

func (v NullableSsoEncryptAlgorithmTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoEncryptAlgorithmTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

