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

// checks if the DnsSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsSettingsRequestsDto{}

// DnsSettingsRequestsDto The request parameters for managing the DNS (Domain Name System) settings.
type DnsSettingsRequestsDto struct {
	// The DNS (Domain Name System) configuration name.
	DnsName NullableString `json:"dnsName,omitempty"`
	// Specifies whether the DNS settings are enabled.
	Enable *bool `json:"enable,omitempty"`
}

// NewDnsSettingsRequestsDto instantiates a new DnsSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSettingsRequestsDto() *DnsSettingsRequestsDto {
	this := DnsSettingsRequestsDto{}
	return &this
}

// NewDnsSettingsRequestsDtoWithDefaults instantiates a new DnsSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSettingsRequestsDtoWithDefaults() *DnsSettingsRequestsDto {
	this := DnsSettingsRequestsDto{}
	return &this
}

// GetDnsName returns the DnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsSettingsRequestsDto) GetDnsName() string {
	if o == nil || IsNil(o.DnsName.Get()) {
		var ret string
		return ret
	}
	return *o.DnsName.Get()
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsSettingsRequestsDto) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsName.Get(), o.DnsName.IsSet()
}

// HasDnsName returns a boolean if a field has been set.
func (o *DnsSettingsRequestsDto) IsDnsNameSet() bool {
	if o != nil && o.DnsName.IsSet() {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given NullableString and assigns it to the DnsName field.
func (o *DnsSettingsRequestsDto) SetDnsName(v string) {
	o.DnsName.Set(&v)
}
// SetDnsNameNil sets the value for DnsName to be an explicit nil
func (o *DnsSettingsRequestsDto) SetDnsNameNil() {
	o.DnsName.Set(nil)
}

// UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
func (o *DnsSettingsRequestsDto) UnsetDnsName() {
	o.DnsName.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *DnsSettingsRequestsDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSettingsRequestsDto) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *DnsSettingsRequestsDto) IsEnableSet() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *DnsSettingsRequestsDto) SetEnable(v bool) {
	o.Enable = &v
}

func (o DnsSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsName.IsSet() {
		toSerialize["dnsName"] = o.DnsName.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableDnsSettingsRequestsDto struct {
	value *DnsSettingsRequestsDto
	isSet bool
}

func (v NullableDnsSettingsRequestsDto) Get() *DnsSettingsRequestsDto {
	return v.value
}

func (v *NullableDnsSettingsRequestsDto) Set(val *DnsSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSettingsRequestsDto(val *DnsSettingsRequestsDto) *NullableDnsSettingsRequestsDto {
	return &NullableDnsSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableDnsSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

