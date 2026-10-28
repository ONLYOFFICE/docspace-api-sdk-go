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

// checks if the CspRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CspRequestDto{}

// CspRequestDto The external sources the portal Content Security Policy is to trust.
type CspRequestDto struct {
	// The domains the policy trusts, as the complete list that is to hold afterwards rather than a list of  additions: send the domains already trusted together with the new one to add one, leave one out to withdraw  it, and send an empty list to fall back to the portal built-in policy. An entry may be a bare host, a host  with a scheme, or a wildcard host such as `*.example.com`; it has to form a valid absolute address and may  contain ASCII characters only. Every entry becomes an allowed source for scripts, styles, images, fonts,  frames, media and connections at once - the directives cannot be set apart here.
	Domains []string `json:"domains,omitempty"`
}

// NewCspRequestDto instantiates a new CspRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCspRequestDto() *CspRequestDto {
	this := CspRequestDto{}
	return &this
}

// NewCspRequestDtoWithDefaults instantiates a new CspRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCspRequestDtoWithDefaults() *CspRequestDto {
	this := CspRequestDto{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CspRequestDto) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CspRequestDto) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *CspRequestDto) IsDomainsSet() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *CspRequestDto) SetDomains(v []string) {
	o.Domains = v
}

func (o CspRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CspRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	return toSerialize, nil
}

type NullableCspRequestDto struct {
	value *CspRequestDto
	isSet bool
}

func (v NullableCspRequestDto) Get() *CspRequestDto {
	return v.value
}

func (v *NullableCspRequestDto) Set(val *CspRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCspRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCspRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCspRequestDto(val *CspRequestDto) *NullableCspRequestDto {
	return &NullableCspRequestDto{value: val, isSet: true}
}

func (v NullableCspRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCspRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

