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

// checks if the EncryptionKeysConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionKeysConfig{}

// EncryptionKeysConfig The encryption keys of the editor configuration.
type EncryptionKeysConfig struct {
	// The crypto engine ID of the encryption key.
	CryptoEngineId NullableString `json:"cryptoEngineId,omitempty"`
	// The private key.
	PrivateKeyEnc NullableString `json:"privateKeyEnc,omitempty"`
	// The public key.
	PublicKey NullableString `json:"publicKey,omitempty"`
}

// NewEncryptionKeysConfig instantiates a new EncryptionKeysConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionKeysConfig() *EncryptionKeysConfig {
	this := EncryptionKeysConfig{}
	return &this
}

// NewEncryptionKeysConfigWithDefaults instantiates a new EncryptionKeysConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeysConfigWithDefaults() *EncryptionKeysConfig {
	this := EncryptionKeysConfig{}
	return &this
}

// GetCryptoEngineId returns the CryptoEngineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeysConfig) GetCryptoEngineId() string {
	if o == nil || IsNil(o.CryptoEngineId.Get()) {
		var ret string
		return ret
	}
	return *o.CryptoEngineId.Get()
}

// GetCryptoEngineIdOk returns a tuple with the CryptoEngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeysConfig) GetCryptoEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CryptoEngineId.Get(), o.CryptoEngineId.IsSet()
}

// HasCryptoEngineId returns a boolean if a field has been set.
func (o *EncryptionKeysConfig) IsCryptoEngineIdSet() bool {
	if o != nil && o.CryptoEngineId.IsSet() {
		return true
	}

	return false
}

// SetCryptoEngineId gets a reference to the given NullableString and assigns it to the CryptoEngineId field.
func (o *EncryptionKeysConfig) SetCryptoEngineId(v string) {
	o.CryptoEngineId.Set(&v)
}
// SetCryptoEngineIdNil sets the value for CryptoEngineId to be an explicit nil
func (o *EncryptionKeysConfig) SetCryptoEngineIdNil() {
	o.CryptoEngineId.Set(nil)
}

// UnsetCryptoEngineId ensures that no value is present for CryptoEngineId, not even an explicit nil
func (o *EncryptionKeysConfig) UnsetCryptoEngineId() {
	o.CryptoEngineId.Unset()
}

// GetPrivateKeyEnc returns the PrivateKeyEnc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeysConfig) GetPrivateKeyEnc() string {
	if o == nil || IsNil(o.PrivateKeyEnc.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKeyEnc.Get()
}

// GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeysConfig) GetPrivateKeyEncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyEnc.Get(), o.PrivateKeyEnc.IsSet()
}

// HasPrivateKeyEnc returns a boolean if a field has been set.
func (o *EncryptionKeysConfig) IsPrivateKeyEncSet() bool {
	if o != nil && o.PrivateKeyEnc.IsSet() {
		return true
	}

	return false
}

// SetPrivateKeyEnc gets a reference to the given NullableString and assigns it to the PrivateKeyEnc field.
func (o *EncryptionKeysConfig) SetPrivateKeyEnc(v string) {
	o.PrivateKeyEnc.Set(&v)
}
// SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil
func (o *EncryptionKeysConfig) SetPrivateKeyEncNil() {
	o.PrivateKeyEnc.Set(nil)
}

// UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
func (o *EncryptionKeysConfig) UnsetPrivateKeyEnc() {
	o.PrivateKeyEnc.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeysConfig) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeysConfig) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EncryptionKeysConfig) IsPublicKeySet() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *EncryptionKeysConfig) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *EncryptionKeysConfig) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *EncryptionKeysConfig) UnsetPublicKey() {
	o.PublicKey.Unset()
}

func (o EncryptionKeysConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionKeysConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CryptoEngineId.IsSet() {
		toSerialize["cryptoEngineId"] = o.CryptoEngineId.Get()
	}
	if o.PrivateKeyEnc.IsSet() {
		toSerialize["privateKeyEnc"] = o.PrivateKeyEnc.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["publicKey"] = o.PublicKey.Get()
	}
	return toSerialize, nil
}

type NullableEncryptionKeysConfig struct {
	value *EncryptionKeysConfig
	isSet bool
}

func (v NullableEncryptionKeysConfig) Get() *EncryptionKeysConfig {
	return v.value
}

func (v *NullableEncryptionKeysConfig) Set(val *EncryptionKeysConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionKeysConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionKeysConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionKeysConfig(val *EncryptionKeysConfig) *NullableEncryptionKeysConfig {
	return &NullableEncryptionKeysConfig{value: val, isSet: true}
}

func (v NullableEncryptionKeysConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionKeysConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

