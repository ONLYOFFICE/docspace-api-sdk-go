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

// checks if the SsoSpCertificateAdvanced type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSpCertificateAdvanced{}

// SsoSpCertificateAdvanced The SP advanced certificate parameters.
type SsoSpCertificateAdvanced struct {
	// The certificate signing algorithm.
	SigningAlgorithm NullableString `json:"signingAlgorithm,omitempty"`
	// Specifies if SP will sign the SAML authentication requests sent to IdP or not.
	SignAuthRequests *bool `json:"signAuthRequests,omitempty"`
	// Specifies if SP will sign the SAML logout requests sent to IdP or not.
	SignLogoutRequests *bool `json:"signLogoutRequests,omitempty"`
	// Specifies if SP will sign the SAML logout responses sent to IdP or not.
	SignLogoutResponses *bool `json:"signLogoutResponses,omitempty"`
	// The certificate encryption algorithm.
	EncryptAlgorithm NullableString `json:"encryptAlgorithm,omitempty"`
	// The certificate decryption algorithm.
	DecryptAlgorithm NullableString `json:"decryptAlgorithm,omitempty"`
	// Specifies if the assertions will be encrypted or not.
	EncryptAssertions *bool `json:"encryptAssertions,omitempty"`
}

// NewSsoSpCertificateAdvanced instantiates a new SsoSpCertificateAdvanced object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSpCertificateAdvanced() *SsoSpCertificateAdvanced {
	this := SsoSpCertificateAdvanced{}
	return &this
}

// NewSsoSpCertificateAdvancedWithDefaults instantiates a new SsoSpCertificateAdvanced object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSpCertificateAdvancedWithDefaults() *SsoSpCertificateAdvanced {
	this := SsoSpCertificateAdvanced{}
	return &this
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvanced) GetSigningAlgorithm() string {
	if o == nil || IsNil(o.SigningAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm.Get()
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvanced) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningAlgorithm.Get(), o.SigningAlgorithm.IsSet()
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsSigningAlgorithmSet() bool {
	if o != nil && o.SigningAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given NullableString and assigns it to the SigningAlgorithm field.
func (o *SsoSpCertificateAdvanced) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm.Set(&v)
}
// SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvanced) SetSigningAlgorithmNil() {
	o.SigningAlgorithm.Set(nil)
}

// UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvanced) UnsetSigningAlgorithm() {
	o.SigningAlgorithm.Unset()
}

// GetSignAuthRequests returns the SignAuthRequests field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvanced) GetSignAuthRequests() bool {
	if o == nil || IsNil(o.SignAuthRequests) {
		var ret bool
		return ret
	}
	return *o.SignAuthRequests
}

// GetSignAuthRequestsOk returns a tuple with the SignAuthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvanced) GetSignAuthRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthRequests) {
		return nil, false
	}
	return o.SignAuthRequests, true
}

// HasSignAuthRequests returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsSignAuthRequestsSet() bool {
	if o != nil && !IsNil(o.SignAuthRequests) {
		return true
	}

	return false
}

// SetSignAuthRequests gets a reference to the given bool and assigns it to the SignAuthRequests field.
func (o *SsoSpCertificateAdvanced) SetSignAuthRequests(v bool) {
	o.SignAuthRequests = &v
}

// GetSignLogoutRequests returns the SignLogoutRequests field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvanced) GetSignLogoutRequests() bool {
	if o == nil || IsNil(o.SignLogoutRequests) {
		var ret bool
		return ret
	}
	return *o.SignLogoutRequests
}

// GetSignLogoutRequestsOk returns a tuple with the SignLogoutRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvanced) GetSignLogoutRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLogoutRequests) {
		return nil, false
	}
	return o.SignLogoutRequests, true
}

// HasSignLogoutRequests returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsSignLogoutRequestsSet() bool {
	if o != nil && !IsNil(o.SignLogoutRequests) {
		return true
	}

	return false
}

// SetSignLogoutRequests gets a reference to the given bool and assigns it to the SignLogoutRequests field.
func (o *SsoSpCertificateAdvanced) SetSignLogoutRequests(v bool) {
	o.SignLogoutRequests = &v
}

// GetSignLogoutResponses returns the SignLogoutResponses field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvanced) GetSignLogoutResponses() bool {
	if o == nil || IsNil(o.SignLogoutResponses) {
		var ret bool
		return ret
	}
	return *o.SignLogoutResponses
}

// GetSignLogoutResponsesOk returns a tuple with the SignLogoutResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvanced) GetSignLogoutResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLogoutResponses) {
		return nil, false
	}
	return o.SignLogoutResponses, true
}

// HasSignLogoutResponses returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsSignLogoutResponsesSet() bool {
	if o != nil && !IsNil(o.SignLogoutResponses) {
		return true
	}

	return false
}

// SetSignLogoutResponses gets a reference to the given bool and assigns it to the SignLogoutResponses field.
func (o *SsoSpCertificateAdvanced) SetSignLogoutResponses(v bool) {
	o.SignLogoutResponses = &v
}

// GetEncryptAlgorithm returns the EncryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvanced) GetEncryptAlgorithm() string {
	if o == nil || IsNil(o.EncryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptAlgorithm.Get()
}

// GetEncryptAlgorithmOk returns a tuple with the EncryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvanced) GetEncryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptAlgorithm.Get(), o.EncryptAlgorithm.IsSet()
}

// HasEncryptAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsEncryptAlgorithmSet() bool {
	if o != nil && o.EncryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetEncryptAlgorithm gets a reference to the given NullableString and assigns it to the EncryptAlgorithm field.
func (o *SsoSpCertificateAdvanced) SetEncryptAlgorithm(v string) {
	o.EncryptAlgorithm.Set(&v)
}
// SetEncryptAlgorithmNil sets the value for EncryptAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvanced) SetEncryptAlgorithmNil() {
	o.EncryptAlgorithm.Set(nil)
}

// UnsetEncryptAlgorithm ensures that no value is present for EncryptAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvanced) UnsetEncryptAlgorithm() {
	o.EncryptAlgorithm.Unset()
}

// GetDecryptAlgorithm returns the DecryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvanced) GetDecryptAlgorithm() string {
	if o == nil || IsNil(o.DecryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.DecryptAlgorithm.Get()
}

// GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvanced) GetDecryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecryptAlgorithm.Get(), o.DecryptAlgorithm.IsSet()
}

// HasDecryptAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsDecryptAlgorithmSet() bool {
	if o != nil && o.DecryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetDecryptAlgorithm gets a reference to the given NullableString and assigns it to the DecryptAlgorithm field.
func (o *SsoSpCertificateAdvanced) SetDecryptAlgorithm(v string) {
	o.DecryptAlgorithm.Set(&v)
}
// SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvanced) SetDecryptAlgorithmNil() {
	o.DecryptAlgorithm.Set(nil)
}

// UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvanced) UnsetDecryptAlgorithm() {
	o.DecryptAlgorithm.Unset()
}

// GetEncryptAssertions returns the EncryptAssertions field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvanced) GetEncryptAssertions() bool {
	if o == nil || IsNil(o.EncryptAssertions) {
		var ret bool
		return ret
	}
	return *o.EncryptAssertions
}

// GetEncryptAssertionsOk returns a tuple with the EncryptAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvanced) GetEncryptAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptAssertions) {
		return nil, false
	}
	return o.EncryptAssertions, true
}

// HasEncryptAssertions returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvanced) IsEncryptAssertionsSet() bool {
	if o != nil && !IsNil(o.EncryptAssertions) {
		return true
	}

	return false
}

// SetEncryptAssertions gets a reference to the given bool and assigns it to the EncryptAssertions field.
func (o *SsoSpCertificateAdvanced) SetEncryptAssertions(v bool) {
	o.EncryptAssertions = &v
}

func (o SsoSpCertificateAdvanced) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSpCertificateAdvanced) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SigningAlgorithm.IsSet() {
		toSerialize["signingAlgorithm"] = o.SigningAlgorithm.Get()
	}
	if !IsNil(o.SignAuthRequests) {
		toSerialize["signAuthRequests"] = o.SignAuthRequests
	}
	if !IsNil(o.SignLogoutRequests) {
		toSerialize["signLogoutRequests"] = o.SignLogoutRequests
	}
	if !IsNil(o.SignLogoutResponses) {
		toSerialize["signLogoutResponses"] = o.SignLogoutResponses
	}
	if o.EncryptAlgorithm.IsSet() {
		toSerialize["encryptAlgorithm"] = o.EncryptAlgorithm.Get()
	}
	if o.DecryptAlgorithm.IsSet() {
		toSerialize["decryptAlgorithm"] = o.DecryptAlgorithm.Get()
	}
	if !IsNil(o.EncryptAssertions) {
		toSerialize["encryptAssertions"] = o.EncryptAssertions
	}
	return toSerialize, nil
}

type NullableSsoSpCertificateAdvanced struct {
	value *SsoSpCertificateAdvanced
	isSet bool
}

func (v NullableSsoSpCertificateAdvanced) Get() *SsoSpCertificateAdvanced {
	return v.value
}

func (v *NullableSsoSpCertificateAdvanced) Set(val *SsoSpCertificateAdvanced) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSpCertificateAdvanced) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSpCertificateAdvanced) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSpCertificateAdvanced(val *SsoSpCertificateAdvanced) *NullableSsoSpCertificateAdvanced {
	return &NullableSsoSpCertificateAdvanced{value: val, isSet: true}
}

func (v NullableSsoSpCertificateAdvanced) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSpCertificateAdvanced) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

