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

// checks if the DnsSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsSettingsRequestDto{}

// DnsSettingsRequestDto The custom domain the portal answers on, and whether that mapping is in force.
type DnsSettingsRequestDto struct {
	// The domain the portal is to be reachable under, as a bare hostname without a scheme. It must not collide with  the reserved base domain of the installation, and a name that fails validation is refused without disturbing  the mapping in force. It is read only while `enable` is true.
	DnsName NullableString `json:"dnsName,omitempty"`
	// Whether the custom domain is put in force. Setting it false clears the mapping and ignores `dnsName`; setting  it true also stops the previous domain from answering and rewrites any Content Security Policy entry that  named it.
	Enable *bool `json:"enable,omitempty"`
}

// NewDnsSettingsRequestDto instantiates a new DnsSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSettingsRequestDto() *DnsSettingsRequestDto {
	this := DnsSettingsRequestDto{}
	return &this
}

// NewDnsSettingsRequestDtoWithDefaults instantiates a new DnsSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSettingsRequestDtoWithDefaults() *DnsSettingsRequestDto {
	this := DnsSettingsRequestDto{}
	return &this
}

// GetDnsName returns the DnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsSettingsRequestDto) GetDnsName() string {
	if o == nil || IsNil(o.DnsName.Get()) {
		var ret string
		return ret
	}
	return *o.DnsName.Get()
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsSettingsRequestDto) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsName.Get(), o.DnsName.IsSet()
}

// HasDnsName returns a boolean if a field has been set.
func (o *DnsSettingsRequestDto) IsDnsNameSet() bool {
	if o != nil && o.DnsName.IsSet() {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given NullableString and assigns it to the DnsName field.
func (o *DnsSettingsRequestDto) SetDnsName(v string) {
	o.DnsName.Set(&v)
}
// SetDnsNameNil sets the value for DnsName to be an explicit nil
func (o *DnsSettingsRequestDto) SetDnsNameNil() {
	o.DnsName.Set(nil)
}

// UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
func (o *DnsSettingsRequestDto) UnsetDnsName() {
	o.DnsName.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *DnsSettingsRequestDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSettingsRequestDto) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *DnsSettingsRequestDto) IsEnableSet() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *DnsSettingsRequestDto) SetEnable(v bool) {
	o.Enable = &v
}

func (o DnsSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsName.IsSet() {
		toSerialize["dnsName"] = o.DnsName.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableDnsSettingsRequestDto struct {
	value *DnsSettingsRequestDto
	isSet bool
}

func (v NullableDnsSettingsRequestDto) Get() *DnsSettingsRequestDto {
	return v.value
}

func (v *NullableDnsSettingsRequestDto) Set(val *DnsSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSettingsRequestDto(val *DnsSettingsRequestDto) *NullableDnsSettingsRequestDto {
	return &NullableDnsSettingsRequestDto{value: val, isSet: true}
}

func (v NullableDnsSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

