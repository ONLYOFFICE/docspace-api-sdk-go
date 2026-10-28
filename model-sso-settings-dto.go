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
	"time"
)

// checks if the SsoSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSettingsDto{}

// SsoSettingsDto The SAML single sign-on configuration of the portal.
type SsoSettingsDto struct {
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// Specifies if the SSO settings are enabled or not.
	EnableSso NullableBool `json:"enableSso,omitempty"`
	// The SSO IdP settings.
	IdpSettings *SsoIdpSettingsDto `json:"idpSettings,omitempty"`
	// The list of the IdP certificates.
	IdpCertificates []SsoCertificateDto `json:"idpCertificates,omitempty"`
	// The IdP advanced certificate.
	IdpCertificateAdvanced *SsoIdpCertificateAdvancedDto `json:"idpCertificateAdvanced,omitempty"`
	// The SP login label.
	SpLoginLabel NullableString `json:"spLoginLabel,omitempty"`
	// The list of the SP certificates.
	SpCertificates []SsoCertificateDto `json:"spCertificates,omitempty"`
	// The SP advanced certificate.
	SpCertificateAdvanced *SsoSpCertificateAdvancedDto `json:"spCertificateAdvanced,omitempty"`
	// The SSO field mapping.
	FieldMapping *SsoFieldMappingDto `json:"fieldMapping,omitempty"`
	// Specifies if the authentication page will be hidden or not.
	HideAuthPage *bool `json:"hideAuthPage,omitempty"`
	// The user type.
	UsersType *int32 `json:"usersType,omitempty"`
	// Specifies if the email verification is disabled or not.
	DisableEmailVerification *bool `json:"disableEmailVerification,omitempty"`
}

// NewSsoSettingsDto instantiates a new SsoSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSettingsDto() *SsoSettingsDto {
	this := SsoSettingsDto{}
	return &this
}

// NewSsoSettingsDtoWithDefaults instantiates a new SsoSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSettingsDtoWithDefaults() *SsoSettingsDto {
	this := SsoSettingsDto{}
	return &this
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *SsoSettingsDto) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetEnableSso returns the EnableSso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsDto) GetEnableSso() bool {
	if o == nil || IsNil(o.EnableSso.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSso.Get()
}

// GetEnableSsoOk returns a tuple with the EnableSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsDto) GetEnableSsoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSso.Get(), o.EnableSso.IsSet()
}

// HasEnableSso returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsEnableSsoSet() bool {
	if o != nil && o.EnableSso.IsSet() {
		return true
	}

	return false
}

// SetEnableSso gets a reference to the given NullableBool and assigns it to the EnableSso field.
func (o *SsoSettingsDto) SetEnableSso(v bool) {
	o.EnableSso.Set(&v)
}
// SetEnableSsoNil sets the value for EnableSso to be an explicit nil
func (o *SsoSettingsDto) SetEnableSsoNil() {
	o.EnableSso.Set(nil)
}

// UnsetEnableSso ensures that no value is present for EnableSso, not even an explicit nil
func (o *SsoSettingsDto) UnsetEnableSso() {
	o.EnableSso.Unset()
}

// GetIdpSettings returns the IdpSettings field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetIdpSettings() SsoIdpSettingsDto {
	if o == nil || IsNil(o.IdpSettings) {
		var ret SsoIdpSettingsDto
		return ret
	}
	return *o.IdpSettings
}

// GetIdpSettingsOk returns a tuple with the IdpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetIdpSettingsOk() (*SsoIdpSettingsDto, bool) {
	if o == nil || IsNil(o.IdpSettings) {
		return nil, false
	}
	return o.IdpSettings, true
}

// HasIdpSettings returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsIdpSettingsSet() bool {
	if o != nil && !IsNil(o.IdpSettings) {
		return true
	}

	return false
}

// SetIdpSettings gets a reference to the given SsoIdpSettingsDto and assigns it to the IdpSettings field.
func (o *SsoSettingsDto) SetIdpSettings(v SsoIdpSettingsDto) {
	o.IdpSettings = &v
}

// GetIdpCertificates returns the IdpCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsDto) GetIdpCertificates() []SsoCertificateDto {
	if o == nil {
		var ret []SsoCertificateDto
		return ret
	}
	return o.IdpCertificates
}

// GetIdpCertificatesOk returns a tuple with the IdpCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsDto) GetIdpCertificatesOk() ([]SsoCertificateDto, bool) {
	if o == nil || IsNil(o.IdpCertificates) {
		return nil, false
	}
	return o.IdpCertificates, true
}

// HasIdpCertificates returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsIdpCertificatesSet() bool {
	if o != nil && !IsNil(o.IdpCertificates) {
		return true
	}

	return false
}

// SetIdpCertificates gets a reference to the given []SsoCertificateDto and assigns it to the IdpCertificates field.
func (o *SsoSettingsDto) SetIdpCertificates(v []SsoCertificateDto) {
	o.IdpCertificates = v
}

// GetIdpCertificateAdvanced returns the IdpCertificateAdvanced field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetIdpCertificateAdvanced() SsoIdpCertificateAdvancedDto {
	if o == nil || IsNil(o.IdpCertificateAdvanced) {
		var ret SsoIdpCertificateAdvancedDto
		return ret
	}
	return *o.IdpCertificateAdvanced
}

// GetIdpCertificateAdvancedOk returns a tuple with the IdpCertificateAdvanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetIdpCertificateAdvancedOk() (*SsoIdpCertificateAdvancedDto, bool) {
	if o == nil || IsNil(o.IdpCertificateAdvanced) {
		return nil, false
	}
	return o.IdpCertificateAdvanced, true
}

// HasIdpCertificateAdvanced returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsIdpCertificateAdvancedSet() bool {
	if o != nil && !IsNil(o.IdpCertificateAdvanced) {
		return true
	}

	return false
}

// SetIdpCertificateAdvanced gets a reference to the given SsoIdpCertificateAdvancedDto and assigns it to the IdpCertificateAdvanced field.
func (o *SsoSettingsDto) SetIdpCertificateAdvanced(v SsoIdpCertificateAdvancedDto) {
	o.IdpCertificateAdvanced = &v
}

// GetSpLoginLabel returns the SpLoginLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsDto) GetSpLoginLabel() string {
	if o == nil || IsNil(o.SpLoginLabel.Get()) {
		var ret string
		return ret
	}
	return *o.SpLoginLabel.Get()
}

// GetSpLoginLabelOk returns a tuple with the SpLoginLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsDto) GetSpLoginLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpLoginLabel.Get(), o.SpLoginLabel.IsSet()
}

// HasSpLoginLabel returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsSpLoginLabelSet() bool {
	if o != nil && o.SpLoginLabel.IsSet() {
		return true
	}

	return false
}

// SetSpLoginLabel gets a reference to the given NullableString and assigns it to the SpLoginLabel field.
func (o *SsoSettingsDto) SetSpLoginLabel(v string) {
	o.SpLoginLabel.Set(&v)
}
// SetSpLoginLabelNil sets the value for SpLoginLabel to be an explicit nil
func (o *SsoSettingsDto) SetSpLoginLabelNil() {
	o.SpLoginLabel.Set(nil)
}

// UnsetSpLoginLabel ensures that no value is present for SpLoginLabel, not even an explicit nil
func (o *SsoSettingsDto) UnsetSpLoginLabel() {
	o.SpLoginLabel.Unset()
}

// GetSpCertificates returns the SpCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsDto) GetSpCertificates() []SsoCertificateDto {
	if o == nil {
		var ret []SsoCertificateDto
		return ret
	}
	return o.SpCertificates
}

// GetSpCertificatesOk returns a tuple with the SpCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsDto) GetSpCertificatesOk() ([]SsoCertificateDto, bool) {
	if o == nil || IsNil(o.SpCertificates) {
		return nil, false
	}
	return o.SpCertificates, true
}

// HasSpCertificates returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsSpCertificatesSet() bool {
	if o != nil && !IsNil(o.SpCertificates) {
		return true
	}

	return false
}

// SetSpCertificates gets a reference to the given []SsoCertificateDto and assigns it to the SpCertificates field.
func (o *SsoSettingsDto) SetSpCertificates(v []SsoCertificateDto) {
	o.SpCertificates = v
}

// GetSpCertificateAdvanced returns the SpCertificateAdvanced field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetSpCertificateAdvanced() SsoSpCertificateAdvancedDto {
	if o == nil || IsNil(o.SpCertificateAdvanced) {
		var ret SsoSpCertificateAdvancedDto
		return ret
	}
	return *o.SpCertificateAdvanced
}

// GetSpCertificateAdvancedOk returns a tuple with the SpCertificateAdvanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetSpCertificateAdvancedOk() (*SsoSpCertificateAdvancedDto, bool) {
	if o == nil || IsNil(o.SpCertificateAdvanced) {
		return nil, false
	}
	return o.SpCertificateAdvanced, true
}

// HasSpCertificateAdvanced returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsSpCertificateAdvancedSet() bool {
	if o != nil && !IsNil(o.SpCertificateAdvanced) {
		return true
	}

	return false
}

// SetSpCertificateAdvanced gets a reference to the given SsoSpCertificateAdvancedDto and assigns it to the SpCertificateAdvanced field.
func (o *SsoSettingsDto) SetSpCertificateAdvanced(v SsoSpCertificateAdvancedDto) {
	o.SpCertificateAdvanced = &v
}

// GetFieldMapping returns the FieldMapping field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetFieldMapping() SsoFieldMappingDto {
	if o == nil || IsNil(o.FieldMapping) {
		var ret SsoFieldMappingDto
		return ret
	}
	return *o.FieldMapping
}

// GetFieldMappingOk returns a tuple with the FieldMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetFieldMappingOk() (*SsoFieldMappingDto, bool) {
	if o == nil || IsNil(o.FieldMapping) {
		return nil, false
	}
	return o.FieldMapping, true
}

// HasFieldMapping returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsFieldMappingSet() bool {
	if o != nil && !IsNil(o.FieldMapping) {
		return true
	}

	return false
}

// SetFieldMapping gets a reference to the given SsoFieldMappingDto and assigns it to the FieldMapping field.
func (o *SsoSettingsDto) SetFieldMapping(v SsoFieldMappingDto) {
	o.FieldMapping = &v
}

// GetHideAuthPage returns the HideAuthPage field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetHideAuthPage() bool {
	if o == nil || IsNil(o.HideAuthPage) {
		var ret bool
		return ret
	}
	return *o.HideAuthPage
}

// GetHideAuthPageOk returns a tuple with the HideAuthPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetHideAuthPageOk() (*bool, bool) {
	if o == nil || IsNil(o.HideAuthPage) {
		return nil, false
	}
	return o.HideAuthPage, true
}

// HasHideAuthPage returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsHideAuthPageSet() bool {
	if o != nil && !IsNil(o.HideAuthPage) {
		return true
	}

	return false
}

// SetHideAuthPage gets a reference to the given bool and assigns it to the HideAuthPage field.
func (o *SsoSettingsDto) SetHideAuthPage(v bool) {
	o.HideAuthPage = &v
}

// GetUsersType returns the UsersType field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetUsersType() int32 {
	if o == nil || IsNil(o.UsersType) {
		var ret int32
		return ret
	}
	return *o.UsersType
}

// GetUsersTypeOk returns a tuple with the UsersType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetUsersTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.UsersType) {
		return nil, false
	}
	return o.UsersType, true
}

// HasUsersType returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsUsersTypeSet() bool {
	if o != nil && !IsNil(o.UsersType) {
		return true
	}

	return false
}

// SetUsersType gets a reference to the given int32 and assigns it to the UsersType field.
func (o *SsoSettingsDto) SetUsersType(v int32) {
	o.UsersType = &v
}

// GetDisableEmailVerification returns the DisableEmailVerification field value if set, zero value otherwise.
func (o *SsoSettingsDto) GetDisableEmailVerification() bool {
	if o == nil || IsNil(o.DisableEmailVerification) {
		var ret bool
		return ret
	}
	return *o.DisableEmailVerification
}

// GetDisableEmailVerificationOk returns a tuple with the DisableEmailVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsDto) GetDisableEmailVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableEmailVerification) {
		return nil, false
	}
	return o.DisableEmailVerification, true
}

// HasDisableEmailVerification returns a boolean if a field has been set.
func (o *SsoSettingsDto) IsDisableEmailVerificationSet() bool {
	if o != nil && !IsNil(o.DisableEmailVerification) {
		return true
	}

	return false
}

// SetDisableEmailVerification gets a reference to the given bool and assigns it to the DisableEmailVerification field.
func (o *SsoSettingsDto) SetDisableEmailVerification(v bool) {
	o.DisableEmailVerification = &v
}

func (o SsoSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.EnableSso.IsSet() {
		toSerialize["enableSso"] = o.EnableSso.Get()
	}
	if !IsNil(o.IdpSettings) {
		toSerialize["idpSettings"] = o.IdpSettings
	}
	if o.IdpCertificates != nil {
		toSerialize["idpCertificates"] = o.IdpCertificates
	}
	if !IsNil(o.IdpCertificateAdvanced) {
		toSerialize["idpCertificateAdvanced"] = o.IdpCertificateAdvanced
	}
	if o.SpLoginLabel.IsSet() {
		toSerialize["spLoginLabel"] = o.SpLoginLabel.Get()
	}
	if o.SpCertificates != nil {
		toSerialize["spCertificates"] = o.SpCertificates
	}
	if !IsNil(o.SpCertificateAdvanced) {
		toSerialize["spCertificateAdvanced"] = o.SpCertificateAdvanced
	}
	if !IsNil(o.FieldMapping) {
		toSerialize["fieldMapping"] = o.FieldMapping
	}
	if !IsNil(o.HideAuthPage) {
		toSerialize["hideAuthPage"] = o.HideAuthPage
	}
	if !IsNil(o.UsersType) {
		toSerialize["usersType"] = o.UsersType
	}
	if !IsNil(o.DisableEmailVerification) {
		toSerialize["disableEmailVerification"] = o.DisableEmailVerification
	}
	return toSerialize, nil
}

type NullableSsoSettingsDto struct {
	value *SsoSettingsDto
	isSet bool
}

func (v NullableSsoSettingsDto) Get() *SsoSettingsDto {
	return v.value
}

func (v *NullableSsoSettingsDto) Set(val *SsoSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSettingsDto(val *SsoSettingsDto) *NullableSsoSettingsDto {
	return &NullableSsoSettingsDto{value: val, isSet: true}
}

func (v NullableSsoSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

