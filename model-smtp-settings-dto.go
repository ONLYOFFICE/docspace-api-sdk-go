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

// checks if the SmtpSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpSettingsDto{}

// SmtpSettingsDto The mail server the portal sends its letters through.
type SmtpSettingsDto struct {
	// The host name or address of the mail server. On a cloud portal that has saved no relay of its own every  field of this object comes back empty, because the installation's own server is not disclosed - only  `isDefaultSettings` is set there.
	Host NullableString `json:"host,omitempty"`
	// The port the mail server is reached on - conventionally 25 or 587 without encryption from the start, 465  with it. It is empty when no port was stored, in which case the portal falls back to its own default.
	Port NullableInt32 `json:"port,omitempty"`
	// The address the letters are sent from, which appears in the From header and is what a reply goes to.
	SenderAddress NullableString `json:"senderAddress,omitempty"`
	// The name shown beside that address in a recipient's mailbox.
	SenderDisplayName NullableString `json:"senderDisplayName,omitempty"`
	// The account the portal signs in to the mail server as, meaningful only while `enableAuth` is `true`.
	CredentialsUserName NullableString `json:"credentialsUserName,omitempty"`
	// Always empty here: the stored password is never returned, so a client that sends these settings back has  to supply it again rather than echoing what it read.
	CredentialsUserPassword NullableString `json:"credentialsUserPassword,omitempty"`
	// Whether the connection to the mail server is encrypted.
	EnableSSL *bool `json:"enableSSL,omitempty"`
	// Whether the portal signs in to the mail server at all. While it is `false` the credentials above are  ignored and the server is expected to accept mail unauthenticated.
	EnableAuth *bool `json:"enableAuth,omitempty"`
	// Always `false` here: the flag is accepted when settings are saved but is not stored, so it never comes  back set and says nothing about how the portal authenticates.
	UseNtlm *bool `json:"useNtlm,omitempty"`
	// Whether the portal is still on the mail configuration of the installation rather than on a relay of its  own. `DELETE api/2.0/smtpsettings/smtp` puts it back to `true`, and while it is `true` on a cloud portal  the fields above are blank rather than showing the installation's server.
	IsDefaultSettings *bool `json:"isDefaultSettings,omitempty"`
}

// NewSmtpSettingsDto instantiates a new SmtpSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpSettingsDto() *SmtpSettingsDto {
	this := SmtpSettingsDto{}
	return &this
}

// NewSmtpSettingsDtoWithDefaults instantiates a new SmtpSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpSettingsDtoWithDefaults() *SmtpSettingsDto {
	this := SmtpSettingsDto{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsHostSet() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *SmtpSettingsDto) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *SmtpSettingsDto) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *SmtpSettingsDto) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsPortSet() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *SmtpSettingsDto) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *SmtpSettingsDto) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *SmtpSettingsDto) UnsetPort() {
	o.Port.Unset()
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetSenderAddress() string {
	if o == nil || IsNil(o.SenderAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SenderAddress.Get()
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetSenderAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderAddress.Get(), o.SenderAddress.IsSet()
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsSenderAddressSet() bool {
	if o != nil && o.SenderAddress.IsSet() {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given NullableString and assigns it to the SenderAddress field.
func (o *SmtpSettingsDto) SetSenderAddress(v string) {
	o.SenderAddress.Set(&v)
}
// SetSenderAddressNil sets the value for SenderAddress to be an explicit nil
func (o *SmtpSettingsDto) SetSenderAddressNil() {
	o.SenderAddress.Set(nil)
}

// UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
func (o *SmtpSettingsDto) UnsetSenderAddress() {
	o.SenderAddress.Unset()
}

// GetSenderDisplayName returns the SenderDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetSenderDisplayName() string {
	if o == nil || IsNil(o.SenderDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.SenderDisplayName.Get()
}

// GetSenderDisplayNameOk returns a tuple with the SenderDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetSenderDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderDisplayName.Get(), o.SenderDisplayName.IsSet()
}

// HasSenderDisplayName returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsSenderDisplayNameSet() bool {
	if o != nil && o.SenderDisplayName.IsSet() {
		return true
	}

	return false
}

// SetSenderDisplayName gets a reference to the given NullableString and assigns it to the SenderDisplayName field.
func (o *SmtpSettingsDto) SetSenderDisplayName(v string) {
	o.SenderDisplayName.Set(&v)
}
// SetSenderDisplayNameNil sets the value for SenderDisplayName to be an explicit nil
func (o *SmtpSettingsDto) SetSenderDisplayNameNil() {
	o.SenderDisplayName.Set(nil)
}

// UnsetSenderDisplayName ensures that no value is present for SenderDisplayName, not even an explicit nil
func (o *SmtpSettingsDto) UnsetSenderDisplayName() {
	o.SenderDisplayName.Unset()
}

// GetCredentialsUserName returns the CredentialsUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetCredentialsUserName() string {
	if o == nil || IsNil(o.CredentialsUserName.Get()) {
		var ret string
		return ret
	}
	return *o.CredentialsUserName.Get()
}

// GetCredentialsUserNameOk returns a tuple with the CredentialsUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetCredentialsUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialsUserName.Get(), o.CredentialsUserName.IsSet()
}

// HasCredentialsUserName returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsCredentialsUserNameSet() bool {
	if o != nil && o.CredentialsUserName.IsSet() {
		return true
	}

	return false
}

// SetCredentialsUserName gets a reference to the given NullableString and assigns it to the CredentialsUserName field.
func (o *SmtpSettingsDto) SetCredentialsUserName(v string) {
	o.CredentialsUserName.Set(&v)
}
// SetCredentialsUserNameNil sets the value for CredentialsUserName to be an explicit nil
func (o *SmtpSettingsDto) SetCredentialsUserNameNil() {
	o.CredentialsUserName.Set(nil)
}

// UnsetCredentialsUserName ensures that no value is present for CredentialsUserName, not even an explicit nil
func (o *SmtpSettingsDto) UnsetCredentialsUserName() {
	o.CredentialsUserName.Unset()
}

// GetCredentialsUserPassword returns the CredentialsUserPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpSettingsDto) GetCredentialsUserPassword() string {
	if o == nil || IsNil(o.CredentialsUserPassword.Get()) {
		var ret string
		return ret
	}
	return *o.CredentialsUserPassword.Get()
}

// GetCredentialsUserPasswordOk returns a tuple with the CredentialsUserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpSettingsDto) GetCredentialsUserPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialsUserPassword.Get(), o.CredentialsUserPassword.IsSet()
}

// HasCredentialsUserPassword returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsCredentialsUserPasswordSet() bool {
	if o != nil && o.CredentialsUserPassword.IsSet() {
		return true
	}

	return false
}

// SetCredentialsUserPassword gets a reference to the given NullableString and assigns it to the CredentialsUserPassword field.
func (o *SmtpSettingsDto) SetCredentialsUserPassword(v string) {
	o.CredentialsUserPassword.Set(&v)
}
// SetCredentialsUserPasswordNil sets the value for CredentialsUserPassword to be an explicit nil
func (o *SmtpSettingsDto) SetCredentialsUserPasswordNil() {
	o.CredentialsUserPassword.Set(nil)
}

// UnsetCredentialsUserPassword ensures that no value is present for CredentialsUserPassword, not even an explicit nil
func (o *SmtpSettingsDto) UnsetCredentialsUserPassword() {
	o.CredentialsUserPassword.Unset()
}

// GetEnableSSL returns the EnableSSL field value if set, zero value otherwise.
func (o *SmtpSettingsDto) GetEnableSSL() bool {
	if o == nil || IsNil(o.EnableSSL) {
		var ret bool
		return ret
	}
	return *o.EnableSSL
}

// GetEnableSSLOk returns a tuple with the EnableSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpSettingsDto) GetEnableSSLOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSSL) {
		return nil, false
	}
	return o.EnableSSL, true
}

// HasEnableSSL returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsEnableSSLSet() bool {
	if o != nil && !IsNil(o.EnableSSL) {
		return true
	}

	return false
}

// SetEnableSSL gets a reference to the given bool and assigns it to the EnableSSL field.
func (o *SmtpSettingsDto) SetEnableSSL(v bool) {
	o.EnableSSL = &v
}

// GetEnableAuth returns the EnableAuth field value if set, zero value otherwise.
func (o *SmtpSettingsDto) GetEnableAuth() bool {
	if o == nil || IsNil(o.EnableAuth) {
		var ret bool
		return ret
	}
	return *o.EnableAuth
}

// GetEnableAuthOk returns a tuple with the EnableAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpSettingsDto) GetEnableAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAuth) {
		return nil, false
	}
	return o.EnableAuth, true
}

// HasEnableAuth returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsEnableAuthSet() bool {
	if o != nil && !IsNil(o.EnableAuth) {
		return true
	}

	return false
}

// SetEnableAuth gets a reference to the given bool and assigns it to the EnableAuth field.
func (o *SmtpSettingsDto) SetEnableAuth(v bool) {
	o.EnableAuth = &v
}

// GetUseNtlm returns the UseNtlm field value if set, zero value otherwise.
func (o *SmtpSettingsDto) GetUseNtlm() bool {
	if o == nil || IsNil(o.UseNtlm) {
		var ret bool
		return ret
	}
	return *o.UseNtlm
}

// GetUseNtlmOk returns a tuple with the UseNtlm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpSettingsDto) GetUseNtlmOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNtlm) {
		return nil, false
	}
	return o.UseNtlm, true
}

// HasUseNtlm returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsUseNtlmSet() bool {
	if o != nil && !IsNil(o.UseNtlm) {
		return true
	}

	return false
}

// SetUseNtlm gets a reference to the given bool and assigns it to the UseNtlm field.
func (o *SmtpSettingsDto) SetUseNtlm(v bool) {
	o.UseNtlm = &v
}

// GetIsDefaultSettings returns the IsDefaultSettings field value if set, zero value otherwise.
func (o *SmtpSettingsDto) GetIsDefaultSettings() bool {
	if o == nil || IsNil(o.IsDefaultSettings) {
		var ret bool
		return ret
	}
	return *o.IsDefaultSettings
}

// GetIsDefaultSettingsOk returns a tuple with the IsDefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpSettingsDto) GetIsDefaultSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultSettings) {
		return nil, false
	}
	return o.IsDefaultSettings, true
}

// HasIsDefaultSettings returns a boolean if a field has been set.
func (o *SmtpSettingsDto) IsIsDefaultSettingsSet() bool {
	if o != nil && !IsNil(o.IsDefaultSettings) {
		return true
	}

	return false
}

// SetIsDefaultSettings gets a reference to the given bool and assigns it to the IsDefaultSettings field.
func (o *SmtpSettingsDto) SetIsDefaultSettings(v bool) {
	o.IsDefaultSettings = &v
}

func (o SmtpSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.SenderAddress.IsSet() {
		toSerialize["senderAddress"] = o.SenderAddress.Get()
	}
	if o.SenderDisplayName.IsSet() {
		toSerialize["senderDisplayName"] = o.SenderDisplayName.Get()
	}
	if o.CredentialsUserName.IsSet() {
		toSerialize["credentialsUserName"] = o.CredentialsUserName.Get()
	}
	if o.CredentialsUserPassword.IsSet() {
		toSerialize["credentialsUserPassword"] = o.CredentialsUserPassword.Get()
	}
	if !IsNil(o.EnableSSL) {
		toSerialize["enableSSL"] = o.EnableSSL
	}
	if !IsNil(o.EnableAuth) {
		toSerialize["enableAuth"] = o.EnableAuth
	}
	if !IsNil(o.UseNtlm) {
		toSerialize["useNtlm"] = o.UseNtlm
	}
	if !IsNil(o.IsDefaultSettings) {
		toSerialize["isDefaultSettings"] = o.IsDefaultSettings
	}
	return toSerialize, nil
}

type NullableSmtpSettingsDto struct {
	value *SmtpSettingsDto
	isSet bool
}

func (v NullableSmtpSettingsDto) Get() *SmtpSettingsDto {
	return v.value
}

func (v *NullableSmtpSettingsDto) Set(val *SmtpSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpSettingsDto(val *SmtpSettingsDto) *NullableSmtpSettingsDto {
	return &NullableSmtpSettingsDto{value: val, isSet: true}
}

func (v NullableSmtpSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

