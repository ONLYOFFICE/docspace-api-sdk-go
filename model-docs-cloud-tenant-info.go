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

// checks if the DocsCloudTenantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudTenantInfo{}

// DocsCloudTenantInfo Represents the license and server information of a DocsCloud tenant, with usage statistics for the current period.
type DocsCloudTenantInfo struct {
	// The license information.
	License *DocsCloudLicenseInfo `json:"license,omitempty"`
	// The DocsCloud server information.
	Server *DocsCloudServerInfo `json:"server,omitempty"`
	// The user limits of the license.
	UsersLimit *DocsCloudUsersLimit `json:"usersLimit,omitempty"`
	// The usage statistics for the current period.
	Stats *DocsCloudStats `json:"stats,omitempty"`
}

// NewDocsCloudTenantInfo instantiates a new DocsCloudTenantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudTenantInfo() *DocsCloudTenantInfo {
	this := DocsCloudTenantInfo{}
	return &this
}

// NewDocsCloudTenantInfoWithDefaults instantiates a new DocsCloudTenantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudTenantInfoWithDefaults() *DocsCloudTenantInfo {
	this := DocsCloudTenantInfo{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *DocsCloudTenantInfo) GetLicense() DocsCloudLicenseInfo {
	if o == nil || IsNil(o.License) {
		var ret DocsCloudLicenseInfo
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfo) GetLicenseOk() (*DocsCloudLicenseInfo, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *DocsCloudTenantInfo) IsLicenseSet() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given DocsCloudLicenseInfo and assigns it to the License field.
func (o *DocsCloudTenantInfo) SetLicense(v DocsCloudLicenseInfo) {
	o.License = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DocsCloudTenantInfo) GetServer() DocsCloudServerInfo {
	if o == nil || IsNil(o.Server) {
		var ret DocsCloudServerInfo
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfo) GetServerOk() (*DocsCloudServerInfo, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DocsCloudTenantInfo) IsServerSet() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given DocsCloudServerInfo and assigns it to the Server field.
func (o *DocsCloudTenantInfo) SetServer(v DocsCloudServerInfo) {
	o.Server = &v
}

// GetUsersLimit returns the UsersLimit field value if set, zero value otherwise.
func (o *DocsCloudTenantInfo) GetUsersLimit() DocsCloudUsersLimit {
	if o == nil || IsNil(o.UsersLimit) {
		var ret DocsCloudUsersLimit
		return ret
	}
	return *o.UsersLimit
}

// GetUsersLimitOk returns a tuple with the UsersLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfo) GetUsersLimitOk() (*DocsCloudUsersLimit, bool) {
	if o == nil || IsNil(o.UsersLimit) {
		return nil, false
	}
	return o.UsersLimit, true
}

// HasUsersLimit returns a boolean if a field has been set.
func (o *DocsCloudTenantInfo) IsUsersLimitSet() bool {
	if o != nil && !IsNil(o.UsersLimit) {
		return true
	}

	return false
}

// SetUsersLimit gets a reference to the given DocsCloudUsersLimit and assigns it to the UsersLimit field.
func (o *DocsCloudTenantInfo) SetUsersLimit(v DocsCloudUsersLimit) {
	o.UsersLimit = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *DocsCloudTenantInfo) GetStats() DocsCloudStats {
	if o == nil || IsNil(o.Stats) {
		var ret DocsCloudStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudTenantInfo) GetStatsOk() (*DocsCloudStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *DocsCloudTenantInfo) IsStatsSet() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given DocsCloudStats and assigns it to the Stats field.
func (o *DocsCloudTenantInfo) SetStats(v DocsCloudStats) {
	o.Stats = &v
}

func (o DocsCloudTenantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudTenantInfo) ToMap() (map[string]interface{}, error) {
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

type NullableDocsCloudTenantInfo struct {
	value *DocsCloudTenantInfo
	isSet bool
}

func (v NullableDocsCloudTenantInfo) Get() *DocsCloudTenantInfo {
	return v.value
}

func (v *NullableDocsCloudTenantInfo) Set(val *DocsCloudTenantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudTenantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudTenantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudTenantInfo(val *DocsCloudTenantInfo) *NullableDocsCloudTenantInfo {
	return &NullableDocsCloudTenantInfo{value: val, isSet: true}
}

func (v NullableDocsCloudTenantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudTenantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

