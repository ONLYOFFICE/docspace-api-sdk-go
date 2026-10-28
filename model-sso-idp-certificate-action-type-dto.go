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

// checks if the SsoIdpCertificateActionTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoIdpCertificateActionTypeDto{}

// SsoIdpCertificateActionTypeDto What the identity provider's certificate may be used for, as the `action` of an identity provider certificate.
type SsoIdpCertificateActionTypeDto struct {
	// The certificate verifies the signatures on what the provider sends, and nothing else - the counterpart of  the service provider's signing action.
	Verification NullableString `json:"verification,omitempty"`
	// The certificate is used to decrypt what the provider sends, but verifies no signature.
	Decrypt NullableString `json:"decrypt,omitempty"`
	// The certificate does both, which is what a single provider certificate has to be set to.
	VerificationAndDecrypt NullableString `json:"verificationAndDecrypt,omitempty"`
}

// NewSsoIdpCertificateActionTypeDto instantiates a new SsoIdpCertificateActionTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoIdpCertificateActionTypeDto() *SsoIdpCertificateActionTypeDto {
	this := SsoIdpCertificateActionTypeDto{}
	return &this
}

// NewSsoIdpCertificateActionTypeDtoWithDefaults instantiates a new SsoIdpCertificateActionTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoIdpCertificateActionTypeDtoWithDefaults() *SsoIdpCertificateActionTypeDto {
	this := SsoIdpCertificateActionTypeDto{}
	return &this
}

// GetVerification returns the Verification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateActionTypeDto) GetVerification() string {
	if o == nil || IsNil(o.Verification.Get()) {
		var ret string
		return ret
	}
	return *o.Verification.Get()
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateActionTypeDto) GetVerificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verification.Get(), o.Verification.IsSet()
}

// HasVerification returns a boolean if a field has been set.
func (o *SsoIdpCertificateActionTypeDto) IsVerificationSet() bool {
	if o != nil && o.Verification.IsSet() {
		return true
	}

	return false
}

// SetVerification gets a reference to the given NullableString and assigns it to the Verification field.
func (o *SsoIdpCertificateActionTypeDto) SetVerification(v string) {
	o.Verification.Set(&v)
}
// SetVerificationNil sets the value for Verification to be an explicit nil
func (o *SsoIdpCertificateActionTypeDto) SetVerificationNil() {
	o.Verification.Set(nil)
}

// UnsetVerification ensures that no value is present for Verification, not even an explicit nil
func (o *SsoIdpCertificateActionTypeDto) UnsetVerification() {
	o.Verification.Unset()
}

// GetDecrypt returns the Decrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateActionTypeDto) GetDecrypt() string {
	if o == nil || IsNil(o.Decrypt.Get()) {
		var ret string
		return ret
	}
	return *o.Decrypt.Get()
}

// GetDecryptOk returns a tuple with the Decrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateActionTypeDto) GetDecryptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Decrypt.Get(), o.Decrypt.IsSet()
}

// HasDecrypt returns a boolean if a field has been set.
func (o *SsoIdpCertificateActionTypeDto) IsDecryptSet() bool {
	if o != nil && o.Decrypt.IsSet() {
		return true
	}

	return false
}

// SetDecrypt gets a reference to the given NullableString and assigns it to the Decrypt field.
func (o *SsoIdpCertificateActionTypeDto) SetDecrypt(v string) {
	o.Decrypt.Set(&v)
}
// SetDecryptNil sets the value for Decrypt to be an explicit nil
func (o *SsoIdpCertificateActionTypeDto) SetDecryptNil() {
	o.Decrypt.Set(nil)
}

// UnsetDecrypt ensures that no value is present for Decrypt, not even an explicit nil
func (o *SsoIdpCertificateActionTypeDto) UnsetDecrypt() {
	o.Decrypt.Unset()
}

// GetVerificationAndDecrypt returns the VerificationAndDecrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateActionTypeDto) GetVerificationAndDecrypt() string {
	if o == nil || IsNil(o.VerificationAndDecrypt.Get()) {
		var ret string
		return ret
	}
	return *o.VerificationAndDecrypt.Get()
}

// GetVerificationAndDecryptOk returns a tuple with the VerificationAndDecrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateActionTypeDto) GetVerificationAndDecryptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationAndDecrypt.Get(), o.VerificationAndDecrypt.IsSet()
}

// HasVerificationAndDecrypt returns a boolean if a field has been set.
func (o *SsoIdpCertificateActionTypeDto) IsVerificationAndDecryptSet() bool {
	if o != nil && o.VerificationAndDecrypt.IsSet() {
		return true
	}

	return false
}

// SetVerificationAndDecrypt gets a reference to the given NullableString and assigns it to the VerificationAndDecrypt field.
func (o *SsoIdpCertificateActionTypeDto) SetVerificationAndDecrypt(v string) {
	o.VerificationAndDecrypt.Set(&v)
}
// SetVerificationAndDecryptNil sets the value for VerificationAndDecrypt to be an explicit nil
func (o *SsoIdpCertificateActionTypeDto) SetVerificationAndDecryptNil() {
	o.VerificationAndDecrypt.Set(nil)
}

// UnsetVerificationAndDecrypt ensures that no value is present for VerificationAndDecrypt, not even an explicit nil
func (o *SsoIdpCertificateActionTypeDto) UnsetVerificationAndDecrypt() {
	o.VerificationAndDecrypt.Unset()
}

func (o SsoIdpCertificateActionTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoIdpCertificateActionTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Verification.IsSet() {
		toSerialize["verification"] = o.Verification.Get()
	}
	if o.Decrypt.IsSet() {
		toSerialize["decrypt"] = o.Decrypt.Get()
	}
	if o.VerificationAndDecrypt.IsSet() {
		toSerialize["verificationAndDecrypt"] = o.VerificationAndDecrypt.Get()
	}
	return toSerialize, nil
}

type NullableSsoIdpCertificateActionTypeDto struct {
	value *SsoIdpCertificateActionTypeDto
	isSet bool
}

func (v NullableSsoIdpCertificateActionTypeDto) Get() *SsoIdpCertificateActionTypeDto {
	return v.value
}

func (v *NullableSsoIdpCertificateActionTypeDto) Set(val *SsoIdpCertificateActionTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoIdpCertificateActionTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoIdpCertificateActionTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoIdpCertificateActionTypeDto(val *SsoIdpCertificateActionTypeDto) *NullableSsoIdpCertificateActionTypeDto {
	return &NullableSsoIdpCertificateActionTypeDto{value: val, isSet: true}
}

func (v NullableSsoIdpCertificateActionTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoIdpCertificateActionTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

