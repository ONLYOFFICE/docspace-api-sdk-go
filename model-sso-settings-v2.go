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

// checks if the SsoSettingsV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSettingsV2{}

// SsoSettingsV2 The SSO portal settings.
type SsoSettingsV2 struct {
	// The timestamp indicating when the settings were last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// Specifies if the SSO settings are enabled or not.
	EnableSso NullableBool `json:"enableSso,omitempty"`
	IdpSettings *SsoIdpSettings `json:"idpSettings,omitempty"`
	// The list of the IdP certificates.
	IdpCertificates []SsoCertificate `json:"idpCertificates,omitempty"`
	IdpCertificateAdvanced *SsoIdpCertificateAdvanced `json:"idpCertificateAdvanced,omitempty"`
	// The SP login label.
	SpLoginLabel NullableString `json:"spLoginLabel,omitempty"`
	// The list of the SP certificates.
	SpCertificates []SsoCertificate `json:"spCertificates,omitempty"`
	SpCertificateAdvanced *SsoSpCertificateAdvanced `json:"spCertificateAdvanced,omitempty"`
	FieldMapping *SsoFieldMapping `json:"fieldMapping,omitempty"`
	// Specifies if the authentication page will be hidden or not.
	HideAuthPage *bool `json:"hideAuthPage,omitempty"`
	// The user type.
	UsersType *int32 `json:"usersType,omitempty"`
	// Specifies if the email verification is disabled or not.
	DisableEmailVerification *bool `json:"disableEmailVerification,omitempty"`
}

// NewSsoSettingsV2 instantiates a new SsoSettingsV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSettingsV2() *SsoSettingsV2 {
	this := SsoSettingsV2{}
	return &this
}

// NewSsoSettingsV2WithDefaults instantiates a new SsoSettingsV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSettingsV2WithDefaults() *SsoSettingsV2 {
	this := SsoSettingsV2{}
	return &this
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *SsoSettingsV2) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetEnableSso returns the EnableSso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsV2) GetEnableSso() bool {
	if o == nil || IsNil(o.EnableSso.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSso.Get()
}

// GetEnableSsoOk returns a tuple with the EnableSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsV2) GetEnableSsoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSso.Get(), o.EnableSso.IsSet()
}

// HasEnableSso returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsEnableSsoSet() bool {
	if o != nil && o.EnableSso.IsSet() {
		return true
	}

	return false
}

// SetEnableSso gets a reference to the given NullableBool and assigns it to the EnableSso field.
func (o *SsoSettingsV2) SetEnableSso(v bool) {
	o.EnableSso.Set(&v)
}
// SetEnableSsoNil sets the value for EnableSso to be an explicit nil
func (o *SsoSettingsV2) SetEnableSsoNil() {
	o.EnableSso.Set(nil)
}

// UnsetEnableSso ensures that no value is present for EnableSso, not even an explicit nil
func (o *SsoSettingsV2) UnsetEnableSso() {
	o.EnableSso.Unset()
}

// GetIdpSettings returns the IdpSettings field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetIdpSettings() SsoIdpSettings {
	if o == nil || IsNil(o.IdpSettings) {
		var ret SsoIdpSettings
		return ret
	}
	return *o.IdpSettings
}

// GetIdpSettingsOk returns a tuple with the IdpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetIdpSettingsOk() (*SsoIdpSettings, bool) {
	if o == nil || IsNil(o.IdpSettings) {
		return nil, false
	}
	return o.IdpSettings, true
}

// HasIdpSettings returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsIdpSettingsSet() bool {
	if o != nil && !IsNil(o.IdpSettings) {
		return true
	}

	return false
}

// SetIdpSettings gets a reference to the given SsoIdpSettings and assigns it to the IdpSettings field.
func (o *SsoSettingsV2) SetIdpSettings(v SsoIdpSettings) {
	o.IdpSettings = &v
}

// GetIdpCertificates returns the IdpCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsV2) GetIdpCertificates() []SsoCertificate {
	if o == nil {
		var ret []SsoCertificate
		return ret
	}
	return o.IdpCertificates
}

// GetIdpCertificatesOk returns a tuple with the IdpCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsV2) GetIdpCertificatesOk() ([]SsoCertificate, bool) {
	if o == nil || IsNil(o.IdpCertificates) {
		return nil, false
	}
	return o.IdpCertificates, true
}

// HasIdpCertificates returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsIdpCertificatesSet() bool {
	if o != nil && !IsNil(o.IdpCertificates) {
		return true
	}

	return false
}

// SetIdpCertificates gets a reference to the given []SsoCertificate and assigns it to the IdpCertificates field.
func (o *SsoSettingsV2) SetIdpCertificates(v []SsoCertificate) {
	o.IdpCertificates = v
}

// GetIdpCertificateAdvanced returns the IdpCertificateAdvanced field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetIdpCertificateAdvanced() SsoIdpCertificateAdvanced {
	if o == nil || IsNil(o.IdpCertificateAdvanced) {
		var ret SsoIdpCertificateAdvanced
		return ret
	}
	return *o.IdpCertificateAdvanced
}

// GetIdpCertificateAdvancedOk returns a tuple with the IdpCertificateAdvanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetIdpCertificateAdvancedOk() (*SsoIdpCertificateAdvanced, bool) {
	if o == nil || IsNil(o.IdpCertificateAdvanced) {
		return nil, false
	}
	return o.IdpCertificateAdvanced, true
}

// HasIdpCertificateAdvanced returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsIdpCertificateAdvancedSet() bool {
	if o != nil && !IsNil(o.IdpCertificateAdvanced) {
		return true
	}

	return false
}

// SetIdpCertificateAdvanced gets a reference to the given SsoIdpCertificateAdvanced and assigns it to the IdpCertificateAdvanced field.
func (o *SsoSettingsV2) SetIdpCertificateAdvanced(v SsoIdpCertificateAdvanced) {
	o.IdpCertificateAdvanced = &v
}

// GetSpLoginLabel returns the SpLoginLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsV2) GetSpLoginLabel() string {
	if o == nil || IsNil(o.SpLoginLabel.Get()) {
		var ret string
		return ret
	}
	return *o.SpLoginLabel.Get()
}

// GetSpLoginLabelOk returns a tuple with the SpLoginLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsV2) GetSpLoginLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpLoginLabel.Get(), o.SpLoginLabel.IsSet()
}

// HasSpLoginLabel returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsSpLoginLabelSet() bool {
	if o != nil && o.SpLoginLabel.IsSet() {
		return true
	}

	return false
}

// SetSpLoginLabel gets a reference to the given NullableString and assigns it to the SpLoginLabel field.
func (o *SsoSettingsV2) SetSpLoginLabel(v string) {
	o.SpLoginLabel.Set(&v)
}
// SetSpLoginLabelNil sets the value for SpLoginLabel to be an explicit nil
func (o *SsoSettingsV2) SetSpLoginLabelNil() {
	o.SpLoginLabel.Set(nil)
}

// UnsetSpLoginLabel ensures that no value is present for SpLoginLabel, not even an explicit nil
func (o *SsoSettingsV2) UnsetSpLoginLabel() {
	o.SpLoginLabel.Unset()
}

// GetSpCertificates returns the SpCertificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoSettingsV2) GetSpCertificates() []SsoCertificate {
	if o == nil {
		var ret []SsoCertificate
		return ret
	}
	return o.SpCertificates
}

// GetSpCertificatesOk returns a tuple with the SpCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoSettingsV2) GetSpCertificatesOk() ([]SsoCertificate, bool) {
	if o == nil || IsNil(o.SpCertificates) {
		return nil, false
	}
	return o.SpCertificates, true
}

// HasSpCertificates returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsSpCertificatesSet() bool {
	if o != nil && !IsNil(o.SpCertificates) {
		return true
	}

	return false
}

// SetSpCertificates gets a reference to the given []SsoCertificate and assigns it to the SpCertificates field.
func (o *SsoSettingsV2) SetSpCertificates(v []SsoCertificate) {
	o.SpCertificates = v
}

// GetSpCertificateAdvanced returns the SpCertificateAdvanced field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetSpCertificateAdvanced() SsoSpCertificateAdvanced {
	if o == nil || IsNil(o.SpCertificateAdvanced) {
		var ret SsoSpCertificateAdvanced
		return ret
	}
	return *o.SpCertificateAdvanced
}

// GetSpCertificateAdvancedOk returns a tuple with the SpCertificateAdvanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetSpCertificateAdvancedOk() (*SsoSpCertificateAdvanced, bool) {
	if o == nil || IsNil(o.SpCertificateAdvanced) {
		return nil, false
	}
	return o.SpCertificateAdvanced, true
}

// HasSpCertificateAdvanced returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsSpCertificateAdvancedSet() bool {
	if o != nil && !IsNil(o.SpCertificateAdvanced) {
		return true
	}

	return false
}

// SetSpCertificateAdvanced gets a reference to the given SsoSpCertificateAdvanced and assigns it to the SpCertificateAdvanced field.
func (o *SsoSettingsV2) SetSpCertificateAdvanced(v SsoSpCertificateAdvanced) {
	o.SpCertificateAdvanced = &v
}

// GetFieldMapping returns the FieldMapping field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetFieldMapping() SsoFieldMapping {
	if o == nil || IsNil(o.FieldMapping) {
		var ret SsoFieldMapping
		return ret
	}
	return *o.FieldMapping
}

// GetFieldMappingOk returns a tuple with the FieldMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetFieldMappingOk() (*SsoFieldMapping, bool) {
	if o == nil || IsNil(o.FieldMapping) {
		return nil, false
	}
	return o.FieldMapping, true
}

// HasFieldMapping returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsFieldMappingSet() bool {
	if o != nil && !IsNil(o.FieldMapping) {
		return true
	}

	return false
}

// SetFieldMapping gets a reference to the given SsoFieldMapping and assigns it to the FieldMapping field.
func (o *SsoSettingsV2) SetFieldMapping(v SsoFieldMapping) {
	o.FieldMapping = &v
}

// GetHideAuthPage returns the HideAuthPage field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetHideAuthPage() bool {
	if o == nil || IsNil(o.HideAuthPage) {
		var ret bool
		return ret
	}
	return *o.HideAuthPage
}

// GetHideAuthPageOk returns a tuple with the HideAuthPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetHideAuthPageOk() (*bool, bool) {
	if o == nil || IsNil(o.HideAuthPage) {
		return nil, false
	}
	return o.HideAuthPage, true
}

// HasHideAuthPage returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsHideAuthPageSet() bool {
	if o != nil && !IsNil(o.HideAuthPage) {
		return true
	}

	return false
}

// SetHideAuthPage gets a reference to the given bool and assigns it to the HideAuthPage field.
func (o *SsoSettingsV2) SetHideAuthPage(v bool) {
	o.HideAuthPage = &v
}

// GetUsersType returns the UsersType field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetUsersType() int32 {
	if o == nil || IsNil(o.UsersType) {
		var ret int32
		return ret
	}
	return *o.UsersType
}

// GetUsersTypeOk returns a tuple with the UsersType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetUsersTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.UsersType) {
		return nil, false
	}
	return o.UsersType, true
}

// HasUsersType returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsUsersTypeSet() bool {
	if o != nil && !IsNil(o.UsersType) {
		return true
	}

	return false
}

// SetUsersType gets a reference to the given int32 and assigns it to the UsersType field.
func (o *SsoSettingsV2) SetUsersType(v int32) {
	o.UsersType = &v
}

// GetDisableEmailVerification returns the DisableEmailVerification field value if set, zero value otherwise.
func (o *SsoSettingsV2) GetDisableEmailVerification() bool {
	if o == nil || IsNil(o.DisableEmailVerification) {
		var ret bool
		return ret
	}
	return *o.DisableEmailVerification
}

// GetDisableEmailVerificationOk returns a tuple with the DisableEmailVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSettingsV2) GetDisableEmailVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableEmailVerification) {
		return nil, false
	}
	return o.DisableEmailVerification, true
}

// HasDisableEmailVerification returns a boolean if a field has been set.
func (o *SsoSettingsV2) IsDisableEmailVerificationSet() bool {
	if o != nil && !IsNil(o.DisableEmailVerification) {
		return true
	}

	return false
}

// SetDisableEmailVerification gets a reference to the given bool and assigns it to the DisableEmailVerification field.
func (o *SsoSettingsV2) SetDisableEmailVerification(v bool) {
	o.DisableEmailVerification = &v
}

func (o SsoSettingsV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSettingsV2) ToMap() (map[string]interface{}, error) {
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

type NullableSsoSettingsV2 struct {
	value *SsoSettingsV2
	isSet bool
}

func (v NullableSsoSettingsV2) Get() *SsoSettingsV2 {
	return v.value
}

func (v *NullableSsoSettingsV2) Set(val *SsoSettingsV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSettingsV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSettingsV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSettingsV2(val *SsoSettingsV2) *NullableSsoSettingsV2 {
	return &NullableSsoSettingsV2{value: val, isSet: true}
}

func (v NullableSsoSettingsV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSettingsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

