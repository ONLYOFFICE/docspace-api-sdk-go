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

// checks if the EncryptionKeyRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionKeyRequestDto{}

// EncryptionKeyRequestDto The request parameters for storing the encryption key pair of a user.
type EncryptionKeyRequestDto struct {
	// The identifier of the key pair.
	Id *string `json:"id,omitempty"`
	// The public key of the pair, used to encrypt the file keys.
	PublicKey NullableString `json:"publicKey,omitempty"`
	// The private key of the pair, encrypted with the user password.
	PrivateKeyEnc NullableString `json:"privateKeyEnc,omitempty"`
}

// NewEncryptionKeyRequestDto instantiates a new EncryptionKeyRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionKeyRequestDto() *EncryptionKeyRequestDto {
	this := EncryptionKeyRequestDto{}
	return &this
}

// NewEncryptionKeyRequestDtoWithDefaults instantiates a new EncryptionKeyRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeyRequestDtoWithDefaults() *EncryptionKeyRequestDto {
	this := EncryptionKeyRequestDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EncryptionKeyRequestDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionKeyRequestDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EncryptionKeyRequestDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EncryptionKeyRequestDto) SetId(v string) {
	o.Id = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeyRequestDto) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeyRequestDto) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EncryptionKeyRequestDto) IsPublicKeySet() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *EncryptionKeyRequestDto) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *EncryptionKeyRequestDto) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *EncryptionKeyRequestDto) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetPrivateKeyEnc returns the PrivateKeyEnc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EncryptionKeyRequestDto) GetPrivateKeyEnc() string {
	if o == nil || IsNil(o.PrivateKeyEnc.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKeyEnc.Get()
}

// GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EncryptionKeyRequestDto) GetPrivateKeyEncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyEnc.Get(), o.PrivateKeyEnc.IsSet()
}

// HasPrivateKeyEnc returns a boolean if a field has been set.
func (o *EncryptionKeyRequestDto) IsPrivateKeyEncSet() bool {
	if o != nil && o.PrivateKeyEnc.IsSet() {
		return true
	}

	return false
}

// SetPrivateKeyEnc gets a reference to the given NullableString and assigns it to the PrivateKeyEnc field.
func (o *EncryptionKeyRequestDto) SetPrivateKeyEnc(v string) {
	o.PrivateKeyEnc.Set(&v)
}
// SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil
func (o *EncryptionKeyRequestDto) SetPrivateKeyEncNil() {
	o.PrivateKeyEnc.Set(nil)
}

// UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
func (o *EncryptionKeyRequestDto) UnsetPrivateKeyEnc() {
	o.PrivateKeyEnc.Unset()
}

func (o EncryptionKeyRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionKeyRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.PublicKey.IsSet() {
		toSerialize["publicKey"] = o.PublicKey.Get()
	}
	if o.PrivateKeyEnc.IsSet() {
		toSerialize["privateKeyEnc"] = o.PrivateKeyEnc.Get()
	}
	return toSerialize, nil
}

type NullableEncryptionKeyRequestDto struct {
	value *EncryptionKeyRequestDto
	isSet bool
}

func (v NullableEncryptionKeyRequestDto) Get() *EncryptionKeyRequestDto {
	return v.value
}

func (v *NullableEncryptionKeyRequestDto) Set(val *EncryptionKeyRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionKeyRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionKeyRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionKeyRequestDto(val *EncryptionKeyRequestDto) *NullableEncryptionKeyRequestDto {
	return &NullableEncryptionKeyRequestDto{value: val, isSet: true}
}

func (v NullableEncryptionKeyRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionKeyRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

