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

// checks if the SsoSpCertificateAdvancedDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSpCertificateAdvancedDto{}

// SsoSpCertificateAdvancedDto How the portal signs and encrypts what it sends to the identity provider.
type SsoSpCertificateAdvancedDto struct {
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

// NewSsoSpCertificateAdvancedDto instantiates a new SsoSpCertificateAdvancedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSpCertificateAdvancedDto() *SsoSpCertificateAdvancedDto {
	this := SsoSpCertificateAdvancedDto{}
	return &this
}

// NewSsoSpCertificateAdvancedDtoWithDefaults instantiates a new SsoSpCertificateAdvancedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSpCertificateAdvancedDtoWithDefaults() *SsoSpCertificateAdvancedDto {
	this := SsoSpCertificateAdvancedDto{}
	return &this
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvancedDto) GetSigningAlgorithm() string {
	if o == nil || IsNil(o.SigningAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm.Get()
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvancedDto) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningAlgorithm.Get(), o.SigningAlgorithm.IsSet()
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsSigningAlgorithmSet() bool {
	if o != nil && o.SigningAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given NullableString and assigns it to the SigningAlgorithm field.
func (o *SsoSpCertificateAdvancedDto) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm.Set(&v)
}
// SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvancedDto) SetSigningAlgorithmNil() {
	o.SigningAlgorithm.Set(nil)
}

// UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvancedDto) UnsetSigningAlgorithm() {
	o.SigningAlgorithm.Unset()
}

// GetSignAuthRequests returns the SignAuthRequests field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvancedDto) GetSignAuthRequests() bool {
	if o == nil || IsNil(o.SignAuthRequests) {
		var ret bool
		return ret
	}
	return *o.SignAuthRequests
}

// GetSignAuthRequestsOk returns a tuple with the SignAuthRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvancedDto) GetSignAuthRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthRequests) {
		return nil, false
	}
	return o.SignAuthRequests, true
}

// HasSignAuthRequests returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsSignAuthRequestsSet() bool {
	if o != nil && !IsNil(o.SignAuthRequests) {
		return true
	}

	return false
}

// SetSignAuthRequests gets a reference to the given bool and assigns it to the SignAuthRequests field.
func (o *SsoSpCertificateAdvancedDto) SetSignAuthRequests(v bool) {
	o.SignAuthRequests = &v
}

// GetSignLogoutRequests returns the SignLogoutRequests field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvancedDto) GetSignLogoutRequests() bool {
	if o == nil || IsNil(o.SignLogoutRequests) {
		var ret bool
		return ret
	}
	return *o.SignLogoutRequests
}

// GetSignLogoutRequestsOk returns a tuple with the SignLogoutRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvancedDto) GetSignLogoutRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLogoutRequests) {
		return nil, false
	}
	return o.SignLogoutRequests, true
}

// HasSignLogoutRequests returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsSignLogoutRequestsSet() bool {
	if o != nil && !IsNil(o.SignLogoutRequests) {
		return true
	}

	return false
}

// SetSignLogoutRequests gets a reference to the given bool and assigns it to the SignLogoutRequests field.
func (o *SsoSpCertificateAdvancedDto) SetSignLogoutRequests(v bool) {
	o.SignLogoutRequests = &v
}

// GetSignLogoutResponses returns the SignLogoutResponses field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvancedDto) GetSignLogoutResponses() bool {
	if o == nil || IsNil(o.SignLogoutResponses) {
		var ret bool
		return ret
	}
	return *o.SignLogoutResponses
}

// GetSignLogoutResponsesOk returns a tuple with the SignLogoutResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvancedDto) GetSignLogoutResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.SignLogoutResponses) {
		return nil, false
	}
	return o.SignLogoutResponses, true
}

// HasSignLogoutResponses returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsSignLogoutResponsesSet() bool {
	if o != nil && !IsNil(o.SignLogoutResponses) {
		return true
	}

	return false
}

// SetSignLogoutResponses gets a reference to the given bool and assigns it to the SignLogoutResponses field.
func (o *SsoSpCertificateAdvancedDto) SetSignLogoutResponses(v bool) {
	o.SignLogoutResponses = &v
}

// GetEncryptAlgorithm returns the EncryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvancedDto) GetEncryptAlgorithm() string {
	if o == nil || IsNil(o.EncryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.EncryptAlgorithm.Get()
}

// GetEncryptAlgorithmOk returns a tuple with the EncryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvancedDto) GetEncryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptAlgorithm.Get(), o.EncryptAlgorithm.IsSet()
}

// HasEncryptAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsEncryptAlgorithmSet() bool {
	if o != nil && o.EncryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetEncryptAlgorithm gets a reference to the given NullableString and assigns it to the EncryptAlgorithm field.
func (o *SsoSpCertificateAdvancedDto) SetEncryptAlgorithm(v string) {
	o.EncryptAlgorithm.Set(&v)
}
// SetEncryptAlgorithmNil sets the value for EncryptAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvancedDto) SetEncryptAlgorithmNil() {
	o.EncryptAlgorithm.Set(nil)
}

// UnsetEncryptAlgorithm ensures that no value is present for EncryptAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvancedDto) UnsetEncryptAlgorithm() {
	o.EncryptAlgorithm.Unset()
}

// GetDecryptAlgorithm returns the DecryptAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSpCertificateAdvancedDto) GetDecryptAlgorithm() string {
	if o == nil || IsNil(o.DecryptAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.DecryptAlgorithm.Get()
}

// GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSpCertificateAdvancedDto) GetDecryptAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecryptAlgorithm.Get(), o.DecryptAlgorithm.IsSet()
}

// HasDecryptAlgorithm returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsDecryptAlgorithmSet() bool {
	if o != nil && o.DecryptAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetDecryptAlgorithm gets a reference to the given NullableString and assigns it to the DecryptAlgorithm field.
func (o *SsoSpCertificateAdvancedDto) SetDecryptAlgorithm(v string) {
	o.DecryptAlgorithm.Set(&v)
}
// SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil
func (o *SsoSpCertificateAdvancedDto) SetDecryptAlgorithmNil() {
	o.DecryptAlgorithm.Set(nil)
}

// UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
func (o *SsoSpCertificateAdvancedDto) UnsetDecryptAlgorithm() {
	o.DecryptAlgorithm.Unset()
}

// GetEncryptAssertions returns the EncryptAssertions field value if set, zero value otherwise.
func (o *SsoSpCertificateAdvancedDto) GetEncryptAssertions() bool {
	if o == nil || IsNil(o.EncryptAssertions) {
		var ret bool
		return ret
	}
	return *o.EncryptAssertions
}

// GetEncryptAssertionsOk returns a tuple with the EncryptAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSpCertificateAdvancedDto) GetEncryptAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptAssertions) {
		return nil, false
	}
	return o.EncryptAssertions, true
}

// HasEncryptAssertions returns a boolean if a field has been set.
func (o *SsoSpCertificateAdvancedDto) IsEncryptAssertionsSet() bool {
	if o != nil && !IsNil(o.EncryptAssertions) {
		return true
	}

	return false
}

// SetEncryptAssertions gets a reference to the given bool and assigns it to the EncryptAssertions field.
func (o *SsoSpCertificateAdvancedDto) SetEncryptAssertions(v bool) {
	o.EncryptAssertions = &v
}

func (o SsoSpCertificateAdvancedDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSpCertificateAdvancedDto) ToMap() (map[string]interface{}, error) {
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

type NullableSsoSpCertificateAdvancedDto struct {
	value *SsoSpCertificateAdvancedDto
	isSet bool
}

func (v NullableSsoSpCertificateAdvancedDto) Get() *SsoSpCertificateAdvancedDto {
	return v.value
}

func (v *NullableSsoSpCertificateAdvancedDto) Set(val *SsoSpCertificateAdvancedDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSpCertificateAdvancedDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSpCertificateAdvancedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSpCertificateAdvancedDto(val *SsoSpCertificateAdvancedDto) *NullableSsoSpCertificateAdvancedDto {
	return &NullableSsoSpCertificateAdvancedDto{value: val, isSet: true}
}

func (v NullableSsoSpCertificateAdvancedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSpCertificateAdvancedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

