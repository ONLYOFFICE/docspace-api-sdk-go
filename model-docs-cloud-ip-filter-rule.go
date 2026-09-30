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

// checks if the DocsCloudIpFilterRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudIpFilterRule{}

// DocsCloudIpFilterRule Represents the IP filter rule of a Docs Connect tenant.
type DocsCloudIpFilterRule struct {
	// The IP address.
	Address NullableString `json:"address,omitempty"`
	// Whether the IP address is allowed.
	Allowed *bool `json:"allowed,omitempty"`
}

// NewDocsCloudIpFilterRule instantiates a new DocsCloudIpFilterRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudIpFilterRule() *DocsCloudIpFilterRule {
	this := DocsCloudIpFilterRule{}
	return &this
}

// NewDocsCloudIpFilterRuleWithDefaults instantiates a new DocsCloudIpFilterRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudIpFilterRuleWithDefaults() *DocsCloudIpFilterRule {
	this := DocsCloudIpFilterRule{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudIpFilterRule) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudIpFilterRule) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *DocsCloudIpFilterRule) IsAddressSet() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *DocsCloudIpFilterRule) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *DocsCloudIpFilterRule) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *DocsCloudIpFilterRule) UnsetAddress() {
	o.Address.Unset()
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *DocsCloudIpFilterRule) GetAllowed() bool {
	if o == nil || IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudIpFilterRule) GetAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *DocsCloudIpFilterRule) IsAllowedSet() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *DocsCloudIpFilterRule) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o DocsCloudIpFilterRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudIpFilterRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableDocsCloudIpFilterRule struct {
	value *DocsCloudIpFilterRule
	isSet bool
}

func (v NullableDocsCloudIpFilterRule) Get() *DocsCloudIpFilterRule {
	return v.value
}

func (v *NullableDocsCloudIpFilterRule) Set(val *DocsCloudIpFilterRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudIpFilterRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudIpFilterRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudIpFilterRule(val *DocsCloudIpFilterRule) *NullableDocsCloudIpFilterRule {
	return &NullableDocsCloudIpFilterRule{value: val, isSet: true}
}

func (v NullableDocsCloudIpFilterRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudIpFilterRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

