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

// checks if the DocsCloudIpFilterConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudIpFilterConfigDto{}

// DocsCloudIpFilterConfigDto Represents the IP filter configuration of a Docs Connect tenant.
type DocsCloudIpFilterConfigDto struct {
	// The IP filter rules.
	Rules []DocsCloudIpFilterRuleDto `json:"rules,omitempty"`
}

// NewDocsCloudIpFilterConfigDto instantiates a new DocsCloudIpFilterConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudIpFilterConfigDto() *DocsCloudIpFilterConfigDto {
	this := DocsCloudIpFilterConfigDto{}
	return &this
}

// NewDocsCloudIpFilterConfigDtoWithDefaults instantiates a new DocsCloudIpFilterConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudIpFilterConfigDtoWithDefaults() *DocsCloudIpFilterConfigDto {
	this := DocsCloudIpFilterConfigDto{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudIpFilterConfigDto) GetRules() []DocsCloudIpFilterRuleDto {
	if o == nil {
		var ret []DocsCloudIpFilterRuleDto
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudIpFilterConfigDto) GetRulesOk() ([]DocsCloudIpFilterRuleDto, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *DocsCloudIpFilterConfigDto) IsRulesSet() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []DocsCloudIpFilterRuleDto and assigns it to the Rules field.
func (o *DocsCloudIpFilterConfigDto) SetRules(v []DocsCloudIpFilterRuleDto) {
	o.Rules = v
}

func (o DocsCloudIpFilterConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudIpFilterConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableDocsCloudIpFilterConfigDto struct {
	value *DocsCloudIpFilterConfigDto
	isSet bool
}

func (v NullableDocsCloudIpFilterConfigDto) Get() *DocsCloudIpFilterConfigDto {
	return v.value
}

func (v *NullableDocsCloudIpFilterConfigDto) Set(val *DocsCloudIpFilterConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudIpFilterConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudIpFilterConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudIpFilterConfigDto(val *DocsCloudIpFilterConfigDto) *NullableDocsCloudIpFilterConfigDto {
	return &NullableDocsCloudIpFilterConfigDto{value: val, isSet: true}
}

func (v NullableDocsCloudIpFilterConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudIpFilterConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

