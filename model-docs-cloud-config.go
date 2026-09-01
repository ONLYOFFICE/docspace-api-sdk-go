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

// checks if the DocsCloudConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudConfig{}

// DocsCloudConfig Represents the configuration of a DocsCloud tenant.
type DocsCloudConfig struct {
	// The tenant name.
	TenantName NullableString `json:"tenantName,omitempty"`
	// The security configuration.
	Security *DocsCloudSecurityConfig `json:"security,omitempty"`
	// The server configuration.
	Server *DocsCloudServerConfig `json:"server,omitempty"`
	// The WOPI configuration.
	Wopi *DocsCloudWopiConfig `json:"wopi,omitempty"`
	// The IP filter configuration.
	IpFilter *DocsCloudIpFilterConfig `json:"ipFilter,omitempty"`
}

// NewDocsCloudConfig instantiates a new DocsCloudConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudConfig() *DocsCloudConfig {
	this := DocsCloudConfig{}
	return &this
}

// NewDocsCloudConfigWithDefaults instantiates a new DocsCloudConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudConfigWithDefaults() *DocsCloudConfig {
	this := DocsCloudConfig{}
	return &this
}

// GetTenantName returns the TenantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudConfig) GetTenantName() string {
	if o == nil || IsNil(o.TenantName.Get()) {
		var ret string
		return ret
	}
	return *o.TenantName.Get()
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudConfig) GetTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantName.Get(), o.TenantName.IsSet()
}

// HasTenantName returns a boolean if a field has been set.
func (o *DocsCloudConfig) IsTenantNameSet() bool {
	if o != nil && o.TenantName.IsSet() {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given NullableString and assigns it to the TenantName field.
func (o *DocsCloudConfig) SetTenantName(v string) {
	o.TenantName.Set(&v)
}
// SetTenantNameNil sets the value for TenantName to be an explicit nil
func (o *DocsCloudConfig) SetTenantNameNil() {
	o.TenantName.Set(nil)
}

// UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
func (o *DocsCloudConfig) UnsetTenantName() {
	o.TenantName.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *DocsCloudConfig) GetSecurity() DocsCloudSecurityConfig {
	if o == nil || IsNil(o.Security) {
		var ret DocsCloudSecurityConfig
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfig) GetSecurityOk() (*DocsCloudSecurityConfig, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *DocsCloudConfig) IsSecuritySet() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given DocsCloudSecurityConfig and assigns it to the Security field.
func (o *DocsCloudConfig) SetSecurity(v DocsCloudSecurityConfig) {
	o.Security = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DocsCloudConfig) GetServer() DocsCloudServerConfig {
	if o == nil || IsNil(o.Server) {
		var ret DocsCloudServerConfig
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfig) GetServerOk() (*DocsCloudServerConfig, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DocsCloudConfig) IsServerSet() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given DocsCloudServerConfig and assigns it to the Server field.
func (o *DocsCloudConfig) SetServer(v DocsCloudServerConfig) {
	o.Server = &v
}

// GetWopi returns the Wopi field value if set, zero value otherwise.
func (o *DocsCloudConfig) GetWopi() DocsCloudWopiConfig {
	if o == nil || IsNil(o.Wopi) {
		var ret DocsCloudWopiConfig
		return ret
	}
	return *o.Wopi
}

// GetWopiOk returns a tuple with the Wopi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfig) GetWopiOk() (*DocsCloudWopiConfig, bool) {
	if o == nil || IsNil(o.Wopi) {
		return nil, false
	}
	return o.Wopi, true
}

// HasWopi returns a boolean if a field has been set.
func (o *DocsCloudConfig) IsWopiSet() bool {
	if o != nil && !IsNil(o.Wopi) {
		return true
	}

	return false
}

// SetWopi gets a reference to the given DocsCloudWopiConfig and assigns it to the Wopi field.
func (o *DocsCloudConfig) SetWopi(v DocsCloudWopiConfig) {
	o.Wopi = &v
}

// GetIpFilter returns the IpFilter field value if set, zero value otherwise.
func (o *DocsCloudConfig) GetIpFilter() DocsCloudIpFilterConfig {
	if o == nil || IsNil(o.IpFilter) {
		var ret DocsCloudIpFilterConfig
		return ret
	}
	return *o.IpFilter
}

// GetIpFilterOk returns a tuple with the IpFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfig) GetIpFilterOk() (*DocsCloudIpFilterConfig, bool) {
	if o == nil || IsNil(o.IpFilter) {
		return nil, false
	}
	return o.IpFilter, true
}

// HasIpFilter returns a boolean if a field has been set.
func (o *DocsCloudConfig) IsIpFilterSet() bool {
	if o != nil && !IsNil(o.IpFilter) {
		return true
	}

	return false
}

// SetIpFilter gets a reference to the given DocsCloudIpFilterConfig and assigns it to the IpFilter field.
func (o *DocsCloudConfig) SetIpFilter(v DocsCloudIpFilterConfig) {
	o.IpFilter = &v
}

func (o DocsCloudConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantName.IsSet() {
		toSerialize["tenantName"] = o.TenantName.Get()
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Wopi) {
		toSerialize["wopi"] = o.Wopi
	}
	if !IsNil(o.IpFilter) {
		toSerialize["ipFilter"] = o.IpFilter
	}
	return toSerialize, nil
}

type NullableDocsCloudConfig struct {
	value *DocsCloudConfig
	isSet bool
}

func (v NullableDocsCloudConfig) Get() *DocsCloudConfig {
	return v.value
}

func (v *NullableDocsCloudConfig) Set(val *DocsCloudConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudConfig(val *DocsCloudConfig) *NullableDocsCloudConfig {
	return &NullableDocsCloudConfig{value: val, isSet: true}
}

func (v NullableDocsCloudConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

