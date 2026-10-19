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

// checks if the SsoIdpCertificateAdvancedDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoIdpCertificateAdvancedDto{}

// SsoIdpCertificateAdvancedDto How the portal checks and decrypts what the identity provider sends.
type SsoIdpCertificateAdvancedDto struct {
	// The certificate verification algorithm.
	VerifyAlgorithm NullableString `json:"verifyAlgorithm,omitempty"`
	// Specifies if the signatures of the SAML authentication responses sent to SP will be verified or not.
	VerifyAuthResponsesSign *bool `json:"verifyAuthResponsesSign,omitempty"`
	// Specifies if the signatures of the SAML logout requests sent to SP will be verified or not.
	VerifyLogoutRequestsSign *bool `json:"verifyLogoutRequestsSign,omitempty"`
	// Specifies if the signatures of the SAML logout responses sent to SP will be verified or not.
	VerifyLogoutResponsesSign *bool `json:"verifyLogoutResponsesSign,omitempty"`
	// The certificate decryption algorithm.
	DecryptAlgorithm NullableString `json:"decryptAlgorithm,omitempty"`
	// Specifies if the assertions will be decrypted or not.
	DecryptAssertions *bool `json:"decryptAssertions,omitempty"`
}

// NewSsoIdpCertificateAdvancedDto instantiates a new SsoIdpCertificateAdvancedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoIdpCertificateAdvancedDto() *SsoIdpCertificateAdvancedDto {
	this := SsoIdpCertificateAdvancedDto{}
	return &this
}

// NewSsoIdpCertificateAdvancedDtoWithDefaults instantiates a new SsoIdpCertificateAdvancedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoIdpCertificateAdvancedDtoWithDefaults() *SsoIdpCertificateAdvancedDto {
	this := SsoIdpCertificateAdvancedDto{}
	return &this
}

// GetVerifyAlgorithm returns the VerifyAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateAdvancedDto) GetVerifyAlgorithm() string {
	if o == nil || IsNil(o.VerifyAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.VerifyAlgorithm.Get()
}

// GetVerifyAlgorithmOk returns a tuple with the VerifyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateAdvancedDto) GetVerifyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifyAlgorithm.Get(), o.VerifyAlgorithm.IsSet()
}

// HasVerifyAlgorithm returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsVerifyAlgorithmSet() bool {
	if o != nil && o.VerifyAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetVerifyAlgorithm gets a reference to the given NullableString and assigns it to the VerifyAlgorithm field.
func (o *SsoIdpCertificateAdvancedDto) SetVerifyAlgorithm(v string) {
	o.VerifyAlgorithm.Set(&v)
}
// SetVerifyAlgorithmNil sets the value for VerifyAlgorithm to be an explicit nil
func (o *SsoIdpCertificateAdvancedDto) SetVerifyAlgorithmNil() {
	o.VerifyAlgorithm.Set(nil)
}

// UnsetVerifyAlgorithm ensures that no value is present for VerifyAlgorithm, not even an explicit nil
func (o *SsoIdpCertificateAdvancedDto) UnsetVerifyAlgorithm() {
	o.VerifyAlgorithm.Unset()
}

// GetVerifyAuthResponsesSign returns the VerifyAuthResponsesSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyAuthResponsesSign() bool {
	if o == nil || IsNil(o.VerifyAuthResponsesSign) {
		var ret bool
		return ret
	}
	return *o.VerifyAuthResponsesSign
}

// GetVerifyAuthResponsesSignOk returns a tuple with the VerifyAuthResponsesSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyAuthResponsesSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyAuthResponsesSign) {
		return nil, false
	}
	return o.VerifyAuthResponsesSign, true
}

// HasVerifyAuthResponsesSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsVerifyAuthResponsesSignSet() bool {
	if o != nil && !IsNil(o.VerifyAuthResponsesSign) {
		return true
	}

	return false
}

// SetVerifyAuthResponsesSign gets a reference to the given bool and assigns it to the VerifyAuthResponsesSign field.
func (o *SsoIdpCertificateAdvancedDto) SetVerifyAuthResponsesSign(v bool) {
	o.VerifyAuthResponsesSign = &v
}

// GetVerifyLogoutRequestsSign returns the VerifyLogoutRequestsSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutRequestsSign() bool {
	if o == nil || IsNil(o.VerifyLogoutRequestsSign) {
		var ret bool
		return ret
	}
	return *o.VerifyLogoutRequestsSign
}

// GetVerifyLogoutRequestsSignOk returns a tuple with the VerifyLogoutRequestsSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutRequestsSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyLogoutRequestsSign) {
		return nil, false
	}
	return o.VerifyLogoutRequestsSign, true
}

// HasVerifyLogoutRequestsSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsVerifyLogoutRequestsSignSet() bool {
	if o != nil && !IsNil(o.VerifyLogoutRequestsSign) {
		return true
	}

	return false
}

// SetVerifyLogoutRequestsSign gets a reference to the given bool and assigns it to the VerifyLogoutRequestsSign field.
func (o *SsoIdpCertificateAdvancedDto) SetVerifyLogoutRequestsSign(v bool) {
	o.VerifyLogoutRequestsSign = &v
}

// GetVerifyLogoutResponsesSign returns the VerifyLogoutResponsesSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutResponsesSign() bool {
	if o == nil || IsNil(o.VerifyLogoutResponsesSign) {
		var ret bool
		return ret
	}
	return *o.VerifyLogoutResponsesSign
}

// GetVerifyLogoutResponsesSignOk returns a tuple with the VerifyLogoutResponsesSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutResponsesSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyLogoutResponsesSign) {
		return nil, false
	}
	return o.VerifyLogoutResponsesSign, true
}

// HasVerifyLogoutResponsesSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsVerifyLogoutResponsesSignSet() bool {
	if o != nil && !IsNil(o.VerifyLogoutResponsesSign) {
		return true
	}

	return false
}

// SetVerifyLogoutResponsesSign gets a reference to the given bool and assigns it to the VerifyLogoutResponsesSign field.
func (o *SsoIdpCertificateAdvancedDto) SetVerifyLogoutResponsesSign(v bool) {
	o.VerifyLogoutResponsesSign = &v
}

// GetDecryptAlgorithm returns the DecryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateAdvancedDto) GetDecryptAlgorithm() string {
	if o == nil || IsNil(o.DecryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.DecryptAlgorithm.Get()
}

// GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateAdvancedDto) GetDecryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecryptAlgorithm.Get(), o.DecryptAlgorithm.IsSet()
}

// HasDecryptAlgorithm returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsDecryptAlgorithmSet() bool {
	if o != nil && o.DecryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetDecryptAlgorithm gets a reference to the given NullableString and assigns it to the DecryptAlgorithm field.
func (o *SsoIdpCertificateAdvancedDto) SetDecryptAlgorithm(v string) {
	o.DecryptAlgorithm.Set(&v)
}
// SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil
func (o *SsoIdpCertificateAdvancedDto) SetDecryptAlgorithmNil() {
	o.DecryptAlgorithm.Set(nil)
}

// UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
func (o *SsoIdpCertificateAdvancedDto) UnsetDecryptAlgorithm() {
	o.DecryptAlgorithm.Unset()
}

// GetDecryptAssertions returns the DecryptAssertions field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvancedDto) GetDecryptAssertions() bool {
	if o == nil || IsNil(o.DecryptAssertions) {
		var ret bool
		return ret
	}
	return *o.DecryptAssertions
}

// GetDecryptAssertionsOk returns a tuple with the DecryptAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvancedDto) GetDecryptAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.DecryptAssertions) {
		return nil, false
	}
	return o.DecryptAssertions, true
}

// HasDecryptAssertions returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvancedDto) IsDecryptAssertionsSet() bool {
	if o != nil && !IsNil(o.DecryptAssertions) {
		return true
	}

	return false
}

// SetDecryptAssertions gets a reference to the given bool and assigns it to the DecryptAssertions field.
func (o *SsoIdpCertificateAdvancedDto) SetDecryptAssertions(v bool) {
	o.DecryptAssertions = &v
}

func (o SsoIdpCertificateAdvancedDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoIdpCertificateAdvancedDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VerifyAlgorithm.IsSet() {
		toSerialize["verifyAlgorithm"] = o.VerifyAlgorithm.Get()
	}
	if !IsNil(o.VerifyAuthResponsesSign) {
		toSerialize["verifyAuthResponsesSign"] = o.VerifyAuthResponsesSign
	}
	if !IsNil(o.VerifyLogoutRequestsSign) {
		toSerialize["verifyLogoutRequestsSign"] = o.VerifyLogoutRequestsSign
	}
	if !IsNil(o.VerifyLogoutResponsesSign) {
		toSerialize["verifyLogoutResponsesSign"] = o.VerifyLogoutResponsesSign
	}
	if o.DecryptAlgorithm.IsSet() {
		toSerialize["decryptAlgorithm"] = o.DecryptAlgorithm.Get()
	}
	if !IsNil(o.DecryptAssertions) {
		toSerialize["decryptAssertions"] = o.DecryptAssertions
	}
	return toSerialize, nil
}

type NullableSsoIdpCertificateAdvancedDto struct {
	value *SsoIdpCertificateAdvancedDto
	isSet bool
}

func (v NullableSsoIdpCertificateAdvancedDto) Get() *SsoIdpCertificateAdvancedDto {
	return v.value
}

func (v *NullableSsoIdpCertificateAdvancedDto) Set(val *SsoIdpCertificateAdvancedDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoIdpCertificateAdvancedDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoIdpCertificateAdvancedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoIdpCertificateAdvancedDto(val *SsoIdpCertificateAdvancedDto) *NullableSsoIdpCertificateAdvancedDto {
	return &NullableSsoIdpCertificateAdvancedDto{value: val, isSet: true}
}

func (v NullableSsoIdpCertificateAdvancedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoIdpCertificateAdvancedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

