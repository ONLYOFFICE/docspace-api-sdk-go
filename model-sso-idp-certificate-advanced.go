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

// checks if the SsoIdpCertificateAdvanced type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoIdpCertificateAdvanced{}

// SsoIdpCertificateAdvanced The IdP advanced certificate parameters.
type SsoIdpCertificateAdvanced struct {
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

// NewSsoIdpCertificateAdvanced instantiates a new SsoIdpCertificateAdvanced object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoIdpCertificateAdvanced() *SsoIdpCertificateAdvanced {
	this := SsoIdpCertificateAdvanced{}
	return &this
}

// NewSsoIdpCertificateAdvancedWithDefaults instantiates a new SsoIdpCertificateAdvanced object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoIdpCertificateAdvancedWithDefaults() *SsoIdpCertificateAdvanced {
	this := SsoIdpCertificateAdvanced{}
	return &this
}

// GetVerifyAlgorithm returns the VerifyAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateAdvanced) GetVerifyAlgorithm() string {
	if o == nil || IsNil(o.VerifyAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.VerifyAlgorithm.Get()
}

// GetVerifyAlgorithmOk returns a tuple with the VerifyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateAdvanced) GetVerifyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifyAlgorithm.Get(), o.VerifyAlgorithm.IsSet()
}

// HasVerifyAlgorithm returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsVerifyAlgorithmSet() bool {
	if o != nil && o.VerifyAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetVerifyAlgorithm gets a reference to the given NullableString and assigns it to the VerifyAlgorithm field.
func (o *SsoIdpCertificateAdvanced) SetVerifyAlgorithm(v string) {
	o.VerifyAlgorithm.Set(&v)
}
// SetVerifyAlgorithmNil sets the value for VerifyAlgorithm to be an explicit nil
func (o *SsoIdpCertificateAdvanced) SetVerifyAlgorithmNil() {
	o.VerifyAlgorithm.Set(nil)
}

// UnsetVerifyAlgorithm ensures that no value is present for VerifyAlgorithm, not even an explicit nil
func (o *SsoIdpCertificateAdvanced) UnsetVerifyAlgorithm() {
	o.VerifyAlgorithm.Unset()
}

// GetVerifyAuthResponsesSign returns the VerifyAuthResponsesSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvanced) GetVerifyAuthResponsesSign() bool {
	if o == nil || IsNil(o.VerifyAuthResponsesSign) {
		var ret bool
		return ret
	}
	return *o.VerifyAuthResponsesSign
}

// GetVerifyAuthResponsesSignOk returns a tuple with the VerifyAuthResponsesSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvanced) GetVerifyAuthResponsesSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyAuthResponsesSign) {
		return nil, false
	}
	return o.VerifyAuthResponsesSign, true
}

// HasVerifyAuthResponsesSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsVerifyAuthResponsesSignSet() bool {
	if o != nil && !IsNil(o.VerifyAuthResponsesSign) {
		return true
	}

	return false
}

// SetVerifyAuthResponsesSign gets a reference to the given bool and assigns it to the VerifyAuthResponsesSign field.
func (o *SsoIdpCertificateAdvanced) SetVerifyAuthResponsesSign(v bool) {
	o.VerifyAuthResponsesSign = &v
}

// GetVerifyLogoutRequestsSign returns the VerifyLogoutRequestsSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutRequestsSign() bool {
	if o == nil || IsNil(o.VerifyLogoutRequestsSign) {
		var ret bool
		return ret
	}
	return *o.VerifyLogoutRequestsSign
}

// GetVerifyLogoutRequestsSignOk returns a tuple with the VerifyLogoutRequestsSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutRequestsSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyLogoutRequestsSign) {
		return nil, false
	}
	return o.VerifyLogoutRequestsSign, true
}

// HasVerifyLogoutRequestsSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsVerifyLogoutRequestsSignSet() bool {
	if o != nil && !IsNil(o.VerifyLogoutRequestsSign) {
		return true
	}

	return false
}

// SetVerifyLogoutRequestsSign gets a reference to the given bool and assigns it to the VerifyLogoutRequestsSign field.
func (o *SsoIdpCertificateAdvanced) SetVerifyLogoutRequestsSign(v bool) {
	o.VerifyLogoutRequestsSign = &v
}

// GetVerifyLogoutResponsesSign returns the VerifyLogoutResponsesSign field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutResponsesSign() bool {
	if o == nil || IsNil(o.VerifyLogoutResponsesSign) {
		var ret bool
		return ret
	}
	return *o.VerifyLogoutResponsesSign
}

// GetVerifyLogoutResponsesSignOk returns a tuple with the VerifyLogoutResponsesSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutResponsesSignOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyLogoutResponsesSign) {
		return nil, false
	}
	return o.VerifyLogoutResponsesSign, true
}

// HasVerifyLogoutResponsesSign returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsVerifyLogoutResponsesSignSet() bool {
	if o != nil && !IsNil(o.VerifyLogoutResponsesSign) {
		return true
	}

	return false
}

// SetVerifyLogoutResponsesSign gets a reference to the given bool and assigns it to the VerifyLogoutResponsesSign field.
func (o *SsoIdpCertificateAdvanced) SetVerifyLogoutResponsesSign(v bool) {
	o.VerifyLogoutResponsesSign = &v
}

// GetDecryptAlgorithm returns the DecryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoIdpCertificateAdvanced) GetDecryptAlgorithm() string {
	if o == nil || IsNil(o.DecryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.DecryptAlgorithm.Get()
}

// GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoIdpCertificateAdvanced) GetDecryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecryptAlgorithm.Get(), o.DecryptAlgorithm.IsSet()
}

// HasDecryptAlgorithm returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsDecryptAlgorithmSet() bool {
	if o != nil && o.DecryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetDecryptAlgorithm gets a reference to the given NullableString and assigns it to the DecryptAlgorithm field.
func (o *SsoIdpCertificateAdvanced) SetDecryptAlgorithm(v string) {
	o.DecryptAlgorithm.Set(&v)
}
// SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil
func (o *SsoIdpCertificateAdvanced) SetDecryptAlgorithmNil() {
	o.DecryptAlgorithm.Set(nil)
}

// UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
func (o *SsoIdpCertificateAdvanced) UnsetDecryptAlgorithm() {
	o.DecryptAlgorithm.Unset()
}

// GetDecryptAssertions returns the DecryptAssertions field value if set, zero value otherwise.
func (o *SsoIdpCertificateAdvanced) GetDecryptAssertions() bool {
	if o == nil || IsNil(o.DecryptAssertions) {
		var ret bool
		return ret
	}
	return *o.DecryptAssertions
}

// GetDecryptAssertionsOk returns a tuple with the DecryptAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoIdpCertificateAdvanced) GetDecryptAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.DecryptAssertions) {
		return nil, false
	}
	return o.DecryptAssertions, true
}

// HasDecryptAssertions returns a boolean if a field has been set.
func (o *SsoIdpCertificateAdvanced) IsDecryptAssertionsSet() bool {
	if o != nil && !IsNil(o.DecryptAssertions) {
		return true
	}

	return false
}

// SetDecryptAssertions gets a reference to the given bool and assigns it to the DecryptAssertions field.
func (o *SsoIdpCertificateAdvanced) SetDecryptAssertions(v bool) {
	o.DecryptAssertions = &v
}

func (o SsoIdpCertificateAdvanced) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoIdpCertificateAdvanced) ToMap() (map[string]interface{}, error) {
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

type NullableSsoIdpCertificateAdvanced struct {
	value *SsoIdpCertificateAdvanced
	isSet bool
}

func (v NullableSsoIdpCertificateAdvanced) Get() *SsoIdpCertificateAdvanced {
	return v.value
}

func (v *NullableSsoIdpCertificateAdvanced) Set(val *SsoIdpCertificateAdvanced) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoIdpCertificateAdvanced) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoIdpCertificateAdvanced) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoIdpCertificateAdvanced(val *SsoIdpCertificateAdvanced) *NullableSsoIdpCertificateAdvanced {
	return &NullableSsoIdpCertificateAdvanced{value: val, isSet: true}
}

func (v NullableSsoIdpCertificateAdvanced) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoIdpCertificateAdvanced) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

