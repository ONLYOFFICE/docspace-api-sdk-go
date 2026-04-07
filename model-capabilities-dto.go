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
	"bytes"
	"fmt"
)

// checks if the CapabilitiesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesDto{}

// CapabilitiesDto The capabilities parameters.
type CapabilitiesDto struct {
	// Specifies if the LDAP settings are enabled or not.
	LdapEnabled bool `json:"ldapEnabled"`
	// The LDAP domain.
	LdapDomain NullableString `json:"ldapDomain,omitempty"`
	// The list of providers.
	Providers []string `json:"providers"`
	// The SP login label.
	SsoLabel NullableString `json:"ssoLabel"`
	// Specifies if OAuth is enabled or not.
	OauthEnabled bool `json:"oauthEnabled"`
	// The SSO URL. If this parameter is empty, then the SSO settings are disabled.
	SsoUrl NullableString `json:"ssoUrl"`
	// Specifies if an identity server is enabled or not.
	IdentityServerEnabled bool `json:"identityServerEnabled"`
}

type _CapabilitiesDto CapabilitiesDto

// NewCapabilitiesDto instantiates a new CapabilitiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesDto(ldapEnabled bool, providers []string, ssoLabel NullableString, oauthEnabled bool, ssoUrl NullableString, identityServerEnabled bool) *CapabilitiesDto {
	this := CapabilitiesDto{}
	this.LdapEnabled = ldapEnabled
	this.Providers = providers
	this.SsoLabel = ssoLabel
	this.OauthEnabled = oauthEnabled
	this.SsoUrl = ssoUrl
	this.IdentityServerEnabled = identityServerEnabled
	return &this
}

// NewCapabilitiesDtoWithDefaults instantiates a new CapabilitiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesDtoWithDefaults() *CapabilitiesDto {
	this := CapabilitiesDto{}
	return &this
}

// GetLdapEnabled returns the LdapEnabled field value
func (o *CapabilitiesDto) GetLdapEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LdapEnabled
}

// GetLdapEnabledOk returns a tuple with the LdapEnabled field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesDto) GetLdapEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapEnabled, true
}

// SetLdapEnabled sets field value
func (o *CapabilitiesDto) SetLdapEnabled(v bool) {
	o.LdapEnabled = v
}

// GetLdapDomain returns the LdapDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilitiesDto) GetLdapDomain() string {
	if o == nil || IsNil(o.LdapDomain.Get()) {
		var ret string
		return ret
	}
	return *o.LdapDomain.Get()
}

// GetLdapDomainOk returns a tuple with the LdapDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilitiesDto) GetLdapDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LdapDomain.Get(), o.LdapDomain.IsSet()
}

// HasLdapDomain returns a boolean if a field has been set.
func (o *CapabilitiesDto) IsLdapDomainSet() bool {
	if o != nil && o.LdapDomain.IsSet() {
		return true
	}

	return false
}

// SetLdapDomain gets a reference to the given NullableString and assigns it to the LdapDomain field.
func (o *CapabilitiesDto) SetLdapDomain(v string) {
	o.LdapDomain.Set(&v)
}
// SetLdapDomainNil sets the value for LdapDomain to be an explicit nil
func (o *CapabilitiesDto) SetLdapDomainNil() {
	o.LdapDomain.Set(nil)
}

// UnsetLdapDomain ensures that no value is present for LdapDomain, not even an explicit nil
func (o *CapabilitiesDto) UnsetLdapDomain() {
	o.LdapDomain.Unset()
}

// GetProviders returns the Providers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CapabilitiesDto) GetProviders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilitiesDto) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *CapabilitiesDto) SetProviders(v []string) {
	o.Providers = v
}

// GetSsoLabel returns the SsoLabel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CapabilitiesDto) GetSsoLabel() string {
	if o == nil || o.SsoLabel.Get() == nil {
		var ret string
		return ret
	}

	return *o.SsoLabel.Get()
}

// GetSsoLabelOk returns a tuple with the SsoLabel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilitiesDto) GetSsoLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoLabel.Get(), o.SsoLabel.IsSet()
}

// SetSsoLabel sets field value
func (o *CapabilitiesDto) SetSsoLabel(v string) {
	o.SsoLabel.Set(&v)
}

// GetOauthEnabled returns the OauthEnabled field value
func (o *CapabilitiesDto) GetOauthEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OauthEnabled
}

// GetOauthEnabledOk returns a tuple with the OauthEnabled field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesDto) GetOauthEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OauthEnabled, true
}

// SetOauthEnabled sets field value
func (o *CapabilitiesDto) SetOauthEnabled(v bool) {
	o.OauthEnabled = v
}

// GetSsoUrl returns the SsoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CapabilitiesDto) GetSsoUrl() string {
	if o == nil || o.SsoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilitiesDto) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// SetSsoUrl sets field value
func (o *CapabilitiesDto) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}

// GetIdentityServerEnabled returns the IdentityServerEnabled field value
func (o *CapabilitiesDto) GetIdentityServerEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IdentityServerEnabled
}

// GetIdentityServerEnabledOk returns a tuple with the IdentityServerEnabled field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesDto) GetIdentityServerEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityServerEnabled, true
}

// SetIdentityServerEnabled sets field value
func (o *CapabilitiesDto) SetIdentityServerEnabled(v bool) {
	o.IdentityServerEnabled = v
}

func (o CapabilitiesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ldapEnabled"] = o.LdapEnabled
	if o.LdapDomain.IsSet() {
		toSerialize["ldapDomain"] = o.LdapDomain.Get()
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	toSerialize["ssoLabel"] = o.SsoLabel.Get()
	toSerialize["oauthEnabled"] = o.OauthEnabled
	toSerialize["ssoUrl"] = o.SsoUrl.Get()
	toSerialize["identityServerEnabled"] = o.IdentityServerEnabled
	return toSerialize, nil
}

func (o *CapabilitiesDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ldapEnabled",
		"providers",
		"ssoLabel",
		"oauthEnabled",
		"ssoUrl",
		"identityServerEnabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCapabilitiesDto := _CapabilitiesDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCapabilitiesDto)

	if err != nil {
		return err
	}

	*o = CapabilitiesDto(varCapabilitiesDto)

	return err
}

type NullableCapabilitiesDto struct {
	value *CapabilitiesDto
	isSet bool
}

func (v NullableCapabilitiesDto) Get() *CapabilitiesDto {
	return v.value
}

func (v *NullableCapabilitiesDto) Set(val *CapabilitiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesDto(val *CapabilitiesDto) *NullableCapabilitiesDto {
	return &NullableCapabilitiesDto{value: val, isSet: true}
}

func (v NullableCapabilitiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

