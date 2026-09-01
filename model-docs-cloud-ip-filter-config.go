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

// checks if the DocsCloudIpFilterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudIpFilterConfig{}

// DocsCloudIpFilterConfig Represents the IP filter configuration of a DocsCloud tenant.
type DocsCloudIpFilterConfig struct {
	// The IP filter rules.
	Rules []DocsCloudIpFilterRule `json:"rules,omitempty"`
}

// NewDocsCloudIpFilterConfig instantiates a new DocsCloudIpFilterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudIpFilterConfig() *DocsCloudIpFilterConfig {
	this := DocsCloudIpFilterConfig{}
	return &this
}

// NewDocsCloudIpFilterConfigWithDefaults instantiates a new DocsCloudIpFilterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudIpFilterConfigWithDefaults() *DocsCloudIpFilterConfig {
	this := DocsCloudIpFilterConfig{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudIpFilterConfig) GetRules() []DocsCloudIpFilterRule {
	if o == nil {
		var ret []DocsCloudIpFilterRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudIpFilterConfig) GetRulesOk() ([]DocsCloudIpFilterRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *DocsCloudIpFilterConfig) IsRulesSet() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []DocsCloudIpFilterRule and assigns it to the Rules field.
func (o *DocsCloudIpFilterConfig) SetRules(v []DocsCloudIpFilterRule) {
	o.Rules = v
}

func (o DocsCloudIpFilterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudIpFilterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableDocsCloudIpFilterConfig struct {
	value *DocsCloudIpFilterConfig
	isSet bool
}

func (v NullableDocsCloudIpFilterConfig) Get() *DocsCloudIpFilterConfig {
	return v.value
}

func (v *NullableDocsCloudIpFilterConfig) Set(val *DocsCloudIpFilterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudIpFilterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudIpFilterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudIpFilterConfig(val *DocsCloudIpFilterConfig) *NullableDocsCloudIpFilterConfig {
	return &NullableDocsCloudIpFilterConfig{value: val, isSet: true}
}

func (v NullableDocsCloudIpFilterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudIpFilterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

