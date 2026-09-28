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

// checks if the DocsCloudSecurityConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudSecurityConfig{}

// DocsCloudSecurityConfig Represents the security configuration of a Docs Connect tenant.
type DocsCloudSecurityConfig struct {
	// The security secret.
	Secret NullableString `json:"secret,omitempty"`
	// The security header name.
	Header NullableString `json:"header,omitempty"`
}

// NewDocsCloudSecurityConfig instantiates a new DocsCloudSecurityConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudSecurityConfig() *DocsCloudSecurityConfig {
	this := DocsCloudSecurityConfig{}
	return &this
}

// NewDocsCloudSecurityConfigWithDefaults instantiates a new DocsCloudSecurityConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudSecurityConfigWithDefaults() *DocsCloudSecurityConfig {
	this := DocsCloudSecurityConfig{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudSecurityConfig) GetSecret() string {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret string
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudSecurityConfig) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *DocsCloudSecurityConfig) IsSecretSet() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableString and assigns it to the Secret field.
func (o *DocsCloudSecurityConfig) SetSecret(v string) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *DocsCloudSecurityConfig) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *DocsCloudSecurityConfig) UnsetSecret() {
	o.Secret.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudSecurityConfig) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudSecurityConfig) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *DocsCloudSecurityConfig) IsHeaderSet() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *DocsCloudSecurityConfig) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *DocsCloudSecurityConfig) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *DocsCloudSecurityConfig) UnsetHeader() {
	o.Header.Unset()
}

func (o DocsCloudSecurityConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudSecurityConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	return toSerialize, nil
}

type NullableDocsCloudSecurityConfig struct {
	value *DocsCloudSecurityConfig
	isSet bool
}

func (v NullableDocsCloudSecurityConfig) Get() *DocsCloudSecurityConfig {
	return v.value
}

func (v *NullableDocsCloudSecurityConfig) Set(val *DocsCloudSecurityConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudSecurityConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudSecurityConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudSecurityConfig(val *DocsCloudSecurityConfig) *NullableDocsCloudSecurityConfig {
	return &NullableDocsCloudSecurityConfig{value: val, isSet: true}
}

func (v NullableDocsCloudSecurityConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudSecurityConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

