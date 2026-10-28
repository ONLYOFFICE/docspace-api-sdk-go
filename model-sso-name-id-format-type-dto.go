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

// checks if the SsoNameIdFormatTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoNameIdFormatTypeDto{}

// SsoNameIdFormatTypeDto The SAML name ID formats the SSO settings accept.
type SsoNameIdFormatTypeDto struct {
	// The SAML 1.1 unspecified name ID format.
	Saml11Unspecified NullableString `json:"saml11Unspecified,omitempty"`
	// The SAML 1.1 email address name ID format.
	Saml11EmailAddress NullableString `json:"saml11EmailAddress,omitempty"`
	// The SAML 2.0 entity name ID format.
	Saml20Entity NullableString `json:"saml20Entity,omitempty"`
	// The SAML 2.0 transient name ID format, whose identifier differs from one session to the next. It is what  the built-in configuration uses.
	Saml20Transient NullableString `json:"saml20Transient,omitempty"`
	// The SAML 2.0 persistent name ID format, whose identifier stays the same for one person across sessions.
	Saml20Persistent NullableString `json:"saml20Persistent,omitempty"`
	// The SAML 2.0 encrypted name ID format.
	Saml20Encrypted NullableString `json:"saml20Encrypted,omitempty"`
	// The SAML 2.0 unspecified name ID format.
	Saml20Unspecified NullableString `json:"saml20Unspecified,omitempty"`
	// The SAML 1.1 X.509 subject name name ID format.
	Saml11X509SubjectName NullableString `json:"saml11X509SubjectName,omitempty"`
	// The SAML 1.1 Windows domain qualified name name ID format.
	Saml11WindowsDomainQualifiedName NullableString `json:"saml11WindowsDomainQualifiedName,omitempty"`
	// The SAML 2.0 Kerberos name ID format.
	Saml20Kerberos NullableString `json:"saml20Kerberos,omitempty"`
}

// NewSsoNameIdFormatTypeDto instantiates a new SsoNameIdFormatTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoNameIdFormatTypeDto() *SsoNameIdFormatTypeDto {
	this := SsoNameIdFormatTypeDto{}
	return &this
}

// NewSsoNameIdFormatTypeDtoWithDefaults instantiates a new SsoNameIdFormatTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoNameIdFormatTypeDtoWithDefaults() *SsoNameIdFormatTypeDto {
	this := SsoNameIdFormatTypeDto{}
	return &this
}

// GetSaml11Unspecified returns the Saml11Unspecified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml11Unspecified() string {
	if o == nil || IsNil(o.Saml11Unspecified.Get()) {
		var ret string
		return ret
	}
	return *o.Saml11Unspecified.Get()
}

// GetSaml11UnspecifiedOk returns a tuple with the Saml11Unspecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml11UnspecifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml11Unspecified.Get(), o.Saml11Unspecified.IsSet()
}

// HasSaml11Unspecified returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml11UnspecifiedSet() bool {
	if o != nil && o.Saml11Unspecified.IsSet() {
		return true
	}

	return false
}

// SetSaml11Unspecified gets a reference to the given NullableString and assigns it to the Saml11Unspecified field.
func (o *SsoNameIdFormatTypeDto) SetSaml11Unspecified(v string) {
	o.Saml11Unspecified.Set(&v)
}
// SetSaml11UnspecifiedNil sets the value for Saml11Unspecified to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml11UnspecifiedNil() {
	o.Saml11Unspecified.Set(nil)
}

// UnsetSaml11Unspecified ensures that no value is present for Saml11Unspecified, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml11Unspecified() {
	o.Saml11Unspecified.Unset()
}

// GetSaml11EmailAddress returns the Saml11EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml11EmailAddress() string {
	if o == nil || IsNil(o.Saml11EmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.Saml11EmailAddress.Get()
}

// GetSaml11EmailAddressOk returns a tuple with the Saml11EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml11EmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml11EmailAddress.Get(), o.Saml11EmailAddress.IsSet()
}

// HasSaml11EmailAddress returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml11EmailAddressSet() bool {
	if o != nil && o.Saml11EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetSaml11EmailAddress gets a reference to the given NullableString and assigns it to the Saml11EmailAddress field.
func (o *SsoNameIdFormatTypeDto) SetSaml11EmailAddress(v string) {
	o.Saml11EmailAddress.Set(&v)
}
// SetSaml11EmailAddressNil sets the value for Saml11EmailAddress to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml11EmailAddressNil() {
	o.Saml11EmailAddress.Set(nil)
}

// UnsetSaml11EmailAddress ensures that no value is present for Saml11EmailAddress, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml11EmailAddress() {
	o.Saml11EmailAddress.Unset()
}

// GetSaml20Entity returns the Saml20Entity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Entity() string {
	if o == nil || IsNil(o.Saml20Entity.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Entity.Get()
}

// GetSaml20EntityOk returns a tuple with the Saml20Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20EntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Entity.Get(), o.Saml20Entity.IsSet()
}

// HasSaml20Entity returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20EntitySet() bool {
	if o != nil && o.Saml20Entity.IsSet() {
		return true
	}

	return false
}

// SetSaml20Entity gets a reference to the given NullableString and assigns it to the Saml20Entity field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Entity(v string) {
	o.Saml20Entity.Set(&v)
}
// SetSaml20EntityNil sets the value for Saml20Entity to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20EntityNil() {
	o.Saml20Entity.Set(nil)
}

// UnsetSaml20Entity ensures that no value is present for Saml20Entity, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Entity() {
	o.Saml20Entity.Unset()
}

// GetSaml20Transient returns the Saml20Transient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Transient() string {
	if o == nil || IsNil(o.Saml20Transient.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Transient.Get()
}

// GetSaml20TransientOk returns a tuple with the Saml20Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20TransientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Transient.Get(), o.Saml20Transient.IsSet()
}

// HasSaml20Transient returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20TransientSet() bool {
	if o != nil && o.Saml20Transient.IsSet() {
		return true
	}

	return false
}

// SetSaml20Transient gets a reference to the given NullableString and assigns it to the Saml20Transient field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Transient(v string) {
	o.Saml20Transient.Set(&v)
}
// SetSaml20TransientNil sets the value for Saml20Transient to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20TransientNil() {
	o.Saml20Transient.Set(nil)
}

// UnsetSaml20Transient ensures that no value is present for Saml20Transient, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Transient() {
	o.Saml20Transient.Unset()
}

// GetSaml20Persistent returns the Saml20Persistent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Persistent() string {
	if o == nil || IsNil(o.Saml20Persistent.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Persistent.Get()
}

// GetSaml20PersistentOk returns a tuple with the Saml20Persistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20PersistentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Persistent.Get(), o.Saml20Persistent.IsSet()
}

// HasSaml20Persistent returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20PersistentSet() bool {
	if o != nil && o.Saml20Persistent.IsSet() {
		return true
	}

	return false
}

// SetSaml20Persistent gets a reference to the given NullableString and assigns it to the Saml20Persistent field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Persistent(v string) {
	o.Saml20Persistent.Set(&v)
}
// SetSaml20PersistentNil sets the value for Saml20Persistent to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20PersistentNil() {
	o.Saml20Persistent.Set(nil)
}

// UnsetSaml20Persistent ensures that no value is present for Saml20Persistent, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Persistent() {
	o.Saml20Persistent.Unset()
}

// GetSaml20Encrypted returns the Saml20Encrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Encrypted() string {
	if o == nil || IsNil(o.Saml20Encrypted.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Encrypted.Get()
}

// GetSaml20EncryptedOk returns a tuple with the Saml20Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20EncryptedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Encrypted.Get(), o.Saml20Encrypted.IsSet()
}

// HasSaml20Encrypted returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20EncryptedSet() bool {
	if o != nil && o.Saml20Encrypted.IsSet() {
		return true
	}

	return false
}

// SetSaml20Encrypted gets a reference to the given NullableString and assigns it to the Saml20Encrypted field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Encrypted(v string) {
	o.Saml20Encrypted.Set(&v)
}
// SetSaml20EncryptedNil sets the value for Saml20Encrypted to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20EncryptedNil() {
	o.Saml20Encrypted.Set(nil)
}

// UnsetSaml20Encrypted ensures that no value is present for Saml20Encrypted, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Encrypted() {
	o.Saml20Encrypted.Unset()
}

// GetSaml20Unspecified returns the Saml20Unspecified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Unspecified() string {
	if o == nil || IsNil(o.Saml20Unspecified.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Unspecified.Get()
}

// GetSaml20UnspecifiedOk returns a tuple with the Saml20Unspecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20UnspecifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Unspecified.Get(), o.Saml20Unspecified.IsSet()
}

// HasSaml20Unspecified returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20UnspecifiedSet() bool {
	if o != nil && o.Saml20Unspecified.IsSet() {
		return true
	}

	return false
}

// SetSaml20Unspecified gets a reference to the given NullableString and assigns it to the Saml20Unspecified field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Unspecified(v string) {
	o.Saml20Unspecified.Set(&v)
}
// SetSaml20UnspecifiedNil sets the value for Saml20Unspecified to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20UnspecifiedNil() {
	o.Saml20Unspecified.Set(nil)
}

// UnsetSaml20Unspecified ensures that no value is present for Saml20Unspecified, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Unspecified() {
	o.Saml20Unspecified.Unset()
}

// GetSaml11X509SubjectName returns the Saml11X509SubjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml11X509SubjectName() string {
	if o == nil || IsNil(o.Saml11X509SubjectName.Get()) {
		var ret string
		return ret
	}
	return *o.Saml11X509SubjectName.Get()
}

// GetSaml11X509SubjectNameOk returns a tuple with the Saml11X509SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml11X509SubjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml11X509SubjectName.Get(), o.Saml11X509SubjectName.IsSet()
}

// HasSaml11X509SubjectName returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml11X509SubjectNameSet() bool {
	if o != nil && o.Saml11X509SubjectName.IsSet() {
		return true
	}

	return false
}

// SetSaml11X509SubjectName gets a reference to the given NullableString and assigns it to the Saml11X509SubjectName field.
func (o *SsoNameIdFormatTypeDto) SetSaml11X509SubjectName(v string) {
	o.Saml11X509SubjectName.Set(&v)
}
// SetSaml11X509SubjectNameNil sets the value for Saml11X509SubjectName to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml11X509SubjectNameNil() {
	o.Saml11X509SubjectName.Set(nil)
}

// UnsetSaml11X509SubjectName ensures that no value is present for Saml11X509SubjectName, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml11X509SubjectName() {
	o.Saml11X509SubjectName.Unset()
}

// GetSaml11WindowsDomainQualifiedName returns the Saml11WindowsDomainQualifiedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml11WindowsDomainQualifiedName() string {
	if o == nil || IsNil(o.Saml11WindowsDomainQualifiedName.Get()) {
		var ret string
		return ret
	}
	return *o.Saml11WindowsDomainQualifiedName.Get()
}

// GetSaml11WindowsDomainQualifiedNameOk returns a tuple with the Saml11WindowsDomainQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml11WindowsDomainQualifiedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml11WindowsDomainQualifiedName.Get(), o.Saml11WindowsDomainQualifiedName.IsSet()
}

// HasSaml11WindowsDomainQualifiedName returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml11WindowsDomainQualifiedNameSet() bool {
	if o != nil && o.Saml11WindowsDomainQualifiedName.IsSet() {
		return true
	}

	return false
}

// SetSaml11WindowsDomainQualifiedName gets a reference to the given NullableString and assigns it to the Saml11WindowsDomainQualifiedName field.
func (o *SsoNameIdFormatTypeDto) SetSaml11WindowsDomainQualifiedName(v string) {
	o.Saml11WindowsDomainQualifiedName.Set(&v)
}
// SetSaml11WindowsDomainQualifiedNameNil sets the value for Saml11WindowsDomainQualifiedName to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml11WindowsDomainQualifiedNameNil() {
	o.Saml11WindowsDomainQualifiedName.Set(nil)
}

// UnsetSaml11WindowsDomainQualifiedName ensures that no value is present for Saml11WindowsDomainQualifiedName, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml11WindowsDomainQualifiedName() {
	o.Saml11WindowsDomainQualifiedName.Unset()
}

// GetSaml20Kerberos returns the Saml20Kerberos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoNameIdFormatTypeDto) GetSaml20Kerberos() string {
	if o == nil || IsNil(o.Saml20Kerberos.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20Kerberos.Get()
}

// GetSaml20KerberosOk returns a tuple with the Saml20Kerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoNameIdFormatTypeDto) GetSaml20KerberosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20Kerberos.Get(), o.Saml20Kerberos.IsSet()
}

// HasSaml20Kerberos returns a boolean if a field has been set.
func (o *SsoNameIdFormatTypeDto) IsSaml20KerberosSet() bool {
	if o != nil && o.Saml20Kerberos.IsSet() {
		return true
	}

	return false
}

// SetSaml20Kerberos gets a reference to the given NullableString and assigns it to the Saml20Kerberos field.
func (o *SsoNameIdFormatTypeDto) SetSaml20Kerberos(v string) {
	o.Saml20Kerberos.Set(&v)
}
// SetSaml20KerberosNil sets the value for Saml20Kerberos to be an explicit nil
func (o *SsoNameIdFormatTypeDto) SetSaml20KerberosNil() {
	o.Saml20Kerberos.Set(nil)
}

// UnsetSaml20Kerberos ensures that no value is present for Saml20Kerberos, not even an explicit nil
func (o *SsoNameIdFormatTypeDto) UnsetSaml20Kerberos() {
	o.Saml20Kerberos.Unset()
}

func (o SsoNameIdFormatTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoNameIdFormatTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Saml11Unspecified.IsSet() {
		toSerialize["saml11Unspecified"] = o.Saml11Unspecified.Get()
	}
	if o.Saml11EmailAddress.IsSet() {
		toSerialize["saml11EmailAddress"] = o.Saml11EmailAddress.Get()
	}
	if o.Saml20Entity.IsSet() {
		toSerialize["saml20Entity"] = o.Saml20Entity.Get()
	}
	if o.Saml20Transient.IsSet() {
		toSerialize["saml20Transient"] = o.Saml20Transient.Get()
	}
	if o.Saml20Persistent.IsSet() {
		toSerialize["saml20Persistent"] = o.Saml20Persistent.Get()
	}
	if o.Saml20Encrypted.IsSet() {
		toSerialize["saml20Encrypted"] = o.Saml20Encrypted.Get()
	}
	if o.Saml20Unspecified.IsSet() {
		toSerialize["saml20Unspecified"] = o.Saml20Unspecified.Get()
	}
	if o.Saml11X509SubjectName.IsSet() {
		toSerialize["saml11X509SubjectName"] = o.Saml11X509SubjectName.Get()
	}
	if o.Saml11WindowsDomainQualifiedName.IsSet() {
		toSerialize["saml11WindowsDomainQualifiedName"] = o.Saml11WindowsDomainQualifiedName.Get()
	}
	if o.Saml20Kerberos.IsSet() {
		toSerialize["saml20Kerberos"] = o.Saml20Kerberos.Get()
	}
	return toSerialize, nil
}

type NullableSsoNameIdFormatTypeDto struct {
	value *SsoNameIdFormatTypeDto
	isSet bool
}

func (v NullableSsoNameIdFormatTypeDto) Get() *SsoNameIdFormatTypeDto {
	return v.value
}

func (v *NullableSsoNameIdFormatTypeDto) Set(val *SsoNameIdFormatTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoNameIdFormatTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoNameIdFormatTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoNameIdFormatTypeDto(val *SsoNameIdFormatTypeDto) *NullableSsoNameIdFormatTypeDto {
	return &NullableSsoNameIdFormatTypeDto{value: val, isSet: true}
}

func (v NullableSsoNameIdFormatTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoNameIdFormatTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

