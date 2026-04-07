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

// checks if the CspRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CspRequestsDto{}

// CspRequestsDto The request parameters for configuring the Content Security Policy (CSP) settings.
type CspRequestsDto struct {
	// The collection of allowed domains in the Content Security Policy (CSP).
	Domains []string `json:"domains,omitempty"`
}

// NewCspRequestsDto instantiates a new CspRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCspRequestsDto() *CspRequestsDto {
	this := CspRequestsDto{}
	return &this
}

// NewCspRequestsDtoWithDefaults instantiates a new CspRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCspRequestsDtoWithDefaults() *CspRequestsDto {
	this := CspRequestsDto{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CspRequestsDto) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CspRequestsDto) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *CspRequestsDto) IsDomainsSet() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *CspRequestsDto) SetDomains(v []string) {
	o.Domains = v
}

func (o CspRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CspRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	return toSerialize, nil
}

type NullableCspRequestsDto struct {
	value *CspRequestsDto
	isSet bool
}

func (v NullableCspRequestsDto) Get() *CspRequestsDto {
	return v.value
}

func (v *NullableCspRequestsDto) Set(val *CspRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCspRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCspRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCspRequestsDto(val *CspRequestsDto) *NullableCspRequestsDto {
	return &NullableCspRequestsDto{value: val, isSet: true}
}

func (v NullableCspRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCspRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

