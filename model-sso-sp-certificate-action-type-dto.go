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

// checks if the SsoSpCertificateActionTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSpCertificateActionTypeDto{}

// SsoSpCertificateActionTypeDto What the portal's own key pair may be used for, as the `action` of a service provider certificate.
type SsoSpCertificateActionTypeDto struct {
	// The key pair signs the requests the portal sends and nothing else.
	Signing NullableString `json:"signing,omitempty"`
	// The key pair encrypts what the portal sends and decrypts what comes back, but signs nothing.
	Encrypt NullableString `json:"encrypt,omitempty"`
	// The key pair does both, which is what one pair configured on its own has to be set to.
	SigningAndEncrypt NullableString `json:"signingAndEncrypt,omitempty"`
}

// NewSsoSpCertificateActionTypeDto instantiates a new SsoSpCertificateActionTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSpCertificateActionTypeDto() *SsoSpCertificateActionTypeDto {
	this := SsoSpCertificateActionTypeDto{}
	return &this
}

// NewSsoSpCertificateActionTypeDtoWithDefaults instantiates a new SsoSpCertificateActionTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSpCertificateActionTypeDtoWithDefaults() *SsoSpCertificateActionTypeDto {
	this := SsoSpCertificateActionTypeDto{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateActionTypeDto) GetSigning() string {
	if o == nil || IsNil(o.Signing.Get()) {
		var ret string
		return ret
	}
	return *o.Signing.Get()
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateActionTypeDto) GetSigningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signing.Get(), o.Signing.IsSet()
}

// HasSigning returns a boolean if a field has been set.
func (o *SsoSpCertificateActionTypeDto) IsSigningSet() bool {
	if o != nil && o.Signing.IsSet() {
		return true
	}

	return false
}

// SetSigning gets a reference to the given NullableString and assigns it to the Signing field.
func (o *SsoSpCertificateActionTypeDto) SetSigning(v string) {
	o.Signing.Set(&v)
}
// SetSigningNil sets the value for Signing to be an explicit nil
func (o *SsoSpCertificateActionTypeDto) SetSigningNil() {
	o.Signing.Set(nil)
}

// UnsetSigning ensures that no value is present for Signing, not even an explicit nil
func (o *SsoSpCertificateActionTypeDto) UnsetSigning() {
	o.Signing.Unset()
}

// GetEncrypt returns the Encrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateActionTypeDto) GetEncrypt() string {
	if o == nil || IsNil(o.Encrypt.Get()) {
		var ret string
		return ret
	}
	return *o.Encrypt.Get()
}

// GetEncryptOk returns a tuple with the Encrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateActionTypeDto) GetEncryptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encrypt.Get(), o.Encrypt.IsSet()
}

// HasEncrypt returns a boolean if a field has been set.
func (o *SsoSpCertificateActionTypeDto) IsEncryptSet() bool {
	if o != nil && o.Encrypt.IsSet() {
		return true
	}

	return false
}

// SetEncrypt gets a reference to the given NullableString and assigns it to the Encrypt field.
func (o *SsoSpCertificateActionTypeDto) SetEncrypt(v string) {
	o.Encrypt.Set(&v)
}
// SetEncryptNil sets the value for Encrypt to be an explicit nil
func (o *SsoSpCertificateActionTypeDto) SetEncryptNil() {
	o.Encrypt.Set(nil)
}

// UnsetEncrypt ensures that no value is present for Encrypt, not even an explicit nil
func (o *SsoSpCertificateActionTypeDto) UnsetEncrypt() {
	o.Encrypt.Unset()
}

// GetSigningAndEncrypt returns the SigningAndEncrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateActionTypeDto) GetSigningAndEncrypt() string {
	if o == nil || IsNil(o.SigningAndEncrypt.Get()) {
		var ret string
		return ret
	}
	return *o.SigningAndEncrypt.Get()
}

// GetSigningAndEncryptOk returns a tuple with the SigningAndEncrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateActionTypeDto) GetSigningAndEncryptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningAndEncrypt.Get(), o.SigningAndEncrypt.IsSet()
}

// HasSigningAndEncrypt returns a boolean if a field has been set.
func (o *SsoSpCertificateActionTypeDto) IsSigningAndEncryptSet() bool {
	if o != nil && o.SigningAndEncrypt.IsSet() {
		return true
	}

	return false
}

// SetSigningAndEncrypt gets a reference to the given NullableString and assigns it to the SigningAndEncrypt field.
func (o *SsoSpCertificateActionTypeDto) SetSigningAndEncrypt(v string) {
	o.SigningAndEncrypt.Set(&v)
}
// SetSigningAndEncryptNil sets the value for SigningAndEncrypt to be an explicit nil
func (o *SsoSpCertificateActionTypeDto) SetSigningAndEncryptNil() {
	o.SigningAndEncrypt.Set(nil)
}

// UnsetSigningAndEncrypt ensures that no value is present for SigningAndEncrypt, not even an explicit nil
func (o *SsoSpCertificateActionTypeDto) UnsetSigningAndEncrypt() {
	o.SigningAndEncrypt.Unset()
}

func (o SsoSpCertificateActionTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSpCertificateActionTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Signing.IsSet() {
		toSerialize["signing"] = o.Signing.Get()
	}
	if o.Encrypt.IsSet() {
		toSerialize["encrypt"] = o.Encrypt.Get()
	}
	if o.SigningAndEncrypt.IsSet() {
		toSerialize["signingAndEncrypt"] = o.SigningAndEncrypt.Get()
	}
	return toSerialize, nil
}

type NullableSsoSpCertificateActionTypeDto struct {
	value *SsoSpCertificateActionTypeDto
	isSet bool
}

func (v NullableSsoSpCertificateActionTypeDto) Get() *SsoSpCertificateActionTypeDto {
	return v.value
}

func (v *NullableSsoSpCertificateActionTypeDto) Set(val *SsoSpCertificateActionTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSpCertificateActionTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSpCertificateActionTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSpCertificateActionTypeDto(val *SsoSpCertificateActionTypeDto) *NullableSsoSpCertificateActionTypeDto {
	return &NullableSsoSpCertificateActionTypeDto{value: val, isSet: true}
}

func (v NullableSsoSpCertificateActionTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSpCertificateActionTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

