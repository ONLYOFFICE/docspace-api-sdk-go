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

// checks if the DocsCloudServerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudServerConfig{}

// DocsCloudServerConfig Represents the server configuration of a Docs Connect tenant.
type DocsCloudServerConfig struct {
	// Whether anonymous access is supported.
	IsAnonymousSupport *bool `json:"isAnonymousSupport,omitempty"`
	// The maximum file size in bytes.
	FileSizeLimit *int64 `json:"fileSizeLimit,omitempty"`
}

// NewDocsCloudServerConfig instantiates a new DocsCloudServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudServerConfig() *DocsCloudServerConfig {
	this := DocsCloudServerConfig{}
	return &this
}

// NewDocsCloudServerConfigWithDefaults instantiates a new DocsCloudServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudServerConfigWithDefaults() *DocsCloudServerConfig {
	this := DocsCloudServerConfig{}
	return &this
}

// GetIsAnonymousSupport returns the IsAnonymousSupport field value if set, zero value otherwise.
func (o *DocsCloudServerConfig) GetIsAnonymousSupport() bool {
	if o == nil || IsNil(o.IsAnonymousSupport) {
		var ret bool
		return ret
	}
	return *o.IsAnonymousSupport
}

// GetIsAnonymousSupportOk returns a tuple with the IsAnonymousSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudServerConfig) GetIsAnonymousSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymousSupport) {
		return nil, false
	}
	return o.IsAnonymousSupport, true
}

// HasIsAnonymousSupport returns a boolean if a field has been set.
func (o *DocsCloudServerConfig) IsIsAnonymousSupportSet() bool {
	if o != nil && !IsNil(o.IsAnonymousSupport) {
		return true
	}

	return false
}

// SetIsAnonymousSupport gets a reference to the given bool and assigns it to the IsAnonymousSupport field.
func (o *DocsCloudServerConfig) SetIsAnonymousSupport(v bool) {
	o.IsAnonymousSupport = &v
}

// GetFileSizeLimit returns the FileSizeLimit field value if set, zero value otherwise.
func (o *DocsCloudServerConfig) GetFileSizeLimit() int64 {
	if o == nil || IsNil(o.FileSizeLimit) {
		var ret int64
		return ret
	}
	return *o.FileSizeLimit
}

// GetFileSizeLimitOk returns a tuple with the FileSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudServerConfig) GetFileSizeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSizeLimit) {
		return nil, false
	}
	return o.FileSizeLimit, true
}

// HasFileSizeLimit returns a boolean if a field has been set.
func (o *DocsCloudServerConfig) IsFileSizeLimitSet() bool {
	if o != nil && !IsNil(o.FileSizeLimit) {
		return true
	}

	return false
}

// SetFileSizeLimit gets a reference to the given int64 and assigns it to the FileSizeLimit field.
func (o *DocsCloudServerConfig) SetFileSizeLimit(v int64) {
	o.FileSizeLimit = &v
}

func (o DocsCloudServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudServerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAnonymousSupport) {
		toSerialize["isAnonymousSupport"] = o.IsAnonymousSupport
	}
	if !IsNil(o.FileSizeLimit) {
		toSerialize["fileSizeLimit"] = o.FileSizeLimit
	}
	return toSerialize, nil
}

type NullableDocsCloudServerConfig struct {
	value *DocsCloudServerConfig
	isSet bool
}

func (v NullableDocsCloudServerConfig) Get() *DocsCloudServerConfig {
	return v.value
}

func (v *NullableDocsCloudServerConfig) Set(val *DocsCloudServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudServerConfig(val *DocsCloudServerConfig) *NullableDocsCloudServerConfig {
	return &NullableDocsCloudServerConfig{value: val, isSet: true}
}

func (v NullableDocsCloudServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

