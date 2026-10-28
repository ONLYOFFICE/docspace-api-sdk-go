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

// checks if the SsoSettingsConstantsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSettingsConstantsDto{}

// SsoSettingsConstantsDto The SSO settings constants: every value the settings accept, by name.
type SsoSettingsConstantsDto struct {
	// The values the `nameIdFormat` of the identity provider settings accepts. The built-in configuration uses  the SAML 2.0 transient format.
	SsoNameIdFormatType *SsoNameIdFormatTypeDto `json:"ssoNameIdFormatType,omitempty"`
	// The values the `ssoBinding` and `sloBinding` of the identity provider settings accept - how the portal  sends its sign-in and sign-out requests. The built-in configuration uses HTTP POST for both.
	SsoBindingType *SsoBindingTypeDto `json:"ssoBindingType,omitempty"`
	// The values the `signingAlgorithm` of the service provider certificate and the `verifyAlgorithm` of the  identity provider certificate accept. The built-in configuration uses RSA-SHA1 for both.
	SsoSigningAlgorithmType *SsoSigningAlgorithmTypeDto `json:"ssoSigningAlgorithmType,omitempty"`
	// The values the `encryptAlgorithm` and `decryptAlgorithm` of the certificate settings accept. The built-in  configuration uses AES-128 everywhere.
	SsoEncryptAlgorithmType *SsoEncryptAlgorithmTypeDto `json:"ssoEncryptAlgorithmType,omitempty"`
	// The values the `action` of a service provider certificate accepts, which is what the portal's own key  pair may be used for.
	SsoSpCertificateActionType *SsoSpCertificateActionTypeDto `json:"ssoSpCertificateActionType,omitempty"`
	// The values the `action` of an identity provider certificate accepts, which is what the provider's  certificate may be used for - the mirror image of the service provider actions.
	SsoIdpCertificateActionType *SsoIdpCertificateActionTypeDto `json:"ssoIdpCertificateActionType,omitempty"`
}

// NewSsoSettingsConstantsDto instantiates a new SsoSettingsConstantsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSettingsConstantsDto() *SsoSettingsConstantsDto {
	this := SsoSettingsConstantsDto{}
	return &this
}

// NewSsoSettingsConstantsDtoWithDefaults instantiates a new SsoSettingsConstantsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSettingsConstantsDtoWithDefaults() *SsoSettingsConstantsDto {
	this := SsoSettingsConstantsDto{}
	return &this
}

// GetSsoNameIdFormatType returns the SsoNameIdFormatType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoNameIdFormatType() SsoNameIdFormatTypeDto {
	if o == nil || IsNil(o.SsoNameIdFormatType) {
		var ret SsoNameIdFormatTypeDto
		return ret
	}
	return *o.SsoNameIdFormatType
}

// GetSsoNameIdFormatTypeOk returns a tuple with the SsoNameIdFormatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoNameIdFormatTypeOk() (*SsoNameIdFormatTypeDto, bool) {
	if o == nil || IsNil(o.SsoNameIdFormatType) {
		return nil, false
	}
	return o.SsoNameIdFormatType, true
}

// HasSsoNameIdFormatType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoNameIdFormatTypeSet() bool {
	if o != nil && !IsNil(o.SsoNameIdFormatType) {
		return true
	}

	return false
}

// SetSsoNameIdFormatType gets a reference to the given SsoNameIdFormatTypeDto and assigns it to the SsoNameIdFormatType field.
func (o *SsoSettingsConstantsDto) SetSsoNameIdFormatType(v SsoNameIdFormatTypeDto) {
	o.SsoNameIdFormatType = &v
}

// GetSsoBindingType returns the SsoBindingType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoBindingType() SsoBindingTypeDto {
	if o == nil || IsNil(o.SsoBindingType) {
		var ret SsoBindingTypeDto
		return ret
	}
	return *o.SsoBindingType
}

// GetSsoBindingTypeOk returns a tuple with the SsoBindingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoBindingTypeOk() (*SsoBindingTypeDto, bool) {
	if o == nil || IsNil(o.SsoBindingType) {
		return nil, false
	}
	return o.SsoBindingType, true
}

// HasSsoBindingType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoBindingTypeSet() bool {
	if o != nil && !IsNil(o.SsoBindingType) {
		return true
	}

	return false
}

// SetSsoBindingType gets a reference to the given SsoBindingTypeDto and assigns it to the SsoBindingType field.
func (o *SsoSettingsConstantsDto) SetSsoBindingType(v SsoBindingTypeDto) {
	o.SsoBindingType = &v
}

// GetSsoSigningAlgorithmType returns the SsoSigningAlgorithmType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoSigningAlgorithmType() SsoSigningAlgorithmTypeDto {
	if o == nil || IsNil(o.SsoSigningAlgorithmType) {
		var ret SsoSigningAlgorithmTypeDto
		return ret
	}
	return *o.SsoSigningAlgorithmType
}

// GetSsoSigningAlgorithmTypeOk returns a tuple with the SsoSigningAlgorithmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoSigningAlgorithmTypeOk() (*SsoSigningAlgorithmTypeDto, bool) {
	if o == nil || IsNil(o.SsoSigningAlgorithmType) {
		return nil, false
	}
	return o.SsoSigningAlgorithmType, true
}

// HasSsoSigningAlgorithmType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoSigningAlgorithmTypeSet() bool {
	if o != nil && !IsNil(o.SsoSigningAlgorithmType) {
		return true
	}

	return false
}

// SetSsoSigningAlgorithmType gets a reference to the given SsoSigningAlgorithmTypeDto and assigns it to the SsoSigningAlgorithmType field.
func (o *SsoSettingsConstantsDto) SetSsoSigningAlgorithmType(v SsoSigningAlgorithmTypeDto) {
	o.SsoSigningAlgorithmType = &v
}

// GetSsoEncryptAlgorithmType returns the SsoEncryptAlgorithmType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoEncryptAlgorithmType() SsoEncryptAlgorithmTypeDto {
	if o == nil || IsNil(o.SsoEncryptAlgorithmType) {
		var ret SsoEncryptAlgorithmTypeDto
		return ret
	}
	return *o.SsoEncryptAlgorithmType
}

// GetSsoEncryptAlgorithmTypeOk returns a tuple with the SsoEncryptAlgorithmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoEncryptAlgorithmTypeOk() (*SsoEncryptAlgorithmTypeDto, bool) {
	if o == nil || IsNil(o.SsoEncryptAlgorithmType) {
		return nil, false
	}
	return o.SsoEncryptAlgorithmType, true
}

// HasSsoEncryptAlgorithmType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoEncryptAlgorithmTypeSet() bool {
	if o != nil && !IsNil(o.SsoEncryptAlgorithmType) {
		return true
	}

	return false
}

// SetSsoEncryptAlgorithmType gets a reference to the given SsoEncryptAlgorithmTypeDto and assigns it to the SsoEncryptAlgorithmType field.
func (o *SsoSettingsConstantsDto) SetSsoEncryptAlgorithmType(v SsoEncryptAlgorithmTypeDto) {
	o.SsoEncryptAlgorithmType = &v
}

// GetSsoSpCertificateActionType returns the SsoSpCertificateActionType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoSpCertificateActionType() SsoSpCertificateActionTypeDto {
	if o == nil || IsNil(o.SsoSpCertificateActionType) {
		var ret SsoSpCertificateActionTypeDto
		return ret
	}
	return *o.SsoSpCertificateActionType
}

// GetSsoSpCertificateActionTypeOk returns a tuple with the SsoSpCertificateActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoSpCertificateActionTypeOk() (*SsoSpCertificateActionTypeDto, bool) {
	if o == nil || IsNil(o.SsoSpCertificateActionType) {
		return nil, false
	}
	return o.SsoSpCertificateActionType, true
}

// HasSsoSpCertificateActionType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoSpCertificateActionTypeSet() bool {
	if o != nil && !IsNil(o.SsoSpCertificateActionType) {
		return true
	}

	return false
}

// SetSsoSpCertificateActionType gets a reference to the given SsoSpCertificateActionTypeDto and assigns it to the SsoSpCertificateActionType field.
func (o *SsoSettingsConstantsDto) SetSsoSpCertificateActionType(v SsoSpCertificateActionTypeDto) {
	o.SsoSpCertificateActionType = &v
}

// GetSsoIdpCertificateActionType returns the SsoIdpCertificateActionType field value if set, zero value otherwise.
func (o *SsoSettingsConstantsDto) GetSsoIdpCertificateActionType() SsoIdpCertificateActionTypeDto {
	if o == nil || IsNil(o.SsoIdpCertificateActionType) {
		var ret SsoIdpCertificateActionTypeDto
		return ret
	}
	return *o.SsoIdpCertificateActionType
}

// GetSsoIdpCertificateActionTypeOk returns a tuple with the SsoIdpCertificateActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsConstantsDto) GetSsoIdpCertificateActionTypeOk() (*SsoIdpCertificateActionTypeDto, bool) {
	if o == nil || IsNil(o.SsoIdpCertificateActionType) {
		return nil, false
	}
	return o.SsoIdpCertificateActionType, true
}

// HasSsoIdpCertificateActionType returns a boolean if a field has been set.
func (o *SsoSettingsConstantsDto) IsSsoIdpCertificateActionTypeSet() bool {
	if o != nil && !IsNil(o.SsoIdpCertificateActionType) {
		return true
	}

	return false
}

// SetSsoIdpCertificateActionType gets a reference to the given SsoIdpCertificateActionTypeDto and assigns it to the SsoIdpCertificateActionType field.
func (o *SsoSettingsConstantsDto) SetSsoIdpCertificateActionType(v SsoIdpCertificateActionTypeDto) {
	o.SsoIdpCertificateActionType = &v
}

func (o SsoSettingsConstantsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSettingsConstantsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsoNameIdFormatType) {
		toSerialize["ssoNameIdFormatType"] = o.SsoNameIdFormatType
	}
	if !IsNil(o.SsoBindingType) {
		toSerialize["ssoBindingType"] = o.SsoBindingType
	}
	if !IsNil(o.SsoSigningAlgorithmType) {
		toSerialize["ssoSigningAlgorithmType"] = o.SsoSigningAlgorithmType
	}
	if !IsNil(o.SsoEncryptAlgorithmType) {
		toSerialize["ssoEncryptAlgorithmType"] = o.SsoEncryptAlgorithmType
	}
	if !IsNil(o.SsoSpCertificateActionType) {
		toSerialize["ssoSpCertificateActionType"] = o.SsoSpCertificateActionType
	}
	if !IsNil(o.SsoIdpCertificateActionType) {
		toSerialize["ssoIdpCertificateActionType"] = o.SsoIdpCertificateActionType
	}
	return toSerialize, nil
}

type NullableSsoSettingsConstantsDto struct {
	value *SsoSettingsConstantsDto
	isSet bool
}

func (v NullableSsoSettingsConstantsDto) Get() *SsoSettingsConstantsDto {
	return v.value
}

func (v *NullableSsoSettingsConstantsDto) Set(val *SsoSettingsConstantsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSettingsConstantsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSettingsConstantsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSettingsConstantsDto(val *SsoSettingsConstantsDto) *NullableSsoSettingsConstantsDto {
	return &NullableSsoSettingsConstantsDto{value: val, isSet: true}
}

func (v NullableSsoSettingsConstantsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSettingsConstantsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

