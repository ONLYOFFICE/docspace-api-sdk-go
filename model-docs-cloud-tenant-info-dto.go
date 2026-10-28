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

// checks if the DocsCloudTenantInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudTenantInfoDto{}

// DocsCloudTenantInfoDto Represents the license and server information of a Docs Connect tenant, with usage statistics for the current period.
type DocsCloudTenantInfoDto struct {
	// The license information.
	License *DocsCloudLicenseInfoDto `json:"license,omitempty"`
	// The Docs Connect server information.
	Server *DocsCloudServerInfoDto `json:"server,omitempty"`
	// The user limits of the license.
	UsersLimit *DocsCloudUsersLimitDto `json:"usersLimit,omitempty"`
	// The usage statistics for the current period.
	Stats *DocsCloudStatsDto `json:"stats,omitempty"`
}

// NewDocsCloudTenantInfoDto instantiates a new DocsCloudTenantInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudTenantInfoDto() *DocsCloudTenantInfoDto {
	this := DocsCloudTenantInfoDto{}
	return &this
}

// NewDocsCloudTenantInfoDtoWithDefaults instantiates a new DocsCloudTenantInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudTenantInfoDtoWithDefaults() *DocsCloudTenantInfoDto {
	this := DocsCloudTenantInfoDto{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *DocsCloudTenantInfoDto) GetLicense() DocsCloudLicenseInfoDto {
	if o == nil || IsNil(o.License) {
		var ret DocsCloudLicenseInfoDto
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfoDto) GetLicenseOk() (*DocsCloudLicenseInfoDto, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *DocsCloudTenantInfoDto) IsLicenseSet() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given DocsCloudLicenseInfoDto and assigns it to the License field.
func (o *DocsCloudTenantInfoDto) SetLicense(v DocsCloudLicenseInfoDto) {
	o.License = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DocsCloudTenantInfoDto) GetServer() DocsCloudServerInfoDto {
	if o == nil || IsNil(o.Server) {
		var ret DocsCloudServerInfoDto
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfoDto) GetServerOk() (*DocsCloudServerInfoDto, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DocsCloudTenantInfoDto) IsServerSet() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given DocsCloudServerInfoDto and assigns it to the Server field.
func (o *DocsCloudTenantInfoDto) SetServer(v DocsCloudServerInfoDto) {
	o.Server = &v
}

// GetUsersLimit returns the UsersLimit field value if set, zero value otherwise.
func (o *DocsCloudTenantInfoDto) GetUsersLimit() DocsCloudUsersLimitDto {
	if o == nil || IsNil(o.UsersLimit) {
		var ret DocsCloudUsersLimitDto
		return ret
	}
	return *o.UsersLimit
}

// GetUsersLimitOk returns a tuple with the UsersLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfoDto) GetUsersLimitOk() (*DocsCloudUsersLimitDto, bool) {
	if o == nil || IsNil(o.UsersLimit) {
		return nil, false
	}
	return o.UsersLimit, true
}

// HasUsersLimit returns a boolean if a field has been set.
func (o *DocsCloudTenantInfoDto) IsUsersLimitSet() bool {
	if o != nil && !IsNil(o.UsersLimit) {
		return true
	}

	return false
}

// SetUsersLimit gets a reference to the given DocsCloudUsersLimitDto and assigns it to the UsersLimit field.
func (o *DocsCloudTenantInfoDto) SetUsersLimit(v DocsCloudUsersLimitDto) {
	o.UsersLimit = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *DocsCloudTenantInfoDto) GetStats() DocsCloudStatsDto {
	if o == nil || IsNil(o.Stats) {
		var ret DocsCloudStatsDto
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfoDto) GetStatsOk() (*DocsCloudStatsDto, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *DocsCloudTenantInfoDto) IsStatsSet() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given DocsCloudStatsDto and assigns it to the Stats field.
func (o *DocsCloudTenantInfoDto) SetStats(v DocsCloudStatsDto) {
	o.Stats = &v
}

func (o DocsCloudTenantInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudTenantInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.UsersLimit) {
		toSerialize["usersLimit"] = o.UsersLimit
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	return toSerialize, nil
}

type NullableDocsCloudTenantInfoDto struct {
	value *DocsCloudTenantInfoDto
	isSet bool
}

func (v NullableDocsCloudTenantInfoDto) Get() *DocsCloudTenantInfoDto {
	return v.value
}

func (v *NullableDocsCloudTenantInfoDto) Set(val *DocsCloudTenantInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudTenantInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudTenantInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudTenantInfoDto(val *DocsCloudTenantInfoDto) *NullableDocsCloudTenantInfoDto {
	return &NullableDocsCloudTenantInfoDto{value: val, isSet: true}
}

func (v NullableDocsCloudTenantInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudTenantInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

