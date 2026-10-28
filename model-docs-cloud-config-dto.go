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

// checks if the DocsCloudConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudConfigDto{}

// DocsCloudConfigDto Represents the configuration of a Docs Connect tenant.
type DocsCloudConfigDto struct {
	// The tenant name.
	TenantName NullableString `json:"tenantName,omitempty"`
	// The security configuration.
	Security *DocsCloudSecurityConfigDto `json:"security,omitempty"`
	// The server configuration.
	Server *DocsCloudServerConfigDto `json:"server,omitempty"`
	// The WOPI configuration.
	Wopi *DocsCloudWopiConfigDto `json:"wopi,omitempty"`
	// The IP filter configuration.
	IpFilter *DocsCloudIpFilterConfigDto `json:"ipFilter,omitempty"`
}

// NewDocsCloudConfigDto instantiates a new DocsCloudConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudConfigDto() *DocsCloudConfigDto {
	this := DocsCloudConfigDto{}
	return &this
}

// NewDocsCloudConfigDtoWithDefaults instantiates a new DocsCloudConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudConfigDtoWithDefaults() *DocsCloudConfigDto {
	this := DocsCloudConfigDto{}
	return &this
}

// GetTenantName returns the TenantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudConfigDto) GetTenantName() string {
	if o == nil || IsNil(o.TenantName.Get()) {
		var ret string
		return ret
	}
	return *o.TenantName.Get()
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudConfigDto) GetTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantName.Get(), o.TenantName.IsSet()
}

// HasTenantName returns a boolean if a field has been set.
func (o *DocsCloudConfigDto) IsTenantNameSet() bool {
	if o != nil && o.TenantName.IsSet() {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given NullableString and assigns it to the TenantName field.
func (o *DocsCloudConfigDto) SetTenantName(v string) {
	o.TenantName.Set(&v)
}
// SetTenantNameNil sets the value for TenantName to be an explicit nil
func (o *DocsCloudConfigDto) SetTenantNameNil() {
	o.TenantName.Set(nil)
}

// UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
func (o *DocsCloudConfigDto) UnsetTenantName() {
	o.TenantName.Unset()
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *DocsCloudConfigDto) GetSecurity() DocsCloudSecurityConfigDto {
	if o == nil || IsNil(o.Security) {
		var ret DocsCloudSecurityConfigDto
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfigDto) GetSecurityOk() (*DocsCloudSecurityConfigDto, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *DocsCloudConfigDto) IsSecuritySet() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given DocsCloudSecurityConfigDto and assigns it to the Security field.
func (o *DocsCloudConfigDto) SetSecurity(v DocsCloudSecurityConfigDto) {
	o.Security = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DocsCloudConfigDto) GetServer() DocsCloudServerConfigDto {
	if o == nil || IsNil(o.Server) {
		var ret DocsCloudServerConfigDto
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfigDto) GetServerOk() (*DocsCloudServerConfigDto, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DocsCloudConfigDto) IsServerSet() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given DocsCloudServerConfigDto and assigns it to the Server field.
func (o *DocsCloudConfigDto) SetServer(v DocsCloudServerConfigDto) {
	o.Server = &v
}

// GetWopi returns the Wopi field value if set, zero value otherwise.
func (o *DocsCloudConfigDto) GetWopi() DocsCloudWopiConfigDto {
	if o == nil || IsNil(o.Wopi) {
		var ret DocsCloudWopiConfigDto
		return ret
	}
	return *o.Wopi
}

// GetWopiOk returns a tuple with the Wopi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfigDto) GetWopiOk() (*DocsCloudWopiConfigDto, bool) {
	if o == nil || IsNil(o.Wopi) {
		return nil, false
	}
	return o.Wopi, true
}

// HasWopi returns a boolean if a field has been set.
func (o *DocsCloudConfigDto) IsWopiSet() bool {
	if o != nil && !IsNil(o.Wopi) {
		return true
	}

	return false
}

// SetWopi gets a reference to the given DocsCloudWopiConfigDto and assigns it to the Wopi field.
func (o *DocsCloudConfigDto) SetWopi(v DocsCloudWopiConfigDto) {
	o.Wopi = &v
}

// GetIpFilter returns the IpFilter field value if set, zero value otherwise.
func (o *DocsCloudConfigDto) GetIpFilter() DocsCloudIpFilterConfigDto {
	if o == nil || IsNil(o.IpFilter) {
		var ret DocsCloudIpFilterConfigDto
		return ret
	}
	return *o.IpFilter
}

// GetIpFilterOk returns a tuple with the IpFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudConfigDto) GetIpFilterOk() (*DocsCloudIpFilterConfigDto, bool) {
	if o == nil || IsNil(o.IpFilter) {
		return nil, false
	}
	return o.IpFilter, true
}

// HasIpFilter returns a boolean if a field has been set.
func (o *DocsCloudConfigDto) IsIpFilterSet() bool {
	if o != nil && !IsNil(o.IpFilter) {
		return true
	}

	return false
}

// SetIpFilter gets a reference to the given DocsCloudIpFilterConfigDto and assigns it to the IpFilter field.
func (o *DocsCloudConfigDto) SetIpFilter(v DocsCloudIpFilterConfigDto) {
	o.IpFilter = &v
}

func (o DocsCloudConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudConfigDto) ToMap() (map[string]interface{}, error) {
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

type NullableDocsCloudConfigDto struct {
	value *DocsCloudConfigDto
	isSet bool
}

func (v NullableDocsCloudConfigDto) Get() *DocsCloudConfigDto {
	return v.value
}

func (v *NullableDocsCloudConfigDto) Set(val *DocsCloudConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudConfigDto(val *DocsCloudConfigDto) *NullableDocsCloudConfigDto {
	return &NullableDocsCloudConfigDto{value: val, isSet: true}
}

func (v NullableDocsCloudConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

